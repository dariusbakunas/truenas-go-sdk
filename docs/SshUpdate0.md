# SshUpdate0

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bindiface** | Pointer to **[]string** |  | [optional] 
**Tcpport** | Pointer to **int32** |  | [optional] 
**Rootlogin** | Pointer to **bool** |  | [optional] 
**Passwordauth** | Pointer to **bool** |  | [optional] 
**Kerberosauth** | Pointer to **bool** |  | [optional] 
**Tcpfwd** | Pointer to **bool** |  | [optional] 
**Compression** | Pointer to **bool** |  | [optional] 
**SftpLogLevel** | Pointer to **string** |  | [optional] 
**SftpLogFacility** | Pointer to **string** |  | [optional] 
**WeakCiphers** | Pointer to **[]string** |  | [optional] 
**Options** | Pointer to **string** |  | [optional] 

## Methods

### NewSshUpdate0

`func NewSshUpdate0() *SshUpdate0`

NewSshUpdate0 instantiates a new SshUpdate0 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSshUpdate0WithDefaults

`func NewSshUpdate0WithDefaults() *SshUpdate0`

NewSshUpdate0WithDefaults instantiates a new SshUpdate0 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBindiface

`func (o *SshUpdate0) GetBindiface() []string`

GetBindiface returns the Bindiface field if non-nil, zero value otherwise.

### GetBindifaceOk

`func (o *SshUpdate0) GetBindifaceOk() (*[]string, bool)`

GetBindifaceOk returns a tuple with the Bindiface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindiface

`func (o *SshUpdate0) SetBindiface(v []string)`

SetBindiface sets Bindiface field to given value.

### HasBindiface

`func (o *SshUpdate0) HasBindiface() bool`

HasBindiface returns a boolean if a field has been set.

### GetTcpport

`func (o *SshUpdate0) GetTcpport() int32`

GetTcpport returns the Tcpport field if non-nil, zero value otherwise.

### GetTcpportOk

`func (o *SshUpdate0) GetTcpportOk() (*int32, bool)`

GetTcpportOk returns a tuple with the Tcpport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpport

`func (o *SshUpdate0) SetTcpport(v int32)`

SetTcpport sets Tcpport field to given value.

### HasTcpport

`func (o *SshUpdate0) HasTcpport() bool`

HasTcpport returns a boolean if a field has been set.

### GetRootlogin

`func (o *SshUpdate0) GetRootlogin() bool`

GetRootlogin returns the Rootlogin field if non-nil, zero value otherwise.

### GetRootloginOk

`func (o *SshUpdate0) GetRootloginOk() (*bool, bool)`

GetRootloginOk returns a tuple with the Rootlogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootlogin

`func (o *SshUpdate0) SetRootlogin(v bool)`

SetRootlogin sets Rootlogin field to given value.

### HasRootlogin

`func (o *SshUpdate0) HasRootlogin() bool`

HasRootlogin returns a boolean if a field has been set.

### GetPasswordauth

`func (o *SshUpdate0) GetPasswordauth() bool`

GetPasswordauth returns the Passwordauth field if non-nil, zero value otherwise.

### GetPasswordauthOk

`func (o *SshUpdate0) GetPasswordauthOk() (*bool, bool)`

GetPasswordauthOk returns a tuple with the Passwordauth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordauth

`func (o *SshUpdate0) SetPasswordauth(v bool)`

SetPasswordauth sets Passwordauth field to given value.

### HasPasswordauth

`func (o *SshUpdate0) HasPasswordauth() bool`

HasPasswordauth returns a boolean if a field has been set.

### GetKerberosauth

`func (o *SshUpdate0) GetKerberosauth() bool`

GetKerberosauth returns the Kerberosauth field if non-nil, zero value otherwise.

### GetKerberosauthOk

`func (o *SshUpdate0) GetKerberosauthOk() (*bool, bool)`

GetKerberosauthOk returns a tuple with the Kerberosauth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKerberosauth

`func (o *SshUpdate0) SetKerberosauth(v bool)`

SetKerberosauth sets Kerberosauth field to given value.

### HasKerberosauth

`func (o *SshUpdate0) HasKerberosauth() bool`

HasKerberosauth returns a boolean if a field has been set.

### GetTcpfwd

`func (o *SshUpdate0) GetTcpfwd() bool`

GetTcpfwd returns the Tcpfwd field if non-nil, zero value otherwise.

### GetTcpfwdOk

`func (o *SshUpdate0) GetTcpfwdOk() (*bool, bool)`

GetTcpfwdOk returns a tuple with the Tcpfwd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpfwd

`func (o *SshUpdate0) SetTcpfwd(v bool)`

SetTcpfwd sets Tcpfwd field to given value.

### HasTcpfwd

`func (o *SshUpdate0) HasTcpfwd() bool`

HasTcpfwd returns a boolean if a field has been set.

### GetCompression

`func (o *SshUpdate0) GetCompression() bool`

GetCompression returns the Compression field if non-nil, zero value otherwise.

### GetCompressionOk

`func (o *SshUpdate0) GetCompressionOk() (*bool, bool)`

GetCompressionOk returns a tuple with the Compression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompression

`func (o *SshUpdate0) SetCompression(v bool)`

SetCompression sets Compression field to given value.

### HasCompression

`func (o *SshUpdate0) HasCompression() bool`

HasCompression returns a boolean if a field has been set.

### GetSftpLogLevel

`func (o *SshUpdate0) GetSftpLogLevel() string`

GetSftpLogLevel returns the SftpLogLevel field if non-nil, zero value otherwise.

### GetSftpLogLevelOk

`func (o *SshUpdate0) GetSftpLogLevelOk() (*string, bool)`

GetSftpLogLevelOk returns a tuple with the SftpLogLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSftpLogLevel

`func (o *SshUpdate0) SetSftpLogLevel(v string)`

SetSftpLogLevel sets SftpLogLevel field to given value.

### HasSftpLogLevel

`func (o *SshUpdate0) HasSftpLogLevel() bool`

HasSftpLogLevel returns a boolean if a field has been set.

### GetSftpLogFacility

`func (o *SshUpdate0) GetSftpLogFacility() string`

GetSftpLogFacility returns the SftpLogFacility field if non-nil, zero value otherwise.

### GetSftpLogFacilityOk

`func (o *SshUpdate0) GetSftpLogFacilityOk() (*string, bool)`

GetSftpLogFacilityOk returns a tuple with the SftpLogFacility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSftpLogFacility

`func (o *SshUpdate0) SetSftpLogFacility(v string)`

SetSftpLogFacility sets SftpLogFacility field to given value.

### HasSftpLogFacility

`func (o *SshUpdate0) HasSftpLogFacility() bool`

HasSftpLogFacility returns a boolean if a field has been set.

### GetWeakCiphers

`func (o *SshUpdate0) GetWeakCiphers() []string`

GetWeakCiphers returns the WeakCiphers field if non-nil, zero value otherwise.

### GetWeakCiphersOk

`func (o *SshUpdate0) GetWeakCiphersOk() (*[]string, bool)`

GetWeakCiphersOk returns a tuple with the WeakCiphers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeakCiphers

`func (o *SshUpdate0) SetWeakCiphers(v []string)`

SetWeakCiphers sets WeakCiphers field to given value.

### HasWeakCiphers

`func (o *SshUpdate0) HasWeakCiphers() bool`

HasWeakCiphers returns a boolean if a field has been set.

### GetOptions

`func (o *SshUpdate0) GetOptions() string`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *SshUpdate0) GetOptionsOk() (*string, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *SshUpdate0) SetOptions(v string)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *SshUpdate0) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


