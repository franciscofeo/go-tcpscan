package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/projects/tcpscan"
)

func main() {
	url := os.Args[1]
	workersString := os.Args[2]
	portsString := os.Args[3]

	workers, _ := strconv.Atoi(workersString)
	ports, _ := strconv.Atoi(portsString)

	fmt.Println("****** TCP PORT SCAN ******")
	portsList := tcpscan.TcpScan(workers, ports, url)
	tcpscan.PortsToTxt(portsList, url)
	fmt.Println("**************************")

}
