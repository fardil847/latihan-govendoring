package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// h8HelperRand "github.com/novalagung/gubrak/v2"

func main() {
	// fmt.Println(h8HelperRand.RandomInt(10, 20))

	// res, err := http.Get("https://jsonplaceholder.typecode.com/posts/1")
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// fmt.Println(res.Body)
	// body, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// defer res.Body.Close()
	// sb := string(body)
	// log.Println(sb)

	data := map[string]interface{}{
		"title":  "Airell",
		"body":   "Jordan",
		"userId": 1,
	}

	requestJson, err := json.Marshal(data)
	client := &http.Client{}
	if err != nil {
		log.Fatalln(err)
	}
	req, err := http.NewRequest("POST", "https://jsonplaceholder.typecode.com/posts",
		bytes.NewBuffer(requestJson))
	req.Header.Set("Content-type", "application/json")
	if err != nil {
		log.Fatalln(err)
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(body)

}
