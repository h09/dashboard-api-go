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

// NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIp Additional IP attributes of the packet.
type NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIp struct {
	// IP ID of the packet.
	Id *string `json:"id,omitempty"`
	// IP version of the packet.
	Version *int32 `json:"version,omitempty"`
	// IP length of the packet.
	Length *int32 `json:"length,omitempty"`
	// IP header length of the packet.
	HeaderLength *int32 `json:"headerLength,omitempty"`
	// IP protocol number of the packet.
	Protocol *int32 `json:"protocol,omitempty"`
	// Time to live of the packet.
	Ttl *int32 `json:"ttl,omitempty"`
	Dscp *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIpDscp `json:"dscp,omitempty"`
}

// NewNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIp instantiates a new NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIp() *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIp {
	this := NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIp{}
	return &this
}

// NewNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIpWithDefaults instantiates a new NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIpWithDefaults() *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIp {
	this := NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIp{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIp) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIp) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIp) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIp) SetId(v string) {
	o.Id = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIp) GetVersion() int32 {
	if o == nil || isNil(o.Version) {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIp) GetVersionOk() (*int32, bool) {
	if o == nil || isNil(o.Version) {
    return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIp) HasVersion() bool {
	if o != nil && !isNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIp) SetVersion(v int32) {
	o.Version = &v
}

// GetLength returns the Length field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIp) GetLength() int32 {
	if o == nil || isNil(o.Length) {
		var ret int32
		return ret
	}
	return *o.Length
}

// GetLengthOk returns a tuple with the Length field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIp) GetLengthOk() (*int32, bool) {
	if o == nil || isNil(o.Length) {
    return nil, false
	}
	return o.Length, true
}

// HasLength returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIp) HasLength() bool {
	if o != nil && !isNil(o.Length) {
		return true
	}

	return false
}

// SetLength gets a reference to the given int32 and assigns it to the Length field.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIp) SetLength(v int32) {
	o.Length = &v
}

// GetHeaderLength returns the HeaderLength field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIp) GetHeaderLength() int32 {
	if o == nil || isNil(o.HeaderLength) {
		var ret int32
		return ret
	}
	return *o.HeaderLength
}

// GetHeaderLengthOk returns a tuple with the HeaderLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIp) GetHeaderLengthOk() (*int32, bool) {
	if o == nil || isNil(o.HeaderLength) {
    return nil, false
	}
	return o.HeaderLength, true
}

// HasHeaderLength returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIp) HasHeaderLength() bool {
	if o != nil && !isNil(o.HeaderLength) {
		return true
	}

	return false
}

// SetHeaderLength gets a reference to the given int32 and assigns it to the HeaderLength field.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIp) SetHeaderLength(v int32) {
	o.HeaderLength = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIp) GetProtocol() int32 {
	if o == nil || isNil(o.Protocol) {
		var ret int32
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIp) GetProtocolOk() (*int32, bool) {
	if o == nil || isNil(o.Protocol) {
    return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIp) HasProtocol() bool {
	if o != nil && !isNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given int32 and assigns it to the Protocol field.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIp) SetProtocol(v int32) {
	o.Protocol = &v
}

// GetTtl returns the Ttl field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIp) GetTtl() int32 {
	if o == nil || isNil(o.Ttl) {
		var ret int32
		return ret
	}
	return *o.Ttl
}

// GetTtlOk returns a tuple with the Ttl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIp) GetTtlOk() (*int32, bool) {
	if o == nil || isNil(o.Ttl) {
    return nil, false
	}
	return o.Ttl, true
}

// HasTtl returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIp) HasTtl() bool {
	if o != nil && !isNil(o.Ttl) {
		return true
	}

	return false
}

// SetTtl gets a reference to the given int32 and assigns it to the Ttl field.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIp) SetTtl(v int32) {
	o.Ttl = &v
}

// GetDscp returns the Dscp field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIp) GetDscp() NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIpDscp {
	if o == nil || isNil(o.Dscp) {
		var ret NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIpDscp
		return ret
	}
	return *o.Dscp
}

// GetDscpOk returns a tuple with the Dscp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIp) GetDscpOk() (*NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIpDscp, bool) {
	if o == nil || isNil(o.Dscp) {
    return nil, false
	}
	return o.Dscp, true
}

// HasDscp returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIp) HasDscp() bool {
	if o != nil && !isNil(o.Dscp) {
		return true
	}

	return false
}

// SetDscp gets a reference to the given NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIpDscp and assigns it to the Dscp field.
func (o *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIp) SetDscp(v NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIpDscp) {
	o.Dscp = &v
}

func (o NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !isNil(o.Length) {
		toSerialize["length"] = o.Length
	}
	if !isNil(o.HeaderLength) {
		toSerialize["headerLength"] = o.HeaderLength
	}
	if !isNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}
	if !isNil(o.Ttl) {
		toSerialize["ttl"] = o.Ttl
	}
	if !isNil(o.Dscp) {
		toSerialize["dscp"] = o.Dscp
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIp struct {
	value *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIp
	isSet bool
}

func (v NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIp) Get() *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIp {
	return v.value
}

func (v *NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIp) Set(val *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIp) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIp) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIp(val *NetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIp) *NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIp {
	return &NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIp{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdSwitchDhcpV4ServersSeenLastPacketIp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


