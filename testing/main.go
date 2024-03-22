package main

import (
	"encoding/json"
	"fmt"
	"gosu"
)

func main() {

	client, err := gosu.NewClient(28457, "KtJnJPSXA0VMFPGg5kaLAB9jsXVEUuyhRv5FFvpD")
	if err != nil {
		panic(err)
	}

	score, err := client.GetUserRecentActivity(11517157).Build()
	if err != nil {
		panic(err)
	}

	s, _ := json.MarshalIndent(score, "", "\t")
	fmt.Print(string(s))

}
