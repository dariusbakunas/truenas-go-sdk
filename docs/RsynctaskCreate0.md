# RsynctaskCreate0

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

### NewRsynctaskCreate0

`func NewRsynctaskCreate0() *RsynctaskCreate0`

NewRsynctaskCreate0 instantiates a new RsynctaskCreate0 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRsynctaskCreate0WithDefaults

`func NewRsynctaskCreate0WithDefaults() *RsynctaskCreate0`

NewRsynctaskCreate0WithDefaults instantiates a new RsynctaskCreate0 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *RsynctaskCreate0) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *RsynctaskCreate0) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *RsynctaskCreate0) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *RsynctaskCreate0) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetUser

`func (o *RsynctaskCreate0) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *RsynctaskCreate0) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *RsynctaskCreate0) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *RsynctaskCreate0) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetRemotehost

`func (o *RsynctaskCreate0) GetRemotehost() string`

GetRemotehost returns the Remotehost field if non-nil, zero value otherwise.

### GetRemotehostOk

`func (o *RsynctaskCreate0) GetRemotehostOk() (*string, bool)`

GetRemotehostOk returns a tuple with the Remotehost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotehost

`func (o *RsynctaskCreate0) SetRemotehost(v string)`

SetRemotehost sets Remotehost field to given value.

### HasRemotehost

`func (o *RsynctaskCreate0) HasRemotehost() bool`

HasRemotehost returns a boolean if a field has been set.

### GetRemoteport

`func (o *RsynctaskCreate0) GetRemoteport() int32`

GetRemoteport returns the Remoteport field if non-nil, zero value otherwise.

### GetRemoteportOk

`func (o *RsynctaskCreate0) GetRemoteportOk() (*int32, bool)`

GetRemoteportOk returns a tuple with the Remoteport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteport

`func (o *RsynctaskCreate0) SetRemoteport(v int32)`

SetRemoteport sets Remoteport field to given value.

### HasRemoteport

`func (o *RsynctaskCreate0) HasRemoteport() bool`

HasRemoteport returns a boolean if a field has been set.

### GetMode

`func (o *RsynctaskCreate0) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *RsynctaskCreate0) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *RsynctaskCreate0) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *RsynctaskCreate0) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetRemotemodule

`func (o *RsynctaskCreate0) GetRemotemodule() string`

GetRemotemodule returns the Remotemodule field if non-nil, zero value otherwise.

### GetRemotemoduleOk

`func (o *RsynctaskCreate0) GetRemotemoduleOk() (*string, bool)`

GetRemotemoduleOk returns a tuple with the Remotemodule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotemodule

`func (o *RsynctaskCreate0) SetRemotemodule(v string)`

SetRemotemodule sets Remotemodule field to given value.

### HasRemotemodule

`func (o *RsynctaskCreate0) HasRemotemodule() bool`

HasRemotemodule returns a boolean if a field has been set.

### GetRemotepath

`func (o *RsynctaskCreate0) GetRemotepath() string`

GetRemotepath returns the Remotepath field if non-nil, zero value otherwise.

### GetRemotepathOk

`func (o *RsynctaskCreate0) GetRemotepathOk() (*string, bool)`

GetRemotepathOk returns a tuple with the Remotepath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotepath

`func (o *RsynctaskCreate0) SetRemotepath(v string)`

SetRemotepath sets Remotepath field to given value.

### HasRemotepath

`func (o *RsynctaskCreate0) HasRemotepath() bool`

HasRemotepath returns a boolean if a field has been set.

### GetValidateRpath

`func (o *RsynctaskCreate0) GetValidateRpath() bool`

GetValidateRpath returns the ValidateRpath field if non-nil, zero value otherwise.

### GetValidateRpathOk

`func (o *RsynctaskCreate0) GetValidateRpathOk() (*bool, bool)`

GetValidateRpathOk returns a tuple with the ValidateRpath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidateRpath

`func (o *RsynctaskCreate0) SetValidateRpath(v bool)`

SetValidateRpath sets ValidateRpath field to given value.

### HasValidateRpath

`func (o *RsynctaskCreate0) HasValidateRpath() bool`

HasValidateRpath returns a boolean if a field has been set.

### GetDirection

`func (o *RsynctaskCreate0) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *RsynctaskCreate0) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *RsynctaskCreate0) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *RsynctaskCreate0) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetDesc

`func (o *RsynctaskCreate0) GetDesc() string`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *RsynctaskCreate0) GetDescOk() (*string, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *RsynctaskCreate0) SetDesc(v string)`

SetDesc sets Desc field to given value.

### HasDesc

`func (o *RsynctaskCreate0) HasDesc() bool`

HasDesc returns a boolean if a field has been set.

### GetSchedule

