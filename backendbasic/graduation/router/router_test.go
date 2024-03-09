package router

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

func TestRouterError(t *testing.T) {
	var r Router
	if err := r.AddRoute("", nil, ""); err == nil {
		t.Errorf("AddRoute(%q, %v, %q) returned nil, want error", "", nil, "")
	}
	if err := r.AddRoute("/{a:.+}/{a:.+}", nil, ""); err == nil {
		t.Errorf("AddRoute(%q, %v, %q) returned nil, want error", "/{a:.+}/{a:.+}", nil, "")
	}
}

func TestRouter(t *testing.T) {
	var r Router
	for _, route := range []struct {
		pattern, method string
		handler         http.HandlerFunc
	}{
		{"/", "GET", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, "Hello, world!\r\n")
		}},
		{"/echo/{a:.+}/{b:.+}/{c:.+}", "GET",
			func(w http.ResponseWriter, r *http.Request) {
				vars, _ := CtxVal[PathVars](r.Context())
				_ = json.NewEncoder(w).Encode(vars)
			}},
		{
			"/hello/{name:[a-zA-Z]+}", "GET", func(w http.ResponseWriter, r *http.Request) {
				vars, _ := CtxVal[PathVars](r.Context())
				fmt.Fprintf(w, "Hello, %s!\r\n", vars["name"])
			},
		},
	} {
		if err := r.AddRoute(route.pattern, route.handler, route.method); err != nil {
			t.Fatalf("AddRoute(%q, %v, %q) returned error: %v", route.pattern, route.handler, route.method, err)
		}
	}
	for _, tt := range []struct {
		path, method, want string
	}{
		{"/", "GET", "Hello, world!\r\n"},
		{"/hello/efron", "GET", "Hello, efron!\r\n"},
		{"/hello/efron", "POST", "404 page not found\n"},
		{"/hello/efron", "PUT", "404 page not found\n"},
		{"/echo/first/second/third", "GET", `{"a":"first","b":"second","c":"third"}` + "\n"},
	} {
		rec := httptest.NewRecorder()
		req := httptest.NewRequest(tt.method, tt.path, nil)
		r.ServeHTTP(rec, req)
		if got := rec.Body.String(); !strings.Contains(got, tt.want) {
			t.Errorf("r.ServeHTTP(%q, %q) returned %q, want %q", tt.method, tt.path, got, tt.want)
		}

	}

}

func TestRouteVars(t *testing.T) {
	for name, tt := range map[string]struct {
		pattern string
		path    string
		want    PathVars
		wantErr bool
	}{
		"no path params, no regexp": {
			pattern: "/chess/replay",
			path:    "/chess/replay",
		},
		"regexp, no path params": {
			pattern: "/rng/seed/{[0-9]+}",
			path:    "/rng/seed/1234",
		},
		"regexp w/ path params": {
			pattern: "/chess/replay/{white:[a-zA-Z]+}/{black:[a-zA-Z]+}/{id:[0-9]+}",
			path:    "/chess/replay/efronlicht/bobross/1234",
			want:    PathVars{"white": "efronlicht", "black": "bobross", "id": "1234"},
		},
		"bad regexp": {
			pattern: "/badregexp/{[-}",
			path:    "/badregexp/1234",
			wantErr: true,
		},
	} {
		t.Run(name, func(t *testing.T) {
			re, names, err := buildRoute(tt.pattern)
			if err != nil && tt.wantErr {
				return
			} else if err != nil {
				t.Fatalf("buildRoute(%q) returned error: %v", tt.pattern, err)
			}
			vars := pathVars(re, names, tt.path)
			if tt.want == nil {
				tt.want = make(PathVars)
			}
			if !reflect.DeepEqual(vars, tt.want) {
				t.Errorf("pathVars(%q, %q) returned %v, want %v", tt.pattern, tt.path, vars, tt.want)
			}

		})
		t.Run("no match", func(t *testing.T) {
			pattern := "/chess/replay/{white:[a-zA-Z]+}/{black:[a-zA-Z]+}/{id:[0-9]+}"
			path := "/chess/replay/efronlicht/bobross/aaa"
			re, names, err := buildRoute(pattern)
			if err != nil {
				t.Fatalf("buildRoute(%q) returned error: %v", pattern, err)
			}
			defer func() {
				if r := recover(); r == nil {
					t.Errorf("pathVars(%q, %q) did not panic, want panic", pattern, path)
				}
			}()
			_ = pathVars(re, names, path)
		})

	}
}
