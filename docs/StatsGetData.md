# StatsGetData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StatsList** | Pointer to **[]map[string]interface{}** |  | [optional] 
**StatsFilter** | Pointer to [**StatsGetData1**](StatsGetData1.md) |  | [optional] [default to {}]

## Methods

### NewStatsGetData

`func NewStatsGetData() *StatsGetData`

NewStatsGetData instantiates a new StatsGetData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatsGetDataWithDefaults

`func NewStatsGetDataWithDefaults() *StatsGetData`

NewStatsGetDataWithDefaults instantiates a new StatsGetData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatsList

`func (o *StatsGetData) GetStatsList() []map[string]interface{}`

GetStatsList returns the StatsList field if non-nil, zero value otherwise.

### GetStatsListOk

`func (o *StatsGetData) GetStatsListOk() (*[]map[string]interface{}, bool)`

GetStatsListOk returns a tuple with the StatsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatsList

`func (o *StatsGetData) SetStatsList(v []map[string]interface{})`

SetStatsList sets StatsList field to given value.

### HasStatsList

`func (o *StatsGetData) HasStatsList() bool`

HasStatsList returns a boolean if a field has been set.

### GetStatsFilter

`func (o *StatsGetData) GetStatsFilter() StatsGetData1`

GetStatsFilter returns the StatsFilter field if non-nil, zero value otherwise.

### GetStatsFilterOk

`func (o *StatsGetData) GetStatsFilterOk() (*StatsGetData1, bool)`

GetStatsFilterOk returns a tuple with the StatsFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatsFilter

`func (o *StatsGetData) SetStatsFilter(v StatsGetData1)`

SetStatsFilter sets StatsFilter field to given value.

### HasStatsFilter

`func (o *StatsGetData) HasStatsFilter() bool`

HasStatsFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


