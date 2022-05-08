package tcpscan

import (
	"fmt"
	"net"
)

func worker(ports, results chan int, addressInput string) {
	for p := range ports {
		address := fmt.Sprintf("%s:%d", addressInput, p)
		conn, err := net.Dial("tcp", address)

		if err != nil {
			results <- 0
			continue
		}

		conn.Close()
		results <- p

	}
}

func TcpScan(workersQuantity int, portsQuantity int, addressInput string) []int {

	if portsQuantity < 1 {
		fmt.Println("Ports quantity cannot be less than 1. Actual: ", portsQuantity)
	}

	ports := make(chan int, workersQuantity)
	results := make(chan int)
	var openPorts []int

	for i := 0; i < cap(ports); i++ {
		go worker(ports, results, addressInput)
	}

	go func() {
		for i := 1; i <= portsQuantity; i++ {
			ports <- i
		}
	}()

	for i := 0; i < (portsQuantity - 1); i++ {
		port := <-results

		if port != 0 {
			openPorts = append(openPorts, port)
		}

	}

	close(ports)
	close(results)

	return openPorts

}


