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

// InlineResponse200112 struct for InlineResponse200112
type InlineResponse200112 struct {
	// Toggle depicting if SAML SSO settings are enabled
	Enabled *bool `json:"enabled,omitempty"`
}

// NewInlineResponse200112 instantiates a new InlineResponse200112 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200112() *InlineResponse200112 {
	this := InlineResponse200112{}
	return &this
}

// NewInlineResponse200112WithDefaults instantiates a new InlineResponse200112 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200112WithDefaults() *InlineResponse200112 {
	this := InlineResponse200112{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineResponse200112) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200112) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineResponse200112) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineResponse200112) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o InlineResponse200112) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200112 struct {
	value *InlineResponse200112
	isSet bool
}

func (v NullableInlineResponse200112) Get() *InlineResponse200112 {
	return v.value
}

func (v *NullableInlineResponse200112) Set(val *InlineResponse200112) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200112) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200112) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200112(val *InlineResponse200112) *NullableInlineResponse200112 {
	return &NullableInlineResponse200112{value: val, isSet: true}
}

func (v NullableInlineResponse200112) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200112) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

