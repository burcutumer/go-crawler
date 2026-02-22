package storage

import "github.com/burcutumer/go-crawler/models"

type Storage interface {
    Save(article models.Article) error
    GetAll() ([]models.Article, error)
}