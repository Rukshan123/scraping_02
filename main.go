package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gocolly/colly"
)

func main() {
	district := District()
	category := Categery()

	scraping := colly.NewCollector()

	dataBase, _ := sql.Open("mysql", "root:ijse@tcp(127.0.0.1:3306)/ikman")

	scraping.OnHTML(".gtm-normal-ad", func(element *colly.HTMLElement) {
		model := element.ChildText(".heading--2eONR")
		descr := element.ChildText(".description--2-ez3")
		price := element.ChildText(".price--3SnqI")
		fmt.Println("\n")
		fmt.Println("\tmodel: ", model)
		fmt.Println("\tprice: ", price)
		fmt.Println("\tdescr: ", descr)

		insert, err := dataBase.Query("INSERT INTO Adds (district, category, model, price, descr) VALUES (?, ?, ?, ?, ?)", district, category, model, price, descr)
		check(err)
		defer insert.Close()
	})

	_ = scraping.Visit("https://ikman.lk/en/ads/" + district + "/" + category)

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
	} else if district == 2 {
		return districts[district-1]
	} else if district == 3 {
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
