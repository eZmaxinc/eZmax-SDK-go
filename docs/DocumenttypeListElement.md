# DocumenttypeListElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiDocumenttypeID** | **int32** | The unique ID of the Documenttype | 
**SDocumenttypeNameX** | **string** | The name of the Documenttype in the language of the requester | 
**SDocumenttypeDefaultfilename** | **string** | The defaultfilename of the Documenttype | 
**BDocumenttypeShowagent** | **bool** | Whether if it&#39;s an showagent | 
**BDocumenttypeIsactive** | **bool** | Whether the documenttype is active or not | 
**BDocumenttypeIsrequired** | **bool** | Whether if it&#39;s an isrequired | 
**BDocumenttypeSale** | **bool** | Whether if it&#39;s an sale | 
**BDocumenttypeRental** | **bool** | Whether if it&#39;s an rental | 
**BDocumenttypeReference** | **bool** | Whether if it&#39;s an reference | 
**BDocumenttypeMortgagereference** | **bool** | Whether if it&#39;s an mortgagereference | 
**BDocumenttypeEvaluation** | **bool** | Whether if it&#39;s an evaluation | 
**BDocumenttypeConsultation** | **bool** | Whether if it&#39;s an consultation | 
**BDocumenttypeOtherservices** | **bool** | Whether if it&#39;s an otherservices | 
**BDocumenttypeForeclosedsale** | **bool** | Whether if it&#39;s an foreclosedsale | 
**BDocumenttypeCommercialfunding** | **bool** | Whether if it&#39;s an commercialfunding | 
**BDocumenttypeCanafe** | **bool** | Whether if it&#39;s an canafe | 
**EDocumenttypeEzcom** | Pointer to [**FieldEDocumenttypeEzcom**](FieldEDocumenttypeEzcom.md) |  | [optional] 
**EDocumenttypePrivacydefault** | [**FieldEDocumenttypePrivacydefault**](FieldEDocumenttypePrivacydefault.md) |  | 
**EDocumenttypeStatusdefault** | [**FieldEDocumenttypeStatusdefault**](FieldEDocumenttypeStatusdefault.md) |  | 
**EDocumenttypeSalebyowner** | [**FieldEDocumenttypeSalebyowner**](FieldEDocumenttypeSalebyowner.md) |  | 
**EDocumenttypeSaledisplaybyrole** | [**FieldEDocumenttypeSaledisplaybyrole**](FieldEDocumenttypeSaledisplaybyrole.md) |  | 
**EDocumenttypeTransactiontype** | [**FieldEDocumenttypeTransactiontype**](FieldEDocumenttypeTransactiontype.md) |  | 
**EDocumenttypeRepossession** | [**FieldEDocumenttypeRepossession**](FieldEDocumenttypeRepossession.md) |  | 

## Methods

### NewDocumenttypeListElement

`func NewDocumenttypeListElement(pkiDocumenttypeID int32, sDocumenttypeNameX string, sDocumenttypeDefaultfilename string, bDocumenttypeShowagent bool, bDocumenttypeIsactive bool, bDocumenttypeIsrequired bool, bDocumenttypeSale bool, bDocumenttypeRental bool, bDocumenttypeReference bool, bDocumenttypeMortgagereference bool, bDocumenttypeEvaluation bool, bDocumenttypeConsultation bool, bDocumenttypeOtherservices bool, bDocumenttypeForeclosedsale bool, bDocumenttypeCommercialfunding bool, bDocumenttypeCanafe bool, eDocumenttypePrivacydefault FieldEDocumenttypePrivacydefault, eDocumenttypeStatusdefault FieldEDocumenttypeStatusdefault, eDocumenttypeSalebyowner FieldEDocumenttypeSalebyowner, eDocumenttypeSaledisplaybyrole FieldEDocumenttypeSaledisplaybyrole, eDocumenttypeTransactiontype FieldEDocumenttypeTransactiontype, eDocumenttypeRepossession FieldEDocumenttypeRepossession, ) *DocumenttypeListElement`

NewDocumenttypeListElement instantiates a new DocumenttypeListElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumenttypeListElementWithDefaults

`func NewDocumenttypeListElementWithDefaults() *DocumenttypeListElement`

NewDocumenttypeListElementWithDefaults instantiates a new DocumenttypeListElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiDocumenttypeID

`func (o *DocumenttypeListElement) GetPkiDocumenttypeID() int32`

