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

// checks if the GetNetworkSwitchDhcpServerPolicy200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkSwitchDhcpServerPolicy200Response{}

// GetNetworkSwitchDhcpServerPolicy200Response struct for GetNetworkSwitchDhcpServerPolicy200Response
type GetNetworkSwitchDhcpServerPolicy200Response struct {
	Alerts *GetNetworkSwitchDhcpServerPolicy200ResponseAlerts `json:"alerts,omitempty"`
	// 'allow' or 'block' new DHCP servers. Default value is 'allow'.
	DefaultPolicy *string `json:"defaultPolicy,omitempty"`
	// List the MAC addresses of DHCP servers to block on the network when defaultPolicy is set       to allow.An empty array will clear the entries.
	BlockedServers []string `json:"blockedServers,omitempty"`
	// List the MAC addresses of DHCP servers to permit on the network when defaultPolicy is set       to block.An empty array will clear the entries.
	AllowedServers []string `json:"allowedServers,omitempty"`
	ArpInspection *GetNetworkSwitchDhcpServerPolicy200ResponseArpInspection `json:"arpInspection,omitempty"`
}

// NewGetNetworkSwitchDhcpServerPolicy200Response instantiates a new GetNetworkSwitchDhcpServerPolicy200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkSwitchDhcpServerPolicy200Response() *GetNetworkSwitchDhcpServerPolicy200Response {
	this := GetNetworkSwitchDhcpServerPolicy200Response{}
	return &this
}

// NewGetNetworkSwitchDhcpServerPolicy200ResponseWithDefaults instantiates a new GetNetworkSwitchDhcpServerPolicy200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkSwitchDhcpServerPolicy200ResponseWithDefaults() *GetNetworkSwitchDhcpServerPolicy200Response {
	this := GetNetworkSwitchDhcpServerPolicy200Response{}
	return &this
}

// GetAlerts returns the Alerts field value if set, zero value otherwise.
func (o *GetNetworkSwitchDhcpServerPolicy200Response) GetAlerts() GetNetworkSwitchDhcpServerPolicy200ResponseAlerts {
	if o == nil || IsNil(o.Alerts) {
		var ret GetNetworkSwitchDhcpServerPolicy200ResponseAlerts
		return ret
	}
	return *o.Alerts
}

// GetAlertsOk returns a tuple with the Alerts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchDhcpServerPolicy200Response) GetAlertsOk() (*GetNetworkSwitchDhcpServerPolicy200ResponseAlerts, bool) {
	if o == nil || IsNil(o.Alerts) {
		return nil, false
	}
	return o.Alerts, true
}

// HasAlerts returns a boolean if a field has been set.
func (o *GetNetworkSwitchDhcpServerPolicy200Response) HasAlerts() bool {
	if o != nil && !IsNil(o.Alerts) {
		return true
	}

	return false
}

// SetAlerts gets a reference to the given GetNetworkSwitchDhcpServerPolicy200ResponseAlerts and assigns it to the Alerts field.
func (o *GetNetworkSwitchDhcpServerPolicy200Response) SetAlerts(v GetNetworkSwitchDhcpServerPolicy200ResponseAlerts) {
	o.Alerts = &v
}

// GetDefaultPolicy returns the DefaultPolicy field value if set, zero value otherwise.
func (o *GetNetworkSwitchDhcpServerPolicy200Response) GetDefaultPolicy() string {
	if o == nil || IsNil(o.DefaultPolicy) {
		var ret string
		return ret
	}
	return *o.DefaultPolicy
}

// GetDefaultPolicyOk returns a tuple with the DefaultPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchDhcpServerPolicy200Response) GetDefaultPolicyOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultPolicy) {
		return nil, false
	}
	return o.DefaultPolicy, true
}

// HasDefaultPolicy returns a boolean if a field has been set.
func (o *GetNetworkSwitchDhcpServerPolicy200Response) HasDefaultPolicy() bool {
	if o != nil && !IsNil(o.DefaultPolicy) {
		return true
	}

	return false
}

// SetDefaultPolicy gets a reference to the given string and assigns it to the DefaultPolicy field.
func (o *GetNetworkSwitchDhcpServerPolicy200Response) SetDefaultPolicy(v string) {
	o.DefaultPolicy = &v
}

// GetBlockedServers returns the BlockedServers field value if set, zero value otherwise.
func (o *GetNetworkSwitchDhcpServerPolicy200Response) GetBlockedServers() []string {
	if o == nil || IsNil(o.BlockedServers) {
		var ret []string
		return ret
	}
	return o.BlockedServers
}

// GetBlockedServersOk returns a tuple with the BlockedServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchDhcpServerPolicy200Response) GetBlockedServersOk() ([]string, bool) {
	if o == nil || IsNil(o.BlockedServers) {
		return nil, false
	}
	return o.BlockedServers, true
}

