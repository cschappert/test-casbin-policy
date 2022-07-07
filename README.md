# Test Casbin Policy
A **very** simple framework for testing a Casbin authorization policy with `go test`

# Try it out
1. Replace `casbin/policy.csv` with your own policy file.
2. Replace `casbin/model.conf` with your own model.
3. Edit `authz_test.go`, adding your test cases to `var tests []test`
```go
	var tests = []test{
		{"admin", "file", "write", true},
		{"guest", "file", "write", false},
		// add additional test cases here
	}
```
4. In the project root, run `go test`