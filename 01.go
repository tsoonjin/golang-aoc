package main

import (
    "fmt"
    "os"
    "log"
    "strings"
    "strconv"

)

func calcWeight(weight int) int {
    return weight / 3 - 2
}

func calcWeightRecur(weight int) int {
    remaining := calcWeight(weight)
    if remaining <= 0 {
        return 0
    }
    return remaining + calcWeightRecur(remaining)
}

func main() {
    content, err := os.ReadFile("01.txt")

    if err != nil {
        log.Fatal(err)
    }
    lines := strings.Split(string(content), "\n")

    var sum1 = 0;
    var sum2 = 0;
    for _, s := range lines {
        if s != "" {
            weight, err := strconv.Atoi(s)
            if err != nil {
                log.Fatal("Failed to convert");
            }
            sum1 += calcWeight(weight)
            sum2 += calcWeightRecur(weight)
        }
    }
    fmt.Println(fmt.Sprintf("Sum1: %v, Sum2: %v", sum1, sum2))
}
