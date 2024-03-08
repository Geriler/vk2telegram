package user

type User struct {
	ID        int    `db:"id"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	ChatID    int64  `db:"chat_id"`
}
