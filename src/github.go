package github

import (
	"encoding/json"
	"fmt"
)

func GetStatus(url string) (*Status, error) {
	url := url
	response := httpget(url)

	// For Debugging:
	fmt.Println(string(response))

	//JSON Parsing:
	var status = &Status{}

	if err := json.Unmarshal(response, status); err != nil {
		return nil, fmt.Errorf("this is sad! boo")
	}

	//happy flow
	return status, nil
}

func httpget(url string) []byte {
	fmt.Println(url)
	return []byte(mockData)
}
