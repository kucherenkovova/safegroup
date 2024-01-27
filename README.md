# safegroup
[![Go Reference](https://pkg.go.dev/badge/github.com/kucherenkovova/safegroup.svg)](https://pkg.go.dev/github.com/kucherenkovova/safegroup)

Panic-safe drop-in replacement for [x/sync/errgroup](https://pkg.go.dev/golang.org/x/sync/errgroup).

### Installation
```bash
go get github.com/kucherenkovova/safegroup@v1.0.1
```

### Usage

```go
func handleStuff(w http.ResponseWriter, req *http.Request) {
	g, ctx := safegroup.WithContext(req.Context())

	g.Go(func() error {
		// query database
	})

	for _, api := range apis {
		api := api
		g.Go(func() error {
			// call api
		})
	}
	
	// Wait for all tasks to complete.
	if err := g.Wait(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, "successfully done")
}
```
