package gosu

type WikiPage struct {
	AvailableLocales []string `json:"available_locales"`
	Layout           string   `json:"layout"`
	Locale           string   `json:"locale"`
	Markdown         string   `json:"markdown"`
	Path             string   `json:"path"`
	Subtitle         *string  `json:"subtitle"`
	Tags             []string `json:"tags"`
	Title            string   `json:"title"`
}

type WikiPageRequest struct {
	client *Client
	Locale string
	Path   string
}

// GetWikiPage retyrns a page from the wiki.
func (c *Client) GetWikiPage(locale string, path string) *WikiPageRequest {
	return &WikiPageRequest{client: c, Locale: locale, Path: path}
}

func (r *WikiPageRequest) Build() (*WikiPage, error) {
	resp, err := r.client.httpClient.R().SetResult(&WikiPage{}).SetPathParams(map[string]string{
		"locale": r.Locale,
		"path":   r.Path,
	}).Get("wiki/{locale}/{path}")
	if err != nil {
		return nil, err
	}

	return resp.Result().(*WikiPage), nil
}
