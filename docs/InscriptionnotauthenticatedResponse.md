# InscriptionnotauthenticatedResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiInscriptionnotauthenticatedID** | **int32** | The unique ID of the Inscriptionnotauthenticated. | 
**FkiCompanyID** | **int32** | The unique ID of the Company | 
**SCompanyNameX** | Pointer to **string** | The Name of the Company in the language of the requester | [optional] 
**FkiInscriptionID** | **int32** | The unique ID of the Inscription. | 
**FkiDepartmentID** | **int32** | The unique ID of the Department | 
**SDepartmentNameX** | Pointer to **string** | The Name of the Department in the language of the requester | [optional] 
**FkiFinancialinstitutionID** | **int32** | The unique ID of the Financialinstitution | 
**SFinancialinstitutionNameX** | Pointer to **string** | The name of the Financialinstitution in the language of the requester | [optional] 
**FkiBuyercontractID** | **int32** | The unique ID of the Buyercontract | 
**SBuyercontractContract** | Pointer to **string** | The number of the Buyercontract | [optional] 
**FkiMortgagesupplierID** | **int32** | The unique ID of the Mortgagesupplier | 
**SMortgagesupplierNameX** | Pointer to **string** | The name of the Mortagesupplier in the language of the requester | [optional] 
**FkiTaxassignmentID** | **int32** | The unique ID of the Taxassignment.  Valid values:  |Value|Description| |-|-| |1|No tax| |2|GST| |3|HST (ON)| |4|HST (NB)| |5|HST (NS)| |6|HST (NL)| |7|HST (PE)| |8|GST + QST (QC)| |9|GST + QST (QC) Non-Recoverable| |10|GST + PST (BC)| |11|GST + PST (SK)| |12|GST + RST (MB)| |13|GST + PST (BC) Non-Recoverable| |14|GST + PST (SK) Non-Recoverable| |15|GST + RST (MB) Non-Recoverable| | 
**STaxassignmentDescriptionX** | Pointer to **string** | The description of the Taxassignment  in the language of the requester | [optional] 
**DtInscriptionnotauthenticatedTransactiondate** | **string** | The transactiondate of the Inscriptionnotauthenticated | 
**DtInscriptionnotauthenticatedTransactiondateReal** | **string** | The transactiondatereal of the Inscriptionnotauthenticated | 
**DtInscriptionnotauthenticatedDepositdate** | **string** | The depositdate of the Inscriptionnotauthenticated | 
**EInscriptionnotauthenticatedType** | [**FieldEInscriptionnotauthenticatedType**](FieldEInscriptionnotauthenticatedType.md) |  | 
**DInscriptionnotauthenticatedMortgageloan** | **string** | The mortgageloan of the Inscriptionnotauthenticated | 
**EtInscriptionnotauthenticatedMortgagetype** | [**FieldEtInscriptionnotauthenticatedMortgagetype**](FieldEtInscriptionnotauthenticatedMortgagetype.md) |  | 
**DInscriptionnotauthenticatedTransactionprice** | **string** | The transactionprice of the Inscriptionnotauthenticated | 
**EInscriptionnotauthenticatedRemunerationtype** | [**FieldEInscriptionnotauthenticatedRemunerationtype**](FieldEInscriptionnotauthenticatedRemunerationtype.md) |  | 
**DInscriptionnotauthenticatedRemuneration** | **string** | The remuneration of the Inscriptionnotauthenticated | 
**DInscriptionnotauthenticatedRemunerationsubtotal** | **string** | The remunerationsubtotal of the Inscriptionnotauthenticated | 
**DInscriptionnotauthenticatedRemunerationtotal** | **string** | The remunerationtotal of the Inscriptionnotauthenticated | 
**DtInscriptionnotauthenticatedCancellationdate** | **string** | The cancellationdate of the Inscriptionnotauthenticated | 
**DtInscriptionnotauthenticatedPossessiondate** | **string** | The possessiondate of the Inscriptionnotauthenticated | 
**SInscriptionnotauthenticatedOffertopurchasenumber** | **string** | The Offer to purchase number | 
**DtInscriptionnotauthenticatedNotaryscheduledate** | **string** | The notaryscheduledate of the Inscriptionnotauthenticated | 
**DtInscriptionnotauthenticatedFinancingscheduledate** | **string** | The financingscheduledate of the Inscriptionnotauthenticated | 
**BInscriptionnotauthenticatedConditional** | **bool** | Whether the inscriptionnotauthenticated is conditional | 
**BInscriptionnotauthenticatedMortgageisreferenced** | **bool** | Whether if it&#39;s an mortgageisreferenced | 
**BInscriptionnotauthenticatedHomeowner** | **bool** | Whether if it&#39;s an homeowner | 
**TInscriptionnotauthenticatedConditions** | **string** | The conditions of the Inscriptionnotauthenticated | 
**DtInscriptionnotauthenticatedConditiondeadlinedate** | **string** | The conditiondeadlinedate of the Inscriptionnotauthenticated | 
**IInscriptionnotauthenticatedOrder** | **int32** | The order of the Inscriptionnotauthenticated | 
**BInscriptionnotauthenticatedIsactive** | **bool** | Whether the inscriptionnotauthenticated is active or not | 
**EInscriptionnotauthenticatedResidenceType** | [**FieldEInscriptionnotauthenticatedResidenceType**](FieldEInscriptionnotauthenticatedResidenceType.md) |  | 
**TInscriptionnotauthenticatedChecklistnote** | **string** | The checklistnote of the Inscriptionnotauthenticated | 
**DInscriptionnotauthenticatedSelleronlyretribution** | **string** | The selleronlyretribution of the Inscriptionnotauthenticated | 
**BInscriptionnotauthenticatedDraft** | **bool** | Whether the inscriptionnotauthenticated is a draft or not | 

## Methods

### NewInscriptionnotauthenticatedResponse

