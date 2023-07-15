/*
Meraki Dashboard API

Testing InsightApiService

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

func Test_client_InsightApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test InsightApiService CreateOrganizationInsightMonitoredMediaServer", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string

		resp, httpRes, err := apiClient.InsightApi.CreateOrganizationInsightMonitoredMediaServer(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InsightApiService DeleteOrganizationInsightMonitoredMediaServer", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string
		var monitoredMediaServerId string

		httpRes, err := apiClient.InsightApi.DeleteOrganizationInsightMonitoredMediaServer(context.Background(), organizationId, monitoredMediaServerId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InsightApiService GetNetworkInsightApplicationHealthByTime", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string
		var applicationId string

		resp, httpRes, err := apiClient.InsightApi.GetNetworkInsightApplicationHealthByTime(context.Background(), networkId, applicationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InsightApiService GetOrganizationInsightApplications", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string

		resp, httpRes, err := apiClient.InsightApi.GetOrganizationInsightApplications(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InsightApiService GetOrganizationInsightMonitoredMediaServer", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string
		var monitoredMediaServerId string

		resp, httpRes, err := apiClient.InsightApi.GetOrganizationInsightMonitoredMediaServer(context.Background(), organizationId, monitoredMediaServerId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InsightApiService GetOrganizationInsightMonitoredMediaServers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string

		resp, httpRes, err := apiClient.InsightApi.GetOrganizationInsightMonitoredMediaServers(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InsightApiService UpdateOrganizationInsightMonitoredMediaServer", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string
		var monitoredMediaServerId string

		resp, httpRes, err := apiClient.InsightApi.UpdateOrganizationInsightMonitoredMediaServer(context.Background(), organizationId, monitoredMediaServerId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}