/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 February, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: v1.30.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInnerPoe PoE details object for the port
type GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInnerPoe struct {
	// The PoE Standard for the port. Can be '802.3at', '802.3af', '802.3bt', or null
	Standard *string `json:"standard,omitempty"`
}

// NewGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInnerPoe instantiates a new GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInnerPoe object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInnerPoe() *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInnerPoe {
	this := GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInnerPoe{}
	return &this
}

// NewGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInnerPoeWithDefaults instantiates a new GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInnerPoe object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInnerPoeWithDefaults() *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInnerPoe {
	this := GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInnerPoe{}
	return &this
}

// GetStandard returns the Standard field value if set, zero value otherwise.
func (o *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInnerPoe) GetStandard() string {
	if o == nil || isNil(o.Standard) {
		var ret string
		return ret
	}
	return *o.Standard
}

// GetStandardOk returns a tuple with the Standard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInnerPoe) GetStandardOk() (*string, bool) {
	if o == nil || isNil(o.Standard) {
    return nil, false
	}
	return o.Standard, true
}

// HasStandard returns a boolean if a field has been set.
func (o *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInnerPoe) HasStandard() bool {
	if o != nil && !isNil(o.Standard) {
		return true
	}

	return false
}

// SetStandard gets a reference to the given string and assigns it to the Standard field.
func (o *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInnerPoe) SetStandard(v string) {
	o.Standard = &v
}

func (o GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInnerPoe) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Standard) {
		toSerialize["standard"] = o.Standard
	}
	return json.Marshal(toSerialize)
}

type NullableGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInnerPoe struct {
	value *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInnerPoe
	isSet bool
}

func (v NullableGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInnerPoe) Get() *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInnerPoe {
	return v.value
}

func (v *NullableGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInnerPoe) Set(val *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInnerPoe) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInnerPoe) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInnerPoe) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInnerPoe(val *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInnerPoe) *NullableGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInnerPoe {
	return &NullableGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInnerPoe{value: val, isSet: true}
}

func (v NullableGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInnerPoe) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInnerPoe) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