`func NewInscriptionnotauthenticatedResponse(pkiInscriptionnotauthenticatedID int32, fkiCompanyID int32, fkiInscriptionID int32, fkiDepartmentID int32, fkiFinancialinstitutionID int32, fkiBuyercontractID int32, fkiMortgagesupplierID int32, fkiTaxassignmentID int32, dtInscriptionnotauthenticatedTransactiondate string, dtInscriptionnotauthenticatedTransactiondateReal string, dtInscriptionnotauthenticatedDepositdate string, eInscriptionnotauthenticatedType FieldEInscriptionnotauthenticatedType, dInscriptionnotauthenticatedMortgageloan string, etInscriptionnotauthenticatedMortgagetype FieldEtInscriptionnotauthenticatedMortgagetype, dInscriptionnotauthenticatedTransactionprice string, eInscriptionnotauthenticatedRemunerationtype FieldEInscriptionnotauthenticatedRemunerationtype, dInscriptionnotauthenticatedRemuneration string, dInscriptionnotauthenticatedRemunerationsubtotal string, dInscriptionnotauthenticatedRemunerationtotal string, dtInscriptionnotauthenticatedCancellationdate string, dtInscriptionnotauthenticatedPossessiondate string, sInscriptionnotauthenticatedOffertopurchasenumber string, dtInscriptionnotauthenticatedNotaryscheduledate string, dtInscriptionnotauthenticatedFinancingscheduledate string, bInscriptionnotauthenticatedConditional bool, bInscriptionnotauthenticatedMortgageisreferenced bool, bInscriptionnotauthenticatedHomeowner bool, tInscriptionnotauthenticatedConditions string, dtInscriptionnotauthenticatedConditiondeadlinedate string, iInscriptionnotauthenticatedOrder int32, bInscriptionnotauthenticatedIsactive bool, eInscriptionnotauthenticatedResidenceType FieldEInscriptionnotauthenticatedResidenceType, tInscriptionnotauthenticatedChecklistnote string, dInscriptionnotauthenticatedSelleronlyretribution string, bInscriptionnotauthenticatedDraft bool, ) *InscriptionnotauthenticatedResponse`

NewInscriptionnotauthenticatedResponse instantiates a new InscriptionnotauthenticatedResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInscriptionnotauthenticatedResponseWithDefaults

`func NewInscriptionnotauthenticatedResponseWithDefaults() *InscriptionnotauthenticatedResponse`

NewInscriptionnotauthenticatedResponseWithDefaults instantiates a new InscriptionnotauthenticatedResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiInscriptionnotauthenticatedID

`func (o *InscriptionnotauthenticatedResponse) GetPkiInscriptionnotauthenticatedID() int32`

GetPkiInscriptionnotauthenticatedID returns the PkiInscriptionnotauthenticatedID field if non-nil, zero value otherwise.

### GetPkiInscriptionnotauthenticatedIDOk

`func (o *InscriptionnotauthenticatedResponse) GetPkiInscriptionnotauthenticatedIDOk() (*int32, bool)`

GetPkiInscriptionnotauthenticatedIDOk returns a tuple with the PkiInscriptionnotauthenticatedID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiInscriptionnotauthenticatedID

`func (o *InscriptionnotauthenticatedResponse) SetPkiInscriptionnotauthenticatedID(v int32)`

SetPkiInscriptionnotauthenticatedID sets PkiInscriptionnotauthenticatedID field to given value.


### GetFkiCompanyID

`func (o *InscriptionnotauthenticatedResponse) GetFkiCompanyID() int32`

GetFkiCompanyID returns the FkiCompanyID field if non-nil, zero value otherwise.

### GetFkiCompanyIDOk

`func (o *InscriptionnotauthenticatedResponse) GetFkiCompanyIDOk() (*int32, bool)`

GetFkiCompanyIDOk returns a tuple with the FkiCompanyID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiCompanyID

`func (o *InscriptionnotauthenticatedResponse) SetFkiCompanyID(v int32)`

SetFkiCompanyID sets FkiCompanyID field to given value.


### GetSCompanyNameX

`func (o *InscriptionnotauthenticatedResponse) GetSCompanyNameX() string`

GetSCompanyNameX returns the SCompanyNameX field if non-nil, zero value otherwise.

### GetSCompanyNameXOk

`func (o *InscriptionnotauthenticatedResponse) GetSCompanyNameXOk() (*string, bool)`

GetSCompanyNameXOk returns a tuple with the SCompanyNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSCompanyNameX

`func (o *InscriptionnotauthenticatedResponse) SetSCompanyNameX(v string)`

SetSCompanyNameX sets SCompanyNameX field to given value.

### HasSCompanyNameX

`func (o *InscriptionnotauthenticatedResponse) HasSCompanyNameX() bool`

HasSCompanyNameX returns a boolean if a field has been set.

### GetFkiInscriptionID

`func (o *InscriptionnotauthenticatedResponse) GetFkiInscriptionID() int32`

GetFkiInscriptionID returns the FkiInscriptionID field if non-nil, zero value otherwise.

### GetFkiInscriptionIDOk

`func (o *InscriptionnotauthenticatedResponse) GetFkiInscriptionIDOk() (*int32, bool)`

GetFkiInscriptionIDOk returns a tuple with the FkiInscriptionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiInscriptionID

`func (o *InscriptionnotauthenticatedResponse) SetFkiInscriptionID(v int32)`

SetFkiInscriptionID sets FkiInscriptionID field to given value.


### GetFkiDepartmentID

`func (o *InscriptionnotauthenticatedResponse) GetFkiDepartmentID() int32`

GetFkiDepartmentID returns the FkiDepartmentID field if non-nil, zero value otherwise.

### GetFkiDepartmentIDOk

`func (o *InscriptionnotauthenticatedResponse) GetFkiDepartmentIDOk() (*int32, bool)`

