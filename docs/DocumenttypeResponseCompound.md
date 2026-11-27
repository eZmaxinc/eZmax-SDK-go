# DocumenttypeResponseCompound

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

### NewDocumenttypeResponseCompound

`func NewDocumenttypeResponseCompound(pkiDocumenttypeID int32, objDocumenttypeName MultilingualDocumenttypeName, eDocumenttypeStep ComputedEDocumenttypeStep, sDocumenttypeDefaultfilename string, bDocumenttypeShowagent bool, bDocumenttypeIsactive bool, aFkiInscriptiontypeID []int32, ) *DocumenttypeResponseCompound`

NewDocumenttypeResponseCompound instantiates a new DocumenttypeResponseCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumenttypeResponseCompoundWithDefaults

`func NewDocumenttypeResponseCompoundWithDefaults() *DocumenttypeResponseCompound`

NewDocumenttypeResponseCompoundWithDefaults instantiates a new DocumenttypeResponseCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiDocumenttypeID

`func (o *DocumenttypeResponseCompound) GetPkiDocumenttypeID() int32`

GetPkiDocumenttypeID returns the PkiDocumenttypeID field if non-nil, zero value otherwise.

### GetPkiDocumenttypeIDOk

`func (o *DocumenttypeResponseCompound) GetPkiDocumenttypeIDOk() (*int32, bool)`

GetPkiDocumenttypeIDOk returns a tuple with the PkiDocumenttypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiDocumenttypeID

`func (o *DocumenttypeResponseCompound) SetPkiDocumenttypeID(v int32)`

SetPkiDocumenttypeID sets PkiDocumenttypeID field to given value.


### GetObjDocumenttypeName

`func (o *DocumenttypeResponseCompound) GetObjDocumenttypeName() MultilingualDocumenttypeName`

GetObjDocumenttypeName returns the ObjDocumenttypeName field if non-nil, zero value otherwise.

### GetObjDocumenttypeNameOk

`func (o *DocumenttypeResponseCompound) GetObjDocumenttypeNameOk() (*MultilingualDocumenttypeName, bool)`

GetObjDocumenttypeNameOk returns a tuple with the ObjDocumenttypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjDocumenttypeName

`func (o *DocumenttypeResponseCompound) SetObjDocumenttypeName(v MultilingualDocumenttypeName)`

SetObjDocumenttypeName sets ObjDocumenttypeName field to given value.


### GetEDocumenttypeStep

`func (o *DocumenttypeResponseCompound) GetEDocumenttypeStep() ComputedEDocumenttypeStep`

GetEDocumenttypeStep returns the EDocumenttypeStep field if non-nil, zero value otherwise.

### GetEDocumenttypeStepOk

`func (o *DocumenttypeResponseCompound) GetEDocumenttypeStepOk() (*ComputedEDocumenttypeStep, bool)`

GetEDocumenttypeStepOk returns a tuple with the EDocumenttypeStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEDocumenttypeStep

`func (o *DocumenttypeResponseCompound) SetEDocumenttypeStep(v ComputedEDocumenttypeStep)`

SetEDocumenttypeStep sets EDocumenttypeStep field to given value.


### GetSDocumenttypeDefaultfilename

`func (o *DocumenttypeResponseCompound) GetSDocumenttypeDefaultfilename() string`

GetSDocumenttypeDefaultfilename returns the SDocumenttypeDefaultfilename field if non-nil, zero value otherwise.

### GetSDocumenttypeDefaultfilenameOk

`func (o *DocumenttypeResponseCompound) GetSDocumenttypeDefaultfilenameOk() (*string, bool)`

GetSDocumenttypeDefaultfilenameOk returns a tuple with the SDocumenttypeDefaultfilename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSDocumenttypeDefaultfilename

`func (o *DocumenttypeResponseCompound) SetSDocumenttypeDefaultfilename(v string)`

SetSDocumenttypeDefaultfilename sets SDocumenttypeDefaultfilename field to given value.


