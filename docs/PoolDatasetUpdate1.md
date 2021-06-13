# PoolDatasetUpdate1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Volsize** | Pointer to **int32** |  | [optional] 
**ForceSize** | Pointer to **bool** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Sync** | Pointer to **string** |  | [optional] 
**Compression** | Pointer to **string** |  | [optional] 
**Atime** | Pointer to **string** |  | [optional] 
**Exec** | Pointer to **string** |  | [optional] 
**Managedby** | Pointer to **string** |  | [optional] 
**Quota** | Pointer to **NullableInt32** |  | [optional] 
**QuotaWarning** | Pointer to [**AnyOfintegerstring**](anyOf&lt;integer,string&gt;.md) |  | [optional] 
**QuotaCritical** | Pointer to [**AnyOfintegerstring**](anyOf&lt;integer,string&gt;.md) |  | [optional] 
**Refquota** | Pointer to **NullableInt32** |  | [optional] 
**RefquotaWarning** | Pointer to [**AnyOfintegerstring**](anyOf&lt;integer,string&gt;.md) |  | [optional] 
**RefquotaCritical** | Pointer to [**AnyOfintegerstring**](anyOf&lt;integer,string&gt;.md) |  | [optional] 
**Reservation** | Pointer to **int32** |  | [optional] 
**Refreservation** | Pointer to **int32** |  | [optional] 
**SpecialSmallBlockSize** | Pointer to **int32** |  | [optional] 
**Copies** | Pointer to **int32** |  | [optional] 
**Snapdir** | Pointer to **string** |  | [optional] 
**Deduplication** | Pointer to **string** |  | [optional] 
**Readonly** | Pointer to **string** |  | [optional] 
**Recordsize** | Pointer to **string** |  | [optional] 
**Aclmode** | Pointer to **string** |  | [optional] 
**Acltype** | Pointer to **string** |  | [optional] 
**Xattr** | Pointer to **string** |  | [optional] 

## Methods

### NewPoolDatasetUpdate1

`func NewPoolDatasetUpdate1() *PoolDatasetUpdate1`

NewPoolDatasetUpdate1 instantiates a new PoolDatasetUpdate1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoolDatasetUpdate1WithDefaults

`func NewPoolDatasetUpdate1WithDefaults() *PoolDatasetUpdate1`

NewPoolDatasetUpdate1WithDefaults instantiates a new PoolDatasetUpdate1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVolsize

`func (o *PoolDatasetUpdate1) GetVolsize() int32`

GetVolsize returns the Volsize field if non-nil, zero value otherwise.

### GetVolsizeOk

`func (o *PoolDatasetUpdate1) GetVolsizeOk() (*int32, bool)`

GetVolsizeOk returns a tuple with the Volsize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolsize

`func (o *PoolDatasetUpdate1) SetVolsize(v int32)`

SetVolsize sets Volsize field to given value.

### HasVolsize

`func (o *PoolDatasetUpdate1) HasVolsize() bool`

HasVolsize returns a boolean if a field has been set.

### GetForceSize

`func (o *PoolDatasetUpdate1) GetForceSize() bool`

GetForceSize returns the ForceSize field if non-nil, zero value otherwise.

### GetForceSizeOk

`func (o *PoolDatasetUpdate1) GetForceSizeOk() (*bool, bool)`

GetForceSizeOk returns a tuple with the ForceSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceSize

`func (o *PoolDatasetUpdate1) SetForceSize(v bool)`

SetForceSize sets ForceSize field to given value.

### HasForceSize

`func (o *PoolDatasetUpdate1) HasForceSize() bool`

HasForceSize returns a boolean if a field has been set.

### GetComments

