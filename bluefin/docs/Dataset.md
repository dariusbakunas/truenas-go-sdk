# Dataset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Pool** | **string** |  | 
**Type** | **string** |  | 
**Mountpoint** | Pointer to **string** |  | [optional] 
**Encrypted** | Pointer to **bool** |  | [optional] 
**EncryptionRoot** | Pointer to **string** |  | [optional] 
**KeyLoaded** | Pointer to **bool** |  | [optional] 
**Locked** | Pointer to **bool** |  | [optional] 
**EncryptionAlgorithm** | Pointer to [**CompositeValue**](CompositeValue.md) |  | [optional] 
**Aclmode** | Pointer to [**CompositeValue**](CompositeValue.md) |  | [optional] 
**Acltype** | Pointer to [**CompositeValue**](CompositeValue.md) |  | [optional] 
**Atime** | Pointer to [**CompositeValue**](CompositeValue.md) |  | [optional] 
**Available** | Pointer to [**CompositeValue**](CompositeValue.md) |  | [optional] 
**Casesensitivity** | Pointer to [**CompositeValue**](CompositeValue.md) |  | [optional] 
**Comments** | Pointer to [**CompositeValue**](CompositeValue.md) |  | [optional] 
**Compression** | Pointer to [**CompositeValue**](CompositeValue.md) |  | [optional] 
**Deduplication** | Pointer to [**CompositeValue**](CompositeValue.md) |  | [optional] 
**Exec** | Pointer to [**CompositeValue**](CompositeValue.md) |  | [optional] 
**KeyFormat** | Pointer to [**CompositeValue**](CompositeValue.md) |  | [optional] 
**Managedby** | Pointer to [**CompositeValue**](CompositeValue.md) |  | [optional] 
**Copies** | Pointer to [**CompositeValue**](CompositeValue.md) |  | [optional] 
**Quota** | Pointer to [**CompositeValue**](CompositeValue.md) |  | [optional] 
**QuotaCritical** | Pointer to [**CompositeValue**](CompositeValue.md) |  | [optional] 
**QuotaWarning** | Pointer to [**CompositeValue**](CompositeValue.md) |  | [optional] 
**Reservation** | Pointer to [**CompositeValue**](CompositeValue.md) |  | [optional] 
**Refreservation** | Pointer to [**CompositeValue**](CompositeValue.md) |  | [optional] 
**Refquota** | Pointer to [**CompositeValue**](CompositeValue.md) |  | [optional] 
**RefquotaCritical** | Pointer to [**CompositeValue**](CompositeValue.md) |  | [optional] 
**RefquotaWarning** | Pointer to [**CompositeValue**](CompositeValue.md) |  | [optional] 
**Readonly** | Pointer to [**CompositeValue**](CompositeValue.md) |  | [optional] 
**Recordsize** | Pointer to [**CompositeValue**](CompositeValue.md) |  | [optional] 
**Sync** | Pointer to [**CompositeValue**](CompositeValue.md) |  | [optional] 
**Snapdir** | Pointer to [**CompositeValue**](CompositeValue.md) |  | [optional] 
**Used** | Pointer to [**CompositeValue**](CompositeValue.md) |  | [optional] 
**Pbkdf2iters** | Pointer to [**CompositeValue**](CompositeValue.md) |  | [optional] 
**Origin** | Pointer to [**CompositeValue**](CompositeValue.md) |  | [optional] 
**Xattr** | Pointer to [**CompositeValue**](CompositeValue.md) |  | [optional] 
**Volsize** | Pointer to [**CompositeValue**](CompositeValue.md) |  | [optional] 
**Volblocksize** | Pointer to [**CompositeValue**](CompositeValue.md) |  | [optional] 

## Methods

### NewDataset

`func NewDataset(id string, name string, pool string, type_ string, ) *Dataset`

NewDataset instantiates a new Dataset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatasetWithDefaults

`func NewDatasetWithDefaults() *Dataset`

