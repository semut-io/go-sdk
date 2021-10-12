package common

// Endpoint An API endpoint mapped by name to its path
type Endpoint map[string]string

// Endpoints all endpoints availble across Platform API versions
type Endpoints map[Version]Endpoint

// GetSupportedVersions will get versions supported by the SDK, currently v1 only
func GetSupportedVersions() Versions {
	return []Version{
		"v1",
	}
}

// SetVersion sets the version of the API that is to be used, globally
func SetVersion(apiVersion Version) (bool, error) {
	// if the version already set is the same as the
	// version being set
	if APIVersion == apiVersion {
		return true, nil
	} else {
		// check if passed version is valid
		for _, version := range GetSupportedVersions() {
			if version == apiVersion {
				APIVersion = apiVersion
				return true, nil
			}
		}
		// if version is invalid, return error
		return false, ErrVersionNotAvailable
	}
}

// GetEndpoints returns all supported endpoints in a version
func GetEndpoints(version Version) (Endpoint, error) {

	var endpoints Endpoints = make(Endpoints)
	endpoints["v1"] = Endpoint{
		// Applications
		"ApplicationInvoke": "/v1/application/invoke",

		// Cron
		"CronCreate": "/v1/cron/create",
		"CronDelete": "/v1/cron/delete",
		"CronList":   "/v1/cron/list",
		"CronUpdate": "/v1/cron/update",

		// Database (ADB KV)
		"DatabaseDelete": "/v1/database/delete",
		"DatabaseGet":    "/v1/database/get",
		"DatabaseList":   "/v1/database/list",
		"DatabaseRename": "/v1/database/rename",
		"DatabaseSet":    "/v1/database/set",

		// Deployments
		"DeploymentDescribe":  "/v1/deployment/describe",
		"DeploymentLaunch":    "/v1/deployment/launch",
		"DeploymentStart":     "/v1/deployment/start",
		"DeploymentStop":      "/v1/deployment/stop",
		"DeploymentTerminate": "/v1/deployment/terminate",

		// Encryption
		"EncryptionDecryptContent": "/v1/encryption/decryptcontent",
		"EncryptionDecryptFile":    "/v1/encryption/decrypt",
		"EncryptionDeleteKey":      "/v1/encryption/deleteencryptionkey",
		"EncryptionEncryptContent": "/v1/encryption/encryptcontent",
		"EncryptionEncryptFile":    "/v1/encryption/encrypt",

		// Events
		"EventSubscriptionsDelete": "/v1/subscriptions/delete",
		"EventSubscriptionsList":   "/v1/subscriptions/list",
		"EventSubscriptionsNew":    "/v1/subscriptions/new",
		"EventsDelete":             "/v1/events/delete",
		"EventsList":               "/v1/events/list",
		"EventsNew":                "/v1/events/new",
		"EventsUpdate":             "/v1/events/update",

		// Exec
		"Exec":     "/v1/exec/exec",
		"ExecKill": "/v1/exec/kill",
		"ExecRun":  "/v1/exec/run",

		// Images
		"ImagesDelete": "/v1/image/delete",
		"ImagesGet":    "/v1/image/get",
		"ImagesList":   "/v1/image/list",
		"ImagesNew":    "/v1/image/new",

		// Log Export
		"LogExportDisable": "/v1/logexport/disable",
		"LogExportEnable":  "/v1/logexport/enable",
		"LogExportList":    "/v1/logexport/list",
		"LogExportUpdate":  "/v1/logexport/update",

		// Metrics
		"MetricsQuery": "/v1/metrics/query",

		// Notifications
		"NotificationsAddEmail":            "/v1/notifications/addemail",
		"NotificationsDeleteEmail":         "/v1/notifications/deleteemail",
		"NotificationsListEmail":           "/v1/notifications/listemail",
		"NotificationVerifiedEmailIDsList": "/v1/notifications/listverifiedemailids",
		"NotificationsSendEmail":           "/v1/notifications/sendemail",
		"NotificationsUpdateEmail":         "/v1/notifications/updateemail",
		"NotificationVerifyEmailID":        "/v1/notifications/verifyemailid",

		// Object Store
		"ObjectStoreDelete":   "/v1/objectstore/delete",
		"ObjectStoreDescribe": "/v1/objectstore/describe",
		"ObjectStoreFetch":    "/v1/objectstore/fetch",
		"ObjectStoreList":     "/v1/objectstore/list",
		"ObjectStorePost":     "/v1/objectstore/post",

		// Secrets
		"SecretsDelete":                   "/v1/secrets/delete",
		"SecretsGenerateCredentials":      "/v1/secrets/generatecredentials",
		"SecretsGenerateStoreCredentials": "/v1/secrets/generateandstorecredentials",
		"SecretsRetrieve":                 "/v1/secrets/retrieve",
		"SecretsStore":                    "/v1/secrets/store",

		// Snapshots
		"SnapshotsDelete":   "/v1/snapshot/delete",
		"SnapshotsDescribe": "/v1/snapshot/describe",
		"SnapshotsTake":     "/v1/snapshot/initiate",

		// Volumes
		"VolumesAttach":       "/v1/volume/attach",
		"VolumesCopy":         "/v1/volume/copy",
		"VolumesCreate":       "/v1/volume/create",
		"VolumesCreateAttach": "/v1/volume/createandattach",
		"VolumesDelete":       "/v1/volume/delete",
		"VolumesDescribe":     "/v1/volume/describe",
		"VolumesDetach":       "/v1/volume/detach",

		// Workers
		"WorkerDescribe":             "/v1/worker/describe",
		"WorkerHealthStatus":         "/v1/worker/describehealth",
		"WorkerLaunch":               "/v1/worker/launch",
		"WorkerMarkHealthy":          "/v1/worker/markhealthy",
		"WorkerMarkUnhealthy":        "/v1/worker/markunhealthy",
		"WorkerStart":                "/v1/worker/start",
		"WorkerStop":                 "/v1/worker/stop",
		"WorkerUpdate":               "/v1/worker/update",
		"WorkerUpdateResourceLimits": "/v1/worker/updateresourcelimits",
		"WorkerTerminate":            "/v1/worker/terminate",

		// Worker Groups
		"WorkerGroupChangeStrategy":       "/v1/workergroup/changestrategy",
		"WorkerGroupDescribe":             "/v1/workergroup/describe",
		"WorkerGroupHealthStatus":         "/v1/workergroup/describehealth",
		"WorkerGroupLaunch":               "/v1/workergroup/launch",
		"WorkerGroupMarkHealthy":          "/v1/workergroup/markhealthy",
		"WorkerGroupMarkUnhealthy":        "/v1/workergroup/markunhealthy",
		"WorkerGroupScale":                "/v1/workergroup/scale",
		"WorkerGroupStart":                "/v1/workergroup/start",
		"WorkerGroupStop":                 "/v1/workergroup/stop",
		"WorkerGroupUpdate":               "/v1/workergroup/update",
		"WorkerGroupUpdateResourceLimits": "/v1/workergroup/updateresourcelimits",
		"WorkerGroupTerminate":            "/v1/workergroup/terminate",
	}

	endpoint, ok := endpoints[version]

	if !ok {
		return nil, ErrVersionNotAvailable
	}

	return endpoint, nil
}

// GetEndpoint returns the enpoint for a given version and endpoint name
func GetEndpoint(name string) (string, error) {
	endpoints, err := GetEndpoints(APIVersion)

	if err != nil {
		return "", ErrVersionNotAvailable
	}

	endpoint, ok := endpoints[name]

	if !ok {
		return "", ErrEndpointNotAvailable
	}

	return endpoint, nil
}
