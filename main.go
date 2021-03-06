package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {

//	maxCardNumber := getUserInput("Enter maximum discount card number (10000 for example) ",
//						"Maximum discount number is:")
//	maxNumberOfItems := getUserInput("Enter maximum number of items (10 for example): ",
//						"Maximum number of items is:")
//	maxItemId := getUserInput("Enter maximum item ID (100 for example): ", "Maximum item ID is:")
//	oneItemMaxQuantity := getUserInput("Enter maximum quantity of one item (10 for example): ",
//						"Maximum one item quantity is:")
//	oneItemMaxPrice := getUserInput("Enter maximum price of one item: (100 for example) ",
//						"One item maximum price is:")
//	promotionDiscount := getUserInput("Enter promotion discount (10 for exmple): ",
//						"Promotion discount is:")

	recAuto := newReceiptAuto(10000, 10, 100, 10, 100, 10) //maxCardNumber, maxNumberOfItems,
				//maxItemId, oneItemMaxQuantity, oneItemMaxPrice, promotionDiscount
//	printReceiptAuto(recAut)

//	recMan := newReceipt(100, 12,100,10,100,10.0)
//	printReceipt(recMan)

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		http.ServeFile(writer, request, "guitarshop.html")})

	http.HandleFunc("/instruments", func(writer http.ResponseWriter, request *http.Request) {
		http.ServeFile(writer, request, "instruments.html")
	})

	http.HandleFunc("/receipt", func(writer http.ResponseWriter, request *http.Request) {
		http.ServeFile(writer, request, "receipt.html")
	})

	http.HandleFunc("/autorec", func(writer http.ResponseWriter, request *http.Request) {
		templ, _ := template.ParseFiles("autorec.html")
		templ.Execute(writer, recAuto)
	})

	fmt.Println("Server is listening...")
	http.ListenAndServe("localhost:8181", nil)
}


