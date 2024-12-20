package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("URL in GO")

	// String into URL
	myURL := "http://host.company.com/showCompanyInfo?name=C%26H%20Sugar"
	parsedURL, err := url.Parse(myURL)
	if err != nil {
		fmt.Println("Error : ", err)
		return
	}

	fmt.Printf("Type of URL : %T ", parsedURL)

	/* Things that can retrive from URL
	- Scheme : http / https
	- Host : specifies domain name & optionally port no
	- Path : represent path comp in URL
	- RawQuery : Includes query params
	*/

	fmt.Println()
	fmt.Println("Scheme : ", parsedURL.Scheme)
	fmt.Println("Host : ", parsedURL.Host)
	fmt.Println("Path : ", parsedURL.Path)
	fmt.Println("RawQuery : ", parsedURL.RawQuery)

}
