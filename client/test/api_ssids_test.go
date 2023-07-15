/*
Meraki Dashboard API

Testing SsidsApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package client

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/meraki/dashboard-api-go/client"
)

func Test_client_SsidsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SsidsApiService CreateNetworkWirelessSsidIdentityPsk", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var number string

		resp, httpRes, err := apiClient.SsidsApi.CreateNetworkWirelessSsidIdentityPsk(context.Background(), networkId, number).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SsidsApiService DeleteNetworkWirelessSsidIdentityPsk", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var number string
		var identityPskId string

		httpRes, err := apiClient.SsidsApi.DeleteNetworkWirelessSsidIdentityPsk(context.Background(), networkId, number, identityPskId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SsidsApiService GetNetworkApplianceSsid", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var number string

		resp, httpRes, err := apiClient.SsidsApi.GetNetworkApplianceSsid(context.Background(), networkId, number).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SsidsApiService GetNetworkApplianceSsids", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.SsidsApi.GetNetworkApplianceSsids(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SsidsApiService GetNetworkWirelessSsid", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var number string

		resp, httpRes, err := apiClient.SsidsApi.GetNetworkWirelessSsid(context.Background(), networkId, number).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SsidsApiService GetNetworkWirelessSsidBonjourForwarding", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var number string

		resp, httpRes, err := apiClient.SsidsApi.GetNetworkWirelessSsidBonjourForwarding(context.Background(), networkId, number).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SsidsApiService GetNetworkWirelessSsidDeviceTypeGroupPolicies", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var number string

		resp, httpRes, err := apiClient.SsidsApi.GetNetworkWirelessSsidDeviceTypeGroupPolicies(context.Background(), networkId, number).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SsidsApiService GetNetworkWirelessSsidEapOverride", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var number string

		resp, httpRes, err := apiClient.SsidsApi.GetNetworkWirelessSsidEapOverride(context.Background(), networkId, number).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SsidsApiService GetNetworkWirelessSsidFirewallL3FirewallRules", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var number string

		resp, httpRes, err := apiClient.SsidsApi.GetNetworkWirelessSsidFirewallL3FirewallRules(context.Background(), networkId, number).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SsidsApiService GetNetworkWirelessSsidFirewallL7FirewallRules", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var number string

		resp, httpRes, err := apiClient.SsidsApi.GetNetworkWirelessSsidFirewallL7FirewallRules(context.Background(), networkId, number).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SsidsApiService GetNetworkWirelessSsidHotspot20", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var number string

		resp, httpRes, err := apiClient.SsidsApi.GetNetworkWirelessSsidHotspot20(context.Background(), networkId, number).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SsidsApiService GetNetworkWirelessSsidIdentityPsk", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var number string
		var identityPskId string

		resp, httpRes, err := apiClient.SsidsApi.GetNetworkWirelessSsidIdentityPsk(context.Background(), networkId, number, identityPskId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SsidsApiService GetNetworkWirelessSsidIdentityPsks", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var number string

		resp, httpRes, err := apiClient.SsidsApi.GetNetworkWirelessSsidIdentityPsks(context.Background(), networkId, number).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SsidsApiService GetNetworkWirelessSsidSchedules", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var number string

		resp, httpRes, err := apiClient.SsidsApi.GetNetworkWirelessSsidSchedules(context.Background(), networkId, number).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SsidsApiService GetNetworkWirelessSsidSplashSettings", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var number string

		resp, httpRes, err := apiClient.SsidsApi.GetNetworkWirelessSsidSplashSettings(context.Background(), networkId, number).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SsidsApiService GetNetworkWirelessSsidTrafficShapingRules", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var number string

		resp, httpRes, err := apiClient.SsidsApi.GetNetworkWirelessSsidTrafficShapingRules(context.Background(), networkId, number).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SsidsApiService GetNetworkWirelessSsidVpn", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var number string

		resp, httpRes, err := apiClient.SsidsApi.GetNetworkWirelessSsidVpn(context.Background(), networkId, number).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SsidsApiService GetNetworkWirelessSsids", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.SsidsApi.GetNetworkWirelessSsids(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SsidsApiService GetOrganizationSummaryTopSsidsByUsage", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string

		resp, httpRes, err := apiClient.SsidsApi.GetOrganizationSummaryTopSsidsByUsage(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SsidsApiService UpdateNetworkApplianceSsid", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var number string

		resp, httpRes, err := apiClient.SsidsApi.UpdateNetworkApplianceSsid(context.Background(), networkId, number).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SsidsApiService UpdateNetworkWirelessSsid", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var number string

		resp, httpRes, err := apiClient.SsidsApi.UpdateNetworkWirelessSsid(context.Background(), networkId, number).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SsidsApiService UpdateNetworkWirelessSsidBonjourForwarding", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var number string

		resp, httpRes, err := apiClient.SsidsApi.UpdateNetworkWirelessSsidBonjourForwarding(context.Background(), networkId, number).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SsidsApiService UpdateNetworkWirelessSsidDeviceTypeGroupPolicies", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var number string

		resp, httpRes, err := apiClient.SsidsApi.UpdateNetworkWirelessSsidDeviceTypeGroupPolicies(context.Background(), networkId, number).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SsidsApiService UpdateNetworkWirelessSsidEapOverride", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var number string

		resp, httpRes, err := apiClient.SsidsApi.UpdateNetworkWirelessSsidEapOverride(context.Background(), networkId, number).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SsidsApiService UpdateNetworkWirelessSsidFirewallL3FirewallRules", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var number string

		resp, httpRes, err := apiClient.SsidsApi.UpdateNetworkWirelessSsidFirewallL3FirewallRules(context.Background(), networkId, number).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SsidsApiService UpdateNetworkWirelessSsidFirewallL7FirewallRules", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var number string

		resp, httpRes, err := apiClient.SsidsApi.UpdateNetworkWirelessSsidFirewallL7FirewallRules(context.Background(), networkId, number).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SsidsApiService UpdateNetworkWirelessSsidHotspot20", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var number string

		resp, httpRes, err := apiClient.SsidsApi.UpdateNetworkWirelessSsidHotspot20(context.Background(), networkId, number).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SsidsApiService UpdateNetworkWirelessSsidIdentityPsk", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var number string
		var identityPskId string

		resp, httpRes, err := apiClient.SsidsApi.UpdateNetworkWirelessSsidIdentityPsk(context.Background(), networkId, number, identityPskId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SsidsApiService UpdateNetworkWirelessSsidSchedules", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var number string

		resp, httpRes, err := apiClient.SsidsApi.UpdateNetworkWirelessSsidSchedules(context.Background(), networkId, number).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SsidsApiService UpdateNetworkWirelessSsidSplashSettings", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var number string

		resp, httpRes, err := apiClient.SsidsApi.UpdateNetworkWirelessSsidSplashSettings(context.Background(), networkId, number).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SsidsApiService UpdateNetworkWirelessSsidTrafficShapingRules", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var number string

		resp, httpRes, err := apiClient.SsidsApi.UpdateNetworkWirelessSsidTrafficShapingRules(context.Background(), networkId, number).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SsidsApiService UpdateNetworkWirelessSsidVpn", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var number string

		resp, httpRes, err := apiClient.SsidsApi.UpdateNetworkWirelessSsidVpn(context.Background(), networkId, number).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}