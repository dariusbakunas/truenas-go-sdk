# SharingSmbUpdate1

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

### NewSharingSmbUpdate1

`func NewSharingSmbUpdate1() *SharingSmbUpdate1`

NewSharingSmbUpdate1 instantiates a new SharingSmbUpdate1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSharingSmbUpdate1WithDefaults

`func NewSharingSmbUpdate1WithDefaults() *SharingSmbUpdate1`

NewSharingSmbUpdate1WithDefaults instantiates a new SharingSmbUpdate1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPurpose

`func (o *SharingSmbUpdate1) GetPurpose() string`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *SharingSmbUpdate1) GetPurposeOk() (*string, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *SharingSmbUpdate1) SetPurpose(v string)`

SetPurpose sets Purpose field to given value.

### HasPurpose

`func (o *SharingSmbUpdate1) HasPurpose() bool`

HasPurpose returns a boolean if a field has been set.

### GetPath

`func (o *SharingSmbUpdate1) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *SharingSmbUpdate1) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *SharingSmbUpdate1) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *SharingSmbUpdate1) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetPathSuffix

`func (o *SharingSmbUpdate1) GetPathSuffix() string`

GetPathSuffix returns the PathSuffix field if non-nil, zero value otherwise.

### GetPathSuffixOk

`func (o *SharingSmbUpdate1) GetPathSuffixOk() (*string, bool)`

GetPathSuffixOk returns a tuple with the PathSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathSuffix

`func (o *SharingSmbUpdate1) SetPathSuffix(v string)`

SetPathSuffix sets PathSuffix field to given value.

### HasPathSuffix

`func (o *SharingSmbUpdate1) HasPathSuffix() bool`

HasPathSuffix returns a boolean if a field has been set.

### GetHome

`func (o *SharingSmbUpdate1) GetHome() bool`

GetHome returns the Home field if non-nil, zero value otherwise.

### GetHomeOk

`func (o *SharingSmbUpdate1) GetHomeOk() (*bool, bool)`

GetHomeOk returns a tuple with the Home field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHome

`func (o *SharingSmbUpdate1) SetHome(v bool)`

SetHome sets Home field to given value.

### HasHome

`func (o *SharingSmbUpdate1) HasHome() bool`

HasHome returns a boolean if a field has been set.

### GetName

`func (o *SharingSmbUpdate1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SharingSmbUpdate1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SharingSmbUpdate1) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SharingSmbUpdate1) HasName() bool`

HasName returns a boolean if a field has been set.

### GetComment

