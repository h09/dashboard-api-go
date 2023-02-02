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

// InlineObject75 struct for InlineObject75
type InlineObject75 struct {
	// The size of the vMX you claim. It can be one of: small, medium, large, 100
	Size string `json:"size"`
}

// NewInlineObject75 instantiates a new InlineObject75 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject75(size string) *InlineObject75 {
	this := InlineObject75{}
	this.Size = size
	return &this
}

// NewInlineObject75WithDefaults instantiates a new InlineObject75 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject75WithDefaults() *InlineObject75 {
	this := InlineObject75{}
	return &this
}

// GetSize returns the Size field value
func (o *InlineObject75) GetSize() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *InlineObject75) GetSizeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *InlineObject75) SetSize(v string) {
	o.Size = v
}

func (o InlineObject75) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["size"] = o.Size
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject75 struct {
	value *InlineObject75
	isSet bool
}

func (v NullableInlineObject75) Get() *InlineObject75 {
	return v.value
}

func (v *NullableInlineObject75) Set(val *InlineObject75) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject75) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject75) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject75(val *InlineObject75) *NullableInlineObject75 {
	return &NullableInlineObject75{value: val, isSet: true}
}

func (v NullableInlineObject75) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject75) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


