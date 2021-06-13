# CoreJobUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**JobUpdate** | Pointer to [**CoreJobUpdate1**](CoreJobUpdate1.md) |  | [optional] [default to {}]

## Methods

### NewCoreJobUpdate

`func NewCoreJobUpdate() *CoreJobUpdate`

NewCoreJobUpdate instantiates a new CoreJobUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCoreJobUpdateWithDefaults

`func NewCoreJobUpdateWithDefaults() *CoreJobUpdate`

NewCoreJobUpdateWithDefaults instantiates a new CoreJobUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CoreJobUpdate) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CoreJobUpdate) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CoreJobUpdate) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CoreJobUpdate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetJobUpdate

`func (o *CoreJobUpdate) GetJobUpdate() CoreJobUpdate1`

GetJobUpdate returns the JobUpdate field if non-nil, zero value otherwise.

### GetJobUpdateOk

`func (o *CoreJobUpdate) GetJobUpdateOk() (*CoreJobUpdate1, bool)`

GetJobUpdateOk returns a tuple with the JobUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobUpdate

`func (o *CoreJobUpdate) SetJobUpdate(v CoreJobUpdate1)`

SetJobUpdate sets JobUpdate field to given value.

### HasJobUpdate

`func (o *CoreJobUpdate) HasJobUpdate() bool`

HasJobUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


