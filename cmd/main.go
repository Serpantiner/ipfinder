package main

import (
	"fmt"
	"log"

	"github.com/Serpantiner/ipfinder"
)

func main() {
	finder := ipfinder.NewIPFinder("https://api.ipify.org")
	ip, err := finder.GetIP()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Your IP address is:", ip)
}
