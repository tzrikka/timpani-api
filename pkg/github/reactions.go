package github

// Reactions is based on:
// https://docs.github.com/en/rest/reactions/reactions?apiVersion=2022-11-28
type Reactions struct {
	// URL string `json:"url"` // Unnecessary.

	TotalCount int `json:"total_count"`
	PlusOne    int `json:"+1"`
	MinusOne   int `json:"-1"`
	Laugh      int `json:"laugh"`
	Hooray     int `json:"hooray"`
	Confused   int `json:"confused"`
	Heart      int `json:"heart"`
	Rocket     int `json:"rocket"`
	Eyes       int `json:"eyes"`
}
