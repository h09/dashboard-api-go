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

// InlineObject157 struct for InlineObject157
type InlineObject157 struct {
	// An ordered array of the firewall rules for this SSID (not including the local LAN access rule or the default rule)
	Rules []NetworksNetworkIdWirelessSsidsNumberFirewallL3FirewallRulesRules `json:"rules,omitempty"`
	// Allow wireless client access to local LAN (boolean value - true allows access and false denies access) (optional)
	AllowLanAccess *bool `json:"allowLanAccess,omitempty"`
}

// NewInlineObject157 instantiates a new InlineObject157 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject157() *InlineObject157 {
	this := InlineObject157{}
	return &this
}

// NewInlineObject157WithDefaults instantiates a new InlineObject157 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject157WithDefaults() *InlineObject157 {
	this := InlineObject157{}
	return &this
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *InlineObject157) GetRules() []NetworksNetworkIdWirelessSsidsNumberFirewallL3FirewallRulesRules {
	if o == nil || isNil(o.Rules) {
		var ret []NetworksNetworkIdWirelessSsidsNumberFirewallL3FirewallRulesRules
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject157) GetRulesOk() ([]NetworksNetworkIdWirelessSsidsNumberFirewallL3FirewallRulesRules, bool) {
	if o == nil || isNil(o.Rules) {
    return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *InlineObject157) HasRules() bool {
	if o != nil && !isNil(o.Rules) {
		return true
	}

	return false
}

// SetRules gets a reference to the given []NetworksNetworkIdWirelessSsidsNumberFirewallL3FirewallRulesRules and assigns it to the Rules field.
func (o *InlineObject157) SetRules(v []NetworksNetworkIdWirelessSsidsNumberFirewallL3FirewallRulesRules) {
	o.Rules = v
}

// GetAllowLanAccess returns the AllowLanAccess field value if set, zero value otherwise.
func (o *InlineObject157) GetAllowLanAccess() bool {
	if o == nil || isNil(o.AllowLanAccess) {
		var ret bool
		return ret
	}
	return *o.AllowLanAccess
}

// GetAllowLanAccessOk returns a tuple with the AllowLanAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject157) GetAllowLanAccessOk() (*bool, bool) {
	if o == nil || isNil(o.AllowLanAccess) {
    return nil, false
	}
	return o.AllowLanAccess, true
}

// HasAllowLanAccess returns a boolean if a field has been set.
func (o *InlineObject157) HasAllowLanAccess() bool {
	if o != nil && !isNil(o.AllowLanAccess) {
		return true
	}

	return false
}

// SetAllowLanAccess gets a reference to the given bool and assigns it to the AllowLanAccess field.
func (o *InlineObject157) SetAllowLanAccess(v bool) {
	o.AllowLanAccess = &v
}

func (o InlineObject157) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Rules) {
		toSerialize["rules"] = o.Rules
	}
	if !isNil(o.AllowLanAccess) {
		toSerialize["allowLanAccess"] = o.AllowLanAccess
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject157 struct {
	value *InlineObject157
	isSet bool
}

func (v NullableInlineObject157) Get() *InlineObject157 {
	return v.value
}

func (v *NullableInlineObject157) Set(val *InlineObject157) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject157) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject157) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject157(val *InlineObject157) *NullableInlineObject157 {
	return &NullableInlineObject157{value: val, isSet: true}
}

func (v NullableInlineObject157) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject157) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

