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

// checks if the GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInnerDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInnerDetails{}

// GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInnerDetails Details about the status changes
type GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInnerDetails struct {
	// Details about the old status
	Old []GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInnerDetailsOldInner `json:"old,omitempty"`
	// Details about the new status
	New []GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInnerDetailsOldInner `json:"new,omitempty"`
}

// NewGetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInnerDetails instantiates a new GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInnerDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInnerDetails() *GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInnerDetails {
	this := GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInnerDetails{}
	return &this
}

// NewGetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInnerDetailsWithDefaults instantiates a new GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInnerDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInnerDetailsWithDefaults() *GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInnerDetails {
	this := GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInnerDetails{}
	return &this
}

// GetOld returns the Old field value if set, zero value otherwise.
func (o *GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInnerDetails) GetOld() []GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInnerDetailsOldInner {
	if o == nil || IsNil(o.Old) {
		var ret []GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInnerDetailsOldInner
		return ret
	}
	return o.Old
}

// GetOldOk returns a tuple with the Old field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInnerDetails) GetOldOk() ([]GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInnerDetailsOldInner, bool) {
	if o == nil || IsNil(o.Old) {
		return nil, false
	}
	return o.Old, true
}

// HasOld returns a boolean if a field has been set.
func (o *GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInnerDetails) HasOld() bool {
	if o != nil && !IsNil(o.Old) {
		return true
	}

	return false
}

// SetOld gets a reference to the given []GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInnerDetailsOldInner and assigns it to the Old field.
func (o *GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInnerDetails) SetOld(v []GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInnerDetailsOldInner) {
	o.Old = v
}

// GetNew returns the New field value if set, zero value otherwise.
func (o *GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInnerDetails) GetNew() []GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInnerDetailsOldInner {
	if o == nil || IsNil(o.New) {
		var ret []GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInnerDetailsOldInner
		return ret
	}
	return o.New
}

// GetNewOk returns a tuple with the New field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInnerDetails) GetNewOk() ([]GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInnerDetailsOldInner, bool) {
	if o == nil || IsNil(o.New) {
		return nil, false
	}
	return o.New, true
}

// HasNew returns a boolean if a field has been set.
func (o *GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInnerDetails) HasNew() bool {
	if o != nil && !IsNil(o.New) {
		return true
	}

	return false
}

// SetNew gets a reference to the given []GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInnerDetailsOldInner and assigns it to the New field.
func (o *GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInnerDetails) SetNew(v []GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInnerDetailsOldInner) {
	o.New = v
}

func (o GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInnerDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInnerDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Old) {
		toSerialize["old"] = o.Old
	}
	if !IsNil(o.New) {
		toSerialize["new"] = o.New
	}
	return toSerialize, nil
}

type NullableGetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInnerDetails struct {
	value *GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInnerDetails
	isSet bool
}

func (v NullableGetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInnerDetails) Get() *GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInnerDetails {
	return v.value
}

func (v *NullableGetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInnerDetails) Set(val *GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInnerDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInnerDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInnerDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInnerDetails(val *GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInnerDetails) *NullableGetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInnerDetails {
	return &NullableGetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInnerDetails{value: val, isSet: true}
}

func (v NullableGetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInnerDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInnerDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

