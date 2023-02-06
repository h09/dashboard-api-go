# InlineResponse20078

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SsidNumber** | Pointer to **int32** | SSID number | [optional] 
**SplashPage** | Pointer to **string** | The type of splash page for this SSID | [optional] 
**UseSplashUrl** | Pointer to **bool** | Boolean indicating whether the users will be redirected to the custom splash url | [optional] 
**SplashUrl** | Pointer to **string** | The custom splash URL of the click-through splash page. | [optional] 
**SplashTimeout** | Pointer to **int32** | Splash timeout in minutes. | [optional] 
**RedirectUrl** | Pointer to **string** | The custom redirect URL where the users will go after the splash page. | [optional] 
**UseRedirectUrl** | Pointer to **bool** | The Boolean indicating whether the the user will be redirected to the custom redirect URL after the splash page. | [optional] 
**WelcomeMessage** | Pointer to **string** | The welcome message for the users on the splash page. | [optional] 
**SplashLogo** | Pointer to [**InlineResponse20078SplashLogo**](InlineResponse20078SplashLogo.md) |  | [optional] 
**SplashImage** | Pointer to [**InlineResponse20078SplashImage**](InlineResponse20078SplashImage.md) |  | [optional] 
**SplashPrepaidFront** | Pointer to [**InlineResponse20078SplashPrepaidFront**](InlineResponse20078SplashPrepaidFront.md) |  | [optional] 
**GuestSponsorship** | Pointer to [**InlineResponse20078GuestSponsorship**](InlineResponse20078GuestSponsorship.md) |  | [optional] 
**BlockAllTrafficBeforeSignOn** | Pointer to **bool** | How restricted allowing traffic should be. If true, all traffic types are blocked until the splash page is acknowledged. If false, all non-HTTP traffic is allowed before the splash page is acknowledged. | [optional] 
**ControllerDisconnectionBehavior** | Pointer to **string** | How login attempts should be handled when the controller is unreachable. | [optional] 
**AllowSimultaneousLogins** | Pointer to **bool** | Whether or not to allow simultaneous logins from different devices. | [optional] 
**Billing** | Pointer to [**InlineResponse20078Billing**](InlineResponse20078Billing.md) |  | [optional] 
**SentryEnrollment** | Pointer to [**InlineResponse20078SentryEnrollment**](InlineResponse20078SentryEnrollment.md) |  | [optional] 
**SelfRegistration** | Pointer to [**InlineResponse20078SelfRegistration**](InlineResponse20078SelfRegistration.md) |  | [optional] 

## Methods

### NewInlineResponse20078

`func NewInlineResponse20078() *InlineResponse20078`

NewInlineResponse20078 instantiates a new InlineResponse20078 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20078WithDefaults

`func NewInlineResponse20078WithDefaults() *InlineResponse20078`

NewInlineResponse20078WithDefaults instantiates a new InlineResponse20078 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSsidNumber

`func (o *InlineResponse20078) GetSsidNumber() int32`

GetSsidNumber returns the SsidNumber field if non-nil, zero value otherwise.

### GetSsidNumberOk

`func (o *InlineResponse20078) GetSsidNumberOk() (*int32, bool)`

GetSsidNumberOk returns a tuple with the SsidNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsidNumber

`func (o *InlineResponse20078) SetSsidNumber(v int32)`

SetSsidNumber sets SsidNumber field to given value.

### HasSsidNumber

`func (o *InlineResponse20078) HasSsidNumber() bool`

HasSsidNumber returns a boolean if a field has been set.

### GetSplashPage

`func (o *InlineResponse20078) GetSplashPage() string`

GetSplashPage returns the SplashPage field if non-nil, zero value otherwise.

### GetSplashPageOk

`func (o *InlineResponse20078) GetSplashPageOk() (*string, bool)`

GetSplashPageOk returns a tuple with the SplashPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplashPage

`func (o *InlineResponse20078) SetSplashPage(v string)`

SetSplashPage sets SplashPage field to given value.

### HasSplashPage

