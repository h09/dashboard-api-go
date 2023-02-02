/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 February, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: v1.30.0
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


// EthernetApiService EthernetApi service
type EthernetApiService service

type EthernetApiGetOrganizationWirelessDevicesEthernetStatusesRequest struct {
	ctx context.Context
	ApiService *EthernetApiService
	organizationId string
	perPage *int32
	startingAfter *string
	endingBefore *string
	networkIds *[]string
}

// The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100.
func (r EthernetApiGetOrganizationWirelessDevicesEthernetStatusesRequest) PerPage(perPage int32) EthernetApiGetOrganizationWirelessDevicesEthernetStatusesRequest {
	r.perPage = &perPage
	return r
}

// A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r EthernetApiGetOrganizationWirelessDevicesEthernetStatusesRequest) StartingAfter(startingAfter string) EthernetApiGetOrganizationWirelessDevicesEthernetStatusesRequest {
	r.startingAfter = &startingAfter
	return r
}

// A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r EthernetApiGetOrganizationWirelessDevicesEthernetStatusesRequest) EndingBefore(endingBefore string) EthernetApiGetOrganizationWirelessDevicesEthernetStatusesRequest {
	r.endingBefore = &endingBefore
	return r
}

// A list of Meraki network IDs to filter results to contain only specified networks. E.g.: networkIds[]&#x3D;N_12345678&amp;networkIds[]&#x3D;L_3456
func (r EthernetApiGetOrganizationWirelessDevicesEthernetStatusesRequest) NetworkIds(networkIds []string) EthernetApiGetOrganizationWirelessDevicesEthernetStatusesRequest {
	r.networkIds = &networkIds
	return r
}

func (r EthernetApiGetOrganizationWirelessDevicesEthernetStatusesRequest) Execute() ([]InlineResponse200129, *http.Response, error) {
	return r.ApiService.GetOrganizationWirelessDevicesEthernetStatusesExecute(r)
}

/*
GetOrganizationWirelessDevicesEthernetStatuses Endpoint to see power status for wireless devices

Endpoint to see power status for wireless devices

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId
 @return EthernetApiGetOrganizationWirelessDevicesEthernetStatusesRequest
*/
func (a *EthernetApiService) GetOrganizationWirelessDevicesEthernetStatuses(ctx context.Context, organizationId string) EthernetApiGetOrganizationWirelessDevicesEthernetStatusesRequest {
	return EthernetApiGetOrganizationWirelessDevicesEthernetStatusesRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []InlineResponse200129
func (a *EthernetApiService) GetOrganizationWirelessDevicesEthernetStatusesExecute(r EthernetApiGetOrganizationWirelessDevicesEthernetStatusesRequest) ([]InlineResponse200129, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []InlineResponse200129
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EthernetApiService.GetOrganizationWirelessDevicesEthernetStatuses")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/wireless/devices/ethernet/statuses"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterToString(r.organizationId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.perPage != nil {
		localVarQueryParams.Add("perPage", parameterToString(*r.perPage, ""))
	}
	if r.startingAfter != nil {
		localVarQueryParams.Add("startingAfter", parameterToString(*r.startingAfter, ""))
	}
	if r.endingBefore != nil {
		localVarQueryParams.Add("endingBefore", parameterToString(*r.endingBefore, ""))
	}
	if r.networkIds != nil {
		localVarQueryParams.Add("networkIds", parameterToString(*r.networkIds, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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
