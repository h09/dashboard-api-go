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

// checks if the UpdateNetworkWirelessSsidRequestLdapCredentials type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkWirelessSsidRequestLdapCredentials{}

// UpdateNetworkWirelessSsidRequestLdapCredentials (Optional) The credentials of the user account to be used by the AP to bind to your LDAP server. The LDAP account should have permissions on all your LDAP servers.
type UpdateNetworkWirelessSsidRequestLdapCredentials struct {
	// The distinguished name of the LDAP user account (example: cn=user,dc=meraki,dc=com).
	DistinguishedName *string `json:"distinguishedName,omitempty"`
	// The password of the LDAP user account.
	Password *string `json:"password,omitempty"`
}

// NewUpdateNetworkWirelessSsidRequestLdapCredentials instantiates a new UpdateNetworkWirelessSsidRequestLdapCredentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkWirelessSsidRequestLdapCredentials() *UpdateNetworkWirelessSsidRequestLdapCredentials {
	this := UpdateNetworkWirelessSsidRequestLdapCredentials{}
	return &this
}

// NewUpdateNetworkWirelessSsidRequestLdapCredentialsWithDefaults instantiates a new UpdateNetworkWirelessSsidRequestLdapCredentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkWirelessSsidRequestLdapCredentialsWithDefaults() *UpdateNetworkWirelessSsidRequestLdapCredentials {
	this := UpdateNetworkWirelessSsidRequestLdapCredentials{}
	return &this
}

// GetDistinguishedName returns the DistinguishedName field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidRequestLdapCredentials) GetDistinguishedName() string {
	if o == nil || IsNil(o.DistinguishedName) {
		var ret string
		return ret
	}
	return *o.DistinguishedName
}

// GetDistinguishedNameOk returns a tuple with the DistinguishedName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidRequestLdapCredentials) GetDistinguishedNameOk() (*string, bool) {
	if o == nil || IsNil(o.DistinguishedName) {
		return nil, false
	}
	return o.DistinguishedName, true
}

// HasDistinguishedName returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidRequestLdapCredentials) HasDistinguishedName() bool {
	if o != nil && !IsNil(o.DistinguishedName) {
		return true
	}

	return false
}

// SetDistinguishedName gets a reference to the given string and assigns it to the DistinguishedName field.
func (o *UpdateNetworkWirelessSsidRequestLdapCredentials) SetDistinguishedName(v string) {
	o.DistinguishedName = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidRequestLdapCredentials) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidRequestLdapCredentials) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidRequestLdapCredentials) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *UpdateNetworkWirelessSsidRequestLdapCredentials) SetPassword(v string) {
	o.Password = &v
}

func (o UpdateNetworkWirelessSsidRequestLdapCredentials) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkWirelessSsidRequestLdapCredentials) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DistinguishedName) {
		toSerialize["distinguishedName"] = o.DistinguishedName
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	return toSerialize, nil
}

type NullableUpdateNetworkWirelessSsidRequestLdapCredentials struct {
	value *UpdateNetworkWirelessSsidRequestLdapCredentials
	isSet bool
}

func (v NullableUpdateNetworkWirelessSsidRequestLdapCredentials) Get() *UpdateNetworkWirelessSsidRequestLdapCredentials {
	return v.value
}

func (v *NullableUpdateNetworkWirelessSsidRequestLdapCredentials) Set(val *UpdateNetworkWirelessSsidRequestLdapCredentials) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkWirelessSsidRequestLdapCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkWirelessSsidRequestLdapCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkWirelessSsidRequestLdapCredentials(val *UpdateNetworkWirelessSsidRequestLdapCredentials) *NullableUpdateNetworkWirelessSsidRequestLdapCredentials {
	return &NullableUpdateNetworkWirelessSsidRequestLdapCredentials{value: val, isSet: true}
}

func (v NullableUpdateNetworkWirelessSsidRequestLdapCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkWirelessSsidRequestLdapCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


