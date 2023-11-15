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

// checks if the UpdateNetworkApplianceVpnBgpRequestNeighborsInnerTtlSecurity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkApplianceVpnBgpRequestNeighborsInnerTtlSecurity{}

// UpdateNetworkApplianceVpnBgpRequestNeighborsInnerTtlSecurity Settings for BGP TTL security to protect BGP peering sessions from forged IP attacks.
type UpdateNetworkApplianceVpnBgpRequestNeighborsInnerTtlSecurity struct {
	// Boolean value to enable or disable BGP TTL security.
	Enabled *bool `json:"enabled,omitempty"`
}

// NewUpdateNetworkApplianceVpnBgpRequestNeighborsInnerTtlSecurity instantiates a new UpdateNetworkApplianceVpnBgpRequestNeighborsInnerTtlSecurity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkApplianceVpnBgpRequestNeighborsInnerTtlSecurity() *UpdateNetworkApplianceVpnBgpRequestNeighborsInnerTtlSecurity {
	this := UpdateNetworkApplianceVpnBgpRequestNeighborsInnerTtlSecurity{}
	return &this
}

// NewUpdateNetworkApplianceVpnBgpRequestNeighborsInnerTtlSecurityWithDefaults instantiates a new UpdateNetworkApplianceVpnBgpRequestNeighborsInnerTtlSecurity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkApplianceVpnBgpRequestNeighborsInnerTtlSecurityWithDefaults() *UpdateNetworkApplianceVpnBgpRequestNeighborsInnerTtlSecurity {
	this := UpdateNetworkApplianceVpnBgpRequestNeighborsInnerTtlSecurity{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceVpnBgpRequestNeighborsInnerTtlSecurity) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceVpnBgpRequestNeighborsInnerTtlSecurity) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceVpnBgpRequestNeighborsInnerTtlSecurity) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *UpdateNetworkApplianceVpnBgpRequestNeighborsInnerTtlSecurity) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o UpdateNetworkApplianceVpnBgpRequestNeighborsInnerTtlSecurity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkApplianceVpnBgpRequestNeighborsInnerTtlSecurity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return toSerialize, nil
}

type NullableUpdateNetworkApplianceVpnBgpRequestNeighborsInnerTtlSecurity struct {
	value *UpdateNetworkApplianceVpnBgpRequestNeighborsInnerTtlSecurity
	isSet bool
}

func (v NullableUpdateNetworkApplianceVpnBgpRequestNeighborsInnerTtlSecurity) Get() *UpdateNetworkApplianceVpnBgpRequestNeighborsInnerTtlSecurity {
	return v.value
}

func (v *NullableUpdateNetworkApplianceVpnBgpRequestNeighborsInnerTtlSecurity) Set(val *UpdateNetworkApplianceVpnBgpRequestNeighborsInnerTtlSecurity) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkApplianceVpnBgpRequestNeighborsInnerTtlSecurity) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkApplianceVpnBgpRequestNeighborsInnerTtlSecurity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkApplianceVpnBgpRequestNeighborsInnerTtlSecurity(val *UpdateNetworkApplianceVpnBgpRequestNeighborsInnerTtlSecurity) *NullableUpdateNetworkApplianceVpnBgpRequestNeighborsInnerTtlSecurity {
	return &NullableUpdateNetworkApplianceVpnBgpRequestNeighborsInnerTtlSecurity{value: val, isSet: true}
}

func (v NullableUpdateNetworkApplianceVpnBgpRequestNeighborsInnerTtlSecurity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkApplianceVpnBgpRequestNeighborsInnerTtlSecurity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


