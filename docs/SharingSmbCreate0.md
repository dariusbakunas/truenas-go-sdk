# SharingSmbCreate0

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Purpose** | Pointer to **string** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**PathSuffix** | Pointer to **string** |  | [optional] 
**Home** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 
**Ro** | Pointer to **bool** |  | [optional] 
**Browsable** | Pointer to **bool** |  | [optional] 
**Timemachine** | Pointer to **bool** |  | [optional] 
**Recyclebin** | Pointer to **bool** |  | [optional] 
**Guestok** | Pointer to **bool** |  | [optional] 
**Abe** | Pointer to **bool** |  | [optional] 
**Hostsallow** | Pointer to **[]interface{}** |  | [optional] 
**Hostsdeny** | Pointer to **[]interface{}** |  | [optional] 
**AaplNameMangling** | Pointer to **bool** |  | [optional] 
**Acl** | Pointer to **bool** |  | [optional] 
**Durablehandle** | Pointer to **bool** |  | [optional] 
**Shadowcopy** | Pointer to **bool** |  | [optional] 
**Streams** | Pointer to **bool** |  | [optional] 
**Fsrvp** | Pointer to **bool** |  | [optional] 
**Auxsmbconf** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewSharingSmbCreate0

`func NewSharingSmbCreate0() *SharingSmbCreate0`

NewSharingSmbCreate0 instantiates a new SharingSmbCreate0 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSharingSmbCreate0WithDefaults

`func NewSharingSmbCreate0WithDefaults() *SharingSmbCreate0`

NewSharingSmbCreate0WithDefaults instantiates a new SharingSmbCreate0 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPurpose

`func (o *SharingSmbCreate0) GetPurpose() string`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *SharingSmbCreate0) GetPurposeOk() (*string, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *SharingSmbCreate0) SetPurpose(v string)`

SetPurpose sets Purpose field to given value.

### HasPurpose

`func (o *SharingSmbCreate0) HasPurpose() bool`

HasPurpose returns a boolean if a field has been set.

### GetPath

`func (o *SharingSmbCreate0) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *SharingSmbCreate0) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *SharingSmbCreate0) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *SharingSmbCreate0) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetPathSuffix

`func (o *SharingSmbCreate0) GetPathSuffix() string`

GetPathSuffix returns the PathSuffix field if non-nil, zero value otherwise.

### GetPathSuffixOk

`func (o *SharingSmbCreate0) GetPathSuffixOk() (*string, bool)`

GetPathSuffixOk returns a tuple with the PathSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathSuffix

`func (o *SharingSmbCreate0) SetPathSuffix(v string)`

SetPathSuffix sets PathSuffix field to given value.

### HasPathSuffix

`func (o *SharingSmbCreate0) HasPathSuffix() bool`

HasPathSuffix returns a boolean if a field has been set.

### GetHome

`func (o *SharingSmbCreate0) GetHome() bool`

GetHome returns the Home field if non-nil, zero value otherwise.

### GetHomeOk

`func (o *SharingSmbCreate0) GetHomeOk() (*bool, bool)`

GetHomeOk returns a tuple with the Home field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHome

`func (o *SharingSmbCreate0) SetHome(v bool)`

SetHome sets Home field to given value.

### HasHome

`func (o *SharingSmbCreate0) HasHome() bool`

HasHome returns a boolean if a field has been set.

### GetName

`func (o *SharingSmbCreate0) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SharingSmbCreate0) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SharingSmbCreate0) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SharingSmbCreate0) HasName() bool`

HasName returns a boolean if a field has been set.

### GetComment