`func (o *PoolDatasetUpdate1) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *PoolDatasetUpdate1) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *PoolDatasetUpdate1) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *PoolDatasetUpdate1) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetSync

`func (o *PoolDatasetUpdate1) GetSync() string`

GetSync returns the Sync field if non-nil, zero value otherwise.

### GetSyncOk

`func (o *PoolDatasetUpdate1) GetSyncOk() (*string, bool)`

GetSyncOk returns a tuple with the Sync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSync

`func (o *PoolDatasetUpdate1) SetSync(v string)`

SetSync sets Sync field to given value.

### HasSync

`func (o *PoolDatasetUpdate1) HasSync() bool`

HasSync returns a boolean if a field has been set.

### GetCompression

`func (o *PoolDatasetUpdate1) GetCompression() string`

GetCompression returns the Compression field if non-nil, zero value otherwise.

### GetCompressionOk

`func (o *PoolDatasetUpdate1) GetCompressionOk() (*string, bool)`

GetCompressionOk returns a tuple with the Compression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompression

`func (o *PoolDatasetUpdate1) SetCompression(v string)`

SetCompression sets Compression field to given value.

### HasCompression

`func (o *PoolDatasetUpdate1) HasCompression() bool`

HasCompression returns a boolean if a field has been set.

### GetAtime

`func (o *PoolDatasetUpdate1) GetAtime() string`

GetAtime returns the Atime field if non-nil, zero value otherwise.

### GetAtimeOk

`func (o *PoolDatasetUpdate1) GetAtimeOk() (*string, bool)`

GetAtimeOk returns a tuple with the Atime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtime

`func (o *PoolDatasetUpdate1) SetAtime(v string)`

SetAtime sets Atime field to given value.

### HasAtime

`func (o *PoolDatasetUpdate1) HasAtime() bool`

HasAtime returns a boolean if a field has been set.

### GetExec

`func (o *PoolDatasetUpdate1) GetExec() string`

GetExec returns the Exec field if non-nil, zero value otherwise.

### GetExecOk

`func (o *PoolDatasetUpdate1) GetExecOk() (*string, bool)`

GetExecOk returns a tuple with the Exec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExec

`func (o *PoolDatasetUpdate1) SetExec(v string)`

SetExec sets Exec field to given value.

### HasExec

`func (o *PoolDatasetUpdate1) HasExec() bool`

HasExec returns a boolean if a field has been set.

### GetManagedby

`func (o *PoolDatasetUpdate1) GetManagedby() string`

GetManagedby returns the Managedby field if non-nil, zero value otherwise.

### GetManagedbyOk

`func (o *PoolDatasetUpdate1) GetManagedbyOk() (*string, bool)`

GetManagedbyOk returns a tuple with the Managedby field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedby

`func (o *PoolDatasetUpdate1) SetManagedby(v string)`

SetManagedby sets Managedby field to given value.

### HasManagedby

`func (o *PoolDatasetUpdate1) HasManagedby() bool`

HasManagedby returns a boolean if a field has been set.

### GetQuota

`func (o *PoolDatasetUpdate1) GetQuota() int32`

GetQuota returns the Quota field if non-nil, zero value otherwise.

### GetQuotaOk

`func (o *PoolDatasetUpdate1) GetQuotaOk() (*int32, bool)`

GetQuotaOk returns a tuple with the Quota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuota

`func (o *PoolDatasetUpdate1) SetQuota(v int32)`

SetQuota sets Quota field to given value.

### HasQuota

`func (o *PoolDatasetUpdate1) HasQuota() bool`

HasQuota returns a boolean if a field has been set.

### SetQuotaNil

`func (o *PoolDatasetUpdate1) SetQuotaNil(b bool)`

 SetQuotaNil sets the value for Quota to be an explicit nil

### UnsetQuota
`func (o *PoolDatasetUpdate1) UnsetQuota()`

UnsetQuota ensures that no value is present for Quota, not even an explicit nil
### GetQuotaWarning

`func (o *PoolDatasetUpdate1) GetQuotaWarning() AnyOfintegerstring`

GetQuotaWarning returns the QuotaWarning field if non-nil, zero value otherwise.

### GetQuotaWarningOk

`func (o *PoolDatasetUpdate1) GetQuotaWarningOk() (*AnyOfintegerstring, bool)`

GetQuotaWarningOk returns a tuple with the QuotaWarning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaWarning

`func (o *PoolDatasetUpdate1) SetQuotaWarning(v AnyOfintegerstring)`

SetQuotaWarning sets QuotaWarning field to given value.

### HasQuotaWarning

`func (o *PoolDatasetUpdate1) HasQuotaWarning() bool`

HasQuotaWarning returns a boolean if a field has been set.

### GetQuotaCritical

`func (o *PoolDatasetUpdate1) GetQuotaCritical() AnyOfintegerstring`

GetQuotaCritical returns the QuotaCritical field if non-nil, zero value otherwise.

### GetQuotaCriticalOk

`func (o *PoolDatasetUpdate1) GetQuotaCriticalOk() (*AnyOfintegerstring, bool)`

GetQuotaCriticalOk returns a tuple with the QuotaCritical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaCritical

`func (o *PoolDatasetUpdate1) SetQuotaCritical(v AnyOfintegerstring)`

SetQuotaCritical sets QuotaCritical field to given value.

### HasQuotaCritical

`func (o *PoolDatasetUpdate1) HasQuotaCritical() bool`

HasQuotaCritical returns a boolean if a field has been set.

### GetRefquota

`func (o *PoolDatasetUpdate1) GetRefquota() int32`

GetRefquota returns the Refquota field if non-nil, zero value otherwise.

### GetRefquotaOk

`func (o *PoolDatasetUpdate1) GetRefquotaOk() (*int32, bool)`

GetRefquotaOk returns a tuple with the Refquota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefquota

`func (o *PoolDatasetUpdate1) SetRefquota(v int32)`

SetRefquota sets Refquota field to given value.

### HasRefquota

`func (o *PoolDatasetUpdate1) HasRefquota() bool`

HasRefquota returns a boolean if a field has been set.

### SetRefquotaNil

`func (o *PoolDatasetUpdate1) SetRefquotaNil(b bool)`

 SetRefquotaNil sets the value for Refquota to be an explicit nil

### UnsetRefquota
`func (o *PoolDatasetUpdate1) UnsetRefquota()`

UnsetRefquota ensures that no value is present for Refquota, not even an explicit nil
### GetRefquotaWarning

`func (o *PoolDatasetUpdate1) GetRefquotaWarning() AnyOfintegerstring`

GetRefquotaWarning returns the RefquotaWarning field if non-nil, zero value otherwise.

### GetRefquotaWarningOk

`func (o *PoolDatasetUpdate1) GetRefquotaWarningOk() (*AnyOfintegerstring, bool)`

GetRefquotaWarningOk returns a tuple with the RefquotaWarning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefquotaWarning

`func (o *PoolDatasetUpdate1) SetRefquotaWarning(v AnyOfintegerstring)`

SetRefquotaWarning sets RefquotaWarning field to given value.

### HasRefquotaWarning

`func (o *PoolDatasetUpdate1) HasRefquotaWarning() bool`

HasRefquotaWarning returns a boolean if a field has been set.

### GetRefquotaCritical

`func (o *PoolDatasetUpdate1) GetRefquotaCritical() AnyOfintegerstring`

GetRefquotaCritical returns the RefquotaCritical field if non-nil, zero value otherwise.

### GetRefquotaCriticalOk

`func (o *PoolDatasetUpdate1) GetRefquotaCriticalOk() (*AnyOfintegerstring, bool)`

GetRefquotaCriticalOk returns a tuple with the RefquotaCritical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefquotaCritical

`func (o *PoolDatasetUpdate1) SetRefquotaCritical(v AnyOfintegerstring)`

SetRefquotaCritical sets RefquotaCritical field to given value.

### HasRefquotaCritical

`func (o *PoolDatasetUpdate1) HasRefquotaCritical() bool`

HasRefquotaCritical returns a boolean if a field has been set.

### GetReservation

`func (o *PoolDatasetUpdate1) GetReservation() int32`

GetReservation returns the Reservation field if non-nil, zero value otherwise.

### GetReservationOk

`func (o *PoolDatasetUpdate1) GetReservationOk() (*int32, bool)`

GetReservationOk returns a tuple with the Reservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservation

`func (o *PoolDatasetUpdate1) SetReservation(v int32)`

SetReservation sets Reservation field to given value.

### HasReservation

`func (o *PoolDatasetUpdate1) HasReservation() bool`

HasReservation returns a boolean if a field has been set.

### GetRefreservation

`func (o *PoolDatasetUpdate1) GetRefreservation() int32`

GetRefreservation returns the Refreservation field if non-nil, zero value otherwise.

### GetRefreservationOk

`func (o *PoolDatasetUpdate1) GetRefreservationOk() (*int32, bool)`

GetRefreservationOk returns a tuple with the Refreservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreservation

`func (o *PoolDatasetUpdate1) SetRefreservation(v int32)`

SetRefreservation sets Refreservation field to given value.

### HasRefreservation

`func (o *PoolDatasetUpdate1) HasRefreservation() bool`

HasRefreservation returns a boolean if a field has been set.

### GetSpecialSmallBlockSize

`func (o *PoolDatasetUpdate1) GetSpecialSmallBlockSize() int32`

GetSpecialSmallBlockSize returns the SpecialSmallBlockSize field if non-nil, zero value otherwise.

### GetSpecialSmallBlockSizeOk

`func (o *PoolDatasetUpdate1) GetSpecialSmallBlockSizeOk() (*int32, bool)`

GetSpecialSmallBlockSizeOk returns a tuple with the SpecialSmallBlockSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialSmallBlockSize

`func (o *PoolDatasetUpdate1) SetSpecialSmallBlockSize(v int32)`

SetSpecialSmallBlockSize sets SpecialSmallBlockSize field to given value.

### HasSpecialSmallBlockSize

`func (o *PoolDatasetUpdate1) HasSpecialSmallBlockSize() bool`

HasSpecialSmallBlockSize returns a boolean if a field has been set.

### GetCopies

`func (o *PoolDatasetUpdate1) GetCopies() int32`

GetCopies returns the Copies field if non-nil, zero value otherwise.

### GetCopiesOk

`func (o *PoolDatasetUpdate1) GetCopiesOk() (*int32, bool)`

GetCopiesOk returns a tuple with the Copies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopies

`func (o *PoolDatasetUpdate1) SetCopies(v int32)`

SetCopies sets Copies field to given value.

### HasCopies

`func (o *PoolDatasetUpdate1) HasCopies() bool`

HasCopies returns a boolean if a field has been set.

### GetSnapdir

`func (o *PoolDatasetUpdate1) GetSnapdir() string`

GetSnapdir returns the Snapdir field if non-nil, zero value otherwise.

### GetSnapdirOk

`func (o *PoolDatasetUpdate1) GetSnapdirOk() (*string, bool)`

GetSnapdirOk returns a tuple with the Snapdir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapdir

`func (o *PoolDatasetUpdate1) SetSnapdir(v string)`

SetSnapdir sets Snapdir field to given value.

### HasSnapdir

`func (o *PoolDatasetUpdate1) HasSnapdir() bool`

HasSnapdir returns a boolean if a field has been set.

### GetDeduplication

`func (o *PoolDatasetUpdate1) GetDeduplication() string`

GetDeduplication returns the Deduplication field if non-nil, zero value otherwise.

### GetDeduplicationOk

`func (o *PoolDatasetUpdate1) GetDeduplicationOk() (*string, bool)`

GetDeduplicationOk returns a tuple with the Deduplication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeduplication

`func (o *PoolDatasetUpdate1) SetDeduplication(v string)`

SetDeduplication sets Deduplication field to given value.

### HasDeduplication

`func (o *PoolDatasetUpdate1) HasDeduplication() bool`

HasDeduplication returns a boolean if a field has been set.

### GetReadonly

`func (o *PoolDatasetUpdate1) GetReadonly() string`

GetReadonly returns the Readonly field if non-nil, zero value otherwise.

### GetReadonlyOk

`func (o *PoolDatasetUpdate1) GetReadonlyOk() (*string, bool)`

GetReadonlyOk returns a tuple with the Readonly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadonly

`func (o *PoolDatasetUpdate1) SetReadonly(v string)`

SetReadonly sets Readonly field to given value.

### HasReadonly

`func (o *PoolDatasetUpdate1) HasReadonly() bool`

HasReadonly returns a boolean if a field has been set.

### GetRecordsize

`func (o *PoolDatasetUpdate1) GetRecordsize() string`

GetRecordsize returns the Recordsize field if non-nil, zero value otherwise.

### GetRecordsizeOk

`func (o *PoolDatasetUpdate1) GetRecordsizeOk() (*string, bool)`

GetRecordsizeOk returns a tuple with the Recordsize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordsize

`func (o *PoolDatasetUpdate1) SetRecordsize(v string)`

SetRecordsize sets Recordsize field to given value.

### HasRecordsize

`func (o *PoolDatasetUpdate1) HasRecordsize() bool`

HasRecordsize returns a boolean if a field has been set.

### GetAclmode

`func (o *PoolDatasetUpdate1) GetAclmode() string`

GetAclmode returns the Aclmode field if non-nil, zero value otherwise.

### GetAclmodeOk

`func (o *PoolDatasetUpdate1) GetAclmodeOk() (*string, bool)`

GetAclmodeOk returns a tuple with the Aclmode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAclmode

`func (o *PoolDatasetUpdate1) SetAclmode(v string)`

SetAclmode sets Aclmode field to given value.

### HasAclmode

`func (o *PoolDatasetUpdate1) HasAclmode() bool`

HasAclmode returns a boolean if a field has been set.

### GetAcltype

`func (o *PoolDatasetUpdate1) GetAcltype() string`

GetAcltype returns the Acltype field if non-nil, zero value otherwise.

### GetAcltypeOk

`func (o *PoolDatasetUpdate1) GetAcltypeOk() (*string, bool)`

GetAcltypeOk returns a tuple with the Acltype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcltype

`func (o *PoolDatasetUpdate1) SetAcltype(v string)`

SetAcltype sets Acltype field to given value.

### HasAcltype

`func (o *PoolDatasetUpdate1) HasAcltype() bool`

HasAcltype returns a boolean if a field has been set.

### GetXattr

`func (o *PoolDatasetUpdate1) GetXattr() string`

GetXattr returns the Xattr field if non-nil, zero value otherwise.

### GetXattrOk

`func (o *PoolDatasetUpdate1) GetXattrOk() (*string, bool)`

GetXattrOk returns a tuple with the Xattr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXattr

`func (o *PoolDatasetUpdate1) SetXattr(v string)`

SetXattr sets Xattr field to given value.

### HasXattr

`func (o *PoolDatasetUpdate1) HasXattr() bool`

HasXattr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


