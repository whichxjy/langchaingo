package tools

import "context"

// Tool is a tool for the llm agent to interact with different applications.
type Tool interface {
	Name() string
	Description() string
	Parameters() any
	Call(ctx context.Context, input string) (string, error)
}
