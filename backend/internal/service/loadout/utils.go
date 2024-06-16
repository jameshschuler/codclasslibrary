package loadout

import (
	"backendv2/gen/postgres/public/model"
	"net/http"

	"github.com/go-chi/render"
)

func NewLoadoutResponse(loadout model.Loadouts) LoadoutResponse {
	resp := LoadoutResponse{
		ID:             loadout.ID,
		CreatedAt:      loadout.CreatedAt,
		UpdatedAt:      loadout.UpdatedAt,
		Title:          loadout.Title,
		Source:         loadout.Source,
		SourceURL:      loadout.SourceURL,
		WeaponName:     loadout.WeaponName,
		WeaponCategory: loadout.WeaponCategory,
		CreatedBy:      loadout.CreatedBy,
		GameID:         loadout.GameID,
		Attachments:    loadout.Attachments,
	}
	return resp
}

func NewLoadoutsListResponse(loadouts []model.Loadouts) []render.Renderer {
	list := []render.Renderer{}
	for _, loadout := range loadouts {
		list = append(list, NewLoadoutResponse(loadout))
	}
	return list
}

func (rd LoadoutResponse) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	return nil
}
