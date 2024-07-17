package loadout

import (
	"backend/gen/postgres/public/model"
	"backend/internal/common"
	"net/http"

	"github.com/go-chi/render"
)

type Uuid string

func (s Uuid) String() string {
	return string(s)
}

func NewLoadoutResponse(loadout *model.Loadouts) *LoadoutResponse {
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
	}
	return &resp
}

func NewLoadoutDetailResponse(loadoutDetail *LoadoutDetail) *LoadoutResponse {
	resp := LoadoutResponse{
		ID:             loadoutDetail.ID,
		CreatedAt:      loadoutDetail.CreatedAt,
		UpdatedAt:      loadoutDetail.UpdatedAt,
		Title:          loadoutDetail.Title,
		Source:         loadoutDetail.Source,
		SourceURL:      loadoutDetail.SourceURL,
		WeaponName:     loadoutDetail.WeaponName,
		WeaponCategory: loadoutDetail.WeaponCategory,
		CreatedBy:      loadoutDetail.CreatedBy,
		GameID:         loadoutDetail.GameID,
	}

	if loadoutDetail.Attachments != nil {
		resp.Attachments = NewAttachmentsListResponse(&loadoutDetail.Attachments)
	}

	return &resp
}

func NewLoadoutsListResponse(loadouts *[]model.Loadouts) []render.Renderer {
	list := []render.Renderer{}
	for _, loadout := range *loadouts {
		list = append(list, NewLoadoutResponse(&loadout))
	}
	return list
}

func (rd LoadoutResponse) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	return nil
}

func NewAttachmentResponse(attachment *model.Attachments) *AttachmentResponse {
	resp := AttachmentResponse{
		ID:       attachment.ID,
		Name:     attachment.Name,
		Category: attachment.Category,
	}
	return &resp
}

func NewAttachmentsListResponse(attachments *[]model.Attachments) []render.Renderer {
	list := []render.Renderer{}
	for _, attachment := range *attachments {
		list = append(list, NewAttachmentResponse(&attachment))
	}
	return list
}

func (rd AttachmentResponse) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	return nil
}

func NewListLoadoutsResponse(loadouts *[]model.Loadouts, totalCount int, hasNextPage bool) *ListLoadoutResponse {
	resp := &ListLoadoutResponse{
		Loadouts: NewLoadoutsListResponse(loadouts),
		Pagination: &common.PaginationResponse{
			TotalCount:  totalCount,
			HasNextPage: hasNextPage,
		},
	}
	return resp
}

func (rd ListLoadoutResponse) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	return nil
}
