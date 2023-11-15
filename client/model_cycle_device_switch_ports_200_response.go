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

// checks if the CycleDeviceSwitchPorts200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CycleDeviceSwitchPorts200Response{}

// CycleDeviceSwitchPorts200Response struct for CycleDeviceSwitchPorts200Response
type CycleDeviceSwitchPorts200Response struct {
	// List of switch ports
	Ports []string `json:"ports,omitempty"`
}

// NewCycleDeviceSwitchPorts200Response instantiates a new CycleDeviceSwitchPorts200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCycleDeviceSwitchPorts200Response() *CycleDeviceSwitchPorts200Response {
	this := CycleDeviceSwitchPorts200Response{}
	return &this
}

// NewCycleDeviceSwitchPorts200ResponseWithDefaults instantiates a new CycleDeviceSwitchPorts200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCycleDeviceSwitchPorts200ResponseWithDefaults() *CycleDeviceSwitchPorts200Response {
	this := CycleDeviceSwitchPorts200Response{}
	return &this
}

// GetPorts returns the Ports field value if set, zero value otherwise.
func (o *CycleDeviceSwitchPorts200Response) GetPorts() []string {
	if o == nil || IsNil(o.Ports) {
		var ret []string
		return ret
	}
	return o.Ports
}

// GetPortsOk returns a tuple with the Ports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CycleDeviceSwitchPorts200Response) GetPortsOk() ([]string, bool) {
	if o == nil || IsNil(o.Ports) {
		return nil, false
	}
	return o.Ports, true
}

// HasPorts returns a boolean if a field has been set.
func (o *CycleDeviceSwitchPorts200Response) HasPorts() bool {
	if o != nil && !IsNil(o.Ports) {
		return true
	}

	return false
}

// SetPorts gets a reference to the given []string and assigns it to the Ports field.
func (o *CycleDeviceSwitchPorts200Response) SetPorts(v []string) {
	o.Ports = v
}

func (o CycleDeviceSwitchPorts200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CycleDeviceSwitchPorts200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ports) {
		toSerialize["ports"] = o.Ports
	}
	return toSerialize, nil
}

type NullableCycleDeviceSwitchPorts200Response struct {
	value *CycleDeviceSwitchPorts200Response
	isSet bool
}

func (v NullableCycleDeviceSwitchPorts200Response) Get() *CycleDeviceSwitchPorts200Response {
	return v.value
}

func (v *NullableCycleDeviceSwitchPorts200Response) Set(val *CycleDeviceSwitchPorts200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCycleDeviceSwitchPorts200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCycleDeviceSwitchPorts200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCycleDeviceSwitchPorts200Response(val *CycleDeviceSwitchPorts200Response) *NullableCycleDeviceSwitchPorts200Response {
	return &NullableCycleDeviceSwitchPorts200Response{value: val, isSet: true}
}

func (v NullableCycleDeviceSwitchPorts200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCycleDeviceSwitchPorts200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


