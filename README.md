# getjava

## What is it?
A cross platform utility for downloading and decompressing Java packages.

## How does it work?
If you would like to download a JRE package:
```bash
$ getjava -u http://download.oracle.com/otn-pub/java/jdk/8u161-b12/2f38c3b165be4555a1fa6e98c45e0808/jre-8u161-linux-x64.tar.gz
```

Or, if you would like to download and decompress it to your desktop:
```bash
$ getjava \
  -u http://download.oracle.com/otn-pub/java/jdk/8u161-b12/2f38c3b165be4555a1fa6e98c45e0808/jre-8u161-linux-x64.tar.gz \
  -d ~/Desktop
```

You can also view application help and usage examples by running:
```bash
$ getjava -h
```

## Where do I download it from?
Check out the tags section.

## How do I build it?
You can use the included Gradle project:
```bash
$ ./gradlew buildApplication
```

You can also use `go build` if you like.