# loopdefer

[![pkg.go.dev][gopkg-badge]][gopkg]

`loopdefer` finds using defer in a loop.

```go
func f() error {
	for i := 0; i < 10; i++ {
		defer println("miss")
	}
}
```

<!-- links -->
[gopkg]: https://pkg.go.dev/github.com/gostaticanalysis/loopdefer
[gopkg-badge]: https://pkg.go.dev/badge/github.com/gostaticanalysis/loopdefer?status.svg
