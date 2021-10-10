package repository

import (
	"api/src/models"
	"database/sql"
)

type Post struct {
	db *sql.DB
}

//NewPostRepository creates a new posts repository
func NewPostRepository(db *sql.DB) *Post {
	return &Post{db}
}

//Create inserts a new post in the DB
func (repository Post) Create(post models.Post) (uint64, error) {
	statement, err := repository.db.Prepare(
		"insert into posts (title, content, poster_id) values (?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(post.Title, post.Content, post.PosterID)
	if err != nil {
		return 0, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastId), nil
}

//Search gets the user and the user's following users from the DB
func (repository Post) Search(userId uint64) ([]models.Post, error) {
	lines, err := repository.db.Query(`
	select distinct p.*, u.nick from posts p
	inner join users u on u.id = p.poster_id
	inner join followers f on p.poster_id = f.ID_user
	where u.id = ? or f.ID_follower = ?
	order by 1 desc`,
		userId, userId)
	if err != nil {
		return nil, err
	}
	defer lines.Close()

	var posts []models.Post
	for lines.Next() {
		var post models.Post
		if err = lines.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.PosterID,
			&post.Likes,
			&post.CreatedOn,
			&post.PosterNick,
		); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

//SearchById gets a specif post from the DB
func (repository Post) SearchById(postId uint64) (models.Post, error) {
	line, err := repository.db.Query(`
		select p.*, u.nick from
		posts p inner join users u
		on u.id = p.poster_id where p.id = ?
	`, postId)
	if err != nil {
		return models.Post{}, err
	}
	defer line.Close()

	var post models.Post
	if line.Next() {
		if err = line.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.PosterID,
			&post.Likes,
			&post.CreatedOn,
			&post.PosterNick,
		); err != nil {
			return models.Post{}, err
		}
	}

	return post, nil
}

//Update changes the data of a post from the DB
func (repository Post) Update(postId uint64, post models.Post) error {
	statement, err := repository.db.Prepare("update posts set title = ?, content = ? where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(post.Title, post.Content, postId); err != nil {
		return err
	}

	return nil
}

//Delete deletes a post from the DB
func (repository Post) Delete(postId uint64) error {
	statement, err := repository.db.Prepare("delete from posts where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(postId); err != nil {
		return err
	}

	return nil
}

////SearchById gets all the posts of a specific user from the DB
func (repository Post) SearchByUser(userId uint64) ([]models.Post, error) {
	lines, err := repository.db.Query(`
	select p.*, u.nick from posts p
	join users u on u.id = p.poster_id
	where p.poster_id = ?
	`, userId)
	if err != nil {
		return nil, err
	}
	defer lines.Close()

	var posts []models.Post

	for lines.Next() {
		var post models.Post
		if err = lines.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.PosterID,
			&post.Likes,
			&post.CreatedOn,
			&post.PosterNick,
		); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

//Like adds a like in the post
func (repository Post) Like(postId uint64) error {
	statement, err := repository.db.Prepare("update posts set likes = likes + 1 where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(postId); err != nil {
		return err
	}

	return nil
}

//Dislike adds a like in the post
func (repository Post) Dislike(postId uint64) error {
	statement, err := repository.db.Prepare(`update posts set likes = 
		CASE WHEN likes > 0 THEN likes - 1 Else 0 END
		where id = ?
	`)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(postId); err != nil {
		return err
	}

	return nil
}
