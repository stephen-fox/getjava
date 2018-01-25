package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/stephen-fox/getjava"
)

const (
	applicationName = "getjava"

	javaPackageUrlArg           = "u"
	downloadDestPathArg         = "o"
	archiveDecompressDirPathArg = "d"
	printHelpArg                = "h"
	versionArg                  = "v"

	examples = `	Download a package into the current directory:
	` + applicationName + ` -` + javaPackageUrlArg + ` http://download.oracle.com/jre.tar.gz

	Download a package into a different directory with a custom name:
	` + applicationName + ` -` + javaPackageUrlArg + ` http://download.oracle.com/jre.tar.gz` + ` -` + downloadDestPathArg + ` ~/Desktop/jre.tar.gz

	Download a package into a different directory with a custom name:
	` + applicationName + ` -` + javaPackageUrlArg + ` http://download.oracle.com/jre.tar.gz` + ` -` + downloadDestPathArg + ` ~/Desktop/jre.tar.gz

	Download a package and decompress it:
	` + applicationName + ` -` + javaPackageUrlArg + ` http://download.oracle.com/jre.tar.gz` + ` -` + archiveDecompressDirPathArg + ` ~/Desktop`
)

var (
	version string

	javaPackageUrl           = flag.String(javaPackageUrlArg, "", "The URL of the Java package")
	downloadDestPath         = flag.String(downloadDestPathArg, "", "The path to download the archive to")
	archiveDecompressDirPath = flag.String(archiveDecompressDirPathArg, "", "The path to decompress the archive to")

	shouldPrintHelp    = flag.Bool(printHelpArg, false, "Display this help page")
	shouldPrintVersion = flag.Bool(versionArg, false, "Print the version")
)

func main() {
	flag.Parse()

	if *shouldPrintHelp || len(os.Args) == 1 {
		fmt.Println(applicationName, version)
		fmt.Println()
		fmt.Println("[ABOUT]")
		fmt.Println("Utility for downloading Oracle Java packages.")
		fmt.Println()
		fmt.Println("[USAGE]")
		flag.PrintDefaults()
		fmt.Println()
		fmt.Println("[EXAMPLES]")
		fmt.Println(examples)
		os.Exit(0)
	}

	if *shouldPrintVersion {
		fmt.Println(version)
		os.Exit(0)
	}

	if len(strings.TrimSpace(*javaPackageUrl)) == 0 {
		log.Fatal("Please specify a Java package URL")
	}

	if len(strings.TrimSpace(*downloadDestPath)) == 0 {
		*downloadDestPath = "./" + path.Base(*javaPackageUrl)
	}

	log.Println("Downloading", *javaPackageUrl, "to", *downloadDestPath + "...")

	err := getjava.Download(*javaPackageUrl, *downloadDestPath)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Download complete")

	if len(strings.TrimSpace(*archiveDecompressDirPath)) > 0 {
		log.Println("Decompressing", *downloadDestPath, "to", *archiveDecompressDirPath + "...")

		err := getjava.Decompress(*downloadDestPath, *archiveDecompressDirPath)
		if err != nil {
			log.Fatal(err.Error())
		}

		log.Println("Decompression complete")
	}
}