GetPkiDocumenttypeID returns the PkiDocumenttypeID field if non-nil, zero value otherwise.

### GetPkiDocumenttypeIDOk

`func (o *DocumenttypeListElement) GetPkiDocumenttypeIDOk() (*int32, bool)`

GetPkiDocumenttypeIDOk returns a tuple with the PkiDocumenttypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiDocumenttypeID

`func (o *DocumenttypeListElement) SetPkiDocumenttypeID(v int32)`

SetPkiDocumenttypeID sets PkiDocumenttypeID field to given value.


### GetSDocumenttypeNameX

`func (o *DocumenttypeListElement) GetSDocumenttypeNameX() string`

GetSDocumenttypeNameX returns the SDocumenttypeNameX field if non-nil, zero value otherwise.

### GetSDocumenttypeNameXOk

`func (o *DocumenttypeListElement) GetSDocumenttypeNameXOk() (*string, bool)`

GetSDocumenttypeNameXOk returns a tuple with the SDocumenttypeNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSDocumenttypeNameX

`func (o *DocumenttypeListElement) SetSDocumenttypeNameX(v string)`

SetSDocumenttypeNameX sets SDocumenttypeNameX field to given value.


### GetSDocumenttypeDefaultfilename

`func (o *DocumenttypeListElement) GetSDocumenttypeDefaultfilename() string`

GetSDocumenttypeDefaultfilename returns the SDocumenttypeDefaultfilename field if non-nil, zero value otherwise.

### GetSDocumenttypeDefaultfilenameOk

`func (o *DocumenttypeListElement) GetSDocumenttypeDefaultfilenameOk() (*string, bool)`

GetSDocumenttypeDefaultfilenameOk returns a tuple with the SDocumenttypeDefaultfilename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSDocumenttypeDefaultfilename

`func (o *DocumenttypeListElement) SetSDocumenttypeDefaultfilename(v string)`

SetSDocumenttypeDefaultfilename sets SDocumenttypeDefaultfilename field to given value.


### GetBDocumenttypeShowagent

`func (o *DocumenttypeListElement) GetBDocumenttypeShowagent() bool`

GetBDocumenttypeShowagent returns the BDocumenttypeShowagent field if non-nil, zero value otherwise.

### GetBDocumenttypeShowagentOk

`func (o *DocumenttypeListElement) GetBDocumenttypeShowagentOk() (*bool, bool)`

GetBDocumenttypeShowagentOk returns a tuple with the BDocumenttypeShowagent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBDocumenttypeShowagent

`func (o *DocumenttypeListElement) SetBDocumenttypeShowagent(v bool)`

SetBDocumenttypeShowagent sets BDocumenttypeShowagent field to given value.


### GetBDocumenttypeIsactive

`func (o *DocumenttypeListElement) GetBDocumenttypeIsactive() bool`

GetBDocumenttypeIsactive returns the BDocumenttypeIsactive field if non-nil, zero value otherwise.

### GetBDocumenttypeIsactiveOk

`func (o *DocumenttypeListElement) GetBDocumenttypeIsactiveOk() (*bool, bool)`

GetBDocumenttypeIsactiveOk returns a tuple with the BDocumenttypeIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBDocumenttypeIsactive

`func (o *DocumenttypeListElement) SetBDocumenttypeIsactive(v bool)`

SetBDocumenttypeIsactive sets BDocumenttypeIsactive field to given value.


### GetBDocumenttypeIsrequired

`func (o *DocumenttypeListElement) GetBDocumenttypeIsrequired() bool`

GetBDocumenttypeIsrequired returns the BDocumenttypeIsrequired field if non-nil, zero value otherwise.

### GetBDocumenttypeIsrequiredOk

`func (o *DocumenttypeListElement) GetBDocumenttypeIsrequiredOk() (*bool, bool)`

GetBDocumenttypeIsrequiredOk returns a tuple with the BDocumenttypeIsrequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBDocumenttypeIsrequired

`func (o *DocumenttypeListElement) SetBDocumenttypeIsrequired(v bool)`

SetBDocumenttypeIsrequired sets BDocumenttypeIsrequired field to given value.


### GetBDocumenttypeSale

`func (o *DocumenttypeListElement) GetBDocumenttypeSale() bool`

GetBDocumenttypeSale returns the BDocumenttypeSale field if non-nil, zero value otherwise.

### GetBDocumenttypeSaleOk

