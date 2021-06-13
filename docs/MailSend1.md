# MailSend1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fromemail** | Pointer to **string** |  | [optional] 
**Fromname** | Pointer to **string** |  | [optional] 
**Outgoingserver** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**Security** | Pointer to **string** |  | [optional] 
**Smtp** | Pointer to **bool** |  | [optional] 
**User** | Pointer to **string** |  | [optional] 
**Pass** | Pointer to **string** |  | [optional] 
**Oauth** | Pointer to [**MailUpdate0Oauth**](MailUpdate0Oauth.md) |  | [optional] 

## Methods

### NewMailSend1

`func NewMailSend1() *MailSend1`

NewMailSend1 instantiates a new MailSend1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMailSend1WithDefaults

`func NewMailSend1WithDefaults() *MailSend1`

NewMailSend1WithDefaults instantiates a new MailSend1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFromemail

`func (o *MailSend1) GetFromemail() string`

GetFromemail returns the Fromemail field if non-nil, zero value otherwise.

### GetFromemailOk

`func (o *MailSend1) GetFromemailOk() (*string, bool)`

GetFromemailOk returns a tuple with the Fromemail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromemail

`func (o *MailSend1) SetFromemail(v string)`

SetFromemail sets Fromemail field to given value.

### HasFromemail

`func (o *MailSend1) HasFromemail() bool`

HasFromemail returns a boolean if a field has been set.

### GetFromname

`func (o *MailSend1) GetFromname() string`

GetFromname returns the Fromname field if non-nil, zero value otherwise.

### GetFromnameOk

`func (o *MailSend1) GetFromnameOk() (*string, bool)`

GetFromnameOk returns a tuple with the Fromname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromname

`func (o *MailSend1) SetFromname(v string)`

SetFromname sets Fromname field to given value.

### HasFromname

`func (o *MailSend1) HasFromname() bool`

HasFromname returns a boolean if a field has been set.

### GetOutgoingserver

`func (o *MailSend1) GetOutgoingserver() string`

GetOutgoingserver returns the Outgoingserver field if non-nil, zero value otherwise.

### GetOutgoingserverOk

`func (o *MailSend1) GetOutgoingserverOk() (*string, bool)`

GetOutgoingserverOk returns a tuple with the Outgoingserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutgoingserver

`func (o *MailSend1) SetOutgoingserver(v string)`

SetOutgoingserver sets Outgoingserver field to given value.

### HasOutgoingserver

`func (o *MailSend1) HasOutgoingserver() bool`

HasOutgoingserver returns a boolean if a field has been set.

### GetPort

`func (o *MailSend1) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *MailSend1) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *MailSend1) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *MailSend1) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetSecurity

`func (o *MailSend1) GetSecurity() string`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *MailSend1) GetSecurityOk() (*string, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *MailSend1) SetSecurity(v string)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *MailSend1) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### GetSmtp

`func (o *MailSend1) GetSmtp() bool`

GetSmtp returns the Smtp field if non-nil, zero value otherwise.

### GetSmtpOk

`func (o *MailSend1) GetSmtpOk() (*bool, bool)`

GetSmtpOk returns a tuple with the Smtp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtp

`func (o *MailSend1) SetSmtp(v bool)`

SetSmtp sets Smtp field to given value.

### HasSmtp

`func (o *MailSend1) HasSmtp() bool`

HasSmtp returns a boolean if a field has been set.

### GetUser

`func (o *MailSend1) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *MailSend1) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *MailSend1) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *MailSend1) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetPass

`func (o *MailSend1) GetPass() string`

GetPass returns the Pass field if non-nil, zero value otherwise.

### GetPassOk

`func (o *MailSend1) GetPassOk() (*string, bool)`

GetPassOk returns a tuple with the Pass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPass

`func (o *MailSend1) SetPass(v string)`

SetPass sets Pass field to given value.

### HasPass

`func (o *MailSend1) HasPass() bool`

HasPass returns a boolean if a field has been set.

### GetOauth

`func (o *MailSend1) GetOauth() MailUpdate0Oauth`

GetOauth returns the Oauth field if non-nil, zero value otherwise.

### GetOauthOk

`func (o *MailSend1) GetOauthOk() (*MailUpdate0Oauth, bool)`

GetOauthOk returns a tuple with the Oauth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth

`func (o *MailSend1) SetOauth(v MailUpdate0Oauth)`

SetOauth sets Oauth field to given value.

### HasOauth

`func (o *MailSend1) HasOauth() bool`

HasOauth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


