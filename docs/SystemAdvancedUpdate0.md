# SystemAdvancedUpdate0

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Advancedmode** | Pointer to **bool** |  | [optional] 
**Autotune** | Pointer to **bool** |  | [optional] 
**BootScrub** | Pointer to **int32** |  | [optional] 
**Consolemenu** | Pointer to **bool** |  | [optional] 
**Consolemsg** | Pointer to **bool** |  | [optional] 
**Debugkernel** | Pointer to **bool** |  | [optional] 
**FqdnSyslog** | Pointer to **bool** |  | [optional] 
**Motd** | Pointer to **string** |  | [optional] 
**Powerdaemon** | Pointer to **bool** |  | [optional] 
**Serialconsole** | Pointer to **bool** |  | [optional] 
**Serialport** | Pointer to **string** |  | [optional] 
**Serialspeed** | Pointer to **string** |  | [optional] 
**Swapondrive** | Pointer to **int32** |  | [optional] 
**Overprovision** | Pointer to **NullableInt32** |  | [optional] 
**Traceback** | Pointer to **bool** |  | [optional] 
**Uploadcrash** | Pointer to **bool** |  | [optional] 
**Anonstats** | Pointer to **bool** |  | [optional] 
**SedUser** | Pointer to **string** |  | [optional] 
**SedPasswd** | Pointer to **string** |  | [optional] 
**Sysloglevel** | Pointer to **string** |  | [optional] 
**Syslogserver** | Pointer to **string** |  | [optional] 
**SyslogTransport** | Pointer to **string** |  | [optional] 
**SyslogTlsCertificate** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewSystemAdvancedUpdate0

`func NewSystemAdvancedUpdate0() *SystemAdvancedUpdate0`

NewSystemAdvancedUpdate0 instantiates a new SystemAdvancedUpdate0 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemAdvancedUpdate0WithDefaults

`func NewSystemAdvancedUpdate0WithDefaults() *SystemAdvancedUpdate0`

NewSystemAdvancedUpdate0WithDefaults instantiates a new SystemAdvancedUpdate0 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdvancedmode

`func (o *SystemAdvancedUpdate0) GetAdvancedmode() bool`

GetAdvancedmode returns the Advancedmode field if non-nil, zero value otherwise.

### GetAdvancedmodeOk

`func (o *SystemAdvancedUpdate0) GetAdvancedmodeOk() (*bool, bool)`

GetAdvancedmodeOk returns a tuple with the Advancedmode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedmode

`func (o *SystemAdvancedUpdate0) SetAdvancedmode(v bool)`

SetAdvancedmode sets Advancedmode field to given value.

### HasAdvancedmode

`func (o *SystemAdvancedUpdate0) HasAdvancedmode() bool`

HasAdvancedmode returns a boolean if a field has been set.

### GetAutotune

`func (o *SystemAdvancedUpdate0) GetAutotune() bool`

GetAutotune returns the Autotune field if non-nil, zero value otherwise.

### GetAutotuneOk

`func (o *SystemAdvancedUpdate0) GetAutotuneOk() (*bool, bool)`

GetAutotuneOk returns a tuple with the Autotune field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutotune

`func (o *SystemAdvancedUpdate0) SetAutotune(v bool)`

SetAutotune sets Autotune field to given value.

### HasAutotune

`func (o *SystemAdvancedUpdate0) HasAutotune() bool`

HasAutotune returns a boolean if a field has been set.

### GetBootScrub

`func (o *SystemAdvancedUpdate0) GetBootScrub() int32`

GetBootScrub returns the BootScrub field if non-nil, zero value otherwise.

### GetBootScrubOk

`func (o *SystemAdvancedUpdate0) GetBootScrubOk() (*int32, bool)`

GetBootScrubOk returns a tuple with the BootScrub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootScrub

`func (o *SystemAdvancedUpdate0) SetBootScrub(v int32)`

SetBootScrub sets BootScrub field to given value.

### HasBootScrub

`func (o *SystemAdvancedUpdate0) HasBootScrub() bool`

HasBootScrub returns a boolean if a field has been set.

### GetConsolemenu

`func (o *SystemAdvancedUpdate0) GetConsolemenu() bool`

GetConsolemenu returns the Consolemenu field if non-nil, zero value otherwise.

### GetConsolemenuOk

`func (o *SystemAdvancedUpdate0) GetConsolemenuOk() (*bool, bool)`

GetConsolemenuOk returns a tuple with the Consolemenu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsolemenu

`func (o *SystemAdvancedUpdate0) SetConsolemenu(v bool)`

SetConsolemenu sets Consolemenu field to given value.

### HasConsolemenu

`func (o *SystemAdvancedUpdate0) HasConsolemenu() bool`

