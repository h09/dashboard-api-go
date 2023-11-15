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

// checks if the GetOrganizationSamlRoles200ResponseInnerTagsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationSamlRoles200ResponseInnerTagsInner{}

// GetOrganizationSamlRoles200ResponseInnerTagsInner struct for GetOrganizationSamlRoles200ResponseInnerTagsInner
type GetOrganizationSamlRoles200ResponseInnerTagsInner struct {
	// The name of the tag
	Tag *string `json:"tag,omitempty"`
	// The privilege of the SAML administrator on the tag
	Access *string `json:"access,omitempty"`
}

// NewGetOrganizationSamlRoles200ResponseInnerTagsInner instantiates a new GetOrganizationSamlRoles200ResponseInnerTagsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationSamlRoles200ResponseInnerTagsInner() *GetOrganizationSamlRoles200ResponseInnerTagsInner {
	this := GetOrganizationSamlRoles200ResponseInnerTagsInner{}
	return &this
}

// NewGetOrganizationSamlRoles200ResponseInnerTagsInnerWithDefaults instantiates a new GetOrganizationSamlRoles200ResponseInnerTagsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationSamlRoles200ResponseInnerTagsInnerWithDefaults() *GetOrganizationSamlRoles200ResponseInnerTagsInner {
	this := GetOrganizationSamlRoles200ResponseInnerTagsInner{}
	return &this
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *GetOrganizationSamlRoles200ResponseInnerTagsInner) GetTag() string {
	if o == nil || IsNil(o.Tag) {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSamlRoles200ResponseInnerTagsInner) GetTagOk() (*string, bool) {
	if o == nil || IsNil(o.Tag) {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *GetOrganizationSamlRoles200ResponseInnerTagsInner) HasTag() bool {
	if o != nil && !IsNil(o.Tag) {
		return true
	}

	return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
func (o *GetOrganizationSamlRoles200ResponseInnerTagsInner) SetTag(v string) {
	o.Tag = &v
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *GetOrganizationSamlRoles200ResponseInnerTagsInner) GetAccess() string {
	if o == nil || IsNil(o.Access) {
		var ret string
		return ret
	}
	return *o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSamlRoles200ResponseInnerTagsInner) GetAccessOk() (*string, bool) {
	if o == nil || IsNil(o.Access) {
		return nil, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *GetOrganizationSamlRoles200ResponseInnerTagsInner) HasAccess() bool {
	if o != nil && !IsNil(o.Access) {
		return true
	}

	return false
}

// SetAccess gets a reference to the given string and assigns it to the Access field.
func (o *GetOrganizationSamlRoles200ResponseInnerTagsInner) SetAccess(v string) {
	o.Access = &v
}

func (o GetOrganizationSamlRoles200ResponseInnerTagsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationSamlRoles200ResponseInnerTagsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Tag) {
		toSerialize["tag"] = o.Tag
	}
	if !IsNil(o.Access) {
		toSerialize["access"] = o.Access
	}
	return toSerialize, nil
}

type NullableGetOrganizationSamlRoles200ResponseInnerTagsInner struct {
	value *GetOrganizationSamlRoles200ResponseInnerTagsInner
	isSet bool
}

func (v NullableGetOrganizationSamlRoles200ResponseInnerTagsInner) Get() *GetOrganizationSamlRoles200ResponseInnerTagsInner {
	return v.value
}

func (v *NullableGetOrganizationSamlRoles200ResponseInnerTagsInner) Set(val *GetOrganizationSamlRoles200ResponseInnerTagsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationSamlRoles200ResponseInnerTagsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationSamlRoles200ResponseInnerTagsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationSamlRoles200ResponseInnerTagsInner(val *GetOrganizationSamlRoles200ResponseInnerTagsInner) *NullableGetOrganizationSamlRoles200ResponseInnerTagsInner {
	return &NullableGetOrganizationSamlRoles200ResponseInnerTagsInner{value: val, isSet: true}
}

func (v NullableGetOrganizationSamlRoles200ResponseInnerTagsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationSamlRoles200ResponseInnerTagsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


