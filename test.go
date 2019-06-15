package main


import (
	"fmt"
	"net/url"
	"os"

	"github.com/goveeam/goveeam"
)

type Config struct {
	User     string
	Password string
	Org      string
	Href     string
	VDC      string
	Insecure bool
}

func (c *Config) Client() (*goveeam.VeeamClient, error) {
	u, err := url.ParseRequestURI(c.Href)
	if err != nil {
		return nil, fmt.Errorf("unable to pass url: %s", err)
	}

	veeamclient := goveeam.NewVeeamClient(*u, c.Insecure)
	err = veeamclient.Authenticate(c.User, c.Password, c.Org)
	if err != nil {
		return nil, fmt.Errorf("unable to authenticate: %s", err)
	}
	return veeamclient, nil
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Syntax: example user password VEEAM_IP ")
		os.Exit(1)
	}
	config := Config{
		User:     os.Args[1],
		Password: os.Args[2],
		Href:     fmt.Sprintf("http://%s/api", os.Args[3]),
		Insecure: true,
	}

	client, err := config.Client() // We now have a client
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Print(client.Client.ENTHREF)
}