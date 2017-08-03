package encode

import (
	"errors"
	"fmt"
	"io"

	"github.com/axgle/mahonia"
)

// Convert converts the content to the encoding of to from the encoding of from.
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

// ToUtf8 converts the content to the encoding of UTF8 from the encoding of from.
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

// ReaderToUtf8 is the same as ToUtf8, but reading the content from io.Reader.
func ReaderToUtf8(from string, r io.Reader) (io.Reader, error) {
	dec := mahonia.NewDecoder(from)
	if dec == nil {
		arg := fmt.Sprintf("Not support the charset: %v", from)
		return nil, errors.New(arg)
	}

	return dec.NewReader(r), nil
}

// FromUtf8 converts the content to the encoding of to from the encoding of UTF-8.
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
