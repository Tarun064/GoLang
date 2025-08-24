package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("Learning the url handlinng in GOlang")
	myUrl := "https://www.youtube.com/watch?v=vu6ZQ-t1sUk&list=PLzjZaW71kMwSEVpdbHPr0nPo5zdzbDulm&index=30"
	fmt.Printf("type of myUrl is %T\n", myUrl)

	parsedUrl, err := url.Parse(myUrl)
	if err != nil {
		println("You got the error ", err)
	}
	fmt.Printf("Now type of url after parsing is %T\n", parsedUrl)

	fmt.Println("Scheme of URL is: ", parsedUrl.Scheme)
	fmt.Println("Host of Url is: ", parsedUrl.Host)
	fmt.Println("Path of url is: ", parsedUrl.Path)
	fmt.Println("RawQuery of url is: ", parsedUrl.RawQuery)

	//modify the url
	parsedUrl.Path = "/newPath"
	parsedUrl.RawQuery = "username=iamtarun"

	newUrl := parsedUrl.String()
	fmt.Println("new Url is: ", newUrl)
}
