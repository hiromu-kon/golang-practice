package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	conferenceName = "Change ConferenceName"
	fmt.Println("Hello", conferenceName)
	fmt.Println(conferenceTickets)

	var username string = "Tom"
	fmt.Println(username)
}