### GetBDocumenttypeShowagent

`func (o *DocumenttypeResponseCompound) GetBDocumenttypeShowagent() bool`

GetBDocumenttypeShowagent returns the BDocumenttypeShowagent field if non-nil, zero value otherwise.

### GetBDocumenttypeShowagentOk

`func (o *DocumenttypeResponseCompound) GetBDocumenttypeShowagentOk() (*bool, bool)`

GetBDocumenttypeShowagentOk returns a tuple with the BDocumenttypeShowagent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBDocumenttypeShowagent

`func (o *DocumenttypeResponseCompound) SetBDocumenttypeShowagent(v bool)`

SetBDocumenttypeShowagent sets BDocumenttypeShowagent field to given value.


### GetBDocumenttypeIsactive

`func (o *DocumenttypeResponseCompound) GetBDocumenttypeIsactive() bool`

GetBDocumenttypeIsactive returns the BDocumenttypeIsactive field if non-nil, zero value otherwise.

### GetBDocumenttypeIsactiveOk

`func (o *DocumenttypeResponseCompound) GetBDocumenttypeIsactiveOk() (*bool, bool)`

GetBDocumenttypeIsactiveOk returns a tuple with the BDocumenttypeIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBDocumenttypeIsactive

`func (o *DocumenttypeResponseCompound) SetBDocumenttypeIsactive(v bool)`

SetBDocumenttypeIsactive sets BDocumenttypeIsactive field to given value.


### GetBDocumenttypeSale

`func (o *DocumenttypeResponseCompound) GetBDocumenttypeSale() bool`

GetBDocumenttypeSale returns the BDocumenttypeSale field if non-nil, zero value otherwise.

### GetBDocumenttypeSaleOk

`func (o *DocumenttypeResponseCompound) GetBDocumenttypeSaleOk() (*bool, bool)`

GetBDocumenttypeSaleOk returns a tuple with the BDocumenttypeSale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBDocumenttypeSale

`func (o *DocumenttypeResponseCompound) SetBDocumenttypeSale(v bool)`

SetBDocumenttypeSale sets BDocumenttypeSale field to given value.

### HasBDocumenttypeSale

`func (o *DocumenttypeResponseCompound) HasBDocumenttypeSale() bool`

HasBDocumenttypeSale returns a boolean if a field has been set.

### GetBDocumenttypeRental

`func (o *DocumenttypeResponseCompound) GetBDocumenttypeRental() bool`

GetBDocumenttypeRental returns the BDocumenttypeRental field if non-nil, zero value otherwise.

### GetBDocumenttypeRentalOk

`func (o *DocumenttypeResponseCompound) GetBDocumenttypeRentalOk() (*bool, bool)`

GetBDocumenttypeRentalOk returns a tuple with the BDocumenttypeRental field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBDocumenttypeRental

`func (o *DocumenttypeResponseCompound) SetBDocumenttypeRental(v bool)`

SetBDocumenttypeRental sets BDocumenttypeRental field to given value.

### HasBDocumenttypeRental

`func (o *DocumenttypeResponseCompound) HasBDocumenttypeRental() bool`

HasBDocumenttypeRental returns a boolean if a field has been set.

### GetBDocumenttypeReference

`func (o *DocumenttypeResponseCompound) GetBDocumenttypeReference() bool`

GetBDocumenttypeReference returns the BDocumenttypeReference field if non-nil, zero value otherwise.

### GetBDocumenttypeReferenceOk

`func (o *DocumenttypeResponseCompound) GetBDocumenttypeReferenceOk() (*bool, bool)`

GetBDocumenttypeReferenceOk returns a tuple with the BDocumenttypeReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBDocumenttypeReference

`func (o *DocumenttypeResponseCompound) SetBDocumenttypeReference(v bool)`

SetBDocumenttypeReference sets BDocumenttypeReference field to given value.

### HasBDocumenttypeReference

