package gosu

import (
	"github.com/mitchellh/mapstructure"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type Playstyle string

const (
	Mouse    Playstyle = "mouse"
	Keyboard Playstyle = "keyboard"
	Tablet   Playstyle = "tablet"
	Touch    Playstyle = "touch"
)

type ProfilePageSection string

const (
	Me             ProfilePageSection = "me"
	RecentActivity ProfilePageSection = "recent_activity"
	Beatmaps       ProfilePageSection = "beatmaps"
	Historical     ProfilePageSection = "historical"
	Kudosu         ProfilePageSection = "kudosu"
	TopRanks       ProfilePageSection = "top_ranks"
	Medals         ProfilePageSection = "medals"
)

type UserAccountHistoryType string

const (
	Note        UserAccountHistoryType = "note"
	Restriction UserAccountHistoryType = "restriction"
	Silence     UserAccountHistoryType = "silence"
)

type ScoreType string

const (
	Best   ScoreType = "best"
	Firsts ScoreType = "firsts"
	Pinned ScoreType = "pinned"
	Recent ScoreType = "recent"
)

type UserCompact struct {
	AvatarURL     string     `json:"avatar_url"`
	CountryCode   string     `json:"country_code"`
	DefaultGroup  *string    `json:"default_group"`
	ID            int        `json:"id"`
	IsActive      bool       `json:"is_active"`
	IsBot         bool       `json:"is_bot"`
	IsDeleted     bool       `json:"is_deleted"`
	IsOnline      bool       `json:"is_online"`
	IsSupporter   bool       `json:"is_supporter"`
	LastVisit     *time.Time `json:"last_visit"`
	PmFriendsOnly bool       `json:"pm_friends_only"`
	ProfileColour *string    `json:"profile_colour"`
	Username      string     `json:"username"`
}

type Cover struct {
	CustomURL *string `json:"custom_url"`
	URL       string  `json:"url"`
	ID        *string `json:"id"`
}

type Country struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type UserKudosu struct {
	Available int `json:"available"`
	Total     int `json:"total"`
}

type User struct {
	UserCompact
	Country      Country              `json:"country"`
	Cover        Cover                `json:"cover"`
	Discord      *string              `json:"discord"`
	HasSupported bool                 `json:"has_supported"`
	Interests    *string              `json:"interests"`
	JoinDate     time.Time            `json:"join_date"`
	Kudosu       UserKudosu           `json:"kudosu"`
	Location     string               `json:"location"`
	MaxBlocks    int                  `json:"max_blocks"`
	MaxFriends   int                  `json:"max_friends"`
	Occupation   *string              `json:"occupation"`
	Playmode     Ruleset              `json:"playmode"`
	Playstyle    []Playstyle          `json:"playstyle"`
	PostCount    int                  `json:"post_count"`
	ProfileOrder []ProfilePageSection `json:"profile_order"`
	Title        *string              `json:"title"`
	TitleURL     *string              `json:"title_url"`
	Twitter      *string              `json:"twitter"`
	Website      *string              `json:"website"`
}

type UserAccountHistory struct {
	Description string                 `json:"description"`
	ID          int                    `json:"id"`
	Length      int                    `json:"length"`
	Permanent   bool                   `json:"permanent"`
	Timestamp   string                 `json:"timestamp"`
	Type        UserAccountHistoryType `json:"type"`
}

type UserActiveTournamentBanner struct {
	ID           int    `json:"id"`
	TournamentID int    `json:"tournament_id"`
	Image        string `json:"image"`
}

type UserBadge struct {
	AwardedAt   time.Time `json:"awarded_at"`
	Description string    `json:"description"`
	ImageURL    string    `json:"image_url"`
	URL         string    `json:"url"`
}

type Page struct {
	HTML string `json:"html"`
	Raw  string `json:"raw"`
}

type Group struct {
	Colour         string `json:"colour"`
	HasListing     bool   `json:"has_listing"`
	HasPlaymodes   bool   `json:"has_playmodes"`
	ID             int    `json:"id"`
	Identifier     string `json:"identifier"`
	IsProbationary bool   `json:"is_probationary"`
	Name           string `json:"name"`
	ShortName      string `json:"short_name"`
}

type UserGroup struct {
	Group
	Playmodes []Ruleset `json:"playmodes"`
}

type MonthlyPlaycount struct {
	StartDate string `json:"start_date"`
	Count     int    `json:"count"`
}

type RankHighest struct {
	Rank      int       `json:"rank"`
	UpdatedAt time.Time `json:"updated_at"`
}

type RankHistory struct {
	Mode Ruleset `json:"mode"`
	Data []int   `json:"data"`
}

type GradeCounts struct {
	A   int `json:"a"`
	S   int `json:"s"`
	Sh  int `json:"sh"`
	Ss  int `json:"ss"`
	Ssh int `json:"ssh"`
}

type UserLevel struct {
	Current  int `json:"current"`
	Progress int `json:"progress"`
}

type UserStatistics struct {
	Count100               int         `json:"count_100"`
	Count300               int         `json:"count_300"`
	Count50                int         `json:"count_50"`
	CountMiss              int         `json:"count_miss"`
	GradeCounts            GradeCounts `json:"grade_counts"`
	HitAccuracy            float64     `json:"hit_accuracy"`
	IsRanked               bool        `json:"is_ranked"`
	Level                  UserLevel   `json:"level"`
	MaximumCombo           int         `json:"maximum_combo"`
	PlayCount              int         `json:"play_count"`
	PlayTime               int         `json:"play_time"`
	PP                     float64     `json:"pp"`
	GlobalRank             *int        `json:"global_rank"`
	RankedScore            int         `json:"ranked_score"`
	ReplaysWatchedByOthers int         `json:"replays_watched_by_others"`
	TotalHits              int         `json:"total_hits"`
	TotalScore             int         `json:"total_score"`
	CountryRank            *int        `json:"country_rank"`
}

type UserAchievement struct {
	AchievedAt    time.Time `json:"achieved_at"`
	AchievementID int       `json:"achievement_id"`
}

type UserExtended struct {
	User
	AccountHistory           []UserAccountHistory        `json:"account_history"`
	ActiveTournamentBanner   *UserActiveTournamentBanner `json:"active_tournament_banner"`
	Badges                   []UserBadge                 `json:"badges"`
	BeatmapPlaycountsCount   int                         `json:"beatmap_playcounts_count"`
	FavouriteBeatmapsetCount int                         `json:"favourite_beatmapset_count"`
	FollowerCount            int                         `json:"follower_count"`
	GraveyardBeatmapsetCount int                         `json:"graveyard_beatmapset_count"`
	Groups                   []UserGroup                 `json:"groups"`
	LovedBeatmapsetCount     int                         `json:"loved_beatmapset_count"`
	MappingFollowerCount     int                         `json:"mapping_follower_count"`
	MonthlyPlaycounts        []MonthlyPlaycount          `json:"monthly_playcounts"`
	Page                     Page                        `json:"page"`
	PendingBeatmapsetCount   int                         `json:"pending_beatmapset_count"`
	PreviousUsernames        []string                    `json:"previous_usernames"`
	RankHighest              *RankHighest                `json:"rank_highest"`
	RankHistory              RankHistory                 `json:"rank_history"`
	RankedBeatmapsetCount    int                         `json:"ranked_beatmapset_count"`
	ReplaysWatchedCounts     []MonthlyPlaycount          `json:"replays_watched_counts"`
	ScoresBestCount          int                         `json:"scores_best_count"`
	ScoresFirstCount         int                         `json:"scores_first_count"`
	ScoresRecentCount        int                         `json:"scores_recent_count"`
	Statistics               UserStatistics              `json:"statistics"`
	SupportLevel             int                         `json:"support_level"`
	UserAchievements         []UserAchievement           `json:"user_achievements"`
}

type Medal struct {
	Description  string   `json:"description"`
	Grouping     string   `json:"grouping"`
	IconURL      string   `json:"icon_url"`
	Instructions string   `json:"instructions"`
	MedalID      uint32   `json:"id"`
	Mode         *Ruleset `json:"mode"`
	Name         string   `json:"name"`
	Ordering     uint32   `json:"ordering"`
	Slug         string   `json:"slug"`
}

type UserScore struct {
	Score
	Beatmap struct {
		Beatmap
		Checksum *string `json:"checksum"`
	} `json:"beatmap"`
	Beatmapset BeatmapsetCompact `json:"beatmapset"`
	User       UserCompact       `json:"user"`
	Weight     *ScoreWeight      `json:"weight"`
}

type UserBeatmapset struct {
	Beatmapset
	Beatmaps []struct {
		Beatmap
		Checksum string `json:"checksum"`
	} `json:"beatmaps"`
}

type StatisticsRulesets map[Ruleset]*UserStatistics

type GetUsersResponse struct {
	Users []struct {
		UserCompact
		Country            `json:"country"`
		Cover              `json:"cover"`
		Groups             []UserGroup        `json:"groups"`
		StatisticsRulesets StatisticsRulesets `json:"statistics_rulesets"`
	} `json:"users"`
}

type GetUserKudosuRequest struct {
	client *Client
	User   int
	Limit  *int
	Offset *int
}

func (c *Client) GetUserKudosu(user int) *GetUserKudosuRequest {
	return &GetUserKudosuRequest{client: c, User: user}
}

func (r *GetUserKudosuRequest) SetLimit(limit int) *GetUserKudosuRequest {
	r.Limit = &limit
	return r
}

func (r *GetUserKudosuRequest) SetOffset(offset int) *GetUserKudosuRequest {
	r.Offset = &offset
	return r
}

func (r *GetUserKudosuRequest) Build() (*[]KudosuHistory, error) {
	req := r.client.httpClient.R().SetResult(&[]KudosuHistory{})

	req.SetPathParam("user", strconv.Itoa(r.User))

	if r.Limit != nil {
		req.SetQueryParam("limit", strconv.Itoa(*r.Limit))
	}

	if r.Offset != nil {
		req.SetQueryParam("offset", strconv.Itoa(*r.Offset))
	}

	resp, err := req.Get("/users/{user}/kudosu")
	if err != nil {
		return nil, err
	}

	return resp.Result().(*[]KudosuHistory), nil
}

type GetUserScoresRequest struct {
	client       *Client
	User         int
	Type         ScoreType
	IncludeFails *bool
	Mode         *Ruleset
	Limit        *int
	Offset       *int
}

func (c *Client) GetUserScores(user int) *GetUserScoresRequest {
	return &GetUserScoresRequest{client: c, User: user}
}

func (r *GetUserScoresRequest) Best() *GetUserScoresRequest {
	r.Type = Best
	return r
}

func (r *GetUserScoresRequest) Firsts() *GetUserScoresRequest {
	r.Type = Firsts
	return r
}

func (r *GetUserScoresRequest) Pinned() *GetUserScoresRequest {
	r.Type = Pinned
	return r
}

func (r *GetUserScoresRequest) Recent() *GetUserScoresRequest {
	r.Type = Recent
	return r
}

func (r *GetUserScoresRequest) SetIncludeFails(includeFails bool) *GetUserScoresRequest {
	r.IncludeFails = &includeFails
	return r
}

func (r *GetUserScoresRequest) SetLimit(limit int) *GetUserScoresRequest {
	r.Limit = &limit
	return r
}

func (r *GetUserScoresRequest) SetOffset(offset int) *GetUserScoresRequest {
	r.Offset = &offset
	return r
}

func (r *GetUserScoresRequest) Build() (*[]UserScore, error) {
	req := r.client.httpClient.R().SetResult(&[]UserScore{})

	req.SetPathParams(map[string]string{
		"user": strconv.Itoa(r.User),
		"type": "best",
	})

	if reflect.ValueOf(r.Type).IsValid() {
		req.SetQueryParam("type", string(r.Type))
	}

	if r.IncludeFails != nil {
		i := 0
		if *r.IncludeFails {
			i = 1
		}
		req.SetPathParam("include_fails", strconv.Itoa(i))
	}

	if r.Mode != nil {
		req.SetQueryParam("mode", string(*r.Mode))
	}

	if r.Limit != nil {
		req.SetQueryParam("limit", strconv.Itoa(*r.Limit))
	}

	if r.Offset != nil {
		req.SetQueryParam("offset", strconv.Itoa(*r.Offset))
	}

	resp, err := req.Get("/users/{user}/scores/{type}")
	if err != nil {
		return nil, err
	}

	return resp.Result().(*[]UserScore), nil
}

type GetUserBeatmapsetsRequest struct {
	client  *Client
	User    int
	MapType string
	Limit   *int
	Offset  *int
}

func (c *Client) GetUserBeatmapsets(user int) *GetUserBeatmapsetsRequest {
	return &GetUserBeatmapsetsRequest{client: c, User: user, MapType: "ranked"}
}

func (r *GetUserBeatmapsetsRequest) SetStatus(mapType RankStatus) *GetUserBeatmapsetsRequest {
	switch mapType {
	case RankStatusApproved, RankStatusRanked:
		r.MapType = "ranked"
	case RankStatusGraveyard:
		r.MapType = "graveyard"
	case RankStatusPending, RankStatusWIP, RankStatusQualified:
		r.MapType = "pending"
	case RankStatusLoved:
		r.MapType = "loved"
	}
	return r
}

func (r *GetUserBeatmapsetsRequest) Ranked() *GetUserBeatmapsetsRequest {
	r.MapType = "ranked"
	return r
}

func (r *GetUserBeatmapsetsRequest) Loved() *GetUserBeatmapsetsRequest {
	r.MapType = "loved"
	return r
}

func (r *GetUserBeatmapsetsRequest) Pending() *GetUserBeatmapsetsRequest {
	r.MapType = "pending"
	return r
}

func (r *GetUserBeatmapsetsRequest) Graveyard() *GetUserBeatmapsetsRequest {
	r.MapType = "graveyard"
	return r
}

func (r *GetUserBeatmapsetsRequest) SetLimit(limit int) *GetUserBeatmapsetsRequest {
	r.Limit = &limit
	return r
}

func (r *GetUserBeatmapsetsRequest) SetOffset(offset int) *GetUserBeatmapsetsRequest {
	r.Offset = &offset
	return r
}

func (r *GetUserBeatmapsetsRequest) Build() (*[]UserBeatmapset, error) {
	req := r.client.httpClient.R().SetResult(&[]UserBeatmapset{})

	req.SetPathParams(map[string]string{
		"user": strconv.Itoa(r.User),
		"type": r.MapType,
	})

	if r.Limit != nil {
		req.SetQueryParam("limit", strconv.Itoa(*r.Limit))
	}

	if r.Offset != nil {
		req.SetQueryParam("offset", strconv.Itoa(*r.Offset))
	}

	resp, err := req.Get("/users/{user}/beatmapsets/{type}")
	if err != nil {
		return nil, err
	}

	return resp.Result().(*[]UserBeatmapset), nil

}

type GetUserMostPlayedRequest struct {
	client *Client
	User   int
	Limit  *int
	Offset *int
}

func (c *Client) GetUserMostPlayed(user int) *GetUserMostPlayedRequest {
	return &GetUserMostPlayedRequest{client: c, User: user}
}

func (r *GetUserMostPlayedRequest) SetLimit(limit int) *GetUserMostPlayedRequest {
	r.Limit = &limit
	return r
}

func (r *GetUserMostPlayedRequest) SetOffset(offset int) *GetUserMostPlayedRequest {
	r.Offset = &offset
	return r
}

func (r *GetUserMostPlayedRequest) Build() (*[]GetUserMostPlayedResponse, error) {
	req := r.client.httpClient.R().SetResult(&[]GetUserMostPlayedResponse{})

	req.SetPathParams(map[string]string{
		"user": strconv.Itoa(r.User),
		"type": "most_played",
	})

	if r.Limit != nil {
		req.SetQueryParam("limit", strconv.Itoa(*r.Limit))
	}

	if r.Offset != nil {
		req.SetQueryParam("offset", strconv.Itoa(*r.Offset))
	}

	resp, err := req.Get("/users/{user}/beatmapsets/{type}")
	if err != nil {
		return nil, err
	}

	return resp.Result().(*[]GetUserMostPlayedResponse), nil
}

type GetUserRecentActivityRequest struct {
	client *Client
	User   int
	Limit  *int
	Offset *int
}

func (c *Client) GetUserRecentActivity(user int) *GetUserRecentActivityRequest {
	return &GetUserRecentActivityRequest{client: c, User: user}
}

func (r *GetUserRecentActivityRequest) SetLimit(limit int) *GetUserRecentActivityRequest {
	r.Limit = &limit
	return r
}

func (r *GetUserRecentActivityRequest) SetOffset(offset int) *GetUserRecentActivityRequest {
	r.Offset = &offset
	return r
}

func (r *GetUserRecentActivityRequest) Build() (*[]EventBase, error) {
	req := r.client.httpClient.R().SetResult(&[]map[string]interface{}{})

	req.SetPathParam("user", strconv.Itoa(r.User))

	if r.Limit != nil {
		req.SetQueryParam("limit", strconv.Itoa(*r.Limit))
	}

	if r.Offset != nil {
		req.SetQueryParam("offset", strconv.Itoa(*r.Offset))
	}

	resp, err := req.Get("/users/{user}/recent_activity")
	if err != nil {
		return nil, err
	}

	event := resp.Result().(*[]map[string]interface{})

	var result []EventBase
	config := &mapstructure.DecoderConfig{
		DecodeHook: mapstructure.StringToTimeHookFunc(time.RFC3339),
		Result:     &result,
		TagName:    "json",
	}

	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return nil, err
	}

	err = decoder.Decode(*event)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

type UserKey string

const (
	Username UserKey = "username"
	ID       UserKey = "id"
)

type GetUserRequest struct {
	client *Client
	User   string
	Mode   *Ruleset
	Key    *UserKey
}

func (c *Client) GetUser(user string) *GetUserRequest {
	return &GetUserRequest{client: c, User: user}
}

func (r *GetUserRequest) SetMode(mode Ruleset) *GetUserRequest {
	r.Mode = &mode
	return r
}

func (r *GetUserRequest) SetKey(key UserKey) *GetUserRequest {
	r.Key = &key
	return r
}

func (r *GetUserRequest) Build() (*UserExtended, error) {
	req := r.client.httpClient.R().SetResult(&UserExtended{}).SetPathParam("user", r.User)

	var url strings.Builder
	url.WriteString("/users/{user}")

	if r.Mode != nil {
		url.WriteString("/" + string(*r.Mode))
	}

	if r.Key != nil {
		println("hello!")
	}

	resp, err := req.Get(url.String())
	if err != nil {
		println()
		return nil, err
	}

	return resp.Result().(*UserExtended), nil
}

type GetUsersRequest struct {
	client *Client
	Users  []int
}

func (c *Client) GetUsers(userIds []int) *GetUsersRequest {
	return &GetUsersRequest{client: c, Users: userIds}
}

func (r *GetUsersRequest) Build() (*GetUsersResponse, error) {
	req := r.client.httpClient.R().SetResult(&GetUsersResponse{})

	for _, id := range r.Users {
		req.QueryParam.Add("ids[]", strconv.Itoa(id))
	}

	resp, err := req.Get("/users")
	if err != nil {
		return nil, err
	}

	return resp.Result().(*GetUsersResponse), nil
}
