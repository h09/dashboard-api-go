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

// NetworksNetworkIdFirmwareUpgradesProductsWirelessNextUpgrade The pending firmware upgrade if it exists
type NetworksNetworkIdFirmwareUpgradesProductsWirelessNextUpgrade struct {
	// The time of the last successful upgrade
	Time *string `json:"time,omitempty"`
	ToVersion *NetworksNetworkIdFirmwareUpgradesProductsWirelessNextUpgradeToVersion `json:"toVersion,omitempty"`
}

// NewNetworksNetworkIdFirmwareUpgradesProductsWirelessNextUpgrade instantiates a new NetworksNetworkIdFirmwareUpgradesProductsWirelessNextUpgrade object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdFirmwareUpgradesProductsWirelessNextUpgrade() *NetworksNetworkIdFirmwareUpgradesProductsWirelessNextUpgrade {
	this := NetworksNetworkIdFirmwareUpgradesProductsWirelessNextUpgrade{}
	return &this
}

// NewNetworksNetworkIdFirmwareUpgradesProductsWirelessNextUpgradeWithDefaults instantiates a new NetworksNetworkIdFirmwareUpgradesProductsWirelessNextUpgrade object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdFirmwareUpgradesProductsWirelessNextUpgradeWithDefaults() *NetworksNetworkIdFirmwareUpgradesProductsWirelessNextUpgrade {
	this := NetworksNetworkIdFirmwareUpgradesProductsWirelessNextUpgrade{}
	return &this
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *NetworksNetworkIdFirmwareUpgradesProductsWirelessNextUpgrade) GetTime() string {
	if o == nil || isNil(o.Time) {
		var ret string
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdFirmwareUpgradesProductsWirelessNextUpgrade) GetTimeOk() (*string, bool) {
	if o == nil || isNil(o.Time) {
    return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *NetworksNetworkIdFirmwareUpgradesProductsWirelessNextUpgrade) HasTime() bool {
	if o != nil && !isNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given string and assigns it to the Time field.
func (o *NetworksNetworkIdFirmwareUpgradesProductsWirelessNextUpgrade) SetTime(v string) {
	o.Time = &v
}

// GetToVersion returns the ToVersion field value if set, zero value otherwise.
func (o *NetworksNetworkIdFirmwareUpgradesProductsWirelessNextUpgrade) GetToVersion() NetworksNetworkIdFirmwareUpgradesProductsWirelessNextUpgradeToVersion {
	if o == nil || isNil(o.ToVersion) {
		var ret NetworksNetworkIdFirmwareUpgradesProductsWirelessNextUpgradeToVersion
		return ret
	}
	return *o.ToVersion
}

// GetToVersionOk returns a tuple with the ToVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdFirmwareUpgradesProductsWirelessNextUpgrade) GetToVersionOk() (*NetworksNetworkIdFirmwareUpgradesProductsWirelessNextUpgradeToVersion, bool) {
	if o == nil || isNil(o.ToVersion) {
    return nil, false
	}
	return o.ToVersion, true
}

// HasToVersion returns a boolean if a field has been set.
func (o *NetworksNetworkIdFirmwareUpgradesProductsWirelessNextUpgrade) HasToVersion() bool {
	if o != nil && !isNil(o.ToVersion) {
		return true
	}

	return false
}

// SetToVersion gets a reference to the given NetworksNetworkIdFirmwareUpgradesProductsWirelessNextUpgradeToVersion and assigns it to the ToVersion field.
func (o *NetworksNetworkIdFirmwareUpgradesProductsWirelessNextUpgrade) SetToVersion(v NetworksNetworkIdFirmwareUpgradesProductsWirelessNextUpgradeToVersion) {
	o.ToVersion = &v
}

func (o NetworksNetworkIdFirmwareUpgradesProductsWirelessNextUpgrade) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	if !isNil(o.ToVersion) {
		toSerialize["toVersion"] = o.ToVersion
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdFirmwareUpgradesProductsWirelessNextUpgrade struct {
	value *NetworksNetworkIdFirmwareUpgradesProductsWirelessNextUpgrade
	isSet bool
}

func (v NullableNetworksNetworkIdFirmwareUpgradesProductsWirelessNextUpgrade) Get() *NetworksNetworkIdFirmwareUpgradesProductsWirelessNextUpgrade {
	return v.value
}

func (v *NullableNetworksNetworkIdFirmwareUpgradesProductsWirelessNextUpgrade) Set(val *NetworksNetworkIdFirmwareUpgradesProductsWirelessNextUpgrade) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdFirmwareUpgradesProductsWirelessNextUpgrade) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdFirmwareUpgradesProductsWirelessNextUpgrade) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdFirmwareUpgradesProductsWirelessNextUpgrade(val *NetworksNetworkIdFirmwareUpgradesProductsWirelessNextUpgrade) *NullableNetworksNetworkIdFirmwareUpgradesProductsWirelessNextUpgrade {
	return &NullableNetworksNetworkIdFirmwareUpgradesProductsWirelessNextUpgrade{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdFirmwareUpgradesProductsWirelessNextUpgrade) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdFirmwareUpgradesProductsWirelessNextUpgrade) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


