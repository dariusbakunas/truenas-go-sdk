# CreateShareSMBParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Purpose** | Pointer to **string** |  | [optional] 
**Path** | **string** |  | 
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
**Hostsallow** | Pointer to **[]string** |  | [optional] 
**Hostsdeny** | Pointer to **[]string** |  | [optional] 
**AaplNameMangling** | Pointer to **bool** |  | [optional] 
**Acl** | Pointer to **bool** |  | [optional] 
**Durablehandle** | Pointer to **bool** |  | [optional] 
**Shadowcopy** | Pointer to **bool** |  | [optional] 
**Streams** | Pointer to **bool** |  | [optional] 
**Fsrvp** | Pointer to **bool** |  | [optional] 
**Auxsmbconf** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewCreateShareSMBParams

`func NewCreateShareSMBParams(path string, ) *CreateShareSMBParams`

NewCreateShareSMBParams instantiates a new CreateShareSMBParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateShareSMBParamsWithDefaults

`func NewCreateShareSMBParamsWithDefaults() *CreateShareSMBParams`

NewCreateShareSMBParamsWithDefaults instantiates a new CreateShareSMBParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPurpose

`func (o *CreateShareSMBParams) GetPurpose() string`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *CreateShareSMBParams) GetPurposeOk() (*string, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *CreateShareSMBParams) SetPurpose(v string)`

SetPurpose sets Purpose field to given value.

### HasPurpose

`func (o *CreateShareSMBParams) HasPurpose() bool`

HasPurpose returns a boolean if a field has been set.

### GetPath

`func (o *CreateShareSMBParams) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *CreateShareSMBParams) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *CreateShareSMBParams) SetPath(v string)`

SetPath sets Path field to given value.


### GetPathSuffix

`func (o *CreateShareSMBParams) GetPathSuffix() string`

GetPathSuffix returns the PathSuffix field if non-nil, zero value otherwise.

### GetPathSuffixOk

`func (o *CreateShareSMBParams) GetPathSuffixOk() (*string, bool)`

GetPathSuffixOk returns a tuple with the PathSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathSuffix

`func (o *CreateShareSMBParams) SetPathSuffix(v string)`

SetPathSuffix sets PathSuffix field to given value.

### HasPathSuffix

`func (o *CreateShareSMBParams) HasPathSuffix() bool`

HasPathSuffix returns a boolean if a field has been set.

### GetHome

`func (o *CreateShareSMBParams) GetHome() bool`

GetHome returns the Home field if non-nil, zero value otherwise.

### GetHomeOk

`func (o *CreateShareSMBParams) GetHomeOk() (*bool, bool)`

GetHomeOk returns a tuple with the Home field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHome

`func (o *CreateShareSMBParams) SetHome(v bool)`

SetHome sets Home field to given value.

### HasHome

`func (o *CreateShareSMBParams) HasHome() bool`

HasHome returns a boolean if a field has been set.

### GetName

`func (o *CreateShareSMBParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateShareSMBParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateShareSMBParams) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateShareSMBParams) HasName() bool`

HasName returns a boolean if a field has been set.

### GetComment

`func (o *CreateShareSMBParams) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *CreateShareSMBParams) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *CreateShareSMBParams) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *CreateShareSMBParams) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetRo

`func (o *CreateShareSMBParams) GetRo() bool`

GetRo returns the Ro field if non-nil, zero value otherwise.

### GetRoOk

`func (o *CreateShareSMBParams) GetRoOk() (*bool, bool)`

GetRoOk returns a tuple with the Ro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRo

`func (o *CreateShareSMBParams) SetRo(v bool)`

SetRo sets Ro field to given value.

### HasRo

`func (o *CreateShareSMBParams) HasRo() bool`

HasRo returns a boolean if a field has been set.

### GetBrowsable

`func (o *CreateShareSMBParams) GetBrowsable() bool`

GetBrowsable returns the Browsable field if non-nil, zero value otherwise.

### GetBrowsableOk

`func (o *CreateShareSMBParams) GetBrowsableOk() (*bool, bool)`

GetBrowsableOk returns a tuple with the Browsable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowsable

`func (o *CreateShareSMBParams) SetBrowsable(v bool)`

SetBrowsable sets Browsable field to given value.

### HasBrowsable

