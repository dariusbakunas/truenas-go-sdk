# CreateDatasetParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Atime** | Pointer to **string** |  | [optional] 
**Aclmode** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Comments** | Pointer to **string** |  | [optional] 
**Compression** | Pointer to **string** |  | [optional] 
**Casesensitivity** | Pointer to **string** |  | [optional] 
**Copies** | Pointer to **int32** |  | [optional] 
**Deduplication** | Pointer to **string** |  | [optional] 
**Encryption** | Pointer to **bool** |  | [optional] 
**EncryptionOptions** | Pointer to [**CreateDatasetParamsEncryptionOptions**](CreateDatasetParamsEncryptionOptions.md) |  | [optional] 
**Exec** | Pointer to **string** |  | [optional] 
**ForceSize** | Pointer to **bool** |  | [optional] 
**InheritEncryption** | Pointer to **bool** |  | [optional] 
**Quota** | Pointer to **int64** |  | [optional] 
**QuotaCritical** | Pointer to **int64** |  | [optional] 
**QuotaWarning** | Pointer to **int64** |  | [optional] 
**Volsize** | Pointer to **int64** |  | [optional] 
**Volblocksize** | Pointer to **string** |  | [optional] 
**Readonly** | Pointer to **string** |  | [optional] 
**Recordsize** | Pointer to **string** |  | [optional] 
**Refquota** | Pointer to **int64** |  | [optional] 
**RefquotaCritical** | Pointer to **int64** |  | [optional] 
**RefquotaWarning** | Pointer to **int64** |  | [optional] 
**Refreservation** | Pointer to **int64** |  | [optional] 
**Reservation** | Pointer to **int64** |  | [optional] 
**ShareType** | Pointer to **string** |  | [optional] 
**Snapdir** | Pointer to **string** |  | [optional] 
**Sync** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateDatasetParams

`func NewCreateDatasetParams(name string, ) *CreateDatasetParams`

NewCreateDatasetParams instantiates a new CreateDatasetParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDatasetParamsWithDefaults

`func NewCreateDatasetParamsWithDefaults() *CreateDatasetParams`

NewCreateDatasetParamsWithDefaults instantiates a new CreateDatasetParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAtime

`func (o *CreateDatasetParams) GetAtime() string`

GetAtime returns the Atime field if non-nil, zero value otherwise.

### GetAtimeOk

`func (o *CreateDatasetParams) GetAtimeOk() (*string, bool)`

GetAtimeOk returns a tuple with the Atime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtime

`func (o *CreateDatasetParams) SetAtime(v string)`

SetAtime sets Atime field to given value.

### HasAtime

`func (o *CreateDatasetParams) HasAtime() bool`

HasAtime returns a boolean if a field has been set.

### GetAclmode

`func (o *CreateDatasetParams) GetAclmode() string`

GetAclmode returns the Aclmode field if non-nil, zero value otherwise.

### GetAclmodeOk

`func (o *CreateDatasetParams) GetAclmodeOk() (*string, bool)`

GetAclmodeOk returns a tuple with the Aclmode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAclmode

`func (o *CreateDatasetParams) SetAclmode(v string)`

SetAclmode sets Aclmode field to given value.

### HasAclmode

`func (o *CreateDatasetParams) HasAclmode() bool`

HasAclmode returns a boolean if a field has been set.

### GetName

`func (o *CreateDatasetParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateDatasetParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateDatasetParams) SetName(v string)`

SetName sets Name field to given value.


### GetComments

