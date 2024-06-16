package loadout

import (
	"backendv2/internal/types"
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
	r.Get("/loadouts", h.HandleListLoadouts)
}

func (h *Handler) HandleListLoadouts(w http.ResponseWriter, r *http.Request) {
	loadouts, err := h.store.ListLoadouts()
	if err != nil {
		render.Render(w, r, types.ErrRender(err))
		return
	}

	if err := render.RenderList(w, r, NewLoadoutsListResponse(loadouts)); err != nil {
		render.Render(w, r, types.ErrRender(err))
		return
	}
}
