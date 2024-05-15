package gosu

import (
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/mitchellh/mapstructure"
)

type Playstyle string

const (
	PlaystyleMouse    Playstyle = "mouse"
	PlaystyleKeyboard Playstyle = "keyboard"
	PlaystyleTablet   Playstyle = "tablet"
	PlaystyleTouch    Playstyle = "touch"
)

type ProfilePage string

const (
	ProfilePageMe             ProfilePage = "me"
	ProfilePageRecentActivity ProfilePage = "recent_activity"
	ProfilePageBeatmaps       ProfilePage = "beatmaps"
	ProfilePageHistorical     ProfilePage = "historical"
	ProfilePageKudosu         ProfilePage = "kudosu"
	ProfilePageTopRanks       ProfilePage = "top_ranks"
	ProfilePageMedals         ProfilePage = "medals"
)

type HistoryType string

const (
	HistoryTypeNote          HistoryType = "note"
	HistoryTypeRestriction   HistoryType = "restriction"
	HistoryTypeSilence       HistoryType = "silence"
	HistoryTypeTournamentBan HistoryType = "tournament_ban"
)

type ScoreType string

const (
	ScoreTypeBest   ScoreType = "best"
	ScoreTypeFirsts ScoreType = "firsts"
	ScoreTypePinned ScoreType = "pinned"
	ScoreTypeRecent ScoreType = "recent"
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
	Country      Country       `json:"country"`
	Cover        Cover         `json:"cover"`
	Discord      *string       `json:"discord"`
	HasSupported bool          `json:"has_supported"`
	Interests    *string       `json:"interests"`
	JoinDate     time.Time     `json:"join_date"`
	Kudosu       UserKudosu    `json:"kudosu"`
	Location     string        `json:"location"`
	MaxBlocks    int           `json:"max_blocks"`
	MaxFriends   int           `json:"max_friends"`
	Occupation   *string       `json:"occupation"`
	Playmode     Ruleset       `json:"playmode"`
	Playstyle    []Playstyle   `json:"playstyle"`
	PostCount    int           `json:"post_count"`
	ProfileOrder []ProfilePage `json:"profile_order"`
	Title        *string       `json:"title"`
	TitleURL     *string       `json:"title_url"`
	Twitter      *string       `json:"twitter"`
	Website      *string       `json:"website"`
}

type UserAccountHistory struct {
	Description string      `json:"description"`
	ID          int         `json:"id"`
	Length      int         `json:"length"`
	Permanent   bool        `json:"permanent"`
	Timestamp   string      `json:"timestamp"`
	Type        HistoryType `json:"type"`
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

type UserKudosuRequest struct {
	client *Client
	User   int
	Limit  *int
	Offset *int
}

func (c *Client) GetUserKudosu(user int) *UserKudosuRequest {
	return &UserKudosuRequest{client: c, User: user}
}

func (r *UserKudosuRequest) SetLimit(limit int) *UserKudosuRequest {
	r.Limit = &limit
	return r
}

func (r *UserKudosuRequest) SetOffset(offset int) *UserKudosuRequest {
	r.Offset = &offset
	return r
}

func (r *UserKudosuRequest) Build() (*[]KudosuHistory, error) {
	req := r.client.httpClient.R().SetResult(&[]KudosuHistory{})

	req.SetPathParam("user", strconv.Itoa(r.User))

	if r.Limit != nil {
		req.SetQueryParam("limit", strconv.Itoa(*r.Limit))
	}

	if r.Offset != nil {
		req.SetQueryParam("offset", strconv.Itoa(*r.Offset))
	}

	resp, err := req.Get("users/{user}/kudosu")
	if err != nil {
		return nil, err
	}

	return resp.Result().(*[]KudosuHistory), nil
}

type UserScoresRequest struct {
	client       *Client
	User         int
	Type         ScoreType
	IncludeFails *bool
	Mode         *Ruleset
	Limit        *int
	Offset       *int
}

func (c *Client) GetUserScores(user int) *UserScoresRequest {
	return &UserScoresRequest{client: c, User: user}
}

func (r *UserScoresRequest) Best() *UserScoresRequest {
	r.Type = ScoreTypeBest
	return r
}

func (r *UserScoresRequest) Firsts() *UserScoresRequest {
	r.Type = ScoreTypeFirsts
	return r
}

func (r *UserScoresRequest) Pinned() *UserScoresRequest {
	r.Type = ScoreTypePinned
	return r
}

func (r *UserScoresRequest) Recent() *UserScoresRequest {
	r.Type = ScoreTypeRecent
	return r
}

func (r *UserScoresRequest) SetIncludeFails(includeFails bool) *UserScoresRequest {
	r.IncludeFails = &includeFails
	return r
}

func (r *UserScoresRequest) SetLimit(limit int) *UserScoresRequest {
	r.Limit = &limit
	return r
}

func (r *UserScoresRequest) SetOffset(offset int) *UserScoresRequest {
	r.Offset = &offset
	return r
}

func (r *UserScoresRequest) Build() (*[]UserScore, error) {
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
		req.SetQueryParam("mode", r.Mode.String())
	}

	if r.Limit != nil {
		req.SetQueryParam("limit", strconv.Itoa(*r.Limit))
	}

	if r.Offset != nil {
		req.SetQueryParam("offset", strconv.Itoa(*r.Offset))
	}

	resp, err := req.Get("users/{user}/scores/{type}")
	if err != nil {
		return nil, err
	}

	return resp.Result().(*[]UserScore), nil
}

