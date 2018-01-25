# getjava

## What is it?
A cross platform utility for downloading and decompressing Java packages.

## How does it work?
If you would like to only download a JRE package:

```bash
$ getjava -u http://download.oracle.com/otn-pub/java/jdk/8u161-b12/2f38c3b165be4555a1fa6e98c45e0808/jre-8u161-linux-x64.tar.gz
```

Or, if you would also like to decompress it:
```bash
$ getjava -u http://download.oracle.com/otn-pub/java/jdk/8u161-b12/2f38c3b165be4555a1fa6e98c45e0808/jre-8u161-linux-x64.tar.gz -d 
```