`func (o *DocumenttypeResponseCompound) HasBDocumenttypeReference() bool`

HasBDocumenttypeReference returns a boolean if a field has been set.

### GetBDocumenttypeMortgagereference

`func (o *DocumenttypeResponseCompound) GetBDocumenttypeMortgagereference() bool`

GetBDocumenttypeMortgagereference returns the BDocumenttypeMortgagereference field if non-nil, zero value otherwise.

### GetBDocumenttypeMortgagereferenceOk

`func (o *DocumenttypeResponseCompound) GetBDocumenttypeMortgagereferenceOk() (*bool, bool)`

GetBDocumenttypeMortgagereferenceOk returns a tuple with the BDocumenttypeMortgagereference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBDocumenttypeMortgagereference

`func (o *DocumenttypeResponseCompound) SetBDocumenttypeMortgagereference(v bool)`

SetBDocumenttypeMortgagereference sets BDocumenttypeMortgagereference field to given value.

### HasBDocumenttypeMortgagereference

`func (o *DocumenttypeResponseCompound) HasBDocumenttypeMortgagereference() bool`

HasBDocumenttypeMortgagereference returns a boolean if a field has been set.

### GetBDocumenttypeEvaluation

`func (o *DocumenttypeResponseCompound) GetBDocumenttypeEvaluation() bool`

GetBDocumenttypeEvaluation returns the BDocumenttypeEvaluation field if non-nil, zero value otherwise.

### GetBDocumenttypeEvaluationOk

`func (o *DocumenttypeResponseCompound) GetBDocumenttypeEvaluationOk() (*bool, bool)`

GetBDocumenttypeEvaluationOk returns a tuple with the BDocumenttypeEvaluation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBDocumenttypeEvaluation

`func (o *DocumenttypeResponseCompound) SetBDocumenttypeEvaluation(v bool)`

SetBDocumenttypeEvaluation sets BDocumenttypeEvaluation field to given value.

### HasBDocumenttypeEvaluation

`func (o *DocumenttypeResponseCompound) HasBDocumenttypeEvaluation() bool`

HasBDocumenttypeEvaluation returns a boolean if a field has been set.

### GetBDocumenttypeConsultation

`func (o *DocumenttypeResponseCompound) GetBDocumenttypeConsultation() bool`

GetBDocumenttypeConsultation returns the BDocumenttypeConsultation field if non-nil, zero value otherwise.

### GetBDocumenttypeConsultationOk

`func (o *DocumenttypeResponseCompound) GetBDocumenttypeConsultationOk() (*bool, bool)`

GetBDocumenttypeConsultationOk returns a tuple with the BDocumenttypeConsultation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBDocumenttypeConsultation

`func (o *DocumenttypeResponseCompound) SetBDocumenttypeConsultation(v bool)`

SetBDocumenttypeConsultation sets BDocumenttypeConsultation field to given value.

### HasBDocumenttypeConsultation

`func (o *DocumenttypeResponseCompound) HasBDocumenttypeConsultation() bool`

HasBDocumenttypeConsultation returns a boolean if a field has been set.

### GetBDocumenttypeOtherservices

`func (o *DocumenttypeResponseCompound) GetBDocumenttypeOtherservices() bool`

GetBDocumenttypeOtherservices returns the BDocumenttypeOtherservices field if non-nil, zero value otherwise.

### GetBDocumenttypeOtherservicesOk

`func (o *DocumenttypeResponseCompound) GetBDocumenttypeOtherservicesOk() (*bool, bool)`

GetBDocumenttypeOtherservicesOk returns a tuple with the BDocumenttypeOtherservices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBDocumenttypeOtherservices

`func (o *DocumenttypeResponseCompound) SetBDocumenttypeOtherservices(v bool)`

SetBDocumenttypeOtherservices sets BDocumenttypeOtherservices field to given value.

### HasBDocumenttypeOtherservices

