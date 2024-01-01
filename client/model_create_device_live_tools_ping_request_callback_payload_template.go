/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 18 December, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.41.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the CreateDeviceLiveToolsPingRequestCallbackPayloadTemplate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateDeviceLiveToolsPingRequestCallbackPayloadTemplate{}

// CreateDeviceLiveToolsPingRequestCallbackPayloadTemplate The payload template of the webhook used for the callback
type CreateDeviceLiveToolsPingRequestCallbackPayloadTemplate struct {
	// The ID of the payload template. Defaults to 'wpt_00005' for the Callback (included) template.
	Id *string `json:"id,omitempty"`
}

// NewCreateDeviceLiveToolsPingRequestCallbackPayloadTemplate instantiates a new CreateDeviceLiveToolsPingRequestCallbackPayloadTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateDeviceLiveToolsPingRequestCallbackPayloadTemplate() *CreateDeviceLiveToolsPingRequestCallbackPayloadTemplate {
	this := CreateDeviceLiveToolsPingRequestCallbackPayloadTemplate{}
	return &this
}

// NewCreateDeviceLiveToolsPingRequestCallbackPayloadTemplateWithDefaults instantiates a new CreateDeviceLiveToolsPingRequestCallbackPayloadTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateDeviceLiveToolsPingRequestCallbackPayloadTemplateWithDefaults() *CreateDeviceLiveToolsPingRequestCallbackPayloadTemplate {
	this := CreateDeviceLiveToolsPingRequestCallbackPayloadTemplate{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CreateDeviceLiveToolsPingRequestCallbackPayloadTemplate) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDeviceLiveToolsPingRequestCallbackPayloadTemplate) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CreateDeviceLiveToolsPingRequestCallbackPayloadTemplate) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CreateDeviceLiveToolsPingRequestCallbackPayloadTemplate) SetId(v string) {
	o.Id = &v
}

func (o CreateDeviceLiveToolsPingRequestCallbackPayloadTemplate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateDeviceLiveToolsPingRequestCallbackPayloadTemplate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableCreateDeviceLiveToolsPingRequestCallbackPayloadTemplate struct {
	value *CreateDeviceLiveToolsPingRequestCallbackPayloadTemplate
	isSet bool
}

func (v NullableCreateDeviceLiveToolsPingRequestCallbackPayloadTemplate) Get() *CreateDeviceLiveToolsPingRequestCallbackPayloadTemplate {
	return v.value
}

func (v *NullableCreateDeviceLiveToolsPingRequestCallbackPayloadTemplate) Set(val *CreateDeviceLiveToolsPingRequestCallbackPayloadTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateDeviceLiveToolsPingRequestCallbackPayloadTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateDeviceLiveToolsPingRequestCallbackPayloadTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateDeviceLiveToolsPingRequestCallbackPayloadTemplate(val *CreateDeviceLiveToolsPingRequestCallbackPayloadTemplate) *NullableCreateDeviceLiveToolsPingRequestCallbackPayloadTemplate {
	return &NullableCreateDeviceLiveToolsPingRequestCallbackPayloadTemplate{value: val, isSet: true}
}

func (v NullableCreateDeviceLiveToolsPingRequestCallbackPayloadTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateDeviceLiveToolsPingRequestCallbackPayloadTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

