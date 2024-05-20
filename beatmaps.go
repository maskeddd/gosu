package gosu

import (
	"encoding/json"
	"github.com/mitchellh/mapstructure"
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
	Beatmaps []BeatmapResponse `json:"beatmaps"`
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

type BeatmapResponse struct {
	Beatmap
	Beatmapset struct {
		Beatmapset
		Ratings []int `json:"ratings"`
	} `json:"beatmapset"`
	Checksum  *string   `json:"checksum"`
	FailTimes FailTimes `json:"failtimes"`
	MaxCombo  int       `json:"max_combo"`
}

type UserMostPlayedResponse struct {
	BeatmapID  int                `json:"beatmap_id"`
	Beatmap    *BeatmapCompact    `json:"beatmap"`
	Beatmapset *BeatmapsetCompact `json:"beatmapset"`
	Count      int                `json:"count"`
}

type UserBeatmapScoreRequest struct {
	client  *Client
	Beatmap int
	User    int
	Mode    *Ruleset
}

// GetUserBeatmapScore returns a user's score on a beatmap.
func (c *Client) GetUserBeatmapScore(beatmap int, user int) *UserBeatmapScoreRequest {
	return &UserBeatmapScoreRequest{client: c, Beatmap: beatmap, User: user}
}

func (r *UserBeatmapScoreRequest) SetMode(mode Ruleset) *UserBeatmapScoreRequest {
	r.Mode = &mode
	return r
}

func (r *UserBeatmapScoreRequest) Build() (*UserBeatmapScore, error) {
	req := r.client.httpClient.R().SetResult(&UserBeatmapScore{})

	req.SetPathParams(map[string]string{
		"beatmap": strconv.Itoa(r.Beatmap),
		"user":    strconv.Itoa(r.User),
	})

	if r.Mode != nil {
		req.SetQueryParam("mode", r.Mode.String())
	}

	resp, err := req.Get("beatmaps/{beatmap}/scores/users/{user}")
	if err != nil {
		return nil, err
	}

	return resp.Result().(*UserBeatmapScore), nil
}

type UserBeatmapScoresRequest struct {
	client  *Client
	Beatmap int
	User    int
	Mode    *Ruleset
}

// GetUserBeatmapScores returns a user's scores on a beatmap.
func (c *Client) GetUserBeatmapScores(beatmap int, user int) *UserBeatmapScoresRequest {
	return &UserBeatmapScoresRequest{client: c, Beatmap: beatmap, User: user}
}

func (r *UserBeatmapScoresRequest) SetMode(mode Ruleset) *UserBeatmapScoresRequest {
	r.Mode = &mode
	return r
}

func (r *UserBeatmapScoresRequest) Build() (*UserBeatmapScores, error) {
	req := r.client.httpClient.R().SetResult(&UserBeatmapScores{})

	req.SetPathParams(map[string]string{
		"beatmap": strconv.Itoa(r.Beatmap),
		"user":    strconv.Itoa(r.User),
	})

	if r.Mode != nil {
		req.SetQueryParam("mode", r.Mode.String())
	}

	resp, err := req.Get("beatmaps/{beatmap}/scores/users/{user}/all")
	if err != nil {
		return nil, err
	}

	return resp.Result().(*UserBeatmapScores), nil
}

type BeatmapScoresRequest struct {
	client  *Client
	Beatmap int
	Mode    *Ruleset
}

// GetBeatmapScores returns the top scores for a beatmap. Depending on user preferences, this may only show legacy scores.
func (c *Client) GetBeatmapScores(beatmap int) *BeatmapScoresRequest {
	return &BeatmapScoresRequest{client: c, Beatmap: beatmap}
}

func (r *BeatmapScoresRequest) SetMode(mode Ruleset) *BeatmapScoresRequest {
	r.Mode = &mode
	return r
}

func (r *BeatmapScoresRequest) Build() (*BeatmapScores, error) {
	req := r.client.httpClient.R().SetResult(&BeatmapScores{})

	req.SetPathParams(map[string]string{
		"beatmap": strconv.Itoa(r.Beatmap),
	})

	if r.Mode != nil {
		req.SetQueryParam("mode", r.Mode.String())
	}

	resp, err := req.Get("beatmaps/{beatmap}/scores")
	if err != nil {
		return nil, err
	}

	return resp.Result().(*BeatmapScores), nil
}

type BeatmapsRequest struct {
	client   *Client
	Beatmaps []int
}

// GetBeatmaps returns a list of beatmaps.
func (c *Client) GetBeatmaps(beatmaps []int) *BeatmapsRequest {
	return &BeatmapsRequest{client: c, Beatmaps: beatmaps}
}

func (r *BeatmapsRequest) Build() (*GetBeatmapsResponse, error) {
	req := r.client.httpClient.R().SetResult(&GetBeatmapsResponse{})

	for _, id := range r.Beatmaps {
		req.QueryParam.Add("ids[]", strconv.Itoa(id))
	}

	resp, err := req.Get("beatmaps")
	if err != nil {
		return nil, err
	}

	return resp.Result().(*GetBeatmapsResponse), nil
}

type BeatmapRequest struct {
	client  *Client
	Beatmap int
}

// GetBeatmap returns beatmap data for the specified beatmap ID.
func (c *Client) GetBeatmap(beatmap int) *BeatmapRequest {
	return &BeatmapRequest{client: c, Beatmap: beatmap}
}

func (r *BeatmapRequest) Build() (*BeatmapResponse, error) {
	req := r.client.httpClient.R().SetResult(&BeatmapResponse{})

	req.SetPathParam("id", strconv.Itoa(r.Beatmap))

	resp, err := req.Get("beatmaps/{id}")
	if err != nil {
		return nil, err
	}

	return resp.Result().(*BeatmapResponse), nil
}

type LookupBeatmapRequest struct {
	client   *Client
	Checksum *string
	Filename *string
	ID       *int
}

// LookupBeatmap returns beatmap.
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

func (r *LookupBeatmapRequest) Build() (*BeatmapResponse, error) {
	req := r.client.httpClient.R().SetResult(&BeatmapResponse{})

	if r.Checksum != nil {
		req.SetQueryParam("checksum", *r.Checksum)
	}

	if r.Filename != nil {
		req.SetQueryParam("filename", *r.Filename)
	}

	if r.ID != nil {
		req.SetQueryParam("id", strconv.Itoa(*r.ID))
	}

	resp, err := req.Get("beatmaps/lookup")
	if err != nil {
		return nil, err
	}

	return resp.Result().(*BeatmapResponse), nil
}

type BaseDifficultyAttributes struct {
	Attributes struct {
		MaxCombo          int                `json:"max_combo"`
		StarRating        float32            `json:"star_rating"`
		RulesetAttributes map[string]float32 `json:"ruleset_attributes,remain"`
	} `json:"attributes"`
}

type OsuDifficultyAttributes struct {
	AimDifficulty        float32 `json:"aim_difficulty"`
	SpeedDifficulty      float32 `json:"speed_difficulty"`
	SpeedNoteCount       float32 `json:"speed_note_count"`
	FlashlightDifficulty float32 `json:"flashlight_difficulty"`
	SliderFactor         float32 `json:"slider_factor"`
	ApproachRate         float32 `json:"approach_rate"`
	OverallDifficulty    float32 `json:"overall_difficulty"`
}

type TaikoDifficultyAttributes struct {
	StaminaDifficulty float32 `json:"stamina_difficulty"`
	RhythmDifficulty  float32 `json:"rhythm_difficulty"`
	ColourDifficulty  float32 `json:"colour_difficulty"`
	PeakDifficulty    float32 `json:"peak_difficulty"`
	GreatHitWindow    float32 `json:"great_hit_window"`
}

type FruitsDifficultyAttributes struct {
	ApproachRate float32 `json:"approach_rate"`
}

type ManiaDifficultyAttributes struct {
	GreatHitWindow float32 `json:"great_hit_window"`
}

type BeatmapAttributesRequest struct {
	client  *Client
	Beatmap int
	Mods    Mod
	Mode    *Ruleset
}

// GetBeatmapAttributes returns difficulty attributes of beatmap with specific mode and mods combination.
func (c *Client) GetBeatmapAttributes(beatmap int) *BeatmapAttributesRequest {
	return &BeatmapAttributesRequest{client: c, Beatmap: beatmap}
}

func (r *BeatmapAttributesRequest) SetMods(mods Mod) *BeatmapAttributesRequest {
	r.Mods = mods
	return r
}

func (r *BeatmapAttributesRequest) AddMods(mods Mod) *BeatmapAttributesRequest {
	r.Mods |= mods
	return r
}

func (r *BeatmapAttributesRequest) SetMode(mode Ruleset) *BeatmapAttributesRequest {
	r.Mode = &mode
	return r
}

type GenericDifficultyAttributes struct {
	Attributes json.RawMessage `json:"attributes"`
}

func (r *BeatmapAttributesRequest) Build() (*BaseDifficultyAttributes, error) {
	req := r.client.httpClient.R().SetResult(&map[string]interface{}{}).SetPathParam("beatmap", strconv.Itoa(r.Beatmap))

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

	attributes := resp.Result().(*map[string]interface{})

	var result BaseDifficultyAttributes
	config := &mapstructure.DecoderConfig{
		Result:  &result,
		TagName: "json",
	}

	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return nil, err
	}

	err = decoder.Decode(*attributes)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
