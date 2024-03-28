package gosu

import (
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
	"reflect"
	"strconv"
	"time"
)

type BeatmapCompact struct {
	BeatmapsetID     int     `json:"beatmapset_id"`
	DifficultyRating float32 `json:"difficulty_rating"`
	ID               int     `json:"id"`
	Mode             Ruleset `json:"mode"`
	Status           string  `json:"status"`
	TotalLength      int     `json:"total_length"`
	UserID           int     `json:"user_id"`
	Version          string  `json:"version"`
}

type Beatmap struct {
	BeatmapCompact
	Accuracy      float32    `json:"accuracy"`
	AR            float32    `json:"ar"`
	BeatmapsetID  int        `json:"beatmapset_id"`
	BPM           *float32   `json:"bpm"`
	Convert       bool       `json:"convert"`
	CountCircles  int        `json:"count_circles"`
	CountSliders  int        `json:"count_sliders"`
	CountSpinners int        `json:"count_spinners"`
	CS            float32    `json:"cs"`
	DeletedAt     *time.Time `json:"deleted_at"`
	Drain         float32    `json:"drain"`
	HitLength     int        `json:"hit_length"`
	IsScoreable   bool       `json:"is_scoreable"`
	LastUpdated   time.Time  `json:"last_updated"`
	ModeInt       int        `json:"mode_int"`
	Passcount     int        `json:"passcount"`
	Playcount     int        `json:"playcount"`
	Ranked        int        `json:"ranked"`
	URL           string     `json:"url"`
}

type GetBeatmapsResponse struct {
	Beatmaps []GetBeatmapResponse `json:"beatmaps"`
}

type FailTimes struct {
	Exit *[]int `json:"exit"`
	Fail *[]int `json:"fail"`
}

type BeatmapScores struct {
	Scores []struct {
		Score
		User struct {
			Country `json:"country"`
			Cover   `json:"cover"`
		} `json:"user"`
	} `json:"scores"`
}

type UserBeatmapScores struct {
	Scores []Score `json:"scores"`
}

type UserBeatmapScore struct {
	Position int `json:"position"`
	Score    struct {
		Score
		Beatmap struct {
			Beatmap
			Checksum string `json:"checksum"`
		} `json:"beatmap"`
		User struct {
			Country `json:"country"`
			Cover   `json:"cover"`
		} `json:"user"`
	} `json:"score"`
}

type GetBeatmapResponse struct {
	Beatmap
	Beatmapset struct {
		Beatmapset
		Ratings []int `json:"ratings"`
	} `json:"beatmapset"`
	Checksum  *string   `json:"checksum"`
	FailTimes FailTimes `json:"failtimes"`
	MaxCombo  int       `json:"max_combo"`
}

type GetUserMostPlayedResponse struct {
	BeatmapID  int                `json:"beatmap_id"`
	Beatmap    *BeatmapCompact    `json:"beatmap"`
	Beatmapset *BeatmapsetCompact `json:"beatmapset"`
	Count      int                `json:"count"`
}

type LookupBeatmapRequest struct {
	client   *Client
	Checksum *string
	Filename *string
	ID       *int
}

// LookupBeatmap returns beatmap.
// https://osu.ppy.sh/docs/index.html#beatmaps
func (c *Client) LookupBeatmap() *LookupBeatmapRequest {
	return &LookupBeatmapRequest{client: c}
}

func (r *LookupBeatmapRequest) SetChecksum(checksum string) *LookupBeatmapRequest {
	r.Checksum = &checksum
	return r
}

func (r *LookupBeatmapRequest) SetFilename(filename string) *LookupBeatmapRequest {
	r.Filename = &filename
	return r
}

func (r *LookupBeatmapRequest) SetID(id int) *LookupBeatmapRequest {
	r.ID = &id
	return r
}

func (r *LookupBeatmapRequest) Build() (*GetBeatmapResponse, error) {
	req := r.client.httpClient.R().SetResult(&GetBeatmapResponse{})

	if r.Checksum != nil {
		req.SetQueryParam("checksum", *r.Checksum)
	}

	if r.Filename != nil {
		req.SetQueryParam("filename", *r.Filename)
	}

	if r.ID != nil {
		req.SetQueryParam("id", strconv.Itoa(*r.ID))
	}

	resp, err := req.Get("/beatmaps/lookup")
	if err != nil {
		return nil, err
	}

	return resp.Result().(*GetBeatmapResponse), nil
}

type GetUserBeatmapScoreRequest struct {
	client  *Client
	Beatmap int
	User    int
	Mode    *Ruleset
}

