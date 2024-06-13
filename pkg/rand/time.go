package rand

import "time"

func RandomTime() string {
	now := time.Now()
	n := Int64(1000000)
	now = now.Add(time.Duration(n) * -time.Second)

	return now.Format(time.RFC3339)
}
