package vkutil

type Post struct {
	Id        int       `json:"id"`
	OwnerId   int       `json:"owner_id"`
	FromId    int       `json:"from_id"`
	Date      EpochTime `json:"date"`
	AccessKey string    `json:"access_key"`
	Likes     struct {
		Count int `json:"count"`
	} `json:"count"`
	Views struct {
		Count int `json:"count"`
	} `json:"views"`
	Reposts struct {
		Count int `json:"count"`
	} `json:"reposts"`
	Text     string `json:"text"`
	PostType string `json:"post_type"`

	CopyHistory []Post     `json:"copy_history"`
	PostSource  PostSource `json:"post_source"`

	SignerID int `json:"signer_id"`

	Attachments []Attachment `json:"attachments"`
}

type PostSource struct {
	Type string `json:"type"`
}
