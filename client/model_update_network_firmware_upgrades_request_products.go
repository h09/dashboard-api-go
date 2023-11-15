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

// checks if the UpdateNetworkFirmwareUpgradesRequestProducts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkFirmwareUpgradesRequestProducts{}

// UpdateNetworkFirmwareUpgradesRequestProducts Contains information about the network to update
type UpdateNetworkFirmwareUpgradesRequestProducts struct {
	Wireless *UpdateNetworkFirmwareUpgradesRequestProductsWireless `json:"wireless,omitempty"`
	Appliance *UpdateNetworkFirmwareUpgradesRequestProductsWireless `json:"appliance,omitempty"`
	Switch *UpdateNetworkFirmwareUpgradesRequestProductsWireless `json:"switch,omitempty"`
	Camera *UpdateNetworkFirmwareUpgradesRequestProductsWireless `json:"camera,omitempty"`
	CellularGateway *UpdateNetworkFirmwareUpgradesRequestProductsWireless `json:"cellularGateway,omitempty"`
	Sensor *UpdateNetworkFirmwareUpgradesRequestProductsWireless `json:"sensor,omitempty"`
	SwitchCatalyst *UpdateNetworkFirmwareUpgradesRequestProductsWireless `json:"switchCatalyst,omitempty"`
}

// NewUpdateNetworkFirmwareUpgradesRequestProducts instantiates a new UpdateNetworkFirmwareUpgradesRequestProducts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkFirmwareUpgradesRequestProducts() *UpdateNetworkFirmwareUpgradesRequestProducts {
	this := UpdateNetworkFirmwareUpgradesRequestProducts{}
	return &this
}

// NewUpdateNetworkFirmwareUpgradesRequestProductsWithDefaults instantiates a new UpdateNetworkFirmwareUpgradesRequestProducts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkFirmwareUpgradesRequestProductsWithDefaults() *UpdateNetworkFirmwareUpgradesRequestProducts {
	this := UpdateNetworkFirmwareUpgradesRequestProducts{}
	return &this
}

// GetWireless returns the Wireless field value if set, zero value otherwise.
func (o *UpdateNetworkFirmwareUpgradesRequestProducts) GetWireless() UpdateNetworkFirmwareUpgradesRequestProductsWireless {
	if o == nil || IsNil(o.Wireless) {
		var ret UpdateNetworkFirmwareUpgradesRequestProductsWireless
		return ret
	}
	return *o.Wireless
}

// GetWirelessOk returns a tuple with the Wireless field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkFirmwareUpgradesRequestProducts) GetWirelessOk() (*UpdateNetworkFirmwareUpgradesRequestProductsWireless, bool) {
	if o == nil || IsNil(o.Wireless) {
		return nil, false
	}
	return o.Wireless, true
}

// HasWireless returns a boolean if a field has been set.
func (o *UpdateNetworkFirmwareUpgradesRequestProducts) HasWireless() bool {
	if o != nil && !IsNil(o.Wireless) {
		return true
	}

	return false
}

// SetWireless gets a reference to the given UpdateNetworkFirmwareUpgradesRequestProductsWireless and assigns it to the Wireless field.
func (o *UpdateNetworkFirmwareUpgradesRequestProducts) SetWireless(v UpdateNetworkFirmwareUpgradesRequestProductsWireless) {
	o.Wireless = &v
}

// GetAppliance returns the Appliance field value if set, zero value otherwise.
func (o *UpdateNetworkFirmwareUpgradesRequestProducts) GetAppliance() UpdateNetworkFirmwareUpgradesRequestProductsWireless {
	if o == nil || IsNil(o.Appliance) {
		var ret UpdateNetworkFirmwareUpgradesRequestProductsWireless
		return ret
	}
	return *o.Appliance
}

// GetApplianceOk returns a tuple with the Appliance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkFirmwareUpgradesRequestProducts) GetApplianceOk() (*UpdateNetworkFirmwareUpgradesRequestProductsWireless, bool) {
	if o == nil || IsNil(o.Appliance) {
		return nil, false
	}
	return o.Appliance, true
}

// HasAppliance returns a boolean if a field has been set.
func (o *UpdateNetworkFirmwareUpgradesRequestProducts) HasAppliance() bool {
	if o != nil && !IsNil(o.Appliance) {
		return true
	}

	return false
}

