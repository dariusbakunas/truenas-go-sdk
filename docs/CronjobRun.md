# CronjobRun

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**SkipDisabled** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewCronjobRun

`func NewCronjobRun() *CronjobRun`

NewCronjobRun instantiates a new CronjobRun object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCronjobRunWithDefaults

`func NewCronjobRunWithDefaults() *CronjobRun`

NewCronjobRunWithDefaults instantiates a new CronjobRun object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CronjobRun) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CronjobRun) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CronjobRun) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CronjobRun) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSkipDisabled

`func (o *CronjobRun) GetSkipDisabled() bool`

GetSkipDisabled returns the SkipDisabled field if non-nil, zero value otherwise.

### GetSkipDisabledOk

`func (o *CronjobRun) GetSkipDisabledOk() (*bool, bool)`

GetSkipDisabledOk returns a tuple with the SkipDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipDisabled

`func (o *CronjobRun) SetSkipDisabled(v bool)`

SetSkipDisabled sets SkipDisabled field to given value.

### HasSkipDisabled

`func (o *CronjobRun) HasSkipDisabled() bool`

HasSkipDisabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


