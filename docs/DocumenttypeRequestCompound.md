# DocumenttypeRequestCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiDocumenttypeID** | Pointer to **int32** | The unique ID of the Documenttype | [optional] 
**ObjDocumenttypeName** | [**MultilingualDocumenttypeName**](MultilingualDocumenttypeName.md) |  | 
**EDocumenttypeStep** | [**ComputedEDocumenttypeStep**](ComputedEDocumenttypeStep.md) |  | 
**SDocumenttypeDefaultfilename** | **string** | The default filename for the document type | 
**BDocumenttypeShowagent** | **bool** | Whether the document type should be shown to the agent | 
**BDocumenttypeIsactive** | **bool** | Whether the documenttype is active or not | 
**BDocumenttypeSale** | **bool** | Whether the document type applies to sale transactions | 
**BDocumenttypeRental** | **bool** | Whether the document type applies to rental transactions | 
**BDocumenttypeReference** | **bool** | Whether the document type is a reference document | 
**BDocumenttypeMortgagereference** | **bool** | Whether the document type is a mortgage reference | 
**BDocumenttypeEvaluation** | **bool** | Whether the document type applies to evaluation transactions | 
**BDocumenttypeConsultation** | **bool** | Whether the document type applies to consultation transactions | 
**BDocumenttypeOtherservices** | **bool** | Whether the document type applies to other services | 
**BDocumenttypeForeclosedsale** | **bool** | Whether the document type applies to foreclosed sale transactions | 
**BDocumenttypeCommercialfunding** | **bool** | Whether the document type applies to commercial funding transactions | 
**BDocumenttypeCanafe** | **bool** | Whether the document type applies to Canafe transactions | 
**EDocumenttypeEzcom** | Pointer to [**FieldEDocumenttypeEzcom**](FieldEDocumenttypeEzcom.md) |  | [optional] 
**EDocumenttypePrivacydefault** | [**FieldEDocumenttypePrivacydefault**](FieldEDocumenttypePrivacydefault.md) |  | 
**EDocumenttypeStatusdefault** | [**FieldEDocumenttypeStatusdefault**](FieldEDocumenttypeStatusdefault.md) |  | 
**EDocumenttypeSalebyowner** | [**FieldEDocumenttypeSalebyowner**](FieldEDocumenttypeSalebyowner.md) |  | 
**EDocumenttypeSaledisplaybyrole** | [**FieldEDocumenttypeSaledisplaybyrole**](FieldEDocumenttypeSaledisplaybyrole.md) |  | 
**EDocumenttypeTransactiontype** | [**FieldEDocumenttypeTransactiontype**](FieldEDocumenttypeTransactiontype.md) |  | 
**EDocumenttypeRepossession** | [**FieldEDocumenttypeRepossession**](FieldEDocumenttypeRepossession.md) |  | 
**AFkiInscriptiontypeID** | **[]int32** |  | 

## Methods

### NewDocumenttypeRequestCompound

`func NewDocumenttypeRequestCompound(objDocumenttypeName MultilingualDocumenttypeName, eDocumenttypeStep ComputedEDocumenttypeStep, sDocumenttypeDefaultfilename string, bDocumenttypeShowagent bool, bDocumenttypeIsactive bool, bDocumenttypeSale bool, bDocumenttypeRental bool, bDocumenttypeReference bool, bDocumenttypeMortgagereference bool, bDocumenttypeEvaluation bool, bDocumenttypeConsultation bool, bDocumenttypeOtherservices bool, bDocumenttypeForeclosedsale bool, bDocumenttypeCommercialfunding bool, bDocumenttypeCanafe bool, eDocumenttypePrivacydefault FieldEDocumenttypePrivacydefault, eDocumenttypeStatusdefault FieldEDocumenttypeStatusdefault, eDocumenttypeSalebyowner FieldEDocumenttypeSalebyowner, eDocumenttypeSaledisplaybyrole FieldEDocumenttypeSaledisplaybyrole, eDocumenttypeTransactiontype FieldEDocumenttypeTransactiontype, eDocumenttypeRepossession FieldEDocumenttypeRepossession, aFkiInscriptiontypeID []int32, ) *DocumenttypeRequestCompound`

