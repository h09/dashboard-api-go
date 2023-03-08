/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 22 February, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: v1.31.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20026Reasons struct for InlineResponse20026Reasons
type InlineResponse20026Reasons struct {
	// Reason for the rollback
	Category *string `json:"category,omitempty"`
	// Additional comment about the rollback
	Comment *string `json:"comment,omitempty"`
}

// NewInlineResponse20026Reasons instantiates a new InlineResponse20026Reasons object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20026Reasons() *InlineResponse20026Reasons {
	this := InlineResponse20026Reasons{}
	return &this
}

// NewInlineResponse20026ReasonsWithDefaults instantiates a new InlineResponse20026Reasons object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20026ReasonsWithDefaults() *InlineResponse20026Reasons {
	this := InlineResponse20026Reasons{}
	return &this
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *InlineResponse20026Reasons) GetCategory() string {
	if o == nil || isNil(o.Category) {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20026Reasons) GetCategoryOk() (*string, bool) {
	if o == nil || isNil(o.Category) {
    return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *InlineResponse20026Reasons) HasCategory() bool {
	if o != nil && !isNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *InlineResponse20026Reasons) SetCategory(v string) {
	o.Category = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *InlineResponse20026Reasons) GetComment() string {
	if o == nil || isNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20026Reasons) GetCommentOk() (*string, bool) {
	if o == nil || isNil(o.Comment) {
    return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *InlineResponse20026Reasons) HasComment() bool {
	if o != nil && !isNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *InlineResponse20026Reasons) SetComment(v string) {
	o.Comment = &v
}

func (o InlineResponse20026Reasons) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	if !isNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20026Reasons struct {
	value *InlineResponse20026Reasons
	isSet bool
}

func (v NullableInlineResponse20026Reasons) Get() *InlineResponse20026Reasons {
	return v.value
}

func (v *NullableInlineResponse20026Reasons) Set(val *InlineResponse20026Reasons) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20026Reasons) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20026Reasons) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20026Reasons(val *InlineResponse20026Reasons) *NullableInlineResponse20026Reasons {
	return &NullableInlineResponse20026Reasons{value: val, isSet: true}
}

func (v NullableInlineResponse20026Reasons) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20026Reasons) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

