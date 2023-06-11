package tests

import (
	stdlog "log"
	"os"
	"reflect"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/gogo/protobuf/proto"
	pb "github.com/rpcxio/rpcx-benchmark/proto"
	"github.com/rpcxio/rpcx-benchmark/stat"
	"github.com/smallnest/rpcx/log"
	"github.com/yangleyings/ServiceMeshTest/kitex/pb/hello"
	"go.uber.org/ratelimit"
	"golang.org/x/net/context"
)

func ThroughPut(concurrency int, total int, host string, pool int, rate int) {

	// flag.Parse()
	// 读取测试用例文件
	// a := make(map[string]map[string]int)

	log.SetLogger(log.NewDefaultLogger(os.Stdout, "", stdlog.LstdFlags|stdlog.Lshortfile, log.LvInfo))

	var rl ratelimit.Limiter
	if rate > 0 {
		rl = ratelimit.New(rate)
	}

	// 并发goroutine数.模拟客户端
	n := concurrency
	// 每个客户端需要发送的请求数
	m := total / n
	log.Infof("concurrency: %d\nrequests per client: %d\n\n", n, m)

	servers := strings.Split(host, ",")
	log.Infof("Servers: %+v\n\n", host)
	if pool > 1 {
		log.Warnf("Notice: Kitex doesn't need the 'pool' param, not set is suggested")
	}

	args := prepareArgs()

	// 请求消息大小
	b, _ := proto.Marshal(args)
	log.Infof("message size: %d bytes\n\n", len(b))

	// 等待所有测试完成
	var wg sync.WaitGroup
	wg.Add(n * m)

	// 总请求数
	var trans uint64
	// 返回正常的总请求数
	var transOK uint64

	// 每个goroutine的耗时记录
	d := make([][]int64, n, n)

	// The kitex client uses a built-in connection pool with multiplexing.
	client := hello.MustNewClient("echo",
		client.WithHostPorts(servers...),
		client.WithMuxConnection(pool),
	)
	// warmup
	var warmWg sync.WaitGroup
	for i := 0; i < pool; i++ {
		warmWg.Add(1)
		go func() {
			defer warmWg.Done()
			for j := 0; j < 5; j++ {
				client.Say(context.Background(), args)
			}
		}()
	}
	warmWg.Wait()

	// Fence, control client starts testing at the same time
	var startWg sync.WaitGroup
	startWg.Add(n + 1) // The +1 is because there is a goroutine that records the start time

	// Create the client goroutine and start testing
	startTime := time.Now().UnixNano()
	go func() {
		startWg.Done()
		startWg.Wait()
		startTime = time.Now().UnixNano()
	}()
	for i := 0; i < n; i++ {
		dt := make([]int64, 0, m)
		d = append(d, dt)

		go func(i int) {
			startWg.Done()
			startWg.Wait()

			for j := 0; j < m; j++ {
				// Current limiting: The current limiting time is not included in the waiting time.
				if rl != nil {
					rl.Take()
				}

				t := time.Now().UnixNano()

				reply, err := client.Say(context.Background(), args)
				// Waiting time + service time. Waiting time is the waiting time scheduled
				// by the client and the time read and scheduled by the server.
				// Service time is the actual time when the request is processed by the service.
				t = time.Now().UnixNano() - t

				d[i] = append(d[i], t)

				if err == nil && reply.Field1 == "OK" {
					atomic.AddUint64(&transOK, 1)
				}

				atomic.AddUint64(&trans, 1)
				wg.Done()
			}

		}(i)
	}

	wg.Wait()
	// Statistic
	stat.Stats(startTime, total, d, trans, transOK)
}

func prepareArgs() *pb.BenchmarkMessage {
	b := true
	var i int32 = 100000
	var i64 int64 = 100000
	s := "This is an test message."

	var args pb.BenchmarkMessage

	v := reflect.ValueOf(&args).Elem()
	num := v.NumField()
	for k := 0; k < num; k++ {
		field := v.Field(k)
		if !field.CanSet() {
			continue
		}
		if field.Type().Kind() == reflect.Pointer {
			switch v.Field(k).Type().Elem().Kind() {
			case reflect.Int, reflect.Int32:
				field.Set(reflect.ValueOf(&i))
			case reflect.Int64:
				field.Set(reflect.ValueOf(&i64))
			case reflect.Bool:
				field.Set(reflect.ValueOf(&b))
			case reflect.String:
				field.Set(reflect.ValueOf(&s))
			}
		} else {
			switch field.Kind() {
			case reflect.Int, reflect.Int32, reflect.Int64:
				field.SetInt(100000)
			case reflect.Bool:
				field.SetBool(true)
			case reflect.String:
				field.SetString(s)
			}
		}

	}
	return &args
}