HasConsolemenu returns a boolean if a field has been set.

### GetConsolemsg

`func (o *SystemAdvancedUpdate0) GetConsolemsg() bool`

GetConsolemsg returns the Consolemsg field if non-nil, zero value otherwise.

### GetConsolemsgOk

`func (o *SystemAdvancedUpdate0) GetConsolemsgOk() (*bool, bool)`

GetConsolemsgOk returns a tuple with the Consolemsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsolemsg

`func (o *SystemAdvancedUpdate0) SetConsolemsg(v bool)`

SetConsolemsg sets Consolemsg field to given value.

### HasConsolemsg

`func (o *SystemAdvancedUpdate0) HasConsolemsg() bool`

HasConsolemsg returns a boolean if a field has been set.

### GetDebugkernel

`func (o *SystemAdvancedUpdate0) GetDebugkernel() bool`

GetDebugkernel returns the Debugkernel field if non-nil, zero value otherwise.

### GetDebugkernelOk

`func (o *SystemAdvancedUpdate0) GetDebugkernelOk() (*bool, bool)`

GetDebugkernelOk returns a tuple with the Debugkernel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebugkernel

`func (o *SystemAdvancedUpdate0) SetDebugkernel(v bool)`

SetDebugkernel sets Debugkernel field to given value.

### HasDebugkernel

`func (o *SystemAdvancedUpdate0) HasDebugkernel() bool`

HasDebugkernel returns a boolean if a field has been set.

### GetFqdnSyslog

`func (o *SystemAdvancedUpdate0) GetFqdnSyslog() bool`

GetFqdnSyslog returns the FqdnSyslog field if non-nil, zero value otherwise.

### GetFqdnSyslogOk

`func (o *SystemAdvancedUpdate0) GetFqdnSyslogOk() (*bool, bool)`

GetFqdnSyslogOk returns a tuple with the FqdnSyslog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdnSyslog

`func (o *SystemAdvancedUpdate0) SetFqdnSyslog(v bool)`

SetFqdnSyslog sets FqdnSyslog field to given value.

### HasFqdnSyslog

`func (o *SystemAdvancedUpdate0) HasFqdnSyslog() bool`

HasFqdnSyslog returns a boolean if a field has been set.

### GetMotd

`func (o *SystemAdvancedUpdate0) GetMotd() string`

GetMotd returns the Motd field if non-nil, zero value otherwise.

### GetMotdOk

`func (o *SystemAdvancedUpdate0) GetMotdOk() (*string, bool)`

GetMotdOk returns a tuple with the Motd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMotd

`func (o *SystemAdvancedUpdate0) SetMotd(v string)`

SetMotd sets Motd field to given value.

### HasMotd

`func (o *SystemAdvancedUpdate0) HasMotd() bool`

HasMotd returns a boolean if a field has been set.

### GetPowerdaemon

`func (o *SystemAdvancedUpdate0) GetPowerdaemon() bool`

GetPowerdaemon returns the Powerdaemon field if non-nil, zero value otherwise.

### GetPowerdaemonOk

`func (o *SystemAdvancedUpdate0) GetPowerdaemonOk() (*bool, bool)`

GetPowerdaemonOk returns a tuple with the Powerdaemon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerdaemon

`func (o *SystemAdvancedUpdate0) SetPowerdaemon(v bool)`

SetPowerdaemon sets Powerdaemon field to given value.

### HasPowerdaemon

`func (o *SystemAdvancedUpdate0) HasPowerdaemon() bool`

HasPowerdaemon returns a boolean if a field has been set.

### GetSerialconsole

`func (o *SystemAdvancedUpdate0) GetSerialconsole() bool`

GetSerialconsole returns the Serialconsole field if non-nil, zero value otherwise.

### GetSerialconsoleOk

`func (o *SystemAdvancedUpdate0) GetSerialconsoleOk() (*bool, bool)`

GetSerialconsoleOk returns a tuple with the Serialconsole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialconsole

`func (o *SystemAdvancedUpdate0) SetSerialconsole(v bool)`

SetSerialconsole sets Serialconsole field to given value.

### HasSerialconsole

`func (o *SystemAdvancedUpdate0) HasSerialconsole() bool`

HasSerialconsole returns a boolean if a field has been set.

### GetSerialport

`func (o *SystemAdvancedUpdate0) GetSerialport() string`

GetSerialport returns the Serialport field if non-nil, zero value otherwise.

### GetSerialportOk

`func (o *SystemAdvancedUpdate0) GetSerialportOk() (*string, bool)`

GetSerialportOk returns a tuple with the Serialport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialport

`func (o *SystemAdvancedUpdate0) SetSerialport(v string)`

