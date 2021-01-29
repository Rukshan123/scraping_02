package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	district := District()
	category := Categery()

	scraping := colly.NewCollector()

	scraping.OnHTML(".gtm-normal-ad", func(element *colly.HTMLElement) {
		fmt.Println("\n")
		fmt.Println("\tModel: ", element.ChildText(".heading--2eONR"))
		fmt.Println("\tPrice: ", element.ChildText(".price--3SnqI"))
		fmt.Println("\tDesc: ", element.ChildText(".description--2-ez3"))
	})

	scraping.OnRequest(func(r *colly.Request) {
		fmt.Println(r.URL.String())
	})
	scraping.OnResponse(func(r *colly.Response) {
		fmt.Println(r.Request.URL)

	})

	_ = scraping.Visit("https://ikman.lk/en/ads/" + district + "/" + category)
}

func District() string {
	districts := [3]string{"colombo", "galle", "gampaha"}
	for i, d := range districts {
		fmt.Println(i+1, ". ", d)
	}
	fmt.Print("district: ")
	var district int
	_, err := fmt.Scanf("%d", &district)
	check(err)
	if district == 1 {
		return districts[district-1]
	} else {
		return "Cannot Found"
	}
}

func check(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

func Categery() string {
	categeries := [3]string{"electronics", "vehicles", "education"}
	for i, c := range categeries {
		fmt.Println("\t", i+1, ". ", c)
	}

	fmt.Print("category: ")
	var category int
	_, err := fmt.Scanf("%d", &category)
	check(err)

	if category == 1 {
		return categeries[category-1]
	} else if category == 2 {
		return categeries[category-1]
	} else if category == 3 {
		return categeries[category-1]
	} else {
		return "Cannot Found"
	}
}
