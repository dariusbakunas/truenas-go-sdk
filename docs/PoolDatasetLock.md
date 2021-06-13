# PoolDatasetLock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**LockOptions** | Pointer to [**PoolDatasetLock1**](PoolDatasetLock1.md) |  | [optional] [default to {}]

## Methods

### NewPoolDatasetLock

`func NewPoolDatasetLock() *PoolDatasetLock`

NewPoolDatasetLock instantiates a new PoolDatasetLock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoolDatasetLockWithDefaults

`func NewPoolDatasetLockWithDefaults() *PoolDatasetLock`

NewPoolDatasetLockWithDefaults instantiates a new PoolDatasetLock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PoolDatasetLock) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PoolDatasetLock) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PoolDatasetLock) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PoolDatasetLock) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLockOptions

`func (o *PoolDatasetLock) GetLockOptions() PoolDatasetLock1`

GetLockOptions returns the LockOptions field if non-nil, zero value otherwise.

### GetLockOptionsOk

`func (o *PoolDatasetLock) GetLockOptionsOk() (*PoolDatasetLock1, bool)`

GetLockOptionsOk returns a tuple with the LockOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockOptions

`func (o *PoolDatasetLock) SetLockOptions(v PoolDatasetLock1)`

SetLockOptions sets LockOptions field to given value.

### HasLockOptions

`func (o *PoolDatasetLock) HasLockOptions() bool`

HasLockOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


