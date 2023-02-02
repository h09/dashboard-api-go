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

// InlineObject6 struct for InlineObject6
type InlineObject6 struct {
	// Boolean indicating if sense(license) is enabled(true) or disabled(false) on the camera
	SenseEnabled *bool `json:"senseEnabled,omitempty"`
	// The ID of the MQTT broker to be enabled on the camera. A value of null will disable MQTT on the camera
	MqttBrokerId *string `json:"mqttBrokerId,omitempty"`
	AudioDetection *DevicesSerialCameraSenseAudioDetection `json:"audioDetection,omitempty"`
	// The ID of the object detection model
	DetectionModelId *string `json:"detectionModelId,omitempty"`
}

// NewInlineObject6 instantiates a new InlineObject6 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject6() *InlineObject6 {
	this := InlineObject6{}
	return &this
}

// NewInlineObject6WithDefaults instantiates a new InlineObject6 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject6WithDefaults() *InlineObject6 {
	this := InlineObject6{}
	return &this
}

// GetSenseEnabled returns the SenseEnabled field value if set, zero value otherwise.
func (o *InlineObject6) GetSenseEnabled() bool {
	if o == nil || isNil(o.SenseEnabled) {
		var ret bool
		return ret
	}
	return *o.SenseEnabled
}

// GetSenseEnabledOk returns a tuple with the SenseEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject6) GetSenseEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.SenseEnabled) {
    return nil, false
	}
	return o.SenseEnabled, true
}

// HasSenseEnabled returns a boolean if a field has been set.
func (o *InlineObject6) HasSenseEnabled() bool {
	if o != nil && !isNil(o.SenseEnabled) {
		return true
	}

	return false
}

// SetSenseEnabled gets a reference to the given bool and assigns it to the SenseEnabled field.
func (o *InlineObject6) SetSenseEnabled(v bool) {
	o.SenseEnabled = &v
}

// GetMqttBrokerId returns the MqttBrokerId field value if set, zero value otherwise.
func (o *InlineObject6) GetMqttBrokerId() string {
	if o == nil || isNil(o.MqttBrokerId) {
		var ret string
		return ret
	}
	return *o.MqttBrokerId
}

// GetMqttBrokerIdOk returns a tuple with the MqttBrokerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject6) GetMqttBrokerIdOk() (*string, bool) {
	if o == nil || isNil(o.MqttBrokerId) {
    return nil, false
	}
	return o.MqttBrokerId, true
}

// HasMqttBrokerId returns a boolean if a field has been set.
func (o *InlineObject6) HasMqttBrokerId() bool {
	if o != nil && !isNil(o.MqttBrokerId) {
		return true
	}

	return false
}

// SetMqttBrokerId gets a reference to the given string and assigns it to the MqttBrokerId field.
func (o *InlineObject6) SetMqttBrokerId(v string) {
	o.MqttBrokerId = &v
}

// GetAudioDetection returns the AudioDetection field value if set, zero value otherwise.
func (o *InlineObject6) GetAudioDetection() DevicesSerialCameraSenseAudioDetection {
	if o == nil || isNil(o.AudioDetection) {
		var ret DevicesSerialCameraSenseAudioDetection
		return ret
	}
	return *o.AudioDetection
}

// GetAudioDetectionOk returns a tuple with the AudioDetection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject6) GetAudioDetectionOk() (*DevicesSerialCameraSenseAudioDetection, bool) {
	if o == nil || isNil(o.AudioDetection) {
    return nil, false
	}
	return o.AudioDetection, true
}

// HasAudioDetection returns a boolean if a field has been set.
func (o *InlineObject6) HasAudioDetection() bool {
	if o != nil && !isNil(o.AudioDetection) {
		return true
	}

	return false
}

// SetAudioDetection gets a reference to the given DevicesSerialCameraSenseAudioDetection and assigns it to the AudioDetection field.
func (o *InlineObject6) SetAudioDetection(v DevicesSerialCameraSenseAudioDetection) {
	o.AudioDetection = &v
}

// GetDetectionModelId returns the DetectionModelId field value if set, zero value otherwise.
func (o *InlineObject6) GetDetectionModelId() string {
	if o == nil || isNil(o.DetectionModelId) {
		var ret string
		return ret
	}
	return *o.DetectionModelId
}

// GetDetectionModelIdOk returns a tuple with the DetectionModelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject6) GetDetectionModelIdOk() (*string, bool) {
	if o == nil || isNil(o.DetectionModelId) {
    return nil, false
	}
	return o.DetectionModelId, true
}

// HasDetectionModelId returns a boolean if a field has been set.
func (o *InlineObject6) HasDetectionModelId() bool {
	if o != nil && !isNil(o.DetectionModelId) {
		return true
	}

	return false
}

// SetDetectionModelId gets a reference to the given string and assigns it to the DetectionModelId field.
func (o *InlineObject6) SetDetectionModelId(v string) {
	o.DetectionModelId = &v
}

func (o InlineObject6) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.SenseEnabled) {
		toSerialize["senseEnabled"] = o.SenseEnabled
	}
	if !isNil(o.MqttBrokerId) {
		toSerialize["mqttBrokerId"] = o.MqttBrokerId
	}
	if !isNil(o.AudioDetection) {
		toSerialize["audioDetection"] = o.AudioDetection
	}
	if !isNil(o.DetectionModelId) {
		toSerialize["detectionModelId"] = o.DetectionModelId
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject6 struct {
	value *InlineObject6
	isSet bool
}

func (v NullableInlineObject6) Get() *InlineObject6 {
	return v.value
}

func (v *NullableInlineObject6) Set(val *InlineObject6) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject6) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject6) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject6(val *InlineObject6) *NullableInlineObject6 {
	return &NullableInlineObject6{value: val, isSet: true}
}

func (v NullableInlineObject6) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject6) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


