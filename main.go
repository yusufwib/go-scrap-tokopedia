package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/debug"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandomString() string {
	b := make([]byte, rand.Intn(10)+10)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func main() {
	// init the filename
	fileName := "tokped.csv"

	file, err := os.Create(fileName)
	// check if error
	if err != nil {
		log.Fatalf("err: :%q", err)
		return
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	//create the header of csv file
	writer.Write([]string{"Name", "Price", "Rating", "Store", "ProductLink"})

	collector := colly.NewCollector(
		colly.UserAgent(RandomString()),
		colly.Debugger(&debug.LogDebugger{}),
	)

	collector.OnHTML(".css-bk6tzz", func(element *colly.HTMLElement) {

		//find element that contains the value of the product

		productName := element.ChildText(".css-11s9vse span")
		productPrice := element.ChildText(".css-4u82jy span")
		produtRate := strconv.Itoa(len(element.ChildAttrs(".css-177n1u3", "alt")))
		productStore := element.ChildText(".css-vbihp9 span")
		productLink := element.ChildAttrs("a", "href")[0]

		writer.Write([]string{
			productName,
			productPrice,
			produtRate,
			productStore,
			productLink,
		})
	})

	for i := 1; i <= 10; i++ {
		collector.Visit(fmt.Sprintf("https://www.tokopedia.com/p/handphone-tablet/handphone?page=%d", i))
	}
}
