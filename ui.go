package main

import "fmt"

func Logo() {
	fmt.Println()
	blue.Println(`
	██████████████████████████████████████████████
	█▄─▄─▀█▄─▄▄─█▄─▀█▀─▄█─▄▄─█▄─▀█▄─▄█▄─▄▄─█▄─█─▄█
	██─▄─▀██─▄█▀██─█▄█─██─██─██─█▄▀─███─▄█▀██▄─▄██
	▀▄▄▄▄▀▀▄▄▄▄▄▀▄▄▄▀▄▄▄▀▄▄▄▄▀▄▄▄▀▀▄▄▀▄▄▄▄▄▀▀▄▄▄▀▀
	                eBay View Bot`)
}

func Menu() int {
	magenta.Println("\n\n\t\t\t 1. Start\n\n\t\t\t 2. Exit")
	var input int
	blue.Print("\n[#] Enter a option number:  ")
	fmt.Scanln(&input)
	switch input {
	case 1:
		return 1
	case 2:
		return 2
	default:
		return 2
	}
}
