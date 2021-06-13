# OpenvpnClientUpdate0

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nobind** | Pointer to **bool** |  | [optional] 
**TlsCryptAuthEnabled** | Pointer to **bool** |  | [optional] 
**ClientCertificate** | Pointer to **NullableInt32** |  | [optional] 
**RootCa** | Pointer to **NullableInt32** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**AdditionalParameters** | Pointer to **string** |  | [optional] 
**AuthenticationAlgorithm** | Pointer to **NullableString** |  | [optional] 
**Cipher** | Pointer to **NullableString** |  | [optional] 
**Compression** | Pointer to **NullableString** |  | [optional] 
**DeviceType** | Pointer to **string** |  | [optional] 
**Protocol** | Pointer to **string** |  | [optional] 
**Remote** | Pointer to **string** |  | [optional] 
**TlsCryptAuth** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewOpenvpnClientUpdate0

`func NewOpenvpnClientUpdate0() *OpenvpnClientUpdate0`

NewOpenvpnClientUpdate0 instantiates a new OpenvpnClientUpdate0 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenvpnClientUpdate0WithDefaults

`func NewOpenvpnClientUpdate0WithDefaults() *OpenvpnClientUpdate0`

NewOpenvpnClientUpdate0WithDefaults instantiates a new OpenvpnClientUpdate0 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNobind

`func (o *OpenvpnClientUpdate0) GetNobind() bool`

GetNobind returns the Nobind field if non-nil, zero value otherwise.

### GetNobindOk

`func (o *OpenvpnClientUpdate0) GetNobindOk() (*bool, bool)`

GetNobindOk returns a tuple with the Nobind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNobind

`func (o *OpenvpnClientUpdate0) SetNobind(v bool)`

SetNobind sets Nobind field to given value.

### HasNobind

`func (o *OpenvpnClientUpdate0) HasNobind() bool`

HasNobind returns a boolean if a field has been set.

### GetTlsCryptAuthEnabled

`func (o *OpenvpnClientUpdate0) GetTlsCryptAuthEnabled() bool`

GetTlsCryptAuthEnabled returns the TlsCryptAuthEnabled field if non-nil, zero value otherwise.

### GetTlsCryptAuthEnabledOk

`func (o *OpenvpnClientUpdate0) GetTlsCryptAuthEnabledOk() (*bool, bool)`

GetTlsCryptAuthEnabledOk returns a tuple with the TlsCryptAuthEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsCryptAuthEnabled

`func (o *OpenvpnClientUpdate0) SetTlsCryptAuthEnabled(v bool)`

SetTlsCryptAuthEnabled sets TlsCryptAuthEnabled field to given value.

### HasTlsCryptAuthEnabled

`func (o *OpenvpnClientUpdate0) HasTlsCryptAuthEnabled() bool`

HasTlsCryptAuthEnabled returns a boolean if a field has been set.

### GetClientCertificate

`func (o *OpenvpnClientUpdate0) GetClientCertificate() int32`

GetClientCertificate returns the ClientCertificate field if non-nil, zero value otherwise.

### GetClientCertificateOk

`func (o *OpenvpnClientUpdate0) GetClientCertificateOk() (*int32, bool)`

GetClientCertificateOk returns a tuple with the ClientCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertificate

`func (o *OpenvpnClientUpdate0) SetClientCertificate(v int32)`

SetClientCertificate sets ClientCertificate field to given value.

### HasClientCertificate

`func (o *OpenvpnClientUpdate0) HasClientCertificate() bool`

HasClientCertificate returns a boolean if a field has been set.

### SetClientCertificateNil

`func (o *OpenvpnClientUpdate0) SetClientCertificateNil(b bool)`

 SetClientCertificateNil sets the value for ClientCertificate to be an explicit nil

### UnsetClientCertificate
`func (o *OpenvpnClientUpdate0) UnsetClientCertificate()`

UnsetClientCertificate ensures that no value is present for ClientCertificate, not even an explicit nil
### GetRootCa

`func (o *OpenvpnClientUpdate0) GetRootCa() int32`

GetRootCa returns the RootCa field if non-nil, zero value otherwise.

### GetRootCaOk

`func (o *OpenvpnClientUpdate0) GetRootCaOk() (*int32, bool)`

GetRootCaOk returns a tuple with the RootCa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootCa

`func (o *OpenvpnClientUpdate0) SetRootCa(v int32)`

SetRootCa sets RootCa field to given value.

### HasRootCa

`func (o *OpenvpnClientUpdate0) HasRootCa() bool`

HasRootCa returns a boolean if a field has been set.

### SetRootCaNil

`func (o *OpenvpnClientUpdate0) SetRootCaNil(b bool)`

 SetRootCaNil sets the value for RootCa to be an explicit nil

### UnsetRootCa
`func (o *OpenvpnClientUpdate0) UnsetRootCa()`

UnsetRootCa ensures that no value is present for RootCa, not even an explicit nil
### GetPort

