package utils

import (
	"net/http"
	"strings"
)

func GetAndTrim(r *http.Request, key string) string {
	return strings.TrimSpace(r.FormValue(key))
}
