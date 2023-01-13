# UpdateDatasetParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Atime** | Pointer to **string** |  | [optional] 
**Aclmode** | Pointer to **string** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Compression** | Pointer to **string** |  | [optional] 
**Copies** | Pointer to **int32** |  | [optional] 
**Deduplication** | Pointer to **string** |  | [optional] 
**Exec** | Pointer to **string** |  | [optional] 
**ForceSize** | Pointer to **bool** |  | [optional] 
**Quota** | Pointer to **int64** |  | [optional] 
**Readonly** | Pointer to **string** |  | [optional] 
**Recordsize** | Pointer to **string** |  | [optional] 
**Refquota** | Pointer to **int64** |  | [optional] 
**Refreservation** | Pointer to **int64** |  | [optional] 
**Volsize** | Pointer to **int64** |  | [optional] 
**Snapdir** | Pointer to **string** |  | [optional] 
**Sync** | Pointer to **string** |  | [optional] 

## Methods

### NewUpdateDatasetParams

`func NewUpdateDatasetParams() *UpdateDatasetParams`

NewUpdateDatasetParams instantiates a new UpdateDatasetParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDatasetParamsWithDefaults

`func NewUpdateDatasetParamsWithDefaults() *UpdateDatasetParams`

NewUpdateDatasetParamsWithDefaults instantiates a new UpdateDatasetParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAtime

`func (o *UpdateDatasetParams) GetAtime() string`

GetAtime returns the Atime field if non-nil, zero value otherwise.

### GetAtimeOk

`func (o *UpdateDatasetParams) GetAtimeOk() (*string, bool)`

GetAtimeOk returns a tuple with the Atime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtime

`func (o *UpdateDatasetParams) SetAtime(v string)`

SetAtime sets Atime field to given value.

### HasAtime

`func (o *UpdateDatasetParams) HasAtime() bool`

HasAtime returns a boolean if a field has been set.

### GetAclmode

`func (o *UpdateDatasetParams) GetAclmode() string`

GetAclmode returns the Aclmode field if non-nil, zero value otherwise.

### GetAclmodeOk

`func (o *UpdateDatasetParams) GetAclmodeOk() (*string, bool)`

GetAclmodeOk returns a tuple with the Aclmode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAclmode

`func (o *UpdateDatasetParams) SetAclmode(v string)`

SetAclmode sets Aclmode field to given value.

### HasAclmode

`func (o *UpdateDatasetParams) HasAclmode() bool`

HasAclmode returns a boolean if a field has been set.

### GetComments

