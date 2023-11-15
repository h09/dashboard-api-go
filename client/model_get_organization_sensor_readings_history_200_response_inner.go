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

// checks if the GetOrganizationSensorReadingsHistory200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationSensorReadingsHistory200ResponseInner{}

// GetOrganizationSensorReadingsHistory200ResponseInner struct for GetOrganizationSensorReadingsHistory200ResponseInner
type GetOrganizationSensorReadingsHistory200ResponseInner struct {
	// Serial number of the sensor that took the reading.
	Serial *string `json:"serial,omitempty"`
	Network *GetOrganizationSensorReadingsHistory200ResponseInnerNetwork `json:"network,omitempty"`
	// Time at which the reading occurred, in ISO8601 format.
	Ts *string `json:"ts,omitempty"`
	// Type of sensor reading.
	Metric *string `json:"metric,omitempty"`
	Battery *GetOrganizationSensorReadingsHistory200ResponseInnerBattery `json:"battery,omitempty"`
	Button *GetOrganizationSensorReadingsHistory200ResponseInnerButton `json:"button,omitempty"`
	Door *GetOrganizationSensorReadingsHistory200ResponseInnerDoor `json:"door,omitempty"`
	Humidity *GetOrganizationSensorReadingsHistory200ResponseInnerHumidity `json:"humidity,omitempty"`
	IndoorAirQuality *GetOrganizationSensorReadingsHistory200ResponseInnerIndoorAirQuality `json:"indoorAirQuality,omitempty"`
	Noise *GetOrganizationSensorReadingsHistory200ResponseInnerNoise `json:"noise,omitempty"`
	Pm25 *GetOrganizationSensorReadingsHistory200ResponseInnerPm25 `json:"pm25,omitempty"`
	Temperature *GetOrganizationSensorReadingsHistory200ResponseInnerTemperature `json:"temperature,omitempty"`
	Tvoc *GetOrganizationSensorReadingsHistory200ResponseInnerTvoc `json:"tvoc,omitempty"`
	Water *GetOrganizationSensorReadingsHistory200ResponseInnerWater `json:"water,omitempty"`
}

// NewGetOrganizationSensorReadingsHistory200ResponseInner instantiates a new GetOrganizationSensorReadingsHistory200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationSensorReadingsHistory200ResponseInner() *GetOrganizationSensorReadingsHistory200ResponseInner {
	this := GetOrganizationSensorReadingsHistory200ResponseInner{}
	return &this
}

