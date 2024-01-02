/*
Russian ntp servers: https://www.ntp-servers.net/servers.html
*/
package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	ntpAddress := flag.String("s", "ntp0.ntp-servers.net", "set ntp server's address")
	timeout := flag.Duration("t", 30*time.Second, "set timeout")
	flag.Parse()

	options := ntp.QueryOptions{Timeout: *timeout}
	response, err := ntp.QueryWithOptions(*ntpAddress, options)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	time := time.Now().Add(response.ClockOffset)
	fmt.Println(time)
}
