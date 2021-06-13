# JailFetch0

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Release** | Pointer to **string** |  | [optional] 
**Server** | Pointer to **string** |  | [optional] 
**User** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**JailName** | Pointer to **string** |  | [optional] 
**Accept** | Pointer to **bool** |  | [optional] 
**Https** | Pointer to **bool** |  | [optional] 
**Props** | Pointer to **[]interface{}** |  | [optional] 
**Files** | Pointer to **[]interface{}** |  | [optional] 
**Branch** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewJailFetch0

`func NewJailFetch0() *JailFetch0`

NewJailFetch0 instantiates a new JailFetch0 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJailFetch0WithDefaults

`func NewJailFetch0WithDefaults() *JailFetch0`

NewJailFetch0WithDefaults instantiates a new JailFetch0 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRelease

`func (o *JailFetch0) GetRelease() string`

GetRelease returns the Release field if non-nil, zero value otherwise.

### GetReleaseOk

`func (o *JailFetch0) GetReleaseOk() (*string, bool)`

GetReleaseOk returns a tuple with the Release field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelease

`func (o *JailFetch0) SetRelease(v string)`

SetRelease sets Release field to given value.

### HasRelease

`func (o *JailFetch0) HasRelease() bool`

HasRelease returns a boolean if a field has been set.

### GetServer

`func (o *JailFetch0) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *JailFetch0) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *JailFetch0) SetServer(v string)`

SetServer sets Server field to given value.

### HasServer

`func (o *JailFetch0) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetUser

`func (o *JailFetch0) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *JailFetch0) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *JailFetch0) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *JailFetch0) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetPassword

`func (o *JailFetch0) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *JailFetch0) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *JailFetch0) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *JailFetch0) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetName

`func (o *JailFetch0) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *JailFetch0) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *JailFetch0) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *JailFetch0) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *JailFetch0) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *JailFetch0) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetJailName

`func (o *JailFetch0) GetJailName() string`

GetJailName returns the JailName field if non-nil, zero value otherwise.

### GetJailNameOk

`func (o *JailFetch0) GetJailNameOk() (*string, bool)`

GetJailNameOk returns a tuple with the JailName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJailName

`func (o *JailFetch0) SetJailName(v string)`

SetJailName sets JailName field to given value.

### HasJailName

`func (o *JailFetch0) HasJailName() bool`

HasJailName returns a boolean if a field has been set.

### GetAccept

`func (o *JailFetch0) GetAccept() bool`

GetAccept returns the Accept field if non-nil, zero value otherwise.

### GetAcceptOk

`func (o *JailFetch0) GetAcceptOk() (*bool, bool)`

GetAcceptOk returns a tuple with the Accept field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccept

`func (o *JailFetch0) SetAccept(v bool)`

SetAccept sets Accept field to given value.

### HasAccept

`func (o *JailFetch0) HasAccept() bool`

HasAccept returns a boolean if a field has been set.

### GetHttps

`func (o *JailFetch0) GetHttps() bool`

GetHttps returns the Https field if non-nil, zero value otherwise.

### GetHttpsOk

`func (o *JailFetch0) GetHttpsOk() (*bool, bool)`

GetHttpsOk returns a tuple with the Https field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttps

`func (o *JailFetch0) SetHttps(v bool)`

SetHttps sets Https field to given value.

### HasHttps

`func (o *JailFetch0) HasHttps() bool`

HasHttps returns a boolean if a field has been set.

### GetProps

`func (o *JailFetch0) GetProps() []interface{}`

GetProps returns the Props field if non-nil, zero value otherwise.

### GetPropsOk

`func (o *JailFetch0) GetPropsOk() (*[]interface{}, bool)`

GetPropsOk returns a tuple with the Props field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProps

`func (o *JailFetch0) SetProps(v []interface{})`

SetProps sets Props field to given value.

### HasProps

`func (o *JailFetch0) HasProps() bool`

HasProps returns a boolean if a field has been set.

### GetFiles

`func (o *JailFetch0) GetFiles() []interface{}`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *JailFetch0) GetFilesOk() (*[]interface{}, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *JailFetch0) SetFiles(v []interface{})`

SetFiles sets Files field to given value.

### HasFiles

`func (o *JailFetch0) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetBranch

`func (o *JailFetch0) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *JailFetch0) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *JailFetch0) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *JailFetch0) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### SetBranchNil

`func (o *JailFetch0) SetBranchNil(b bool)`

 SetBranchNil sets the value for Branch to be an explicit nil

### UnsetBranch
`func (o *JailFetch0) UnsetBranch()`

UnsetBranch ensures that no value is present for Branch, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


