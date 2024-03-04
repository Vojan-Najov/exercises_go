package libhttp

import (
  "log"
  "strconv"
  "bytes"
  "encoding"
  "io"
  "errors"
  "strings"
  "fmt"
  "net/http"
)

// Header represents a HTTP header. A HTTP header is a key-value pair,
// separated by a colon (:);
// the key should be formatted in Title-Case.
// Use Request.AddHeader() or Response.AddHeader() to add headers to a request or
// response and guarantee title-casing of the key.

type Header struct {
  Key, Value string
}

// Request represents a HTTP 1.1 request

type Request struct {
  Method  string                      // e.g, GET, POST, PUT, DELETE
  Path    string                      // e.g, /index.html
  //Headers []struct{Key, Value string} // e.g, Host: eblog.fly.dev
  Headers    []Header   // e.g, Content-Type: text/html
  Body    string                      // e.g, <html><body><h1>Hi</h1></body></html>
}

type Response struct {
  StatusCode int                         // e.g, 200
//  Headers    []struct{Key, Value string} // e.g, Content-Type: text/html
  Headers    []Header   // e.g, Content-Type: text/html
  Body       string                      // e.g, <html><body><h1>Hi</h1></body></html>
}

func NewRequest(method, path, host, body string) (*Request, error) {
  switch {
  case method == "":
    return nil, errors.New("missing required argument: method")
  case path == "":
    return nil, errors.New("missing required argument: path")
  case !strings.HasPrefix(path, "/"):
    return nil, errors.New("path must start with /")
  case host == "":
    return nil, errors.New("missing required argument: host")
  default:
    headers := make([]Header, 2)
    headers[0] = Header{"Host", host}
    if body != "" {
      headers = append(
        headers,
        Header{"Content-Length", fmt.Sprintf("%d", len(body))},
      )
    }
    return &Request{Method: method, Path: path, Headers: headers, Body: body}, nil
  }
}

func NewResponse(status int, body string) (*Response, error) {
  switch {
  case status < 100 || status > 599:
    return nil, errors.New("invalid status code")
  default:
    if body == "" {
      body = http.StatusText(status)
    }
    headers := []Header{Header{"Content-Length", fmt.Sprintf("%d", len(body))}}
    return &Response{
      StatusCode: status,
      Headers: headers,
      Body: body,
    }, nil
  }
}

func (resp *Response) WithHeader(key, value string) *Response {
  resp.Headers = append(resp.Headers, Header{AsTitle(key), value})
  return resp
}

func (r *Request) WithHeader(key, value string) *Request {
  r.Headers = append(r.Headers, Header{AsTitle(key), value})
  return r
}

// AsTitle returns the given header key as title case;
// e.g. "content-type" -> "Content-Type"
// It will panic if the key is empty.

func AsTitle(key string) string {
  // design note --- an empty string could be considered 'in title case',
  // but in practice it's probably probably programmer error.
  // rather than guess, we'll panic
  if key == "" {
    panic("empty header key")
  }
  if isTitleCase(key) {
    return key
  }

  // allocation is very expensive, while iteration throught strings is very cheap
  return newTitleCase(key)
}

// newTitleCase returns the given header key as title case;
// e.g. "content-type" -> "Content-Type";
// it always allocate a new string.
func newTitleCase(key string) string {
  var b strings.Builder
  b.Grow(len(key))
  for i := range key {
    if i == 0 || key[i-1] == '-' {
      b.WriteByte(upper(key[i]))
    } else {
      b.WriteByte(lower(key[i]))
    }
  }
  return b.String()
}

// straight from K&R C, 2nd edition, page 43. Some classic never go out of style/
func lower(c byte) byte {
  if c >= 'A' && c <= 'Z' {
    return c + 'a' - 'A'
  }
  return c
}

func upper(c byte) byte {
  if c >= 'a' && c <= 'z' {
    return c + 'A' - 'a'
  }
  return c
}

// isTitleCase returns true if the given header key is already title case;
// i.e, it is of the form "Content-Type" or "Content-Length, "Some-Odd-Header", etc.
func isTitleCase(key string) bool {
  // check if this is already title case.
  for i := range key {
    if i == 0 || key[i-1] == '-' {
      if key[i] >= 'a' && key[i] <= 'z' {
        return false
      }
    } else if key[i] >= 'A' && key[i] <= 'Z' {
      return false
    }
  }
  return true
}

// Write writes the Request to the given io.Writer.
func (r *Request) WriteTo(w io.Writer) (n int64, err error) {
  // write & count bytes written.
  // using small closures like this to cut down on repetition
  // can be nice; byt you sometimes pay a perfomance penalty.
  printf := func(format string, args ...any) error {
    m, err := fmt.Fprintf(w, format, args...)
    n += int64(m)
    return err
  }
  // remember, a HTTP request looks like this:
  // <METHOD>  <PATH>  <PROTOCOL/VERSION>
  // <HEADER>: <VALUE>
  // <HEADER>: <VALUE>
  //
  // <REQUEST BODY>

  // write the request line: like "GET /index.html HTTP/1.1"
  if err := printf("%s %s HTTP/1.1\r\n", r.Method, r.Path); err != nil {
    return n, err
  }

  // write the headers. we don't do anything to order them or combine/merge
  // duplicate headers; this is just an example.
  for _, h := range r.Headers {
    if err := printf("%s: %s\r\n", h.Key, h.Value); err != nil {
      return n, err
    }
  }
  printf("\r\n") // write the empty line that separates the headers from the body
  err = printf("%s\r\n", r.Body) // write the body and terminate with a newline
  return n, err
}

