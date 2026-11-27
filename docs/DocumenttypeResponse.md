# DocumenttypeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiDocumenttypeID** | **int32** | The unique ID of the Documenttype | 
**ObjDocumenttypeName** | [**MultilingualDocumenttypeName**](MultilingualDocumenttypeName.md) |  | 
**EDocumenttypeStep** | [**ComputedEDocumenttypeStep**](ComputedEDocumenttypeStep.md) |  | 
**SDocumenttypeDefaultfilename** | **string** | The default filename for the document type | 
**BDocumenttypeShowagent** | **bool** | Whether the document type should be shown to the agent | 
**BDocumenttypeIsactive** | **bool** | Whether the documenttype is active or not | 
**BDocumenttypeSale** | Pointer to **bool** | Whether the document type applies to sale transactions | [optional] 
**BDocumenttypeRental** | Pointer to **bool** | Whether the document type applies to rental transactions | [optional] 
**BDocumenttypeReference** | Pointer to **bool** | Whether the document type is a reference document | [optional] 
**BDocumenttypeMortgagereference** | Pointer to **bool** | Whether the document type is a mortgage reference | [optional] 
**BDocumenttypeEvaluation** | Pointer to **bool** | Whether the document type applies to evaluation transactions | [optional] 
**BDocumenttypeConsultation** | Pointer to **bool** | Whether the document type applies to consultation transactions | [optional] 
**BDocumenttypeOtherservices** | Pointer to **bool** | Whether the document type applies to other services | [optional] 
**BDocumenttypeForeclosedsale** | Pointer to **bool** | Whether the document type applies to foreclosed sale transactions | [optional] 
**BDocumenttypeCommercialfunding** | Pointer to **bool** | Whether the document type applies to commercial funding transactions | [optional] 
**BDocumenttypeCanafe** | Pointer to **bool** | Whether the document type applies to Canafe transactions | [optional] 
**EDocumenttypeEzcom** | Pointer to [**FieldEDocumenttypeEzcom**](FieldEDocumenttypeEzcom.md) |  | [optional] 
**EDocumenttypePrivacydefault** | Pointer to [**FieldEDocumenttypePrivacydefault**](FieldEDocumenttypePrivacydefault.md) |  | [optional] 
**EDocumenttypeStatusdefault** | Pointer to [**FieldEDocumenttypeStatusdefault**](FieldEDocumenttypeStatusdefault.md) |  | [optional] 
**EDocumenttypeSalebyowner** | Pointer to [**FieldEDocumenttypeSalebyowner**](FieldEDocumenttypeSalebyowner.md) |  | [optional] 
**EDocumenttypeSaledisplaybyrole** | Pointer to [**FieldEDocumenttypeSaledisplaybyrole**](FieldEDocumenttypeSaledisplaybyrole.md) |  | [optional] 
**EDocumenttypeTransactiontype** | Pointer to [**FieldEDocumenttypeTransactiontype**](FieldEDocumenttypeTransactiontype.md) |  | [optional] 
**EDocumenttypeRepossession** | Pointer to [**FieldEDocumenttypeRepossession**](FieldEDocumenttypeRepossession.md) |  | [optional] 
**AFkiInscriptiontypeID** | **[]int32** |  | 

## Methods

### NewDocumenttypeResponse

`func NewDocumenttypeResponse(pkiDocumenttypeID int32, objDocumenttypeName MultilingualDocumenttypeName, eDocumenttypeStep ComputedEDocumenttypeStep, sDocumenttypeDefaultfilename string, bDocumenttypeShowagent bool, bDocumenttypeIsactive bool, aFkiInscriptiontypeID []int32, ) *DocumenttypeResponse`

NewDocumenttypeResponse instantiates a new DocumenttypeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumenttypeResponseWithDefaults

`func NewDocumenttypeResponseWithDefaults() *DocumenttypeResponse`

NewDocumenttypeResponseWithDefaults instantiates a new DocumenttypeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiDocumenttypeID

`func (o *DocumenttypeResponse) GetPkiDocumenttypeID() int32`

GetPkiDocumenttypeID returns the PkiDocumenttypeID field if non-nil, zero value otherwise.

### GetPkiDocumenttypeIDOk

`func (o *DocumenttypeResponse) GetPkiDocumenttypeIDOk() (*int32, bool)`