// SetAppliance gets a reference to the given UpdateNetworkFirmwareUpgradesRequestProductsWireless and assigns it to the Appliance field.
func (o *UpdateNetworkFirmwareUpgradesRequestProducts) SetAppliance(v UpdateNetworkFirmwareUpgradesRequestProductsWireless) {
	o.Appliance = &v
}

// GetSwitch returns the Switch field value if set, zero value otherwise.
func (o *UpdateNetworkFirmwareUpgradesRequestProducts) GetSwitch() UpdateNetworkFirmwareUpgradesRequestProductsWireless {
	if o == nil || IsNil(o.Switch) {
		var ret UpdateNetworkFirmwareUpgradesRequestProductsWireless
		return ret
	}
	return *o.Switch
}

// GetSwitchOk returns a tuple with the Switch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkFirmwareUpgradesRequestProducts) GetSwitchOk() (*UpdateNetworkFirmwareUpgradesRequestProductsWireless, bool) {
	if o == nil || IsNil(o.Switch) {
		return nil, false
	}
	return o.Switch, true
}

// HasSwitch returns a boolean if a field has been set.
func (o *UpdateNetworkFirmwareUpgradesRequestProducts) HasSwitch() bool {
	if o != nil && !IsNil(o.Switch) {
		return true
	}

	return false
}

// SetSwitch gets a reference to the given UpdateNetworkFirmwareUpgradesRequestProductsWireless and assigns it to the Switch field.
func (o *UpdateNetworkFirmwareUpgradesRequestProducts) SetSwitch(v UpdateNetworkFirmwareUpgradesRequestProductsWireless) {
	o.Switch = &v
}

// GetCamera returns the Camera field value if set, zero value otherwise.
func (o *UpdateNetworkFirmwareUpgradesRequestProducts) GetCamera() UpdateNetworkFirmwareUpgradesRequestProductsWireless {
	if o == nil || IsNil(o.Camera) {
		var ret UpdateNetworkFirmwareUpgradesRequestProductsWireless
		return ret
	}
	return *o.Camera
}

// GetCameraOk returns a tuple with the Camera field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkFirmwareUpgradesRequestProducts) GetCameraOk() (*UpdateNetworkFirmwareUpgradesRequestProductsWireless, bool) {
	if o == nil || IsNil(o.Camera) {
		return nil, false
	}
	return o.Camera, true
}

// HasCamera returns a boolean if a field has been set.
func (o *UpdateNetworkFirmwareUpgradesRequestProducts) HasCamera() bool {
	if o != nil && !IsNil(o.Camera) {
		return true
	}

	return false
}

// SetCamera gets a reference to the given UpdateNetworkFirmwareUpgradesRequestProductsWireless and assigns it to the Camera field.
func (o *UpdateNetworkFirmwareUpgradesRequestProducts) SetCamera(v UpdateNetworkFirmwareUpgradesRequestProductsWireless) {
	o.Camera = &v
}

// GetCellularGateway returns the CellularGateway field value if set, zero value otherwise.
func (o *UpdateNetworkFirmwareUpgradesRequestProducts) GetCellularGateway() UpdateNetworkFirmwareUpgradesRequestProductsWireless {
	if o == nil || IsNil(o.CellularGateway) {
		var ret UpdateNetworkFirmwareUpgradesRequestProductsWireless
		return ret
	}
	return *o.CellularGateway
}

// GetCellularGatewayOk returns a tuple with the CellularGateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkFirmwareUpgradesRequestProducts) GetCellularGatewayOk() (*UpdateNetworkFirmwareUpgradesRequestProductsWireless, bool) {
	if o == nil || IsNil(o.CellularGateway) {
		return nil, false
	}
	return o.CellularGateway, true
}

// HasCellularGateway returns a boolean if a field has been set.
func (o *UpdateNetworkFirmwareUpgradesRequestProducts) HasCellularGateway() bool {
	if o != nil && !IsNil(o.CellularGateway) {
		return true
	}

	return false
}

// SetCellularGateway gets a reference to the given UpdateNetworkFirmwareUpgradesRequestProductsWireless and assigns it to the CellularGateway field.
func (o *UpdateNetworkFirmwareUpgradesRequestProducts) SetCellularGateway(v UpdateNetworkFirmwareUpgradesRequestProductsWireless) {
	o.CellularGateway = &v
}

// GetSensor returns the Sensor field value if set, zero value otherwise.
func (o *UpdateNetworkFirmwareUpgradesRequestProducts) GetSensor() UpdateNetworkFirmwareUpgradesRequestProductsWireless {
	if o == nil || IsNil(o.Sensor) {
		var ret UpdateNetworkFirmwareUpgradesRequestProductsWireless
		return ret
	}
	return *o.Sensor
}