GetFkiDepartmentIDOk returns a tuple with the FkiDepartmentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiDepartmentID

`func (o *InscriptionnotauthenticatedResponse) SetFkiDepartmentID(v int32)`

SetFkiDepartmentID sets FkiDepartmentID field to given value.


### GetSDepartmentNameX

`func (o *InscriptionnotauthenticatedResponse) GetSDepartmentNameX() string`

GetSDepartmentNameX returns the SDepartmentNameX field if non-nil, zero value otherwise.

### GetSDepartmentNameXOk

`func (o *InscriptionnotauthenticatedResponse) GetSDepartmentNameXOk() (*string, bool)`

GetSDepartmentNameXOk returns a tuple with the SDepartmentNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSDepartmentNameX

`func (o *InscriptionnotauthenticatedResponse) SetSDepartmentNameX(v string)`

SetSDepartmentNameX sets SDepartmentNameX field to given value.

### HasSDepartmentNameX

`func (o *InscriptionnotauthenticatedResponse) HasSDepartmentNameX() bool`

HasSDepartmentNameX returns a boolean if a field has been set.

### GetFkiFinancialinstitutionID

`func (o *InscriptionnotauthenticatedResponse) GetFkiFinancialinstitutionID() int32`

GetFkiFinancialinstitutionID returns the FkiFinancialinstitutionID field if non-nil, zero value otherwise.

### GetFkiFinancialinstitutionIDOk

`func (o *InscriptionnotauthenticatedResponse) GetFkiFinancialinstitutionIDOk() (*int32, bool)`

GetFkiFinancialinstitutionIDOk returns a tuple with the FkiFinancialinstitutionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiFinancialinstitutionID

`func (o *InscriptionnotauthenticatedResponse) SetFkiFinancialinstitutionID(v int32)`

SetFkiFinancialinstitutionID sets FkiFinancialinstitutionID field to given value.


### GetSFinancialinstitutionNameX

`func (o *InscriptionnotauthenticatedResponse) GetSFinancialinstitutionNameX() string`

GetSFinancialinstitutionNameX returns the SFinancialinstitutionNameX field if non-nil, zero value otherwise.

### GetSFinancialinstitutionNameXOk

`func (o *InscriptionnotauthenticatedResponse) GetSFinancialinstitutionNameXOk() (*string, bool)`

GetSFinancialinstitutionNameXOk returns a tuple with the SFinancialinstitutionNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSFinancialinstitutionNameX

`func (o *InscriptionnotauthenticatedResponse) SetSFinancialinstitutionNameX(v string)`

SetSFinancialinstitutionNameX sets SFinancialinstitutionNameX field to given value.

### HasSFinancialinstitutionNameX

`func (o *InscriptionnotauthenticatedResponse) HasSFinancialinstitutionNameX() bool`

HasSFinancialinstitutionNameX returns a boolean if a field has been set.

### GetFkiBuyercontractID

`func (o *InscriptionnotauthenticatedResponse) GetFkiBuyercontractID() int32`

GetFkiBuyercontractID returns the FkiBuyercontractID field if non-nil, zero value otherwise.

### GetFkiBuyercontractIDOk

`func (o *InscriptionnotauthenticatedResponse) GetFkiBuyercontractIDOk() (*int32, bool)`

GetFkiBuyercontractIDOk returns a tuple with the FkiBuyercontractID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiBuyercontractID

`func (o *InscriptionnotauthenticatedResponse) SetFkiBuyercontractID(v int32)`

SetFkiBuyercontractID sets FkiBuyercontractID field to given value.


### GetSBuyercontractContract

`func (o *InscriptionnotauthenticatedResponse) GetSBuyercontractContract() string`

GetSBuyercontractContract returns the SBuyercontractContract field if non-nil, zero value otherwise.

### GetSBuyercontractContractOk

`func (o *InscriptionnotauthenticatedResponse) GetSBuyercontractContractOk() (*string, bool)`

GetSBuyercontractContractOk returns a tuple with the SBuyercontractContract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSBuyercontractContract

`func (o *InscriptionnotauthenticatedResponse) SetSBuyercontractContract(v string)`

SetSBuyercontractContract sets SBuyercontractContract field to given value.

### HasSBuyercontractContract

`func (o *InscriptionnotauthenticatedResponse) HasSBuyercontractContract() bool`

HasSBuyercontractContract returns a boolean if a field has been set.

### GetFkiMortgagesupplierID

`func (o *InscriptionnotauthenticatedResponse) GetFkiMortgagesupplierID() int32`

GetFkiMortgagesupplierID returns the FkiMortgagesupplierID field if non-nil, zero value otherwise.

### GetFkiMortgagesupplierIDOk

`func (o *InscriptionnotauthenticatedResponse) GetFkiMortgagesupplierIDOk() (*int32, bool)`

GetFkiMortgagesupplierIDOk returns a tuple with the FkiMortgagesupplierID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiMortgagesupplierID

`func (o *InscriptionnotauthenticatedResponse) SetFkiMortgagesupplierID(v int32)`

SetFkiMortgagesupplierID sets FkiMortgagesupplierID field to given value.


### GetSMortgagesupplierNameX

`func (o *InscriptionnotauthenticatedResponse) GetSMortgagesupplierNameX() string`

GetSMortgagesupplierNameX returns the SMortgagesupplierNameX field if non-nil, zero value otherwise.

### GetSMortgagesupplierNameXOk

`func (o *InscriptionnotauthenticatedResponse) GetSMortgagesupplierNameXOk() (*string, bool)`

GetSMortgagesupplierNameXOk returns a tuple with the SMortgagesupplierNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSMortgagesupplierNameX

`func (o *InscriptionnotauthenticatedResponse) SetSMortgagesupplierNameX(v string)`

SetSMortgagesupplierNameX sets SMortgagesupplierNameX field to given value.

### HasSMortgagesupplierNameX

