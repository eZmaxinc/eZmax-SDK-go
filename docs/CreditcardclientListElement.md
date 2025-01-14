# CreditcardclientListElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiCreditcardclientID** | **int32** | The unique ID of the Creditcardclient | 
**FkiCreditcarddetailID** | **int32** | The unique ID of the Creditcarddetail | 
**FkiCreditcardtypeID** | **int32** | The unique ID of the Creditcardtype | 
**BCreditcardclientrelationIsdefault** | **bool** | Whether if it&#39;s the creditcardclient is the default one | 
**SCreditcardclientDescription** | **string** | The description of the Creditcardclient | 
**BCreditcardclientAllowedcompanypayment** | **bool** | Whether if it&#39;s an allowedagencypayment | 
**BCreditcardclientAllowedtranquillit** | **bool** | Whether if it&#39;s an allowedtranquillit | 
**ICreditcarddetailExpirationmonth** | **int32** | The expirationmonth of the Creditcarddetail | 
**ICreditcarddetailExpirationyear** | **int32** | The expirationyear of the Creditcarddetail | 
**ICreditcarddetailLastdigits** | **int32** | The last digits of the Creditcarddetail | 

## Methods

### NewCreditcardclientListElement

`func NewCreditcardclientListElement(pkiCreditcardclientID int32, fkiCreditcarddetailID int32, fkiCreditcardtypeID int32, bCreditcardclientrelationIsdefault bool, sCreditcardclientDescription string, bCreditcardclientAllowedcompanypayment bool, bCreditcardclientAllowedtranquillit bool, iCreditcarddetailExpirationmonth int32, iCreditcarddetailExpirationyear int32, iCreditcarddetailLastdigits int32, ) *CreditcardclientListElement`

NewCreditcardclientListElement instantiates a new CreditcardclientListElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditcardclientListElementWithDefaults

`func NewCreditcardclientListElementWithDefaults() *CreditcardclientListElement`

NewCreditcardclientListElementWithDefaults instantiates a new CreditcardclientListElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiCreditcardclientID

`func (o *CreditcardclientListElement) GetPkiCreditcardclientID() int32`

GetPkiCreditcardclientID returns the PkiCreditcardclientID field if non-nil, zero value otherwise.

### GetPkiCreditcardclientIDOk

`func (o *CreditcardclientListElement) GetPkiCreditcardclientIDOk() (*int32, bool)`

GetPkiCreditcardclientIDOk returns a tuple with the PkiCreditcardclientID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiCreditcardclientID

`func (o *CreditcardclientListElement) SetPkiCreditcardclientID(v int32)`

SetPkiCreditcardclientID sets PkiCreditcardclientID field to given value.


### GetFkiCreditcarddetailID

`func (o *CreditcardclientListElement) GetFkiCreditcarddetailID() int32`

GetFkiCreditcarddetailID returns the FkiCreditcarddetailID field if non-nil, zero value otherwise.

### GetFkiCreditcarddetailIDOk

`func (o *CreditcardclientListElement) GetFkiCreditcarddetailIDOk() (*int32, bool)`

GetFkiCreditcarddetailIDOk returns a tuple with the FkiCreditcarddetailID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiCreditcarddetailID

`func (o *CreditcardclientListElement) SetFkiCreditcarddetailID(v int32)`

SetFkiCreditcarddetailID sets FkiCreditcarddetailID field to given value.


### GetFkiCreditcardtypeID

`func (o *CreditcardclientListElement) GetFkiCreditcardtypeID() int32`

GetFkiCreditcardtypeID returns the FkiCreditcardtypeID field if non-nil, zero value otherwise.

### GetFkiCreditcardtypeIDOk

`func (o *CreditcardclientListElement) GetFkiCreditcardtypeIDOk() (*int32, bool)`

GetFkiCreditcardtypeIDOk returns a tuple with the FkiCreditcardtypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiCreditcardtypeID

`func (o *CreditcardclientListElement) SetFkiCreditcardtypeID(v int32)`

SetFkiCreditcardtypeID sets FkiCreditcardtypeID field to given value.


### GetBCreditcardclientrelationIsdefault

`func (o *CreditcardclientListElement) GetBCreditcardclientrelationIsdefault() bool`

GetBCreditcardclientrelationIsdefault returns the BCreditcardclientrelationIsdefault field if non-nil, zero value otherwise.

### GetBCreditcardclientrelationIsdefaultOk

`func (o *CreditcardclientListElement) GetBCreditcardclientrelationIsdefaultOk() (*bool, bool)`

GetBCreditcardclientrelationIsdefaultOk returns a tuple with the BCreditcardclientrelationIsdefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBCreditcardclientrelationIsdefault

`func (o *CreditcardclientListElement) SetBCreditcardclientrelationIsdefault(v bool)`