NewDatasetWithDefaults instantiates a new Dataset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Dataset) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Dataset) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Dataset) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Dataset) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Dataset) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Dataset) SetName(v string)`

SetName sets Name field to given value.


### GetPool

`func (o *Dataset) GetPool() string`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *Dataset) GetPoolOk() (*string, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *Dataset) SetPool(v string)`

SetPool sets Pool field to given value.


### GetType

`func (o *Dataset) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Dataset) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Dataset) SetType(v string)`

SetType sets Type field to given value.


### GetMountpoint

`func (o *Dataset) GetMountpoint() string`

GetMountpoint returns the Mountpoint field if non-nil, zero value otherwise.

### GetMountpointOk

`func (o *Dataset) GetMountpointOk() (*string, bool)`

GetMountpointOk returns a tuple with the Mountpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountpoint

`func (o *Dataset) SetMountpoint(v string)`

SetMountpoint sets Mountpoint field to given value.

### HasMountpoint

`func (o *Dataset) HasMountpoint() bool`

HasMountpoint returns a boolean if a field has been set.

### GetEncrypted

`func (o *Dataset) GetEncrypted() bool`

GetEncrypted returns the Encrypted field if non-nil, zero value otherwise.

### GetEncryptedOk

`func (o *Dataset) GetEncryptedOk() (*bool, bool)`

GetEncryptedOk returns a tuple with the Encrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncrypted

`func (o *Dataset) SetEncrypted(v bool)`

SetEncrypted sets Encrypted field to given value.

### HasEncrypted

`func (o *Dataset) HasEncrypted() bool`

HasEncrypted returns a boolean if a field has been set.

### GetEncryptionRoot

`func (o *Dataset) GetEncryptionRoot() string`

GetEncryptionRoot returns the EncryptionRoot field if non-nil, zero value otherwise.

### GetEncryptionRootOk

`func (o *Dataset) GetEncryptionRootOk() (*string, bool)`

GetEncryptionRootOk returns a tuple with the EncryptionRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionRoot

`func (o *Dataset) SetEncryptionRoot(v string)`

SetEncryptionRoot sets EncryptionRoot field to given value.

### HasEncryptionRoot

`func (o *Dataset) HasEncryptionRoot() bool`

HasEncryptionRoot returns a boolean if a field has been set.

### GetKeyLoaded

`func (o *Dataset) GetKeyLoaded() bool`

GetKeyLoaded returns the KeyLoaded field if non-nil, zero value otherwise.

### GetKeyLoadedOk

`func (o *Dataset) GetKeyLoadedOk() (*bool, bool)`

GetKeyLoadedOk returns a tuple with the KeyLoaded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyLoaded

`func (o *Dataset) SetKeyLoaded(v bool)`

SetKeyLoaded sets KeyLoaded field to given value.

### HasKeyLoaded

`func (o *Dataset) HasKeyLoaded() bool`

HasKeyLoaded returns a boolean if a field has been set.

### GetLocked

