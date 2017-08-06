package boltdb_app_server

// MarkAsSuccessful
func (database *Database) MarkAsSuccessful(taskName string, instanceID string) error {

	srcBucketName := bucket(taskName)
	dstBucketName := []byte(BUCKET_SUCCESSFUL)
	instancesIDs := []string{instanceID}

	_, err := database.moveInstances(srcBucketName, dstBucketName, instancesIDs)

	return err

}
