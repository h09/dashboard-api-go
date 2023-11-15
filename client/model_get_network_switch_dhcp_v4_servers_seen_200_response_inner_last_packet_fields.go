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

// checks if the GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields{}

// GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields DHCP-specific fields of the packet.
type GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields struct {
	// Operation code of the packet.
	Op *int32 `json:"op,omitempty"`
	// Hardware type code of the packet.
	Htype *int32 `json:"htype,omitempty"`
	// Hardware length of the packet.
	Hlen *int32 `json:"hlen,omitempty"`
	// Number of hops the packet took.
	Hops *int32 `json:"hops,omitempty"`
	// Transaction id of the packet.
	Xid *string `json:"xid,omitempty"`
	// Number of seconds since receiving the packet.
	Secs *int32 `json:"secs,omitempty"`
	// Packet flags.
	Flags *string `json:"flags,omitempty"`
	// Client IP address of the packet.
	Ciaddr *string `json:"ciaddr,omitempty"`
	// Assigned IP address of the packet.
	Yiaddr *string `json:"yiaddr,omitempty"`
	// Server IP address of the packet.
	Siaddr *string `json:"siaddr,omitempty"`
	// Gateway IP address of the packet.
	Giaddr *string `json:"giaddr,omitempty"`
	// Client hardware address of the packet.
	Chaddr *string `json:"chaddr,omitempty"`
	// Server identifier address of the packet.
	Sname *string `json:"sname,omitempty"`
	// Magic cookie of the packet.
	MagicCookie *string `json:"magicCookie,omitempty"`
	// Additional DHCP options of the packet.
	Options []GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFieldsOptionsInner `json:"options,omitempty"`
}

// NewGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields instantiates a new GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields() *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields {
	this := GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields{}
	return &this
}

// NewGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFieldsWithDefaults instantiates a new GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFieldsWithDefaults() *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields {
	this := GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields{}
	return &this
}

// GetOp returns the Op field value if set, zero value otherwise.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) GetOp() int32 {
	if o == nil || IsNil(o.Op) {
		var ret int32
		return ret
	}
	return *o.Op
}

// GetOpOk returns a tuple with the Op field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) GetOpOk() (*int32, bool) {
	if o == nil || IsNil(o.Op) {
		return nil, false
	}
	return o.Op, true
}

// HasOp returns a boolean if a field has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) HasOp() bool {
	if o != nil && !IsNil(o.Op) {
		return true
	}

	return false
}

// SetOp gets a reference to the given int32 and assigns it to the Op field.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) SetOp(v int32) {
	o.Op = &v
}

// GetHtype returns the Htype field value if set, zero value otherwise.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) GetHtype() int32 {
	if o == nil || IsNil(o.Htype) {
		var ret int32
		return ret
	}
	return *o.Htype
}

// GetHtypeOk returns a tuple with the Htype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) GetHtypeOk() (*int32, bool) {
	if o == nil || IsNil(o.Htype) {
		return nil, false
	}
	return o.Htype, true
}

// HasHtype returns a boolean if a field has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) HasHtype() bool {
	if o != nil && !IsNil(o.Htype) {
		return true
	}

	return false
}

// SetHtype gets a reference to the given int32 and assigns it to the Htype field.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) SetHtype(v int32) {
	o.Htype = &v
}

// GetHlen returns the Hlen field value if set, zero value otherwise.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) GetHlen() int32 {
	if o == nil || IsNil(o.Hlen) {
		var ret int32
		return ret
	}
	return *o.Hlen
}

// GetHlenOk returns a tuple with the Hlen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) GetHlenOk() (*int32, bool) {
	if o == nil || IsNil(o.Hlen) {
		return nil, false
	}
	return o.Hlen, true
}

// HasHlen returns a boolean if a field has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) HasHlen() bool {
	if o != nil && !IsNil(o.Hlen) {
		return true
	}

	return false
}

// SetHlen gets a reference to the given int32 and assigns it to the Hlen field.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) SetHlen(v int32) {
	o.Hlen = &v
}

// GetHops returns the Hops field value if set, zero value otherwise.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) GetHops() int32 {
	if o == nil || IsNil(o.Hops) {
		var ret int32
		return ret
	}
	return *o.Hops
}

// GetHopsOk returns a tuple with the Hops field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) GetHopsOk() (*int32, bool) {
	if o == nil || IsNil(o.Hops) {
		return nil, false
	}
	return o.Hops, true
}

// HasHops returns a boolean if a field has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) HasHops() bool {
	if o != nil && !IsNil(o.Hops) {
		return true
	}

	return false
}

// SetHops gets a reference to the given int32 and assigns it to the Hops field.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) SetHops(v int32) {
	o.Hops = &v
}

