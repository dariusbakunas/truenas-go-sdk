# IscsiExtentUpdate1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Disk** | Pointer to **NullableString** |  | [optional] 
**Serial** | Pointer to **NullableString** |  | [optional] 
**Path** | Pointer to **NullableString** |  | [optional] 
**Filesize** | Pointer to **int32** |  | [optional] 
**Blocksize** | Pointer to **int32** |  | [optional] 
**Pblocksize** | Pointer to **bool** |  | [optional] 
**AvailThreshold** | Pointer to **NullableInt32** |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 
**InsecureTpc** | Pointer to **bool** |  | [optional] 
**Xen** | Pointer to **bool** |  | [optional] 
**Rpm** | Pointer to **string** |  | [optional] 
**Ro** | Pointer to **bool** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewIscsiExtentUpdate1

`func NewIscsiExtentUpdate1() *IscsiExtentUpdate1`

NewIscsiExtentUpdate1 instantiates a new IscsiExtentUpdate1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIscsiExtentUpdate1WithDefaults

`func NewIscsiExtentUpdate1WithDefaults() *IscsiExtentUpdate1`

NewIscsiExtentUpdate1WithDefaults instantiates a new IscsiExtentUpdate1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *IscsiExtentUpdate1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IscsiExtentUpdate1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IscsiExtentUpdate1) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IscsiExtentUpdate1) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *IscsiExtentUpdate1) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IscsiExtentUpdate1) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IscsiExtentUpdate1) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IscsiExtentUpdate1) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDisk

`func (o *IscsiExtentUpdate1) GetDisk() string`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *IscsiExtentUpdate1) GetDiskOk() (*string, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *IscsiExtentUpdate1) SetDisk(v string)`

SetDisk sets Disk field to given value.

### HasDisk

`func (o *IscsiExtentUpdate1) HasDisk() bool`

HasDisk returns a boolean if a field has been set.

### SetDiskNil

`func (o *IscsiExtentUpdate1) SetDiskNil(b bool)`

 SetDiskNil sets the value for Disk to be an explicit nil

### UnsetDisk
`func (o *IscsiExtentUpdate1) UnsetDisk()`

UnsetDisk ensures that no value is present for Disk, not even an explicit nil
### GetSerial

`func (o *IscsiExtentUpdate1) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *IscsiExtentUpdate1) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *IscsiExtentUpdate1) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *IscsiExtentUpdate1) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### SetSerialNil

`func (o *IscsiExtentUpdate1) SetSerialNil(b bool)`

 SetSerialNil sets the value for Serial to be an explicit nil

### UnsetSerial
`func (o *IscsiExtentUpdate1) UnsetSerial()`

UnsetSerial ensures that no value is present for Serial, not even an explicit nil
### GetPath