NewDocumenttypeRequestCompound instantiates a new DocumenttypeRequestCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumenttypeRequestCompoundWithDefaults

`func NewDocumenttypeRequestCompoundWithDefaults() *DocumenttypeRequestCompound`

NewDocumenttypeRequestCompoundWithDefaults instantiates a new DocumenttypeRequestCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiDocumenttypeID

`func (o *DocumenttypeRequestCompound) GetPkiDocumenttypeID() int32`

GetPkiDocumenttypeID returns the PkiDocumenttypeID field if non-nil, zero value otherwise.

### GetPkiDocumenttypeIDOk

`func (o *DocumenttypeRequestCompound) GetPkiDocumenttypeIDOk() (*int32, bool)`

GetPkiDocumenttypeIDOk returns a tuple with the PkiDocumenttypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiDocumenttypeID

`func (o *DocumenttypeRequestCompound) SetPkiDocumenttypeID(v int32)`

SetPkiDocumenttypeID sets PkiDocumenttypeID field to given value.

### HasPkiDocumenttypeID

`func (o *DocumenttypeRequestCompound) HasPkiDocumenttypeID() bool`

HasPkiDocumenttypeID returns a boolean if a field has been set.

### GetObjDocumenttypeName

`func (o *DocumenttypeRequestCompound) GetObjDocumenttypeName() MultilingualDocumenttypeName`

GetObjDocumenttypeName returns the ObjDocumenttypeName field if non-nil, zero value otherwise.

### GetObjDocumenttypeNameOk

`func (o *DocumenttypeRequestCompound) GetObjDocumenttypeNameOk() (*MultilingualDocumenttypeName, bool)`

GetObjDocumenttypeNameOk returns a tuple with the ObjDocumenttypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjDocumenttypeName

`func (o *DocumenttypeRequestCompound) SetObjDocumenttypeName(v MultilingualDocumenttypeName)`

SetObjDocumenttypeName sets ObjDocumenttypeName field to given value.


### GetEDocumenttypeStep

`func (o *DocumenttypeRequestCompound) GetEDocumenttypeStep() ComputedEDocumenttypeStep`

GetEDocumenttypeStep returns the EDocumenttypeStep field if non-nil, zero value otherwise.

### GetEDocumenttypeStepOk

`func (o *DocumenttypeRequestCompound) GetEDocumenttypeStepOk() (*ComputedEDocumenttypeStep, bool)`

GetEDocumenttypeStepOk returns a tuple with the EDocumenttypeStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEDocumenttypeStep

`func (o *DocumenttypeRequestCompound) SetEDocumenttypeStep(v ComputedEDocumenttypeStep)`

SetEDocumenttypeStep sets EDocumenttypeStep field to given value.


### GetSDocumenttypeDefaultfilename

`func (o *DocumenttypeRequestCompound) GetSDocumenttypeDefaultfilename() string`

GetSDocumenttypeDefaultfilename returns the SDocumenttypeDefaultfilename field if non-nil, zero value otherwise.

### GetSDocumenttypeDefaultfilenameOk

`func (o *DocumenttypeRequestCompound) GetSDocumenttypeDefaultfilenameOk() (*string, bool)`

GetSDocumenttypeDefaultfilenameOk returns a tuple with the SDocumenttypeDefaultfilename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSDocumenttypeDefaultfilename

`func (o *DocumenttypeRequestCompound) SetSDocumenttypeDefaultfilename(v string)`

SetSDocumenttypeDefaultfilename sets SDocumenttypeDefaultfilename field to given value.


### GetBDocumenttypeShowagent

`func (o *DocumenttypeRequestCompound) GetBDocumenttypeShowagent() bool`

GetBDocumenttypeShowagent returns the BDocumenttypeShowagent field if non-nil, zero value otherwise.

