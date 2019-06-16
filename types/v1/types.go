package types

import "fmt"

type BackupServer struct {
	HREF string `xml:"Href,attr,omitempty"`
	Type string `xml:"Type,attr,omitempty"`
	UID  string `xml:"UID,attr,omitempty"`
	Name string `xml:"Name,attr,omitempty"`

	Links       LinkList `xml:"Links,omitempty"`
	Description string  `xml:"Description,omitempty"`
	Port        int     `xml:"Port,omitempty"`
	Version     float32 `xml:"Version,omitempty"`
}

type LinkList []*Links

type Links struct {
	Link Link `xml:"Link"`
}

type Link struct {
	HREF string `xml:"Href,attr"`
	Type string `xml:"Type,attr,omitempty"`
	Rel  string `xml:"Rel,attr,omitempty"`
}

type CredentialsInfo struct {
	HREF  string  `xml:"Href,attr,omitempty"`
	Type  string  `xml:"Type,attr,omitempty"`
	Links LinkList `xml:"Links,omitempty"`

	Id          string `xml:"Id"`
	Username    string `xml:"Username"`
	Description string `xml:"Description,omitempty"`
	Password    string `xml:"Password,omitempty"`
}

type PasswordKeyInfo struct {
	HREF  string  `xml:"Href,attr,omitempty"`
	Type  string  `xml:"Type,attr,omitempty"`
	Links []*Link `xml:"Links,omitempty"`

	Id                   string `xml:"Id"`
	Hint                 string `xml:"Hint,omitempty"`
	LastModificationDate string `xml:"LastModificationDate,omitempty"`
}

type ManagedServer struct {
	HREF  string  `xml:"Href,attr,omitempty"`
	Type  string  `xml:"Type,attr,omitempty"`
	Links LinkList `xml:"Links,omitempty"`

	Description       string `xml:"Description,omitempty"`
	ManagedServerType string `xml:"ManagedServerType,omitempty"`
}

type Job struct {
	HREF        string  `xml:"Href,attr,omitempty"`
	Type        string  `xml:"Type,attr,omitempty"`
	UID         string  `xml:"UID,attr,omitempty"`
	Name        string  `xml:"Name,attr,omitempty"`
	Description string  `xml:"Description,omitempty"`
	Links       LinkList `xml:"Links,omitempty"`

	JobType            string              `xml:"JobType"`
	Platform           string              `xml:"Platform"`
	ScheduleConfigured bool                `xml:"ScheduleConfigured,omitempty"`
	ScheduleEnabled    bool                `xml:"ScheduleEnabled,omitempty"`
	JobScheduleOptions *JobScheduleOptions `xml:"JobScheduleOptions,omitempty"`
	JobsInfo           *JobInfo            `xml:"JobInfo,omitempty"`
}

type JobScheduleOptions struct {
	RetryOption                       *RetryOptions         `xml:"RetryOptions,omitempty"`
	WaitForBackupCompletion           bool                  `xml:"WaitForBackupCompletion,omitempty"`
	BackupCompetitionWaitingPeriodMin int                   `xml:"BackupCompetitionWaitingPeriodMin,omitempty"`
	OptionsDaily                      *DailyOptions         `xml:"OptionsDaily,omitempty"`
	OptionsMonthly                    *MonthlyOptions       `xml:"OptionsMonthly,omitempty"`
	OptionsPeriodically               *PeriodicOptions      `xml:"OptionsPeriodically,omitempty"`
	OptionsBackupWindow               *BackupWindowOptions  `xml:"OptionsBackupWindow,omitempty"`
	OptionsDaisyChaining              *DaisyChainingOptions `xml:"OptionsDaisyChaining,omitempty"`
}

type JobInfo struct {
	BackupJobInfo *BackupJobInfo `xml:"BackupJobInfo"`
}

type BackupJobInfo struct {
	Includes               *Includes               `xml:"Includes,omitempty"`
	GuestProcessingOptions *GuestProcessingOptions `xml:"GuestProcessingOptions,omitempty"`
	AdvancedStorageOptions *AdvancedStorageOptions `xml:"AdvancedStorageOptions,omitempty"`
}

type RetryOptions struct {
	RetryTimes     int  `xml:"RetryTimes,omitempty"`
	RetryTimeout   int  `xml:"RetryTimeout,omitempty"`
	RetrySpecified bool `xml:"RetrySpecified,omitempty"`
}

// TODO: review day structs and how they are referenced, might be better likd FSRoot style
type DailyOptions struct {
	Kind string `xml:"Kind"`
	Days string `xml:"Days"`
	Time string `xml:"Time"`
}

type Day struct {
	Name string `xml:"Name,attr"`
}

type MonthlyOptions struct {
	Time             string `xml:"Time"`
	DayNumberInMonth int    `xml:"DayNumberInMonth"`
	DayOfWeek        string `xml:"DayOfWeek"`
	Months           string `xml:"Months"`
}

type PeriodicOptions struct {
	Kind       string `xml:"Kind"`
	FullPeriod int    `xml:"FullPeriod"`
	Schedule   []*Day `xml:"Schedule"`
}

