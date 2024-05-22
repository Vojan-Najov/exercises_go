r := chi.NewRouter()

r.Get("/articles/{date}-{slug}", getArticle)

func getArticle(w http.ResponseWriter, r *http.Request) {
	dateParam := chi.URLParam(r, "date")
	slugParam := chi.URLParam(r, "slug")
	article, err := database.GetArticle(date, slug)

	if err != nil {
		w.WriteHeader(422)
		w.Write(
			[]byte(
				fmt.Sprintf(
					"error fetching article %s-%s: %v",
					dataParam,
					slugParam,
					err,
				)
			)
		)
		return
	}

	if article == nil {
		w.WriteHeader(404)
		w.Write([]byte("article not found"))
		return
	}

	w.Write([]byte(article.Text()))
}

// as you can see above, the url parameters are defined using the curly
// brackets {} with the parameter name in between, as {date} and {slug}.

// When a HTTP request is sent to the server and handled by the chi router, if
// the URL path matches the format of /articles/{date}-{slug}, then the
// getArticle function will be called to send a response to the client.

// For instance, URL paths like /articles/20200109-this-is-so-cool will match
// the route, however, /articles/1 will not.

// We can also use regex in url patterns
// For Example:

r := chi.NewRouter()
r.Get("/articles/{rid:^[0-9]{5,6}}", getArticle)