`func (o *IscsiExtentUpdate1) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *IscsiExtentUpdate1) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *IscsiExtentUpdate1) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *IscsiExtentUpdate1) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *IscsiExtentUpdate1) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *IscsiExtentUpdate1) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetFilesize

`func (o *IscsiExtentUpdate1) GetFilesize() int32`

GetFilesize returns the Filesize field if non-nil, zero value otherwise.

### GetFilesizeOk

`func (o *IscsiExtentUpdate1) GetFilesizeOk() (*int32, bool)`

GetFilesizeOk returns a tuple with the Filesize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesize

`func (o *IscsiExtentUpdate1) SetFilesize(v int32)`

SetFilesize sets Filesize field to given value.

### HasFilesize

`func (o *IscsiExtentUpdate1) HasFilesize() bool`

HasFilesize returns a boolean if a field has been set.

### GetBlocksize

`func (o *IscsiExtentUpdate1) GetBlocksize() int32`

GetBlocksize returns the Blocksize field if non-nil, zero value otherwise.

### GetBlocksizeOk

`func (o *IscsiExtentUpdate1) GetBlocksizeOk() (*int32, bool)`

GetBlocksizeOk returns a tuple with the Blocksize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocksize

`func (o *IscsiExtentUpdate1) SetBlocksize(v int32)`

SetBlocksize sets Blocksize field to given value.

### HasBlocksize

`func (o *IscsiExtentUpdate1) HasBlocksize() bool`

HasBlocksize returns a boolean if a field has been set.

### GetPblocksize

`func (o *IscsiExtentUpdate1) GetPblocksize() bool`

GetPblocksize returns the Pblocksize field if non-nil, zero value otherwise.

### GetPblocksizeOk

`func (o *IscsiExtentUpdate1) GetPblocksizeOk() (*bool, bool)`

GetPblocksizeOk returns a tuple with the Pblocksize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPblocksize

`func (o *IscsiExtentUpdate1) SetPblocksize(v bool)`

SetPblocksize sets Pblocksize field to given value.

### HasPblocksize

`func (o *IscsiExtentUpdate1) HasPblocksize() bool`

HasPblocksize returns a boolean if a field has been set.

### GetAvailThreshold

`func (o *IscsiExtentUpdate1) GetAvailThreshold() int32`

GetAvailThreshold returns the AvailThreshold field if non-nil, zero value otherwise.

### GetAvailThresholdOk

`func (o *IscsiExtentUpdate1) GetAvailThresholdOk() (*int32, bool)`

GetAvailThresholdOk returns a tuple with the AvailThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailThreshold

`func (o *IscsiExtentUpdate1) SetAvailThreshold(v int32)`

SetAvailThreshold sets AvailThreshold field to given value.

### HasAvailThreshold

`func (o *IscsiExtentUpdate1) HasAvailThreshold() bool`

HasAvailThreshold returns a boolean if a field has been set.

### SetAvailThresholdNil

`func (o *IscsiExtentUpdate1) SetAvailThresholdNil(b bool)`

 SetAvailThresholdNil sets the value for AvailThreshold to be an explicit nil

### UnsetAvailThreshold
`func (o *IscsiExtentUpdate1) UnsetAvailThreshold()`

UnsetAvailThreshold ensures that no value is present for AvailThreshold, not even an explicit nil
### GetComment

`func (o *IscsiExtentUpdate1) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *IscsiExtentUpdate1) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *IscsiExtentUpdate1) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *IscsiExtentUpdate1) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetInsecureTpc

`func (o *IscsiExtentUpdate1) GetInsecureTpc() bool`

GetInsecureTpc returns the InsecureTpc field if non-nil, zero value otherwise.

### GetInsecureTpcOk

`func (o *IscsiExtentUpdate1) GetInsecureTpcOk() (*bool, bool)`

GetInsecureTpcOk returns a tuple with the InsecureTpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsecureTpc

`func (o *IscsiExtentUpdate1) SetInsecureTpc(v bool)`

SetInsecureTpc sets InsecureTpc field to given value.

### HasInsecureTpc

`func (o *IscsiExtentUpdate1) HasInsecureTpc() bool`

HasInsecureTpc returns a boolean if a field has been set.

### GetXen

`func (o *IscsiExtentUpdate1) GetXen() bool`

GetXen returns the Xen field if non-nil, zero value otherwise.

### GetXenOk

`func (o *IscsiExtentUpdate1) GetXenOk() (*bool, bool)`

GetXenOk returns a tuple with the Xen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXen

`func (o *IscsiExtentUpdate1) SetXen(v bool)`

SetXen sets Xen field to given value.

### HasXen

`func (o *IscsiExtentUpdate1) HasXen() bool`

HasXen returns a boolean if a field has been set.

### GetRpm

`func (o *IscsiExtentUpdate1) GetRpm() string`

GetRpm returns the Rpm field if non-nil, zero value otherwise.

### GetRpmOk

`func (o *IscsiExtentUpdate1) GetRpmOk() (*string, bool)`

GetRpmOk returns a tuple with the Rpm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpm

`func (o *IscsiExtentUpdate1) SetRpm(v string)`

SetRpm sets Rpm field to given value.

### HasRpm

`func (o *IscsiExtentUpdate1) HasRpm() bool`

HasRpm returns a boolean if a field has been set.

### GetRo

`func (o *IscsiExtentUpdate1) GetRo() bool`

GetRo returns the Ro field if non-nil, zero value otherwise.

### GetRoOk

`func (o *IscsiExtentUpdate1) GetRoOk() (*bool, bool)`

GetRoOk returns a tuple with the Ro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRo

`func (o *IscsiExtentUpdate1) SetRo(v bool)`

SetRo sets Ro field to given value.

### HasRo

`func (o *IscsiExtentUpdate1) HasRo() bool`

HasRo returns a boolean if a field has been set.

### GetEnabled

`func (o *IscsiExtentUpdate1) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *IscsiExtentUpdate1) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *IscsiExtentUpdate1) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *IscsiExtentUpdate1) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


