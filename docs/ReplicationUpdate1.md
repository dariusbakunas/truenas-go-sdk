# ReplicationUpdate1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Direction** | Pointer to **string** |  | [optional] 
**Transport** | Pointer to **string** |  | [optional] 
**SshCredentials** | Pointer to **NullableInt32** |  | [optional] 
**NetcatActiveSide** | Pointer to **NullableString** |  | [optional] 
**NetcatActiveSideListenAddress** | Pointer to **NullableString** |  | [optional] 
**NetcatActiveSidePortMin** | Pointer to **NullableInt32** |  | [optional] 
**NetcatActiveSidePortMax** | Pointer to **NullableInt32** |  | [optional] 
**NetcatPassiveSideConnectAddress** | Pointer to **NullableString** |  | [optional] 
**SourceDatasets** | Pointer to **[]string** |  | [optional] 
**TargetDataset** | Pointer to **string** |  | [optional] 
**Recursive** | Pointer to **bool** |  | [optional] 
**Exclude** | Pointer to **[]string** |  | [optional] 
**Properties** | Pointer to **bool** |  | [optional] 
**PropertiesExclude** | Pointer to **[]string** |  | [optional] 
**Replicate** | Pointer to **bool** |  | [optional] 
**Encryption** | Pointer to **bool** |  | [optional] 
**EncryptionKey** | Pointer to **NullableString** |  | [optional] 
**EncryptionKeyFormat** | Pointer to **NullableString** |  | [optional] 
**EncryptionKeyLocation** | Pointer to **NullableString** |  | [optional] 
**PeriodicSnapshotTasks** | Pointer to **[]int32** |  | [optional] 
**NamingSchema** | Pointer to **[]string** |  | [optional] 
**AlsoIncludeNamingSchema** | Pointer to **[]string** |  | [optional] 
**Auto** | Pointer to **bool** |  | [optional] 
**Schedule** | Pointer to [**PoolSnapshottaskCreate0Schedule**](PoolSnapshottaskCreate0Schedule.md) |  | [optional] 
**RestrictSchedule** | Pointer to [**PoolSnapshottaskCreate0Schedule**](PoolSnapshottaskCreate0Schedule.md) |  | [optional] 
**OnlyMatchingSchedule** | Pointer to **bool** |  | [optional] 
**AllowFromScratch** | Pointer to **bool** |  | [optional] 
**Readonly** | Pointer to **string** |  | [optional] 
**HoldPendingSnapshots** | Pointer to **bool** |  | [optional] 
**RetentionPolicy** | Pointer to **string** |  | [optional] 
**LifetimeValue** | Pointer to **NullableInt32** |  | [optional] 
**LifetimeUnit** | Pointer to **NullableString** |  | [optional] 
**Compression** | Pointer to **NullableString** |  | [optional] 
**SpeedLimit** | Pointer to **NullableInt32** |  | [optional] 
**LargeBlock** | Pointer to **bool** |  | [optional] 
**Embed** | Pointer to **bool** |  | [optional] 
**Compressed** | Pointer to **bool** |  | [optional] 
**Retries** | Pointer to **int32** |  | [optional] 
**LoggingLevel** | Pointer to **NullableString** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewReplicationUpdate1

`func NewReplicationUpdate1() *ReplicationUpdate1`

NewReplicationUpdate1 instantiates a new ReplicationUpdate1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplicationUpdate1WithDefaults

`func NewReplicationUpdate1WithDefaults() *ReplicationUpdate1`

NewReplicationUpdate1WithDefaults instantiates a new ReplicationUpdate1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ReplicationUpdate1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ReplicationUpdate1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ReplicationUpdate1) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ReplicationUpdate1) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDirection

`func (o *ReplicationUpdate1) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *ReplicationUpdate1) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *ReplicationUpdate1) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *ReplicationUpdate1) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetTransport

`func (o *ReplicationUpdate1) GetTransport() string`

GetTransport returns the Transport field if non-nil, zero value otherwise.

### GetTransportOk

`func (o *ReplicationUpdate1) GetTransportOk() (*string, bool)`

GetTransportOk returns a tuple with the Transport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransport

`func (o *ReplicationUpdate1) SetTransport(v string)`

SetTransport sets Transport field to given value.

### HasTransport

`func (o *ReplicationUpdate1) HasTransport() bool`

HasTransport returns a boolean if a field has been set.

### GetSshCredentials

`func (o *ReplicationUpdate1) GetSshCredentials() int32`

GetSshCredentials returns the SshCredentials field if non-nil, zero value otherwise.

### GetSshCredentialsOk

`func (o *ReplicationUpdate1) GetSshCredentialsOk() (*int32, bool)`

GetSshCredentialsOk returns a tuple with the SshCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshCredentials

`func (o *ReplicationUpdate1) SetSshCredentials(v int32)`

SetSshCredentials sets SshCredentials field to given value.

### HasSshCredentials

`func (o *ReplicationUpdate1) HasSshCredentials() bool`

HasSshCredentials returns a boolean if a field has been set.

### SetSshCredentialsNil

`func (o *ReplicationUpdate1) SetSshCredentialsNil(b bool)`

 SetSshCredentialsNil sets the value for SshCredentials to be an explicit nil

### UnsetSshCredentials
`func (o *ReplicationUpdate1) UnsetSshCredentials()`

UnsetSshCredentials ensures that no value is present for SshCredentials, not even an explicit nil
### GetNetcatActiveSide

`func (o *ReplicationUpdate1) GetNetcatActiveSide() string`

GetNetcatActiveSide returns the NetcatActiveSide field if non-nil, zero value otherwise.

### GetNetcatActiveSideOk

`func (o *ReplicationUpdate1) GetNetcatActiveSideOk() (*string, bool)`

GetNetcatActiveSideOk returns a tuple with the NetcatActiveSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetcatActiveSide

`func (o *ReplicationUpdate1) SetNetcatActiveSide(v string)`

SetNetcatActiveSide sets NetcatActiveSide field to given value.

### HasNetcatActiveSide

`func (o *ReplicationUpdate1) HasNetcatActiveSide() bool`

HasNetcatActiveSide returns a boolean if a field has been set.

### SetNetcatActiveSideNil

`func (o *ReplicationUpdate1) SetNetcatActiveSideNil(b bool)`

 SetNetcatActiveSideNil sets the value for NetcatActiveSide to be an explicit nil

### UnsetNetcatActiveSide
`func (o *ReplicationUpdate1) UnsetNetcatActiveSide()`

UnsetNetcatActiveSide ensures that no value is present for NetcatActiveSide, not even an explicit nil
### GetNetcatActiveSideListenAddress

`func (o *ReplicationUpdate1) GetNetcatActiveSideListenAddress() string`

GetNetcatActiveSideListenAddress returns the NetcatActiveSideListenAddress field if non-nil, zero value otherwise.

### GetNetcatActiveSideListenAddressOk

`func (o *ReplicationUpdate1) GetNetcatActiveSideListenAddressOk() (*string, bool)`

GetNetcatActiveSideListenAddressOk returns a tuple with the NetcatActiveSideListenAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetcatActiveSideListenAddress

`func (o *ReplicationUpdate1) SetNetcatActiveSideListenAddress(v string)`

SetNetcatActiveSideListenAddress sets NetcatActiveSideListenAddress field to given value.

### HasNetcatActiveSideListenAddress

`func (o *ReplicationUpdate1) HasNetcatActiveSideListenAddress() bool`

HasNetcatActiveSideListenAddress returns a boolean if a field has been set.

### SetNetcatActiveSideListenAddressNil

`func (o *ReplicationUpdate1) SetNetcatActiveSideListenAddressNil(b bool)`

 SetNetcatActiveSideListenAddressNil sets the value for NetcatActiveSideListenAddress to be an explicit nil

### UnsetNetcatActiveSideListenAddress
`func (o *ReplicationUpdate1) UnsetNetcatActiveSideListenAddress()`

UnsetNetcatActiveSideListenAddress ensures that no value is present for NetcatActiveSideListenAddress, not even an explicit nil
### GetNetcatActiveSidePortMin

`func (o *ReplicationUpdate1) GetNetcatActiveSidePortMin() int32`

GetNetcatActiveSidePortMin returns the NetcatActiveSidePortMin field if non-nil, zero value otherwise.

### GetNetcatActiveSidePortMinOk

`func (o *ReplicationUpdate1) GetNetcatActiveSidePortMinOk() (*int32, bool)`

GetNetcatActiveSidePortMinOk returns a tuple with the NetcatActiveSidePortMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetcatActiveSidePortMin

`func (o *ReplicationUpdate1) SetNetcatActiveSidePortMin(v int32)`

SetNetcatActiveSidePortMin sets NetcatActiveSidePortMin field to given value.

### HasNetcatActiveSidePortMin

`func (o *ReplicationUpdate1) HasNetcatActiveSidePortMin() bool`

HasNetcatActiveSidePortMin returns a boolean if a field has been set.

### SetNetcatActiveSidePortMinNil

`func (o *ReplicationUpdate1) SetNetcatActiveSidePortMinNil(b bool)`

 SetNetcatActiveSidePortMinNil sets the value for NetcatActiveSidePortMin to be an explicit nil

### UnsetNetcatActiveSidePortMin
`func (o *ReplicationUpdate1) UnsetNetcatActiveSidePortMin()`

UnsetNetcatActiveSidePortMin ensures that no value is present for NetcatActiveSidePortMin, not even an explicit nil
### GetNetcatActiveSidePortMax

`func (o *ReplicationUpdate1) GetNetcatActiveSidePortMax() int32`

GetNetcatActiveSidePortMax returns the NetcatActiveSidePortMax field if non-nil, zero value otherwise.

### GetNetcatActiveSidePortMaxOk

`func (o *ReplicationUpdate1) GetNetcatActiveSidePortMaxOk() (*int32, bool)`

GetNetcatActiveSidePortMaxOk returns a tuple with the NetcatActiveSidePortMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetcatActiveSidePortMax

`func (o *ReplicationUpdate1) SetNetcatActiveSidePortMax(v int32)`

SetNetcatActiveSidePortMax sets NetcatActiveSidePortMax field to given value.

### HasNetcatActiveSidePortMax

`func (o *ReplicationUpdate1) HasNetcatActiveSidePortMax() bool`

HasNetcatActiveSidePortMax returns a boolean if a field has been set.

### SetNetcatActiveSidePortMaxNil

`func (o *ReplicationUpdate1) SetNetcatActiveSidePortMaxNil(b bool)`

 SetNetcatActiveSidePortMaxNil sets the value for NetcatActiveSidePortMax to be an explicit nil

### UnsetNetcatActiveSidePortMax
`func (o *ReplicationUpdate1) UnsetNetcatActiveSidePortMax()`

UnsetNetcatActiveSidePortMax ensures that no value is present for NetcatActiveSidePortMax, not even an explicit nil
### GetNetcatPassiveSideConnectAddress

`func (o *ReplicationUpdate1) GetNetcatPassiveSideConnectAddress() string`

GetNetcatPassiveSideConnectAddress returns the NetcatPassiveSideConnectAddress field if non-nil, zero value otherwise.

### GetNetcatPassiveSideConnectAddressOk

`func (o *ReplicationUpdate1) GetNetcatPassiveSideConnectAddressOk() (*string, bool)`

GetNetcatPassiveSideConnectAddressOk returns a tuple with the NetcatPassiveSideConnectAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetcatPassiveSideConnectAddress

`func (o *ReplicationUpdate1) SetNetcatPassiveSideConnectAddress(v string)`

SetNetcatPassiveSideConnectAddress sets NetcatPassiveSideConnectAddress field to given value.

### HasNetcatPassiveSideConnectAddress

`func (o *ReplicationUpdate1) HasNetcatPassiveSideConnectAddress() bool`

HasNetcatPassiveSideConnectAddress returns a boolean if a field has been set.

### SetNetcatPassiveSideConnectAddressNil

`func (o *ReplicationUpdate1) SetNetcatPassiveSideConnectAddressNil(b bool)`

 SetNetcatPassiveSideConnectAddressNil sets the value for NetcatPassiveSideConnectAddress to be an explicit nil

### UnsetNetcatPassiveSideConnectAddress
`func (o *ReplicationUpdate1) UnsetNetcatPassiveSideConnectAddress()`

UnsetNetcatPassiveSideConnectAddress ensures that no value is present for NetcatPassiveSideConnectAddress, not even an explicit nil
### GetSourceDatasets

`func (o *ReplicationUpdate1) GetSourceDatasets() []string`

GetSourceDatasets returns the SourceDatasets field if non-nil, zero value otherwise.

### GetSourceDatasetsOk

`func (o *ReplicationUpdate1) GetSourceDatasetsOk() (*[]string, bool)`

GetSourceDatasetsOk returns a tuple with the SourceDatasets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDatasets

`func (o *ReplicationUpdate1) SetSourceDatasets(v []string)`

SetSourceDatasets sets SourceDatasets field to given value.

### HasSourceDatasets

`func (o *ReplicationUpdate1) HasSourceDatasets() bool`

HasSourceDatasets returns a boolean if a field has been set.

### GetTargetDataset

`func (o *ReplicationUpdate1) GetTargetDataset() string`

GetTargetDataset returns the TargetDataset field if non-nil, zero value otherwise.

### GetTargetDatasetOk

`func (o *ReplicationUpdate1) GetTargetDatasetOk() (*string, bool)`

GetTargetDatasetOk returns a tuple with the TargetDataset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetDataset

`func (o *ReplicationUpdate1) SetTargetDataset(v string)`

SetTargetDataset sets TargetDataset field to given value.

### HasTargetDataset

`func (o *ReplicationUpdate1) HasTargetDataset() bool`

HasTargetDataset returns a boolean if a field has been set.

### GetRecursive

`func (o *ReplicationUpdate1) GetRecursive() bool`

GetRecursive returns the Recursive field if non-nil, zero value otherwise.

### GetRecursiveOk

`func (o *ReplicationUpdate1) GetRecursiveOk() (*bool, bool)`

GetRecursiveOk returns a tuple with the Recursive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecursive

`func (o *ReplicationUpdate1) SetRecursive(v bool)`

SetRecursive sets Recursive field to given value.

### HasRecursive

`func (o *ReplicationUpdate1) HasRecursive() bool`

HasRecursive returns a boolean if a field has been set.

### GetExclude

`func (o *ReplicationUpdate1) GetExclude() []string`

GetExclude returns the Exclude field if non-nil, zero value otherwise.

### GetExcludeOk

`func (o *ReplicationUpdate1) GetExcludeOk() (*[]string, bool)`

GetExcludeOk returns a tuple with the Exclude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclude

`func (o *ReplicationUpdate1) SetExclude(v []string)`

SetExclude sets Exclude field to given value.

### HasExclude

`func (o *ReplicationUpdate1) HasExclude() bool`

HasExclude returns a boolean if a field has been set.

### GetProperties

`func (o *ReplicationUpdate1) GetProperties() bool`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *ReplicationUpdate1) GetPropertiesOk() (*bool, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *ReplicationUpdate1) SetProperties(v bool)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *ReplicationUpdate1) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetPropertiesExclude

`func (o *ReplicationUpdate1) GetPropertiesExclude() []string`

GetPropertiesExclude returns the PropertiesExclude field if non-nil, zero value otherwise.

### GetPropertiesExcludeOk

`func (o *ReplicationUpdate1) GetPropertiesExcludeOk() (*[]string, bool)`

GetPropertiesExcludeOk returns a tuple with the PropertiesExclude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertiesExclude

`func (o *ReplicationUpdate1) SetPropertiesExclude(v []string)`

SetPropertiesExclude sets PropertiesExclude field to given value.

### HasPropertiesExclude

`func (o *ReplicationUpdate1) HasPropertiesExclude() bool`

HasPropertiesExclude returns a boolean if a field has been set.

### GetReplicate

`func (o *ReplicationUpdate1) GetReplicate() bool`

GetReplicate returns the Replicate field if non-nil, zero value otherwise.

### GetReplicateOk

`func (o *ReplicationUpdate1) GetReplicateOk() (*bool, bool)`

GetReplicateOk returns a tuple with the Replicate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicate

`func (o *ReplicationUpdate1) SetReplicate(v bool)`

SetReplicate sets Replicate field to given value.

### HasReplicate

`func (o *ReplicationUpdate1) HasReplicate() bool`

HasReplicate returns a boolean if a field has been set.

### GetEncryption

`func (o *ReplicationUpdate1) GetEncryption() bool`

GetEncryption returns the Encryption field if non-nil, zero value otherwise.

### GetEncryptionOk

`func (o *ReplicationUpdate1) GetEncryptionOk() (*bool, bool)`

GetEncryptionOk returns a tuple with the Encryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryption

`func (o *ReplicationUpdate1) SetEncryption(v bool)`

SetEncryption sets Encryption field to given value.

### HasEncryption

`func (o *ReplicationUpdate1) HasEncryption() bool`

HasEncryption returns a boolean if a field has been set.

### GetEncryptionKey

`func (o *ReplicationUpdate1) GetEncryptionKey() string`

GetEncryptionKey returns the EncryptionKey field if non-nil, zero value otherwise.

### GetEncryptionKeyOk

`func (o *ReplicationUpdate1) GetEncryptionKeyOk() (*string, bool)`

GetEncryptionKeyOk returns a tuple with the EncryptionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionKey

`func (o *ReplicationUpdate1) SetEncryptionKey(v string)`

SetEncryptionKey sets EncryptionKey field to given value.

### HasEncryptionKey

`func (o *ReplicationUpdate1) HasEncryptionKey() bool`

HasEncryptionKey returns a boolean if a field has been set.

### SetEncryptionKeyNil

`func (o *ReplicationUpdate1) SetEncryptionKeyNil(b bool)`

 SetEncryptionKeyNil sets the value for EncryptionKey to be an explicit nil

### UnsetEncryptionKey
`func (o *ReplicationUpdate1) UnsetEncryptionKey()`

UnsetEncryptionKey ensures that no value is present for EncryptionKey, not even an explicit nil
### GetEncryptionKeyFormat

`func (o *ReplicationUpdate1) GetEncryptionKeyFormat() string`

GetEncryptionKeyFormat returns the EncryptionKeyFormat field if non-nil, zero value otherwise.

### GetEncryptionKeyFormatOk

`func (o *ReplicationUpdate1) GetEncryptionKeyFormatOk() (*string, bool)`

GetEncryptionKeyFormatOk returns a tuple with the EncryptionKeyFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionKeyFormat

`func (o *ReplicationUpdate1) SetEncryptionKeyFormat(v string)`

SetEncryptionKeyFormat sets EncryptionKeyFormat field to given value.

### HasEncryptionKeyFormat

`func (o *ReplicationUpdate1) HasEncryptionKeyFormat() bool`

HasEncryptionKeyFormat returns a boolean if a field has been set.

### SetEncryptionKeyFormatNil

`func (o *ReplicationUpdate1) SetEncryptionKeyFormatNil(b bool)`

 SetEncryptionKeyFormatNil sets the value for EncryptionKeyFormat to be an explicit nil

### UnsetEncryptionKeyFormat
`func (o *ReplicationUpdate1) UnsetEncryptionKeyFormat()`

UnsetEncryptionKeyFormat ensures that no value is present for EncryptionKeyFormat, not even an explicit nil
### GetEncryptionKeyLocation

`func (o *ReplicationUpdate1) GetEncryptionKeyLocation() string`

GetEncryptionKeyLocation returns the EncryptionKeyLocation field if non-nil, zero value otherwise.

### GetEncryptionKeyLocationOk

`func (o *ReplicationUpdate1) GetEncryptionKeyLocationOk() (*string, bool)`

GetEncryptionKeyLocationOk returns a tuple with the EncryptionKeyLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionKeyLocation

`func (o *ReplicationUpdate1) SetEncryptionKeyLocation(v string)`

SetEncryptionKeyLocation sets EncryptionKeyLocation field to given value.

### HasEncryptionKeyLocation

`func (o *ReplicationUpdate1) HasEncryptionKeyLocation() bool`

HasEncryptionKeyLocation returns a boolean if a field has been set.

### SetEncryptionKeyLocationNil

`func (o *ReplicationUpdate1) SetEncryptionKeyLocationNil(b bool)`

 SetEncryptionKeyLocationNil sets the value for EncryptionKeyLocation to be an explicit nil

### UnsetEncryptionKeyLocation
`func (o *ReplicationUpdate1) UnsetEncryptionKeyLocation()`

UnsetEncryptionKeyLocation ensures that no value is present for EncryptionKeyLocation, not even an explicit nil
### GetPeriodicSnapshotTasks

`func (o *ReplicationUpdate1) GetPeriodicSnapshotTasks() []int32`

GetPeriodicSnapshotTasks returns the PeriodicSnapshotTasks field if non-nil, zero value otherwise.

### GetPeriodicSnapshotTasksOk

`func (o *ReplicationUpdate1) GetPeriodicSnapshotTasksOk() (*[]int32, bool)`

GetPeriodicSnapshotTasksOk returns a tuple with the PeriodicSnapshotTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodicSnapshotTasks

`func (o *ReplicationUpdate1) SetPeriodicSnapshotTasks(v []int32)`

SetPeriodicSnapshotTasks sets PeriodicSnapshotTasks field to given value.

### HasPeriodicSnapshotTasks

`func (o *ReplicationUpdate1) HasPeriodicSnapshotTasks() bool`

HasPeriodicSnapshotTasks returns a boolean if a field has been set.

### GetNamingSchema

`func (o *ReplicationUpdate1) GetNamingSchema() []string`

GetNamingSchema returns the NamingSchema field if non-nil, zero value otherwise.

### GetNamingSchemaOk

`func (o *ReplicationUpdate1) GetNamingSchemaOk() (*[]string, bool)`

GetNamingSchemaOk returns a tuple with the NamingSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamingSchema

`func (o *ReplicationUpdate1) SetNamingSchema(v []string)`

SetNamingSchema sets NamingSchema field to given value.

### HasNamingSchema

`func (o *ReplicationUpdate1) HasNamingSchema() bool`

HasNamingSchema returns a boolean if a field has been set.

### GetAlsoIncludeNamingSchema

`func (o *ReplicationUpdate1) GetAlsoIncludeNamingSchema() []string`

GetAlsoIncludeNamingSchema returns the AlsoIncludeNamingSchema field if non-nil, zero value otherwise.

### GetAlsoIncludeNamingSchemaOk

`func (o *ReplicationUpdate1) GetAlsoIncludeNamingSchemaOk() (*[]string, bool)`

GetAlsoIncludeNamingSchemaOk returns a tuple with the AlsoIncludeNamingSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlsoIncludeNamingSchema

`func (o *ReplicationUpdate1) SetAlsoIncludeNamingSchema(v []string)`

SetAlsoIncludeNamingSchema sets AlsoIncludeNamingSchema field to given value.

### HasAlsoIncludeNamingSchema

`func (o *ReplicationUpdate1) HasAlsoIncludeNamingSchema() bool`

HasAlsoIncludeNamingSchema returns a boolean if a field has been set.

### GetAuto

`func (o *ReplicationUpdate1) GetAuto() bool`

GetAuto returns the Auto field if non-nil, zero value otherwise.

### GetAutoOk

`func (o *ReplicationUpdate1) GetAutoOk() (*bool, bool)`

GetAutoOk returns a tuple with the Auto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuto

`func (o *ReplicationUpdate1) SetAuto(v bool)`

SetAuto sets Auto field to given value.

### HasAuto

`func (o *ReplicationUpdate1) HasAuto() bool`

HasAuto returns a boolean if a field has been set.

### GetSchedule

`func (o *ReplicationUpdate1) GetSchedule() PoolSnapshottaskCreate0Schedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *ReplicationUpdate1) GetScheduleOk() (*PoolSnapshottaskCreate0Schedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *ReplicationUpdate1) SetSchedule(v PoolSnapshottaskCreate0Schedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *ReplicationUpdate1) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetRestrictSchedule

