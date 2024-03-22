package gosu

import "time"

type EventType string

const (
	Achievement           EventType = "achievement"
	BeatmapPlaycountEvent EventType = "beatmapPlaycount"
	BeatmapsetApprove     EventType = "beatmapsetApprove"
	BeatmapsetDelete      EventType = "beatmapsetDelete"
	BeatmapsetRevive      EventType = "beatmapsetRevive"
	BeatmapsetUpdate      EventType = "beatmapsetUpdate"
	BeatmapsetUpload      EventType = "beatmapsetUpload"
	Rank                  EventType = "rank"
	RankLost              EventType = "rankLost"
	UserSupportAgain      EventType = "userSupportAgain"
	UserSupportFirst      EventType = "userSupportFirst"
	UserSupportGift       EventType = "userSupportGift"
	UsernameChange        EventType = "usernameChange"
)

type EventBase struct {
	CreatedAt time.Time `json:"created_at"`
	ID        int
	Type      EventType `json:"type"`
}
