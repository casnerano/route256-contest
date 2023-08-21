package main

import (
    "bufio"
    "fmt"
    "math"
    "os"
)

type min struct {
    value float64
    index int
}

func main() {
    in := bufio.NewReader(os.Stdin)
    out := bufio.NewWriter(os.Stdout)
    defer out.Flush()

    var t, n int
    var p [50]int

    fmt.Fscan(in, &t)

    for i := 0; i < t; i++ {
        fmt.Fscan(in, &n)

        for j := 0; j < n; j++ {
            fmt.Fscan(in, &p[j])
        }

        for j := 0; j < n-1; j++ {
            if p[j] == 0 {
                continue
            }

            var _min *min

            for k := j + 1; k < n; k++ {
                if p[k] == 0 {
                    continue
                }

                if _min == nil {
                    _min = &min{math.Abs(float64(p[j] - p[k])), k}
                }

                m := math.Abs(float64(p[j] - p[k]))
                if _min.value > m {
                    _min.value = m
                    _min.index = k
                }
            }

            if _min != nil {
                p[j] = 0
                p[_min.index] = 0

                fmt.Fprintf(out, "%d %d\n", j+1, _min.index+1)
            }
        }

        fmt.Fprintln(out)
    }
}
