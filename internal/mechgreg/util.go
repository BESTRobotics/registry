package mechgreg

import (
	"github.com/BESTRobotics/registry/internal/models"
)

func patchUserSlice(in []models.User, insert bool, user models.User) []models.User {
	if in == nil && insert {
		return []models.User{user}
	}

	done := false
	var out []models.User
	for i := range in {
		if in[i].ID == user.ID && !insert {
			continue
		} else if in[i].ID == user.ID && insert && !done {
			done = true
		}
		out = append(out, models.User{ID: user.ID})
	}
	return out
}
