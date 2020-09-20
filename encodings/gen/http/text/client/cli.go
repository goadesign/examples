// Code generated by goa v2.2.4, DO NOT EDIT.
//
// text HTTP client CLI support package
//
// Command:
// $ goa gen goa.design/examples/encodings/design -o
// $(GOPATH)/src/goa.design/examples/encodings

package client

import (
	text "goa.design/examples/encodings/gen/text"
)

// BuildConcatstringsPayload builds the payload for the text concatstrings
// endpoint from CLI flags.
func BuildConcatstringsPayload(textConcatstringsA string, textConcatstringsB string) (*text.ConcatstringsPayload, error) {
	var a string
	{
		a = textConcatstringsA
	}
	var b string
	{
		b = textConcatstringsB
	}
	v := &text.ConcatstringsPayload{}
	v.A = a
	v.B = b

	return v, nil
}

// BuildConcatbytesPayload builds the payload for the text concatbytes endpoint
// from CLI flags.
func BuildConcatbytesPayload(textConcatbytesA string, textConcatbytesB string) (*text.ConcatbytesPayload, error) {
	var a string
	{
		a = textConcatbytesA
	}
	var b string
	{
		b = textConcatbytesB
	}
	v := &text.ConcatbytesPayload{}
	v.A = a
	v.B = b

	return v, nil
}

// BuildConcatstringfieldPayload builds the payload for the text
// concatstringfield endpoint from CLI flags.
func BuildConcatstringfieldPayload(textConcatstringfieldA string, textConcatstringfieldB string) (*text.ConcatstringfieldPayload, error) {
	var a string
	{
		a = textConcatstringfieldA
	}
	var b string
	{
		b = textConcatstringfieldB
	}
	v := &text.ConcatstringfieldPayload{}
	v.A = a
	v.B = b

	return v, nil
}

// BuildConcatbytesfieldPayload builds the payload for the text
// concatbytesfield endpoint from CLI flags.
func BuildConcatbytesfieldPayload(textConcatbytesfieldA string, textConcatbytesfieldB string) (*text.ConcatbytesfieldPayload, error) {
	var a string
	{
		a = textConcatbytesfieldA
	}
	var b string
	{
		b = textConcatbytesfieldB
	}
	v := &text.ConcatbytesfieldPayload{}
	v.A = a
	v.B = b

	return v, nil
}
