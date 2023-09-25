import (
    "fmt"
    "sync"
)

func main() {
    var mu sync.Mutex

    go func() {
        mu.Lock()
        // Не выполняется Unlock, что приводит к deadlock
    }()

    mu.Lock() // Здесь горутина заблокировала mu, и программа зависает
    fmt.Println("This will never be printed")
    mu.Unlock()
}