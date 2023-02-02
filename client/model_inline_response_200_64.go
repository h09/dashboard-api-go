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

// InlineResponse20064 struct for InlineResponse20064
type InlineResponse20064 struct {
	// Broadcast threshold.
	BroadcastThreshold *int32 `json:"broadcastThreshold,omitempty"`
	// Multicast threshold.
	MulticastThreshold *int32 `json:"multicastThreshold,omitempty"`
	// Unknown Unicast threshold.
	UnknownUnicastThreshold *int32 `json:"unknownUnicastThreshold,omitempty"`
}

// NewInlineResponse20064 instantiates a new InlineResponse20064 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20064() *InlineResponse20064 {
	this := InlineResponse20064{}
	return &this
}

// NewInlineResponse20064WithDefaults instantiates a new InlineResponse20064 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20064WithDefaults() *InlineResponse20064 {
	this := InlineResponse20064{}
	return &this
}

// GetBroadcastThreshold returns the BroadcastThreshold field value if set, zero value otherwise.
func (o *InlineResponse20064) GetBroadcastThreshold() int32 {
	if o == nil || isNil(o.BroadcastThreshold) {
		var ret int32
		return ret
	}
	return *o.BroadcastThreshold
}

// GetBroadcastThresholdOk returns a tuple with the BroadcastThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20064) GetBroadcastThresholdOk() (*int32, bool) {
	if o == nil || isNil(o.BroadcastThreshold) {
    return nil, false
	}
	return o.BroadcastThreshold, true
}

// HasBroadcastThreshold returns a boolean if a field has been set.
func (o *InlineResponse20064) HasBroadcastThreshold() bool {
	if o != nil && !isNil(o.BroadcastThreshold) {
		return true
	}

	return false
}

// SetBroadcastThreshold gets a reference to the given int32 and assigns it to the BroadcastThreshold field.
func (o *InlineResponse20064) SetBroadcastThreshold(v int32) {
	o.BroadcastThreshold = &v
}

// GetMulticastThreshold returns the MulticastThreshold field value if set, zero value otherwise.
func (o *InlineResponse20064) GetMulticastThreshold() int32 {
	if o == nil || isNil(o.MulticastThreshold) {
		var ret int32
		return ret
	}
	return *o.MulticastThreshold
}

// GetMulticastThresholdOk returns a tuple with the MulticastThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20064) GetMulticastThresholdOk() (*int32, bool) {
	if o == nil || isNil(o.MulticastThreshold) {
    return nil, false
	}
	return o.MulticastThreshold, true
}

// HasMulticastThreshold returns a boolean if a field has been set.
func (o *InlineResponse20064) HasMulticastThreshold() bool {
	if o != nil && !isNil(o.MulticastThreshold) {
		return true
	}

	return false
}

// SetMulticastThreshold gets a reference to the given int32 and assigns it to the MulticastThreshold field.
func (o *InlineResponse20064) SetMulticastThreshold(v int32) {
	o.MulticastThreshold = &v
}

// GetUnknownUnicastThreshold returns the UnknownUnicastThreshold field value if set, zero value otherwise.
func (o *InlineResponse20064) GetUnknownUnicastThreshold() int32 {
	if o == nil || isNil(o.UnknownUnicastThreshold) {
		var ret int32
		return ret
	}
	return *o.UnknownUnicastThreshold
}

// GetUnknownUnicastThresholdOk returns a tuple with the UnknownUnicastThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20064) GetUnknownUnicastThresholdOk() (*int32, bool) {
	if o == nil || isNil(o.UnknownUnicastThreshold) {
    return nil, false
	}
	return o.UnknownUnicastThreshold, true
}

// HasUnknownUnicastThreshold returns a boolean if a field has been set.
func (o *InlineResponse20064) HasUnknownUnicastThreshold() bool {
	if o != nil && !isNil(o.UnknownUnicastThreshold) {
		return true
	}

	return false
}

// SetUnknownUnicastThreshold gets a reference to the given int32 and assigns it to the UnknownUnicastThreshold field.
func (o *InlineResponse20064) SetUnknownUnicastThreshold(v int32) {
	o.UnknownUnicastThreshold = &v
}

func (o InlineResponse20064) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.BroadcastThreshold) {
		toSerialize["broadcastThreshold"] = o.BroadcastThreshold
	}
	if !isNil(o.MulticastThreshold) {
		toSerialize["multicastThreshold"] = o.MulticastThreshold
	}
	if !isNil(o.UnknownUnicastThreshold) {
		toSerialize["unknownUnicastThreshold"] = o.UnknownUnicastThreshold
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20064 struct {
	value *InlineResponse20064
	isSet bool
}

func (v NullableInlineResponse20064) Get() *InlineResponse20064 {
	return v.value
}

func (v *NullableInlineResponse20064) Set(val *InlineResponse20064) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20064) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20064) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20064(val *InlineResponse20064) *NullableInlineResponse20064 {
	return &NullableInlineResponse20064{value: val, isSet: true}
}

func (v NullableInlineResponse20064) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20064) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


