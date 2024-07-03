package utils

import (
	"backend/internal/types"

	"github.com/go-chi/render"
)

func ErrInvalidRequest(err error) render.Renderer {
	return &types.ErrResponse{
		Err:            err,
		HTTPStatusCode: 400,
		StatusText:     "Invalid request.",
		ErrorText:      err.Error(),
	}
}

func ErrRender(err error) render.Renderer {
	return &types.ErrResponse{
		Err:            err,
		HTTPStatusCode: 422,
		StatusText:     "Error rendering response.",
		ErrorText:      err.Error(),
	}
}

var ErrNotFound = &types.ErrResponse{HTTPStatusCode: 404, StatusText: "Resource not found."}
