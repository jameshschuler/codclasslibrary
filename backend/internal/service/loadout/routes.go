package loadout

import (
	"backend/internal/middleware"
	"backend/internal/types"
	"backend/internal/utils"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type Handler struct {
	store types.LoadoutStore
}

func NewHandler(store types.LoadoutStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(r *chi.Mux) {
	r.With(middleware.Paginate).Get("/community/loadouts", h.HandleListCommunityLoadouts)
	r.With(middleware.AuthMiddleware).With(middleware.Paginate).Get("/loadouts", h.HandleListLoadouts)
}

func (handler *Handler) HandleListCommunityLoadouts(w http.ResponseWriter, r *http.Request) {
	pagination := r.Context().Value(middleware.PaginationCtxKey("pagination")).(*middleware.Pagination)

	loadouts, err := handler.store.ListCommunityLoadouts(pagination.Page, pagination.PageSize)
	if err != nil {
		render.Render(w, r, utils.ErrRender(err))
		return
	}

	if err := render.RenderList(w, r, NewLoadoutsListResponse(loadouts)); err != nil {
		render.Render(w, r, utils.ErrRender(err))
		return
	}
}

func (handler *Handler) HandleListLoadouts(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(middleware.UserIdCtxKey("userId")).(string)

	loadouts, err := handler.store.ListLoadoutsByUser(userId)
	if err != nil {
		render.Render(w, r, utils.ErrRender(err))
		return
	}

	if err := render.RenderList(w, r, NewLoadoutsListResponse(loadouts)); err != nil {
		render.Render(w, r, utils.ErrRender(err))
		return
	}
}
