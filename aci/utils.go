package aci

import (
	"context"
	"fmt"

	"steampipe-plugin-aci/client"

	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func connect(ctx context.Context, d *plugin.QueryData) (*client.Client, error) {
	aciConfig := GetConfig(d.Connection)

	// Initial values. Env vars will be overridden by configuration if values are set in there
	//vsphereServer := os.Getenv("VSPHERE_SERVER")
	//user := os.Getenv("VSPHERE_USER")
	//password := os.Getenv("VSPHERE_PASSWORD")
	allowUnverifiedSSL := false
	clusterURI := ""
	user := ""
	password := ""
	loginDomain := "DefaultAuth"

	// Override potential env values with config values
	if aciConfig.AllowUnverifiedSSL != nil {
		allowUnverifiedSSL = *aciConfig.AllowUnverifiedSSL
	}

	if aciConfig.ClusterURI != nil {
		clusterURI = *aciConfig.ClusterURI
	}

	if aciConfig.User != nil {
		user = *aciConfig.User
	}

	if aciConfig.Password != nil {
		password = *aciConfig.Password
	}

	if aciConfig.LoginDomain != nil {
		loginDomain = *aciConfig.LoginDomain
	}
	// Make sure we have all required arguments set via either env or config
	if clusterURI == "" || user == "" || password == "" || loginDomain == "" {
		errorMsg := ""
		if clusterURI == "" {
			errorMsg += "Missing clusterURI from config'\n"
		}
		if user == "" {
			errorMsg += "Missing user from config'\n"
		}
		if password == "" {
			errorMsg += "Missing password from config'\n"
		}
		if loginDomain == "" {
			errorMsg += "Missing loginDomain from config'\n"
		}
		return nil, fmt.Errorf("Error in configuraiton: %s", errorMsg)
	}

	aciclient := client.GetClient(clusterURI, user, client.Password(password), client.Insecure(allowUnverifiedSSL))

	return aciclient, nil
}
