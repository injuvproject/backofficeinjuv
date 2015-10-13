package binding

import "net/http"

type Binding map[string]interface{}

func GetDefault(r *http.Request) Binding {
	return Binding{
		"PageTitle":  "Back office INJUV",
		"CurrentURL": r.URL.Path,
	}
}
