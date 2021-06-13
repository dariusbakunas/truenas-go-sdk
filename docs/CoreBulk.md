# CoreBulk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Method** | Pointer to **string** |  | [optional] 
**Params** | Pointer to **[]interface{}** |  | [optional] [default to []]
**Description** | Pointer to **NullableString** |  | [optional] [default to "null"]

## Methods

### NewCoreBulk

`func NewCoreBulk() *CoreBulk`

NewCoreBulk instantiates a new CoreBulk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCoreBulkWithDefaults

`func NewCoreBulkWithDefaults() *CoreBulk`

NewCoreBulkWithDefaults instantiates a new CoreBulk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethod

`func (o *CoreBulk) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *CoreBulk) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *CoreBulk) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *CoreBulk) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetParams

`func (o *CoreBulk) GetParams() []interface{}`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *CoreBulk) GetParamsOk() (*[]interface{}, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *CoreBulk) SetParams(v []interface{})`

SetParams sets Params field to given value.

### HasParams

`func (o *CoreBulk) HasParams() bool`

HasParams returns a boolean if a field has been set.

### GetDescription

`func (o *CoreBulk) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CoreBulk) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CoreBulk) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CoreBulk) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CoreBulk) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CoreBulk) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


