# DiskTemperatures

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Names** | Pointer to **[]string** |  | [optional] 
**Powermode** | Pointer to [**DiskTemperatures1**](DiskTemperatures1.md) |  | [optional] [default to NEVER]

## Methods

### NewDiskTemperatures

`func NewDiskTemperatures() *DiskTemperatures`

NewDiskTemperatures instantiates a new DiskTemperatures object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiskTemperaturesWithDefaults

`func NewDiskTemperaturesWithDefaults() *DiskTemperatures`

NewDiskTemperaturesWithDefaults instantiates a new DiskTemperatures object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNames

`func (o *DiskTemperatures) GetNames() []string`

GetNames returns the Names field if non-nil, zero value otherwise.

### GetNamesOk

`func (o *DiskTemperatures) GetNamesOk() (*[]string, bool)`

GetNamesOk returns a tuple with the Names field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNames

`func (o *DiskTemperatures) SetNames(v []string)`

SetNames sets Names field to given value.

### HasNames

`func (o *DiskTemperatures) HasNames() bool`

HasNames returns a boolean if a field has been set.

### GetPowermode

`func (o *DiskTemperatures) GetPowermode() DiskTemperatures1`

GetPowermode returns the Powermode field if non-nil, zero value otherwise.

### GetPowermodeOk

`func (o *DiskTemperatures) GetPowermodeOk() (*DiskTemperatures1, bool)`

GetPowermodeOk returns a tuple with the Powermode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowermode

`func (o *DiskTemperatures) SetPowermode(v DiskTemperatures1)`

SetPowermode sets Powermode field to given value.

### HasPowermode

`func (o *DiskTemperatures) HasPowermode() bool`

HasPowermode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


