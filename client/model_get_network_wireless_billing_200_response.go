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

// checks if the GetNetworkWirelessBilling200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkWirelessBilling200Response{}

// GetNetworkWirelessBilling200Response struct for GetNetworkWirelessBilling200Response
type GetNetworkWirelessBilling200Response struct {
	// The currency code of this node group's billing plans
	Currency *string `json:"currency,omitempty"`
	// Array of billing plans in the node group. (Can configure a maximum of 5)
	Plans []GetNetworkWirelessBilling200ResponsePlansInner `json:"plans,omitempty"`
}

// NewGetNetworkWirelessBilling200Response instantiates a new GetNetworkWirelessBilling200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkWirelessBilling200Response() *GetNetworkWirelessBilling200Response {
	this := GetNetworkWirelessBilling200Response{}
	return &this
}

// NewGetNetworkWirelessBilling200ResponseWithDefaults instantiates a new GetNetworkWirelessBilling200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkWirelessBilling200ResponseWithDefaults() *GetNetworkWirelessBilling200Response {
	this := GetNetworkWirelessBilling200Response{}
	return &this
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *GetNetworkWirelessBilling200Response) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessBilling200Response) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *GetNetworkWirelessBilling200Response) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *GetNetworkWirelessBilling200Response) SetCurrency(v string) {
	o.Currency = &v
}

// GetPlans returns the Plans field value if set, zero value otherwise.
func (o *GetNetworkWirelessBilling200Response) GetPlans() []GetNetworkWirelessBilling200ResponsePlansInner {
	if o == nil || IsNil(o.Plans) {
		var ret []GetNetworkWirelessBilling200ResponsePlansInner
		return ret
	}
	return o.Plans
}

// GetPlansOk returns a tuple with the Plans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessBilling200Response) GetPlansOk() ([]GetNetworkWirelessBilling200ResponsePlansInner, bool) {
	if o == nil || IsNil(o.Plans) {
		return nil, false
	}
	return o.Plans, true
}

// HasPlans returns a boolean if a field has been set.
func (o *GetNetworkWirelessBilling200Response) HasPlans() bool {
	if o != nil && !IsNil(o.Plans) {
		return true
	}

	return false
}

// SetPlans gets a reference to the given []GetNetworkWirelessBilling200ResponsePlansInner and assigns it to the Plans field.
func (o *GetNetworkWirelessBilling200Response) SetPlans(v []GetNetworkWirelessBilling200ResponsePlansInner) {
	o.Plans = v
}

func (o GetNetworkWirelessBilling200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkWirelessBilling200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.Plans) {
		toSerialize["plans"] = o.Plans
	}
	return toSerialize, nil
}

type NullableGetNetworkWirelessBilling200Response struct {
	value *GetNetworkWirelessBilling200Response
	isSet bool
}

func (v NullableGetNetworkWirelessBilling200Response) Get() *GetNetworkWirelessBilling200Response {
	return v.value
}

func (v *NullableGetNetworkWirelessBilling200Response) Set(val *GetNetworkWirelessBilling200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkWirelessBilling200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkWirelessBilling200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkWirelessBilling200Response(val *GetNetworkWirelessBilling200Response) *NullableGetNetworkWirelessBilling200Response {
	return &NullableGetNetworkWirelessBilling200Response{value: val, isSet: true}
}

func (v NullableGetNetworkWirelessBilling200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkWirelessBilling200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