`func (o *ReplicationUpdate1) GetRestrictSchedule() PoolSnapshottaskCreate0Schedule`

GetRestrictSchedule returns the RestrictSchedule field if non-nil, zero value otherwise.

### GetRestrictScheduleOk

`func (o *ReplicationUpdate1) GetRestrictScheduleOk() (*PoolSnapshottaskCreate0Schedule, bool)`

GetRestrictScheduleOk returns a tuple with the RestrictSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictSchedule

`func (o *ReplicationUpdate1) SetRestrictSchedule(v PoolSnapshottaskCreate0Schedule)`

SetRestrictSchedule sets RestrictSchedule field to given value.

### HasRestrictSchedule

`func (o *ReplicationUpdate1) HasRestrictSchedule() bool`

HasRestrictSchedule returns a boolean if a field has been set.

### GetOnlyMatchingSchedule

`func (o *ReplicationUpdate1) GetOnlyMatchingSchedule() bool`

GetOnlyMatchingSchedule returns the OnlyMatchingSchedule field if non-nil, zero value otherwise.

### GetOnlyMatchingScheduleOk

`func (o *ReplicationUpdate1) GetOnlyMatchingScheduleOk() (*bool, bool)`

GetOnlyMatchingScheduleOk returns a tuple with the OnlyMatchingSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlyMatchingSchedule

`func (o *ReplicationUpdate1) SetOnlyMatchingSchedule(v bool)`

SetOnlyMatchingSchedule sets OnlyMatchingSchedule field to given value.

### HasOnlyMatchingSchedule

`func (o *ReplicationUpdate1) HasOnlyMatchingSchedule() bool`

HasOnlyMatchingSchedule returns a boolean if a field has been set.

### GetAllowFromScratch

`func (o *ReplicationUpdate1) GetAllowFromScratch() bool`

GetAllowFromScratch returns the AllowFromScratch field if non-nil, zero value otherwise.

### GetAllowFromScratchOk

`func (o *ReplicationUpdate1) GetAllowFromScratchOk() (*bool, bool)`

GetAllowFromScratchOk returns a tuple with the AllowFromScratch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowFromScratch

`func (o *ReplicationUpdate1) SetAllowFromScratch(v bool)`

SetAllowFromScratch sets AllowFromScratch field to given value.

### HasAllowFromScratch

`func (o *ReplicationUpdate1) HasAllowFromScratch() bool`

HasAllowFromScratch returns a boolean if a field has been set.

### GetReadonly

`func (o *ReplicationUpdate1) GetReadonly() string`

GetReadonly returns the Readonly field if non-nil, zero value otherwise.

### GetReadonlyOk

