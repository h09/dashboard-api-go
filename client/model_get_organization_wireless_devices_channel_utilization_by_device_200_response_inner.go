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

// checks if the GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInner{}

// GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInner struct for GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInner
type GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInner struct {
	// The serial number for the device.
	Serial *string `json:"serial,omitempty"`
	// The MAC address of the device.
	Mac *string `json:"mac,omitempty"`
	Network *GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerNetwork `json:"network,omitempty"`
	// Channel utilization broken down by band.
	ByBand []GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInner `json:"byBand,omitempty"`
}

// NewGetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInner instantiates a new GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInner() *GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInner {
	this := GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInner{}
	return &this
}

// NewGetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerWithDefaults instantiates a new GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerWithDefaults() *GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInner {
	this := GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInner{}
	return &this
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInner) GetSerial() string {
	if o == nil || IsNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInner) GetSerialOk() (*string, bool) {
	if o == nil || IsNil(o.Serial) {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInner) HasSerial() bool {
	if o != nil && !IsNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInner) SetSerial(v string) {
	o.Serial = &v
}

// GetMac returns the Mac field value if set, zero value otherwise.
func (o *GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInner) GetMac() string {
	if o == nil || IsNil(o.Mac) {
		var ret string
		return ret
	}
	return *o.Mac
}

// GetMacOk returns a tuple with the Mac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInner) GetMacOk() (*string, bool) {
	if o == nil || IsNil(o.Mac) {
		return nil, false
	}
	return o.Mac, true
}

// HasMac returns a boolean if a field has been set.
func (o *GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInner) HasMac() bool {
	if o != nil && !IsNil(o.Mac) {
		return true
	}

	return false
}

// SetMac gets a reference to the given string and assigns it to the Mac field.
func (o *GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInner) SetMac(v string) {
	o.Mac = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInner) GetNetwork() GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerNetwork {
	if o == nil || IsNil(o.Network) {
		var ret GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerNetwork
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInner) GetNetworkOk() (*GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerNetwork, bool) {
	if o == nil || IsNil(o.Network) {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInner) HasNetwork() bool {
	if o != nil && !IsNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerNetwork and assigns it to the Network field.
func (o *GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInner) SetNetwork(v GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerNetwork) {
	o.Network = &v
}

// GetByBand returns the ByBand field value if set, zero value otherwise.
func (o *GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInner) GetByBand() []GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInner {
	if o == nil || IsNil(o.ByBand) {
		var ret []GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInner
		return ret
	}
	return o.ByBand
}

// GetByBandOk returns a tuple with the ByBand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInner) GetByBandOk() ([]GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInner, bool) {
	if o == nil || IsNil(o.ByBand) {
		return nil, false
	}
	return o.ByBand, true
}

// HasByBand returns a boolean if a field has been set.
func (o *GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInner) HasByBand() bool {
	if o != nil && !IsNil(o.ByBand) {
		return true
	}

	return false
}

// SetByBand gets a reference to the given []GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInner and assigns it to the ByBand field.
func (o *GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInner) SetByBand(v []GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInner) {
	o.ByBand = v
}

func (o GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !IsNil(o.Mac) {
		toSerialize["mac"] = o.Mac
	}
	if !IsNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !IsNil(o.ByBand) {
		toSerialize["byBand"] = o.ByBand
	}
	return toSerialize, nil
}

type NullableGetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInner struct {
	value *GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInner
	isSet bool
}

func (v NullableGetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInner) Get() *GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInner {
	return v.value
}

func (v *NullableGetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInner) Set(val *GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInner(val *GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInner) *NullableGetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInner {
	return &NullableGetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInner{value: val, isSet: true}
}

func (v NullableGetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


