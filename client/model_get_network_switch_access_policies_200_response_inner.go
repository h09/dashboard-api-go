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

// checks if the GetNetworkSwitchAccessPolicies200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkSwitchAccessPolicies200ResponseInner{}

// GetNetworkSwitchAccessPolicies200ResponseInner struct for GetNetworkSwitchAccessPolicies200ResponseInner
type GetNetworkSwitchAccessPolicies200ResponseInner struct {
	// Name of the access policy
	Name *string `json:"name,omitempty"`
	// List of RADIUS servers to require connecting devices to authenticate against before granting network access
	RadiusServers []GetNetworkSwitchAccessPolicies200ResponseInnerRadiusServersInner `json:"radiusServers,omitempty"`
	Radius *GetNetworkSwitchAccessPolicies200ResponseInnerRadius `json:"radius,omitempty"`
	// If enabled, Meraki devices will periodically send access-request messages to these RADIUS servers
	GuestPortBouncing *bool `json:"guestPortBouncing,omitempty"`
	// If enabled, Meraki devices will periodically send access-request messages to these RADIUS servers
	RadiusTestingEnabled *bool `json:"radiusTestingEnabled,omitempty"`
	// Change of authentication for RADIUS re-authentication and disconnection
	RadiusCoaSupportEnabled *bool `json:"radiusCoaSupportEnabled,omitempty"`
	// Enable to send start, interim-update and stop messages to a configured RADIUS accounting server for tracking connected clients
	RadiusAccountingEnabled *bool `json:"radiusAccountingEnabled,omitempty"`
	// List of RADIUS accounting servers to require connecting devices to authenticate against before granting network access
	RadiusAccountingServers []GetNetworkSwitchAccessPolicies200ResponseInnerRadiusAccountingServersInner `json:"radiusAccountingServers,omitempty"`
	// Acceptable values are `\"\"` for None, or `\"11\"` for Group Policies ACL
	RadiusGroupAttribute *string `json:"radiusGroupAttribute,omitempty"`
	// Choose the Host Mode for the access policy.
	HostMode *string `json:"hostMode,omitempty"`
	// Access Type of the policy. Automatically 'Hybrid authentication' when hostMode is 'Multi-Domain'.
	AccessPolicyType *string `json:"accessPolicyType,omitempty"`
	// Enabling this option will make switches execute 802.1X and MAC-bypass authentication simultaneously so that clients authenticate faster. Only required when accessPolicyType is 'Hybrid Authentication.
	IncreaseAccessSpeed *bool `json:"increaseAccessSpeed,omitempty"`
	// ID for the guest VLAN allow unauthorized devices access to limited network resources
	GuestVlanId *int32 `json:"guestVlanId,omitempty"`
	Dot1x *GetNetworkSwitchAccessPolicies200ResponseInnerDot1x `json:"dot1x,omitempty"`
	// CDP/LLDP capable voice clients will be able to use this VLAN. Automatically true when hostMode is 'Multi-Domain'.
	VoiceVlanClients *bool `json:"voiceVlanClients,omitempty"`
	// Enable to restrict access for clients to a response_objectific set of IP addresses or hostnames prior to authentication
	UrlRedirectWalledGardenEnabled *bool `json:"urlRedirectWalledGardenEnabled,omitempty"`
	// IP address ranges, in CIDR notation, to restrict access for clients to a specific set of IP addresses or hostnames prior to authentication
	UrlRedirectWalledGardenRanges []string `json:"urlRedirectWalledGardenRanges,omitempty"`
}

// NewGetNetworkSwitchAccessPolicies200ResponseInner instantiates a new GetNetworkSwitchAccessPolicies200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkSwitchAccessPolicies200ResponseInner() *GetNetworkSwitchAccessPolicies200ResponseInner {
	this := GetNetworkSwitchAccessPolicies200ResponseInner{}
	return &this
}

