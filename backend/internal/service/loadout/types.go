package loadout

import (
	"time"

	"github.com/google/uuid"
)

type LoadoutResponse struct {
	ID             uuid.UUID  `json:"id"`
	CreatedAt      time.Time  `json:"createdAt"`
	UpdatedAt      *time.Time `json:"updatedAt,omitempty"`
	Title          string     `json:"title"`
	Source         *string    `json:"source,omitempty"`
	SourceURL      *string    `json:"sourceUrl,omitempty"`
	WeaponName     string     `json:"weaponName"`
	WeaponCategory string     `json:"weaponCategory"`
	CreatedBy      uuid.UUID  `json:"createdBy"`
	GameID         uuid.UUID  `json:"gameId"`
	Attachments    *string    `json:"attachments,omitempty"`
}
