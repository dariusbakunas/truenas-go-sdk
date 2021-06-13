# PoolUpdate1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EncryptionOptions** | Pointer to [**PoolCreate0EncryptionOptions**](PoolCreate0EncryptionOptions.md) |  | [optional] 
**Topology** | Pointer to [**PoolCreate0Topology**](PoolCreate0Topology.md) |  | [optional] 
**Autotrim** | Pointer to **string** |  | [optional] 

## Methods

### NewPoolUpdate1

`func NewPoolUpdate1() *PoolUpdate1`

NewPoolUpdate1 instantiates a new PoolUpdate1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoolUpdate1WithDefaults

`func NewPoolUpdate1WithDefaults() *PoolUpdate1`

NewPoolUpdate1WithDefaults instantiates a new PoolUpdate1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEncryptionOptions

`func (o *PoolUpdate1) GetEncryptionOptions() PoolCreate0EncryptionOptions`

GetEncryptionOptions returns the EncryptionOptions field if non-nil, zero value otherwise.

### GetEncryptionOptionsOk

`func (o *PoolUpdate1) GetEncryptionOptionsOk() (*PoolCreate0EncryptionOptions, bool)`

GetEncryptionOptionsOk returns a tuple with the EncryptionOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionOptions

`func (o *PoolUpdate1) SetEncryptionOptions(v PoolCreate0EncryptionOptions)`

SetEncryptionOptions sets EncryptionOptions field to given value.

### HasEncryptionOptions

`func (o *PoolUpdate1) HasEncryptionOptions() bool`

HasEncryptionOptions returns a boolean if a field has been set.

### GetTopology

`func (o *PoolUpdate1) GetTopology() PoolCreate0Topology`

GetTopology returns the Topology field if non-nil, zero value otherwise.

### GetTopologyOk

`func (o *PoolUpdate1) GetTopologyOk() (*PoolCreate0Topology, bool)`

GetTopologyOk returns a tuple with the Topology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopology

`func (o *PoolUpdate1) SetTopology(v PoolCreate0Topology)`

SetTopology sets Topology field to given value.

### HasTopology

`func (o *PoolUpdate1) HasTopology() bool`

HasTopology returns a boolean if a field has been set.

### GetAutotrim

`func (o *PoolUpdate1) GetAutotrim() string`

GetAutotrim returns the Autotrim field if non-nil, zero value otherwise.

### GetAutotrimOk

`func (o *PoolUpdate1) GetAutotrimOk() (*string, bool)`

GetAutotrimOk returns a tuple with the Autotrim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutotrim

`func (o *PoolUpdate1) SetAutotrim(v string)`

SetAutotrim sets Autotrim field to given value.

### HasAutotrim

`func (o *PoolUpdate1) HasAutotrim() bool`

HasAutotrim returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


