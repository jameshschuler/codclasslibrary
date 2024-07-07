package loadout

import (
	"backend/gen/postgres/public/model"
	"time"

	"github.com/go-chi/render"
	"github.com/google/uuid"
)

type LoadoutResponse struct {
	ID             uuid.UUID         `json:"id"`
	CreatedAt      time.Time         `json:"createdAt"`
	UpdatedAt      *time.Time        `json:"updatedAt,omitempty"`
	Title          string            `json:"title"`
	Source         *string           `json:"source,omitempty"`
	SourceURL      *string           `json:"sourceUrl,omitempty"`
	WeaponName     string            `json:"weaponName"`
	WeaponCategory string            `json:"weaponCategory"`
	CreatedBy      uuid.UUID         `json:"createdBy"`
	GameID         uuid.UUID         `json:"gameId"`
	Attachments    []render.Renderer `json:"attachments"`
}

type AttachmentResponse struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Category string    `json:"category"`
}

type LoadoutDetail struct {
	model.Loadouts
	Attachments []model.Attachments
}
