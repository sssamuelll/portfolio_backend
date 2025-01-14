package services

import (
	"encoding/json"

	"github.com/portfolio-backend/models"
	"github.com/portfolio-backend/storage"
)

func GetAllPosts() ([]models.Post, error) {
	rows, err := storage.DB.Query("SELECT * FROM posts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		var tags, media string
		if err := rows.Scan(&post.ID, &post.Image, &post.Name, &post.Description, &post.Category, &tags, &media, &post.StartDate, &post.EndDate, &post.Link); err != nil {
			return nil, err
		}

		// Convertir JSON strings a arrays
		json.Unmarshal([]byte(tags), &post.Tags)
		json.Unmarshal([]byte(media), &post.Media)

		posts = append(posts, post)
	}
	return posts, nil
}

func CreatePost(post *models.Post) error {
	tagsJSON, _ := json.Marshal(post.Tags)
	mediaJSON, _ := json.Marshal(post.Media)

	query := `
	INSERT INTO posts (image, name, description, category, tags, media, start_date, end_date, link)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
	_, err := storage.DB.Exec(query, post.Image, post.Name, post.Description, post.Category, tagsJSON, mediaJSON, post.StartDate, post.EndDate, post.Link)
	return err
}
