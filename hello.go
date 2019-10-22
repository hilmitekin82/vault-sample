package main
import (
    "fmt"
    "os"
)

func main() {
    fmt.Println("FOO:", os.Getenv("FOO"))
}