package main

import (
	"fmt"
	"reflect"

	//"github.com/ethereum/go-ethereum/p2p"
	//"github.com/ethereum/go-ethereum/rpc"
	"study-go/helper"
) /*
type Service interface {
	// Protocols retrieves the P2P protocols the service wishes to start.
	Protocols() []p2p.Protocol

	// APIs retrieves the list of RPC descriptors the service provides
	APIs() []rpc.API

	// Start is called after all services have been constructed and the networking
	// layer was also initialized to spawn any goroutines required by the service.
	Init(server *p2p.Server) error

	// Stop terminates all goroutines belonging to the service, blocking until they
	// are all terminated.
	Stop() error
}

type SI struct {
	Name string
}

func (s SI)APIs() {
	//a:= []rpc.API{}
	//return a
}
*/

func main() {
	/*services := make(map[reflect.Type]Service)
	for kind, s := range services { // copy needed for threaded access
		fmt.Println(kind, s)
		fmt.Println("999")
	}*/
	//kind := reflect.TypeOf(service)
	//go fmt.Println("ok")
	a := helper.Helper{}
	fmt.Println(a)
	fmt.Println(reflect.TypeOf(a).Name())
	fmt.Println(reflect.TypeOf(a).PkgPath())
	fmt.Println(reflect.TypeOf(a).PkgPath())
}
