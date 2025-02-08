// package main

// import (
//     "fmt"
//     "net"
//     "time"
// )

// func scanPort(protocol, target string, port int, timeout time.Duration) bool {
//     address := fmt.Sprintf("%s:%d", target, port)
//     conn, err := net.DialTimeout(protocol, address, timeout)
//     if err != nil {
//         return false
//     }
//     conn.Close()
//     return true
// }

// func main() {
//     var target string
//     fmt.Print("Enter target IP or domain: ")
//     fmt.Scan(&target)

//     fmt.Println("Scanning ports...")

//     startTime := time.Now()

//     startPort := 1
//     endPort := 1024
//     timeout := 1 * time.Second

//     openPorts := []int{}

//     for port := startPort; port <= endPort; port++ {
//         if scanPort("tcp", target, port, timeout) {
//             fmt.Printf("[+] Port %d is open\n", port)
//             openPorts = append(openPorts, port)
//         }
//     }

//     elapsed := time.Since(startTime)

//     if len(openPorts) == 0 {
//         fmt.Println("No open ports found.")
//     } else {
//         fmt.Println("Open ports:", openPorts)
//     }

//     fmt.Printf("\nScan completed in %s\n", elapsed)
// }


// The above approach is not efficient because it scans ports sequentially.
// Down approach improved the performance by scanning ports concurrently using goroutines and channels.



package main

import (
    "fmt"
    "net"
    "sync"
    "time"
)

const WORKER_POOL_SIZE = 100

type ScanResult struct {
    Port    int
    IsOpen  bool
}

func scanPort(protocol, target string, port int, timeout time.Duration) bool {
    address := fmt.Sprintf("%s:%d", target, port)
    conn, err := net.DialTimeout(protocol, address, timeout)
    if err != nil {
        return false
    }
    conn.Close()
    return true
}

func worker(protocol, target string, timeout time.Duration, jobs <-chan int, results chan<- ScanResult, wg *sync.WaitGroup) {
    defer wg.Done()

    for port := range jobs {
        isOpen := scanPort(protocol, target, port, timeout)
        results <- ScanResult{Port: port, IsOpen: isOpen}
    }
}

func main() {
    var target string
    fmt.Print("Enter target IP or domain: ")
    fmt.Scan(&target)

    fmt.Println("Scanning ports...")
    startTime := time.Now()

    startPort := 1
    endPort := 1024
    timeout := 500 * time.Millisecond

    numPorts := endPort - startPort + 1
    jobs := make(chan int, numPorts)
    results := make(chan ScanResult, numPorts)

    var wg sync.WaitGroup

    for i := 0; i < WORKER_POOL_SIZE; i++ {
        wg.Add(1)
        go worker("tcp", target, timeout, jobs, results, &wg)
    }

    go func() {
        for port := startPort; port <= endPort; port++ {
            jobs <- port
        }
        close(jobs)
    }()

    go func() {
        wg.Wait()
        close(results)
    }()

    openPorts := []int{}
    for result := range results {
        if result.IsOpen {
            fmt.Printf("[+] Port %d is open\n", result.Port)
            openPorts = append(openPorts, result.Port)
        }
    }

    elapsed := time.Since(startTime)

    if len(openPorts) == 0 {
        fmt.Println("No open ports found.")
    } else {
        fmt.Println("Open ports:", openPorts)
    }

    fmt.Printf("\nScan completed in %s\n", elapsed)
}

