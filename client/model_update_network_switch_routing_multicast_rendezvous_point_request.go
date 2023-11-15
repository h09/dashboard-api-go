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

// checks if the UpdateNetworkSwitchRoutingMulticastRendezvousPointRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkSwitchRoutingMulticastRendezvousPointRequest{}

// UpdateNetworkSwitchRoutingMulticastRendezvousPointRequest struct for UpdateNetworkSwitchRoutingMulticastRendezvousPointRequest
type UpdateNetworkSwitchRoutingMulticastRendezvousPointRequest struct {
	// The IP address of the interface to use
	InterfaceIp string `json:"interfaceIp"`
	// 'Any', or the IP address of a multicast group
	MulticastGroup string `json:"multicastGroup"`
}

// NewUpdateNetworkSwitchRoutingMulticastRendezvousPointRequest instantiates a new UpdateNetworkSwitchRoutingMulticastRendezvousPointRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkSwitchRoutingMulticastRendezvousPointRequest(interfaceIp string, multicastGroup string) *UpdateNetworkSwitchRoutingMulticastRendezvousPointRequest {
	this := UpdateNetworkSwitchRoutingMulticastRendezvousPointRequest{}
	this.InterfaceIp = interfaceIp
	this.MulticastGroup = multicastGroup
	return &this
}

// NewUpdateNetworkSwitchRoutingMulticastRendezvousPointRequestWithDefaults instantiates a new UpdateNetworkSwitchRoutingMulticastRendezvousPointRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkSwitchRoutingMulticastRendezvousPointRequestWithDefaults() *UpdateNetworkSwitchRoutingMulticastRendezvousPointRequest {
	this := UpdateNetworkSwitchRoutingMulticastRendezvousPointRequest{}
	return &this
}

// GetInterfaceIp returns the InterfaceIp field value
func (o *UpdateNetworkSwitchRoutingMulticastRendezvousPointRequest) GetInterfaceIp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InterfaceIp
}

// GetInterfaceIpOk returns a tuple with the InterfaceIp field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSwitchRoutingMulticastRendezvousPointRequest) GetInterfaceIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InterfaceIp, true
}

// SetInterfaceIp sets field value
func (o *UpdateNetworkSwitchRoutingMulticastRendezvousPointRequest) SetInterfaceIp(v string) {
	o.InterfaceIp = v
}

// GetMulticastGroup returns the MulticastGroup field value
func (o *UpdateNetworkSwitchRoutingMulticastRendezvousPointRequest) GetMulticastGroup() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MulticastGroup
}

// GetMulticastGroupOk returns a tuple with the MulticastGroup field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSwitchRoutingMulticastRendezvousPointRequest) GetMulticastGroupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MulticastGroup, true
}

// SetMulticastGroup sets field value
func (o *UpdateNetworkSwitchRoutingMulticastRendezvousPointRequest) SetMulticastGroup(v string) {
	o.MulticastGroup = v
}

func (o UpdateNetworkSwitchRoutingMulticastRendezvousPointRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkSwitchRoutingMulticastRendezvousPointRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["interfaceIp"] = o.InterfaceIp
	toSerialize["multicastGroup"] = o.MulticastGroup
	return toSerialize, nil
}

type NullableUpdateNetworkSwitchRoutingMulticastRendezvousPointRequest struct {
	value *UpdateNetworkSwitchRoutingMulticastRendezvousPointRequest
	isSet bool
}

func (v NullableUpdateNetworkSwitchRoutingMulticastRendezvousPointRequest) Get() *UpdateNetworkSwitchRoutingMulticastRendezvousPointRequest {
	return v.value
}

func (v *NullableUpdateNetworkSwitchRoutingMulticastRendezvousPointRequest) Set(val *UpdateNetworkSwitchRoutingMulticastRendezvousPointRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkSwitchRoutingMulticastRendezvousPointRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkSwitchRoutingMulticastRendezvousPointRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkSwitchRoutingMulticastRendezvousPointRequest(val *UpdateNetworkSwitchRoutingMulticastRendezvousPointRequest) *NullableUpdateNetworkSwitchRoutingMulticastRendezvousPointRequest {
	return &NullableUpdateNetworkSwitchRoutingMulticastRendezvousPointRequest{value: val, isSet: true}
}

func (v NullableUpdateNetworkSwitchRoutingMulticastRendezvousPointRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkSwitchRoutingMulticastRendezvousPointRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


