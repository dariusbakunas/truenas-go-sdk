# PoolImportDisk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Device** | Pointer to **string** |  | [optional] 
**FsType** | Pointer to **string** |  | [optional] 
**FsOptions** | Pointer to **map[string]map[string]interface{}** |  | [optional] [default to {}]
**DstPath** | Pointer to **string** |  | [optional] 

## Methods

### NewPoolImportDisk

`func NewPoolImportDisk() *PoolImportDisk`

NewPoolImportDisk instantiates a new PoolImportDisk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoolImportDiskWithDefaults

`func NewPoolImportDiskWithDefaults() *PoolImportDisk`

NewPoolImportDiskWithDefaults instantiates a new PoolImportDisk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevice

`func (o *PoolImportDisk) GetDevice() string`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *PoolImportDisk) GetDeviceOk() (*string, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *PoolImportDisk) SetDevice(v string)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *PoolImportDisk) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetFsType

`func (o *PoolImportDisk) GetFsType() string`

GetFsType returns the FsType field if non-nil, zero value otherwise.

### GetFsTypeOk

`func (o *PoolImportDisk) GetFsTypeOk() (*string, bool)`

GetFsTypeOk returns a tuple with the FsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsType

`func (o *PoolImportDisk) SetFsType(v string)`

SetFsType sets FsType field to given value.

### HasFsType

`func (o *PoolImportDisk) HasFsType() bool`

HasFsType returns a boolean if a field has been set.

### GetFsOptions

`func (o *PoolImportDisk) GetFsOptions() map[string]map[string]interface{}`

GetFsOptions returns the FsOptions field if non-nil, zero value otherwise.

### GetFsOptionsOk

`func (o *PoolImportDisk) GetFsOptionsOk() (*map[string]map[string]interface{}, bool)`

GetFsOptionsOk returns a tuple with the FsOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsOptions

`func (o *PoolImportDisk) SetFsOptions(v map[string]map[string]interface{})`

SetFsOptions sets FsOptions field to given value.

### HasFsOptions

`func (o *PoolImportDisk) HasFsOptions() bool`

HasFsOptions returns a boolean if a field has been set.

### GetDstPath

`func (o *PoolImportDisk) GetDstPath() string`

GetDstPath returns the DstPath field if non-nil, zero value otherwise.

### GetDstPathOk

`func (o *PoolImportDisk) GetDstPathOk() (*string, bool)`

GetDstPathOk returns a tuple with the DstPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDstPath

`func (o *PoolImportDisk) SetDstPath(v string)`

SetDstPath sets DstPath field to given value.

### HasDstPath

`func (o *PoolImportDisk) HasDstPath() bool`

HasDstPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


