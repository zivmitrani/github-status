package github

type Page struct {
	id         string `json:"id"`
	name       string `json:"name"`
	url        string `json:"url"`
	updated_at string `json:"updated_at"`
}

type Status struct {
	description string `json:"description"`
	indicator   string `json:"indicator"`
}

type GHStatus struct {
	Statuses Status `json:"status"`
	Pages    Page   `json:"page"`
}

var mockData = `
{
	"page":{
	  "id":"kctbh9vrtdwd",
	  "name":"GitHub",
	  "url":"https://www.githubstatus.com",
	  "updated_at": "2023-02-07T17:26:48Z"
	},
	"status": {
	  "description": "Partial System Outage",
	  "indicator": "major"
	}
  }

`
