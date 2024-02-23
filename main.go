package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

var logger *log.Logger

func init() {
	logger = log.New(os.Stdout, "", log.LstdFlags)
}

func main() {

	ipFlag := flag.String("backend", "", "comma-separated list of backend IP addresses")
	flag.Parse()

	if *ipFlag == "" {
		fmt.Println("Please provide backend IP address")
		return
	}

	ips := strings.Split(*ipFlag, ",")

	for i, ip := range ips {
		ips[i] = strings.TrimSpace(ip)
	}

	servers := make([]*Server, len(ips))
	weight := 1

	for i, ip := range ips {
		servers[i] = &Server{URL: "https://" + ip, Weight: weight, Alive: true}
		weight++
	}

	for _, server := range servers {
		fmt.Printf("Server: URL=%s, Weight=%d, Alive=%t\n", server.URL, server.Weight, server.Alive)
	}

	roundRobin := NewRoundRobin(servers)

	http.Handle("/your-end-point", roundRobin)

	fmt.Println("Loadbalancer started.")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Printf("Error starting server: %s\n", err.Error())
	}
}
