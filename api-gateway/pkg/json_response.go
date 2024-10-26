package pkg

import (
	"encoding/json"
	"net/http"
)

func WriteJSONResponse(w http.ResponseWriter, data interface{}, responseCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(responseCode)
	
	jsonData, err := json.Marshal(data)
	if err != nil {
		json.NewEncoder(w).Encode(&struct{
			Success		bool
			ErrMessage		string
		} {
			Success: false,
			ErrMessage: "unable to encode json data",
		})
	}

	w.Write(jsonData)
}