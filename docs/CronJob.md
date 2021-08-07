# CronJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**User** | Pointer to **string** |  | [optional] 
**Command** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Stdout** | Pointer to **bool** |  | [optional] 
**Stderr** | Pointer to **bool** |  | [optional] 
**Schedule** | Pointer to [**CronJobSchedule**](CronJobSchedule.md) |  | [optional] 

## Methods

### NewCronJob

`func NewCronJob() *CronJob`

NewCronJob instantiates a new CronJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCronJobWithDefaults

`func NewCronJobWithDefaults() *CronJob`

NewCronJobWithDefaults instantiates a new CronJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CronJob) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CronJob) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CronJob) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CronJob) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUser

`func (o *CronJob) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *CronJob) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *CronJob) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *CronJob) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetCommand

`func (o *CronJob) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *CronJob) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *CronJob) SetCommand(v string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *CronJob) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### GetDescription

`func (o *CronJob) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CronJob) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CronJob) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CronJob) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *CronJob) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CronJob) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CronJob) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CronJob) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetStdout

`func (o *CronJob) GetStdout() bool`

GetStdout returns the Stdout field if non-nil, zero value otherwise.

### GetStdoutOk

`func (o *CronJob) GetStdoutOk() (*bool, bool)`

GetStdoutOk returns a tuple with the Stdout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStdout

`func (o *CronJob) SetStdout(v bool)`

SetStdout sets Stdout field to given value.

### HasStdout

`func (o *CronJob) HasStdout() bool`

HasStdout returns a boolean if a field has been set.

### GetStderr

`func (o *CronJob) GetStderr() bool`

GetStderr returns the Stderr field if non-nil, zero value otherwise.

### GetStderrOk

`func (o *CronJob) GetStderrOk() (*bool, bool)`

GetStderrOk returns a tuple with the Stderr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStderr

`func (o *CronJob) SetStderr(v bool)`

SetStderr sets Stderr field to given value.

### HasStderr

`func (o *CronJob) HasStderr() bool`

HasStderr returns a boolean if a field has been set.

### GetSchedule

`func (o *CronJob) GetSchedule() CronJobSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *CronJob) GetScheduleOk() (*CronJobSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *CronJob) SetSchedule(v CronJobSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *CronJob) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


