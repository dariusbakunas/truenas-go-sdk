# DiskTemperature

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Powermode** | Pointer to [**DiskTemperature1**](DiskTemperature1.md) |  | [optional] [default to NEVER]

## Methods

### NewDiskTemperature

`func NewDiskTemperature() *DiskTemperature`

NewDiskTemperature instantiates a new DiskTemperature object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiskTemperatureWithDefaults

`func NewDiskTemperatureWithDefaults() *DiskTemperature`

NewDiskTemperatureWithDefaults instantiates a new DiskTemperature object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DiskTemperature) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DiskTemperature) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DiskTemperature) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DiskTemperature) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPowermode

`func (o *DiskTemperature) GetPowermode() DiskTemperature1`

GetPowermode returns the Powermode field if non-nil, zero value otherwise.

### GetPowermodeOk

`func (o *DiskTemperature) GetPowermodeOk() (*DiskTemperature1, bool)`

GetPowermodeOk returns a tuple with the Powermode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowermode

`func (o *DiskTemperature) SetPowermode(v DiskTemperature1)`

SetPowermode sets Powermode field to given value.

### HasPowermode

`func (o *DiskTemperature) HasPowermode() bool`

HasPowermode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


