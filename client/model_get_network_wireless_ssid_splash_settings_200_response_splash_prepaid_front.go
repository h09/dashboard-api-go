/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 05 July, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.35.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the GetNetworkWirelessSsidSplashSettings200ResponseSplashPrepaidFront type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkWirelessSsidSplashSettings200ResponseSplashPrepaidFront{}

// GetNetworkWirelessSsidSplashSettings200ResponseSplashPrepaidFront The prepaid front image used in the splash page.
type GetNetworkWirelessSsidSplashSettings200ResponseSplashPrepaidFront struct {
	// The MD5 value of the prepaid front image file.
	Md5 *string `json:"md5,omitempty"`
	// The extension of the prepaid front image file.
	Extension *string `json:"extension,omitempty"`
}

// NewGetNetworkWirelessSsidSplashSettings200ResponseSplashPrepaidFront instantiates a new GetNetworkWirelessSsidSplashSettings200ResponseSplashPrepaidFront object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkWirelessSsidSplashSettings200ResponseSplashPrepaidFront() *GetNetworkWirelessSsidSplashSettings200ResponseSplashPrepaidFront {
	this := GetNetworkWirelessSsidSplashSettings200ResponseSplashPrepaidFront{}
	return &this
}

// NewGetNetworkWirelessSsidSplashSettings200ResponseSplashPrepaidFrontWithDefaults instantiates a new GetNetworkWirelessSsidSplashSettings200ResponseSplashPrepaidFront object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkWirelessSsidSplashSettings200ResponseSplashPrepaidFrontWithDefaults() *GetNetworkWirelessSsidSplashSettings200ResponseSplashPrepaidFront {
	this := GetNetworkWirelessSsidSplashSettings200ResponseSplashPrepaidFront{}
	return &this
}

// GetMd5 returns the Md5 field value if set, zero value otherwise.
func (o *GetNetworkWirelessSsidSplashSettings200ResponseSplashPrepaidFront) GetMd5() string {
	if o == nil || IsNil(o.Md5) {
		var ret string
		return ret
	}
	return *o.Md5
}

// GetMd5Ok returns a tuple with the Md5 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessSsidSplashSettings200ResponseSplashPrepaidFront) GetMd5Ok() (*string, bool) {
	if o == nil || IsNil(o.Md5) {
		return nil, false
	}
	return o.Md5, true
}

// HasMd5 returns a boolean if a field has been set.
func (o *GetNetworkWirelessSsidSplashSettings200ResponseSplashPrepaidFront) HasMd5() bool {
	if o != nil && !IsNil(o.Md5) {
		return true
	}

	return false
}

// SetMd5 gets a reference to the given string and assigns it to the Md5 field.
func (o *GetNetworkWirelessSsidSplashSettings200ResponseSplashPrepaidFront) SetMd5(v string) {
	o.Md5 = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *GetNetworkWirelessSsidSplashSettings200ResponseSplashPrepaidFront) GetExtension() string {
	if o == nil || IsNil(o.Extension) {
		var ret string
		return ret
	}
	return *o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessSsidSplashSettings200ResponseSplashPrepaidFront) GetExtensionOk() (*string, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *GetNetworkWirelessSsidSplashSettings200ResponseSplashPrepaidFront) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given string and assigns it to the Extension field.
func (o *GetNetworkWirelessSsidSplashSettings200ResponseSplashPrepaidFront) SetExtension(v string) {
	o.Extension = &v
}

func (o GetNetworkWirelessSsidSplashSettings200ResponseSplashPrepaidFront) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkWirelessSsidSplashSettings200ResponseSplashPrepaidFront) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Md5) {
		toSerialize["md5"] = o.Md5
	}
	if !IsNil(o.Extension) {
		toSerialize["extension"] = o.Extension
	}
	return toSerialize, nil
}

type NullableGetNetworkWirelessSsidSplashSettings200ResponseSplashPrepaidFront struct {
	value *GetNetworkWirelessSsidSplashSettings200ResponseSplashPrepaidFront
	isSet bool
}

func (v NullableGetNetworkWirelessSsidSplashSettings200ResponseSplashPrepaidFront) Get() *GetNetworkWirelessSsidSplashSettings200ResponseSplashPrepaidFront {
	return v.value
}

func (v *NullableGetNetworkWirelessSsidSplashSettings200ResponseSplashPrepaidFront) Set(val *GetNetworkWirelessSsidSplashSettings200ResponseSplashPrepaidFront) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkWirelessSsidSplashSettings200ResponseSplashPrepaidFront) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkWirelessSsidSplashSettings200ResponseSplashPrepaidFront) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkWirelessSsidSplashSettings200ResponseSplashPrepaidFront(val *GetNetworkWirelessSsidSplashSettings200ResponseSplashPrepaidFront) *NullableGetNetworkWirelessSsidSplashSettings200ResponseSplashPrepaidFront {
	return &NullableGetNetworkWirelessSsidSplashSettings200ResponseSplashPrepaidFront{value: val, isSet: true}
}

func (v NullableGetNetworkWirelessSsidSplashSettings200ResponseSplashPrepaidFront) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkWirelessSsidSplashSettings200ResponseSplashPrepaidFront) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

