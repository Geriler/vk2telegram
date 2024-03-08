package messenger

type Group struct {
	ID         int    `json:"id" db:"id"`
	Name       string `json:"name" db:"name"`
	ScreenName string `json:"screen_name" db:"screen_name"`
	Photo      string `json:"photo_200"`
}

type WrapperGroup struct {
	Response ResponseGroup `json:"response"`
}

type ResponseGroup struct {
	Count int     `json:"count"`
	Items []Group `json:"items"`
}