SetSerialport sets Serialport field to given value.

### HasSerialport

`func (o *SystemAdvancedUpdate0) HasSerialport() bool`

HasSerialport returns a boolean if a field has been set.

### GetSerialspeed

`func (o *SystemAdvancedUpdate0) GetSerialspeed() string`

GetSerialspeed returns the Serialspeed field if non-nil, zero value otherwise.

### GetSerialspeedOk

`func (o *SystemAdvancedUpdate0) GetSerialspeedOk() (*string, bool)`

GetSerialspeedOk returns a tuple with the Serialspeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialspeed

`func (o *SystemAdvancedUpdate0) SetSerialspeed(v string)`

SetSerialspeed sets Serialspeed field to given value.

### HasSerialspeed

`func (o *SystemAdvancedUpdate0) HasSerialspeed() bool`

HasSerialspeed returns a boolean if a field has been set.

### GetSwapondrive

`func (o *SystemAdvancedUpdate0) GetSwapondrive() int32`

GetSwapondrive returns the Swapondrive field if non-nil, zero value otherwise.

### GetSwapondriveOk

`func (o *SystemAdvancedUpdate0) GetSwapondriveOk() (*int32, bool)`

GetSwapondriveOk returns a tuple with the Swapondrive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwapondrive

`func (o *SystemAdvancedUpdate0) SetSwapondrive(v int32)`

SetSwapondrive sets Swapondrive field to given value.

### HasSwapondrive

`func (o *SystemAdvancedUpdate0) HasSwapondrive() bool`

HasSwapondrive returns a boolean if a field has been set.

### GetOverprovision

`func (o *SystemAdvancedUpdate0) GetOverprovision() int32`

GetOverprovision returns the Overprovision field if non-nil, zero value otherwise.

### GetOverprovisionOk

`func (o *SystemAdvancedUpdate0) GetOverprovisionOk() (*int32, bool)`

GetOverprovisionOk returns a tuple with the Overprovision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverprovision

`func (o *SystemAdvancedUpdate0) SetOverprovision(v int32)`

SetOverprovision sets Overprovision field to given value.

### HasOverprovision

`func (o *SystemAdvancedUpdate0) HasOverprovision() bool`

HasOverprovision returns a boolean if a field has been set.

### SetOverprovisionNil

`func (o *SystemAdvancedUpdate0) SetOverprovisionNil(b bool)`

 SetOverprovisionNil sets the value for Overprovision to be an explicit nil

### UnsetOverprovision
`func (o *SystemAdvancedUpdate0) UnsetOverprovision()`

UnsetOverprovision ensures that no value is present for Overprovision, not even an explicit nil
### GetTraceback

`func (o *SystemAdvancedUpdate0) GetTraceback() bool`

GetTraceback returns the Traceback field if non-nil, zero value otherwise.

### GetTracebackOk

`func (o *SystemAdvancedUpdate0) GetTracebackOk() (*bool, bool)`

GetTracebackOk returns a tuple with the Traceback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceback

`func (o *SystemAdvancedUpdate0) SetTraceback(v bool)`

SetTraceback sets Traceback field to given value.

### HasTraceback

`func (o *SystemAdvancedUpdate0) HasTraceback() bool`

HasTraceback returns a boolean if a field has been set.

### GetUploadcrash

`func (o *SystemAdvancedUpdate0) GetUploadcrash() bool`

GetUploadcrash returns the Uploadcrash field if non-nil, zero value otherwise.

### GetUploadcrashOk

`func (o *SystemAdvancedUpdate0) GetUploadcrashOk() (*bool, bool)`

GetUploadcrashOk returns a tuple with the Uploadcrash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadcrash

`func (o *SystemAdvancedUpdate0) SetUploadcrash(v bool)`

SetUploadcrash sets Uploadcrash field to given value.

### HasUploadcrash

`func (o *SystemAdvancedUpdate0) HasUploadcrash() bool`

HasUploadcrash returns a boolean if a field has been set.

### GetAnonstats

`func (o *SystemAdvancedUpdate0) GetAnonstats() bool`

GetAnonstats returns the Anonstats field if non-nil, zero value otherwise.

### GetAnonstatsOk

`func (o *SystemAdvancedUpdate0) GetAnonstatsOk() (*bool, bool)`

GetAnonstatsOk returns a tuple with the Anonstats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnonstats

`func (o *SystemAdvancedUpdate0) SetAnonstats(v bool)`

SetAnonstats sets Anonstats field to given value.

### HasAnonstats

`func (o *SystemAdvancedUpdate0) HasAnonstats() bool`

HasAnonstats returns a boolean if a field has been set.

### GetSedUser

`func (o *SystemAdvancedUpdate0) GetSedUser() string`

