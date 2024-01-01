/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 18 December, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.41.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the GetNetworkApplianceVlans200ResponseInnerFixedIpAssignmentsValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkApplianceVlans200ResponseInnerFixedIpAssignmentsValue{}

// GetNetworkApplianceVlans200ResponseInnerFixedIpAssignmentsValue IP assignment information, keyed by MAC address of the device
type GetNetworkApplianceVlans200ResponseInnerFixedIpAssignmentsValue struct {
	// IP address of the assignment
	Ip *string `json:"ip,omitempty"`
	// Name of the IP assignment
	Name *string `json:"name,omitempty"`
}

// NewGetNetworkApplianceVlans200ResponseInnerFixedIpAssignmentsValue instantiates a new GetNetworkApplianceVlans200ResponseInnerFixedIpAssignmentsValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkApplianceVlans200ResponseInnerFixedIpAssignmentsValue() *GetNetworkApplianceVlans200ResponseInnerFixedIpAssignmentsValue {
	this := GetNetworkApplianceVlans200ResponseInnerFixedIpAssignmentsValue{}
	return &this
}

// NewGetNetworkApplianceVlans200ResponseInnerFixedIpAssignmentsValueWithDefaults instantiates a new GetNetworkApplianceVlans200ResponseInnerFixedIpAssignmentsValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkApplianceVlans200ResponseInnerFixedIpAssignmentsValueWithDefaults() *GetNetworkApplianceVlans200ResponseInnerFixedIpAssignmentsValue {
	this := GetNetworkApplianceVlans200ResponseInnerFixedIpAssignmentsValue{}
	return &this
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *GetNetworkApplianceVlans200ResponseInnerFixedIpAssignmentsValue) GetIp() string {
	if o == nil || IsNil(o.Ip) {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceVlans200ResponseInnerFixedIpAssignmentsValue) GetIpOk() (*string, bool) {
	if o == nil || IsNil(o.Ip) {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *GetNetworkApplianceVlans200ResponseInnerFixedIpAssignmentsValue) HasIp() bool {
	if o != nil && !IsNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *GetNetworkApplianceVlans200ResponseInnerFixedIpAssignmentsValue) SetIp(v string) {
	o.Ip = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetNetworkApplianceVlans200ResponseInnerFixedIpAssignmentsValue) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceVlans200ResponseInnerFixedIpAssignmentsValue) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetNetworkApplianceVlans200ResponseInnerFixedIpAssignmentsValue) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetNetworkApplianceVlans200ResponseInnerFixedIpAssignmentsValue) SetName(v string) {
	o.Name = &v
}

func (o GetNetworkApplianceVlans200ResponseInnerFixedIpAssignmentsValue) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkApplianceVlans200ResponseInnerFixedIpAssignmentsValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ip) {
		toSerialize["ip"] = o.Ip
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableGetNetworkApplianceVlans200ResponseInnerFixedIpAssignmentsValue struct {
	value *GetNetworkApplianceVlans200ResponseInnerFixedIpAssignmentsValue
	isSet bool
}

func (v NullableGetNetworkApplianceVlans200ResponseInnerFixedIpAssignmentsValue) Get() *GetNetworkApplianceVlans200ResponseInnerFixedIpAssignmentsValue {
	return v.value
}

func (v *NullableGetNetworkApplianceVlans200ResponseInnerFixedIpAssignmentsValue) Set(val *GetNetworkApplianceVlans200ResponseInnerFixedIpAssignmentsValue) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkApplianceVlans200ResponseInnerFixedIpAssignmentsValue) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkApplianceVlans200ResponseInnerFixedIpAssignmentsValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkApplianceVlans200ResponseInnerFixedIpAssignmentsValue(val *GetNetworkApplianceVlans200ResponseInnerFixedIpAssignmentsValue) *NullableGetNetworkApplianceVlans200ResponseInnerFixedIpAssignmentsValue {
	return &NullableGetNetworkApplianceVlans200ResponseInnerFixedIpAssignmentsValue{value: val, isSet: true}
}

func (v NullableGetNetworkApplianceVlans200ResponseInnerFixedIpAssignmentsValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkApplianceVlans200ResponseInnerFixedIpAssignmentsValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

