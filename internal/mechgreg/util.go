package mechgreg

import (
	"github.com/BESTRobotics/registry/internal/models"
)

func patchUserSlice(in []models.User, insert bool, user models.User) []models.User {
	if (in == nil || len(in) == 0) && insert {
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
		out = append(out, models.User{ID: in[i].ID})
	}
	if insert && !done {
		out = append(out, models.User{ID: user.ID})
	}
	return out
}

func patchStudentSlice(in []models.Student, insert bool, user models.Student) []models.Student {
	if (in == nil || len(in) == 0) && insert {
		return []models.Student{user}
	}

	done := false
	var out []models.Student
	for i := range in {
		if in[i].ID == user.ID && !insert {
			continue
		} else if in[i].ID == user.ID && insert && !done {
			done = true
		}
		out = append(out, models.Student{ID: in[i].ID})
	}
	if insert && !done {
		out = append(out, models.Student{ID: user.ID})
	}
	return out
}