type BackupWindowOptions struct {
	TimePeriod []*Day `xml:"TimePeriod"`
}

type DaisyChainingOptions struct {
	PreviousJobUid string `xml:"PreviousJobUid,omitempty"`
}

type Includes struct {
	ObjectInJob *ObjectInJob `xml:"ObjectInJob"`
}

type ObjectInJob struct {
	ObjectInJobId          string                  `xml:"ObjectInJobId"`
	HierarchyObjRef        string                  `xml:"HierarchyObjRef"`
	Name                   string                  `xml:"Name"`
	DisplayName            string                  `xml:"DisplayName"`
	Order                  int                     `xml:"Order"`
	GuestProcessingOptions *GuestProcessingOptions `xml:"GuestProcessingOptions,omitempty"`
}

type GuestProcessingOptions struct {
	VssSnapshotOptions            *VssSnapshotOptions            `xml:"VssSnapshotOptions,omitempty"`
	WindowsGuestFSIndexingOptions *WindowsGuestFSIndexingOptions `xml:"WindowsGuestFSIndexingOptions,omitempty"`
	LinuxGuestFSIndexingOptions   *LinuxGuestFSIndexingOptions   `xml:"LinuxGuestFSIndexingOptions,omitempty"`
	SqlBackupOptions              *SqlBackupOptions              `xml:"SqlBackupOptions,omitempty"`
	WindowsCredentialsId          string                         `xml:"WindowsCredentialsId,omitempty"`
	LinuxCredentialsId            string                         `xml:"LinuxCredentialsId,omitempty"`
}

type VssSnapshotOptions struct {
	VssSnapshotMode string `xml:"VssSnapshotMode"`
	IsCopyOnly      bool   `xml:"IsCopyOnly"`
}

type WindowsGuestFSIndexingOptions struct {
	FileSystemIndexingMode  string  `xml:"FileSystemIndexingMode"`
	IncludedIndexingFolders string  `xml:"IncludedIndexingFolders,omitempty"`
	ExcludedIndexingFolders []*Path `xml:"ExcludedIndexingFolders,omitempty"`
}

type LinuxGuestFSIndexingOptions struct {
	FileSystemIndexingMode  string  `xml:"FileSystemIndexingMode"`
	IncludedIndexingFolders string  `xml:"IncludedIndexingFolders,omitempty"`
	ExcludedIndexingFolders []*Path `xml:"ExcludedIndexingFolders,omitempty"`
}

type SqlBackupOptions struct {
	TransactionLogsProcessing string `xml:"TransactionLogsProcessing"`
	BackupLogsFrequencyMin    int    `xml:"BackupLogsFrequencyMin"`
	UseDbBackupRetention      bool   `xml:"UseDbBackupRetention"`
	RetainDays                int    `xml:"RetainDays"`
}

type Path struct {
	Path string `xml:"Path,omitempty"`
}

type AdvancedStorageOptions struct {
	PasswordKeyInfo string `xml:"PasswordKeyId"`
}

type EntityRef struct {
	HREF  string  `xml:"Href,attr,omitempty"`
	Type  string  `xml:"Type,attr,omitempty"`
	Name  string  `xml:"Name,attr,omitempty"`
	Links LinkList `xml:"Links,omitempty"`
}

type Task struct {
	HREF  string  `xml:"Href,attr,omitempty"`
	Type  string  `xml:"Type,attr"`
	Links LinkList `xml:"Links,omitempty"`
}

type HierarchyRoot struct {
	HREF  string  `xml:"Href,attr,omitempty"`
	Type  string  `xml:"Type,attr"`
	UID   string  `xml:"UID,attr,omitempty"`
	Name  string  `xml:"Name,attr,omitempty"`
	Links LinkList `xml:"Links,omitempty"`

	HierarchyRootId string `xml:"HierarchyRootId"`
	UniqueId        string `xml:"UniqueId"`
	HostType        string `xml:"HostType"`
}

type Repository struct {
	HREF  string  `xml:"Href,attr,omitempty"`
	Type  string  `xml:"Type,attr"`
	UID   string  `xml:"UID,attr,omitempty"`
	Name  string  `xml:"Name,attr,omitempty"`
	Links LinkList `xml:"Links,omitempty"`

	Capacity  int64  `xml:"Capacity"`
	FreeSpace int64  `xml:"FreeSpace"`
	Kind      string `xml:"Kind"`
}

type Backup struct {
	HREF  string  `xml:"Href,attr,omitempty"`
	Type  string  `xml:"Type,attr"`
	UID   string  `xml:"UID,attr,omitempty"`
	Name  string  `xml:"Name,attr,omitempty"`
	Links LinkList `xml:"Links,omitempty"`

	Platform   string `xml:"Platform"`
	BackupType string `xml:"BackupType"`
}

type RestorePoint struct {
	HREF  string  `xml:"Href,attr,omitempty"`
	Type  string  `xml:"Type,attr"`
	UID   string  `xml:"UID,attr,omitempty"`
	Name  string  `xml:"Name,attr,omitempty"`
	Links LinkList `xml:"Links,omitempty"`

	BackupDateUTC string `xml:"BackupDateUTC"`
}