GetSedUser returns the SedUser field if non-nil, zero value otherwise.

### GetSedUserOk

`func (o *SystemAdvancedUpdate0) GetSedUserOk() (*string, bool)`

GetSedUserOk returns a tuple with the SedUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSedUser

`func (o *SystemAdvancedUpdate0) SetSedUser(v string)`

SetSedUser sets SedUser field to given value.

### HasSedUser

`func (o *SystemAdvancedUpdate0) HasSedUser() bool`

HasSedUser returns a boolean if a field has been set.

### GetSedPasswd

`func (o *SystemAdvancedUpdate0) GetSedPasswd() string`

GetSedPasswd returns the SedPasswd field if non-nil, zero value otherwise.

### GetSedPasswdOk

`func (o *SystemAdvancedUpdate0) GetSedPasswdOk() (*string, bool)`

GetSedPasswdOk returns a tuple with the SedPasswd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSedPasswd

`func (o *SystemAdvancedUpdate0) SetSedPasswd(v string)`

SetSedPasswd sets SedPasswd field to given value.

### HasSedPasswd

`func (o *SystemAdvancedUpdate0) HasSedPasswd() bool`

HasSedPasswd returns a boolean if a field has been set.

### GetSysloglevel

`func (o *SystemAdvancedUpdate0) GetSysloglevel() string`

GetSysloglevel returns the Sysloglevel field if non-nil, zero value otherwise.

### GetSysloglevelOk

`func (o *SystemAdvancedUpdate0) GetSysloglevelOk() (*string, bool)`

GetSysloglevelOk returns a tuple with the Sysloglevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSysloglevel

`func (o *SystemAdvancedUpdate0) SetSysloglevel(v string)`

SetSysloglevel sets Sysloglevel field to given value.

### HasSysloglevel

`func (o *SystemAdvancedUpdate0) HasSysloglevel() bool`

HasSysloglevel returns a boolean if a field has been set.

### GetSyslogserver

`func (o *SystemAdvancedUpdate0) GetSyslogserver() string`

GetSyslogserver returns the Syslogserver field if non-nil, zero value otherwise.

### GetSyslogserverOk

`func (o *SystemAdvancedUpdate0) GetSyslogserverOk() (*string, bool)`

GetSyslogserverOk returns a tuple with the Syslogserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogserver

`func (o *SystemAdvancedUpdate0) SetSyslogserver(v string)`

SetSyslogserver sets Syslogserver field to given value.

### HasSyslogserver

`func (o *SystemAdvancedUpdate0) HasSyslogserver() bool`

HasSyslogserver returns a boolean if a field has been set.

### GetSyslogTransport

`func (o *SystemAdvancedUpdate0) GetSyslogTransport() string`

GetSyslogTransport returns the SyslogTransport field if non-nil, zero value otherwise.

### GetSyslogTransportOk

`func (o *SystemAdvancedUpdate0) GetSyslogTransportOk() (*string, bool)`

GetSyslogTransportOk returns a tuple with the SyslogTransport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogTransport

`func (o *SystemAdvancedUpdate0) SetSyslogTransport(v string)`

SetSyslogTransport sets SyslogTransport field to given value.

### HasSyslogTransport

`func (o *SystemAdvancedUpdate0) HasSyslogTransport() bool`

HasSyslogTransport returns a boolean if a field has been set.

### GetSyslogTlsCertificate

`func (o *SystemAdvancedUpdate0) GetSyslogTlsCertificate() int32`

GetSyslogTlsCertificate returns the SyslogTlsCertificate field if non-nil, zero value otherwise.

### GetSyslogTlsCertificateOk

`func (o *SystemAdvancedUpdate0) GetSyslogTlsCertificateOk() (*int32, bool)`

GetSyslogTlsCertificateOk returns a tuple with the SyslogTlsCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogTlsCertificate

`func (o *SystemAdvancedUpdate0) SetSyslogTlsCertificate(v int32)`

SetSyslogTlsCertificate sets SyslogTlsCertificate field to given value.

### HasSyslogTlsCertificate

`func (o *SystemAdvancedUpdate0) HasSyslogTlsCertificate() bool`

HasSyslogTlsCertificate returns a boolean if a field has been set.

### SetSyslogTlsCertificateNil

`func (o *SystemAdvancedUpdate0) SetSyslogTlsCertificateNil(b bool)`

 SetSyslogTlsCertificateNil sets the value for SyslogTlsCertificate to be an explicit nil

### UnsetSyslogTlsCertificate
`func (o *SystemAdvancedUpdate0) UnsetSyslogTlsCertificate()`

UnsetSyslogTlsCertificate ensures that no value is present for SyslogTlsCertificate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


