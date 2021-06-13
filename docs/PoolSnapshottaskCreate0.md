# PoolSnapshottaskCreate0

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dataset** | Pointer to **string** |  | [optional] 
**Recursive** | Pointer to **bool** |  | [optional] 
**Exclude** | Pointer to **[]string** |  | [optional] 
**LifetimeValue** | Pointer to **int32** |  | [optional] 
**LifetimeUnit** | Pointer to **string** |  | [optional] 
**NamingSchema** | Pointer to **string** |  | [optional] 
**Schedule** | Pointer to [**PoolSnapshottaskCreate0Schedule**](PoolSnapshottaskCreate0Schedule.md) |  | [optional] 
**AllowEmpty** | Pointer to **bool** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewPoolSnapshottaskCreate0

`func NewPoolSnapshottaskCreate0() *PoolSnapshottaskCreate0`

NewPoolSnapshottaskCreate0 instantiates a new PoolSnapshottaskCreate0 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoolSnapshottaskCreate0WithDefaults

`func NewPoolSnapshottaskCreate0WithDefaults() *PoolSnapshottaskCreate0`

NewPoolSnapshottaskCreate0WithDefaults instantiates a new PoolSnapshottaskCreate0 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataset

`func (o *PoolSnapshottaskCreate0) GetDataset() string`

GetDataset returns the Dataset field if non-nil, zero value otherwise.

### GetDatasetOk

`func (o *PoolSnapshottaskCreate0) GetDatasetOk() (*string, bool)`

GetDatasetOk returns a tuple with the Dataset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataset

`func (o *PoolSnapshottaskCreate0) SetDataset(v string)`

SetDataset sets Dataset field to given value.

### HasDataset

`func (o *PoolSnapshottaskCreate0) HasDataset() bool`

HasDataset returns a boolean if a field has been set.

### GetRecursive

`func (o *PoolSnapshottaskCreate0) GetRecursive() bool`

GetRecursive returns the Recursive field if non-nil, zero value otherwise.

### GetRecursiveOk

`func (o *PoolSnapshottaskCreate0) GetRecursiveOk() (*bool, bool)`

GetRecursiveOk returns a tuple with the Recursive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecursive

`func (o *PoolSnapshottaskCreate0) SetRecursive(v bool)`

SetRecursive sets Recursive field to given value.

### HasRecursive

`func (o *PoolSnapshottaskCreate0) HasRecursive() bool`

HasRecursive returns a boolean if a field has been set.

### GetExclude

`func (o *PoolSnapshottaskCreate0) GetExclude() []string`

GetExclude returns the Exclude field if non-nil, zero value otherwise.

### GetExcludeOk

`func (o *PoolSnapshottaskCreate0) GetExcludeOk() (*[]string, bool)`

GetExcludeOk returns a tuple with the Exclude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclude

`func (o *PoolSnapshottaskCreate0) SetExclude(v []string)`

SetExclude sets Exclude field to given value.

### HasExclude

`func (o *PoolSnapshottaskCreate0) HasExclude() bool`

HasExclude returns a boolean if a field has been set.

### GetLifetimeValue

`func (o *PoolSnapshottaskCreate0) GetLifetimeValue() int32`

GetLifetimeValue returns the LifetimeValue field if non-nil, zero value otherwise.

### GetLifetimeValueOk

`func (o *PoolSnapshottaskCreate0) GetLifetimeValueOk() (*int32, bool)`

GetLifetimeValueOk returns a tuple with the LifetimeValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifetimeValue

`func (o *PoolSnapshottaskCreate0) SetLifetimeValue(v int32)`

SetLifetimeValue sets LifetimeValue field to given value.

### HasLifetimeValue

`func (o *PoolSnapshottaskCreate0) HasLifetimeValue() bool`

HasLifetimeValue returns a boolean if a field has been set.

### GetLifetimeUnit

`func (o *PoolSnapshottaskCreate0) GetLifetimeUnit() string`

GetLifetimeUnit returns the LifetimeUnit field if non-nil, zero value otherwise.

### GetLifetimeUnitOk

`func (o *PoolSnapshottaskCreate0) GetLifetimeUnitOk() (*string, bool)`

GetLifetimeUnitOk returns a tuple with the LifetimeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifetimeUnit

`func (o *PoolSnapshottaskCreate0) SetLifetimeUnit(v string)`

SetLifetimeUnit sets LifetimeUnit field to given value.

### HasLifetimeUnit

`func (o *PoolSnapshottaskCreate0) HasLifetimeUnit() bool`

HasLifetimeUnit returns a boolean if a field has been set.

### GetNamingSchema

`func (o *PoolSnapshottaskCreate0) GetNamingSchema() string`

GetNamingSchema returns the NamingSchema field if non-nil, zero value otherwise.

### GetNamingSchemaOk

`func (o *PoolSnapshottaskCreate0) GetNamingSchemaOk() (*string, bool)`

GetNamingSchemaOk returns a tuple with the NamingSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamingSchema

`func (o *PoolSnapshottaskCreate0) SetNamingSchema(v string)`

SetNamingSchema sets NamingSchema field to given value.

### HasNamingSchema

`func (o *PoolSnapshottaskCreate0) HasNamingSchema() bool`

HasNamingSchema returns a boolean if a field has been set.

### GetSchedule

`func (o *PoolSnapshottaskCreate0) GetSchedule() PoolSnapshottaskCreate0Schedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *PoolSnapshottaskCreate0) GetScheduleOk() (*PoolSnapshottaskCreate0Schedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *PoolSnapshottaskCreate0) SetSchedule(v PoolSnapshottaskCreate0Schedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *PoolSnapshottaskCreate0) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetAllowEmpty

`func (o *PoolSnapshottaskCreate0) GetAllowEmpty() bool`

GetAllowEmpty returns the AllowEmpty field if non-nil, zero value otherwise.

### GetAllowEmptyOk

`func (o *PoolSnapshottaskCreate0) GetAllowEmptyOk() (*bool, bool)`

GetAllowEmptyOk returns a tuple with the AllowEmpty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowEmpty

`func (o *PoolSnapshottaskCreate0) SetAllowEmpty(v bool)`

SetAllowEmpty sets AllowEmpty field to given value.

### HasAllowEmpty

`func (o *PoolSnapshottaskCreate0) HasAllowEmpty() bool`

HasAllowEmpty returns a boolean if a field has been set.

### GetEnabled

`func (o *PoolSnapshottaskCreate0) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PoolSnapshottaskCreate0) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PoolSnapshottaskCreate0) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PoolSnapshottaskCreate0) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


