package gosu

import (
	"strconv"
	"time"
)

type SpotlightType string

const (
	SpotlightTypeMonthly   SpotlightType = "monthly"
	SpotlightTypeSpotlight SpotlightType = "spotlight"
	SpotlightTypeTheme     SpotlightType = "theme"
	SpotlightTypeSpecial   SpotlightType = "special"
	SpotlightTypeBestOf    SpotlightType = "bestof"
)

type RankingType string

const (
	RankingTypeCharts      RankingType = "charts"
	RankingTypeCountry     RankingType = "country"
	RankingTypePerformance RankingType = "performance"
	RankingTypeScore       RankingType = "score"
)

type RankingFilter string

const (
	RankingFilterAll     RankingFilter = "all"
	RankingFilterFriends RankingFilter = "friends"
)

type RankingVariant string

const (
	RankingVariant4k = "4k"
	RankingVariant7k = "7k"
)

type Spotlight struct {
	EndDate          time.Time     `json:"end_date"`
	ID               int           `json:"id"`
	ModeSpecific     bool          `json:"mode_specific"`
	ParticipantCount *int          `json:"participant_count,omitempty"`
	Name             string        `json:"name"`
	StartDate        time.Time     `json:"start_date"`
	Type             SpotlightType `json:"type"`
}

type KudosuRankingResponse struct {
	Ranking []struct {
		UserCompact
		Kudosu UserKudosu `json:"kudosu"`
	} `json:"ranking"`
}

type KudosuRankingRequest struct {
	client *Client
	Page   *int `json:"page"`
}

type Rankings struct {
	Beatmapsets *[]Beatmapset `json:"beatmapsets,omitempty"`
	Cursor      Cursor        `json:"cursor"`
	Ranking     []struct {
		UserStatistics
		User struct {
			UserCompact
			Country Country `json:"country"`
			Cover   Cover   `json:"cover"`
		} `json:"user"`
	} `json:"ranking"`
	Spotlight *Spotlight `json:"spotlight,omitempty"`
	Total     int        `json:"total"`
}

type Spotlights struct {
	Spotlights []Spotlight `json:"spotlights"`
}

// GetKudosuRanking returns the kudosu ranking.
func (c *Client) GetKudosuRanking() *KudosuRankingRequest {
	return &KudosuRankingRequest{
		client: c,
	}
}

func (r *KudosuRankingRequest) SetPage(page int) *KudosuRankingRequest {
	r.Page = &page
	return r
}

func (r *KudosuRankingRequest) Build() (*KudosuRankingResponse, error) {
	req := r.client.httpClient.R().SetResult(&KudosuRankingResponse{})

	if r.Page != nil {
		req.SetQueryParam("page", strconv.Itoa(*r.Page))
	}

	resp, err := req.Get("rankings/kudosu")
	if err != nil {
		return nil, err
	}

	return resp.Result().(*KudosuRankingResponse), nil
}

type RankingRequest struct {
	client    *Client
	Mode      Ruleset
	Type      RankingType
	Country   *string
	Filter    *RankingFilter
	Spotlight *int
	Variant   *RankingVariant
}

// GetRanking returns the current ranking for the specified type and game mode.
func (c *Client) GetRanking(mode Ruleset, rankingType RankingType) *RankingRequest {
	return &RankingRequest{
		client: c,
		Mode:   mode,
		Type:   rankingType,
	}
}

func (r *RankingRequest) SetCountry(country string) *RankingRequest {
	r.Country = &country
	return r
}

func (r *RankingRequest) SetFilter(filter RankingFilter) *RankingRequest {
	r.Filter = &filter
	return r
}

func (r *RankingRequest) SetSpotlight(spotlight int) *RankingRequest {
	r.Spotlight = &spotlight
	return r
}

func (r *RankingRequest) SetVariant(variant RankingVariant) *RankingRequest {
	r.Variant = &variant
	return r
}

func (r *RankingRequest) Build() (*Rankings, error) {
	req := r.client.httpClient.R().SetResult(&Rankings{}).SetPathParams(map[string]string{
		"mode": r.Mode.String(),
		"type": string(r.Type),
	})

	if r.Country != nil {
		req.SetQueryParam("country", *r.Country)
	}

	if r.Filter != nil {
		req.SetQueryParam("filter", string(*r.Filter))
	}

	if r.Spotlight != nil && r.Type == RankingTypeCharts {
		req.SetQueryParam("spotlight", strconv.Itoa(*r.Spotlight))
	}

	if r.Variant != nil && r.Mode == RulesetMania {
		req.SetQueryParam("variant", string(*r.Variant))
	}

	resp, err := req.Get("rankings/{mode}/{type}")
	if err != nil {
		return nil, err
	}

	return resp.Result().(*Rankings), nil
}

type SpotlightsRequest struct {
	client *Client
}

// GetSpotlights returns the list of spotlights.
func (c *Client) GetSpotlights() *SpotlightsRequest {
	return &SpotlightsRequest{client: c}
}

func (r *SpotlightsRequest) Build() (*Spotlights, error) {
	resp, err := r.client.httpClient.R().SetResult(&Spotlights{}).Get("spotlights")
	if err != nil {
		return nil, err
	}

	return resp.Result().(*Spotlights), nil
}
