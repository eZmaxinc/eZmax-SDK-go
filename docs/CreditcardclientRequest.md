# CreditcardclientRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiCreditcardclientID** | Pointer to **int32** | The unique ID of the Creditcardclient | [optional] 
**FksCreditcardtokenID** | Pointer to **string** | The creditcard token identifier | [optional] 
**BCreditcardclientrelationIsdefault** | **bool** | Whether if it&#39;s the creditcardclient is the default one | 
**SCreditcardclientDescription** | **string** | The description of the Creditcardclient | 
**BCreditcardclientAllowedcompanypayment** | **bool** | Whether if it&#39;s an allowedagencypayment | 
**BCreditcardclientAllowedezsign** | **bool** | Whether if it&#39;s an allowedroyallepageprotection | 
**BCreditcardclientAllowedtranquillit** | **bool** | Whether if it&#39;s an allowedtranquillit | 
**ObjCreditcarddetail** | [**CreditcarddetailRequest**](CreditcarddetailRequest.md) |  | 
**SCreditcardclientCVV** | **string** | The creditcard card CVV | 

## Methods

### NewCreditcardclientRequest

`func NewCreditcardclientRequest(bCreditcardclientrelationIsdefault bool, sCreditcardclientDescription string, bCreditcardclientAllowedcompanypayment bool, bCreditcardclientAllowedezsign bool, bCreditcardclientAllowedtranquillit bool, objCreditcarddetail CreditcarddetailRequest, sCreditcardclientCVV string, ) *CreditcardclientRequest`

NewCreditcardclientRequest instantiates a new CreditcardclientRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditcardclientRequestWithDefaults

`func NewCreditcardclientRequestWithDefaults() *CreditcardclientRequest`

NewCreditcardclientRequestWithDefaults instantiates a new CreditcardclientRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiCreditcardclientID

`func (o *CreditcardclientRequest) GetPkiCreditcardclientID() int32`

GetPkiCreditcardclientID returns the PkiCreditcardclientID field if non-nil, zero value otherwise.

### GetPkiCreditcardclientIDOk

`func (o *CreditcardclientRequest) GetPkiCreditcardclientIDOk() (*int32, bool)`

GetPkiCreditcardclientIDOk returns a tuple with the PkiCreditcardclientID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiCreditcardclientID

`func (o *CreditcardclientRequest) SetPkiCreditcardclientID(v int32)`

SetPkiCreditcardclientID sets PkiCreditcardclientID field to given value.

### HasPkiCreditcardclientID

`func (o *CreditcardclientRequest) HasPkiCreditcardclientID() bool`

HasPkiCreditcardclientID returns a boolean if a field has been set.

### GetFksCreditcardtokenID

`func (o *CreditcardclientRequest) GetFksCreditcardtokenID() string`

GetFksCreditcardtokenID returns the FksCreditcardtokenID field if non-nil, zero value otherwise.

### GetFksCreditcardtokenIDOk

`func (o *CreditcardclientRequest) GetFksCreditcardtokenIDOk() (*string, bool)`

GetFksCreditcardtokenIDOk returns a tuple with the FksCreditcardtokenID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFksCreditcardtokenID

`func (o *CreditcardclientRequest) SetFksCreditcardtokenID(v string)`

SetFksCreditcardtokenID sets FksCreditcardtokenID field to given value.

### HasFksCreditcardtokenID

`func (o *CreditcardclientRequest) HasFksCreditcardtokenID() bool`

HasFksCreditcardtokenID returns a boolean if a field has been set.

### GetBCreditcardclientrelationIsdefault

`func (o *CreditcardclientRequest) GetBCreditcardclientrelationIsdefault() bool`

GetBCreditcardclientrelationIsdefault returns the BCreditcardclientrelationIsdefault field if non-nil, zero value otherwise.

### GetBCreditcardclientrelationIsdefaultOk

`func (o *CreditcardclientRequest) GetBCreditcardclientrelationIsdefaultOk() (*bool, bool)`

GetBCreditcardclientrelationIsdefaultOk returns a tuple with the BCreditcardclientrelationIsdefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBCreditcardclientrelationIsdefault

`func (o *CreditcardclientRequest) SetBCreditcardclientrelationIsdefault(v bool)`

SetBCreditcardclientrelationIsdefault sets BCreditcardclientrelationIsdefault field to given value.


### GetSCreditcardclientDescription

`func (o *CreditcardclientRequest) GetSCreditcardclientDescription() string`