### GetBDocumenttypeShowagentOk

`func (o *DocumenttypeRequestCompound) GetBDocumenttypeShowagentOk() (*bool, bool)`

GetBDocumenttypeShowagentOk returns a tuple with the BDocumenttypeShowagent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBDocumenttypeShowagent

`func (o *DocumenttypeRequestCompound) SetBDocumenttypeShowagent(v bool)`

SetBDocumenttypeShowagent sets BDocumenttypeShowagent field to given value.


### GetBDocumenttypeIsactive

`func (o *DocumenttypeRequestCompound) GetBDocumenttypeIsactive() bool`

GetBDocumenttypeIsactive returns the BDocumenttypeIsactive field if non-nil, zero value otherwise.

### GetBDocumenttypeIsactiveOk

`func (o *DocumenttypeRequestCompound) GetBDocumenttypeIsactiveOk() (*bool, bool)`

GetBDocumenttypeIsactiveOk returns a tuple with the BDocumenttypeIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBDocumenttypeIsactive

`func (o *DocumenttypeRequestCompound) SetBDocumenttypeIsactive(v bool)`

SetBDocumenttypeIsactive sets BDocumenttypeIsactive field to given value.


### GetBDocumenttypeSale

`func (o *DocumenttypeRequestCompound) GetBDocumenttypeSale() bool`

GetBDocumenttypeSale returns the BDocumenttypeSale field if non-nil, zero value otherwise.

### GetBDocumenttypeSaleOk

`func (o *DocumenttypeRequestCompound) GetBDocumenttypeSaleOk() (*bool, bool)`

GetBDocumenttypeSaleOk returns a tuple with the BDocumenttypeSale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBDocumenttypeSale

`func (o *DocumenttypeRequestCompound) SetBDocumenttypeSale(v bool)`

SetBDocumenttypeSale sets BDocumenttypeSale field to given value.


### GetBDocumenttypeRental

`func (o *DocumenttypeRequestCompound) GetBDocumenttypeRental() bool`

GetBDocumenttypeRental returns the BDocumenttypeRental field if non-nil, zero value otherwise.

### GetBDocumenttypeRentalOk

`func (o *DocumenttypeRequestCompound) GetBDocumenttypeRentalOk() (*bool, bool)`

GetBDocumenttypeRentalOk returns a tuple with the BDocumenttypeRental field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBDocumenttypeRental

`func (o *DocumenttypeRequestCompound) SetBDocumenttypeRental(v bool)`

SetBDocumenttypeRental sets BDocumenttypeRental field to given value.


### GetBDocumenttypeReference

`func (o *DocumenttypeRequestCompound) GetBDocumenttypeReference() bool`

GetBDocumenttypeReference returns the BDocumenttypeReference field if non-nil, zero value otherwise.

### GetBDocumenttypeReferenceOk

`func (o *DocumenttypeRequestCompound) GetBDocumenttypeReferenceOk() (*bool, bool)`

GetBDocumenttypeReferenceOk returns a tuple with the BDocumenttypeReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBDocumenttypeReference

`func (o *DocumenttypeRequestCompound) SetBDocumenttypeReference(v bool)`

SetBDocumenttypeReference sets BDocumenttypeReference field to given value.


### GetBDocumenttypeMortgagereference

`func (o *DocumenttypeRequestCompound) GetBDocumenttypeMortgagereference() bool`

GetBDocumenttypeMortgagereference returns the BDocumenttypeMortgagereference field if non-nil, zero value otherwise.

### GetBDocumenttypeMortgagereferenceOk

`func (o *DocumenttypeRequestCompound) GetBDocumenttypeMortgagereferenceOk() (*bool, bool)`

GetBDocumenttypeMortgagereferenceOk returns a tuple with the BDocumenttypeMortgagereference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBDocumenttypeMortgagereference

`func (o *DocumenttypeRequestCompound) SetBDocumenttypeMortgagereference(v bool)`

