package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	"html"

	"git.01.alem.school/bbaktyke/forum.git/pkg/models"
)

type PostModel struct {
	DB *sql.DB
}

func (m *PostModel) Insert(uid int, title, description, name, tag string) (int, error) {
	description = html.EscapeString(description)
	title = html.EscapeString(title)
	stmt := `INSERT INTO post (title, description, userid,username,category, created_at)
	VALUES(?, ?,?,?,?, datetime('now'))`
	result, err := m.DB.Exec(stmt, title, description, uid, name, tag)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

// This will return a specific snippet based on its id.
func (m *PostModel) Get(id int) (*models.Posts, error) {
	stmt := `SELECT id, title, userid, username,category, description, created_at FROM post
	WHERE id = ?`
	row := m.DB.QueryRow(stmt, id)
	s := &models.Posts{}

	err := row.Scan(&s.ID, &s.Title, &s.UserId, &s.UserName, &s.Category, &s.Description, &s.Created_At)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}
	return s, nil
}

func (m *PostModel) GetCategory(category string) ([]*models.Posts, error) {
	stmt := `SELECT id, title, userid, username,category, description, created_at FROM post
	WHERE category = ?`
	posts := []*models.Posts{}
	row, err := m.DB.Query(stmt, category)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	for row.Next() {
		s := &models.Posts{}
		err := row.Scan(&s.ID, &s.Title, &s.UserId, &s.UserName, &s.Category, &s.Description, &s.Created_At)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, models.ErrNoRecord
			} else {
				return nil, err
			}
		}
		posts = append(posts, s)
	}
	if err = row.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

func (m *PostModel) Latest() ([]*models.Posts, error) {
	stmt := `select * from post order by created_at DESC LIMIT 10`
	posts := []*models.Posts{}
	row, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	for row.Next() {
		s := &models.Posts{}
		err := row.Scan(&s.ID, &s.Title, &s.UserId, &s.UserName, &s.Category, &s.Description, &s.Created_At)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, models.ErrNoRecord
			} else {
				return nil, err
			}
		}
		posts = append(posts, s)
	}
	if err = row.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

func (m *PostModel) SeeAll() ([]*models.Posts, error) {
	stmt := `select * from post order by created_at DESC`
	posts := []*models.Posts{}
	row, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	for row.Next() {
		s := &models.Posts{}
		err := row.Scan(&s.ID, &s.Title, &s.Description, &s.Created_At)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, models.ErrNoRecord
			} else {
				return nil, err
			}
		}
		posts = append(posts, s)
	}
	if err = row.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

func (m *PostModel) MyPosts(userid int) ([]*models.Posts, error) {
	stmt := `select * from post where userid=?`
	posts := []*models.Posts{}
	row, err := m.DB.Query(stmt, userid)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	for row.Next() {
		s := &models.Posts{}
		err := row.Scan(&s.ID, &s.Title, &s.UserId, &s.UserName, &s.Category, &s.Description, &s.Created_At)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, models.ErrNoRecord
			} else {
				return nil, err
			}
		}
		posts = append(posts, s)
	}
	if err = row.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

func (m *PostModel) LikedPosts(userid int) ([]*models.Posts, error) {
	stmt := `select post.id, title, post.userid,post.username, post.category, post.description, created_at from post inner join reaction on post.id=reaction.postid where reaction.userid=19 and reaction.like=1`
	posts := []*models.Posts{}
	row, err := m.DB.Query(stmt, userid)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	for row.Next() {
		s := &models.Posts{}
		err := row.Scan(&s.ID, &s.Title, &s.UserId, &s.UserName, &s.Category, &s.Description, &s.Created_At)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, models.ErrNoRecord
			} else {
				return nil, err
			}
		}
		fmt.Println(s.ID)
		posts = append(posts, s)
	}
	if err = row.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

func (m *PostModel) PostAuthor(id int) (int, error) {
	stmt := `select userid from post where id=?`
	row := m.DB.QueryRow(stmt, id)
	var userid int
	err := row.Scan(&userid)
	if err != nil {
		return 0, err
	}
	return userid, nil
}