// NewGetNetworkSwitchAccessPolicies200ResponseInnerWithDefaults instantiates a new GetNetworkSwitchAccessPolicies200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkSwitchAccessPolicies200ResponseInnerWithDefaults() *GetNetworkSwitchAccessPolicies200ResponseInner {
	this := GetNetworkSwitchAccessPolicies200ResponseInner{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetNetworkSwitchAccessPolicies200ResponseInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchAccessPolicies200ResponseInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetNetworkSwitchAccessPolicies200ResponseInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetNetworkSwitchAccessPolicies200ResponseInner) SetName(v string) {
	o.Name = &v
}

// GetRadiusServers returns the RadiusServers field value if set, zero value otherwise.
func (o *GetNetworkSwitchAccessPolicies200ResponseInner) GetRadiusServers() []GetNetworkSwitchAccessPolicies200ResponseInnerRadiusServersInner {
	if o == nil || IsNil(o.RadiusServers) {
		var ret []GetNetworkSwitchAccessPolicies200ResponseInnerRadiusServersInner
		return ret
	}
	return o.RadiusServers
}

// GetRadiusServersOk returns a tuple with the RadiusServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchAccessPolicies200ResponseInner) GetRadiusServersOk() ([]GetNetworkSwitchAccessPolicies200ResponseInnerRadiusServersInner, bool) {
	if o == nil || IsNil(o.RadiusServers) {
		return nil, false
	}
	return o.RadiusServers, true
}

// HasRadiusServers returns a boolean if a field has been set.
func (o *GetNetworkSwitchAccessPolicies200ResponseInner) HasRadiusServers() bool {
	if o != nil && !IsNil(o.RadiusServers) {
		return true
	}

	return false
}

// SetRadiusServers gets a reference to the given []GetNetworkSwitchAccessPolicies200ResponseInnerRadiusServersInner and assigns it to the RadiusServers field.
func (o *GetNetworkSwitchAccessPolicies200ResponseInner) SetRadiusServers(v []GetNetworkSwitchAccessPolicies200ResponseInnerRadiusServersInner) {
	o.RadiusServers = v
}

// GetRadius returns the Radius field value if set, zero value otherwise.
func (o *GetNetworkSwitchAccessPolicies200ResponseInner) GetRadius() GetNetworkSwitchAccessPolicies200ResponseInnerRadius {
	if o == nil || IsNil(o.Radius) {
		var ret GetNetworkSwitchAccessPolicies200ResponseInnerRadius
		return ret
	}
	return *o.Radius
}

// GetRadiusOk returns a tuple with the Radius field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchAccessPolicies200ResponseInner) GetRadiusOk() (*GetNetworkSwitchAccessPolicies200ResponseInnerRadius, bool) {
	if o == nil || IsNil(o.Radius) {
		return nil, false
	}
	return o.Radius, true
}

// HasRadius returns a boolean if a field has been set.
func (o *GetNetworkSwitchAccessPolicies200ResponseInner) HasRadius() bool {
	if o != nil && !IsNil(o.Radius) {
		return true
	}

	return false
}

// SetRadius gets a reference to the given GetNetworkSwitchAccessPolicies200ResponseInnerRadius and assigns it to the Radius field.
func (o *GetNetworkSwitchAccessPolicies200ResponseInner) SetRadius(v GetNetworkSwitchAccessPolicies200ResponseInnerRadius) {
	o.Radius = &v
}

// GetGuestPortBouncing returns the GuestPortBouncing field value if set, zero value otherwise.
func (o *GetNetworkSwitchAccessPolicies200ResponseInner) GetGuestPortBouncing() bool {
	if o == nil || IsNil(o.GuestPortBouncing) {
		var ret bool
		return ret
	}
	return *o.GuestPortBouncing
}

// GetGuestPortBouncingOk returns a tuple with the GuestPortBouncing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchAccessPolicies200ResponseInner) GetGuestPortBouncingOk() (*bool, bool) {
	if o == nil || IsNil(o.GuestPortBouncing) {
		return nil, false
	}
	return o.GuestPortBouncing, true
}

// HasGuestPortBouncing returns a boolean if a field has been set.
func (o *GetNetworkSwitchAccessPolicies200ResponseInner) HasGuestPortBouncing() bool {
	if o != nil && !IsNil(o.GuestPortBouncing) {
		return true
	}

	return false
}

