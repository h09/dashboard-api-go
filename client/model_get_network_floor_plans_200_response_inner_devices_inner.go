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

// checks if the GetNetworkFloorPlans200ResponseInnerDevicesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkFloorPlans200ResponseInnerDevicesInner{}

// GetNetworkFloorPlans200ResponseInnerDevicesInner struct for GetNetworkFloorPlans200ResponseInnerDevicesInner
type GetNetworkFloorPlans200ResponseInnerDevicesInner struct {
	// Name of the device
	Name *string `json:"name,omitempty"`
	// Latitude of the device
	Lat *float32 `json:"lat,omitempty"`
	// Longitude of the device
	Lng *float32 `json:"lng,omitempty"`
	// Physical address of the device
	Address *string `json:"address,omitempty"`
	// Notes for the device, limited to 255 characters
	Notes *string `json:"notes,omitempty"`
	// List of tags assigned to the device
	Tags []string `json:"tags,omitempty"`
	// ID of the network the device belongs to
	NetworkId *string `json:"networkId,omitempty"`
	// Serial number of the device
	Serial *string `json:"serial,omitempty"`
	// Model of the device
	Model *string `json:"model,omitempty"`
	// MAC address of the device
	Mac *string `json:"mac,omitempty"`
	// LAN IP address of the device
	LanIp *string `json:"lanIp,omitempty"`
	// Firmware version of the device
	Firmware *string `json:"firmware,omitempty"`
	// Product type of the device
	ProductType *string `json:"productType,omitempty"`
	// Additional device information
	Details []GetNetworkFloorPlans200ResponseInnerDevicesInnerDetailsInner `json:"details,omitempty"`
}

// NewGetNetworkFloorPlans200ResponseInnerDevicesInner instantiates a new GetNetworkFloorPlans200ResponseInnerDevicesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkFloorPlans200ResponseInnerDevicesInner() *GetNetworkFloorPlans200ResponseInnerDevicesInner {
	this := GetNetworkFloorPlans200ResponseInnerDevicesInner{}
	return &this
}

// NewGetNetworkFloorPlans200ResponseInnerDevicesInnerWithDefaults instantiates a new GetNetworkFloorPlans200ResponseInnerDevicesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkFloorPlans200ResponseInnerDevicesInnerWithDefaults() *GetNetworkFloorPlans200ResponseInnerDevicesInner {
	this := GetNetworkFloorPlans200ResponseInnerDevicesInner{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetNetworkFloorPlans200ResponseInnerDevicesInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkFloorPlans200ResponseInnerDevicesInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetNetworkFloorPlans200ResponseInnerDevicesInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetNetworkFloorPlans200ResponseInnerDevicesInner) SetName(v string) {
	o.Name = &v
}

// GetLat returns the Lat field value if set, zero value otherwise.
func (o *GetNetworkFloorPlans200ResponseInnerDevicesInner) GetLat() float32 {
	if o == nil || IsNil(o.Lat) {
		var ret float32
		return ret
	}
	return *o.Lat
}

// GetLatOk returns a tuple with the Lat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkFloorPlans200ResponseInnerDevicesInner) GetLatOk() (*float32, bool) {
	if o == nil || IsNil(o.Lat) {
		return nil, false
	}
	return o.Lat, true
}

// HasLat returns a boolean if a field has been set.
func (o *GetNetworkFloorPlans200ResponseInnerDevicesInner) HasLat() bool {
	if o != nil && !IsNil(o.Lat) {
		return true
	}

	return false
}

// SetLat gets a reference to the given float32 and assigns it to the Lat field.
func (o *GetNetworkFloorPlans200ResponseInnerDevicesInner) SetLat(v float32) {
	o.Lat = &v
}

// GetLng returns the Lng field value if set, zero value otherwise.
func (o *GetNetworkFloorPlans200ResponseInnerDevicesInner) GetLng() float32 {
	if o == nil || IsNil(o.Lng) {
		var ret float32
		return ret
	}
	return *o.Lng
}

// GetLngOk returns a tuple with the Lng field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkFloorPlans200ResponseInnerDevicesInner) GetLngOk() (*float32, bool) {
	if o == nil || IsNil(o.Lng) {
		return nil, false
	}
	return o.Lng, true
}

// HasLng returns a boolean if a field has been set.
func (o *GetNetworkFloorPlans200ResponseInnerDevicesInner) HasLng() bool {
	if o != nil && !IsNil(o.Lng) {
		return true
	}

	return false
}

