package gosu

import (
	"strconv"
	"time"
)

type ChangelogStream string

const (
	ChangelogStreamStable40    = "stable40"
	ChangelogStreamBeta40      = "beta40"
	ChangelogStreamCuttingEdge = "cuttingedge"
	ChangelogStreamLazer       = "lazer"
	ChangelogStreamWeb         = "web"
)

type MessageFormat string

const (
	MessageFormatHTML     = "html"
	MessageFormatMarkdown = "markdown"
)

type UpdateStream struct {
	DisplayName *string `json:"display_name"`
	ID          int     `json:"id"`
	IsFeatured  bool    `json:"is_featured"`
	Name        string  `json:"name"`
}

type Build struct {
	CreatedAt      time.Time     `json:"created_at"`
	DisplayVersion string        `json:"display_version"`
	ID             int           `json:"id"`
	UpdateStream   *UpdateStream `json:"update_stream"`
	Users          int           `json:"users"`
	Version        *string       `json:"version"`
	YoutubeID      *string       `json:"youtube_id"`
}

type BuildVersions struct {
	Next     *Build `json:"next"`
	Previous *Build `json:"previous"`
}

type ChangelogEntry struct {
	Category            string     `json:"category"`
	CreatedAt           *time.Time `json:"created_at"`
	GithubPullRequestID *int       `json:"github_pull_request_id"`
	GithubURL           *string    `json:"github_url"`
	ID                  *int       `json:"id"`
	Major               bool       `json:"major"`
	Repository          *string    `json:"repository"`
	Title               *string    `json:"title"`
	Type                string     `json:"type"`
	URL                 *string    `json:"url"`
}

type GithubUser struct {
	DisplayName    string  `json:"display_name"`
	GithubURL      *string `json:"github_url"`
	GithubUsername *string `json:"github_username"`
	ID             *int    `json:"id"`
	OsuUsername    *string `json:"osu_username"`
	UserID         *int    `json:"user_id"`
	UserURL        *string `json:"user_url"`
}

type ChangelogBuildResponse struct {
	Build
	ChangelogEntries []struct {
		ChangelogEntry
		GithubUser  GithubUser `json:"github_user"`
		Message     *string    `json:"message"`
		MessageHTML *string    `json:"message_html"`
	} `json:"changelog_entries"`
	Versions BuildVersions `json:"versions"`
}

type ChangelogListingResponse struct {
	Builds []struct {
		ChangelogEntries []struct {
			ChangelogEntry
			GithubUser  GithubUser `json:"github_user"`
			Message     *string    `json:"message"`
			MessageHTML *string    `json:"message_html"`
		} `json:"changelog_entries"`
		CreatedAt      time.Time    `json:"created_at"`
		DisplayVersion string       `json:"display_version"`
		ID             int          `json:"id"`
		UpdateStream   UpdateStream `json:"update_stream"`
		Users          int          `json:"users"`
		Version        string       `json:"version"`
		YoutubeID      *string      `json:"youtube_id"`
	} `json:"builds"`
	Search struct {
		From   *string `json:"from"`
		Limit  int     `json:"limit"`
		MaxID  *int    `json:"max_id"`
		Stream *string `json:"stream"`
		To     *string `json:"to"`
	} `json:"search"`
	Streams []struct {
		UpdateStream
		LatestBuild Build `json:"latest_build"`
		UserCount   int   `json:"user_count"`
	} `json:"streams"`
}

type LookupChangelogBuildResponse struct {
	Build
	ChangelogEntries []struct {
		ChangelogEntry
		GithubUser  GithubUser `json:"github_user"`
		Message     *string    `json:"message"`
		MessageHTML *string    `json:"message_html"`
	} `json:"changelog_entries"`
	Versions BuildVersions `json:"versions"`
}

type ChangelogBuildRequest struct {
	client       *Client
	Stream       string
	BuildVersion string
}

