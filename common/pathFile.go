package common

import "fmt"

func PathFile(path string) string {
	return fmt.Sprintf("provider/%s/resident.xml", path)
}
func PathFileReplica(path string) string {
	return fmt.Sprintf("provider/%s/resident_replica.xml", path)
}
