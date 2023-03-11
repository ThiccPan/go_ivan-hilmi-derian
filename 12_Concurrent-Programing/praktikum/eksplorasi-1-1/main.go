package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
)

type Products struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Image       string  `json:"image"`
	Rating      Rating  `json:"rating"`
}

func (p Products) String() string {
	return fmt.Sprintf("===\ntitle: %v\nprice: %v\ncategory: %v\n===\n", p.Title, p.Price, p.Category)
}

type Rating struct {
	Rate  float64 `json:"rate"`
	Count int     `json:"count"`
}

var wg sync.WaitGroup

func main() {
	chRes := make(chan *http.Response, 3)

	// fetch data based on id
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			idStr := fmt.Sprint(id)

			res, err := http.Get("https://fakestoreapi.com/products/" + idStr)
			if err != nil {
				return
			}
			chRes <- res
		}(i)
	}
	wg.Wait()
	close(chRes)

	productList := make([]Products, 0)

	for {
		var product Products
		res, b := <-chRes
		fmt.Println(b)
		if !b {
			break
		}

		resBody, err2 := io.ReadAll(res.Body)
		if err2 != nil {
			os.Exit(2)
		}

		err3 := json.Unmarshal(resBody, &product)
		if err3 != nil {
			os.Exit(3)
		}
		productList = append(productList, product)
	}

	fmt.Println(productList)
}