type UserBeatmapsetsRequest struct {
	client  *Client
	User    int
	MapType string
	Limit   *int
	Offset  *int
}

func (c *Client) GetUserBeatmapsets(user int) *UserBeatmapsetsRequest {
	return &UserBeatmapsetsRequest{client: c, User: user, MapType: "ranked"}
}

func (r *UserBeatmapsetsRequest) SetStatus(mapType RankStatus) *UserBeatmapsetsRequest {
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

func (r *UserBeatmapsetsRequest) Ranked() *UserBeatmapsetsRequest {
	r.MapType = "ranked"
	return r
}

func (r *UserBeatmapsetsRequest) Loved() *UserBeatmapsetsRequest {
	r.MapType = "loved"
	return r
}

func (r *UserBeatmapsetsRequest) Pending() *UserBeatmapsetsRequest {
	r.MapType = "pending"
	return r
}

func (r *UserBeatmapsetsRequest) Graveyard() *UserBeatmapsetsRequest {
	r.MapType = "graveyard"
	return r
}

func (r *UserBeatmapsetsRequest) SetLimit(limit int) *UserBeatmapsetsRequest {
	r.Limit = &limit
	return r
}

func (r *UserBeatmapsetsRequest) SetOffset(offset int) *UserBeatmapsetsRequest {
	r.Offset = &offset
	return r
}

func (r *UserBeatmapsetsRequest) Build() (*[]UserBeatmapset, error) {
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

	resp, err := req.Get("users/{user}/beatmapsets/{type}")
	if err != nil {
		return nil, err
	}

	return resp.Result().(*[]UserBeatmapset), nil

}

type UserMostPlayedRequest struct {
	client *Client
	User   int
	Limit  *int
	Offset *int
}

func (c *Client) GetUserMostPlayed(user int) *UserMostPlayedRequest {
	return &UserMostPlayedRequest{client: c, User: user}
}

func (r *UserMostPlayedRequest) SetLimit(limit int) *UserMostPlayedRequest {
	r.Limit = &limit
	return r
}

func (r *UserMostPlayedRequest) SetOffset(offset int) *UserMostPlayedRequest {
	r.Offset = &offset
	return r
}

func (r *UserMostPlayedRequest) Build() (*[]GetUserMostPlayedResponse, error) {
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

	resp, err := req.Get("users/{user}/beatmapsets/{type}")
	if err != nil {
		return nil, err
	}

	return resp.Result().(*[]GetUserMostPlayedResponse), nil
}

type UserRecentActivityRequest struct {
	client *Client
	User   int
	Limit  *int
	Offset *int
}

func (c *Client) GetUserRecentActivity(user int) *UserRecentActivityRequest {
	return &UserRecentActivityRequest{client: c, User: user}
}

func (r *UserRecentActivityRequest) SetLimit(limit int) *UserRecentActivityRequest {
	r.Limit = &limit
	return r
}

func (r *UserRecentActivityRequest) SetOffset(offset int) *UserRecentActivityRequest {
	r.Offset = &offset
	return r
}

func (r *UserRecentActivityRequest) Build() (*[]EventBase, error) {
	req := r.client.httpClient.R().SetResult(&[]map[string]interface{}{})

	req.SetPathParam("user", strconv.Itoa(r.User))

	if r.Limit != nil {
		req.SetQueryParam("limit", strconv.Itoa(*r.Limit))
	}

	if r.Offset != nil {
		req.SetQueryParam("offset", strconv.Itoa(*r.Offset))
	}

	resp, err := req.Get("users/{user}/recent_activity")
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

type UserRequest struct {
	client *Client
	User   string
	Mode   *Ruleset
}

func (c *Client) GetUser(user string) *UserRequest {
	return &UserRequest{client: c, User: user}
}

func (r *UserRequest) SetMode(mode Ruleset) *UserRequest {
	r.Mode = &mode
	return r
}

func (r *UserRequest) Build() (*UserExtended, error) {
	req := r.client.httpClient.R().SetResult(&UserExtended{}).SetPathParam("user", r.User)

	var url strings.Builder
	url.WriteString("/users/{user}")

	if r.Mode != nil {
		url.WriteString("/" + r.Mode.String())
	}

	resp, err := req.Get(url.String())
	if err != nil {
		return nil, err
	}

	return resp.Result().(*UserExtended), nil
}

type UsersRequest struct {
	client *Client
	Users  []int
}

func (c *Client) GetUsers(userIds []int) *UsersRequest {
	return &UsersRequest{client: c, Users: userIds}
}

func (r *UsersRequest) Build() (*GetUsersResponse, error) {
	req := r.client.httpClient.R().SetResult(&GetUsersResponse{})

	for _, id := range r.Users {
		req.QueryParam.Add("ids[]", strconv.Itoa(id))
	}

	resp, err := req.Get("users")
	if err != nil {
		return nil, err
	}

	return resp.Result().(*GetUsersResponse), nil
}
