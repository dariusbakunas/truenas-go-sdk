# ServiceReload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Service** | Pointer to **string** |  | [optional] 
**ServiceControl** | Pointer to [**ServiceReload1**](ServiceReload1.md) |  | [optional] [default to {}]

## Methods

### NewServiceReload

`func NewServiceReload() *ServiceReload`

NewServiceReload instantiates a new ServiceReload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceReloadWithDefaults

`func NewServiceReloadWithDefaults() *ServiceReload`

NewServiceReloadWithDefaults instantiates a new ServiceReload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetService

`func (o *ServiceReload) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *ServiceReload) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *ServiceReload) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *ServiceReload) HasService() bool`

HasService returns a boolean if a field has been set.

### GetServiceControl

`func (o *ServiceReload) GetServiceControl() ServiceReload1`

GetServiceControl returns the ServiceControl field if non-nil, zero value otherwise.

### GetServiceControlOk

`func (o *ServiceReload) GetServiceControlOk() (*ServiceReload1, bool)`

GetServiceControlOk returns a tuple with the ServiceControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceControl

`func (o *ServiceReload) SetServiceControl(v ServiceReload1)`

SetServiceControl sets ServiceControl field to given value.

### HasServiceControl

`func (o *ServiceReload) HasServiceControl() bool`

HasServiceControl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