`func (o *DocumenttypeListElement) GetBDocumenttypeSaleOk() (*bool, bool)`

GetBDocumenttypeSaleOk returns a tuple with the BDocumenttypeSale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBDocumenttypeSale

`func (o *DocumenttypeListElement) SetBDocumenttypeSale(v bool)`

SetBDocumenttypeSale sets BDocumenttypeSale field to given value.


### GetBDocumenttypeRental

`func (o *DocumenttypeListElement) GetBDocumenttypeRental() bool`

GetBDocumenttypeRental returns the BDocumenttypeRental field if non-nil, zero value otherwise.

### GetBDocumenttypeRentalOk

`func (o *DocumenttypeListElement) GetBDocumenttypeRentalOk() (*bool, bool)`

GetBDocumenttypeRentalOk returns a tuple with the BDocumenttypeRental field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBDocumenttypeRental

`func (o *DocumenttypeListElement) SetBDocumenttypeRental(v bool)`

SetBDocumenttypeRental sets BDocumenttypeRental field to given value.


### GetBDocumenttypeReference

`func (o *DocumenttypeListElement) GetBDocumenttypeReference() bool`

GetBDocumenttypeReference returns the BDocumenttypeReference field if non-nil, zero value otherwise.

### GetBDocumenttypeReferenceOk

`func (o *DocumenttypeListElement) GetBDocumenttypeReferenceOk() (*bool, bool)`

GetBDocumenttypeReferenceOk returns a tuple with the BDocumenttypeReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBDocumenttypeReference

`func (o *DocumenttypeListElement) SetBDocumenttypeReference(v bool)`

SetBDocumenttypeReference sets BDocumenttypeReference field to given value.


### GetBDocumenttypeMortgagereference

`func (o *DocumenttypeListElement) GetBDocumenttypeMortgagereference() bool`

GetBDocumenttypeMortgagereference returns the BDocumenttypeMortgagereference field if non-nil, zero value otherwise.

### GetBDocumenttypeMortgagereferenceOk

`func (o *DocumenttypeListElement) GetBDocumenttypeMortgagereferenceOk() (*bool, bool)`

GetBDocumenttypeMortgagereferenceOk returns a tuple with the BDocumenttypeMortgagereference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBDocumenttypeMortgagereference

`func (o *DocumenttypeListElement) SetBDocumenttypeMortgagereference(v bool)`

SetBDocumenttypeMortgagereference sets BDocumenttypeMortgagereference field to given value.


### GetBDocumenttypeEvaluation

`func (o *DocumenttypeListElement) GetBDocumenttypeEvaluation() bool`

GetBDocumenttypeEvaluation returns the BDocumenttypeEvaluation field if non-nil, zero value otherwise.

### GetBDocumenttypeEvaluationOk

`func (o *DocumenttypeListElement) GetBDocumenttypeEvaluationOk() (*bool, bool)`

GetBDocumenttypeEvaluationOk returns a tuple with the BDocumenttypeEvaluation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBDocumenttypeEvaluation

`func (o *DocumenttypeListElement) SetBDocumenttypeEvaluation(v bool)`

SetBDocumenttypeEvaluation sets BDocumenttypeEvaluation field to given value.


### GetBDocumenttypeConsultation

`func (o *DocumenttypeListElement) GetBDocumenttypeConsultation() bool`

GetBDocumenttypeConsultation returns the BDocumenttypeConsultation field if non-nil, zero value otherwise.

### GetBDocumenttypeConsultationOk

`func (o *DocumenttypeListElement) GetBDocumenttypeConsultationOk() (*bool, bool)`

GetBDocumenttypeConsultationOk returns a tuple with the BDocumenttypeConsultation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBDocumenttypeConsultation

`func (o *DocumenttypeListElement) SetBDocumenttypeConsultation(v bool)`

SetBDocumenttypeConsultation sets BDocumenttypeConsultation field to given value.


### GetBDocumenttypeOtherservices

`func (o *DocumenttypeListElement) GetBDocumenttypeOtherservices() bool`

GetBDocumenttypeOtherservices returns the BDocumenttypeOtherservices field if non-nil, zero value otherwise.

### GetBDocumenttypeOtherservicesOk

`func (o *DocumenttypeListElement) GetBDocumenttypeOtherservicesOk() (*bool, bool)`

GetBDocumenttypeOtherservicesOk returns a tuple with the BDocumenttypeOtherservices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBDocumenttypeOtherservices

`func (o *DocumenttypeListElement) SetBDocumenttypeOtherservices(v bool)`

SetBDocumenttypeOtherservices sets BDocumenttypeOtherservices field to given value.


### GetBDocumenttypeForeclosedsale

`func (o *DocumenttypeListElement) GetBDocumenttypeForeclosedsale() bool`

GetBDocumenttypeForeclosedsale returns the BDocumenttypeForeclosedsale field if non-nil, zero value otherwise.

### GetBDocumenttypeForeclosedsaleOk

`func (o *DocumenttypeListElement) GetBDocumenttypeForeclosedsaleOk() (*bool, bool)`

GetBDocumenttypeForeclosedsaleOk returns a tuple with the BDocumenttypeForeclosedsale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBDocumenttypeForeclosedsale

`func (o *DocumenttypeListElement) SetBDocumenttypeForeclosedsale(v bool)`

SetBDocumenttypeForeclosedsale sets BDocumenttypeForeclosedsale field to given value.


### GetBDocumenttypeCommercialfunding

`func (o *DocumenttypeListElement) GetBDocumenttypeCommercialfunding() bool`

GetBDocumenttypeCommercialfunding returns the BDocumenttypeCommercialfunding field if non-nil, zero value otherwise.

### GetBDocumenttypeCommercialfundingOk

`func (o *DocumenttypeListElement) GetBDocumenttypeCommercialfundingOk() (*bool, bool)`

GetBDocumenttypeCommercialfundingOk returns a tuple with the BDocumenttypeCommercialfunding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBDocumenttypeCommercialfunding

`func (o *DocumenttypeListElement) SetBDocumenttypeCommercialfunding(v bool)`

SetBDocumenttypeCommercialfunding sets BDocumenttypeCommercialfunding field to given value.


### GetBDocumenttypeCanafe

`func (o *DocumenttypeListElement) GetBDocumenttypeCanafe() bool`

GetBDocumenttypeCanafe returns the BDocumenttypeCanafe field if non-nil, zero value otherwise.

### GetBDocumenttypeCanafeOk

`func (o *DocumenttypeListElement) GetBDocumenttypeCanafeOk() (*bool, bool)`

GetBDocumenttypeCanafeOk returns a tuple with the BDocumenttypeCanafe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBDocumenttypeCanafe

`func (o *DocumenttypeListElement) SetBDocumenttypeCanafe(v bool)`

SetBDocumenttypeCanafe sets BDocumenttypeCanafe field to given value.


### GetEDocumenttypeEzcom

`func (o *DocumenttypeListElement) GetEDocumenttypeEzcom() FieldEDocumenttypeEzcom`

GetEDocumenttypeEzcom returns the EDocumenttypeEzcom field if non-nil, zero value otherwise.

### GetEDocumenttypeEzcomOk

`func (o *DocumenttypeListElement) GetEDocumenttypeEzcomOk() (*FieldEDocumenttypeEzcom, bool)`

GetEDocumenttypeEzcomOk returns a tuple with the EDocumenttypeEzcom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEDocumenttypeEzcom

`func (o *DocumenttypeListElement) SetEDocumenttypeEzcom(v FieldEDocumenttypeEzcom)`

SetEDocumenttypeEzcom sets EDocumenttypeEzcom field to given value.

### HasEDocumenttypeEzcom

`func (o *DocumenttypeListElement) HasEDocumenttypeEzcom() bool`

HasEDocumenttypeEzcom returns a boolean if a field has been set.

### GetEDocumenttypePrivacydefault

`func (o *DocumenttypeListElement) GetEDocumenttypePrivacydefault() FieldEDocumenttypePrivacydefault`

GetEDocumenttypePrivacydefault returns the EDocumenttypePrivacydefault field if non-nil, zero value otherwise.

### GetEDocumenttypePrivacydefaultOk

`func (o *DocumenttypeListElement) GetEDocumenttypePrivacydefaultOk() (*FieldEDocumenttypePrivacydefault, bool)`

GetEDocumenttypePrivacydefaultOk returns a tuple with the EDocumenttypePrivacydefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEDocumenttypePrivacydefault

`func (o *DocumenttypeListElement) SetEDocumenttypePrivacydefault(v FieldEDocumenttypePrivacydefault)`

SetEDocumenttypePrivacydefault sets EDocumenttypePrivacydefault field to given value.


### GetEDocumenttypeStatusdefault

