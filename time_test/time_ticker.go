package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
    log.SetFlags(log.Lmicroseconds)
    ticker := time.NewTicker(time.Millisecond * 100)
    defer ticker.Stop()
		count := 0
		fmt.Printf("%v", ticker.C)
    for {
        select {
        case <-ticker.C:
            log.Printf("count=%d\n", count)
            count++
            doPeriodically()
        }
    }
}

func doPeriodically() {
    /* do something */
}
