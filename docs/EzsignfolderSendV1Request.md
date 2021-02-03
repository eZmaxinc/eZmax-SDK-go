# EzsignfolderSendV1Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TExtraMessage** | **string** | A custom text message that will be added to the email sent to signatories inviting them to sign.  You can send an empty string and only the generic message will be sent. | 

## Methods

### NewEzsignfolderSendV1Request

`func NewEzsignfolderSendV1Request(tExtraMessage string, ) *EzsignfolderSendV1Request`

NewEzsignfolderSendV1Request instantiates a new EzsignfolderSendV1Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsignfolderSendV1RequestWithDefaults

`func NewEzsignfolderSendV1RequestWithDefaults() *EzsignfolderSendV1Request`

NewEzsignfolderSendV1RequestWithDefaults instantiates a new EzsignfolderSendV1Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTExtraMessage

`func (o *EzsignfolderSendV1Request) GetTExtraMessage() string`

GetTExtraMessage returns the TExtraMessage field if non-nil, zero value otherwise.

### GetTExtraMessageOk

`func (o *EzsignfolderSendV1Request) GetTExtraMessageOk() (*string, bool)`

GetTExtraMessageOk returns a tuple with the TExtraMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTExtraMessage

`func (o *EzsignfolderSendV1Request) SetTExtraMessage(v string)`

SetTExtraMessage sets TExtraMessage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


