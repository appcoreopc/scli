package Services

type VersioningService struct {
	CurrentVersion string
}

func (v *VersioningService) CheckVersion(InstalledVersion, CurrentVersion string) {

}

type IVersioningService interface {
	CheckVersion(InstalledVersion, CurrentVersion string)
}
