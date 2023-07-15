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

// checks if the GetNetworkWirelessBluetoothSettings200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkWirelessBluetoothSettings200Response{}

// GetNetworkWirelessBluetoothSettings200Response struct for GetNetworkWirelessBluetoothSettings200Response
type GetNetworkWirelessBluetoothSettings200Response struct {
	// Whether APs will scan for Bluetooth enabled clients.
	ScanningEnabled *bool `json:"scanningEnabled,omitempty"`
	// Whether APs will advertise beacons.
	AdvertisingEnabled *bool `json:"advertisingEnabled,omitempty"`
	// The UUID to be used in the beacon identifier.
	Uuid *string `json:"uuid,omitempty"`
	// The way major and minor number should be assigned to nodes in the network. ('Unique', 'Non-unique')
	MajorMinorAssignmentMode *string `json:"majorMinorAssignmentMode,omitempty"`
	// The major number to be used in the beacon identifier. Only valid in 'Non-unique' mode.
	Major *int32 `json:"major,omitempty"`
	// The minor number to be used in the beacon identifier. Only valid in 'Non-unique' mode.
	Minor *int32 `json:"minor,omitempty"`
	// Whether ESL is enabled on this network.
	EslEnabled *bool `json:"eslEnabled,omitempty"`
}

// NewGetNetworkWirelessBluetoothSettings200Response instantiates a new GetNetworkWirelessBluetoothSettings200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkWirelessBluetoothSettings200Response() *GetNetworkWirelessBluetoothSettings200Response {
	this := GetNetworkWirelessBluetoothSettings200Response{}
	return &this
}

// NewGetNetworkWirelessBluetoothSettings200ResponseWithDefaults instantiates a new GetNetworkWirelessBluetoothSettings200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkWirelessBluetoothSettings200ResponseWithDefaults() *GetNetworkWirelessBluetoothSettings200Response {
	this := GetNetworkWirelessBluetoothSettings200Response{}
	return &this
}

// GetScanningEnabled returns the ScanningEnabled field value if set, zero value otherwise.
func (o *GetNetworkWirelessBluetoothSettings200Response) GetScanningEnabled() bool {
	if o == nil || IsNil(o.ScanningEnabled) {
		var ret bool
		return ret
	}
	return *o.ScanningEnabled
}

// GetScanningEnabledOk returns a tuple with the ScanningEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessBluetoothSettings200Response) GetScanningEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.ScanningEnabled) {
		return nil, false
	}
	return o.ScanningEnabled, true
}

// HasScanningEnabled returns a boolean if a field has been set.
func (o *GetNetworkWirelessBluetoothSettings200Response) HasScanningEnabled() bool {
	if o != nil && !IsNil(o.ScanningEnabled) {
		return true
	}

	return false
}

// SetScanningEnabled gets a reference to the given bool and assigns it to the ScanningEnabled field.
func (o *GetNetworkWirelessBluetoothSettings200Response) SetScanningEnabled(v bool) {
	o.ScanningEnabled = &v
}

// GetAdvertisingEnabled returns the AdvertisingEnabled field value if set, zero value otherwise.
func (o *GetNetworkWirelessBluetoothSettings200Response) GetAdvertisingEnabled() bool {
	if o == nil || IsNil(o.AdvertisingEnabled) {
		var ret bool
		return ret
	}
	return *o.AdvertisingEnabled
}

// GetAdvertisingEnabledOk returns a tuple with the AdvertisingEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessBluetoothSettings200Response) GetAdvertisingEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.AdvertisingEnabled) {
		return nil, false
	}
	return o.AdvertisingEnabled, true
}

// HasAdvertisingEnabled returns a boolean if a field has been set.
func (o *GetNetworkWirelessBluetoothSettings200Response) HasAdvertisingEnabled() bool {
	if o != nil && !IsNil(o.AdvertisingEnabled) {
		return true
	}

	return false
}

// SetAdvertisingEnabled gets a reference to the given bool and assigns it to the AdvertisingEnabled field.
func (o *GetNetworkWirelessBluetoothSettings200Response) SetAdvertisingEnabled(v bool) {
	o.AdvertisingEnabled = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *GetNetworkWirelessBluetoothSettings200Response) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessBluetoothSettings200Response) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *GetNetworkWirelessBluetoothSettings200Response) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *GetNetworkWirelessBluetoothSettings200Response) SetUuid(v string) {
	o.Uuid = &v
}

// GetMajorMinorAssignmentMode returns the MajorMinorAssignmentMode field value if set, zero value otherwise.
func (o *GetNetworkWirelessBluetoothSettings200Response) GetMajorMinorAssignmentMode() string {
	if o == nil || IsNil(o.MajorMinorAssignmentMode) {
		var ret string
		return ret
	}
	return *o.MajorMinorAssignmentMode
}

// GetMajorMinorAssignmentModeOk returns a tuple with the MajorMinorAssignmentMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessBluetoothSettings200Response) GetMajorMinorAssignmentModeOk() (*string, bool) {
	if o == nil || IsNil(o.MajorMinorAssignmentMode) {
		return nil, false
	}
	return o.MajorMinorAssignmentMode, true
}