// NewGetOrganizationSensorReadingsHistory200ResponseInnerWithDefaults instantiates a new GetOrganizationSensorReadingsHistory200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationSensorReadingsHistory200ResponseInnerWithDefaults() *GetOrganizationSensorReadingsHistory200ResponseInner {
	this := GetOrganizationSensorReadingsHistory200ResponseInner{}
	return &this
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *GetOrganizationSensorReadingsHistory200ResponseInner) GetSerial() string {
	if o == nil || IsNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSensorReadingsHistory200ResponseInner) GetSerialOk() (*string, bool) {
	if o == nil || IsNil(o.Serial) {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *GetOrganizationSensorReadingsHistory200ResponseInner) HasSerial() bool {
	if o != nil && !IsNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *GetOrganizationSensorReadingsHistory200ResponseInner) SetSerial(v string) {
	o.Serial = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *GetOrganizationSensorReadingsHistory200ResponseInner) GetNetwork() GetOrganizationSensorReadingsHistory200ResponseInnerNetwork {
	if o == nil || IsNil(o.Network) {
		var ret GetOrganizationSensorReadingsHistory200ResponseInnerNetwork
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSensorReadingsHistory200ResponseInner) GetNetworkOk() (*GetOrganizationSensorReadingsHistory200ResponseInnerNetwork, bool) {
	if o == nil || IsNil(o.Network) {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *GetOrganizationSensorReadingsHistory200ResponseInner) HasNetwork() bool {
	if o != nil && !IsNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given GetOrganizationSensorReadingsHistory200ResponseInnerNetwork and assigns it to the Network field.
func (o *GetOrganizationSensorReadingsHistory200ResponseInner) SetNetwork(v GetOrganizationSensorReadingsHistory200ResponseInnerNetwork) {
	o.Network = &v
}

// GetTs returns the Ts field value if set, zero value otherwise.
func (o *GetOrganizationSensorReadingsHistory200ResponseInner) GetTs() string {
	if o == nil || IsNil(o.Ts) {
		var ret string
		return ret
	}
	return *o.Ts
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSensorReadingsHistory200ResponseInner) GetTsOk() (*string, bool) {
	if o == nil || IsNil(o.Ts) {
		return nil, false
	}
	return o.Ts, true
}

// HasTs returns a boolean if a field has been set.
func (o *GetOrganizationSensorReadingsHistory200ResponseInner) HasTs() bool {
	if o != nil && !IsNil(o.Ts) {
		return true
	}

	return false
}

// SetTs gets a reference to the given string and assigns it to the Ts field.
func (o *GetOrganizationSensorReadingsHistory200ResponseInner) SetTs(v string) {
	o.Ts = &v
}

// GetMetric returns the Metric field value if set, zero value otherwise.
func (o *GetOrganizationSensorReadingsHistory200ResponseInner) GetMetric() string {
	if o == nil || IsNil(o.Metric) {
		var ret string
		return ret
	}
	return *o.Metric
}

// GetMetricOk returns a tuple with the Metric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSensorReadingsHistory200ResponseInner) GetMetricOk() (*string, bool) {
	if o == nil || IsNil(o.Metric) {
		return nil, false
	}
	return o.Metric, true
}

// HasMetric returns a boolean if a field has been set.
func (o *GetOrganizationSensorReadingsHistory200ResponseInner) HasMetric() bool {
	if o != nil && !IsNil(o.Metric) {
		return true
	}

	return false
}

// SetMetric gets a reference to the given string and assigns it to the Metric field.
func (o *GetOrganizationSensorReadingsHistory200ResponseInner) SetMetric(v string) {
	o.Metric = &v
}

// GetBattery returns the Battery field value if set, zero value otherwise.
func (o *GetOrganizationSensorReadingsHistory200ResponseInner) GetBattery() GetOrganizationSensorReadingsHistory200ResponseInnerBattery {
	if o == nil || IsNil(o.Battery) {
		var ret GetOrganizationSensorReadingsHistory200ResponseInnerBattery
		return ret
	}
	return *o.Battery
}

// GetBatteryOk returns a tuple with the Battery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSensorReadingsHistory200ResponseInner) GetBatteryOk() (*GetOrganizationSensorReadingsHistory200ResponseInnerBattery, bool) {
	if o == nil || IsNil(o.Battery) {
		return nil, false
	}
	return o.Battery, true
}

// HasBattery returns a boolean if a field has been set.
func (o *GetOrganizationSensorReadingsHistory200ResponseInner) HasBattery() bool {
	if o != nil && !IsNil(o.Battery) {
		return true
	}

	return false
}

// SetBattery gets a reference to the given GetOrganizationSensorReadingsHistory200ResponseInnerBattery and assigns it to the Battery field.
func (o *GetOrganizationSensorReadingsHistory200ResponseInner) SetBattery(v GetOrganizationSensorReadingsHistory200ResponseInnerBattery) {
	o.Battery = &v
}

// GetButton returns the Button field value if set, zero value otherwise.
func (o *GetOrganizationSensorReadingsHistory200ResponseInner) GetButton() GetOrganizationSensorReadingsHistory200ResponseInnerButton {
	if o == nil || IsNil(o.Button) {
		var ret GetOrganizationSensorReadingsHistory200ResponseInnerButton
		return ret
	}
	return *o.Button
}

// GetButtonOk returns a tuple with the Button field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSensorReadingsHistory200ResponseInner) GetButtonOk() (*GetOrganizationSensorReadingsHistory200ResponseInnerButton, bool) {
	if o == nil || IsNil(o.Button) {
		return nil, false
	}
	return o.Button, true
}

// HasButton returns a boolean if a field has been set.
func (o *GetOrganizationSensorReadingsHistory200ResponseInner) HasButton() bool {
	if o != nil && !IsNil(o.Button) {
		return true
	}

	return false
}

// SetButton gets a reference to the given GetOrganizationSensorReadingsHistory200ResponseInnerButton and assigns it to the Button field.
func (o *GetOrganizationSensorReadingsHistory200ResponseInner) SetButton(v GetOrganizationSensorReadingsHistory200ResponseInnerButton) {
	o.Button = &v
}

// GetDoor returns the Door field value if set, zero value otherwise.
func (o *GetOrganizationSensorReadingsHistory200ResponseInner) GetDoor() GetOrganizationSensorReadingsHistory200ResponseInnerDoor {
	if o == nil || IsNil(o.Door) {
		var ret GetOrganizationSensorReadingsHistory200ResponseInnerDoor
		return ret
	}
	return *o.Door
}

// GetDoorOk returns a tuple with the Door field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSensorReadingsHistory200ResponseInner) GetDoorOk() (*GetOrganizationSensorReadingsHistory200ResponseInnerDoor, bool) {
	if o == nil || IsNil(o.Door) {
		return nil, false
	}
	return o.Door, true
}

