# SharingAfpCreate0

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to **string** |  | [optional] 
**Home** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 
**Allow** | Pointer to **[]interface{}** |  | [optional] 
**Deny** | Pointer to **[]interface{}** |  | [optional] 
**Ro** | Pointer to **[]interface{}** |  | [optional] 
**Rw** | Pointer to **[]interface{}** |  | [optional] 
**Timemachine** | Pointer to **bool** |  | [optional] 
**TimemachineQuota** | Pointer to **int32** |  | [optional] 
**Nodev** | Pointer to **bool** |  | [optional] 
**Nostat** | Pointer to **bool** |  | [optional] 
**Upriv** | Pointer to **bool** |  | [optional] 
**Fperm** | Pointer to **string** |  | [optional] 
**Dperm** | Pointer to **string** |  | [optional] 
**Umask** | Pointer to **string** |  | [optional] 
**Hostsallow** | Pointer to **[]interface{}** |  | [optional] 
**Hostsdeny** | Pointer to **[]interface{}** |  | [optional] 
**Vuid** | Pointer to **NullableString** |  | [optional] 
**Auxparams** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewSharingAfpCreate0

`func NewSharingAfpCreate0() *SharingAfpCreate0`

NewSharingAfpCreate0 instantiates a new SharingAfpCreate0 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSharingAfpCreate0WithDefaults

`func NewSharingAfpCreate0WithDefaults() *SharingAfpCreate0`

NewSharingAfpCreate0WithDefaults instantiates a new SharingAfpCreate0 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *SharingAfpCreate0) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *SharingAfpCreate0) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *SharingAfpCreate0) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *SharingAfpCreate0) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetHome

`func (o *SharingAfpCreate0) GetHome() bool`

GetHome returns the Home field if non-nil, zero value otherwise.

### GetHomeOk

`func (o *SharingAfpCreate0) GetHomeOk() (*bool, bool)`

GetHomeOk returns a tuple with the Home field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHome

`func (o *SharingAfpCreate0) SetHome(v bool)`

SetHome sets Home field to given value.

### HasHome

`func (o *SharingAfpCreate0) HasHome() bool`

HasHome returns a boolean if a field has been set.

### GetName

`func (o *SharingAfpCreate0) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SharingAfpCreate0) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SharingAfpCreate0) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SharingAfpCreate0) HasName() bool`

HasName returns a boolean if a field has been set.

### GetComment

