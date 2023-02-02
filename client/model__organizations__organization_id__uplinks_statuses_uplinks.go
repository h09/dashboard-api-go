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

// OrganizationsOrganizationIdUplinksStatusesUplinks struct for OrganizationsOrganizationIdUplinksStatusesUplinks
type OrganizationsOrganizationIdUplinksStatusesUplinks struct {
	// Uplink interface
	Interface *string `json:"interface,omitempty"`
	// Uplink status
	Status *string `json:"status,omitempty"`
	// Uplink IP
	Ip *string `json:"ip,omitempty"`
	// Gateway IP
	Gateway *string `json:"gateway,omitempty"`
	// Public IP
	PublicIp *string `json:"publicIp,omitempty"`
	// Primary DNS IP
	PrimaryDns *string `json:"primaryDns,omitempty"`
	// Secondary DNS IP
	SecondaryDns *string `json:"secondaryDns,omitempty"`
	// The way in which the IP is assigned
	IpAssignedBy *string `json:"ipAssignedBy,omitempty"`
	// Network Provider
	Provider *string `json:"provider,omitempty"`
	SignalStat *OrganizationsOrganizationIdCellularGatewayUplinkStatusesSignalStat `json:"signalStat,omitempty"`
	// Connection Type
	ConnectionType *string `json:"connectionType,omitempty"`
	// Access Point Name
	Apn *string `json:"apn,omitempty"`
	// Primary DNS IP
	Dns1 *string `json:"dns1,omitempty"`
	// Secondary DNS IP
	Dns2 *string `json:"dns2,omitempty"`
	// Signal Type
	SignalType *string `json:"signalType,omitempty"`
	// Integrated Circuit Card Identification Number
	Iccid *string `json:"iccid,omitempty"`
}

// NewOrganizationsOrganizationIdUplinksStatusesUplinks instantiates a new OrganizationsOrganizationIdUplinksStatusesUplinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdUplinksStatusesUplinks() *OrganizationsOrganizationIdUplinksStatusesUplinks {
	this := OrganizationsOrganizationIdUplinksStatusesUplinks{}
	return &this
}

// NewOrganizationsOrganizationIdUplinksStatusesUplinksWithDefaults instantiates a new OrganizationsOrganizationIdUplinksStatusesUplinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdUplinksStatusesUplinksWithDefaults() *OrganizationsOrganizationIdUplinksStatusesUplinks {
	this := OrganizationsOrganizationIdUplinksStatusesUplinks{}
	return &this
}

// GetInterface returns the Interface field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdUplinksStatusesUplinks) GetInterface() string {
	if o == nil || isNil(o.Interface) {
		var ret string
		return ret
	}
	return *o.Interface
}

// GetInterfaceOk returns a tuple with the Interface field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdUplinksStatusesUplinks) GetInterfaceOk() (*string, bool) {
	if o == nil || isNil(o.Interface) {
    return nil, false
	}
	return o.Interface, true
}

// HasInterface returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdUplinksStatusesUplinks) HasInterface() bool {
	if o != nil && !isNil(o.Interface) {
		return true
	}

	return false
}

// SetInterface gets a reference to the given string and assigns it to the Interface field.
func (o *OrganizationsOrganizationIdUplinksStatusesUplinks) SetInterface(v string) {
	o.Interface = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdUplinksStatusesUplinks) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdUplinksStatusesUplinks) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdUplinksStatusesUplinks) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *OrganizationsOrganizationIdUplinksStatusesUplinks) SetStatus(v string) {
	o.Status = &v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdUplinksStatusesUplinks) GetIp() string {
	if o == nil || isNil(o.Ip) {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdUplinksStatusesUplinks) GetIpOk() (*string, bool) {
	if o == nil || isNil(o.Ip) {
    return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdUplinksStatusesUplinks) HasIp() bool {
	if o != nil && !isNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *OrganizationsOrganizationIdUplinksStatusesUplinks) SetIp(v string) {
	o.Ip = &v
}

// GetGateway returns the Gateway field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdUplinksStatusesUplinks) GetGateway() string {
	if o == nil || isNil(o.Gateway) {
		var ret string
		return ret
	}
	return *o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdUplinksStatusesUplinks) GetGatewayOk() (*string, bool) {
	if o == nil || isNil(o.Gateway) {
    return nil, false
	}
	return o.Gateway, true
}

// HasGateway returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdUplinksStatusesUplinks) HasGateway() bool {
	if o != nil && !isNil(o.Gateway) {
		return true
	}

	return false
}

