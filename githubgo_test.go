package githubgo

import (
	"testing"
)

func TestGetUser(t *testing.T) {
	username := "l3gacyb3ta"
	wantID := 58434499

	user, err := GetUser(username)

	if user.ID != int64(wantID) || err != nil {
		t.Fatalf(`GetUser("l3gacyb3ta") = %v, %v, want match for %#q, nil`, user, err, wantID)
	}
}
