# ZfsSnapshotRollback

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Options** | Pointer to [**ZfsSnapshotRollback1**](ZfsSnapshotRollback1.md) |  | [optional] [default to {}]

## Methods

### NewZfsSnapshotRollback

`func NewZfsSnapshotRollback() *ZfsSnapshotRollback`

NewZfsSnapshotRollback instantiates a new ZfsSnapshotRollback object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZfsSnapshotRollbackWithDefaults

`func NewZfsSnapshotRollbackWithDefaults() *ZfsSnapshotRollback`

NewZfsSnapshotRollbackWithDefaults instantiates a new ZfsSnapshotRollback object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ZfsSnapshotRollback) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ZfsSnapshotRollback) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ZfsSnapshotRollback) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ZfsSnapshotRollback) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOptions

`func (o *ZfsSnapshotRollback) GetOptions() ZfsSnapshotRollback1`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *ZfsSnapshotRollback) GetOptionsOk() (*ZfsSnapshotRollback1, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *ZfsSnapshotRollback) SetOptions(v ZfsSnapshotRollback1)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *ZfsSnapshotRollback) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


