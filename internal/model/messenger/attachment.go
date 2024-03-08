package messenger

type Attachment struct {
	Type  AttachmentType `json:"type"`
	Photo *Photo         `json:"photo,omitempty"`
	Link  *Link          `json:"link,omitempty"`
	Doc   *Document      `json:"doc,omitempty"`
}

type AttachmentType string

const (
	AttachmentTypePhoto       AttachmentType = "photo"
	AttachmentTypeVideo       AttachmentType = "video"
	AttachmentTypeAudio       AttachmentType = "audio"
	AttachmentTypeDocument    AttachmentType = "doc"
	AttachmentTypeLink        AttachmentType = "link"
	AttachmentTypeNote        AttachmentType = "note"
	AttachmentTypePoll        AttachmentType = "poll"
	AttachmentTypePage        AttachmentType = "page"
	AttachmentTypeAlbum       AttachmentType = "album"
	AttachmentTypePhotosList  AttachmentType = "photos_list"
	AttachmentTypeMarket      AttachmentType = "market"
	AttachmentTypeSticker     AttachmentType = "sticker"
	AttachmentTypePrettyCards AttachmentType = "pretty_cards"
	AttachmentTypeEvent       AttachmentType = "event"
)
