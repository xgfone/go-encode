package encode

import (
	"errors"
	"fmt"

	"github.com/axgle/mahonia"
)

// Convert the content to the encoding of to from the encoding of from.
//
// The charset of content is 'from', and the result is 'to'.
func Convert(from, to, content string) (string, error) {
	dec := mahonia.NewDecoder(from)
	enc := mahonia.NewEncoder(to)
	if dec == nil || enc == nil {
		charset := from
		if enc == nil {
			charset = to
		}
		arg := fmt.Sprintf("Not support the charset: %v", charset)
		return "", errors.New(arg)
	}
	_dec := dec.ConvertString(content)
	_enc := enc.ConvertString(_dec)
	return _enc, nil
}

// Convert the content to the encoding of UTF8 from the encoding of from.
//
// The charset of content is 'from', and the result is UTF-8.
func ToUtf8(from, content string) (string, error) {
	dec := mahonia.NewDecoder(from)
	if dec == nil {
		arg := fmt.Sprintf("Not support the charset: %v", from)
		return "", errors.New(arg)
	}
	return dec.ConvertString(content), nil
}

// Convert the content to the encoding of to from the encoding of UTF-8.
//
// The charset of content is UTF-8, and the result is 'to'.
func FromUtf8(to, content string) (string, error) {
	enc := mahonia.NewEncoder(to)
	if enc == nil {
		arg := fmt.Sprintf("Not support the charset: %v", to)
		return "", errors.New(arg)
	}
	return enc.ConvertString(content), nil
}
