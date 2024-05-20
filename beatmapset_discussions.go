package gosu

import (
	"strconv"
	"time"
)

type MessageType string

const (
	MessageTypeSuggestion = "suggestion"
	MessageTypeProblem    = "problem"
	MessageTypeMapperNote = "mapper_note"
	MessageTypePraise     = "praise"
	MessageTypeHype       = "hype"
	MessageTypeReview     = "review"
)

type DiscussionPostType string

const (
	DiscussionPostTypeFirst  = "first"
	DiscussionPostTypeReply  = "reply"
	DiscussionPostTypeSystem = "system"
)

type DiscussionVoteDirection string

const (
	DiscussionVoteDirectionUp   = "1"
	DiscussionVoteDirectionDown = "-1"
)

type DiscussionStatus string

const (
	DiscussionStatusAll            DiscussionStatus = "all"
	DiscussionStatusRanked         DiscussionStatus = "ranked"
	DiscussionStatusQualified      DiscussionStatus = "qualified"
	DiscussionStatusDisqualified   DiscussionStatus = "disqualified"
	DiscussionStatusNeverQualified DiscussionStatus = "never_qualified"
)

type BeatmapsetDiscussion struct {
	BeatmapID      *int        `json:"beatmap_id"`
	BeatmapsetID   *int        `json:"beatmapset_id"`
	CanBeResolved  bool        `json:"can_be_resolved"`
	CanGrantKudosu bool        `json:"can_grant_kudosu"`
	CreatedAt      time.Time   `json:"created_at"`
	DeletedAt      *time.Time  `json:"deleted_at"`
	DeletedByID    *int        `json:"deleted_by_id"`
	ID             int         `json:"id"`
	KudosuDenied   bool        `json:"kudosu_denied"`
	LastPostAt     time.Time   `json:"last_post_at"`
	MessageType    MessageType `json:"message_type"`
	ParentID       *int        `json:"parent_id"`
	Resolved       bool        `json:"resolved"`
	Timestamp      *int        `json:"timestamp"`
	UpdatedAt      time.Time   `json:"updated_at"`
	UserID         *int        `json:"user_id"`
}

type DiscussionPost struct {
	BeatmapsetDiscussionID int         `json:"beatmapset_discussion_id"`
	CreatedAt              time.Time   `json:"created_at"`
	DeletedAt              *time.Time  `json:"deleted_at"`
	DeletedByID            *int        `json:"deleted_by_id"`
	ID                     int         `json:"id"`
	LastEditorID           *int        `json:"last_editor_id"`
	Message                interface{} `json:"message"`
	System                 bool        `json:"system"`
	UpdatedAt              time.Time   `json:"updated_at"`
	UserID                 int         `json:"user_id"`
}

type DiscussionVote struct {
	BeatmapsetDiscussionID int       `json:"beatmapset_discussion_id"`
	CreatedAt              time.Time `json:"created_at"`
	ID                     int       `json:"id"`
	Score                  int       `json:"score"`
	UpdatedAt              time.Time `json:"updated_at"`
	UserID                 int       `json:"user_id"`
}

type DiscussionPostsResponse struct {
	Beatmapsets  []BeatmapsetCompact    `json:"beatmapsets"`
	CursorString Cursor                 `json:"cursor_string"`
	Discussions  []BeatmapsetDiscussion `json:"discussions"`
	Posts        []DiscussionPost       `json:"posts"`
	Users        []UserCompact          `json:"users"`
}

type DiscussionVotesResponse struct {
	CursorString Cursor                 `json:"cursor_string"`
	Discussions  []BeatmapsetDiscussion `json:"discussions"`
	Users        []UserCompact          `json:"users"`
	Votes        []DiscussionVote       `json:"votes"`
}

type DiscussionsResponse struct {
	CursorString        Cursor                 `json:"cursor_string"`
	Users               []UserCompact          `json:"users"`
	Discussions         []BeatmapsetDiscussion `json:"discussions"`
	IncludedDiscussions []BeatmapsetDiscussion `json:"included_discussions"`
	Beatmapsets         []Beatmapset           `json:"beatmapsets"`
	Beatmaps            []struct {
		Beatmap
		Checksum *string `json:"checksum"`
	} `json:"beatmaps"`
	ReviewsConfig struct {
		MaxBlocks int `json:"max_blocks"`
	} `json:"reviews_config"`
}

