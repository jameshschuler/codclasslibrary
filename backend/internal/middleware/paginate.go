package middleware

import (
	"context"
	"net/http"
	"strconv"

	"backend/internal/common"

	"github.com/go-chi/render"
)

type PaginationCtxKey string

type Pagination struct {
	PageSize int64
	Page     int64
}

func Paginate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var err error
		var pageSize int64 = 10
		var page int64 = 1

		if r.URL.Query().Get("pageSize") != "" {
			pageSize, err = strconv.ParseInt(r.URL.Query().Get("pageSize"), 10, 64)

			// TODO: return better error message
			if err != nil {
				render.Render(w, r, common.ErrBadRequest)
				return
			}

			if pageSize <= 0 || pageSize > 100 {
				render.Render(w, r, common.ErrBadRequest)
				return
			}
		}

		if r.URL.Query().Get("page") != "" {
			page, err = strconv.ParseInt(r.URL.Query().Get("page"), 10, 64)

			// TODO: return better error message
			if err != nil {
				render.Render(w, r, common.ErrBadRequest)
				return
			}

			if page <= 0 {
				render.Render(w, r, common.ErrBadRequest)
				return
			}
		}

		ctx := context.WithValue(r.Context(), PaginationCtxKey("pagination"), &Pagination{pageSize, page})
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
