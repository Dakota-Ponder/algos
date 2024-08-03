package main

import (
	"fmt"
)

func main() {
	games := []string{"1:0","2:0","3:0","4:0","2:1","3:1","4:1","3:2","4:2","4:3"}
    fmt.Println(Points(games))
}
