# EzsignfolderSendV2Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TEzsignfolderMessage** | **string** | A custom text message that will be added to the email sent. | 
**AFkiEzsignfoldersignerassociationID** | **[]int32** |  | 
**AObjEzsignfoldersignerassociationmessage** | [**[]CustomEzsignfoldersignerassociationmessageRequest**](CustomEzsignfoldersignerassociationmessageRequest.md) |  | 

## Methods

### NewEzsignfolderSendV2Request

`func NewEzsignfolderSendV2Request(tEzsignfolderMessage string, aFkiEzsignfoldersignerassociationID []int32, aObjEzsignfoldersignerassociationmessage []CustomEzsignfoldersignerassociationmessageRequest, ) *EzsignfolderSendV2Request`

NewEzsignfolderSendV2Request instantiates a new EzsignfolderSendV2Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsignfolderSendV2RequestWithDefaults

`func NewEzsignfolderSendV2RequestWithDefaults() *EzsignfolderSendV2Request`

NewEzsignfolderSendV2RequestWithDefaults instantiates a new EzsignfolderSendV2Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTEzsignfolderMessage

`func (o *EzsignfolderSendV2Request) GetTEzsignfolderMessage() string`

GetTEzsignfolderMessage returns the TEzsignfolderMessage field if non-nil, zero value otherwise.

### GetTEzsignfolderMessageOk

`func (o *EzsignfolderSendV2Request) GetTEzsignfolderMessageOk() (*string, bool)`

GetTEzsignfolderMessageOk returns a tuple with the TEzsignfolderMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTEzsignfolderMessage

`func (o *EzsignfolderSendV2Request) SetTEzsignfolderMessage(v string)`

SetTEzsignfolderMessage sets TEzsignfolderMessage field to given value.


### GetAFkiEzsignfoldersignerassociationID

`func (o *EzsignfolderSendV2Request) GetAFkiEzsignfoldersignerassociationID() []int32`

GetAFkiEzsignfoldersignerassociationID returns the AFkiEzsignfoldersignerassociationID field if non-nil, zero value otherwise.

### GetAFkiEzsignfoldersignerassociationIDOk

`func (o *EzsignfolderSendV2Request) GetAFkiEzsignfoldersignerassociationIDOk() (*[]int32, bool)`

GetAFkiEzsignfoldersignerassociationIDOk returns a tuple with the AFkiEzsignfoldersignerassociationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAFkiEzsignfoldersignerassociationID

`func (o *EzsignfolderSendV2Request) SetAFkiEzsignfoldersignerassociationID(v []int32)`

SetAFkiEzsignfoldersignerassociationID sets AFkiEzsignfoldersignerassociationID field to given value.


### GetAObjEzsignfoldersignerassociationmessage

`func (o *EzsignfolderSendV2Request) GetAObjEzsignfoldersignerassociationmessage() []CustomEzsignfoldersignerassociationmessageRequest`

GetAObjEzsignfoldersignerassociationmessage returns the AObjEzsignfoldersignerassociationmessage field if non-nil, zero value otherwise.

### GetAObjEzsignfoldersignerassociationmessageOk

`func (o *EzsignfolderSendV2Request) GetAObjEzsignfoldersignerassociationmessageOk() (*[]CustomEzsignfoldersignerassociationmessageRequest, bool)`

GetAObjEzsignfoldersignerassociationmessageOk returns a tuple with the AObjEzsignfoldersignerassociationmessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjEzsignfoldersignerassociationmessage

`func (o *EzsignfolderSendV2Request) SetAObjEzsignfoldersignerassociationmessage(v []CustomEzsignfoldersignerassociationmessageRequest)`

SetAObjEzsignfoldersignerassociationmessage sets AObjEzsignfoldersignerassociationmessage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


