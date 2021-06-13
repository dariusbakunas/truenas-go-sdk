# SmbStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InfoLevel** | Pointer to [**SmbStatus0**](SmbStatus0.md) |  | [optional] [default to ALL]
**QueryFilters** | Pointer to **[]interface{}** |  | [optional] [default to []]
**QueryOptions** | Pointer to [**SmbStatus2**](SmbStatus2.md) |  | [optional] [default to {}]
**StatusOptions** | Pointer to [**SmbStatus3**](SmbStatus3.md) |  | [optional] [default to {}]

## Methods

### NewSmbStatus

`func NewSmbStatus() *SmbStatus`

NewSmbStatus instantiates a new SmbStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmbStatusWithDefaults

`func NewSmbStatusWithDefaults() *SmbStatus`

NewSmbStatusWithDefaults instantiates a new SmbStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInfoLevel

`func (o *SmbStatus) GetInfoLevel() SmbStatus0`

GetInfoLevel returns the InfoLevel field if non-nil, zero value otherwise.

### GetInfoLevelOk

`func (o *SmbStatus) GetInfoLevelOk() (*SmbStatus0, bool)`

GetInfoLevelOk returns a tuple with the InfoLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfoLevel

`func (o *SmbStatus) SetInfoLevel(v SmbStatus0)`

SetInfoLevel sets InfoLevel field to given value.

### HasInfoLevel

`func (o *SmbStatus) HasInfoLevel() bool`

HasInfoLevel returns a boolean if a field has been set.

### GetQueryFilters

`func (o *SmbStatus) GetQueryFilters() []interface{}`

GetQueryFilters returns the QueryFilters field if non-nil, zero value otherwise.

### GetQueryFiltersOk

`func (o *SmbStatus) GetQueryFiltersOk() (*[]interface{}, bool)`

GetQueryFiltersOk returns a tuple with the QueryFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryFilters

`func (o *SmbStatus) SetQueryFilters(v []interface{})`

SetQueryFilters sets QueryFilters field to given value.

### HasQueryFilters

`func (o *SmbStatus) HasQueryFilters() bool`

HasQueryFilters returns a boolean if a field has been set.

### GetQueryOptions

`func (o *SmbStatus) GetQueryOptions() SmbStatus2`

GetQueryOptions returns the QueryOptions field if non-nil, zero value otherwise.

### GetQueryOptionsOk

`func (o *SmbStatus) GetQueryOptionsOk() (*SmbStatus2, bool)`

GetQueryOptionsOk returns a tuple with the QueryOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryOptions

`func (o *SmbStatus) SetQueryOptions(v SmbStatus2)`

SetQueryOptions sets QueryOptions field to given value.

### HasQueryOptions

`func (o *SmbStatus) HasQueryOptions() bool`

HasQueryOptions returns a boolean if a field has been set.

### GetStatusOptions

`func (o *SmbStatus) GetStatusOptions() SmbStatus3`

GetStatusOptions returns the StatusOptions field if non-nil, zero value otherwise.

### GetStatusOptionsOk

`func (o *SmbStatus) GetStatusOptionsOk() (*SmbStatus3, bool)`

GetStatusOptionsOk returns a tuple with the StatusOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusOptions

`func (o *SmbStatus) SetStatusOptions(v SmbStatus3)`

SetStatusOptions sets StatusOptions field to given value.

### HasStatusOptions

`func (o *SmbStatus) HasStatusOptions() bool`

HasStatusOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


