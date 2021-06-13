# ReplicationCountEligibleManualSnapshots

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datasets** | Pointer to **[]string** |  | [optional] 
**NamingSchema** | Pointer to **[]string** |  | [optional] 
**Transport** | Pointer to [**ReplicationCountEligibleManualSnapshots2**](ReplicationCountEligibleManualSnapshots2.md) |  | [optional] 
**SshCredentials** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewReplicationCountEligibleManualSnapshots

`func NewReplicationCountEligibleManualSnapshots() *ReplicationCountEligibleManualSnapshots`

NewReplicationCountEligibleManualSnapshots instantiates a new ReplicationCountEligibleManualSnapshots object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplicationCountEligibleManualSnapshotsWithDefaults

`func NewReplicationCountEligibleManualSnapshotsWithDefaults() *ReplicationCountEligibleManualSnapshots`

NewReplicationCountEligibleManualSnapshotsWithDefaults instantiates a new ReplicationCountEligibleManualSnapshots object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatasets

`func (o *ReplicationCountEligibleManualSnapshots) GetDatasets() []string`

GetDatasets returns the Datasets field if non-nil, zero value otherwise.

### GetDatasetsOk

`func (o *ReplicationCountEligibleManualSnapshots) GetDatasetsOk() (*[]string, bool)`

GetDatasetsOk returns a tuple with the Datasets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasets

`func (o *ReplicationCountEligibleManualSnapshots) SetDatasets(v []string)`

SetDatasets sets Datasets field to given value.

### HasDatasets

`func (o *ReplicationCountEligibleManualSnapshots) HasDatasets() bool`

HasDatasets returns a boolean if a field has been set.

### GetNamingSchema

`func (o *ReplicationCountEligibleManualSnapshots) GetNamingSchema() []string`

GetNamingSchema returns the NamingSchema field if non-nil, zero value otherwise.

### GetNamingSchemaOk

`func (o *ReplicationCountEligibleManualSnapshots) GetNamingSchemaOk() (*[]string, bool)`

GetNamingSchemaOk returns a tuple with the NamingSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamingSchema

`func (o *ReplicationCountEligibleManualSnapshots) SetNamingSchema(v []string)`

SetNamingSchema sets NamingSchema field to given value.

### HasNamingSchema

`func (o *ReplicationCountEligibleManualSnapshots) HasNamingSchema() bool`

HasNamingSchema returns a boolean if a field has been set.

### GetTransport

`func (o *ReplicationCountEligibleManualSnapshots) GetTransport() ReplicationCountEligibleManualSnapshots2`

GetTransport returns the Transport field if non-nil, zero value otherwise.

### GetTransportOk

`func (o *ReplicationCountEligibleManualSnapshots) GetTransportOk() (*ReplicationCountEligibleManualSnapshots2, bool)`

GetTransportOk returns a tuple with the Transport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransport

`func (o *ReplicationCountEligibleManualSnapshots) SetTransport(v ReplicationCountEligibleManualSnapshots2)`

SetTransport sets Transport field to given value.

### HasTransport

`func (o *ReplicationCountEligibleManualSnapshots) HasTransport() bool`

HasTransport returns a boolean if a field has been set.

### GetSshCredentials

`func (o *ReplicationCountEligibleManualSnapshots) GetSshCredentials() int32`

GetSshCredentials returns the SshCredentials field if non-nil, zero value otherwise.

### GetSshCredentialsOk

`func (o *ReplicationCountEligibleManualSnapshots) GetSshCredentialsOk() (*int32, bool)`

GetSshCredentialsOk returns a tuple with the SshCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshCredentials

`func (o *ReplicationCountEligibleManualSnapshots) SetSshCredentials(v int32)`

SetSshCredentials sets SshCredentials field to given value.

### HasSshCredentials

`func (o *ReplicationCountEligibleManualSnapshots) HasSshCredentials() bool`

HasSshCredentials returns a boolean if a field has been set.

### SetSshCredentialsNil

`func (o *ReplicationCountEligibleManualSnapshots) SetSshCredentialsNil(b bool)`

 SetSshCredentialsNil sets the value for SshCredentials to be an explicit nil

### UnsetSshCredentials
`func (o *ReplicationCountEligibleManualSnapshots) UnsetSshCredentials()`

UnsetSshCredentials ensures that no value is present for SshCredentials, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


