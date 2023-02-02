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

// OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeFromVersion The initial version of the device
type OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeFromVersion struct {
	// ID of the initial firmware version
	Id *string `json:"id,omitempty"`
	// Firmware version short name
	ShortName *string `json:"shortName,omitempty"`
	// Release type of the firmware version
	ReleaseType *string `json:"releaseType,omitempty"`
	// Release date of the firmware version
	ReleaseDate *string `json:"releaseDate,omitempty"`
}

// NewOrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeFromVersion instantiates a new OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeFromVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeFromVersion() *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeFromVersion {
	this := OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeFromVersion{}
	return &this
}

// NewOrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeFromVersionWithDefaults instantiates a new OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeFromVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeFromVersionWithDefaults() *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeFromVersion {
	this := OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeFromVersion{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeFromVersion) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeFromVersion) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeFromVersion) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeFromVersion) SetId(v string) {
	o.Id = &v
}

// GetShortName returns the ShortName field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeFromVersion) GetShortName() string {
	if o == nil || isNil(o.ShortName) {
		var ret string
		return ret
	}
	return *o.ShortName
}

// GetShortNameOk returns a tuple with the ShortName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeFromVersion) GetShortNameOk() (*string, bool) {
	if o == nil || isNil(o.ShortName) {
    return nil, false
	}
	return o.ShortName, true
}

// HasShortName returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeFromVersion) HasShortName() bool {
	if o != nil && !isNil(o.ShortName) {
		return true
	}

	return false
}

// SetShortName gets a reference to the given string and assigns it to the ShortName field.
func (o *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeFromVersion) SetShortName(v string) {
	o.ShortName = &v
}

// GetReleaseType returns the ReleaseType field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeFromVersion) GetReleaseType() string {
	if o == nil || isNil(o.ReleaseType) {
		var ret string
		return ret
	}
	return *o.ReleaseType
}

// GetReleaseTypeOk returns a tuple with the ReleaseType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeFromVersion) GetReleaseTypeOk() (*string, bool) {
	if o == nil || isNil(o.ReleaseType) {
    return nil, false
	}
	return o.ReleaseType, true
}

// HasReleaseType returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeFromVersion) HasReleaseType() bool {
	if o != nil && !isNil(o.ReleaseType) {
		return true
	}

	return false
}

// SetReleaseType gets a reference to the given string and assigns it to the ReleaseType field.
func (o *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeFromVersion) SetReleaseType(v string) {
	o.ReleaseType = &v
}

// GetReleaseDate returns the ReleaseDate field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeFromVersion) GetReleaseDate() string {
	if o == nil || isNil(o.ReleaseDate) {
		var ret string
		return ret
	}
	return *o.ReleaseDate
}

// GetReleaseDateOk returns a tuple with the ReleaseDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeFromVersion) GetReleaseDateOk() (*string, bool) {
	if o == nil || isNil(o.ReleaseDate) {
    return nil, false
	}
	return o.ReleaseDate, true
}

// HasReleaseDate returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeFromVersion) HasReleaseDate() bool {
	if o != nil && !isNil(o.ReleaseDate) {
		return true
	}

	return false
}

// SetReleaseDate gets a reference to the given string and assigns it to the ReleaseDate field.
func (o *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeFromVersion) SetReleaseDate(v string) {
	o.ReleaseDate = &v
}

func (o OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeFromVersion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.ShortName) {
		toSerialize["shortName"] = o.ShortName
	}
	if !isNil(o.ReleaseType) {
		toSerialize["releaseType"] = o.ReleaseType
	}
	if !isNil(o.ReleaseDate) {
		toSerialize["releaseDate"] = o.ReleaseDate
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeFromVersion struct {
	value *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeFromVersion
	isSet bool
}

func (v NullableOrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeFromVersion) Get() *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeFromVersion {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeFromVersion) Set(val *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeFromVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeFromVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeFromVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeFromVersion(val *OrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeFromVersion) *NullableOrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeFromVersion {
	return &NullableOrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeFromVersion{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeFromVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdFirmwareUpgradesByDeviceUpgradeFromVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


