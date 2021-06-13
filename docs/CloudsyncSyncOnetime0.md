# CloudsyncSyncOnetime0

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Direction** | Pointer to **string** |  | [optional] 
**TransferMode** | Pointer to **string** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**Credentials** | Pointer to **int32** |  | [optional] 
**Encryption** | Pointer to **bool** |  | [optional] 
**FilenameEncryption** | Pointer to **bool** |  | [optional] 
**EncryptionPassword** | Pointer to **string** |  | [optional] 
**EncryptionSalt** | Pointer to **string** |  | [optional] 
**Schedule** | Pointer to [**CloudsyncCreate0Schedule**](CloudsyncCreate0Schedule.md) |  | [optional] 
**FollowSymlinks** | Pointer to **bool** |  | [optional] 
**Transfers** | Pointer to **NullableInt32** |  | [optional] 
**Bwlimit** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Exclude** | Pointer to **[]string** |  | [optional] 
**Attributes** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Snapshot** | Pointer to **bool** |  | [optional] 
**PreScript** | Pointer to **string** |  | [optional] 
**PostScript** | Pointer to **string** |  | [optional] 
**Args** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewCloudsyncSyncOnetime0

`func NewCloudsyncSyncOnetime0() *CloudsyncSyncOnetime0`

NewCloudsyncSyncOnetime0 instantiates a new CloudsyncSyncOnetime0 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudsyncSyncOnetime0WithDefaults

`func NewCloudsyncSyncOnetime0WithDefaults() *CloudsyncSyncOnetime0`

NewCloudsyncSyncOnetime0WithDefaults instantiates a new CloudsyncSyncOnetime0 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CloudsyncSyncOnetime0) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CloudsyncSyncOnetime0) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CloudsyncSyncOnetime0) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CloudsyncSyncOnetime0) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDirection

`func (o *CloudsyncSyncOnetime0) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *CloudsyncSyncOnetime0) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *CloudsyncSyncOnetime0) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *CloudsyncSyncOnetime0) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetTransferMode

`func (o *CloudsyncSyncOnetime0) GetTransferMode() string`

GetTransferMode returns the TransferMode field if non-nil, zero value otherwise.

### GetTransferModeOk

`func (o *CloudsyncSyncOnetime0) GetTransferModeOk() (*string, bool)`

GetTransferModeOk returns a tuple with the TransferMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferMode

`func (o *CloudsyncSyncOnetime0) SetTransferMode(v string)`

SetTransferMode sets TransferMode field to given value.

### HasTransferMode

`func (o *CloudsyncSyncOnetime0) HasTransferMode() bool`

HasTransferMode returns a boolean if a field has been set.

### GetPath

`func (o *CloudsyncSyncOnetime0) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *CloudsyncSyncOnetime0) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *CloudsyncSyncOnetime0) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *CloudsyncSyncOnetime0) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetCredentials

`func (o *CloudsyncSyncOnetime0) GetCredentials() int32`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *CloudsyncSyncOnetime0) GetCredentialsOk() (*int32, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *CloudsyncSyncOnetime0) SetCredentials(v int32)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *CloudsyncSyncOnetime0) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetEncryption

`func (o *CloudsyncSyncOnetime0) GetEncryption() bool`

GetEncryption returns the Encryption field if non-nil, zero value otherwise.

### GetEncryptionOk

`func (o *CloudsyncSyncOnetime0) GetEncryptionOk() (*bool, bool)`

GetEncryptionOk returns a tuple with the Encryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryption

`func (o *CloudsyncSyncOnetime0) SetEncryption(v bool)`

SetEncryption sets Encryption field to given value.

### HasEncryption

`func (o *CloudsyncSyncOnetime0) HasEncryption() bool`

HasEncryption returns a boolean if a field has been set.

### GetFilenameEncryption

`func (o *CloudsyncSyncOnetime0) GetFilenameEncryption() bool`

GetFilenameEncryption returns the FilenameEncryption field if non-nil, zero value otherwise.

### GetFilenameEncryptionOk

`func (o *CloudsyncSyncOnetime0) GetFilenameEncryptionOk() (*bool, bool)`

