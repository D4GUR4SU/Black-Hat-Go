package main

import (
	"fmt"
	"net"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	for i := 0; i <= 1024; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			address := fmt.Sprintf("scanme.nmap.org:%d", j)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				return
			}

			conn.Close()
			fmt.Printf("PORT: %d ===> OPEN\n", j)
		}(i)
	}
	wg.Wait()
}
