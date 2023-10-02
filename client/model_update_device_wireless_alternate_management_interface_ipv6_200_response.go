/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 06 September, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.37.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the UpdateDeviceWirelessAlternateManagementInterfaceIpv6200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateDeviceWirelessAlternateManagementInterfaceIpv6200Response{}

// UpdateDeviceWirelessAlternateManagementInterfaceIpv6200Response struct for UpdateDeviceWirelessAlternateManagementInterfaceIpv6200Response
type UpdateDeviceWirelessAlternateManagementInterfaceIpv6200Response struct {
	// configured alternate management interface addresses
	Addresses []UpdateDeviceWirelessAlternateManagementInterfaceIpv6200ResponseAddressesInner `json:"addresses,omitempty"`
}

// NewUpdateDeviceWirelessAlternateManagementInterfaceIpv6200Response instantiates a new UpdateDeviceWirelessAlternateManagementInterfaceIpv6200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateDeviceWirelessAlternateManagementInterfaceIpv6200Response() *UpdateDeviceWirelessAlternateManagementInterfaceIpv6200Response {
	this := UpdateDeviceWirelessAlternateManagementInterfaceIpv6200Response{}
	return &this
}

// NewUpdateDeviceWirelessAlternateManagementInterfaceIpv6200ResponseWithDefaults instantiates a new UpdateDeviceWirelessAlternateManagementInterfaceIpv6200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateDeviceWirelessAlternateManagementInterfaceIpv6200ResponseWithDefaults() *UpdateDeviceWirelessAlternateManagementInterfaceIpv6200Response {
	this := UpdateDeviceWirelessAlternateManagementInterfaceIpv6200Response{}
	return &this
}

// GetAddresses returns the Addresses field value if set, zero value otherwise.
func (o *UpdateDeviceWirelessAlternateManagementInterfaceIpv6200Response) GetAddresses() []UpdateDeviceWirelessAlternateManagementInterfaceIpv6200ResponseAddressesInner {
	if o == nil || IsNil(o.Addresses) {
		var ret []UpdateDeviceWirelessAlternateManagementInterfaceIpv6200ResponseAddressesInner
		return ret
	}
	return o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceWirelessAlternateManagementInterfaceIpv6200Response) GetAddressesOk() ([]UpdateDeviceWirelessAlternateManagementInterfaceIpv6200ResponseAddressesInner, bool) {
	if o == nil || IsNil(o.Addresses) {
		return nil, false
	}
	return o.Addresses, true
}

// HasAddresses returns a boolean if a field has been set.
func (o *UpdateDeviceWirelessAlternateManagementInterfaceIpv6200Response) HasAddresses() bool {
	if o != nil && !IsNil(o.Addresses) {
		return true
	}

	return false
}

// SetAddresses gets a reference to the given []UpdateDeviceWirelessAlternateManagementInterfaceIpv6200ResponseAddressesInner and assigns it to the Addresses field.
func (o *UpdateDeviceWirelessAlternateManagementInterfaceIpv6200Response) SetAddresses(v []UpdateDeviceWirelessAlternateManagementInterfaceIpv6200ResponseAddressesInner) {
	o.Addresses = v
}

func (o UpdateDeviceWirelessAlternateManagementInterfaceIpv6200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateDeviceWirelessAlternateManagementInterfaceIpv6200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Addresses) {
		toSerialize["addresses"] = o.Addresses
	}
	return toSerialize, nil
}

type NullableUpdateDeviceWirelessAlternateManagementInterfaceIpv6200Response struct {
	value *UpdateDeviceWirelessAlternateManagementInterfaceIpv6200Response
	isSet bool
}

func (v NullableUpdateDeviceWirelessAlternateManagementInterfaceIpv6200Response) Get() *UpdateDeviceWirelessAlternateManagementInterfaceIpv6200Response {
	return v.value
}

func (v *NullableUpdateDeviceWirelessAlternateManagementInterfaceIpv6200Response) Set(val *UpdateDeviceWirelessAlternateManagementInterfaceIpv6200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateDeviceWirelessAlternateManagementInterfaceIpv6200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateDeviceWirelessAlternateManagementInterfaceIpv6200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateDeviceWirelessAlternateManagementInterfaceIpv6200Response(val *UpdateDeviceWirelessAlternateManagementInterfaceIpv6200Response) *NullableUpdateDeviceWirelessAlternateManagementInterfaceIpv6200Response {
	return &NullableUpdateDeviceWirelessAlternateManagementInterfaceIpv6200Response{value: val, isSet: true}
}

func (v NullableUpdateDeviceWirelessAlternateManagementInterfaceIpv6200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateDeviceWirelessAlternateManagementInterfaceIpv6200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