// SetGuestPortBouncing gets a reference to the given bool and assigns it to the GuestPortBouncing field.
func (o *GetNetworkSwitchAccessPolicies200ResponseInner) SetGuestPortBouncing(v bool) {
	o.GuestPortBouncing = &v
}

// GetRadiusTestingEnabled returns the RadiusTestingEnabled field value if set, zero value otherwise.
func (o *GetNetworkSwitchAccessPolicies200ResponseInner) GetRadiusTestingEnabled() bool {
	if o == nil || IsNil(o.RadiusTestingEnabled) {
		var ret bool
		return ret
	}
	return *o.RadiusTestingEnabled
}

// GetRadiusTestingEnabledOk returns a tuple with the RadiusTestingEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchAccessPolicies200ResponseInner) GetRadiusTestingEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.RadiusTestingEnabled) {
		return nil, false
	}
	return o.RadiusTestingEnabled, true
}

// HasRadiusTestingEnabled returns a boolean if a field has been set.
func (o *GetNetworkSwitchAccessPolicies200ResponseInner) HasRadiusTestingEnabled() bool {
	if o != nil && !IsNil(o.RadiusTestingEnabled) {
		return true
	}

	return false
}

// SetRadiusTestingEnabled gets a reference to the given bool and assigns it to the RadiusTestingEnabled field.
func (o *GetNetworkSwitchAccessPolicies200ResponseInner) SetRadiusTestingEnabled(v bool) {
	o.RadiusTestingEnabled = &v
}

// GetRadiusCoaSupportEnabled returns the RadiusCoaSupportEnabled field value if set, zero value otherwise.
func (o *GetNetworkSwitchAccessPolicies200ResponseInner) GetRadiusCoaSupportEnabled() bool {
	if o == nil || IsNil(o.RadiusCoaSupportEnabled) {
		var ret bool
		return ret
	}
	return *o.RadiusCoaSupportEnabled
}

// GetRadiusCoaSupportEnabledOk returns a tuple with the RadiusCoaSupportEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchAccessPolicies200ResponseInner) GetRadiusCoaSupportEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.RadiusCoaSupportEnabled) {
		return nil, false
	}
	return o.RadiusCoaSupportEnabled, true
}

// HasRadiusCoaSupportEnabled returns a boolean if a field has been set.
func (o *GetNetworkSwitchAccessPolicies200ResponseInner) HasRadiusCoaSupportEnabled() bool {
	if o != nil && !IsNil(o.RadiusCoaSupportEnabled) {
		return true
	}

	return false
}

// SetRadiusCoaSupportEnabled gets a reference to the given bool and assigns it to the RadiusCoaSupportEnabled field.
func (o *GetNetworkSwitchAccessPolicies200ResponseInner) SetRadiusCoaSupportEnabled(v bool) {
	o.RadiusCoaSupportEnabled = &v
}

// GetRadiusAccountingEnabled returns the RadiusAccountingEnabled field value if set, zero value otherwise.
func (o *GetNetworkSwitchAccessPolicies200ResponseInner) GetRadiusAccountingEnabled() bool {
	if o == nil || IsNil(o.RadiusAccountingEnabled) {
		var ret bool
		return ret
	}
	return *o.RadiusAccountingEnabled
}

// GetRadiusAccountingEnabledOk returns a tuple with the RadiusAccountingEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchAccessPolicies200ResponseInner) GetRadiusAccountingEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.RadiusAccountingEnabled) {
		return nil, false
	}
	return o.RadiusAccountingEnabled, true
}

// HasRadiusAccountingEnabled returns a boolean if a field has been set.
func (o *GetNetworkSwitchAccessPolicies200ResponseInner) HasRadiusAccountingEnabled() bool {
	if o != nil && !IsNil(o.RadiusAccountingEnabled) {
		return true
	}

	return false
}

// SetRadiusAccountingEnabled gets a reference to the given bool and assigns it to the RadiusAccountingEnabled field.
func (o *GetNetworkSwitchAccessPolicies200ResponseInner) SetRadiusAccountingEnabled(v bool) {
	o.RadiusAccountingEnabled = &v
}

