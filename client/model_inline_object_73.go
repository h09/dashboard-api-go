/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 February, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: v1.30.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject73 struct for InlineObject73
type InlineObject73 struct {
	Ssids NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids `json:"ssids"`
}

// NewInlineObject73 instantiates a new InlineObject73 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject73(ssids NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids) *InlineObject73 {
	this := InlineObject73{}
	this.Ssids = ssids
	return &this
}

// NewInlineObject73WithDefaults instantiates a new InlineObject73 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject73WithDefaults() *InlineObject73 {
	this := InlineObject73{}
	return &this
}

// GetSsids returns the Ssids field value
func (o *InlineObject73) GetSsids() NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids {
	if o == nil {
		var ret NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids
		return ret
	}

	return o.Ssids
}

// GetSsidsOk returns a tuple with the Ssids field value
// and a boolean to check if the value has been set.
func (o *InlineObject73) GetSsidsOk() (*NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Ssids, true
}

// SetSsids sets field value
func (o *InlineObject73) SetSsids(v NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids) {
	o.Ssids = v
}

func (o InlineObject73) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ssids"] = o.Ssids
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject73 struct {
	value *InlineObject73
	isSet bool
}

func (v NullableInlineObject73) Get() *InlineObject73 {
	return v.value
}

func (v *NullableInlineObject73) Set(val *InlineObject73) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject73) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject73) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject73(val *InlineObject73) *NullableInlineObject73 {
	return &NullableInlineObject73{value: val, isSet: true}
}

func (v NullableInlineObject73) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject73) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