// SetGateway gets a reference to the given string and assigns it to the Gateway field.
func (o *OrganizationsOrganizationIdUplinksStatusesUplinks) SetGateway(v string) {
	o.Gateway = &v
}

// GetPublicIp returns the PublicIp field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdUplinksStatusesUplinks) GetPublicIp() string {
	if o == nil || isNil(o.PublicIp) {
		var ret string
		return ret
	}
	return *o.PublicIp
}

// GetPublicIpOk returns a tuple with the PublicIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdUplinksStatusesUplinks) GetPublicIpOk() (*string, bool) {
	if o == nil || isNil(o.PublicIp) {
    return nil, false
	}
	return o.PublicIp, true
}

// HasPublicIp returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdUplinksStatusesUplinks) HasPublicIp() bool {
	if o != nil && !isNil(o.PublicIp) {
		return true
	}

	return false
}

// SetPublicIp gets a reference to the given string and assigns it to the PublicIp field.
func (o *OrganizationsOrganizationIdUplinksStatusesUplinks) SetPublicIp(v string) {
	o.PublicIp = &v
}

// GetPrimaryDns returns the PrimaryDns field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdUplinksStatusesUplinks) GetPrimaryDns() string {
	if o == nil || isNil(o.PrimaryDns) {
		var ret string
		return ret
	}
	return *o.PrimaryDns
}

// GetPrimaryDnsOk returns a tuple with the PrimaryDns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdUplinksStatusesUplinks) GetPrimaryDnsOk() (*string, bool) {
	if o == nil || isNil(o.PrimaryDns) {
    return nil, false
	}
	return o.PrimaryDns, true
}

// HasPrimaryDns returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdUplinksStatusesUplinks) HasPrimaryDns() bool {
	if o != nil && !isNil(o.PrimaryDns) {
		return true
	}

	return false
}

// SetPrimaryDns gets a reference to the given string and assigns it to the PrimaryDns field.
func (o *OrganizationsOrganizationIdUplinksStatusesUplinks) SetPrimaryDns(v string) {
	o.PrimaryDns = &v
}

// GetSecondaryDns returns the SecondaryDns field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdUplinksStatusesUplinks) GetSecondaryDns() string {
	if o == nil || isNil(o.SecondaryDns) {
		var ret string
		return ret
	}
	return *o.SecondaryDns
}

// GetSecondaryDnsOk returns a tuple with the SecondaryDns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdUplinksStatusesUplinks) GetSecondaryDnsOk() (*string, bool) {
	if o == nil || isNil(o.SecondaryDns) {
    return nil, false
	}
	return o.SecondaryDns, true
}

// HasSecondaryDns returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdUplinksStatusesUplinks) HasSecondaryDns() bool {
	if o != nil && !isNil(o.SecondaryDns) {
		return true
	}

	return false
}

// SetSecondaryDns gets a reference to the given string and assigns it to the SecondaryDns field.
func (o *OrganizationsOrganizationIdUplinksStatusesUplinks) SetSecondaryDns(v string) {
	o.SecondaryDns = &v
}

// GetIpAssignedBy returns the IpAssignedBy field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdUplinksStatusesUplinks) GetIpAssignedBy() string {
	if o == nil || isNil(o.IpAssignedBy) {
		var ret string
		return ret
	}
	return *o.IpAssignedBy
}

// GetIpAssignedByOk returns a tuple with the IpAssignedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdUplinksStatusesUplinks) GetIpAssignedByOk() (*string, bool) {
	if o == nil || isNil(o.IpAssignedBy) {
    return nil, false
	}
	return o.IpAssignedBy, true
}

// HasIpAssignedBy returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdUplinksStatusesUplinks) HasIpAssignedBy() bool {
	if o != nil && !isNil(o.IpAssignedBy) {
		return true
	}

	return false
}

