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

// UpdateNetworkWirelessSsidFirewallL3FirewallRulesRequestRulesInner struct for UpdateNetworkWirelessSsidFirewallL3FirewallRulesRequestRulesInner
type UpdateNetworkWirelessSsidFirewallL3FirewallRulesRequestRulesInner struct {
	// Description of the rule (optional)
	Comment *string `json:"comment,omitempty"`
	// 'allow' or 'deny' traffic specified by this rule
	Policy string `json:"policy"`
	// The type of protocol (must be 'tcp', 'udp', 'icmp', 'icmp6' or 'any')
	Protocol string `json:"protocol"`
	// Comma-separated list of destination port(s) (integer in the range 1-65535), or 'any'
	DestPort *string `json:"destPort,omitempty"`
	// Comma-separated list of destination IP address(es) (in IP or CIDR notation), fully-qualified domain names (FQDN) or 'any'
	DestCidr string `json:"destCidr"`
}

// NewUpdateNetworkWirelessSsidFirewallL3FirewallRulesRequestRulesInner instantiates a new UpdateNetworkWirelessSsidFirewallL3FirewallRulesRequestRulesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkWirelessSsidFirewallL3FirewallRulesRequestRulesInner(policy string, protocol string, destCidr string) *UpdateNetworkWirelessSsidFirewallL3FirewallRulesRequestRulesInner {
	this := UpdateNetworkWirelessSsidFirewallL3FirewallRulesRequestRulesInner{}
	this.Policy = policy
	this.Protocol = protocol
	this.DestCidr = destCidr
	return &this
}

// NewUpdateNetworkWirelessSsidFirewallL3FirewallRulesRequestRulesInnerWithDefaults instantiates a new UpdateNetworkWirelessSsidFirewallL3FirewallRulesRequestRulesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkWirelessSsidFirewallL3FirewallRulesRequestRulesInnerWithDefaults() *UpdateNetworkWirelessSsidFirewallL3FirewallRulesRequestRulesInner {
	this := UpdateNetworkWirelessSsidFirewallL3FirewallRulesRequestRulesInner{}
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidFirewallL3FirewallRulesRequestRulesInner) GetComment() string {
	if o == nil || isNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidFirewallL3FirewallRulesRequestRulesInner) GetCommentOk() (*string, bool) {
	if o == nil || isNil(o.Comment) {
    return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidFirewallL3FirewallRulesRequestRulesInner) HasComment() bool {
	if o != nil && !isNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *UpdateNetworkWirelessSsidFirewallL3FirewallRulesRequestRulesInner) SetComment(v string) {
	o.Comment = &v
}

// GetPolicy returns the Policy field value
func (o *UpdateNetworkWirelessSsidFirewallL3FirewallRulesRequestRulesInner) GetPolicy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Policy
}

// GetPolicyOk returns a tuple with the Policy field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidFirewallL3FirewallRulesRequestRulesInner) GetPolicyOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Policy, true
}

// SetPolicy sets field value
func (o *UpdateNetworkWirelessSsidFirewallL3FirewallRulesRequestRulesInner) SetPolicy(v string) {
	o.Policy = v
}

// GetProtocol returns the Protocol field value
func (o *UpdateNetworkWirelessSsidFirewallL3FirewallRulesRequestRulesInner) GetProtocol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidFirewallL3FirewallRulesRequestRulesInner) GetProtocolOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Protocol, true
}

// SetProtocol sets field value
func (o *UpdateNetworkWirelessSsidFirewallL3FirewallRulesRequestRulesInner) SetProtocol(v string) {
	o.Protocol = v
}

// GetDestPort returns the DestPort field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidFirewallL3FirewallRulesRequestRulesInner) GetDestPort() string {
	if o == nil || isNil(o.DestPort) {
		var ret string
		return ret
	}
	return *o.DestPort
}

// GetDestPortOk returns a tuple with the DestPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidFirewallL3FirewallRulesRequestRulesInner) GetDestPortOk() (*string, bool) {
	if o == nil || isNil(o.DestPort) {
    return nil, false
	}
	return o.DestPort, true
}

// HasDestPort returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidFirewallL3FirewallRulesRequestRulesInner) HasDestPort() bool {
	if o != nil && !isNil(o.DestPort) {
		return true
	}

	return false
}

// SetDestPort gets a reference to the given string and assigns it to the DestPort field.
func (o *UpdateNetworkWirelessSsidFirewallL3FirewallRulesRequestRulesInner) SetDestPort(v string) {
	o.DestPort = &v
}

// GetDestCidr returns the DestCidr field value
func (o *UpdateNetworkWirelessSsidFirewallL3FirewallRulesRequestRulesInner) GetDestCidr() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DestCidr
}

// GetDestCidrOk returns a tuple with the DestCidr field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidFirewallL3FirewallRulesRequestRulesInner) GetDestCidrOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.DestCidr, true
}

// SetDestCidr sets field value
func (o *UpdateNetworkWirelessSsidFirewallL3FirewallRulesRequestRulesInner) SetDestCidr(v string) {
	o.DestCidr = v
}

func (o UpdateNetworkWirelessSsidFirewallL3FirewallRulesRequestRulesInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if true {
		toSerialize["policy"] = o.Policy
	}
	if true {
		toSerialize["protocol"] = o.Protocol
	}
	if !isNil(o.DestPort) {
		toSerialize["destPort"] = o.DestPort
	}
	if true {
		toSerialize["destCidr"] = o.DestCidr
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateNetworkWirelessSsidFirewallL3FirewallRulesRequestRulesInner struct {
	value *UpdateNetworkWirelessSsidFirewallL3FirewallRulesRequestRulesInner
	isSet bool
}

func (v NullableUpdateNetworkWirelessSsidFirewallL3FirewallRulesRequestRulesInner) Get() *UpdateNetworkWirelessSsidFirewallL3FirewallRulesRequestRulesInner {
	return v.value
}

func (v *NullableUpdateNetworkWirelessSsidFirewallL3FirewallRulesRequestRulesInner) Set(val *UpdateNetworkWirelessSsidFirewallL3FirewallRulesRequestRulesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkWirelessSsidFirewallL3FirewallRulesRequestRulesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkWirelessSsidFirewallL3FirewallRulesRequestRulesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkWirelessSsidFirewallL3FirewallRulesRequestRulesInner(val *UpdateNetworkWirelessSsidFirewallL3FirewallRulesRequestRulesInner) *NullableUpdateNetworkWirelessSsidFirewallL3FirewallRulesRequestRulesInner {
	return &NullableUpdateNetworkWirelessSsidFirewallL3FirewallRulesRequestRulesInner{value: val, isSet: true}
}

func (v NullableUpdateNetworkWirelessSsidFirewallL3FirewallRulesRequestRulesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkWirelessSsidFirewallL3FirewallRulesRequestRulesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

