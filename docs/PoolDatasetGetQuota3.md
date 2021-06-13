# PoolDatasetGetQuota3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Relationships** | Pointer to **bool** |  | [optional] 
**Extend** | Pointer to **NullableString** |  | [optional] 
**ExtendContext** | Pointer to **NullableString** |  | [optional] 
**Prefix** | Pointer to **NullableString** |  | [optional] 
**Extra** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**OrderBy** | Pointer to **[]interface{}** |  | [optional] 
**Select** | Pointer to **[]interface{}** |  | [optional] 
**Count** | Pointer to **bool** |  | [optional] 
**Get** | Pointer to **bool** |  | [optional] 
**Offset** | Pointer to **int32** |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 

## Methods

### NewPoolDatasetGetQuota3

`func NewPoolDatasetGetQuota3() *PoolDatasetGetQuota3`

NewPoolDatasetGetQuota3 instantiates a new PoolDatasetGetQuota3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoolDatasetGetQuota3WithDefaults

`func NewPoolDatasetGetQuota3WithDefaults() *PoolDatasetGetQuota3`

NewPoolDatasetGetQuota3WithDefaults instantiates a new PoolDatasetGetQuota3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRelationships

`func (o *PoolDatasetGetQuota3) GetRelationships() bool`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PoolDatasetGetQuota3) GetRelationshipsOk() (*bool, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PoolDatasetGetQuota3) SetRelationships(v bool)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PoolDatasetGetQuota3) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetExtend

`func (o *PoolDatasetGetQuota3) GetExtend() string`

GetExtend returns the Extend field if non-nil, zero value otherwise.

### GetExtendOk

`func (o *PoolDatasetGetQuota3) GetExtendOk() (*string, bool)`

GetExtendOk returns a tuple with the Extend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtend

`func (o *PoolDatasetGetQuota3) SetExtend(v string)`

SetExtend sets Extend field to given value.

### HasExtend

`func (o *PoolDatasetGetQuota3) HasExtend() bool`

HasExtend returns a boolean if a field has been set.

### SetExtendNil

`func (o *PoolDatasetGetQuota3) SetExtendNil(b bool)`

 SetExtendNil sets the value for Extend to be an explicit nil

### UnsetExtend
`func (o *PoolDatasetGetQuota3) UnsetExtend()`

UnsetExtend ensures that no value is present for Extend, not even an explicit nil
### GetExtendContext

`func (o *PoolDatasetGetQuota3) GetExtendContext() string`

GetExtendContext returns the ExtendContext field if non-nil, zero value otherwise.

### GetExtendContextOk

`func (o *PoolDatasetGetQuota3) GetExtendContextOk() (*string, bool)`

GetExtendContextOk returns a tuple with the ExtendContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendContext

`func (o *PoolDatasetGetQuota3) SetExtendContext(v string)`

SetExtendContext sets ExtendContext field to given value.

### HasExtendContext

`func (o *PoolDatasetGetQuota3) HasExtendContext() bool`

HasExtendContext returns a boolean if a field has been set.

### SetExtendContextNil

`func (o *PoolDatasetGetQuota3) SetExtendContextNil(b bool)`

 SetExtendContextNil sets the value for ExtendContext to be an explicit nil

### UnsetExtendContext
`func (o *PoolDatasetGetQuota3) UnsetExtendContext()`

UnsetExtendContext ensures that no value is present for ExtendContext, not even an explicit nil
### GetPrefix

`func (o *PoolDatasetGetQuota3) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *PoolDatasetGetQuota3) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *PoolDatasetGetQuota3) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *PoolDatasetGetQuota3) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### SetPrefixNil

`func (o *PoolDatasetGetQuota3) SetPrefixNil(b bool)`

 SetPrefixNil sets the value for Prefix to be an explicit nil

### UnsetPrefix
`func (o *PoolDatasetGetQuota3) UnsetPrefix()`

UnsetPrefix ensures that no value is present for Prefix, not even an explicit nil
### GetExtra

`func (o *PoolDatasetGetQuota3) GetExtra() map[string]map[string]interface{}`

GetExtra returns the Extra field if non-nil, zero value otherwise.

### GetExtraOk

`func (o *PoolDatasetGetQuota3) GetExtraOk() (*map[string]map[string]interface{}, bool)`

GetExtraOk returns a tuple with the Extra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtra

`func (o *PoolDatasetGetQuota3) SetExtra(v map[string]map[string]interface{})`

SetExtra sets Extra field to given value.

### HasExtra

`func (o *PoolDatasetGetQuota3) HasExtra() bool`

HasExtra returns a boolean if a field has been set.

### GetOrderBy

`func (o *PoolDatasetGetQuota3) GetOrderBy() []interface{}`

GetOrderBy returns the OrderBy field if non-nil, zero value otherwise.

### GetOrderByOk

`func (o *PoolDatasetGetQuota3) GetOrderByOk() (*[]interface{}, bool)`

GetOrderByOk returns a tuple with the OrderBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBy

`func (o *PoolDatasetGetQuota3) SetOrderBy(v []interface{})`

SetOrderBy sets OrderBy field to given value.

### HasOrderBy

`func (o *PoolDatasetGetQuota3) HasOrderBy() bool`

HasOrderBy returns a boolean if a field has been set.

### GetSelect

`func (o *PoolDatasetGetQuota3) GetSelect() []interface{}`

GetSelect returns the Select field if non-nil, zero value otherwise.

### GetSelectOk

`func (o *PoolDatasetGetQuota3) GetSelectOk() (*[]interface{}, bool)`

GetSelectOk returns a tuple with the Select field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelect

`func (o *PoolDatasetGetQuota3) SetSelect(v []interface{})`

SetSelect sets Select field to given value.

### HasSelect

`func (o *PoolDatasetGetQuota3) HasSelect() bool`

HasSelect returns a boolean if a field has been set.

### GetCount

`func (o *PoolDatasetGetQuota3) GetCount() bool`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PoolDatasetGetQuota3) GetCountOk() (*bool, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PoolDatasetGetQuota3) SetCount(v bool)`

SetCount sets Count field to given value.

### HasCount

`func (o *PoolDatasetGetQuota3) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetGet

`func (o *PoolDatasetGetQuota3) GetGet() bool`

GetGet returns the Get field if non-nil, zero value otherwise.

### GetGetOk

`func (o *PoolDatasetGetQuota3) GetGetOk() (*bool, bool)`

GetGetOk returns a tuple with the Get field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGet

`func (o *PoolDatasetGetQuota3) SetGet(v bool)`

SetGet sets Get field to given value.

### HasGet

`func (o *PoolDatasetGetQuota3) HasGet() bool`

HasGet returns a boolean if a field has been set.

### GetOffset

`func (o *PoolDatasetGetQuota3) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *PoolDatasetGetQuota3) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *PoolDatasetGetQuota3) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *PoolDatasetGetQuota3) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetLimit

`func (o *PoolDatasetGetQuota3) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *PoolDatasetGetQuota3) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *PoolDatasetGetQuota3) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *PoolDatasetGetQuota3) HasLimit() bool`

HasLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


