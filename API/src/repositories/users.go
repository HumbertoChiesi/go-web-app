package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

type Users struct {
	db *sql.DB
}

//newUsersRepository creates an users repository
func NewUsersRepository(db *sql.DB) *Users {
	return &Users{db}
}

func (repository Users) Create(user models.User) (uint64, error) {
	statement, err := repository.db.Prepare(
		"insert into usuarios (name, nick, email, password) values (?, ?, ?, ?)",
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
		"select id, name, nick, email, createdON from usuarios where name LIKE ? or nick LIKE ?",
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

//SearchByID brings the specific user that has the ID passed by parameter
func (repository Users) SearchByID(ID uint64) (models.User, error) {
	lines, err := repository.db.Query(
		"Select id, name, nick, email, createdOn from usuarios where ID = ?", ID,
	)
	if err != nil {
		return models.User{}, nil
	}

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

//Update updates an user's content in the DB
func (repository Users) Update(ID uint64, user models.User) error {
	statement, err := repository.db.Prepare(
		"update usuarios set name = ?, nick = ?, email = ? where id = ?",
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

//Delete excludes the content of an user from the DB
func (repository Users) Delete(ID uint64) error {
	statement, err := repository.db.Prepare("delete from usuarios where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(ID); err != nil {
		return err
	}

	return nil
}
