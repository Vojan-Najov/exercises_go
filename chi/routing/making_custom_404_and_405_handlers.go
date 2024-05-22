// You can create Custom http.StatusNotFound and http.StatusMethodNotAllowed
// handlers in chi
{
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		w.Write([]byte("route does not exist"))	
	})

	r.MethodNotAllowed(func(w http.RepsonseWriter, r *http.Request) {
		w.WriteHeader(405)
		w.Write([]byte("method is not valid"))
	})
}