`func (o *DocumenttypeListElement) GetEDocumenttypeStatusdefault() FieldEDocumenttypeStatusdefault`

GetEDocumenttypeStatusdefault returns the EDocumenttypeStatusdefault field if non-nil, zero value otherwise.

### GetEDocumenttypeStatusdefaultOk

`func (o *DocumenttypeListElement) GetEDocumenttypeStatusdefaultOk() (*FieldEDocumenttypeStatusdefault, bool)`

GetEDocumenttypeStatusdefaultOk returns a tuple with the EDocumenttypeStatusdefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEDocumenttypeStatusdefault

`func (o *DocumenttypeListElement) SetEDocumenttypeStatusdefault(v FieldEDocumenttypeStatusdefault)`

SetEDocumenttypeStatusdefault sets EDocumenttypeStatusdefault field to given value.


### GetEDocumenttypeSalebyowner

`func (o *DocumenttypeListElement) GetEDocumenttypeSalebyowner() FieldEDocumenttypeSalebyowner`

GetEDocumenttypeSalebyowner returns the EDocumenttypeSalebyowner field if non-nil, zero value otherwise.

### GetEDocumenttypeSalebyownerOk

`func (o *DocumenttypeListElement) GetEDocumenttypeSalebyownerOk() (*FieldEDocumenttypeSalebyowner, bool)`

GetEDocumenttypeSalebyownerOk returns a tuple with the EDocumenttypeSalebyowner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEDocumenttypeSalebyowner

`func (o *DocumenttypeListElement) SetEDocumenttypeSalebyowner(v FieldEDocumenttypeSalebyowner)`

SetEDocumenttypeSalebyowner sets EDocumenttypeSalebyowner field to given value.


### GetEDocumenttypeSaledisplaybyrole

`func (o *DocumenttypeListElement) GetEDocumenttypeSaledisplaybyrole() FieldEDocumenttypeSaledisplaybyrole`

GetEDocumenttypeSaledisplaybyrole returns the EDocumenttypeSaledisplaybyrole field if non-nil, zero value otherwise.

### GetEDocumenttypeSaledisplaybyroleOk

`func (o *DocumenttypeListElement) GetEDocumenttypeSaledisplaybyroleOk() (*FieldEDocumenttypeSaledisplaybyrole, bool)`

GetEDocumenttypeSaledisplaybyroleOk returns a tuple with the EDocumenttypeSaledisplaybyrole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEDocumenttypeSaledisplaybyrole

`func (o *DocumenttypeListElement) SetEDocumenttypeSaledisplaybyrole(v FieldEDocumenttypeSaledisplaybyrole)`

SetEDocumenttypeSaledisplaybyrole sets EDocumenttypeSaledisplaybyrole field to given value.


### GetEDocumenttypeTransactiontype

`func (o *DocumenttypeListElement) GetEDocumenttypeTransactiontype() FieldEDocumenttypeTransactiontype`

GetEDocumenttypeTransactiontype returns the EDocumenttypeTransactiontype field if non-nil, zero value otherwise.

### GetEDocumenttypeTransactiontypeOk

`func (o *DocumenttypeListElement) GetEDocumenttypeTransactiontypeOk() (*FieldEDocumenttypeTransactiontype, bool)`

GetEDocumenttypeTransactiontypeOk returns a tuple with the EDocumenttypeTransactiontype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEDocumenttypeTransactiontype

`func (o *DocumenttypeListElement) SetEDocumenttypeTransactiontype(v FieldEDocumenttypeTransactiontype)`

SetEDocumenttypeTransactiontype sets EDocumenttypeTransactiontype field to given value.


### GetEDocumenttypeRepossession

`func (o *DocumenttypeListElement) GetEDocumenttypeRepossession() FieldEDocumenttypeRepossession`

GetEDocumenttypeRepossession returns the EDocumenttypeRepossession field if non-nil, zero value otherwise.

### GetEDocumenttypeRepossessionOk

`func (o *DocumenttypeListElement) GetEDocumenttypeRepossessionOk() (*FieldEDocumenttypeRepossession, bool)`

GetEDocumenttypeRepossessionOk returns a tuple with the EDocumenttypeRepossession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEDocumenttypeRepossession

`func (o *DocumenttypeListElement) SetEDocumenttypeRepossession(v FieldEDocumenttypeRepossession)`

SetEDocumenttypeRepossession sets EDocumenttypeRepossession field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