// SetIpAssignedBy gets a reference to the given string and assigns it to the IpAssignedBy field.
func (o *OrganizationsOrganizationIdUplinksStatusesUplinks) SetIpAssignedBy(v string) {
	o.IpAssignedBy = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdUplinksStatusesUplinks) GetProvider() string {
	if o == nil || isNil(o.Provider) {
		var ret string
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdUplinksStatusesUplinks) GetProviderOk() (*string, bool) {
	if o == nil || isNil(o.Provider) {
    return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdUplinksStatusesUplinks) HasProvider() bool {
	if o != nil && !isNil(o.Provider) {
		return true
	}

	return false
}

// SetProvider gets a reference to the given string and assigns it to the Provider field.
func (o *OrganizationsOrganizationIdUplinksStatusesUplinks) SetProvider(v string) {
	o.Provider = &v
}

// GetSignalStat returns the SignalStat field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdUplinksStatusesUplinks) GetSignalStat() OrganizationsOrganizationIdCellularGatewayUplinkStatusesSignalStat {
	if o == nil || isNil(o.SignalStat) {
		var ret OrganizationsOrganizationIdCellularGatewayUplinkStatusesSignalStat
		return ret
	}
	return *o.SignalStat
}

// GetSignalStatOk returns a tuple with the SignalStat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdUplinksStatusesUplinks) GetSignalStatOk() (*OrganizationsOrganizationIdCellularGatewayUplinkStatusesSignalStat, bool) {
	if o == nil || isNil(o.SignalStat) {
    return nil, false
	}
	return o.SignalStat, true
}

// HasSignalStat returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdUplinksStatusesUplinks) HasSignalStat() bool {
	if o != nil && !isNil(o.SignalStat) {
		return true
	}

	return false
}

// SetSignalStat gets a reference to the given OrganizationsOrganizationIdCellularGatewayUplinkStatusesSignalStat and assigns it to the SignalStat field.
func (o *OrganizationsOrganizationIdUplinksStatusesUplinks) SetSignalStat(v OrganizationsOrganizationIdCellularGatewayUplinkStatusesSignalStat) {
	o.SignalStat = &v
}

// GetConnectionType returns the ConnectionType field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdUplinksStatusesUplinks) GetConnectionType() string {
	if o == nil || isNil(o.ConnectionType) {
		var ret string
		return ret
	}
	return *o.ConnectionType
}

// GetConnectionTypeOk returns a tuple with the ConnectionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdUplinksStatusesUplinks) GetConnectionTypeOk() (*string, bool) {
	if o == nil || isNil(o.ConnectionType) {
    return nil, false
	}
	return o.ConnectionType, true
}

// HasConnectionType returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdUplinksStatusesUplinks) HasConnectionType() bool {
	if o != nil && !isNil(o.ConnectionType) {
		return true
	}

	return false
}

// SetConnectionType gets a reference to the given string and assigns it to the ConnectionType field.
func (o *OrganizationsOrganizationIdUplinksStatusesUplinks) SetConnectionType(v string) {
	o.ConnectionType = &v
}

// GetApn returns the Apn field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdUplinksStatusesUplinks) GetApn() string {
	if o == nil || isNil(o.Apn) {
		var ret string
		return ret
	}
	return *o.Apn
}

// GetApnOk returns a tuple with the Apn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdUplinksStatusesUplinks) GetApnOk() (*string, bool) {
	if o == nil || isNil(o.Apn) {
    return nil, false
	}
	return o.Apn, true
}

// HasApn returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdUplinksStatusesUplinks) HasApn() bool {
	if o != nil && !isNil(o.Apn) {
		return true
	}

	return false
}

// SetApn gets a reference to the given string and assigns it to the Apn field.
func (o *OrganizationsOrganizationIdUplinksStatusesUplinks) SetApn(v string) {
	o.Apn = &v
}

// GetDns1 returns the Dns1 field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdUplinksStatusesUplinks) GetDns1() string {
	if o == nil || isNil(o.Dns1) {
		var ret string
		return ret
	}
	return *o.Dns1
}

// GetDns1Ok returns a tuple with the Dns1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdUplinksStatusesUplinks) GetDns1Ok() (*string, bool) {
	if o == nil || isNil(o.Dns1) {
    return nil, false
	}
	return o.Dns1, true
}

// HasDns1 returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdUplinksStatusesUplinks) HasDns1() bool {
	if o != nil && !isNil(o.Dns1) {
		return true
	}

	return false
}

// SetDns1 gets a reference to the given string and assigns it to the Dns1 field.
func (o *OrganizationsOrganizationIdUplinksStatusesUplinks) SetDns1(v string) {
	o.Dns1 = &v
}

// GetDns2 returns the Dns2 field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdUplinksStatusesUplinks) GetDns2() string {
	if o == nil || isNil(o.Dns2) {
		var ret string
		return ret
	}
	return *o.Dns2
}

// GetDns2Ok returns a tuple with the Dns2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdUplinksStatusesUplinks) GetDns2Ok() (*string, bool) {
	if o == nil || isNil(o.Dns2) {
    return nil, false
	}
	return o.Dns2, true
}

