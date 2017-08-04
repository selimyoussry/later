package boltdb

import "fmt"

func bucket(taskName string) []byte {
	return []byte(fmt.Sprintf("Instances.%s", taskName))
}
