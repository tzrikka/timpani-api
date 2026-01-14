package github

// Reactions is based on:
// https://docs.github.com/en/rest/reactions/reactions?apiVersion=2022-11-28
type Reactions struct {
	// URL string `json:"url"` // Unnecessary.

	TotalCount int `json:"total_count,omitzero"`
	PlusOne    int `json:"+1,omitzero"`
	MinusOne   int `json:"-1,omitzero"`
	Laugh      int `json:"laugh,omitzero"`
	Hooray     int `json:"hooray,omitzero"`
	Confused   int `json:"confused,omitzero"`
	Heart      int `json:"heart,omitzero"`
	Rocket     int `json:"rocket,omitzero"`
	Eyes       int `json:"eyes,omitzero"`
}
