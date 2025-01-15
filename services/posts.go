package services

import (
	"encoding/json"
	"errors"

	"github.com/sssamuelll/portfolio_backend/models"
	"github.com/sssamuelll/portfolio_backend/storage"
)

// GetAllPosts obtiene todos los posts desde la base de datos
func GetAllPosts() ([]models.Post, error) {
	var posts []models.Post
	err := storage.DB.Find(&posts).Error
	if err != nil {
		return nil, err
	}

	// Convertir cadenas JSON a arreglos en Tags y Media
	for i := range posts {
		if posts[i].TagsJSON != "" {
			if err := json.Unmarshal([]byte(posts[i].TagsJSON), &posts[i].Tags); err != nil {
				return nil, errors.New("failed to parse tags JSON")
			}
		}
		if posts[i].MediaJSON != "" {
			if err := json.Unmarshal([]byte(posts[i].MediaJSON), &posts[i].Media); err != nil {
				return nil, errors.New("failed to parse media JSON")
			}
		}
	}
	return posts, nil
}

// CreatePost crea un nuevo post en la base de datos
func CreatePost(post *models.Post) error {
	// Convertir Tags y Media a JSON
	tagsJSON, err := json.Marshal(post.Tags)
	if err != nil {
		return errors.New("failed to convert tags to JSON")
	}
	mediaJSON, err := json.Marshal(post.Media)
	if err != nil {
		return errors.New("failed to convert media to JSON")
	}

	post.TagsJSON = string(tagsJSON)
	post.MediaJSON = string(mediaJSON)

	// Guardar el post en la base de datos
	return storage.DB.Create(post).Error
}