`func (o *UpdateDatasetParams) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *UpdateDatasetParams) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *UpdateDatasetParams) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *UpdateDatasetParams) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetCompression

`func (o *UpdateDatasetParams) GetCompression() string`

GetCompression returns the Compression field if non-nil, zero value otherwise.

### GetCompressionOk

`func (o *UpdateDatasetParams) GetCompressionOk() (*string, bool)`

GetCompressionOk returns a tuple with the Compression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompression

`func (o *UpdateDatasetParams) SetCompression(v string)`

SetCompression sets Compression field to given value.

### HasCompression

`func (o *UpdateDatasetParams) HasCompression() bool`

HasCompression returns a boolean if a field has been set.

### GetCopies

`func (o *UpdateDatasetParams) GetCopies() int32`

GetCopies returns the Copies field if non-nil, zero value otherwise.

### GetCopiesOk

`func (o *UpdateDatasetParams) GetCopiesOk() (*int32, bool)`

GetCopiesOk returns a tuple with the Copies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopies

`func (o *UpdateDatasetParams) SetCopies(v int32)`

SetCopies sets Copies field to given value.

### HasCopies

`func (o *UpdateDatasetParams) HasCopies() bool`

HasCopies returns a boolean if a field has been set.

### GetDeduplication

`func (o *UpdateDatasetParams) GetDeduplication() string`

GetDeduplication returns the Deduplication field if non-nil, zero value otherwise.

### GetDeduplicationOk

`func (o *UpdateDatasetParams) GetDeduplicationOk() (*string, bool)`

GetDeduplicationOk returns a tuple with the Deduplication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeduplication

`func (o *UpdateDatasetParams) SetDeduplication(v string)`

SetDeduplication sets Deduplication field to given value.

### HasDeduplication

`func (o *UpdateDatasetParams) HasDeduplication() bool`

HasDeduplication returns a boolean if a field has been set.

### GetExec

`func (o *UpdateDatasetParams) GetExec() string`

GetExec returns the Exec field if non-nil, zero value otherwise.

### GetExecOk

`func (o *UpdateDatasetParams) GetExecOk() (*string, bool)`

GetExecOk returns a tuple with the Exec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExec

`func (o *UpdateDatasetParams) SetExec(v string)`

SetExec sets Exec field to given value.

### HasExec

`func (o *UpdateDatasetParams) HasExec() bool`

HasExec returns a boolean if a field has been set.

### GetForceSize

`func (o *UpdateDatasetParams) GetForceSize() bool`

GetForceSize returns the ForceSize field if non-nil, zero value otherwise.

### GetForceSizeOk

`func (o *UpdateDatasetParams) GetForceSizeOk() (*bool, bool)`

GetForceSizeOk returns a tuple with the ForceSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceSize

`func (o *UpdateDatasetParams) SetForceSize(v bool)`

SetForceSize sets ForceSize field to given value.

### HasForceSize

`func (o *UpdateDatasetParams) HasForceSize() bool`

HasForceSize returns a boolean if a field has been set.

### GetQuota

`func (o *UpdateDatasetParams) GetQuota() int64`

GetQuota returns the Quota field if non-nil, zero value otherwise.

### GetQuotaOk

`func (o *UpdateDatasetParams) GetQuotaOk() (*int64, bool)`

GetQuotaOk returns a tuple with the Quota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuota

`func (o *UpdateDatasetParams) SetQuota(v int64)`

SetQuota sets Quota field to given value.

### HasQuota

`func (o *UpdateDatasetParams) HasQuota() bool`

HasQuota returns a boolean if a field has been set.

### GetReadonly

`func (o *UpdateDatasetParams) GetReadonly() string`

GetReadonly returns the Readonly field if non-nil, zero value otherwise.

### GetReadonlyOk

`func (o *UpdateDatasetParams) GetReadonlyOk() (*string, bool)`

GetReadonlyOk returns a tuple with the Readonly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadonly

`func (o *UpdateDatasetParams) SetReadonly(v string)`

SetReadonly sets Readonly field to given value.

### HasReadonly

`func (o *UpdateDatasetParams) HasReadonly() bool`

HasReadonly returns a boolean if a field has been set.

### GetRecordsize

`func (o *UpdateDatasetParams) GetRecordsize() string`

GetRecordsize returns the Recordsize field if non-nil, zero value otherwise.

### GetRecordsizeOk

`func (o *UpdateDatasetParams) GetRecordsizeOk() (*string, bool)`

GetRecordsizeOk returns a tuple with the Recordsize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordsize

`func (o *UpdateDatasetParams) SetRecordsize(v string)`

SetRecordsize sets Recordsize field to given value.

### HasRecordsize

`func (o *UpdateDatasetParams) HasRecordsize() bool`

HasRecordsize returns a boolean if a field has been set.

### GetRefquota

`func (o *UpdateDatasetParams) GetRefquota() int64`

GetRefquota returns the Refquota field if non-nil, zero value otherwise.

### GetRefquotaOk

`func (o *UpdateDatasetParams) GetRefquotaOk() (*int64, bool)`

GetRefquotaOk returns a tuple with the Refquota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefquota

`func (o *UpdateDatasetParams) SetRefquota(v int64)`

SetRefquota sets Refquota field to given value.

### HasRefquota

`func (o *UpdateDatasetParams) HasRefquota() bool`

HasRefquota returns a boolean if a field has been set.

### GetRefreservation

`func (o *UpdateDatasetParams) GetRefreservation() int64`

GetRefreservation returns the Refreservation field if non-nil, zero value otherwise.

### GetRefreservationOk

`func (o *UpdateDatasetParams) GetRefreservationOk() (*int64, bool)`

GetRefreservationOk returns a tuple with the Refreservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreservation

`func (o *UpdateDatasetParams) SetRefreservation(v int64)`

SetRefreservation sets Refreservation field to given value.

### HasRefreservation

`func (o *UpdateDatasetParams) HasRefreservation() bool`

HasRefreservation returns a boolean if a field has been set.

### GetVolsize

`func (o *UpdateDatasetParams) GetVolsize() int64`

GetVolsize returns the Volsize field if non-nil, zero value otherwise.

### GetVolsizeOk

`func (o *UpdateDatasetParams) GetVolsizeOk() (*int64, bool)`

GetVolsizeOk returns a tuple with the Volsize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolsize

`func (o *UpdateDatasetParams) SetVolsize(v int64)`

SetVolsize sets Volsize field to given value.

### HasVolsize

`func (o *UpdateDatasetParams) HasVolsize() bool`

HasVolsize returns a boolean if a field has been set.

### GetSnapdir

`func (o *UpdateDatasetParams) GetSnapdir() string`

GetSnapdir returns the Snapdir field if non-nil, zero value otherwise.

### GetSnapdirOk

`func (o *UpdateDatasetParams) GetSnapdirOk() (*string, bool)`

GetSnapdirOk returns a tuple with the Snapdir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapdir

`func (o *UpdateDatasetParams) SetSnapdir(v string)`

SetSnapdir sets Snapdir field to given value.

### HasSnapdir

`func (o *UpdateDatasetParams) HasSnapdir() bool`

HasSnapdir returns a boolean if a field has been set.

### GetSync

`func (o *UpdateDatasetParams) GetSync() string`

GetSync returns the Sync field if non-nil, zero value otherwise.

### GetSyncOk

`func (o *UpdateDatasetParams) GetSyncOk() (*string, bool)`

GetSyncOk returns a tuple with the Sync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSync

`func (o *UpdateDatasetParams) SetSync(v string)`

SetSync sets Sync field to given value.

### HasSync

`func (o *UpdateDatasetParams) HasSync() bool`

HasSync returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


