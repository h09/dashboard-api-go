/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 02 November, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.27.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


// RollbacksApiService RollbacksApi service
type RollbacksApiService service

type RollbacksApiCreateNetworkFirmwareUpgradesRollbackRequest struct {
	ctx context.Context
	ApiService *RollbacksApiService
	networkId string
	createNetworkFirmwareUpgradesRollback *CreateNetworkFirmwareUpgradesRollbackRequest
}

func (r RollbacksApiCreateNetworkFirmwareUpgradesRollbackRequest) CreateNetworkFirmwareUpgradesRollback(createNetworkFirmwareUpgradesRollback CreateNetworkFirmwareUpgradesRollbackRequest) RollbacksApiCreateNetworkFirmwareUpgradesRollbackRequest {
	r.createNetworkFirmwareUpgradesRollback = &createNetworkFirmwareUpgradesRollback
	return r
}

func (r RollbacksApiCreateNetworkFirmwareUpgradesRollbackRequest) Execute() (*CreateNetworkFirmwareUpgradesRollback200Response, *http.Response, error) {
	return r.ApiService.CreateNetworkFirmwareUpgradesRollbackExecute(r)
}

/*
CreateNetworkFirmwareUpgradesRollback Rollback a Firmware Upgrade For A Network

Rollback a Firmware Upgrade For A Network

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId
 @return RollbacksApiCreateNetworkFirmwareUpgradesRollbackRequest
*/
func (a *RollbacksApiService) CreateNetworkFirmwareUpgradesRollback(ctx context.Context, networkId string) RollbacksApiCreateNetworkFirmwareUpgradesRollbackRequest {
	return RollbacksApiCreateNetworkFirmwareUpgradesRollbackRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return CreateNetworkFirmwareUpgradesRollback200Response
func (a *RollbacksApiService) CreateNetworkFirmwareUpgradesRollbackExecute(r RollbacksApiCreateNetworkFirmwareUpgradesRollbackRequest) (*CreateNetworkFirmwareUpgradesRollback200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CreateNetworkFirmwareUpgradesRollback200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RollbacksApiService.CreateNetworkFirmwareUpgradesRollback")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/firmwareUpgrades/rollbacks"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterToString(r.networkId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createNetworkFirmwareUpgradesRollback == nil {
		return localVarReturnValue, nil, reportError("createNetworkFirmwareUpgradesRollback is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.createNetworkFirmwareUpgradesRollback
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["meraki_api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Cisco-Meraki-API-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}