`func (o *InlineResponse20078) HasSplashPage() bool`

HasSplashPage returns a boolean if a field has been set.

### GetUseSplashUrl

`func (o *InlineResponse20078) GetUseSplashUrl() bool`

GetUseSplashUrl returns the UseSplashUrl field if non-nil, zero value otherwise.

### GetUseSplashUrlOk

`func (o *InlineResponse20078) GetUseSplashUrlOk() (*bool, bool)`

GetUseSplashUrlOk returns a tuple with the UseSplashUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSplashUrl

`func (o *InlineResponse20078) SetUseSplashUrl(v bool)`

SetUseSplashUrl sets UseSplashUrl field to given value.

### HasUseSplashUrl

`func (o *InlineResponse20078) HasUseSplashUrl() bool`

HasUseSplashUrl returns a boolean if a field has been set.

### GetSplashUrl

`func (o *InlineResponse20078) GetSplashUrl() string`

GetSplashUrl returns the SplashUrl field if non-nil, zero value otherwise.

### GetSplashUrlOk

`func (o *InlineResponse20078) GetSplashUrlOk() (*string, bool)`

GetSplashUrlOk returns a tuple with the SplashUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplashUrl

`func (o *InlineResponse20078) SetSplashUrl(v string)`

SetSplashUrl sets SplashUrl field to given value.

### HasSplashUrl

`func (o *InlineResponse20078) HasSplashUrl() bool`

HasSplashUrl returns a boolean if a field has been set.

### GetSplashTimeout

`func (o *InlineResponse20078) GetSplashTimeout() int32`

GetSplashTimeout returns the SplashTimeout field if non-nil, zero value otherwise.

### GetSplashTimeoutOk

`func (o *InlineResponse20078) GetSplashTimeoutOk() (*int32, bool)`

GetSplashTimeoutOk returns a tuple with the SplashTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplashTimeout

`func (o *InlineResponse20078) SetSplashTimeout(v int32)`

SetSplashTimeout sets SplashTimeout field to given value.

### HasSplashTimeout

`func (o *InlineResponse20078) HasSplashTimeout() bool`

HasSplashTimeout returns a boolean if a field has been set.

### GetRedirectUrl

