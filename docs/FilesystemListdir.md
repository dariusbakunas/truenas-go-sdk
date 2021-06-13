# FilesystemListdir

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to **string** |  | [optional] 
**QueryFilters** | Pointer to **[]interface{}** |  | [optional] [default to []]
**QueryOptions** | Pointer to [**FilesystemListdir2**](FilesystemListdir2.md) |  | [optional] [default to {}]

## Methods

### NewFilesystemListdir

`func NewFilesystemListdir() *FilesystemListdir`

NewFilesystemListdir instantiates a new FilesystemListdir object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilesystemListdirWithDefaults

`func NewFilesystemListdirWithDefaults() *FilesystemListdir`

NewFilesystemListdirWithDefaults instantiates a new FilesystemListdir object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *FilesystemListdir) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *FilesystemListdir) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *FilesystemListdir) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *FilesystemListdir) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetQueryFilters

`func (o *FilesystemListdir) GetQueryFilters() []interface{}`

GetQueryFilters returns the QueryFilters field if non-nil, zero value otherwise.

### GetQueryFiltersOk

`func (o *FilesystemListdir) GetQueryFiltersOk() (*[]interface{}, bool)`

GetQueryFiltersOk returns a tuple with the QueryFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryFilters

`func (o *FilesystemListdir) SetQueryFilters(v []interface{})`

SetQueryFilters sets QueryFilters field to given value.

### HasQueryFilters

`func (o *FilesystemListdir) HasQueryFilters() bool`

HasQueryFilters returns a boolean if a field has been set.

### GetQueryOptions

`func (o *FilesystemListdir) GetQueryOptions() FilesystemListdir2`

GetQueryOptions returns the QueryOptions field if non-nil, zero value otherwise.

### GetQueryOptionsOk

`func (o *FilesystemListdir) GetQueryOptionsOk() (*FilesystemListdir2, bool)`

GetQueryOptionsOk returns a tuple with the QueryOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryOptions

`func (o *FilesystemListdir) SetQueryOptions(v FilesystemListdir2)`

SetQueryOptions sets QueryOptions field to given value.

### HasQueryOptions

`func (o *FilesystemListdir) HasQueryOptions() bool`

HasQueryOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


