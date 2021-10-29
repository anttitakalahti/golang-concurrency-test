package main

import (
    "fmt"
    "strconv"
    "sync"
    "time"
)

type Gate struct {
    mu sync.Mutex
}

var gateInstance *Gate

func init() {
    gateInstance = new(Gate)
}

func (gate Gate) Enter(id string) {
    gateInstance.mu.Lock()
    defer gateInstance.mu.Unlock()

    fmt.Println("person with id: " + id + " is at the gate")
    time.Sleep(100 * time.Millisecond)
    fmt.Println("person with id: " + id + " passed through")
}

func enter(n int) {
    gate := Gate{}
    gate.Enter(strconv.Itoa(n))
}

func main() {
    for i := 0; i < 10; i++ {
        go enter(i)
    }
    var input string
    fmt.Scanln(&input)
}
