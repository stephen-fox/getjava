package getjava

import (
	"errors"
	"io"
	"os"
	"net/http"
	"strconv"
	"time"

	"github.com/mholt/archiver"
)

const (
	OracleCookie = "oraclelicense=accept-securebackup-cookie"
)

func Download(url string, outputPath string) error {
	headers := make(map[string]string)
	headers["Cookie"] = OracleCookie

	err := downloadFile(url, headers, outputPath)
	if err != nil {
		return err
	}

	return nil
}

func downloadFile(url string, optionalRequestHeaders map[string]string, filePath string) error {
	response, err := getHttpResponse(optionalRequestHeaders, url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	_, statErr := os.Stat(filePath)
	if statErr == nil {
		os.Remove(filePath)
	}

	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, response.Body)
	if err != nil {
		return err
	}

	return nil
}

func getHttpResponse(requestHeaders map[string]string, url string) (*http.Response, error) {
	i := 0
	response := &http.Response{}
	client := &http.Client{
		Timeout: time.Minute * 5,
		// Do not permit redirects because http.Client does not forward
		// certain HTTP headers when a redirect from another
		// site occurs.
		CheckRedirect: func(request *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	nextUrl := url

	for i < 100 {
		request, err := http.NewRequest("GET", nextUrl, nil)
		if err != nil {
			return response, err
		}

		if requestHeaders != nil {
			for k, v := range requestHeaders {
				request.Header.Add(k, v)
			}
		}

		response, err = client.Do(request)
		if err != nil {
			return response, err
		}

		if response.StatusCode == 200 {
			break
		} else {
			response.Body.Close()
			nextUrl = response.Header.Get("Location")
		}

		i++
	}

	if response.StatusCode != 200 {
		return response, errors.New("Failed to get HTTP code 200. " +
			"Final HTTP status code: " + strconv.Itoa(response.StatusCode))
	}

	return response, nil
}

func Decompress(archivePath string, destinationPath string) error {
	return archiver.TarGz.Open(archivePath, destinationPath)
}