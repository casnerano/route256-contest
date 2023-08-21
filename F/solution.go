package main

import (
    "bufio"
    "errors"
    "fmt"
    "os"
    "sort"
    "strings"
    "time"
)

const timeFormat = "15:04:05"

type timeRange struct {
    start time.Time
    end   time.Time
}

func parseTime(strDate string) (tRange timeRange, err error) {
    stRes := strings.Split(strDate, "-")

    if tRange.start, err = time.Parse(timeFormat, stRes[0]); err != nil {
        return
    }

    if tRange.end, err = time.Parse(timeFormat, stRes[1]); err != nil {
        return
    }

    if tRange.start.After(tRange.end) {
        err = errors.New("invalid time range")
        return
    }

    return
}

func checkTimeRanges(timeRanges []timeRange) bool {
    for i := 0; i < len(timeRanges)-1; i++ {
        if !timeRanges[i].end.Before(timeRanges[i+1].start) {
            return false
        }
    }
    return true
}

func main() {
    in := bufio.NewReader(os.Stdin)
    out := bufio.NewWriter(os.Stdout)
    defer out.Flush()

    var t, n int
    var tRange timeRange
    var st string
    var err error

    fmt.Fscan(in, &t)

    for q := 0; q < t; q++ {
        fmt.Fscan(in, &n)
        var invalid bool

        timeRanges := make([]timeRange, 0, n)

        for i := 0; i < n; i++ {
            fmt.Fscan(in, &st)

            tRange, err = parseTime(st)
            if err != nil {
                invalid = true
            }

            timeRanges = append(timeRanges, tRange)
        }

        if !invalid {
            sort.Slice(timeRanges, func(i, j int) bool {
                return timeRanges[i].start.Before(timeRanges[j].start)
            })

            if !invalid && !checkTimeRanges(timeRanges) {
                invalid = true
            }
        }

        if invalid {
            fmt.Fprintln(out, "NO")
        } else {
            fmt.Fprintln(out, "YES")
        }
    }
}
