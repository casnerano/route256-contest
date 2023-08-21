package main

import (
    "bufio"
    "fmt"
    "os"
    "sort"
)

func main() {
    in := bufio.NewReader(os.Stdin)
    out := bufio.NewWriter(os.Stdout)
    defer out.Flush()

    var n, m int

    fmt.Fscan(in, &n, &m)

    g := make(map[int][]int)
    ans := make(map[int][]int)

    var a, b int
    for i := 0; i < m; i++ {
        fmt.Fscan(in, &a, &b)
        g[a-1] = append(g[a-1], b-1)
        g[b-1] = append(g[b-1], a-1)
    }

    for index, friends := range g {
        f2List := make(map[int]int)
        for j := 0; j < len(friends); j++ {
            f2 := g[friends[j]]
            for k := 0; k < len(f2); k++ {
                if f2[k] != index {
                    f2List[f2[k]]++
                }
            }
        }

        for _, friend := range friends {
            if _, exist := f2List[friend]; exist {
                delete(f2List, friend)
            }
        }

        f2Max := -1
        for _, f2Count := range f2List {
            if f2Max < 0 || f2Count > f2Max {
                f2Max = f2Count
            }
        }

        for f2Index, f2Count := range f2List {
            if f2Max > -1 && f2Max == f2Count {
                ans[index] = append(ans[index], f2Index)
            }
        }
    }

    for i := 0; i < n; i++ {
        if _, exist := ans[i]; exist {
            sort.Sort(sort.IntSlice(ans[i]))
            for _, fValue := range ans[i] {
                fmt.Fprint(out, fValue+1, " ")
            }
        } else {
            fmt.Fprint(out, 0)
        }
        fmt.Fprintln(out)
    }
}