`func (o *ReplicationUpdate1) GetReadonlyOk() (*string, bool)`

GetReadonlyOk returns a tuple with the Readonly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadonly

`func (o *ReplicationUpdate1) SetReadonly(v string)`

SetReadonly sets Readonly field to given value.

### HasReadonly

`func (o *ReplicationUpdate1) HasReadonly() bool`

HasReadonly returns a boolean if a field has been set.

### GetHoldPendingSnapshots

`func (o *ReplicationUpdate1) GetHoldPendingSnapshots() bool`

GetHoldPendingSnapshots returns the HoldPendingSnapshots field if non-nil, zero value otherwise.

### GetHoldPendingSnapshotsOk

`func (o *ReplicationUpdate1) GetHoldPendingSnapshotsOk() (*bool, bool)`

GetHoldPendingSnapshotsOk returns a tuple with the HoldPendingSnapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldPendingSnapshots

`func (o *ReplicationUpdate1) SetHoldPendingSnapshots(v bool)`

SetHoldPendingSnapshots sets HoldPendingSnapshots field to given value.

### HasHoldPendingSnapshots

`func (o *ReplicationUpdate1) HasHoldPendingSnapshots() bool`

HasHoldPendingSnapshots returns a boolean if a field has been set.

### GetRetentionPolicy