// GetSensorOk returns a tuple with the Sensor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkFirmwareUpgradesRequestProducts) GetSensorOk() (*UpdateNetworkFirmwareUpgradesRequestProductsWireless, bool) {
	if o == nil || IsNil(o.Sensor) {
		return nil, false
	}
	return o.Sensor, true
}

// HasSensor returns a boolean if a field has been set.
func (o *UpdateNetworkFirmwareUpgradesRequestProducts) HasSensor() bool {
	if o != nil && !IsNil(o.Sensor) {
		return true
	}

	return false
}

// SetSensor gets a reference to the given UpdateNetworkFirmwareUpgradesRequestProductsWireless and assigns it to the Sensor field.
func (o *UpdateNetworkFirmwareUpgradesRequestProducts) SetSensor(v UpdateNetworkFirmwareUpgradesRequestProductsWireless) {
	o.Sensor = &v
}

// GetSwitchCatalyst returns the SwitchCatalyst field value if set, zero value otherwise.
func (o *UpdateNetworkFirmwareUpgradesRequestProducts) GetSwitchCatalyst() UpdateNetworkFirmwareUpgradesRequestProductsWireless {
	if o == nil || IsNil(o.SwitchCatalyst) {
		var ret UpdateNetworkFirmwareUpgradesRequestProductsWireless
		return ret
	}
	return *o.SwitchCatalyst
}

// GetSwitchCatalystOk returns a tuple with the SwitchCatalyst field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkFirmwareUpgradesRequestProducts) GetSwitchCatalystOk() (*UpdateNetworkFirmwareUpgradesRequestProductsWireless, bool) {
	if o == nil || IsNil(o.SwitchCatalyst) {
		return nil, false
	}
	return o.SwitchCatalyst, true
}

// HasSwitchCatalyst returns a boolean if a field has been set.
func (o *UpdateNetworkFirmwareUpgradesRequestProducts) HasSwitchCatalyst() bool {
	if o != nil && !IsNil(o.SwitchCatalyst) {
		return true
	}

	return false
}

// SetSwitchCatalyst gets a reference to the given UpdateNetworkFirmwareUpgradesRequestProductsWireless and assigns it to the SwitchCatalyst field.
func (o *UpdateNetworkFirmwareUpgradesRequestProducts) SetSwitchCatalyst(v UpdateNetworkFirmwareUpgradesRequestProductsWireless) {
	o.SwitchCatalyst = &v
}

func (o UpdateNetworkFirmwareUpgradesRequestProducts) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkFirmwareUpgradesRequestProducts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Wireless) {
		toSerialize["wireless"] = o.Wireless
	}
	if !IsNil(o.Appliance) {
		toSerialize["appliance"] = o.Appliance
	}
	if !IsNil(o.Switch) {
		toSerialize["switch"] = o.Switch
	}
	if !IsNil(o.Camera) {
		toSerialize["camera"] = o.Camera
	}
	if !IsNil(o.CellularGateway) {
		toSerialize["cellularGateway"] = o.CellularGateway
	}
	if !IsNil(o.Sensor) {
		toSerialize["sensor"] = o.Sensor
	}
	if !IsNil(o.SwitchCatalyst) {
		toSerialize["switchCatalyst"] = o.SwitchCatalyst
	}
	return toSerialize, nil
}

type NullableUpdateNetworkFirmwareUpgradesRequestProducts struct {
	value *UpdateNetworkFirmwareUpgradesRequestProducts
	isSet bool
}

func (v NullableUpdateNetworkFirmwareUpgradesRequestProducts) Get() *UpdateNetworkFirmwareUpgradesRequestProducts {
	return v.value
}

func (v *NullableUpdateNetworkFirmwareUpgradesRequestProducts) Set(val *UpdateNetworkFirmwareUpgradesRequestProducts) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkFirmwareUpgradesRequestProducts) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkFirmwareUpgradesRequestProducts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkFirmwareUpgradesRequestProducts(val *UpdateNetworkFirmwareUpgradesRequestProducts) *NullableUpdateNetworkFirmwareUpgradesRequestProducts {
	return &NullableUpdateNetworkFirmwareUpgradesRequestProducts{value: val, isSet: true}
}

func (v NullableUpdateNetworkFirmwareUpgradesRequestProducts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkFirmwareUpgradesRequestProducts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


