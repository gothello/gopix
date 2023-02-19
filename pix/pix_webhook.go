package pix

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func WebHookNotify() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodPost {

			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var input map[string]interface{}

		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		fmt.Println(input)
	}
}