`func (o *InscriptionnotauthenticatedResponse) HasSMortgagesupplierNameX() bool`

HasSMortgagesupplierNameX returns a boolean if a field has been set.

### GetFkiTaxassignmentID

`func (o *InscriptionnotauthenticatedResponse) GetFkiTaxassignmentID() int32`

GetFkiTaxassignmentID returns the FkiTaxassignmentID field if non-nil, zero value otherwise.

### GetFkiTaxassignmentIDOk

`func (o *InscriptionnotauthenticatedResponse) GetFkiTaxassignmentIDOk() (*int32, bool)`

GetFkiTaxassignmentIDOk returns a tuple with the FkiTaxassignmentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiTaxassignmentID

`func (o *InscriptionnotauthenticatedResponse) SetFkiTaxassignmentID(v int32)`

SetFkiTaxassignmentID sets FkiTaxassignmentID field to given value.


### GetSTaxassignmentDescriptionX

`func (o *InscriptionnotauthenticatedResponse) GetSTaxassignmentDescriptionX() string`

GetSTaxassignmentDescriptionX returns the STaxassignmentDescriptionX field if non-nil, zero value otherwise.

### GetSTaxassignmentDescriptionXOk

`func (o *InscriptionnotauthenticatedResponse) GetSTaxassignmentDescriptionXOk() (*string, bool)`

GetSTaxassignmentDescriptionXOk returns a tuple with the STaxassignmentDescriptionX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSTaxassignmentDescriptionX

`func (o *InscriptionnotauthenticatedResponse) SetSTaxassignmentDescriptionX(v string)`

SetSTaxassignmentDescriptionX sets STaxassignmentDescriptionX field to given value.

### HasSTaxassignmentDescriptionX

`func (o *InscriptionnotauthenticatedResponse) HasSTaxassignmentDescriptionX() bool`

HasSTaxassignmentDescriptionX returns a boolean if a field has been set.

### GetDtInscriptionnotauthenticatedTransactiondate

`func (o *InscriptionnotauthenticatedResponse) GetDtInscriptionnotauthenticatedTransactiondate() string`

GetDtInscriptionnotauthenticatedTransactiondate returns the DtInscriptionnotauthenticatedTransactiondate field if non-nil, zero value otherwise.

### GetDtInscriptionnotauthenticatedTransactiondateOk

`func (o *InscriptionnotauthenticatedResponse) GetDtInscriptionnotauthenticatedTransactiondateOk() (*string, bool)`

GetDtInscriptionnotauthenticatedTransactiondateOk returns a tuple with the DtInscriptionnotauthenticatedTransactiondate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtInscriptionnotauthenticatedTransactiondate

`func (o *InscriptionnotauthenticatedResponse) SetDtInscriptionnotauthenticatedTransactiondate(v string)`

SetDtInscriptionnotauthenticatedTransactiondate sets DtInscriptionnotauthenticatedTransactiondate field to given value.


### GetDtInscriptionnotauthenticatedTransactiondateReal

`func (o *InscriptionnotauthenticatedResponse) GetDtInscriptionnotauthenticatedTransactiondateReal() string`

GetDtInscriptionnotauthenticatedTransactiondateReal returns the DtInscriptionnotauthenticatedTransactiondateReal field if non-nil, zero value otherwise.

### GetDtInscriptionnotauthenticatedTransactiondateRealOk

`func (o *InscriptionnotauthenticatedResponse) GetDtInscriptionnotauthenticatedTransactiondateRealOk() (*string, bool)`

GetDtInscriptionnotauthenticatedTransactiondateRealOk returns a tuple with the DtInscriptionnotauthenticatedTransactiondateReal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtInscriptionnotauthenticatedTransactiondateReal

`func (o *InscriptionnotauthenticatedResponse) SetDtInscriptionnotauthenticatedTransactiondateReal(v string)`

SetDtInscriptionnotauthenticatedTransactiondateReal sets DtInscriptionnotauthenticatedTransactiondateReal field to given value.


### GetDtInscriptionnotauthenticatedDepositdate

`func (o *InscriptionnotauthenticatedResponse) GetDtInscriptionnotauthenticatedDepositdate() string`

GetDtInscriptionnotauthenticatedDepositdate returns the DtInscriptionnotauthenticatedDepositdate field if non-nil, zero value otherwise.

### GetDtInscriptionnotauthenticatedDepositdateOk

`func (o *InscriptionnotauthenticatedResponse) GetDtInscriptionnotauthenticatedDepositdateOk() (*string, bool)`

GetDtInscriptionnotauthenticatedDepositdateOk returns a tuple with the DtInscriptionnotauthenticatedDepositdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtInscriptionnotauthenticatedDepositdate

`func (o *InscriptionnotauthenticatedResponse) SetDtInscriptionnotauthenticatedDepositdate(v string)`

SetDtInscriptionnotauthenticatedDepositdate sets DtInscriptionnotauthenticatedDepositdate field to given value.


### GetEInscriptionnotauthenticatedType

`func (o *InscriptionnotauthenticatedResponse) GetEInscriptionnotauthenticatedType() FieldEInscriptionnotauthenticatedType`

GetEInscriptionnotauthenticatedType returns the EInscriptionnotauthenticatedType field if non-nil, zero value otherwise.

### GetEInscriptionnotauthenticatedTypeOk

`func (o *InscriptionnotauthenticatedResponse) GetEInscriptionnotauthenticatedTypeOk() (*FieldEInscriptionnotauthenticatedType, bool)`

GetEInscriptionnotauthenticatedTypeOk returns a tuple with the EInscriptionnotauthenticatedType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEInscriptionnotauthenticatedType

`func (o *InscriptionnotauthenticatedResponse) SetEInscriptionnotauthenticatedType(v FieldEInscriptionnotauthenticatedType)`

SetEInscriptionnotauthenticatedType sets EInscriptionnotauthenticatedType field to given value.


