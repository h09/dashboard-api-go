/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 01 November, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.39.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the UninstallNetworkSmDeviceAppsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UninstallNetworkSmDeviceAppsRequest{}

// UninstallNetworkSmDeviceAppsRequest struct for UninstallNetworkSmDeviceAppsRequest
type UninstallNetworkSmDeviceAppsRequest struct {
	// ids of applications to be uninstalled
	AppIds []string `json:"appIds"`
}

// NewUninstallNetworkSmDeviceAppsRequest instantiates a new UninstallNetworkSmDeviceAppsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUninstallNetworkSmDeviceAppsRequest(appIds []string) *UninstallNetworkSmDeviceAppsRequest {
	this := UninstallNetworkSmDeviceAppsRequest{}
	this.AppIds = appIds
	return &this
}

// NewUninstallNetworkSmDeviceAppsRequestWithDefaults instantiates a new UninstallNetworkSmDeviceAppsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUninstallNetworkSmDeviceAppsRequestWithDefaults() *UninstallNetworkSmDeviceAppsRequest {
	this := UninstallNetworkSmDeviceAppsRequest{}
	return &this
}

// GetAppIds returns the AppIds field value
func (o *UninstallNetworkSmDeviceAppsRequest) GetAppIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AppIds
}

// GetAppIdsOk returns a tuple with the AppIds field value
// and a boolean to check if the value has been set.
func (o *UninstallNetworkSmDeviceAppsRequest) GetAppIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AppIds, true
}

// SetAppIds sets field value
func (o *UninstallNetworkSmDeviceAppsRequest) SetAppIds(v []string) {
	o.AppIds = v
}

func (o UninstallNetworkSmDeviceAppsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UninstallNetworkSmDeviceAppsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["appIds"] = o.AppIds
	return toSerialize, nil
}

type NullableUninstallNetworkSmDeviceAppsRequest struct {
	value *UninstallNetworkSmDeviceAppsRequest
	isSet bool
}

func (v NullableUninstallNetworkSmDeviceAppsRequest) Get() *UninstallNetworkSmDeviceAppsRequest {
	return v.value
}

func (v *NullableUninstallNetworkSmDeviceAppsRequest) Set(val *UninstallNetworkSmDeviceAppsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUninstallNetworkSmDeviceAppsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUninstallNetworkSmDeviceAppsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUninstallNetworkSmDeviceAppsRequest(val *UninstallNetworkSmDeviceAppsRequest) *NullableUninstallNetworkSmDeviceAppsRequest {
	return &NullableUninstallNetworkSmDeviceAppsRequest{value: val, isSet: true}
}

func (v NullableUninstallNetworkSmDeviceAppsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUninstallNetworkSmDeviceAppsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