// GetUserBeatmapScore returns a User's score on a beatmap.
// https://osu.ppy.sh/docs/index.html#get-a-user-beatmap-score
func (c *Client) GetUserBeatmapScore(beatmap int, user int) *GetUserBeatmapScoreRequest {
	return &GetUserBeatmapScoreRequest{client: c, Beatmap: beatmap, User: user}
}

func (r *GetUserBeatmapScoreRequest) SetMode(mode Ruleset) *GetUserBeatmapScoreRequest {
	r.Mode = &mode
	return r
}

func (r *GetUserBeatmapScoreRequest) Build() (*UserBeatmapScore, error) {
	req := r.client.httpClient.R().SetResult(&UserBeatmapScore{})

	req.SetPathParams(map[string]string{
		"beatmap": strconv.Itoa(r.Beatmap),
		"user":    strconv.Itoa(r.User),
	})

	if r.Mode != nil {
		req.SetQueryParam("mode", string(*r.Mode))
	}

	resp, err := req.Get("/beatmaps/{beatmap}/scores/users/{user}")
	if err != nil {
		return nil, err
	}

	return resp.Result().(*UserBeatmapScore), nil
}

type GetUserBeatmapScoresRequest struct {
	client  *Client
	Beatmap int
	User    int
	Mode    *Ruleset
}

// GetUserBeatmapScores returns a User's scores on a beatmap.
// https://osu.ppy.sh/docs/index.html#get-a-user-beatmap-scores
func (c *Client) GetUserBeatmapScores(beatmap int, user int) *GetUserBeatmapScoresRequest {
	return &GetUserBeatmapScoresRequest{client: c, Beatmap: beatmap, User: user}
}

func (r *GetUserBeatmapScoresRequest) SetMode(mode Ruleset) *GetUserBeatmapScoresRequest {
	r.Mode = &mode
	return r
}

func (r *GetUserBeatmapScoresRequest) Build() (*UserBeatmapScores, error) {
	req := r.client.httpClient.R().SetResult(&UserBeatmapScores{})

	req.SetPathParams(map[string]string{
		"beatmap": strconv.Itoa(r.Beatmap),
		"user":    strconv.Itoa(r.User),
	})

	if r.Mode != nil {
		req.SetQueryParam("mode", string(*r.Mode))
	}

	resp, err := req.Get("/beatmaps/{beatmap}/scores/users/{user}/all")
	if err != nil {
		return nil, err
	}

	return resp.Result().(*UserBeatmapScores), nil
}

type GetBeatmapScoresRequest struct {
	client  *Client
	Beatmap int
	Mode    *Ruleset
}

// GetBeatmapScores returns a User's scores on a beatmap.
// https://osu.ppy.sh/docs/index.html#get-a-user-beatmap-scores
func (c *Client) GetBeatmapScores(beatmap int) *GetBeatmapScoresRequest {
	return &GetBeatmapScoresRequest{client: c, Beatmap: beatmap}
}

func (r *GetBeatmapScoresRequest) SetMode(mode Ruleset) *GetBeatmapScoresRequest {
	r.Mode = &mode
	return r
}

func (r *GetBeatmapScoresRequest) Build() (*BeatmapScores, error) {
	req := r.client.httpClient.R().SetResult(&BeatmapScores{})

	req.SetPathParams(map[string]string{
		"beatmap": strconv.Itoa(r.Beatmap),
	})

	if r.Mode != nil {
		req.SetQueryParam("mode", string(*r.Mode))
	}

	resp, err := req.Get("/beatmaps/{beatmap}/scores")
	if err != nil {
		return nil, err
	}

	println(resp.String())

	return resp.Result().(*BeatmapScores), nil
}

type GetBeatmapsRequest struct {
	client   *Client
	Beatmaps []int
}

// GetBeatmaps returns a list of beatmaps.
// https://osu.ppy.sh/docs/index.html#get-beatmaps
func (c *Client) GetBeatmaps(beatmaps []int) *GetBeatmapsRequest {
	return &GetBeatmapsRequest{client: c, Beatmaps: beatmaps}
}

func (r *GetBeatmapsRequest) Build() (*GetBeatmapsResponse, error) {
	req := r.client.httpClient.R().SetResult(&GetBeatmapsResponse{})

	for _, id := range r.Beatmaps {
		req.QueryParam.Add("ids[]", strconv.Itoa(id))
	}

	resp, err := req.Get("/beatmaps")
	if err != nil {
		return nil, err
	}

	return resp.Result().(*GetBeatmapsResponse), nil
}

