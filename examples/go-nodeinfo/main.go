// SPDX-FileCopyrightText: 2020 Brett Smith <xbcsmith@gmail.com>
// SPDX-License-Identifier: Apache-2.0

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
