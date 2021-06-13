# OpenvpnServerUpdate0

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TlsCryptAuthEnabled** | Pointer to **bool** |  | [optional] 
**Netmask** | Pointer to **int32** |  | [optional] 
**ServerCertificate** | Pointer to **NullableInt32** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**RootCa** | Pointer to **NullableInt32** |  | [optional] 
**Server** | Pointer to **string** |  | [optional] 
**AdditionalParameters** | Pointer to **string** |  | [optional] 
**AuthenticationAlgorithm** | Pointer to **NullableString** |  | [optional] 
**Cipher** | Pointer to **NullableString** |  | [optional] 
**Compression** | Pointer to **NullableString** |  | [optional] 
**DeviceType** | Pointer to **string** |  | [optional] 
**Protocol** | Pointer to **string** |  | [optional] 
**TlsCryptAuth** | Pointer to **NullableString** |  | [optional] 
**Topology** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewOpenvpnServerUpdate0

`func NewOpenvpnServerUpdate0() *OpenvpnServerUpdate0`

NewOpenvpnServerUpdate0 instantiates a new OpenvpnServerUpdate0 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenvpnServerUpdate0WithDefaults

`func NewOpenvpnServerUpdate0WithDefaults() *OpenvpnServerUpdate0`

NewOpenvpnServerUpdate0WithDefaults instantiates a new OpenvpnServerUpdate0 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTlsCryptAuthEnabled

`func (o *OpenvpnServerUpdate0) GetTlsCryptAuthEnabled() bool`

GetTlsCryptAuthEnabled returns the TlsCryptAuthEnabled field if non-nil, zero value otherwise.

### GetTlsCryptAuthEnabledOk

`func (o *OpenvpnServerUpdate0) GetTlsCryptAuthEnabledOk() (*bool, bool)`

GetTlsCryptAuthEnabledOk returns a tuple with the TlsCryptAuthEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsCryptAuthEnabled

`func (o *OpenvpnServerUpdate0) SetTlsCryptAuthEnabled(v bool)`

SetTlsCryptAuthEnabled sets TlsCryptAuthEnabled field to given value.

### HasTlsCryptAuthEnabled

`func (o *OpenvpnServerUpdate0) HasTlsCryptAuthEnabled() bool`

HasTlsCryptAuthEnabled returns a boolean if a field has been set.

### GetNetmask

`func (o *OpenvpnServerUpdate0) GetNetmask() int32`

GetNetmask returns the Netmask field if non-nil, zero value otherwise.

### GetNetmaskOk

`func (o *OpenvpnServerUpdate0) GetNetmaskOk() (*int32, bool)`

GetNetmaskOk returns a tuple with the Netmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask

`func (o *OpenvpnServerUpdate0) SetNetmask(v int32)`

SetNetmask sets Netmask field to given value.

### HasNetmask

`func (o *OpenvpnServerUpdate0) HasNetmask() bool`

HasNetmask returns a boolean if a field has been set.

### GetServerCertificate

`func (o *OpenvpnServerUpdate0) GetServerCertificate() int32`

GetServerCertificate returns the ServerCertificate field if non-nil, zero value otherwise.

### GetServerCertificateOk

`func (o *OpenvpnServerUpdate0) GetServerCertificateOk() (*int32, bool)`

GetServerCertificateOk returns a tuple with the ServerCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerCertificate

`func (o *OpenvpnServerUpdate0) SetServerCertificate(v int32)`

SetServerCertificate sets ServerCertificate field to given value.

### HasServerCertificate

`func (o *OpenvpnServerUpdate0) HasServerCertificate() bool`

HasServerCertificate returns a boolean if a field has been set.

### SetServerCertificateNil

`func (o *OpenvpnServerUpdate0) SetServerCertificateNil(b bool)`

 SetServerCertificateNil sets the value for ServerCertificate to be an explicit nil

### UnsetServerCertificate
`func (o *OpenvpnServerUpdate0) UnsetServerCertificate()`

UnsetServerCertificate ensures that no value is present for ServerCertificate, not even an explicit nil
### GetPort