GetPkiDocumenttypeIDOk returns a tuple with the PkiDocumenttypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiDocumenttypeID

`func (o *DocumenttypeResponse) SetPkiDocumenttypeID(v int32)`

SetPkiDocumenttypeID sets PkiDocumenttypeID field to given value.


### GetObjDocumenttypeName

`func (o *DocumenttypeResponse) GetObjDocumenttypeName() MultilingualDocumenttypeName`

GetObjDocumenttypeName returns the ObjDocumenttypeName field if non-nil, zero value otherwise.

### GetObjDocumenttypeNameOk

`func (o *DocumenttypeResponse) GetObjDocumenttypeNameOk() (*MultilingualDocumenttypeName, bool)`

GetObjDocumenttypeNameOk returns a tuple with the ObjDocumenttypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjDocumenttypeName

`func (o *DocumenttypeResponse) SetObjDocumenttypeName(v MultilingualDocumenttypeName)`

SetObjDocumenttypeName sets ObjDocumenttypeName field to given value.


### GetEDocumenttypeStep

`func (o *DocumenttypeResponse) GetEDocumenttypeStep() ComputedEDocumenttypeStep`

GetEDocumenttypeStep returns the EDocumenttypeStep field if non-nil, zero value otherwise.

### GetEDocumenttypeStepOk

`func (o *DocumenttypeResponse) GetEDocumenttypeStepOk() (*ComputedEDocumenttypeStep, bool)`

GetEDocumenttypeStepOk returns a tuple with the EDocumenttypeStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEDocumenttypeStep

`func (o *DocumenttypeResponse) SetEDocumenttypeStep(v ComputedEDocumenttypeStep)`

SetEDocumenttypeStep sets EDocumenttypeStep field to given value.


### GetSDocumenttypeDefaultfilename

`func (o *DocumenttypeResponse) GetSDocumenttypeDefaultfilename() string`

GetSDocumenttypeDefaultfilename returns the SDocumenttypeDefaultfilename field if non-nil, zero value otherwise.

### GetSDocumenttypeDefaultfilenameOk

`func (o *DocumenttypeResponse) GetSDocumenttypeDefaultfilenameOk() (*string, bool)`

GetSDocumenttypeDefaultfilenameOk returns a tuple with the SDocumenttypeDefaultfilename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSDocumenttypeDefaultfilename

`func (o *DocumenttypeResponse) SetSDocumenttypeDefaultfilename(v string)`

SetSDocumenttypeDefaultfilename sets SDocumenttypeDefaultfilename field to given value.


### GetBDocumenttypeShowagent

`func (o *DocumenttypeResponse) GetBDocumenttypeShowagent() bool`

GetBDocumenttypeShowagent returns the BDocumenttypeShowagent field if non-nil, zero value otherwise.

### GetBDocumenttypeShowagentOk

`func (o *DocumenttypeResponse) GetBDocumenttypeShowagentOk() (*bool, bool)`

GetBDocumenttypeShowagentOk returns a tuple with the BDocumenttypeShowagent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBDocumenttypeShowagent

`func (o *DocumenttypeResponse) SetBDocumenttypeShowagent(v bool)`

SetBDocumenttypeShowagent sets BDocumenttypeShowagent field to given value.


### GetBDocumenttypeIsactive

`func (o *DocumenttypeResponse) GetBDocumenttypeIsactive() bool`

GetBDocumenttypeIsactive returns the BDocumenttypeIsactive field if non-nil, zero value otherwise.

### GetBDocumenttypeIsactiveOk

`func (o *DocumenttypeResponse) GetBDocumenttypeIsactiveOk() (*bool, bool)`

GetBDocumenttypeIsactiveOk returns a tuple with the BDocumenttypeIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBDocumenttypeIsactive

`func (o *DocumenttypeResponse) SetBDocumenttypeIsactive(v bool)`

SetBDocumenttypeIsactive sets BDocumenttypeIsactive field to given value.


### GetBDocumenttypeSale

`func (o *DocumenttypeResponse) GetBDocumenttypeSale() bool`

GetBDocumenttypeSale returns the BDocumenttypeSale field if non-nil, zero value otherwise.

### GetBDocumenttypeSaleOk

`func (o *DocumenttypeResponse) GetBDocumenttypeSaleOk() (*bool, bool)`

