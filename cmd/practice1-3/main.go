package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	url := "https://jsonplaceholder.typicode.com/posts"
	response, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		body, err := ioutil.ReadAll(response.Body)

		if err != nil {
			log.Fatal(err)
		}

		bodyString := string(body)
		fmt.Println(bodyString)
	}

}