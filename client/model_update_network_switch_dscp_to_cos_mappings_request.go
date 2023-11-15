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

// checks if the UpdateNetworkSwitchDscpToCosMappingsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkSwitchDscpToCosMappingsRequest{}

// UpdateNetworkSwitchDscpToCosMappingsRequest struct for UpdateNetworkSwitchDscpToCosMappingsRequest
type UpdateNetworkSwitchDscpToCosMappingsRequest struct {
	// An array of DSCP to CoS mappings. An empty array will reset the mappings to default.
	Mappings []UpdateNetworkSwitchDscpToCosMappingsRequestMappingsInner `json:"mappings"`
}

// NewUpdateNetworkSwitchDscpToCosMappingsRequest instantiates a new UpdateNetworkSwitchDscpToCosMappingsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkSwitchDscpToCosMappingsRequest(mappings []UpdateNetworkSwitchDscpToCosMappingsRequestMappingsInner) *UpdateNetworkSwitchDscpToCosMappingsRequest {
	this := UpdateNetworkSwitchDscpToCosMappingsRequest{}
	this.Mappings = mappings
	return &this
}

// NewUpdateNetworkSwitchDscpToCosMappingsRequestWithDefaults instantiates a new UpdateNetworkSwitchDscpToCosMappingsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkSwitchDscpToCosMappingsRequestWithDefaults() *UpdateNetworkSwitchDscpToCosMappingsRequest {
	this := UpdateNetworkSwitchDscpToCosMappingsRequest{}
	return &this
}

// GetMappings returns the Mappings field value
func (o *UpdateNetworkSwitchDscpToCosMappingsRequest) GetMappings() []UpdateNetworkSwitchDscpToCosMappingsRequestMappingsInner {
	if o == nil {
		var ret []UpdateNetworkSwitchDscpToCosMappingsRequestMappingsInner
		return ret
	}

	return o.Mappings
}

// GetMappingsOk returns a tuple with the Mappings field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSwitchDscpToCosMappingsRequest) GetMappingsOk() ([]UpdateNetworkSwitchDscpToCosMappingsRequestMappingsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Mappings, true
}

// SetMappings sets field value
func (o *UpdateNetworkSwitchDscpToCosMappingsRequest) SetMappings(v []UpdateNetworkSwitchDscpToCosMappingsRequestMappingsInner) {
	o.Mappings = v
}

func (o UpdateNetworkSwitchDscpToCosMappingsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkSwitchDscpToCosMappingsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mappings"] = o.Mappings
	return toSerialize, nil
}

type NullableUpdateNetworkSwitchDscpToCosMappingsRequest struct {
	value *UpdateNetworkSwitchDscpToCosMappingsRequest
	isSet bool
}

func (v NullableUpdateNetworkSwitchDscpToCosMappingsRequest) Get() *UpdateNetworkSwitchDscpToCosMappingsRequest {
	return v.value
}

func (v *NullableUpdateNetworkSwitchDscpToCosMappingsRequest) Set(val *UpdateNetworkSwitchDscpToCosMappingsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkSwitchDscpToCosMappingsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkSwitchDscpToCosMappingsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkSwitchDscpToCosMappingsRequest(val *UpdateNetworkSwitchDscpToCosMappingsRequest) *NullableUpdateNetworkSwitchDscpToCosMappingsRequest {
	return &NullableUpdateNetworkSwitchDscpToCosMappingsRequest{value: val, isSet: true}
}

func (v NullableUpdateNetworkSwitchDscpToCosMappingsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkSwitchDscpToCosMappingsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


