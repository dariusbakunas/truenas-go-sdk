# MailSend

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MailMessage** | Pointer to [**MailSend0**](MailSend0.md) |  | [optional] [default to {}]
**MailUpdate** | Pointer to [**MailSend1**](MailSend1.md) |  | [optional] [default to {}]

## Methods

### NewMailSend

`func NewMailSend() *MailSend`

NewMailSend instantiates a new MailSend object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMailSendWithDefaults

`func NewMailSendWithDefaults() *MailSend`

NewMailSendWithDefaults instantiates a new MailSend object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMailMessage

`func (o *MailSend) GetMailMessage() MailSend0`

GetMailMessage returns the MailMessage field if non-nil, zero value otherwise.

### GetMailMessageOk

`func (o *MailSend) GetMailMessageOk() (*MailSend0, bool)`

GetMailMessageOk returns a tuple with the MailMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailMessage

`func (o *MailSend) SetMailMessage(v MailSend0)`

SetMailMessage sets MailMessage field to given value.

### HasMailMessage

`func (o *MailSend) HasMailMessage() bool`

HasMailMessage returns a boolean if a field has been set.

### GetMailUpdate

`func (o *MailSend) GetMailUpdate() MailSend1`

GetMailUpdate returns the MailUpdate field if non-nil, zero value otherwise.

### GetMailUpdateOk

`func (o *MailSend) GetMailUpdateOk() (*MailSend1, bool)`

GetMailUpdateOk returns a tuple with the MailUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailUpdate

`func (o *MailSend) SetMailUpdate(v MailSend1)`

SetMailUpdate sets MailUpdate field to given value.

### HasMailUpdate

`func (o *MailSend) HasMailUpdate() bool`

HasMailUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