type VmRestorePoint struct {
	HREF  string  `xml:"Href,attr,omitempty"`
	Type  string  `xml:"Type,attr"`
	UID   string  `xml:"UID,attr,omitempty"`
	Name  string  `xml:"Name,attr,omitempty"`
	Links LinkList `xml:"Links,omitempty"`

	CreationTimeUTC string `xml:"CreationTimeUTC,omitempty"`
	Algorithm       string `xml:"Algorithm"`
	PointType       string `xml:"PointType"`
	HierarchyObjRef string `xml:"HierarchyObjRef"`
}

type VmRestorePointMount struct {
	HREF  string  `xml:"Href,attr,omitempty"`
	Type  string  `xml:"Type,attr"`
	Links LinkList `xml:"Links,omitempty"`

	FSRoot []*DirectoryEntry `xml:"FSRoot"`
}

type DirectoryEntry struct {
	HREF string `xml:"Href,attr,omitempty"`
	Type string `xml:"Type,attr"`

	Path string `xml:"Path"`
	Name string `xml:"Name"`
}

type FileSystemEntries struct {
	HREF        string            `xml:"Href,attr,omitempty"`
	Directories []*DirectoryEntry `xml:"Directories"`
}

type VAppRestorePoint struct {
	HREF            string  `xml:"Href,attr,omitempty"`
	Type            string  `xml:"Type,attr"`
	UID             string  `xml:"UID,attr,omitempty"`
	Name            string  `xml:"Name,attr,omitempty"`
	VAppDisplayName string  `xml:"VAppDisplayName,omitempty"`
	Links           LinkList `xml:"Links,omitempty"`

	CreationTimeUTC string `xml:"CreationTimeUTC,omitempty"`
	Algorithm       string `xml:"Algorithm"`
	PointType       string `xml:"PointType"`
	HierarchyObjRef string `xml:"HierarchyObjRef"`
}

type Replica struct {
	HREF  string  `xml:"Href,attr,omitempty"`
	Type  string  `xml:"Type,attr"`
	UID   string  `xml:"UID,attr,omitempty"`
	Name  string  `xml:"Name,attr,omitempty"`
	Links LinkList `xml:"Links,omitempty"`

	Platform string `xml:"Platform"`
}

type VmReplicaPoint struct {
	HREF  string  `xml:"Href,attr,omitempty"`
	Type  string  `xml:"Type,attr"`
	UID   string  `xml:"UID,attr,omitempty"`
	Name  string  `xml:"Name,attr,omitempty"`
	Links LinkList `xml:"Links,omitempty"`

	CreationTimeUTC string `xml:"CreationTimeUTC,omitempty"`
	Algorithm       string `xml:"Algorithm"`
	PointType       string `xml:"PointType"`
}

type VmReplicaPointMount struct {
	HREF  string  `xml:"Href,attr,omitempty"`
	Type  string  `xml:"Type,attr"`
	Links LinkList `xml:"Links,omitempty"`

	FSRoot []*DirectoryEntry `xml:"FSRoot"`
}

type FileSystemEntry struct {
	DirectoryEntry *DirectoryEntry `xml:"DirectoryEntry"`
}

type CatalogVm struct {
	HREF  string  `xml:"Href,attr,omitempty"`
	Type  string  `xml:"Type,attr"`
	UID   string  `xml:"UID,attr,omitempty"`
	Name  string  `xml:"Name,attr,omitempty"`
	Links LinkList `xml:"Links,omitempty"`
}

type CatalogVmRestorePoint struct {
	HREF  string  `xml:"Href,attr,omitempty"`
	Type  string  `xml:"Type,attr"`
	UID   string  `xml:"UID,attr,omitempty"`
	Name  string  `xml:"Name,attr,omitempty"`
	Links LinkList `xml:"Links,omitempty"`

	BackupDateUTC string `xml:"BackupDateUTC,omitempty"`
}

type BackupJobSession struct {
	HREF  string  `xml:"Href,attr,omitempty"`
	Type  string  `xml:"Type,attr"`
	UID   string  `xml:"UID,attr,omitempty"`
	Name  string  `xml:"Name,attr,omitempty"`
	Links LinkList `xml:"Links,omitempty"`

	JobUid          string `xml:"JobUid"`
	JobName         string `xml:"JobName"`
	JobType         string `xml:"JobType"`
	CreationTimeUTC string `xml:"CreationTimeUTC,omitempty"`
	EndTimeUTC      string `xml:"EndTimeUTC,omitempty"`
	State           string `xml:"State,omitempty"`
	Result          string `xml:"Result,omitempty"`
	Progress        int    `xml:"Progress,omitempty"`
	IsRetry         bool   `xml:"IsRetry"`
}

