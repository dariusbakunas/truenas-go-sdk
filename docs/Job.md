# Job

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Abortable** | Pointer to **bool** |  | [optional] 
**Arguments** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**Exception** | Pointer to **string** |  | [optional] 
**LogsExcerpt** | Pointer to **string** |  | [optional] 
**LogsPath** | Pointer to **string** |  | [optional] 
**Method** | Pointer to **string** |  | [optional] 
**Progress** | Pointer to [**JobProgress**](JobProgress.md) |  | [optional] 
**Result** | Pointer to **map[string]interface{}** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 

## Methods

### NewJob

`func NewJob(id int32, ) *Job`

NewJob instantiates a new Job object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobWithDefaults

`func NewJobWithDefaults() *Job`

NewJobWithDefaults instantiates a new Job object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Job) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Job) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Job) SetId(v int32)`

SetId sets Id field to given value.


### GetAbortable

`func (o *Job) GetAbortable() bool`

GetAbortable returns the Abortable field if non-nil, zero value otherwise.

### GetAbortableOk

`func (o *Job) GetAbortableOk() (*bool, bool)`

GetAbortableOk returns a tuple with the Abortable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbortable

`func (o *Job) SetAbortable(v bool)`

SetAbortable sets Abortable field to given value.

### HasAbortable

`func (o *Job) HasAbortable() bool`

HasAbortable returns a boolean if a field has been set.

### GetArguments

`func (o *Job) GetArguments() []map[string]interface{}`

GetArguments returns the Arguments field if non-nil, zero value otherwise.

### GetArgumentsOk

`func (o *Job) GetArgumentsOk() (*[]map[string]interface{}, bool)`

GetArgumentsOk returns a tuple with the Arguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArguments

`func (o *Job) SetArguments(v []map[string]interface{})`

SetArguments sets Arguments field to given value.

### HasArguments

`func (o *Job) HasArguments() bool`

HasArguments returns a boolean if a field has been set.

### GetDescription

`func (o *Job) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Job) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Job) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Job) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetError

`func (o *Job) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *Job) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *Job) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *Job) HasError() bool`

HasError returns a boolean if a field has been set.

### GetException

`func (o *Job) GetException() string`

GetException returns the Exception field if non-nil, zero value otherwise.

### GetExceptionOk

`func (o *Job) GetExceptionOk() (*string, bool)`

GetExceptionOk returns a tuple with the Exception field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetException

`func (o *Job) SetException(v string)`

SetException sets Exception field to given value.

### HasException

`func (o *Job) HasException() bool`

HasException returns a boolean if a field has been set.

### GetLogsExcerpt

`func (o *Job) GetLogsExcerpt() string`

GetLogsExcerpt returns the LogsExcerpt field if non-nil, zero value otherwise.

### GetLogsExcerptOk

`func (o *Job) GetLogsExcerptOk() (*string, bool)`

GetLogsExcerptOk returns a tuple with the LogsExcerpt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogsExcerpt

`func (o *Job) SetLogsExcerpt(v string)`

SetLogsExcerpt sets LogsExcerpt field to given value.

### HasLogsExcerpt

`func (o *Job) HasLogsExcerpt() bool`

HasLogsExcerpt returns a boolean if a field has been set.

### GetLogsPath

`func (o *Job) GetLogsPath() string`

GetLogsPath returns the LogsPath field if non-nil, zero value otherwise.

### GetLogsPathOk

`func (o *Job) GetLogsPathOk() (*string, bool)`

GetLogsPathOk returns a tuple with the LogsPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogsPath

`func (o *Job) SetLogsPath(v string)`

SetLogsPath sets LogsPath field to given value.

### HasLogsPath

`func (o *Job) HasLogsPath() bool`

HasLogsPath returns a boolean if a field has been set.

### GetMethod

`func (o *Job) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *Job) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *Job) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *Job) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetProgress

`func (o *Job) GetProgress() JobProgress`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *Job) GetProgressOk() (*JobProgress, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *Job) SetProgress(v JobProgress)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *Job) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetResult

`func (o *Job) GetResult() map[string]interface{}`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *Job) GetResultOk() (*map[string]interface{}, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *Job) SetResult(v map[string]interface{})`

SetResult sets Result field to given value.

### HasResult

`func (o *Job) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetState

`func (o *Job) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Job) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Job) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Job) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


