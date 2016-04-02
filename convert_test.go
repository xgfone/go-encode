package encode_test

import (
	"testing"

	"github.com/xgfone/go-encode"
)

func TestEncodeConvert(t *testing.T) {
	gbk := string([]byte{'\xd6', '\xd0', '\xb9', '\xfa'})

	if utf8, err := encode.ToUtf8("gbk", gbk); err != nil || utf8 != "中国" {
		t.Fail()
	}

	if utf8, err := encode.Convert("gbk", "utf8", gbk); err != nil || utf8 != "中国" {
		t.Fail()
	}

	if _gbk, err := encode.FromUtf8("gbk", "中国"); err != nil || _gbk != gbk {
		t.Fail()
	}
}