// HasDoor returns a boolean if a field has been set.
func (o *GetOrganizationSensorReadingsHistory200ResponseInner) HasDoor() bool {
	if o != nil && !IsNil(o.Door) {
		return true
	}

	return false
}

// SetDoor gets a reference to the given GetOrganizationSensorReadingsHistory200ResponseInnerDoor and assigns it to the Door field.
func (o *GetOrganizationSensorReadingsHistory200ResponseInner) SetDoor(v GetOrganizationSensorReadingsHistory200ResponseInnerDoor) {
	o.Door = &v
}

// GetHumidity returns the Humidity field value if set, zero value otherwise.
func (o *GetOrganizationSensorReadingsHistory200ResponseInner) GetHumidity() GetOrganizationSensorReadingsHistory200ResponseInnerHumidity {
	if o == nil || IsNil(o.Humidity) {
		var ret GetOrganizationSensorReadingsHistory200ResponseInnerHumidity
		return ret
	}
	return *o.Humidity
}

// GetHumidityOk returns a tuple with the Humidity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSensorReadingsHistory200ResponseInner) GetHumidityOk() (*GetOrganizationSensorReadingsHistory200ResponseInnerHumidity, bool) {
	if o == nil || IsNil(o.Humidity) {
		return nil, false
	}
	return o.Humidity, true
}

// HasHumidity returns a boolean if a field has been set.
func (o *GetOrganizationSensorReadingsHistory200ResponseInner) HasHumidity() bool {
	if o != nil && !IsNil(o.Humidity) {
		return true
	}

	return false
}

// SetHumidity gets a reference to the given GetOrganizationSensorReadingsHistory200ResponseInnerHumidity and assigns it to the Humidity field.
func (o *GetOrganizationSensorReadingsHistory200ResponseInner) SetHumidity(v GetOrganizationSensorReadingsHistory200ResponseInnerHumidity) {
	o.Humidity = &v
}

// GetIndoorAirQuality returns the IndoorAirQuality field value if set, zero value otherwise.
func (o *GetOrganizationSensorReadingsHistory200ResponseInner) GetIndoorAirQuality() GetOrganizationSensorReadingsHistory200ResponseInnerIndoorAirQuality {
	if o == nil || IsNil(o.IndoorAirQuality) {
		var ret GetOrganizationSensorReadingsHistory200ResponseInnerIndoorAirQuality
		return ret
	}
	return *o.IndoorAirQuality
}

// GetIndoorAirQualityOk returns a tuple with the IndoorAirQuality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSensorReadingsHistory200ResponseInner) GetIndoorAirQualityOk() (*GetOrganizationSensorReadingsHistory200ResponseInnerIndoorAirQuality, bool) {
	if o == nil || IsNil(o.IndoorAirQuality) {
		return nil, false
	}
	return o.IndoorAirQuality, true
}

// HasIndoorAirQuality returns a boolean if a field has been set.
func (o *GetOrganizationSensorReadingsHistory200ResponseInner) HasIndoorAirQuality() bool {
	if o != nil && !IsNil(o.IndoorAirQuality) {
		return true
	}

	return false
}

// SetIndoorAirQuality gets a reference to the given GetOrganizationSensorReadingsHistory200ResponseInnerIndoorAirQuality and assigns it to the IndoorAirQuality field.
func (o *GetOrganizationSensorReadingsHistory200ResponseInner) SetIndoorAirQuality(v GetOrganizationSensorReadingsHistory200ResponseInnerIndoorAirQuality) {
	o.IndoorAirQuality = &v
}

