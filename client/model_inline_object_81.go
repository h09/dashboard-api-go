/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 22 February, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: v1.31.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject81 struct for InlineObject81
type InlineObject81 struct {
	// All completed or in-progress stages in the network with their new start times. All pending stages will be canceled
	Stages []NetworksNetworkIdFirmwareUpgradesStagedEventsStages `json:"stages"`
	// The reason for rolling back the staged upgrade
	Reasons []NetworksNetworkIdFirmwareUpgradesRollbacksReasons `json:"reasons,omitempty"`
}

// NewInlineObject81 instantiates a new InlineObject81 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject81(stages []NetworksNetworkIdFirmwareUpgradesStagedEventsStages) *InlineObject81 {
	this := InlineObject81{}
	this.Stages = stages
	return &this
}

// NewInlineObject81WithDefaults instantiates a new InlineObject81 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject81WithDefaults() *InlineObject81 {
	this := InlineObject81{}
	return &this
}

// GetStages returns the Stages field value
func (o *InlineObject81) GetStages() []NetworksNetworkIdFirmwareUpgradesStagedEventsStages {
	if o == nil {
		var ret []NetworksNetworkIdFirmwareUpgradesStagedEventsStages
		return ret
	}

	return o.Stages
}

// GetStagesOk returns a tuple with the Stages field value
// and a boolean to check if the value has been set.
func (o *InlineObject81) GetStagesOk() ([]NetworksNetworkIdFirmwareUpgradesStagedEventsStages, bool) {
	if o == nil {
    return nil, false
	}
	return o.Stages, true
}

// SetStages sets field value
func (o *InlineObject81) SetStages(v []NetworksNetworkIdFirmwareUpgradesStagedEventsStages) {
	o.Stages = v
}

// GetReasons returns the Reasons field value if set, zero value otherwise.
func (o *InlineObject81) GetReasons() []NetworksNetworkIdFirmwareUpgradesRollbacksReasons {
	if o == nil || isNil(o.Reasons) {
		var ret []NetworksNetworkIdFirmwareUpgradesRollbacksReasons
		return ret
	}
	return o.Reasons
}

// GetReasonsOk returns a tuple with the Reasons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject81) GetReasonsOk() ([]NetworksNetworkIdFirmwareUpgradesRollbacksReasons, bool) {
	if o == nil || isNil(o.Reasons) {
    return nil, false
	}
	return o.Reasons, true
}

// HasReasons returns a boolean if a field has been set.
func (o *InlineObject81) HasReasons() bool {
	if o != nil && !isNil(o.Reasons) {
		return true
	}

	return false
}

// SetReasons gets a reference to the given []NetworksNetworkIdFirmwareUpgradesRollbacksReasons and assigns it to the Reasons field.
func (o *InlineObject81) SetReasons(v []NetworksNetworkIdFirmwareUpgradesRollbacksReasons) {
	o.Reasons = v
}

func (o InlineObject81) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["stages"] = o.Stages
	}
	if !isNil(o.Reasons) {
		toSerialize["reasons"] = o.Reasons
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject81 struct {
	value *InlineObject81
	isSet bool
}

func (v NullableInlineObject81) Get() *InlineObject81 {
	return v.value
}

func (v *NullableInlineObject81) Set(val *InlineObject81) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject81) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject81) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject81(val *InlineObject81) *NullableInlineObject81 {
	return &NullableInlineObject81{value: val, isSet: true}
}

func (v NullableInlineObject81) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject81) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

