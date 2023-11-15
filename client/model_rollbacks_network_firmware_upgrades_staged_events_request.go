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

// checks if the RollbacksNetworkFirmwareUpgradesStagedEventsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RollbacksNetworkFirmwareUpgradesStagedEventsRequest{}

// RollbacksNetworkFirmwareUpgradesStagedEventsRequest struct for RollbacksNetworkFirmwareUpgradesStagedEventsRequest
type RollbacksNetworkFirmwareUpgradesStagedEventsRequest struct {
	// All completed or in-progress stages in the network with their new start times. All pending stages will be canceled
	Stages []UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInner `json:"stages"`
	// The reason for rolling back the staged upgrade
	Reasons []CreateNetworkFirmwareUpgradesRollbackRequestReasonsInner `json:"reasons,omitempty"`
}

// NewRollbacksNetworkFirmwareUpgradesStagedEventsRequest instantiates a new RollbacksNetworkFirmwareUpgradesStagedEventsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRollbacksNetworkFirmwareUpgradesStagedEventsRequest(stages []UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInner) *RollbacksNetworkFirmwareUpgradesStagedEventsRequest {
	this := RollbacksNetworkFirmwareUpgradesStagedEventsRequest{}
	this.Stages = stages
	return &this
}

// NewRollbacksNetworkFirmwareUpgradesStagedEventsRequestWithDefaults instantiates a new RollbacksNetworkFirmwareUpgradesStagedEventsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRollbacksNetworkFirmwareUpgradesStagedEventsRequestWithDefaults() *RollbacksNetworkFirmwareUpgradesStagedEventsRequest {
	this := RollbacksNetworkFirmwareUpgradesStagedEventsRequest{}
	return &this
}

// GetStages returns the Stages field value
func (o *RollbacksNetworkFirmwareUpgradesStagedEventsRequest) GetStages() []UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInner {
	if o == nil {
		var ret []UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInner
		return ret
	}

	return o.Stages
}

// GetStagesOk returns a tuple with the Stages field value
// and a boolean to check if the value has been set.
func (o *RollbacksNetworkFirmwareUpgradesStagedEventsRequest) GetStagesOk() ([]UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Stages, true
}

// SetStages sets field value
func (o *RollbacksNetworkFirmwareUpgradesStagedEventsRequest) SetStages(v []UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInner) {
	o.Stages = v
}

// GetReasons returns the Reasons field value if set, zero value otherwise.
func (o *RollbacksNetworkFirmwareUpgradesStagedEventsRequest) GetReasons() []CreateNetworkFirmwareUpgradesRollbackRequestReasonsInner {
	if o == nil || IsNil(o.Reasons) {
		var ret []CreateNetworkFirmwareUpgradesRollbackRequestReasonsInner
		return ret
	}
	return o.Reasons
}

// GetReasonsOk returns a tuple with the Reasons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RollbacksNetworkFirmwareUpgradesStagedEventsRequest) GetReasonsOk() ([]CreateNetworkFirmwareUpgradesRollbackRequestReasonsInner, bool) {
	if o == nil || IsNil(o.Reasons) {
		return nil, false
	}
	return o.Reasons, true
}

// HasReasons returns a boolean if a field has been set.
func (o *RollbacksNetworkFirmwareUpgradesStagedEventsRequest) HasReasons() bool {
	if o != nil && !IsNil(o.Reasons) {
		return true
	}

	return false
}

// SetReasons gets a reference to the given []CreateNetworkFirmwareUpgradesRollbackRequestReasonsInner and assigns it to the Reasons field.
func (o *RollbacksNetworkFirmwareUpgradesStagedEventsRequest) SetReasons(v []CreateNetworkFirmwareUpgradesRollbackRequestReasonsInner) {
	o.Reasons = v
}

func (o RollbacksNetworkFirmwareUpgradesStagedEventsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RollbacksNetworkFirmwareUpgradesStagedEventsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["stages"] = o.Stages
	if !IsNil(o.Reasons) {
		toSerialize["reasons"] = o.Reasons
	}
	return toSerialize, nil
}

type NullableRollbacksNetworkFirmwareUpgradesStagedEventsRequest struct {
	value *RollbacksNetworkFirmwareUpgradesStagedEventsRequest
	isSet bool
}

func (v NullableRollbacksNetworkFirmwareUpgradesStagedEventsRequest) Get() *RollbacksNetworkFirmwareUpgradesStagedEventsRequest {
	return v.value
}

func (v *NullableRollbacksNetworkFirmwareUpgradesStagedEventsRequest) Set(val *RollbacksNetworkFirmwareUpgradesStagedEventsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRollbacksNetworkFirmwareUpgradesStagedEventsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRollbacksNetworkFirmwareUpgradesStagedEventsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRollbacksNetworkFirmwareUpgradesStagedEventsRequest(val *RollbacksNetworkFirmwareUpgradesStagedEventsRequest) *NullableRollbacksNetworkFirmwareUpgradesStagedEventsRequest {
	return &NullableRollbacksNetworkFirmwareUpgradesStagedEventsRequest{value: val, isSet: true}
}

func (v NullableRollbacksNetworkFirmwareUpgradesStagedEventsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRollbacksNetworkFirmwareUpgradesStagedEventsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


