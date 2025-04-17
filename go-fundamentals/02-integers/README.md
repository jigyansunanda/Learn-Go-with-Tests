## 🧪 Test Behaviour

-   By default, `go test` hides all output unless the test fails or you run in verbose mode.
-   So even though you're calling `fmt.Printf("✅ %d/%d test cases passed\n", passedTests, totalTests)` inside `t.Cleanup()`, you won’t see that output unless you explicitly tell go test to show it.

## ✅ Solution: Use the `-v` (verbose) flag

```go
go test -v ./go-fundamentals/02-integers
```

## 🔁 Optional: Disable caching if you're testing output

If you’re not seeing changes, you can force re-running tests (avoiding cache) with:

```go
go test -v -count=1 ./go-fundamentals/02-integers
```

### 📝 TL;DR

-   ✅ Use -v to see print output from tests
-   ✅ Use -count=1 to force re-running and avoid the cache
