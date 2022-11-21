Usual go builds without any flags (except output) yields a **dynamically linked** 9.8MB binary.

```sh
$ go build -o pebble ./cmd/pebble/
```
```sh
```sh
$ file pebble | tr , '\n'
pebble: ELF 64-bit LSB executable
 x86-64
 version 1 (SYSV)
 dynamically linked
 interpreter /lib64/ld-linux-x86-64.so.2
 Go BuildID=lO0KJfsEFcS676KFl7vi/NGnln2gj6K5R8GdTrzbv/m-vHuPyfKWp9cSdjsXX5/6YDtyivHXzDNoTqv_F9M
 not stripped
```
```sh
$ stat pebble
  File: pebble
  Size: 10233024  	Blocks: 19992      IO Block: 4096   regular file
Device: 801h/2049d	Inode: 22917       Links: 1
Access: (0775/-rwxrwxr-x)  Uid: ( 1000/  ubuntu)   Gid: ( 1000/  ubuntu)
```

We can strip the binary of symbol table, debugging info and the DWARF symbol table, by passing `-ldflags="-s -w"` in the build command. It results in a 6.8MB sized **dynamically linked** pebble binary. [reference](https://pkg.go.dev/cmd/link)
```sh
$ go build -ldflags="-s -w" -o pebble ./cmd/pebble/
```
```sh
$ file pebble | tr , '\n'
pebble: ELF 64-bit LSB executable
 x86-64
 version 1 (SYSV)
 dynamically linked
 interpreter /lib64/ld-linux-x86-64.so.2
 Go BuildID=0YORf-Ttz9APxdhueOtt/f6BFL7n8DfXFQJGdy4bd/m-vHuPyfKWp9cSdjsXX5/U6s6NKbF4lew0dc2ilha
 stripped
```
```sh
$ stat pebble
  File: pebble
  Size: 7094272   	Blocks: 13856      IO Block: 4096   regular file
Device: 801h/2049d	Inode: 22917       Links: 1
Access: (0775/-rwxrwxr-x)  Uid: ( 1000/  ubuntu)   Gid: ( 1000/  ubuntu)
```

We can further compress the binary with [upx](https://upx.github.io/). However, the binary can be a bit slower and might consume more ram. In our case, it results to a 2.9MB sized **statically linked** pebble binary.
```sh
$ upx pebble 
                       Ultimate Packer for eXecutables
                          Copyright (C) 1996 - 2020
UPX 3.96        Markus Oberhumer, Laszlo Molnar & John Reiser   Jan 23rd 2020

        File size         Ratio      Format      Name
   --------------------   ------   -----------   -----------
   7094272 ->   2933084   41.34%   linux/amd64   pebble                        

Packed 1 file.
```
```sh
$ file pebble | tr , '\n'
pebble: ELF 64-bit LSB executable
 x86-64
 version 1 (SYSV)
 statically linked
 no section header
```
```sh
$ stat pebble
  File: pebble
  Size: 2933084   	Blocks: 5736       IO Block: 4096   regular file
Device: 801h/2049d	Inode: 279428      Links: 1
Access: (0775/-rwxrwxr-x)  Uid: ( 1000/  ubuntu)   Gid: ( 1000/  ubuntu)
```

Helpful resources:
- https://words.filippo.io/shrink-your-go-binaries-with-this-one-weird-trick/
- https://github.com/xaionaro/documentation/blob/master/golang/reduce-binary-size.md
