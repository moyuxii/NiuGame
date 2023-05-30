package tool

import (
	"NiuGame/main/Auth"
	"fmt"
	"reflect"
)

var (
	errorType          = reflect.TypeOf((*error)(nil)).Elem()
	CustomerClaimsType = reflect.TypeOf((*Auth.CustomerClaims)(nil)).Elem()
)

func indirectToClaimsOrError(a any) any {
	if a == nil {
		return nil
	}
	v := reflect.ValueOf(a)
	for !v.Type().Implements(CustomerClaimsType) && !v.Type().Implements(errorType) &&
		v.Kind() == reflect.Pointer && !v.IsNil() {
		v = v.Elem()
	}
	return v.Interface()
}

func ToClaims(T any) (Auth.CustomerClaims, error) {
	T = indirectToClaimsOrError(T)
	switch t := T.(type) {
	case Auth.CustomerClaims:
		return t, nil
	case nil:
		return Auth.CustomerClaims{}, fmt.Errorf("cust info is block")
	case error:
		return Auth.CustomerClaims{}, fmt.Errorf(t.Error())
	default:
		return Auth.CustomerClaims{}, fmt.Errorf("unable to cast %#v of type %T to string", t, t)
	}
}