`func (o *RsynctaskCreate0) GetSchedule() CloudsyncCreate0Schedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *RsynctaskCreate0) GetScheduleOk() (*CloudsyncCreate0Schedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *RsynctaskCreate0) SetSchedule(v CloudsyncCreate0Schedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *RsynctaskCreate0) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetRecursive

`func (o *RsynctaskCreate0) GetRecursive() bool`

GetRecursive returns the Recursive field if non-nil, zero value otherwise.

### GetRecursiveOk

`func (o *RsynctaskCreate0) GetRecursiveOk() (*bool, bool)`

GetRecursiveOk returns a tuple with the Recursive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecursive

`func (o *RsynctaskCreate0) SetRecursive(v bool)`

SetRecursive sets Recursive field to given value.

### HasRecursive

`func (o *RsynctaskCreate0) HasRecursive() bool`

HasRecursive returns a boolean if a field has been set.

### GetTimes

`func (o *RsynctaskCreate0) GetTimes() bool`

GetTimes returns the Times field if non-nil, zero value otherwise.

### GetTimesOk

`func (o *RsynctaskCreate0) GetTimesOk() (*bool, bool)`

GetTimesOk returns a tuple with the Times field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimes

`func (o *RsynctaskCreate0) SetTimes(v bool)`

SetTimes sets Times field to given value.

### HasTimes

`func (o *RsynctaskCreate0) HasTimes() bool`

HasTimes returns a boolean if a field has been set.

### GetCompress

`func (o *RsynctaskCreate0) GetCompress() bool`

GetCompress returns the Compress field if non-nil, zero value otherwise.

### GetCompressOk

`func (o *RsynctaskCreate0) GetCompressOk() (*bool, bool)`

GetCompressOk returns a tuple with the Compress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompress

`func (o *RsynctaskCreate0) SetCompress(v bool)`

SetCompress sets Compress field to given value.

### HasCompress

`func (o *RsynctaskCreate0) HasCompress() bool`

HasCompress returns a boolean if a field has been set.

### GetArchive

`func (o *RsynctaskCreate0) GetArchive() bool`

GetArchive returns the Archive field if non-nil, zero value otherwise.

### GetArchiveOk

`func (o *RsynctaskCreate0) GetArchiveOk() (*bool, bool)`

GetArchiveOk returns a tuple with the Archive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchive

`func (o *RsynctaskCreate0) SetArchive(v bool)`

SetArchive sets Archive field to given value.

### HasArchive

`func (o *RsynctaskCreate0) HasArchive() bool`

HasArchive returns a boolean if a field has been set.

### GetDelete

`func (o *RsynctaskCreate0) GetDelete() bool`

GetDelete returns the Delete field if non-nil, zero value otherwise.

### GetDeleteOk

`func (o *RsynctaskCreate0) GetDeleteOk() (*bool, bool)`

GetDeleteOk returns a tuple with the Delete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelete

`func (o *RsynctaskCreate0) SetDelete(v bool)`

SetDelete sets Delete field to given value.

### HasDelete

`func (o *RsynctaskCreate0) HasDelete() bool`

HasDelete returns a boolean if a field has been set.

### GetQuiet

`func (o *RsynctaskCreate0) GetQuiet() bool`

GetQuiet returns the Quiet field if non-nil, zero value otherwise.

### GetQuietOk

`func (o *RsynctaskCreate0) GetQuietOk() (*bool, bool)`

GetQuietOk returns a tuple with the Quiet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuiet

`func (o *RsynctaskCreate0) SetQuiet(v bool)`

SetQuiet sets Quiet field to given value.

### HasQuiet

`func (o *RsynctaskCreate0) HasQuiet() bool`

HasQuiet returns a boolean if a field has been set.

### GetPreserveperm

`func (o *RsynctaskCreate0) GetPreserveperm() bool`

GetPreserveperm returns the Preserveperm field if non-nil, zero value otherwise.

### GetPreservepermOk

`func (o *RsynctaskCreate0) GetPreservepermOk() (*bool, bool)`

GetPreservepermOk returns a tuple with the Preserveperm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveperm

`func (o *RsynctaskCreate0) SetPreserveperm(v bool)`

SetPreserveperm sets Preserveperm field to given value.

### HasPreserveperm

`func (o *RsynctaskCreate0) HasPreserveperm() bool`

HasPreserveperm returns a boolean if a field has been set.

### GetPreserveattr

`func (o *RsynctaskCreate0) GetPreserveattr() bool`

GetPreserveattr returns the Preserveattr field if non-nil, zero value otherwise.

### GetPreserveattrOk

`func (o *RsynctaskCreate0) GetPreserveattrOk() (*bool, bool)`

GetPreserveattrOk returns a tuple with the Preserveattr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveattr

`func (o *RsynctaskCreate0) SetPreserveattr(v bool)`

SetPreserveattr sets Preserveattr field to given value.

### HasPreserveattr

`func (o *RsynctaskCreate0) HasPreserveattr() bool`

HasPreserveattr returns a boolean if a field has been set.

### GetDelayupdates

`func (o *RsynctaskCreate0) GetDelayupdates() bool`

GetDelayupdates returns the Delayupdates field if non-nil, zero value otherwise.

### GetDelayupdatesOk

`func (o *RsynctaskCreate0) GetDelayupdatesOk() (*bool, bool)`

GetDelayupdatesOk returns a tuple with the Delayupdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayupdates

`func (o *RsynctaskCreate0) SetDelayupdates(v bool)`

SetDelayupdates sets Delayupdates field to given value.

### HasDelayupdates

`func (o *RsynctaskCreate0) HasDelayupdates() bool`

HasDelayupdates returns a boolean if a field has been set.

### GetExtra

`func (o *RsynctaskCreate0) GetExtra() []string`

GetExtra returns the Extra field if non-nil, zero value otherwise.

### GetExtraOk

`func (o *RsynctaskCreate0) GetExtraOk() (*[]string, bool)`

GetExtraOk returns a tuple with the Extra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtra

`func (o *RsynctaskCreate0) SetExtra(v []string)`

SetExtra sets Extra field to given value.

### HasExtra

`func (o *RsynctaskCreate0) HasExtra() bool`

HasExtra returns a boolean if a field has been set.

### GetEnabled

`func (o *RsynctaskCreate0) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *RsynctaskCreate0) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *RsynctaskCreate0) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *RsynctaskCreate0) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