// HasMajorMinorAssignmentMode returns a boolean if a field has been set.
func (o *GetNetworkWirelessBluetoothSettings200Response) HasMajorMinorAssignmentMode() bool {
	if o != nil && !IsNil(o.MajorMinorAssignmentMode) {
		return true
	}

	return false
}

// SetMajorMinorAssignmentMode gets a reference to the given string and assigns it to the MajorMinorAssignmentMode field.
func (o *GetNetworkWirelessBluetoothSettings200Response) SetMajorMinorAssignmentMode(v string) {
	o.MajorMinorAssignmentMode = &v
}

// GetMajor returns the Major field value if set, zero value otherwise.
func (o *GetNetworkWirelessBluetoothSettings200Response) GetMajor() int32 {
	if o == nil || IsNil(o.Major) {
		var ret int32
		return ret
	}
	return *o.Major
}

// GetMajorOk returns a tuple with the Major field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessBluetoothSettings200Response) GetMajorOk() (*int32, bool) {
	if o == nil || IsNil(o.Major) {
		return nil, false
	}
	return o.Major, true
}

// HasMajor returns a boolean if a field has been set.
func (o *GetNetworkWirelessBluetoothSettings200Response) HasMajor() bool {
	if o != nil && !IsNil(o.Major) {
		return true
	}

	return false
}

// SetMajor gets a reference to the given int32 and assigns it to the Major field.
func (o *GetNetworkWirelessBluetoothSettings200Response) SetMajor(v int32) {
	o.Major = &v
}

// GetMinor returns the Minor field value if set, zero value otherwise.
func (o *GetNetworkWirelessBluetoothSettings200Response) GetMinor() int32 {
	if o == nil || IsNil(o.Minor) {
		var ret int32
		return ret
	}
	return *o.Minor
}

// GetMinorOk returns a tuple with the Minor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessBluetoothSettings200Response) GetMinorOk() (*int32, bool) {
	if o == nil || IsNil(o.Minor) {
		return nil, false
	}
	return o.Minor, true
}

// HasMinor returns a boolean if a field has been set.
func (o *GetNetworkWirelessBluetoothSettings200Response) HasMinor() bool {
	if o != nil && !IsNil(o.Minor) {
		return true
	}

	return false
}

// SetMinor gets a reference to the given int32 and assigns it to the Minor field.
func (o *GetNetworkWirelessBluetoothSettings200Response) SetMinor(v int32) {
	o.Minor = &v
}

// GetEslEnabled returns the EslEnabled field value if set, zero value otherwise.
func (o *GetNetworkWirelessBluetoothSettings200Response) GetEslEnabled() bool {
	if o == nil || IsNil(o.EslEnabled) {
		var ret bool
		return ret
	}
	return *o.EslEnabled
}

// GetEslEnabledOk returns a tuple with the EslEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessBluetoothSettings200Response) GetEslEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.EslEnabled) {
		return nil, false
	}
	return o.EslEnabled, true
}

// HasEslEnabled returns a boolean if a field has been set.
func (o *GetNetworkWirelessBluetoothSettings200Response) HasEslEnabled() bool {
	if o != nil && !IsNil(o.EslEnabled) {
		return true
	}

	return false
}

// SetEslEnabled gets a reference to the given bool and assigns it to the EslEnabled field.
func (o *GetNetworkWirelessBluetoothSettings200Response) SetEslEnabled(v bool) {
	o.EslEnabled = &v
}

func (o GetNetworkWirelessBluetoothSettings200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkWirelessBluetoothSettings200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ScanningEnabled) {
		toSerialize["scanningEnabled"] = o.ScanningEnabled
	}
	if !IsNil(o.AdvertisingEnabled) {
		toSerialize["advertisingEnabled"] = o.AdvertisingEnabled
	}
	if !IsNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}
	if !IsNil(o.MajorMinorAssignmentMode) {
		toSerialize["majorMinorAssignmentMode"] = o.MajorMinorAssignmentMode
	}
	if !IsNil(o.Major) {
		toSerialize["major"] = o.Major
	}
	if !IsNil(o.Minor) {
		toSerialize["minor"] = o.Minor
	}
	if !IsNil(o.EslEnabled) {
		toSerialize["eslEnabled"] = o.EslEnabled
	}
	return toSerialize, nil
}

type NullableGetNetworkWirelessBluetoothSettings200Response struct {
	value *GetNetworkWirelessBluetoothSettings200Response
	isSet bool
}

func (v NullableGetNetworkWirelessBluetoothSettings200Response) Get() *GetNetworkWirelessBluetoothSettings200Response {
	return v.value
}

func (v *NullableGetNetworkWirelessBluetoothSettings200Response) Set(val *GetNetworkWirelessBluetoothSettings200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkWirelessBluetoothSettings200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkWirelessBluetoothSettings200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkWirelessBluetoothSettings200Response(val *GetNetworkWirelessBluetoothSettings200Response) *NullableGetNetworkWirelessBluetoothSettings200Response {
	return &NullableGetNetworkWirelessBluetoothSettings200Response{value: val, isSet: true}
}

func (v NullableGetNetworkWirelessBluetoothSettings200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkWirelessBluetoothSettings200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

