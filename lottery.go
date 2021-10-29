package main

import (
    "fmt"
    "bufio"
	"os"
    "strings"
)

func contains(a []string, i string) bool {
    for _, item := range a {
        if item == i {
            return true
        }
    }
    return false
}

func processLine(line string) uint {
    var ret uint = 0
    items := strings.Split(line, " ")
    winner := strings.Split("76 73 19 88 78", " ")
    //winner := strings.Split("49 35 37 34 68", " ")
    for _, item := range items {
        if contains(winner, item) {
            ret += 1
        }
    }
    return ret
}

func main() {
    results := make(map[uint]int32)
    //path := "./test.txt"
    path := "./10m-v2.txt"
	file, _ := os.Open(path)
    defer file.Close()
	/* handle error */
    scanner := bufio.NewScanner(file)
	for scanner.Scan() {
        result := processLine(scanner.Text())
        if result > 1 {
            results[result] += 1
        }
	}
    fmt.Printf("%+v", results)
}
