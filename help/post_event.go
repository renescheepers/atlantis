// Simple script to test the race condition
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

func main() {

	wg := sync.WaitGroup{}

	wg.Add(1)
	go sendToURL("http://localhost:4141/api/plan", &wg)
	wg.Add(1)
	go sendToURL("http://localhost:4242/api/plan", &wg)

	wg.Wait()
}

func sendToURL(url string, wg *sync.WaitGroup) {
	defer wg.Done()

	data := map[string]interface{}{
		"Repository": "renescheeepers-test-org/atlantis-test",
		"Ref":        "renescheepers/project-2",
		"Type":       "Github",
		"PR":         3,
		"Paths": []map[string]interface{}{
			{
				"Directory": "terraform/project-2",
				"Workspace": "default",
			},
		},
	}

	jsonBytes, _ := json.Marshal(data)

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBytes))
	if err != nil {
		log.Fatal(err)
	}

	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	request.Header.Set("X-Atlantis-Token", "secret")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Println("response Status:", response.Status)
	fmt.Println("response Headers:", response.Header)
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println("response Body:", string(body))

}
