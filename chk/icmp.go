package chk

import "github.com/go-ping/ping"

func CheckICMP(address string, count int) (*ping.Statistics, error) {
	pinger, err := ping.NewPinger(address)
	if err != nil {
		return nil, err
	}
	pinger.Count = count
	err = pinger.Run()
	if err != nil {
		return nil, err
	}
	return pinger.Statistics(), nil
}
