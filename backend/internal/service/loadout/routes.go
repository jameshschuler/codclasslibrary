package loadout

import (
	"backend/gen/postgres/public/model"
	"backend/internal/common"
	"backend/internal/middleware"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/go-jet/jet/v2/qrm"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Handler struct {
	store     LoadoutStore
	validator *validator.Validate
}

func NewHandler(store LoadoutStore, validator *validator.Validate) *Handler {
	return &Handler{store: store, validator: validator}
}

func (h *Handler) RegisterRoutes(r *chi.Mux) {
	r.With(middleware.Paginate).Get("/community/loadouts", h.HandleListCommunityLoadouts)
	r.Get("/community/loadout/{loadoutId}", h.HandleGetCommunityLoadout)
	r.With(middleware.AuthMiddleware).With(middleware.Paginate).Get("/me/loadouts", h.HandleListLoadouts)
	r.With(middleware.AuthMiddleware).Get("/me/loadout/{loadoutId}", h.HandleGetLoadout)
	r.With(middleware.AuthMiddleware).Post("/me/loadout", h.HandleCreateLoadout)
}

func (handler *Handler) HandleListCommunityLoadouts(w http.ResponseWriter, r *http.Request) {
	pagination := r.Context().Value(middleware.PaginationCtxKey("pagination")).(*middleware.Pagination)

	loadouts, err := handler.store.ListCommunityLoadouts(pagination.Page, pagination.PageSize)
	if err != nil {
		render.Render(w, r, common.ErrRender(err))
		return
	}

	if err := render.RenderList(w, r, NewLoadoutsListResponse(loadouts)); err != nil {
		render.Render(w, r, common.ErrRender(err))
		return
	}
}

func (handler *Handler) HandleListLoadouts(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(middleware.UserIdCtxKey("userId")).(string)

	loadouts, err := handler.store.ListLoadoutsByUser(userId)
	if err != nil {
		render.Render(w, r, common.ErrRender(err))
		return
	}

	if err := render.RenderList(w, r, NewLoadoutsListResponse(loadouts)); err != nil {
		render.Render(w, r, common.ErrRender(err))
		return
	}
}

func (handler *Handler) HandleGetLoadout(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(middleware.UserIdCtxKey("userId")).(string)
	loadoutId := chi.URLParam(r, "loadoutId")

	if loadoutId == "" {
		render.Render(w, r, common.ErrBadRequest)
		return
	}

	loadout, err := handler.store.GetLoadoutByUser(userId, loadoutId)

	if err != nil {
		if err == qrm.ErrNoRows {
			render.Render(w, r, common.ErrNotFound)
			return
		}

		render.Render(w, r, common.ErrRender(err))
		return
	}

	if err := render.Render(w, r, NewLoadoutDetailResponse(loadout)); err != nil {
		render.Render(w, r, common.ErrRender(err))
		return
	}
}

func (handler *Handler) HandleGetCommunityLoadout(w http.ResponseWriter, r *http.Request) {
	loadoutId := chi.URLParam(r, "loadoutId")

	if loadoutId == "" {
		render.Render(w, r, common.ErrBadRequest)
		return
	}

	loadout, err := handler.store.GetCommunityLoadout(loadoutId)
	if err != nil {
		if err == qrm.ErrNoRows {
			render.Render(w, r, common.ErrNotFound)
			return
		}

		render.Render(w, r, common.ErrRender(err))
		return
	}

	if err := render.Render(w, r, NewLoadoutDetailResponse(loadout)); err != nil {
		render.Render(w, r, common.ErrRender(err))
		return
	}
}

func (handler *Handler) HandleCreateLoadout(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(middleware.UserIdCtxKey("userId")).(string)

	requestData := &CreateLoadoutRequest{}
	if err := render.Bind(r, requestData); err != nil {
		render.Render(w, r, common.ErrInvalidRequest(err, nil))
		return
	}

	err := handler.validator.Struct(requestData)
	if err != nil {
		render.Render(w, r, common.ErrInvalidRequest(err, common.ParseValidationErrors(err)))
		return
	}

	parsedUserId, err := uuid.Parse(userId)

	if err != nil {
		render.Render(w, r, common.ErrRender(err))
		return
	}

	loadout := &model.Loadouts{
		Title:          requestData.Title,
		Source:         &requestData.Source,
		SourceURL:      &requestData.SourceUrl,
		WeaponName:     requestData.WeaponName,
		WeaponCategory: requestData.WeaponCategory,
		GameID:         requestData.GameId,
		CreatedBy:      parsedUserId,
	}

	createdLoadoutId, err := handler.store.CreateLoadout(loadout)
	if err != nil {
		render.Render(w, r, common.ErrRender(err))
		return
	}

	w.Header().Set("Location", fmt.Sprintf("/me/loadout/%s", createdLoadoutId))
	w.WriteHeader(http.StatusCreated)
}
