# CreditcardclientRequestCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiCreditcardclientID** | Pointer to **int32** | The unique ID of the Creditcardclient | [optional] 
**FksCreditcardtokenID** | Pointer to **string** | The creditcard token identifier | [optional] 
**BCreditcardclientrelationIsdefault** | **bool** | Whether if it&#39;s an relationisdefault | 
**SCreditcardclientDescription** | **string** | The description of the Creditcardclient | 
**BCreditcardclientIsactive** | **bool** | Whether the creditcardclient is active or not | 
**BCreditcardclientAllowedagencypayment** | **bool** | Whether if it&#39;s an allowedagencypayment | 
**BCreditcardclientAllowedroyallepageprotection** | **bool** | Whether if it&#39;s an allowedroyallepageprotection | 
**BCreditcardclientAllowedtranquillit** | **bool** | Whether if it&#39;s an allowedtranquillit | 
**ObjCreditcarddetail** | [**CreditcarddetailRequest**](CreditcarddetailRequest.md) |  | 
**SCreditcardclientCVV** | **string** | The creditcard card CVV | 

## Methods

### NewCreditcardclientRequestCompound

`func NewCreditcardclientRequestCompound(bCreditcardclientrelationIsdefault bool, sCreditcardclientDescription string, bCreditcardclientIsactive bool, bCreditcardclientAllowedagencypayment bool, bCreditcardclientAllowedroyallepageprotection bool, bCreditcardclientAllowedtranquillit bool, objCreditcarddetail CreditcarddetailRequest, sCreditcardclientCVV string, ) *CreditcardclientRequestCompound`

NewCreditcardclientRequestCompound instantiates a new CreditcardclientRequestCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditcardclientRequestCompoundWithDefaults

`func NewCreditcardclientRequestCompoundWithDefaults() *CreditcardclientRequestCompound`

NewCreditcardclientRequestCompoundWithDefaults instantiates a new CreditcardclientRequestCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiCreditcardclientID

`func (o *CreditcardclientRequestCompound) GetPkiCreditcardclientID() int32`

GetPkiCreditcardclientID returns the PkiCreditcardclientID field if non-nil, zero value otherwise.

### GetPkiCreditcardclientIDOk

`func (o *CreditcardclientRequestCompound) GetPkiCreditcardclientIDOk() (*int32, bool)`

GetPkiCreditcardclientIDOk returns a tuple with the PkiCreditcardclientID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiCreditcardclientID

`func (o *CreditcardclientRequestCompound) SetPkiCreditcardclientID(v int32)`

SetPkiCreditcardclientID sets PkiCreditcardclientID field to given value.

### HasPkiCreditcardclientID

`func (o *CreditcardclientRequestCompound) HasPkiCreditcardclientID() bool`

HasPkiCreditcardclientID returns a boolean if a field has been set.

### GetFksCreditcardtokenID

`func (o *CreditcardclientRequestCompound) GetFksCreditcardtokenID() string`

GetFksCreditcardtokenID returns the FksCreditcardtokenID field if non-nil, zero value otherwise.

### GetFksCreditcardtokenIDOk

`func (o *CreditcardclientRequestCompound) GetFksCreditcardtokenIDOk() (*string, bool)`

GetFksCreditcardtokenIDOk returns a tuple with the FksCreditcardtokenID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFksCreditcardtokenID

`func (o *CreditcardclientRequestCompound) SetFksCreditcardtokenID(v string)`

SetFksCreditcardtokenID sets FksCreditcardtokenID field to given value.

### HasFksCreditcardtokenID

`func (o *CreditcardclientRequestCompound) HasFksCreditcardtokenID() bool`

HasFksCreditcardtokenID returns a boolean if a field has been set.

### GetBCreditcardclientrelationIsdefault

`func (o *CreditcardclientRequestCompound) GetBCreditcardclientrelationIsdefault() bool`

GetBCreditcardclientrelationIsdefault returns the BCreditcardclientrelationIsdefault field if non-nil, zero value otherwise.

### GetBCreditcardclientrelationIsdefaultOk

`func (o *CreditcardclientRequestCompound) GetBCreditcardclientrelationIsdefaultOk() (*bool, bool)`

GetBCreditcardclientrelationIsdefaultOk returns a tuple with the BCreditcardclientrelationIsdefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBCreditcardclientrelationIsdefault

`func (o *CreditcardclientRequestCompound) SetBCreditcardclientrelationIsdefault(v bool)`

SetBCreditcardclientrelationIsdefault sets BCreditcardclientrelationIsdefault field to given value.


### GetSCreditcardclientDescription

`func (o *CreditcardclientRequestCompound) GetSCreditcardclientDescription() string`

GetSCreditcardclientDescription returns the SCreditcardclientDescription field if non-nil, zero value otherwise.

### GetSCreditcardclientDescriptionOk

`func (o *CreditcardclientRequestCompound) GetSCreditcardclientDescriptionOk() (*string, bool)`

GetSCreditcardclientDescriptionOk returns a tuple with the SCreditcardclientDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSCreditcardclientDescription

`func (o *CreditcardclientRequestCompound) SetSCreditcardclientDescription(v string)`

SetSCreditcardclientDescription sets SCreditcardclientDescription field to given value.


### GetBCreditcardclientIsactive

`func (o *CreditcardclientRequestCompound) GetBCreditcardclientIsactive() bool`

GetBCreditcardclientIsactive returns the BCreditcardclientIsactive field if non-nil, zero value otherwise.

### GetBCreditcardclientIsactiveOk

`func (o *CreditcardclientRequestCompound) GetBCreditcardclientIsactiveOk() (*bool, bool)`

GetBCreditcardclientIsactiveOk returns a tuple with the BCreditcardclientIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBCreditcardclientIsactive

`func (o *CreditcardclientRequestCompound) SetBCreditcardclientIsactive(v bool)`

SetBCreditcardclientIsactive sets BCreditcardclientIsactive field to given value.


### GetBCreditcardclientAllowedagencypayment

`func (o *CreditcardclientRequestCompound) GetBCreditcardclientAllowedagencypayment() bool`

GetBCreditcardclientAllowedagencypayment returns the BCreditcardclientAllowedagencypayment field if non-nil, zero value otherwise.

### GetBCreditcardclientAllowedagencypaymentOk

`func (o *CreditcardclientRequestCompound) GetBCreditcardclientAllowedagencypaymentOk() (*bool, bool)`

GetBCreditcardclientAllowedagencypaymentOk returns a tuple with the BCreditcardclientAllowedagencypayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBCreditcardclientAllowedagencypayment

`func (o *CreditcardclientRequestCompound) SetBCreditcardclientAllowedagencypayment(v bool)`

SetBCreditcardclientAllowedagencypayment sets BCreditcardclientAllowedagencypayment field to given value.


### GetBCreditcardclientAllowedroyallepageprotection

`func (o *CreditcardclientRequestCompound) GetBCreditcardclientAllowedroyallepageprotection() bool`

GetBCreditcardclientAllowedroyallepageprotection returns the BCreditcardclientAllowedroyallepageprotection field if non-nil, zero value otherwise.

### GetBCreditcardclientAllowedroyallepageprotectionOk

`func (o *CreditcardclientRequestCompound) GetBCreditcardclientAllowedroyallepageprotectionOk() (*bool, bool)`

GetBCreditcardclientAllowedroyallepageprotectionOk returns a tuple with the BCreditcardclientAllowedroyallepageprotection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBCreditcardclientAllowedroyallepageprotection

`func (o *CreditcardclientRequestCompound) SetBCreditcardclientAllowedroyallepageprotection(v bool)`

SetBCreditcardclientAllowedroyallepageprotection sets BCreditcardclientAllowedroyallepageprotection field to given value.


### GetBCreditcardclientAllowedtranquillit

`func (o *CreditcardclientRequestCompound) GetBCreditcardclientAllowedtranquillit() bool`

GetBCreditcardclientAllowedtranquillit returns the BCreditcardclientAllowedtranquillit field if non-nil, zero value otherwise.

### GetBCreditcardclientAllowedtranquillitOk

`func (o *CreditcardclientRequestCompound) GetBCreditcardclientAllowedtranquillitOk() (*bool, bool)`

GetBCreditcardclientAllowedtranquillitOk returns a tuple with the BCreditcardclientAllowedtranquillit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBCreditcardclientAllowedtranquillit

`func (o *CreditcardclientRequestCompound) SetBCreditcardclientAllowedtranquillit(v bool)`

SetBCreditcardclientAllowedtranquillit sets BCreditcardclientAllowedtranquillit field to given value.


### GetObjCreditcarddetail

`func (o *CreditcardclientRequestCompound) GetObjCreditcarddetail() CreditcarddetailRequest`

GetObjCreditcarddetail returns the ObjCreditcarddetail field if non-nil, zero value otherwise.

### GetObjCreditcarddetailOk

`func (o *CreditcardclientRequestCompound) GetObjCreditcarddetailOk() (*CreditcarddetailRequest, bool)`

GetObjCreditcarddetailOk returns a tuple with the ObjCreditcarddetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjCreditcarddetail

`func (o *CreditcardclientRequestCompound) SetObjCreditcarddetail(v CreditcarddetailRequest)`

SetObjCreditcarddetail sets ObjCreditcarddetail field to given value.


### GetSCreditcardclientCVV

`func (o *CreditcardclientRequestCompound) GetSCreditcardclientCVV() string`

GetSCreditcardclientCVV returns the SCreditcardclientCVV field if non-nil, zero value otherwise.

### GetSCreditcardclientCVVOk

`func (o *CreditcardclientRequestCompound) GetSCreditcardclientCVVOk() (*string, bool)`

GetSCreditcardclientCVVOk returns a tuple with the SCreditcardclientCVV field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSCreditcardclientCVV

`func (o *CreditcardclientRequestCompound) SetSCreditcardclientCVV(v string)`

SetSCreditcardclientCVV sets SCreditcardclientCVV field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


