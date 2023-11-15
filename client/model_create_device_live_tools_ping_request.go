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

// checks if the CreateDeviceLiveToolsPingRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateDeviceLiveToolsPingRequest{}

// CreateDeviceLiveToolsPingRequest struct for CreateDeviceLiveToolsPingRequest
type CreateDeviceLiveToolsPingRequest struct {
	// FQDN, IPv4 or IPv6 address
	Target string `json:"target"`
	// Count parameter to pass to ping. [1..5], default 5
	Count *int32 `json:"count,omitempty"`
}

// NewCreateDeviceLiveToolsPingRequest instantiates a new CreateDeviceLiveToolsPingRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateDeviceLiveToolsPingRequest(target string) *CreateDeviceLiveToolsPingRequest {
	this := CreateDeviceLiveToolsPingRequest{}
	this.Target = target
	return &this
}

// NewCreateDeviceLiveToolsPingRequestWithDefaults instantiates a new CreateDeviceLiveToolsPingRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateDeviceLiveToolsPingRequestWithDefaults() *CreateDeviceLiveToolsPingRequest {
	this := CreateDeviceLiveToolsPingRequest{}
	return &this
}

// GetTarget returns the Target field value
func (o *CreateDeviceLiveToolsPingRequest) GetTarget() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *CreateDeviceLiveToolsPingRequest) GetTargetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value
func (o *CreateDeviceLiveToolsPingRequest) SetTarget(v string) {
	o.Target = v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *CreateDeviceLiveToolsPingRequest) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDeviceLiveToolsPingRequest) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *CreateDeviceLiveToolsPingRequest) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *CreateDeviceLiveToolsPingRequest) SetCount(v int32) {
	o.Count = &v
}

func (o CreateDeviceLiveToolsPingRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateDeviceLiveToolsPingRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["target"] = o.Target
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	return toSerialize, nil
}

type NullableCreateDeviceLiveToolsPingRequest struct {
	value *CreateDeviceLiveToolsPingRequest
	isSet bool
}

func (v NullableCreateDeviceLiveToolsPingRequest) Get() *CreateDeviceLiveToolsPingRequest {
	return v.value
}

func (v *NullableCreateDeviceLiveToolsPingRequest) Set(val *CreateDeviceLiveToolsPingRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateDeviceLiveToolsPingRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateDeviceLiveToolsPingRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateDeviceLiveToolsPingRequest(val *CreateDeviceLiveToolsPingRequest) *NullableCreateDeviceLiveToolsPingRequest {
	return &NullableCreateDeviceLiveToolsPingRequest{value: val, isSet: true}
}

func (v NullableCreateDeviceLiveToolsPingRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateDeviceLiveToolsPingRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


