# LdapUpdate0

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hostname** | Pointer to **[]interface{}** |  | [optional] 
**Basedn** | Pointer to **string** |  | [optional] 
**Binddn** | Pointer to **string** |  | [optional] 
**Bindpw** | Pointer to **string** |  | [optional] 
**Anonbind** | Pointer to **bool** |  | [optional] 
**Ssl** | Pointer to **string** |  | [optional] 
**Certificate** | Pointer to **NullableInt32** |  | [optional] 
**ValidateCertificates** | Pointer to **bool** |  | [optional] 
**DisableFreenasCache** | Pointer to **bool** |  | [optional] 
**Timeout** | Pointer to **int32** |  | [optional] 
**DnsTimeout** | Pointer to **int32** |  | [optional] 
**KerberosRealm** | Pointer to **NullableInt32** |  | [optional] 
**KerberosPrincipal** | Pointer to **string** |  | [optional] 
**HasSambaSchema** | Pointer to **bool** |  | [optional] 
**AuxiliaryParameters** | Pointer to **string** |  | [optional] 
**Schema** | Pointer to **string** |  | [optional] 
**Enable** | Pointer to **bool** |  | [optional] 

## Methods

### NewLdapUpdate0

`func NewLdapUpdate0() *LdapUpdate0`

NewLdapUpdate0 instantiates a new LdapUpdate0 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLdapUpdate0WithDefaults

`func NewLdapUpdate0WithDefaults() *LdapUpdate0`

NewLdapUpdate0WithDefaults instantiates a new LdapUpdate0 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostname

`func (o *LdapUpdate0) GetHostname() []interface{}`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *LdapUpdate0) GetHostnameOk() (*[]interface{}, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *LdapUpdate0) SetHostname(v []interface{})`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *LdapUpdate0) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetBasedn

`func (o *LdapUpdate0) GetBasedn() string`

GetBasedn returns the Basedn field if non-nil, zero value otherwise.

### GetBasednOk

`func (o *LdapUpdate0) GetBasednOk() (*string, bool)`

GetBasednOk returns a tuple with the Basedn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasedn

`func (o *LdapUpdate0) SetBasedn(v string)`

SetBasedn sets Basedn field to given value.

### HasBasedn

`func (o *LdapUpdate0) HasBasedn() bool`

HasBasedn returns a boolean if a field has been set.

### GetBinddn

`func (o *LdapUpdate0) GetBinddn() string`

GetBinddn returns the Binddn field if non-nil, zero value otherwise.

### GetBinddnOk

`func (o *LdapUpdate0) GetBinddnOk() (*string, bool)`

GetBinddnOk returns a tuple with the Binddn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinddn

`func (o *LdapUpdate0) SetBinddn(v string)`

SetBinddn sets Binddn field to given value.

### HasBinddn

`func (o *LdapUpdate0) HasBinddn() bool`

HasBinddn returns a boolean if a field has been set.

### GetBindpw

`func (o *LdapUpdate0) GetBindpw() string`

GetBindpw returns the Bindpw field if non-nil, zero value otherwise.

### GetBindpwOk

`func (o *LdapUpdate0) GetBindpwOk() (*string, bool)`

GetBindpwOk returns a tuple with the Bindpw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindpw

`func (o *LdapUpdate0) SetBindpw(v string)`

SetBindpw sets Bindpw field to given value.

### HasBindpw

`func (o *LdapUpdate0) HasBindpw() bool`

HasBindpw returns a boolean if a field has been set.

### GetAnonbind

`func (o *LdapUpdate0) GetAnonbind() bool`

GetAnonbind returns the Anonbind field if non-nil, zero value otherwise.

### GetAnonbindOk

`func (o *LdapUpdate0) GetAnonbindOk() (*bool, bool)`

GetAnonbindOk returns a tuple with the Anonbind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnonbind

`func (o *LdapUpdate0) SetAnonbind(v bool)`

SetAnonbind sets Anonbind field to given value.

### HasAnonbind

`func (o *LdapUpdate0) HasAnonbind() bool`

HasAnonbind returns a boolean if a field has been set.

### GetSsl

`func (o *LdapUpdate0) GetSsl() string`

GetSsl returns the Ssl field if non-nil, zero value otherwise.

### GetSslOk

`func (o *LdapUpdate0) GetSslOk() (*string, bool)`

GetSslOk returns a tuple with the Ssl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsl

`func (o *LdapUpdate0) SetSsl(v string)`

SetSsl sets Ssl field to given value.

### HasSsl

`func (o *LdapUpdate0) HasSsl() bool`

HasSsl returns a boolean if a field has been set.

### GetCertificate

`func (o *LdapUpdate0) GetCertificate() int32`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *LdapUpdate0) GetCertificateOk() (*int32, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *LdapUpdate0) SetCertificate(v int32)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *LdapUpdate0) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### SetCertificateNil

`func (o *LdapUpdate0) SetCertificateNil(b bool)`

 SetCertificateNil sets the value for Certificate to be an explicit nil

### UnsetCertificate
`func (o *LdapUpdate0) UnsetCertificate()`

UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil
### GetValidateCertificates

`func (o *LdapUpdate0) GetValidateCertificates() bool`

GetValidateCertificates returns the ValidateCertificates field if non-nil, zero value otherwise.

### GetValidateCertificatesOk

`func (o *LdapUpdate0) GetValidateCertificatesOk() (*bool, bool)`

GetValidateCertificatesOk returns a tuple with the ValidateCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidateCertificates

`func (o *LdapUpdate0) SetValidateCertificates(v bool)`

SetValidateCertificates sets ValidateCertificates field to given value.

### HasValidateCertificates

`func (o *LdapUpdate0) HasValidateCertificates() bool`

HasValidateCertificates returns a boolean if a field has been set.

### GetDisableFreenasCache

`func (o *LdapUpdate0) GetDisableFreenasCache() bool`

GetDisableFreenasCache returns the DisableFreenasCache field if non-nil, zero value otherwise.

### GetDisableFreenasCacheOk

`func (o *LdapUpdate0) GetDisableFreenasCacheOk() (*bool, bool)`

GetDisableFreenasCacheOk returns a tuple with the DisableFreenasCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableFreenasCache

`func (o *LdapUpdate0) SetDisableFreenasCache(v bool)`

SetDisableFreenasCache sets DisableFreenasCache field to given value.

### HasDisableFreenasCache

`func (o *LdapUpdate0) HasDisableFreenasCache() bool`

HasDisableFreenasCache returns a boolean if a field has been set.

### GetTimeout

`func (o *LdapUpdate0) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *LdapUpdate0) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *LdapUpdate0) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *LdapUpdate0) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetDnsTimeout

`func (o *LdapUpdate0) GetDnsTimeout() int32`

GetDnsTimeout returns the DnsTimeout field if non-nil, zero value otherwise.

### GetDnsTimeoutOk

`func (o *LdapUpdate0) GetDnsTimeoutOk() (*int32, bool)`

GetDnsTimeoutOk returns a tuple with the DnsTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsTimeout

`func (o *LdapUpdate0) SetDnsTimeout(v int32)`

SetDnsTimeout sets DnsTimeout field to given value.

### HasDnsTimeout

`func (o *LdapUpdate0) HasDnsTimeout() bool`

HasDnsTimeout returns a boolean if a field has been set.

### GetKerberosRealm

`func (o *LdapUpdate0) GetKerberosRealm() int32`

GetKerberosRealm returns the KerberosRealm field if non-nil, zero value otherwise.

### GetKerberosRealmOk

`func (o *LdapUpdate0) GetKerberosRealmOk() (*int32, bool)`

GetKerberosRealmOk returns a tuple with the KerberosRealm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKerberosRealm

`func (o *LdapUpdate0) SetKerberosRealm(v int32)`

SetKerberosRealm sets KerberosRealm field to given value.

### HasKerberosRealm

`func (o *LdapUpdate0) HasKerberosRealm() bool`

HasKerberosRealm returns a boolean if a field has been set.

### SetKerberosRealmNil

`func (o *LdapUpdate0) SetKerberosRealmNil(b bool)`

 SetKerberosRealmNil sets the value for KerberosRealm to be an explicit nil

### UnsetKerberosRealm
`func (o *LdapUpdate0) UnsetKerberosRealm()`

UnsetKerberosRealm ensures that no value is present for KerberosRealm, not even an explicit nil
### GetKerberosPrincipal

`func (o *LdapUpdate0) GetKerberosPrincipal() string`

GetKerberosPrincipal returns the KerberosPrincipal field if non-nil, zero value otherwise.

### GetKerberosPrincipalOk

`func (o *LdapUpdate0) GetKerberosPrincipalOk() (*string, bool)`

GetKerberosPrincipalOk returns a tuple with the KerberosPrincipal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKerberosPrincipal

`func (o *LdapUpdate0) SetKerberosPrincipal(v string)`

SetKerberosPrincipal sets KerberosPrincipal field to given value.

### HasKerberosPrincipal

`func (o *LdapUpdate0) HasKerberosPrincipal() bool`

HasKerberosPrincipal returns a boolean if a field has been set.

### GetHasSambaSchema

`func (o *LdapUpdate0) GetHasSambaSchema() bool`

GetHasSambaSchema returns the HasSambaSchema field if non-nil, zero value otherwise.

### GetHasSambaSchemaOk

`func (o *LdapUpdate0) GetHasSambaSchemaOk() (*bool, bool)`

GetHasSambaSchemaOk returns a tuple with the HasSambaSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSambaSchema

`func (o *LdapUpdate0) SetHasSambaSchema(v bool)`

SetHasSambaSchema sets HasSambaSchema field to given value.

### HasHasSambaSchema

`func (o *LdapUpdate0) HasHasSambaSchema() bool`

HasHasSambaSchema returns a boolean if a field has been set.

### GetAuxiliaryParameters

`func (o *LdapUpdate0) GetAuxiliaryParameters() string`

GetAuxiliaryParameters returns the AuxiliaryParameters field if non-nil, zero value otherwise.

### GetAuxiliaryParametersOk

`func (o *LdapUpdate0) GetAuxiliaryParametersOk() (*string, bool)`

GetAuxiliaryParametersOk returns a tuple with the AuxiliaryParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuxiliaryParameters

`func (o *LdapUpdate0) SetAuxiliaryParameters(v string)`

SetAuxiliaryParameters sets AuxiliaryParameters field to given value.

### HasAuxiliaryParameters

`func (o *LdapUpdate0) HasAuxiliaryParameters() bool`

HasAuxiliaryParameters returns a boolean if a field has been set.

### GetSchema

`func (o *LdapUpdate0) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *LdapUpdate0) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *LdapUpdate0) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *LdapUpdate0) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetEnable

`func (o *LdapUpdate0) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *LdapUpdate0) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *LdapUpdate0) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *LdapUpdate0) HasEnable() bool`

HasEnable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


