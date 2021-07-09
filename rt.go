package main

import (
	"errors"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

type NewRT struct {
	tr http.RoundTripper
	limit int64
	tTime time.Duration
	prefix []string
	flag bool

	timStatus int32
	tim *time.Timer
	counter int64
	sync.Mutex
	ch chan struct{}
}

func (c *NewRT) RoundTrip(req *http.Request ) (*http.Response, error) {


	if c.limit == 0 || checkUrl(req.URL.Path,c) == 1 {
		return c.tr.RoundTrip(req)
	}

	c.Lock()
	if atomic.LoadInt64(&c.counter)>=atomic.LoadInt64(&c.limit) {
		if c.flag {
			return nil, errors.New("limit return")
		}
		<-c.ch
	}

	atomic.AddInt64(&c.counter,1)

	if atomic.LoadInt32(&c.timStatus) == 0 {
		atomic.AddInt32(&c.timStatus,1)
		time.AfterFunc(c.tTime, func() {

			atomic.AddInt64(&c.counter,-atomic.LoadInt64(&c.counter))
			atomic.AddInt32(&c.timStatus,-1)
			c.ch<- struct{}{}

		})
	}
	c.Unlock()

	return c.tr.RoundTrip(req)
}

func NewThrottler (tripper http.RoundTripper,limit int64,tTime time.Duration, prefix []string,flag bool) NewRT{
	return NewRT{
		tr:     tripper,
		limit:  limit,
		tTime:  tTime,
		prefix: prefix,
		flag:   flag,
		timStatus: 0,
		ch: make(chan struct{}, limit),
	}
}