// GetRadiusAccountingServers returns the RadiusAccountingServers field value if set, zero value otherwise.
func (o *GetNetworkSwitchAccessPolicies200ResponseInner) GetRadiusAccountingServers() []GetNetworkSwitchAccessPolicies200ResponseInnerRadiusAccountingServersInner {
	if o == nil || IsNil(o.RadiusAccountingServers) {
		var ret []GetNetworkSwitchAccessPolicies200ResponseInnerRadiusAccountingServersInner
		return ret
	}
	return o.RadiusAccountingServers
}

// GetRadiusAccountingServersOk returns a tuple with the RadiusAccountingServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchAccessPolicies200ResponseInner) GetRadiusAccountingServersOk() ([]GetNetworkSwitchAccessPolicies200ResponseInnerRadiusAccountingServersInner, bool) {
	if o == nil || IsNil(o.RadiusAccountingServers) {
		return nil, false
	}
	return o.RadiusAccountingServers, true
}

// HasRadiusAccountingServers returns a boolean if a field has been set.
func (o *GetNetworkSwitchAccessPolicies200ResponseInner) HasRadiusAccountingServers() bool {
	if o != nil && !IsNil(o.RadiusAccountingServers) {
		return true
	}

	return false
}

// SetRadiusAccountingServers gets a reference to the given []GetNetworkSwitchAccessPolicies200ResponseInnerRadiusAccountingServersInner and assigns it to the RadiusAccountingServers field.
func (o *GetNetworkSwitchAccessPolicies200ResponseInner) SetRadiusAccountingServers(v []GetNetworkSwitchAccessPolicies200ResponseInnerRadiusAccountingServersInner) {
	o.RadiusAccountingServers = v
}

// GetRadiusGroupAttribute returns the RadiusGroupAttribute field value if set, zero value otherwise.
func (o *GetNetworkSwitchAccessPolicies200ResponseInner) GetRadiusGroupAttribute() string {
	if o == nil || IsNil(o.RadiusGroupAttribute) {
		var ret string
		return ret
	}
	return *o.RadiusGroupAttribute
}

// GetRadiusGroupAttributeOk returns a tuple with the RadiusGroupAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchAccessPolicies200ResponseInner) GetRadiusGroupAttributeOk() (*string, bool) {
	if o == nil || IsNil(o.RadiusGroupAttribute) {
		return nil, false
	}
	return o.RadiusGroupAttribute, true
}

// HasRadiusGroupAttribute returns a boolean if a field has been set.
func (o *GetNetworkSwitchAccessPolicies200ResponseInner) HasRadiusGroupAttribute() bool {
	if o != nil && !IsNil(o.RadiusGroupAttribute) {
		return true
	}

	return false
}

// SetRadiusGroupAttribute gets a reference to the given string and assigns it to the RadiusGroupAttribute field.
func (o *GetNetworkSwitchAccessPolicies200ResponseInner) SetRadiusGroupAttribute(v string) {
	o.RadiusGroupAttribute = &v
}

// GetHostMode returns the HostMode field value if set, zero value otherwise.
func (o *GetNetworkSwitchAccessPolicies200ResponseInner) GetHostMode() string {
	if o == nil || IsNil(o.HostMode) {
		var ret string
		return ret
	}
	return *o.HostMode
}

// GetHostModeOk returns a tuple with the HostMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchAccessPolicies200ResponseInner) GetHostModeOk() (*string, bool) {
	if o == nil || IsNil(o.HostMode) {
		return nil, false
	}
	return o.HostMode, true
}

// HasHostMode returns a boolean if a field has been set.
func (o *GetNetworkSwitchAccessPolicies200ResponseInner) HasHostMode() bool {
	if o != nil && !IsNil(o.HostMode) {
		return true
	}

	return false
}

// SetHostMode gets a reference to the given string and assigns it to the HostMode field.
func (o *GetNetworkSwitchAccessPolicies200ResponseInner) SetHostMode(v string) {
	o.HostMode = &v
}

