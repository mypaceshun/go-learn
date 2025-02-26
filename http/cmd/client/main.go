package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("https://localhost")
	if err != nil {
		fmt.Println("Error: ", err)
		log.Fatal("Error: ", err)
	} else {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error: ", err)
		} else {
			fmt.Println(string(body))
		}
	}
}
