# DiskWipe

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dev** | Pointer to **string** |  | [optional] 
**Mode** | Pointer to [**DiskWipe1**](DiskWipe1.md) |  | [optional] 
**Synccache** | Pointer to **bool** |  | [optional] [default to true]
**SwapRemovalOptions** | Pointer to [**DiskWipe3**](DiskWipe3.md) |  | [optional] [default to {}]

## Methods

### NewDiskWipe

`func NewDiskWipe() *DiskWipe`

NewDiskWipe instantiates a new DiskWipe object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiskWipeWithDefaults

`func NewDiskWipeWithDefaults() *DiskWipe`

NewDiskWipeWithDefaults instantiates a new DiskWipe object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDev

`func (o *DiskWipe) GetDev() string`

GetDev returns the Dev field if non-nil, zero value otherwise.

### GetDevOk

`func (o *DiskWipe) GetDevOk() (*string, bool)`

GetDevOk returns a tuple with the Dev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDev

`func (o *DiskWipe) SetDev(v string)`

SetDev sets Dev field to given value.

### HasDev

`func (o *DiskWipe) HasDev() bool`

HasDev returns a boolean if a field has been set.

### GetMode

`func (o *DiskWipe) GetMode() DiskWipe1`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *DiskWipe) GetModeOk() (*DiskWipe1, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *DiskWipe) SetMode(v DiskWipe1)`

SetMode sets Mode field to given value.

### HasMode

`func (o *DiskWipe) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetSynccache

`func (o *DiskWipe) GetSynccache() bool`

GetSynccache returns the Synccache field if non-nil, zero value otherwise.

### GetSynccacheOk

`func (o *DiskWipe) GetSynccacheOk() (*bool, bool)`

GetSynccacheOk returns a tuple with the Synccache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynccache

`func (o *DiskWipe) SetSynccache(v bool)`

SetSynccache sets Synccache field to given value.

### HasSynccache

`func (o *DiskWipe) HasSynccache() bool`

HasSynccache returns a boolean if a field has been set.

### GetSwapRemovalOptions

`func (o *DiskWipe) GetSwapRemovalOptions() DiskWipe3`

GetSwapRemovalOptions returns the SwapRemovalOptions field if non-nil, zero value otherwise.

### GetSwapRemovalOptionsOk

`func (o *DiskWipe) GetSwapRemovalOptionsOk() (*DiskWipe3, bool)`

GetSwapRemovalOptionsOk returns a tuple with the SwapRemovalOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwapRemovalOptions

`func (o *DiskWipe) SetSwapRemovalOptions(v DiskWipe3)`

SetSwapRemovalOptions sets SwapRemovalOptions field to given value.

### HasSwapRemovalOptions

`func (o *DiskWipe) HasSwapRemovalOptions() bool`

HasSwapRemovalOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


