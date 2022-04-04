package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/itchio/ox/syscallex"
)

func main() {
	intDomainname, _ := strconv.ParseUint(os.Args[1], 0, 16)
	intUsername, _ := strconv.ParseUint(os.Args[2], 0, 16)
	intOldpasswd, _ := strconv.ParseUint(os.Args[3], 0, 16)
	intNewpasswd, _ := strconv.ParseUint(os.Args[4], 0, 16)

	domainname := uint16(intDomainname)
	username := uint16(intUsername)
	oldpasswd := uint16(intOldpasswd)
	newpasswd := uint16(intNewpasswd)

	err := syscallex.NetUserChangePassword(&domainname, &username, &oldpasswd, &newpasswd)

	if err != nil {
		fmt.Println(err)
	}
}