func (resp *Response) WriteTo(w io.Writer) (n int64, err error) {
  printf := func(format string, args ...any) error {
    m, err := fmt.Fprintf(w, format, args...)
    n += int64(m)
    return err
  }
  err = printf("HTTP/1.1 %d %s\r\n", resp.StatusCode, http.StatusText(resp.StatusCode))
  if  err != nil {
    return n, err
  }
  for _, h := range resp.Headers {
    if err := printf("%s: %s\r\n", h.Key, h.Value); err != nil {
      return n, err
    }
  }
  err = printf("\r\n%s\r\n", resp.Body)
  return n, nil
}

// compile-time check that Request and Response implement fmt.Stringer
var _, _ fmt.Stringer = (*Request)(nil), (*Response)(nil)
var _, _ encoding.TextMarshaler = (*Request)(nil), (*Response)(nil)

func (r *Request) String() string {
  b := new(strings.Builder)
  r.WriteTo(b)
  return b.String()
}

func (resp *Response) String() string {
  b := new(strings.Builder)
  resp.WriteTo(b)
  return b.String()
}

func (r *Request) MarshalText() ([]byte, error) {
  b := new(bytes.Buffer)
  r.WriteTo(b)
  return b.Bytes(), nil
}

func (resp *Response) MarshalText() ([]byte, error) {
  b := new(bytes.Buffer)
  resp.WriteTo(b)
  return b.Bytes(), nil
}

// ParseRequest parses a HTTP request from the given text.
func ParseRequest(raw string) (r Request, err error) {
  // request has three parts:
  // 1. Request linedd
  // 2. Headers
  // 3. Body (optional)
  lines := splitLines(raw)

  log.Println(lines)
  if len(lines) < 3 {
    return Request{}, fmt.Errorf("malformed request: should have at least 3 lines")
  }
  // First line is special.
  first := strings.Fields(lines[0])
  r.Method, r.Path = first[0], first[1]
  if !strings.HasPrefix(r.Path, "/") {
    return Request{}, fmt.Errorf("malformed request: path should start with /")
  }
  if !strings.Contains(first[2], "HTTP") {
    return Request{}, fmt.Errorf(
      "malformed request: first line should contain HTTP version",
    )
  }

  var foundhost bool
  var bodyStart int
  //then we have headers, up until the an empty line.
  for i := 1; i < len(lines); i++ {
    if lines[i] == "" {
      bodyStart = i + 1
      break
    }
    key, val, ok := strings.Cut(lines[i], ": ")
    if !ok {
      return Request{}, fmt.Errorf(
        "malformed request: header %q should be of form 'key: value'", 
        lines[i],
      )
    }
    if key == "Host" {
      foundhost = true
    }
    key = AsTitle(key)
    r.Headers = append(r.Headers, Header{key, val})
  }
  // recombine the body using normal newlines; skip the last empty line.
  end := len(lines) - 1
  r.Body = strings.Join(lines[bodyStart:end], "\r\n")
  if !foundhost {
    return Request{}, fmt.Errorf("malformed request: missing Host header")
  }
  return r, nil
}

// ParseResponse parses the given HTTP/1.1 response string into the Response.
// It returns an error of the Response is invalid,
// - not a valid integet
// - missing status text
// - invalid status code
// - invalid headers
// it doesn't properly handle multi-line headers, headers with multiple values,
// or html-encoding, etc.zzs
func ParseResponse(raw string) (resp *Response, err error) {
  // response has three parts:
  // 1. Response line
  // 2. Headers
  // 3. Body (optional)
  lines := splitLines(raw)
  log.Println(lines)

  // First line is special
  first := strings.SplitN(lines[0], " ", 3)
  if !strings.Contains(first[0], "HTTP") {
    return nil, fmt.Errorf("malformed response: first line should contain HTTP version")
  }
  resp = new(Response)
  resp.StatusCode, err = strconv.Atoi(first[1])
  if err != nil {
    return nil, fmt.Errorf(
      "malformed response: expected status code to be an integer, got %q",
      first[1],
    )
  }
  if first[2] == "" || http.StatusText(resp.StatusCode) != first[2] {
    log.Printf(
      "missing or incorrect status text for status code %d: expected %q, but got %q",
      resp.StatusCode,
      http.StatusText(resp.StatusCode),
      first[2],
    )
  }

  var bodyStart int
  // then we have headers, up until the an empty line.
  for i := 1; i < len(lines); i++ {
    log.Println(i, lines[i])
    if lines[i] == "" {
      bodyStart = i + 1
      break
    }
    key, val, ok := strings.Cut(lines[i], ": ")
    if !ok {
      return nil, fmt.Errorf(
        "malformed response: header %q should be of form 'key: value'",
        lines[i],
      )
    }
    key = AsTitle(key)
    resp.Headers = append(resp.Headers, Header{key, val})
  }
  // recombine the body using normal newlines
  resp.Body = strings.TrimSpace(strings.Join(lines[bodyStart:], "\r\n"))
  return resp, nil
}

// splitLines on the "\r\n" sequence; multiple separators in a row are NOT collapsed.
func splitLines(s string) []string {
  if s == "" {
    return nil
  }
  var lines []string
  i := 0
  for {
    j := strings.Index(s[i:], "\r\n")
    if j == -1 {
      lines = append(lines, s[i:])
      return lines
    }
    lines = append(lines, s[i:i+j]) // up to but not including the \r\n
    i += j + 2  // skip the \r\n
  }
}















