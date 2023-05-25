package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/yangleyings/ServiceMeshTest/kitex/client/tests"
)

var (
	testNames   []string
	concurrency = 1                // concurrency
	total       = 10000            // total requests for all clients
	host        = "127.0.0.1:8972" // server ip and port
	pool        = 1                // shared kitex clients
	rate        = 0                // throughputs
	// concurrency = flag.Int("c", 1, "concurrency")
	// total       = flag.Int("n", 10000, "total requests for all clients")
	// host        = flag.String("s", "127.0.0.1:8972", "server ip and port")
	// pool        = flag.Int("pool", 1, " shared kitex clients")
	// rate        = flag.Int("r", 0, "throughputs")
)

func main() {
	// flag.Parse()
	// 读取测试用例文件
	file := os.Args[1]
	f, err := os.Open(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	} else {
		input := bufio.NewScanner(f)
		for input.Scan() {
			arg := input.Text()
			if len(arg) < 3 {
				fmt.Fprintf(os.Stderr, "Error: Test format is wrong!")
			}

			sign := arg[:3]
			if sign == "===" {
				if len(arg) < 5 {
					testNames = append(testNames, "ThroughPut")
				}
				testNames = append(testNames, arg[4:])
				if arg[4:] == "ThroughPut" {
					tests.ThroughPut()
				}
			}
		}
		f.Close()
	}

}
