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

    var t, n, current, prev int
    fmt.Fscan(in, &t)

    for q := 0; q < t; q++ {
        fmt.Fscan(in, &n)

        set := make(map[int]bool)
        invalid := false

        for i := 0; i < n; i++ {
            fmt.Fscan(in, &current)
            if i == 0 {
                prev = current
            }

            if _, exist := set[current]; !exist {
                set[current] = true
            } else {
                if prev != current {
                    invalid = true
                }
            }

            prev = current
        }

        if invalid {
            fmt.Fprintln(out, "NO")
        } else {
            fmt.Fprintln(out, "YES")
        }
    }
}
