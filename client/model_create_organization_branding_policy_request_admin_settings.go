/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 02 November, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.27.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// CreateOrganizationBrandingPolicyRequestAdminSettings Settings for describing which kinds of admins this policy applies to.
type CreateOrganizationBrandingPolicyRequestAdminSettings struct {
	// Which kinds of admins this policy applies to. Can be one of 'All organization admins', 'All enterprise admins', 'All network admins', 'All admins of networks...', 'All admins of networks tagged...', 'Specific admins...', 'All admins' or 'All SAML admins'.
	AppliesTo *string `json:"appliesTo,omitempty"`
	//     If 'appliesTo' is set to one of 'Specific admins...', 'All admins of networks...' or 'All admins of networks tagged...', then you must specify this 'values' property to provide the set of     entities to apply the branding policy to. For 'Specific admins...', specify an array of admin IDs. For 'All admins of     networks...', specify an array of network IDs and/or configuration template IDs. For 'All admins of networks tagged...',     specify an array of tag names. 
	Values []string `json:"values,omitempty"`
}

// NewCreateOrganizationBrandingPolicyRequestAdminSettings instantiates a new CreateOrganizationBrandingPolicyRequestAdminSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrganizationBrandingPolicyRequestAdminSettings() *CreateOrganizationBrandingPolicyRequestAdminSettings {
	this := CreateOrganizationBrandingPolicyRequestAdminSettings{}
	return &this
}

// NewCreateOrganizationBrandingPolicyRequestAdminSettingsWithDefaults instantiates a new CreateOrganizationBrandingPolicyRequestAdminSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrganizationBrandingPolicyRequestAdminSettingsWithDefaults() *CreateOrganizationBrandingPolicyRequestAdminSettings {
	this := CreateOrganizationBrandingPolicyRequestAdminSettings{}
	return &this
}

// GetAppliesTo returns the AppliesTo field value if set, zero value otherwise.
func (o *CreateOrganizationBrandingPolicyRequestAdminSettings) GetAppliesTo() string {
	if o == nil || isNil(o.AppliesTo) {
		var ret string
		return ret
	}
	return *o.AppliesTo
}

// GetAppliesToOk returns a tuple with the AppliesTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationBrandingPolicyRequestAdminSettings) GetAppliesToOk() (*string, bool) {
	if o == nil || isNil(o.AppliesTo) {
    return nil, false
	}
	return o.AppliesTo, true
}

// HasAppliesTo returns a boolean if a field has been set.
func (o *CreateOrganizationBrandingPolicyRequestAdminSettings) HasAppliesTo() bool {
	if o != nil && !isNil(o.AppliesTo) {
		return true
	}

	return false
}

// SetAppliesTo gets a reference to the given string and assigns it to the AppliesTo field.
func (o *CreateOrganizationBrandingPolicyRequestAdminSettings) SetAppliesTo(v string) {
	o.AppliesTo = &v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *CreateOrganizationBrandingPolicyRequestAdminSettings) GetValues() []string {
	if o == nil || isNil(o.Values) {
		var ret []string
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationBrandingPolicyRequestAdminSettings) GetValuesOk() ([]string, bool) {
	if o == nil || isNil(o.Values) {
    return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *CreateOrganizationBrandingPolicyRequestAdminSettings) HasValues() bool {
	if o != nil && !isNil(o.Values) {
		return true
	}

	return false
}

// SetValues gets a reference to the given []string and assigns it to the Values field.
func (o *CreateOrganizationBrandingPolicyRequestAdminSettings) SetValues(v []string) {
	o.Values = v
}

func (o CreateOrganizationBrandingPolicyRequestAdminSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AppliesTo) {
		toSerialize["appliesTo"] = o.AppliesTo
	}
	if !isNil(o.Values) {
		toSerialize["values"] = o.Values
	}
	return json.Marshal(toSerialize)
}

type NullableCreateOrganizationBrandingPolicyRequestAdminSettings struct {
	value *CreateOrganizationBrandingPolicyRequestAdminSettings
	isSet bool
}

func (v NullableCreateOrganizationBrandingPolicyRequestAdminSettings) Get() *CreateOrganizationBrandingPolicyRequestAdminSettings {
	return v.value
}

func (v *NullableCreateOrganizationBrandingPolicyRequestAdminSettings) Set(val *CreateOrganizationBrandingPolicyRequestAdminSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrganizationBrandingPolicyRequestAdminSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrganizationBrandingPolicyRequestAdminSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrganizationBrandingPolicyRequestAdminSettings(val *CreateOrganizationBrandingPolicyRequestAdminSettings) *NullableCreateOrganizationBrandingPolicyRequestAdminSettings {
	return &NullableCreateOrganizationBrandingPolicyRequestAdminSettings{value: val, isSet: true}
}

func (v NullableCreateOrganizationBrandingPolicyRequestAdminSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrganizationBrandingPolicyRequestAdminSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

