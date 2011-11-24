
## Colors ##
Colors is a simple golang package that provides basic functions on colorful outputing in terminal.
![Golang with colors](http://farm7.staticflickr.com/6051/6382022437_1f60b4130f.jpg)

## Usage ##
```go
package main

import (
        "github.com/wsxiaoys/colors"
)

func main() {
        colors.Red.Println("Hello world")
        colors.Printf("%CHello %Cworld", colors.Red, colors.Blue)
}
```

## TODO ##
* More pre-defined color schema