`func (o *ReplicationUpdate1) GetRetentionPolicy() string`

GetRetentionPolicy returns the RetentionPolicy field if non-nil, zero value otherwise.

### GetRetentionPolicyOk

`func (o *ReplicationUpdate1) GetRetentionPolicyOk() (*string, bool)`

GetRetentionPolicyOk returns a tuple with the RetentionPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionPolicy

`func (o *ReplicationUpdate1) SetRetentionPolicy(v string)`

SetRetentionPolicy sets RetentionPolicy field to given value.

### HasRetentionPolicy

`func (o *ReplicationUpdate1) HasRetentionPolicy() bool`

HasRetentionPolicy returns a boolean if a field has been set.

### GetLifetimeValue

`func (o *ReplicationUpdate1) GetLifetimeValue() int32`

GetLifetimeValue returns the LifetimeValue field if non-nil, zero value otherwise.

### GetLifetimeValueOk

`func (o *ReplicationUpdate1) GetLifetimeValueOk() (*int32, bool)`

GetLifetimeValueOk returns a tuple with the LifetimeValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifetimeValue

`func (o *ReplicationUpdate1) SetLifetimeValue(v int32)`

SetLifetimeValue sets LifetimeValue field to given value.

### HasLifetimeValue

`func (o *ReplicationUpdate1) HasLifetimeValue() bool`