type RestoreSession struct {
	HREF          string  `xml:"Href,attr,omitempty"`
	Type          string  `xml:"Type,attr"`
	UID           string  `xml:"UID,attr,omitempty"`
	Name          string  `xml:"Name,attr,omitempty"`
	VmDisplayName string  `xml:"VmDisplayName,omitempty"`
	Links         LinkList `xml:"Links,omitempty"`

	JobType         string `xml:"JobType"`
	CreationTimeUTC string `xml:"CreationTimeUTC,omitempty"`
	EndTimeUTC      string `xml:"EndTimeUTC,omitempty"`
	State           string `xml:"State,omitempty"`
	Result          string `xml:"Result,omitempty"`
	Progress        int    `xml:"Progress,omitempty"`
}

type ReplicaJobSession struct {
	HREF  string  `xml:"Href,attr,omitempty"`
	Type  string  `xml:"Type,attr"`
	UID   string  `xml:"UID,attr,omitempty"`
	Name  string  `xml:"Name,attr,omitempty"`
	Links LinkList `xml:"Links,omitempty"`

	JobUid          string `xml:"JobUid"`
	JobName         string `xml:"JobName"`
	JobType         string `xml:"JobType"`
	CreationTimeUTC string `xml:"CreationTimeUTC,omitempty"`
	EndTimeUTC      string `xml:"EndTimeUTC,omitempty"`
	State           string `xml:"State,omitempty"`
	Result          string `xml:"Result,omitempty"`
	Progress        int    `xml:"Progress,omitempty"`
	IsRetry         bool   `xml:"IsRetry"`
}

type BackupTaskSession struct {
	HREF          string  `xml:"Href,attr,omitempty"`
	Type          string  `xml:"Type,attr"`
	UID           string  `xml:"UID,attr,omitempty"`
	Name          string  `xml:"Name,attr,omitempty"`
	VmDisplayName string  `xml:"VmDisplayName,omitempty"`
	Links         LinkList `xml:"Links,omitempty"`

	JobSessionUid   string `xml:"JobSessionUid"`
	CreationTimeUTC string `xml:"CreationTimeUTC,omitempty"`
	EndTimeUTC      string `xml:"EndTimeUTC,omitempty"`
	State           string `xml:"State,omitempty"`
	Result          string `xml:"Result,omitempty"`
	Reason          string `xml:"Reason,omitempty"`
	TotalSize       int64  `xml:"TotalSize"`
}

type ReplicaTaskSession struct {
	HREF          string  `xml:"Href,attr,omitempty"`
	Type          string  `xml:"Type,attr"`
	UID           string  `xml:"UID,attr,omitempty"`
	Name          string  `xml:"Name,attr,omitempty"`
	VmDisplayName string  `xml:"VmDisplayName,omitempty"`
	Links         LinkList `xml:"Links,omitempty"`

	JobSessionUid   string `xml:"JobSessionUid"`
	CreationTimeUTC string `xml:"CreationTimeUTC,omitempty"`
	EndTimeUTC      string `xml:"EndTimeUTC,omitempty"`
	State           string `xml:"State,omitempty"`
	Result          string `xml:"Result,omitempty"`
	Reason          string `xml:"Reason,omitempty"`
	TotalSize       int64  `xml:"TotalSize"`
}

type EnterpriseRole struct {
	HREF  string  `xml:"Href,attr,omitempty"`
	Type  string  `xml:"Type,attr"`
	UID   string  `xml:"UID,attr,omitempty"`
	Name  string  `xml:"Name,attr,omitempty"`
	Links LinkList `xml:"Links,omitempty"`
}

type EnterpriseAccount struct {
	HREF  string  `xml:"Href,attr,omitempty"`
	Type  string  `xml:"Type,attr"`
	UID   string  `xml:"UID,attr,omitempty"`
	Name  string  `xml:"Name,attr,omitempty"`
	Links LinkList `xml:"Links,omitempty"`

	AccountType        string `xml:"AccountType"`
	AllowRestoreAllVms bool   `xml:"AllowRestoreAllVms"`
}

type EnterpriseAccountInRole struct {
	HREF  string  `xml:"Href,attr,omitempty"`
	Links LinkList `xml:"Links,omitempty"`

	RoleName string `xml:"RoleName"`
}

type EnterpriseAccountHierarchyScope struct {
	HREF  string  `xml:"Href,attr,omitempty"`
	Links LinkList `xml:"Links,omitempty"`

	Name                string `xml:"Name"`
	HierarchyRootName   string `xml:"HierarchyRootName"`
	Platform            string `xml:"Platform"`
	HierarchyObjectType string `xml:"HierarchyObjectType"`
	State               string `xml:"State"`
}

type WanAccelerator struct {
	HREF  string  `xml:"Href,attr,omitempty"`
	Type  string  `xml:"Type,attr"`
	UID   string  `xml:"UID,attr,omitempty"`
	Name  string  `xml:"Name,attr,omitempty"`
	Links LinkList `xml:"Links,omitempty"`

	Description      string  `xml:"Description,omitempty"`
	OutOfDate        bool    `xml:"OutOfDate"`
	Version          float32 `xml:"Version"`
	Capacity         int64   `xml:"Capacity"`
	TrafficPort      int     `xml:"TrafficPort"`
	ConnectionsCount int     `xml:"ConnectionsCount"`
	CachePath        string  `xml:"CachePath"`
}

