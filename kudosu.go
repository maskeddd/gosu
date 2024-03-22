package gosu

import "time"

type KudosuHistory struct {
	ID        int          `json:"id"`
	Action    string       `json:"action"`
	Amount    int          `json:"amount"`
	Model     string       `json:"model"`
	CreatedAt time.Time    `json:"created_at"`
	Giver     *KudosuGiver `json:"giver"`
	Post      KudosuPost   `json:"post"`
}

type KudosuGiver struct {
	URL      string `json:"url"`
	Username string `json:"username"`
}

type KudosuPost struct {
	URL   *string `json:"url"`
	Title string  `json:"title"`
}
