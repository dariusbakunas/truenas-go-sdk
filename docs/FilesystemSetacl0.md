# FilesystemSetacl0

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to **string** |  | [optional] 
**Uid** | Pointer to **NullableInt32** |  | [optional] 
**Gid** | Pointer to **NullableInt32** |  | [optional] 
**Dacl** | Pointer to [**[]OneOfobjectobject**](OneOfobjectobject.md) |  | [optional] 
**Nfs41Flags** | Pointer to [**FilesystemSetacl0Nfs41Flags**](FilesystemSetacl0Nfs41Flags.md) |  | [optional] 
**Acltype** | Pointer to **string** |  | [optional] 
**Options** | Pointer to [**FilesystemSetacl0Options**](FilesystemSetacl0Options.md) |  | [optional] 

## Methods

### NewFilesystemSetacl0

`func NewFilesystemSetacl0() *FilesystemSetacl0`

NewFilesystemSetacl0 instantiates a new FilesystemSetacl0 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilesystemSetacl0WithDefaults

`func NewFilesystemSetacl0WithDefaults() *FilesystemSetacl0`

NewFilesystemSetacl0WithDefaults instantiates a new FilesystemSetacl0 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *FilesystemSetacl0) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *FilesystemSetacl0) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *FilesystemSetacl0) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *FilesystemSetacl0) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetUid

`func (o *FilesystemSetacl0) GetUid() int32`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *FilesystemSetacl0) GetUidOk() (*int32, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *FilesystemSetacl0) SetUid(v int32)`

SetUid sets Uid field to given value.

### HasUid

`func (o *FilesystemSetacl0) HasUid() bool`

HasUid returns a boolean if a field has been set.

### SetUidNil

`func (o *FilesystemSetacl0) SetUidNil(b bool)`

 SetUidNil sets the value for Uid to be an explicit nil

### UnsetUid
`func (o *FilesystemSetacl0) UnsetUid()`

UnsetUid ensures that no value is present for Uid, not even an explicit nil
### GetGid

`func (o *FilesystemSetacl0) GetGid() int32`

GetGid returns the Gid field if non-nil, zero value otherwise.

### GetGidOk

`func (o *FilesystemSetacl0) GetGidOk() (*int32, bool)`

GetGidOk returns a tuple with the Gid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGid

`func (o *FilesystemSetacl0) SetGid(v int32)`

SetGid sets Gid field to given value.

### HasGid

`func (o *FilesystemSetacl0) HasGid() bool`

HasGid returns a boolean if a field has been set.

### SetGidNil

`func (o *FilesystemSetacl0) SetGidNil(b bool)`

 SetGidNil sets the value for Gid to be an explicit nil

### UnsetGid
`func (o *FilesystemSetacl0) UnsetGid()`

UnsetGid ensures that no value is present for Gid, not even an explicit nil
### GetDacl

`func (o *FilesystemSetacl0) GetDacl() []OneOfobjectobject`

GetDacl returns the Dacl field if non-nil, zero value otherwise.

### GetDaclOk

`func (o *FilesystemSetacl0) GetDaclOk() (*[]OneOfobjectobject, bool)`

GetDaclOk returns a tuple with the Dacl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDacl

`func (o *FilesystemSetacl0) SetDacl(v []OneOfobjectobject)`

SetDacl sets Dacl field to given value.

### HasDacl

`func (o *FilesystemSetacl0) HasDacl() bool`

HasDacl returns a boolean if a field has been set.

### GetNfs41Flags

`func (o *FilesystemSetacl0) GetNfs41Flags() FilesystemSetacl0Nfs41Flags`

GetNfs41Flags returns the Nfs41Flags field if non-nil, zero value otherwise.

### GetNfs41FlagsOk

`func (o *FilesystemSetacl0) GetNfs41FlagsOk() (*FilesystemSetacl0Nfs41Flags, bool)`

GetNfs41FlagsOk returns a tuple with the Nfs41Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfs41Flags

`func (o *FilesystemSetacl0) SetNfs41Flags(v FilesystemSetacl0Nfs41Flags)`

SetNfs41Flags sets Nfs41Flags field to given value.

### HasNfs41Flags

`func (o *FilesystemSetacl0) HasNfs41Flags() bool`

HasNfs41Flags returns a boolean if a field has been set.

### GetAcltype

`func (o *FilesystemSetacl0) GetAcltype() string`

GetAcltype returns the Acltype field if non-nil, zero value otherwise.

### GetAcltypeOk

`func (o *FilesystemSetacl0) GetAcltypeOk() (*string, bool)`

GetAcltypeOk returns a tuple with the Acltype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcltype

`func (o *FilesystemSetacl0) SetAcltype(v string)`

SetAcltype sets Acltype field to given value.

### HasAcltype

`func (o *FilesystemSetacl0) HasAcltype() bool`

HasAcltype returns a boolean if a field has been set.

### GetOptions

`func (o *FilesystemSetacl0) GetOptions() FilesystemSetacl0Options`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *FilesystemSetacl0) GetOptionsOk() (*FilesystemSetacl0Options, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *FilesystemSetacl0) SetOptions(v FilesystemSetacl0Options)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *FilesystemSetacl0) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


