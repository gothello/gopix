package utils

import (
	"errors"
	"fmt"
	"net/http"
	"testing"
)

func ParseIdTest(t *testing.T) {
	cases := []struct {
		Name     string
		ID       string
		Expected error
	}{
		{
			Name:     "Empty",
			ID:       "0",
			Expected: errors.New("send number id valid"),
		},
		{
			Name:     "Invalid",
			ID:       "123a",
			Expected: errors.New("send number id valid"),
		},
		{
			Name:     "Valid",
			ID:       "1",
			Expected: nil,
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {

			var r *http.Request
			r.URL.Query().Add("id", c.ID)

			_, err := ParseID(r)
			if err != c.Expected {
				t.Errorf("ParseID returned %v exepcted %v\n", err, c.Expected)
				return
			}

			fmt.Printf("ParseID returned %v expected %v\n", err, c.Expected)

		})
	}
}
