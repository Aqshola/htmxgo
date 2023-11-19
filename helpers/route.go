package helpers

import (
	"htmxgo/types"
	"net/http"
)

func GetField(r *http.Request, index int) string {
	fields := r.Context().Value(types.CtrKey{}).([]string)
	return fields[index]
}
