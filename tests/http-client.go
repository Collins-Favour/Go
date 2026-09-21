package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("https://gobyexample.com/http-client")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println("Response Status:", resp.Status)
}