// GetNoise returns the Noise field value if set, zero value otherwise.
func (o *GetOrganizationSensorReadingsHistory200ResponseInner) GetNoise() GetOrganizationSensorReadingsHistory200ResponseInnerNoise {
	if o == nil || IsNil(o.Noise) {
		var ret GetOrganizationSensorReadingsHistory200ResponseInnerNoise
		return ret
	}
	return *o.Noise
}

// GetNoiseOk returns a tuple with the Noise field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSensorReadingsHistory200ResponseInner) GetNoiseOk() (*GetOrganizationSensorReadingsHistory200ResponseInnerNoise, bool) {
	if o == nil || IsNil(o.Noise) {
		return nil, false
	}
	return o.Noise, true
}

// HasNoise returns a boolean if a field has been set.
func (o *GetOrganizationSensorReadingsHistory200ResponseInner) HasNoise() bool {
	if o != nil && !IsNil(o.Noise) {
		return true
	}

	return false
}

// SetNoise gets a reference to the given GetOrganizationSensorReadingsHistory200ResponseInnerNoise and assigns it to the Noise field.
func (o *GetOrganizationSensorReadingsHistory200ResponseInner) SetNoise(v GetOrganizationSensorReadingsHistory200ResponseInnerNoise) {
	o.Noise = &v
}

// GetPm25 returns the Pm25 field value if set, zero value otherwise.
func (o *GetOrganizationSensorReadingsHistory200ResponseInner) GetPm25() GetOrganizationSensorReadingsHistory200ResponseInnerPm25 {
	if o == nil || IsNil(o.Pm25) {
		var ret GetOrganizationSensorReadingsHistory200ResponseInnerPm25
		return ret
	}
	return *o.Pm25
}

// GetPm25Ok returns a tuple with the Pm25 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSensorReadingsHistory200ResponseInner) GetPm25Ok() (*GetOrganizationSensorReadingsHistory200ResponseInnerPm25, bool) {
	if o == nil || IsNil(o.Pm25) {
		return nil, false
	}
	return o.Pm25, true
}

// HasPm25 returns a boolean if a field has been set.
func (o *GetOrganizationSensorReadingsHistory200ResponseInner) HasPm25() bool {
	if o != nil && !IsNil(o.Pm25) {
		return true
	}

	return false
}

// SetPm25 gets a reference to the given GetOrganizationSensorReadingsHistory200ResponseInnerPm25 and assigns it to the Pm25 field.
func (o *GetOrganizationSensorReadingsHistory200ResponseInner) SetPm25(v GetOrganizationSensorReadingsHistory200ResponseInnerPm25) {
	o.Pm25 = &v
}

// GetTemperature returns the Temperature field value if set, zero value otherwise.
func (o *GetOrganizationSensorReadingsHistory200ResponseInner) GetTemperature() GetOrganizationSensorReadingsHistory200ResponseInnerTemperature {
	if o == nil || IsNil(o.Temperature) {
		var ret GetOrganizationSensorReadingsHistory200ResponseInnerTemperature
		return ret
	}
	return *o.Temperature
}

// GetTemperatureOk returns a tuple with the Temperature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSensorReadingsHistory200ResponseInner) GetTemperatureOk() (*GetOrganizationSensorReadingsHistory200ResponseInnerTemperature, bool) {
	if o == nil || IsNil(o.Temperature) {
		return nil, false
	}
	return o.Temperature, true
}

// HasTemperature returns a boolean if a field has been set.
func (o *GetOrganizationSensorReadingsHistory200ResponseInner) HasTemperature() bool {
	if o != nil && !IsNil(o.Temperature) {
		return true
	}

	return false
}

// SetTemperature gets a reference to the given GetOrganizationSensorReadingsHistory200ResponseInnerTemperature and assigns it to the Temperature field.
func (o *GetOrganizationSensorReadingsHistory200ResponseInner) SetTemperature(v GetOrganizationSensorReadingsHistory200ResponseInnerTemperature) {
	o.Temperature = &v
}

// GetTvoc returns the Tvoc field value if set, zero value otherwise.
func (o *GetOrganizationSensorReadingsHistory200ResponseInner) GetTvoc() GetOrganizationSensorReadingsHistory200ResponseInnerTvoc {
	if o == nil || IsNil(o.Tvoc) {
		var ret GetOrganizationSensorReadingsHistory200ResponseInnerTvoc
		return ret
	}
	return *o.Tvoc
}