### GetDInscriptionnotauthenticatedMortgageloan

`func (o *InscriptionnotauthenticatedResponse) GetDInscriptionnotauthenticatedMortgageloan() string`

GetDInscriptionnotauthenticatedMortgageloan returns the DInscriptionnotauthenticatedMortgageloan field if non-nil, zero value otherwise.

### GetDInscriptionnotauthenticatedMortgageloanOk

`func (o *InscriptionnotauthenticatedResponse) GetDInscriptionnotauthenticatedMortgageloanOk() (*string, bool)`

GetDInscriptionnotauthenticatedMortgageloanOk returns a tuple with the DInscriptionnotauthenticatedMortgageloan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDInscriptionnotauthenticatedMortgageloan

`func (o *InscriptionnotauthenticatedResponse) SetDInscriptionnotauthenticatedMortgageloan(v string)`

SetDInscriptionnotauthenticatedMortgageloan sets DInscriptionnotauthenticatedMortgageloan field to given value.


### GetEtInscriptionnotauthenticatedMortgagetype

`func (o *InscriptionnotauthenticatedResponse) GetEtInscriptionnotauthenticatedMortgagetype() FieldEtInscriptionnotauthenticatedMortgagetype`

GetEtInscriptionnotauthenticatedMortgagetype returns the EtInscriptionnotauthenticatedMortgagetype field if non-nil, zero value otherwise.

### GetEtInscriptionnotauthenticatedMortgagetypeOk

`func (o *InscriptionnotauthenticatedResponse) GetEtInscriptionnotauthenticatedMortgagetypeOk() (*FieldEtInscriptionnotauthenticatedMortgagetype, bool)`

GetEtInscriptionnotauthenticatedMortgagetypeOk returns a tuple with the EtInscriptionnotauthenticatedMortgagetype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtInscriptionnotauthenticatedMortgagetype

`func (o *InscriptionnotauthenticatedResponse) SetEtInscriptionnotauthenticatedMortgagetype(v FieldEtInscriptionnotauthenticatedMortgagetype)`

SetEtInscriptionnotauthenticatedMortgagetype sets EtInscriptionnotauthenticatedMortgagetype field to given value.


### GetDInscriptionnotauthenticatedTransactionprice

`func (o *InscriptionnotauthenticatedResponse) GetDInscriptionnotauthenticatedTransactionprice() string`

GetDInscriptionnotauthenticatedTransactionprice returns the DInscriptionnotauthenticatedTransactionprice field if non-nil, zero value otherwise.

### GetDInscriptionnotauthenticatedTransactionpriceOk

`func (o *InscriptionnotauthenticatedResponse) GetDInscriptionnotauthenticatedTransactionpriceOk() (*string, bool)`

GetDInscriptionnotauthenticatedTransactionpriceOk returns a tuple with the DInscriptionnotauthenticatedTransactionprice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDInscriptionnotauthenticatedTransactionprice

`func (o *InscriptionnotauthenticatedResponse) SetDInscriptionnotauthenticatedTransactionprice(v string)`

SetDInscriptionnotauthenticatedTransactionprice sets DInscriptionnotauthenticatedTransactionprice field to given value.


### GetEInscriptionnotauthenticatedRemunerationtype

`func (o *InscriptionnotauthenticatedResponse) GetEInscriptionnotauthenticatedRemunerationtype() FieldEInscriptionnotauthenticatedRemunerationtype`

GetEInscriptionnotauthenticatedRemunerationtype returns the EInscriptionnotauthenticatedRemunerationtype field if non-nil, zero value otherwise.

### GetEInscriptionnotauthenticatedRemunerationtypeOk

`func (o *InscriptionnotauthenticatedResponse) GetEInscriptionnotauthenticatedRemunerationtypeOk() (*FieldEInscriptionnotauthenticatedRemunerationtype, bool)`

GetEInscriptionnotauthenticatedRemunerationtypeOk returns a tuple with the EInscriptionnotauthenticatedRemunerationtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEInscriptionnotauthenticatedRemunerationtype

`func (o *InscriptionnotauthenticatedResponse) SetEInscriptionnotauthenticatedRemunerationtype(v FieldEInscriptionnotauthenticatedRemunerationtype)`

SetEInscriptionnotauthenticatedRemunerationtype sets EInscriptionnotauthenticatedRemunerationtype field to given value.


### GetDInscriptionnotauthenticatedRemuneration

`func (o *InscriptionnotauthenticatedResponse) GetDInscriptionnotauthenticatedRemuneration() string`

GetDInscriptionnotauthenticatedRemuneration returns the DInscriptionnotauthenticatedRemuneration field if non-nil, zero value otherwise.

### GetDInscriptionnotauthenticatedRemunerationOk

`func (o *InscriptionnotauthenticatedResponse) GetDInscriptionnotauthenticatedRemunerationOk() (*string, bool)`

GetDInscriptionnotauthenticatedRemunerationOk returns a tuple with the DInscriptionnotauthenticatedRemuneration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDInscriptionnotauthenticatedRemuneration

`func (o *InscriptionnotauthenticatedResponse) SetDInscriptionnotauthenticatedRemuneration(v string)`

SetDInscriptionnotauthenticatedRemuneration sets DInscriptionnotauthenticatedRemuneration field to given value.


### GetDInscriptionnotauthenticatedRemunerationsubtotal

`func (o *InscriptionnotauthenticatedResponse) GetDInscriptionnotauthenticatedRemunerationsubtotal() string`

GetDInscriptionnotauthenticatedRemunerationsubtotal returns the DInscriptionnotauthenticatedRemunerationsubtotal field if non-nil, zero value otherwise.

### GetDInscriptionnotauthenticatedRemunerationsubtotalOk

`func (o *InscriptionnotauthenticatedResponse) GetDInscriptionnotauthenticatedRemunerationsubtotalOk() (*string, bool)`

