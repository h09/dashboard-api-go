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

// GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInner struct for GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInner
type GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInner struct {
	// Label of the port
	Name *string `json:"name,omitempty"`
	Poe *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInnerPoe `json:"poe,omitempty"`
	LinkNegotiation *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInnerLinkNegotiation `json:"linkNegotiation,omitempty"`
}

// NewGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInner instantiates a new GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInner() *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInner {
	this := GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInner{}
	return &this
}

// NewGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInnerWithDefaults instantiates a new GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInnerWithDefaults() *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInner {
	this := GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInner{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInner) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInner) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInner) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInner) SetName(v string) {
	o.Name = &v
}

// GetPoe returns the Poe field value if set, zero value otherwise.
func (o *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInner) GetPoe() GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInnerPoe {
	if o == nil || isNil(o.Poe) {
		var ret GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInnerPoe
		return ret
	}
	return *o.Poe
}

// GetPoeOk returns a tuple with the Poe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInner) GetPoeOk() (*GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInnerPoe, bool) {
	if o == nil || isNil(o.Poe) {
    return nil, false
	}
	return o.Poe, true
}

// HasPoe returns a boolean if a field has been set.
func (o *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInner) HasPoe() bool {
	if o != nil && !isNil(o.Poe) {
		return true
	}

	return false
}

// SetPoe gets a reference to the given GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInnerPoe and assigns it to the Poe field.
func (o *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInner) SetPoe(v GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInnerPoe) {
	o.Poe = &v
}

// GetLinkNegotiation returns the LinkNegotiation field value if set, zero value otherwise.
func (o *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInner) GetLinkNegotiation() GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInnerLinkNegotiation {
	if o == nil || isNil(o.LinkNegotiation) {
		var ret GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInnerLinkNegotiation
		return ret
	}
	return *o.LinkNegotiation
}

// GetLinkNegotiationOk returns a tuple with the LinkNegotiation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInner) GetLinkNegotiationOk() (*GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInnerLinkNegotiation, bool) {
	if o == nil || isNil(o.LinkNegotiation) {
    return nil, false
	}
	return o.LinkNegotiation, true
}

// HasLinkNegotiation returns a boolean if a field has been set.
func (o *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInner) HasLinkNegotiation() bool {
	if o != nil && !isNil(o.LinkNegotiation) {
		return true
	}

	return false
}

// SetLinkNegotiation gets a reference to the given GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInnerLinkNegotiation and assigns it to the LinkNegotiation field.
func (o *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInner) SetLinkNegotiation(v GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInnerLinkNegotiation) {
	o.LinkNegotiation = &v
}

func (o GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Poe) {
		toSerialize["poe"] = o.Poe
	}
	if !isNil(o.LinkNegotiation) {
		toSerialize["linkNegotiation"] = o.LinkNegotiation
	}
	return json.Marshal(toSerialize)
}

type NullableGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInner struct {
	value *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInner
	isSet bool
}

func (v NullableGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInner) Get() *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInner {
	return v.value
}

func (v *NullableGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInner) Set(val *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInner(val *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInner) *NullableGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInner {
	return &NullableGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInner{value: val, isSet: true}
}

func (v NullableGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPortsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

