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

// checks if the GetOrganizationSmVppAccounts200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationSmVppAccounts200ResponseInner{}

// GetOrganizationSmVppAccounts200ResponseInner struct for GetOrganizationSmVppAccounts200ResponseInner
type GetOrganizationSmVppAccounts200ResponseInner struct {
	// The id of the VPP Account
	Id *string `json:"id,omitempty"`
	// The VPP Account's Service Token
	VppServiceToken *string `json:"vppServiceToken,omitempty"`
}

// NewGetOrganizationSmVppAccounts200ResponseInner instantiates a new GetOrganizationSmVppAccounts200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationSmVppAccounts200ResponseInner() *GetOrganizationSmVppAccounts200ResponseInner {
	this := GetOrganizationSmVppAccounts200ResponseInner{}
	return &this
}

// NewGetOrganizationSmVppAccounts200ResponseInnerWithDefaults instantiates a new GetOrganizationSmVppAccounts200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationSmVppAccounts200ResponseInnerWithDefaults() *GetOrganizationSmVppAccounts200ResponseInner {
	this := GetOrganizationSmVppAccounts200ResponseInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetOrganizationSmVppAccounts200ResponseInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSmVppAccounts200ResponseInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetOrganizationSmVppAccounts200ResponseInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetOrganizationSmVppAccounts200ResponseInner) SetId(v string) {
	o.Id = &v
}

// GetVppServiceToken returns the VppServiceToken field value if set, zero value otherwise.
func (o *GetOrganizationSmVppAccounts200ResponseInner) GetVppServiceToken() string {
	if o == nil || IsNil(o.VppServiceToken) {
		var ret string
		return ret
	}
	return *o.VppServiceToken
}

// GetVppServiceTokenOk returns a tuple with the VppServiceToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSmVppAccounts200ResponseInner) GetVppServiceTokenOk() (*string, bool) {
	if o == nil || IsNil(o.VppServiceToken) {
		return nil, false
	}
	return o.VppServiceToken, true
}

// HasVppServiceToken returns a boolean if a field has been set.
func (o *GetOrganizationSmVppAccounts200ResponseInner) HasVppServiceToken() bool {
	if o != nil && !IsNil(o.VppServiceToken) {
		return true
	}

	return false
}

// SetVppServiceToken gets a reference to the given string and assigns it to the VppServiceToken field.
func (o *GetOrganizationSmVppAccounts200ResponseInner) SetVppServiceToken(v string) {
	o.VppServiceToken = &v
}

func (o GetOrganizationSmVppAccounts200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationSmVppAccounts200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.VppServiceToken) {
		toSerialize["vppServiceToken"] = o.VppServiceToken
	}
	return toSerialize, nil
}

type NullableGetOrganizationSmVppAccounts200ResponseInner struct {
	value *GetOrganizationSmVppAccounts200ResponseInner
	isSet bool
}

func (v NullableGetOrganizationSmVppAccounts200ResponseInner) Get() *GetOrganizationSmVppAccounts200ResponseInner {
	return v.value
}

func (v *NullableGetOrganizationSmVppAccounts200ResponseInner) Set(val *GetOrganizationSmVppAccounts200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationSmVppAccounts200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationSmVppAccounts200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationSmVppAccounts200ResponseInner(val *GetOrganizationSmVppAccounts200ResponseInner) *NullableGetOrganizationSmVppAccounts200ResponseInner {
	return &NullableGetOrganizationSmVppAccounts200ResponseInner{value: val, isSet: true}
}

func (v NullableGetOrganizationSmVppAccounts200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationSmVppAccounts200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

