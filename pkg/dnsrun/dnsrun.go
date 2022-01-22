package dnsrun

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/gookit/config"
)

func DNSCheck(ctx context.Context) {
	tick := time.Tick(30 * time.Second)
	DNS_URL_CHECK1, _ := config.String("DNS_URL_CHECK1")
	DNS_URL_CHECK2, _ := config.String("DNS_URL_CHECK2")
	linkschecks := []string{
		DNS_URL_CHECK1,
		DNS_URL_CHECK2,
	}

	c := make(chan string)

	for {
		select {
		case <-tick:
		case <-ctx.Done():
			return
		}
		fmt.Println("DNS Check ...")
		for _, linkscheck := range linkschecks {
			go CheckCname(linkscheck, c)
		}
	}
}

func DnsCheck() {
	// Cncurrent
	links := []string{
		"twitter.com",
		"amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
		go CheckCname(link, c)
		log.Println("Check ", link)
	}

	for l := range c {
		time.Sleep(3 * time.Second)
		go checkLink(l, c)
		// fmt.Println(<- c)
	}
}

func CheckCname(link string, c chan string) {
	cname, _ := net.LookupCNAME(link)
	fmt.Println("cname: "+link+" ", cname)
	iprecords, _ := net.LookupIP(link)
	for _, ip := range iprecords {
		fmt.Println("IP Address: "+link+" ", ip)
	}
}

// Check Domain URL
func checkLink(link string, c chan string) {
	client := http.Client{Timeout: 5 * time.Second}
	_, err := client.Get(link)
	if err != nil {
		fmt.Println("Error: ", link, " might be down ", err)
		c <- link
	}
	fmt.Println(link + " is Up")
	c <- link
}