`func (o *InlineResponse20078) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *InlineResponse20078) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *InlineResponse20078) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *InlineResponse20078) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.

### GetUseRedirectUrl

`func (o *InlineResponse20078) GetUseRedirectUrl() bool`

GetUseRedirectUrl returns the UseRedirectUrl field if non-nil, zero value otherwise.

### GetUseRedirectUrlOk

`func (o *InlineResponse20078) GetUseRedirectUrlOk() (*bool, bool)`

GetUseRedirectUrlOk returns a tuple with the UseRedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseRedirectUrl

`func (o *InlineResponse20078) SetUseRedirectUrl(v bool)`

SetUseRedirectUrl sets UseRedirectUrl field to given value.

### HasUseRedirectUrl

`func (o *InlineResponse20078) HasUseRedirectUrl() bool`

HasUseRedirectUrl returns a boolean if a field has been set.

### GetWelcomeMessage

`func (o *InlineResponse20078) GetWelcomeMessage() string`

GetWelcomeMessage returns the WelcomeMessage field if non-nil, zero value otherwise.

### GetWelcomeMessageOk

`func (o *InlineResponse20078) GetWelcomeMessageOk() (*string, bool)`

GetWelcomeMessageOk returns a tuple with the WelcomeMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWelcomeMessage

`func (o *InlineResponse20078) SetWelcomeMessage(v string)`

SetWelcomeMessage sets WelcomeMessage field to given value.

### HasWelcomeMessage

`func (o *InlineResponse20078) HasWelcomeMessage() bool`

HasWelcomeMessage returns a boolean if a field has been set.

### GetSplashLogo

`func (o *InlineResponse20078) GetSplashLogo() InlineResponse20078SplashLogo`

GetSplashLogo returns the SplashLogo field if non-nil, zero value otherwise.

### GetSplashLogoOk

`func (o *InlineResponse20078) GetSplashLogoOk() (*InlineResponse20078SplashLogo, bool)`

GetSplashLogoOk returns a tuple with the SplashLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplashLogo

`func (o *InlineResponse20078) SetSplashLogo(v InlineResponse20078SplashLogo)`

SetSplashLogo sets SplashLogo field to given value.

### HasSplashLogo

`func (o *InlineResponse20078) HasSplashLogo() bool`

HasSplashLogo returns a boolean if a field has been set.

### GetSplashImage

`func (o *InlineResponse20078) GetSplashImage() InlineResponse20078SplashImage`

GetSplashImage returns the SplashImage field if non-nil, zero value otherwise.

### GetSplashImageOk

`func (o *InlineResponse20078) GetSplashImageOk() (*InlineResponse20078SplashImage, bool)`

GetSplashImageOk returns a tuple with the SplashImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplashImage

`func (o *InlineResponse20078) SetSplashImage(v InlineResponse20078SplashImage)`

SetSplashImage sets SplashImage field to given value.

### HasSplashImage

`func (o *InlineResponse20078) HasSplashImage() bool`

HasSplashImage returns a boolean if a field has been set.

### GetSplashPrepaidFront

`func (o *InlineResponse20078) GetSplashPrepaidFront() InlineResponse20078SplashPrepaidFront`

GetSplashPrepaidFront returns the SplashPrepaidFront field if non-nil, zero value otherwise.

### GetSplashPrepaidFrontOk

`func (o *InlineResponse20078) GetSplashPrepaidFrontOk() (*InlineResponse20078SplashPrepaidFront, bool)`

GetSplashPrepaidFrontOk returns a tuple with the SplashPrepaidFront field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplashPrepaidFront

`func (o *InlineResponse20078) SetSplashPrepaidFront(v InlineResponse20078SplashPrepaidFront)`

SetSplashPrepaidFront sets SplashPrepaidFront field to given value.

### HasSplashPrepaidFront

`func (o *InlineResponse20078) HasSplashPrepaidFront() bool`

HasSplashPrepaidFront returns a boolean if a field has been set.

### GetGuestSponsorship

`func (o *InlineResponse20078) GetGuestSponsorship() InlineResponse20078GuestSponsorship`

GetGuestSponsorship returns the GuestSponsorship field if non-nil, zero value otherwise.

### GetGuestSponsorshipOk

`func (o *InlineResponse20078) GetGuestSponsorshipOk() (*InlineResponse20078GuestSponsorship, bool)`

GetGuestSponsorshipOk returns a tuple with the GuestSponsorship field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestSponsorship

`func (o *InlineResponse20078) SetGuestSponsorship(v InlineResponse20078GuestSponsorship)`

SetGuestSponsorship sets GuestSponsorship field to given value.

### HasGuestSponsorship

`func (o *InlineResponse20078) HasGuestSponsorship() bool`

HasGuestSponsorship returns a boolean if a field has been set.

### GetBlockAllTrafficBeforeSignOn

`func (o *InlineResponse20078) GetBlockAllTrafficBeforeSignOn() bool`

GetBlockAllTrafficBeforeSignOn returns the BlockAllTrafficBeforeSignOn field if non-nil, zero value otherwise.

### GetBlockAllTrafficBeforeSignOnOk

`func (o *InlineResponse20078) GetBlockAllTrafficBeforeSignOnOk() (*bool, bool)`

GetBlockAllTrafficBeforeSignOnOk returns a tuple with the BlockAllTrafficBeforeSignOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockAllTrafficBeforeSignOn

`func (o *InlineResponse20078) SetBlockAllTrafficBeforeSignOn(v bool)`

SetBlockAllTrafficBeforeSignOn sets BlockAllTrafficBeforeSignOn field to given value.

### HasBlockAllTrafficBeforeSignOn

`func (o *InlineResponse20078) HasBlockAllTrafficBeforeSignOn() bool`

HasBlockAllTrafficBeforeSignOn returns a boolean if a field has been set.

### GetControllerDisconnectionBehavior

`func (o *InlineResponse20078) GetControllerDisconnectionBehavior() string`

GetControllerDisconnectionBehavior returns the ControllerDisconnectionBehavior field if non-nil, zero value otherwise.

### GetControllerDisconnectionBehaviorOk

`func (o *InlineResponse20078) GetControllerDisconnectionBehaviorOk() (*string, bool)`

GetControllerDisconnectionBehaviorOk returns a tuple with the ControllerDisconnectionBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerDisconnectionBehavior

`func (o *InlineResponse20078) SetControllerDisconnectionBehavior(v string)`

SetControllerDisconnectionBehavior sets ControllerDisconnectionBehavior field to given value.

### HasControllerDisconnectionBehavior

`func (o *InlineResponse20078) HasControllerDisconnectionBehavior() bool`

HasControllerDisconnectionBehavior returns a boolean if a field has been set.

### GetAllowSimultaneousLogins

`func (o *InlineResponse20078) GetAllowSimultaneousLogins() bool`

GetAllowSimultaneousLogins returns the AllowSimultaneousLogins field if non-nil, zero value otherwise.

### GetAllowSimultaneousLoginsOk

`func (o *InlineResponse20078) GetAllowSimultaneousLoginsOk() (*bool, bool)`

GetAllowSimultaneousLoginsOk returns a tuple with the AllowSimultaneousLogins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSimultaneousLogins

`func (o *InlineResponse20078) SetAllowSimultaneousLogins(v bool)`

SetAllowSimultaneousLogins sets AllowSimultaneousLogins field to given value.

### HasAllowSimultaneousLogins

`func (o *InlineResponse20078) HasAllowSimultaneousLogins() bool`

HasAllowSimultaneousLogins returns a boolean if a field has been set.

### GetBilling

`func (o *InlineResponse20078) GetBilling() InlineResponse20078Billing`

GetBilling returns the Billing field if non-nil, zero value otherwise.

### GetBillingOk

`func (o *InlineResponse20078) GetBillingOk() (*InlineResponse20078Billing, bool)`

GetBillingOk returns a tuple with the Billing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBilling

`func (o *InlineResponse20078) SetBilling(v InlineResponse20078Billing)`

SetBilling sets Billing field to given value.

### HasBilling

`func (o *InlineResponse20078) HasBilling() bool`

HasBilling returns a boolean if a field has been set.

### GetSentryEnrollment

`func (o *InlineResponse20078) GetSentryEnrollment() InlineResponse20078SentryEnrollment`

GetSentryEnrollment returns the SentryEnrollment field if non-nil, zero value otherwise.

### GetSentryEnrollmentOk

`func (o *InlineResponse20078) GetSentryEnrollmentOk() (*InlineResponse20078SentryEnrollment, bool)`

GetSentryEnrollmentOk returns a tuple with the SentryEnrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentryEnrollment

`func (o *InlineResponse20078) SetSentryEnrollment(v InlineResponse20078SentryEnrollment)`

SetSentryEnrollment sets SentryEnrollment field to given value.

### HasSentryEnrollment

`func (o *InlineResponse20078) HasSentryEnrollment() bool`

HasSentryEnrollment returns a boolean if a field has been set.

### GetSelfRegistration

`func (o *InlineResponse20078) GetSelfRegistration() InlineResponse20078SelfRegistration`

GetSelfRegistration returns the SelfRegistration field if non-nil, zero value otherwise.

### GetSelfRegistrationOk

`func (o *InlineResponse20078) GetSelfRegistrationOk() (*InlineResponse20078SelfRegistration, bool)`

GetSelfRegistrationOk returns a tuple with the SelfRegistration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfRegistration

`func (o *InlineResponse20078) SetSelfRegistration(v InlineResponse20078SelfRegistration)`

SetSelfRegistration sets SelfRegistration field to given value.

### HasSelfRegistration

`func (o *InlineResponse20078) HasSelfRegistration() bool`

HasSelfRegistration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

