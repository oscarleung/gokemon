package gokemon

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Get - id is optional, only first one will be used
func Get(types string, id ...int) (*http.Response, error) {
	//func Get(types string, id ...int) (*bytes.Buffer, error) {
	resp, err := http.Get(BaseURL + types)
	if err != nil {
		fmt.Println("Failed request")
	}
	fmt.Println(BaseURL + types)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("fail to read data")
	}

	ToJSON(body)
	//fmt.Printf("%s\n", string(body))
	return resp, err
}

// ToJSON -
func ToJSON(body []byte) Pokedex {
	var jsonResult Pokedex
	jsonReader := bytes.NewReader(body)
	dec := json.NewDecoder(jsonReader)

	if err := dec.Decode(&jsonResult); err != nil {
		// bad things happen
		fmt.Println("Fail to decode")
	}
	return jsonResult
}
