# UpsUpdate0

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Emailnotify** | Pointer to **bool** |  | [optional] 
**Powerdown** | Pointer to **bool** |  | [optional] 
**Rmonitor** | Pointer to **bool** |  | [optional] 
**Nocommwarntime** | Pointer to **NullableInt32** |  | [optional] 
**Remoteport** | Pointer to **int32** |  | [optional] 
**Shutdowntimer** | Pointer to **int32** |  | [optional] 
**Hostsync** | Pointer to **int32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Driver** | Pointer to **string** |  | [optional] 
**Extrausers** | Pointer to **string** |  | [optional] 
**Identifier** | Pointer to **string** |  | [optional] 
**Mode** | Pointer to **string** |  | [optional] 
**Monpwd** | Pointer to **string** |  | [optional] 
**Monuser** | Pointer to **string** |  | [optional] 
**Options** | Pointer to **string** |  | [optional] 
**Optionsupsd** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **string** |  | [optional] 
**Remotehost** | Pointer to **string** |  | [optional] 
**Shutdown** | Pointer to **string** |  | [optional] 
**Shutdowncmd** | Pointer to **NullableString** |  | [optional] 
**Subject** | Pointer to **string** |  | [optional] 
**Toemail** | Pointer to **[]string** |  | [optional] 

## Methods

### NewUpsUpdate0

`func NewUpsUpdate0() *UpsUpdate0`

NewUpsUpdate0 instantiates a new UpsUpdate0 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpsUpdate0WithDefaults

`func NewUpsUpdate0WithDefaults() *UpsUpdate0`

NewUpsUpdate0WithDefaults instantiates a new UpsUpdate0 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailnotify

`func (o *UpsUpdate0) GetEmailnotify() bool`

GetEmailnotify returns the Emailnotify field if non-nil, zero value otherwise.

### GetEmailnotifyOk

`func (o *UpsUpdate0) GetEmailnotifyOk() (*bool, bool)`

GetEmailnotifyOk returns a tuple with the Emailnotify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailnotify

`func (o *UpsUpdate0) SetEmailnotify(v bool)`

SetEmailnotify sets Emailnotify field to given value.

### HasEmailnotify

`func (o *UpsUpdate0) HasEmailnotify() bool`

HasEmailnotify returns a boolean if a field has been set.

### GetPowerdown

`func (o *UpsUpdate0) GetPowerdown() bool`

GetPowerdown returns the Powerdown field if non-nil, zero value otherwise.

### GetPowerdownOk

`func (o *UpsUpdate0) GetPowerdownOk() (*bool, bool)`

GetPowerdownOk returns a tuple with the Powerdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerdown

`func (o *UpsUpdate0) SetPowerdown(v bool)`

SetPowerdown sets Powerdown field to given value.

### HasPowerdown

`func (o *UpsUpdate0) HasPowerdown() bool`

HasPowerdown returns a boolean if a field has been set.

### GetRmonitor

`func (o *UpsUpdate0) GetRmonitor() bool`

GetRmonitor returns the Rmonitor field if non-nil, zero value otherwise.

### GetRmonitorOk

`func (o *UpsUpdate0) GetRmonitorOk() (*bool, bool)`

GetRmonitorOk returns a tuple with the Rmonitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRmonitor

`func (o *UpsUpdate0) SetRmonitor(v bool)`

SetRmonitor sets Rmonitor field to given value.

### HasRmonitor

`func (o *UpsUpdate0) HasRmonitor() bool`

HasRmonitor returns a boolean if a field has been set.

### GetNocommwarntime

`func (o *UpsUpdate0) GetNocommwarntime() int32`

GetNocommwarntime returns the Nocommwarntime field if non-nil, zero value otherwise.

### GetNocommwarntimeOk

`func (o *UpsUpdate0) GetNocommwarntimeOk() (*int32, bool)`

GetNocommwarntimeOk returns a tuple with the Nocommwarntime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNocommwarntime

`func (o *UpsUpdate0) SetNocommwarntime(v int32)`

SetNocommwarntime sets Nocommwarntime field to given value.

### HasNocommwarntime

`func (o *UpsUpdate0) HasNocommwarntime() bool`

HasNocommwarntime returns a boolean if a field has been set.

### SetNocommwarntimeNil

`func (o *UpsUpdate0) SetNocommwarntimeNil(b bool)`

 SetNocommwarntimeNil sets the value for Nocommwarntime to be an explicit nil

