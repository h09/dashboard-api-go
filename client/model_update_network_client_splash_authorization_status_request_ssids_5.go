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

// checks if the UpdateNetworkClientSplashAuthorizationStatusRequestSsids5 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkClientSplashAuthorizationStatusRequestSsids5{}

// UpdateNetworkClientSplashAuthorizationStatusRequestSsids5 Splash authorization for SSID 5
type UpdateNetworkClientSplashAuthorizationStatusRequestSsids5 struct {
	// New authorization status for the SSID (true, false).
	IsAuthorized *bool `json:"isAuthorized,omitempty"`
}

// NewUpdateNetworkClientSplashAuthorizationStatusRequestSsids5 instantiates a new UpdateNetworkClientSplashAuthorizationStatusRequestSsids5 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkClientSplashAuthorizationStatusRequestSsids5() *UpdateNetworkClientSplashAuthorizationStatusRequestSsids5 {
	this := UpdateNetworkClientSplashAuthorizationStatusRequestSsids5{}
	return &this
}

// NewUpdateNetworkClientSplashAuthorizationStatusRequestSsids5WithDefaults instantiates a new UpdateNetworkClientSplashAuthorizationStatusRequestSsids5 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkClientSplashAuthorizationStatusRequestSsids5WithDefaults() *UpdateNetworkClientSplashAuthorizationStatusRequestSsids5 {
	this := UpdateNetworkClientSplashAuthorizationStatusRequestSsids5{}
	return &this
}

// GetIsAuthorized returns the IsAuthorized field value if set, zero value otherwise.
func (o *UpdateNetworkClientSplashAuthorizationStatusRequestSsids5) GetIsAuthorized() bool {
	if o == nil || IsNil(o.IsAuthorized) {
		var ret bool
		return ret
	}
	return *o.IsAuthorized
}

// GetIsAuthorizedOk returns a tuple with the IsAuthorized field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkClientSplashAuthorizationStatusRequestSsids5) GetIsAuthorizedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAuthorized) {
		return nil, false
	}
	return o.IsAuthorized, true
}

// HasIsAuthorized returns a boolean if a field has been set.
func (o *UpdateNetworkClientSplashAuthorizationStatusRequestSsids5) HasIsAuthorized() bool {
	if o != nil && !IsNil(o.IsAuthorized) {
		return true
	}

	return false
}

// SetIsAuthorized gets a reference to the given bool and assigns it to the IsAuthorized field.
func (o *UpdateNetworkClientSplashAuthorizationStatusRequestSsids5) SetIsAuthorized(v bool) {
	o.IsAuthorized = &v
}

func (o UpdateNetworkClientSplashAuthorizationStatusRequestSsids5) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkClientSplashAuthorizationStatusRequestSsids5) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IsAuthorized) {
		toSerialize["isAuthorized"] = o.IsAuthorized
	}
	return toSerialize, nil
}

type NullableUpdateNetworkClientSplashAuthorizationStatusRequestSsids5 struct {
	value *UpdateNetworkClientSplashAuthorizationStatusRequestSsids5
	isSet bool
}

func (v NullableUpdateNetworkClientSplashAuthorizationStatusRequestSsids5) Get() *UpdateNetworkClientSplashAuthorizationStatusRequestSsids5 {
	return v.value
}

func (v *NullableUpdateNetworkClientSplashAuthorizationStatusRequestSsids5) Set(val *UpdateNetworkClientSplashAuthorizationStatusRequestSsids5) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkClientSplashAuthorizationStatusRequestSsids5) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkClientSplashAuthorizationStatusRequestSsids5) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkClientSplashAuthorizationStatusRequestSsids5(val *UpdateNetworkClientSplashAuthorizationStatusRequestSsids5) *NullableUpdateNetworkClientSplashAuthorizationStatusRequestSsids5 {
	return &NullableUpdateNetworkClientSplashAuthorizationStatusRequestSsids5{value: val, isSet: true}
}

func (v NullableUpdateNetworkClientSplashAuthorizationStatusRequestSsids5) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkClientSplashAuthorizationStatusRequestSsids5) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