`func (o *OpenvpnClientUpdate0) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *OpenvpnClientUpdate0) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *OpenvpnClientUpdate0) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *OpenvpnClientUpdate0) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetAdditionalParameters

`func (o *OpenvpnClientUpdate0) GetAdditionalParameters() string`

GetAdditionalParameters returns the AdditionalParameters field if non-nil, zero value otherwise.

### GetAdditionalParametersOk

`func (o *OpenvpnClientUpdate0) GetAdditionalParametersOk() (*string, bool)`

GetAdditionalParametersOk returns a tuple with the AdditionalParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalParameters

`func (o *OpenvpnClientUpdate0) SetAdditionalParameters(v string)`

SetAdditionalParameters sets AdditionalParameters field to given value.

### HasAdditionalParameters

`func (o *OpenvpnClientUpdate0) HasAdditionalParameters() bool`

HasAdditionalParameters returns a boolean if a field has been set.

### GetAuthenticationAlgorithm

`func (o *OpenvpnClientUpdate0) GetAuthenticationAlgorithm() string`

GetAuthenticationAlgorithm returns the AuthenticationAlgorithm field if non-nil, zero value otherwise.

### GetAuthenticationAlgorithmOk

`func (o *OpenvpnClientUpdate0) GetAuthenticationAlgorithmOk() (*string, bool)`

GetAuthenticationAlgorithmOk returns a tuple with the AuthenticationAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationAlgorithm

`func (o *OpenvpnClientUpdate0) SetAuthenticationAlgorithm(v string)`

SetAuthenticationAlgorithm sets AuthenticationAlgorithm field to given value.

### HasAuthenticationAlgorithm

`func (o *OpenvpnClientUpdate0) HasAuthenticationAlgorithm() bool`

HasAuthenticationAlgorithm returns a boolean if a field has been set.

### SetAuthenticationAlgorithmNil

`func (o *OpenvpnClientUpdate0) SetAuthenticationAlgorithmNil(b bool)`

 SetAuthenticationAlgorithmNil sets the value for AuthenticationAlgorithm to be an explicit nil

### UnsetAuthenticationAlgorithm
`func (o *OpenvpnClientUpdate0) UnsetAuthenticationAlgorithm()`

UnsetAuthenticationAlgorithm ensures that no value is present for AuthenticationAlgorithm, not even an explicit nil
### GetCipher

`func (o *OpenvpnClientUpdate0) GetCipher() string`

GetCipher returns the Cipher field if non-nil, zero value otherwise.

### GetCipherOk

`func (o *OpenvpnClientUpdate0) GetCipherOk() (*string, bool)`

GetCipherOk returns a tuple with the Cipher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCipher

`func (o *OpenvpnClientUpdate0) SetCipher(v string)`

SetCipher sets Cipher field to given value.

### HasCipher

`func (o *OpenvpnClientUpdate0) HasCipher() bool`

HasCipher returns a boolean if a field has been set.

### SetCipherNil

`func (o *OpenvpnClientUpdate0) SetCipherNil(b bool)`

 SetCipherNil sets the value for Cipher to be an explicit nil

### UnsetCipher
`func (o *OpenvpnClientUpdate0) UnsetCipher()`

UnsetCipher ensures that no value is present for Cipher, not even an explicit nil
### GetCompression

`func (o *OpenvpnClientUpdate0) GetCompression() string`

GetCompression returns the Compression field if non-nil, zero value otherwise.

### GetCompressionOk

`func (o *OpenvpnClientUpdate0) GetCompressionOk() (*string, bool)`

GetCompressionOk returns a tuple with the Compression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompression

`func (o *OpenvpnClientUpdate0) SetCompression(v string)`

SetCompression sets Compression field to given value.

### HasCompression

`func (o *OpenvpnClientUpdate0) HasCompression() bool`

HasCompression returns a boolean if a field has been set.

### SetCompressionNil

`func (o *OpenvpnClientUpdate0) SetCompressionNil(b bool)`

 SetCompressionNil sets the value for Compression to be an explicit nil

### UnsetCompression
`func (o *OpenvpnClientUpdate0) UnsetCompression()`

UnsetCompression ensures that no value is present for Compression, not even an explicit nil
### GetDeviceType

`func (o *OpenvpnClientUpdate0) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *OpenvpnClientUpdate0) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *OpenvpnClientUpdate0) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *OpenvpnClientUpdate0) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetProtocol

`func (o *OpenvpnClientUpdate0) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *OpenvpnClientUpdate0) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *OpenvpnClientUpdate0) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *OpenvpnClientUpdate0) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetRemote

`func (o *OpenvpnClientUpdate0) GetRemote() string`

GetRemote returns the Remote field if non-nil, zero value otherwise.

### GetRemoteOk

`func (o *OpenvpnClientUpdate0) GetRemoteOk() (*string, bool)`

GetRemoteOk returns a tuple with the Remote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemote

`func (o *OpenvpnClientUpdate0) SetRemote(v string)`

SetRemote sets Remote field to given value.

### HasRemote

`func (o *OpenvpnClientUpdate0) HasRemote() bool`

HasRemote returns a boolean if a field has been set.

### GetTlsCryptAuth

`func (o *OpenvpnClientUpdate0) GetTlsCryptAuth() string`

GetTlsCryptAuth returns the TlsCryptAuth field if non-nil, zero value otherwise.

### GetTlsCryptAuthOk

`func (o *OpenvpnClientUpdate0) GetTlsCryptAuthOk() (*string, bool)`

GetTlsCryptAuthOk returns a tuple with the TlsCryptAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsCryptAuth

`func (o *OpenvpnClientUpdate0) SetTlsCryptAuth(v string)`

SetTlsCryptAuth sets TlsCryptAuth field to given value.

### HasTlsCryptAuth

`func (o *OpenvpnClientUpdate0) HasTlsCryptAuth() bool`

HasTlsCryptAuth returns a boolean if a field has been set.

### SetTlsCryptAuthNil

`func (o *OpenvpnClientUpdate0) SetTlsCryptAuthNil(b bool)`

 SetTlsCryptAuthNil sets the value for TlsCryptAuth to be an explicit nil

### UnsetTlsCryptAuth
`func (o *OpenvpnClientUpdate0) UnsetTlsCryptAuth()`

UnsetTlsCryptAuth ensures that no value is present for TlsCryptAuth, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


