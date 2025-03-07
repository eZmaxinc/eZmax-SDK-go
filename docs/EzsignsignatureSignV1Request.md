# EzsignsignatureSignV1Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FkiEzsignsigningreasonID** | Pointer to **int32** | The unique ID of the Ezsignsigningreason | [optional] 
**FkiFontID** | Pointer to **int32** | The unique ID of the Font | [optional] 
**SValue** | Pointer to **string** | The value required for the Ezsignsignature.  This can only be set if eEzsignsignatureType is **City**, **FieldText** or **FieldTextarea** | [optional] 
**EAttachmentsConfirmationDecision** | Pointer to **string** | Whether the attachment are accepted or refused.  This can only be set if eEzsignsignatureType is **AttachmentsConfirmation** | [optional] 
**SAttachmentsRefusalReason** | Pointer to **string** | The reason of refused.  This can only be set if eEzsignsignatureType is **AttachmentsConfirmation** | [optional] 
**SSvg** | Pointer to **string** | The SVG of the signature.  This can only be set if eEzsignsignatureType is **Signature**_/_**Initials** and **bIsAutomatic** is false | [optional] 
**AObjFile** | Pointer to [**[]CommonFile**](CommonFile.md) |  | [optional] 
**ObjCreditcard** | Pointer to [**CustomCreditcardRequest**](CustomCreditcardRequest.md) |  | [optional] 
**BIsAutomatic** | **bool** | Indicates if the Ezsignsignature was part of an automatic process or not.  This can only be true if eEzsignsignatureType is **Acknowledgement**, **City**, **Signature**, **Initials** or **Stamp**.  | 

## Methods

### NewEzsignsignatureSignV1Request

`func NewEzsignsignatureSignV1Request(bIsAutomatic bool, ) *EzsignsignatureSignV1Request`

NewEzsignsignatureSignV1Request instantiates a new EzsignsignatureSignV1Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsignsignatureSignV1RequestWithDefaults

`func NewEzsignsignatureSignV1RequestWithDefaults() *EzsignsignatureSignV1Request`

NewEzsignsignatureSignV1RequestWithDefaults instantiates a new EzsignsignatureSignV1Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFkiEzsignsigningreasonID

`func (o *EzsignsignatureSignV1Request) GetFkiEzsignsigningreasonID() int32`

GetFkiEzsignsigningreasonID returns the FkiEzsignsigningreasonID field if non-nil, zero value otherwise.

### GetFkiEzsignsigningreasonIDOk

`func (o *EzsignsignatureSignV1Request) GetFkiEzsignsigningreasonIDOk() (*int32, bool)`

GetFkiEzsignsigningreasonIDOk returns a tuple with the FkiEzsignsigningreasonID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignsigningreasonID

`func (o *EzsignsignatureSignV1Request) SetFkiEzsignsigningreasonID(v int32)`

SetFkiEzsignsigningreasonID sets FkiEzsignsigningreasonID field to given value.

### HasFkiEzsignsigningreasonID

`func (o *EzsignsignatureSignV1Request) HasFkiEzsignsigningreasonID() bool`

HasFkiEzsignsigningreasonID returns a boolean if a field has been set.

### GetFkiFontID

`func (o *EzsignsignatureSignV1Request) GetFkiFontID() int32`

GetFkiFontID returns the FkiFontID field if non-nil, zero value otherwise.

### GetFkiFontIDOk

`func (o *EzsignsignatureSignV1Request) GetFkiFontIDOk() (*int32, bool)`

GetFkiFontIDOk returns a tuple with the FkiFontID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiFontID

`func (o *EzsignsignatureSignV1Request) SetFkiFontID(v int32)`

SetFkiFontID sets FkiFontID field to given value.

### HasFkiFontID

`func (o *EzsignsignatureSignV1Request) HasFkiFontID() bool`

HasFkiFontID returns a boolean if a field has been set.

### GetSValue

`func (o *EzsignsignatureSignV1Request) GetSValue() string`

GetSValue returns the SValue field if non-nil, zero value otherwise.

### GetSValueOk

`func (o *EzsignsignatureSignV1Request) GetSValueOk() (*string, bool)`

GetSValueOk returns a tuple with the SValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSValue

`func (o *EzsignsignatureSignV1Request) SetSValue(v string)`

SetSValue sets SValue field to given value.

### HasSValue

`func (o *EzsignsignatureSignV1Request) HasSValue() bool`

HasSValue returns a boolean if a field has been set.

### GetEAttachmentsConfirmationDecision

`func (o *EzsignsignatureSignV1Request) GetEAttachmentsConfirmationDecision() string`

