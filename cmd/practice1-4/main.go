package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

func getPost(n int, wg *sync.WaitGroup) {
	defer wg.Done()

	s := fmt.Sprintf("%d", n)
	url := "https://jsonplaceholder.typicode.com/posts/" + s
	response, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		body, err := ioutil.ReadAll(response.Body)

		if err != nil {
			fmt.Println(err)
		}

		bodyString := string(body)
		fmt.Println(bodyString)
	}


}

func main() {
	postsAmount := 100
	var wg sync.WaitGroup
	wg.Add(postsAmount)

	for i := 1; i <= postsAmount; i++ {
		go getPost(i, &wg)
	}

	wg.Wait()
}