type DiscussionBaseRequest struct {
	client *Client
	Limit  *int
	Page   *int
	Sort   *Sort
}

type DiscussionPostsRequest struct {
	DiscussionBaseRequest
	BeatmapsetDiscussionID *int
	Types                  []DiscussionPostType
	User                   *int
}

// GetDiscussionPosts returns the posts of beatmapset discussions.
func (c *Client) GetDiscussionPosts() *DiscussionPostsRequest {
	return &DiscussionPostsRequest{
		DiscussionBaseRequest: DiscussionBaseRequest{client: c},
	}
}

func (r *DiscussionPostsRequest) SetDiscussionID(discussionID int) *DiscussionPostsRequest {
	r.BeatmapsetDiscussionID = &discussionID
	return r
}

func (r *DiscussionPostsRequest) AddType(types DiscussionPostType) *DiscussionPostsRequest {
	r.Types = append(r.Types, types)
	return r
}

func (r *DiscussionPostsRequest) SetTypes(types []DiscussionPostType) *DiscussionPostsRequest {
	r.Types = types
	return r
}

func (r *DiscussionPostsRequest) SetUser(userID int) *DiscussionPostsRequest {
	r.User = &userID
	return r
}

func (r *DiscussionPostsRequest) SetLimit(limit int) *DiscussionPostsRequest {
	r.Limit = &limit
	return r
}

func (r *DiscussionPostsRequest) SetPage(page int) *DiscussionPostsRequest {
	r.Page = &page
	return r
}

func (r *DiscussionPostsRequest) SetSort(sort Sort) *DiscussionPostsRequest {
	r.Sort = &sort
	return r
}

func (r *DiscussionPostsRequest) Build() (*DiscussionPostsResponse, error) {
	req := r.client.httpClient.R().SetResult(&DiscussionPostsResponse{})

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

	if r.Types != nil {
		for _, postType := range r.Types {
			req.QueryParam.Add("types[]", string(postType))
		}
	}

	resp, err := req.Get("beatmapsets/discussions/posts")
	if err != nil {
		return nil, err
	}

	return resp.Result().(*DiscussionPostsResponse), nil
}

type DiscussionVotesRequest struct {
	DiscussionBaseRequest
	BeatmapsetDiscussionID *int
	Receiver               *int                     `json:"receiver"`
	Score                  *DiscussionVoteDirection `json:"score"`
	User                   *int                     `json:"user"`
}

// GetDiscussionVotes returns the votes given to beatmapset discussions.
func (c *Client) GetDiscussionVotes() *DiscussionVotesRequest {
	return &DiscussionVotesRequest{
		DiscussionBaseRequest: DiscussionBaseRequest{client: c},
	}
}

func (r *DiscussionVotesRequest) SetDiscussionID(discussionID int) *DiscussionVotesRequest {
	r.BeatmapsetDiscussionID = &discussionID
	return r
}

func (r *DiscussionVotesRequest) SetReceiver(receiverID int) *DiscussionVotesRequest {
	r.Receiver = &receiverID
	return r
}

func (r *DiscussionVotesRequest) SetScore(score DiscussionVoteDirection) *DiscussionVotesRequest {
	r.Score = &score
	return r
}

func (r *DiscussionVotesRequest) SetUser(userID int) *DiscussionVotesRequest {
	r.User = &userID
	return r
}

func (r *DiscussionVotesRequest) SetLimit(limit int) *DiscussionVotesRequest {
	r.Limit = &limit
	return r
}

func (r *DiscussionVotesRequest) SetPage(page int) *DiscussionVotesRequest {
	r.Page = &page
	return r
}

func (r *DiscussionVotesRequest) SetSort(sort Sort) *DiscussionVotesRequest {
	r.Sort = &sort
	return r
}

