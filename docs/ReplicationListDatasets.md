# ReplicationListDatasets

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Transport** | Pointer to [**ReplicationListDatasets0**](ReplicationListDatasets0.md) |  | [optional] 
**SshCredentials** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewReplicationListDatasets

`func NewReplicationListDatasets() *ReplicationListDatasets`

NewReplicationListDatasets instantiates a new ReplicationListDatasets object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplicationListDatasetsWithDefaults

`func NewReplicationListDatasetsWithDefaults() *ReplicationListDatasets`

NewReplicationListDatasetsWithDefaults instantiates a new ReplicationListDatasets object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransport

`func (o *ReplicationListDatasets) GetTransport() ReplicationListDatasets0`

GetTransport returns the Transport field if non-nil, zero value otherwise.

### GetTransportOk

`func (o *ReplicationListDatasets) GetTransportOk() (*ReplicationListDatasets0, bool)`

GetTransportOk returns a tuple with the Transport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransport

`func (o *ReplicationListDatasets) SetTransport(v ReplicationListDatasets0)`

SetTransport sets Transport field to given value.

### HasTransport

`func (o *ReplicationListDatasets) HasTransport() bool`

HasTransport returns a boolean if a field has been set.

### GetSshCredentials

`func (o *ReplicationListDatasets) GetSshCredentials() int32`

GetSshCredentials returns the SshCredentials field if non-nil, zero value otherwise.

### GetSshCredentialsOk

`func (o *ReplicationListDatasets) GetSshCredentialsOk() (*int32, bool)`

GetSshCredentialsOk returns a tuple with the SshCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshCredentials

`func (o *ReplicationListDatasets) SetSshCredentials(v int32)`

SetSshCredentials sets SshCredentials field to given value.

### HasSshCredentials

`func (o *ReplicationListDatasets) HasSshCredentials() bool`

HasSshCredentials returns a boolean if a field has been set.

### SetSshCredentialsNil

`func (o *ReplicationListDatasets) SetSshCredentialsNil(b bool)`

 SetSshCredentialsNil sets the value for SshCredentials to be an explicit nil

### UnsetSshCredentials
`func (o *ReplicationListDatasets) UnsetSshCredentials()`

UnsetSshCredentials ensures that no value is present for SshCredentials, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


