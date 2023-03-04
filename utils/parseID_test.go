package utils

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

func TestParsedID(t *testing.T) {
	cases := []struct {
		Name     string
		ID       string
		Expected error
	}{
		{
			Name:     "Empty",
			ID:       "",
			Expected: errors.New("send number id invalid"),
		},
		{
			Name:     "Invalid",
			ID:       "123a",
			Expected: errors.New("send number id invalid"),
		},
		{
			Name:     "Valid",
			ID:       "1",
			Expected: nil,
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {

			r := &http.Request{URL: &url.URL{}}
			r.URL.Query().Add("id", c.ID)

			_, err := ParseID(r)
			if err == c.Expected {
				t.Errorf("ParseID returned %v exepcted %v\n", err, c.Expected)

			} else {
				fmt.Printf("ParseID returned wrong status %v expected %v\n", err, c.Expected)
			}

		})
	}
}
