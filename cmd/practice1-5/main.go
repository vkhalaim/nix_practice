package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
)

func getPost(n int, wg *sync.WaitGroup) {
	defer wg.Done()

	postId := fmt.Sprintf("%d", n)
	url := "https://jsonplaceholder.typicode.com/posts/" + postId
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


		postToFile(body, postId)
	}


}

func postToFile(post []byte, postNumber string) {
	directory := "././storage/posts/"
	fileName := directory + postNumber + ".txt"
	file, err := os.Create(fileName)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	err = ioutil.WriteFile(fileName, post, os.FileMode(0644))

	if err != nil {
		fmt.Println(err)
		return
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

