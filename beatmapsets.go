package gosu

import (
	"strconv"
	"time"
)

type RankStatus int

const (
	Graveyard RankStatus = iota - 2
	WIP
	Pending
	Ranked
	Approved
	Qualified
	Loved
)

type BeatmapsetCompact struct {
	Artist         string          `json:"artist"`
	ArtistUnicode  *string         `json:"artist_unicode"`
	Covers         BeatmapsetCover `json:"covers"`
	Creator        string          `json:"creator"`
	FavouriteCount int             `json:"favourite_count"`
	Hype           *BeatmapsetHype `json:"hype"`
	ID             int             `json:"id"`
	NSFW           bool            `json:"nsfw"`
	Offset         int             `json:"offset"`
	PlayCount      int             `json:"play_count"`
	PreviewURL     string          `json:"preview_url"`
	Source         string          `json:"source"`
	Status         string          `json:"status"`
	Spotlight      bool            `json:"spotlight"`
	Title          string          `json:"title"`
	TitleUnicode   string          `json:"title_unicode"`
	UserID         int             `json:"user_id"`
	Video          bool            `json:"video"`
}

type Beatmapset struct {
	BeatmapsetCompact
	Availability       BeatmapsetAvailability `json:"availability"`
	BPM                float32                `json:"bpm"`
	CanBeHyped         bool                   `json:"can_be_hyped"`
	DeletedAt          *time.Time             `json:"deleted_at"`
	DiscussionLocked   bool                   `json:"discussion_locked"`
	Hype               *BeatmapsetHype        `json:"hype"`
	IsScoreable        bool                   `json:"is_scoreable"`
	LastUpdated        time.Time              `json:"last_updated"`
	LegacyThreadURL    *string                `json:"legacy_thread_url"`
	NominationsSummary BeatmapsetNominations  `json:"nominations_summary"`
	Ranked             RankStatus             `json:"ranked"`
	RankedDate         *time.Time             `json:"ranked_date"`
	Storyboard         bool                   `json:"storyboard"`
	SubmittedDate      *time.Time             `json:"submitted_date"`
	Tags               string                 `json:"tags"`
}

type BeatmapsetCover struct {
	Cover       string `json:"cover"`
	Cover2X     string `json:"cover@2x"`
	Card        string `json:"card"`
	Card2X      string `json:"card@2x"`
	List        string `json:"list"`
	List2X      string `json:"list@2x"`
	Slimcover   string `json:"slimcover"`
	Slimcover2X string `json:"slimcover@2x"`
}

type BeatmapsetHype struct {
	Current  int `json:"current"`
	Required int `json:"required"`
}

type BeatmapsetNominations struct {
	Current  int `json:"current"`
	Required int `json:"required"`
}

type Nomination struct {
	BeatmapsetID int       `json:"beatmapset_id"`
	Rulesets     []Ruleset `json:"rulesets"`
	Reset        bool      `json:"reset"`
	UserID       int       `json:"user_id"`
}

type BeatmapsetAvailability struct {
	DownloadDisabled bool    `json:"download_disabled"`
	MoreInformation  *string `json:"more_information"`
}
type BeatmapsetGenre struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type BeatmapsetLanguage struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type LookupBeatmapsetResponse struct {
	Beatmapset
	Beatmaps []struct {
		Beatmap
		Checksum  string `json:"checksum"`
		FailTimes `json:"failtimes"`
		MaxCombo  int `json:"max_combo"`
	} `json:"beatmaps"`
	Converts []struct {
		Beatmap
		Checksum  string `json:"checksum"`
		FailTimes `json:"failtimes"`
	} `json:"converts"`
	CurrentNominations []Nomination `json:"current_nominations"`
	Description        struct {
		Description string `json:"description"`
	} `json:"description"`
	BeatmapsetGenre    `json:"genre"`
	BeatmapsetLanguage `json:"language"`
	PackTags           []string      `json:"pack_tags"`
	Ratings            []int         `json:"ratings"`
	RecentFavourites   []UserCompact `json:"recent_favourites"`
	RelatedUsers       []UserCompact `json:"related_users"`
	User               UserCompact   `json:"user"`
}

type LookupBeatmapsetRequest struct {
	client       *Client
	BeatmapsetID int
}

func (c *Client) GetBeatmapsetFromID(beatmapsetID int) *LookupBeatmapsetRequest {
	return &LookupBeatmapsetRequest{client: c, BeatmapsetID: beatmapsetID}
}

func (r *LookupBeatmapsetRequest) Build() (*LookupBeatmapsetResponse, error) {
	req := r.client.httpClient.R().SetResult(&LookupBeatmapsetResponse{})
	req.SetQueryParam("beatmap_id", strconv.Itoa(r.BeatmapsetID))

	resp, err := req.Get("/beatmapsets/lookup")
	if err != nil {
		return nil, err
	}
	return resp.Result().(*LookupBeatmapsetResponse), nil
}

//type SearchRankStatus struct {
//	Any      bool
//	Specific RankStatus
//}
//
//type BeatmapsetSearchRequest struct {
//	client     *Client
//	Query      *string
//	Mode       *int
//	Status     *SearchRankStatus
//	Genre      *int
//	Language   *int
//	Video      bool
//	Storyboard bool
//	NSFW       bool
//	Descending bool
//	Cursor     *int
//}