### UnsetNocommwarntime
`func (o *UpsUpdate0) UnsetNocommwarntime()`

UnsetNocommwarntime ensures that no value is present for Nocommwarntime, not even an explicit nil
### GetRemoteport

`func (o *UpsUpdate0) GetRemoteport() int32`

GetRemoteport returns the Remoteport field if non-nil, zero value otherwise.

### GetRemoteportOk

`func (o *UpsUpdate0) GetRemoteportOk() (*int32, bool)`

GetRemoteportOk returns a tuple with the Remoteport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteport

`func (o *UpsUpdate0) SetRemoteport(v int32)`

SetRemoteport sets Remoteport field to given value.

### HasRemoteport

`func (o *UpsUpdate0) HasRemoteport() bool`

HasRemoteport returns a boolean if a field has been set.

### GetShutdowntimer

`func (o *UpsUpdate0) GetShutdowntimer() int32`

GetShutdowntimer returns the Shutdowntimer field if non-nil, zero value otherwise.

### GetShutdowntimerOk

`func (o *UpsUpdate0) GetShutdowntimerOk() (*int32, bool)`

GetShutdowntimerOk returns a tuple with the Shutdowntimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdowntimer

`func (o *UpsUpdate0) SetShutdowntimer(v int32)`

SetShutdowntimer sets Shutdowntimer field to given value.

### HasShutdowntimer

`func (o *UpsUpdate0) HasShutdowntimer() bool`

HasShutdowntimer returns a boolean if a field has been set.

### GetHostsync

`func (o *UpsUpdate0) GetHostsync() int32`

GetHostsync returns the Hostsync field if non-nil, zero value otherwise.

### GetHostsyncOk

`func (o *UpsUpdate0) GetHostsyncOk() (*int32, bool)`

GetHostsyncOk returns a tuple with the Hostsync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostsync

`func (o *UpsUpdate0) SetHostsync(v int32)`

SetHostsync sets Hostsync field to given value.

### HasHostsync

`func (o *UpsUpdate0) HasHostsync() bool`

HasHostsync returns a boolean if a field has been set.

### GetDescription

`func (o *UpsUpdate0) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpsUpdate0) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpsUpdate0) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpsUpdate0) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDriver

`func (o *UpsUpdate0) GetDriver() string`

GetDriver returns the Driver field if non-nil, zero value otherwise.

### GetDriverOk

`func (o *UpsUpdate0) GetDriverOk() (*string, bool)`

GetDriverOk returns a tuple with the Driver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriver

`func (o *UpsUpdate0) SetDriver(v string)`

SetDriver sets Driver field to given value.

### HasDriver

`func (o *UpsUpdate0) HasDriver() bool`

HasDriver returns a boolean if a field has been set.

### GetExtrausers

`func (o *UpsUpdate0) GetExtrausers() string`

GetExtrausers returns the Extrausers field if non-nil, zero value otherwise.

### GetExtrausersOk

`func (o *UpsUpdate0) GetExtrausersOk() (*string, bool)`

GetExtrausersOk returns a tuple with the Extrausers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtrausers

`func (o *UpsUpdate0) SetExtrausers(v string)`

SetExtrausers sets Extrausers field to given value.

### HasExtrausers

`func (o *UpsUpdate0) HasExtrausers() bool`

HasExtrausers returns a boolean if a field has been set.

### GetIdentifier

`func (o *UpsUpdate0) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *UpsUpdate0) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *UpsUpdate0) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *UpsUpdate0) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetMode

