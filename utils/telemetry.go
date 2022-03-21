package utils

import (
	"fmt"

	"github.com/showwin/speedtest-go/speedtest"
)

// GetBandwidth returns the bandwidth of the system currently
// Bandwdidth is a measure of performance and here it refers to
// upload speed and download speed
func GetBandwidth() string {
	user, _ := speedtest.FetchUserInfo()

	serverList, _ := speedtest.FetchServers(user)
	targets, _ := serverList.FindServer([]int{})

	for _, s := range targets {
		s.PingTest()
		s.DownloadTest(false)
		s.UploadTest(false)

		return fmt.Sprintf("%s,%f,%f ", s.Latency, s.DLSpeed, s.ULSpeed)
	}

	return ""
}