`func (o *SharingSmbCreate0) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *SharingSmbCreate0) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *SharingSmbCreate0) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *SharingSmbCreate0) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetRo

`func (o *SharingSmbCreate0) GetRo() bool`

GetRo returns the Ro field if non-nil, zero value otherwise.

### GetRoOk

`func (o *SharingSmbCreate0) GetRoOk() (*bool, bool)`

GetRoOk returns a tuple with the Ro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRo

`func (o *SharingSmbCreate0) SetRo(v bool)`

SetRo sets Ro field to given value.

### HasRo

`func (o *SharingSmbCreate0) HasRo() bool`

HasRo returns a boolean if a field has been set.

### GetBrowsable

`func (o *SharingSmbCreate0) GetBrowsable() bool`

GetBrowsable returns the Browsable field if non-nil, zero value otherwise.

### GetBrowsableOk

`func (o *SharingSmbCreate0) GetBrowsableOk() (*bool, bool)`

GetBrowsableOk returns a tuple with the Browsable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowsable

`func (o *SharingSmbCreate0) SetBrowsable(v bool)`

SetBrowsable sets Browsable field to given value.

### HasBrowsable

`func (o *SharingSmbCreate0) HasBrowsable() bool`

HasBrowsable returns a boolean if a field has been set.

### GetTimemachine

`func (o *SharingSmbCreate0) GetTimemachine() bool`

GetTimemachine returns the Timemachine field if non-nil, zero value otherwise.

### GetTimemachineOk

`func (o *SharingSmbCreate0) GetTimemachineOk() (*bool, bool)`

GetTimemachineOk returns a tuple with the Timemachine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimemachine

`func (o *SharingSmbCreate0) SetTimemachine(v bool)`

SetTimemachine sets Timemachine field to given value.

### HasTimemachine

`func (o *SharingSmbCreate0) HasTimemachine() bool`

HasTimemachine returns a boolean if a field has been set.

### GetRecyclebin

`func (o *SharingSmbCreate0) GetRecyclebin() bool`

GetRecyclebin returns the Recyclebin field if non-nil, zero value otherwise.

### GetRecyclebinOk

`func (o *SharingSmbCreate0) GetRecyclebinOk() (*bool, bool)`

GetRecyclebinOk returns a tuple with the Recyclebin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecyclebin

`func (o *SharingSmbCreate0) SetRecyclebin(v bool)`

SetRecyclebin sets Recyclebin field to given value.

### HasRecyclebin

`func (o *SharingSmbCreate0) HasRecyclebin() bool`

HasRecyclebin returns a boolean if a field has been set.

### GetGuestok

`func (o *SharingSmbCreate0) GetGuestok() bool`

GetGuestok returns the Guestok field if non-nil, zero value otherwise.

### GetGuestokOk

`func (o *SharingSmbCreate0) GetGuestokOk() (*bool, bool)`

GetGuestokOk returns a tuple with the Guestok field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestok

`func (o *SharingSmbCreate0) SetGuestok(v bool)`

SetGuestok sets Guestok field to given value.

### HasGuestok

`func (o *SharingSmbCreate0) HasGuestok() bool`

HasGuestok returns a boolean if a field has been set.

### GetAbe

`func (o *SharingSmbCreate0) GetAbe() bool`

GetAbe returns the Abe field if non-nil, zero value otherwise.

### GetAbeOk

`func (o *SharingSmbCreate0) GetAbeOk() (*bool, bool)`

GetAbeOk returns a tuple with the Abe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbe

`func (o *SharingSmbCreate0) SetAbe(v bool)`

SetAbe sets Abe field to given value.

### HasAbe

`func (o *SharingSmbCreate0) HasAbe() bool`

HasAbe returns a boolean if a field has been set.

### GetHostsallow

`func (o *SharingSmbCreate0) GetHostsallow() []interface{}`

GetHostsallow returns the Hostsallow field if non-nil, zero value otherwise.

### GetHostsallowOk

`func (o *SharingSmbCreate0) GetHostsallowOk() (*[]interface{}, bool)`

GetHostsallowOk returns a tuple with the Hostsallow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostsallow

`func (o *SharingSmbCreate0) SetHostsallow(v []interface{})`

SetHostsallow sets Hostsallow field to given value.

### HasHostsallow

`func (o *SharingSmbCreate0) HasHostsallow() bool`

HasHostsallow returns a boolean if a field has been set.

### GetHostsdeny

`func (o *SharingSmbCreate0) GetHostsdeny() []interface{}`

GetHostsdeny returns the Hostsdeny field if non-nil, zero value otherwise.

### GetHostsdenyOk

`func (o *SharingSmbCreate0) GetHostsdenyOk() (*[]interface{}, bool)`

GetHostsdenyOk returns a tuple with the Hostsdeny field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostsdeny

`func (o *SharingSmbCreate0) SetHostsdeny(v []interface{})`

SetHostsdeny sets Hostsdeny field to given value.

### HasHostsdeny

`func (o *SharingSmbCreate0) HasHostsdeny() bool`

HasHostsdeny returns a boolean if a field has been set.

### GetAaplNameMangling

`func (o *SharingSmbCreate0) GetAaplNameMangling() bool`

GetAaplNameMangling returns the AaplNameMangling field if non-nil, zero value otherwise.

### GetAaplNameManglingOk

`func (o *SharingSmbCreate0) GetAaplNameManglingOk() (*bool, bool)`

GetAaplNameManglingOk returns a tuple with the AaplNameMangling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAaplNameMangling

`func (o *SharingSmbCreate0) SetAaplNameMangling(v bool)`

SetAaplNameMangling sets AaplNameMangling field to given value.

### HasAaplNameMangling

`func (o *SharingSmbCreate0) HasAaplNameMangling() bool`

HasAaplNameMangling returns a boolean if a field has been set.

### GetAcl

`func (o *SharingSmbCreate0) GetAcl() bool`

GetAcl returns the Acl field if non-nil, zero value otherwise.

### GetAclOk

`func (o *SharingSmbCreate0) GetAclOk() (*bool, bool)`

GetAclOk returns a tuple with the Acl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcl

`func (o *SharingSmbCreate0) SetAcl(v bool)`

SetAcl sets Acl field to given value.

### HasAcl

`func (o *SharingSmbCreate0) HasAcl() bool`

HasAcl returns a boolean if a field has been set.

### GetDurablehandle

`func (o *SharingSmbCreate0) GetDurablehandle() bool`

GetDurablehandle returns the Durablehandle field if non-nil, zero value otherwise.

### GetDurablehandleOk

`func (o *SharingSmbCreate0) GetDurablehandleOk() (*bool, bool)`

GetDurablehandleOk returns a tuple with the Durablehandle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurablehandle

`func (o *SharingSmbCreate0) SetDurablehandle(v bool)`

SetDurablehandle sets Durablehandle field to given value.

### HasDurablehandle

`func (o *SharingSmbCreate0) HasDurablehandle() bool`

HasDurablehandle returns a boolean if a field has been set.

### GetShadowcopy

`func (o *SharingSmbCreate0) GetShadowcopy() bool`

GetShadowcopy returns the Shadowcopy field if non-nil, zero value otherwise.

### GetShadowcopyOk

`func (o *SharingSmbCreate0) GetShadowcopyOk() (*bool, bool)`

GetShadowcopyOk returns a tuple with the Shadowcopy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShadowcopy

`func (o *SharingSmbCreate0) SetShadowcopy(v bool)`

SetShadowcopy sets Shadowcopy field to given value.

### HasShadowcopy

`func (o *SharingSmbCreate0) HasShadowcopy() bool`

HasShadowcopy returns a boolean if a field has been set.

### GetStreams

`func (o *SharingSmbCreate0) GetStreams() bool`

GetStreams returns the Streams field if non-nil, zero value otherwise.

### GetStreamsOk

`func (o *SharingSmbCreate0) GetStreamsOk() (*bool, bool)`

GetStreamsOk returns a tuple with the Streams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreams

`func (o *SharingSmbCreate0) SetStreams(v bool)`

SetStreams sets Streams field to given value.

### HasStreams

`func (o *SharingSmbCreate0) HasStreams() bool`

HasStreams returns a boolean if a field has been set.

### GetFsrvp

`func (o *SharingSmbCreate0) GetFsrvp() bool`

GetFsrvp returns the Fsrvp field if non-nil, zero value otherwise.

### GetFsrvpOk

`func (o *SharingSmbCreate0) GetFsrvpOk() (*bool, bool)`

GetFsrvpOk returns a tuple with the Fsrvp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsrvp

`func (o *SharingSmbCreate0) SetFsrvp(v bool)`

SetFsrvp sets Fsrvp field to given value.

### HasFsrvp

`func (o *SharingSmbCreate0) HasFsrvp() bool`

HasFsrvp returns a boolean if a field has been set.

### GetAuxsmbconf

`func (o *SharingSmbCreate0) GetAuxsmbconf() string`

GetAuxsmbconf returns the Auxsmbconf field if non-nil, zero value otherwise.

### GetAuxsmbconfOk

`func (o *SharingSmbCreate0) GetAuxsmbconfOk() (*string, bool)`

GetAuxsmbconfOk returns a tuple with the Auxsmbconf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuxsmbconf

`func (o *SharingSmbCreate0) SetAuxsmbconf(v string)`

SetAuxsmbconf sets Auxsmbconf field to given value.

### HasAuxsmbconf

`func (o *SharingSmbCreate0) HasAuxsmbconf() bool`

HasAuxsmbconf returns a boolean if a field has been set.

### GetEnabled

`func (o *SharingSmbCreate0) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SharingSmbCreate0) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SharingSmbCreate0) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SharingSmbCreate0) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


