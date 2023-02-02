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

// InlineResponse20016TrafficFilters struct for InlineResponse20016TrafficFilters
type InlineResponse20016TrafficFilters struct {
	// Traffic filter type. Must be \"custom\"
	Type string `json:"type"`
	Value InlineResponse20016Value `json:"value"`
}

// NewInlineResponse20016TrafficFilters instantiates a new InlineResponse20016TrafficFilters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20016TrafficFilters(type_ string, value InlineResponse20016Value) *InlineResponse20016TrafficFilters {
	this := InlineResponse20016TrafficFilters{}
	this.Type = type_
	this.Value = value
	return &this
}

// NewInlineResponse20016TrafficFiltersWithDefaults instantiates a new InlineResponse20016TrafficFilters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20016TrafficFiltersWithDefaults() *InlineResponse20016TrafficFilters {
	this := InlineResponse20016TrafficFilters{}
	return &this
}

// GetType returns the Type field value
func (o *InlineResponse20016TrafficFilters) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20016TrafficFilters) GetTypeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InlineResponse20016TrafficFilters) SetType(v string) {
	o.Type = v
}

// GetValue returns the Value field value
func (o *InlineResponse20016TrafficFilters) GetValue() InlineResponse20016Value {
	if o == nil {
		var ret InlineResponse20016Value
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20016TrafficFilters) GetValueOk() (*InlineResponse20016Value, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *InlineResponse20016TrafficFilters) SetValue(v InlineResponse20016Value) {
	o.Value = v
}

func (o InlineResponse20016TrafficFilters) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20016TrafficFilters struct {
	value *InlineResponse20016TrafficFilters
	isSet bool
}

func (v NullableInlineResponse20016TrafficFilters) Get() *InlineResponse20016TrafficFilters {
	return v.value
}

func (v *NullableInlineResponse20016TrafficFilters) Set(val *InlineResponse20016TrafficFilters) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20016TrafficFilters) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20016TrafficFilters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20016TrafficFilters(val *InlineResponse20016TrafficFilters) *NullableInlineResponse20016TrafficFilters {
	return &NullableInlineResponse20016TrafficFilters{value: val, isSet: true}
}

func (v NullableInlineResponse20016TrafficFilters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20016TrafficFilters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


