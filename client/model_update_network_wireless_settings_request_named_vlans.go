/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 04 October, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.38.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the UpdateNetworkWirelessSettingsRequestNamedVlans type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkWirelessSettingsRequestNamedVlans{}

// UpdateNetworkWirelessSettingsRequestNamedVlans Named VLAN settings for wireless networks.
type UpdateNetworkWirelessSettingsRequestNamedVlans struct {
	PoolDhcpMonitoring *UpdateNetworkWirelessSettingsRequestNamedVlansPoolDhcpMonitoring `json:"poolDhcpMonitoring,omitempty"`
}

// NewUpdateNetworkWirelessSettingsRequestNamedVlans instantiates a new UpdateNetworkWirelessSettingsRequestNamedVlans object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkWirelessSettingsRequestNamedVlans() *UpdateNetworkWirelessSettingsRequestNamedVlans {
	this := UpdateNetworkWirelessSettingsRequestNamedVlans{}
	return &this
}

// NewUpdateNetworkWirelessSettingsRequestNamedVlansWithDefaults instantiates a new UpdateNetworkWirelessSettingsRequestNamedVlans object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkWirelessSettingsRequestNamedVlansWithDefaults() *UpdateNetworkWirelessSettingsRequestNamedVlans {
	this := UpdateNetworkWirelessSettingsRequestNamedVlans{}
	return &this
}

// GetPoolDhcpMonitoring returns the PoolDhcpMonitoring field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSettingsRequestNamedVlans) GetPoolDhcpMonitoring() UpdateNetworkWirelessSettingsRequestNamedVlansPoolDhcpMonitoring {
	if o == nil || IsNil(o.PoolDhcpMonitoring) {
		var ret UpdateNetworkWirelessSettingsRequestNamedVlansPoolDhcpMonitoring
		return ret
	}
	return *o.PoolDhcpMonitoring
}

// GetPoolDhcpMonitoringOk returns a tuple with the PoolDhcpMonitoring field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSettingsRequestNamedVlans) GetPoolDhcpMonitoringOk() (*UpdateNetworkWirelessSettingsRequestNamedVlansPoolDhcpMonitoring, bool) {
	if o == nil || IsNil(o.PoolDhcpMonitoring) {
		return nil, false
	}
	return o.PoolDhcpMonitoring, true
}

// HasPoolDhcpMonitoring returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSettingsRequestNamedVlans) HasPoolDhcpMonitoring() bool {
	if o != nil && !IsNil(o.PoolDhcpMonitoring) {
		return true
	}

	return false
}

// SetPoolDhcpMonitoring gets a reference to the given UpdateNetworkWirelessSettingsRequestNamedVlansPoolDhcpMonitoring and assigns it to the PoolDhcpMonitoring field.
func (o *UpdateNetworkWirelessSettingsRequestNamedVlans) SetPoolDhcpMonitoring(v UpdateNetworkWirelessSettingsRequestNamedVlansPoolDhcpMonitoring) {
	o.PoolDhcpMonitoring = &v
}

func (o UpdateNetworkWirelessSettingsRequestNamedVlans) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkWirelessSettingsRequestNamedVlans) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PoolDhcpMonitoring) {
		toSerialize["poolDhcpMonitoring"] = o.PoolDhcpMonitoring
	}
	return toSerialize, nil
}

type NullableUpdateNetworkWirelessSettingsRequestNamedVlans struct {
	value *UpdateNetworkWirelessSettingsRequestNamedVlans
	isSet bool
}

func (v NullableUpdateNetworkWirelessSettingsRequestNamedVlans) Get() *UpdateNetworkWirelessSettingsRequestNamedVlans {
	return v.value
}

func (v *NullableUpdateNetworkWirelessSettingsRequestNamedVlans) Set(val *UpdateNetworkWirelessSettingsRequestNamedVlans) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkWirelessSettingsRequestNamedVlans) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkWirelessSettingsRequestNamedVlans) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkWirelessSettingsRequestNamedVlans(val *UpdateNetworkWirelessSettingsRequestNamedVlans) *NullableUpdateNetworkWirelessSettingsRequestNamedVlans {
	return &NullableUpdateNetworkWirelessSettingsRequestNamedVlans{value: val, isSet: true}
}

func (v NullableUpdateNetworkWirelessSettingsRequestNamedVlans) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkWirelessSettingsRequestNamedVlans) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