SetBDocumenttypeMortgagereference sets BDocumenttypeMortgagereference field to given value.


### GetBDocumenttypeEvaluation

`func (o *DocumenttypeRequestCompound) GetBDocumenttypeEvaluation() bool`

GetBDocumenttypeEvaluation returns the BDocumenttypeEvaluation field if non-nil, zero value otherwise.

### GetBDocumenttypeEvaluationOk

`func (o *DocumenttypeRequestCompound) GetBDocumenttypeEvaluationOk() (*bool, bool)`

GetBDocumenttypeEvaluationOk returns a tuple with the BDocumenttypeEvaluation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBDocumenttypeEvaluation

`func (o *DocumenttypeRequestCompound) SetBDocumenttypeEvaluation(v bool)`

SetBDocumenttypeEvaluation sets BDocumenttypeEvaluation field to given value.


### GetBDocumenttypeConsultation

`func (o *DocumenttypeRequestCompound) GetBDocumenttypeConsultation() bool`

GetBDocumenttypeConsultation returns the BDocumenttypeConsultation field if non-nil, zero value otherwise.

### GetBDocumenttypeConsultationOk

`func (o *DocumenttypeRequestCompound) GetBDocumenttypeConsultationOk() (*bool, bool)`

GetBDocumenttypeConsultationOk returns a tuple with the BDocumenttypeConsultation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBDocumenttypeConsultation

`func (o *DocumenttypeRequestCompound) SetBDocumenttypeConsultation(v bool)`

SetBDocumenttypeConsultation sets BDocumenttypeConsultation field to given value.


### GetBDocumenttypeOtherservices

`func (o *DocumenttypeRequestCompound) GetBDocumenttypeOtherservices() bool`

GetBDocumenttypeOtherservices returns the BDocumenttypeOtherservices field if non-nil, zero value otherwise.

### GetBDocumenttypeOtherservicesOk

`func (o *DocumenttypeRequestCompound) GetBDocumenttypeOtherservicesOk() (*bool, bool)`

GetBDocumenttypeOtherservicesOk returns a tuple with the BDocumenttypeOtherservices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBDocumenttypeOtherservices

`func (o *DocumenttypeRequestCompound) SetBDocumenttypeOtherservices(v bool)`

SetBDocumenttypeOtherservices sets BDocumenttypeOtherservices field to given value.


### GetBDocumenttypeForeclosedsale

`func (o *DocumenttypeRequestCompound) GetBDocumenttypeForeclosedsale() bool`

GetBDocumenttypeForeclosedsale returns the BDocumenttypeForeclosedsale field if non-nil, zero value otherwise.

### GetBDocumenttypeForeclosedsaleOk

`func (o *DocumenttypeRequestCompound) GetBDocumenttypeForeclosedsaleOk() (*bool, bool)`

GetBDocumenttypeForeclosedsaleOk returns a tuple with the BDocumenttypeForeclosedsale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBDocumenttypeForeclosedsale

`func (o *DocumenttypeRequestCompound) SetBDocumenttypeForeclosedsale(v bool)`

SetBDocumenttypeForeclosedsale sets BDocumenttypeForeclosedsale field to given value.


### GetBDocumenttypeCommercialfunding

`func (o *DocumenttypeRequestCompound) GetBDocumenttypeCommercialfunding() bool`

GetBDocumenttypeCommercialfunding returns the BDocumenttypeCommercialfunding field if non-nil, zero value otherwise.

### GetBDocumenttypeCommercialfundingOk

`func (o *DocumenttypeRequestCompound) GetBDocumenttypeCommercialfundingOk() (*bool, bool)`

GetBDocumenttypeCommercialfundingOk returns a tuple with the BDocumenttypeCommercialfunding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBDocumenttypeCommercialfunding

`func (o *DocumenttypeRequestCompound) SetBDocumenttypeCommercialfunding(v bool)`

SetBDocumenttypeCommercialfunding sets BDocumenttypeCommercialfunding field to given value.


### GetBDocumenttypeCanafe

