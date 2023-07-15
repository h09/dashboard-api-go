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

// checks if the UpdateNetworkFirmwareUpgradesStagedEventsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkFirmwareUpgradesStagedEventsRequest{}

// UpdateNetworkFirmwareUpgradesStagedEventsRequest struct for UpdateNetworkFirmwareUpgradesStagedEventsRequest
type UpdateNetworkFirmwareUpgradesStagedEventsRequest struct {
	// All firmware upgrade stages in the network with their start time.
	Stages []UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInner `json:"stages"`
}

// NewUpdateNetworkFirmwareUpgradesStagedEventsRequest instantiates a new UpdateNetworkFirmwareUpgradesStagedEventsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkFirmwareUpgradesStagedEventsRequest(stages []UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInner) *UpdateNetworkFirmwareUpgradesStagedEventsRequest {
	this := UpdateNetworkFirmwareUpgradesStagedEventsRequest{}
	this.Stages = stages
	return &this
}

// NewUpdateNetworkFirmwareUpgradesStagedEventsRequestWithDefaults instantiates a new UpdateNetworkFirmwareUpgradesStagedEventsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkFirmwareUpgradesStagedEventsRequestWithDefaults() *UpdateNetworkFirmwareUpgradesStagedEventsRequest {
	this := UpdateNetworkFirmwareUpgradesStagedEventsRequest{}
	return &this
}

// GetStages returns the Stages field value
func (o *UpdateNetworkFirmwareUpgradesStagedEventsRequest) GetStages() []UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInner {
	if o == nil {
		var ret []UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInner
		return ret
	}

	return o.Stages
}

// GetStagesOk returns a tuple with the Stages field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkFirmwareUpgradesStagedEventsRequest) GetStagesOk() ([]UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Stages, true
}

// SetStages sets field value
func (o *UpdateNetworkFirmwareUpgradesStagedEventsRequest) SetStages(v []UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInner) {
	o.Stages = v
}

func (o UpdateNetworkFirmwareUpgradesStagedEventsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkFirmwareUpgradesStagedEventsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["stages"] = o.Stages
	return toSerialize, nil
}

type NullableUpdateNetworkFirmwareUpgradesStagedEventsRequest struct {
	value *UpdateNetworkFirmwareUpgradesStagedEventsRequest
	isSet bool
}

func (v NullableUpdateNetworkFirmwareUpgradesStagedEventsRequest) Get() *UpdateNetworkFirmwareUpgradesStagedEventsRequest {
	return v.value
}

func (v *NullableUpdateNetworkFirmwareUpgradesStagedEventsRequest) Set(val *UpdateNetworkFirmwareUpgradesStagedEventsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkFirmwareUpgradesStagedEventsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkFirmwareUpgradesStagedEventsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkFirmwareUpgradesStagedEventsRequest(val *UpdateNetworkFirmwareUpgradesStagedEventsRequest) *NullableUpdateNetworkFirmwareUpgradesStagedEventsRequest {
	return &NullableUpdateNetworkFirmwareUpgradesStagedEventsRequest{value: val, isSet: true}
}

func (v NullableUpdateNetworkFirmwareUpgradesStagedEventsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkFirmwareUpgradesStagedEventsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

