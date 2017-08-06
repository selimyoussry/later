package boltdb_app_server

// MarkAsSuccessful
func (database *Database) MarkAsSuccessful(instanceID string) error {

	srcBucketName := []byte(BUCKET_PENDING)
	dstBucketName := []byte(BUCKET_SUCCESSFUL)
	instancesIDs := []string{instanceID}

	_, err := database.moveInstances(srcBucketName, dstBucketName, instancesIDs)

	return err

}
