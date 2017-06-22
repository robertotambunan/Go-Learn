package main

import "fmt"
import "net/http"
import "log"
import "encoding/json"
import "os"
import "bufio"
import "net/url"

type ReturnAPI struct {
	Product []Product `json:"products"`
}
type Product struct {
	Id             string `json:"id"`
	Name           string `json:"name"`
	City           string `json:"city"`
	Price          int    `json:"price"`
	Category       string `json:"category"`
	SellerUsername string `json:"seller_username"`
	SellerName     string `json:"seller_name"`
	Province       string `json:"province"`
	Url            string `json:"url"`
	Weight         int    `json:"weight"`
	Stock          int    `json:"stock"`
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan keyword barang yang dicari: ")
	text, _ := reader.ReadString('\n')
	safeText := url.QueryEscape(text)

	url := fmt.Sprintf("https://api.bukalapak.com/v2/products.json?keywords=%s&page=1&per_page=24", safeText)
	//Build the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}
	// For control over HTTP client headers,
	// redirect policy, and other settings,
	// create a Client
	// A Client is "an HTTP client
	client := &http.Client{}

	// Send the request via a client
	// Do sends an HTTP request and
	// returns an HTTP response
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}

	// Callers should close resp.Body
	// when done reading from it
	// Defer the closing of the body
	defer resp.Body.Close()

	var record ReturnAPI
	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println(err)
	}
	idx := 0
	for _, element := range record.Product {
		fmt.Println("Produk No. ", idx)
		fmt.Println("ID Produk\t\t= ", element.Id)
		fmt.Println("Nama Produk\t\t= ", element.Name)
		fmt.Println("Kota Asal Produk\t= ", element.City)
		fmt.Println("Provinsi\t\t= ", element.Province)
		fmt.Println("Harga\t\t\t= ", element.Price)
		fmt.Println("Username Pelapak\t= ", element.SellerUsername)
		fmt.Println("Nama Pelapak\t\t= ", element.Province)
		fmt.Println("URL Produk\t\t= ", element.Url)
		fmt.Println("Berat Produk\t\t= ", element.Weight)
		fmt.Println("Stok\t\t\t= ", element.Stock)
		fmt.Println("\n")
		idx++
	}
}
