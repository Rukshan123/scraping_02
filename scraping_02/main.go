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

	dataBase, _ := sql.Open("mysql", "root:ijse@tcp(127.0.0.1:3306)/ikman_02")

	var model string
	var descr string
	var price string
	var contact string
	var AddDesc string

	//get a details in selected type and district
	scraping.OnHTML(".gtm-normal-ad", func(element *colly.HTMLElement) {
		model = element.ChildText(".heading--2eONR")
		descr = element.ChildText(".description--2-ez3")
		price = element.ChildText(".price--3SnqI")
		url := element.ChildAttr(".card-link--3ssYv", "href")
		fmt.Println("\n")
		fmt.Println("\tmodel: ", model)
		fmt.Println("\tprice: ", price)
		fmt.Println("\tdescr: ", descr)
		fmt.Println("\turl: ", url)

		err := element.Request.Visit(url)
		check(err)

	})

	//get a contact details in the add
	scraping.OnHTML(".contact-name--m97Sb", func(element *colly.HTMLElement) {
		contact = element.Text
		fmt.Println("\tContact: ", contact)
	})

	//get a full description  in the add
	scraping.OnHTML(".description-section--oR57b > div > .description--1nRbz", func(element *colly.HTMLElement) {
		AddDesc = element.Text
		fmt.Println("\tAdd Desc: ", AddDesc)
		fmt.Println("}")
	})

	//save all data in database
	scraping.OnScraped(func(r *colly.Response) {
		insert, err := dataBase.Query("INSERT INTO Adds (district, category, model, price, descr, contact, AddDesc) VALUES (?, ?, ?, ?, ?, ?, ?)",
			district, category, model, price, descr, contact, AddDesc)
		check(err)
		defer insert.Close()
	})

	//Adds related to the selected district and type
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

func check(err error) {
	if err != nil {
		fmt.Println(err)
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
