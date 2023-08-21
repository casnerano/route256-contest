package main

import (
    "bufio"
    "fmt"
    "os"
    "sort"
)

const maxN = 30

type table struct {
    cells               [maxN]*[maxN]int
    height              int
    width               int
    sortableColumnIndex int
}

func (t *table) setSortableColumn(index int) *table {
    t.sortableColumnIndex = index
    return t
}

func (t *table) Len() int {
    return t.height
}

func (t *table) Less(i, j int) bool {
    return t.cells[i][t.sortableColumnIndex] < t.cells[j][t.sortableColumnIndex]
}

func (t *table) Swap(i, j int) {
    t.cells[i], t.cells[j] = t.cells[j], t.cells[i]
}

func main() {
    in := bufio.NewReader(os.Stdin)
    out := bufio.NewWriter(os.Stdout)
    defer out.Flush()

    var t int
    fmt.Fscan(in, &t)

    var _table table

    for q := 0; q < t; q++ {
        fmt.Fscan(in, &_table.height, &_table.width)
        for i := 0; i < _table.height; i++ {
            if _table.cells[i] == nil {
                _table.cells[i] = new([maxN]int)
            }
            for j := 0; j < _table.width; j++ {
                fmt.Fscan(in, &_table.cells[i][j])
            }
        }

        var k, c int
        fmt.Fscan(in, &k)

        for i := 0; i < k; i++ {
            fmt.Fscan(in, &c)
            sort.Stable(_table.setSortableColumn(c - 1))
        }

        for i := 0; i < _table.height; i++ {
            for j := 0; j < _table.width; j++ {
                fmt.Fprint(out, _table.cells[i][j], " ")
            }
            fmt.Fprintln(out)
        }

        fmt.Fprintln(out)
    }
}
