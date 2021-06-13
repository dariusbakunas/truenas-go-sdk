# ZfsSnapshotCreate0

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dataset** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NamingSchema** | Pointer to **string** |  | [optional] 
**Recursive** | Pointer to **bool** |  | [optional] 
**VmwareSync** | Pointer to **bool** |  | [optional] 
**Properties** | Pointer to **map[string]map[string]interface{}** |  | [optional] 

## Methods

### NewZfsSnapshotCreate0

`func NewZfsSnapshotCreate0() *ZfsSnapshotCreate0`

NewZfsSnapshotCreate0 instantiates a new ZfsSnapshotCreate0 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZfsSnapshotCreate0WithDefaults

`func NewZfsSnapshotCreate0WithDefaults() *ZfsSnapshotCreate0`

NewZfsSnapshotCreate0WithDefaults instantiates a new ZfsSnapshotCreate0 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataset

`func (o *ZfsSnapshotCreate0) GetDataset() string`

GetDataset returns the Dataset field if non-nil, zero value otherwise.

### GetDatasetOk

`func (o *ZfsSnapshotCreate0) GetDatasetOk() (*string, bool)`

GetDatasetOk returns a tuple with the Dataset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataset

`func (o *ZfsSnapshotCreate0) SetDataset(v string)`

SetDataset sets Dataset field to given value.

### HasDataset

`func (o *ZfsSnapshotCreate0) HasDataset() bool`

HasDataset returns a boolean if a field has been set.

### GetName

`func (o *ZfsSnapshotCreate0) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ZfsSnapshotCreate0) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ZfsSnapshotCreate0) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ZfsSnapshotCreate0) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamingSchema

`func (o *ZfsSnapshotCreate0) GetNamingSchema() string`

GetNamingSchema returns the NamingSchema field if non-nil, zero value otherwise.

### GetNamingSchemaOk

`func (o *ZfsSnapshotCreate0) GetNamingSchemaOk() (*string, bool)`

GetNamingSchemaOk returns a tuple with the NamingSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamingSchema

`func (o *ZfsSnapshotCreate0) SetNamingSchema(v string)`

SetNamingSchema sets NamingSchema field to given value.

### HasNamingSchema

`func (o *ZfsSnapshotCreate0) HasNamingSchema() bool`

HasNamingSchema returns a boolean if a field has been set.

### GetRecursive

`func (o *ZfsSnapshotCreate0) GetRecursive() bool`

GetRecursive returns the Recursive field if non-nil, zero value otherwise.

### GetRecursiveOk

`func (o *ZfsSnapshotCreate0) GetRecursiveOk() (*bool, bool)`

GetRecursiveOk returns a tuple with the Recursive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecursive

`func (o *ZfsSnapshotCreate0) SetRecursive(v bool)`

SetRecursive sets Recursive field to given value.

### HasRecursive

`func (o *ZfsSnapshotCreate0) HasRecursive() bool`

HasRecursive returns a boolean if a field has been set.

### GetVmwareSync

`func (o *ZfsSnapshotCreate0) GetVmwareSync() bool`

GetVmwareSync returns the VmwareSync field if non-nil, zero value otherwise.

### GetVmwareSyncOk

`func (o *ZfsSnapshotCreate0) GetVmwareSyncOk() (*bool, bool)`

GetVmwareSyncOk returns a tuple with the VmwareSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmwareSync

`func (o *ZfsSnapshotCreate0) SetVmwareSync(v bool)`

SetVmwareSync sets VmwareSync field to given value.

### HasVmwareSync

`func (o *ZfsSnapshotCreate0) HasVmwareSync() bool`

HasVmwareSync returns a boolean if a field has been set.

### GetProperties

`func (o *ZfsSnapshotCreate0) GetProperties() map[string]map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *ZfsSnapshotCreate0) GetPropertiesOk() (*map[string]map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *ZfsSnapshotCreate0) SetProperties(v map[string]map[string]interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *ZfsSnapshotCreate0) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