// GetXid returns the Xid field value if set, zero value otherwise.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) GetXid() string {
	if o == nil || IsNil(o.Xid) {
		var ret string
		return ret
	}
	return *o.Xid
}

// GetXidOk returns a tuple with the Xid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) GetXidOk() (*string, bool) {
	if o == nil || IsNil(o.Xid) {
		return nil, false
	}
	return o.Xid, true
}

// HasXid returns a boolean if a field has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) HasXid() bool {
	if o != nil && !IsNil(o.Xid) {
		return true
	}

	return false
}

// SetXid gets a reference to the given string and assigns it to the Xid field.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) SetXid(v string) {
	o.Xid = &v
}

// GetSecs returns the Secs field value if set, zero value otherwise.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) GetSecs() int32 {
	if o == nil || IsNil(o.Secs) {
		var ret int32
		return ret
	}
	return *o.Secs
}

// GetSecsOk returns a tuple with the Secs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) GetSecsOk() (*int32, bool) {
	if o == nil || IsNil(o.Secs) {
		return nil, false
	}
	return o.Secs, true
}

// HasSecs returns a boolean if a field has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) HasSecs() bool {
	if o != nil && !IsNil(o.Secs) {
		return true
	}

	return false
}

// SetSecs gets a reference to the given int32 and assigns it to the Secs field.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) SetSecs(v int32) {
	o.Secs = &v
}

// GetFlags returns the Flags field value if set, zero value otherwise.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) GetFlags() string {
	if o == nil || IsNil(o.Flags) {
		var ret string
		return ret
	}
	return *o.Flags
}

// GetFlagsOk returns a tuple with the Flags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) GetFlagsOk() (*string, bool) {
	if o == nil || IsNil(o.Flags) {
		return nil, false
	}
	return o.Flags, true
}

// HasFlags returns a boolean if a field has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) HasFlags() bool {
	if o != nil && !IsNil(o.Flags) {
		return true
	}

	return false
}

// SetFlags gets a reference to the given string and assigns it to the Flags field.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) SetFlags(v string) {
	o.Flags = &v
}

// GetCiaddr returns the Ciaddr field value if set, zero value otherwise.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) GetCiaddr() string {
	if o == nil || IsNil(o.Ciaddr) {
		var ret string
		return ret
	}
	return *o.Ciaddr
}

// GetCiaddrOk returns a tuple with the Ciaddr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) GetCiaddrOk() (*string, bool) {
	if o == nil || IsNil(o.Ciaddr) {
		return nil, false
	}
	return o.Ciaddr, true
}

// HasCiaddr returns a boolean if a field has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) HasCiaddr() bool {
	if o != nil && !IsNil(o.Ciaddr) {
		return true
	}

	return false
}

// SetCiaddr gets a reference to the given string and assigns it to the Ciaddr field.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) SetCiaddr(v string) {
	o.Ciaddr = &v
}

// GetYiaddr returns the Yiaddr field value if set, zero value otherwise.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) GetYiaddr() string {
	if o == nil || IsNil(o.Yiaddr) {
		var ret string
		return ret
	}
	return *o.Yiaddr
}

// GetYiaddrOk returns a tuple with the Yiaddr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) GetYiaddrOk() (*string, bool) {
	if o == nil || IsNil(o.Yiaddr) {
		return nil, false
	}
	return o.Yiaddr, true
}

// HasYiaddr returns a boolean if a field has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) HasYiaddr() bool {
	if o != nil && !IsNil(o.Yiaddr) {
		return true
	}

	return false
}

// SetYiaddr gets a reference to the given string and assigns it to the Yiaddr field.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) SetYiaddr(v string) {
	o.Yiaddr = &v
}

// GetSiaddr returns the Siaddr field value if set, zero value otherwise.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) GetSiaddr() string {
	if o == nil || IsNil(o.Siaddr) {
		var ret string
		return ret
	}
	return *o.Siaddr
}

// GetSiaddrOk returns a tuple with the Siaddr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) GetSiaddrOk() (*string, bool) {
	if o == nil || IsNil(o.Siaddr) {
		return nil, false
	}
	return o.Siaddr, true
}

// HasSiaddr returns a boolean if a field has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) HasSiaddr() bool {
	if o != nil && !IsNil(o.Siaddr) {
		return true
	}

	return false
}

// SetSiaddr gets a reference to the given string and assigns it to the Siaddr field.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) SetSiaddr(v string) {
	o.Siaddr = &v
}

// GetGiaddr returns the Giaddr field value if set, zero value otherwise.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) GetGiaddr() string {
	if o == nil || IsNil(o.Giaddr) {
		var ret string
		return ret
	}
	return *o.Giaddr
}