// SetLng gets a reference to the given float32 and assigns it to the Lng field.
func (o *GetNetworkFloorPlans200ResponseInnerDevicesInner) SetLng(v float32) {
	o.Lng = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *GetNetworkFloorPlans200ResponseInnerDevicesInner) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkFloorPlans200ResponseInnerDevicesInner) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *GetNetworkFloorPlans200ResponseInnerDevicesInner) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *GetNetworkFloorPlans200ResponseInnerDevicesInner) SetAddress(v string) {
	o.Address = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *GetNetworkFloorPlans200ResponseInnerDevicesInner) GetNotes() string {
	if o == nil || IsNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkFloorPlans200ResponseInnerDevicesInner) GetNotesOk() (*string, bool) {
	if o == nil || IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *GetNetworkFloorPlans200ResponseInnerDevicesInner) HasNotes() bool {
	if o != nil && !IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *GetNetworkFloorPlans200ResponseInnerDevicesInner) SetNotes(v string) {
	o.Notes = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *GetNetworkFloorPlans200ResponseInnerDevicesInner) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkFloorPlans200ResponseInnerDevicesInner) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *GetNetworkFloorPlans200ResponseInnerDevicesInner) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *GetNetworkFloorPlans200ResponseInnerDevicesInner) SetTags(v []string) {
	o.Tags = v
}

// GetNetworkId returns the NetworkId field value if set, zero value otherwise.
func (o *GetNetworkFloorPlans200ResponseInnerDevicesInner) GetNetworkId() string {
	if o == nil || IsNil(o.NetworkId) {
		var ret string
		return ret
	}
	return *o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkFloorPlans200ResponseInnerDevicesInner) GetNetworkIdOk() (*string, bool) {
	if o == nil || IsNil(o.NetworkId) {
		return nil, false
	}
	return o.NetworkId, true
}

// HasNetworkId returns a boolean if a field has been set.
func (o *GetNetworkFloorPlans200ResponseInnerDevicesInner) HasNetworkId() bool {
	if o != nil && !IsNil(o.NetworkId) {
		return true
	}

	return false
}

// SetNetworkId gets a reference to the given string and assigns it to the NetworkId field.
func (o *GetNetworkFloorPlans200ResponseInnerDevicesInner) SetNetworkId(v string) {
	o.NetworkId = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *GetNetworkFloorPlans200ResponseInnerDevicesInner) GetSerial() string {
	if o == nil || IsNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkFloorPlans200ResponseInnerDevicesInner) GetSerialOk() (*string, bool) {
	if o == nil || IsNil(o.Serial) {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *GetNetworkFloorPlans200ResponseInnerDevicesInner) HasSerial() bool {
	if o != nil && !IsNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *GetNetworkFloorPlans200ResponseInnerDevicesInner) SetSerial(v string) {
	o.Serial = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *GetNetworkFloorPlans200ResponseInnerDevicesInner) GetModel() string {
	if o == nil || IsNil(o.Model) {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkFloorPlans200ResponseInnerDevicesInner) GetModelOk() (*string, bool) {
	if o == nil || IsNil(o.Model) {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *GetNetworkFloorPlans200ResponseInnerDevicesInner) HasModel() bool {
	if o != nil && !IsNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *GetNetworkFloorPlans200ResponseInnerDevicesInner) SetModel(v string) {
	o.Model = &v
}

// GetMac returns the Mac field value if set, zero value otherwise.
func (o *GetNetworkFloorPlans200ResponseInnerDevicesInner) GetMac() string {
	if o == nil || IsNil(o.Mac) {
		var ret string
		return ret
	}
	return *o.Mac
}

// GetMacOk returns a tuple with the Mac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkFloorPlans200ResponseInnerDevicesInner) GetMacOk() (*string, bool) {
	if o == nil || IsNil(o.Mac) {
		return nil, false
	}
	return o.Mac, true
}

// HasMac returns a boolean if a field has been set.
func (o *GetNetworkFloorPlans200ResponseInnerDevicesInner) HasMac() bool {
	if o != nil && !IsNil(o.Mac) {
		return true
	}

	return false
}

// SetMac gets a reference to the given string and assigns it to the Mac field.
func (o *GetNetworkFloorPlans200ResponseInnerDevicesInner) SetMac(v string) {
	o.Mac = &v
}

// GetLanIp returns the LanIp field value if set, zero value otherwise.
func (o *GetNetworkFloorPlans200ResponseInnerDevicesInner) GetLanIp() string {
	if o == nil || IsNil(o.LanIp) {
		var ret string
		return ret
	}
	return *o.LanIp
}

// GetLanIpOk returns a tuple with the LanIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkFloorPlans200ResponseInnerDevicesInner) GetLanIpOk() (*string, bool) {
	if o == nil || IsNil(o.LanIp) {
		return nil, false
	}
	return o.LanIp, true
}

// HasLanIp returns a boolean if a field has been set.
func (o *GetNetworkFloorPlans200ResponseInnerDevicesInner) HasLanIp() bool {
	if o != nil && !IsNil(o.LanIp) {
		return true
	}

	return false
}

// SetLanIp gets a reference to the given string and assigns it to the LanIp field.
func (o *GetNetworkFloorPlans200ResponseInnerDevicesInner) SetLanIp(v string) {
	o.LanIp = &v
}

// GetFirmware returns the Firmware field value if set, zero value otherwise.
func (o *GetNetworkFloorPlans200ResponseInnerDevicesInner) GetFirmware() string {
	if o == nil || IsNil(o.Firmware) {
		var ret string
		return ret
	}
	return *o.Firmware
}

// GetFirmwareOk returns a tuple with the Firmware field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkFloorPlans200ResponseInnerDevicesInner) GetFirmwareOk() (*string, bool) {
	if o == nil || IsNil(o.Firmware) {
		return nil, false
	}
	return o.Firmware, true
}

