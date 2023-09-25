import (
	"fmt"
	"sync"
	"time"
)

var (
	data  map[string]string
	mutex sync.RWMutex
)

func init() {
	data = make(map[string]string)
}

func readData(key string) string {
	mutex.RLock()
	defer mutex.RUnlock()
	return data[key]
}

func writeData(key, value string) {
	mutex.Lock()
	defer mutex.Unlock()
	data[key] = value
}

func main() {
	go func() {
		for i := 0; i < 5; i++ {
			writeData(fmt.Sprintf("key%d", i), fmt.Sprintf("value%d", i))
			time.Sleep(100 * time.Millisecond)
		}
	}()

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println(readData(fmt.Sprintf("key%d", i)))
			time.Sleep(200 * time.Millisecond)
		}
	}()

	time.Sleep(2 * time.Second)
}