`func (o *DocumenttypeRequestCompound) GetBDocumenttypeCanafe() bool`

GetBDocumenttypeCanafe returns the BDocumenttypeCanafe field if non-nil, zero value otherwise.

### GetBDocumenttypeCanafeOk

`func (o *DocumenttypeRequestCompound) GetBDocumenttypeCanafeOk() (*bool, bool)`

GetBDocumenttypeCanafeOk returns a tuple with the BDocumenttypeCanafe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBDocumenttypeCanafe

`func (o *DocumenttypeRequestCompound) SetBDocumenttypeCanafe(v bool)`

SetBDocumenttypeCanafe sets BDocumenttypeCanafe field to given value.


### GetEDocumenttypeEzcom

`func (o *DocumenttypeRequestCompound) GetEDocumenttypeEzcom() FieldEDocumenttypeEzcom`

GetEDocumenttypeEzcom returns the EDocumenttypeEzcom field if non-nil, zero value otherwise.

### GetEDocumenttypeEzcomOk

`func (o *DocumenttypeRequestCompound) GetEDocumenttypeEzcomOk() (*FieldEDocumenttypeEzcom, bool)`

GetEDocumenttypeEzcomOk returns a tuple with the EDocumenttypeEzcom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEDocumenttypeEzcom

`func (o *DocumenttypeRequestCompound) SetEDocumenttypeEzcom(v FieldEDocumenttypeEzcom)`

SetEDocumenttypeEzcom sets EDocumenttypeEzcom field to given value.

### HasEDocumenttypeEzcom

`func (o *DocumenttypeRequestCompound) HasEDocumenttypeEzcom() bool`

HasEDocumenttypeEzcom returns a boolean if a field has been set.

### GetEDocumenttypePrivacydefault

`func (o *DocumenttypeRequestCompound) GetEDocumenttypePrivacydefault() FieldEDocumenttypePrivacydefault`

GetEDocumenttypePrivacydefault returns the EDocumenttypePrivacydefault field if non-nil, zero value otherwise.

### GetEDocumenttypePrivacydefaultOk

`func (o *DocumenttypeRequestCompound) GetEDocumenttypePrivacydefaultOk() (*FieldEDocumenttypePrivacydefault, bool)`

GetEDocumenttypePrivacydefaultOk returns a tuple with the EDocumenttypePrivacydefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEDocumenttypePrivacydefault

`func (o *DocumenttypeRequestCompound) SetEDocumenttypePrivacydefault(v FieldEDocumenttypePrivacydefault)`

SetEDocumenttypePrivacydefault sets EDocumenttypePrivacydefault field to given value.


### GetEDocumenttypeStatusdefault

`func (o *DocumenttypeRequestCompound) GetEDocumenttypeStatusdefault() FieldEDocumenttypeStatusdefault`

GetEDocumenttypeStatusdefault returns the EDocumenttypeStatusdefault field if non-nil, zero value otherwise.

### GetEDocumenttypeStatusdefaultOk

`func (o *DocumenttypeRequestCompound) GetEDocumenttypeStatusdefaultOk() (*FieldEDocumenttypeStatusdefault, bool)`

GetEDocumenttypeStatusdefaultOk returns a tuple with the EDocumenttypeStatusdefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEDocumenttypeStatusdefault

`func (o *DocumenttypeRequestCompound) SetEDocumenttypeStatusdefault(v FieldEDocumenttypeStatusdefault)`

SetEDocumenttypeStatusdefault sets EDocumenttypeStatusdefault field to given value.


### GetEDocumenttypeSalebyowner

`func (o *DocumenttypeRequestCompound) GetEDocumenttypeSalebyowner() FieldEDocumenttypeSalebyowner`

GetEDocumenttypeSalebyowner returns the EDocumenttypeSalebyowner field if non-nil, zero value otherwise.

### GetEDocumenttypeSalebyownerOk

`func (o *DocumenttypeRequestCompound) GetEDocumenttypeSalebyownerOk() (*FieldEDocumenttypeSalebyowner, bool)`