`func (o *UpsUpdate0) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *UpsUpdate0) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *UpsUpdate0) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *UpsUpdate0) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetMonpwd

`func (o *UpsUpdate0) GetMonpwd() string`

GetMonpwd returns the Monpwd field if non-nil, zero value otherwise.

### GetMonpwdOk

`func (o *UpsUpdate0) GetMonpwdOk() (*string, bool)`

GetMonpwdOk returns a tuple with the Monpwd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonpwd

`func (o *UpsUpdate0) SetMonpwd(v string)`

SetMonpwd sets Monpwd field to given value.

### HasMonpwd

`func (o *UpsUpdate0) HasMonpwd() bool`

HasMonpwd returns a boolean if a field has been set.

### GetMonuser

`func (o *UpsUpdate0) GetMonuser() string`

GetMonuser returns the Monuser field if non-nil, zero value otherwise.

### GetMonuserOk

`func (o *UpsUpdate0) GetMonuserOk() (*string, bool)`

GetMonuserOk returns a tuple with the Monuser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonuser

`func (o *UpsUpdate0) SetMonuser(v string)`

SetMonuser sets Monuser field to given value.

### HasMonuser

`func (o *UpsUpdate0) HasMonuser() bool`

HasMonuser returns a boolean if a field has been set.

### GetOptions

`func (o *UpsUpdate0) GetOptions() string`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *UpsUpdate0) GetOptionsOk() (*string, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *UpsUpdate0) SetOptions(v string)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *UpsUpdate0) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetOptionsupsd

`func (o *UpsUpdate0) GetOptionsupsd() string`

GetOptionsupsd returns the Optionsupsd field if non-nil, zero value otherwise.

### GetOptionsupsdOk

`func (o *UpsUpdate0) GetOptionsupsdOk() (*string, bool)`

GetOptionsupsdOk returns a tuple with the Optionsupsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionsupsd

`func (o *UpsUpdate0) SetOptionsupsd(v string)`

SetOptionsupsd sets Optionsupsd field to given value.

### HasOptionsupsd

`func (o *UpsUpdate0) HasOptionsupsd() bool`

HasOptionsupsd returns a boolean if a field has been set.

### GetPort

`func (o *UpsUpdate0) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *UpsUpdate0) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *UpsUpdate0) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *UpsUpdate0) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetRemotehost

`func (o *UpsUpdate0) GetRemotehost() string`

GetRemotehost returns the Remotehost field if non-nil, zero value otherwise.

### GetRemotehostOk

`func (o *UpsUpdate0) GetRemotehostOk() (*string, bool)`

GetRemotehostOk returns a tuple with the Remotehost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotehost

`func (o *UpsUpdate0) SetRemotehost(v string)`

SetRemotehost sets Remotehost field to given value.

### HasRemotehost

`func (o *UpsUpdate0) HasRemotehost() bool`

HasRemotehost returns a boolean if a field has been set.

### GetShutdown

`func (o *UpsUpdate0) GetShutdown() string`

GetShutdown returns the Shutdown field if non-nil, zero value otherwise.

### GetShutdownOk

`func (o *UpsUpdate0) GetShutdownOk() (*string, bool)`

GetShutdownOk returns a tuple with the Shutdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdown

`func (o *UpsUpdate0) SetShutdown(v string)`

SetShutdown sets Shutdown field to given value.

### HasShutdown

`func (o *UpsUpdate0) HasShutdown() bool`

HasShutdown returns a boolean if a field has been set.

### GetShutdowncmd

`func (o *UpsUpdate0) GetShutdowncmd() string`

GetShutdowncmd returns the Shutdowncmd field if non-nil, zero value otherwise.

### GetShutdowncmdOk

`func (o *UpsUpdate0) GetShutdowncmdOk() (*string, bool)`

GetShutdowncmdOk returns a tuple with the Shutdowncmd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdowncmd

`func (o *UpsUpdate0) SetShutdowncmd(v string)`

SetShutdowncmd sets Shutdowncmd field to given value.

### HasShutdowncmd

`func (o *UpsUpdate0) HasShutdowncmd() bool`

HasShutdowncmd returns a boolean if a field has been set.

### SetShutdowncmdNil

`func (o *UpsUpdate0) SetShutdowncmdNil(b bool)`

 SetShutdowncmdNil sets the value for Shutdowncmd to be an explicit nil

### UnsetShutdowncmd
`func (o *UpsUpdate0) UnsetShutdowncmd()`

UnsetShutdowncmd ensures that no value is present for Shutdowncmd, not even an explicit nil
### GetSubject

`func (o *UpsUpdate0) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *UpsUpdate0) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *UpsUpdate0) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *UpsUpdate0) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetToemail

`func (o *UpsUpdate0) GetToemail() []string`

GetToemail returns the Toemail field if non-nil, zero value otherwise.

### GetToemailOk

`func (o *UpsUpdate0) GetToemailOk() (*[]string, bool)`

GetToemailOk returns a tuple with the Toemail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToemail

`func (o *UpsUpdate0) SetToemail(v []string)`

SetToemail sets Toemail field to given value.

### HasToemail

`func (o *UpsUpdate0) HasToemail() bool`

HasToemail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


