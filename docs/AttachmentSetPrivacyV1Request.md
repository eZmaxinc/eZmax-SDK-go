# AttachmentSetPrivacyV1Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EAttachmentPrivacy** | [**FieldEAttachmentPrivacy**](FieldEAttachmentPrivacy.md) |  | 
**FkiUserIDSpecific** | Pointer to **int32** | The unique ID of the User | [optional] 

## Methods

### NewAttachmentSetPrivacyV1Request

`func NewAttachmentSetPrivacyV1Request(eAttachmentPrivacy FieldEAttachmentPrivacy, ) *AttachmentSetPrivacyV1Request`

NewAttachmentSetPrivacyV1Request instantiates a new AttachmentSetPrivacyV1Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttachmentSetPrivacyV1RequestWithDefaults

`func NewAttachmentSetPrivacyV1RequestWithDefaults() *AttachmentSetPrivacyV1Request`

NewAttachmentSetPrivacyV1RequestWithDefaults instantiates a new AttachmentSetPrivacyV1Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEAttachmentPrivacy

`func (o *AttachmentSetPrivacyV1Request) GetEAttachmentPrivacy() FieldEAttachmentPrivacy`

GetEAttachmentPrivacy returns the EAttachmentPrivacy field if non-nil, zero value otherwise.

### GetEAttachmentPrivacyOk

`func (o *AttachmentSetPrivacyV1Request) GetEAttachmentPrivacyOk() (*FieldEAttachmentPrivacy, bool)`

GetEAttachmentPrivacyOk returns a tuple with the EAttachmentPrivacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEAttachmentPrivacy

`func (o *AttachmentSetPrivacyV1Request) SetEAttachmentPrivacy(v FieldEAttachmentPrivacy)`

SetEAttachmentPrivacy sets EAttachmentPrivacy field to given value.


### GetFkiUserIDSpecific

`func (o *AttachmentSetPrivacyV1Request) GetFkiUserIDSpecific() int32`

GetFkiUserIDSpecific returns the FkiUserIDSpecific field if non-nil, zero value otherwise.

### GetFkiUserIDSpecificOk

`func (o *AttachmentSetPrivacyV1Request) GetFkiUserIDSpecificOk() (*int32, bool)`

GetFkiUserIDSpecificOk returns a tuple with the FkiUserIDSpecific field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUserIDSpecific

`func (o *AttachmentSetPrivacyV1Request) SetFkiUserIDSpecific(v int32)`

SetFkiUserIDSpecific sets FkiUserIDSpecific field to given value.

### HasFkiUserIDSpecific

`func (o *AttachmentSetPrivacyV1Request) HasFkiUserIDSpecific() bool`

HasFkiUserIDSpecific returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


