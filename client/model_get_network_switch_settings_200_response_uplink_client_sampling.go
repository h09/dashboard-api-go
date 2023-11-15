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

// checks if the GetNetworkSwitchSettings200ResponseUplinkClientSampling type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkSwitchSettings200ResponseUplinkClientSampling{}

// GetNetworkSwitchSettings200ResponseUplinkClientSampling Uplink client sampling
type GetNetworkSwitchSettings200ResponseUplinkClientSampling struct {
	// Enable client sampling on uplink
	Enabled *bool `json:"enabled,omitempty"`
}

// NewGetNetworkSwitchSettings200ResponseUplinkClientSampling instantiates a new GetNetworkSwitchSettings200ResponseUplinkClientSampling object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkSwitchSettings200ResponseUplinkClientSampling() *GetNetworkSwitchSettings200ResponseUplinkClientSampling {
	this := GetNetworkSwitchSettings200ResponseUplinkClientSampling{}
	return &this
}

// NewGetNetworkSwitchSettings200ResponseUplinkClientSamplingWithDefaults instantiates a new GetNetworkSwitchSettings200ResponseUplinkClientSampling object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkSwitchSettings200ResponseUplinkClientSamplingWithDefaults() *GetNetworkSwitchSettings200ResponseUplinkClientSampling {
	this := GetNetworkSwitchSettings200ResponseUplinkClientSampling{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *GetNetworkSwitchSettings200ResponseUplinkClientSampling) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchSettings200ResponseUplinkClientSampling) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *GetNetworkSwitchSettings200ResponseUplinkClientSampling) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *GetNetworkSwitchSettings200ResponseUplinkClientSampling) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o GetNetworkSwitchSettings200ResponseUplinkClientSampling) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkSwitchSettings200ResponseUplinkClientSampling) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return toSerialize, nil
}

type NullableGetNetworkSwitchSettings200ResponseUplinkClientSampling struct {
	value *GetNetworkSwitchSettings200ResponseUplinkClientSampling
	isSet bool
}

func (v NullableGetNetworkSwitchSettings200ResponseUplinkClientSampling) Get() *GetNetworkSwitchSettings200ResponseUplinkClientSampling {
	return v.value
}

func (v *NullableGetNetworkSwitchSettings200ResponseUplinkClientSampling) Set(val *GetNetworkSwitchSettings200ResponseUplinkClientSampling) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkSwitchSettings200ResponseUplinkClientSampling) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkSwitchSettings200ResponseUplinkClientSampling) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkSwitchSettings200ResponseUplinkClientSampling(val *GetNetworkSwitchSettings200ResponseUplinkClientSampling) *NullableGetNetworkSwitchSettings200ResponseUplinkClientSampling {
	return &NullableGetNetworkSwitchSettings200ResponseUplinkClientSampling{value: val, isSet: true}
}

func (v NullableGetNetworkSwitchSettings200ResponseUplinkClientSampling) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkSwitchSettings200ResponseUplinkClientSampling) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


