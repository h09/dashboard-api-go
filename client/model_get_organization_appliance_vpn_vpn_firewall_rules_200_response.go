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

// checks if the GetOrganizationApplianceVpnVpnFirewallRules200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationApplianceVpnVpnFirewallRules200Response{}

// GetOrganizationApplianceVpnVpnFirewallRules200Response struct for GetOrganizationApplianceVpnVpnFirewallRules200Response
type GetOrganizationApplianceVpnVpnFirewallRules200Response struct {
	// An ordered array of the firewall rules (not including the default rule)
	Rules []GetNetworkApplianceFirewallInboundFirewallRules200ResponseRulesInner `json:"rules,omitempty"`
}

// NewGetOrganizationApplianceVpnVpnFirewallRules200Response instantiates a new GetOrganizationApplianceVpnVpnFirewallRules200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationApplianceVpnVpnFirewallRules200Response() *GetOrganizationApplianceVpnVpnFirewallRules200Response {
	this := GetOrganizationApplianceVpnVpnFirewallRules200Response{}
	return &this
}

// NewGetOrganizationApplianceVpnVpnFirewallRules200ResponseWithDefaults instantiates a new GetOrganizationApplianceVpnVpnFirewallRules200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationApplianceVpnVpnFirewallRules200ResponseWithDefaults() *GetOrganizationApplianceVpnVpnFirewallRules200Response {
	this := GetOrganizationApplianceVpnVpnFirewallRules200Response{}
	return &this
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *GetOrganizationApplianceVpnVpnFirewallRules200Response) GetRules() []GetNetworkApplianceFirewallInboundFirewallRules200ResponseRulesInner {
	if o == nil || IsNil(o.Rules) {
		var ret []GetNetworkApplianceFirewallInboundFirewallRules200ResponseRulesInner
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationApplianceVpnVpnFirewallRules200Response) GetRulesOk() ([]GetNetworkApplianceFirewallInboundFirewallRules200ResponseRulesInner, bool) {
	if o == nil || IsNil(o.Rules) {
		return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *GetOrganizationApplianceVpnVpnFirewallRules200Response) HasRules() bool {
	if o != nil && !IsNil(o.Rules) {
		return true
	}

	return false
}

// SetRules gets a reference to the given []GetNetworkApplianceFirewallInboundFirewallRules200ResponseRulesInner and assigns it to the Rules field.
func (o *GetOrganizationApplianceVpnVpnFirewallRules200Response) SetRules(v []GetNetworkApplianceFirewallInboundFirewallRules200ResponseRulesInner) {
	o.Rules = v
}

func (o GetOrganizationApplianceVpnVpnFirewallRules200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationApplianceVpnVpnFirewallRules200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Rules) {
		toSerialize["rules"] = o.Rules
	}
	return toSerialize, nil
}

type NullableGetOrganizationApplianceVpnVpnFirewallRules200Response struct {
	value *GetOrganizationApplianceVpnVpnFirewallRules200Response
	isSet bool
}

func (v NullableGetOrganizationApplianceVpnVpnFirewallRules200Response) Get() *GetOrganizationApplianceVpnVpnFirewallRules200Response {
	return v.value
}

func (v *NullableGetOrganizationApplianceVpnVpnFirewallRules200Response) Set(val *GetOrganizationApplianceVpnVpnFirewallRules200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationApplianceVpnVpnFirewallRules200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationApplianceVpnVpnFirewallRules200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationApplianceVpnVpnFirewallRules200Response(val *GetOrganizationApplianceVpnVpnFirewallRules200Response) *NullableGetOrganizationApplianceVpnVpnFirewallRules200Response {
	return &NullableGetOrganizationApplianceVpnVpnFirewallRules200Response{value: val, isSet: true}
}

func (v NullableGetOrganizationApplianceVpnVpnFirewallRules200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationApplianceVpnVpnFirewallRules200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


