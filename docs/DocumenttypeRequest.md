# DocumenttypeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiDocumenttypeID** | Pointer to **int32** | The unique ID of the Documenttype | [optional] 
**ObjDocumenttypeName** | [**MultilingualDocumenttypeName**](MultilingualDocumenttypeName.md) |  | 
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
**EDocumenttypeEzcom** | [**FieldEDocumenttypeEzcom**](FieldEDocumenttypeEzcom.md) |  | 
**EDocumenttypePrivacydefault** | [**FieldEDocumenttypePrivacydefault**](FieldEDocumenttypePrivacydefault.md) |  | 
**EDocumenttypeStatusdefault** | [**FieldEDocumenttypeStatusdefault**](FieldEDocumenttypeStatusdefault.md) |  | 
**EDocumenttypeSalebyowner** | [**FieldEDocumenttypeSalebyowner**](FieldEDocumenttypeSalebyowner.md) |  | 
**EDocumenttypeSaledisplaybyrole** | [**FieldEDocumenttypeSaledisplaybyrole**](FieldEDocumenttypeSaledisplaybyrole.md) |  | 
**EDocumenttypeTransactiontype** | [**FieldEDocumenttypeTransactiontype**](FieldEDocumenttypeTransactiontype.md) |  | 
**EDocumenttypeRepossession** | [**FieldEDocumenttypeRepossession**](FieldEDocumenttypeRepossession.md) |  | 

## Methods

### NewDocumenttypeRequest

`func NewDocumenttypeRequest(objDocumenttypeName MultilingualDocumenttypeName, sDocumenttypeDefaultfilename string, bDocumenttypeShowagent bool, bDocumenttypeIsactive bool, bDocumenttypeIsrequired bool, bDocumenttypeSale bool, bDocumenttypeRental bool, bDocumenttypeReference bool, bDocumenttypeMortgagereference bool, bDocumenttypeEvaluation bool, bDocumenttypeConsultation bool, bDocumenttypeOtherservices bool, bDocumenttypeForeclosedsale bool, bDocumenttypeCommercialfunding bool, bDocumenttypeCanafe bool, eDocumenttypeEzcom FieldEDocumenttypeEzcom, eDocumenttypePrivacydefault FieldEDocumenttypePrivacydefault, eDocumenttypeStatusdefault FieldEDocumenttypeStatusdefault, eDocumenttypeSalebyowner FieldEDocumenttypeSalebyowner, eDocumenttypeSaledisplaybyrole FieldEDocumenttypeSaledisplaybyrole, eDocumenttypeTransactiontype FieldEDocumenttypeTransactiontype, eDocumenttypeRepossession FieldEDocumenttypeRepossession, ) *DocumenttypeRequest`

NewDocumenttypeRequest instantiates a new DocumenttypeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumenttypeRequestWithDefaults

`func NewDocumenttypeRequestWithDefaults() *DocumenttypeRequest`

NewDocumenttypeRequestWithDefaults instantiates a new DocumenttypeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiDocumenttypeID

`func (o *DocumenttypeRequest) GetPkiDocumenttypeID() int32`

GetPkiDocumenttypeID returns the PkiDocumenttypeID field if non-nil, zero value otherwise.

### GetPkiDocumenttypeIDOk

`func (o *DocumenttypeRequest) GetPkiDocumenttypeIDOk() (*int32, bool)`

GetPkiDocumenttypeIDOk returns a tuple with the PkiDocumenttypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiDocumenttypeID

`func (o *DocumenttypeRequest) SetPkiDocumenttypeID(v int32)`

SetPkiDocumenttypeID sets PkiDocumenttypeID field to given value.

### HasPkiDocumenttypeID

`func (o *DocumenttypeRequest) HasPkiDocumenttypeID() bool`

HasPkiDocumenttypeID returns a boolean if a field has been set.

### GetObjDocumenttypeName

`func (o *DocumenttypeRequest) GetObjDocumenttypeName() MultilingualDocumenttypeName`

GetObjDocumenttypeName returns the ObjDocumenttypeName field if non-nil, zero value otherwise.

### GetObjDocumenttypeNameOk

