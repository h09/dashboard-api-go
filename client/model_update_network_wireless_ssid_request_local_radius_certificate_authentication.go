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

// checks if the UpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthentication type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthentication{}

// UpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthentication The current setting for certificate verification.
type UpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthentication struct {
	// Whether or not to use EAP-TLS certificate-based authentication to validate wireless clients.
	Enabled *bool `json:"enabled,omitempty"`
	// Whether or not to verify the certificate with LDAP.
	UseLdap *bool `json:"useLdap,omitempty"`
	// Whether or not to verify the certificate with OCSP.
	UseOcsp *bool `json:"useOcsp,omitempty"`
	// (Optional) The URL of the OCSP responder to verify client certificate status.
	OcspResponderUrl *string `json:"ocspResponderUrl,omitempty"`
	ClientRootCaCertificate *UpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthenticationClientRootCaCertificate `json:"clientRootCaCertificate,omitempty"`
}

// NewUpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthentication instantiates a new UpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthentication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthentication() *UpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthentication {
	this := UpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthentication{}
	return &this
}

// NewUpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthenticationWithDefaults instantiates a new UpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthentication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthenticationWithDefaults() *UpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthentication {
	this := UpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthentication{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthentication) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthentication) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthentication) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *UpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthentication) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetUseLdap returns the UseLdap field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthentication) GetUseLdap() bool {
	if o == nil || IsNil(o.UseLdap) {
		var ret bool
		return ret
	}
	return *o.UseLdap
}

// GetUseLdapOk returns a tuple with the UseLdap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthentication) GetUseLdapOk() (*bool, bool) {
	if o == nil || IsNil(o.UseLdap) {
		return nil, false
	}
	return o.UseLdap, true
}

// HasUseLdap returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthentication) HasUseLdap() bool {
	if o != nil && !IsNil(o.UseLdap) {
		return true
	}

	return false
}

// SetUseLdap gets a reference to the given bool and assigns it to the UseLdap field.
func (o *UpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthentication) SetUseLdap(v bool) {
	o.UseLdap = &v
}

// GetUseOcsp returns the UseOcsp field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthentication) GetUseOcsp() bool {
	if o == nil || IsNil(o.UseOcsp) {
		var ret bool
		return ret
	}
	return *o.UseOcsp
}

// GetUseOcspOk returns a tuple with the UseOcsp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthentication) GetUseOcspOk() (*bool, bool) {
	if o == nil || IsNil(o.UseOcsp) {
		return nil, false
	}
	return o.UseOcsp, true
}

// HasUseOcsp returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthentication) HasUseOcsp() bool {
	if o != nil && !IsNil(o.UseOcsp) {
		return true
	}

	return false
}

// SetUseOcsp gets a reference to the given bool and assigns it to the UseOcsp field.
func (o *UpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthentication) SetUseOcsp(v bool) {
	o.UseOcsp = &v
}

// GetOcspResponderUrl returns the OcspResponderUrl field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthentication) GetOcspResponderUrl() string {
	if o == nil || IsNil(o.OcspResponderUrl) {
		var ret string
		return ret
	}
	return *o.OcspResponderUrl
}

// GetOcspResponderUrlOk returns a tuple with the OcspResponderUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthentication) GetOcspResponderUrlOk() (*string, bool) {
	if o == nil || IsNil(o.OcspResponderUrl) {
		return nil, false
	}
	return o.OcspResponderUrl, true
}

// HasOcspResponderUrl returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthentication) HasOcspResponderUrl() bool {
	if o != nil && !IsNil(o.OcspResponderUrl) {
		return true
	}

	return false
}

// SetOcspResponderUrl gets a reference to the given string and assigns it to the OcspResponderUrl field.
func (o *UpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthentication) SetOcspResponderUrl(v string) {
	o.OcspResponderUrl = &v
}

// GetClientRootCaCertificate returns the ClientRootCaCertificate field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthentication) GetClientRootCaCertificate() UpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthenticationClientRootCaCertificate {
	if o == nil || IsNil(o.ClientRootCaCertificate) {
		var ret UpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthenticationClientRootCaCertificate
		return ret
	}
	return *o.ClientRootCaCertificate
}

// GetClientRootCaCertificateOk returns a tuple with the ClientRootCaCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthentication) GetClientRootCaCertificateOk() (*UpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthenticationClientRootCaCertificate, bool) {
	if o == nil || IsNil(o.ClientRootCaCertificate) {
		return nil, false
	}
	return o.ClientRootCaCertificate, true
}

// HasClientRootCaCertificate returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthentication) HasClientRootCaCertificate() bool {
	if o != nil && !IsNil(o.ClientRootCaCertificate) {
		return true
	}

	return false
}

// SetClientRootCaCertificate gets a reference to the given UpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthenticationClientRootCaCertificate and assigns it to the ClientRootCaCertificate field.
func (o *UpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthentication) SetClientRootCaCertificate(v UpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthenticationClientRootCaCertificate) {
	o.ClientRootCaCertificate = &v
}

func (o UpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthentication) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthentication) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.UseLdap) {
		toSerialize["useLdap"] = o.UseLdap
	}
	if !IsNil(o.UseOcsp) {
		toSerialize["useOcsp"] = o.UseOcsp
	}
	if !IsNil(o.OcspResponderUrl) {
		toSerialize["ocspResponderUrl"] = o.OcspResponderUrl
	}
	if !IsNil(o.ClientRootCaCertificate) {
		toSerialize["clientRootCaCertificate"] = o.ClientRootCaCertificate
	}
	return toSerialize, nil
}

type NullableUpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthentication struct {
	value *UpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthentication
	isSet bool
}

func (v NullableUpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthentication) Get() *UpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthentication {
	return v.value
}

func (v *NullableUpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthentication) Set(val *UpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthentication) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthentication) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthentication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthentication(val *UpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthentication) *NullableUpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthentication {
	return &NullableUpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthentication{value: val, isSet: true}
}

func (v NullableUpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthentication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkWirelessSsidRequestLocalRadiusCertificateAuthentication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