type VCloudOrganizationConfig struct {
	HREF  string  `xml:"Href,attr,omitempty"`
	Type  string  `xml:"Type,attr"`
	UID   string  `xml:"UID,attr,omitempty"`
	Name  string  `xml:"Name,attr,omitempty"`
	Links LinkList `xml:"Links,omitempty"`

	BackupServerUid string       `xml:"BackupServerUid"`
	RepositoryUid   string       `xml:"RepositoryUid"`
	QuotaGb         int64        `xml:"QuotaGb"`
	JobSettings     *JobSettings `xml:"JobSettings"`
	IsDisabled      bool         `xml:"IsDisabled"`
}

type JobSettings struct {
	CustomSettings   bool   `xml:"CustomSettings"`
	JobSchedulerType string `xml:"JobSchedulerType"`
}

type BackupFile struct {
	HREF  string  `xml:"Href,attr,omitempty"`
	Type  string  `xml:"Type,attr"`
	UID   string  `xml:"UID,attr,omitempty"`
	Name  string  `xml:"Name,attr,omitempty"`
	Links LinkList `xml:"Links,omitempty"`

	FilePath           string `xml:"FilePath"`
	BackupSize         int64  `xml:"BackupSize"`
	DataSize           int64  `xml:"DataSize"`
	DeduplicationRatio int    `xml:"DeduplicationRatio"`
	CompressRatio      int    `xml:"CompressRatio"`
	CreationTimeUtc    string `xml:"CreationTimeUtc"`
	FileType           string `xml:"FileType"`
}

type ExternalRepository struct {
	HREF  string  `xml:"Href,attr,omitempty"`
	Type  string  `xml:"Type,attr"`
	UID   string  `xml:"UID,attr,omitempty"`
	Name  string  `xml:"Name,attr,omitempty"`
	Links LinkList `xml:"Links,omitempty"`

	RepositoryType string `xml:"RepositoryType"`
	Path           string `xml:"Path"`
	UsedSpace      int64  `xml:"UsedSpace"`
	Description    string `xml:"Description,omitempty"`
}

type AgentBackupJob struct {
	HREF  string  `xml:"Href,attr,omitempty"`
	Type  string  `xml:"Type,attr"`
	UID   string  `xml:"UID,attr,omitempty"`
	Name  string  `xml:"Name,attr,omitempty"`
	Links LinkList `xml:"Links,omitempty"`

	JobType            string              `xml:"JobType"`
	Platform           string              `xml:"Platform"`
	Description        string              `xml:"Description,omitempty"`
	ScheduleConfigured bool                `xml:"ScheduleConfigured"`
	ScheduleEnabled    bool                `xml:"ScheduleEnabled"`
	JobScheduleOptions *JobScheduleOptions `xml:"JobScheduleOptions"`
	JobInfo            *JobInfo            `xml:"JobInfo"`
}

type AgentRestorePoint struct {
	HREF  string  `xml:"Href,attr,omitempty"`
	Type  string  `xml:"Type,attr"`
	UID   string  `xml:"UID,attr,omitempty"`
	Name  string  `xml:"Name,attr,omitempty"`
	Links LinkList `xml:"Links,omitempty"`

	CreationTimeUTC string `xml:"CreationTimeUTC,omitempty"`
	Algorithm       string `xml:"Algorithm"`
	PointType       string `xml:"PointType"`
}

type AgentRestorePointMount struct {
	HREF string `xml:"Href,attr,omitempty"`
	Type string `xml:"Type,attr"`

	FSRoots []*DirectoryEntry `xml:"FSRoots"`
}

type AgentProtectionGroup struct {
	HREF  string  `xml:"Href,attr,omitempty"`
	Type  string  `xml:"Type,attr"`
	UID   string  `xml:"UID,attr,omitempty"`
	Name  string  `xml:"Name,attr,omitempty"`
	Links LinkList `xml:"Links,omitempty"`

	RescanScheduleEnabled bool `xml:"RescanScheduleEnabled"`
}

type DiscoveredComputer struct {
	HREF  string  `xml:"Href,attr,omitempty"`
	Type  string  `xml:"Type,attr"`
	UID   string  `xml:"UID,attr,omitempty"`
	Name  string  `xml:"Name,attr,omitempty"`
	Links LinkList `xml:"Links,omitempty"`

	HostStatus   string `xml:"HostStatus"`
	AgentVersion string `xml:"AgentVersion"`
	AgentStatus  string `xml:"AgentStatus"`
	OsVersion    string `xml:"OsVersion"`
	IpAddress    string `xml:"IpAddress"`
}

