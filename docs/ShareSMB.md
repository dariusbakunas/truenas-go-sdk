# ShareSMB

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Path** | **string** |  | 
**PathSuffix** | Pointer to **string** |  | [optional] 
**Purpose** | Pointer to **string** |  | [optional] 
**Home** | Pointer to **bool** |  | [optional] 
**Timemachine** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 
**Ro** | Pointer to **bool** |  | [optional] 
**Browsable** | Pointer to **bool** |  | [optional] 
**Recyclebin** | Pointer to **bool** |  | [optional] 
**Shadowcopy** | Pointer to **bool** |  | [optional] 
**Guestok** | Pointer to **bool** |  | [optional] 
**Abe** | Pointer to **bool** |  | [optional] 
**Hostsallow** | Pointer to **[]string** |  | [optional] 
**Hostsdeny** | Pointer to **[]string** |  | [optional] 
**AaplNameMangling** | Pointer to **bool** |  | [optional] 
**Acl** | Pointer to **bool** |  | [optional] 
**Durablehandle** | Pointer to **bool** |  | [optional] 
**Streams** | Pointer to **bool** |  | [optional] 
**Fsrvp** | Pointer to **bool** |  | [optional] 
**Auxsmbconf** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Locked** | Pointer to **bool** |  | [optional] 
**Vuid** | Pointer to **string** |  | [optional] 

## Methods

### NewShareSMB

`func NewShareSMB(id int32, path string, ) *ShareSMB`

NewShareSMB instantiates a new ShareSMB object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShareSMBWithDefaults

`func NewShareSMBWithDefaults() *ShareSMB`

NewShareSMBWithDefaults instantiates a new ShareSMB object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ShareSMB) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ShareSMB) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ShareSMB) SetId(v int32)`

SetId sets Id field to given value.


### GetPath

`func (o *ShareSMB) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ShareSMB) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ShareSMB) SetPath(v string)`

SetPath sets Path field to given value.


### GetPathSuffix

`func (o *ShareSMB) GetPathSuffix() string`

GetPathSuffix returns the PathSuffix field if non-nil, zero value otherwise.

### GetPathSuffixOk

`func (o *ShareSMB) GetPathSuffixOk() (*string, bool)`

GetPathSuffixOk returns a tuple with the PathSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathSuffix

`func (o *ShareSMB) SetPathSuffix(v string)`

SetPathSuffix sets PathSuffix field to given value.

### HasPathSuffix

`func (o *ShareSMB) HasPathSuffix() bool`

HasPathSuffix returns a boolean if a field has been set.

### GetPurpose

`func (o *ShareSMB) GetPurpose() string`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *ShareSMB) GetPurposeOk() (*string, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *ShareSMB) SetPurpose(v string)`

SetPurpose sets Purpose field to given value.

### HasPurpose

`func (o *ShareSMB) HasPurpose() bool`

HasPurpose returns a boolean if a field has been set.

### GetHome

`func (o *ShareSMB) GetHome() bool`

GetHome returns the Home field if non-nil, zero value otherwise.

### GetHomeOk

`func (o *ShareSMB) GetHomeOk() (*bool, bool)`

GetHomeOk returns a tuple with the Home field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHome

`func (o *ShareSMB) SetHome(v bool)`

SetHome sets Home field to given value.

### HasHome

`func (o *ShareSMB) HasHome() bool`

HasHome returns a boolean if a field has been set.

### GetTimemachine

`func (o *ShareSMB) GetTimemachine() bool`

GetTimemachine returns the Timemachine field if non-nil, zero value otherwise.

### GetTimemachineOk

`func (o *ShareSMB) GetTimemachineOk() (*bool, bool)`

GetTimemachineOk returns a tuple with the Timemachine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimemachine

`func (o *ShareSMB) SetTimemachine(v bool)`

SetTimemachine sets Timemachine field to given value.

### HasTimemachine

`func (o *ShareSMB) HasTimemachine() bool`

HasTimemachine returns a boolean if a field has been set.

### GetName

`func (o *ShareSMB) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ShareSMB) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ShareSMB) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ShareSMB) HasName() bool`

HasName returns a boolean if a field has been set.

### GetComment

