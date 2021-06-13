# DyndnsUpdate0

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Provider** | Pointer to **string** |  | [optional] 
**CheckipSsl** | Pointer to **bool** |  | [optional] 
**CheckipServer** | Pointer to **string** |  | [optional] 
**CheckipPath** | Pointer to **string** |  | [optional] 
**Ssl** | Pointer to **bool** |  | [optional] 
**CustomDdnsServer** | Pointer to **string** |  | [optional] 
**CustomDdnsPath** | Pointer to **string** |  | [optional] 
**Domain** | Pointer to **[]string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**Period** | Pointer to **int32** |  | [optional] 

## Methods

### NewDyndnsUpdate0

`func NewDyndnsUpdate0() *DyndnsUpdate0`

NewDyndnsUpdate0 instantiates a new DyndnsUpdate0 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDyndnsUpdate0WithDefaults

`func NewDyndnsUpdate0WithDefaults() *DyndnsUpdate0`

NewDyndnsUpdate0WithDefaults instantiates a new DyndnsUpdate0 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvider

`func (o *DyndnsUpdate0) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *DyndnsUpdate0) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *DyndnsUpdate0) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *DyndnsUpdate0) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetCheckipSsl

`func (o *DyndnsUpdate0) GetCheckipSsl() bool`

GetCheckipSsl returns the CheckipSsl field if non-nil, zero value otherwise.

### GetCheckipSslOk

`func (o *DyndnsUpdate0) GetCheckipSslOk() (*bool, bool)`

GetCheckipSslOk returns a tuple with the CheckipSsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckipSsl

`func (o *DyndnsUpdate0) SetCheckipSsl(v bool)`

SetCheckipSsl sets CheckipSsl field to given value.

### HasCheckipSsl

`func (o *DyndnsUpdate0) HasCheckipSsl() bool`

HasCheckipSsl returns a boolean if a field has been set.

### GetCheckipServer

`func (o *DyndnsUpdate0) GetCheckipServer() string`

GetCheckipServer returns the CheckipServer field if non-nil, zero value otherwise.

### GetCheckipServerOk

`func (o *DyndnsUpdate0) GetCheckipServerOk() (*string, bool)`

GetCheckipServerOk returns a tuple with the CheckipServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckipServer

`func (o *DyndnsUpdate0) SetCheckipServer(v string)`

SetCheckipServer sets CheckipServer field to given value.

### HasCheckipServer

`func (o *DyndnsUpdate0) HasCheckipServer() bool`

HasCheckipServer returns a boolean if a field has been set.

### GetCheckipPath

`func (o *DyndnsUpdate0) GetCheckipPath() string`

GetCheckipPath returns the CheckipPath field if non-nil, zero value otherwise.

### GetCheckipPathOk

`func (o *DyndnsUpdate0) GetCheckipPathOk() (*string, bool)`

GetCheckipPathOk returns a tuple with the CheckipPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckipPath

`func (o *DyndnsUpdate0) SetCheckipPath(v string)`

SetCheckipPath sets CheckipPath field to given value.

### HasCheckipPath

`func (o *DyndnsUpdate0) HasCheckipPath() bool`

HasCheckipPath returns a boolean if a field has been set.

### GetSsl

`func (o *DyndnsUpdate0) GetSsl() bool`

GetSsl returns the Ssl field if non-nil, zero value otherwise.

### GetSslOk

`func (o *DyndnsUpdate0) GetSslOk() (*bool, bool)`

GetSslOk returns a tuple with the Ssl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsl

`func (o *DyndnsUpdate0) SetSsl(v bool)`

SetSsl sets Ssl field to given value.

### HasSsl

`func (o *DyndnsUpdate0) HasSsl() bool`

HasSsl returns a boolean if a field has been set.

### GetCustomDdnsServer

`func (o *DyndnsUpdate0) GetCustomDdnsServer() string`

GetCustomDdnsServer returns the CustomDdnsServer field if non-nil, zero value otherwise.

### GetCustomDdnsServerOk

`func (o *DyndnsUpdate0) GetCustomDdnsServerOk() (*string, bool)`

GetCustomDdnsServerOk returns a tuple with the CustomDdnsServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDdnsServer

`func (o *DyndnsUpdate0) SetCustomDdnsServer(v string)`

SetCustomDdnsServer sets CustomDdnsServer field to given value.

### HasCustomDdnsServer

`func (o *DyndnsUpdate0) HasCustomDdnsServer() bool`

HasCustomDdnsServer returns a boolean if a field has been set.

### GetCustomDdnsPath

`func (o *DyndnsUpdate0) GetCustomDdnsPath() string`

GetCustomDdnsPath returns the CustomDdnsPath field if non-nil, zero value otherwise.

### GetCustomDdnsPathOk

`func (o *DyndnsUpdate0) GetCustomDdnsPathOk() (*string, bool)`

GetCustomDdnsPathOk returns a tuple with the CustomDdnsPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDdnsPath

`func (o *DyndnsUpdate0) SetCustomDdnsPath(v string)`

SetCustomDdnsPath sets CustomDdnsPath field to given value.

### HasCustomDdnsPath

`func (o *DyndnsUpdate0) HasCustomDdnsPath() bool`

HasCustomDdnsPath returns a boolean if a field has been set.

### GetDomain

`func (o *DyndnsUpdate0) GetDomain() []string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *DyndnsUpdate0) GetDomainOk() (*[]string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *DyndnsUpdate0) SetDomain(v []string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *DyndnsUpdate0) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetUsername

`func (o *DyndnsUpdate0) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *DyndnsUpdate0) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *DyndnsUpdate0) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *DyndnsUpdate0) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *DyndnsUpdate0) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *DyndnsUpdate0) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *DyndnsUpdate0) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *DyndnsUpdate0) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPeriod

`func (o *DyndnsUpdate0) GetPeriod() int32`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *DyndnsUpdate0) GetPeriodOk() (*int32, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *DyndnsUpdate0) SetPeriod(v int32)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *DyndnsUpdate0) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


