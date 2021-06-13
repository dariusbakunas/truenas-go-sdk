# RsynctaskUpdate1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to **string** |  | [optional] 
**User** | Pointer to **string** |  | [optional] 
**Remotehost** | Pointer to **string** |  | [optional] 
**Remoteport** | Pointer to **int32** |  | [optional] 
**Mode** | Pointer to **string** |  | [optional] 
**Remotemodule** | Pointer to **string** |  | [optional] 
**Remotepath** | Pointer to **string** |  | [optional] 
**ValidateRpath** | Pointer to **bool** |  | [optional] 
**Direction** | Pointer to **string** |  | [optional] 
**Desc** | Pointer to **string** |  | [optional] 
**Schedule** | Pointer to [**CloudsyncCreate0Schedule**](CloudsyncCreate0Schedule.md) |  | [optional] 
**Recursive** | Pointer to **bool** |  | [optional] 
**Times** | Pointer to **bool** |  | [optional] 
**Compress** | Pointer to **bool** |  | [optional] 
**Archive** | Pointer to **bool** |  | [optional] 
**Delete** | Pointer to **bool** |  | [optional] 
**Quiet** | Pointer to **bool** |  | [optional] 
**Preserveperm** | Pointer to **bool** |  | [optional] 
**Preserveattr** | Pointer to **bool** |  | [optional] 
**Delayupdates** | Pointer to **bool** |  | [optional] 
**Extra** | Pointer to **[]string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewRsynctaskUpdate1

`func NewRsynctaskUpdate1() *RsynctaskUpdate1`

NewRsynctaskUpdate1 instantiates a new RsynctaskUpdate1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRsynctaskUpdate1WithDefaults

`func NewRsynctaskUpdate1WithDefaults() *RsynctaskUpdate1`

NewRsynctaskUpdate1WithDefaults instantiates a new RsynctaskUpdate1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *RsynctaskUpdate1) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *RsynctaskUpdate1) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *RsynctaskUpdate1) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *RsynctaskUpdate1) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetUser

`func (o *RsynctaskUpdate1) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *RsynctaskUpdate1) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *RsynctaskUpdate1) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *RsynctaskUpdate1) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetRemotehost

`func (o *RsynctaskUpdate1) GetRemotehost() string`

GetRemotehost returns the Remotehost field if non-nil, zero value otherwise.

### GetRemotehostOk

`func (o *RsynctaskUpdate1) GetRemotehostOk() (*string, bool)`

GetRemotehostOk returns a tuple with the Remotehost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotehost

`func (o *RsynctaskUpdate1) SetRemotehost(v string)`

SetRemotehost sets Remotehost field to given value.

### HasRemotehost

`func (o *RsynctaskUpdate1) HasRemotehost() bool`

HasRemotehost returns a boolean if a field has been set.

### GetRemoteport

`func (o *RsynctaskUpdate1) GetRemoteport() int32`

GetRemoteport returns the Remoteport field if non-nil, zero value otherwise.

### GetRemoteportOk

`func (o *RsynctaskUpdate1) GetRemoteportOk() (*int32, bool)`

GetRemoteportOk returns a tuple with the Remoteport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteport

`func (o *RsynctaskUpdate1) SetRemoteport(v int32)`

SetRemoteport sets Remoteport field to given value.

### HasRemoteport

`func (o *RsynctaskUpdate1) HasRemoteport() bool`

HasRemoteport returns a boolean if a field has been set.

### GetMode

`func (o *RsynctaskUpdate1) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *RsynctaskUpdate1) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *RsynctaskUpdate1) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *RsynctaskUpdate1) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetRemotemodule

`func (o *RsynctaskUpdate1) GetRemotemodule() string`

GetRemotemodule returns the Remotemodule field if non-nil, zero value otherwise.

### GetRemotemoduleOk

`func (o *RsynctaskUpdate1) GetRemotemoduleOk() (*string, bool)`

GetRemotemoduleOk returns a tuple with the Remotemodule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotemodule

`func (o *RsynctaskUpdate1) SetRemotemodule(v string)`

SetRemotemodule sets Remotemodule field to given value.

### HasRemotemodule

`func (o *RsynctaskUpdate1) HasRemotemodule() bool`

