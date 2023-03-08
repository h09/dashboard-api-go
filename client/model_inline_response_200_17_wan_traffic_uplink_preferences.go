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

// InlineResponse20017WanTrafficUplinkPreferences struct for InlineResponse20017WanTrafficUplinkPreferences
type InlineResponse20017WanTrafficUplinkPreferences struct {
	// Traffic filters
	TrafficFilters []InlineResponse20017TrafficFilters `json:"trafficFilters"`
	// Preferred uplink for uplink preference rule. Must be one of: 'wan1' or 'wan2'
	PreferredUplink string `json:"preferredUplink"`
}

// NewInlineResponse20017WanTrafficUplinkPreferences instantiates a new InlineResponse20017WanTrafficUplinkPreferences object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20017WanTrafficUplinkPreferences(trafficFilters []InlineResponse20017TrafficFilters, preferredUplink string) *InlineResponse20017WanTrafficUplinkPreferences {
	this := InlineResponse20017WanTrafficUplinkPreferences{}
	this.TrafficFilters = trafficFilters
	this.PreferredUplink = preferredUplink
	return &this
}

// NewInlineResponse20017WanTrafficUplinkPreferencesWithDefaults instantiates a new InlineResponse20017WanTrafficUplinkPreferences object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20017WanTrafficUplinkPreferencesWithDefaults() *InlineResponse20017WanTrafficUplinkPreferences {
	this := InlineResponse20017WanTrafficUplinkPreferences{}
	return &this
}

// GetTrafficFilters returns the TrafficFilters field value
func (o *InlineResponse20017WanTrafficUplinkPreferences) GetTrafficFilters() []InlineResponse20017TrafficFilters {
	if o == nil {
		var ret []InlineResponse20017TrafficFilters
		return ret
	}

	return o.TrafficFilters
}

// GetTrafficFiltersOk returns a tuple with the TrafficFilters field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20017WanTrafficUplinkPreferences) GetTrafficFiltersOk() ([]InlineResponse20017TrafficFilters, bool) {
	if o == nil {
    return nil, false
	}
	return o.TrafficFilters, true
}

// SetTrafficFilters sets field value
func (o *InlineResponse20017WanTrafficUplinkPreferences) SetTrafficFilters(v []InlineResponse20017TrafficFilters) {
	o.TrafficFilters = v
}

// GetPreferredUplink returns the PreferredUplink field value
func (o *InlineResponse20017WanTrafficUplinkPreferences) GetPreferredUplink() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PreferredUplink
}

// GetPreferredUplinkOk returns a tuple with the PreferredUplink field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20017WanTrafficUplinkPreferences) GetPreferredUplinkOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.PreferredUplink, true
}

// SetPreferredUplink sets field value
func (o *InlineResponse20017WanTrafficUplinkPreferences) SetPreferredUplink(v string) {
	o.PreferredUplink = v
}

func (o InlineResponse20017WanTrafficUplinkPreferences) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["trafficFilters"] = o.TrafficFilters
	}
	if true {
		toSerialize["preferredUplink"] = o.PreferredUplink
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20017WanTrafficUplinkPreferences struct {
	value *InlineResponse20017WanTrafficUplinkPreferences
	isSet bool
}

func (v NullableInlineResponse20017WanTrafficUplinkPreferences) Get() *InlineResponse20017WanTrafficUplinkPreferences {
	return v.value
}

func (v *NullableInlineResponse20017WanTrafficUplinkPreferences) Set(val *InlineResponse20017WanTrafficUplinkPreferences) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20017WanTrafficUplinkPreferences) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20017WanTrafficUplinkPreferences) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20017WanTrafficUplinkPreferences(val *InlineResponse20017WanTrafficUplinkPreferences) *NullableInlineResponse20017WanTrafficUplinkPreferences {
	return &NullableInlineResponse20017WanTrafficUplinkPreferences{value: val, isSet: true}
}

func (v NullableInlineResponse20017WanTrafficUplinkPreferences) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20017WanTrafficUplinkPreferences) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

