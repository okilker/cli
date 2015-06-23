package plugin_models

type SpaceQuotaFields struct {
	Guid                    string
	Name                    string
	MemoryLimit             int64
	InstanceMemoryLimit     int64
	RoutesLimit             int
	ServicesLimit           int
	NonBasicServicesAllowed bool
}
