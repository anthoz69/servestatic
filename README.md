# servestatic

[![Go Report Card](https://goreportcard.com/badge/github.com/StpDevelo/servestatic)](https://goreportcard.com/report/github.com/StpDevelo/servestatic)

Serve Static is Goland handle for handle static file. And if path is directory, function will return not found.

## How to use
```go
http.Handle("/-/", http.StripPrefix("/-", servestatic.New("assets")))
```

## Reference
[webstatic](https://github.com/acoshift/webstatic) by Acoshift


## License
MIT