`func (o *OpenvpnServerUpdate0) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *OpenvpnServerUpdate0) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *OpenvpnServerUpdate0) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *OpenvpnServerUpdate0) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetRootCa

`func (o *OpenvpnServerUpdate0) GetRootCa() int32`

GetRootCa returns the RootCa field if non-nil, zero value otherwise.

### GetRootCaOk

`func (o *OpenvpnServerUpdate0) GetRootCaOk() (*int32, bool)`

GetRootCaOk returns a tuple with the RootCa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootCa

`func (o *OpenvpnServerUpdate0) SetRootCa(v int32)`

SetRootCa sets RootCa field to given value.

### HasRootCa

`func (o *OpenvpnServerUpdate0) HasRootCa() bool`

HasRootCa returns a boolean if a field has been set.

### SetRootCaNil

`func (o *OpenvpnServerUpdate0) SetRootCaNil(b bool)`

 SetRootCaNil sets the value for RootCa to be an explicit nil

### UnsetRootCa
`func (o *OpenvpnServerUpdate0) UnsetRootCa()`

UnsetRootCa ensures that no value is present for RootCa, not even an explicit nil
### GetServer

`func (o *OpenvpnServerUpdate0) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *OpenvpnServerUpdate0) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *OpenvpnServerUpdate0) SetServer(v string)`

SetServer sets Server field to given value.

### HasServer

`func (o *OpenvpnServerUpdate0) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetAdditionalParameters

`func (o *OpenvpnServerUpdate0) GetAdditionalParameters() string`

GetAdditionalParameters returns the AdditionalParameters field if non-nil, zero value otherwise.

### GetAdditionalParametersOk

`func (o *OpenvpnServerUpdate0) GetAdditionalParametersOk() (*string, bool)`

GetAdditionalParametersOk returns a tuple with the AdditionalParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalParameters

`func (o *OpenvpnServerUpdate0) SetAdditionalParameters(v string)`

SetAdditionalParameters sets AdditionalParameters field to given value.

### HasAdditionalParameters

`func (o *OpenvpnServerUpdate0) HasAdditionalParameters() bool`

HasAdditionalParameters returns a boolean if a field has been set.

### GetAuthenticationAlgorithm

`func (o *OpenvpnServerUpdate0) GetAuthenticationAlgorithm() string`

GetAuthenticationAlgorithm returns the AuthenticationAlgorithm field if non-nil, zero value otherwise.

### GetAuthenticationAlgorithmOk

`func (o *OpenvpnServerUpdate0) GetAuthenticationAlgorithmOk() (*string, bool)`

GetAuthenticationAlgorithmOk returns a tuple with the AuthenticationAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationAlgorithm

`func (o *OpenvpnServerUpdate0) SetAuthenticationAlgorithm(v string)`

SetAuthenticationAlgorithm sets AuthenticationAlgorithm field to given value.

### HasAuthenticationAlgorithm

`func (o *OpenvpnServerUpdate0) HasAuthenticationAlgorithm() bool`

HasAuthenticationAlgorithm returns a boolean if a field has been set.

### SetAuthenticationAlgorithmNil

`func (o *OpenvpnServerUpdate0) SetAuthenticationAlgorithmNil(b bool)`

 SetAuthenticationAlgorithmNil sets the value for AuthenticationAlgorithm to be an explicit nil

### UnsetAuthenticationAlgorithm
`func (o *OpenvpnServerUpdate0) UnsetAuthenticationAlgorithm()`

UnsetAuthenticationAlgorithm ensures that no value is present for AuthenticationAlgorithm, not even an explicit nil
### GetCipher

`func (o *OpenvpnServerUpdate0) GetCipher() string`

GetCipher returns the Cipher field if non-nil, zero value otherwise.

### GetCipherOk

`func (o *OpenvpnServerUpdate0) GetCipherOk() (*string, bool)`

GetCipherOk returns a tuple with the Cipher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCipher

`func (o *OpenvpnServerUpdate0) SetCipher(v string)`

