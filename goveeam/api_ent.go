package goveeam

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
	"sync"
	"time"
)

// VeeamClientOption defines signature for customizing VeeamClient using
// functional options pattern.
type VeeamClientOption func(*VeeamClient) error

type VeeamClient struct {
	Client            Client  // Client for the underlying Ent instance
	sessionHREF       url.URL // HREF for the session API
	QueryHREF         url.URL // HREF for the query API
	Mutex             sync.Mutex
	supportedVersions SupportedVersions // Versions from /api endpoint
}

func (veeamCli *VeeamClient) veeamloginurl() error {
	if err := veeamCli.validateAPIVersion(); err != nil {
		return fmt.Errorf("could not find valid version for login: %s", err)
	}

	// find login address matching the API version
	var neededVersion VersionInfo
	for _, versionInfo := range veeamCli.supportedVersions.SupportedVersions.Versions {
		if versionInfo.Name == veeamCli.Client.APIVersion {
			neededVersion = versionInfo
			break
		}
	}

	loginUrl, err := url.Parse(neededVersion.Links[0].Link.HREF)
	if err != nil {
		return fmt.Errorf("couldn't find a LoginUrl for version %s", veeamCli.Client.APIVersion)
	}

	veeamCli.sessionHREF = *loginUrl
	return nil
}

func (veeamCli *VeeamClient) veeamauthorize(user, pass string) error {
	var missingItems []string
	if user == "" {
		missingItems = append(missingItems, "user")
	}
	if pass == "" {
		missingItems = append(missingItems, "password")
	}
	if len(missingItems) > 0 {
		return fmt.Errorf("authorization is not possible because of these missing items: %v", missingItems)
	}
	// No point in checking for errors here
	req := veeamCli.Client.NewRequest(map[string]string{}, http.MethodPost, veeamCli.sessionHREF, nil)
	// Set Basic Authentication Header
	req.SetBasicAuth(user, pass)
	// Add the Accept header for vCA
	req.Header.Add("Accept", "application/xml")
	resp, err := checkResp(veeamCli.Client.Http.Do(req))
	if err != nil {
		return fmt.Errorf("unable to check resposne: %s", err)
	}
	defer resp.Body.Close()
	// Store the authentication header
	veeamCli.Client.VeeamEntToken = resp.Header.Get("X-RestSvcSessionId")
	veeamCli.Client.VeeamEntAuthHeader = "X-RestSvcSessionId"

	// Get query href
	veeamCli.QueryHREF = veeamCli.Client.ENTHREF
	veeamCli.QueryHREF.Path += "/query"
	return nil
}

// NewVeeamClient initializes Veeam Enterprise Manager client with reasonable defaults.
// It accepts functions of type VeeamClientOption for adjusting defaults.
func NewVeeamClient(veeamEndpoint url.URL, insecure bool, options ...VeeamClientOption) *VeeamClient {
	// Setting defaults
	veeamClient := &VeeamClient{
		Client: Client{
			APIVersion: "v1_4",
			ENTHREF:    veeamEndpoint,
			Http: http.Client{
				Transport: &http.Transport{
					TLSClientConfig: &tls.Config{
						InsecureSkipVerify: insecure,
					},
					Proxy:               http.ProxyFromEnvironment,
					TLSHandshakeTimeout: 120 * time.Second,
				},
			},
			MaxRetryTimeout: 60, // Default timeout in seconds for Client
		},
	}

	// Override defaults with functional options
	for _, option := range options {
		err := option(veeamClient)
		if err != nil {
			// We do not have error in return of this function signature.
			// To avoid breaking API the only thing we can do is panic.
			panic(fmt.Sprintf("unable to initialize veem ent client: %s", err))
		}
	}
	return veeamClient
}

// Authenticate is an helper function that performs a login in vCloud Director.
func (veeamCli *VeeamClient) Authenticate(username, password string) error {

	// LoginUrl
	err := veeamCli.veeamloginurl()
	if err != nil {
		return fmt.Errorf("error finding LoginUrl: %s", err)
	}
	// Authorize
	err = veeamCli.veeamauthorize(username, password)
	if err != nil {
		return fmt.Errorf("error authorizing: %s", err)
	}
	return nil
}

// Disconnect performs a disconnection from the vCloud Director API endpoint.
func (veeamCli *VeeamClient) Disconnect() error {
	if veeamCli.Client.VeeamEntToken == "" && veeamCli.Client.VeeamEntAuthHeader == "" {
		return fmt.Errorf("cannot disconnect, client is not authenticated")
	}
	req := veeamCli.Client.NewRequest(map[string]string{}, http.MethodDelete, veeamCli.sessionHREF, nil)
	// Add the Accept header for ent mgr
	req.Header.Add("Accept", "application/xml;version="+veeamCli.Client.APIVersion)
	// Set Authorization Header
	req.Header.Add(veeamCli.Client.VeeamEntAuthHeader, veeamCli.Client.VeeamEntToken)
	if _, err := checkResp(veeamCli.Client.Http.Do(req)); err != nil {
		return fmt.Errorf("error processing session delete for veeam enterprise manager: %s", err)
	}
	return nil
}

// WithAPIVersion allows to override default API version. Please be cautious
// about changing the version as the default specified is the most tested.
func WithAPIVersion(version string) VeeamClientOption {
	return func(veeamClient *VeeamClient) error {
		veeamClient.Client.APIVersion = version
		return nil
	}
}

// WithMaxRetryTimeout allows default vCDClient MaxRetryTimeout value override
func WithMaxRetryTimeout(timeoutSeconds int) VeeamClientOption {
	return func(veeamClient *VeeamClient) error {
		veeamClient.Client.MaxRetryTimeout = timeoutSeconds
		return nil
	}
}