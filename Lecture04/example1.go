import (
	"fmt"
	"sync"
)

type MySyncMap struct {
	mu    sync.Mutex
	data  map[string]interface{}
}

func NewMySyncMap() *MySyncMap {
	return &MySyncMap{
		data: make(map[string]interface{}),
	}
}

func (m *MySyncMap) Load(key string) (interface{}, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	value, ok := m.data[key]
	return value, ok
}

func (m *MySyncMap) Store(key string, value interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data[key] = value
}

func (m *MySyncMap) Delete(key string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.data, key)
}

func main() {
	myMap := NewMySyncMap()

	myMap.Store("name", "Alice")
	myMap.Store("age", 30)

	name, ok := myMap.Load("name")
	if ok {
		fmt.Println("Name:", name)
	}

	age, ok := myMap.Load("age")
	if ok {
		fmt.Println("Age:", age)
	}

	myMap.Delete("age")

	_, ok = myMap.Load("age")
	if !ok {
		fmt.Println("Age is deleted")
	}
}