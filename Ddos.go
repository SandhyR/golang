package main

import "fmt"

func main()  {
	var ip string
	fmt.Println("IP:")
	fmt.Scanln(&ip)
	for {
		fmt.Println("DDOS TO", ip)
	}
}