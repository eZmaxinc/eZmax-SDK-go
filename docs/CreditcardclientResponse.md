# CreditcardclientResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiCreditcardclientID** | **int32** | The unique ID of the Creditcardclient | 
**FkiCreditcarddetailID** | **int32** | The unique ID of the Creditcarddetail | 
**BCreditcardclientrelationIsdefault** | **bool** | Whether if it&#39;s the creditcardclient is the default one | 
**SCreditcardclientDescription** | **string** | The description of the Creditcardclient | 
**BCreditcardclientAllowedcompanypayment** | **bool** | Whether if it&#39;s an allowedagencypayment | 
**BCreditcardclientAllowedtranquillit** | **bool** | Whether if it&#39;s an allowedtranquillit | 
**ObjCreditcarddetail** | [**CreditcarddetailResponseCompound**](CreditcarddetailResponseCompound.md) |  | 

## Methods

### NewCreditcardclientResponse

`func NewCreditcardclientResponse(pkiCreditcardclientID int32, fkiCreditcarddetailID int32, bCreditcardclientrelationIsdefault bool, sCreditcardclientDescription string, bCreditcardclientAllowedcompanypayment bool, bCreditcardclientAllowedtranquillit bool, objCreditcarddetail CreditcarddetailResponseCompound, ) *CreditcardclientResponse`

NewCreditcardclientResponse instantiates a new CreditcardclientResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditcardclientResponseWithDefaults

`func NewCreditcardclientResponseWithDefaults() *CreditcardclientResponse`

NewCreditcardclientResponseWithDefaults instantiates a new CreditcardclientResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiCreditcardclientID

`func (o *CreditcardclientResponse) GetPkiCreditcardclientID() int32`

GetPkiCreditcardclientID returns the PkiCreditcardclientID field if non-nil, zero value otherwise.

### GetPkiCreditcardclientIDOk

`func (o *CreditcardclientResponse) GetPkiCreditcardclientIDOk() (*int32, bool)`

GetPkiCreditcardclientIDOk returns a tuple with the PkiCreditcardclientID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiCreditcardclientID

`func (o *CreditcardclientResponse) SetPkiCreditcardclientID(v int32)`

SetPkiCreditcardclientID sets PkiCreditcardclientID field to given value.


### GetFkiCreditcarddetailID

`func (o *CreditcardclientResponse) GetFkiCreditcarddetailID() int32`

GetFkiCreditcarddetailID returns the FkiCreditcarddetailID field if non-nil, zero value otherwise.

### GetFkiCreditcarddetailIDOk

`func (o *CreditcardclientResponse) GetFkiCreditcarddetailIDOk() (*int32, bool)`

GetFkiCreditcarddetailIDOk returns a tuple with the FkiCreditcarddetailID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiCreditcarddetailID

`func (o *CreditcardclientResponse) SetFkiCreditcarddetailID(v int32)`

SetFkiCreditcarddetailID sets FkiCreditcarddetailID field to given value.


### GetBCreditcardclientrelationIsdefault

`func (o *CreditcardclientResponse) GetBCreditcardclientrelationIsdefault() bool`

GetBCreditcardclientrelationIsdefault returns the BCreditcardclientrelationIsdefault field if non-nil, zero value otherwise.

### GetBCreditcardclientrelationIsdefaultOk

`func (o *CreditcardclientResponse) GetBCreditcardclientrelationIsdefaultOk() (*bool, bool)`

GetBCreditcardclientrelationIsdefaultOk returns a tuple with the BCreditcardclientrelationIsdefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBCreditcardclientrelationIsdefault

`func (o *CreditcardclientResponse) SetBCreditcardclientrelationIsdefault(v bool)`

SetBCreditcardclientrelationIsdefault sets BCreditcardclientrelationIsdefault field to given value.


### GetSCreditcardclientDescription

`func (o *CreditcardclientResponse) GetSCreditcardclientDescription() string`

GetSCreditcardclientDescription returns the SCreditcardclientDescription field if non-nil, zero value otherwise.

### GetSCreditcardclientDescriptionOk

`func (o *CreditcardclientResponse) GetSCreditcardclientDescriptionOk() (*string, bool)`

GetSCreditcardclientDescriptionOk returns a tuple with the SCreditcardclientDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSCreditcardclientDescription

`func (o *CreditcardclientResponse) SetSCreditcardclientDescription(v string)`

SetSCreditcardclientDescription sets SCreditcardclientDescription field to given value.


### GetBCreditcardclientAllowedcompanypayment

`func (o *CreditcardclientResponse) GetBCreditcardclientAllowedcompanypayment() bool`

GetBCreditcardclientAllowedcompanypayment returns the BCreditcardclientAllowedcompanypayment field if non-nil, zero value otherwise.

### GetBCreditcardclientAllowedcompanypaymentOk

`func (o *CreditcardclientResponse) GetBCreditcardclientAllowedcompanypaymentOk() (*bool, bool)`

GetBCreditcardclientAllowedcompanypaymentOk returns a tuple with the BCreditcardclientAllowedcompanypayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBCreditcardclientAllowedcompanypayment

`func (o *CreditcardclientResponse) SetBCreditcardclientAllowedcompanypayment(v bool)`

SetBCreditcardclientAllowedcompanypayment sets BCreditcardclientAllowedcompanypayment field to given value.


### GetBCreditcardclientAllowedtranquillit

`func (o *CreditcardclientResponse) GetBCreditcardclientAllowedtranquillit() bool`

GetBCreditcardclientAllowedtranquillit returns the BCreditcardclientAllowedtranquillit field if non-nil, zero value otherwise.

### GetBCreditcardclientAllowedtranquillitOk

`func (o *CreditcardclientResponse) GetBCreditcardclientAllowedtranquillitOk() (*bool, bool)`

GetBCreditcardclientAllowedtranquillitOk returns a tuple with the BCreditcardclientAllowedtranquillit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBCreditcardclientAllowedtranquillit

`func (o *CreditcardclientResponse) SetBCreditcardclientAllowedtranquillit(v bool)`

SetBCreditcardclientAllowedtranquillit sets BCreditcardclientAllowedtranquillit field to given value.


### GetObjCreditcarddetail

`func (o *CreditcardclientResponse) GetObjCreditcarddetail() CreditcarddetailResponseCompound`

GetObjCreditcarddetail returns the ObjCreditcarddetail field if non-nil, zero value otherwise.

### GetObjCreditcarddetailOk

`func (o *CreditcardclientResponse) GetObjCreditcarddetailOk() (*CreditcarddetailResponseCompound, bool)`

GetObjCreditcarddetailOk returns a tuple with the ObjCreditcarddetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjCreditcarddetail

`func (o *CreditcardclientResponse) SetObjCreditcarddetail(v CreditcarddetailResponseCompound)`

SetObjCreditcarddetail sets ObjCreditcarddetail field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


