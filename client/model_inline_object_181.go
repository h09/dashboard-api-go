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

// InlineObject181 struct for InlineObject181
type InlineObject181 struct {
	// Sets a list of specific SNORT signatures to allow
	AllowedRules []OrganizationsOrganizationIdApplianceSecurityIntrusionAllowedRules `json:"allowedRules"`
}

// NewInlineObject181 instantiates a new InlineObject181 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject181(allowedRules []OrganizationsOrganizationIdApplianceSecurityIntrusionAllowedRules) *InlineObject181 {
	this := InlineObject181{}
	this.AllowedRules = allowedRules
	return &this
}

// NewInlineObject181WithDefaults instantiates a new InlineObject181 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject181WithDefaults() *InlineObject181 {
	this := InlineObject181{}
	return &this
}

// GetAllowedRules returns the AllowedRules field value
func (o *InlineObject181) GetAllowedRules() []OrganizationsOrganizationIdApplianceSecurityIntrusionAllowedRules {
	if o == nil {
		var ret []OrganizationsOrganizationIdApplianceSecurityIntrusionAllowedRules
		return ret
	}

	return o.AllowedRules
}

// GetAllowedRulesOk returns a tuple with the AllowedRules field value
// and a boolean to check if the value has been set.
func (o *InlineObject181) GetAllowedRulesOk() ([]OrganizationsOrganizationIdApplianceSecurityIntrusionAllowedRules, bool) {
	if o == nil {
    return nil, false
	}
	return o.AllowedRules, true
}

// SetAllowedRules sets field value
func (o *InlineObject181) SetAllowedRules(v []OrganizationsOrganizationIdApplianceSecurityIntrusionAllowedRules) {
	o.AllowedRules = v
}

func (o InlineObject181) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["allowedRules"] = o.AllowedRules
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject181 struct {
	value *InlineObject181
	isSet bool
}

func (v NullableInlineObject181) Get() *InlineObject181 {
	return v.value
}

func (v *NullableInlineObject181) Set(val *InlineObject181) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject181) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject181) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject181(val *InlineObject181) *NullableInlineObject181 {
	return &NullableInlineObject181{value: val, isSet: true}
}

func (v NullableInlineObject181) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject181) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


