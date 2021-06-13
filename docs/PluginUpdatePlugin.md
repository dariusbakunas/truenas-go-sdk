# PluginUpdatePlugin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Jail** | Pointer to **string** |  | [optional] 
**UpdateJail** | Pointer to **bool** |  | [optional] [default to true]

## Methods

### NewPluginUpdatePlugin

`func NewPluginUpdatePlugin() *PluginUpdatePlugin`

NewPluginUpdatePlugin instantiates a new PluginUpdatePlugin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPluginUpdatePluginWithDefaults

`func NewPluginUpdatePluginWithDefaults() *PluginUpdatePlugin`

NewPluginUpdatePluginWithDefaults instantiates a new PluginUpdatePlugin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJail

`func (o *PluginUpdatePlugin) GetJail() string`

GetJail returns the Jail field if non-nil, zero value otherwise.

### GetJailOk

`func (o *PluginUpdatePlugin) GetJailOk() (*string, bool)`

GetJailOk returns a tuple with the Jail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJail

`func (o *PluginUpdatePlugin) SetJail(v string)`

SetJail sets Jail field to given value.

### HasJail

`func (o *PluginUpdatePlugin) HasJail() bool`

HasJail returns a boolean if a field has been set.

### GetUpdateJail

`func (o *PluginUpdatePlugin) GetUpdateJail() bool`

GetUpdateJail returns the UpdateJail field if non-nil, zero value otherwise.

### GetUpdateJailOk

`func (o *PluginUpdatePlugin) GetUpdateJailOk() (*bool, bool)`

GetUpdateJailOk returns a tuple with the UpdateJail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateJail

`func (o *PluginUpdatePlugin) SetUpdateJail(v bool)`

SetUpdateJail sets UpdateJail field to given value.

### HasUpdateJail

`func (o *PluginUpdatePlugin) HasUpdateJail() bool`

HasUpdateJail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


