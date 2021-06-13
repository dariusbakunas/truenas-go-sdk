# ZfsSnapshotRemove0

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dataset** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**DeferDelete** | Pointer to **bool** |  | [optional] 

## Methods

### NewZfsSnapshotRemove0

`func NewZfsSnapshotRemove0() *ZfsSnapshotRemove0`

NewZfsSnapshotRemove0 instantiates a new ZfsSnapshotRemove0 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZfsSnapshotRemove0WithDefaults

`func NewZfsSnapshotRemove0WithDefaults() *ZfsSnapshotRemove0`

NewZfsSnapshotRemove0WithDefaults instantiates a new ZfsSnapshotRemove0 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataset

`func (o *ZfsSnapshotRemove0) GetDataset() string`

GetDataset returns the Dataset field if non-nil, zero value otherwise.

### GetDatasetOk

`func (o *ZfsSnapshotRemove0) GetDatasetOk() (*string, bool)`

GetDatasetOk returns a tuple with the Dataset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataset

`func (o *ZfsSnapshotRemove0) SetDataset(v string)`

SetDataset sets Dataset field to given value.

### HasDataset

`func (o *ZfsSnapshotRemove0) HasDataset() bool`

HasDataset returns a boolean if a field has been set.

### GetName

`func (o *ZfsSnapshotRemove0) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ZfsSnapshotRemove0) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ZfsSnapshotRemove0) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ZfsSnapshotRemove0) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDeferDelete

`func (o *ZfsSnapshotRemove0) GetDeferDelete() bool`

GetDeferDelete returns the DeferDelete field if non-nil, zero value otherwise.

### GetDeferDeleteOk

`func (o *ZfsSnapshotRemove0) GetDeferDeleteOk() (*bool, bool)`

GetDeferDeleteOk returns a tuple with the DeferDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeferDelete

`func (o *ZfsSnapshotRemove0) SetDeferDelete(v bool)`

SetDeferDelete sets DeferDelete field to given value.

### HasDeferDelete

`func (o *ZfsSnapshotRemove0) HasDeferDelete() bool`

HasDeferDelete returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


