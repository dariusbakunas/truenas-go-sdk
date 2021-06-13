# JailUpdateToLatestPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Jail** | Pointer to **string** |  | [optional] 
**UpdatePkgs** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewJailUpdateToLatestPatch

`func NewJailUpdateToLatestPatch() *JailUpdateToLatestPatch`

NewJailUpdateToLatestPatch instantiates a new JailUpdateToLatestPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJailUpdateToLatestPatchWithDefaults

`func NewJailUpdateToLatestPatchWithDefaults() *JailUpdateToLatestPatch`

NewJailUpdateToLatestPatchWithDefaults instantiates a new JailUpdateToLatestPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJail

`func (o *JailUpdateToLatestPatch) GetJail() string`

GetJail returns the Jail field if non-nil, zero value otherwise.

### GetJailOk

`func (o *JailUpdateToLatestPatch) GetJailOk() (*string, bool)`

GetJailOk returns a tuple with the Jail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJail

`func (o *JailUpdateToLatestPatch) SetJail(v string)`

SetJail sets Jail field to given value.

### HasJail

`func (o *JailUpdateToLatestPatch) HasJail() bool`

HasJail returns a boolean if a field has been set.

### GetUpdatePkgs

`func (o *JailUpdateToLatestPatch) GetUpdatePkgs() bool`

GetUpdatePkgs returns the UpdatePkgs field if non-nil, zero value otherwise.

### GetUpdatePkgsOk

`func (o *JailUpdateToLatestPatch) GetUpdatePkgsOk() (*bool, bool)`

GetUpdatePkgsOk returns a tuple with the UpdatePkgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatePkgs

`func (o *JailUpdateToLatestPatch) SetUpdatePkgs(v bool)`

SetUpdatePkgs sets UpdatePkgs field to given value.

### HasUpdatePkgs

`func (o *JailUpdateToLatestPatch) HasUpdatePkgs() bool`

HasUpdatePkgs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