`func (o *Dataset) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *Dataset) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *Dataset) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *Dataset) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetEncryptionAlgorithm

`func (o *Dataset) GetEncryptionAlgorithm() CompositeValue`

GetEncryptionAlgorithm returns the EncryptionAlgorithm field if non-nil, zero value otherwise.

### GetEncryptionAlgorithmOk

`func (o *Dataset) GetEncryptionAlgorithmOk() (*CompositeValue, bool)`

GetEncryptionAlgorithmOk returns a tuple with the EncryptionAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionAlgorithm

`func (o *Dataset) SetEncryptionAlgorithm(v CompositeValue)`

SetEncryptionAlgorithm sets EncryptionAlgorithm field to given value.

### HasEncryptionAlgorithm

`func (o *Dataset) HasEncryptionAlgorithm() bool`

HasEncryptionAlgorithm returns a boolean if a field has been set.

### GetAclmode

`func (o *Dataset) GetAclmode() CompositeValue`

GetAclmode returns the Aclmode field if non-nil, zero value otherwise.

### GetAclmodeOk

`func (o *Dataset) GetAclmodeOk() (*CompositeValue, bool)`

GetAclmodeOk returns a tuple with the Aclmode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAclmode

`func (o *Dataset) SetAclmode(v CompositeValue)`

SetAclmode sets Aclmode field to given value.

### HasAclmode

`func (o *Dataset) HasAclmode() bool`

HasAclmode returns a boolean if a field has been set.

### GetAcltype

`func (o *Dataset) GetAcltype() CompositeValue`

GetAcltype returns the Acltype field if non-nil, zero value otherwise.

### GetAcltypeOk

`func (o *Dataset) GetAcltypeOk() (*CompositeValue, bool)`

GetAcltypeOk returns a tuple with the Acltype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcltype

`func (o *Dataset) SetAcltype(v CompositeValue)`

SetAcltype sets Acltype field to given value.

### HasAcltype

`func (o *Dataset) HasAcltype() bool`

HasAcltype returns a boolean if a field has been set.

### GetAtime

`func (o *Dataset) GetAtime() CompositeValue`

GetAtime returns the Atime field if non-nil, zero value otherwise.

### GetAtimeOk

`func (o *Dataset) GetAtimeOk() (*CompositeValue, bool)`

GetAtimeOk returns a tuple with the Atime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtime

`func (o *Dataset) SetAtime(v CompositeValue)`

SetAtime sets Atime field to given value.

### HasAtime

`func (o *Dataset) HasAtime() bool`

HasAtime returns a boolean if a field has been set.

### GetAvailable

`func (o *Dataset) GetAvailable() CompositeValue`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *Dataset) GetAvailableOk() (*CompositeValue, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *Dataset) SetAvailable(v CompositeValue)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *Dataset) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetCasesensitivity

`func (o *Dataset) GetCasesensitivity() CompositeValue`

GetCasesensitivity returns the Casesensitivity field if non-nil, zero value otherwise.

### GetCasesensitivityOk

`func (o *Dataset) GetCasesensitivityOk() (*CompositeValue, bool)`

GetCasesensitivityOk returns a tuple with the Casesensitivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCasesensitivity

`func (o *Dataset) SetCasesensitivity(v CompositeValue)`

SetCasesensitivity sets Casesensitivity field to given value.

### HasCasesensitivity

`func (o *Dataset) HasCasesensitivity() bool`

HasCasesensitivity returns a boolean if a field has been set.

### GetComments

`func (o *Dataset) GetComments() CompositeValue`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *Dataset) GetCommentsOk() (*CompositeValue, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *Dataset) SetComments(v CompositeValue)`

SetComments sets Comments field to given value.

### HasComments

