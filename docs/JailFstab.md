# JailFstab

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Jail** | Pointer to **string** |  | [optional] 
**Options** | Pointer to [**JailFstab1**](JailFstab1.md) |  | [optional] [default to {}]

## Methods

### NewJailFstab

`func NewJailFstab() *JailFstab`

NewJailFstab instantiates a new JailFstab object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJailFstabWithDefaults

`func NewJailFstabWithDefaults() *JailFstab`

NewJailFstabWithDefaults instantiates a new JailFstab object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJail

`func (o *JailFstab) GetJail() string`

GetJail returns the Jail field if non-nil, zero value otherwise.

### GetJailOk

`func (o *JailFstab) GetJailOk() (*string, bool)`

GetJailOk returns a tuple with the Jail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJail

`func (o *JailFstab) SetJail(v string)`

SetJail sets Jail field to given value.

### HasJail

`func (o *JailFstab) HasJail() bool`

HasJail returns a boolean if a field has been set.

### GetOptions

`func (o *JailFstab) GetOptions() JailFstab1`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *JailFstab) GetOptionsOk() (*JailFstab1, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *JailFstab) SetOptions(v JailFstab1)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *JailFstab) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


