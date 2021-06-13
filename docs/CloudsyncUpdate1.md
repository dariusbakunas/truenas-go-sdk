# CloudsyncUpdate1

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

### NewCloudsyncUpdate1

`func NewCloudsyncUpdate1() *CloudsyncUpdate1`

NewCloudsyncUpdate1 instantiates a new CloudsyncUpdate1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudsyncUpdate1WithDefaults

`func NewCloudsyncUpdate1WithDefaults() *CloudsyncUpdate1`

NewCloudsyncUpdate1WithDefaults instantiates a new CloudsyncUpdate1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CloudsyncUpdate1) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CloudsyncUpdate1) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CloudsyncUpdate1) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CloudsyncUpdate1) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDirection

`func (o *CloudsyncUpdate1) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *CloudsyncUpdate1) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *CloudsyncUpdate1) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *CloudsyncUpdate1) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetTransferMode

`func (o *CloudsyncUpdate1) GetTransferMode() string`

GetTransferMode returns the TransferMode field if non-nil, zero value otherwise.

### GetTransferModeOk

`func (o *CloudsyncUpdate1) GetTransferModeOk() (*string, bool)`

GetTransferModeOk returns a tuple with the TransferMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferMode

`func (o *CloudsyncUpdate1) SetTransferMode(v string)`

SetTransferMode sets TransferMode field to given value.

### HasTransferMode

`func (o *CloudsyncUpdate1) HasTransferMode() bool`

HasTransferMode returns a boolean if a field has been set.

### GetPath

`func (o *CloudsyncUpdate1) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *CloudsyncUpdate1) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *CloudsyncUpdate1) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *CloudsyncUpdate1) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetCredentials

`func (o *CloudsyncUpdate1) GetCredentials() int32`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *CloudsyncUpdate1) GetCredentialsOk() (*int32, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *CloudsyncUpdate1) SetCredentials(v int32)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *CloudsyncUpdate1) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetEncryption

`func (o *CloudsyncUpdate1) GetEncryption() bool`

GetEncryption returns the Encryption field if non-nil, zero value otherwise.

### GetEncryptionOk

`func (o *CloudsyncUpdate1) GetEncryptionOk() (*bool, bool)`

GetEncryptionOk returns a tuple with the Encryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryption

`func (o *CloudsyncUpdate1) SetEncryption(v bool)`

SetEncryption sets Encryption field to given value.

### HasEncryption

`func (o *CloudsyncUpdate1) HasEncryption() bool`

HasEncryption returns a boolean if a field has been set.

### GetFilenameEncryption

`func (o *CloudsyncUpdate1) GetFilenameEncryption() bool`

GetFilenameEncryption returns the FilenameEncryption field if non-nil, zero value otherwise.

### GetFilenameEncryptionOk

`func (o *CloudsyncUpdate1) GetFilenameEncryptionOk() (*bool, bool)`

GetFilenameEncryptionOk returns a tuple with the FilenameEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilenameEncryption

`func (o *CloudsyncUpdate1) SetFilenameEncryption(v bool)`

SetFilenameEncryption sets FilenameEncryption field to given value.

### HasFilenameEncryption

`func (o *CloudsyncUpdate1) HasFilenameEncryption() bool`

HasFilenameEncryption returns a boolean if a field has been set.

### GetEncryptionPassword

`func (o *CloudsyncUpdate1) GetEncryptionPassword() string`

GetEncryptionPassword returns the EncryptionPassword field if non-nil, zero value otherwise.

### GetEncryptionPasswordOk

`func (o *CloudsyncUpdate1) GetEncryptionPasswordOk() (*string, bool)`

GetEncryptionPasswordOk returns a tuple with the EncryptionPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionPassword

`func (o *CloudsyncUpdate1) SetEncryptionPassword(v string)`

SetEncryptionPassword sets EncryptionPassword field to given value.

### HasEncryptionPassword

`func (o *CloudsyncUpdate1) HasEncryptionPassword() bool`

HasEncryptionPassword returns a boolean if a field has been set.

### GetEncryptionSalt

`func (o *CloudsyncUpdate1) GetEncryptionSalt() string`

GetEncryptionSalt returns the EncryptionSalt field if non-nil, zero value otherwise.

### GetEncryptionSaltOk

`func (o *CloudsyncUpdate1) GetEncryptionSaltOk() (*string, bool)`

GetEncryptionSaltOk returns a tuple with the EncryptionSalt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionSalt

`func (o *CloudsyncUpdate1) SetEncryptionSalt(v string)`

SetEncryptionSalt sets EncryptionSalt field to given value.

### HasEncryptionSalt

`func (o *CloudsyncUpdate1) HasEncryptionSalt() bool`

HasEncryptionSalt returns a boolean if a field has been set.

### GetSchedule