// GetAccessPolicyType returns the AccessPolicyType field value if set, zero value otherwise.
func (o *GetNetworkSwitchAccessPolicies200ResponseInner) GetAccessPolicyType() string {
	if o == nil || IsNil(o.AccessPolicyType) {
		var ret string
		return ret
	}
	return *o.AccessPolicyType
}

// GetAccessPolicyTypeOk returns a tuple with the AccessPolicyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchAccessPolicies200ResponseInner) GetAccessPolicyTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AccessPolicyType) {
		return nil, false
	}
	return o.AccessPolicyType, true
}

// HasAccessPolicyType returns a boolean if a field has been set.
func (o *GetNetworkSwitchAccessPolicies200ResponseInner) HasAccessPolicyType() bool {
	if o != nil && !IsNil(o.AccessPolicyType) {
		return true
	}

	return false
}

// SetAccessPolicyType gets a reference to the given string and assigns it to the AccessPolicyType field.
func (o *GetNetworkSwitchAccessPolicies200ResponseInner) SetAccessPolicyType(v string) {
	o.AccessPolicyType = &v
}

// GetIncreaseAccessSpeed returns the IncreaseAccessSpeed field value if set, zero value otherwise.
func (o *GetNetworkSwitchAccessPolicies200ResponseInner) GetIncreaseAccessSpeed() bool {
	if o == nil || IsNil(o.IncreaseAccessSpeed) {
		var ret bool
		return ret
	}
	return *o.IncreaseAccessSpeed
}

// GetIncreaseAccessSpeedOk returns a tuple with the IncreaseAccessSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchAccessPolicies200ResponseInner) GetIncreaseAccessSpeedOk() (*bool, bool) {
	if o == nil || IsNil(o.IncreaseAccessSpeed) {
		return nil, false
	}
	return o.IncreaseAccessSpeed, true
}

// HasIncreaseAccessSpeed returns a boolean if a field has been set.
func (o *GetNetworkSwitchAccessPolicies200ResponseInner) HasIncreaseAccessSpeed() bool {
	if o != nil && !IsNil(o.IncreaseAccessSpeed) {
		return true
	}

	return false
}

// SetIncreaseAccessSpeed gets a reference to the given bool and assigns it to the IncreaseAccessSpeed field.
func (o *GetNetworkSwitchAccessPolicies200ResponseInner) SetIncreaseAccessSpeed(v bool) {
	o.IncreaseAccessSpeed = &v
}

// GetGuestVlanId returns the GuestVlanId field value if set, zero value otherwise.
func (o *GetNetworkSwitchAccessPolicies200ResponseInner) GetGuestVlanId() int32 {
	if o == nil || IsNil(o.GuestVlanId) {
		var ret int32
		return ret
	}
	return *o.GuestVlanId
}

// GetGuestVlanIdOk returns a tuple with the GuestVlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchAccessPolicies200ResponseInner) GetGuestVlanIdOk() (*int32, bool) {
	if o == nil || IsNil(o.GuestVlanId) {
		return nil, false
	}
	return o.GuestVlanId, true
}

// HasGuestVlanId returns a boolean if a field has been set.
func (o *GetNetworkSwitchAccessPolicies200ResponseInner) HasGuestVlanId() bool {
	if o != nil && !IsNil(o.GuestVlanId) {
		return true
	}

	return false
}

// SetGuestVlanId gets a reference to the given int32 and assigns it to the GuestVlanId field.
func (o *GetNetworkSwitchAccessPolicies200ResponseInner) SetGuestVlanId(v int32) {
	o.GuestVlanId = &v
}

// GetDot1x returns the Dot1x field value if set, zero value otherwise.
func (o *GetNetworkSwitchAccessPolicies200ResponseInner) GetDot1x() GetNetworkSwitchAccessPolicies200ResponseInnerDot1x {
	if o == nil || IsNil(o.Dot1x) {
		var ret GetNetworkSwitchAccessPolicies200ResponseInnerDot1x
		return ret
	}
	return *o.Dot1x
}

// GetDot1xOk returns a tuple with the Dot1x field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchAccessPolicies200ResponseInner) GetDot1xOk() (*GetNetworkSwitchAccessPolicies200ResponseInnerDot1x, bool) {
	if o == nil || IsNil(o.Dot1x) {
		return nil, false
	}
	return o.Dot1x, true
}

