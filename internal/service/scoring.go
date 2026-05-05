package service

import (
	"biami-engine/internal/model"
	"time"
)

// ScoringService handles the evaluation of survey integrity
type ScoringService struct{}

// NewScoringService creates and returns a new scoring service instance.
func NewScoringService() *ScoringService {
	// Return a pointer so methods can be called without copying the struct.
	return &ScoringService{}
}

// CalculateScore computes a quality score for a survey response session.
func (s *ScoringService) CalculateScore(response model.SurveyResponse, startTime time.Time) int {
	// Start with a perfect score and subtract points for suspicious patterns.
	score := 100

	// 1) Velocity check: measure how long the respondent took to submit.
	duration := time.Since(startTime).Seconds()
	// If the response is too fast, reduce confidence in response quality.
	if duration < 10 { // 10 seconds is the current least expected duration.
		score -= 40
	}

	// 2) Placeholder for consistency checks between related answers.
	// Example: compare Q1 and Q5 if they should logically match.
	_ = response // Keep parameter for upcoming scoring checks that use response data.

	// Return the final calculated score.
	return score
}
