package main

import (
	"ServiceMeshTest/kitex_gen/api/echo"
	"github.com/cloudwego/kitex/client"
	"ServiceMeshTest/kitex_gen/api"
	"github.com/cloudwego/kitex/client/callopt"

	"context"
	"log"
	"time"
)

func main(){
	c, err := echo.NewClient("ServiceMeshTest", client.WithHostPorts("0.0.0.0:8888"))
if err != nil {
  log.Fatal(err)}
 req := &api.Request{Message: "my request"}
resp, err := c.Echo(context.Background(), req, callopt.WithRPCTimeout(3*time.Second))

if err != nil {
  log.Fatal(err)
}
log.Println(resp)
}