HasRemotemodule returns a boolean if a field has been set.

### GetRemotepath

`func (o *RsynctaskUpdate1) GetRemotepath() string`

GetRemotepath returns the Remotepath field if non-nil, zero value otherwise.

### GetRemotepathOk

`func (o *RsynctaskUpdate1) GetRemotepathOk() (*string, bool)`

GetRemotepathOk returns a tuple with the Remotepath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotepath

`func (o *RsynctaskUpdate1) SetRemotepath(v string)`

SetRemotepath sets Remotepath field to given value.

### HasRemotepath

`func (o *RsynctaskUpdate1) HasRemotepath() bool`

HasRemotepath returns a boolean if a field has been set.

### GetValidateRpath

`func (o *RsynctaskUpdate1) GetValidateRpath() bool`

GetValidateRpath returns the ValidateRpath field if non-nil, zero value otherwise.

### GetValidateRpathOk

`func (o *RsynctaskUpdate1) GetValidateRpathOk() (*bool, bool)`

GetValidateRpathOk returns a tuple with the ValidateRpath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidateRpath

`func (o *RsynctaskUpdate1) SetValidateRpath(v bool)`

SetValidateRpath sets ValidateRpath field to given value.

### HasValidateRpath

`func (o *RsynctaskUpdate1) HasValidateRpath() bool`

HasValidateRpath returns a boolean if a field has been set.

### GetDirection

`func (o *RsynctaskUpdate1) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *RsynctaskUpdate1) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *RsynctaskUpdate1) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *RsynctaskUpdate1) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetDesc

`func (o *RsynctaskUpdate1) GetDesc() string`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *RsynctaskUpdate1) GetDescOk() (*string, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *RsynctaskUpdate1) SetDesc(v string)`

SetDesc sets Desc field to given value.

### HasDesc

`func (o *RsynctaskUpdate1) HasDesc() bool`

HasDesc returns a boolean if a field has been set.

### GetSchedule