// HasDot1x returns a boolean if a field has been set.
func (o *GetNetworkSwitchAccessPolicies200ResponseInner) HasDot1x() bool {
	if o != nil && !IsNil(o.Dot1x) {
		return true
	}

	return false
}

// SetDot1x gets a reference to the given GetNetworkSwitchAccessPolicies200ResponseInnerDot1x and assigns it to the Dot1x field.
func (o *GetNetworkSwitchAccessPolicies200ResponseInner) SetDot1x(v GetNetworkSwitchAccessPolicies200ResponseInnerDot1x) {
	o.Dot1x = &v
}

// GetVoiceVlanClients returns the VoiceVlanClients field value if set, zero value otherwise.
func (o *GetNetworkSwitchAccessPolicies200ResponseInner) GetVoiceVlanClients() bool {
	if o == nil || IsNil(o.VoiceVlanClients) {
		var ret bool
		return ret
	}
	return *o.VoiceVlanClients
}

// GetVoiceVlanClientsOk returns a tuple with the VoiceVlanClients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchAccessPolicies200ResponseInner) GetVoiceVlanClientsOk() (*bool, bool) {
	if o == nil || IsNil(o.VoiceVlanClients) {
		return nil, false
	}
	return o.VoiceVlanClients, true
}

// HasVoiceVlanClients returns a boolean if a field has been set.
func (o *GetNetworkSwitchAccessPolicies200ResponseInner) HasVoiceVlanClients() bool {
	if o != nil && !IsNil(o.VoiceVlanClients) {
		return true
	}

	return false
}

// SetVoiceVlanClients gets a reference to the given bool and assigns it to the VoiceVlanClients field.
func (o *GetNetworkSwitchAccessPolicies200ResponseInner) SetVoiceVlanClients(v bool) {
	o.VoiceVlanClients = &v
}

// GetUrlRedirectWalledGardenEnabled returns the UrlRedirectWalledGardenEnabled field value if set, zero value otherwise.
func (o *GetNetworkSwitchAccessPolicies200ResponseInner) GetUrlRedirectWalledGardenEnabled() bool {
	if o == nil || IsNil(o.UrlRedirectWalledGardenEnabled) {
		var ret bool
		return ret
	}
	return *o.UrlRedirectWalledGardenEnabled
}

// GetUrlRedirectWalledGardenEnabledOk returns a tuple with the UrlRedirectWalledGardenEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchAccessPolicies200ResponseInner) GetUrlRedirectWalledGardenEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.UrlRedirectWalledGardenEnabled) {
		return nil, false
	}
	return o.UrlRedirectWalledGardenEnabled, true
}

// HasUrlRedirectWalledGardenEnabled returns a boolean if a field has been set.
func (o *GetNetworkSwitchAccessPolicies200ResponseInner) HasUrlRedirectWalledGardenEnabled() bool {
	if o != nil && !IsNil(o.UrlRedirectWalledGardenEnabled) {
		return true
	}

	return false
}

// SetUrlRedirectWalledGardenEnabled gets a reference to the given bool and assigns it to the UrlRedirectWalledGardenEnabled field.
func (o *GetNetworkSwitchAccessPolicies200ResponseInner) SetUrlRedirectWalledGardenEnabled(v bool) {
	o.UrlRedirectWalledGardenEnabled = &v
}

// GetUrlRedirectWalledGardenRanges returns the UrlRedirectWalledGardenRanges field value if set, zero value otherwise.
func (o *GetNetworkSwitchAccessPolicies200ResponseInner) GetUrlRedirectWalledGardenRanges() []string {
	if o == nil || IsNil(o.UrlRedirectWalledGardenRanges) {
		var ret []string
		return ret
	}
	return o.UrlRedirectWalledGardenRanges
}

// GetUrlRedirectWalledGardenRangesOk returns a tuple with the UrlRedirectWalledGardenRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchAccessPolicies200ResponseInner) GetUrlRedirectWalledGardenRangesOk() ([]string, bool) {
	if o == nil || IsNil(o.UrlRedirectWalledGardenRanges) {
		return nil, false
	}
	return o.UrlRedirectWalledGardenRanges, true
}