GetBDocumenttypeSaleOk returns a tuple with the BDocumenttypeSale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBDocumenttypeSale

`func (o *DocumenttypeResponse) SetBDocumenttypeSale(v bool)`

SetBDocumenttypeSale sets BDocumenttypeSale field to given value.

### HasBDocumenttypeSale

`func (o *DocumenttypeResponse) HasBDocumenttypeSale() bool`

HasBDocumenttypeSale returns a boolean if a field has been set.

### GetBDocumenttypeRental

`func (o *DocumenttypeResponse) GetBDocumenttypeRental() bool`

GetBDocumenttypeRental returns the BDocumenttypeRental field if non-nil, zero value otherwise.

### GetBDocumenttypeRentalOk

`func (o *DocumenttypeResponse) GetBDocumenttypeRentalOk() (*bool, bool)`

GetBDocumenttypeRentalOk returns a tuple with the BDocumenttypeRental field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBDocumenttypeRental

`func (o *DocumenttypeResponse) SetBDocumenttypeRental(v bool)`

SetBDocumenttypeRental sets BDocumenttypeRental field to given value.

### HasBDocumenttypeRental

`func (o *DocumenttypeResponse) HasBDocumenttypeRental() bool`

HasBDocumenttypeRental returns a boolean if a field has been set.

### GetBDocumenttypeReference

`func (o *DocumenttypeResponse) GetBDocumenttypeReference() bool`

GetBDocumenttypeReference returns the BDocumenttypeReference field if non-nil, zero value otherwise.

### GetBDocumenttypeReferenceOk

`func (o *DocumenttypeResponse) GetBDocumenttypeReferenceOk() (*bool, bool)`

GetBDocumenttypeReferenceOk returns a tuple with the BDocumenttypeReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBDocumenttypeReference

`func (o *DocumenttypeResponse) SetBDocumenttypeReference(v bool)`

SetBDocumenttypeReference sets BDocumenttypeReference field to given value.

### HasBDocumenttypeReference

`func (o *DocumenttypeResponse) HasBDocumenttypeReference() bool`

HasBDocumenttypeReference returns a boolean if a field has been set.

### GetBDocumenttypeMortgagereference

`func (o *DocumenttypeResponse) GetBDocumenttypeMortgagereference() bool`

GetBDocumenttypeMortgagereference returns the BDocumenttypeMortgagereference field if non-nil, zero value otherwise.

### GetBDocumenttypeMortgagereferenceOk

`func (o *DocumenttypeResponse) GetBDocumenttypeMortgagereferenceOk() (*bool, bool)`

GetBDocumenttypeMortgagereferenceOk returns a tuple with the BDocumenttypeMortgagereference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBDocumenttypeMortgagereference

`func (o *DocumenttypeResponse) SetBDocumenttypeMortgagereference(v bool)`

SetBDocumenttypeMortgagereference sets BDocumenttypeMortgagereference field to given value.

### HasBDocumenttypeMortgagereference

`func (o *DocumenttypeResponse) HasBDocumenttypeMortgagereference() bool`

HasBDocumenttypeMortgagereference returns a boolean if a field has been set.

### GetBDocumenttypeEvaluation

`func (o *DocumenttypeResponse) GetBDocumenttypeEvaluation() bool`

GetBDocumenttypeEvaluation returns the BDocumenttypeEvaluation field if non-nil, zero value otherwise.

### GetBDocumenttypeEvaluationOk

`func (o *DocumenttypeResponse) GetBDocumenttypeEvaluationOk() (*bool, bool)`

GetBDocumenttypeEvaluationOk returns a tuple with the BDocumenttypeEvaluation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBDocumenttypeEvaluation

`func (o *DocumenttypeResponse) SetBDocumenttypeEvaluation(v bool)`

SetBDocumenttypeEvaluation sets BDocumenttypeEvaluation field to given value.

### HasBDocumenttypeEvaluation

`func (o *DocumenttypeResponse) HasBDocumenttypeEvaluation() bool`

HasBDocumenttypeEvaluation returns a boolean if a field has been set.

### GetBDocumenttypeConsultation

`func (o *DocumenttypeResponse) GetBDocumenttypeConsultation() bool`

GetBDocumenttypeConsultation returns the BDocumenttypeConsultation field if non-nil, zero value otherwise.

