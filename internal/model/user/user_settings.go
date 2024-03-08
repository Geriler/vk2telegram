package user

type UserSettings struct {
	UserID  int    `db:"user_id"`
	VkToken string `db:"vk_token"`
}
