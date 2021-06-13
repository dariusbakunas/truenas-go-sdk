# ServiceStart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Service** | Pointer to **string** |  | [optional] 
**ServiceControl** | Pointer to [**ServiceStart1**](ServiceStart1.md) |  | [optional] [default to {}]

## Methods

### NewServiceStart

`func NewServiceStart() *ServiceStart`

NewServiceStart instantiates a new ServiceStart object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceStartWithDefaults

`func NewServiceStartWithDefaults() *ServiceStart`

NewServiceStartWithDefaults instantiates a new ServiceStart object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetService

`func (o *ServiceStart) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *ServiceStart) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *ServiceStart) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *ServiceStart) HasService() bool`

HasService returns a boolean if a field has been set.

### GetServiceControl

`func (o *ServiceStart) GetServiceControl() ServiceStart1`

GetServiceControl returns the ServiceControl field if non-nil, zero value otherwise.

### GetServiceControlOk

`func (o *ServiceStart) GetServiceControlOk() (*ServiceStart1, bool)`

GetServiceControlOk returns a tuple with the ServiceControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceControl

`func (o *ServiceStart) SetServiceControl(v ServiceStart1)`

SetServiceControl sets ServiceControl field to given value.

### HasServiceControl

`func (o *ServiceStart) HasServiceControl() bool`

HasServiceControl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


