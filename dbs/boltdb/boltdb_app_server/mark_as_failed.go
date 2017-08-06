package boltdb_app_server

// MarkAsFailed
func (database *Database) MarkAsFailed(instanceID string) error {

	srcBucketName := []byte(BUCKET_PENDING)
	dstBucketName := []byte(BUCKET_FAILED)
	instancesIDs := []string{instanceID}

	_, err := database.moveInstances(srcBucketName, dstBucketName, instancesIDs)

	return err

}
