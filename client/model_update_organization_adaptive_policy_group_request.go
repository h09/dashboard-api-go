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

// checks if the UpdateOrganizationAdaptivePolicyGroupRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateOrganizationAdaptivePolicyGroupRequest{}

// UpdateOrganizationAdaptivePolicyGroupRequest struct for UpdateOrganizationAdaptivePolicyGroupRequest
type UpdateOrganizationAdaptivePolicyGroupRequest struct {
	// Name of the group
	Name *string `json:"name,omitempty"`
	// SGT value of the group
	Sgt *int32 `json:"sgt,omitempty"`
	// Description of the group
	Description *string `json:"description,omitempty"`
	// The policy objects that belong to this group; traffic from addresses specified by these policy objects will be tagged with this group's SGT value if no other tagging scheme is being used (each requires one unique attribute)
	PolicyObjects []CreateOrganizationAdaptivePolicyGroupRequestPolicyObjectsInner `json:"policyObjects,omitempty"`
}

// NewUpdateOrganizationAdaptivePolicyGroupRequest instantiates a new UpdateOrganizationAdaptivePolicyGroupRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateOrganizationAdaptivePolicyGroupRequest() *UpdateOrganizationAdaptivePolicyGroupRequest {
	this := UpdateOrganizationAdaptivePolicyGroupRequest{}
	return &this
}

// NewUpdateOrganizationAdaptivePolicyGroupRequestWithDefaults instantiates a new UpdateOrganizationAdaptivePolicyGroupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateOrganizationAdaptivePolicyGroupRequestWithDefaults() *UpdateOrganizationAdaptivePolicyGroupRequest {
	this := UpdateOrganizationAdaptivePolicyGroupRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateOrganizationAdaptivePolicyGroupRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationAdaptivePolicyGroupRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateOrganizationAdaptivePolicyGroupRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateOrganizationAdaptivePolicyGroupRequest) SetName(v string) {
	o.Name = &v
}

// GetSgt returns the Sgt field value if set, zero value otherwise.
func (o *UpdateOrganizationAdaptivePolicyGroupRequest) GetSgt() int32 {
	if o == nil || IsNil(o.Sgt) {
		var ret int32
		return ret
	}
	return *o.Sgt
}

// GetSgtOk returns a tuple with the Sgt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationAdaptivePolicyGroupRequest) GetSgtOk() (*int32, bool) {
	if o == nil || IsNil(o.Sgt) {
		return nil, false
	}
	return o.Sgt, true
}

// HasSgt returns a boolean if a field has been set.
func (o *UpdateOrganizationAdaptivePolicyGroupRequest) HasSgt() bool {
	if o != nil && !IsNil(o.Sgt) {
		return true
	}

	return false
}

// SetSgt gets a reference to the given int32 and assigns it to the Sgt field.
func (o *UpdateOrganizationAdaptivePolicyGroupRequest) SetSgt(v int32) {
	o.Sgt = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateOrganizationAdaptivePolicyGroupRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationAdaptivePolicyGroupRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateOrganizationAdaptivePolicyGroupRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateOrganizationAdaptivePolicyGroupRequest) SetDescription(v string) {
	o.Description = &v
}

// GetPolicyObjects returns the PolicyObjects field value if set, zero value otherwise.
func (o *UpdateOrganizationAdaptivePolicyGroupRequest) GetPolicyObjects() []CreateOrganizationAdaptivePolicyGroupRequestPolicyObjectsInner {
	if o == nil || IsNil(o.PolicyObjects) {
		var ret []CreateOrganizationAdaptivePolicyGroupRequestPolicyObjectsInner
		return ret
	}
	return o.PolicyObjects
}

// GetPolicyObjectsOk returns a tuple with the PolicyObjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationAdaptivePolicyGroupRequest) GetPolicyObjectsOk() ([]CreateOrganizationAdaptivePolicyGroupRequestPolicyObjectsInner, bool) {
	if o == nil || IsNil(o.PolicyObjects) {
		return nil, false
	}
	return o.PolicyObjects, true
}

// HasPolicyObjects returns a boolean if a field has been set.
func (o *UpdateOrganizationAdaptivePolicyGroupRequest) HasPolicyObjects() bool {
	if o != nil && !IsNil(o.PolicyObjects) {
		return true
	}

	return false
}

// SetPolicyObjects gets a reference to the given []CreateOrganizationAdaptivePolicyGroupRequestPolicyObjectsInner and assigns it to the PolicyObjects field.
func (o *UpdateOrganizationAdaptivePolicyGroupRequest) SetPolicyObjects(v []CreateOrganizationAdaptivePolicyGroupRequestPolicyObjectsInner) {
	o.PolicyObjects = v
}

func (o UpdateOrganizationAdaptivePolicyGroupRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateOrganizationAdaptivePolicyGroupRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Sgt) {
		toSerialize["sgt"] = o.Sgt
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.PolicyObjects) {
		toSerialize["policyObjects"] = o.PolicyObjects
	}
	return toSerialize, nil
}

type NullableUpdateOrganizationAdaptivePolicyGroupRequest struct {
	value *UpdateOrganizationAdaptivePolicyGroupRequest
	isSet bool
}

func (v NullableUpdateOrganizationAdaptivePolicyGroupRequest) Get() *UpdateOrganizationAdaptivePolicyGroupRequest {
	return v.value
}

func (v *NullableUpdateOrganizationAdaptivePolicyGroupRequest) Set(val *UpdateOrganizationAdaptivePolicyGroupRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateOrganizationAdaptivePolicyGroupRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateOrganizationAdaptivePolicyGroupRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateOrganizationAdaptivePolicyGroupRequest(val *UpdateOrganizationAdaptivePolicyGroupRequest) *NullableUpdateOrganizationAdaptivePolicyGroupRequest {
	return &NullableUpdateOrganizationAdaptivePolicyGroupRequest{value: val, isSet: true}
}

func (v NullableUpdateOrganizationAdaptivePolicyGroupRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateOrganizationAdaptivePolicyGroupRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