// GetTvocOk returns a tuple with the Tvoc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSensorReadingsHistory200ResponseInner) GetTvocOk() (*GetOrganizationSensorReadingsHistory200ResponseInnerTvoc, bool) {
	if o == nil || IsNil(o.Tvoc) {
		return nil, false
	}
	return o.Tvoc, true
}

// HasTvoc returns a boolean if a field has been set.
func (o *GetOrganizationSensorReadingsHistory200ResponseInner) HasTvoc() bool {
	if o != nil && !IsNil(o.Tvoc) {
		return true
	}

	return false
}

// SetTvoc gets a reference to the given GetOrganizationSensorReadingsHistory200ResponseInnerTvoc and assigns it to the Tvoc field.
func (o *GetOrganizationSensorReadingsHistory200ResponseInner) SetTvoc(v GetOrganizationSensorReadingsHistory200ResponseInnerTvoc) {
	o.Tvoc = &v
}

// GetWater returns the Water field value if set, zero value otherwise.
func (o *GetOrganizationSensorReadingsHistory200ResponseInner) GetWater() GetOrganizationSensorReadingsHistory200ResponseInnerWater {
	if o == nil || IsNil(o.Water) {
		var ret GetOrganizationSensorReadingsHistory200ResponseInnerWater
		return ret
	}
	return *o.Water
}

// GetWaterOk returns a tuple with the Water field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSensorReadingsHistory200ResponseInner) GetWaterOk() (*GetOrganizationSensorReadingsHistory200ResponseInnerWater, bool) {
	if o == nil || IsNil(o.Water) {
		return nil, false
	}
	return o.Water, true
}

// HasWater returns a boolean if a field has been set.
func (o *GetOrganizationSensorReadingsHistory200ResponseInner) HasWater() bool {
	if o != nil && !IsNil(o.Water) {
		return true
	}

	return false
}

// SetWater gets a reference to the given GetOrganizationSensorReadingsHistory200ResponseInnerWater and assigns it to the Water field.
func (o *GetOrganizationSensorReadingsHistory200ResponseInner) SetWater(v GetOrganizationSensorReadingsHistory200ResponseInnerWater) {
	o.Water = &v
}

func (o GetOrganizationSensorReadingsHistory200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationSensorReadingsHistory200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !IsNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !IsNil(o.Ts) {
		toSerialize["ts"] = o.Ts
	}
	if !IsNil(o.Metric) {
		toSerialize["metric"] = o.Metric
	}
	if !IsNil(o.Battery) {
		toSerialize["battery"] = o.Battery
	}
	if !IsNil(o.Button) {
		toSerialize["button"] = o.Button
	}
	if !IsNil(o.Door) {
		toSerialize["door"] = o.Door
	}
	if !IsNil(o.Humidity) {
		toSerialize["humidity"] = o.Humidity
	}
	if !IsNil(o.IndoorAirQuality) {
		toSerialize["indoorAirQuality"] = o.IndoorAirQuality
	}
	if !IsNil(o.Noise) {
		toSerialize["noise"] = o.Noise
	}
	if !IsNil(o.Pm25) {
		toSerialize["pm25"] = o.Pm25
	}
	if !IsNil(o.Temperature) {
		toSerialize["temperature"] = o.Temperature
	}
	if !IsNil(o.Tvoc) {
		toSerialize["tvoc"] = o.Tvoc
	}
	if !IsNil(o.Water) {
		toSerialize["water"] = o.Water
	}
	return toSerialize, nil
}

type NullableGetOrganizationSensorReadingsHistory200ResponseInner struct {
	value *GetOrganizationSensorReadingsHistory200ResponseInner
	isSet bool
}

func (v NullableGetOrganizationSensorReadingsHistory200ResponseInner) Get() *GetOrganizationSensorReadingsHistory200ResponseInner {
	return v.value
}

func (v *NullableGetOrganizationSensorReadingsHistory200ResponseInner) Set(val *GetOrganizationSensorReadingsHistory200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationSensorReadingsHistory200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationSensorReadingsHistory200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationSensorReadingsHistory200ResponseInner(val *GetOrganizationSensorReadingsHistory200ResponseInner) *NullableGetOrganizationSensorReadingsHistory200ResponseInner {
	return &NullableGetOrganizationSensorReadingsHistory200ResponseInner{value: val, isSet: true}
}

func (v NullableGetOrganizationSensorReadingsHistory200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationSensorReadingsHistory200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