`func (o *DocumenttypeRequest) GetObjDocumenttypeNameOk() (*MultilingualDocumenttypeName, bool)`

GetObjDocumenttypeNameOk returns a tuple with the ObjDocumenttypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjDocumenttypeName

`func (o *DocumenttypeRequest) SetObjDocumenttypeName(v MultilingualDocumenttypeName)`

SetObjDocumenttypeName sets ObjDocumenttypeName field to given value.


### GetSDocumenttypeDefaultfilename

`func (o *DocumenttypeRequest) GetSDocumenttypeDefaultfilename() string`

GetSDocumenttypeDefaultfilename returns the SDocumenttypeDefaultfilename field if non-nil, zero value otherwise.

### GetSDocumenttypeDefaultfilenameOk

`func (o *DocumenttypeRequest) GetSDocumenttypeDefaultfilenameOk() (*string, bool)`

GetSDocumenttypeDefaultfilenameOk returns a tuple with the SDocumenttypeDefaultfilename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSDocumenttypeDefaultfilename

`func (o *DocumenttypeRequest) SetSDocumenttypeDefaultfilename(v string)`

SetSDocumenttypeDefaultfilename sets SDocumenttypeDefaultfilename field to given value.


### GetBDocumenttypeShowagent

`func (o *DocumenttypeRequest) GetBDocumenttypeShowagent() bool`

GetBDocumenttypeShowagent returns the BDocumenttypeShowagent field if non-nil, zero value otherwise.

### GetBDocumenttypeShowagentOk

`func (o *DocumenttypeRequest) GetBDocumenttypeShowagentOk() (*bool, bool)`

GetBDocumenttypeShowagentOk returns a tuple with the BDocumenttypeShowagent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBDocumenttypeShowagent

`func (o *DocumenttypeRequest) SetBDocumenttypeShowagent(v bool)`

SetBDocumenttypeShowagent sets BDocumenttypeShowagent field to given value.


### GetBDocumenttypeIsactive

`func (o *DocumenttypeRequest) GetBDocumenttypeIsactive() bool`

GetBDocumenttypeIsactive returns the BDocumenttypeIsactive field if non-nil, zero value otherwise.

### GetBDocumenttypeIsactiveOk

`func (o *DocumenttypeRequest) GetBDocumenttypeIsactiveOk() (*bool, bool)`

GetBDocumenttypeIsactiveOk returns a tuple with the BDocumenttypeIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBDocumenttypeIsactive

`func (o *DocumenttypeRequest) SetBDocumenttypeIsactive(v bool)`

SetBDocumenttypeIsactive sets BDocumenttypeIsactive field to given value.


### GetBDocumenttypeIsrequired

`func (o *DocumenttypeRequest) GetBDocumenttypeIsrequired() bool`

GetBDocumenttypeIsrequired returns the BDocumenttypeIsrequired field if non-nil, zero value otherwise.

### GetBDocumenttypeIsrequiredOk

`func (o *DocumenttypeRequest) GetBDocumenttypeIsrequiredOk() (*bool, bool)`

GetBDocumenttypeIsrequiredOk returns a tuple with the BDocumenttypeIsrequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBDocumenttypeIsrequired

`func (o *DocumenttypeRequest) SetBDocumenttypeIsrequired(v bool)`

SetBDocumenttypeIsrequired sets BDocumenttypeIsrequired field to given value.


### GetBDocumenttypeSale

`func (o *DocumenttypeRequest) GetBDocumenttypeSale() bool`

GetBDocumenttypeSale returns the BDocumenttypeSale field if non-nil, zero value otherwise.

### GetBDocumenttypeSaleOk

`func (o *DocumenttypeRequest) GetBDocumenttypeSaleOk() (*bool, bool)`

GetBDocumenttypeSaleOk returns a tuple with the BDocumenttypeSale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBDocumenttypeSale

`func (o *DocumenttypeRequest) SetBDocumenttypeSale(v bool)`

SetBDocumenttypeSale sets BDocumenttypeSale field to given value.


### GetBDocumenttypeRental

`func (o *DocumenttypeRequest) GetBDocumenttypeRental() bool`

GetBDocumenttypeRental returns the BDocumenttypeRental field if non-nil, zero value otherwise.