HasLifetimeValue returns a boolean if a field has been set.

### SetLifetimeValueNil

`func (o *ReplicationUpdate1) SetLifetimeValueNil(b bool)`

 SetLifetimeValueNil sets the value for LifetimeValue to be an explicit nil

### UnsetLifetimeValue
`func (o *ReplicationUpdate1) UnsetLifetimeValue()`

UnsetLifetimeValue ensures that no value is present for LifetimeValue, not even an explicit nil
### GetLifetimeUnit

`func (o *ReplicationUpdate1) GetLifetimeUnit() string`

GetLifetimeUnit returns the LifetimeUnit field if non-nil, zero value otherwise.

### GetLifetimeUnitOk

`func (o *ReplicationUpdate1) GetLifetimeUnitOk() (*string, bool)`

GetLifetimeUnitOk returns a tuple with the LifetimeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifetimeUnit

`func (o *ReplicationUpdate1) SetLifetimeUnit(v string)`

SetLifetimeUnit sets LifetimeUnit field to given value.

### HasLifetimeUnit

`func (o *ReplicationUpdate1) HasLifetimeUnit() bool`

HasLifetimeUnit returns a boolean if a field has been set.

### SetLifetimeUnitNil

`func (o *ReplicationUpdate1) SetLifetimeUnitNil(b bool)`

 SetLifetimeUnitNil sets the value for LifetimeUnit to be an explicit nil

### UnsetLifetimeUnit
`func (o *ReplicationUpdate1) UnsetLifetimeUnit()`

UnsetLifetimeUnit ensures that no value is present for LifetimeUnit, not even an explicit nil
### GetCompression

`func (o *ReplicationUpdate1) GetCompression() string`

GetCompression returns the Compression field if non-nil, zero value otherwise.

### GetCompressionOk

`func (o *ReplicationUpdate1) GetCompressionOk() (*string, bool)`

GetCompressionOk returns a tuple with the Compression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompression

`func (o *ReplicationUpdate1) SetCompression(v string)`

SetCompression sets Compression field to given value.

### HasCompression

`func (o *ReplicationUpdate1) HasCompression() bool`

HasCompression returns a boolean if a field has been set.

### SetCompressionNil

`func (o *ReplicationUpdate1) SetCompressionNil(b bool)`

 SetCompressionNil sets the value for Compression to be an explicit nil

### UnsetCompression
`func (o *ReplicationUpdate1) UnsetCompression()`

UnsetCompression ensures that no value is present for Compression, not even an explicit nil
### GetSpeedLimit

`func (o *ReplicationUpdate1) GetSpeedLimit() int32`