type VSphereSelfServiceConfig struct {
	HREF  string  `xml:"Href,attr,omitempty"`
	Type  string  `xml:"Type,attr"`
	UID   string  `xml:"UID,attr,omitempty"`
	Name  string  `xml:"Name,attr,omitempty"`
	Links LinkList `xml:"Links,omitempty"`

	BackupServerUid string       `xml:"BackupServerUid"`
	RepositoryUid   string       `xml:"RepositoryUid"`
	QuotaGb         int64        `xml:"QuotaGb"`
	PerUser         bool         `xml:"PerUser"`
	Account         *Account     `xml:"Account"`
	JobSettings     *JobSettings `xml:"JobSettings"`
	Tags            string       `xml:"Tags"`
}

type Account struct {
	AccountName string `xml:"AccountName"`
	AccountSid  string `xml:"AccountSid"`
	AccountType string `xml:"AccounType"`
}

type CloudGateway struct {
	HREF  string  `xml:"Href,attr,omitempty"`
	Type  string  `xml:"Type,attr"`
	UID   string  `xml:"UID,attr,omitempty"`
	Name  string  `xml:"Name,attr,omitempty"`
	Links LinkList `xml:"Links,omitempty"`

	Enabled      bool   `xml:"Enabled"`
	NetworkMode  string `xml:"NetworkMode"`
	ExternalIP   string `xml:"ExternalIP"`
	ExternalPort int    `xml:"ExternalPort"`
	InternalPort int    `xml:"InternalPort"`
	Description  string `xml:"Description,omitempty"`
}

type CloudTenant struct {
	HREF  string  `xml:"Href,attr,omitempty"`
	Type  string  `xml:"Type,attr"`
	UID   string  `xml:"UID,attr,omitempty"`
	Name  string  `xml:"Name,attr,omitempty"`
	Links LinkList `xml:"Links,omitempty"`

	Password                string           `xml:"Password,omitempty"`
	Description             string           `xml:"Description,omitempty"`
	Enabled                 bool             `xml:"Enabled"`
	LeaseOptions            bool             `xml:"LeaseOptions,attr,omitempty"`
	Resources               []*Resources     `xml:"Resources,omitempty"`
	LastResult              string           `xml:"LastResult,omitempty"`
	LastActive              string           `xml:"LastActive,omitempty"`
	ComputeResources        *ComputeResource `xml:"ComputeResources,omitempty"`
	ThrottlingEnabled       bool             `xml:"ThrottlingEnabled"`
	ThrottlingSpeedLimit    int              `xml:"ThrottlingSpeedLimit"`
	ThrottlingSpeedUnit     string           `xml:"ThrottlingSpeedUnit"`
	PublicIpCount           int              `xml:"PublicIpCount"`
	BackupCount             int              `xml:"BackupCount"`
	ReplicaCount            int              `xml:"ReplicaCount"`
	MaxConcurrentTasks      int              `xml:"MaxConcurrentTasks"`
	WorkStationBackupCount  int              `xml:"WorkStationBackupCount"`
	ServerBackupCount       int              `xml:"ServerBackupCount"`
	BackupProtectionEnabled bool             `xml:"BackupProtectionEnabled"`
	BackupProtectionPeriod  int              `xml:"BackupProtectionPeriod"`
	TenatType               *TenantType      `xml:"TenantType"`
}

type Resources struct {
	CloudTenantResources []*CloudTenantResource `xml:"CloudTenantResource"`
}

type CloudTenantResource struct {
	HREF string `xml:"Href,attr,omitempty"`
	Type string `xml:"Type,attr"`
	UID  string `xml:"UID,attr,omitempty"`

	RepositoryQuota *RepositoryQuota `xml:"RepositoryQuota"`
}

type RepositoryQuota struct {
	DisplayName       string `xml:"DisplayName"`
	RepositoryUid     string `xml:"RepositoryUid"`
	WanAcceleratorUid string `xml:"WanAcceleratorUid"`
	Quota             int64  `xml:"Quota"`
}

type ComputeResource struct {
	CloudTenantComputeResource *CloudTenantComputeResource `xml:"CloudTenantComputeResource"`
}

// TODO: Change Id to UID or ID to follow established convention
type CloudTenantComputeResource struct {
	HREF string `xml:"Href,attr,omitempty"`
	Type string `xml:"Type,attr"`
	Id   string `xml:"Id,attr,omitempty"`

	CloudHardwarePlanUid        string                `xml:"CloudHardwarePlanUid"`
	WanAcceleratorUid           string                `xml:"WanAcceleratorUid"`
	PlatformType                string                `xml:"PlatformType"`
	UseNetworkFailoverResources bool                  `xml:"UseNetworkFailoverResources"`
	NetworkAppliance            *NetworkAppliance     `xml:"NetworkAppliance"`
	ComputeResourceStats        *ComputeResourceStats `xml:"ComputeResourceStats"`
}

type NetworkAppliance struct {
	Name                         string `xml:"Name"`
	ProductionNetwork            string `xml:"ProductionNetwork"`
	ObtainIPAddressAutomatically bool   `xml:"ObtainIPAddressAutomatically"`
	ViDistributedSwitchUuid      string `xml:"ViDistributedSwitchUuid"`
	ProductionNetworkUnderDvs    bool   `xml:"ProductionNetworkUnderDvs"`
}

