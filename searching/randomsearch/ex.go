package main
import(
    "fmt"
    "math/rand"
    "time"
)

func random(max int) int {
    rand.Seed(time.Now().Unix())
    return rand.Intn(max)
}

func main() {
    myrand := random(6)
    fmt.Println(myrand)
}
