package main

func main() {
    var port int
    flag.IntVar(&port, "port", 8080, "port for outgoing requests")
    flag.Parse()
    rt = clientmw.Trace(rt)
    rt = clientmw.Log(rt)
    client := &http.Client{
        Transport: rt,
        Timeout:   1 * time.Second,
    }
    req, err := http.NewRequest("GET", fmt.Sprintf("http://localhost:%d/panic", port), nil)
    if err != nil {
        log.Fatal(err)
    }

    resp, err := client.Do(req)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("GET /panic:")
    resp.Write(os.Stdout)

    var paths = []string{
        "/",
        "/echo/first/second/third",
        "/echo/first/second/third",
        "/echo/first/second/third",
        "/time",
        "/time",
        "/time",
        "/time",
    }

    var queries = []map[string]string{
        nil,
        nil,
        {"case": "upper"},
        {"case": "lower"},
        nil,
        {"format": time.RFC1123},
        {"format": time.RFC1123, "tz": "America/New_York"},
        {"format": time.RFC1123, "tz": "America/Los_Angeles"},
        {"format": time.RFC1123, "tz": "America/Chicago"},
    }
    for i := range paths {
        q := make(url.Values)
        for k, v := range queries[i] {
            q.Set(k, v)
        }
        url := fmt.Sprintf("http://localhost:%d%s", port, paths[i])
        if len(q) > 0 {
            url += "?" + q.Encode()
        }
        ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
        req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
        if err != nil {
            log.Fatal(err)
        }
        resp, err := client.Do(req)
        cancel()
        if err != nil {
            log.Fatal(err)
        }
        fmt.Printf("GET %s: \n", url)
        resp.Write(os.Stdout)
        fmt.Println("\n-------")
    }

}

