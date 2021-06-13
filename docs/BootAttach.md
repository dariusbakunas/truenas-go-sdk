# BootAttach

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dev** | Pointer to **string** |  | [optional] 
**Options** | Pointer to [**BootAttach1**](BootAttach1.md) |  | [optional] [default to {}]

## Methods

### NewBootAttach

`func NewBootAttach() *BootAttach`

NewBootAttach instantiates a new BootAttach object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBootAttachWithDefaults

`func NewBootAttachWithDefaults() *BootAttach`

NewBootAttachWithDefaults instantiates a new BootAttach object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDev

`func (o *BootAttach) GetDev() string`

GetDev returns the Dev field if non-nil, zero value otherwise.

### GetDevOk

`func (o *BootAttach) GetDevOk() (*string, bool)`

GetDevOk returns a tuple with the Dev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDev

`func (o *BootAttach) SetDev(v string)`

SetDev sets Dev field to given value.

### HasDev

`func (o *BootAttach) HasDev() bool`

HasDev returns a boolean if a field has been set.

### GetOptions

`func (o *BootAttach) GetOptions() BootAttach1`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *BootAttach) GetOptionsOk() (*BootAttach1, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *BootAttach) SetOptions(v BootAttach1)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *BootAttach) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