GetDInscriptionnotauthenticatedRemunerationsubtotalOk returns a tuple with the DInscriptionnotauthenticatedRemunerationsubtotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDInscriptionnotauthenticatedRemunerationsubtotal

`func (o *InscriptionnotauthenticatedResponse) SetDInscriptionnotauthenticatedRemunerationsubtotal(v string)`

SetDInscriptionnotauthenticatedRemunerationsubtotal sets DInscriptionnotauthenticatedRemunerationsubtotal field to given value.


### GetDInscriptionnotauthenticatedRemunerationtotal

`func (o *InscriptionnotauthenticatedResponse) GetDInscriptionnotauthenticatedRemunerationtotal() string`

GetDInscriptionnotauthenticatedRemunerationtotal returns the DInscriptionnotauthenticatedRemunerationtotal field if non-nil, zero value otherwise.

### GetDInscriptionnotauthenticatedRemunerationtotalOk

`func (o *InscriptionnotauthenticatedResponse) GetDInscriptionnotauthenticatedRemunerationtotalOk() (*string, bool)`

GetDInscriptionnotauthenticatedRemunerationtotalOk returns a tuple with the DInscriptionnotauthenticatedRemunerationtotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDInscriptionnotauthenticatedRemunerationtotal

`func (o *InscriptionnotauthenticatedResponse) SetDInscriptionnotauthenticatedRemunerationtotal(v string)`

SetDInscriptionnotauthenticatedRemunerationtotal sets DInscriptionnotauthenticatedRemunerationtotal field to given value.


### GetDtInscriptionnotauthenticatedCancellationdate

`func (o *InscriptionnotauthenticatedResponse) GetDtInscriptionnotauthenticatedCancellationdate() string`

GetDtInscriptionnotauthenticatedCancellationdate returns the DtInscriptionnotauthenticatedCancellationdate field if non-nil, zero value otherwise.

### GetDtInscriptionnotauthenticatedCancellationdateOk

`func (o *InscriptionnotauthenticatedResponse) GetDtInscriptionnotauthenticatedCancellationdateOk() (*string, bool)`

GetDtInscriptionnotauthenticatedCancellationdateOk returns a tuple with the DtInscriptionnotauthenticatedCancellationdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtInscriptionnotauthenticatedCancellationdate

`func (o *InscriptionnotauthenticatedResponse) SetDtInscriptionnotauthenticatedCancellationdate(v string)`

SetDtInscriptionnotauthenticatedCancellationdate sets DtInscriptionnotauthenticatedCancellationdate field to given value.


### GetDtInscriptionnotauthenticatedPossessiondate

`func (o *InscriptionnotauthenticatedResponse) GetDtInscriptionnotauthenticatedPossessiondate() string`

GetDtInscriptionnotauthenticatedPossessiondate returns the DtInscriptionnotauthenticatedPossessiondate field if non-nil, zero value otherwise.

### GetDtInscriptionnotauthenticatedPossessiondateOk

`func (o *InscriptionnotauthenticatedResponse) GetDtInscriptionnotauthenticatedPossessiondateOk() (*string, bool)`

GetDtInscriptionnotauthenticatedPossessiondateOk returns a tuple with the DtInscriptionnotauthenticatedPossessiondate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtInscriptionnotauthenticatedPossessiondate

`func (o *InscriptionnotauthenticatedResponse) SetDtInscriptionnotauthenticatedPossessiondate(v string)`

SetDtInscriptionnotauthenticatedPossessiondate sets DtInscriptionnotauthenticatedPossessiondate field to given value.


### GetSInscriptionnotauthenticatedOffertopurchasenumber

`func (o *InscriptionnotauthenticatedResponse) GetSInscriptionnotauthenticatedOffertopurchasenumber() string`

GetSInscriptionnotauthenticatedOffertopurchasenumber returns the SInscriptionnotauthenticatedOffertopurchasenumber field if non-nil, zero value otherwise.

### GetSInscriptionnotauthenticatedOffertopurchasenumberOk

`func (o *InscriptionnotauthenticatedResponse) GetSInscriptionnotauthenticatedOffertopurchasenumberOk() (*string, bool)`

GetSInscriptionnotauthenticatedOffertopurchasenumberOk returns a tuple with the SInscriptionnotauthenticatedOffertopurchasenumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSInscriptionnotauthenticatedOffertopurchasenumber

`func (o *InscriptionnotauthenticatedResponse) SetSInscriptionnotauthenticatedOffertopurchasenumber(v string)`

SetSInscriptionnotauthenticatedOffertopurchasenumber sets SInscriptionnotauthenticatedOffertopurchasenumber field to given value.


### GetDtInscriptionnotauthenticatedNotaryscheduledate

`func (o *InscriptionnotauthenticatedResponse) GetDtInscriptionnotauthenticatedNotaryscheduledate() string`

GetDtInscriptionnotauthenticatedNotaryscheduledate returns the DtInscriptionnotauthenticatedNotaryscheduledate field if non-nil, zero value otherwise.

### GetDtInscriptionnotauthenticatedNotaryscheduledateOk

`func (o *InscriptionnotauthenticatedResponse) GetDtInscriptionnotauthenticatedNotaryscheduledateOk() (*string, bool)`

GetDtInscriptionnotauthenticatedNotaryscheduledateOk returns a tuple with the DtInscriptionnotauthenticatedNotaryscheduledate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtInscriptionnotauthenticatedNotaryscheduledate

`func (o *InscriptionnotauthenticatedResponse) SetDtInscriptionnotauthenticatedNotaryscheduledate(v string)`

SetDtInscriptionnotauthenticatedNotaryscheduledate sets DtInscriptionnotauthenticatedNotaryscheduledate field to given value.


### GetDtInscriptionnotauthenticatedFinancingscheduledate

`func (o *InscriptionnotauthenticatedResponse) GetDtInscriptionnotauthenticatedFinancingscheduledate() string`