GetEDocumenttypeSalebyownerOk returns a tuple with the EDocumenttypeSalebyowner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEDocumenttypeSalebyowner

`func (o *DocumenttypeRequestCompound) SetEDocumenttypeSalebyowner(v FieldEDocumenttypeSalebyowner)`

SetEDocumenttypeSalebyowner sets EDocumenttypeSalebyowner field to given value.


### GetEDocumenttypeSaledisplaybyrole

`func (o *DocumenttypeRequestCompound) GetEDocumenttypeSaledisplaybyrole() FieldEDocumenttypeSaledisplaybyrole`

GetEDocumenttypeSaledisplaybyrole returns the EDocumenttypeSaledisplaybyrole field if non-nil, zero value otherwise.

### GetEDocumenttypeSaledisplaybyroleOk

`func (o *DocumenttypeRequestCompound) GetEDocumenttypeSaledisplaybyroleOk() (*FieldEDocumenttypeSaledisplaybyrole, bool)`

GetEDocumenttypeSaledisplaybyroleOk returns a tuple with the EDocumenttypeSaledisplaybyrole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEDocumenttypeSaledisplaybyrole

`func (o *DocumenttypeRequestCompound) SetEDocumenttypeSaledisplaybyrole(v FieldEDocumenttypeSaledisplaybyrole)`

SetEDocumenttypeSaledisplaybyrole sets EDocumenttypeSaledisplaybyrole field to given value.


### GetEDocumenttypeTransactiontype

`func (o *DocumenttypeRequestCompound) GetEDocumenttypeTransactiontype() FieldEDocumenttypeTransactiontype`

GetEDocumenttypeTransactiontype returns the EDocumenttypeTransactiontype field if non-nil, zero value otherwise.

### GetEDocumenttypeTransactiontypeOk

`func (o *DocumenttypeRequestCompound) GetEDocumenttypeTransactiontypeOk() (*FieldEDocumenttypeTransactiontype, bool)`

GetEDocumenttypeTransactiontypeOk returns a tuple with the EDocumenttypeTransactiontype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEDocumenttypeTransactiontype

`func (o *DocumenttypeRequestCompound) SetEDocumenttypeTransactiontype(v FieldEDocumenttypeTransactiontype)`

SetEDocumenttypeTransactiontype sets EDocumenttypeTransactiontype field to given value.


### GetEDocumenttypeRepossession

`func (o *DocumenttypeRequestCompound) GetEDocumenttypeRepossession() FieldEDocumenttypeRepossession`

GetEDocumenttypeRepossession returns the EDocumenttypeRepossession field if non-nil, zero value otherwise.

### GetEDocumenttypeRepossessionOk

`func (o *DocumenttypeRequestCompound) GetEDocumenttypeRepossessionOk() (*FieldEDocumenttypeRepossession, bool)`

GetEDocumenttypeRepossessionOk returns a tuple with the EDocumenttypeRepossession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEDocumenttypeRepossession

`func (o *DocumenttypeRequestCompound) SetEDocumenttypeRepossession(v FieldEDocumenttypeRepossession)`

SetEDocumenttypeRepossession sets EDocumenttypeRepossession field to given value.


### GetAFkiInscriptiontypeID

`func (o *DocumenttypeRequestCompound) GetAFkiInscriptiontypeID() []int32`

GetAFkiInscriptiontypeID returns the AFkiInscriptiontypeID field if non-nil, zero value otherwise.

### GetAFkiInscriptiontypeIDOk

`func (o *DocumenttypeRequestCompound) GetAFkiInscriptiontypeIDOk() (*[]int32, bool)`

GetAFkiInscriptiontypeIDOk returns a tuple with the AFkiInscriptiontypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAFkiInscriptiontypeID

`func (o *DocumenttypeRequestCompound) SetAFkiInscriptiontypeID(v []int32)`

SetAFkiInscriptiontypeID sets AFkiInscriptiontypeID field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


