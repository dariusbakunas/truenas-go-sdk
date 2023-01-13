# CreateCronjobParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | **string** |  | 
**Command** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Stdout** | Pointer to **bool** |  | [optional] 
**Stderr** | Pointer to **bool** |  | [optional] 
**Schedule** | Pointer to [**CronJobSchedule**](CronJobSchedule.md) |  | [optional] 

## Methods

### NewCreateCronjobParams

`func NewCreateCronjobParams(user string, command string, ) *CreateCronjobParams`

NewCreateCronjobParams instantiates a new CreateCronjobParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCronjobParamsWithDefaults

`func NewCreateCronjobParamsWithDefaults() *CreateCronjobParams`

NewCreateCronjobParamsWithDefaults instantiates a new CreateCronjobParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *CreateCronjobParams) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *CreateCronjobParams) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *CreateCronjobParams) SetUser(v string)`

SetUser sets User field to given value.


### GetCommand

`func (o *CreateCronjobParams) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *CreateCronjobParams) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *CreateCronjobParams) SetCommand(v string)`

SetCommand sets Command field to given value.


### GetDescription

`func (o *CreateCronjobParams) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateCronjobParams) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateCronjobParams) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateCronjobParams) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *CreateCronjobParams) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CreateCronjobParams) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CreateCronjobParams) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CreateCronjobParams) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetStdout

`func (o *CreateCronjobParams) GetStdout() bool`

GetStdout returns the Stdout field if non-nil, zero value otherwise.

### GetStdoutOk

`func (o *CreateCronjobParams) GetStdoutOk() (*bool, bool)`

GetStdoutOk returns a tuple with the Stdout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStdout

`func (o *CreateCronjobParams) SetStdout(v bool)`

SetStdout sets Stdout field to given value.

### HasStdout

`func (o *CreateCronjobParams) HasStdout() bool`

HasStdout returns a boolean if a field has been set.

### GetStderr

`func (o *CreateCronjobParams) GetStderr() bool`

GetStderr returns the Stderr field if non-nil, zero value otherwise.

### GetStderrOk

`func (o *CreateCronjobParams) GetStderrOk() (*bool, bool)`

GetStderrOk returns a tuple with the Stderr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStderr

`func (o *CreateCronjobParams) SetStderr(v bool)`

SetStderr sets Stderr field to given value.

### HasStderr

`func (o *CreateCronjobParams) HasStderr() bool`

HasStderr returns a boolean if a field has been set.

### GetSchedule

`func (o *CreateCronjobParams) GetSchedule() CronJobSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *CreateCronjobParams) GetScheduleOk() (*CronJobSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *CreateCronjobParams) SetSchedule(v CronJobSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *CreateCronjobParams) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


