package main

import (
    "bufio"
    "container/heap"
    "fmt"
    "os"
)

type HeapItem struct {
    Priority int
    Payload  int
}
type MinHeap []*HeapItem

func (mh *MinHeap) Len() int {
    return len(*mh)
}
func (mh *MinHeap) Less(i, j int) bool {
    return (*mh)[i].Priority < (*mh)[j].Priority
}
func (mh *MinHeap) Swap(i, j int) {
    (*mh)[i], (*mh)[j] = (*mh)[j], (*mh)[i]
}

func (mh *MinHeap) Push(x interface{}) {
    *mh = append(*mh, x.(*HeapItem))
}

func (mh *MinHeap) Pop() interface{} {
    old := *mh
    n := len(old)
    x := old[n-1]
    *mh = old[0 : n-1]
    return x
}

func (mh *MinHeap) Empty() bool {
    return len(*mh) == 0
}

func (mh *MinHeap) Peak() *HeapItem {
    return (*mh)[0]
}

func main() {
    in := bufio.NewReader(os.Stdin)
    out := bufio.NewWriter(os.Stdout)
    defer out.Flush()

    var n, m, t int

    fmt.Fscan(in, &n, &m)

    idleProcessors := &MinHeap{}
    busyProcessors := &MinHeap{}

    heap.Init(idleProcessors)
    heap.Init(busyProcessors)

    for i := 0; i < n; i++ {
        fmt.Fscan(in, &t)
        heap.Push(idleProcessors, &HeapItem{
            Priority: t,
            Payload:  t,
        })
    }

    var start, length int
    var powerSpent int

    for i := 0; i < m; i++ {
        fmt.Fscan(in, &start, &length)

        for !busyProcessors.Empty() {
            item := busyProcessors.Peak()
            if item.Priority <= start {
                item = heap.Pop(busyProcessors).(*HeapItem)
                item.Priority = item.Payload
                heap.Push(idleProcessors, item)
            } else {
                break
            }
        }

        if !idleProcessors.Empty() {
            item := heap.Pop(idleProcessors).(*HeapItem)
            powerSpent += item.Payload * length
            item.Priority = start + length
            heap.Push(busyProcessors, item)
        }
    }

    fmt.Fprintln(out, powerSpent)
}
