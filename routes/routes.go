package routes

import (
	"context"
	"htmxgo/handlers"
	"htmxgo/types"
	"net/http"
	"regexp"
	"strings"
)

// ADD ROUTES HERE
var Routes = []types.Route{
	newRoute("GET", "/", handlers.Index),
	newRoute("POST", "/add-todo/", handlers.AddTodo),
	newRoute("DELETE", "/remove-todo/([^/]+)", handlers.RemoveTodo),
	newRoute("PUT", "/update-todo/([^/]+)", handlers.UpdateTodo),
}

func newRoute(Method, Pattern string, Handler http.HandlerFunc) types.Route {
	return types.Route{Method, regexp.MustCompile("^" + Pattern + "$"), Handler}
}

func ServeRoute(w http.ResponseWriter, r *http.Request) {
	var allow []string

	for _, route := range Routes {
		matches := route.Regex.FindStringSubmatch(r.URL.Path)
		if len(matches) > 0 {
			if r.Method != route.Method {
				allow = append(allow, route.Method)
				continue
			}
			ctx := context.WithValue(r.Context(), types.CtrKey{}, matches[1:])
			route.Handler(w, r.WithContext(ctx))
			return
		}
	}
	if len(allow) > 0 {
		w.Header().Set("Allow", strings.Join(allow, ", "))
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.NotFound(w, r)
}
