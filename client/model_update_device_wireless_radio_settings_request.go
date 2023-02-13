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

// UpdateDeviceWirelessRadioSettingsRequest struct for UpdateDeviceWirelessRadioSettingsRequest
type UpdateDeviceWirelessRadioSettingsRequest struct {
	// The ID of an RF profile to assign to the device. If the value of this parameter is null, the appropriate basic RF profile (indoor or outdoor) will be assigned to the device. Assigning an RF profile will clear ALL manually configured overrides on the device (channel width, channel, power).
	RfProfileId *string `json:"rfProfileId,omitempty"`
	TwoFourGhzSettings *UpdateDeviceWirelessRadioSettingsRequestTwoFourGhzSettings `json:"twoFourGhzSettings,omitempty"`
	FiveGhzSettings *UpdateDeviceWirelessRadioSettingsRequestFiveGhzSettings `json:"fiveGhzSettings,omitempty"`
}

// NewUpdateDeviceWirelessRadioSettingsRequest instantiates a new UpdateDeviceWirelessRadioSettingsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateDeviceWirelessRadioSettingsRequest() *UpdateDeviceWirelessRadioSettingsRequest {
	this := UpdateDeviceWirelessRadioSettingsRequest{}
	return &this
}

// NewUpdateDeviceWirelessRadioSettingsRequestWithDefaults instantiates a new UpdateDeviceWirelessRadioSettingsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateDeviceWirelessRadioSettingsRequestWithDefaults() *UpdateDeviceWirelessRadioSettingsRequest {
	this := UpdateDeviceWirelessRadioSettingsRequest{}
	return &this
}

// GetRfProfileId returns the RfProfileId field value if set, zero value otherwise.
func (o *UpdateDeviceWirelessRadioSettingsRequest) GetRfProfileId() string {
	if o == nil || isNil(o.RfProfileId) {
		var ret string
		return ret
	}
	return *o.RfProfileId
}

// GetRfProfileIdOk returns a tuple with the RfProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceWirelessRadioSettingsRequest) GetRfProfileIdOk() (*string, bool) {
	if o == nil || isNil(o.RfProfileId) {
    return nil, false
	}
	return o.RfProfileId, true
}

// HasRfProfileId returns a boolean if a field has been set.
func (o *UpdateDeviceWirelessRadioSettingsRequest) HasRfProfileId() bool {
	if o != nil && !isNil(o.RfProfileId) {
		return true
	}

	return false
}

// SetRfProfileId gets a reference to the given string and assigns it to the RfProfileId field.
func (o *UpdateDeviceWirelessRadioSettingsRequest) SetRfProfileId(v string) {
	o.RfProfileId = &v
}

// GetTwoFourGhzSettings returns the TwoFourGhzSettings field value if set, zero value otherwise.
func (o *UpdateDeviceWirelessRadioSettingsRequest) GetTwoFourGhzSettings() UpdateDeviceWirelessRadioSettingsRequestTwoFourGhzSettings {
	if o == nil || isNil(o.TwoFourGhzSettings) {
		var ret UpdateDeviceWirelessRadioSettingsRequestTwoFourGhzSettings
		return ret
	}
	return *o.TwoFourGhzSettings
}

// GetTwoFourGhzSettingsOk returns a tuple with the TwoFourGhzSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceWirelessRadioSettingsRequest) GetTwoFourGhzSettingsOk() (*UpdateDeviceWirelessRadioSettingsRequestTwoFourGhzSettings, bool) {
	if o == nil || isNil(o.TwoFourGhzSettings) {
    return nil, false
	}
	return o.TwoFourGhzSettings, true
}

// HasTwoFourGhzSettings returns a boolean if a field has been set.
func (o *UpdateDeviceWirelessRadioSettingsRequest) HasTwoFourGhzSettings() bool {
	if o != nil && !isNil(o.TwoFourGhzSettings) {
		return true
	}

	return false
}

// SetTwoFourGhzSettings gets a reference to the given UpdateDeviceWirelessRadioSettingsRequestTwoFourGhzSettings and assigns it to the TwoFourGhzSettings field.
func (o *UpdateDeviceWirelessRadioSettingsRequest) SetTwoFourGhzSettings(v UpdateDeviceWirelessRadioSettingsRequestTwoFourGhzSettings) {
	o.TwoFourGhzSettings = &v
}

// GetFiveGhzSettings returns the FiveGhzSettings field value if set, zero value otherwise.
func (o *UpdateDeviceWirelessRadioSettingsRequest) GetFiveGhzSettings() UpdateDeviceWirelessRadioSettingsRequestFiveGhzSettings {
	if o == nil || isNil(o.FiveGhzSettings) {
		var ret UpdateDeviceWirelessRadioSettingsRequestFiveGhzSettings
		return ret
	}
	return *o.FiveGhzSettings
}

// GetFiveGhzSettingsOk returns a tuple with the FiveGhzSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceWirelessRadioSettingsRequest) GetFiveGhzSettingsOk() (*UpdateDeviceWirelessRadioSettingsRequestFiveGhzSettings, bool) {
	if o == nil || isNil(o.FiveGhzSettings) {
    return nil, false
	}
	return o.FiveGhzSettings, true
}

// HasFiveGhzSettings returns a boolean if a field has been set.
func (o *UpdateDeviceWirelessRadioSettingsRequest) HasFiveGhzSettings() bool {
	if o != nil && !isNil(o.FiveGhzSettings) {
		return true
	}

	return false
}

// SetFiveGhzSettings gets a reference to the given UpdateDeviceWirelessRadioSettingsRequestFiveGhzSettings and assigns it to the FiveGhzSettings field.
func (o *UpdateDeviceWirelessRadioSettingsRequest) SetFiveGhzSettings(v UpdateDeviceWirelessRadioSettingsRequestFiveGhzSettings) {
	o.FiveGhzSettings = &v
}

func (o UpdateDeviceWirelessRadioSettingsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.RfProfileId) {
		toSerialize["rfProfileId"] = o.RfProfileId
	}
	if !isNil(o.TwoFourGhzSettings) {
		toSerialize["twoFourGhzSettings"] = o.TwoFourGhzSettings
	}
	if !isNil(o.FiveGhzSettings) {
		toSerialize["fiveGhzSettings"] = o.FiveGhzSettings
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateDeviceWirelessRadioSettingsRequest struct {
	value *UpdateDeviceWirelessRadioSettingsRequest
	isSet bool
}

func (v NullableUpdateDeviceWirelessRadioSettingsRequest) Get() *UpdateDeviceWirelessRadioSettingsRequest {
	return v.value
}

func (v *NullableUpdateDeviceWirelessRadioSettingsRequest) Set(val *UpdateDeviceWirelessRadioSettingsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateDeviceWirelessRadioSettingsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateDeviceWirelessRadioSettingsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateDeviceWirelessRadioSettingsRequest(val *UpdateDeviceWirelessRadioSettingsRequest) *NullableUpdateDeviceWirelessRadioSettingsRequest {
	return &NullableUpdateDeviceWirelessRadioSettingsRequest{value: val, isSet: true}
}

func (v NullableUpdateDeviceWirelessRadioSettingsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateDeviceWirelessRadioSettingsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

