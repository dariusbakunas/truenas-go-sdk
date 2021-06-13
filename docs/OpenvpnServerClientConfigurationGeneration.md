# OpenvpnServerClientConfigurationGeneration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientCertificateId** | Pointer to **int32** |  | [optional] 
**ServerAddress** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewOpenvpnServerClientConfigurationGeneration

`func NewOpenvpnServerClientConfigurationGeneration() *OpenvpnServerClientConfigurationGeneration`

NewOpenvpnServerClientConfigurationGeneration instantiates a new OpenvpnServerClientConfigurationGeneration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenvpnServerClientConfigurationGenerationWithDefaults

`func NewOpenvpnServerClientConfigurationGenerationWithDefaults() *OpenvpnServerClientConfigurationGeneration`

NewOpenvpnServerClientConfigurationGenerationWithDefaults instantiates a new OpenvpnServerClientConfigurationGeneration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientCertificateId

`func (o *OpenvpnServerClientConfigurationGeneration) GetClientCertificateId() int32`

GetClientCertificateId returns the ClientCertificateId field if non-nil, zero value otherwise.

### GetClientCertificateIdOk

`func (o *OpenvpnServerClientConfigurationGeneration) GetClientCertificateIdOk() (*int32, bool)`

GetClientCertificateIdOk returns a tuple with the ClientCertificateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertificateId

`func (o *OpenvpnServerClientConfigurationGeneration) SetClientCertificateId(v int32)`

SetClientCertificateId sets ClientCertificateId field to given value.

### HasClientCertificateId

`func (o *OpenvpnServerClientConfigurationGeneration) HasClientCertificateId() bool`

HasClientCertificateId returns a boolean if a field has been set.

### GetServerAddress

`func (o *OpenvpnServerClientConfigurationGeneration) GetServerAddress() string`

GetServerAddress returns the ServerAddress field if non-nil, zero value otherwise.

### GetServerAddressOk

`func (o *OpenvpnServerClientConfigurationGeneration) GetServerAddressOk() (*string, bool)`

GetServerAddressOk returns a tuple with the ServerAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerAddress

`func (o *OpenvpnServerClientConfigurationGeneration) SetServerAddress(v string)`

SetServerAddress sets ServerAddress field to given value.

### HasServerAddress

`func (o *OpenvpnServerClientConfigurationGeneration) HasServerAddress() bool`

HasServerAddress returns a boolean if a field has been set.

### SetServerAddressNil

`func (o *OpenvpnServerClientConfigurationGeneration) SetServerAddressNil(b bool)`

 SetServerAddressNil sets the value for ServerAddress to be an explicit nil

### UnsetServerAddress
`func (o *OpenvpnServerClientConfigurationGeneration) UnsetServerAddress()`

UnsetServerAddress ensures that no value is present for ServerAddress, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