type ComputeResourceStats struct {
	MemoryUsageMb        int32                  `xml:"MemoryUsageMb"`
	CPUCount             int                    `xml:"CPUCount"`
	StorageResourceStats []*StorageResourceStat `xml:"StorageResourceStat"`
}

type StorageResourceStat struct {
	StorageName    string `xml:"StorageName"`
	StorageUsageGb int    `xml:"StorageUsageGb"`
	StorageLimitGb int    `xml:"StorageUsageGb"`
}

type TenantType struct {
	StandaloneTenant *StandaloneTenant `xml:"StandaloneTenant"`
}

type StandaloneTenant struct {
	TenantCredentials *TenantCredentials `xml:"TenantCredentials"`
}

type TenantCredentials struct {
	Username string `xml:"Username"`
}

type CloudTenantVCloudComputeResource struct {
	HREF  string  `xml:"Href,attr,omitempty"`
	Type  string  `xml:"Type,attr"`
	Id    string  `xml:"Id,attr,omitempty"`
	Links LinkList `xml:"Links,omitempty"`

	VirtualDataCenterName       string            `xml:"VirtualDataCenterName"`
	VirtualDataCenterRef        string            `xml:"VirtualDataCenterRef"`
	Enabled                     bool              `xml:"Enabled"`
	AllocationModel             string            `xml:"AllocationModel"`
	UseNetworkFailoverResources bool              `xml:"UseNetworkFailoverResources"`
	ResourceUsage               *ResourceUsage    `xml:"ResourceUsage"`
	WanAcceleratorUid           string            `xml:"WanAcceleratorUid"`
	NetworkAppliance            *NetworkAppliance `xml:"NetworkAppliance"`
}

type ResourceUsage struct {
	CpuUsageMhz    int `xml:"CpuUsageMhz"`
	MemoryUsageMb  int `xml:"MemoryUsageMb"`
	StorageUsageGb int `xml:"StorageUsageGb"`
}

type CloudSubtenant struct {
	HREF  string  `xml:"Href,attr,omitempty"`
	Type  string  `xml:"Type,attr"`
	Id    string  `xml:"Id,attr,omitempty"`
	Links LinkList `xml:"Links,omitempty"`

	Name            string           `xml:"Name"`
	Description     string           `xml:"Description,omitempty"`
	Password        string           `xml:"Password,omitempty"`
	Enabled         bool             `xml:"Enabled"`
	RepositoryQuota *RepositoryQuota `xml:"RepositoryQuota"`
}

type CloudHardwarePlan struct {
	HREF  string  `xml:"Href,attr,omitempty"`
	Type  string  `xml:"Type,attr"`
	UID   string  `xml:"UID,attr,omitempty"`
	Links LinkList `xml:"Links,omitempty"`

	Description            string               `xml:"Description,omitempty"`
	ProcessorUsageLimitMhz int                  `xml:"ProcessorUsageLimitMhz"`
	MemoryUsageLimitMb     int                  `xml:"MemoryUsageLimitMb"`
	HardwarePlanDetails    *HardwarePlanDetails `xml:"HardwarePlanDetails"`
}

type HardwarePlanDetails struct {
	HvCloudHardwarePlan *HvCloudHardwarePlan `xml:"HvCloudHardwarePlan"`
}

type HvCloudHardwarePlan struct {
	HypervisorHostRef string    `xml:"HypervisorHostRef"`
	Volumes           []*Volume `xml:"Volumes,omitempty"`
	Network           *Network  `xml:"Network,omitempty"`
}

type Volume struct {
	Id           string `xml:"Id,attr"`
	FriendlyName string `xml:"FriendlyName"`
	VolumePath   string `xml:"VolumePath"`
	QuotaGb      int    `xml:"QuotaGb"`
}

type Network struct {
	Id                   string `xml:"Id,attr"`
	CountWithInternet    int    `xml:"CountWithInternet"`
	CountWithoutInternet int    `xml:"CountWithoutInternet"`
}

type CloudPublicIpAddress struct {
	HREF            string  `xml:"Href,attr,omitempty"`
	Type            string  `xml:"Type,attr"`
	UID             string  `xml:"UID,attr,omitempty"`
	BackupServerUid string  `xml:"BackupServerUid"`
	Links           LinkList `xml:"Links,omitempty"`

	IpAddress string `xml:"IpAddress"`
}

type CloudFailoverPlan struct {
	HREF  string  `xml:"Href,attr,omitempty"`
	Type  string  `xml:"Type,attr"`
	Name  string  `xml:"Name,attr"`
	UID   string  `xml:"UID,attr,omitempty"`
	Links LinkList `xml:"Links,omitempty"`

	TenantUid                string                    `xml:"TenantUid"`
	TenantName               string                    `xml:"TenantName"`
	Description              string                    `xml:"Description,omitempty"`
	CloudFailoverPlanOptions *CloudFailoverPlanOptions `xml:"CloudFailoverPlanOptions"`
	CloudFailoverPlanInfo    *CloudFailoverPlanInfo    `xml:"CloudFailoverPlanInfo"`
}

