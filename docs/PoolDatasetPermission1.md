# PoolDatasetPermission1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | Pointer to **string** |  | [optional] 
**Group** | Pointer to **string** |  | [optional] 
**Mode** | Pointer to **NullableString** |  | [optional] 
**Acl** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Options** | Pointer to [**FilesystemSetperm0Options**](FilesystemSetperm0Options.md) |  | [optional] 

## Methods

### NewPoolDatasetPermission1

`func NewPoolDatasetPermission1() *PoolDatasetPermission1`

NewPoolDatasetPermission1 instantiates a new PoolDatasetPermission1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoolDatasetPermission1WithDefaults

`func NewPoolDatasetPermission1WithDefaults() *PoolDatasetPermission1`

NewPoolDatasetPermission1WithDefaults instantiates a new PoolDatasetPermission1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *PoolDatasetPermission1) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *PoolDatasetPermission1) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *PoolDatasetPermission1) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *PoolDatasetPermission1) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetGroup

`func (o *PoolDatasetPermission1) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *PoolDatasetPermission1) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *PoolDatasetPermission1) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *PoolDatasetPermission1) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetMode

`func (o *PoolDatasetPermission1) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *PoolDatasetPermission1) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *PoolDatasetPermission1) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *PoolDatasetPermission1) HasMode() bool`

HasMode returns a boolean if a field has been set.

### SetModeNil

`func (o *PoolDatasetPermission1) SetModeNil(b bool)`

 SetModeNil sets the value for Mode to be an explicit nil

### UnsetMode
`func (o *PoolDatasetPermission1) UnsetMode()`

UnsetMode ensures that no value is present for Mode, not even an explicit nil
### GetAcl

`func (o *PoolDatasetPermission1) GetAcl() []map[string]interface{}`

GetAcl returns the Acl field if non-nil, zero value otherwise.

### GetAclOk

`func (o *PoolDatasetPermission1) GetAclOk() (*[]map[string]interface{}, bool)`

GetAclOk returns a tuple with the Acl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcl

`func (o *PoolDatasetPermission1) SetAcl(v []map[string]interface{})`

SetAcl sets Acl field to given value.

### HasAcl

`func (o *PoolDatasetPermission1) HasAcl() bool`

HasAcl returns a boolean if a field has been set.

### GetOptions

`func (o *PoolDatasetPermission1) GetOptions() FilesystemSetperm0Options`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *PoolDatasetPermission1) GetOptionsOk() (*FilesystemSetperm0Options, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *PoolDatasetPermission1) SetOptions(v FilesystemSetperm0Options)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *PoolDatasetPermission1) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


