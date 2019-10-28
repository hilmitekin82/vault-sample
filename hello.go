package main
import (
    "fmt"
    "os"
    "time"
)

func main() {
    fmt.Println("FOO:", os.Getenv("FOO"))
    time.Sleep(10000 * time.Second)
}