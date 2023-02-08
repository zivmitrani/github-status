package github

import (
	"encoding/json"
	"fmt"
)

func GetStatus(urlCheck string) (*Status, error) {
	response := httpget(urlCheck)

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

func main() {
	fmt.Println("Hello, World!")

	status, _ := GetStatus("https://www.githubstatus.com/api/v2/status.json")
	fmt.Println(status)
}
