package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	fmt.Println("Welcome to Grawler!")
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("A site was not specifed, please specify a site to crawl")
	}

	ReadSite(args[0])
	fmt.Println(args[0])
}

func ReadSite(site string) {

	res, err  := http.Get(site)
	if err != nil {
		panic(err)
		fmt.Println("The site you specified probably does not exist. Try a differnet site")
	}

	body, err := ioutil.ReadAll(res.Body)
	if  err != nil {
		panic(err)
	}

	Find(string(body), "<div")
}

func Find(site, searchedString string) {
	//count := strings.Count(site, "href=")
	
	dex := strings.Index(site, searchedString)
	count := strings.Count(site, searchedString)

	fmt.Println(dex)
	fmt.Println(count)
	if dex > 0 {
		i := 0
 		for {
			res := site[dex+i]
			fmt.Print(string(res))
		
			if i > 500 {
				break
			}

			i++
		}
	}else {
		fmt.Println("The element/string that you're looking for does not exist, please try a different string")
	}
 }

/*
func Contains(elm, searchFor string) {
	//Get this working tomorrow!
}*/

//func PrintHTMLEl()
