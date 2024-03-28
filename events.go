package gosu

import "time"

type EventType string

const (
	Achievement       EventType = "achievement"
	BeatmapPlaycount  EventType = "beatmapPlaycount"
	BeatmapsetApprove EventType = "beatmapsetApprove"
	BeatmapsetDelete  EventType = "beatmapsetDelete"
	BeatmapsetRevive  EventType = "beatmapsetRevive"
	BeatmapsetUpdate  EventType = "beatmapsetUpdate"
	BeatmapsetUpload  EventType = "beatmapsetUpload"
	Rank              EventType = "rank"
	RankLost          EventType = "rankLost"
	UserSupportAgain  EventType = "userSupportAgain"
	UserSupportFirst  EventType = "userSupportFirst"
	UserSupportGift   EventType = "userSupportGift"
	UsernameChange    EventType = "usernameChange"
)

type EventBase struct {
	CreatedAt       time.Time `json:"created_at" json:"created_at"`
	ID              int
	Type            EventType              `json:"type"`
	EventAttributes map[string]interface{} `json:",remain"`
}

type EventBeatmap struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

type EventBeatmapset struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

type BeatmapPlaycountEvent struct {
	Beatmap EventBeatmap `json:"beatmap"`
	Count   uint32       `json:"count"`
}

type BeatmapsetApproveEvent struct {
	Approval   RankStatus      `json:"approval"`
	Beatmapset EventBeatmapset `json:"beatmapset"`
	User       EventUser       `json:"user"`
}

type BeatmapsetDeleteEvent struct {
	Beatmapset EventBeatmapset `json:"beatmapset"`
}

type BeatmapsetReviveEvent struct {
	Beatmapset EventBeatmapset `json:"beatmapset"`
	User       EventUser       `json:"user"`
}

type BeatmapsetUpdateEvent struct {
	Beatmapset EventBeatmapset `json:"beatmapset"`
	User       EventUser       `json:"user"`
}

type BeatmapsetUploadEvent struct {
	Beatmapset EventBeatmapset `json:"beatmapset"`
	User       EventUser       `json:"user"`
}

type MedalEvent struct {
	Achievement Medal     `json:"achievement"`
	User        EventUser `json:"user"`
}

type RankEvent struct {
	Grade   Grade        `json:"scoreRank"`
	Rank    uint32       `json:"rank"`
	Mode    Ruleset      `json:"mode"`
	Beatmap EventBeatmap `json:"beatmap"`
	User    EventUser    `json:"user"`
}

type RankLostEvent struct {
	Mode    Ruleset      `json:"mode"`
	Beatmap EventBeatmap `json:"beatmap"`
	User    EventUser    `json:"user"`
}

type SupportAgainEvent struct {
	User EventUser `json:"user"`
}

type SupportFirstEvent struct {
	User EventUser `json:"user"`
}

type SupportGiftEvent struct {
	User EventUser `json:"user"`
}

type UsernameChangeEvent struct {
	User EventUser `json:"user"`
}

type EventUser struct {
	Username         string  `json:"username"`
	URL              string  `json:"url"`
	PreviousUsername *string `json:"previousUsername,omitempty"`
}