`func (o *CreateDatasetParams) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *CreateDatasetParams) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *CreateDatasetParams) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *CreateDatasetParams) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetCompression

`func (o *CreateDatasetParams) GetCompression() string`

GetCompression returns the Compression field if non-nil, zero value otherwise.

### GetCompressionOk

`func (o *CreateDatasetParams) GetCompressionOk() (*string, bool)`

GetCompressionOk returns a tuple with the Compression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompression

`func (o *CreateDatasetParams) SetCompression(v string)`

SetCompression sets Compression field to given value.

### HasCompression

`func (o *CreateDatasetParams) HasCompression() bool`

HasCompression returns a boolean if a field has been set.

### GetCasesensitivity

`func (o *CreateDatasetParams) GetCasesensitivity() string`

GetCasesensitivity returns the Casesensitivity field if non-nil, zero value otherwise.

### GetCasesensitivityOk

`func (o *CreateDatasetParams) GetCasesensitivityOk() (*string, bool)`

GetCasesensitivityOk returns a tuple with the Casesensitivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCasesensitivity

`func (o *CreateDatasetParams) SetCasesensitivity(v string)`

SetCasesensitivity sets Casesensitivity field to given value.

### HasCasesensitivity

`func (o *CreateDatasetParams) HasCasesensitivity() bool`

HasCasesensitivity returns a boolean if a field has been set.

### GetCopies

`func (o *CreateDatasetParams) GetCopies() int32`

GetCopies returns the Copies field if non-nil, zero value otherwise.

### GetCopiesOk

`func (o *CreateDatasetParams) GetCopiesOk() (*int32, bool)`

GetCopiesOk returns a tuple with the Copies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopies

`func (o *CreateDatasetParams) SetCopies(v int32)`

SetCopies sets Copies field to given value.

### HasCopies

`func (o *CreateDatasetParams) HasCopies() bool`

HasCopies returns a boolean if a field has been set.

### GetDeduplication

`func (o *CreateDatasetParams) GetDeduplication() string`

GetDeduplication returns the Deduplication field if non-nil, zero value otherwise.

### GetDeduplicationOk

`func (o *CreateDatasetParams) GetDeduplicationOk() (*string, bool)`

GetDeduplicationOk returns a tuple with the Deduplication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeduplication

`func (o *CreateDatasetParams) SetDeduplication(v string)`

SetDeduplication sets Deduplication field to given value.

### HasDeduplication

`func (o *CreateDatasetParams) HasDeduplication() bool`

HasDeduplication returns a boolean if a field has been set.

### GetEncryption

`func (o *CreateDatasetParams) GetEncryption() bool`

GetEncryption returns the Encryption field if non-nil, zero value otherwise.

### GetEncryptionOk

`func (o *CreateDatasetParams) GetEncryptionOk() (*bool, bool)`

GetEncryptionOk returns a tuple with the Encryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryption

`func (o *CreateDatasetParams) SetEncryption(v bool)`

SetEncryption sets Encryption field to given value.

### HasEncryption

`func (o *CreateDatasetParams) HasEncryption() bool`

HasEncryption returns a boolean if a field has been set.

### GetEncryptionOptions

`func (o *CreateDatasetParams) GetEncryptionOptions() CreateDatasetParamsEncryptionOptions`

GetEncryptionOptions returns the EncryptionOptions field if non-nil, zero value otherwise.

### GetEncryptionOptionsOk

`func (o *CreateDatasetParams) GetEncryptionOptionsOk() (*CreateDatasetParamsEncryptionOptions, bool)`

GetEncryptionOptionsOk returns a tuple with the EncryptionOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionOptions

`func (o *CreateDatasetParams) SetEncryptionOptions(v CreateDatasetParamsEncryptionOptions)`

SetEncryptionOptions sets EncryptionOptions field to given value.

### HasEncryptionOptions

`func (o *CreateDatasetParams) HasEncryptionOptions() bool`

HasEncryptionOptions returns a boolean if a field has been set.

### GetExec

`func (o *CreateDatasetParams) GetExec() string`

GetExec returns the Exec field if non-nil, zero value otherwise.

### GetExecOk

`func (o *CreateDatasetParams) GetExecOk() (*string, bool)`

GetExecOk returns a tuple with the Exec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExec

`func (o *CreateDatasetParams) SetExec(v string)`

SetExec sets Exec field to given value.

### HasExec

`func (o *CreateDatasetParams) HasExec() bool`

HasExec returns a boolean if a field has been set.

### GetForceSize

`func (o *CreateDatasetParams) GetForceSize() bool`

GetForceSize returns the ForceSize field if non-nil, zero value otherwise.

### GetForceSizeOk

`func (o *CreateDatasetParams) GetForceSizeOk() (*bool, bool)`

GetForceSizeOk returns a tuple with the ForceSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceSize

`func (o *CreateDatasetParams) SetForceSize(v bool)`

SetForceSize sets ForceSize field to given value.

### HasForceSize

`func (o *CreateDatasetParams) HasForceSize() bool`

HasForceSize returns a boolean if a field has been set.

### GetInheritEncryption

`func (o *CreateDatasetParams) GetInheritEncryption() bool`

GetInheritEncryption returns the InheritEncryption field if non-nil, zero value otherwise.

### GetInheritEncryptionOk

`func (o *CreateDatasetParams) GetInheritEncryptionOk() (*bool, bool)`

GetInheritEncryptionOk returns a tuple with the InheritEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInheritEncryption

`func (o *CreateDatasetParams) SetInheritEncryption(v bool)`

SetInheritEncryption sets InheritEncryption field to given value.

### HasInheritEncryption

`func (o *CreateDatasetParams) HasInheritEncryption() bool`

HasInheritEncryption returns a boolean if a field has been set.

### GetQuota

`func (o *CreateDatasetParams) GetQuota() int64`

GetQuota returns the Quota field if non-nil, zero value otherwise.

### GetQuotaOk

`func (o *CreateDatasetParams) GetQuotaOk() (*int64, bool)`

GetQuotaOk returns a tuple with the Quota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuota

`func (o *CreateDatasetParams) SetQuota(v int64)`

SetQuota sets Quota field to given value.

### HasQuota

`func (o *CreateDatasetParams) HasQuota() bool`

HasQuota returns a boolean if a field has been set.

### GetQuotaCritical

`func (o *CreateDatasetParams) GetQuotaCritical() int64`

GetQuotaCritical returns the QuotaCritical field if non-nil, zero value otherwise.

### GetQuotaCriticalOk

`func (o *CreateDatasetParams) GetQuotaCriticalOk() (*int64, bool)`

GetQuotaCriticalOk returns a tuple with the QuotaCritical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaCritical

`func (o *CreateDatasetParams) SetQuotaCritical(v int64)`

SetQuotaCritical sets QuotaCritical field to given value.

### HasQuotaCritical

`func (o *CreateDatasetParams) HasQuotaCritical() bool`

HasQuotaCritical returns a boolean if a field has been set.

### GetQuotaWarning

`func (o *CreateDatasetParams) GetQuotaWarning() int64`

GetQuotaWarning returns the QuotaWarning field if non-nil, zero value otherwise.

### GetQuotaWarningOk

`func (o *CreateDatasetParams) GetQuotaWarningOk() (*int64, bool)`

GetQuotaWarningOk returns a tuple with the QuotaWarning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaWarning

`func (o *CreateDatasetParams) SetQuotaWarning(v int64)`

SetQuotaWarning sets QuotaWarning field to given value.

### HasQuotaWarning

`func (o *CreateDatasetParams) HasQuotaWarning() bool`

HasQuotaWarning returns a boolean if a field has been set.

### GetVolsize

`func (o *CreateDatasetParams) GetVolsize() int64`

GetVolsize returns the Volsize field if non-nil, zero value otherwise.

### GetVolsizeOk

`func (o *CreateDatasetParams) GetVolsizeOk() (*int64, bool)`

GetVolsizeOk returns a tuple with the Volsize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolsize

`func (o *CreateDatasetParams) SetVolsize(v int64)`

SetVolsize sets Volsize field to given value.

### HasVolsize

`func (o *CreateDatasetParams) HasVolsize() bool`

HasVolsize returns a boolean if a field has been set.

### GetVolblocksize

`func (o *CreateDatasetParams) GetVolblocksize() string`

GetVolblocksize returns the Volblocksize field if non-nil, zero value otherwise.

### GetVolblocksizeOk

`func (o *CreateDatasetParams) GetVolblocksizeOk() (*string, bool)`

GetVolblocksizeOk returns a tuple with the Volblocksize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolblocksize

`func (o *CreateDatasetParams) SetVolblocksize(v string)`

SetVolblocksize sets Volblocksize field to given value.

### HasVolblocksize

`func (o *CreateDatasetParams) HasVolblocksize() bool`

HasVolblocksize returns a boolean if a field has been set.

### GetReadonly

`func (o *CreateDatasetParams) GetReadonly() string`

GetReadonly returns the Readonly field if non-nil, zero value otherwise.

### GetReadonlyOk

`func (o *CreateDatasetParams) GetReadonlyOk() (*string, bool)`

GetReadonlyOk returns a tuple with the Readonly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadonly

`func (o *CreateDatasetParams) SetReadonly(v string)`

SetReadonly sets Readonly field to given value.

### HasReadonly

`func (o *CreateDatasetParams) HasReadonly() bool`

HasReadonly returns a boolean if a field has been set.

### GetRecordsize

`func (o *CreateDatasetParams) GetRecordsize() string`

GetRecordsize returns the Recordsize field if non-nil, zero value otherwise.

### GetRecordsizeOk

`func (o *CreateDatasetParams) GetRecordsizeOk() (*string, bool)`

GetRecordsizeOk returns a tuple with the Recordsize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordsize

`func (o *CreateDatasetParams) SetRecordsize(v string)`

SetRecordsize sets Recordsize field to given value.

### HasRecordsize

`func (o *CreateDatasetParams) HasRecordsize() bool`

HasRecordsize returns a boolean if a field has been set.

### GetRefquota

`func (o *CreateDatasetParams) GetRefquota() int64`

GetRefquota returns the Refquota field if non-nil, zero value otherwise.

### GetRefquotaOk

`func (o *CreateDatasetParams) GetRefquotaOk() (*int64, bool)`

GetRefquotaOk returns a tuple with the Refquota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefquota

`func (o *CreateDatasetParams) SetRefquota(v int64)`

SetRefquota sets Refquota field to given value.

### HasRefquota

`func (o *CreateDatasetParams) HasRefquota() bool`

HasRefquota returns a boolean if a field has been set.

### GetRefquotaCritical

`func (o *CreateDatasetParams) GetRefquotaCritical() int64`

GetRefquotaCritical returns the RefquotaCritical field if non-nil, zero value otherwise.

### GetRefquotaCriticalOk

`func (o *CreateDatasetParams) GetRefquotaCriticalOk() (*int64, bool)`

GetRefquotaCriticalOk returns a tuple with the RefquotaCritical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefquotaCritical

`func (o *CreateDatasetParams) SetRefquotaCritical(v int64)`

SetRefquotaCritical sets RefquotaCritical field to given value.

### HasRefquotaCritical

`func (o *CreateDatasetParams) HasRefquotaCritical() bool`

HasRefquotaCritical returns a boolean if a field has been set.

### GetRefquotaWarning

`func (o *CreateDatasetParams) GetRefquotaWarning() int64`

GetRefquotaWarning returns the RefquotaWarning field if non-nil, zero value otherwise.

### GetRefquotaWarningOk

`func (o *CreateDatasetParams) GetRefquotaWarningOk() (*int64, bool)`

GetRefquotaWarningOk returns a tuple with the RefquotaWarning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefquotaWarning

`func (o *CreateDatasetParams) SetRefquotaWarning(v int64)`

SetRefquotaWarning sets RefquotaWarning field to given value.

### HasRefquotaWarning

`func (o *CreateDatasetParams) HasRefquotaWarning() bool`

HasRefquotaWarning returns a boolean if a field has been set.

### GetRefreservation

`func (o *CreateDatasetParams) GetRefreservation() int64`

GetRefreservation returns the Refreservation field if non-nil, zero value otherwise.

### GetRefreservationOk

`func (o *CreateDatasetParams) GetRefreservationOk() (*int64, bool)`

GetRefreservationOk returns a tuple with the Refreservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreservation

`func (o *CreateDatasetParams) SetRefreservation(v int64)`

SetRefreservation sets Refreservation field to given value.

### HasRefreservation

`func (o *CreateDatasetParams) HasRefreservation() bool`

HasRefreservation returns a boolean if a field has been set.

### GetReservation

`func (o *CreateDatasetParams) GetReservation() int64`

GetReservation returns the Reservation field if non-nil, zero value otherwise.

### GetReservationOk

`func (o *CreateDatasetParams) GetReservationOk() (*int64, bool)`

GetReservationOk returns a tuple with the Reservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservation

`func (o *CreateDatasetParams) SetReservation(v int64)`

SetReservation sets Reservation field to given value.

### HasReservation

`func (o *CreateDatasetParams) HasReservation() bool`

HasReservation returns a boolean if a field has been set.

### GetShareType

`func (o *CreateDatasetParams) GetShareType() string`

GetShareType returns the ShareType field if non-nil, zero value otherwise.

### GetShareTypeOk

`func (o *CreateDatasetParams) GetShareTypeOk() (*string, bool)`

GetShareTypeOk returns a tuple with the ShareType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareType

`func (o *CreateDatasetParams) SetShareType(v string)`

SetShareType sets ShareType field to given value.

### HasShareType

`func (o *CreateDatasetParams) HasShareType() bool`

HasShareType returns a boolean if a field has been set.

### GetSnapdir

`func (o *CreateDatasetParams) GetSnapdir() string`

GetSnapdir returns the Snapdir field if non-nil, zero value otherwise.

### GetSnapdirOk

`func (o *CreateDatasetParams) GetSnapdirOk() (*string, bool)`

GetSnapdirOk returns a tuple with the Snapdir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapdir

`func (o *CreateDatasetParams) SetSnapdir(v string)`

SetSnapdir sets Snapdir field to given value.

### HasSnapdir

`func (o *CreateDatasetParams) HasSnapdir() bool`

HasSnapdir returns a boolean if a field has been set.

### GetSync

`func (o *CreateDatasetParams) GetSync() string`

GetSync returns the Sync field if non-nil, zero value otherwise.

### GetSyncOk

`func (o *CreateDatasetParams) GetSyncOk() (*string, bool)`

GetSyncOk returns a tuple with the Sync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSync

`func (o *CreateDatasetParams) SetSync(v string)`

SetSync sets Sync field to given value.

### HasSync

`func (o *CreateDatasetParams) HasSync() bool`

HasSync returns a boolean if a field has been set.

### GetType

`func (o *CreateDatasetParams) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateDatasetParams) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateDatasetParams) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CreateDatasetParams) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


