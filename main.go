package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
	"time"
)

func main() {
	throttled := NewThrottler(
		http.DefaultTransport,
		60,
		time.Minute, // 60 rpm
		[]string{"/servers/*/status", "/network/"}, // except servers status and network operations
		false, // wait on limit
	)
	client := http.Client{
		Transport: &throttled,
	}


	var j int64
	for i:=0;i<100;i++{
		go func() {
			req,_ := http.NewRequest("PUT", "http://apidomain.com/images/reload", nil)
			resp, err:= client.Do(req)
			fmt.Printf("%v %#v %#v\n",atomic.LoadInt64(&j),err,resp)
			atomic.AddInt64(&j,1)
		}()
	}
	fmt.Scanln()


	// ...
	// no throttling
	resp, err:= client.Get("http://apidomain.com/network/routes")
	fmt.Printf("%#v %#v\n",err,resp)
	// ...
	fmt.Printf("%#v %#v\n",err,resp)
	req,_ := http.NewRequest("PUT", "http://apidomain.com/images/reload", nil)
	resp, err= client.Do(req)
	// ...

	// ...
	// no throttling
	resp, err= client.Get("http://apidomain.com/servers/1337/status?simple=true")
	// ...
	fmt.Printf("%#v %#v\n",err,resp)

}