### GetBDocumenttypeConsultationOk

`func (o *DocumenttypeResponse) GetBDocumenttypeConsultationOk() (*bool, bool)`

GetBDocumenttypeConsultationOk returns a tuple with the BDocumenttypeConsultation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBDocumenttypeConsultation

`func (o *DocumenttypeResponse) SetBDocumenttypeConsultation(v bool)`

SetBDocumenttypeConsultation sets BDocumenttypeConsultation field to given value.

### HasBDocumenttypeConsultation

`func (o *DocumenttypeResponse) HasBDocumenttypeConsultation() bool`

HasBDocumenttypeConsultation returns a boolean if a field has been set.

### GetBDocumenttypeOtherservices

`func (o *DocumenttypeResponse) GetBDocumenttypeOtherservices() bool`

GetBDocumenttypeOtherservices returns the BDocumenttypeOtherservices field if non-nil, zero value otherwise.

### GetBDocumenttypeOtherservicesOk

`func (o *DocumenttypeResponse) GetBDocumenttypeOtherservicesOk() (*bool, bool)`

GetBDocumenttypeOtherservicesOk returns a tuple with the BDocumenttypeOtherservices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBDocumenttypeOtherservices

`func (o *DocumenttypeResponse) SetBDocumenttypeOtherservices(v bool)`

SetBDocumenttypeOtherservices sets BDocumenttypeOtherservices field to given value.

### HasBDocumenttypeOtherservices

`func (o *DocumenttypeResponse) HasBDocumenttypeOtherservices() bool`

HasBDocumenttypeOtherservices returns a boolean if a field has been set.

### GetBDocumenttypeForeclosedsale

`func (o *DocumenttypeResponse) GetBDocumenttypeForeclosedsale() bool`

GetBDocumenttypeForeclosedsale returns the BDocumenttypeForeclosedsale field if non-nil, zero value otherwise.

### GetBDocumenttypeForeclosedsaleOk

`func (o *DocumenttypeResponse) GetBDocumenttypeForeclosedsaleOk() (*bool, bool)`

GetBDocumenttypeForeclosedsaleOk returns a tuple with the BDocumenttypeForeclosedsale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBDocumenttypeForeclosedsale

`func (o *DocumenttypeResponse) SetBDocumenttypeForeclosedsale(v bool)`

SetBDocumenttypeForeclosedsale sets BDocumenttypeForeclosedsale field to given value.

### HasBDocumenttypeForeclosedsale

`func (o *DocumenttypeResponse) HasBDocumenttypeForeclosedsale() bool`

HasBDocumenttypeForeclosedsale returns a boolean if a field has been set.

### GetBDocumenttypeCommercialfunding

`func (o *DocumenttypeResponse) GetBDocumenttypeCommercialfunding() bool`

GetBDocumenttypeCommercialfunding returns the BDocumenttypeCommercialfunding field if non-nil, zero value otherwise.

### GetBDocumenttypeCommercialfundingOk

`func (o *DocumenttypeResponse) GetBDocumenttypeCommercialfundingOk() (*bool, bool)`

GetBDocumenttypeCommercialfundingOk returns a tuple with the BDocumenttypeCommercialfunding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBDocumenttypeCommercialfunding

`func (o *DocumenttypeResponse) SetBDocumenttypeCommercialfunding(v bool)`

SetBDocumenttypeCommercialfunding sets BDocumenttypeCommercialfunding field to given value.

### HasBDocumenttypeCommercialfunding

`func (o *DocumenttypeResponse) HasBDocumenttypeCommercialfunding() bool`

HasBDocumenttypeCommercialfunding returns a boolean if a field has been set.

### GetBDocumenttypeCanafe

`func (o *DocumenttypeResponse) GetBDocumenttypeCanafe() bool`

GetBDocumenttypeCanafe returns the BDocumenttypeCanafe field if non-nil, zero value otherwise.

### GetBDocumenttypeCanafeOk

`func (o *DocumenttypeResponse) GetBDocumenttypeCanafeOk() (*bool, bool)`

GetBDocumenttypeCanafeOk returns a tuple with the BDocumenttypeCanafe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBDocumenttypeCanafe

`func (o *DocumenttypeResponse) SetBDocumenttypeCanafe(v bool)`

SetBDocumenttypeCanafe sets BDocumenttypeCanafe field to given value.

