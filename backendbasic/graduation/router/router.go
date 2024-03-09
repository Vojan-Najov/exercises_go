// Router allows you to match HTTP requests to handlers based on the request path
// It use a syntax similar to gorilla/mux:
// /path/{regexp}/{name:captured-regexp}
// AddRoute adds a route to the router.
// Vars returns the path parametrs for the current request, or nil if there are none.
//
//  var r Router
//  func echoHandler(w http.ResponseWriter, r *http.Request) {
//    vars := Vars(r.Context())
//    _ = json.NewEncoder(w).Encode(vars)
//  }
//  r.AddRoute(
//    "/chess/replay/{white:[a-zA-Z]+}/{black:[a-zA-Z]+}/{id:[0-9]+}",
//    myHandler,
//    "GET",
//  )
//  rec := httptest.NewRecorder()
//  ...

package router

import (
	"context"
	_ "encoding/json"
	"fmt"
	"net/http"
	_ "reflect"
	"regexp"
	"sort"
	"strings"
)

type Router struct {
	routes []route
}

type key[T any] struct{}

func CtxVal[T any](ctx context.Context) (T, bool) {
	t, ok := ctx.Value(key[T]{}).(T)
	return t, ok
}

func CtxWithVal[T any](ctx context.Context, t T) context.Context {
	return context.WithValue(ctx, key[T]{}, t)
}

type route struct {
	pattern *regexp.Regexp
	names   []string
	raw     string // the raw pattern string
	method  string // the HTTP method to match; if empty, all methods match.
	handler http.Handler
}

// Vars is a map of path parameters to their values. It is a unique type so that
// ctxutil.Value can be user to rerieve it.
type PathVars map[string]string

var empty = make(PathVars)

// Vars return the path parametrs for the current request, or nil if there are none.
func Vars(ctx context.Context) PathVars {
	v, _ := CtxVal[PathVars](ctx)
	return v
}

// suppose our input is /chess/replay/{white:[a-zA-Z]+}/{black:[a-z][A-Z]+/{id:[0-9]+}
// i.e, we choose the white and black player's names, and the game id.
// we'd like to to match /chess/replay/efronlicht/bobross/1234
// where white=efronlicht, black=bobross, and id=1234
// we will eventually compile this into a regexp that looks like:
//
// "^/chess/replay/(a-zA-Z]+/[a-zA-Z]+)/(0-9]+$"
//
// and a names slice that looks like:
//
// []string{"white", "black", "id"}
func buildRoute(pattern string) (re *regexp.Regexp, names []string, err error) {
	if pattern == "" || pattern[0] != '/' {
		return nil, nil, fmt.Errorf("invalid pattern %s: must begin with '/'", pattern)
	}
	var buf strings.Builder
	buf.WriteByte('^') // match the beginning of the string

	// we gradually build up the regexp, and keep track of the path parametrs we
	// encounter. e.g, on successive iterations, we'll have:
	// FOR {
	// 0: /chess, nil
	// 1: /chess/replay, nil
	// 2: /chess/replay/([a-zA-Z]+), [white]
	// 3: /chess/replay/([a-zA-Z]+)/([a-zA-Z]+), [white, black]
	// 4: /chess/replay/([a-zA-Z]+)/([a-zA-Z]+)/([0-9]+), [white, black, id]
	// }
	for _, f := range strings.Split(pattern, "/")[1:] {
		buf.WriteByte('/') // add the '/' back
		if len(f) >= 2 && f[0] == '{' && f[len(f)-1] == '}' {
			//path parameter
			trimmed := f[1 : len(f)-1] // strip off the '{' and '}'
			// - {white:[a-zA-Z]+} -> [a-zA-Z]+
			if before, after, ok := strings.Cut(trimmed, ":"); ok {
				// it is a regexp-capture group
				names = append(names, before)
				// replace with a capture group: i.e, if we have {id:[0-9]+},
				// we want to replace it with ([0-9]+)
				buf.WriteByte('(')
				buf.WriteString(after)
				buf.WriteByte(')')
				// white:[a-zA-Z]+ -> ([a-zA-Z]+)
			} else {
				// a regular expression, but not a captured one
				buf.WriteString(trimmed)
			}
		} else {
			buf.WriteString(regexp.QuoteMeta(f)) // escape any special characters
		}
	}
	// check for duplicate path parametrs
	for i := range names {
		for j := i + 1; j < len(names); j++ {
			if names[i] == names[j] {
				return nil, nil, fmt.Errorf(
					"duplicate path parameter %s in %q", names[i], pattern,
				)
			}
		}
	}
	buf.WriteByte('$') // match the end of the string
	re, err = regexp.Compile(buf.String())
	if err != nil {
		return nil, nil, fmt.Errorf("invalid regexp %s: %w", buf.String(), err)
	}
	return re, names, nil
}

// AddRoute adds a route to the router. Method is the HTTP method to match;
// if empty, all methods match
// Method will be converted to uppercase; "get", "gEt" and "GET" are all equivalent.
func (r *Router) AddRoute(pattern string, h http.Handler, method string) error {
	re, names, err := buildRoute(pattern)
	if err != nil {
		return err
	}
	r.routes = append(r.routes, route{
		raw:     pattern,
		pattern: re,
		names:   names,
		method:  strings.ToUpper(strings.TrimSpace(method)),
		handler: h,
	})

	// sort the routes by length, so that the longest routes are matched first.
	// sort by length, then lexicographically
	sort.Slice(r.routes, func(i, j int) bool {
		return len(r.routes[i].raw) > len(r.routes[j].raw) ||
			(len(r.routes[i].raw) == len(r.routes[j].raw) &&
				r.routes[i].raw < r.routes[j].raw)
	})
	return nil
}

// pathVars extracts the path paramerers from the path and into a map.
// --- perfomance designe note: ---
// this is pretty inefficient, since we're re-matching the regexp/
// we could instead store the regexp and the names in the route struct,
// just iterate through & check for matches.
// since most paths will have very few path parameters, this will perform
// better and avoid extra allocs.
// additionaly, we could store a small amount of storage for names directly in the
// route struct so as to avoid allocating a slice for each reques.
// even better, we could make a new API for *regexp.FindStingSubmatch that
// _appends_ to an existing slice rather than allocating a new one,
// using a sync.Pool or something to avoid allocations entirely.
// Still, the goal here is to match gorilla/mux's API as simple of an
// implemetation as possible, so we'll leave it as-is
func pathVars(re *regexp.Regexp, names []string, path string) PathVars {
	matches := re.FindStringSubmatch(path)
	if len(matches) != len(names)+1 {
		// +1 because the first match is the entire string
		panic(fmt.Errorf(
			"programmer error: expected regexp %q to match %q", path, re.String(),
		))
	}
	vars := make(PathVars, len(names))
	for i, match := range matches[1:] {
		// again, skip the first match, which is the entire string
		vars[names[i]] = match
	}
	return vars
}

// ServeHTTP implemenets http.Handler, dispatching requests to the appropriate handler
func (rt *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, route := range rt.routes {
		if route.pattern.MatchString(r.URL.Path) &&
			(route.method == "" || route.method == r.Method) {
			vars := pathVars(route.pattern, route.names, r.URL.Path)
			ctx := CtxWithVal(r.Context(), vars)
			route.handler.ServeHTTP(w, r.WithContext(ctx))
			return
		}
	}
	http.NotFound(w, r) // no route matched; serve a 404
}
