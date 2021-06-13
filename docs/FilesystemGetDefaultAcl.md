# FilesystemGetDefaultAcl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AclType** | Pointer to [**FilesystemGetDefaultAcl0**](FilesystemGetDefaultAcl0.md) |  | [optional] [default to OPEN]
**ShareType** | Pointer to [**FilesystemGetDefaultAcl1**](FilesystemGetDefaultAcl1.md) |  | [optional] [default to NONE]

## Methods

### NewFilesystemGetDefaultAcl

`func NewFilesystemGetDefaultAcl() *FilesystemGetDefaultAcl`

NewFilesystemGetDefaultAcl instantiates a new FilesystemGetDefaultAcl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilesystemGetDefaultAclWithDefaults

`func NewFilesystemGetDefaultAclWithDefaults() *FilesystemGetDefaultAcl`

NewFilesystemGetDefaultAclWithDefaults instantiates a new FilesystemGetDefaultAcl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAclType

`func (o *FilesystemGetDefaultAcl) GetAclType() FilesystemGetDefaultAcl0`

GetAclType returns the AclType field if non-nil, zero value otherwise.

### GetAclTypeOk

`func (o *FilesystemGetDefaultAcl) GetAclTypeOk() (*FilesystemGetDefaultAcl0, bool)`

GetAclTypeOk returns a tuple with the AclType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAclType

`func (o *FilesystemGetDefaultAcl) SetAclType(v FilesystemGetDefaultAcl0)`

SetAclType sets AclType field to given value.

### HasAclType

`func (o *FilesystemGetDefaultAcl) HasAclType() bool`

HasAclType returns a boolean if a field has been set.

### GetShareType

`func (o *FilesystemGetDefaultAcl) GetShareType() FilesystemGetDefaultAcl1`

GetShareType returns the ShareType field if non-nil, zero value otherwise.

### GetShareTypeOk

`func (o *FilesystemGetDefaultAcl) GetShareTypeOk() (*FilesystemGetDefaultAcl1, bool)`

GetShareTypeOk returns a tuple with the ShareType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareType

`func (o *FilesystemGetDefaultAcl) SetShareType(v FilesystemGetDefaultAcl1)`

SetShareType sets ShareType field to given value.

### HasShareType

`func (o *FilesystemGetDefaultAcl) HasShareType() bool`

HasShareType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


