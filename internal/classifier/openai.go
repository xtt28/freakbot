package classifier

import (
	"context"
	"errors"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

// Require compliance with interface.
var _ ClassifierService = &OpenAIClassifierService{}

type OpenAIClassifierService struct {
	client openai.Client
}

func (s *OpenAIClassifierService) IsFlagged(message string) (bool, error) {
	ctx := context.Background()
	res, err := s.client.Moderations.New(ctx, openai.ModerationNewParams{
		Input: openai.ModerationNewParamsInputUnion{OfString: openai.String(message)},
		Model: openai.ModerationModelOmniModerationLatest,
	})

	if err != nil {
		return false, err
	}

	if len(res.Results) < 1 {
		return false, errors.New("no results in moderation query")
	}
	result := res.Results[0]
	
	return result.Flagged, nil
}

func NewOpenAIClassifier(apiKey string) *OpenAIClassifierService {
	client := openai.NewClient(
		option.WithAPIKey(apiKey),
	)

	return &OpenAIClassifierService{client: client}
}