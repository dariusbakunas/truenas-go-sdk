# CronjobCreate0

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** |  | [optional] 
**Stderr** | Pointer to **bool** |  | [optional] 
**Stdout** | Pointer to **bool** |  | [optional] 
**Schedule** | Pointer to [**CloudsyncCreate0Schedule**](CloudsyncCreate0Schedule.md) |  | [optional] 
**Command** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**User** | Pointer to **string** |  | [optional] 

## Methods

### NewCronjobCreate0

`func NewCronjobCreate0() *CronjobCreate0`

NewCronjobCreate0 instantiates a new CronjobCreate0 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCronjobCreate0WithDefaults

`func NewCronjobCreate0WithDefaults() *CronjobCreate0`

NewCronjobCreate0WithDefaults instantiates a new CronjobCreate0 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *CronjobCreate0) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CronjobCreate0) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CronjobCreate0) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CronjobCreate0) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetStderr

`func (o *CronjobCreate0) GetStderr() bool`

GetStderr returns the Stderr field if non-nil, zero value otherwise.

### GetStderrOk

`func (o *CronjobCreate0) GetStderrOk() (*bool, bool)`

GetStderrOk returns a tuple with the Stderr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStderr

`func (o *CronjobCreate0) SetStderr(v bool)`

SetStderr sets Stderr field to given value.

### HasStderr

`func (o *CronjobCreate0) HasStderr() bool`

HasStderr returns a boolean if a field has been set.

### GetStdout

`func (o *CronjobCreate0) GetStdout() bool`

GetStdout returns the Stdout field if non-nil, zero value otherwise.

### GetStdoutOk

`func (o *CronjobCreate0) GetStdoutOk() (*bool, bool)`

GetStdoutOk returns a tuple with the Stdout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStdout

`func (o *CronjobCreate0) SetStdout(v bool)`

SetStdout sets Stdout field to given value.

### HasStdout

`func (o *CronjobCreate0) HasStdout() bool`

HasStdout returns a boolean if a field has been set.

### GetSchedule

`func (o *CronjobCreate0) GetSchedule() CloudsyncCreate0Schedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *CronjobCreate0) GetScheduleOk() (*CloudsyncCreate0Schedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *CronjobCreate0) SetSchedule(v CloudsyncCreate0Schedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *CronjobCreate0) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetCommand

`func (o *CronjobCreate0) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *CronjobCreate0) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *CronjobCreate0) SetCommand(v string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *CronjobCreate0) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### GetDescription

`func (o *CronjobCreate0) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CronjobCreate0) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CronjobCreate0) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CronjobCreate0) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUser

`func (o *CronjobCreate0) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *CronjobCreate0) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *CronjobCreate0) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *CronjobCreate0) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


