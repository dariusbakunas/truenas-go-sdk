# FilesystemGetacl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to **string** |  | [optional] 
**Simplified** | Pointer to **bool** |  | [optional] [default to true]

## Methods

### NewFilesystemGetacl

`func NewFilesystemGetacl() *FilesystemGetacl`

NewFilesystemGetacl instantiates a new FilesystemGetacl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilesystemGetaclWithDefaults

`func NewFilesystemGetaclWithDefaults() *FilesystemGetacl`

NewFilesystemGetaclWithDefaults instantiates a new FilesystemGetacl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *FilesystemGetacl) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *FilesystemGetacl) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *FilesystemGetacl) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *FilesystemGetacl) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetSimplified

`func (o *FilesystemGetacl) GetSimplified() bool`

GetSimplified returns the Simplified field if non-nil, zero value otherwise.

### GetSimplifiedOk

`func (o *FilesystemGetacl) GetSimplifiedOk() (*bool, bool)`

GetSimplifiedOk returns a tuple with the Simplified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimplified

`func (o *FilesystemGetacl) SetSimplified(v bool)`

SetSimplified sets Simplified field to given value.

### HasSimplified

`func (o *FilesystemGetacl) HasSimplified() bool`

HasSimplified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


