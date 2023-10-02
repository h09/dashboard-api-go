/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 06 September, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.37.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the CreateOrganizationCameraRoleRequestAppliedOrgWideInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOrganizationCameraRoleRequestAppliedOrgWideInner{}

// CreateOrganizationCameraRoleRequestAppliedOrgWideInner struct for CreateOrganizationCameraRoleRequestAppliedOrgWideInner
type CreateOrganizationCameraRoleRequestAppliedOrgWideInner struct {
	// Permission scope id
	PermissionScopeId string `json:"permissionScopeId"`
}

// NewCreateOrganizationCameraRoleRequestAppliedOrgWideInner instantiates a new CreateOrganizationCameraRoleRequestAppliedOrgWideInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrganizationCameraRoleRequestAppliedOrgWideInner(permissionScopeId string) *CreateOrganizationCameraRoleRequestAppliedOrgWideInner {
	this := CreateOrganizationCameraRoleRequestAppliedOrgWideInner{}
	this.PermissionScopeId = permissionScopeId
	return &this
}

// NewCreateOrganizationCameraRoleRequestAppliedOrgWideInnerWithDefaults instantiates a new CreateOrganizationCameraRoleRequestAppliedOrgWideInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrganizationCameraRoleRequestAppliedOrgWideInnerWithDefaults() *CreateOrganizationCameraRoleRequestAppliedOrgWideInner {
	this := CreateOrganizationCameraRoleRequestAppliedOrgWideInner{}
	return &this
}

// GetPermissionScopeId returns the PermissionScopeId field value
func (o *CreateOrganizationCameraRoleRequestAppliedOrgWideInner) GetPermissionScopeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PermissionScopeId
}

// GetPermissionScopeIdOk returns a tuple with the PermissionScopeId field value
// and a boolean to check if the value has been set.
func (o *CreateOrganizationCameraRoleRequestAppliedOrgWideInner) GetPermissionScopeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PermissionScopeId, true
}

// SetPermissionScopeId sets field value
func (o *CreateOrganizationCameraRoleRequestAppliedOrgWideInner) SetPermissionScopeId(v string) {
	o.PermissionScopeId = v
}

func (o CreateOrganizationCameraRoleRequestAppliedOrgWideInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateOrganizationCameraRoleRequestAppliedOrgWideInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["permissionScopeId"] = o.PermissionScopeId
	return toSerialize, nil
}

type NullableCreateOrganizationCameraRoleRequestAppliedOrgWideInner struct {
	value *CreateOrganizationCameraRoleRequestAppliedOrgWideInner
	isSet bool
}

func (v NullableCreateOrganizationCameraRoleRequestAppliedOrgWideInner) Get() *CreateOrganizationCameraRoleRequestAppliedOrgWideInner {
	return v.value
}

func (v *NullableCreateOrganizationCameraRoleRequestAppliedOrgWideInner) Set(val *CreateOrganizationCameraRoleRequestAppliedOrgWideInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrganizationCameraRoleRequestAppliedOrgWideInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrganizationCameraRoleRequestAppliedOrgWideInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrganizationCameraRoleRequestAppliedOrgWideInner(val *CreateOrganizationCameraRoleRequestAppliedOrgWideInner) *NullableCreateOrganizationCameraRoleRequestAppliedOrgWideInner {
	return &NullableCreateOrganizationCameraRoleRequestAppliedOrgWideInner{value: val, isSet: true}
}

func (v NullableCreateOrganizationCameraRoleRequestAppliedOrgWideInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrganizationCameraRoleRequestAppliedOrgWideInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

