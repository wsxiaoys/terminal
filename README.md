## Colors ##
Colors is a simple golang package that provides basic functions on colorful outputing in terminal.

## Usage ##
        package main

        import (
                "github.com/wsxiaoys/colors"
        )

        func main() {
                colors.Red.Println("Hello world")
                colors.Printf("%CHello %Cworld", colors.Red, colors.Blue)
        }

## TODO ##
* More pre-defined color schema
