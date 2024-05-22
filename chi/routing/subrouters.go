
func main() {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!\n"))
	})

	// Creating a new router
	apiRouter := chi.NewRouter()
	apiRouter.Get("/articles/{date}-{slug}", getArticle)

	// Mounting the new Sub Router on the main router
	r.Mount("/api", apiRouter)
}

// Another way of implementing sub router would be:

{
	r.Route("/articles", func(r chi.Router) {
		// Get /articles
		r.With(paginate).Get("/". listArticles)

		// Get /articles/01-16-2017
		r.With(paginate).Get(
			"/{month}-{day}-{year}",
			listArticlesByDate,
		)

		// POST /articles
		r.Post("/", createArticle)

		// Regexp url parameters
		// GET /articles/home-is-toronto
		r.Get("/{articleSlug:[a-z-]+}", getArticleBySlug)

		// Subrouters:
		r.Route("/{articleID}", func(r chi.Router) {
			r.Use(ArticleCtx)

			// GET /articles/123
			r.Get("/", getArticle)

			// PUT /articles/123
			r.Put("/", updateArticle)

			// Delete /articles/123
			r.Delete("/", deleteArticle) 
		})
	})
}
