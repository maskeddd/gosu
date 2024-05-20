package gosu

import (
	"strconv"
	"strings"
	"time"
)

type Genre int

const (
	GenreAny Genre = iota
	GenreUnspecified
	GenreVideoGame
	GenreAnime
	GenreRock
	GenrePop
	GenreOther
	GenreNovelty
	GenreHipHop Genre = iota + 1
	GenreElectronic
	GenreMetal
	GenreClassical
	GenreFolk
	GenreJazz
)

type Language int

const (
	LanguageAny Language = iota
	LanguageOther
	LanguageEnglish
	LanguageJapanese
	LanguageChinese
	LanguageInstrumental
	LanguageKorean
	LanguageFrench
	LanguageGerman
	LanguageSwedish
	LanguageSpanish
	LanguageItalian
	LanguageRussian
	LanguagePolish
	LanguageUnspecified
)

type BeatmapsetSearchSort string

const (
	SortArtist     BeatmapsetSearchSort = "artist"
	SortFavourites                      = "favourites"
	SortPlaycount                       = "plays"
	SortRankedDate                      = "ranked"
	SortRating                          = "rating"
	SortRelevance                       = "relevance"
	SortStars                           = "stars"
	SortTitle                           = "title"
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
	Ranked             int                    `json:"ranked"`
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

type BeatmapsetSearchResponse struct {
	Cursor      *Cursor `json:"cursor"`
	Beatmapsets []struct {
		Beatmapset
		Beatmaps []struct {
			Beatmap
			Checksum string `json:"checksum"`
			MaxCombo int    `json:"max_combo"`
		} `json:"beatmaps"`
		PackTags []string `json:"pack_tags"`
	} `json:"beatmapsets"`
	Total int `json:"total"`
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

type BeatmapsetWithIDRequest struct {
	client       *Client
	BeatmapsetID int
}

// GetBeatmapsetWithID returns a beatmapset, using its ID.
func (c *Client) GetBeatmapsetWithID(beatmapsetID int) *BeatmapsetWithIDRequest {
	return &BeatmapsetWithIDRequest{client: c, BeatmapsetID: beatmapsetID}
}

func (r *BeatmapsetWithIDRequest) Build() (*LookupBeatmapsetResponse, error) {
	req := r.client.httpClient.R().SetResult(&LookupBeatmapsetResponse{})
	req.SetQueryParam("beatmap_id", strconv.Itoa(r.BeatmapsetID))

	resp, err := req.Get("beatmapsets/lookup")
	if err != nil {
		return nil, err
	}
	return resp.Result().(*LookupBeatmapsetResponse), nil
}

type searchRankStatus struct {
	Any      bool
	Specific *RankStatus
}

type BeatmapsetWithSearchRequest struct {
	client     *Client
	Query      *string
	Mode       *Ruleset
	status     *searchRankStatus
	Genre      *Genre
	Language   *Language
	Video      bool
	Storyboard bool
	NSFW       bool
	Sort       *BeatmapsetSearchSort
	Descending bool
	Cursor     *Cursor
}

// GetBeatmapsetWithSearch returns a beatmapset, using a search query.
func (c *Client) GetBeatmapsetWithSearch() *BeatmapsetWithSearchRequest {
	return &BeatmapsetWithSearchRequest{client: c}
}

func (r *BeatmapsetWithSearchRequest) SetQuery(query string) *BeatmapsetWithSearchRequest {
	r.Query = &query
	return r
}

func (r *BeatmapsetWithSearchRequest) SetMode(mode Ruleset) *BeatmapsetWithSearchRequest {
	r.Mode = &mode
	return r
}

func (r *BeatmapsetWithSearchRequest) AnyStatus() *BeatmapsetWithSearchRequest {
	r.status = &searchRankStatus{Any: true}
	return r
}

func (r *BeatmapsetWithSearchRequest) SetStatus(status RankStatus) *BeatmapsetWithSearchRequest {
	if status == RankStatusRanked {
		temp := RankStatusPending
		r.status = &searchRankStatus{Any: false, Specific: &temp}
	} else {
		r.status = &searchRankStatus{Any: false, Specific: &status}
	}
	return r
}

func (r *BeatmapsetWithSearchRequest) SetGenre(genre Genre) *BeatmapsetWithSearchRequest {
	r.Genre = &genre
	return r
}

func (r *BeatmapsetWithSearchRequest) SetLanguage(language Language) *BeatmapsetWithSearchRequest {
	r.Language = &language
	return r
}

func (r *BeatmapsetWithSearchRequest) SetVideo(video bool) *BeatmapsetWithSearchRequest {
	r.Video = video
	return r
}

func (r *BeatmapsetWithSearchRequest) SetStoryboard(storyboard bool) *BeatmapsetWithSearchRequest {
	r.Storyboard = storyboard
	return r
}

func (r *BeatmapsetWithSearchRequest) SetNSFW(nsfw bool) *BeatmapsetWithSearchRequest {
	r.NSFW = nsfw
	return r
}

func (r *BeatmapsetWithSearchRequest) SortBy(sort BeatmapsetSearchSort, descending bool) *BeatmapsetWithSearchRequest {
	r.Sort = &sort
	r.Descending = descending
	return r
}

func (r *BeatmapsetWithSearchRequest) SetCursor(cursor *Cursor) *BeatmapsetWithSearchRequest {
	r.Cursor = cursor
	return r
}

func (r *BeatmapsetWithSearchRequest) Build() (*BeatmapsetSearchResponse, error) {
	req := r.client.httpClient.R().SetResult(&BeatmapsetSearchResponse{})

	if r.Query != nil {
		req.SetQueryParam("q", *r.Query)
	}

	if r.Mode != nil {
		req.SetQueryParam("m", r.Mode.String())
	}

	if r.status != nil {
		if r.status.Any {
			req.SetQueryParam("s", "any")
		} else if r.status.Specific != nil {
			req.SetQueryParam("s", strings.ToLower(r.status.Specific.String()))
		}
	}

	if r.Genre != nil {
		req.SetQueryParam("g", strconv.Itoa(int(*r.Genre)))
	}

	if r.Language != nil {
		req.SetQueryParam("l", strconv.Itoa(int(*r.Language)))
	}

	var extra string
	switch {
	case !r.Video && r.Storyboard:
		extra = "storyboard"
	case r.Video && !r.Storyboard:
		extra = "video"
	case r.Video && r.Storyboard:
		extra = "storyboard.video"
	default:
		extra = ""
	}

	if extra != "" {
		req.SetQueryParam("e", extra)
	}

	req.SetQueryParam("nsfw", strconv.FormatBool(r.NSFW))

	if r.Sort != nil {
		var sort strings.Builder
		sort.WriteString(string(*r.Sort))

		if r.Descending {
			sort.WriteString("_desc")
		} else {
			sort.WriteString("_asc")
		}

		req.SetQueryParam("sort", sort.String())
	}

	resp, err := req.Get("beatmapsets/search")
	if err != nil {
		return nil, err
	}

	return resp.Result().(*BeatmapsetSearchResponse), nil
}