### GetBDocumenttypeRentalOk

`func (o *DocumenttypeRequest) GetBDocumenttypeRentalOk() (*bool, bool)`

GetBDocumenttypeRentalOk returns a tuple with the BDocumenttypeRental field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBDocumenttypeRental

`func (o *DocumenttypeRequest) SetBDocumenttypeRental(v bool)`

SetBDocumenttypeRental sets BDocumenttypeRental field to given value.


### GetBDocumenttypeReference

`func (o *DocumenttypeRequest) GetBDocumenttypeReference() bool`

GetBDocumenttypeReference returns the BDocumenttypeReference field if non-nil, zero value otherwise.

### GetBDocumenttypeReferenceOk

`func (o *DocumenttypeRequest) GetBDocumenttypeReferenceOk() (*bool, bool)`

GetBDocumenttypeReferenceOk returns a tuple with the BDocumenttypeReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBDocumenttypeReference

`func (o *DocumenttypeRequest) SetBDocumenttypeReference(v bool)`

SetBDocumenttypeReference sets BDocumenttypeReference field to given value.


### GetBDocumenttypeMortgagereference

`func (o *DocumenttypeRequest) GetBDocumenttypeMortgagereference() bool`

GetBDocumenttypeMortgagereference returns the BDocumenttypeMortgagereference field if non-nil, zero value otherwise.

### GetBDocumenttypeMortgagereferenceOk

`func (o *DocumenttypeRequest) GetBDocumenttypeMortgagereferenceOk() (*bool, bool)`

GetBDocumenttypeMortgagereferenceOk returns a tuple with the BDocumenttypeMortgagereference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBDocumenttypeMortgagereference

`func (o *DocumenttypeRequest) SetBDocumenttypeMortgagereference(v bool)`

SetBDocumenttypeMortgagereference sets BDocumenttypeMortgagereference field to given value.


### GetBDocumenttypeEvaluation

`func (o *DocumenttypeRequest) GetBDocumenttypeEvaluation() bool`

GetBDocumenttypeEvaluation returns the BDocumenttypeEvaluation field if non-nil, zero value otherwise.

### GetBDocumenttypeEvaluationOk

`func (o *DocumenttypeRequest) GetBDocumenttypeEvaluationOk() (*bool, bool)`

GetBDocumenttypeEvaluationOk returns a tuple with the BDocumenttypeEvaluation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBDocumenttypeEvaluation

`func (o *DocumenttypeRequest) SetBDocumenttypeEvaluation(v bool)`

SetBDocumenttypeEvaluation sets BDocumenttypeEvaluation field to given value.


### GetBDocumenttypeConsultation

`func (o *DocumenttypeRequest) GetBDocumenttypeConsultation() bool`

GetBDocumenttypeConsultation returns the BDocumenttypeConsultation field if non-nil, zero value otherwise.

### GetBDocumenttypeConsultationOk

`func (o *DocumenttypeRequest) GetBDocumenttypeConsultationOk() (*bool, bool)`

GetBDocumenttypeConsultationOk returns a tuple with the BDocumenttypeConsultation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBDocumenttypeConsultation

`func (o *DocumenttypeRequest) SetBDocumenttypeConsultation(v bool)`

SetBDocumenttypeConsultation sets BDocumenttypeConsultation field to given value.


### GetBDocumenttypeOtherservices

`func (o *DocumenttypeRequest) GetBDocumenttypeOtherservices() bool`

GetBDocumenttypeOtherservices returns the BDocumenttypeOtherservices field if non-nil, zero value otherwise.

### GetBDocumenttypeOtherservicesOk

`func (o *DocumenttypeRequest) GetBDocumenttypeOtherservicesOk() (*bool, bool)`

GetBDocumenttypeOtherservicesOk returns a tuple with the BDocumenttypeOtherservices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBDocumenttypeOtherservices

`func (o *DocumenttypeRequest) SetBDocumenttypeOtherservices(v bool)`

SetBDocumenttypeOtherservices sets BDocumenttypeOtherservices field to given value.


### GetBDocumenttypeForeclosedsale

`func (o *DocumenttypeRequest) GetBDocumenttypeForeclosedsale() bool`

