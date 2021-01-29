package chk

import (
	"time"

	"github.com/go-ping/ping"
)

func CheckICMP(address string, count int, timeout int, interval int) (*ping.Statistics, error) {
	pinger, err := ping.NewPinger(address)
	if err != nil {
		return nil, err
	}
	pinger.Count = count
	pinger.Timeout = time.Duration(timeout * int(time.Second))
	pinger.Interval = time.Duration(interval * int(time.Second))
	err = pinger.Run()
	if err != nil {
		return nil, err
	}
	return pinger.Statistics(), nil
}
