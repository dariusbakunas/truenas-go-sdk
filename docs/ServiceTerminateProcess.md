# ServiceTerminateProcess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pid** | Pointer to **int32** |  | [optional] 
**Timeout** | Pointer to **int32** |  | [optional] [default to 10]

## Methods

### NewServiceTerminateProcess

`func NewServiceTerminateProcess() *ServiceTerminateProcess`

NewServiceTerminateProcess instantiates a new ServiceTerminateProcess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceTerminateProcessWithDefaults

`func NewServiceTerminateProcessWithDefaults() *ServiceTerminateProcess`

NewServiceTerminateProcessWithDefaults instantiates a new ServiceTerminateProcess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPid

`func (o *ServiceTerminateProcess) GetPid() int32`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *ServiceTerminateProcess) GetPidOk() (*int32, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *ServiceTerminateProcess) SetPid(v int32)`

SetPid sets Pid field to given value.

### HasPid

`func (o *ServiceTerminateProcess) HasPid() bool`

HasPid returns a boolean if a field has been set.

### GetTimeout

`func (o *ServiceTerminateProcess) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *ServiceTerminateProcess) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *ServiceTerminateProcess) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *ServiceTerminateProcess) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


