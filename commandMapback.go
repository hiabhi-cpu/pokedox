package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func commandMapBack(con *config) error {
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	var url string
	if con.Previous == "" {
		return fmt.Errorf("No previous locations present")
	} else {
		url = con.Previous
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("There was an error creating request", err)
	}
	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("There was an error in sending request", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Errorf("Erorr in reading json data", err)
	}

	var datamap LocationResponse
	json.Unmarshal([]byte(data), &datamap)

	// fmt.Println(datamap.Next)
	// fmt.Println(datamap.Previous)
	// fmt.Println(datamap.Results)
	locations := datamap.Results

	for _, loc := range locations {
		fmt.Println(loc.Name)
	}

	nextUrl := datamap.Next
	if nextUrl != "nil" {
		con.Next = nextUrl
	} else {
		con.Next = ""
	}

	previousUrl := datamap.Previous
	if previousUrl != "nil" {
		con.Previous = previousUrl
	} else {
		con.Previous = ""
	}

	return nil
}
