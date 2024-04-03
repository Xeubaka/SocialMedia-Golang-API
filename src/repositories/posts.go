package repositories

import (
	"api/src/models"
	"database/sql"
)

// PostsRepository represents a repository of posts
type PostsRepository struct {
	db *sql.DB
}

// NewPostRepository creates a new postRepository
func NewPostRepository(db *sql.DB) *PostsRepository {
	return &PostsRepository{db}
}

// CreatePost inserts a post on the database
func (postsRepository PostsRepository) CreatePost(post models.Post) (userID uint64, err error) {
	statement, err := postsRepository.db.Prepare(
		"insert into posts (title, content, author_id) values (?, ?, ?)",
	)
	if err != nil {
		return
	}
	defer statement.Close()

	result, err := statement.Exec(post.Title, post.Content, post.AuthorID)
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

// SearchByID search a post by its ID
func (postsRepository PostsRepository) SearchByID(postID uint64) (post models.Post, err error) {
	lines, err := postsRepository.db.Query(`
		select p.*, u.nick 
		from posts p 
		inner join users u on u.id = p.author_id
		where p.id = ?`,
		postID,
	)
	if err != nil {
		return
	}
	defer lines.Close()

	if lines.Next() {
		if err = lines.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorNick,
		); err != nil {
			return
		}
	}

	return
}

// Search gets all posts from the user and those that he follows
func (postsRepository PostsRepository) Search(userID uint64) (posts []models.Post, err error) {
	lines, err := postsRepository.db.Query(`
		select distinct p.*, u.nick 
		from posts p 
		inner join users u on u.id = p.author_id 
		left join followers f on p.author_id = f.user_id 
		where u.id = ? or f.follower_id = ?
		order by p.createdAt DESC`,
		userID, userID,
	)
	if err != nil {
		return
	}
	defer lines.Close()

	for lines.Next() {
		var post models.Post

		if err = lines.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorNick,
		); err != nil {
			return
		}

		posts = append(posts, post)
	}

	return
}

// UpdatePost update post's informations
func (postsRepository PostsRepository) UpdatePost(postID uint64, post models.Post) (err error) {
	statement, err := postsRepository.db.Prepare("update posts set title = ?, content = ? where id = ?")
	if err != nil {
		return
	}
	defer statement.Close()

	if _, err = statement.Exec(post.Title, post.Content, postID); err != nil {
		return
	}

	return
}

// DeletePost deletes a post from the Database
func (postsRepository PostsRepository) DeletePost(postID uint64) (err error) {
	statement, err := postsRepository.db.Prepare("delete from posts where id = ?")
	if err != nil {
		return
	}
	defer statement.Close()

	if _, err = statement.Exec(postID); err != nil {
		return
	}

	return
}

// SearchPostsByUser get all posts from an user
func (postsRepository PostsRepository) SearchPostsByUser(userID uint64) (posts []models.Post, err error) {
	lines, err := postsRepository.db.Query(`
		select p.*, u.nick 
		from posts p 
		inner join users u on u.id = p.author_id
		where p.author_id = ?`,
		userID,
	)
	if err != nil {
		return
	}
	defer lines.Close()

	for lines.Next() {
		var post models.Post

		if err = lines.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorNick,
		); err != nil {
			return
		}

		posts = append(posts, post)
	}

	return
}

// Like adds one like on a post
func (postsRepository PostsRepository) Like(postID uint64) (err error) {
	statement, err := postsRepository.db.Prepare(`
		update posts set likes = 
		CASE 
			WHEN likes > 0 THEN likes - 1
			ELSE 0 
		END
		where id = ?
	`)
	if err != nil {
		return
	}
	defer statement.Close()

	if _, err = statement.Exec(postID); err != nil {
		return
	}

	return
}

// Like removes one like on a post
func (postsRepository PostsRepository) UnLike(postID uint64) (err error) {
	statement, err := postsRepository.db.Prepare("update posts set likes = likes - 1 where id = ?")
	if err != nil {
		return
	}
	defer statement.Close()

	if _, err = statement.Exec(postID); err != nil {
		return
	}

	return
}