GetEAttachmentsConfirmationDecision returns the EAttachmentsConfirmationDecision field if non-nil, zero value otherwise.

### GetEAttachmentsConfirmationDecisionOk

`func (o *EzsignsignatureSignV1Request) GetEAttachmentsConfirmationDecisionOk() (*string, bool)`

GetEAttachmentsConfirmationDecisionOk returns a tuple with the EAttachmentsConfirmationDecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEAttachmentsConfirmationDecision

`func (o *EzsignsignatureSignV1Request) SetEAttachmentsConfirmationDecision(v string)`

SetEAttachmentsConfirmationDecision sets EAttachmentsConfirmationDecision field to given value.

### HasEAttachmentsConfirmationDecision

`func (o *EzsignsignatureSignV1Request) HasEAttachmentsConfirmationDecision() bool`

HasEAttachmentsConfirmationDecision returns a boolean if a field has been set.

### GetSAttachmentsRefusalReason

`func (o *EzsignsignatureSignV1Request) GetSAttachmentsRefusalReason() string`

GetSAttachmentsRefusalReason returns the SAttachmentsRefusalReason field if non-nil, zero value otherwise.

### GetSAttachmentsRefusalReasonOk

`func (o *EzsignsignatureSignV1Request) GetSAttachmentsRefusalReasonOk() (*string, bool)`

GetSAttachmentsRefusalReasonOk returns a tuple with the SAttachmentsRefusalReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAttachmentsRefusalReason

`func (o *EzsignsignatureSignV1Request) SetSAttachmentsRefusalReason(v string)`

SetSAttachmentsRefusalReason sets SAttachmentsRefusalReason field to given value.

### HasSAttachmentsRefusalReason

`func (o *EzsignsignatureSignV1Request) HasSAttachmentsRefusalReason() bool`

HasSAttachmentsRefusalReason returns a boolean if a field has been set.

### GetSSvg

`func (o *EzsignsignatureSignV1Request) GetSSvg() string`

GetSSvg returns the SSvg field if non-nil, zero value otherwise.

### GetSSvgOk

`func (o *EzsignsignatureSignV1Request) GetSSvgOk() (*string, bool)`

GetSSvgOk returns a tuple with the SSvg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSSvg

`func (o *EzsignsignatureSignV1Request) SetSSvg(v string)`

SetSSvg sets SSvg field to given value.

### HasSSvg

`func (o *EzsignsignatureSignV1Request) HasSSvg() bool`

HasSSvg returns a boolean if a field has been set.

### GetAObjFile

`func (o *EzsignsignatureSignV1Request) GetAObjFile() []CommonFile`

GetAObjFile returns the AObjFile field if non-nil, zero value otherwise.

### GetAObjFileOk

`func (o *EzsignsignatureSignV1Request) GetAObjFileOk() (*[]CommonFile, bool)`

GetAObjFileOk returns a tuple with the AObjFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjFile

`func (o *EzsignsignatureSignV1Request) SetAObjFile(v []CommonFile)`

SetAObjFile sets AObjFile field to given value.

### HasAObjFile

`func (o *EzsignsignatureSignV1Request) HasAObjFile() bool`

HasAObjFile returns a boolean if a field has been set.

### GetObjCreditcard

`func (o *EzsignsignatureSignV1Request) GetObjCreditcard() CustomCreditcardRequest`

GetObjCreditcard returns the ObjCreditcard field if non-nil, zero value otherwise.

### GetObjCreditcardOk

`func (o *EzsignsignatureSignV1Request) GetObjCreditcardOk() (*CustomCreditcardRequest, bool)`

GetObjCreditcardOk returns a tuple with the ObjCreditcard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjCreditcard

`func (o *EzsignsignatureSignV1Request) SetObjCreditcard(v CustomCreditcardRequest)`

SetObjCreditcard sets ObjCreditcard field to given value.

### HasObjCreditcard

`func (o *EzsignsignatureSignV1Request) HasObjCreditcard() bool`

HasObjCreditcard returns a boolean if a field has been set.

### GetBIsAutomatic

`func (o *EzsignsignatureSignV1Request) GetBIsAutomatic() bool`

GetBIsAutomatic returns the BIsAutomatic field if non-nil, zero value otherwise.

### GetBIsAutomaticOk

`func (o *EzsignsignatureSignV1Request) GetBIsAutomaticOk() (*bool, bool)`

GetBIsAutomaticOk returns a tuple with the BIsAutomatic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBIsAutomatic

`func (o *EzsignsignatureSignV1Request) SetBIsAutomatic(v bool)`

SetBIsAutomatic sets BIsAutomatic field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


