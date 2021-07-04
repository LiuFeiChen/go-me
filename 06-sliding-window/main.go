package main

import (
	"fmt"
	"sync"
	"time"
)

var DefaultSlidingWdSize int64 = 10

type SlidingWd struct {
	Wds   map[int64]*SWDValue
	Mutex *sync.RWMutex
}

type SWDValue struct {
	Value uint64
}

func NewSWD() *SlidingWd {
	return &SlidingWd{
		Wds:   make(map[int64]*SWDValue),
		Mutex: &sync.RWMutex{},
	}
}

func (s *SlidingWd) getNowTimeWd() *SWDValue {
	now := time.Now().Unix()

	var res *SWDValue
	var ok bool

	if res, ok = s.Wds[now]; !ok {
		res = &SWDValue{}
		s.Wds[now] = res
	}
	return res
}

func (s *SlidingWd) removeOldWd() {
	removeTime := time.Now().Unix() - DefaultSlidingWdSize

	for k, _ := range s.Wds {
		if k < removeTime {
			delete(s.Wds, k)
		}
	}
}

func (s *SlidingWd) Add(in uint64) {
	if in == 0 {
		return
	}

	nwt := s.getNowTimeWd()
	nwt.Value += in

	s.removeOldWd()
}

func (s *SlidingWd) Sum(now time.Time) uint64 {
	sum := uint64(0)

	s.Mutex.RLock()
	defer s.Mutex.RUnlock()

	for k, v := range s.Wds {
		if k > now.Unix()-DefaultSlidingWdSize {
			sum += v.Value
		}
	}

	return sum
}

func main() {
	swd := NewSWD()
	for {
		swd.Add(1)
		time.Sleep(time.Second * 1)
		sum := swd.Sum(time.Now())
		fmt.Println(sum)
	}
}
