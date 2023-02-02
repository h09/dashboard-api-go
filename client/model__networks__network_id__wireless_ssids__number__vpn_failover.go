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

// NetworksNetworkIdWirelessSsidsNumberVpnFailover Secondary VPN concentrator settings. This is only used when two VPN concentrators are configured on the SSID.
type NetworksNetworkIdWirelessSsidsNumberVpnFailover struct {
	// IP addressed reserved on DHCP server where SSID will terminate.
	RequestIp *string `json:"requestIp,omitempty"`
	// Idle timer interval in seconds.
	HeartbeatInterval *int32 `json:"heartbeatInterval,omitempty"`
	// Idle timer timeout in seconds.
	IdleTimeout *int32 `json:"idleTimeout,omitempty"`
}

// NewNetworksNetworkIdWirelessSsidsNumberVpnFailover instantiates a new NetworksNetworkIdWirelessSsidsNumberVpnFailover object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdWirelessSsidsNumberVpnFailover() *NetworksNetworkIdWirelessSsidsNumberVpnFailover {
	this := NetworksNetworkIdWirelessSsidsNumberVpnFailover{}
	return &this
}

// NewNetworksNetworkIdWirelessSsidsNumberVpnFailoverWithDefaults instantiates a new NetworksNetworkIdWirelessSsidsNumberVpnFailover object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdWirelessSsidsNumberVpnFailoverWithDefaults() *NetworksNetworkIdWirelessSsidsNumberVpnFailover {
	this := NetworksNetworkIdWirelessSsidsNumberVpnFailover{}
	return &this
}

// GetRequestIp returns the RequestIp field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsNumberVpnFailover) GetRequestIp() string {
	if o == nil || isNil(o.RequestIp) {
		var ret string
		return ret
	}
	return *o.RequestIp
}

// GetRequestIpOk returns a tuple with the RequestIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberVpnFailover) GetRequestIpOk() (*string, bool) {
	if o == nil || isNil(o.RequestIp) {
    return nil, false
	}
	return o.RequestIp, true
}

// HasRequestIp returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberVpnFailover) HasRequestIp() bool {
	if o != nil && !isNil(o.RequestIp) {
		return true
	}

	return false
}

// SetRequestIp gets a reference to the given string and assigns it to the RequestIp field.
func (o *NetworksNetworkIdWirelessSsidsNumberVpnFailover) SetRequestIp(v string) {
	o.RequestIp = &v
}

// GetHeartbeatInterval returns the HeartbeatInterval field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsNumberVpnFailover) GetHeartbeatInterval() int32 {
	if o == nil || isNil(o.HeartbeatInterval) {
		var ret int32
		return ret
	}
	return *o.HeartbeatInterval
}

// GetHeartbeatIntervalOk returns a tuple with the HeartbeatInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberVpnFailover) GetHeartbeatIntervalOk() (*int32, bool) {
	if o == nil || isNil(o.HeartbeatInterval) {
    return nil, false
	}
	return o.HeartbeatInterval, true
}

// HasHeartbeatInterval returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberVpnFailover) HasHeartbeatInterval() bool {
	if o != nil && !isNil(o.HeartbeatInterval) {
		return true
	}

	return false
}

// SetHeartbeatInterval gets a reference to the given int32 and assigns it to the HeartbeatInterval field.
func (o *NetworksNetworkIdWirelessSsidsNumberVpnFailover) SetHeartbeatInterval(v int32) {
	o.HeartbeatInterval = &v
}

// GetIdleTimeout returns the IdleTimeout field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsNumberVpnFailover) GetIdleTimeout() int32 {
	if o == nil || isNil(o.IdleTimeout) {
		var ret int32
		return ret
	}
	return *o.IdleTimeout
}

// GetIdleTimeoutOk returns a tuple with the IdleTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberVpnFailover) GetIdleTimeoutOk() (*int32, bool) {
	if o == nil || isNil(o.IdleTimeout) {
    return nil, false
	}
	return o.IdleTimeout, true
}

// HasIdleTimeout returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberVpnFailover) HasIdleTimeout() bool {
	if o != nil && !isNil(o.IdleTimeout) {
		return true
	}

	return false
}

// SetIdleTimeout gets a reference to the given int32 and assigns it to the IdleTimeout field.
func (o *NetworksNetworkIdWirelessSsidsNumberVpnFailover) SetIdleTimeout(v int32) {
	o.IdleTimeout = &v
}

func (o NetworksNetworkIdWirelessSsidsNumberVpnFailover) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.RequestIp) {
		toSerialize["requestIp"] = o.RequestIp
	}
	if !isNil(o.HeartbeatInterval) {
		toSerialize["heartbeatInterval"] = o.HeartbeatInterval
	}
	if !isNil(o.IdleTimeout) {
		toSerialize["idleTimeout"] = o.IdleTimeout
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdWirelessSsidsNumberVpnFailover struct {
	value *NetworksNetworkIdWirelessSsidsNumberVpnFailover
	isSet bool
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberVpnFailover) Get() *NetworksNetworkIdWirelessSsidsNumberVpnFailover {
	return v.value
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberVpnFailover) Set(val *NetworksNetworkIdWirelessSsidsNumberVpnFailover) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberVpnFailover) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberVpnFailover) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdWirelessSsidsNumberVpnFailover(val *NetworksNetworkIdWirelessSsidsNumberVpnFailover) *NullableNetworksNetworkIdWirelessSsidsNumberVpnFailover {
	return &NullableNetworksNetworkIdWirelessSsidsNumberVpnFailover{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberVpnFailover) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberVpnFailover) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


