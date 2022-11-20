# JobProgress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Percent** | Pointer to **float32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Extra** | Pointer to **string** |  | [optional] 

## Methods

### NewJobProgress

`func NewJobProgress() *JobProgress`

NewJobProgress instantiates a new JobProgress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobProgressWithDefaults

`func NewJobProgressWithDefaults() *JobProgress`

NewJobProgressWithDefaults instantiates a new JobProgress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPercent

`func (o *JobProgress) GetPercent() float32`

GetPercent returns the Percent field if non-nil, zero value otherwise.

### GetPercentOk

`func (o *JobProgress) GetPercentOk() (*float32, bool)`

GetPercentOk returns a tuple with the Percent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercent

`func (o *JobProgress) SetPercent(v float32)`

SetPercent sets Percent field to given value.

### HasPercent

`func (o *JobProgress) HasPercent() bool`

HasPercent returns a boolean if a field has been set.

### GetDescription

`func (o *JobProgress) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *JobProgress) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *JobProgress) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *JobProgress) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExtra

`func (o *JobProgress) GetExtra() string`

GetExtra returns the Extra field if non-nil, zero value otherwise.

### GetExtraOk

`func (o *JobProgress) GetExtraOk() (*string, bool)`

GetExtraOk returns a tuple with the Extra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtra

`func (o *JobProgress) SetExtra(v string)`

SetExtra sets Extra field to given value.

### HasExtra

`func (o *JobProgress) HasExtra() bool`

HasExtra returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


