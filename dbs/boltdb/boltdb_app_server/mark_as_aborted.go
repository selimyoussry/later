package boltdb_app_server

// AbortInstance aborts a single instance
func (database *Database) AbortInstance(taskName string, instanceID string) error {

	srcBucketName := bucket(taskName)
	dstBucketName := []byte(BUCKET_ABORTED)
	instancesIDs := []string{instanceID}

	_, err := database.moveInstances(srcBucketName, dstBucketName, instancesIDs)

	return err

}
