# IscsiPortalCreate0

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** |  | [optional] 
**DiscoveryAuthmethod** | Pointer to **string** |  | [optional] 
**DiscoveryAuthgroup** | Pointer to **NullableInt32** |  | [optional] 
**Listen** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewIscsiPortalCreate0

`func NewIscsiPortalCreate0() *IscsiPortalCreate0`

NewIscsiPortalCreate0 instantiates a new IscsiPortalCreate0 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIscsiPortalCreate0WithDefaults

`func NewIscsiPortalCreate0WithDefaults() *IscsiPortalCreate0`

NewIscsiPortalCreate0WithDefaults instantiates a new IscsiPortalCreate0 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *IscsiPortalCreate0) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *IscsiPortalCreate0) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *IscsiPortalCreate0) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *IscsiPortalCreate0) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDiscoveryAuthmethod

`func (o *IscsiPortalCreate0) GetDiscoveryAuthmethod() string`

GetDiscoveryAuthmethod returns the DiscoveryAuthmethod field if non-nil, zero value otherwise.

### GetDiscoveryAuthmethodOk

`func (o *IscsiPortalCreate0) GetDiscoveryAuthmethodOk() (*string, bool)`

GetDiscoveryAuthmethodOk returns a tuple with the DiscoveryAuthmethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryAuthmethod

`func (o *IscsiPortalCreate0) SetDiscoveryAuthmethod(v string)`

SetDiscoveryAuthmethod sets DiscoveryAuthmethod field to given value.

### HasDiscoveryAuthmethod

`func (o *IscsiPortalCreate0) HasDiscoveryAuthmethod() bool`

HasDiscoveryAuthmethod returns a boolean if a field has been set.

### GetDiscoveryAuthgroup

`func (o *IscsiPortalCreate0) GetDiscoveryAuthgroup() int32`

GetDiscoveryAuthgroup returns the DiscoveryAuthgroup field if non-nil, zero value otherwise.

### GetDiscoveryAuthgroupOk

`func (o *IscsiPortalCreate0) GetDiscoveryAuthgroupOk() (*int32, bool)`

GetDiscoveryAuthgroupOk returns a tuple with the DiscoveryAuthgroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryAuthgroup

`func (o *IscsiPortalCreate0) SetDiscoveryAuthgroup(v int32)`

SetDiscoveryAuthgroup sets DiscoveryAuthgroup field to given value.

### HasDiscoveryAuthgroup

`func (o *IscsiPortalCreate0) HasDiscoveryAuthgroup() bool`

HasDiscoveryAuthgroup returns a boolean if a field has been set.

### SetDiscoveryAuthgroupNil

`func (o *IscsiPortalCreate0) SetDiscoveryAuthgroupNil(b bool)`

 SetDiscoveryAuthgroupNil sets the value for DiscoveryAuthgroup to be an explicit nil

### UnsetDiscoveryAuthgroup
`func (o *IscsiPortalCreate0) UnsetDiscoveryAuthgroup()`

UnsetDiscoveryAuthgroup ensures that no value is present for DiscoveryAuthgroup, not even an explicit nil
### GetListen

`func (o *IscsiPortalCreate0) GetListen() []map[string]interface{}`

GetListen returns the Listen field if non-nil, zero value otherwise.

### GetListenOk

`func (o *IscsiPortalCreate0) GetListenOk() (*[]map[string]interface{}, bool)`

GetListenOk returns a tuple with the Listen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListen

`func (o *IscsiPortalCreate0) SetListen(v []map[string]interface{})`

SetListen sets Listen field to given value.

### HasListen

`func (o *IscsiPortalCreate0) HasListen() bool`

HasListen returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


