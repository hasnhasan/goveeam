package goveeam

import "github.com/hasnhasan/goveeam/types/v1"

type BackupServers []BackupServer

type BackupServer struct {
	BackupServer *types.BackupServer
	client *Client
}

func NewBackupServers(client *Client) *BackupServers {
	return &BackupServers{
		BackupServer{
			BackupServer: new(types.BackupServer),
			client: client,
		},
	}
}