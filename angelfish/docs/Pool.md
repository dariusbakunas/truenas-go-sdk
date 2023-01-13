# Pool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **string** |  | 
**Guid** | Pointer to **string** |  | [optional] 
**Path** | **string** |  | 
**Status** | Pointer to **string** |  | [optional] 
**Healthy** | Pointer to **bool** |  | [optional] 
**IsDecrypted** | Pointer to **bool** |  | [optional] 
**EncryptkeyPath** | Pointer to **string** |  | [optional] 

## Methods

### NewPool

`func NewPool(id int32, name string, path string, ) *Pool`

NewPool instantiates a new Pool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoolWithDefaults

`func NewPoolWithDefaults() *Pool`

NewPoolWithDefaults instantiates a new Pool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Pool) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Pool) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Pool) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *Pool) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Pool) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Pool) SetName(v string)`

SetName sets Name field to given value.


### GetGuid

`func (o *Pool) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *Pool) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *Pool) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *Pool) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetPath

`func (o *Pool) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *Pool) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *Pool) SetPath(v string)`

SetPath sets Path field to given value.


### GetStatus

`func (o *Pool) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Pool) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Pool) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Pool) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetHealthy

`func (o *Pool) GetHealthy() bool`

GetHealthy returns the Healthy field if non-nil, zero value otherwise.

### GetHealthyOk

`func (o *Pool) GetHealthyOk() (*bool, bool)`

GetHealthyOk returns a tuple with the Healthy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthy

`func (o *Pool) SetHealthy(v bool)`

SetHealthy sets Healthy field to given value.

### HasHealthy

`func (o *Pool) HasHealthy() bool`

HasHealthy returns a boolean if a field has been set.

### GetIsDecrypted

`func (o *Pool) GetIsDecrypted() bool`

GetIsDecrypted returns the IsDecrypted field if non-nil, zero value otherwise.

### GetIsDecryptedOk

`func (o *Pool) GetIsDecryptedOk() (*bool, bool)`

GetIsDecryptedOk returns a tuple with the IsDecrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDecrypted

`func (o *Pool) SetIsDecrypted(v bool)`

SetIsDecrypted sets IsDecrypted field to given value.

### HasIsDecrypted

`func (o *Pool) HasIsDecrypted() bool`

HasIsDecrypted returns a boolean if a field has been set.

### GetEncryptkeyPath

`func (o *Pool) GetEncryptkeyPath() string`

GetEncryptkeyPath returns the EncryptkeyPath field if non-nil, zero value otherwise.

### GetEncryptkeyPathOk

`func (o *Pool) GetEncryptkeyPathOk() (*string, bool)`

GetEncryptkeyPathOk returns a tuple with the EncryptkeyPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptkeyPath

`func (o *Pool) SetEncryptkeyPath(v string)`

SetEncryptkeyPath sets EncryptkeyPath field to given value.

### HasEncryptkeyPath

`func (o *Pool) HasEncryptkeyPath() bool`

HasEncryptkeyPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


