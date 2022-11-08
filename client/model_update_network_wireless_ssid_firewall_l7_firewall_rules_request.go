/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 02 November, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.27.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// UpdateNetworkWirelessSsidFirewallL7FirewallRulesRequest struct for UpdateNetworkWirelessSsidFirewallL7FirewallRulesRequest
type UpdateNetworkWirelessSsidFirewallL7FirewallRulesRequest struct {
	// An array of L7 firewall rules for this SSID. Rules will get applied in the same order user has specified in request. Empty array will clear the L7 firewall rule configuration.
	Rules []UpdateNetworkWirelessSsidFirewallL7FirewallRulesRequestRulesInner `json:"rules,omitempty"`
}

// NewUpdateNetworkWirelessSsidFirewallL7FirewallRulesRequest instantiates a new UpdateNetworkWirelessSsidFirewallL7FirewallRulesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkWirelessSsidFirewallL7FirewallRulesRequest() *UpdateNetworkWirelessSsidFirewallL7FirewallRulesRequest {
	this := UpdateNetworkWirelessSsidFirewallL7FirewallRulesRequest{}
	return &this
}

// NewUpdateNetworkWirelessSsidFirewallL7FirewallRulesRequestWithDefaults instantiates a new UpdateNetworkWirelessSsidFirewallL7FirewallRulesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkWirelessSsidFirewallL7FirewallRulesRequestWithDefaults() *UpdateNetworkWirelessSsidFirewallL7FirewallRulesRequest {
	this := UpdateNetworkWirelessSsidFirewallL7FirewallRulesRequest{}
	return &this
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidFirewallL7FirewallRulesRequest) GetRules() []UpdateNetworkWirelessSsidFirewallL7FirewallRulesRequestRulesInner {
	if o == nil || isNil(o.Rules) {
		var ret []UpdateNetworkWirelessSsidFirewallL7FirewallRulesRequestRulesInner
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidFirewallL7FirewallRulesRequest) GetRulesOk() ([]UpdateNetworkWirelessSsidFirewallL7FirewallRulesRequestRulesInner, bool) {
	if o == nil || isNil(o.Rules) {
    return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidFirewallL7FirewallRulesRequest) HasRules() bool {
	if o != nil && !isNil(o.Rules) {
		return true
	}

	return false
}

// SetRules gets a reference to the given []UpdateNetworkWirelessSsidFirewallL7FirewallRulesRequestRulesInner and assigns it to the Rules field.
func (o *UpdateNetworkWirelessSsidFirewallL7FirewallRulesRequest) SetRules(v []UpdateNetworkWirelessSsidFirewallL7FirewallRulesRequestRulesInner) {
	o.Rules = v
}

func (o UpdateNetworkWirelessSsidFirewallL7FirewallRulesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Rules) {
		toSerialize["rules"] = o.Rules
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateNetworkWirelessSsidFirewallL7FirewallRulesRequest struct {
	value *UpdateNetworkWirelessSsidFirewallL7FirewallRulesRequest
	isSet bool
}

func (v NullableUpdateNetworkWirelessSsidFirewallL7FirewallRulesRequest) Get() *UpdateNetworkWirelessSsidFirewallL7FirewallRulesRequest {
	return v.value
}

func (v *NullableUpdateNetworkWirelessSsidFirewallL7FirewallRulesRequest) Set(val *UpdateNetworkWirelessSsidFirewallL7FirewallRulesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkWirelessSsidFirewallL7FirewallRulesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkWirelessSsidFirewallL7FirewallRulesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkWirelessSsidFirewallL7FirewallRulesRequest(val *UpdateNetworkWirelessSsidFirewallL7FirewallRulesRequest) *NullableUpdateNetworkWirelessSsidFirewallL7FirewallRulesRequest {
	return &NullableUpdateNetworkWirelessSsidFirewallL7FirewallRulesRequest{value: val, isSet: true}
}

func (v NullableUpdateNetworkWirelessSsidFirewallL7FirewallRulesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkWirelessSsidFirewallL7FirewallRulesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