`func (o *CloudsyncUpdate1) GetSchedule() CloudsyncCreate0Schedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *CloudsyncUpdate1) GetScheduleOk() (*CloudsyncCreate0Schedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *CloudsyncUpdate1) SetSchedule(v CloudsyncCreate0Schedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *CloudsyncUpdate1) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetFollowSymlinks

`func (o *CloudsyncUpdate1) GetFollowSymlinks() bool`

GetFollowSymlinks returns the FollowSymlinks field if non-nil, zero value otherwise.

### GetFollowSymlinksOk

`func (o *CloudsyncUpdate1) GetFollowSymlinksOk() (*bool, bool)`

GetFollowSymlinksOk returns a tuple with the FollowSymlinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowSymlinks

`func (o *CloudsyncUpdate1) SetFollowSymlinks(v bool)`

SetFollowSymlinks sets FollowSymlinks field to given value.

### HasFollowSymlinks

`func (o *CloudsyncUpdate1) HasFollowSymlinks() bool`

HasFollowSymlinks returns a boolean if a field has been set.

### GetTransfers

`func (o *CloudsyncUpdate1) GetTransfers() int32`

GetTransfers returns the Transfers field if non-nil, zero value otherwise.

### GetTransfersOk

`func (o *CloudsyncUpdate1) GetTransfersOk() (*int32, bool)`

GetTransfersOk returns a tuple with the Transfers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransfers

`func (o *CloudsyncUpdate1) SetTransfers(v int32)`

SetTransfers sets Transfers field to given value.

### HasTransfers

`func (o *CloudsyncUpdate1) HasTransfers() bool`

HasTransfers returns a boolean if a field has been set.

### SetTransfersNil

`func (o *CloudsyncUpdate1) SetTransfersNil(b bool)`

 SetTransfersNil sets the value for Transfers to be an explicit nil

### UnsetTransfers
`func (o *CloudsyncUpdate1) UnsetTransfers()`

UnsetTransfers ensures that no value is present for Transfers, not even an explicit nil
### GetBwlimit

`func (o *CloudsyncUpdate1) GetBwlimit() []map[string]interface{}`

GetBwlimit returns the Bwlimit field if non-nil, zero value otherwise.

### GetBwlimitOk

`func (o *CloudsyncUpdate1) GetBwlimitOk() (*[]map[string]interface{}, bool)`

GetBwlimitOk returns a tuple with the Bwlimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBwlimit

`func (o *CloudsyncUpdate1) SetBwlimit(v []map[string]interface{})`

SetBwlimit sets Bwlimit field to given value.

### HasBwlimit

`func (o *CloudsyncUpdate1) HasBwlimit() bool`

HasBwlimit returns a boolean if a field has been set.

### GetExclude

`func (o *CloudsyncUpdate1) GetExclude() []string`

GetExclude returns the Exclude field if non-nil, zero value otherwise.

### GetExcludeOk

`func (o *CloudsyncUpdate1) GetExcludeOk() (*[]string, bool)`

GetExcludeOk returns a tuple with the Exclude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclude

`func (o *CloudsyncUpdate1) SetExclude(v []string)`

SetExclude sets Exclude field to given value.

### HasExclude

`func (o *CloudsyncUpdate1) HasExclude() bool`

HasExclude returns a boolean if a field has been set.

### GetAttributes

`func (o *CloudsyncUpdate1) GetAttributes() map[string]map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CloudsyncUpdate1) GetAttributesOk() (*map[string]map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CloudsyncUpdate1) SetAttributes(v map[string]map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *CloudsyncUpdate1) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetSnapshot

`func (o *CloudsyncUpdate1) GetSnapshot() bool`

GetSnapshot returns the Snapshot field if non-nil, zero value otherwise.

### GetSnapshotOk

`func (o *CloudsyncUpdate1) GetSnapshotOk() (*bool, bool)`

GetSnapshotOk returns a tuple with the Snapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshot

`func (o *CloudsyncUpdate1) SetSnapshot(v bool)`

SetSnapshot sets Snapshot field to given value.

### HasSnapshot

`func (o *CloudsyncUpdate1) HasSnapshot() bool`

HasSnapshot returns a boolean if a field has been set.

### GetPreScript

`func (o *CloudsyncUpdate1) GetPreScript() string`

GetPreScript returns the PreScript field if non-nil, zero value otherwise.

### GetPreScriptOk

`func (o *CloudsyncUpdate1) GetPreScriptOk() (*string, bool)`

GetPreScriptOk returns a tuple with the PreScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreScript

`func (o *CloudsyncUpdate1) SetPreScript(v string)`

SetPreScript sets PreScript field to given value.

### HasPreScript

`func (o *CloudsyncUpdate1) HasPreScript() bool`

HasPreScript returns a boolean if a field has been set.

### GetPostScript

`func (o *CloudsyncUpdate1) GetPostScript() string`

GetPostScript returns the PostScript field if non-nil, zero value otherwise.

### GetPostScriptOk

`func (o *CloudsyncUpdate1) GetPostScriptOk() (*string, bool)`

GetPostScriptOk returns a tuple with the PostScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostScript

`func (o *CloudsyncUpdate1) SetPostScript(v string)`

SetPostScript sets PostScript field to given value.

### HasPostScript

`func (o *CloudsyncUpdate1) HasPostScript() bool`

HasPostScript returns a boolean if a field has been set.

### GetArgs

`func (o *CloudsyncUpdate1) GetArgs() string`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *CloudsyncUpdate1) GetArgsOk() (*string, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *CloudsyncUpdate1) SetArgs(v string)`

SetArgs sets Args field to given value.

### HasArgs

`func (o *CloudsyncUpdate1) HasArgs() bool`

HasArgs returns a boolean if a field has been set.

### GetEnabled

`func (o *CloudsyncUpdate1) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CloudsyncUpdate1) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CloudsyncUpdate1) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CloudsyncUpdate1) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