`func (o *SharingAfpCreate0) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *SharingAfpCreate0) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *SharingAfpCreate0) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *SharingAfpCreate0) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetAllow

`func (o *SharingAfpCreate0) GetAllow() []interface{}`

GetAllow returns the Allow field if non-nil, zero value otherwise.

### GetAllowOk

`func (o *SharingAfpCreate0) GetAllowOk() (*[]interface{}, bool)`

GetAllowOk returns a tuple with the Allow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllow

`func (o *SharingAfpCreate0) SetAllow(v []interface{})`

SetAllow sets Allow field to given value.

### HasAllow

`func (o *SharingAfpCreate0) HasAllow() bool`

HasAllow returns a boolean if a field has been set.

### GetDeny

`func (o *SharingAfpCreate0) GetDeny() []interface{}`

GetDeny returns the Deny field if non-nil, zero value otherwise.

### GetDenyOk

`func (o *SharingAfpCreate0) GetDenyOk() (*[]interface{}, bool)`

GetDenyOk returns a tuple with the Deny field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeny

`func (o *SharingAfpCreate0) SetDeny(v []interface{})`

SetDeny sets Deny field to given value.

### HasDeny

`func (o *SharingAfpCreate0) HasDeny() bool`

HasDeny returns a boolean if a field has been set.

### GetRo

`func (o *SharingAfpCreate0) GetRo() []interface{}`

GetRo returns the Ro field if non-nil, zero value otherwise.

### GetRoOk

`func (o *SharingAfpCreate0) GetRoOk() (*[]interface{}, bool)`

GetRoOk returns a tuple with the Ro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRo

`func (o *SharingAfpCreate0) SetRo(v []interface{})`

SetRo sets Ro field to given value.

### HasRo

`func (o *SharingAfpCreate0) HasRo() bool`

HasRo returns a boolean if a field has been set.

### GetRw

`func (o *SharingAfpCreate0) GetRw() []interface{}`

GetRw returns the Rw field if non-nil, zero value otherwise.

### GetRwOk

`func (o *SharingAfpCreate0) GetRwOk() (*[]interface{}, bool)`

GetRwOk returns a tuple with the Rw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRw

`func (o *SharingAfpCreate0) SetRw(v []interface{})`

SetRw sets Rw field to given value.

### HasRw

`func (o *SharingAfpCreate0) HasRw() bool`

HasRw returns a boolean if a field has been set.

### GetTimemachine

`func (o *SharingAfpCreate0) GetTimemachine() bool`

GetTimemachine returns the Timemachine field if non-nil, zero value otherwise.

### GetTimemachineOk

`func (o *SharingAfpCreate0) GetTimemachineOk() (*bool, bool)`

GetTimemachineOk returns a tuple with the Timemachine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimemachine

`func (o *SharingAfpCreate0) SetTimemachine(v bool)`

SetTimemachine sets Timemachine field to given value.

### HasTimemachine

`func (o *SharingAfpCreate0) HasTimemachine() bool`

HasTimemachine returns a boolean if a field has been set.

### GetTimemachineQuota

`func (o *SharingAfpCreate0) GetTimemachineQuota() int32`

GetTimemachineQuota returns the TimemachineQuota field if non-nil, zero value otherwise.

### GetTimemachineQuotaOk

`func (o *SharingAfpCreate0) GetTimemachineQuotaOk() (*int32, bool)`

GetTimemachineQuotaOk returns a tuple with the TimemachineQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimemachineQuota

`func (o *SharingAfpCreate0) SetTimemachineQuota(v int32)`

SetTimemachineQuota sets TimemachineQuota field to given value.

### HasTimemachineQuota

`func (o *SharingAfpCreate0) HasTimemachineQuota() bool`

HasTimemachineQuota returns a boolean if a field has been set.

### GetNodev

`func (o *SharingAfpCreate0) GetNodev() bool`

GetNodev returns the Nodev field if non-nil, zero value otherwise.

### GetNodevOk

`func (o *SharingAfpCreate0) GetNodevOk() (*bool, bool)`

GetNodevOk returns a tuple with the Nodev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodev

`func (o *SharingAfpCreate0) SetNodev(v bool)`

SetNodev sets Nodev field to given value.

### HasNodev

`func (o *SharingAfpCreate0) HasNodev() bool`

HasNodev returns a boolean if a field has been set.

### GetNostat

`func (o *SharingAfpCreate0) GetNostat() bool`

GetNostat returns the Nostat field if non-nil, zero value otherwise.

### GetNostatOk

`func (o *SharingAfpCreate0) GetNostatOk() (*bool, bool)`

GetNostatOk returns a tuple with the Nostat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNostat

`func (o *SharingAfpCreate0) SetNostat(v bool)`

SetNostat sets Nostat field to given value.

### HasNostat

`func (o *SharingAfpCreate0) HasNostat() bool`

HasNostat returns a boolean if a field has been set.

### GetUpriv

`func (o *SharingAfpCreate0) GetUpriv() bool`

GetUpriv returns the Upriv field if non-nil, zero value otherwise.

### GetUprivOk

`func (o *SharingAfpCreate0) GetUprivOk() (*bool, bool)`

GetUprivOk returns a tuple with the Upriv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpriv

`func (o *SharingAfpCreate0) SetUpriv(v bool)`

SetUpriv sets Upriv field to given value.

### HasUpriv

`func (o *SharingAfpCreate0) HasUpriv() bool`

HasUpriv returns a boolean if a field has been set.

### GetFperm

`func (o *SharingAfpCreate0) GetFperm() string`

GetFperm returns the Fperm field if non-nil, zero value otherwise.

### GetFpermOk

`func (o *SharingAfpCreate0) GetFpermOk() (*string, bool)`

GetFpermOk returns a tuple with the Fperm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFperm

`func (o *SharingAfpCreate0) SetFperm(v string)`

SetFperm sets Fperm field to given value.

### HasFperm

`func (o *SharingAfpCreate0) HasFperm() bool`

HasFperm returns a boolean if a field has been set.

### GetDperm

`func (o *SharingAfpCreate0) GetDperm() string`

GetDperm returns the Dperm field if non-nil, zero value otherwise.

### GetDpermOk

`func (o *SharingAfpCreate0) GetDpermOk() (*string, bool)`

GetDpermOk returns a tuple with the Dperm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDperm

`func (o *SharingAfpCreate0) SetDperm(v string)`

SetDperm sets Dperm field to given value.

### HasDperm

`func (o *SharingAfpCreate0) HasDperm() bool`

HasDperm returns a boolean if a field has been set.

### GetUmask

`func (o *SharingAfpCreate0) GetUmask() string`

GetUmask returns the Umask field if non-nil, zero value otherwise.

### GetUmaskOk

`func (o *SharingAfpCreate0) GetUmaskOk() (*string, bool)`

GetUmaskOk returns a tuple with the Umask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUmask

`func (o *SharingAfpCreate0) SetUmask(v string)`

SetUmask sets Umask field to given value.

### HasUmask

`func (o *SharingAfpCreate0) HasUmask() bool`

HasUmask returns a boolean if a field has been set.

### GetHostsallow

`func (o *SharingAfpCreate0) GetHostsallow() []interface{}`

GetHostsallow returns the Hostsallow field if non-nil, zero value otherwise.

### GetHostsallowOk

`func (o *SharingAfpCreate0) GetHostsallowOk() (*[]interface{}, bool)`

GetHostsallowOk returns a tuple with the Hostsallow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostsallow

`func (o *SharingAfpCreate0) SetHostsallow(v []interface{})`

SetHostsallow sets Hostsallow field to given value.

### HasHostsallow

`func (o *SharingAfpCreate0) HasHostsallow() bool`

HasHostsallow returns a boolean if a field has been set.

### GetHostsdeny

`func (o *SharingAfpCreate0) GetHostsdeny() []interface{}`

GetHostsdeny returns the Hostsdeny field if non-nil, zero value otherwise.

### GetHostsdenyOk

`func (o *SharingAfpCreate0) GetHostsdenyOk() (*[]interface{}, bool)`

GetHostsdenyOk returns a tuple with the Hostsdeny field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostsdeny

`func (o *SharingAfpCreate0) SetHostsdeny(v []interface{})`

SetHostsdeny sets Hostsdeny field to given value.

### HasHostsdeny

`func (o *SharingAfpCreate0) HasHostsdeny() bool`

HasHostsdeny returns a boolean if a field has been set.

### GetVuid

`func (o *SharingAfpCreate0) GetVuid() string`

GetVuid returns the Vuid field if non-nil, zero value otherwise.

### GetVuidOk

`func (o *SharingAfpCreate0) GetVuidOk() (*string, bool)`

GetVuidOk returns a tuple with the Vuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVuid

`func (o *SharingAfpCreate0) SetVuid(v string)`

SetVuid sets Vuid field to given value.

### HasVuid

`func (o *SharingAfpCreate0) HasVuid() bool`

HasVuid returns a boolean if a field has been set.

### SetVuidNil

`func (o *SharingAfpCreate0) SetVuidNil(b bool)`

 SetVuidNil sets the value for Vuid to be an explicit nil

### UnsetVuid
`func (o *SharingAfpCreate0) UnsetVuid()`

UnsetVuid ensures that no value is present for Vuid, not even an explicit nil
### GetAuxparams

`func (o *SharingAfpCreate0) GetAuxparams() string`

GetAuxparams returns the Auxparams field if non-nil, zero value otherwise.

### GetAuxparamsOk

`func (o *SharingAfpCreate0) GetAuxparamsOk() (*string, bool)`

GetAuxparamsOk returns a tuple with the Auxparams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuxparams

`func (o *SharingAfpCreate0) SetAuxparams(v string)`

SetAuxparams sets Auxparams field to given value.

### HasAuxparams

`func (o *SharingAfpCreate0) HasAuxparams() bool`

HasAuxparams returns a boolean if a field has been set.

### GetEnabled

`func (o *SharingAfpCreate0) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SharingAfpCreate0) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SharingAfpCreate0) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SharingAfpCreate0) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


