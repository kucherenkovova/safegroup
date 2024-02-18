# safegroup
[![Go Reference](https://pkg.go.dev/badge/github.com/kucherenkovova/safegroup.svg)](https://pkg.go.dev/github.com/kucherenkovova/safegroup)
[![Go Report Card](https://goreportcard.com/badge/github.com/kucherenkovova/safegroup)](https://goreportcard.com/report/github.com/kucherenkovova/safegroup)

Panic-safe drop-in replacement for [x/sync/errgroup](https://pkg.go.dev/golang.org/x/sync/errgroup).

Don't let your whole server crash because of a single goroutine.

### Installation
```bash
go get github.com/kucherenkovova/safegroup@v1.0.2
```

### Usage

```go
func handleStuff(w http.ResponseWriter, req *http.Request) {
	g, ctx := safegroup.WithContext(req.Context())

	g.Go(queryDB)
	g.Go(sendEmail)

	g.Go(func() error {
		panic("oops, buggy code")
		return nil
	})

	// Wait for all tasks to complete.
	if err := g.Wait(); err != nil {
		if errors.Is(err, safegroup.ErrPanic) {
			sendAlert(ctx, err)
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, "successfully done")
}
```