`func (o *SharingSmbUpdate1) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *SharingSmbUpdate1) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *SharingSmbUpdate1) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *SharingSmbUpdate1) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetRo

`func (o *SharingSmbUpdate1) GetRo() bool`

GetRo returns the Ro field if non-nil, zero value otherwise.

### GetRoOk

`func (o *SharingSmbUpdate1) GetRoOk() (*bool, bool)`

GetRoOk returns a tuple with the Ro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRo

`func (o *SharingSmbUpdate1) SetRo(v bool)`

SetRo sets Ro field to given value.

### HasRo

`func (o *SharingSmbUpdate1) HasRo() bool`

HasRo returns a boolean if a field has been set.

### GetBrowsable

`func (o *SharingSmbUpdate1) GetBrowsable() bool`

GetBrowsable returns the Browsable field if non-nil, zero value otherwise.

### GetBrowsableOk

`func (o *SharingSmbUpdate1) GetBrowsableOk() (*bool, bool)`

GetBrowsableOk returns a tuple with the Browsable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowsable

`func (o *SharingSmbUpdate1) SetBrowsable(v bool)`

SetBrowsable sets Browsable field to given value.

### HasBrowsable

`func (o *SharingSmbUpdate1) HasBrowsable() bool`

HasBrowsable returns a boolean if a field has been set.

### GetTimemachine

`func (o *SharingSmbUpdate1) GetTimemachine() bool`

GetTimemachine returns the Timemachine field if non-nil, zero value otherwise.

### GetTimemachineOk

`func (o *SharingSmbUpdate1) GetTimemachineOk() (*bool, bool)`

GetTimemachineOk returns a tuple with the Timemachine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimemachine

`func (o *SharingSmbUpdate1) SetTimemachine(v bool)`

SetTimemachine sets Timemachine field to given value.

### HasTimemachine

`func (o *SharingSmbUpdate1) HasTimemachine() bool`

HasTimemachine returns a boolean if a field has been set.

### GetRecyclebin

`func (o *SharingSmbUpdate1) GetRecyclebin() bool`

GetRecyclebin returns the Recyclebin field if non-nil, zero value otherwise.

### GetRecyclebinOk

`func (o *SharingSmbUpdate1) GetRecyclebinOk() (*bool, bool)`

GetRecyclebinOk returns a tuple with the Recyclebin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecyclebin

`func (o *SharingSmbUpdate1) SetRecyclebin(v bool)`

SetRecyclebin sets Recyclebin field to given value.

### HasRecyclebin

`func (o *SharingSmbUpdate1) HasRecyclebin() bool`

HasRecyclebin returns a boolean if a field has been set.

### GetGuestok

`func (o *SharingSmbUpdate1) GetGuestok() bool`

GetGuestok returns the Guestok field if non-nil, zero value otherwise.

### GetGuestokOk

`func (o *SharingSmbUpdate1) GetGuestokOk() (*bool, bool)`

GetGuestokOk returns a tuple with the Guestok field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestok

`func (o *SharingSmbUpdate1) SetGuestok(v bool)`

SetGuestok sets Guestok field to given value.

### HasGuestok

`func (o *SharingSmbUpdate1) HasGuestok() bool`

HasGuestok returns a boolean if a field has been set.

### GetAbe

`func (o *SharingSmbUpdate1) GetAbe() bool`

GetAbe returns the Abe field if non-nil, zero value otherwise.

### GetAbeOk

`func (o *SharingSmbUpdate1) GetAbeOk() (*bool, bool)`

GetAbeOk returns a tuple with the Abe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbe

`func (o *SharingSmbUpdate1) SetAbe(v bool)`

SetAbe sets Abe field to given value.

### HasAbe

`func (o *SharingSmbUpdate1) HasAbe() bool`

HasAbe returns a boolean if a field has been set.

### GetHostsallow

`func (o *SharingSmbUpdate1) GetHostsallow() []interface{}`

GetHostsallow returns the Hostsallow field if non-nil, zero value otherwise.

### GetHostsallowOk

`func (o *SharingSmbUpdate1) GetHostsallowOk() (*[]interface{}, bool)`

GetHostsallowOk returns a tuple with the Hostsallow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostsallow

`func (o *SharingSmbUpdate1) SetHostsallow(v []interface{})`

SetHostsallow sets Hostsallow field to given value.

### HasHostsallow

`func (o *SharingSmbUpdate1) HasHostsallow() bool`

HasHostsallow returns a boolean if a field has been set.

### GetHostsdeny

`func (o *SharingSmbUpdate1) GetHostsdeny() []interface{}`

GetHostsdeny returns the Hostsdeny field if non-nil, zero value otherwise.

### GetHostsdenyOk

`func (o *SharingSmbUpdate1) GetHostsdenyOk() (*[]interface{}, bool)`

GetHostsdenyOk returns a tuple with the Hostsdeny field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostsdeny

`func (o *SharingSmbUpdate1) SetHostsdeny(v []interface{})`

SetHostsdeny sets Hostsdeny field to given value.

### HasHostsdeny

`func (o *SharingSmbUpdate1) HasHostsdeny() bool`

HasHostsdeny returns a boolean if a field has been set.

### GetAaplNameMangling

`func (o *SharingSmbUpdate1) GetAaplNameMangling() bool`

GetAaplNameMangling returns the AaplNameMangling field if non-nil, zero value otherwise.

### GetAaplNameManglingOk

`func (o *SharingSmbUpdate1) GetAaplNameManglingOk() (*bool, bool)`

GetAaplNameManglingOk returns a tuple with the AaplNameMangling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAaplNameMangling

`func (o *SharingSmbUpdate1) SetAaplNameMangling(v bool)`

SetAaplNameMangling sets AaplNameMangling field to given value.

### HasAaplNameMangling

`func (o *SharingSmbUpdate1) HasAaplNameMangling() bool`

HasAaplNameMangling returns a boolean if a field has been set.

### GetAcl

`func (o *SharingSmbUpdate1) GetAcl() bool`

GetAcl returns the Acl field if non-nil, zero value otherwise.

### GetAclOk

`func (o *SharingSmbUpdate1) GetAclOk() (*bool, bool)`

GetAclOk returns a tuple with the Acl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcl

`func (o *SharingSmbUpdate1) SetAcl(v bool)`

SetAcl sets Acl field to given value.

### HasAcl

`func (o *SharingSmbUpdate1) HasAcl() bool`

HasAcl returns a boolean if a field has been set.

### GetDurablehandle

`func (o *SharingSmbUpdate1) GetDurablehandle() bool`

GetDurablehandle returns the Durablehandle field if non-nil, zero value otherwise.

### GetDurablehandleOk

`func (o *SharingSmbUpdate1) GetDurablehandleOk() (*bool, bool)`

GetDurablehandleOk returns a tuple with the Durablehandle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurablehandle

`func (o *SharingSmbUpdate1) SetDurablehandle(v bool)`

SetDurablehandle sets Durablehandle field to given value.

### HasDurablehandle

`func (o *SharingSmbUpdate1) HasDurablehandle() bool`

HasDurablehandle returns a boolean if a field has been set.

### GetShadowcopy

`func (o *SharingSmbUpdate1) GetShadowcopy() bool`

GetShadowcopy returns the Shadowcopy field if non-nil, zero value otherwise.

### GetShadowcopyOk

`func (o *SharingSmbUpdate1) GetShadowcopyOk() (*bool, bool)`

GetShadowcopyOk returns a tuple with the Shadowcopy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShadowcopy

`func (o *SharingSmbUpdate1) SetShadowcopy(v bool)`

SetShadowcopy sets Shadowcopy field to given value.

### HasShadowcopy

`func (o *SharingSmbUpdate1) HasShadowcopy() bool`

HasShadowcopy returns a boolean if a field has been set.

### GetStreams

`func (o *SharingSmbUpdate1) GetStreams() bool`

GetStreams returns the Streams field if non-nil, zero value otherwise.

### GetStreamsOk

`func (o *SharingSmbUpdate1) GetStreamsOk() (*bool, bool)`

GetStreamsOk returns a tuple with the Streams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreams

`func (o *SharingSmbUpdate1) SetStreams(v bool)`

SetStreams sets Streams field to given value.

### HasStreams

`func (o *SharingSmbUpdate1) HasStreams() bool`

HasStreams returns a boolean if a field has been set.

### GetFsrvp

`func (o *SharingSmbUpdate1) GetFsrvp() bool`

GetFsrvp returns the Fsrvp field if non-nil, zero value otherwise.

### GetFsrvpOk

`func (o *SharingSmbUpdate1) GetFsrvpOk() (*bool, bool)`

GetFsrvpOk returns a tuple with the Fsrvp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsrvp

`func (o *SharingSmbUpdate1) SetFsrvp(v bool)`

SetFsrvp sets Fsrvp field to given value.

### HasFsrvp

`func (o *SharingSmbUpdate1) HasFsrvp() bool`

HasFsrvp returns a boolean if a field has been set.

### GetAuxsmbconf

`func (o *SharingSmbUpdate1) GetAuxsmbconf() string`

GetAuxsmbconf returns the Auxsmbconf field if non-nil, zero value otherwise.

### GetAuxsmbconfOk

`func (o *SharingSmbUpdate1) GetAuxsmbconfOk() (*string, bool)`

GetAuxsmbconfOk returns a tuple with the Auxsmbconf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuxsmbconf

`func (o *SharingSmbUpdate1) SetAuxsmbconf(v string)`

SetAuxsmbconf sets Auxsmbconf field to given value.

### HasAuxsmbconf

`func (o *SharingSmbUpdate1) HasAuxsmbconf() bool`

HasAuxsmbconf returns a boolean if a field has been set.

### GetEnabled

`func (o *SharingSmbUpdate1) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SharingSmbUpdate1) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SharingSmbUpdate1) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SharingSmbUpdate1) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


