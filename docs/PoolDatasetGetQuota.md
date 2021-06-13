# PoolDatasetGetQuota

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QuotaType** | Pointer to [**PoolDatasetGetQuota1**](PoolDatasetGetQuota1.md) |  | [optional] 
**QueryFilters** | Pointer to **[]interface{}** |  | [optional] [default to []]
**QueryOptions** | Pointer to [**PoolDatasetGetQuota3**](PoolDatasetGetQuota3.md) |  | [optional] [default to {}]

## Methods

### NewPoolDatasetGetQuota

`func NewPoolDatasetGetQuota() *PoolDatasetGetQuota`

NewPoolDatasetGetQuota instantiates a new PoolDatasetGetQuota object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoolDatasetGetQuotaWithDefaults

`func NewPoolDatasetGetQuotaWithDefaults() *PoolDatasetGetQuota`

NewPoolDatasetGetQuotaWithDefaults instantiates a new PoolDatasetGetQuota object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuotaType

`func (o *PoolDatasetGetQuota) GetQuotaType() PoolDatasetGetQuota1`

GetQuotaType returns the QuotaType field if non-nil, zero value otherwise.

### GetQuotaTypeOk

`func (o *PoolDatasetGetQuota) GetQuotaTypeOk() (*PoolDatasetGetQuota1, bool)`

GetQuotaTypeOk returns a tuple with the QuotaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaType

`func (o *PoolDatasetGetQuota) SetQuotaType(v PoolDatasetGetQuota1)`

SetQuotaType sets QuotaType field to given value.

### HasQuotaType

`func (o *PoolDatasetGetQuota) HasQuotaType() bool`

HasQuotaType returns a boolean if a field has been set.

### GetQueryFilters

`func (o *PoolDatasetGetQuota) GetQueryFilters() []interface{}`

GetQueryFilters returns the QueryFilters field if non-nil, zero value otherwise.

### GetQueryFiltersOk

`func (o *PoolDatasetGetQuota) GetQueryFiltersOk() (*[]interface{}, bool)`

GetQueryFiltersOk returns a tuple with the QueryFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryFilters

`func (o *PoolDatasetGetQuota) SetQueryFilters(v []interface{})`

SetQueryFilters sets QueryFilters field to given value.

### HasQueryFilters

`func (o *PoolDatasetGetQuota) HasQueryFilters() bool`

HasQueryFilters returns a boolean if a field has been set.

### GetQueryOptions

`func (o *PoolDatasetGetQuota) GetQueryOptions() PoolDatasetGetQuota3`

GetQueryOptions returns the QueryOptions field if non-nil, zero value otherwise.

### GetQueryOptionsOk

`func (o *PoolDatasetGetQuota) GetQueryOptionsOk() (*PoolDatasetGetQuota3, bool)`

GetQueryOptionsOk returns a tuple with the QueryOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryOptions

`func (o *PoolDatasetGetQuota) SetQueryOptions(v PoolDatasetGetQuota3)`

SetQueryOptions sets QueryOptions field to given value.

### HasQueryOptions

`func (o *PoolDatasetGetQuota) HasQueryOptions() bool`

HasQueryOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


