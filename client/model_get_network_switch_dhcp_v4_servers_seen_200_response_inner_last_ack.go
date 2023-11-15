/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 01 November, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.39.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// checks if the GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastAck type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastAck{}

// GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastAck Attributes of the server's last ack.
type GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastAck struct {
	// Last time the server was acked.
	Ts *time.Time `json:"ts,omitempty"`
	Ipv4 *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastAckIpv4 `json:"ipv4,omitempty"`
}

// NewGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastAck instantiates a new GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastAck object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastAck() *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastAck {
	this := GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastAck{}
	return &this
}

// NewGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastAckWithDefaults instantiates a new GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastAck object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastAckWithDefaults() *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastAck {
	this := GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastAck{}
	return &this
}

// GetTs returns the Ts field value if set, zero value otherwise.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastAck) GetTs() time.Time {
	if o == nil || IsNil(o.Ts) {
		var ret time.Time
		return ret
	}
	return *o.Ts
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastAck) GetTsOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Ts) {
		return nil, false
	}
	return o.Ts, true
}

// HasTs returns a boolean if a field has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastAck) HasTs() bool {
	if o != nil && !IsNil(o.Ts) {
		return true
	}

	return false
}

// SetTs gets a reference to the given time.Time and assigns it to the Ts field.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastAck) SetTs(v time.Time) {
	o.Ts = &v
}

// GetIpv4 returns the Ipv4 field value if set, zero value otherwise.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastAck) GetIpv4() GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastAckIpv4 {
	if o == nil || IsNil(o.Ipv4) {
		var ret GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastAckIpv4
		return ret
	}
	return *o.Ipv4
}

// GetIpv4Ok returns a tuple with the Ipv4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastAck) GetIpv4Ok() (*GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastAckIpv4, bool) {
	if o == nil || IsNil(o.Ipv4) {
		return nil, false
	}
	return o.Ipv4, true
}

// HasIpv4 returns a boolean if a field has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastAck) HasIpv4() bool {
	if o != nil && !IsNil(o.Ipv4) {
		return true
	}

	return false
}

// SetIpv4 gets a reference to the given GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastAckIpv4 and assigns it to the Ipv4 field.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastAck) SetIpv4(v GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastAckIpv4) {
	o.Ipv4 = &v
}

func (o GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastAck) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastAck) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ts) {
		toSerialize["ts"] = o.Ts
	}
	if !IsNil(o.Ipv4) {
		toSerialize["ipv4"] = o.Ipv4
	}
	return toSerialize, nil
}

type NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastAck struct {
	value *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastAck
	isSet bool
}

func (v NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastAck) Get() *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastAck {
	return v.value
}

func (v *NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastAck) Set(val *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastAck) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastAck) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastAck) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastAck(val *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastAck) *NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastAck {
	return &NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastAck{value: val, isSet: true}
}

func (v NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastAck) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastAck) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


