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

func split(line string) ([5]uint8, error) {
    items := strings.Split(line, " ")
    values := make(map[uint8]bool)
    var ret [5]uint8
    for _, v := range items {
        u64, err := strconv.ParseUint(v, 10, 8)
        if err != nil || u64 < 1 || u64 > 90{
            return ret, fmt.Errorf("Invalid input, numbers should be in the 0,90 range: %s in: %s", v, line)
        }
        if values[uint8(u64)] {
            return ret, fmt.Errorf("Duplicate number: %s in: %s", v, line)
        }
        values[uint8(u64)] = true
    }

    if len(values) != 5 {
        return ret, fmt.Errorf("Invalid numbers. They should contain exactly 5 numbers in the 0,90 range: %s", line)
    }

    i := 0
    for v := range values {
        ret[i] = v
        i++
    }

    if len(ret) < 5 {
        return ret, fmt.Errorf("Not enough valid numbers: %s", line)
    }
    return ret, nil
}

func processLineUint(items [5]uint8, winningNumbers [5]uint8) uint8 {
    var ret uint8
    for _, item := range items {
        if containsUint(winningNumbers, item) {
            ret++
        }
    }
    return ret
}

func containsUint(a [5]uint8, i uint8) bool {
    for _, item := range a {
        if item == i {
            return true
        }
    }
    return false
}

func showDebug() bool {
    return os.Getenv("DEBUG") == "1"
}

func displayOutput(results map[uint8]int32, t time.Duration) {
    if showDebug() {
        fmt.Printf("%d %d %d %d (Time taken: %s)\n", results[2], results[3], results[4], results[5], t)
    } else {
        fmt.Printf("%d %d %d %d\n", results[2], results[3], results[4], results[5])
    }
}

func processInput(path string) [][5]uint8 {
    startPrepare := time.Now()
	file, err := os.Open(path)
    if err != nil {
        log.Fatalf("Can't open file: %s", path)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    numbers := make([][5]uint8, 0, 10000000)
    for scanner.Scan() {
        n, err := split(scanner.Text())
        if err == nil {
            numbers = append(numbers, n)
        } else if showDebug() {
            log.Println(err)
        }
    }

    if showDebug() {
        log.Printf("Finished processing input file in %s", time.Since(startPrepare))
    }
    return numbers
}

func main()  {
    if len(os.Args) < 2 {
        log.Fatalf("Usage: %s input_file_name", os.Args[0]);
    }

    numbers := processInput(os.Args[1])

    fmt.Println("READY")
    stdinScanner := bufio.NewScanner(os.Stdin)
    for stdinScanner.Scan() {
        winningNumbers, err := split(stdinScanner.Text())
        if err != nil || len(winningNumbers) != 5 {
            log.Print("No valid input provided. Exiting.")
            break
        }
        start := time.Now()
        results := make(map[uint8]int32)
        for _, number := range numbers {
            result := processLineUint(number, winningNumbers)
            if result > 1 {
                results[result]++
            }
        }
        displayOutput(results, time.Since(start))
    }

}
