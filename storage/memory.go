package storage

import articleM "github.com/burcutumer/go-crawler/models"

type MemoryStorage struct {
    articles []articleM.Article
}