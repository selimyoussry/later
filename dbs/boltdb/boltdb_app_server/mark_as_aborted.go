package boltdb_app_server

// AbortInstance aborts a single instance
func (database *Database) AbortInstance(instanceID string) error {

	srcBucketName := []byte(BUCKET_PENDING)
	dstBucketName := []byte(BUCKET_ABORTED)
	instancesIDs := []string{instanceID}

	_, err := database.moveInstances(srcBucketName, dstBucketName, instancesIDs)

	return err

}
