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

// checks if the UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInner{}

// UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInner struct for UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInner
type UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInner struct {
	// A descriptive name for the rule
	Name *string `json:"name,omitempty"`
	// The IP address that will be used to access the internal resource from the WAN
	PublicIp *string `json:"publicIp,omitempty"`
	// The IP address of the server or device that hosts the internal resource that you wish to make available on the WAN
	LanIp string `json:"lanIp"`
	// The physical WAN interface on which the traffic will arrive ('internet1' or, if available, 'internet2')
	Uplink *string `json:"uplink,omitempty"`
	// The ports this mapping will provide access on, and the remote IPs that will be allowed access to the resource
	AllowedInbound []UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInnerAllowedInboundInner `json:"allowedInbound,omitempty"`
}

// NewUpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInner instantiates a new UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInner(lanIp string) *UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInner {
	this := UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInner{}
	this.LanIp = lanIp
	return &this
}

// NewUpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInnerWithDefaults instantiates a new UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInnerWithDefaults() *UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInner {
	this := UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInner{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInner) SetName(v string) {
	o.Name = &v
}

// GetPublicIp returns the PublicIp field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInner) GetPublicIp() string {
	if o == nil || IsNil(o.PublicIp) {
		var ret string
		return ret
	}
	return *o.PublicIp
}

// GetPublicIpOk returns a tuple with the PublicIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInner) GetPublicIpOk() (*string, bool) {
	if o == nil || IsNil(o.PublicIp) {
		return nil, false
	}
	return o.PublicIp, true
}

// HasPublicIp returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInner) HasPublicIp() bool {
	if o != nil && !IsNil(o.PublicIp) {
		return true
	}

	return false
}

// SetPublicIp gets a reference to the given string and assigns it to the PublicIp field.
func (o *UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInner) SetPublicIp(v string) {
	o.PublicIp = &v
}

// GetLanIp returns the LanIp field value
func (o *UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInner) GetLanIp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LanIp
}

// GetLanIpOk returns a tuple with the LanIp field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInner) GetLanIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LanIp, true
}

// SetLanIp sets field value
func (o *UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInner) SetLanIp(v string) {
	o.LanIp = v
}

// GetUplink returns the Uplink field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInner) GetUplink() string {
	if o == nil || IsNil(o.Uplink) {
		var ret string
		return ret
	}
	return *o.Uplink
}

// GetUplinkOk returns a tuple with the Uplink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInner) GetUplinkOk() (*string, bool) {
	if o == nil || IsNil(o.Uplink) {
		return nil, false
	}
	return o.Uplink, true
}

// HasUplink returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInner) HasUplink() bool {
	if o != nil && !IsNil(o.Uplink) {
		return true
	}

	return false
}

// SetUplink gets a reference to the given string and assigns it to the Uplink field.
func (o *UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInner) SetUplink(v string) {
	o.Uplink = &v
}

// GetAllowedInbound returns the AllowedInbound field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInner) GetAllowedInbound() []UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInnerAllowedInboundInner {
	if o == nil || IsNil(o.AllowedInbound) {
		var ret []UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInnerAllowedInboundInner
		return ret
	}
	return o.AllowedInbound
}

// GetAllowedInboundOk returns a tuple with the AllowedInbound field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInner) GetAllowedInboundOk() ([]UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInnerAllowedInboundInner, bool) {
	if o == nil || IsNil(o.AllowedInbound) {
		return nil, false
	}
	return o.AllowedInbound, true
}

// HasAllowedInbound returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInner) HasAllowedInbound() bool {
	if o != nil && !IsNil(o.AllowedInbound) {
		return true
	}

	return false
}

// SetAllowedInbound gets a reference to the given []UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInnerAllowedInboundInner and assigns it to the AllowedInbound field.
func (o *UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInner) SetAllowedInbound(v []UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInnerAllowedInboundInner) {
	o.AllowedInbound = v
}

func (o UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.PublicIp) {
		toSerialize["publicIp"] = o.PublicIp
	}
	toSerialize["lanIp"] = o.LanIp
	if !IsNil(o.Uplink) {
		toSerialize["uplink"] = o.Uplink
	}
	if !IsNil(o.AllowedInbound) {
		toSerialize["allowedInbound"] = o.AllowedInbound
	}
	return toSerialize, nil
}

type NullableUpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInner struct {
	value *UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInner
	isSet bool
}

func (v NullableUpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInner) Get() *UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInner {
	return v.value
}

func (v *NullableUpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInner) Set(val *UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInner(val *UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInner) *NullableUpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInner {
	return &NullableUpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInner{value: val, isSet: true}
}

func (v NullableUpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


