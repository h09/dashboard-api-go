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

// InlineResponse20020Usage Usage, sent and received
type InlineResponse20020Usage struct {
	// Usage sent by the client
	Sent *float32 `json:"sent,omitempty"`
	// Usage received by the client
	Recv *float32 `json:"recv,omitempty"`
}

// NewInlineResponse20020Usage instantiates a new InlineResponse20020Usage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20020Usage() *InlineResponse20020Usage {
	this := InlineResponse20020Usage{}
	return &this
}

// NewInlineResponse20020UsageWithDefaults instantiates a new InlineResponse20020Usage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20020UsageWithDefaults() *InlineResponse20020Usage {
	this := InlineResponse20020Usage{}
	return &this
}

// GetSent returns the Sent field value if set, zero value otherwise.
func (o *InlineResponse20020Usage) GetSent() float32 {
	if o == nil || isNil(o.Sent) {
		var ret float32
		return ret
	}
	return *o.Sent
}

// GetSentOk returns a tuple with the Sent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20020Usage) GetSentOk() (*float32, bool) {
	if o == nil || isNil(o.Sent) {
    return nil, false
	}
	return o.Sent, true
}

// HasSent returns a boolean if a field has been set.
func (o *InlineResponse20020Usage) HasSent() bool {
	if o != nil && !isNil(o.Sent) {
		return true
	}

	return false
}

// SetSent gets a reference to the given float32 and assigns it to the Sent field.
func (o *InlineResponse20020Usage) SetSent(v float32) {
	o.Sent = &v
}

// GetRecv returns the Recv field value if set, zero value otherwise.
func (o *InlineResponse20020Usage) GetRecv() float32 {
	if o == nil || isNil(o.Recv) {
		var ret float32
		return ret
	}
	return *o.Recv
}

// GetRecvOk returns a tuple with the Recv field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20020Usage) GetRecvOk() (*float32, bool) {
	if o == nil || isNil(o.Recv) {
    return nil, false
	}
	return o.Recv, true
}

// HasRecv returns a boolean if a field has been set.
func (o *InlineResponse20020Usage) HasRecv() bool {
	if o != nil && !isNil(o.Recv) {
		return true
	}

	return false
}

// SetRecv gets a reference to the given float32 and assigns it to the Recv field.
func (o *InlineResponse20020Usage) SetRecv(v float32) {
	o.Recv = &v
}

func (o InlineResponse20020Usage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Sent) {
		toSerialize["sent"] = o.Sent
	}
	if !isNil(o.Recv) {
		toSerialize["recv"] = o.Recv
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20020Usage struct {
	value *InlineResponse20020Usage
	isSet bool
}

func (v NullableInlineResponse20020Usage) Get() *InlineResponse20020Usage {
	return v.value
}

func (v *NullableInlineResponse20020Usage) Set(val *InlineResponse20020Usage) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20020Usage) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20020Usage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20020Usage(val *InlineResponse20020Usage) *NullableInlineResponse20020Usage {
	return &NullableInlineResponse20020Usage{value: val, isSet: true}
}

func (v NullableInlineResponse20020Usage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20020Usage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


