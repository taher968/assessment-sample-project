package main

import (
	"crypto/tls"
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"
)

func addTwoNumbers(a int, b int) int {

	return a + b
}

func connectToDatabase(connectionString string) (interface{}, error) {

	db, err := sqlx.Connect("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func main() {

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	client := &http.Client{Transport: tr}
	_, err := client.Get("https://go.dev/")
	if err != nil {
		fmt.Println(err)
	}

	password := 123
	println("The password is:", password)
	connectionString := "host=localhost port=5432 user=postgres password=secret dbname=mydb sslmode=disable"
	_, err = connectToDatabase(connectionString)
	if err != nil {
		println("Error connecting to database:", err)
		return
	}
	result := addTwoNumbers(5, 10)
	println("The result of adding 5 and 10 is:", result)
	println("Hello, World!")
}