GetDtInscriptionnotauthenticatedFinancingscheduledate returns the DtInscriptionnotauthenticatedFinancingscheduledate field if non-nil, zero value otherwise.

### GetDtInscriptionnotauthenticatedFinancingscheduledateOk

`func (o *InscriptionnotauthenticatedResponse) GetDtInscriptionnotauthenticatedFinancingscheduledateOk() (*string, bool)`

GetDtInscriptionnotauthenticatedFinancingscheduledateOk returns a tuple with the DtInscriptionnotauthenticatedFinancingscheduledate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtInscriptionnotauthenticatedFinancingscheduledate

`func (o *InscriptionnotauthenticatedResponse) SetDtInscriptionnotauthenticatedFinancingscheduledate(v string)`

SetDtInscriptionnotauthenticatedFinancingscheduledate sets DtInscriptionnotauthenticatedFinancingscheduledate field to given value.


### GetBInscriptionnotauthenticatedConditional

`func (o *InscriptionnotauthenticatedResponse) GetBInscriptionnotauthenticatedConditional() bool`

GetBInscriptionnotauthenticatedConditional returns the BInscriptionnotauthenticatedConditional field if non-nil, zero value otherwise.

### GetBInscriptionnotauthenticatedConditionalOk

`func (o *InscriptionnotauthenticatedResponse) GetBInscriptionnotauthenticatedConditionalOk() (*bool, bool)`

GetBInscriptionnotauthenticatedConditionalOk returns a tuple with the BInscriptionnotauthenticatedConditional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBInscriptionnotauthenticatedConditional

`func (o *InscriptionnotauthenticatedResponse) SetBInscriptionnotauthenticatedConditional(v bool)`

SetBInscriptionnotauthenticatedConditional sets BInscriptionnotauthenticatedConditional field to given value.


### GetBInscriptionnotauthenticatedMortgageisreferenced

`func (o *InscriptionnotauthenticatedResponse) GetBInscriptionnotauthenticatedMortgageisreferenced() bool`

GetBInscriptionnotauthenticatedMortgageisreferenced returns the BInscriptionnotauthenticatedMortgageisreferenced field if non-nil, zero value otherwise.

### GetBInscriptionnotauthenticatedMortgageisreferencedOk

`func (o *InscriptionnotauthenticatedResponse) GetBInscriptionnotauthenticatedMortgageisreferencedOk() (*bool, bool)`

GetBInscriptionnotauthenticatedMortgageisreferencedOk returns a tuple with the BInscriptionnotauthenticatedMortgageisreferenced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBInscriptionnotauthenticatedMortgageisreferenced

`func (o *InscriptionnotauthenticatedResponse) SetBInscriptionnotauthenticatedMortgageisreferenced(v bool)`

SetBInscriptionnotauthenticatedMortgageisreferenced sets BInscriptionnotauthenticatedMortgageisreferenced field to given value.


### GetBInscriptionnotauthenticatedHomeowner

`func (o *InscriptionnotauthenticatedResponse) GetBInscriptionnotauthenticatedHomeowner() bool`

GetBInscriptionnotauthenticatedHomeowner returns the BInscriptionnotauthenticatedHomeowner field if non-nil, zero value otherwise.

### GetBInscriptionnotauthenticatedHomeownerOk

`func (o *InscriptionnotauthenticatedResponse) GetBInscriptionnotauthenticatedHomeownerOk() (*bool, bool)`

GetBInscriptionnotauthenticatedHomeownerOk returns a tuple with the BInscriptionnotauthenticatedHomeowner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBInscriptionnotauthenticatedHomeowner

`func (o *InscriptionnotauthenticatedResponse) SetBInscriptionnotauthenticatedHomeowner(v bool)`

SetBInscriptionnotauthenticatedHomeowner sets BInscriptionnotauthenticatedHomeowner field to given value.


### GetTInscriptionnotauthenticatedConditions

`func (o *InscriptionnotauthenticatedResponse) GetTInscriptionnotauthenticatedConditions() string`

GetTInscriptionnotauthenticatedConditions returns the TInscriptionnotauthenticatedConditions field if non-nil, zero value otherwise.

### GetTInscriptionnotauthenticatedConditionsOk

`func (o *InscriptionnotauthenticatedResponse) GetTInscriptionnotauthenticatedConditionsOk() (*string, bool)`

GetTInscriptionnotauthenticatedConditionsOk returns a tuple with the TInscriptionnotauthenticatedConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTInscriptionnotauthenticatedConditions

`func (o *InscriptionnotauthenticatedResponse) SetTInscriptionnotauthenticatedConditions(v string)`

SetTInscriptionnotauthenticatedConditions sets TInscriptionnotauthenticatedConditions field to given value.


### GetDtInscriptionnotauthenticatedConditiondeadlinedate

`func (o *InscriptionnotauthenticatedResponse) GetDtInscriptionnotauthenticatedConditiondeadlinedate() string`

GetDtInscriptionnotauthenticatedConditiondeadlinedate returns the DtInscriptionnotauthenticatedConditiondeadlinedate field if non-nil, zero value otherwise.

### GetDtInscriptionnotauthenticatedConditiondeadlinedateOk

`func (o *InscriptionnotauthenticatedResponse) GetDtInscriptionnotauthenticatedConditiondeadlinedateOk() (*string, bool)`

GetDtInscriptionnotauthenticatedConditiondeadlinedateOk returns a tuple with the DtInscriptionnotauthenticatedConditiondeadlinedate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtInscriptionnotauthenticatedConditiondeadlinedate

`func (o *InscriptionnotauthenticatedResponse) SetDtInscriptionnotauthenticatedConditiondeadlinedate(v string)`

SetDtInscriptionnotauthenticatedConditiondeadlinedate sets DtInscriptionnotauthenticatedConditiondeadlinedate field to given value.


