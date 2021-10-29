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

func (gate Gate) Enter(id string) {
    gateInstance.mu.Lock()
    defer gateInstance.mu.Unlock()

    fmt.Println("person with id: " + id + " is at the gate")
    time.Sleep(100 * time.Millisecond)
    fmt.Println("person with id: " + id + " passed through")
}

type MultiPersonGate struct {

}

func (multiPersonGate MultiPersonGate) Enter(id string) {
    fmt.Println("    person with id: " + id + " is at the multi person gate")
    time.Sleep(100 * time.Millisecond)
    fmt.Println("    person with id: " + id + " passed through the multi person gate")
}

var gateInstance *Gate
var multiPersonGateInstance *MultiPersonGate

func init() {
    gateInstance = new(Gate)
    multiPersonGateInstance = new (MultiPersonGate)
}

func enter(n int) {
    gate := Gate{}
    gate.Enter(strconv.Itoa(n))
}

func enterMulti(n int) {
    multiPersonGate := MultiPersonGate{}
    multiPersonGate.Enter(strconv.Itoa(n))
}

func main() {
    for i := 0; i < 10; i++ {
        go enter(i)
        go enterMulti(i)
    }

    var input string
    fmt.Scanln(&input)
}