`func (o *DocumenttypeResponseCompound) HasBDocumenttypeOtherservices() bool`

HasBDocumenttypeOtherservices returns a boolean if a field has been set.

### GetBDocumenttypeForeclosedsale

`func (o *DocumenttypeResponseCompound) GetBDocumenttypeForeclosedsale() bool`

GetBDocumenttypeForeclosedsale returns the BDocumenttypeForeclosedsale field if non-nil, zero value otherwise.

### GetBDocumenttypeForeclosedsaleOk

`func (o *DocumenttypeResponseCompound) GetBDocumenttypeForeclosedsaleOk() (*bool, bool)`

GetBDocumenttypeForeclosedsaleOk returns a tuple with the BDocumenttypeForeclosedsale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBDocumenttypeForeclosedsale

`func (o *DocumenttypeResponseCompound) SetBDocumenttypeForeclosedsale(v bool)`

SetBDocumenttypeForeclosedsale sets BDocumenttypeForeclosedsale field to given value.

### HasBDocumenttypeForeclosedsale

`func (o *DocumenttypeResponseCompound) HasBDocumenttypeForeclosedsale() bool`

HasBDocumenttypeForeclosedsale returns a boolean if a field has been set.

### GetBDocumenttypeCommercialfunding

`func (o *DocumenttypeResponseCompound) GetBDocumenttypeCommercialfunding() bool`

GetBDocumenttypeCommercialfunding returns the BDocumenttypeCommercialfunding field if non-nil, zero value otherwise.

### GetBDocumenttypeCommercialfundingOk

`func (o *DocumenttypeResponseCompound) GetBDocumenttypeCommercialfundingOk() (*bool, bool)`

GetBDocumenttypeCommercialfundingOk returns a tuple with the BDocumenttypeCommercialfunding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBDocumenttypeCommercialfunding

`func (o *DocumenttypeResponseCompound) SetBDocumenttypeCommercialfunding(v bool)`

SetBDocumenttypeCommercialfunding sets BDocumenttypeCommercialfunding field to given value.

### HasBDocumenttypeCommercialfunding

`func (o *DocumenttypeResponseCompound) HasBDocumenttypeCommercialfunding() bool`

HasBDocumenttypeCommercialfunding returns a boolean if a field has been set.

### GetBDocumenttypeCanafe

`func (o *DocumenttypeResponseCompound) GetBDocumenttypeCanafe() bool`

GetBDocumenttypeCanafe returns the BDocumenttypeCanafe field if non-nil, zero value otherwise.

### GetBDocumenttypeCanafeOk

`func (o *DocumenttypeResponseCompound) GetBDocumenttypeCanafeOk() (*bool, bool)`

GetBDocumenttypeCanafeOk returns a tuple with the BDocumenttypeCanafe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBDocumenttypeCanafe

`func (o *DocumenttypeResponseCompound) SetBDocumenttypeCanafe(v bool)`

SetBDocumenttypeCanafe sets BDocumenttypeCanafe field to given value.

### HasBDocumenttypeCanafe

`func (o *DocumenttypeResponseCompound) HasBDocumenttypeCanafe() bool`

HasBDocumenttypeCanafe returns a boolean if a field has been set.

### GetEDocumenttypeEzcom

`func (o *DocumenttypeResponseCompound) GetEDocumenttypeEzcom() FieldEDocumenttypeEzcom`

GetEDocumenttypeEzcom returns the EDocumenttypeEzcom field if non-nil, zero value otherwise.

### GetEDocumenttypeEzcomOk

`func (o *DocumenttypeResponseCompound) GetEDocumenttypeEzcomOk() (*FieldEDocumenttypeEzcom, bool)`

GetEDocumenttypeEzcomOk returns a tuple with the EDocumenttypeEzcom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEDocumenttypeEzcom

`func (o *DocumenttypeResponseCompound) SetEDocumenttypeEzcom(v FieldEDocumenttypeEzcom)`

