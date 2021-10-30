package function

import "time"

func getUnixTime() int64 {
	return time.Now().Unix()
}