GetFilenameEncryptionOk returns a tuple with the FilenameEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilenameEncryption

`func (o *CloudsyncSyncOnetime0) SetFilenameEncryption(v bool)`

SetFilenameEncryption sets FilenameEncryption field to given value.

### HasFilenameEncryption

`func (o *CloudsyncSyncOnetime0) HasFilenameEncryption() bool`

HasFilenameEncryption returns a boolean if a field has been set.

### GetEncryptionPassword

`func (o *CloudsyncSyncOnetime0) GetEncryptionPassword() string`

GetEncryptionPassword returns the EncryptionPassword field if non-nil, zero value otherwise.

### GetEncryptionPasswordOk

`func (o *CloudsyncSyncOnetime0) GetEncryptionPasswordOk() (*string, bool)`

GetEncryptionPasswordOk returns a tuple with the EncryptionPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionPassword

`func (o *CloudsyncSyncOnetime0) SetEncryptionPassword(v string)`

SetEncryptionPassword sets EncryptionPassword field to given value.

### HasEncryptionPassword

`func (o *CloudsyncSyncOnetime0) HasEncryptionPassword() bool`

HasEncryptionPassword returns a boolean if a field has been set.

### GetEncryptionSalt

`func (o *CloudsyncSyncOnetime0) GetEncryptionSalt() string`

GetEncryptionSalt returns the EncryptionSalt field if non-nil, zero value otherwise.

### GetEncryptionSaltOk

`func (o *CloudsyncSyncOnetime0) GetEncryptionSaltOk() (*string, bool)`

GetEncryptionSaltOk returns a tuple with the EncryptionSalt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionSalt

`func (o *CloudsyncSyncOnetime0) SetEncryptionSalt(v string)`

SetEncryptionSalt sets EncryptionSalt field to given value.

### HasEncryptionSalt

`func (o *CloudsyncSyncOnetime0) HasEncryptionSalt() bool`

HasEncryptionSalt returns a boolean if a field has been set.

### GetSchedule

`func (o *CloudsyncSyncOnetime0) GetSchedule() CloudsyncCreate0Schedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *CloudsyncSyncOnetime0) GetScheduleOk() (*CloudsyncCreate0Schedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *CloudsyncSyncOnetime0) SetSchedule(v CloudsyncCreate0Schedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *CloudsyncSyncOnetime0) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetFollowSymlinks

`func (o *CloudsyncSyncOnetime0) GetFollowSymlinks() bool`

GetFollowSymlinks returns the FollowSymlinks field if non-nil, zero value otherwise.

### GetFollowSymlinksOk

`func (o *CloudsyncSyncOnetime0) GetFollowSymlinksOk() (*bool, bool)`

GetFollowSymlinksOk returns a tuple with the FollowSymlinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowSymlinks

`func (o *CloudsyncSyncOnetime0) SetFollowSymlinks(v bool)`

SetFollowSymlinks sets FollowSymlinks field to given value.

### HasFollowSymlinks

`func (o *CloudsyncSyncOnetime0) HasFollowSymlinks() bool`

HasFollowSymlinks returns a boolean if a field has been set.

### GetTransfers

`func (o *CloudsyncSyncOnetime0) GetTransfers() int32`

GetTransfers returns the Transfers field if non-nil, zero value otherwise.

### GetTransfersOk

`func (o *CloudsyncSyncOnetime0) GetTransfersOk() (*int32, bool)`

GetTransfersOk returns a tuple with the Transfers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransfers

`func (o *CloudsyncSyncOnetime0) SetTransfers(v int32)`

SetTransfers sets Transfers field to given value.

### HasTransfers

`func (o *CloudsyncSyncOnetime0) HasTransfers() bool`

HasTransfers returns a boolean if a field has been set.

### SetTransfersNil

`func (o *CloudsyncSyncOnetime0) SetTransfersNil(b bool)`

 SetTransfersNil sets the value for Transfers to be an explicit nil

### UnsetTransfers
`func (o *CloudsyncSyncOnetime0) UnsetTransfers()`

UnsetTransfers ensures that no value is present for Transfers, not even an explicit nil
### GetBwlimit

`func (o *CloudsyncSyncOnetime0) GetBwlimit() []map[string]interface{}`

GetBwlimit returns the Bwlimit field if non-nil, zero value otherwise.

### GetBwlimitOk

`func (o *CloudsyncSyncOnetime0) GetBwlimitOk() (*[]map[string]interface{}, bool)`

GetBwlimitOk returns a tuple with the Bwlimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBwlimit

`func (o *CloudsyncSyncOnetime0) SetBwlimit(v []map[string]interface{})`

SetBwlimit sets Bwlimit field to given value.

### HasBwlimit

`func (o *CloudsyncSyncOnetime0) HasBwlimit() bool`

HasBwlimit returns a boolean if a field has been set.

### GetExclude

`func (o *CloudsyncSyncOnetime0) GetExclude() []string`

GetExclude returns the Exclude field if non-nil, zero value otherwise.

### GetExcludeOk

`func (o *CloudsyncSyncOnetime0) GetExcludeOk() (*[]string, bool)`

GetExcludeOk returns a tuple with the Exclude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclude

`func (o *CloudsyncSyncOnetime0) SetExclude(v []string)`

SetExclude sets Exclude field to given value.

### HasExclude

`func (o *CloudsyncSyncOnetime0) HasExclude() bool`

HasExclude returns a boolean if a field has been set.

### GetAttributes

`func (o *CloudsyncSyncOnetime0) GetAttributes() map[string]map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CloudsyncSyncOnetime0) GetAttributesOk() (*map[string]map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CloudsyncSyncOnetime0) SetAttributes(v map[string]map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *CloudsyncSyncOnetime0) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetSnapshot

`func (o *CloudsyncSyncOnetime0) GetSnapshot() bool`

GetSnapshot returns the Snapshot field if non-nil, zero value otherwise.

### GetSnapshotOk

`func (o *CloudsyncSyncOnetime0) GetSnapshotOk() (*bool, bool)`

GetSnapshotOk returns a tuple with the Snapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshot

`func (o *CloudsyncSyncOnetime0) SetSnapshot(v bool)`

SetSnapshot sets Snapshot field to given value.

### HasSnapshot

`func (o *CloudsyncSyncOnetime0) HasSnapshot() bool`

HasSnapshot returns a boolean if a field has been set.

### GetPreScript

`func (o *CloudsyncSyncOnetime0) GetPreScript() string`

GetPreScript returns the PreScript field if non-nil, zero value otherwise.

### GetPreScriptOk

`func (o *CloudsyncSyncOnetime0) GetPreScriptOk() (*string, bool)`

GetPreScriptOk returns a tuple with the PreScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreScript

`func (o *CloudsyncSyncOnetime0) SetPreScript(v string)`

SetPreScript sets PreScript field to given value.

### HasPreScript

`func (o *CloudsyncSyncOnetime0) HasPreScript() bool`

HasPreScript returns a boolean if a field has been set.

### GetPostScript

`func (o *CloudsyncSyncOnetime0) GetPostScript() string`

GetPostScript returns the PostScript field if non-nil, zero value otherwise.

### GetPostScriptOk

`func (o *CloudsyncSyncOnetime0) GetPostScriptOk() (*string, bool)`

GetPostScriptOk returns a tuple with the PostScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostScript

`func (o *CloudsyncSyncOnetime0) SetPostScript(v string)`

SetPostScript sets PostScript field to given value.

### HasPostScript

`func (o *CloudsyncSyncOnetime0) HasPostScript() bool`

HasPostScript returns a boolean if a field has been set.

### GetArgs

`func (o *CloudsyncSyncOnetime0) GetArgs() string`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *CloudsyncSyncOnetime0) GetArgsOk() (*string, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *CloudsyncSyncOnetime0) SetArgs(v string)`

SetArgs sets Args field to given value.

### HasArgs

`func (o *CloudsyncSyncOnetime0) HasArgs() bool`

HasArgs returns a boolean if a field has been set.

### GetEnabled

`func (o *CloudsyncSyncOnetime0) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CloudsyncSyncOnetime0) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CloudsyncSyncOnetime0) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CloudsyncSyncOnetime0) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


