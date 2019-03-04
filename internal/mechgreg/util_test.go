package mechgreg

import (
	"sort"
	"testing"

	"github.com/BESTRobotics/registry/internal/models"
)

func TestPatchUserSlice(t *testing.T) {
	cases := []struct {
		in     []models.User
		insert bool
		user   models.User
		want   []models.User
	}{
		{
			nil,
			true,
			models.User{ID: 1},
			[]models.User{
				models.User{ID: 1},
			},
		},
		{
			[]models.User{},
			true,
			models.User{ID: 1},
			[]models.User{
				models.User{ID: 1},
			},
		},
		{
			[]models.User{
				models.User{ID: 1},
				models.User{ID: 2},
				models.User{ID: 3},
			},
			true,
			models.User{ID: 1},
			[]models.User{
				models.User{ID: 1},
				models.User{ID: 2},
				models.User{ID: 3},
			},
		},
		{
			[]models.User{
				models.User{ID: 2},
				models.User{ID: 3},
			},
			true,
			models.User{ID: 1},
			[]models.User{
				models.User{ID: 1},
				models.User{ID: 2},
				models.User{ID: 3},
			},
		},
		{
			[]models.User{
				models.User{ID: 1},
				models.User{ID: 2},
				models.User{ID: 3},
			},
			false,
			models.User{ID: 1},
			[]models.User{
				models.User{ID: 2},
				models.User{ID: 3},
			},
		},
	}

	for i, c := range cases {
		got := patchUserSlice(c.in, c.insert, c.user)
		if !slicesAreEqual(got, c.want) {
			t.Errorf("Case %d Got: %v; Want: %v", i, got, c.want)
		}
	}
}

func slicesAreEqual(left, right []models.User) bool {
	if len(left) != len(right) {
		return false
	}

	sort.Slice(left, func(i, j int) bool {
		return left[i].ID < left[j].ID
	})

	sort.Slice(right, func(i, j int) bool {
		return right[i].ID < right[j].ID
	})

	for i, v := range left {
		if v.ID != right[i].ID {
			return false
		}
	}
	return true
}

func TestSlicesAreEqual(t *testing.T) {
	cases := []struct {
		left  []models.User
		right []models.User
		want  bool
	}{
		{
			[]models.User{
				models.User{ID: 1},
				models.User{ID: 2},
			},
			[]models.User{
				models.User{ID: 1},
			},
			false,
		},
		{
			[]models.User{
				models.User{ID: 1},
				models.User{ID: 2},
				models.User{ID: 3},
			},
			[]models.User{
				models.User{ID: 1},
				models.User{ID: 2},
				models.User{ID: 3},
			},
			true,
		},
		{
			[]models.User{
				models.User{ID: 1},
				models.User{ID: 2},
				models.User{ID: 3},
			},
			[]models.User{
				models.User{ID: 2},
				models.User{ID: 1},
				models.User{ID: 3},
			},
			true,
		},
	}

	for i, c := range cases {
		if got := slicesAreEqual(c.left, c.right); got != c.want {
			t.Errorf("%d: Got %v; Want %v", i, got, c.want)
		}
	}
}