### HasBDocumenttypeCanafe

`func (o *DocumenttypeResponse) HasBDocumenttypeCanafe() bool`

HasBDocumenttypeCanafe returns a boolean if a field has been set.

### GetEDocumenttypeEzcom

`func (o *DocumenttypeResponse) GetEDocumenttypeEzcom() FieldEDocumenttypeEzcom`

GetEDocumenttypeEzcom returns the EDocumenttypeEzcom field if non-nil, zero value otherwise.

### GetEDocumenttypeEzcomOk

`func (o *DocumenttypeResponse) GetEDocumenttypeEzcomOk() (*FieldEDocumenttypeEzcom, bool)`

GetEDocumenttypeEzcomOk returns a tuple with the EDocumenttypeEzcom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEDocumenttypeEzcom

`func (o *DocumenttypeResponse) SetEDocumenttypeEzcom(v FieldEDocumenttypeEzcom)`

SetEDocumenttypeEzcom sets EDocumenttypeEzcom field to given value.

### HasEDocumenttypeEzcom

`func (o *DocumenttypeResponse) HasEDocumenttypeEzcom() bool`

HasEDocumenttypeEzcom returns a boolean if a field has been set.

### GetEDocumenttypePrivacydefault

`func (o *DocumenttypeResponse) GetEDocumenttypePrivacydefault() FieldEDocumenttypePrivacydefault`

GetEDocumenttypePrivacydefault returns the EDocumenttypePrivacydefault field if non-nil, zero value otherwise.

### GetEDocumenttypePrivacydefaultOk

`func (o *DocumenttypeResponse) GetEDocumenttypePrivacydefaultOk() (*FieldEDocumenttypePrivacydefault, bool)`

GetEDocumenttypePrivacydefaultOk returns a tuple with the EDocumenttypePrivacydefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEDocumenttypePrivacydefault

`func (o *DocumenttypeResponse) SetEDocumenttypePrivacydefault(v FieldEDocumenttypePrivacydefault)`

SetEDocumenttypePrivacydefault sets EDocumenttypePrivacydefault field to given value.

### HasEDocumenttypePrivacydefault

`func (o *DocumenttypeResponse) HasEDocumenttypePrivacydefault() bool`

HasEDocumenttypePrivacydefault returns a boolean if a field has been set.

### GetEDocumenttypeStatusdefault

`func (o *DocumenttypeResponse) GetEDocumenttypeStatusdefault() FieldEDocumenttypeStatusdefault`

GetEDocumenttypeStatusdefault returns the EDocumenttypeStatusdefault field if non-nil, zero value otherwise.

### GetEDocumenttypeStatusdefaultOk

`func (o *DocumenttypeResponse) GetEDocumenttypeStatusdefaultOk() (*FieldEDocumenttypeStatusdefault, bool)`

GetEDocumenttypeStatusdefaultOk returns a tuple with the EDocumenttypeStatusdefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEDocumenttypeStatusdefault

`func (o *DocumenttypeResponse) SetEDocumenttypeStatusdefault(v FieldEDocumenttypeStatusdefault)`

SetEDocumenttypeStatusdefault sets EDocumenttypeStatusdefault field to given value.

### HasEDocumenttypeStatusdefault

`func (o *DocumenttypeResponse) HasEDocumenttypeStatusdefault() bool`

HasEDocumenttypeStatusdefault returns a boolean if a field has been set.

### GetEDocumenttypeSalebyowner

`func (o *DocumenttypeResponse) GetEDocumenttypeSalebyowner() FieldEDocumenttypeSalebyowner`

GetEDocumenttypeSalebyowner returns the EDocumenttypeSalebyowner field if non-nil, zero value otherwise.

### GetEDocumenttypeSalebyownerOk

`func (o *DocumenttypeResponse) GetEDocumenttypeSalebyownerOk() (*FieldEDocumenttypeSalebyowner, bool)`

GetEDocumenttypeSalebyownerOk returns a tuple with the EDocumenttypeSalebyowner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEDocumenttypeSalebyowner

`func (o *DocumenttypeResponse) SetEDocumenttypeSalebyowner(v FieldEDocumenttypeSalebyowner)`

SetEDocumenttypeSalebyowner sets EDocumenttypeSalebyowner field to given value.

