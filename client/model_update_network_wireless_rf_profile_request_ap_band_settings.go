/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 05 July, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.35.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the UpdateNetworkWirelessRfProfileRequestApBandSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkWirelessRfProfileRequestApBandSettings{}

// UpdateNetworkWirelessRfProfileRequestApBandSettings Settings that will be enabled if selectionType is set to 'ap'.
type UpdateNetworkWirelessRfProfileRequestApBandSettings struct {
	// Choice between 'dual', '2.4ghz' or '5ghz'.
	BandOperationMode *string `json:"bandOperationMode,omitempty"`
	// Steers client to most open band. Can be either true or false.
	BandSteeringEnabled *bool `json:"bandSteeringEnabled,omitempty"`
}

// NewUpdateNetworkWirelessRfProfileRequestApBandSettings instantiates a new UpdateNetworkWirelessRfProfileRequestApBandSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkWirelessRfProfileRequestApBandSettings() *UpdateNetworkWirelessRfProfileRequestApBandSettings {
	this := UpdateNetworkWirelessRfProfileRequestApBandSettings{}
	return &this
}

// NewUpdateNetworkWirelessRfProfileRequestApBandSettingsWithDefaults instantiates a new UpdateNetworkWirelessRfProfileRequestApBandSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkWirelessRfProfileRequestApBandSettingsWithDefaults() *UpdateNetworkWirelessRfProfileRequestApBandSettings {
	this := UpdateNetworkWirelessRfProfileRequestApBandSettings{}
	return &this
}

// GetBandOperationMode returns the BandOperationMode field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessRfProfileRequestApBandSettings) GetBandOperationMode() string {
	if o == nil || IsNil(o.BandOperationMode) {
		var ret string
		return ret
	}
	return *o.BandOperationMode
}

// GetBandOperationModeOk returns a tuple with the BandOperationMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessRfProfileRequestApBandSettings) GetBandOperationModeOk() (*string, bool) {
	if o == nil || IsNil(o.BandOperationMode) {
		return nil, false
	}
	return o.BandOperationMode, true
}

// HasBandOperationMode returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessRfProfileRequestApBandSettings) HasBandOperationMode() bool {
	if o != nil && !IsNil(o.BandOperationMode) {
		return true
	}

	return false
}

// SetBandOperationMode gets a reference to the given string and assigns it to the BandOperationMode field.
func (o *UpdateNetworkWirelessRfProfileRequestApBandSettings) SetBandOperationMode(v string) {
	o.BandOperationMode = &v
}

// GetBandSteeringEnabled returns the BandSteeringEnabled field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessRfProfileRequestApBandSettings) GetBandSteeringEnabled() bool {
	if o == nil || IsNil(o.BandSteeringEnabled) {
		var ret bool
		return ret
	}
	return *o.BandSteeringEnabled
}

// GetBandSteeringEnabledOk returns a tuple with the BandSteeringEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessRfProfileRequestApBandSettings) GetBandSteeringEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.BandSteeringEnabled) {
		return nil, false
	}
	return o.BandSteeringEnabled, true
}

// HasBandSteeringEnabled returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessRfProfileRequestApBandSettings) HasBandSteeringEnabled() bool {
	if o != nil && !IsNil(o.BandSteeringEnabled) {
		return true
	}

	return false
}

// SetBandSteeringEnabled gets a reference to the given bool and assigns it to the BandSteeringEnabled field.
func (o *UpdateNetworkWirelessRfProfileRequestApBandSettings) SetBandSteeringEnabled(v bool) {
	o.BandSteeringEnabled = &v
}

func (o UpdateNetworkWirelessRfProfileRequestApBandSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkWirelessRfProfileRequestApBandSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BandOperationMode) {
		toSerialize["bandOperationMode"] = o.BandOperationMode
	}
	if !IsNil(o.BandSteeringEnabled) {
		toSerialize["bandSteeringEnabled"] = o.BandSteeringEnabled
	}
	return toSerialize, nil
}

type NullableUpdateNetworkWirelessRfProfileRequestApBandSettings struct {
	value *UpdateNetworkWirelessRfProfileRequestApBandSettings
	isSet bool
}

func (v NullableUpdateNetworkWirelessRfProfileRequestApBandSettings) Get() *UpdateNetworkWirelessRfProfileRequestApBandSettings {
	return v.value
}

func (v *NullableUpdateNetworkWirelessRfProfileRequestApBandSettings) Set(val *UpdateNetworkWirelessRfProfileRequestApBandSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkWirelessRfProfileRequestApBandSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkWirelessRfProfileRequestApBandSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkWirelessRfProfileRequestApBandSettings(val *UpdateNetworkWirelessRfProfileRequestApBandSettings) *NullableUpdateNetworkWirelessRfProfileRequestApBandSettings {
	return &NullableUpdateNetworkWirelessRfProfileRequestApBandSettings{value: val, isSet: true}
}

func (v NullableUpdateNetworkWirelessRfProfileRequestApBandSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkWirelessRfProfileRequestApBandSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