// GetChangelogBuild returns details of the specified build.
func (c *Client) GetChangelogBuild(stream string, build string) *ChangelogBuildRequest {
	return &ChangelogBuildRequest{client: c, Stream: stream, BuildVersion: build}
}

func (r *ChangelogBuildRequest) Build() (*ChangelogBuildResponse, error) {
	resp, err := r.client.httpClient.R().SetResult(&ChangelogBuildResponse{}).SetPathParams(map[string]string{
		"stream": r.Stream,
		"build":  r.BuildVersion,
	}).Get("changelog/{stream}/{build}")
	if err != nil {
		return nil, err
	}

	return resp.Result().(*ChangelogBuildResponse), nil
}

type ChangelogListingRequest struct {
	client         *Client
	From           *string
	MaxID          *int
	Stream         *ChangelogStream
	To             *string
	MessageFormats []MessageFormat
}

// GetChangelogListing returns a listing of update streams, builds, and changelog entries.
func (c *Client) GetChangelogListing() *ChangelogListingRequest {
	return &ChangelogListingRequest{client: c}
}

func (r *ChangelogListingRequest) SetFrom(from string) *ChangelogListingRequest {
	r.From = &from
	return r
}

func (r *ChangelogListingRequest) SetMaxID(maxID int) *ChangelogListingRequest {
	r.MaxID = &maxID
	return r
}

func (r *ChangelogListingRequest) SetStream(stream ChangelogStream) *ChangelogListingRequest {
	r.Stream = &stream
	return r
}

func (r *ChangelogListingRequest) SetTo(to string) *ChangelogListingRequest {
	r.To = &to
	return r
}

func (r *ChangelogListingRequest) SetMessageFormats(messageFormats []MessageFormat) *ChangelogListingRequest {
	r.MessageFormats = messageFormats
	return r
}

func (r *ChangelogListingRequest) Build() (*ChangelogListingResponse, error) {
	req := r.client.httpClient.R().SetResult(&ChangelogListingResponse{})

	if r.From != nil {
		req.SetQueryParam("from", *r.From)
	}

	if r.MaxID != nil {
		req.SetQueryParam("max_id", strconv.Itoa(*r.MaxID))
	}

	if r.Stream != nil {
		req.SetQueryParam("stream", string(*r.Stream))
	}

	if r.To != nil {
		req.SetQueryParam("to", *r.To)
	}

	if r.MessageFormats != nil {
		for _, messageFormat := range r.MessageFormats {
			req.QueryParam.Add("message_formats[]", string(messageFormat))
		}
	}

	resp, err := req.Get("changelog")
	if err != nil {
		return nil, err
	}

	return resp.Result().(*ChangelogListingResponse), nil
}

type LookupChangelogBuildRequest struct {
	client         *Client
	Changelog      string
	Key            *string
	MessageFormats []MessageFormat
}

// LookupChangelogBuild returns details of the specified build.
func (c *Client) LookupChangelogBuild(changelog string) *LookupChangelogBuildRequest {
	return &LookupChangelogBuildRequest{client: c, Changelog: changelog}
}

func (r *LookupChangelogBuildRequest) SetKey(key string) *LookupChangelogBuildRequest {
	r.Key = &key
	return r
}

func (r *LookupChangelogBuildRequest) SetMessageFormats(messageFormats []MessageFormat) *LookupChangelogBuildRequest {
	r.MessageFormats = messageFormats
	return r
}

func (r *LookupChangelogBuildRequest) Build() (*LookupChangelogBuildResponse, error) {
	req := r.client.httpClient.R().SetResult(&LookupChangelogBuildResponse{}).SetPathParam("changelog", r.Changelog)

	if r.Key != nil {
		req.SetQueryParam("key", *r.Key)
	}

	if r.MessageFormats != nil {
		for _, messageFormat := range r.MessageFormats {
			req.QueryParam.Add("message_formats[]", string(messageFormat))
		}
	}

	resp, err := req.Get("changelog/{changelog}")
	if err != nil {
		return nil, err
	}

	return resp.Result().(*LookupChangelogBuildResponse), nil
}
