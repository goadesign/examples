// Code generated by goa v3.15.1, DO NOT EDIT.
//
// HTTP request path constructors for the calc service.
//
// Command:
// $ goa gen goa.design/examples/basic/design

package server

import (
	"fmt"
)

// MultiplyCalcPath returns the URL path to the calc service multiply HTTP endpoint.
func MultiplyCalcPath(a int, b int) string {
	return fmt.Sprintf("/multiply/%v/%v", a, b)
}