`func (o *Dataset) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetCompression

`func (o *Dataset) GetCompression() CompositeValue`

GetCompression returns the Compression field if non-nil, zero value otherwise.

### GetCompressionOk

`func (o *Dataset) GetCompressionOk() (*CompositeValue, bool)`

GetCompressionOk returns a tuple with the Compression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompression

`func (o *Dataset) SetCompression(v CompositeValue)`

SetCompression sets Compression field to given value.

### HasCompression

`func (o *Dataset) HasCompression() bool`

HasCompression returns a boolean if a field has been set.

### GetDeduplication

`func (o *Dataset) GetDeduplication() CompositeValue`

GetDeduplication returns the Deduplication field if non-nil, zero value otherwise.

### GetDeduplicationOk

`func (o *Dataset) GetDeduplicationOk() (*CompositeValue, bool)`

GetDeduplicationOk returns a tuple with the Deduplication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeduplication

`func (o *Dataset) SetDeduplication(v CompositeValue)`

SetDeduplication sets Deduplication field to given value.

### HasDeduplication

`func (o *Dataset) HasDeduplication() bool`

HasDeduplication returns a boolean if a field has been set.

### GetExec

`func (o *Dataset) GetExec() CompositeValue`

GetExec returns the Exec field if non-nil, zero value otherwise.

### GetExecOk

`func (o *Dataset) GetExecOk() (*CompositeValue, bool)`

GetExecOk returns a tuple with the Exec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExec

`func (o *Dataset) SetExec(v CompositeValue)`

SetExec sets Exec field to given value.

### HasExec

`func (o *Dataset) HasExec() bool`

HasExec returns a boolean if a field has been set.

### GetKeyFormat

`func (o *Dataset) GetKeyFormat() CompositeValue`

GetKeyFormat returns the KeyFormat field if non-nil, zero value otherwise.

### GetKeyFormatOk

`func (o *Dataset) GetKeyFormatOk() (*CompositeValue, bool)`

GetKeyFormatOk returns a tuple with the KeyFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyFormat

`func (o *Dataset) SetKeyFormat(v CompositeValue)`

SetKeyFormat sets KeyFormat field to given value.

### HasKeyFormat

`func (o *Dataset) HasKeyFormat() bool`

HasKeyFormat returns a boolean if a field has been set.

### GetManagedby

`func (o *Dataset) GetManagedby() CompositeValue`

GetManagedby returns the Managedby field if non-nil, zero value otherwise.

### GetManagedbyOk

`func (o *Dataset) GetManagedbyOk() (*CompositeValue, bool)`

GetManagedbyOk returns a tuple with the Managedby field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedby

`func (o *Dataset) SetManagedby(v CompositeValue)`

SetManagedby sets Managedby field to given value.

### HasManagedby

`func (o *Dataset) HasManagedby() bool`

HasManagedby returns a boolean if a field has been set.

### GetCopies

`func (o *Dataset) GetCopies() CompositeValue`

GetCopies returns the Copies field if non-nil, zero value otherwise.

### GetCopiesOk

`func (o *Dataset) GetCopiesOk() (*CompositeValue, bool)`

GetCopiesOk returns a tuple with the Copies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopies

`func (o *Dataset) SetCopies(v CompositeValue)`

SetCopies sets Copies field to given value.

### HasCopies

`func (o *Dataset) HasCopies() bool`

HasCopies returns a boolean if a field has been set.

### GetQuota

`func (o *Dataset) GetQuota() CompositeValue`

GetQuota returns the Quota field if non-nil, zero value otherwise.

### GetQuotaOk

`func (o *Dataset) GetQuotaOk() (*CompositeValue, bool)`

GetQuotaOk returns a tuple with the Quota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuota

`func (o *Dataset) SetQuota(v CompositeValue)`

SetQuota sets Quota field to given value.

### HasQuota

`func (o *Dataset) HasQuota() bool`

HasQuota returns a boolean if a field has been set.

### GetQuotaCritical

`func (o *Dataset) GetQuotaCritical() CompositeValue`

GetQuotaCritical returns the QuotaCritical field if non-nil, zero value otherwise.

### GetQuotaCriticalOk

`func (o *Dataset) GetQuotaCriticalOk() (*CompositeValue, bool)`

GetQuotaCriticalOk returns a tuple with the QuotaCritical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaCritical

`func (o *Dataset) SetQuotaCritical(v CompositeValue)`

SetQuotaCritical sets QuotaCritical field to given value.

### HasQuotaCritical

`func (o *Dataset) HasQuotaCritical() bool`

HasQuotaCritical returns a boolean if a field has been set.

### GetQuotaWarning

`func (o *Dataset) GetQuotaWarning() CompositeValue`

GetQuotaWarning returns the QuotaWarning field if non-nil, zero value otherwise.

### GetQuotaWarningOk

`func (o *Dataset) GetQuotaWarningOk() (*CompositeValue, bool)`

GetQuotaWarningOk returns a tuple with the QuotaWarning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaWarning

`func (o *Dataset) SetQuotaWarning(v CompositeValue)`

SetQuotaWarning sets QuotaWarning field to given value.

### HasQuotaWarning

`func (o *Dataset) HasQuotaWarning() bool`

HasQuotaWarning returns a boolean if a field has been set.

### GetReservation

`func (o *Dataset) GetReservation() CompositeValue`

GetReservation returns the Reservation field if non-nil, zero value otherwise.

### GetReservationOk

`func (o *Dataset) GetReservationOk() (*CompositeValue, bool)`

GetReservationOk returns a tuple with the Reservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservation

`func (o *Dataset) SetReservation(v CompositeValue)`

SetReservation sets Reservation field to given value.

### HasReservation

`func (o *Dataset) HasReservation() bool`

HasReservation returns a boolean if a field has been set.

### GetRefreservation

`func (o *Dataset) GetRefreservation() CompositeValue`

GetRefreservation returns the Refreservation field if non-nil, zero value otherwise.

### GetRefreservationOk

`func (o *Dataset) GetRefreservationOk() (*CompositeValue, bool)`

GetRefreservationOk returns a tuple with the Refreservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreservation

`func (o *Dataset) SetRefreservation(v CompositeValue)`

SetRefreservation sets Refreservation field to given value.

### HasRefreservation

`func (o *Dataset) HasRefreservation() bool`

HasRefreservation returns a boolean if a field has been set.

### GetRefquota

`func (o *Dataset) GetRefquota() CompositeValue`

GetRefquota returns the Refquota field if non-nil, zero value otherwise.

### GetRefquotaOk

`func (o *Dataset) GetRefquotaOk() (*CompositeValue, bool)`

GetRefquotaOk returns a tuple with the Refquota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefquota

`func (o *Dataset) SetRefquota(v CompositeValue)`

SetRefquota sets Refquota field to given value.

### HasRefquota

`func (o *Dataset) HasRefquota() bool`

HasRefquota returns a boolean if a field has been set.

### GetRefquotaCritical

`func (o *Dataset) GetRefquotaCritical() CompositeValue`

GetRefquotaCritical returns the RefquotaCritical field if non-nil, zero value otherwise.

### GetRefquotaCriticalOk

`func (o *Dataset) GetRefquotaCriticalOk() (*CompositeValue, bool)`

GetRefquotaCriticalOk returns a tuple with the RefquotaCritical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefquotaCritical

`func (o *Dataset) SetRefquotaCritical(v CompositeValue)`

SetRefquotaCritical sets RefquotaCritical field to given value.

### HasRefquotaCritical

`func (o *Dataset) HasRefquotaCritical() bool`

HasRefquotaCritical returns a boolean if a field has been set.

### GetRefquotaWarning

`func (o *Dataset) GetRefquotaWarning() CompositeValue`

GetRefquotaWarning returns the RefquotaWarning field if non-nil, zero value otherwise.

### GetRefquotaWarningOk

`func (o *Dataset) GetRefquotaWarningOk() (*CompositeValue, bool)`

GetRefquotaWarningOk returns a tuple with the RefquotaWarning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefquotaWarning

`func (o *Dataset) SetRefquotaWarning(v CompositeValue)`

SetRefquotaWarning sets RefquotaWarning field to given value.

### HasRefquotaWarning

`func (o *Dataset) HasRefquotaWarning() bool`

HasRefquotaWarning returns a boolean if a field has been set.

### GetReadonly

`func (o *Dataset) GetReadonly() CompositeValue`

GetReadonly returns the Readonly field if non-nil, zero value otherwise.

### GetReadonlyOk

`func (o *Dataset) GetReadonlyOk() (*CompositeValue, bool)`

GetReadonlyOk returns a tuple with the Readonly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadonly

`func (o *Dataset) SetReadonly(v CompositeValue)`

SetReadonly sets Readonly field to given value.

### HasReadonly

`func (o *Dataset) HasReadonly() bool`

HasReadonly returns a boolean if a field has been set.

### GetRecordsize

`func (o *Dataset) GetRecordsize() CompositeValue`

GetRecordsize returns the Recordsize field if non-nil, zero value otherwise.

### GetRecordsizeOk

`func (o *Dataset) GetRecordsizeOk() (*CompositeValue, bool)`

GetRecordsizeOk returns a tuple with the Recordsize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordsize

`func (o *Dataset) SetRecordsize(v CompositeValue)`

SetRecordsize sets Recordsize field to given value.

### HasRecordsize

`func (o *Dataset) HasRecordsize() bool`

HasRecordsize returns a boolean if a field has been set.

### GetSync

`func (o *Dataset) GetSync() CompositeValue`

GetSync returns the Sync field if non-nil, zero value otherwise.

### GetSyncOk

`func (o *Dataset) GetSyncOk() (*CompositeValue, bool)`

GetSyncOk returns a tuple with the Sync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSync

`func (o *Dataset) SetSync(v CompositeValue)`

SetSync sets Sync field to given value.

### HasSync

`func (o *Dataset) HasSync() bool`

HasSync returns a boolean if a field has been set.

### GetSnapdir

`func (o *Dataset) GetSnapdir() CompositeValue`

GetSnapdir returns the Snapdir field if non-nil, zero value otherwise.

### GetSnapdirOk

`func (o *Dataset) GetSnapdirOk() (*CompositeValue, bool)`

GetSnapdirOk returns a tuple with the Snapdir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapdir

`func (o *Dataset) SetSnapdir(v CompositeValue)`

SetSnapdir sets Snapdir field to given value.

### HasSnapdir

`func (o *Dataset) HasSnapdir() bool`

HasSnapdir returns a boolean if a field has been set.

### GetUsed

`func (o *Dataset) GetUsed() CompositeValue`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *Dataset) GetUsedOk() (*CompositeValue, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *Dataset) SetUsed(v CompositeValue)`

