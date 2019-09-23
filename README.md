# LRU
Another thread-safe LRU Cache implemented in GO.


## Install
```
$ go get github.com/souravray/lru
``` 
## Sample Usage & Documentation

```go
package main

import (
  "github.com/souravray/lru"
  "fmt"
)

func main() {
  c, err := lru.New(20)
  if err != nil {
    panic(err)
  }
  c.Add("Key1", "ok")
  value, ok := c.Fetch("Key1")
  if !ok {
    panic("Cannot fetch value for Key1")
  }
  fmt.Println("Key1:", value)
  for i:=2; i<22; i++ {
    c.Add(i, "ok")
  }
  if c.Exist("Key1") {
    panic("Key1: Stale key")
  } else {
    fmt.Println("Key1: Evicted")
  }
}
```
```
  Key1: ok
  Key1: Evicted
```

Documentation is available on [Godoc](https://godoc.org/github.com/souravray/lru)


## Unit Test
Unit tests are avilable for list and base lru implementations. You can run unit test-cases by
```
$ go test -v -cover
``` 
Unit test code coverage is 70.2%, and thread safe methods doesn't have separate unit test-cases. You can also check the [test-coverage](http://raysourav.com/lru/cover.html) report.


## Benchmarking
At best-case scinarions the LRU perfoms 180 ns/op (in avarage). You can run the benchmark tests by
```
go test -bench=. -benchtime=1s
``` 

