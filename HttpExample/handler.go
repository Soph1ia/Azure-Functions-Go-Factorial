package main

import (
    "fmt"
    "log"
    "math/big"
    "net/http"
    "os"
    "strconv"
    "time"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    message := "This HTTP triggered function executed successfully. Pass a number in the query string for a personalized response.\n"
    number := r.URL.Query().Get("number")

    if number != "" {
         number, _ := strconv.Atoi(number)
        benchmark("factorial", number, w)
    } else {
		fmt.Fprint(w, message)
	}

}

func main() {
    listenAddr := ":8080"
    if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
        listenAddr = ":" + val
    }
    http.HandleFunc("/api/HttpExample", helloHandler)
    log.Printf("About to listen on %s. Go to https://127.0.0.1%s/", listenAddr, listenAddr)
    log.Fatal(http.ListenAndServe(listenAddr, nil))
}


/**
Method : Benchmark

This method gets the time taken to execute the factorial 40 times.
In total it loops 80 times.
It takes the last 20 execution times.
Gets the average time
Calculates the throughput as time / 40

Prints out the throughput.

returns: none

*/
func benchmark(funcName string, number int, w http.ResponseWriter) {
    listofTime := [41]int64{}

    for j := 0; j < 40; j++ {
        start := time.Now().UnixNano()
        factorial(number)

        // End time
        end := time.Now().UnixNano()
        // Results
        difference := end - start
        listofTime[j] = difference

    }
    // Average Time
    sum := int64(0)
    for i := 0; i < len(listofTime); i++ {
        // adding the values of
        // array to the variable sum
        sum += listofTime[i]
    }
    // avg to find the average
    avg := (float64(sum)) / (float64(len(listofTime)))

    // Throughput Rate
    throughput := 40/avg

    // Response
    fmt.Fprintf(w, "Time taken by %s function is %v ops/ns \n", funcName, throughput)
}

/**
Method: Factorial

Calculates the factorial of the number provided

Returns: pointer to big int
*/
func factorial(n int) *big.Int {
    factVal := big.NewInt(1)
    if n < 0 {
        fmt.Print("Factorial of negative number doesn't exist.")
    } else {
        for i := 1; i <= n; i++ {
            //factVal *= uint64(i) // mismatched types int64 and int
            factVal = factVal.Mul(factVal, big.NewInt(int64(i)))
        }
    }
    return factVal
}