type GetBeatmapRequest struct {
	client  *Client
	Beatmap int
}

// GetBeatmap returns beatmap data for the specified beatmap ID.
// https://osu.ppy.sh/docs/index.html#get-beatmaps
func (c *Client) GetBeatmap(beatmap int) *GetBeatmapRequest {
	return &GetBeatmapRequest{client: c, Beatmap: beatmap}
}

func (r *GetBeatmapRequest) Build() (*GetBeatmapResponse, error) {
	req := r.client.httpClient.R().SetResult(&GetBeatmapResponse{})

	req.SetPathParam("id", strconv.Itoa(r.Beatmap))

	resp, err := req.Get("/beatmaps/{id}")
	if err != nil {
		return nil, err
	}

	return resp.Result().(*GetBeatmapResponse), nil
}

type BaseDifficultyAttributes struct {
	MaxCombo   int     `json:"max_combo"`
	StarRating float64 `json:"star_rating"`
}

type OsuDifficultyAttributes struct {
	BaseDifficultyAttributes
	AimDifficulty        float64 `json:"aim_difficulty"`
	SpeedDifficulty      float64 `json:"speed_difficulty"`
	SpeedNoteCount       float64 `json:"speed_note_count"`
	FlashlightDifficulty int     `json:"flashlight_difficulty"`
	SliderFactor         float64 `json:"slider_factor"`
	ApproachRate         int     `json:"approach_rate"`
	OverallDifficulty    int     `json:"overall_difficulty"`
}

type TaikoDifficultyAttributes struct {
	BaseDifficultyAttributes
	StaminaDifficulty int `json:"stamina_difficulty"`
	RhythmDifficulty  int `json:"rhythm_difficulty"`
	ColourDifficulty  int `json:"colour_difficulty"`
	PeakDifficulty    int `json:"peak_difficulty"`
	GreatHitWindow    int `json:"great_hit_window"`
}

type FruitsDifficultyAttributes struct {
	BaseDifficultyAttributes
	ApproachRate float32 `json:"approach_rate"`
}

type ManiaDifficultyAttributes struct {
	BaseDifficultyAttributes
	GreatHitWindow float32 `json:"great_hit_window"`
}

type GetBeatmapAttributesRequest struct {
	client  *Client
	Beatmap int
	Mods    Mod
	Mode    *Ruleset
}

// GetBeatmapAttributes returns difficulty attributes of beatmap with specific mode and mods combination.
// https://osu.ppy.sh/docs/index.html#get-beatmap-attributes
func (c *Client) GetBeatmapAttributes(beatmap int) *GetBeatmapAttributesRequest {
	return &GetBeatmapAttributesRequest{client: c, Beatmap: beatmap}
}

func (r *GetBeatmapAttributesRequest) SetMods(mods Mod) *GetBeatmapAttributesRequest {
	r.Mods = mods
	return r
}

func (r *GetBeatmapAttributesRequest) AddMods(mods Mod) *GetBeatmapAttributesRequest {
	r.Mods |= mods
	return r
}

func (r *GetBeatmapAttributesRequest) SetMode(mode Ruleset) *GetBeatmapAttributesRequest {
	r.Mode = &mode
	return r
}

type GenericDifficultyAttributes struct {
	Attributes json.RawMessage `json:"attributes"`
}

func (r *GetBeatmapAttributesRequest) Build() (interface{}, error) {
	req := r.client.httpClient.R().SetPathParam("beatmap", strconv.Itoa(r.Beatmap))

	body := make(map[string]interface{})

	if reflect.ValueOf(r.Mods).IsValid() {
		body["mods"] = r.Mods
	}

	if r.Mode != nil {
		body["ruleset"] = r.Mode
	}

	resp, err := req.SetBody(body).Post("/beatmaps/{beatmap}/attributes")
	if err != nil {
		return nil, err
	}

	attributes := gjson.Get(resp.String(), "attributes")

	var result interface{}

	switch { // Response is unknown type, therefore check using present fields
	case attributes.Get("aim_difficulty").Exists():
		result = &OsuDifficultyAttributes{}
	case attributes.Get("stamina_difficulty").Exists():
		result = &TaikoDifficultyAttributes{}
	case attributes.Get("approach_rate").Exists():
		result = &FruitsDifficultyAttributes{}
	case attributes.Get("great_hit_window").Exists():
		result = &ManiaDifficultyAttributes{}
	default:
		return nil, fmt.Errorf("no matching attribute found")
	}

	return result, json.Unmarshal([]byte(attributes.String()), result)
}
