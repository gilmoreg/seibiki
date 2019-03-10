package service

import (
	"github.com/gilmoreg/seibiki/internal/dictionary"
	"go.uber.org/zap"
)

// LookupService - interface for kagome service
type LookupService interface {
	Lookup(query string) []dictionary.Word
}

type lookupService struct {
	logger *zap.SugaredLogger
	repo   dictionary.Repository
}

// New returns a lookupService
func New(logger *zap.SugaredLogger, repo dictionary.Repository) LookupService {
	return &lookupService{
		logger: logger,
		repo:   repo,
	}
}

// Lookup - tokenize and lookup tokens in dictionary
func (s *lookupService) Lookup(query string) []dictionary.Word {
	words := dictionary.Tokenize(query)
	result := make([]dictionary.Word, 0)
	for _, word := range words {
		result = append(result, word.GetEntries(s.repo))
	}
	return result
}
