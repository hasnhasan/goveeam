package goveeam

import (
	"net/http"
)

func GetBackupServers(veeamClient *VeeamClient) (BackupServers, error){
	backupServers := NewBackupServers(&veeamClient.Client)

	_, err := veeamClient.Client.ExecuteRequest("", http.MethodGet, "",
		"error retrieving backup servers: %s", nil, backupServers)
	if err != nil {
		return BackupServers{}, err
	}

	return *backupServers, nil
}