// GetGiaddrOk returns a tuple with the Giaddr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) GetGiaddrOk() (*string, bool) {
	if o == nil || IsNil(o.Giaddr) {
		return nil, false
	}
	return o.Giaddr, true
}

// HasGiaddr returns a boolean if a field has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) HasGiaddr() bool {
	if o != nil && !IsNil(o.Giaddr) {
		return true
	}

	return false
}

// SetGiaddr gets a reference to the given string and assigns it to the Giaddr field.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) SetGiaddr(v string) {
	o.Giaddr = &v
}

// GetChaddr returns the Chaddr field value if set, zero value otherwise.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) GetChaddr() string {
	if o == nil || IsNil(o.Chaddr) {
		var ret string
		return ret
	}
	return *o.Chaddr
}

// GetChaddrOk returns a tuple with the Chaddr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) GetChaddrOk() (*string, bool) {
	if o == nil || IsNil(o.Chaddr) {
		return nil, false
	}
	return o.Chaddr, true
}

// HasChaddr returns a boolean if a field has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) HasChaddr() bool {
	if o != nil && !IsNil(o.Chaddr) {
		return true
	}

	return false
}

// SetChaddr gets a reference to the given string and assigns it to the Chaddr field.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) SetChaddr(v string) {
	o.Chaddr = &v
}

// GetSname returns the Sname field value if set, zero value otherwise.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) GetSname() string {
	if o == nil || IsNil(o.Sname) {
		var ret string
		return ret
	}
	return *o.Sname
}

// GetSnameOk returns a tuple with the Sname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) GetSnameOk() (*string, bool) {
	if o == nil || IsNil(o.Sname) {
		return nil, false
	}
	return o.Sname, true
}

// HasSname returns a boolean if a field has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) HasSname() bool {
	if o != nil && !IsNil(o.Sname) {
		return true
	}

	return false
}

// SetSname gets a reference to the given string and assigns it to the Sname field.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) SetSname(v string) {
	o.Sname = &v
}

// GetMagicCookie returns the MagicCookie field value if set, zero value otherwise.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) GetMagicCookie() string {
	if o == nil || IsNil(o.MagicCookie) {
		var ret string
		return ret
	}
	return *o.MagicCookie
}

// GetMagicCookieOk returns a tuple with the MagicCookie field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) GetMagicCookieOk() (*string, bool) {
	if o == nil || IsNil(o.MagicCookie) {
		return nil, false
	}
	return o.MagicCookie, true
}

// HasMagicCookie returns a boolean if a field has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) HasMagicCookie() bool {
	if o != nil && !IsNil(o.MagicCookie) {
		return true
	}

	return false
}

// SetMagicCookie gets a reference to the given string and assigns it to the MagicCookie field.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) SetMagicCookie(v string) {
	o.MagicCookie = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) GetOptions() []GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFieldsOptionsInner {
	if o == nil || IsNil(o.Options) {
		var ret []GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFieldsOptionsInner
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) GetOptionsOk() ([]GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFieldsOptionsInner, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given []GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFieldsOptionsInner and assigns it to the Options field.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) SetOptions(v []GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFieldsOptionsInner) {
	o.Options = v
}

func (o GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Op) {
		toSerialize["op"] = o.Op
	}
	if !IsNil(o.Htype) {
		toSerialize["htype"] = o.Htype
	}
	if !IsNil(o.Hlen) {
		toSerialize["hlen"] = o.Hlen
	}
	if !IsNil(o.Hops) {
		toSerialize["hops"] = o.Hops
	}
	if !IsNil(o.Xid) {
		toSerialize["xid"] = o.Xid
	}
	if !IsNil(o.Secs) {
		toSerialize["secs"] = o.Secs
	}
	if !IsNil(o.Flags) {
		toSerialize["flags"] = o.Flags
	}
	if !IsNil(o.Ciaddr) {
		toSerialize["ciaddr"] = o.Ciaddr
	}
	if !IsNil(o.Yiaddr) {
		toSerialize["yiaddr"] = o.Yiaddr
	}
	if !IsNil(o.Siaddr) {
		toSerialize["siaddr"] = o.Siaddr
	}
	if !IsNil(o.Giaddr) {
		toSerialize["giaddr"] = o.Giaddr
	}
	if !IsNil(o.Chaddr) {
		toSerialize["chaddr"] = o.Chaddr
	}
	if !IsNil(o.Sname) {
		toSerialize["sname"] = o.Sname
	}
	if !IsNil(o.MagicCookie) {
		toSerialize["magicCookie"] = o.MagicCookie
	}
	if !IsNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	return toSerialize, nil
}

type NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields struct {
	value *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields
	isSet bool
}

func (v NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) Get() *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields {
	return v.value
}

func (v *NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) Set(val *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields(val *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) *NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields {
	return &NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields{value: val, isSet: true}
}

func (v NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


