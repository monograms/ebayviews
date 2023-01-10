package main

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/fatih/color"
)

var (
	blue      = color.New(color.FgBlue)
	magenta   = color.New(color.FgMagenta)
	red       = color.New(color.FgRed)
	hiMagenta = color.New(color.FgHiMagenta)
)

func main() {
	Logo()
	time.Sleep(3 * time.Second)
	blue.Println("\n\t\t\t[*] Initialized.\n")
	for {
		switch Menu() {
		case 1:
			var itemURL string
			blue.Print("\n[$] Enter item url:  ")
			fmt.Scanln(&itemURL)
			fmt.Println()
			var visits int
			blue.Print("\n[#] Enter number of visits:  ")
			fmt.Println()
			fmt.Scanln(&visits)
			var proxyResponse string
			blue.Print("\n[?] Use proxies? (y/n):  ")
			fmt.Scanln(&proxyResponse)
			var userProxy string
			if strings.TrimSpace(proxyResponse) == "y" {
				blue.Print("\n[$] Enter proxy url:  ")
				fmt.Scanln(&userProxy)
				userProxy = strings.TrimSpace(userProxy)
			} else {
				userProxy = ""
			}
			var wgs sync.WaitGroup
			wgs.Add(visits)
			for i := 0; i < visits; i++ {
				go func() {
					err := VisitItemURL(itemURL, userProxy)
					if err != nil {
						red.Println(err)
					}
					blue.Println("[|] Viewed item:", itemURL)
					wgs.Done()
				}()
			}
			wgs.Wait()
			magenta.Println("\n[!] Done.")
		case 2:
			hiMagenta.Println("\n[!] Exiting...")
			time.Sleep(2 * time.Second)
			return
		default:
			hiMagenta.Println("\n[!] Exiting...")
			time.Sleep(2 * time.Second)
			return
		}
	}
}