type CloudFailoverPlanOptions struct {
	PostFailoverPlanCommandEnabled bool   `xml:"PostFailoverPlanCommandEnabled"`
	PostFailoverPlanCommand        string `xml:"PostFailoverPlanCommand,omitempty"`
	PreFailoverPlanCommandEnabled  bool   `xml:"PreFailoverPlanCommandEnabled"`
	PreFailoverPlanCommand         string `xml:"PreFailoverPlanCommand,omitempty"`
}

// TODO: Error here overriding the Includes
type CloudFailoverPlanInfo struct {
	Includes *Includes `xml:"Includes"`
}

type CloudFailoverPlanVm struct {
	HREF string `xml:"Href,attr,omitempty"`
	Type string `xml:"Type,attr"`

	FailOverPlanVMId string `xml:"FailOverPlanId"`
	Name             string `xml:"Name"`
	Order            int    `xml:"Order"`
}

type CloudFailoveredVm struct {
	HREF  string  `xml:"Href,attr,omitempty"`
	Type  string  `xml:"Type,attr"`
	Links LinkList `xml:"Links,omitempty"`

	FailoverPlanVMId string `xml:"FailoverPlanVMId"`
	Name             string `xml:"Name"`
	Order            int    `xml:"Order"`
}

type CloudVmReplicaPoint struct {
	HREF          string  `xml:"Href,attr,omitempty"`
	Type          string  `xml:"Type,attr"`
	VmDisplayName string  `xml:"VmDisplayName"`
	Links         LinkList `xml:"Links,omitempty"`

	CreationTimeUTC string `xml:"CreationTimeUTC"`
	PointType       string `xml:"PointType"`
	State           string `xml:"State"`
}

type CloudReplica struct {
	HREF  string  `xml:"Href,attr,omitempty"`
	Type  string  `xml:"Type,attr"`
	Name  string  `xml:"Name"`
	UID   string  `xml:"UID"`
	Links LinkList `xml:"Links,omitempty"`

	Platform string `xml:"Platform"`
}

type VlanConfiguration struct {
	UID   string  `xml:"UID"`
	Name  string  `xml:"Name"`
	Links LinkList `xml:"Links,omitempty"`

	HostRef                          string `xml:"HostRef"`
	PlatformType                     string `xml:"PlatformType"`
	VlanIdsWithInternetLeftBound     int    `xml:"VlanIdsWithInternetLeftBound"`
	VlanIdsWithInternetRightBound    int    `xml:"VlanIdsWithInternetRightBound"`
	VlanIdsWithoutInternetLeftBound  int    `xml:"VlanIdsWithoutInternetLeftBound"`
	VlanIdsWithoutInternetRightBound int    `xml:"VlanIdsWithoutInternetRightBound"`
	SwitchName                       string `xml:"SwitchName"`
	SwitchId                         string `xml:"SwitchId"`
	SwitchType                       string `xml:"SwitchType"`
}

type CloudFailoverSession struct {
	HREF  string  `xml:"Href,attr,omitempty"`
	Type  string  `xml:"Type,attr"`
	Name  string  `xml:"Name"`
	UID   string  `xml:"UID"`
	Links LinkList `xml:"Links,omitempty"`

	JobType            string              `xml:"JobType"`
	CreationTimeUTC    string              `xml:"CreationTimeUTC"`
	EndTimeUTC         string              `xml:"EndTimeUTC"`
	State              string              `xml:"State"`
	Result             string              `xml:"Result"`
	Progress           int                 `xml:"Progress"`
	CloudFailoverTasks *CloudFailoverTasks `xml:"CloudFailoverTasks"`
}

// TODO: see if this actually works
type CloudFailoverTasks struct {
	Tasks []*Task `xml:"Task"`
}

type CloudGatewayPool struct {
	HREF  string  `xml:"Href,attr,omitempty"`
	Type  string  `xml:"Type,attr"`
	Name  string  `xml:"Name"`
	UID   string  `xml:"UID"`
	Links LinkList `xml:"Links,omitempty"`

	Description   string          `xml:"Description,omitempty"`
	CloudGateways []*CloudGateway `xml:"CloudGateways,omitempty"`
	CloudTenants  []*CloudTenant  `xml:"CloudTenants"`
}

type LogonSession struct {
	HREF  string  `xml:"Href,attr,omitempty"`
	Type  string  `xml:"Type,attr"`
	Links LinkList `xml:"Links,omitempty"`

	Username  string `xml:"Username"`
	SessionId string `xml:"SessionId"`
}

type Error struct {
	Message                 string `xml:"message,attr"`
	MajorErrorCode          int    `xml:"majorErrorCode,attr"`
	MinorErrorCode          string `xml:"minorErrorCode,attr"`
	VendorSpecificErrorCode string `xml:"vendorSpecificErrorCode,attr,omitempty"`
	StackTrace              string `xml:"stackTrace,attr,omitempty"`
}

func (err Error) Error() string {
	return fmt.Sprintf("API Error: %d: %s", err.MajorErrorCode, err.Message)
}
