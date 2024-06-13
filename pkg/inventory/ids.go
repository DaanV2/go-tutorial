package inventory

import "github.com/DaanV2/go-tutorial/pkg/rand"

var ids []string

func init() {
	ids = make([]string, 0)

	for range 42 {
		ids = append(ids, rand.RandomID())
	}
}

func GetRandomId() string {
	return ids[rand.Int64(len(ids))]
}
