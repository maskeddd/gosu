package gosu

import (
	"strconv"
	"time"
)

type NewsPost struct {
	Author      string    `json:"author"`
	EditUrl     string    `json:"edit_url"`
	FirstImage  *string   `json:"first_image"`
	ID          int       `json:"id"`
	PublishedAt time.Time `json:"published_at"`
	Slug        string    `json:"slug"`
	Title       string    `json:"title"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type NewsSidebar struct {
	CurrentYear int        `json:"current_year"`
	NewsPost    []NewsPost `json:"news_posts"`
	Years       []int      `json:"years"`
}

type NewsSearch struct {
	Limit int    `json:"limit"`
	Sort  string `json:"sort"`
	Year  int    `json:"year"`
}

type NewsListingResponse struct {
	CursorString Cursor `json:"cursor_string"`
	NewsPosts    []struct {
		NewsPost
		Preview string `json:"preview"`
	} `json:"news_posts"`
	NewsSidebar NewsSidebar `json:"news_sidebar"`
	Search      NewsSearch  `json:"search"`
}

type NewsNavigation struct {
	Newer *NewsPost `json:"newer,omitempty"`
	Older *NewsPost `json:"older,omitempty"`
}

type NewsPostResponse struct {
	NewsPost
	Content    string         `json:"content"`
	Navigation NewsNavigation `json:"navigation"`
}

type NewsListingRequest struct {
	client *Client
	Limit  *int
	Year   *int
}

// GetNewsListing returns a list of news posts and related metadata.
func (c *Client) GetNewsListing() *NewsListingRequest {
	return &NewsListingRequest{client: c}
}

func (r *NewsListingRequest) SetLimit(limit int) *NewsListingRequest {
	r.Limit = &limit
	return r
}

func (r *NewsListingRequest) SetYear(year int) *NewsListingRequest {
	r.Year = &year
	return r
}

func (r *NewsListingRequest) Build() (*NewsListingResponse, error) {
	req := r.client.httpClient.R().SetResult(&NewsListingResponse{})

	if r.Limit != nil {
		req.SetQueryParam("limit", strconv.Itoa(*r.Limit))
	}

	if r.Year != nil {
		req.SetQueryParam("year", strconv.Itoa(*r.Year))
	}

	resp, err := req.Get("news")
	if err != nil {
		return nil, err
	}

	return resp.Result().(*NewsListingResponse), nil
}

type NewsPostRequest struct {
	client *Client
	News   string
	Key    *string
}

// GetNewsPost returns details of the specified news post.
func (c *Client) GetNewsPost(news string) *NewsPostRequest {
	return &NewsPostRequest{client: c, News: news}
}

func (r *NewsPostRequest) SetKey(key string) *NewsPostRequest {
	r.Key = &key
	return r
}

func (r *NewsPostRequest) Build() (*NewsPostResponse, error) {
	req := r.client.httpClient.R().SetResult(&NewsPostResponse{})

	req.SetPathParam("news", r.News)

	if r.Key != nil {
		req.SetQueryParam("key", *r.Key)
	}

	resp, err := req.Get("news/{news}")
	if err != nil {
		return nil, err
	}

	return resp.Result().(*NewsPostResponse), nil
}
