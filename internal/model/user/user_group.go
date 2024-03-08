package user

type UserGroup struct {
	UserID  int `db:"user_id"`
	GroupID int `db:"group_id"`
}
