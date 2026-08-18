package cmd

import (
	"context"
	"strings"
)

type RealFakeGarminAI struct{}

func (b *RealFakeGarminAI) Complete(ctx context.Context, request GarminAIRequest) (*GarminAICompletion, error) {
	content := request.Messages[len(request.Messages)-1].Content
	var out string

	if strings.Contains(content, "youtube") {
		out += "youtuber :sob:"
	} else if strings.Contains(content, "i love you") {
		out += "nah i'm good :wilted_rose:"
	} else if strings.Contains(content, "say") {
		out += content
	} else if strings.Contains(content, "moyai") {
		out += ":moyai:"
	} else {
		out += "no"
	}

	return &GarminAICompletion{
			GarminAIMessage{
				Reasoning: "no",
				Role:      "yes",
				Content:   out,
			},
			"success"},
		nil
}
