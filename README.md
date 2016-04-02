# go-encode
Convert the string between the charsets.

## Charsets
It supports the common charset encoding. See [Charset](./charsets.md).

**Notice:**

    You need using the alias of the charset.

## Installation
```shell
$ go get -u github.com/xgfone/go-encode
```

## Example
```go
package main

import (
    "fmt"

    "github.com/xgfone/go-encode"
)

func main() {
    gbk := string([]byte{'\xd6', '\xd0', '\xb9', '\xfa'})

    if utf8, err := encode.ToUtf8("gbk", gbk); err == nil {
        fmt.Printf("%v\n", utf8) // 中国
    }

    if utf8, err := encode.Convert("gbk", "utf8", gbk); err == nil {
        fmt.Printf("%v\n", utf8) // 中国
    }

    if _gbk, err := encode.FromUtf8("gbk", "中国"); err == nil {
        fmt.Printf("%x\n", _gbk) // d6d0b9fa
    }
}
```