# DiskUpdate1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Togglesmart** | Pointer to **bool** |  | [optional] 
**Acousticlevel** | Pointer to **string** |  | [optional] 
**Advpowermgmt** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Hddstandby** | Pointer to **string** |  | [optional] 
**HddstandbyForce** | Pointer to **bool** |  | [optional] 
**Passwd** | Pointer to **string** |  | [optional] 
**Smartoptions** | Pointer to **string** |  | [optional] 
**Critical** | Pointer to **NullableInt32** |  | [optional] 
**Difference** | Pointer to **NullableInt32** |  | [optional] 
**Informational** | Pointer to **NullableInt32** |  | [optional] 
**Enclosure** | Pointer to [**DiskUpdate1Enclosure**](DiskUpdate1Enclosure.md) |  | [optional] 

## Methods

### NewDiskUpdate1

`func NewDiskUpdate1() *DiskUpdate1`

NewDiskUpdate1 instantiates a new DiskUpdate1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiskUpdate1WithDefaults

`func NewDiskUpdate1WithDefaults() *DiskUpdate1`

NewDiskUpdate1WithDefaults instantiates a new DiskUpdate1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTogglesmart

`func (o *DiskUpdate1) GetTogglesmart() bool`

GetTogglesmart returns the Togglesmart field if non-nil, zero value otherwise.

### GetTogglesmartOk

`func (o *DiskUpdate1) GetTogglesmartOk() (*bool, bool)`

GetTogglesmartOk returns a tuple with the Togglesmart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTogglesmart

`func (o *DiskUpdate1) SetTogglesmart(v bool)`

SetTogglesmart sets Togglesmart field to given value.

### HasTogglesmart

`func (o *DiskUpdate1) HasTogglesmart() bool`

HasTogglesmart returns a boolean if a field has been set.

### GetAcousticlevel

`func (o *DiskUpdate1) GetAcousticlevel() string`

GetAcousticlevel returns the Acousticlevel field if non-nil, zero value otherwise.

### GetAcousticlevelOk

`func (o *DiskUpdate1) GetAcousticlevelOk() (*string, bool)`

GetAcousticlevelOk returns a tuple with the Acousticlevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcousticlevel

`func (o *DiskUpdate1) SetAcousticlevel(v string)`

SetAcousticlevel sets Acousticlevel field to given value.

### HasAcousticlevel

`func (o *DiskUpdate1) HasAcousticlevel() bool`

HasAcousticlevel returns a boolean if a field has been set.

### GetAdvpowermgmt

`func (o *DiskUpdate1) GetAdvpowermgmt() string`

GetAdvpowermgmt returns the Advpowermgmt field if non-nil, zero value otherwise.

### GetAdvpowermgmtOk

`func (o *DiskUpdate1) GetAdvpowermgmtOk() (*string, bool)`

GetAdvpowermgmtOk returns a tuple with the Advpowermgmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvpowermgmt

`func (o *DiskUpdate1) SetAdvpowermgmt(v string)`

SetAdvpowermgmt sets Advpowermgmt field to given value.

### HasAdvpowermgmt

`func (o *DiskUpdate1) HasAdvpowermgmt() bool`

HasAdvpowermgmt returns a boolean if a field has been set.

### GetDescription

