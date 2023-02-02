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

// InlineResponse2015Status Status of action batch
type InlineResponse2015Status struct {
	// Flag describing whether all actions in the action batch have completed
	Completed *bool `json:"completed,omitempty"`
	// Flag describing whether any actions in the action batch failed
	Failed *bool `json:"failed,omitempty"`
	// List of errors encountered when running actions in the action batch
	Errors []string `json:"errors,omitempty"`
	// Resources created as a result of this action batch
	CreatedResources []InlineResponse2015StatusCreatedResources `json:"createdResources"`
}

// NewInlineResponse2015Status instantiates a new InlineResponse2015Status object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2015Status(createdResources []InlineResponse2015StatusCreatedResources) *InlineResponse2015Status {
	this := InlineResponse2015Status{}
	this.CreatedResources = createdResources
	return &this
}

// NewInlineResponse2015StatusWithDefaults instantiates a new InlineResponse2015Status object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2015StatusWithDefaults() *InlineResponse2015Status {
	this := InlineResponse2015Status{}
	return &this
}

// GetCompleted returns the Completed field value if set, zero value otherwise.
func (o *InlineResponse2015Status) GetCompleted() bool {
	if o == nil || isNil(o.Completed) {
		var ret bool
		return ret
	}
	return *o.Completed
}

// GetCompletedOk returns a tuple with the Completed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2015Status) GetCompletedOk() (*bool, bool) {
	if o == nil || isNil(o.Completed) {
    return nil, false
	}
	return o.Completed, true
}

// HasCompleted returns a boolean if a field has been set.
func (o *InlineResponse2015Status) HasCompleted() bool {
	if o != nil && !isNil(o.Completed) {
		return true
	}

	return false
}

// SetCompleted gets a reference to the given bool and assigns it to the Completed field.
func (o *InlineResponse2015Status) SetCompleted(v bool) {
	o.Completed = &v
}

// GetFailed returns the Failed field value if set, zero value otherwise.
func (o *InlineResponse2015Status) GetFailed() bool {
	if o == nil || isNil(o.Failed) {
		var ret bool
		return ret
	}
	return *o.Failed
}

// GetFailedOk returns a tuple with the Failed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2015Status) GetFailedOk() (*bool, bool) {
	if o == nil || isNil(o.Failed) {
    return nil, false
	}
	return o.Failed, true
}

// HasFailed returns a boolean if a field has been set.
func (o *InlineResponse2015Status) HasFailed() bool {
	if o != nil && !isNil(o.Failed) {
		return true
	}

	return false
}

// SetFailed gets a reference to the given bool and assigns it to the Failed field.
func (o *InlineResponse2015Status) SetFailed(v bool) {
	o.Failed = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *InlineResponse2015Status) GetErrors() []string {
	if o == nil || isNil(o.Errors) {
		var ret []string
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2015Status) GetErrorsOk() ([]string, bool) {
	if o == nil || isNil(o.Errors) {
    return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *InlineResponse2015Status) HasErrors() bool {
	if o != nil && !isNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []string and assigns it to the Errors field.
func (o *InlineResponse2015Status) SetErrors(v []string) {
	o.Errors = v
}

// GetCreatedResources returns the CreatedResources field value
func (o *InlineResponse2015Status) GetCreatedResources() []InlineResponse2015StatusCreatedResources {
	if o == nil {
		var ret []InlineResponse2015StatusCreatedResources
		return ret
	}

	return o.CreatedResources
}

// GetCreatedResourcesOk returns a tuple with the CreatedResources field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2015Status) GetCreatedResourcesOk() ([]InlineResponse2015StatusCreatedResources, bool) {
	if o == nil {
    return nil, false
	}
	return o.CreatedResources, true
}

// SetCreatedResources sets field value
func (o *InlineResponse2015Status) SetCreatedResources(v []InlineResponse2015StatusCreatedResources) {
	o.CreatedResources = v
}

func (o InlineResponse2015Status) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Completed) {
		toSerialize["completed"] = o.Completed
	}
	if !isNil(o.Failed) {
		toSerialize["failed"] = o.Failed
	}
	if !isNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	if true {
		toSerialize["createdResources"] = o.CreatedResources
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2015Status struct {
	value *InlineResponse2015Status
	isSet bool
}

func (v NullableInlineResponse2015Status) Get() *InlineResponse2015Status {
	return v.value
}

func (v *NullableInlineResponse2015Status) Set(val *InlineResponse2015Status) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2015Status) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2015Status) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2015Status(val *InlineResponse2015Status) *NullableInlineResponse2015Status {
	return &NullableInlineResponse2015Status{value: val, isSet: true}
}

func (v NullableInlineResponse2015Status) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2015Status) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