`func (o *RsynctaskUpdate1) GetSchedule() CloudsyncCreate0Schedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *RsynctaskUpdate1) GetScheduleOk() (*CloudsyncCreate0Schedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *RsynctaskUpdate1) SetSchedule(v CloudsyncCreate0Schedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *RsynctaskUpdate1) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetRecursive

`func (o *RsynctaskUpdate1) GetRecursive() bool`

GetRecursive returns the Recursive field if non-nil, zero value otherwise.

### GetRecursiveOk

`func (o *RsynctaskUpdate1) GetRecursiveOk() (*bool, bool)`

GetRecursiveOk returns a tuple with the Recursive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecursive

`func (o *RsynctaskUpdate1) SetRecursive(v bool)`

SetRecursive sets Recursive field to given value.

### HasRecursive

`func (o *RsynctaskUpdate1) HasRecursive() bool`

HasRecursive returns a boolean if a field has been set.

### GetTimes

`func (o *RsynctaskUpdate1) GetTimes() bool`

GetTimes returns the Times field if non-nil, zero value otherwise.

### GetTimesOk

`func (o *RsynctaskUpdate1) GetTimesOk() (*bool, bool)`

GetTimesOk returns a tuple with the Times field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimes

`func (o *RsynctaskUpdate1) SetTimes(v bool)`

SetTimes sets Times field to given value.

### HasTimes

`func (o *RsynctaskUpdate1) HasTimes() bool`

HasTimes returns a boolean if a field has been set.

### GetCompress

`func (o *RsynctaskUpdate1) GetCompress() bool`

GetCompress returns the Compress field if non-nil, zero value otherwise.

### GetCompressOk

`func (o *RsynctaskUpdate1) GetCompressOk() (*bool, bool)`

GetCompressOk returns a tuple with the Compress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompress

`func (o *RsynctaskUpdate1) SetCompress(v bool)`

SetCompress sets Compress field to given value.

### HasCompress

`func (o *RsynctaskUpdate1) HasCompress() bool`

HasCompress returns a boolean if a field has been set.

### GetArchive

`func (o *RsynctaskUpdate1) GetArchive() bool`

GetArchive returns the Archive field if non-nil, zero value otherwise.

### GetArchiveOk

`func (o *RsynctaskUpdate1) GetArchiveOk() (*bool, bool)`

GetArchiveOk returns a tuple with the Archive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchive

`func (o *RsynctaskUpdate1) SetArchive(v bool)`

SetArchive sets Archive field to given value.

### HasArchive

`func (o *RsynctaskUpdate1) HasArchive() bool`

HasArchive returns a boolean if a field has been set.

### GetDelete

`func (o *RsynctaskUpdate1) GetDelete() bool`

GetDelete returns the Delete field if non-nil, zero value otherwise.

### GetDeleteOk

`func (o *RsynctaskUpdate1) GetDeleteOk() (*bool, bool)`

GetDeleteOk returns a tuple with the Delete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelete

`func (o *RsynctaskUpdate1) SetDelete(v bool)`

SetDelete sets Delete field to given value.

### HasDelete

`func (o *RsynctaskUpdate1) HasDelete() bool`

HasDelete returns a boolean if a field has been set.

### GetQuiet

`func (o *RsynctaskUpdate1) GetQuiet() bool`

GetQuiet returns the Quiet field if non-nil, zero value otherwise.

### GetQuietOk

`func (o *RsynctaskUpdate1) GetQuietOk() (*bool, bool)`

GetQuietOk returns a tuple with the Quiet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuiet

`func (o *RsynctaskUpdate1) SetQuiet(v bool)`

SetQuiet sets Quiet field to given value.

### HasQuiet

`func (o *RsynctaskUpdate1) HasQuiet() bool`

HasQuiet returns a boolean if a field has been set.

### GetPreserveperm

`func (o *RsynctaskUpdate1) GetPreserveperm() bool`

GetPreserveperm returns the Preserveperm field if non-nil, zero value otherwise.

### GetPreservepermOk

`func (o *RsynctaskUpdate1) GetPreservepermOk() (*bool, bool)`

GetPreservepermOk returns a tuple with the Preserveperm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveperm

`func (o *RsynctaskUpdate1) SetPreserveperm(v bool)`

SetPreserveperm sets Preserveperm field to given value.

### HasPreserveperm

`func (o *RsynctaskUpdate1) HasPreserveperm() bool`

HasPreserveperm returns a boolean if a field has been set.

### GetPreserveattr

`func (o *RsynctaskUpdate1) GetPreserveattr() bool`

GetPreserveattr returns the Preserveattr field if non-nil, zero value otherwise.

### GetPreserveattrOk

`func (o *RsynctaskUpdate1) GetPreserveattrOk() (*bool, bool)`

GetPreserveattrOk returns a tuple with the Preserveattr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveattr

`func (o *RsynctaskUpdate1) SetPreserveattr(v bool)`

SetPreserveattr sets Preserveattr field to given value.

### HasPreserveattr

`func (o *RsynctaskUpdate1) HasPreserveattr() bool`

HasPreserveattr returns a boolean if a field has been set.

### GetDelayupdates

`func (o *RsynctaskUpdate1) GetDelayupdates() bool`

GetDelayupdates returns the Delayupdates field if non-nil, zero value otherwise.

### GetDelayupdatesOk

`func (o *RsynctaskUpdate1) GetDelayupdatesOk() (*bool, bool)`

GetDelayupdatesOk returns a tuple with the Delayupdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayupdates

`func (o *RsynctaskUpdate1) SetDelayupdates(v bool)`

SetDelayupdates sets Delayupdates field to given value.

### HasDelayupdates

`func (o *RsynctaskUpdate1) HasDelayupdates() bool`

HasDelayupdates returns a boolean if a field has been set.

### GetExtra

`func (o *RsynctaskUpdate1) GetExtra() []string`

GetExtra returns the Extra field if non-nil, zero value otherwise.

### GetExtraOk

`func (o *RsynctaskUpdate1) GetExtraOk() (*[]string, bool)`

GetExtraOk returns a tuple with the Extra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtra

`func (o *RsynctaskUpdate1) SetExtra(v []string)`

SetExtra sets Extra field to given value.

### HasExtra

`func (o *RsynctaskUpdate1) HasExtra() bool`

HasExtra returns a boolean if a field has been set.

### GetEnabled

`func (o *RsynctaskUpdate1) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *RsynctaskUpdate1) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *RsynctaskUpdate1) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *RsynctaskUpdate1) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