`func (o *ShareSMB) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ShareSMB) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ShareSMB) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ShareSMB) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetRo

`func (o *ShareSMB) GetRo() bool`

GetRo returns the Ro field if non-nil, zero value otherwise.

### GetRoOk

`func (o *ShareSMB) GetRoOk() (*bool, bool)`

GetRoOk returns a tuple with the Ro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRo

`func (o *ShareSMB) SetRo(v bool)`

SetRo sets Ro field to given value.

### HasRo

`func (o *ShareSMB) HasRo() bool`

HasRo returns a boolean if a field has been set.

### GetBrowsable

`func (o *ShareSMB) GetBrowsable() bool`

GetBrowsable returns the Browsable field if non-nil, zero value otherwise.

### GetBrowsableOk

`func (o *ShareSMB) GetBrowsableOk() (*bool, bool)`

GetBrowsableOk returns a tuple with the Browsable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowsable

`func (o *ShareSMB) SetBrowsable(v bool)`

SetBrowsable sets Browsable field to given value.

### HasBrowsable

`func (o *ShareSMB) HasBrowsable() bool`

HasBrowsable returns a boolean if a field has been set.

### GetRecyclebin

`func (o *ShareSMB) GetRecyclebin() bool`

GetRecyclebin returns the Recyclebin field if non-nil, zero value otherwise.

### GetRecyclebinOk

`func (o *ShareSMB) GetRecyclebinOk() (*bool, bool)`

GetRecyclebinOk returns a tuple with the Recyclebin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecyclebin

`func (o *ShareSMB) SetRecyclebin(v bool)`

SetRecyclebin sets Recyclebin field to given value.

### HasRecyclebin

`func (o *ShareSMB) HasRecyclebin() bool`

HasRecyclebin returns a boolean if a field has been set.

### GetShadowcopy

`func (o *ShareSMB) GetShadowcopy() bool`

GetShadowcopy returns the Shadowcopy field if non-nil, zero value otherwise.

### GetShadowcopyOk

`func (o *ShareSMB) GetShadowcopyOk() (*bool, bool)`

GetShadowcopyOk returns a tuple with the Shadowcopy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShadowcopy

`func (o *ShareSMB) SetShadowcopy(v bool)`

SetShadowcopy sets Shadowcopy field to given value.

### HasShadowcopy

`func (o *ShareSMB) HasShadowcopy() bool`

HasShadowcopy returns a boolean if a field has been set.

### GetGuestok

`func (o *ShareSMB) GetGuestok() bool`

GetGuestok returns the Guestok field if non-nil, zero value otherwise.

### GetGuestokOk

`func (o *ShareSMB) GetGuestokOk() (*bool, bool)`

GetGuestokOk returns a tuple with the Guestok field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestok

`func (o *ShareSMB) SetGuestok(v bool)`

SetGuestok sets Guestok field to given value.

### HasGuestok

`func (o *ShareSMB) HasGuestok() bool`

HasGuestok returns a boolean if a field has been set.

### GetAbe

`func (o *ShareSMB) GetAbe() bool`

GetAbe returns the Abe field if non-nil, zero value otherwise.

### GetAbeOk

`func (o *ShareSMB) GetAbeOk() (*bool, bool)`

GetAbeOk returns a tuple with the Abe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbe

`func (o *ShareSMB) SetAbe(v bool)`

SetAbe sets Abe field to given value.

### HasAbe

`func (o *ShareSMB) HasAbe() bool`

HasAbe returns a boolean if a field has been set.

### GetHostsallow

`func (o *ShareSMB) GetHostsallow() []string`

GetHostsallow returns the Hostsallow field if non-nil, zero value otherwise.

### GetHostsallowOk

`func (o *ShareSMB) GetHostsallowOk() (*[]string, bool)`

GetHostsallowOk returns a tuple with the Hostsallow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostsallow

`func (o *ShareSMB) SetHostsallow(v []string)`

SetHostsallow sets Hostsallow field to given value.

### HasHostsallow

`func (o *ShareSMB) HasHostsallow() bool`

HasHostsallow returns a boolean if a field has been set.

### GetHostsdeny

`func (o *ShareSMB) GetHostsdeny() []string`

GetHostsdeny returns the Hostsdeny field if non-nil, zero value otherwise.

### GetHostsdenyOk

`func (o *ShareSMB) GetHostsdenyOk() (*[]string, bool)`

GetHostsdenyOk returns a tuple with the Hostsdeny field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostsdeny

`func (o *ShareSMB) SetHostsdeny(v []string)`

SetHostsdeny sets Hostsdeny field to given value.

### HasHostsdeny

`func (o *ShareSMB) HasHostsdeny() bool`

HasHostsdeny returns a boolean if a field has been set.

### GetAaplNameMangling

`func (o *ShareSMB) GetAaplNameMangling() bool`

GetAaplNameMangling returns the AaplNameMangling field if non-nil, zero value otherwise.

### GetAaplNameManglingOk

`func (o *ShareSMB) GetAaplNameManglingOk() (*bool, bool)`

GetAaplNameManglingOk returns a tuple with the AaplNameMangling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAaplNameMangling

`func (o *ShareSMB) SetAaplNameMangling(v bool)`

SetAaplNameMangling sets AaplNameMangling field to given value.

### HasAaplNameMangling

`func (o *ShareSMB) HasAaplNameMangling() bool`

HasAaplNameMangling returns a boolean if a field has been set.

### GetAcl

`func (o *ShareSMB) GetAcl() bool`

GetAcl returns the Acl field if non-nil, zero value otherwise.

### GetAclOk

`func (o *ShareSMB) GetAclOk() (*bool, bool)`

GetAclOk returns a tuple with the Acl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcl

`func (o *ShareSMB) SetAcl(v bool)`

SetAcl sets Acl field to given value.

### HasAcl

`func (o *ShareSMB) HasAcl() bool`

HasAcl returns a boolean if a field has been set.

### GetDurablehandle

`func (o *ShareSMB) GetDurablehandle() bool`

GetDurablehandle returns the Durablehandle field if non-nil, zero value otherwise.

### GetDurablehandleOk

`func (o *ShareSMB) GetDurablehandleOk() (*bool, bool)`

GetDurablehandleOk returns a tuple with the Durablehandle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurablehandle

`func (o *ShareSMB) SetDurablehandle(v bool)`

SetDurablehandle sets Durablehandle field to given value.

### HasDurablehandle

`func (o *ShareSMB) HasDurablehandle() bool`

HasDurablehandle returns a boolean if a field has been set.

### GetStreams

`func (o *ShareSMB) GetStreams() bool`

GetStreams returns the Streams field if non-nil, zero value otherwise.

### GetStreamsOk

`func (o *ShareSMB) GetStreamsOk() (*bool, bool)`

GetStreamsOk returns a tuple with the Streams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreams

`func (o *ShareSMB) SetStreams(v bool)`

SetStreams sets Streams field to given value.

### HasStreams

`func (o *ShareSMB) HasStreams() bool`

HasStreams returns a boolean if a field has been set.

### GetFsrvp

`func (o *ShareSMB) GetFsrvp() bool`

GetFsrvp returns the Fsrvp field if non-nil, zero value otherwise.

### GetFsrvpOk

`func (o *ShareSMB) GetFsrvpOk() (*bool, bool)`

GetFsrvpOk returns a tuple with the Fsrvp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsrvp

`func (o *ShareSMB) SetFsrvp(v bool)`

SetFsrvp sets Fsrvp field to given value.

### HasFsrvp

`func (o *ShareSMB) HasFsrvp() bool`

HasFsrvp returns a boolean if a field has been set.

### GetAuxsmbconf

`func (o *ShareSMB) GetAuxsmbconf() string`

GetAuxsmbconf returns the Auxsmbconf field if non-nil, zero value otherwise.

### GetAuxsmbconfOk

`func (o *ShareSMB) GetAuxsmbconfOk() (*string, bool)`

GetAuxsmbconfOk returns a tuple with the Auxsmbconf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuxsmbconf

`func (o *ShareSMB) SetAuxsmbconf(v string)`

SetAuxsmbconf sets Auxsmbconf field to given value.

### HasAuxsmbconf

`func (o *ShareSMB) HasAuxsmbconf() bool`

HasAuxsmbconf returns a boolean if a field has been set.

### GetEnabled

`func (o *ShareSMB) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ShareSMB) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ShareSMB) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ShareSMB) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetLocked

`func (o *ShareSMB) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *ShareSMB) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *ShareSMB) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *ShareSMB) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetVuid

`func (o *ShareSMB) GetVuid() string`

GetVuid returns the Vuid field if non-nil, zero value otherwise.

### GetVuidOk

`func (o *ShareSMB) GetVuidOk() (*string, bool)`

GetVuidOk returns a tuple with the Vuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVuid

`func (o *ShareSMB) SetVuid(v string)`

SetVuid sets Vuid field to given value.

### HasVuid

`func (o *ShareSMB) HasVuid() bool`

HasVuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


