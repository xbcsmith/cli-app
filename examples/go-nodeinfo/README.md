# Command Line Application

Creating command line apps using various programming languages is easy with Fedora. For this
example we will write a simple go program that gathers node information and prints to stdout.

## Getting started

### Install required libraries

For go programs a text editor and the latest golang is all we need to get started.

Get the latest version of here [go](https://golang.org/dl/).

Install `go` using the following instructions:

```bash
sudo dnf install golang
```

### Create a workspace

```bash
mkdir work && cd work
```

## Command Line Application in Go

For this example we will use the `syscall.Uname` function in [syscall](https://golang.org/pkg/syscall/#Uname) go library to gather information about the system similar to the output of the `uname -a` command and format it into a human readable block.  

The output from our example cli should match the output of `uname` with the corresponding switches.

Here is an example of the `uname` output we are trying to achieve using `cat` to match the format

```bash
$ cat << EOF
Node Info
----------
Nodename: $(uname -n)
Sysname : $(uname -s)
Version: $(uname -v)
Release: $(uname -r)
Machine: $(uname -m)

EOF
```

The results look like this

```bash
Node Info
----------
Nodename: fedora
Sysname : Linux
Version: #1 SMP Fri Jul 17 16:23:37 UTC 2020
Release: 5.7.9-200.fc32.x86_64
Machine: x86_64
```

## Write code for command line application

Our example is a simple one file go program. To get started we need to write our go code into a file named `main.go`

You can write the go file using the editor of your choice.

```bash
vim main.go
```

Here is the example code for `main.go`

```go
package main

import (
 "flag"
 "fmt"
 "os"

 "github.com/fatih/structs"
 "golang.org/x/sys/unix"
)

const version string = "v1.0.0"

func getVersion() {
 fmt.Printf("goinfo version %s\n", version)
}

func fmtValue(i [65]uint8) string {
 b := make([]byte, 0, len(i))
 for _, v := range i {
  if v == 0x00 {
   break
  }
  b = append(b, v)
 }
 return string(b)
}

func main() {
 v := flag.Bool("version", false, "return version of cli")

 flag.Parse()

 if *v {
  getVersion()
  os.Exit(0)
 }

 var utsname unix.Utsname
 err := unix.Uname(&utsname)
 if err != nil {
  panic(err)
 }

 m := structs.Map(utsname)

 fmt.Println("Node Info\n-----------------------------------")
 for k, v := range m {
  fmt.Printf("%s : %s\n", k, fmtValue(v.([65]uint8)))
 }
}
```

Save the file as `main.go`.

## Compile code for command line application

We can compile our cli at the command line using `go build`. To do this open a terminal and navigate to your workspace.

The command is as follows.

```bash
CGO_ENABLED=0 go build -v  -o nodeinfo .
```

The `go build` command flag `-o .nodeinfo` tells `go` to generate a binary named `nodeinfo` in the same directory as `main.go`.

The binary should be executable so give it a spin.

```bash
./nodeinfo
```

And the output should be similar to this

```bash
Node Info
-----------------------------------
Version : #1 SMP Fri Jul 17 16:23:37 UTC 2020
Machine : x86_64
Domainname : (none)
Sysname : Linux
Nodename : fedora
Release : 5.7.9-200.fc32.x86_64

```

## Source Code

Source code for this and more examples

[https://github.com/xbcsmith/cli-app](https://github.com/xbcsmith/cli-app)