func (r *DiscussionVotesRequest) Build() (*DiscussionVotesResponse, error) {
	req := r.client.httpClient.R().SetResult(&DiscussionVotesResponse{})

	if r.BeatmapsetDiscussionID != nil {
		req.SetQueryParam("beatmapset_discussion_id", strconv.Itoa(*r.BeatmapsetDiscussionID))
	}

	if r.Receiver != nil {
		req.SetQueryParam("receiver", strconv.Itoa(*r.Receiver))
	}

	if r.Score != nil {
		req.SetQueryParam("score", string(*r.Score))
	}

	if r.User != nil {
		req.SetQueryParam("user", strconv.Itoa(*r.User))
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

	resp, err := req.Get("beatmapsets/discussions/votes")
	if err != nil {
		return nil, err
	}

	return resp.Result().(*DiscussionVotesResponse), nil
}

type DiscussionsRequest struct {
	DiscussionBaseRequest
	BeatmapID        *int `json:"beatmap_id"`
	BeatmapsetID     *int `json:"beatmapset_id"`
	BeatmapsetStatus *DiscussionStatus
	MessageTypes     []MessageType `json:"message_types"`
	OnlyUnresolved   *bool         `json:"only_unresolved"`
	User             *int          `json:"user"`
}

// GetDiscussions returns a list of beatmapset discussions.
func (c *Client) GetDiscussions() *DiscussionsRequest {
	return &DiscussionsRequest{
		DiscussionBaseRequest: DiscussionBaseRequest{client: c},
	}
}

func (r *DiscussionsRequest) SetBeatmapID(beatmapID int) *DiscussionsRequest {
	r.BeatmapID = &beatmapID
	return r
}

func (r *DiscussionsRequest) SetBeatmapsetID(beatmapsetID int) *DiscussionsRequest {
	r.BeatmapsetID = &beatmapsetID
	return r
}

func (r *DiscussionsRequest) SetBeatmapsetStatus(status DiscussionStatus) *DiscussionsRequest {
	r.BeatmapsetStatus = &status
	return r
}

func (r *DiscussionsRequest) AddMessageType(types MessageType) *DiscussionsRequest {
	r.MessageTypes = append(r.MessageTypes, types)
	return r
}

func (r *DiscussionsRequest) SetMessageTypes(types []MessageType) *DiscussionsRequest {
	r.MessageTypes = types
	return r
}

func (r *DiscussionsRequest) SetOnlyUnresolved(onlyUnresolved bool) *DiscussionsRequest {
	r.OnlyUnresolved = &onlyUnresolved
	return r
}

func (r *DiscussionsRequest) SetUser(userID int) *DiscussionsRequest {
	r.User = &userID
	return r
}

func (r *DiscussionsRequest) SetLimit(limit int) *DiscussionsRequest {
	r.Limit = &limit
	return r
}

func (r *DiscussionsRequest) SetPage(page int) *DiscussionsRequest {
	r.Page = &page
	return r
}

func (r *DiscussionsRequest) SetSort(sort Sort) *DiscussionsRequest {
	r.Sort = &sort
	return r
}

func (r *DiscussionsRequest) Build() (*DiscussionsResponse, error) {
	req := r.client.httpClient.R().SetResult(&DiscussionsResponse{})

	if r.BeatmapID != nil {
		req.SetQueryParam("beatmap_id", strconv.Itoa(*r.BeatmapID))
	}

	if r.BeatmapsetID != nil {
		req.SetQueryParam("beatmapset_id", strconv.Itoa(*r.BeatmapsetID))
	}

	if r.BeatmapsetStatus != nil {
		req.SetQueryParam("beatmapset_status", string(*r.BeatmapsetStatus))
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

	if r.MessageTypes != nil {
		for _, messageType := range r.MessageTypes {
			req.QueryParam.Add("message_types[]", string(messageType))
		}
	}

	if r.OnlyUnresolved != nil {
		req.SetQueryParam("only_unresolved", strconv.FormatBool(*r.OnlyUnresolved))
	}

	resp, err := req.Get("beatmapsets/discussions")
	if err != nil {
		return nil, err
	}

	return resp.Result().(*DiscussionsResponse), nil
}
