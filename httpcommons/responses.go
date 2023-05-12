package httpcommons

import (
	"encoding/json"
	"net/http"
)

func Respond[T any](w http.ResponseWriter, content T, code int) {
	w.Header().Add("Content-Type", "application/json")
	jsonBytes, err := json.Marshal(content)
	if err != nil {
		InternalServerError(w, err)
		return
	}
	w.WriteHeader(code)
	w.Write(jsonBytes)
}

func InternalServerError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(err.Error()))
}
