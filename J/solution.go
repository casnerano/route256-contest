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

    var n, q int
    var st string

    fmt.Fscan(in, &n)

    dict := make([]string, 0, n)
    index := map[string][]int{}

    for i := 0; i < n; i++ {
        fmt.Fscan(in, &st)
        dict = append(dict, st)

        for j := 0; j < len(st); j++ {
            index[st[j:]] = append(index[st[j:]], i)
        }
    }

    fmt.Fscan(in, &q)

    for i := 0; i < q; i++ {
        fmt.Fscan(in, &st)

        wordIndex := -1
        for j := 0; j < len(st); j++ {
            if v, ok := index[st[j:]]; ok {
                x := 0
                found := false
                for x < len(v) {
                    if dict[v[x]] != st {
                        found = true
                        break
                    }
                    x++
                }

                if !found {
                    continue
                }

                wordIndex = v[x]
                break
            }
        }

        if wordIndex < 0 {
            wordIndex = 0
            for wordIndex < len(dict) && dict[wordIndex] == st {
                wordIndex++
            }
        }

        fmt.Fprintln(out, dict[wordIndex])
    }
}
