# JailStop

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Jail** | Pointer to **string** |  | [optional] 
**Force** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewJailStop

`func NewJailStop() *JailStop`

NewJailStop instantiates a new JailStop object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJailStopWithDefaults

`func NewJailStopWithDefaults() *JailStop`

NewJailStopWithDefaults instantiates a new JailStop object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJail

`func (o *JailStop) GetJail() string`

GetJail returns the Jail field if non-nil, zero value otherwise.

### GetJailOk

`func (o *JailStop) GetJailOk() (*string, bool)`

GetJailOk returns a tuple with the Jail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJail

`func (o *JailStop) SetJail(v string)`

SetJail sets Jail field to given value.

### HasJail

`func (o *JailStop) HasJail() bool`

HasJail returns a boolean if a field has been set.

### GetForce

`func (o *JailStop) GetForce() bool`

GetForce returns the Force field if non-nil, zero value otherwise.

### GetForceOk

`func (o *JailStop) GetForceOk() (*bool, bool)`

GetForceOk returns a tuple with the Force field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForce

`func (o *JailStop) SetForce(v bool)`

SetForce sets Force field to given value.

### HasForce

`func (o *JailStop) HasForce() bool`

HasForce returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