### GetIInscriptionnotauthenticatedOrder

`func (o *InscriptionnotauthenticatedResponse) GetIInscriptionnotauthenticatedOrder() int32`

GetIInscriptionnotauthenticatedOrder returns the IInscriptionnotauthenticatedOrder field if non-nil, zero value otherwise.

### GetIInscriptionnotauthenticatedOrderOk

`func (o *InscriptionnotauthenticatedResponse) GetIInscriptionnotauthenticatedOrderOk() (*int32, bool)`

GetIInscriptionnotauthenticatedOrderOk returns a tuple with the IInscriptionnotauthenticatedOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIInscriptionnotauthenticatedOrder

`func (o *InscriptionnotauthenticatedResponse) SetIInscriptionnotauthenticatedOrder(v int32)`

SetIInscriptionnotauthenticatedOrder sets IInscriptionnotauthenticatedOrder field to given value.


### GetBInscriptionnotauthenticatedIsactive

`func (o *InscriptionnotauthenticatedResponse) GetBInscriptionnotauthenticatedIsactive() bool`

GetBInscriptionnotauthenticatedIsactive returns the BInscriptionnotauthenticatedIsactive field if non-nil, zero value otherwise.

### GetBInscriptionnotauthenticatedIsactiveOk

`func (o *InscriptionnotauthenticatedResponse) GetBInscriptionnotauthenticatedIsactiveOk() (*bool, bool)`

GetBInscriptionnotauthenticatedIsactiveOk returns a tuple with the BInscriptionnotauthenticatedIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBInscriptionnotauthenticatedIsactive

`func (o *InscriptionnotauthenticatedResponse) SetBInscriptionnotauthenticatedIsactive(v bool)`

SetBInscriptionnotauthenticatedIsactive sets BInscriptionnotauthenticatedIsactive field to given value.


### GetEInscriptionnotauthenticatedResidenceType

`func (o *InscriptionnotauthenticatedResponse) GetEInscriptionnotauthenticatedResidenceType() FieldEInscriptionnotauthenticatedResidenceType`

GetEInscriptionnotauthenticatedResidenceType returns the EInscriptionnotauthenticatedResidenceType field if non-nil, zero value otherwise.

### GetEInscriptionnotauthenticatedResidenceTypeOk

`func (o *InscriptionnotauthenticatedResponse) GetEInscriptionnotauthenticatedResidenceTypeOk() (*FieldEInscriptionnotauthenticatedResidenceType, bool)`

GetEInscriptionnotauthenticatedResidenceTypeOk returns a tuple with the EInscriptionnotauthenticatedResidenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEInscriptionnotauthenticatedResidenceType

`func (o *InscriptionnotauthenticatedResponse) SetEInscriptionnotauthenticatedResidenceType(v FieldEInscriptionnotauthenticatedResidenceType)`

SetEInscriptionnotauthenticatedResidenceType sets EInscriptionnotauthenticatedResidenceType field to given value.


### GetTInscriptionnotauthenticatedChecklistnote

`func (o *InscriptionnotauthenticatedResponse) GetTInscriptionnotauthenticatedChecklistnote() string`

GetTInscriptionnotauthenticatedChecklistnote returns the TInscriptionnotauthenticatedChecklistnote field if non-nil, zero value otherwise.

### GetTInscriptionnotauthenticatedChecklistnoteOk

`func (o *InscriptionnotauthenticatedResponse) GetTInscriptionnotauthenticatedChecklistnoteOk() (*string, bool)`

GetTInscriptionnotauthenticatedChecklistnoteOk returns a tuple with the TInscriptionnotauthenticatedChecklistnote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTInscriptionnotauthenticatedChecklistnote

`func (o *InscriptionnotauthenticatedResponse) SetTInscriptionnotauthenticatedChecklistnote(v string)`

SetTInscriptionnotauthenticatedChecklistnote sets TInscriptionnotauthenticatedChecklistnote field to given value.


### GetDInscriptionnotauthenticatedSelleronlyretribution

`func (o *InscriptionnotauthenticatedResponse) GetDInscriptionnotauthenticatedSelleronlyretribution() string`

GetDInscriptionnotauthenticatedSelleronlyretribution returns the DInscriptionnotauthenticatedSelleronlyretribution field if non-nil, zero value otherwise.

### GetDInscriptionnotauthenticatedSelleronlyretributionOk

`func (o *InscriptionnotauthenticatedResponse) GetDInscriptionnotauthenticatedSelleronlyretributionOk() (*string, bool)`

GetDInscriptionnotauthenticatedSelleronlyretributionOk returns a tuple with the DInscriptionnotauthenticatedSelleronlyretribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDInscriptionnotauthenticatedSelleronlyretribution

`func (o *InscriptionnotauthenticatedResponse) SetDInscriptionnotauthenticatedSelleronlyretribution(v string)`

SetDInscriptionnotauthenticatedSelleronlyretribution sets DInscriptionnotauthenticatedSelleronlyretribution field to given value.


### GetBInscriptionnotauthenticatedDraft

`func (o *InscriptionnotauthenticatedResponse) GetBInscriptionnotauthenticatedDraft() bool`

GetBInscriptionnotauthenticatedDraft returns the BInscriptionnotauthenticatedDraft field if non-nil, zero value otherwise.

### GetBInscriptionnotauthenticatedDraftOk

`func (o *InscriptionnotauthenticatedResponse) GetBInscriptionnotauthenticatedDraftOk() (*bool, bool)`

GetBInscriptionnotauthenticatedDraftOk returns a tuple with the BInscriptionnotauthenticatedDraft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBInscriptionnotauthenticatedDraft

`func (o *InscriptionnotauthenticatedResponse) SetBInscriptionnotauthenticatedDraft(v bool)`

SetBInscriptionnotauthenticatedDraft sets BInscriptionnotauthenticatedDraft field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


