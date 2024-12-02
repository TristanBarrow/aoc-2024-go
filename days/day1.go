package days

import (
	"sort"
	"strconv"
	"strings"
)

func parseInt(str string) int {
    i, err := strconv.Atoi(str)
    if err != nil {
        panic(err)
    }
    return i
}

func abs(a, b int) int {
    c := a - b
    if c < 0 {
        return c * -1
    }
    return c
}

func makeLists(input []string) ([]int, []int) {
    left := make([]int, 0, 1000)
    right := make([]int, 0, 1000)

    for _, v := range input {
        if len(v) == 0 {
            continue
        }
        s := strings.Split(v, "   ")
        left = append(left, parseInt(s[0]))
        right = append(right, parseInt(s[1]))
    }

    return left, right
}

func Day1Part1(input []string) int {
    left, right := makeLists(input)

    sort.Ints(left)
    sort.Ints(right)

    var sum int

    for i, l := range left {
        sum += abs(l, right[i])
    }
    
    return sum
}

func Day1Part2(input []string) int {
    left, right := makeLists(input)

    m := make(map[int]int)

    for _, v := range right {
        _, ok := m[v]
        if ok {
            m[v] += 1
        } else {
            m[v] = 1
        }
    }
    
    var sum int
    for _, v := range left {
        sum += v * m[v]
    }

    return sum
}


