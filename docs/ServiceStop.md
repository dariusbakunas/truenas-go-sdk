# ServiceStop

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Service** | Pointer to **string** |  | [optional] 
**ServiceControl** | Pointer to [**ServiceStop1**](ServiceStop1.md) |  | [optional] [default to {}]

## Methods

### NewServiceStop

`func NewServiceStop() *ServiceStop`

NewServiceStop instantiates a new ServiceStop object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceStopWithDefaults

`func NewServiceStopWithDefaults() *ServiceStop`

NewServiceStopWithDefaults instantiates a new ServiceStop object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetService

`func (o *ServiceStop) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *ServiceStop) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *ServiceStop) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *ServiceStop) HasService() bool`

HasService returns a boolean if a field has been set.

### GetServiceControl

`func (o *ServiceStop) GetServiceControl() ServiceStop1`

GetServiceControl returns the ServiceControl field if non-nil, zero value otherwise.

### GetServiceControlOk

`func (o *ServiceStop) GetServiceControlOk() (*ServiceStop1, bool)`

GetServiceControlOk returns a tuple with the ServiceControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceControl

`func (o *ServiceStop) SetServiceControl(v ServiceStop1)`

SetServiceControl sets ServiceControl field to given value.

### HasServiceControl

`func (o *ServiceStop) HasServiceControl() bool`

HasServiceControl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


