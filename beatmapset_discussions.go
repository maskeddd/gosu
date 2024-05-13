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

type DiscussionSort string

const (
	DiscussionSortDescending = "id_desc"
	DiscussionSortAscending  = "id_asc"
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
	Sort   *DiscussionSort
}

type GetDiscussionPostsRequest struct {
	DiscussionBaseRequest
	BeatmapsetDiscussionID *int
	Types                  []DiscussionPostType
	User                   *int
}

func (c *Client) GetDiscussionPosts() *GetDiscussionPostsRequest {
	return &GetDiscussionPostsRequest{
		DiscussionBaseRequest: DiscussionBaseRequest{client: c},
	}
}

func (r *GetDiscussionPostsRequest) SetDiscussionID(discussionID int) *GetDiscussionPostsRequest {
	r.BeatmapsetDiscussionID = &discussionID
	return r
}

func (r *GetDiscussionPostsRequest) AddType(types DiscussionPostType) *GetDiscussionPostsRequest {
	r.Types = append(r.Types, types)
	return r
}

func (r *GetDiscussionPostsRequest) SetTypes(types []DiscussionPostType) *GetDiscussionPostsRequest {
	r.Types = types
	return r
}

func (r *GetDiscussionPostsRequest) SetUser(userID int) *GetDiscussionPostsRequest {
	r.User = &userID
	return r
}

func (r *GetDiscussionPostsRequest) SetLimit(limit int) *GetDiscussionPostsRequest {
	r.Limit = &limit
	return r
}

func (r *GetDiscussionPostsRequest) SetPage(page int) *GetDiscussionPostsRequest {
	r.Page = &page
	return r
}

func (r *GetDiscussionPostsRequest) SetSort(sort DiscussionSort) *GetDiscussionPostsRequest {
	r.Sort = &sort
	return r
}

func (r *GetDiscussionPostsRequest) Build() (*DiscussionPostsResponse, error) {
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

type GetDiscussionVotesRequest struct {
	DiscussionBaseRequest
	BeatmapsetDiscussionID *int
	Receiver               *int                     `json:"receiver"`
	Score                  *DiscussionVoteDirection `json:"score"`
	User                   *int                     `json:"user"`
}

func (c *Client) GetDiscussionVotes() *GetDiscussionVotesRequest {
	return &GetDiscussionVotesRequest{
		DiscussionBaseRequest: DiscussionBaseRequest{client: c},
	}
}

func (r *GetDiscussionVotesRequest) SetDiscussionID(discussionID int) *GetDiscussionVotesRequest {
	r.BeatmapsetDiscussionID = &discussionID
	return r
}

func (r *GetDiscussionVotesRequest) SetReceiver(receiverID int) *GetDiscussionVotesRequest {
	r.Receiver = &receiverID
	return r
}

func (r *GetDiscussionVotesRequest) SetScore(score DiscussionVoteDirection) *GetDiscussionVotesRequest {
	r.Score = &score
	return r
}

func (r *GetDiscussionVotesRequest) SetUser(userID int) *GetDiscussionVotesRequest {
	r.User = &userID
	return r
}

func (r *GetDiscussionVotesRequest) SetLimit(limit int) *GetDiscussionVotesRequest {
	r.Limit = &limit
	return r
}

func (r *GetDiscussionVotesRequest) SetPage(page int) *GetDiscussionVotesRequest {
	r.Page = &page
	return r
}

func (r *GetDiscussionVotesRequest) SetSort(sort DiscussionSort) *GetDiscussionVotesRequest {
	r.Sort = &sort
	return r
}

func (r *GetDiscussionVotesRequest) Build() (*DiscussionVotesResponse, error) {
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

type GetDiscussionsRequest struct {
	DiscussionBaseRequest
	BeatmapID        *int `json:"beatmap_id"`
	BeatmapsetID     *int `json:"beatmapset_id"`
	BeatmapsetStatus *DiscussionStatus
	MessageTypes     []MessageType `json:"message_types"`
	OnlyUnresolved   *bool         `json:"only_unresolved"`
	User             *int          `json:"user"`
}

func (c *Client) GetDiscussions() *GetDiscussionsRequest {
	return &GetDiscussionsRequest{
		DiscussionBaseRequest: DiscussionBaseRequest{client: c},
	}
}

func (r *GetDiscussionsRequest) SetBeatmapID(beatmapID int) *GetDiscussionsRequest {
	r.BeatmapID = &beatmapID
	return r
}

func (r *GetDiscussionsRequest) SetBeatmapsetID(beatmapsetID int) *GetDiscussionsRequest {
	r.BeatmapsetID = &beatmapsetID
	return r
}

func (r *GetDiscussionsRequest) SetBeatmapsetStatus(status DiscussionStatus) *GetDiscussionsRequest {
	r.BeatmapsetStatus = &status
	return r
}

func (r *GetDiscussionsRequest) AddMessageType(types MessageType) *GetDiscussionsRequest {
	r.MessageTypes = append(r.MessageTypes, types)
	return r
}

func (r *GetDiscussionsRequest) SetMessageTypes(types []MessageType) *GetDiscussionsRequest {
	r.MessageTypes = types
	return r
}

func (r *GetDiscussionsRequest) SetOnlyUnresolved(onlyUnresolved bool) *GetDiscussionsRequest {
	r.OnlyUnresolved = &onlyUnresolved
	return r
}

func (r *GetDiscussionsRequest) SetUser(userID int) *GetDiscussionsRequest {
	r.User = &userID
	return r
}

func (r *GetDiscussionsRequest) SetLimit(limit int) *GetDiscussionsRequest {
	r.Limit = &limit
	return r
}

func (r *GetDiscussionsRequest) SetPage(page int) *GetDiscussionsRequest {
	r.Page = &page
	return r
}

func (r *GetDiscussionsRequest) SetSort(sort DiscussionSort) *GetDiscussionsRequest {
	r.Sort = &sort
	return r
}

func (r *GetDiscussionsRequest) Build() (*DiscussionsResponse, error) {
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