// HasFirmware returns a boolean if a field has been set.
func (o *GetNetworkFloorPlans200ResponseInnerDevicesInner) HasFirmware() bool {
	if o != nil && !IsNil(o.Firmware) {
		return true
	}

	return false
}

// SetFirmware gets a reference to the given string and assigns it to the Firmware field.
func (o *GetNetworkFloorPlans200ResponseInnerDevicesInner) SetFirmware(v string) {
	o.Firmware = &v
}

// GetProductType returns the ProductType field value if set, zero value otherwise.
func (o *GetNetworkFloorPlans200ResponseInnerDevicesInner) GetProductType() string {
	if o == nil || IsNil(o.ProductType) {
		var ret string
		return ret
	}
	return *o.ProductType
}

// GetProductTypeOk returns a tuple with the ProductType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkFloorPlans200ResponseInnerDevicesInner) GetProductTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ProductType) {
		return nil, false
	}
	return o.ProductType, true
}

// HasProductType returns a boolean if a field has been set.
func (o *GetNetworkFloorPlans200ResponseInnerDevicesInner) HasProductType() bool {
	if o != nil && !IsNil(o.ProductType) {
		return true
	}

	return false
}

// SetProductType gets a reference to the given string and assigns it to the ProductType field.
func (o *GetNetworkFloorPlans200ResponseInnerDevicesInner) SetProductType(v string) {
	o.ProductType = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *GetNetworkFloorPlans200ResponseInnerDevicesInner) GetDetails() []GetNetworkFloorPlans200ResponseInnerDevicesInnerDetailsInner {
	if o == nil || IsNil(o.Details) {
		var ret []GetNetworkFloorPlans200ResponseInnerDevicesInnerDetailsInner
		return ret
	}
	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkFloorPlans200ResponseInnerDevicesInner) GetDetailsOk() ([]GetNetworkFloorPlans200ResponseInnerDevicesInnerDetailsInner, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *GetNetworkFloorPlans200ResponseInnerDevicesInner) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given []GetNetworkFloorPlans200ResponseInnerDevicesInnerDetailsInner and assigns it to the Details field.
func (o *GetNetworkFloorPlans200ResponseInnerDevicesInner) SetDetails(v []GetNetworkFloorPlans200ResponseInnerDevicesInnerDetailsInner) {
	o.Details = v
}

func (o GetNetworkFloorPlans200ResponseInnerDevicesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkFloorPlans200ResponseInnerDevicesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Lat) {
		toSerialize["lat"] = o.Lat
	}
	if !IsNil(o.Lng) {
		toSerialize["lng"] = o.Lng
	}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.NetworkId) {
		toSerialize["networkId"] = o.NetworkId
	}
	if !IsNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !IsNil(o.Model) {
		toSerialize["model"] = o.Model
	}
	if !IsNil(o.Mac) {
		toSerialize["mac"] = o.Mac
	}
	if !IsNil(o.LanIp) {
		toSerialize["lanIp"] = o.LanIp
	}
	if !IsNil(o.Firmware) {
		toSerialize["firmware"] = o.Firmware
	}
	if !IsNil(o.ProductType) {
		toSerialize["productType"] = o.ProductType
	}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	return toSerialize, nil
}

type NullableGetNetworkFloorPlans200ResponseInnerDevicesInner struct {
	value *GetNetworkFloorPlans200ResponseInnerDevicesInner
	isSet bool
}

func (v NullableGetNetworkFloorPlans200ResponseInnerDevicesInner) Get() *GetNetworkFloorPlans200ResponseInnerDevicesInner {
	return v.value
}

func (v *NullableGetNetworkFloorPlans200ResponseInnerDevicesInner) Set(val *GetNetworkFloorPlans200ResponseInnerDevicesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkFloorPlans200ResponseInnerDevicesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkFloorPlans200ResponseInnerDevicesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkFloorPlans200ResponseInnerDevicesInner(val *GetNetworkFloorPlans200ResponseInnerDevicesInner) *NullableGetNetworkFloorPlans200ResponseInnerDevicesInner {
	return &NullableGetNetworkFloorPlans200ResponseInnerDevicesInner{value: val, isSet: true}
}

func (v NullableGetNetworkFloorPlans200ResponseInnerDevicesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkFloorPlans200ResponseInnerDevicesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