GetBDocumenttypeForeclosedsale returns the BDocumenttypeForeclosedsale field if non-nil, zero value otherwise.

### GetBDocumenttypeForeclosedsaleOk

`func (o *DocumenttypeRequest) GetBDocumenttypeForeclosedsaleOk() (*bool, bool)`

GetBDocumenttypeForeclosedsaleOk returns a tuple with the BDocumenttypeForeclosedsale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBDocumenttypeForeclosedsale

`func (o *DocumenttypeRequest) SetBDocumenttypeForeclosedsale(v bool)`

SetBDocumenttypeForeclosedsale sets BDocumenttypeForeclosedsale field to given value.


### GetBDocumenttypeCommercialfunding

`func (o *DocumenttypeRequest) GetBDocumenttypeCommercialfunding() bool`

GetBDocumenttypeCommercialfunding returns the BDocumenttypeCommercialfunding field if non-nil, zero value otherwise.

### GetBDocumenttypeCommercialfundingOk

`func (o *DocumenttypeRequest) GetBDocumenttypeCommercialfundingOk() (*bool, bool)`

GetBDocumenttypeCommercialfundingOk returns a tuple with the BDocumenttypeCommercialfunding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBDocumenttypeCommercialfunding

`func (o *DocumenttypeRequest) SetBDocumenttypeCommercialfunding(v bool)`

SetBDocumenttypeCommercialfunding sets BDocumenttypeCommercialfunding field to given value.


### GetBDocumenttypeCanafe

`func (o *DocumenttypeRequest) GetBDocumenttypeCanafe() bool`

GetBDocumenttypeCanafe returns the BDocumenttypeCanafe field if non-nil, zero value otherwise.

### GetBDocumenttypeCanafeOk

`func (o *DocumenttypeRequest) GetBDocumenttypeCanafeOk() (*bool, bool)`

GetBDocumenttypeCanafeOk returns a tuple with the BDocumenttypeCanafe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBDocumenttypeCanafe

`func (o *DocumenttypeRequest) SetBDocumenttypeCanafe(v bool)`

SetBDocumenttypeCanafe sets BDocumenttypeCanafe field to given value.


### GetEDocumenttypeEzcom

`func (o *DocumenttypeRequest) GetEDocumenttypeEzcom() FieldEDocumenttypeEzcom`

GetEDocumenttypeEzcom returns the EDocumenttypeEzcom field if non-nil, zero value otherwise.

### GetEDocumenttypeEzcomOk

`func (o *DocumenttypeRequest) GetEDocumenttypeEzcomOk() (*FieldEDocumenttypeEzcom, bool)`

GetEDocumenttypeEzcomOk returns a tuple with the EDocumenttypeEzcom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEDocumenttypeEzcom

`func (o *DocumenttypeRequest) SetEDocumenttypeEzcom(v FieldEDocumenttypeEzcom)`

SetEDocumenttypeEzcom sets EDocumenttypeEzcom field to given value.


### GetEDocumenttypePrivacydefault

`func (o *DocumenttypeRequest) GetEDocumenttypePrivacydefault() FieldEDocumenttypePrivacydefault`

GetEDocumenttypePrivacydefault returns the EDocumenttypePrivacydefault field if non-nil, zero value otherwise.

### GetEDocumenttypePrivacydefaultOk

`func (o *DocumenttypeRequest) GetEDocumenttypePrivacydefaultOk() (*FieldEDocumenttypePrivacydefault, bool)`

GetEDocumenttypePrivacydefaultOk returns a tuple with the EDocumenttypePrivacydefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEDocumenttypePrivacydefault

`func (o *DocumenttypeRequest) SetEDocumenttypePrivacydefault(v FieldEDocumenttypePrivacydefault)`

SetEDocumenttypePrivacydefault sets EDocumenttypePrivacydefault field to given value.


### GetEDocumenttypeStatusdefault

`func (o *DocumenttypeRequest) GetEDocumenttypeStatusdefault() FieldEDocumenttypeStatusdefault`

GetEDocumenttypeStatusdefault returns the EDocumenttypeStatusdefault field if non-nil, zero value otherwise.

