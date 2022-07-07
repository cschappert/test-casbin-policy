package authz

import "github.com/casbin/casbin/v2"

func CheckPermissions(sub, obj, act string, e *casbin.Enforcer) (bool, error) {
	ok, err := e.Enforce(sub, obj, act)
	if err != nil {
		return false, err
	}
	return ok, nil
}
