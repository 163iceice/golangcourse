import (
    "fmt"
)

func mergeChannels(channels ...chan int) <-chan int {
    merged := make(chan int)

    for _, ch := range channels {
        go func(c chan int) {
            for val := range c {
                merged <- val
            }
        }(ch)
    }

    return merged
}

func main() {
    ch1 := make(chan int)
    ch2 := make(chan int)

    go func() {
        ch1 <- 1
        ch1 <- 2
        close(ch1)
    }()

    go func() {
        ch2 <- 3
        ch2 <- 4
        close(ch2)
    }()

    merged := mergeChannels(ch1, ch2)

    for val := range merged {
        fmt.Println(val)
    }
}