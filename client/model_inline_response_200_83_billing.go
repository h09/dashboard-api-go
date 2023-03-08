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

// InlineResponse20083Billing Details associated with billing splash
type InlineResponse20083Billing struct {
	FreeAccess *InlineResponse20083BillingFreeAccess `json:"freeAccess,omitempty"`
	// Whether or not billing uses the fast login prepaid access option.
	PrepaidAccessFastLoginEnabled *bool `json:"prepaidAccessFastLoginEnabled,omitempty"`
	// The email address that reeceives replies from clients
	ReplyToEmailAddress *string `json:"replyToEmailAddress,omitempty"`
}

// NewInlineResponse20083Billing instantiates a new InlineResponse20083Billing object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20083Billing() *InlineResponse20083Billing {
	this := InlineResponse20083Billing{}
	return &this
}

// NewInlineResponse20083BillingWithDefaults instantiates a new InlineResponse20083Billing object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20083BillingWithDefaults() *InlineResponse20083Billing {
	this := InlineResponse20083Billing{}
	return &this
}

// GetFreeAccess returns the FreeAccess field value if set, zero value otherwise.
func (o *InlineResponse20083Billing) GetFreeAccess() InlineResponse20083BillingFreeAccess {
	if o == nil || isNil(o.FreeAccess) {
		var ret InlineResponse20083BillingFreeAccess
		return ret
	}
	return *o.FreeAccess
}

// GetFreeAccessOk returns a tuple with the FreeAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20083Billing) GetFreeAccessOk() (*InlineResponse20083BillingFreeAccess, bool) {
	if o == nil || isNil(o.FreeAccess) {
    return nil, false
	}
	return o.FreeAccess, true
}

// HasFreeAccess returns a boolean if a field has been set.
func (o *InlineResponse20083Billing) HasFreeAccess() bool {
	if o != nil && !isNil(o.FreeAccess) {
		return true
	}

	return false
}

// SetFreeAccess gets a reference to the given InlineResponse20083BillingFreeAccess and assigns it to the FreeAccess field.
func (o *InlineResponse20083Billing) SetFreeAccess(v InlineResponse20083BillingFreeAccess) {
	o.FreeAccess = &v
}

// GetPrepaidAccessFastLoginEnabled returns the PrepaidAccessFastLoginEnabled field value if set, zero value otherwise.
func (o *InlineResponse20083Billing) GetPrepaidAccessFastLoginEnabled() bool {
	if o == nil || isNil(o.PrepaidAccessFastLoginEnabled) {
		var ret bool
		return ret
	}
	return *o.PrepaidAccessFastLoginEnabled
}

// GetPrepaidAccessFastLoginEnabledOk returns a tuple with the PrepaidAccessFastLoginEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20083Billing) GetPrepaidAccessFastLoginEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.PrepaidAccessFastLoginEnabled) {
    return nil, false
	}
	return o.PrepaidAccessFastLoginEnabled, true
}

// HasPrepaidAccessFastLoginEnabled returns a boolean if a field has been set.
func (o *InlineResponse20083Billing) HasPrepaidAccessFastLoginEnabled() bool {
	if o != nil && !isNil(o.PrepaidAccessFastLoginEnabled) {
		return true
	}

	return false
}

// SetPrepaidAccessFastLoginEnabled gets a reference to the given bool and assigns it to the PrepaidAccessFastLoginEnabled field.
func (o *InlineResponse20083Billing) SetPrepaidAccessFastLoginEnabled(v bool) {
	o.PrepaidAccessFastLoginEnabled = &v
}

// GetReplyToEmailAddress returns the ReplyToEmailAddress field value if set, zero value otherwise.
func (o *InlineResponse20083Billing) GetReplyToEmailAddress() string {
	if o == nil || isNil(o.ReplyToEmailAddress) {
		var ret string
		return ret
	}
	return *o.ReplyToEmailAddress
}

// GetReplyToEmailAddressOk returns a tuple with the ReplyToEmailAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20083Billing) GetReplyToEmailAddressOk() (*string, bool) {
	if o == nil || isNil(o.ReplyToEmailAddress) {
    return nil, false
	}
	return o.ReplyToEmailAddress, true
}

// HasReplyToEmailAddress returns a boolean if a field has been set.
func (o *InlineResponse20083Billing) HasReplyToEmailAddress() bool {
	if o != nil && !isNil(o.ReplyToEmailAddress) {
		return true
	}

	return false
}

// SetReplyToEmailAddress gets a reference to the given string and assigns it to the ReplyToEmailAddress field.
func (o *InlineResponse20083Billing) SetReplyToEmailAddress(v string) {
	o.ReplyToEmailAddress = &v
}

func (o InlineResponse20083Billing) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.FreeAccess) {
		toSerialize["freeAccess"] = o.FreeAccess
	}
	if !isNil(o.PrepaidAccessFastLoginEnabled) {
		toSerialize["prepaidAccessFastLoginEnabled"] = o.PrepaidAccessFastLoginEnabled
	}
	if !isNil(o.ReplyToEmailAddress) {
		toSerialize["replyToEmailAddress"] = o.ReplyToEmailAddress
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20083Billing struct {
	value *InlineResponse20083Billing
	isSet bool
}

func (v NullableInlineResponse20083Billing) Get() *InlineResponse20083Billing {
	return v.value
}

func (v *NullableInlineResponse20083Billing) Set(val *InlineResponse20083Billing) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20083Billing) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20083Billing) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20083Billing(val *InlineResponse20083Billing) *NullableInlineResponse20083Billing {
	return &NullableInlineResponse20083Billing{value: val, isSet: true}
}

func (v NullableInlineResponse20083Billing) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20083Billing) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

