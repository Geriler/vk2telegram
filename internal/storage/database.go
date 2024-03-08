package storage

import (
	"github.com/jmoiron/sqlx"
	"vk2telegram/internal/errors"
	"vk2telegram/internal/model/messenger"
	"vk2telegram/internal/model/user"

	_ "github.com/lib/pq"
)

type DatabaseStorage struct {
	db *sqlx.DB
}

func Init() (*DatabaseStorage, error) {
	db, err := sqlx.Connect("postgres", "postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		return nil, err
	}

	return &DatabaseStorage{db: db}, nil
}

func (s *DatabaseStorage) Close() error {
	return s.db.Close()
}

func (s *DatabaseStorage) SaveGroup(group messenger.Group) error {
	_, err := s.db.Exec("INSERT INTO groups (name, screen_name) VALUES ($1, $2)", group.Name, group.ScreenName)
	if err != nil {
		return err
	}

	return nil
}

func (s *DatabaseStorage) BindUserGroup(u user.User, group messenger.Group) error {
	_, err := s.db.Exec("INSERT INTO user_group (user_id, group_id) VALUES ($1, $2)", u.ID, group.ID)
	if err != nil {
		return err
	}

	return nil
}

func (s *DatabaseStorage) GetGroups(userID int) ([]*messenger.Group, error) {
	var groups []*messenger.Group
	err := s.db.Select(&groups, "SELECT g.* FROM groups g JOIN user_group ug ON g.id = ug.group_id WHERE ug.user_id = $1", userID)
	if err != nil {
		return nil, err
	}

	return groups, nil
}

func (s *DatabaseStorage) SavePost(post messenger.Post, group messenger.Group) (int64, error) {
	res, err := s.db.Exec("INSERT INTO posts (text, date, post_type, is_pinned, marked, group_id, is_send) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		post.Text, post.Date, post.PostType, post.IsPinned, post.MarkedAsAds, group.ID, false)
	if err != nil {
		return 0, err
	}

	id, _ := res.LastInsertId()

	return id, nil
}

func (s *DatabaseStorage) GetPosts(group messenger.Group) ([]*messenger.Post, error) {
	var posts []*messenger.Post
	err := s.db.Select(&posts, "SELECT p.* FROM posts p WHERE group_id = $1", group.ID)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (s *DatabaseStorage) GetNotSendPostsByUser(u user.User) ([]*messenger.Post, error) {
	var posts []*messenger.Post
	err := s.db.Select(&posts, "SELECT p.* FROM posts p JOIN user_group ug ON ug.group_id = p.group_id WHERE p.is_send = FALSE AND ug.user_id = $1", u.ID)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

//func (s *DatabaseStorage) GetUserSettings(ctx context.Context, userID int) (*user.UserSettings, error) {
//	var settings user.UserSettings
//	err := s.db.GetContext(ctx, &settings, "SELECT * FROM user_settings WHERE user_id =$1", userID)
//	if err != nil {
//		return nil, err
//	}
//
//	return &settings, nil
//}

func (s *DatabaseStorage) GetUser(chatID int64) (*user.User, error) {
	var u user.User
	err := s.db.Get(&u, "SELECT u.* FROM users u WHERE chat_id = $1", chatID)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (s *DatabaseStorage) AddUser(firstName, lastName string, chatID int64) (*user.User, error) {
	u, err := s.GetUser(chatID)
	if u != nil {
		return u, errors.UserAlreadyRegistered
	}

	_, err = s.db.Exec("INSERT INTO users (first_name, last_name, chat_id) VALUES ($1, $2, $3)", firstName, lastName, chatID)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
