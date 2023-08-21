package main

import (
    "bufio"
    "fmt"
    "os"
)

const MAX_N = 20
const POINTS_COUNT = 6

type Point struct {
    x, y int
}

var Points = [POINTS_COUNT]Point{
    {2, 0}, {-2, 0},
    {1, 1}, {-1, 1},
    {-1, -1}, {1, -1},
}

type Board [MAX_N][MAX_N]byte

func main() {

    in := bufio.NewReader(os.Stdin)
    out := bufio.NewWriter(os.Stdout)
    defer out.Flush()

    var t, n, m int
    var board Board

    fmt.Fscan(in, &t)

    for q := 0; q < t; q++ {
        index := make(map[byte]bool)

        fmt.Fscan(in, &n, &m)

        for i := 0; i < n; i++ {
            for j := 0; j < m; j++ {
                var c byte
                fmt.Fscanf(in, "%c", &c)
                if c == '.' || (c >= 'A' && c <= 'Z') {
                    board[i][j] = c
                } else {
                    j--
                    continue
                }
            }
        }

        incorrect := false

        for i := 0; i < n; i++ {
            if incorrect {
                break
            }
            for j := 0; j < m; j++ {
                c := board[i][j]
                if c != '.' && c != '#' {
                    if _, exist := index[c]; !exist {

                        var wave func(int, int)
                        wave = func(a int, b int) {
                            board[a][b] = '#'
                            for _, point := range Points {
                                if a+point.y < n && b+point.x < m && a+point.y >= 0 && b+point.x >= 0 {
                                    if board[a+point.y][b+point.x] == c {
                                        wave(a+point.y, b+point.x)
                                    }
                                }
                            }
                        }

                        index[c] = true
                        wave(i, j)

                    } else {
                        incorrect = true
                        break
                    }
                }
            }
        }

        if incorrect {
            fmt.Fprintln(out, "NO")
        } else {
            fmt.Fprintln(out, "YES")
        }

    }
}