// HasBlockedServers returns a boolean if a field has been set.
func (o *GetNetworkSwitchDhcpServerPolicy200Response) HasBlockedServers() bool {
	if o != nil && !IsNil(o.BlockedServers) {
		return true
	}

	return false
}

// SetBlockedServers gets a reference to the given []string and assigns it to the BlockedServers field.
func (o *GetNetworkSwitchDhcpServerPolicy200Response) SetBlockedServers(v []string) {
	o.BlockedServers = v
}

// GetAllowedServers returns the AllowedServers field value if set, zero value otherwise.
func (o *GetNetworkSwitchDhcpServerPolicy200Response) GetAllowedServers() []string {
	if o == nil || IsNil(o.AllowedServers) {
		var ret []string
		return ret
	}
	return o.AllowedServers
}

// GetAllowedServersOk returns a tuple with the AllowedServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchDhcpServerPolicy200Response) GetAllowedServersOk() ([]string, bool) {
	if o == nil || IsNil(o.AllowedServers) {
		return nil, false
	}
	return o.AllowedServers, true
}

// HasAllowedServers returns a boolean if a field has been set.
func (o *GetNetworkSwitchDhcpServerPolicy200Response) HasAllowedServers() bool {
	if o != nil && !IsNil(o.AllowedServers) {
		return true
	}

	return false
}

// SetAllowedServers gets a reference to the given []string and assigns it to the AllowedServers field.
func (o *GetNetworkSwitchDhcpServerPolicy200Response) SetAllowedServers(v []string) {
	o.AllowedServers = v
}

// GetArpInspection returns the ArpInspection field value if set, zero value otherwise.
func (o *GetNetworkSwitchDhcpServerPolicy200Response) GetArpInspection() GetNetworkSwitchDhcpServerPolicy200ResponseArpInspection {
	if o == nil || IsNil(o.ArpInspection) {
		var ret GetNetworkSwitchDhcpServerPolicy200ResponseArpInspection
		return ret
	}
	return *o.ArpInspection
}

// GetArpInspectionOk returns a tuple with the ArpInspection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchDhcpServerPolicy200Response) GetArpInspectionOk() (*GetNetworkSwitchDhcpServerPolicy200ResponseArpInspection, bool) {
	if o == nil || IsNil(o.ArpInspection) {
		return nil, false
	}
	return o.ArpInspection, true
}

// HasArpInspection returns a boolean if a field has been set.
func (o *GetNetworkSwitchDhcpServerPolicy200Response) HasArpInspection() bool {
	if o != nil && !IsNil(o.ArpInspection) {
		return true
	}

	return false
}

// SetArpInspection gets a reference to the given GetNetworkSwitchDhcpServerPolicy200ResponseArpInspection and assigns it to the ArpInspection field.
func (o *GetNetworkSwitchDhcpServerPolicy200Response) SetArpInspection(v GetNetworkSwitchDhcpServerPolicy200ResponseArpInspection) {
	o.ArpInspection = &v
}

func (o GetNetworkSwitchDhcpServerPolicy200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkSwitchDhcpServerPolicy200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Alerts) {
		toSerialize["alerts"] = o.Alerts
	}
	if !IsNil(o.DefaultPolicy) {
		toSerialize["defaultPolicy"] = o.DefaultPolicy
	}
	if !IsNil(o.BlockedServers) {
		toSerialize["blockedServers"] = o.BlockedServers
	}
	if !IsNil(o.AllowedServers) {
		toSerialize["allowedServers"] = o.AllowedServers
	}
	if !IsNil(o.ArpInspection) {
		toSerialize["arpInspection"] = o.ArpInspection
	}
	return toSerialize, nil
}

type NullableGetNetworkSwitchDhcpServerPolicy200Response struct {
	value *GetNetworkSwitchDhcpServerPolicy200Response
	isSet bool
}

func (v NullableGetNetworkSwitchDhcpServerPolicy200Response) Get() *GetNetworkSwitchDhcpServerPolicy200Response {
	return v.value
}

func (v *NullableGetNetworkSwitchDhcpServerPolicy200Response) Set(val *GetNetworkSwitchDhcpServerPolicy200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkSwitchDhcpServerPolicy200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkSwitchDhcpServerPolicy200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkSwitchDhcpServerPolicy200Response(val *GetNetworkSwitchDhcpServerPolicy200Response) *NullableGetNetworkSwitchDhcpServerPolicy200Response {
	return &NullableGetNetworkSwitchDhcpServerPolicy200Response{value: val, isSet: true}
}

func (v NullableGetNetworkSwitchDhcpServerPolicy200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkSwitchDhcpServerPolicy200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


