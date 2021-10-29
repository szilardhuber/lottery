package main

import (
    "fmt"
    "log"
    "bufio"
	"os"
    "strings"
    "time"
    "strconv"
)

func contains(a []string, i string) bool {
    for _, item := range a {
        if item == i {
            return true
        }
    }
    return false
}

func containsUint(a []uint, i uint) bool {
    for _, item := range a {
        if item == i {
            return true
        }
    }
    return false
}

func split(line string) []uint {
    items := strings.Split(line, " ")
    ret := make([]uint, len(items))
    for i, v := range items {
        u64, _ := strconv.ParseUint(v, 10, 16)
        ret[i] = uint(u64)
    }
    return ret
}

func processLine(line string, winningNumbers []string) uint {
    var ret uint = 0
    items := strings.Split(line, " ")
    for _, item := range items {
        if contains(winningNumbers, item) {
            ret += 1
        }
    }
    return ret
}

func processLineUint(line string, winningNumbers []uint) uint {
    var ret uint = 0
    items := split(line)
    for _, item := range items {
        if containsUint(winningNumbers, item) {
            ret += 1
        }
    }
    return ret
}

func displayOutput(results map[uint]int32, t time.Duration) {
    if os.Getenv("DEBUG") == "1" {
        fmt.Printf("%d %d %d %d (Time taken: %s)\n", results[2], results[3], results[4], results[5], t)
    } else {
        fmt.Printf("%d %d %d %d\n", results[2], results[3], results[4], results[5])
    }
}

func main()  {
    if len(os.Args) < 2 {
        log.Fatalf("Usage: %s input_file_name", os.Args[0]);
    }

    path := os.Args[1]
	file, err := os.Open(path)
    if err != nil {
        log.Fatalf("Can't open file: %s", path)
    }
    defer file.Close()

    fmt.Println("READY")
    stdinScanner := bufio.NewScanner(os.Stdin)
    for stdinScanner.Scan() {
        winningNumbers := strings.Split(stdinScanner.Text(), " ")
        //winningNumbers := split(stdinScanner.Text())
        if len(winningNumbers) != 5 {
            break
        }
        start := time.Now()
        scanner := bufio.NewScanner(file)
        results := make(map[uint]int32)
        for scanner.Scan() {
            result := processLine(scanner.Text(), winningNumbers)
            if result > 1 {
                results[result] += 1
            }
        }
        displayOutput(results, time.Since(start))
        file.Seek(0, 0)
    }

}