SetEDocumenttypeEzcom sets EDocumenttypeEzcom field to given value.

### HasEDocumenttypeEzcom

`func (o *DocumenttypeResponseCompound) HasEDocumenttypeEzcom() bool`

HasEDocumenttypeEzcom returns a boolean if a field has been set.

### GetEDocumenttypePrivacydefault

`func (o *DocumenttypeResponseCompound) GetEDocumenttypePrivacydefault() FieldEDocumenttypePrivacydefault`

GetEDocumenttypePrivacydefault returns the EDocumenttypePrivacydefault field if non-nil, zero value otherwise.

### GetEDocumenttypePrivacydefaultOk

`func (o *DocumenttypeResponseCompound) GetEDocumenttypePrivacydefaultOk() (*FieldEDocumenttypePrivacydefault, bool)`

GetEDocumenttypePrivacydefaultOk returns a tuple with the EDocumenttypePrivacydefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEDocumenttypePrivacydefault

`func (o *DocumenttypeResponseCompound) SetEDocumenttypePrivacydefault(v FieldEDocumenttypePrivacydefault)`

SetEDocumenttypePrivacydefault sets EDocumenttypePrivacydefault field to given value.

### HasEDocumenttypePrivacydefault

`func (o *DocumenttypeResponseCompound) HasEDocumenttypePrivacydefault() bool`

HasEDocumenttypePrivacydefault returns a boolean if a field has been set.

### GetEDocumenttypeStatusdefault

`func (o *DocumenttypeResponseCompound) GetEDocumenttypeStatusdefault() FieldEDocumenttypeStatusdefault`

GetEDocumenttypeStatusdefault returns the EDocumenttypeStatusdefault field if non-nil, zero value otherwise.

### GetEDocumenttypeStatusdefaultOk

`func (o *DocumenttypeResponseCompound) GetEDocumenttypeStatusdefaultOk() (*FieldEDocumenttypeStatusdefault, bool)`

GetEDocumenttypeStatusdefaultOk returns a tuple with the EDocumenttypeStatusdefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEDocumenttypeStatusdefault

`func (o *DocumenttypeResponseCompound) SetEDocumenttypeStatusdefault(v FieldEDocumenttypeStatusdefault)`

SetEDocumenttypeStatusdefault sets EDocumenttypeStatusdefault field to given value.

### HasEDocumenttypeStatusdefault

`func (o *DocumenttypeResponseCompound) HasEDocumenttypeStatusdefault() bool`

HasEDocumenttypeStatusdefault returns a boolean if a field has been set.

### GetEDocumenttypeSalebyowner

`func (o *DocumenttypeResponseCompound) GetEDocumenttypeSalebyowner() FieldEDocumenttypeSalebyowner`

GetEDocumenttypeSalebyowner returns the EDocumenttypeSalebyowner field if non-nil, zero value otherwise.

### GetEDocumenttypeSalebyownerOk

`func (o *DocumenttypeResponseCompound) GetEDocumenttypeSalebyownerOk() (*FieldEDocumenttypeSalebyowner, bool)`

GetEDocumenttypeSalebyownerOk returns a tuple with the EDocumenttypeSalebyowner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEDocumenttypeSalebyowner

`func (o *DocumenttypeResponseCompound) SetEDocumenttypeSalebyowner(v FieldEDocumenttypeSalebyowner)`

SetEDocumenttypeSalebyowner sets EDocumenttypeSalebyowner field to given value.

### HasEDocumenttypeSalebyowner

`func (o *DocumenttypeResponseCompound) HasEDocumenttypeSalebyowner() bool`

HasEDocumenttypeSalebyowner returns a boolean if a field has been set.

### GetEDocumenttypeSaledisplaybyrole

`func (o *DocumenttypeResponseCompound) GetEDocumenttypeSaledisplaybyrole() FieldEDocumenttypeSaledisplaybyrole`

GetEDocumenttypeSaledisplaybyrole returns the EDocumenttypeSaledisplaybyrole field if non-nil, zero value otherwise.

