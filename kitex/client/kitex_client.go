package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/yangleyings/ServiceMeshTest/kitex/client/tests"
)

var (
	testNames   []string           // acquired tests
	test_ok     []bool             // whether tested
	concurrency = 1                // concurrency
	total       = 10000            // total requests for all clients
	host        = "127.0.0.1:8972" // server ip and port
	pool        = 1                // shared kitex clients
	rate        = 0                // throughputs
)

func main() {
	// Read test file
	var file string
	if len(os.Args) < 2 {
		fmt.Println("Please enter TEST FILE (eg. test.txt): ")
		fmt.Scanln(&file)
	} else {
		file = os.Args[1]
	}
	f, err := os.Open(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	} else {
		input := bufio.NewScanner(f)

		// Read test content
		for input.Scan() {
			arg := input.Text()
			if len(arg) < 3 {
				fmt.Fprintf(os.Stderr, "Error: Test format is wrong!")
			} else {
				sign := arg[:3]

				// Get test parameters
				switch sign {
				case "===":
					var testName string

					if len(arg) < 5 { // Failed to get the test name
						fmt.Println("Please enter accessible TEST NAME (ThroughPut OR ...): ")
						fmt.Scanln(&testName)
					} else {
						testName = arg[4:]
					}
					if testName != "ThroughPut" {
						fmt.Fprintf(os.Stderr, "Error: Test name is wrong!")
						os.Exit(0)
					} else {
						testNames = append(testNames, testName)
						test_ok = append(test_ok, false)
					}

				case "---":
					if len(arg) < 7 { // Cannot get the parameter
						fmt.Fprintf(os.Stderr, "Error: Can't get the right parameter!")
						os.Exit(0)
					} else {
						var par string
						par = arg[4:5]
						if par == "c" {
							c, err := strconv.Atoi(arg[6:])
							if err != nil {
								fmt.Println("Failed to get parameter c:", err)
								os.Exit(0)
							} else {
								concurrency = c
							}
						} else if par == "n" {
							n, err := strconv.Atoi(arg[6:])
							if err != nil {
								fmt.Println("Failed to get parameter n:", err)
								os.Exit(0)
							} else {
								total = n
							}
						} else if par == "s" {
							host = arg[6:]
						} else if par == "pool" {
							p, err := strconv.Atoi(arg[6:])
							if err != nil {
								fmt.Println("Failed to get parameter p:", err)
								os.Exit(0)
							} else {
								pool = p
							}
						} else if par == "r" {
							r, err := strconv.Atoi(arg[6:])
							if err != nil {
								fmt.Println("Failed to get parameter r:", err)
								os.Exit(0)
							} else {
								rate = r
							}
						}
					}

				case "END":
					l := len(testNames)
					if l < 1 {
						fmt.Fprintf(os.Stderr, "Error: Can't find the right test!")
						os.Exit(0)
					} else if test_ok[l-1] == true {
						fmt.Fprintf(os.Stderr, "Error: Can't find the right test!")
						os.Exit(0)
					} else {
						test_ok[l-1] = true
						fmt.Printf("Test case %d: %s\n", l, testNames[l-1])
						tests.ThroughPut(concurrency, total, host, pool, rate)
					}
				}
			}
		}
		f.Close()
	}
}
