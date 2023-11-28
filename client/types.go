package client

type Item struct {
	By          string `json:"by"`
	Descendants int    `json:"descendants"`
	ID          int    `json:"id"`
	Kids        []int  `json:"kids"`
	Score       int    `json:"score"`
	Text        string `json:"text"`
	Time        int    `json:"time"`
	Title       string `json:"title"`
	ItemType    string `json:"type"` // "job", "story", "comment", "poll", or "pollopt".
	Deleted     bool   `json:"deleted"`
	Dead        bool   `json:"dead"`
	Parent      int    `json:"parent"`
	Poll        int    `json:"poll"`
	Url         string `json:"url"`
	Parts       []int  `json:"parts"`
}
