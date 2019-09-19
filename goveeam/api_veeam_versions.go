package goveeam

import (
	"fmt"
	"github.com/hasnhasan/goveeam/types/v1"
	"github.com/hasnhasan/goveeam/util"
	"net/http"
)

type VersionInfo struct {
	Name  string         `xml:"Name,attr"`
	Links types.LinkList `xml:"Links"`
}

type VersionInfos struct {
	Versions []VersionInfo `xml:"SupportedVersion"`
}

type SupportedVersions struct {
	Links             types.LinkList `xml:"Links"`
	SupportedVersions VersionInfos   `xml:"SupportedVersions"`
}

// validateAPIVersion fetches API versions
func (veeamCli *VeeamClient) validateAPIVersion() error {
	err := veeamCli.veeamFetchSupportedVersions()
	if err != nil {
		return fmt.Errorf("could not retrieve supported versions: %s", err)
	}

	// Check if version is supported
	if ok, err := veeamCli.veeamCheckSupportedVersion(veeamCli.Client.APIVersion); !ok || err != nil {
		return fmt.Errorf("API version %s is not supported: %s", veeamCli.Client.APIVersion, err)
	}

	return nil
}

// veeamFetchSupportedVersions retrieves list of supported versions from
// /api endpoint and stores them in VeeamClient for future uses.
// It only does it once.
func (veeamCli *VeeamClient) veeamFetchSupportedVersions() error {
	// Only fetch Versions if it is not already stored
	if veeamCli.supportedVersions.SupportedVersions.Versions != nil {
		numVersions := len(veeamCli.supportedVersions.SupportedVersions.Versions)
		util.Logger.Printf("[TRACE] skipping fetch of versions because %d are stored", numVersions)
		return nil
	}

	apiEndpoint := veeamCli.Client.ENTHREF

	entMngr := new(SupportedVersions)
	_, err := veeamCli.Client.ExecuteRequest(apiEndpoint.String(), http.MethodGet,
		"", "error fetching versions: %s", nil, entMngr)

	veeamCli.supportedVersions = *entMngr
	return err
}

// veeamCheckSupportedVersion checks if there is at least one specified version exactly matching listed ones.
// Format example "v1_1"
func (veeamCli *VeeamClient) veeamCheckSupportedVersion(version string) (bool, error) {
	return veeamCli.checkSupportedVersionConstraint(fmt.Sprintf("= %s", version))
}

// Checks if there is at least one specified version matching the list returned by Veeam.
// Constraint format can be in format ">= 1_1, < 1_4",">= 1_4" ,"= 1_4".
func (veeamCli *VeeamClient) checkSupportedVersionConstraint(versionConstraint string) (bool, error) {
	for _, versionInfo := range veeamCli.supportedVersions.SupportedVersions.Versions {
		versionMatch, err := veeamCli.apiVersionMatchesConstraint(versionInfo.Name, versionConstraint)
		if err != nil {
			return false, fmt.Errorf("cannot match version: %s", err)
		}

		if versionMatch {
			return true, nil
		}
	}
	return false, fmt.Errorf("version %s is not supported", versionConstraint)
}

func (veeamCli *VeeamClient) apiVersionMatchesConstraint(version, versionConstraint string) (bool, error) {

	// TODO: Veeam doesnt use dot notation for the version numbers of the ENT Mngr, write check for this
	return true, nil
	/*
		checkVer, err := semver.NewVersion(version)
		if err != nil {
			return false, fmt.Errorf("[ERROR] unable to parse version %s : %s", version, err)
		}
		// Create a provided constraint to check against current max version
		constraints, err := semver.NewConstraint(versionConstraint)
		if err != nil {
			return false, fmt.Errorf("[ERROR] unable to parse given version constraint '%s' : %s", versionConstraint, err)
		}
		if constraints.Check(checkVer) {
			return true, nil
		}

		util.Logger.Printf("[TRACE] API version %s does not satisfy constraints '%s'", checkVer, constraints)
		return false, nil

	*/
}