// HasUrlRedirectWalledGardenRanges returns a boolean if a field has been set.
func (o *GetNetworkSwitchAccessPolicies200ResponseInner) HasUrlRedirectWalledGardenRanges() bool {
	if o != nil && !IsNil(o.UrlRedirectWalledGardenRanges) {
		return true
	}

	return false
}

// SetUrlRedirectWalledGardenRanges gets a reference to the given []string and assigns it to the UrlRedirectWalledGardenRanges field.
func (o *GetNetworkSwitchAccessPolicies200ResponseInner) SetUrlRedirectWalledGardenRanges(v []string) {
	o.UrlRedirectWalledGardenRanges = v
}

func (o GetNetworkSwitchAccessPolicies200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkSwitchAccessPolicies200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.RadiusServers) {
		toSerialize["radiusServers"] = o.RadiusServers
	}
	if !IsNil(o.Radius) {
		toSerialize["radius"] = o.Radius
	}
	if !IsNil(o.GuestPortBouncing) {
		toSerialize["guestPortBouncing"] = o.GuestPortBouncing
	}
	if !IsNil(o.RadiusTestingEnabled) {
		toSerialize["radiusTestingEnabled"] = o.RadiusTestingEnabled
	}
	if !IsNil(o.RadiusCoaSupportEnabled) {
		toSerialize["radiusCoaSupportEnabled"] = o.RadiusCoaSupportEnabled
	}
	if !IsNil(o.RadiusAccountingEnabled) {
		toSerialize["radiusAccountingEnabled"] = o.RadiusAccountingEnabled
	}
	if !IsNil(o.RadiusAccountingServers) {
		toSerialize["radiusAccountingServers"] = o.RadiusAccountingServers
	}
	if !IsNil(o.RadiusGroupAttribute) {
		toSerialize["radiusGroupAttribute"] = o.RadiusGroupAttribute
	}
	if !IsNil(o.HostMode) {
		toSerialize["hostMode"] = o.HostMode
	}
	if !IsNil(o.AccessPolicyType) {
		toSerialize["accessPolicyType"] = o.AccessPolicyType
	}
	if !IsNil(o.IncreaseAccessSpeed) {
		toSerialize["increaseAccessSpeed"] = o.IncreaseAccessSpeed
	}
	if !IsNil(o.GuestVlanId) {
		toSerialize["guestVlanId"] = o.GuestVlanId
	}
	if !IsNil(o.Dot1x) {
		toSerialize["dot1x"] = o.Dot1x
	}
	if !IsNil(o.VoiceVlanClients) {
		toSerialize["voiceVlanClients"] = o.VoiceVlanClients
	}
	if !IsNil(o.UrlRedirectWalledGardenEnabled) {
		toSerialize["urlRedirectWalledGardenEnabled"] = o.UrlRedirectWalledGardenEnabled
	}
	if !IsNil(o.UrlRedirectWalledGardenRanges) {
		toSerialize["urlRedirectWalledGardenRanges"] = o.UrlRedirectWalledGardenRanges
	}
	return toSerialize, nil
}

type NullableGetNetworkSwitchAccessPolicies200ResponseInner struct {
	value *GetNetworkSwitchAccessPolicies200ResponseInner
	isSet bool
}

func (v NullableGetNetworkSwitchAccessPolicies200ResponseInner) Get() *GetNetworkSwitchAccessPolicies200ResponseInner {
	return v.value
}

func (v *NullableGetNetworkSwitchAccessPolicies200ResponseInner) Set(val *GetNetworkSwitchAccessPolicies200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkSwitchAccessPolicies200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkSwitchAccessPolicies200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkSwitchAccessPolicies200ResponseInner(val *GetNetworkSwitchAccessPolicies200ResponseInner) *NullableGetNetworkSwitchAccessPolicies200ResponseInner {
	return &NullableGetNetworkSwitchAccessPolicies200ResponseInner{value: val, isSet: true}
}

func (v NullableGetNetworkSwitchAccessPolicies200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkSwitchAccessPolicies200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


