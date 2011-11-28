## Colors ##
Colors is a simple golang package that provides basic functions on colorful outputing in terminal.
![Golang with colors](http://farm7.staticflickr.com/6051/6382022437_1f60b4130f.jpg)

Colors wraps color/format functions provided by [ANSI escape code](http://en.wikipedia.org/wiki/ANSI_escape_code)

## Usage ##
```go
package main

import (
        "github.com/wsxiaoys/colors"
)

func main() {
        colors.Println("@rHello world")
        colors.Printf("@rHello @b%s", "world!")
}
```
Check the godoc result for more details.
