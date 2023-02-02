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

// InlineResponse2001Interfaces Interface settings.
type InlineResponse2001Interfaces struct {
	Wan1 *InlineResponse2001InterfacesWan1 `json:"wan1,omitempty"`
	Wan2 *InlineResponse2001InterfacesWan2 `json:"wan2,omitempty"`
}

// NewInlineResponse2001Interfaces instantiates a new InlineResponse2001Interfaces object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2001Interfaces() *InlineResponse2001Interfaces {
	this := InlineResponse2001Interfaces{}
	return &this
}

// NewInlineResponse2001InterfacesWithDefaults instantiates a new InlineResponse2001Interfaces object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2001InterfacesWithDefaults() *InlineResponse2001Interfaces {
	this := InlineResponse2001Interfaces{}
	return &this
}

// GetWan1 returns the Wan1 field value if set, zero value otherwise.
func (o *InlineResponse2001Interfaces) GetWan1() InlineResponse2001InterfacesWan1 {
	if o == nil || isNil(o.Wan1) {
		var ret InlineResponse2001InterfacesWan1
		return ret
	}
	return *o.Wan1
}

// GetWan1Ok returns a tuple with the Wan1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2001Interfaces) GetWan1Ok() (*InlineResponse2001InterfacesWan1, bool) {
	if o == nil || isNil(o.Wan1) {
    return nil, false
	}
	return o.Wan1, true
}

// HasWan1 returns a boolean if a field has been set.
func (o *InlineResponse2001Interfaces) HasWan1() bool {
	if o != nil && !isNil(o.Wan1) {
		return true
	}

	return false
}

// SetWan1 gets a reference to the given InlineResponse2001InterfacesWan1 and assigns it to the Wan1 field.
func (o *InlineResponse2001Interfaces) SetWan1(v InlineResponse2001InterfacesWan1) {
	o.Wan1 = &v
}

// GetWan2 returns the Wan2 field value if set, zero value otherwise.
func (o *InlineResponse2001Interfaces) GetWan2() InlineResponse2001InterfacesWan2 {
	if o == nil || isNil(o.Wan2) {
		var ret InlineResponse2001InterfacesWan2
		return ret
	}
	return *o.Wan2
}

// GetWan2Ok returns a tuple with the Wan2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2001Interfaces) GetWan2Ok() (*InlineResponse2001InterfacesWan2, bool) {
	if o == nil || isNil(o.Wan2) {
    return nil, false
	}
	return o.Wan2, true
}

// HasWan2 returns a boolean if a field has been set.
func (o *InlineResponse2001Interfaces) HasWan2() bool {
	if o != nil && !isNil(o.Wan2) {
		return true
	}

	return false
}

// SetWan2 gets a reference to the given InlineResponse2001InterfacesWan2 and assigns it to the Wan2 field.
func (o *InlineResponse2001Interfaces) SetWan2(v InlineResponse2001InterfacesWan2) {
	o.Wan2 = &v
}

func (o InlineResponse2001Interfaces) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Wan1) {
		toSerialize["wan1"] = o.Wan1
	}
	if !isNil(o.Wan2) {
		toSerialize["wan2"] = o.Wan2
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2001Interfaces struct {
	value *InlineResponse2001Interfaces
	isSet bool
}

func (v NullableInlineResponse2001Interfaces) Get() *InlineResponse2001Interfaces {
	return v.value
}

func (v *NullableInlineResponse2001Interfaces) Set(val *InlineResponse2001Interfaces) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2001Interfaces) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2001Interfaces) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2001Interfaces(val *InlineResponse2001Interfaces) *NullableInlineResponse2001Interfaces {
	return &NullableInlineResponse2001Interfaces{value: val, isSet: true}
}

func (v NullableInlineResponse2001Interfaces) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2001Interfaces) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