SetBCreditcardclientrelationIsdefault sets BCreditcardclientrelationIsdefault field to given value.


### GetSCreditcardclientDescription

`func (o *CreditcardclientListElement) GetSCreditcardclientDescription() string`

GetSCreditcardclientDescription returns the SCreditcardclientDescription field if non-nil, zero value otherwise.

### GetSCreditcardclientDescriptionOk

`func (o *CreditcardclientListElement) GetSCreditcardclientDescriptionOk() (*string, bool)`

GetSCreditcardclientDescriptionOk returns a tuple with the SCreditcardclientDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSCreditcardclientDescription

`func (o *CreditcardclientListElement) SetSCreditcardclientDescription(v string)`

SetSCreditcardclientDescription sets SCreditcardclientDescription field to given value.


### GetBCreditcardclientAllowedcompanypayment

`func (o *CreditcardclientListElement) GetBCreditcardclientAllowedcompanypayment() bool`

GetBCreditcardclientAllowedcompanypayment returns the BCreditcardclientAllowedcompanypayment field if non-nil, zero value otherwise.

### GetBCreditcardclientAllowedcompanypaymentOk

`func (o *CreditcardclientListElement) GetBCreditcardclientAllowedcompanypaymentOk() (*bool, bool)`

GetBCreditcardclientAllowedcompanypaymentOk returns a tuple with the BCreditcardclientAllowedcompanypayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBCreditcardclientAllowedcompanypayment

`func (o *CreditcardclientListElement) SetBCreditcardclientAllowedcompanypayment(v bool)`

SetBCreditcardclientAllowedcompanypayment sets BCreditcardclientAllowedcompanypayment field to given value.


### GetBCreditcardclientAllowedtranquillit

`func (o *CreditcardclientListElement) GetBCreditcardclientAllowedtranquillit() bool`

GetBCreditcardclientAllowedtranquillit returns the BCreditcardclientAllowedtranquillit field if non-nil, zero value otherwise.

### GetBCreditcardclientAllowedtranquillitOk

`func (o *CreditcardclientListElement) GetBCreditcardclientAllowedtranquillitOk() (*bool, bool)`

GetBCreditcardclientAllowedtranquillitOk returns a tuple with the BCreditcardclientAllowedtranquillit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBCreditcardclientAllowedtranquillit

`func (o *CreditcardclientListElement) SetBCreditcardclientAllowedtranquillit(v bool)`

SetBCreditcardclientAllowedtranquillit sets BCreditcardclientAllowedtranquillit field to given value.


### GetICreditcarddetailExpirationmonth

`func (o *CreditcardclientListElement) GetICreditcarddetailExpirationmonth() int32`

GetICreditcarddetailExpirationmonth returns the ICreditcarddetailExpirationmonth field if non-nil, zero value otherwise.

### GetICreditcarddetailExpirationmonthOk

`func (o *CreditcardclientListElement) GetICreditcarddetailExpirationmonthOk() (*int32, bool)`

GetICreditcarddetailExpirationmonthOk returns a tuple with the ICreditcarddetailExpirationmonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetICreditcarddetailExpirationmonth

`func (o *CreditcardclientListElement) SetICreditcarddetailExpirationmonth(v int32)`

SetICreditcarddetailExpirationmonth sets ICreditcarddetailExpirationmonth field to given value.


### GetICreditcarddetailExpirationyear

`func (o *CreditcardclientListElement) GetICreditcarddetailExpirationyear() int32`

GetICreditcarddetailExpirationyear returns the ICreditcarddetailExpirationyear field if non-nil, zero value otherwise.

### GetICreditcarddetailExpirationyearOk

`func (o *CreditcardclientListElement) GetICreditcarddetailExpirationyearOk() (*int32, bool)`

GetICreditcarddetailExpirationyearOk returns a tuple with the ICreditcarddetailExpirationyear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetICreditcarddetailExpirationyear

`func (o *CreditcardclientListElement) SetICreditcarddetailExpirationyear(v int32)`

SetICreditcarddetailExpirationyear sets ICreditcarddetailExpirationyear field to given value.


### GetICreditcarddetailLastdigits

`func (o *CreditcardclientListElement) GetICreditcarddetailLastdigits() int32`

GetICreditcarddetailLastdigits returns the ICreditcarddetailLastdigits field if non-nil, zero value otherwise.

### GetICreditcarddetailLastdigitsOk

`func (o *CreditcardclientListElement) GetICreditcarddetailLastdigitsOk() (*int32, bool)`

GetICreditcarddetailLastdigitsOk returns a tuple with the ICreditcarddetailLastdigits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetICreditcarddetailLastdigits

`func (o *CreditcardclientListElement) SetICreditcarddetailLastdigits(v int32)`

SetICreditcarddetailLastdigits sets ICreditcarddetailLastdigits field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