SetUsed sets Used field to given value.

### HasUsed

`func (o *Dataset) HasUsed() bool`

HasUsed returns a boolean if a field has been set.

### GetPbkdf2iters

`func (o *Dataset) GetPbkdf2iters() CompositeValue`

GetPbkdf2iters returns the Pbkdf2iters field if non-nil, zero value otherwise.

### GetPbkdf2itersOk

`func (o *Dataset) GetPbkdf2itersOk() (*CompositeValue, bool)`

GetPbkdf2itersOk returns a tuple with the Pbkdf2iters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPbkdf2iters

`func (o *Dataset) SetPbkdf2iters(v CompositeValue)`

SetPbkdf2iters sets Pbkdf2iters field to given value.

### HasPbkdf2iters

`func (o *Dataset) HasPbkdf2iters() bool`

HasPbkdf2iters returns a boolean if a field has been set.

### GetOrigin

`func (o *Dataset) GetOrigin() CompositeValue`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *Dataset) GetOriginOk() (*CompositeValue, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *Dataset) SetOrigin(v CompositeValue)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *Dataset) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### GetXattr

`func (o *Dataset) GetXattr() CompositeValue`

GetXattr returns the Xattr field if non-nil, zero value otherwise.

### GetXattrOk

`func (o *Dataset) GetXattrOk() (*CompositeValue, bool)`

