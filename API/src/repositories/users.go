package repository

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

type Users struct {
	db *sql.DB
}

//newUsersRepository creates a users repository
func NewUsersRepository(db *sql.DB) *Users {
	return &Users{db}
}

//Create inserts a new user in the DB
func (repository Users) Create(user models.User) (uint64, error) {
	statement, err := repository.db.Prepare(
		"insert into users (name, nick, email, password) values (?, ?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastID), nil
}

//Search brings all the users who match with the name or nick filter
func (repository Users) Search(nameOrNick string) ([]models.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick)

	lines, err := repository.db.Query(
		"select id, name, nick, email, createdON from users where name LIKE ? or nick LIKE ?",
		nameOrNick, nameOrNick,
	)
	if err != nil {
		return nil, err
	}
	defer lines.Close()

	var users []models.User

	for lines.Next() {
		var user models.User
		if err = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedOn,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

//SearchByID brings from the DB the user that has the ID passed by parameter
func (repository Users) SearchByID(ID uint64) (models.User, error) {
	lines, err := repository.db.Query(
		"Select id, name, nick, email, createdOn from users where ID = ?", ID,
	)
	if err != nil {
		return models.User{}, nil
	}
	defer lines.Close()

	var user models.User
	if lines.Next() {
		if err = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedOn,
		); err != nil {
			return models.User{}, nil
		}
	}

	return user, nil
}

//Update updates a user's content in the DB
func (repository Users) Update(ID uint64, user models.User) error {
	statement, err := repository.db.Prepare(
		"update users set name = ?, nick = ?, email = ? where id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(user.Name, user.Nick, user.Email, ID); err != nil {
		return err
	}

	return nil
}

//Delete excludes the content of a user from the DB
func (repository Users) Delete(ID uint64) error {
	statement, err := repository.db.Prepare("delete from users where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(ID); err != nil {
		return err
	}

	return nil
}

//SearchByEmail brings from the DB the user that has the email passed by parameter
func (repository Users) SearchByEmail(email string) (models.User, error) {
	lines, err := repository.db.Query(
		"Select id, password from users where email = ?", email,
	)
	if err != nil {
		return models.User{}, err
	}
	defer lines.Close()

	var user models.User
	if lines.Next() {
		if err = lines.Scan(&user.ID, &user.Password); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

//Follow allows a user to follow aother
func (repository Users) Follow(userId, followerId uint64) error {
	statement, err := repository.db.Prepare(
		"insert ignore into followers (ID_user, ID_follower) values (?, ?)",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(userId, followerId); err != nil {
		return err
	}

	return nil
}

//Unfollow allows a user to unfollow aother
func (repository Users) Unfollow(userId, followerId uint64) error {
	statement, err := repository.db.Prepare(
		"delete from followers where ID_user = ? ad ID_follower = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(userId, followerId); err != nil {
		return err
	}

	return nil
}

//SearchFollowers brings from the DB the user's followers
func (repository Users) SearchFollowers(userId uint64) ([]models.User, error) {
	lines, err := repository.db.Query(`
		select u.id, u.name, u.nick, u.email, u.createdOn
		from users u inner join followers s on u.id = s.ID_follower where s.ID_user = ?
	`, userId)
	if err != nil {
		return nil, err
	}
	defer lines.Close()

	var followers []models.User
	for lines.Next() {
		var follower models.User
		if err = lines.Scan(
			&follower.ID,
			&follower.Name,
			&follower.Nick,
			&follower.Email,
			&follower.CreatedOn,
		); err != nil {
			return nil, err
		}
		followers = append(followers, follower)
	}

	return followers, nil
}

//SearchFollowing brings from the DB the users that a user is following
func (repository Users) SearchFollowing(userId uint64) ([]models.User, error) {
	lines, err := repository.db.Query(`
	    select u.id, u.name, u.nick, u.email, u.createdOn
		from users u inner join followers s on u.id = s.ID_user where s.ID_follower = ?
	`, userId)
	if err != nil {
		return nil, err
	}
	defer lines.Close()

	var users []models.User
	for lines.Next() {
		var user models.User
		if err = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedOn,
		); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, err
}

//SearchPassword gets a user's password by ID
func (repository Users) SearchPassword(userId uint64) (string, error) {
	line, err := repository.db.Query("select password from users where id = ?", userId)
	if err != nil {
		return "", err
	}
	defer line.Close()

	var user models.User
	if line.Next() {
		if err = line.Scan(&user.Password); err != nil {
			return "", err
		}
	}

	return user.Password, nil
}

//UpdatePassword updates a user's password
func (repository Users) UpdatePassword(userId uint64, NewPassword string) error {
	statement, err := repository.db.Prepare("update users set password = ? where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(NewPassword, userId); err != nil {
		return err
	}

	return nil
}
