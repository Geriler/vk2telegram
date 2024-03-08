package messenger

type Post struct {
	ID          int          `json:"id" db:"id"`
	Text        string       `json:"text" db:"text"`
	Date        int          `json:"date" db:"date"`
	PostType    PostType     `json:"post_type" db:"post_type"`
	Attachments []Attachment `json:"attachments"`
	IsPinned    int          `json:"is_pinned" db:"is_pinned"`
	MarkedAsAds int          `json:"marked_as_ads" db:"marked"`
	GroupID     int          `db:"group_id"`
	IsSend      bool         `db:"is_send"`
}

type PostType string

const (
	PostTypePost     PostType = "post"
	PostTypeCopy     PostType = "copy"
	PostTypeReply    PostType = "reply"
	PostTypePostpone PostType = "postpone"
	PostTypeSuggest  PostType = "suggest"
)

type WrapperPost struct {
	Response ResponsePost `json:"response"`
}

type ResponsePost struct {
	Count int    `json:"count"`
	Items []Post `json:"items"`
}
