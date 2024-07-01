package types

import (
	"backend/gen/postgres/public/model"
	"net/http"
)

type LoadoutStore interface {
	ListLoadouts() ([]model.Loadouts, error)
	ListLoadoutsByUser(userId string) ([]model.Loadouts, error)
}

type CreateLoadoutPayload struct {
	Message string `json:"message"`
}

// TODO: move to utils?
// TODO: repurpose for create payload
func (req *CreateLoadoutPayload) Bind(r *http.Request) error {
	// if req.Title == "" {
	// 	return errors.New("") // TODO:
	// }

	// if req.WeaponCategory == "" {
	// 	return errors.New("") // TODO:
	// }

	// if req.WeaponName == "" {
	// 	return errors.New("") // TODO:
	// }

	return nil
}
