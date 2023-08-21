package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    in := bufio.NewReader(os.Stdin)
    out := bufio.NewWriter(os.Stdout)
    defer out.Flush()

    var t int
    fmt.Fscan(in, &t)

    for i := 0; i < t; i++ {

        list := make(map[int]int)
        var k, v int

        fmt.Fscan(in, &k)
        for j := 0; j < k; j++ {
            fmt.Fscan(in, &v)
            list[t] += v
        }

        var ans int
        for a, b := range list {
            ans += b - ((b/a)/3)*a
        }

        fmt.Fprintln(out, ans)
    }
}