GetXattrOk returns a tuple with the Xattr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXattr

`func (o *Dataset) SetXattr(v CompositeValue)`

SetXattr sets Xattr field to given value.

### HasXattr

`func (o *Dataset) HasXattr() bool`

HasXattr returns a boolean if a field has been set.

### GetVolsize

`func (o *Dataset) GetVolsize() CompositeValue`

GetVolsize returns the Volsize field if non-nil, zero value otherwise.

### GetVolsizeOk

`func (o *Dataset) GetVolsizeOk() (*CompositeValue, bool)`

GetVolsizeOk returns a tuple with the Volsize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolsize

`func (o *Dataset) SetVolsize(v CompositeValue)`

SetVolsize sets Volsize field to given value.

### HasVolsize

`func (o *Dataset) HasVolsize() bool`

HasVolsize returns a boolean if a field has been set.

### GetVolblocksize

`func (o *Dataset) GetVolblocksize() CompositeValue`

GetVolblocksize returns the Volblocksize field if non-nil, zero value otherwise.

### GetVolblocksizeOk

`func (o *Dataset) GetVolblocksizeOk() (*CompositeValue, bool)`

GetVolblocksizeOk returns a tuple with the Volblocksize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolblocksize

`func (o *Dataset) SetVolblocksize(v CompositeValue)`

SetVolblocksize sets Volblocksize field to given value.

### HasVolblocksize

`func (o *Dataset) HasVolblocksize() bool`

HasVolblocksize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


