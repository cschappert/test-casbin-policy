package authz

import (
	"fmt"
	"testing"

	"github.com/casbin/casbin/v2"
)

func TestCheckPermissions(t *testing.T) {
	e, err := casbin.NewEnforcer("casbin/model.conf", "casbin/policy.csv")
	if err != nil {
		t.Fatal("Unable to initialize enforcer.", err)
	}

	err = e.LoadPolicy()
	if err != nil {
		t.Fatal("Unable to load policy.", err)
	}

	type test struct {
		sub      string
		obj      string
		act      string
		expected bool
	}

	var tests = []test{
		{"admin", "file", "write", true},
		{"guest", "file", "write", false},
		// add additional test cases here
	}

	for _, test := range tests {
		ok, err := CheckPermissions(test.sub, test.obj, test.act, e)
		if err != nil {
			t.Fatal("CheckPermissions returned an error.", err)
		}
		if ok != test.expected {
			testCase := fmt.Sprintf("%s can %s %s", test.sub, test.act, test.obj)
			t.Errorf("%s -- Output %t not equal to expected %t", testCase, ok, test.expected)
		}
	}
}