### GetEDocumenttypeStatusdefaultOk

`func (o *DocumenttypeRequest) GetEDocumenttypeStatusdefaultOk() (*FieldEDocumenttypeStatusdefault, bool)`

GetEDocumenttypeStatusdefaultOk returns a tuple with the EDocumenttypeStatusdefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEDocumenttypeStatusdefault

`func (o *DocumenttypeRequest) SetEDocumenttypeStatusdefault(v FieldEDocumenttypeStatusdefault)`

SetEDocumenttypeStatusdefault sets EDocumenttypeStatusdefault field to given value.


### GetEDocumenttypeSalebyowner

`func (o *DocumenttypeRequest) GetEDocumenttypeSalebyowner() FieldEDocumenttypeSalebyowner`

GetEDocumenttypeSalebyowner returns the EDocumenttypeSalebyowner field if non-nil, zero value otherwise.

### GetEDocumenttypeSalebyownerOk

`func (o *DocumenttypeRequest) GetEDocumenttypeSalebyownerOk() (*FieldEDocumenttypeSalebyowner, bool)`

GetEDocumenttypeSalebyownerOk returns a tuple with the EDocumenttypeSalebyowner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEDocumenttypeSalebyowner

`func (o *DocumenttypeRequest) SetEDocumenttypeSalebyowner(v FieldEDocumenttypeSalebyowner)`

SetEDocumenttypeSalebyowner sets EDocumenttypeSalebyowner field to given value.


### GetEDocumenttypeSaledisplaybyrole

`func (o *DocumenttypeRequest) GetEDocumenttypeSaledisplaybyrole() FieldEDocumenttypeSaledisplaybyrole`

GetEDocumenttypeSaledisplaybyrole returns the EDocumenttypeSaledisplaybyrole field if non-nil, zero value otherwise.

### GetEDocumenttypeSaledisplaybyroleOk

`func (o *DocumenttypeRequest) GetEDocumenttypeSaledisplaybyroleOk() (*FieldEDocumenttypeSaledisplaybyrole, bool)`

GetEDocumenttypeSaledisplaybyroleOk returns a tuple with the EDocumenttypeSaledisplaybyrole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEDocumenttypeSaledisplaybyrole

`func (o *DocumenttypeRequest) SetEDocumenttypeSaledisplaybyrole(v FieldEDocumenttypeSaledisplaybyrole)`

SetEDocumenttypeSaledisplaybyrole sets EDocumenttypeSaledisplaybyrole field to given value.


### GetEDocumenttypeTransactiontype

`func (o *DocumenttypeRequest) GetEDocumenttypeTransactiontype() FieldEDocumenttypeTransactiontype`

GetEDocumenttypeTransactiontype returns the EDocumenttypeTransactiontype field if non-nil, zero value otherwise.

### GetEDocumenttypeTransactiontypeOk

`func (o *DocumenttypeRequest) GetEDocumenttypeTransactiontypeOk() (*FieldEDocumenttypeTransactiontype, bool)`

GetEDocumenttypeTransactiontypeOk returns a tuple with the EDocumenttypeTransactiontype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEDocumenttypeTransactiontype

`func (o *DocumenttypeRequest) SetEDocumenttypeTransactiontype(v FieldEDocumenttypeTransactiontype)`

SetEDocumenttypeTransactiontype sets EDocumenttypeTransactiontype field to given value.


### GetEDocumenttypeRepossession

`func (o *DocumenttypeRequest) GetEDocumenttypeRepossession() FieldEDocumenttypeRepossession`

GetEDocumenttypeRepossession returns the EDocumenttypeRepossession field if non-nil, zero value otherwise.

### GetEDocumenttypeRepossessionOk

`func (o *DocumenttypeRequest) GetEDocumenttypeRepossessionOk() (*FieldEDocumenttypeRepossession, bool)`

GetEDocumenttypeRepossessionOk returns a tuple with the EDocumenttypeRepossession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEDocumenttypeRepossession

`func (o *DocumenttypeRequest) SetEDocumenttypeRepossession(v FieldEDocumenttypeRepossession)`

SetEDocumenttypeRepossession sets EDocumenttypeRepossession field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