### GetEDocumenttypeSaledisplaybyroleOk

`func (o *DocumenttypeResponseCompound) GetEDocumenttypeSaledisplaybyroleOk() (*FieldEDocumenttypeSaledisplaybyrole, bool)`

GetEDocumenttypeSaledisplaybyroleOk returns a tuple with the EDocumenttypeSaledisplaybyrole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEDocumenttypeSaledisplaybyrole

`func (o *DocumenttypeResponseCompound) SetEDocumenttypeSaledisplaybyrole(v FieldEDocumenttypeSaledisplaybyrole)`

SetEDocumenttypeSaledisplaybyrole sets EDocumenttypeSaledisplaybyrole field to given value.

### HasEDocumenttypeSaledisplaybyrole

`func (o *DocumenttypeResponseCompound) HasEDocumenttypeSaledisplaybyrole() bool`

HasEDocumenttypeSaledisplaybyrole returns a boolean if a field has been set.

### GetEDocumenttypeTransactiontype

`func (o *DocumenttypeResponseCompound) GetEDocumenttypeTransactiontype() FieldEDocumenttypeTransactiontype`

GetEDocumenttypeTransactiontype returns the EDocumenttypeTransactiontype field if non-nil, zero value otherwise.

### GetEDocumenttypeTransactiontypeOk

`func (o *DocumenttypeResponseCompound) GetEDocumenttypeTransactiontypeOk() (*FieldEDocumenttypeTransactiontype, bool)`

GetEDocumenttypeTransactiontypeOk returns a tuple with the EDocumenttypeTransactiontype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEDocumenttypeTransactiontype

`func (o *DocumenttypeResponseCompound) SetEDocumenttypeTransactiontype(v FieldEDocumenttypeTransactiontype)`

SetEDocumenttypeTransactiontype sets EDocumenttypeTransactiontype field to given value.

### HasEDocumenttypeTransactiontype

`func (o *DocumenttypeResponseCompound) HasEDocumenttypeTransactiontype() bool`

HasEDocumenttypeTransactiontype returns a boolean if a field has been set.

### GetEDocumenttypeRepossession

`func (o *DocumenttypeResponseCompound) GetEDocumenttypeRepossession() FieldEDocumenttypeRepossession`

GetEDocumenttypeRepossession returns the EDocumenttypeRepossession field if non-nil, zero value otherwise.

### GetEDocumenttypeRepossessionOk

`func (o *DocumenttypeResponseCompound) GetEDocumenttypeRepossessionOk() (*FieldEDocumenttypeRepossession, bool)`

GetEDocumenttypeRepossessionOk returns a tuple with the EDocumenttypeRepossession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEDocumenttypeRepossession

`func (o *DocumenttypeResponseCompound) SetEDocumenttypeRepossession(v FieldEDocumenttypeRepossession)`

SetEDocumenttypeRepossession sets EDocumenttypeRepossession field to given value.

### HasEDocumenttypeRepossession

`func (o *DocumenttypeResponseCompound) HasEDocumenttypeRepossession() bool`

HasEDocumenttypeRepossession returns a boolean if a field has been set.

### GetAFkiInscriptiontypeID

`func (o *DocumenttypeResponseCompound) GetAFkiInscriptiontypeID() []int32`

GetAFkiInscriptiontypeID returns the AFkiInscriptiontypeID field if non-nil, zero value otherwise.

### GetAFkiInscriptiontypeIDOk

`func (o *DocumenttypeResponseCompound) GetAFkiInscriptiontypeIDOk() (*[]int32, bool)`

GetAFkiInscriptiontypeIDOk returns a tuple with the AFkiInscriptiontypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAFkiInscriptiontypeID

`func (o *DocumenttypeResponseCompound) SetAFkiInscriptiontypeID(v []int32)`

SetAFkiInscriptiontypeID sets AFkiInscriptiontypeID field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