SetCipher sets Cipher field to given value.

### HasCipher

`func (o *OpenvpnServerUpdate0) HasCipher() bool`

HasCipher returns a boolean if a field has been set.

### SetCipherNil

`func (o *OpenvpnServerUpdate0) SetCipherNil(b bool)`

 SetCipherNil sets the value for Cipher to be an explicit nil

### UnsetCipher
`func (o *OpenvpnServerUpdate0) UnsetCipher()`

UnsetCipher ensures that no value is present for Cipher, not even an explicit nil
### GetCompression

`func (o *OpenvpnServerUpdate0) GetCompression() string`

GetCompression returns the Compression field if non-nil, zero value otherwise.

### GetCompressionOk

`func (o *OpenvpnServerUpdate0) GetCompressionOk() (*string, bool)`

GetCompressionOk returns a tuple with the Compression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompression

`func (o *OpenvpnServerUpdate0) SetCompression(v string)`

SetCompression sets Compression field to given value.

### HasCompression

`func (o *OpenvpnServerUpdate0) HasCompression() bool`

HasCompression returns a boolean if a field has been set.

### SetCompressionNil

`func (o *OpenvpnServerUpdate0) SetCompressionNil(b bool)`

 SetCompressionNil sets the value for Compression to be an explicit nil

### UnsetCompression
`func (o *OpenvpnServerUpdate0) UnsetCompression()`

UnsetCompression ensures that no value is present for Compression, not even an explicit nil
### GetDeviceType

`func (o *OpenvpnServerUpdate0) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *OpenvpnServerUpdate0) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *OpenvpnServerUpdate0) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *OpenvpnServerUpdate0) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetProtocol

`func (o *OpenvpnServerUpdate0) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *OpenvpnServerUpdate0) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *OpenvpnServerUpdate0) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *OpenvpnServerUpdate0) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetTlsCryptAuth

`func (o *OpenvpnServerUpdate0) GetTlsCryptAuth() string`

GetTlsCryptAuth returns the TlsCryptAuth field if non-nil, zero value otherwise.

### GetTlsCryptAuthOk

`func (o *OpenvpnServerUpdate0) GetTlsCryptAuthOk() (*string, bool)`

GetTlsCryptAuthOk returns a tuple with the TlsCryptAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsCryptAuth

`func (o *OpenvpnServerUpdate0) SetTlsCryptAuth(v string)`

SetTlsCryptAuth sets TlsCryptAuth field to given value.

### HasTlsCryptAuth

`func (o *OpenvpnServerUpdate0) HasTlsCryptAuth() bool`

HasTlsCryptAuth returns a boolean if a field has been set.

### SetTlsCryptAuthNil

`func (o *OpenvpnServerUpdate0) SetTlsCryptAuthNil(b bool)`

 SetTlsCryptAuthNil sets the value for TlsCryptAuth to be an explicit nil

### UnsetTlsCryptAuth
`func (o *OpenvpnServerUpdate0) UnsetTlsCryptAuth()`

UnsetTlsCryptAuth ensures that no value is present for TlsCryptAuth, not even an explicit nil
### GetTopology

`func (o *OpenvpnServerUpdate0) GetTopology() string`

GetTopology returns the Topology field if non-nil, zero value otherwise.

### GetTopologyOk

`func (o *OpenvpnServerUpdate0) GetTopologyOk() (*string, bool)`

GetTopologyOk returns a tuple with the Topology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopology

`func (o *OpenvpnServerUpdate0) SetTopology(v string)`

SetTopology sets Topology field to given value.

### HasTopology

`func (o *OpenvpnServerUpdate0) HasTopology() bool`

HasTopology returns a boolean if a field has been set.

### SetTopologyNil

`func (o *OpenvpnServerUpdate0) SetTopologyNil(b bool)`

 SetTopologyNil sets the value for Topology to be an explicit nil

### UnsetTopology
`func (o *OpenvpnServerUpdate0) UnsetTopology()`

UnsetTopology ensures that no value is present for Topology, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


