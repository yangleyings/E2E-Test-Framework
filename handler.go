package main

import (
	api "ServiceMeshTest/kitex_gen/api"

//	"github.com/apache/thrift/lib/go/thrift"

	"context"
	"log"
    "io/ioutil"
)

// EchoImpl implements the last service interface defined in the IDL.
type EchoImpl struct{}

// Echo implements the EchoImpl interface.
func (s *EchoImpl) Echo(ctx context.Context, req *api.Request) (resp *api.Response, err error) {
	// TODO: Your code here...
	return &api.Response{Message: req.Message}, nil
}

// Add Serialization Function.
/*func serializeThrift(req *api.AddSerializationRequest)([]byte, error){
	transport := thrift.NewTMemoryBufferLen(1024)
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	err := req.Write(protocolFactory.GetProtocol(transport))
	if err != nil{
		return nil, err
	}
	return transport.Bytes(), nil
}*/


/* Add Deserialization Function.
func deserializeThrift(data []byte) (*api.AddRequest, error) {
    req := &api.AddRequest{}
    transport := thrift.NewTMemoryBufferLen(len(data))
    transport.Write(data)
    protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
    err := req.Read(protocolFactory.GetProtocol(transport))
    if err != nil {
        return nil, err
    }
    return req, nil
}*/

func (s *EchoImpl) Add(ctx context.Context, req *api.AddRequest) (resp *api.AddResponse, err error) {
    /* Serialize request messages
    serializedRequest, err := serializeThrift(req)
    if err != nil {
        return nil, err
    }*/

    /*Log the serialized request fragment
    startIndex := 0
    endIndex := 10
    if endIndex > len(serializedRequest) {
        endIndex = len(serializedRequest)
    }
    serializedRequestFragment := serializedRequest[startIndex:endIndex]
    log.Printf("Serialized request fragment: %s", serializedRequestFragment)*/

    /* Deserialize request messages
    deserializedRequest, err := deserializeThrift(serializedRequest)
    if err != nil {
        return nil, err
    }*/

    // Construct the response message
    fileContent, err := ioutil.ReadFile("file.txt")
    if err !=nil{
        return nil, err
    }
    log.Println(string(fileContent))

    resp = &api.AddResponse{Sum: req.First + req.Second}

    return resp, nil
}