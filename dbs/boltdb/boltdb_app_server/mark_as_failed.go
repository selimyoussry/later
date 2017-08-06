package boltdb_app_server

// MarkAsFailed
func (database *Database) MarkAsFailed(taskName string, instanceID string) error {

	srcBucketName := bucket(taskName)
	dstBucketName := []byte(BUCKET_FAILED)
	instancesIDs := []string{instanceID}

	_, err := database.moveInstances(srcBucketName, dstBucketName, instancesIDs)

	return err

}
