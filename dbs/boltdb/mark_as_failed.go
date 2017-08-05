package boltdb

import (
	"encoding/json"
)

// AbortInstances should be able to abort the instances with either
// - a list of IDs
// - all for a given task
// - between a start and end time
func (database *Database) MarkAsFailed(taskName string, parameters []byte) ([]string, error) {

	var input Input
	err := json.Unmarshal(parameters, &input)
	if err != nil {
		return []string{}, err
	}

	srcBucketName := bucket(taskName)
	dstBucketName := []byte(BUCKET_FAILED)

	return database.moveInstances(srcBucketName, dstBucketName, input.InstancesIDs)

}