### HasEDocumenttypeSalebyowner

`func (o *DocumenttypeResponse) HasEDocumenttypeSalebyowner() bool`

HasEDocumenttypeSalebyowner returns a boolean if a field has been set.

### GetEDocumenttypeSaledisplaybyrole

`func (o *DocumenttypeResponse) GetEDocumenttypeSaledisplaybyrole() FieldEDocumenttypeSaledisplaybyrole`

GetEDocumenttypeSaledisplaybyrole returns the EDocumenttypeSaledisplaybyrole field if non-nil, zero value otherwise.

### GetEDocumenttypeSaledisplaybyroleOk

`func (o *DocumenttypeResponse) GetEDocumenttypeSaledisplaybyroleOk() (*FieldEDocumenttypeSaledisplaybyrole, bool)`

GetEDocumenttypeSaledisplaybyroleOk returns a tuple with the EDocumenttypeSaledisplaybyrole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEDocumenttypeSaledisplaybyrole

`func (o *DocumenttypeResponse) SetEDocumenttypeSaledisplaybyrole(v FieldEDocumenttypeSaledisplaybyrole)`

SetEDocumenttypeSaledisplaybyrole sets EDocumenttypeSaledisplaybyrole field to given value.

### HasEDocumenttypeSaledisplaybyrole

`func (o *DocumenttypeResponse) HasEDocumenttypeSaledisplaybyrole() bool`

HasEDocumenttypeSaledisplaybyrole returns a boolean if a field has been set.

### GetEDocumenttypeTransactiontype

`func (o *DocumenttypeResponse) GetEDocumenttypeTransactiontype() FieldEDocumenttypeTransactiontype`

GetEDocumenttypeTransactiontype returns the EDocumenttypeTransactiontype field if non-nil, zero value otherwise.

### GetEDocumenttypeTransactiontypeOk

`func (o *DocumenttypeResponse) GetEDocumenttypeTransactiontypeOk() (*FieldEDocumenttypeTransactiontype, bool)`

GetEDocumenttypeTransactiontypeOk returns a tuple with the EDocumenttypeTransactiontype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEDocumenttypeTransactiontype

`func (o *DocumenttypeResponse) SetEDocumenttypeTransactiontype(v FieldEDocumenttypeTransactiontype)`

SetEDocumenttypeTransactiontype sets EDocumenttypeTransactiontype field to given value.

### HasEDocumenttypeTransactiontype

`func (o *DocumenttypeResponse) HasEDocumenttypeTransactiontype() bool`

HasEDocumenttypeTransactiontype returns a boolean if a field has been set.

### GetEDocumenttypeRepossession

`func (o *DocumenttypeResponse) GetEDocumenttypeRepossession() FieldEDocumenttypeRepossession`

GetEDocumenttypeRepossession returns the EDocumenttypeRepossession field if non-nil, zero value otherwise.

### GetEDocumenttypeRepossessionOk

`func (o *DocumenttypeResponse) GetEDocumenttypeRepossessionOk() (*FieldEDocumenttypeRepossession, bool)`

GetEDocumenttypeRepossessionOk returns a tuple with the EDocumenttypeRepossession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEDocumenttypeRepossession

`func (o *DocumenttypeResponse) SetEDocumenttypeRepossession(v FieldEDocumenttypeRepossession)`

SetEDocumenttypeRepossession sets EDocumenttypeRepossession field to given value.

### HasEDocumenttypeRepossession

`func (o *DocumenttypeResponse) HasEDocumenttypeRepossession() bool`

HasEDocumenttypeRepossession returns a boolean if a field has been set.

### GetAFkiInscriptiontypeID

`func (o *DocumenttypeResponse) GetAFkiInscriptiontypeID() []int32`

GetAFkiInscriptiontypeID returns the AFkiInscriptiontypeID field if non-nil, zero value otherwise.

### GetAFkiInscriptiontypeIDOk

`func (o *DocumenttypeResponse) GetAFkiInscriptiontypeIDOk() (*[]int32, bool)`

GetAFkiInscriptiontypeIDOk returns a tuple with the AFkiInscriptiontypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAFkiInscriptiontypeID

`func (o *DocumenttypeResponse) SetAFkiInscriptiontypeID(v []int32)`

SetAFkiInscriptiontypeID sets AFkiInscriptiontypeID field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


