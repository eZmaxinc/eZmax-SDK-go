# AttachmentValidateV1Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FkiAttachmentrejectreasonID** | Pointer to **int32** | The unique ID of the Attachmentrejectreason | [optional] 
**EAttachmentVerified** | [**FieldEAttachmentVerified**](FieldEAttachmentVerified.md) |  | 
**TAttachmentRejectioncomment** | Pointer to **string** | The rejectioncomment of the Attachment | [optional] 
**ASNotificationEmailAddress** | Pointer to **[]string** | A list of email addresses to send a rejection notification to. | [optional] 
**TNotificationMessage** | Pointer to **string** | The email body to send to the notification recipients. | [optional] 
**BNotificationIncludeAttachment** | Pointer to **bool** | Should a copy of the attachment be included in the email. | [optional] 

## Methods

### NewAttachmentValidateV1Request

`func NewAttachmentValidateV1Request(eAttachmentVerified FieldEAttachmentVerified, ) *AttachmentValidateV1Request`

NewAttachmentValidateV1Request instantiates a new AttachmentValidateV1Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttachmentValidateV1RequestWithDefaults

`func NewAttachmentValidateV1RequestWithDefaults() *AttachmentValidateV1Request`

NewAttachmentValidateV1RequestWithDefaults instantiates a new AttachmentValidateV1Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFkiAttachmentrejectreasonID

`func (o *AttachmentValidateV1Request) GetFkiAttachmentrejectreasonID() int32`

GetFkiAttachmentrejectreasonID returns the FkiAttachmentrejectreasonID field if non-nil, zero value otherwise.

### GetFkiAttachmentrejectreasonIDOk

`func (o *AttachmentValidateV1Request) GetFkiAttachmentrejectreasonIDOk() (*int32, bool)`

GetFkiAttachmentrejectreasonIDOk returns a tuple with the FkiAttachmentrejectreasonID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiAttachmentrejectreasonID

`func (o *AttachmentValidateV1Request) SetFkiAttachmentrejectreasonID(v int32)`

SetFkiAttachmentrejectreasonID sets FkiAttachmentrejectreasonID field to given value.

### HasFkiAttachmentrejectreasonID

`func (o *AttachmentValidateV1Request) HasFkiAttachmentrejectreasonID() bool`

HasFkiAttachmentrejectreasonID returns a boolean if a field has been set.

### GetEAttachmentVerified

`func (o *AttachmentValidateV1Request) GetEAttachmentVerified() FieldEAttachmentVerified`

GetEAttachmentVerified returns the EAttachmentVerified field if non-nil, zero value otherwise.

### GetEAttachmentVerifiedOk

`func (o *AttachmentValidateV1Request) GetEAttachmentVerifiedOk() (*FieldEAttachmentVerified, bool)`

GetEAttachmentVerifiedOk returns a tuple with the EAttachmentVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEAttachmentVerified

`func (o *AttachmentValidateV1Request) SetEAttachmentVerified(v FieldEAttachmentVerified)`

SetEAttachmentVerified sets EAttachmentVerified field to given value.


### GetTAttachmentRejectioncomment

`func (o *AttachmentValidateV1Request) GetTAttachmentRejectioncomment() string`

GetTAttachmentRejectioncomment returns the TAttachmentRejectioncomment field if non-nil, zero value otherwise.

### GetTAttachmentRejectioncommentOk

`func (o *AttachmentValidateV1Request) GetTAttachmentRejectioncommentOk() (*string, bool)`

GetTAttachmentRejectioncommentOk returns a tuple with the TAttachmentRejectioncomment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTAttachmentRejectioncomment

`func (o *AttachmentValidateV1Request) SetTAttachmentRejectioncomment(v string)`

SetTAttachmentRejectioncomment sets TAttachmentRejectioncomment field to given value.

### HasTAttachmentRejectioncomment

`func (o *AttachmentValidateV1Request) HasTAttachmentRejectioncomment() bool`

HasTAttachmentRejectioncomment returns a boolean if a field has been set.

### GetASNotificationEmailAddress

`func (o *AttachmentValidateV1Request) GetASNotificationEmailAddress() []string`

GetASNotificationEmailAddress returns the ASNotificationEmailAddress field if non-nil, zero value otherwise.

### GetASNotificationEmailAddressOk

`func (o *AttachmentValidateV1Request) GetASNotificationEmailAddressOk() (*[]string, bool)`

GetASNotificationEmailAddressOk returns a tuple with the ASNotificationEmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetASNotificationEmailAddress

`func (o *AttachmentValidateV1Request) SetASNotificationEmailAddress(v []string)`

SetASNotificationEmailAddress sets ASNotificationEmailAddress field to given value.

### HasASNotificationEmailAddress

`func (o *AttachmentValidateV1Request) HasASNotificationEmailAddress() bool`

HasASNotificationEmailAddress returns a boolean if a field has been set.

### GetTNotificationMessage

`func (o *AttachmentValidateV1Request) GetTNotificationMessage() string`

GetTNotificationMessage returns the TNotificationMessage field if non-nil, zero value otherwise.

### GetTNotificationMessageOk

`func (o *AttachmentValidateV1Request) GetTNotificationMessageOk() (*string, bool)`

GetTNotificationMessageOk returns a tuple with the TNotificationMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTNotificationMessage

`func (o *AttachmentValidateV1Request) SetTNotificationMessage(v string)`

SetTNotificationMessage sets TNotificationMessage field to given value.

### HasTNotificationMessage

`func (o *AttachmentValidateV1Request) HasTNotificationMessage() bool`

HasTNotificationMessage returns a boolean if a field has been set.

### GetBNotificationIncludeAttachment

`func (o *AttachmentValidateV1Request) GetBNotificationIncludeAttachment() bool`

GetBNotificationIncludeAttachment returns the BNotificationIncludeAttachment field if non-nil, zero value otherwise.

### GetBNotificationIncludeAttachmentOk

`func (o *AttachmentValidateV1Request) GetBNotificationIncludeAttachmentOk() (*bool, bool)`

GetBNotificationIncludeAttachmentOk returns a tuple with the BNotificationIncludeAttachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBNotificationIncludeAttachment

`func (o *AttachmentValidateV1Request) SetBNotificationIncludeAttachment(v bool)`

SetBNotificationIncludeAttachment sets BNotificationIncludeAttachment field to given value.

### HasBNotificationIncludeAttachment

`func (o *AttachmentValidateV1Request) HasBNotificationIncludeAttachment() bool`

HasBNotificationIncludeAttachment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