`func (o *CreateShareSMBParams) HasBrowsable() bool`

HasBrowsable returns a boolean if a field has been set.

### GetTimemachine

`func (o *CreateShareSMBParams) GetTimemachine() bool`

GetTimemachine returns the Timemachine field if non-nil, zero value otherwise.

### GetTimemachineOk

`func (o *CreateShareSMBParams) GetTimemachineOk() (*bool, bool)`

GetTimemachineOk returns a tuple with the Timemachine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimemachine

`func (o *CreateShareSMBParams) SetTimemachine(v bool)`

SetTimemachine sets Timemachine field to given value.

### HasTimemachine

`func (o *CreateShareSMBParams) HasTimemachine() bool`

HasTimemachine returns a boolean if a field has been set.

### GetRecyclebin

`func (o *CreateShareSMBParams) GetRecyclebin() bool`

GetRecyclebin returns the Recyclebin field if non-nil, zero value otherwise.

### GetRecyclebinOk

`func (o *CreateShareSMBParams) GetRecyclebinOk() (*bool, bool)`

GetRecyclebinOk returns a tuple with the Recyclebin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecyclebin

`func (o *CreateShareSMBParams) SetRecyclebin(v bool)`

SetRecyclebin sets Recyclebin field to given value.

### HasRecyclebin

`func (o *CreateShareSMBParams) HasRecyclebin() bool`

HasRecyclebin returns a boolean if a field has been set.

### GetGuestok

`func (o *CreateShareSMBParams) GetGuestok() bool`

GetGuestok returns the Guestok field if non-nil, zero value otherwise.

### GetGuestokOk

`func (o *CreateShareSMBParams) GetGuestokOk() (*bool, bool)`

GetGuestokOk returns a tuple with the Guestok field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestok

`func (o *CreateShareSMBParams) SetGuestok(v bool)`

SetGuestok sets Guestok field to given value.

### HasGuestok

`func (o *CreateShareSMBParams) HasGuestok() bool`

HasGuestok returns a boolean if a field has been set.

### GetAbe

`func (o *CreateShareSMBParams) GetAbe() bool`

GetAbe returns the Abe field if non-nil, zero value otherwise.

### GetAbeOk

`func (o *CreateShareSMBParams) GetAbeOk() (*bool, bool)`

GetAbeOk returns a tuple with the Abe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbe

`func (o *CreateShareSMBParams) SetAbe(v bool)`

SetAbe sets Abe field to given value.

### HasAbe

`func (o *CreateShareSMBParams) HasAbe() bool`

HasAbe returns a boolean if a field has been set.

### GetHostsallow

`func (o *CreateShareSMBParams) GetHostsallow() []string`

GetHostsallow returns the Hostsallow field if non-nil, zero value otherwise.

### GetHostsallowOk

`func (o *CreateShareSMBParams) GetHostsallowOk() (*[]string, bool)`

GetHostsallowOk returns a tuple with the Hostsallow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostsallow

`func (o *CreateShareSMBParams) SetHostsallow(v []string)`

SetHostsallow sets Hostsallow field to given value.

### HasHostsallow

`func (o *CreateShareSMBParams) HasHostsallow() bool`

HasHostsallow returns a boolean if a field has been set.

### GetHostsdeny

`func (o *CreateShareSMBParams) GetHostsdeny() []string`

GetHostsdeny returns the Hostsdeny field if non-nil, zero value otherwise.

### GetHostsdenyOk

`func (o *CreateShareSMBParams) GetHostsdenyOk() (*[]string, bool)`

GetHostsdenyOk returns a tuple with the Hostsdeny field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostsdeny

`func (o *CreateShareSMBParams) SetHostsdeny(v []string)`

SetHostsdeny sets Hostsdeny field to given value.

### HasHostsdeny

`func (o *CreateShareSMBParams) HasHostsdeny() bool`

HasHostsdeny returns a boolean if a field has been set.

### GetAaplNameMangling

`func (o *CreateShareSMBParams) GetAaplNameMangling() bool`

GetAaplNameMangling returns the AaplNameMangling field if non-nil, zero value otherwise.

### GetAaplNameManglingOk

`func (o *CreateShareSMBParams) GetAaplNameManglingOk() (*bool, bool)`

GetAaplNameManglingOk returns a tuple with the AaplNameMangling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAaplNameMangling

`func (o *CreateShareSMBParams) SetAaplNameMangling(v bool)`

SetAaplNameMangling sets AaplNameMangling field to given value.

### HasAaplNameMangling

`func (o *CreateShareSMBParams) HasAaplNameMangling() bool`

HasAaplNameMangling returns a boolean if a field has been set.

### GetAcl

`func (o *CreateShareSMBParams) GetAcl() bool`

GetAcl returns the Acl field if non-nil, zero value otherwise.

### GetAclOk

`func (o *CreateShareSMBParams) GetAclOk() (*bool, bool)`

GetAclOk returns a tuple with the Acl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcl

`func (o *CreateShareSMBParams) SetAcl(v bool)`

SetAcl sets Acl field to given value.

### HasAcl

`func (o *CreateShareSMBParams) HasAcl() bool`

HasAcl returns a boolean if a field has been set.

### GetDurablehandle

`func (o *CreateShareSMBParams) GetDurablehandle() bool`

GetDurablehandle returns the Durablehandle field if non-nil, zero value otherwise.

### GetDurablehandleOk

`func (o *CreateShareSMBParams) GetDurablehandleOk() (*bool, bool)`

GetDurablehandleOk returns a tuple with the Durablehandle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurablehandle

`func (o *CreateShareSMBParams) SetDurablehandle(v bool)`

SetDurablehandle sets Durablehandle field to given value.

### HasDurablehandle

`func (o *CreateShareSMBParams) HasDurablehandle() bool`

HasDurablehandle returns a boolean if a field has been set.

### GetShadowcopy

`func (o *CreateShareSMBParams) GetShadowcopy() bool`

GetShadowcopy returns the Shadowcopy field if non-nil, zero value otherwise.

### GetShadowcopyOk

`func (o *CreateShareSMBParams) GetShadowcopyOk() (*bool, bool)`

GetShadowcopyOk returns a tuple with the Shadowcopy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShadowcopy

`func (o *CreateShareSMBParams) SetShadowcopy(v bool)`

SetShadowcopy sets Shadowcopy field to given value.

### HasShadowcopy

`func (o *CreateShareSMBParams) HasShadowcopy() bool`

HasShadowcopy returns a boolean if a field has been set.

### GetStreams

`func (o *CreateShareSMBParams) GetStreams() bool`

GetStreams returns the Streams field if non-nil, zero value otherwise.

### GetStreamsOk

`func (o *CreateShareSMBParams) GetStreamsOk() (*bool, bool)`

GetStreamsOk returns a tuple with the Streams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreams

`func (o *CreateShareSMBParams) SetStreams(v bool)`

SetStreams sets Streams field to given value.

### HasStreams

`func (o *CreateShareSMBParams) HasStreams() bool`

HasStreams returns a boolean if a field has been set.

### GetFsrvp

`func (o *CreateShareSMBParams) GetFsrvp() bool`

GetFsrvp returns the Fsrvp field if non-nil, zero value otherwise.

### GetFsrvpOk

`func (o *CreateShareSMBParams) GetFsrvpOk() (*bool, bool)`

GetFsrvpOk returns a tuple with the Fsrvp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsrvp

`func (o *CreateShareSMBParams) SetFsrvp(v bool)`

SetFsrvp sets Fsrvp field to given value.

### HasFsrvp

`func (o *CreateShareSMBParams) HasFsrvp() bool`

HasFsrvp returns a boolean if a field has been set.

### GetAuxsmbconf

`func (o *CreateShareSMBParams) GetAuxsmbconf() string`

GetAuxsmbconf returns the Auxsmbconf field if non-nil, zero value otherwise.

### GetAuxsmbconfOk

`func (o *CreateShareSMBParams) GetAuxsmbconfOk() (*string, bool)`

GetAuxsmbconfOk returns a tuple with the Auxsmbconf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuxsmbconf

`func (o *CreateShareSMBParams) SetAuxsmbconf(v string)`

SetAuxsmbconf sets Auxsmbconf field to given value.

### HasAuxsmbconf

`func (o *CreateShareSMBParams) HasAuxsmbconf() bool`

HasAuxsmbconf returns a boolean if a field has been set.

### GetEnabled

`func (o *CreateShareSMBParams) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CreateShareSMBParams) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CreateShareSMBParams) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CreateShareSMBParams) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


