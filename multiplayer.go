package gosu

import (
	"strconv"
)

type MultiplayerScoresSort string

const (
	MultiplayerScoresSortAscending  = "score_asc"
	MultiplayerScoresSortDescending = "score_desc"
)

type MultiplayerScoresParams struct {
	Limit int                   `json:"limit"`
	Sort  MultiplayerScoresSort `json:"sort"`
}

type MultiplayerScoreMod struct {
	Acronym Mod `json:"acronym"`
}

type MultiplayerScoreStatistics struct {
	ComboBreak int `json:"combo_break"`
	Good       int `json:"good"`
	Great      int `json:"great"`
	IgnoreHit  int `json:"ignore_hit"`
	IgnoreMiss int `json:"ignore_miss"`
	Meh        int `json:"meh"`
	Miss       int `json:"miss"`
	Ok         int `json:"ok"`
	Perfect    int `json:"perfect"`
}

type MultiplayerScore struct {
	ID             int                        `json:"id"`
	UserID         int                        `json:"user_id"`
	RoomID         int                        `json:"room_id"`
	PlaylistItemID int                        `json:"playlist_item_id"`
	BeatmapID      int                        `json:"beatmap_id"`
	Rank           Grade                      `json:"rank"`
	TotalScore     int                        `json:"total_score"`
	Accuracy       float32                    `json:"accuracy"`
	MaxCombo       int                        `json:"max_combo"`
	Mods           []MultiplayerScoreMod      `json:"mods"`
	Statistics     MultiplayerScoreStatistics `json:"statistics"`
	Passed         bool                       `json:"passed"`
	Position       *int                       `json:"position"`
	User           struct {
		UserCompact
		Country Country `json:"country"`
		Cover   Cover   `json:"cover"`
	} `json:"user"`
}

type MultiplayerScores struct {
	CursorString Cursor                  `json:"cursor"`
	Params       MultiplayerScoresParams `json:"params"`
	Scores       []MultiplayerScore      `json:"scores"`
	Total        *int                    `json:"total"`
	UserScore    *MultiplayerScore       `json:"user_score"`
}

type PlaylistScoresRequest struct {
	client       *Client
	Room         int
	Playlist     int
	Limit        *int
	Sort         *MultiplayerScoresSort
	CursorString *string
}

func (c *Client) GetPlaylistScores(room, playlist int) *PlaylistScoresRequest {
	return &PlaylistScoresRequest{client: c, Room: room, Playlist: playlist}
}

func (r *PlaylistScoresRequest) SetLimit(limit int) *PlaylistScoresRequest {
	r.Limit = &limit
	return r
}

func (r *PlaylistScoresRequest) SetSort(sort MultiplayerScoresSort) *PlaylistScoresRequest {
	r.Sort = &sort
	return r
}

func (r *PlaylistScoresRequest) SetCursorString(cursorString string) *PlaylistScoresRequest {
	r.CursorString = &cursorString
	return r
}

func (r *PlaylistScoresRequest) Build() (*MultiplayerScores, error) {
	req := r.client.httpClient.R().SetResult(&MultiplayerScores{}).SetPathParams(map[string]string{
		"room":     strconv.Itoa(r.Room),
		"playlist": strconv.Itoa(r.Playlist),
	})

	if r.Limit != nil {
		req.SetQueryParam("limit", strconv.Itoa(*r.Limit))
	}

	if r.Sort != nil {
		req.SetQueryParam("sort", string(*r.Sort))
	}

	if r.CursorString != nil {
		req.SetQueryParam("cursor", *r.CursorString)
	}

	resp, err := req.Get("rooms/{room}/playlist/{playlist}/scores")
	if err != nil {
		return nil, err
	}

	return resp.Result().(*MultiplayerScores), nil
}