GetSpeedLimit returns the SpeedLimit field if non-nil, zero value otherwise.

### GetSpeedLimitOk

`func (o *ReplicationUpdate1) GetSpeedLimitOk() (*int32, bool)`

GetSpeedLimitOk returns a tuple with the SpeedLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeedLimit

`func (o *ReplicationUpdate1) SetSpeedLimit(v int32)`

SetSpeedLimit sets SpeedLimit field to given value.

### HasSpeedLimit

`func (o *ReplicationUpdate1) HasSpeedLimit() bool`

HasSpeedLimit returns a boolean if a field has been set.

### SetSpeedLimitNil

`func (o *ReplicationUpdate1) SetSpeedLimitNil(b bool)`

 SetSpeedLimitNil sets the value for SpeedLimit to be an explicit nil

### UnsetSpeedLimit
`func (o *ReplicationUpdate1) UnsetSpeedLimit()`

UnsetSpeedLimit ensures that no value is present for SpeedLimit, not even an explicit nil
### GetLargeBlock

`func (o *ReplicationUpdate1) GetLargeBlock() bool`

GetLargeBlock returns the LargeBlock field if non-nil, zero value otherwise.

### GetLargeBlockOk

`func (o *ReplicationUpdate1) GetLargeBlockOk() (*bool, bool)`

GetLargeBlockOk returns a tuple with the LargeBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLargeBlock

`func (o *ReplicationUpdate1) SetLargeBlock(v bool)`

SetLargeBlock sets LargeBlock field to given value.

### HasLargeBlock

`func (o *ReplicationUpdate1) HasLargeBlock() bool`

HasLargeBlock returns a boolean if a field has been set.

### GetEmbed

`func (o *ReplicationUpdate1) GetEmbed() bool`

GetEmbed returns the Embed field if non-nil, zero value otherwise.

### GetEmbedOk

`func (o *ReplicationUpdate1) GetEmbedOk() (*bool, bool)`

GetEmbedOk returns a tuple with the Embed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbed

`func (o *ReplicationUpdate1) SetEmbed(v bool)`

SetEmbed sets Embed field to given value.

### HasEmbed

`func (o *ReplicationUpdate1) HasEmbed() bool`

HasEmbed returns a boolean if a field has been set.

### GetCompressed

`func (o *ReplicationUpdate1) GetCompressed() bool`

GetCompressed returns the Compressed field if non-nil, zero value otherwise.

### GetCompressedOk

`func (o *ReplicationUpdate1) GetCompressedOk() (*bool, bool)`

GetCompressedOk returns a tuple with the Compressed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressed

`func (o *ReplicationUpdate1) SetCompressed(v bool)`

SetCompressed sets Compressed field to given value.

### HasCompressed

`func (o *ReplicationUpdate1) HasCompressed() bool`

HasCompressed returns a boolean if a field has been set.

### GetRetries

`func (o *ReplicationUpdate1) GetRetries() int32`

GetRetries returns the Retries field if non-nil, zero value otherwise.

### GetRetriesOk

`func (o *ReplicationUpdate1) GetRetriesOk() (*int32, bool)`

GetRetriesOk returns a tuple with the Retries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetries

`func (o *ReplicationUpdate1) SetRetries(v int32)`

SetRetries sets Retries field to given value.

### HasRetries

`func (o *ReplicationUpdate1) HasRetries() bool`

HasRetries returns a boolean if a field has been set.

### GetLoggingLevel

`func (o *ReplicationUpdate1) GetLoggingLevel() string`

GetLoggingLevel returns the LoggingLevel field if non-nil, zero value otherwise.

### GetLoggingLevelOk

`func (o *ReplicationUpdate1) GetLoggingLevelOk() (*string, bool)`

GetLoggingLevelOk returns a tuple with the LoggingLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggingLevel

`func (o *ReplicationUpdate1) SetLoggingLevel(v string)`

SetLoggingLevel sets LoggingLevel field to given value.

### HasLoggingLevel

`func (o *ReplicationUpdate1) HasLoggingLevel() bool`

HasLoggingLevel returns a boolean if a field has been set.

### SetLoggingLevelNil

`func (o *ReplicationUpdate1) SetLoggingLevelNil(b bool)`

 SetLoggingLevelNil sets the value for LoggingLevel to be an explicit nil

### UnsetLoggingLevel
`func (o *ReplicationUpdate1) UnsetLoggingLevel()`

UnsetLoggingLevel ensures that no value is present for LoggingLevel, not even an explicit nil
### GetEnabled

`func (o *ReplicationUpdate1) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ReplicationUpdate1) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ReplicationUpdate1) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ReplicationUpdate1) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


