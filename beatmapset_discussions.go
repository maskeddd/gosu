package gosu

import (
	"strconv"
	"time"
)

type DiscussionMessageType string

const (
	DiscussionMessageTypeSuggestion = "suggestion"
	DiscussionMessageTypeProblem    = "problem"
	DiscussionMessageTypeMapperNote = "mapper_note"
	DiscussionMessageTypePraise     = "praise"
	DiscussionMessageTypeHype       = "hype"
	DiscussionMessageTypeReview     = "review"
)

type DiscussionSort string

const (
	DiscussionSortDescending = "id_desc"
	DiscussionSortAscending  = "id_asc"
)

type BeatmapsetDiscussion struct {
	BeatmapID      *int                  `json:"beatmap_id"`
	BeatmapsetID   *int                  `json:"beatmapset_id"`
	CanBeResolved  bool                  `json:"can_be_resolved"`
	CanGrantKudosu bool                  `json:"can_grant_kudosu"`
	CreatedAt      time.Time             `json:"created_at"`
	DeletedAt      *time.Time            `json:"deleted_at"`
	DeletedByID    *int                  `json:"deleted_by_id"`
	ID             int                   `json:"id"`
	KudosuDenied   bool                  `json:"kudosu_denied"`
	LastPostAt     time.Time             `json:"last_post_at"`
	MessageType    DiscussionMessageType `json:"message_type"`
	ParentID       *int                  `json:"parent_id"`
	Resolved       bool                  `json:"resolved"`
	Timestamp      *int                  `json:"timestamp"`
	UpdatedAt      time.Time             `json:"updated_at"`
	UserID         *int                  `json:"user_id"`
}

type BeatmapsetDiscussionPost struct {
	BeatmapsetDiscussionID int        `json:"beatmapset_discussion_id"`
	CreatedAt              time.Time  `json:"created_at"`
	DeletedAt              *time.Time `json:"deleted_at"`
	DeletedByID            *int       `json:"deleted_by_id"`
	ID                     int        `json:"id"`
	LastEditorID           *int       `json:"last_editor_id"`
	Message                string     `json:"message"`
	System                 bool       `json:"system"`
	UpdatedAt              time.Time  `json:"updated_at"`
	UserID                 int        `json:"user_id"`
}

type BeatmapsetDiscussionVote struct {
	BeatmapsetDiscussionID int       `json:"beatmapset_discussion_id"`
	CreatedAt              time.Time `json:"created_at"`
	ID                     int       `json:"id"`
	Score                  int       `json:"score"`
	UpdatedAt              time.Time `json:"updated_at"`
	UserID                 int       `json:"user_id"`
}

type BeatmapsetDiscussionBaseRequest struct {
	client *Client
	Limit  *int
	Page   *int
	Sort   *DiscussionSort
}

func (r *BeatmapsetDiscussionBaseRequest) SetLimit(limit int) *BeatmapsetDiscussionBaseRequest {
	r.Limit = &limit
	return r
}

func (r *BeatmapsetDiscussionBaseRequest) SetPage(page int) *BeatmapsetDiscussionBaseRequest {
	r.Page = &page
	return r
}

func (r *BeatmapsetDiscussionBaseRequest) SetSort(sort DiscussionSort) *BeatmapsetDiscussionBaseRequest {
	r.Sort = &sort
	return r
}

type BeatmapsetDiscussionPostsRequest struct {
	BeatmapsetDiscussionBaseRequest
	BeatmapsetDiscussionID *int
	Types                  *[]string
	User                   *int
}

func (c *Client) GetBeatmapsetDiscussionPosts() *BeatmapsetDiscussionPostsRequest {
	return &BeatmapsetDiscussionPostsRequest{
		BeatmapsetDiscussionBaseRequest: BeatmapsetDiscussionBaseRequest{client: c},
	}
}

func (r *BeatmapsetDiscussionPostsRequest) SetDiscussionID(discussionID int) *BeatmapsetDiscussionPostsRequest {
	r.BeatmapsetDiscussionID = &discussionID
	return r
}

func (r *BeatmapsetDiscussionPostsRequest) SetTypes(types []string) *BeatmapsetDiscussionPostsRequest {
	r.Types = &types
	return r
}

func (r *BeatmapsetDiscussionPostsRequest) SetUser(userID int) *BeatmapsetDiscussionPostsRequest {
	r.User = &userID
	return r
}

func (r *BeatmapsetDiscussionPostsRequest) Build() (*BeatmapsetDiscussionPost, error) {
	req := r.client.httpClient.R().SetResult(&BeatmapsetDiscussionPost{})

	if r.BeatmapsetDiscussionID != nil {
		req.SetQueryParam("beatmapset_discussion_id", strconv.Itoa(*r.BeatmapsetDiscussionID))
	}

	if r.Limit != nil {
		req.SetQueryParam("limit", strconv.Itoa(*r.Limit))
	}

	if r.Page != nil {
		req.SetQueryParam("page", strconv.Itoa(*r.Page))
	}

	if r.Sort != nil {
		req.SetQueryParam("sort", string(*r.Sort))
	}

	if r.User != nil {
		req.SetQueryParam("user", strconv.Itoa(*r.User))
	}

	resp, err := req.Get("beatmapsets/discussions/posts")
	if err != nil {
		return nil, err
	}

	return resp.Result().(*BeatmapsetDiscussionPost), nil
}
