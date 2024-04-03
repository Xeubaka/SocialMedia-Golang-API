package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

// users represents a user repositorie
type UserRepository struct {
	db *sql.DB
}

// NewUserRepository creates a new repositorie of users
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}

// Create insert a new user on database
func (userRepository UserRepository) Create(user models.User) (userID uint64, err error) {
	statement, err := userRepository.db.Prepare("insert into users (name, nick, email, password) values (?,?,?,?)")
	if err != nil {
		return
	}
	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if err != nil {
		return
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return
	}
	userID = uint64(lastInsertID)

	return
}

// Serach for a user given by name or nick
func (userRepository UserRepository) Search(nameOrNick string) (users []models.User, err error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick) // %nameOrNick%

	lines, err := userRepository.db.Query(
		"select id, name, nick, email, createdAt from users where name LIKE ? or nick LIKE ?",
		nameOrNick,
		nameOrNick,
	)
	if err != nil {
		return
	}
	defer lines.Close()

	for lines.Next() {
		var user models.User

		if err = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return
		}

		users = append(users, user)
	}

	return
}

// SearchByID search a user by its ID
func (userRepository UserRepository) SerachByID(ID uint64) (user models.User, err error) {
	lines, err := userRepository.db.Query(
		"select id, name, nick, email, createdAt from users where id = ?",
		ID,
	)
	if err != nil {
		return
	}
	defer lines.Close()

	if lines.Next() {
		if err = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return
		}
	}

	return
}

// SerachByEmail searchs a user by its Email
func (userRepository UserRepository) SearchByEmail(email string) (user models.User, err error) {
	line, err := userRepository.db.Query("select id, password from users where email = ?", email)
	if err != nil {
		return
	}
	defer line.Close()

	if line.Next() {
		if err = line.Scan(&user.ID, &user.Password); err != nil {
			return
		}
	}

	return
}

// Update user information on database
func (userRepository UserRepository) Update(ID uint64, user models.User) (err error) {
	statement, err := userRepository.db.Prepare(
		"update users set name = ?, nick = ?, email = ? where id = ?",
	)
	if err != nil {
		return
	}
	defer statement.Close()

	if _, err = statement.Exec(user.Name, user.Nick, user.Email, ID); err != nil {
		return
	}

	return
}

// Delete from user by ID
func (userRepository UserRepository) Delete(ID uint64) (err error) {
	statement, err := userRepository.db.Prepare("delete from users where id = ?")
	if err != nil {
		return
	}
	defer statement.Close()

	if _, err = statement.Exec(ID); err != nil {
		return
	}

	return
}

// FollowUser permits an user to follow another
func (userRepository UserRepository) FollowUser(userID, followerID uint64) (err error) {
	statement, err := userRepository.db.Prepare(
		"insert ignore into followers (user_id, follower_id) values (?, ?)",
	)
	if err != nil {
		return
	}
	defer statement.Close()

	if _, err = statement.Exec(userID, followerID); err != nil {
		return
	}

	return
}

// UnFollowUser permits an user to follow another
func (userRepository UserRepository) UnFollowUser(userID, followerID uint64) (err error) {
	statement, err := userRepository.db.Prepare(
		"delete from followers where user_id = ? and follower_id = ?",
	)
	if err != nil {
		return
	}
	defer statement.Close()

	if _, err = statement.Exec(userID, followerID); err != nil {
		return
	}

	return
}

// SearchFollowers gets all followers from a user given its ID
func (userRepository UserRepository) SearchFollowers(userID uint64) (users []models.User, err error) {
	lines, err := userRepository.db.Query(`
		select u.id, u.name, u.nick, u.email, u.createdAt
		from users u 
		inner join followers f on u.id = f.follower_id
		where f.user_id = ?
	`, userID)
	if err != nil {
		return
	}
	defer lines.Close()

	for lines.Next() {
		var user models.User

		if err = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return
		}

		users = append(users, user)
	}

	return
}

// SearchFollowers gets all followers from a user given its ID
func (userRepository UserRepository) SearchFollowing(userID uint64) (users []models.User, err error) {
	lines, err := userRepository.db.Query(`
		select u.id, u.name, u.nick, u.email, u.createdAt
		from users u 
		inner join followers f on u.id = f.follower_id
		where f.follower_id = ?
	`, userID)
	if err != nil {
		return
	}
	defer lines.Close()

	for lines.Next() {
		var user models.User

		if err = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return
		}

		users = append(users, user)
	}

	return
}

// SearchPassword gets a Hashed password by user's ID
func (userRepository UserRepository) SearchPassword(userID uint64) (hashedPassword string, err error) {
	line, err := userRepository.db.Query("select password from users where id = ?", userID)
	if err != nil {
		return
	}
	defer line.Close()

	var user models.User

	if line.Next() {
		if err = line.Scan(&user.Password); err != nil {
			return
		}
	}

	hashedPassword = user.Password
	return
}

// UpdatePassword updates users password
func (userRepository UserRepository) UpdatePassword(userID uint64, hashedPassword string) (err error) {
	statement, err := userRepository.db.Prepare("update users set password = ? where id = ?")
	if err != nil {
		return
	}
	defer statement.Close()

	if _, err = statement.Exec(hashedPassword, userID); err != nil {
		return
	}

	return
}