`func (o *DiskUpdate1) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DiskUpdate1) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DiskUpdate1) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DiskUpdate1) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHddstandby

`func (o *DiskUpdate1) GetHddstandby() string`

GetHddstandby returns the Hddstandby field if non-nil, zero value otherwise.

### GetHddstandbyOk

`func (o *DiskUpdate1) GetHddstandbyOk() (*string, bool)`

GetHddstandbyOk returns a tuple with the Hddstandby field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHddstandby

`func (o *DiskUpdate1) SetHddstandby(v string)`

SetHddstandby sets Hddstandby field to given value.

### HasHddstandby

`func (o *DiskUpdate1) HasHddstandby() bool`

HasHddstandby returns a boolean if a field has been set.

### GetHddstandbyForce

`func (o *DiskUpdate1) GetHddstandbyForce() bool`

GetHddstandbyForce returns the HddstandbyForce field if non-nil, zero value otherwise.

### GetHddstandbyForceOk

`func (o *DiskUpdate1) GetHddstandbyForceOk() (*bool, bool)`

GetHddstandbyForceOk returns a tuple with the HddstandbyForce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHddstandbyForce

`func (o *DiskUpdate1) SetHddstandbyForce(v bool)`

SetHddstandbyForce sets HddstandbyForce field to given value.

### HasHddstandbyForce

`func (o *DiskUpdate1) HasHddstandbyForce() bool`

HasHddstandbyForce returns a boolean if a field has been set.

### GetPasswd

`func (o *DiskUpdate1) GetPasswd() string`

GetPasswd returns the Passwd field if non-nil, zero value otherwise.

### GetPasswdOk

`func (o *DiskUpdate1) GetPasswdOk() (*string, bool)`

GetPasswdOk returns a tuple with the Passwd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswd

`func (o *DiskUpdate1) SetPasswd(v string)`

SetPasswd sets Passwd field to given value.

### HasPasswd

`func (o *DiskUpdate1) HasPasswd() bool`

HasPasswd returns a boolean if a field has been set.

### GetSmartoptions

`func (o *DiskUpdate1) GetSmartoptions() string`

GetSmartoptions returns the Smartoptions field if non-nil, zero value otherwise.

### GetSmartoptionsOk

`func (o *DiskUpdate1) GetSmartoptionsOk() (*string, bool)`

GetSmartoptionsOk returns a tuple with the Smartoptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartoptions

`func (o *DiskUpdate1) SetSmartoptions(v string)`

SetSmartoptions sets Smartoptions field to given value.

### HasSmartoptions

`func (o *DiskUpdate1) HasSmartoptions() bool`

HasSmartoptions returns a boolean if a field has been set.

### GetCritical

`func (o *DiskUpdate1) GetCritical() int32`

GetCritical returns the Critical field if non-nil, zero value otherwise.

### GetCriticalOk

`func (o *DiskUpdate1) GetCriticalOk() (*int32, bool)`

GetCriticalOk returns a tuple with the Critical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCritical

`func (o *DiskUpdate1) SetCritical(v int32)`

SetCritical sets Critical field to given value.

### HasCritical

`func (o *DiskUpdate1) HasCritical() bool`

HasCritical returns a boolean if a field has been set.

### SetCriticalNil

`func (o *DiskUpdate1) SetCriticalNil(b bool)`

 SetCriticalNil sets the value for Critical to be an explicit nil

### UnsetCritical
`func (o *DiskUpdate1) UnsetCritical()`

UnsetCritical ensures that no value is present for Critical, not even an explicit nil
### GetDifference

`func (o *DiskUpdate1) GetDifference() int32`

GetDifference returns the Difference field if non-nil, zero value otherwise.

### GetDifferenceOk

`func (o *DiskUpdate1) GetDifferenceOk() (*int32, bool)`

GetDifferenceOk returns a tuple with the Difference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDifference

`func (o *DiskUpdate1) SetDifference(v int32)`

SetDifference sets Difference field to given value.

### HasDifference

`func (o *DiskUpdate1) HasDifference() bool`

HasDifference returns a boolean if a field has been set.

### SetDifferenceNil

`func (o *DiskUpdate1) SetDifferenceNil(b bool)`

 SetDifferenceNil sets the value for Difference to be an explicit nil

### UnsetDifference
`func (o *DiskUpdate1) UnsetDifference()`

UnsetDifference ensures that no value is present for Difference, not even an explicit nil
### GetInformational

`func (o *DiskUpdate1) GetInformational() int32`

GetInformational returns the Informational field if non-nil, zero value otherwise.

### GetInformationalOk

`func (o *DiskUpdate1) GetInformationalOk() (*int32, bool)`

GetInformationalOk returns a tuple with the Informational field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInformational

`func (o *DiskUpdate1) SetInformational(v int32)`

SetInformational sets Informational field to given value.

### HasInformational

`func (o *DiskUpdate1) HasInformational() bool`

HasInformational returns a boolean if a field has been set.

### SetInformationalNil

`func (o *DiskUpdate1) SetInformationalNil(b bool)`

 SetInformationalNil sets the value for Informational to be an explicit nil

### UnsetInformational
`func (o *DiskUpdate1) UnsetInformational()`

UnsetInformational ensures that no value is present for Informational, not even an explicit nil
### GetEnclosure

`func (o *DiskUpdate1) GetEnclosure() DiskUpdate1Enclosure`

GetEnclosure returns the Enclosure field if non-nil, zero value otherwise.

### GetEnclosureOk

`func (o *DiskUpdate1) GetEnclosureOk() (*DiskUpdate1Enclosure, bool)`

GetEnclosureOk returns a tuple with the Enclosure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnclosure

`func (o *DiskUpdate1) SetEnclosure(v DiskUpdate1Enclosure)`

SetEnclosure sets Enclosure field to given value.

### HasEnclosure

`func (o *DiskUpdate1) HasEnclosure() bool`

HasEnclosure returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


