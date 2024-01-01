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

// checks if the BindAdministeredLicensingSubscriptionSubscription200ResponseInsufficientEntitlementsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BindAdministeredLicensingSubscriptionSubscription200ResponseInsufficientEntitlementsInner{}

// BindAdministeredLicensingSubscriptionSubscription200ResponseInsufficientEntitlementsInner struct for BindAdministeredLicensingSubscriptionSubscription200ResponseInsufficientEntitlementsInner
type BindAdministeredLicensingSubscriptionSubscription200ResponseInsufficientEntitlementsInner struct {
	// SKU of the required product
	Sku *string `json:"sku,omitempty"`
	// Number required
	Quantity *int32 `json:"quantity,omitempty"`
}

// NewBindAdministeredLicensingSubscriptionSubscription200ResponseInsufficientEntitlementsInner instantiates a new BindAdministeredLicensingSubscriptionSubscription200ResponseInsufficientEntitlementsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBindAdministeredLicensingSubscriptionSubscription200ResponseInsufficientEntitlementsInner() *BindAdministeredLicensingSubscriptionSubscription200ResponseInsufficientEntitlementsInner {
	this := BindAdministeredLicensingSubscriptionSubscription200ResponseInsufficientEntitlementsInner{}
	return &this
}

// NewBindAdministeredLicensingSubscriptionSubscription200ResponseInsufficientEntitlementsInnerWithDefaults instantiates a new BindAdministeredLicensingSubscriptionSubscription200ResponseInsufficientEntitlementsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBindAdministeredLicensingSubscriptionSubscription200ResponseInsufficientEntitlementsInnerWithDefaults() *BindAdministeredLicensingSubscriptionSubscription200ResponseInsufficientEntitlementsInner {
	this := BindAdministeredLicensingSubscriptionSubscription200ResponseInsufficientEntitlementsInner{}
	return &this
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *BindAdministeredLicensingSubscriptionSubscription200ResponseInsufficientEntitlementsInner) GetSku() string {
	if o == nil || IsNil(o.Sku) {
		var ret string
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BindAdministeredLicensingSubscriptionSubscription200ResponseInsufficientEntitlementsInner) GetSkuOk() (*string, bool) {
	if o == nil || IsNil(o.Sku) {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *BindAdministeredLicensingSubscriptionSubscription200ResponseInsufficientEntitlementsInner) HasSku() bool {
	if o != nil && !IsNil(o.Sku) {
		return true
	}

	return false
}

// SetSku gets a reference to the given string and assigns it to the Sku field.
func (o *BindAdministeredLicensingSubscriptionSubscription200ResponseInsufficientEntitlementsInner) SetSku(v string) {
	o.Sku = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *BindAdministeredLicensingSubscriptionSubscription200ResponseInsufficientEntitlementsInner) GetQuantity() int32 {
	if o == nil || IsNil(o.Quantity) {
		var ret int32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BindAdministeredLicensingSubscriptionSubscription200ResponseInsufficientEntitlementsInner) GetQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *BindAdministeredLicensingSubscriptionSubscription200ResponseInsufficientEntitlementsInner) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int32 and assigns it to the Quantity field.
func (o *BindAdministeredLicensingSubscriptionSubscription200ResponseInsufficientEntitlementsInner) SetQuantity(v int32) {
	o.Quantity = &v
}

func (o BindAdministeredLicensingSubscriptionSubscription200ResponseInsufficientEntitlementsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BindAdministeredLicensingSubscriptionSubscription200ResponseInsufficientEntitlementsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Sku) {
		toSerialize["sku"] = o.Sku
	}
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	return toSerialize, nil
}

type NullableBindAdministeredLicensingSubscriptionSubscription200ResponseInsufficientEntitlementsInner struct {
	value *BindAdministeredLicensingSubscriptionSubscription200ResponseInsufficientEntitlementsInner
	isSet bool
}

func (v NullableBindAdministeredLicensingSubscriptionSubscription200ResponseInsufficientEntitlementsInner) Get() *BindAdministeredLicensingSubscriptionSubscription200ResponseInsufficientEntitlementsInner {
	return v.value
}

func (v *NullableBindAdministeredLicensingSubscriptionSubscription200ResponseInsufficientEntitlementsInner) Set(val *BindAdministeredLicensingSubscriptionSubscription200ResponseInsufficientEntitlementsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableBindAdministeredLicensingSubscriptionSubscription200ResponseInsufficientEntitlementsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableBindAdministeredLicensingSubscriptionSubscription200ResponseInsufficientEntitlementsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBindAdministeredLicensingSubscriptionSubscription200ResponseInsufficientEntitlementsInner(val *BindAdministeredLicensingSubscriptionSubscription200ResponseInsufficientEntitlementsInner) *NullableBindAdministeredLicensingSubscriptionSubscription200ResponseInsufficientEntitlementsInner {
	return &NullableBindAdministeredLicensingSubscriptionSubscription200ResponseInsufficientEntitlementsInner{value: val, isSet: true}
}

func (v NullableBindAdministeredLicensingSubscriptionSubscription200ResponseInsufficientEntitlementsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBindAdministeredLicensingSubscriptionSubscription200ResponseInsufficientEntitlementsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

