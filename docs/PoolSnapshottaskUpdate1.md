# PoolSnapshottaskUpdate1

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

### NewPoolSnapshottaskUpdate1

`func NewPoolSnapshottaskUpdate1() *PoolSnapshottaskUpdate1`

NewPoolSnapshottaskUpdate1 instantiates a new PoolSnapshottaskUpdate1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoolSnapshottaskUpdate1WithDefaults

`func NewPoolSnapshottaskUpdate1WithDefaults() *PoolSnapshottaskUpdate1`

NewPoolSnapshottaskUpdate1WithDefaults instantiates a new PoolSnapshottaskUpdate1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataset

`func (o *PoolSnapshottaskUpdate1) GetDataset() string`

GetDataset returns the Dataset field if non-nil, zero value otherwise.

### GetDatasetOk

`func (o *PoolSnapshottaskUpdate1) GetDatasetOk() (*string, bool)`

GetDatasetOk returns a tuple with the Dataset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataset

`func (o *PoolSnapshottaskUpdate1) SetDataset(v string)`

SetDataset sets Dataset field to given value.

### HasDataset

`func (o *PoolSnapshottaskUpdate1) HasDataset() bool`

HasDataset returns a boolean if a field has been set.

### GetRecursive

`func (o *PoolSnapshottaskUpdate1) GetRecursive() bool`

GetRecursive returns the Recursive field if non-nil, zero value otherwise.

### GetRecursiveOk

`func (o *PoolSnapshottaskUpdate1) GetRecursiveOk() (*bool, bool)`

GetRecursiveOk returns a tuple with the Recursive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecursive

`func (o *PoolSnapshottaskUpdate1) SetRecursive(v bool)`

SetRecursive sets Recursive field to given value.

### HasRecursive

`func (o *PoolSnapshottaskUpdate1) HasRecursive() bool`

HasRecursive returns a boolean if a field has been set.

### GetExclude

`func (o *PoolSnapshottaskUpdate1) GetExclude() []string`

GetExclude returns the Exclude field if non-nil, zero value otherwise.

### GetExcludeOk

`func (o *PoolSnapshottaskUpdate1) GetExcludeOk() (*[]string, bool)`

GetExcludeOk returns a tuple with the Exclude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclude

`func (o *PoolSnapshottaskUpdate1) SetExclude(v []string)`

SetExclude sets Exclude field to given value.

### HasExclude

`func (o *PoolSnapshottaskUpdate1) HasExclude() bool`

HasExclude returns a boolean if a field has been set.

### GetLifetimeValue

`func (o *PoolSnapshottaskUpdate1) GetLifetimeValue() int32`

GetLifetimeValue returns the LifetimeValue field if non-nil, zero value otherwise.

### GetLifetimeValueOk

`func (o *PoolSnapshottaskUpdate1) GetLifetimeValueOk() (*int32, bool)`

GetLifetimeValueOk returns a tuple with the LifetimeValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifetimeValue

`func (o *PoolSnapshottaskUpdate1) SetLifetimeValue(v int32)`

SetLifetimeValue sets LifetimeValue field to given value.

### HasLifetimeValue

`func (o *PoolSnapshottaskUpdate1) HasLifetimeValue() bool`

HasLifetimeValue returns a boolean if a field has been set.

### GetLifetimeUnit

`func (o *PoolSnapshottaskUpdate1) GetLifetimeUnit() string`

GetLifetimeUnit returns the LifetimeUnit field if non-nil, zero value otherwise.

### GetLifetimeUnitOk

`func (o *PoolSnapshottaskUpdate1) GetLifetimeUnitOk() (*string, bool)`

GetLifetimeUnitOk returns a tuple with the LifetimeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifetimeUnit

`func (o *PoolSnapshottaskUpdate1) SetLifetimeUnit(v string)`

SetLifetimeUnit sets LifetimeUnit field to given value.

### HasLifetimeUnit

`func (o *PoolSnapshottaskUpdate1) HasLifetimeUnit() bool`

HasLifetimeUnit returns a boolean if a field has been set.

### GetNamingSchema

`func (o *PoolSnapshottaskUpdate1) GetNamingSchema() string`

GetNamingSchema returns the NamingSchema field if non-nil, zero value otherwise.

### GetNamingSchemaOk

`func (o *PoolSnapshottaskUpdate1) GetNamingSchemaOk() (*string, bool)`

GetNamingSchemaOk returns a tuple with the NamingSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamingSchema

`func (o *PoolSnapshottaskUpdate1) SetNamingSchema(v string)`

SetNamingSchema sets NamingSchema field to given value.

### HasNamingSchema

`func (o *PoolSnapshottaskUpdate1) HasNamingSchema() bool`

HasNamingSchema returns a boolean if a field has been set.

### GetSchedule

`func (o *PoolSnapshottaskUpdate1) GetSchedule() PoolSnapshottaskCreate0Schedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *PoolSnapshottaskUpdate1) GetScheduleOk() (*PoolSnapshottaskCreate0Schedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *PoolSnapshottaskUpdate1) SetSchedule(v PoolSnapshottaskCreate0Schedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *PoolSnapshottaskUpdate1) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetAllowEmpty

`func (o *PoolSnapshottaskUpdate1) GetAllowEmpty() bool`

GetAllowEmpty returns the AllowEmpty field if non-nil, zero value otherwise.

### GetAllowEmptyOk

`func (o *PoolSnapshottaskUpdate1) GetAllowEmptyOk() (*bool, bool)`

GetAllowEmptyOk returns a tuple with the AllowEmpty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowEmpty

`func (o *PoolSnapshottaskUpdate1) SetAllowEmpty(v bool)`

SetAllowEmpty sets AllowEmpty field to given value.

### HasAllowEmpty

`func (o *PoolSnapshottaskUpdate1) HasAllowEmpty() bool`

HasAllowEmpty returns a boolean if a field has been set.

### GetEnabled

`func (o *PoolSnapshottaskUpdate1) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PoolSnapshottaskUpdate1) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PoolSnapshottaskUpdate1) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PoolSnapshottaskUpdate1) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


