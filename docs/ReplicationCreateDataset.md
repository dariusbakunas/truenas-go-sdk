# ReplicationCreateDataset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dataset** | Pointer to **string** |  | [optional] 
**Transport** | Pointer to [**ReplicationCreateDataset1**](ReplicationCreateDataset1.md) |  | [optional] 
**SshCredentials** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewReplicationCreateDataset

`func NewReplicationCreateDataset() *ReplicationCreateDataset`

NewReplicationCreateDataset instantiates a new ReplicationCreateDataset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplicationCreateDatasetWithDefaults

`func NewReplicationCreateDatasetWithDefaults() *ReplicationCreateDataset`

NewReplicationCreateDatasetWithDefaults instantiates a new ReplicationCreateDataset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataset

`func (o *ReplicationCreateDataset) GetDataset() string`

GetDataset returns the Dataset field if non-nil, zero value otherwise.

### GetDatasetOk

`func (o *ReplicationCreateDataset) GetDatasetOk() (*string, bool)`

GetDatasetOk returns a tuple with the Dataset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataset

`func (o *ReplicationCreateDataset) SetDataset(v string)`

SetDataset sets Dataset field to given value.

### HasDataset

`func (o *ReplicationCreateDataset) HasDataset() bool`

HasDataset returns a boolean if a field has been set.

### GetTransport

`func (o *ReplicationCreateDataset) GetTransport() ReplicationCreateDataset1`

GetTransport returns the Transport field if non-nil, zero value otherwise.

### GetTransportOk

`func (o *ReplicationCreateDataset) GetTransportOk() (*ReplicationCreateDataset1, bool)`

GetTransportOk returns a tuple with the Transport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransport

`func (o *ReplicationCreateDataset) SetTransport(v ReplicationCreateDataset1)`

SetTransport sets Transport field to given value.

### HasTransport

`func (o *ReplicationCreateDataset) HasTransport() bool`

HasTransport returns a boolean if a field has been set.

### GetSshCredentials

`func (o *ReplicationCreateDataset) GetSshCredentials() int32`

GetSshCredentials returns the SshCredentials field if non-nil, zero value otherwise.

### GetSshCredentialsOk

`func (o *ReplicationCreateDataset) GetSshCredentialsOk() (*int32, bool)`

GetSshCredentialsOk returns a tuple with the SshCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshCredentials

`func (o *ReplicationCreateDataset) SetSshCredentials(v int32)`

SetSshCredentials sets SshCredentials field to given value.

### HasSshCredentials

`func (o *ReplicationCreateDataset) HasSshCredentials() bool`

HasSshCredentials returns a boolean if a field has been set.

### SetSshCredentialsNil

`func (o *ReplicationCreateDataset) SetSshCredentialsNil(b bool)`

 SetSshCredentialsNil sets the value for SshCredentials to be an explicit nil

### UnsetSshCredentials
`func (o *ReplicationCreateDataset) UnsetSshCredentials()`

UnsetSshCredentials ensures that no value is present for SshCredentials, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


