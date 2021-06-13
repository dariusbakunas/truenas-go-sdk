# ReplicationTargetUnmatchedSnapshots

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Direction** | Pointer to [**ReplicationTargetUnmatchedSnapshots0**](ReplicationTargetUnmatchedSnapshots0.md) |  | [optional] 
**SourceDatasets** | Pointer to **[]string** |  | [optional] 
**TargetDataset** | Pointer to **string** |  | [optional] 
**Transport** | Pointer to [**ReplicationTargetUnmatchedSnapshots3**](ReplicationTargetUnmatchedSnapshots3.md) |  | [optional] 
**SshCredentials** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewReplicationTargetUnmatchedSnapshots

`func NewReplicationTargetUnmatchedSnapshots() *ReplicationTargetUnmatchedSnapshots`

NewReplicationTargetUnmatchedSnapshots instantiates a new ReplicationTargetUnmatchedSnapshots object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplicationTargetUnmatchedSnapshotsWithDefaults

`func NewReplicationTargetUnmatchedSnapshotsWithDefaults() *ReplicationTargetUnmatchedSnapshots`

NewReplicationTargetUnmatchedSnapshotsWithDefaults instantiates a new ReplicationTargetUnmatchedSnapshots object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirection

`func (o *ReplicationTargetUnmatchedSnapshots) GetDirection() ReplicationTargetUnmatchedSnapshots0`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *ReplicationTargetUnmatchedSnapshots) GetDirectionOk() (*ReplicationTargetUnmatchedSnapshots0, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *ReplicationTargetUnmatchedSnapshots) SetDirection(v ReplicationTargetUnmatchedSnapshots0)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *ReplicationTargetUnmatchedSnapshots) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetSourceDatasets

`func (o *ReplicationTargetUnmatchedSnapshots) GetSourceDatasets() []string`

GetSourceDatasets returns the SourceDatasets field if non-nil, zero value otherwise.

### GetSourceDatasetsOk

`func (o *ReplicationTargetUnmatchedSnapshots) GetSourceDatasetsOk() (*[]string, bool)`

GetSourceDatasetsOk returns a tuple with the SourceDatasets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDatasets

`func (o *ReplicationTargetUnmatchedSnapshots) SetSourceDatasets(v []string)`

SetSourceDatasets sets SourceDatasets field to given value.

### HasSourceDatasets

`func (o *ReplicationTargetUnmatchedSnapshots) HasSourceDatasets() bool`

HasSourceDatasets returns a boolean if a field has been set.

### GetTargetDataset

`func (o *ReplicationTargetUnmatchedSnapshots) GetTargetDataset() string`

GetTargetDataset returns the TargetDataset field if non-nil, zero value otherwise.

### GetTargetDatasetOk

`func (o *ReplicationTargetUnmatchedSnapshots) GetTargetDatasetOk() (*string, bool)`

GetTargetDatasetOk returns a tuple with the TargetDataset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetDataset

`func (o *ReplicationTargetUnmatchedSnapshots) SetTargetDataset(v string)`

SetTargetDataset sets TargetDataset field to given value.

### HasTargetDataset

`func (o *ReplicationTargetUnmatchedSnapshots) HasTargetDataset() bool`

HasTargetDataset returns a boolean if a field has been set.

### GetTransport

`func (o *ReplicationTargetUnmatchedSnapshots) GetTransport() ReplicationTargetUnmatchedSnapshots3`

GetTransport returns the Transport field if non-nil, zero value otherwise.

### GetTransportOk

`func (o *ReplicationTargetUnmatchedSnapshots) GetTransportOk() (*ReplicationTargetUnmatchedSnapshots3, bool)`

GetTransportOk returns a tuple with the Transport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransport

`func (o *ReplicationTargetUnmatchedSnapshots) SetTransport(v ReplicationTargetUnmatchedSnapshots3)`

SetTransport sets Transport field to given value.

### HasTransport

`func (o *ReplicationTargetUnmatchedSnapshots) HasTransport() bool`

HasTransport returns a boolean if a field has been set.

### GetSshCredentials

`func (o *ReplicationTargetUnmatchedSnapshots) GetSshCredentials() int32`

GetSshCredentials returns the SshCredentials field if non-nil, zero value otherwise.

### GetSshCredentialsOk

`func (o *ReplicationTargetUnmatchedSnapshots) GetSshCredentialsOk() (*int32, bool)`

GetSshCredentialsOk returns a tuple with the SshCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshCredentials

`func (o *ReplicationTargetUnmatchedSnapshots) SetSshCredentials(v int32)`

SetSshCredentials sets SshCredentials field to given value.

### HasSshCredentials

`func (o *ReplicationTargetUnmatchedSnapshots) HasSshCredentials() bool`

HasSshCredentials returns a boolean if a field has been set.

### SetSshCredentialsNil

`func (o *ReplicationTargetUnmatchedSnapshots) SetSshCredentialsNil(b bool)`

 SetSshCredentialsNil sets the value for SshCredentials to be an explicit nil

### UnsetSshCredentials
`func (o *ReplicationTargetUnmatchedSnapshots) UnsetSshCredentials()`

UnsetSshCredentials ensures that no value is present for SshCredentials, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