// HasDns2 returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdUplinksStatusesUplinks) HasDns2() bool {
	if o != nil && !isNil(o.Dns2) {
		return true
	}

	return false
}

// SetDns2 gets a reference to the given string and assigns it to the Dns2 field.
func (o *OrganizationsOrganizationIdUplinksStatusesUplinks) SetDns2(v string) {
	o.Dns2 = &v
}

// GetSignalType returns the SignalType field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdUplinksStatusesUplinks) GetSignalType() string {
	if o == nil || isNil(o.SignalType) {
		var ret string
		return ret
	}
	return *o.SignalType
}

// GetSignalTypeOk returns a tuple with the SignalType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdUplinksStatusesUplinks) GetSignalTypeOk() (*string, bool) {
	if o == nil || isNil(o.SignalType) {
    return nil, false
	}
	return o.SignalType, true
}

// HasSignalType returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdUplinksStatusesUplinks) HasSignalType() bool {
	if o != nil && !isNil(o.SignalType) {
		return true
	}

	return false
}

// SetSignalType gets a reference to the given string and assigns it to the SignalType field.
func (o *OrganizationsOrganizationIdUplinksStatusesUplinks) SetSignalType(v string) {
	o.SignalType = &v
}

// GetIccid returns the Iccid field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdUplinksStatusesUplinks) GetIccid() string {
	if o == nil || isNil(o.Iccid) {
		var ret string
		return ret
	}
	return *o.Iccid
}

// GetIccidOk returns a tuple with the Iccid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdUplinksStatusesUplinks) GetIccidOk() (*string, bool) {
	if o == nil || isNil(o.Iccid) {
    return nil, false
	}
	return o.Iccid, true
}

// HasIccid returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdUplinksStatusesUplinks) HasIccid() bool {
	if o != nil && !isNil(o.Iccid) {
		return true
	}

	return false
}

// SetIccid gets a reference to the given string and assigns it to the Iccid field.
func (o *OrganizationsOrganizationIdUplinksStatusesUplinks) SetIccid(v string) {
	o.Iccid = &v
}

func (o OrganizationsOrganizationIdUplinksStatusesUplinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Interface) {
		toSerialize["interface"] = o.Interface
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.Ip) {
		toSerialize["ip"] = o.Ip
	}
	if !isNil(o.Gateway) {
		toSerialize["gateway"] = o.Gateway
	}
	if !isNil(o.PublicIp) {
		toSerialize["publicIp"] = o.PublicIp
	}
	if !isNil(o.PrimaryDns) {
		toSerialize["primaryDns"] = o.PrimaryDns
	}
	if !isNil(o.SecondaryDns) {
		toSerialize["secondaryDns"] = o.SecondaryDns
	}
	if !isNil(o.IpAssignedBy) {
		toSerialize["ipAssignedBy"] = o.IpAssignedBy
	}
	if !isNil(o.Provider) {
		toSerialize["provider"] = o.Provider
	}
	if !isNil(o.SignalStat) {
		toSerialize["signalStat"] = o.SignalStat
	}
	if !isNil(o.ConnectionType) {
		toSerialize["connectionType"] = o.ConnectionType
	}
	if !isNil(o.Apn) {
		toSerialize["apn"] = o.Apn
	}
	if !isNil(o.Dns1) {
		toSerialize["dns1"] = o.Dns1
	}
	if !isNil(o.Dns2) {
		toSerialize["dns2"] = o.Dns2
	}
	if !isNil(o.SignalType) {
		toSerialize["signalType"] = o.SignalType
	}
	if !isNil(o.Iccid) {
		toSerialize["iccid"] = o.Iccid
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdUplinksStatusesUplinks struct {
	value *OrganizationsOrganizationIdUplinksStatusesUplinks
	isSet bool
}

func (v NullableOrganizationsOrganizationIdUplinksStatusesUplinks) Get() *OrganizationsOrganizationIdUplinksStatusesUplinks {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdUplinksStatusesUplinks) Set(val *OrganizationsOrganizationIdUplinksStatusesUplinks) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdUplinksStatusesUplinks) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdUplinksStatusesUplinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdUplinksStatusesUplinks(val *OrganizationsOrganizationIdUplinksStatusesUplinks) *NullableOrganizationsOrganizationIdUplinksStatusesUplinks {
	return &NullableOrganizationsOrganizationIdUplinksStatusesUplinks{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdUplinksStatusesUplinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdUplinksStatusesUplinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


