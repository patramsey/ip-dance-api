package main

import (
	"fmt"
	"ipdance"
	"log"
	"os"
)

func main() {
	os.Setenv("IP_DANCE_URL", "https://api.ip.dance")
	os.Setenv("IP_DANCE_USERAGENT", "golang-ip-dance-package-0.1")

	ip, err := ipdance.GetIp()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Current external IP address:", ip)
}
