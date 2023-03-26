package main

import (
	"ServiceMeshTest/kitex_gen/api/echo"
	"github.com/cloudwego/kitex/client"
	"ServiceMeshTest/kitex_gen/api"
	"github.com/cloudwego/kitex/client/callopt"

	"context"
	"log"
	"time"

//	"os"
//	"fmt"
//	"strconv"
)

func main(){
	c, err := echo.NewClient("ServiceMeshTest", client.WithHostPorts("0.0.0.0:8888"))
	if err != nil {
		log.Fatal(err)
	}
 	req := &api.Request{Message: "my request"}
	resp, err := c.Echo(context.Background(), req, callopt.WithRPCTimeout(3*time.Second))
	if err != nil {
  		log.Fatal(err)
	}
	log.Println(resp)

//add request and response
	for {
		req := &api.Request{Message: "my request"}
		resp, err := c.Echo(context.Background(), req)
		if err != nil {
				log.Fatal(err)
		}
	    log.Println(resp)
	    time.Sleep(time.Second)
	
		//original add method 
		addReq := &api.AddRequest{First: 512, Second: 512}
	    addResp, err := c.Add(context.Background(), addReq)
	    if err != nil {
			log.Fatal(err)
	    }
	    log.Println(addResp)
	    time.Sleep(time.Second)

		/*read file

		//read input values from file
		inputFile, err := os.Open("file.txt")
		if err != nil{
			log.Fatal(err)
		}
		defer inputFile.Close()

		var a,b int64
		_, err = fmt.Fscanf(inputFile,"%d%d",&a,&b)
		if err != nil{
			log.Fatal(err)
		}

		//Call Add method with input value
		addReq := &api.AddRequest{First: a, Second: b}
	    addResp, err := c.Add(context.Background(), addReq)
		if err != nil{
			log.Fatal(err)
		}

		//write output value to file
		outputFile, err := os.Create("outfile.txt")
		if err != nil{
			log.Fatal(err)
		}
		defer outputFile.Close()

		_,err = outputFile.WriteString(strconv.FormatInt(addResp.Sum, 10))
		if err != nil{
			log.Fatal(err)
		}

		time.Sleep(time.Second)

		*/
}

}