# UserSetAttribute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** |  | [optional] 
**Value** | Pointer to [**UserSetAttribute2**](UserSetAttribute2.md) |  | [optional] 

## Methods

### NewUserSetAttribute

`func NewUserSetAttribute() *UserSetAttribute`

NewUserSetAttribute instantiates a new UserSetAttribute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSetAttributeWithDefaults

`func NewUserSetAttributeWithDefaults() *UserSetAttribute`

NewUserSetAttributeWithDefaults instantiates a new UserSetAttribute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *UserSetAttribute) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *UserSetAttribute) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *UserSetAttribute) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *UserSetAttribute) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetValue

`func (o *UserSetAttribute) GetValue() UserSetAttribute2`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *UserSetAttribute) GetValueOk() (*UserSetAttribute2, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *UserSetAttribute) SetValue(v UserSetAttribute2)`

SetValue sets Value field to given value.

### HasValue

`func (o *UserSetAttribute) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