GetSCreditcardclientDescription returns the SCreditcardclientDescription field if non-nil, zero value otherwise.

### GetSCreditcardclientDescriptionOk

`func (o *CreditcardclientRequest) GetSCreditcardclientDescriptionOk() (*string, bool)`

GetSCreditcardclientDescriptionOk returns a tuple with the SCreditcardclientDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSCreditcardclientDescription

`func (o *CreditcardclientRequest) SetSCreditcardclientDescription(v string)`

SetSCreditcardclientDescription sets SCreditcardclientDescription field to given value.


### GetBCreditcardclientAllowedcompanypayment

`func (o *CreditcardclientRequest) GetBCreditcardclientAllowedcompanypayment() bool`

GetBCreditcardclientAllowedcompanypayment returns the BCreditcardclientAllowedcompanypayment field if non-nil, zero value otherwise.

### GetBCreditcardclientAllowedcompanypaymentOk

`func (o *CreditcardclientRequest) GetBCreditcardclientAllowedcompanypaymentOk() (*bool, bool)`

GetBCreditcardclientAllowedcompanypaymentOk returns a tuple with the BCreditcardclientAllowedcompanypayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBCreditcardclientAllowedcompanypayment

`func (o *CreditcardclientRequest) SetBCreditcardclientAllowedcompanypayment(v bool)`

SetBCreditcardclientAllowedcompanypayment sets BCreditcardclientAllowedcompanypayment field to given value.


### GetBCreditcardclientAllowedezsign

`func (o *CreditcardclientRequest) GetBCreditcardclientAllowedezsign() bool`

GetBCreditcardclientAllowedezsign returns the BCreditcardclientAllowedezsign field if non-nil, zero value otherwise.

### GetBCreditcardclientAllowedezsignOk

`func (o *CreditcardclientRequest) GetBCreditcardclientAllowedezsignOk() (*bool, bool)`

GetBCreditcardclientAllowedezsignOk returns a tuple with the BCreditcardclientAllowedezsign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBCreditcardclientAllowedezsign

`func (o *CreditcardclientRequest) SetBCreditcardclientAllowedezsign(v bool)`

SetBCreditcardclientAllowedezsign sets BCreditcardclientAllowedezsign field to given value.


### GetBCreditcardclientAllowedtranquillit

`func (o *CreditcardclientRequest) GetBCreditcardclientAllowedtranquillit() bool`

GetBCreditcardclientAllowedtranquillit returns the BCreditcardclientAllowedtranquillit field if non-nil, zero value otherwise.

### GetBCreditcardclientAllowedtranquillitOk

`func (o *CreditcardclientRequest) GetBCreditcardclientAllowedtranquillitOk() (*bool, bool)`

GetBCreditcardclientAllowedtranquillitOk returns a tuple with the BCreditcardclientAllowedtranquillit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBCreditcardclientAllowedtranquillit

`func (o *CreditcardclientRequest) SetBCreditcardclientAllowedtranquillit(v bool)`

SetBCreditcardclientAllowedtranquillit sets BCreditcardclientAllowedtranquillit field to given value.


### GetObjCreditcarddetail

`func (o *CreditcardclientRequest) GetObjCreditcarddetail() CreditcarddetailRequest`

GetObjCreditcarddetail returns the ObjCreditcarddetail field if non-nil, zero value otherwise.

### GetObjCreditcarddetailOk

`func (o *CreditcardclientRequest) GetObjCreditcarddetailOk() (*CreditcarddetailRequest, bool)`

GetObjCreditcarddetailOk returns a tuple with the ObjCreditcarddetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjCreditcarddetail

`func (o *CreditcardclientRequest) SetObjCreditcarddetail(v CreditcarddetailRequest)`

SetObjCreditcarddetail sets ObjCreditcarddetail field to given value.


### GetSCreditcardclientCVV

`func (o *CreditcardclientRequest) GetSCreditcardclientCVV() string`

GetSCreditcardclientCVV returns the SCreditcardclientCVV field if non-nil, zero value otherwise.

### GetSCreditcardclientCVVOk

`func (o *CreditcardclientRequest) GetSCreditcardclientCVVOk() (*string, bool)`

GetSCreditcardclientCVVOk returns a tuple with the SCreditcardclientCVV field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSCreditcardclientCVV

`func (o *CreditcardclientRequest) SetSCreditcardclientCVV(v string)`

SetSCreditcardclientCVV sets SCreditcardclientCVV field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


