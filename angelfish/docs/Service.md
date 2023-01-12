# Service

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Service** | **string** |  | 
**Enable** | Pointer to **bool** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Pids** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewService

`func NewService(id int32, service string, ) *Service`

NewService instantiates a new Service object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceWithDefaults

`func NewServiceWithDefaults() *Service`

NewServiceWithDefaults instantiates a new Service object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Service) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Service) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Service) SetId(v int32)`

SetId sets Id field to given value.


### GetService

`func (o *Service) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *Service) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *Service) SetService(v string)`

SetService sets Service field to given value.


### GetEnable

`func (o *Service) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *Service) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *Service) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *Service) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetState

`func (o *Service) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Service) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Service) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Service) HasState() bool`

HasState returns a boolean if a field has been set.

### GetPids

`func (o *Service) GetPids() []int32`

GetPids returns the Pids field if non-nil, zero value otherwise.

### GetPidsOk

`func (o *Service) GetPidsOk() (*[]int32, bool)`

GetPidsOk returns a tuple with the Pids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPids

`func (o *Service) SetPids(v []int32)`

SetPids sets Pids field to given value.

### HasPids

`func (o *Service) HasPids() bool`

HasPids returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


