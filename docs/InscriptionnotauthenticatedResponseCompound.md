# InscriptionnotauthenticatedResponseCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiInscriptionnotauthenticatedID** | **int32** | The unique ID of the Inscriptionnotauthenticated. | 
**FkiInscriptionID** | **int32** | The unique ID of the Inscription. | 
**FkiDepartmentID** | Pointer to **int32** | The unique ID of the Department | [optional] 
**SDepartmentNameX** | Pointer to **string** | The Name of the Department in the language of the requester | [optional] 
**FkiFinancialinstitutionID** | Pointer to **int32** | The unique ID of the Financialinstitution | [optional] 
**SFinancialinstitutionNameX** | Pointer to **string** | The name of the Financialinstitution in the language of the requester | [optional] 
**FkiBuyercontractID** | Pointer to **int32** | The unique ID of the Buyercontract | [optional] 
**SBuyercontractContract** | Pointer to **string** | The number of the Buyercontract | [optional] 
**FkiMortgagesupplierID** | Pointer to **int32** | The unique ID of the Mortgagesupplier | [optional] 
**SMortgagesupplierNameX** | Pointer to **string** | The name of the Mortagesupplier in the language of the requester | [optional] 
**FkiTaxassignmentID** | **int32** | The unique ID of the Taxassignment.  Valid values:  |Value|Description| |-|-| |1|No tax| |2|GST| |3|HST (ON)| |4|HST (NB)| |5|HST (NS)| |6|HST (NL)| |7|HST (PE)| |8|GST + QST (QC)| |9|GST + QST (QC) Non-Recoverable| |10|GST + PST (BC)| |11|GST + PST (SK)| |12|GST + RST (MB)| |13|GST + PST (BC) Non-Recoverable| |14|GST + PST (SK) Non-Recoverable| |15|GST + RST (MB) Non-Recoverable| | 
**STaxassignmentDescriptionX** | Pointer to **string** | The description of the Taxassignment  in the language of the requester | [optional] 
**DtInscriptionnotauthenticatedTransactiondate** | Pointer to **string** | The transaction date of the Inscriptionnotauthenticated | [optional] 
**DtInscriptionnotauthenticatedTransactiondateReal** | Pointer to **string** | The real transactiondate of the Inscriptionnotauthenticated | [optional] 
**DtInscriptionnotauthenticatedDepositdate** | Pointer to **string** | The deposit date of the Inscriptionnotauthenticated | [optional] 
**EInscriptionnotauthenticatedType** | [**FieldEInscriptionnotauthenticatedType**](FieldEInscriptionnotauthenticatedType.md) |  | 
**DInscriptionnotauthenticatedMortgageloan** | **string** | The amount of the mortgage loan of the Inscriptionnotauthenticated | 
**EtInscriptionnotauthenticatedMortgagetype** | [**FieldEtInscriptionnotauthenticatedMortgagetype**](FieldEtInscriptionnotauthenticatedMortgagetype.md) |  | 
**DInscriptionnotauthenticatedTransactionprice** | **string** | The transaction price of the Inscriptionnotauthenticated | 
**EInscriptionnotauthenticatedRemunerationtype** | [**FieldEInscriptionnotauthenticatedRemunerationtype**](FieldEInscriptionnotauthenticatedRemunerationtype.md) |  | 
**DInscriptionnotauthenticatedRemuneration** | **string** | The amount for the remuneration of the Inscriptionnotauthenticated | 
**DInscriptionnotauthenticatedRemunerationsubtotal** | **string** | The subtotal for the remuneration of the Inscriptionnotauthenticated | 
**DInscriptionnotauthenticatedRemunerationtotal** | **string** | The total for the remuneration of the Inscriptionnotauthenticated | 
**DtInscriptionnotauthenticatedCancellationdate** | Pointer to **string** | The cancellation date of the Inscriptionnotauthenticated | [optional] 
**DtInscriptionnotauthenticatedPossessiondate** | Pointer to **string** | The possession date of the Inscriptionnotauthenticated | [optional] 
**SInscriptionnotauthenticatedOffertopurchasenumber** | **string** | The offer to purchase number of the Inscriptionnotauthenticated | 
**DtInscriptionnotauthenticatedNotaryscheduledate** | Pointer to **string** | The notary schedule date of the Inscriptionnotauthenticated | [optional] 
**DtInscriptionnotauthenticatedFinancingscheduledate** | Pointer to **string** | The financing schedule date of the Inscriptionnotauthenticated | [optional] 
**BInscriptionnotauthenticatedConditional** | **bool** | Whether the inscriptionnotauthenticated is conditional | 
**BInscriptionnotauthenticatedMortgageisreferenced** | **bool** | Whether if the mortgage is referenced | 
**BInscriptionnotauthenticatedHomeowner** | **bool** | Whether if it&#39;s an home owner | 
**TInscriptionnotauthenticatedConditions** | **string** | The conditions of the Inscriptionnotauthenticated | 
**DtInscriptionnotauthenticatedConditiondeadlinedate** | Pointer to **string** | The condition deadline date of the Inscriptionnotauthenticated | [optional] 
**IInscriptionnotauthenticatedOrder** | **int32** | The order of the Inscriptionnotauthenticated | 
**BInscriptionnotauthenticatedIsactive** | **bool** | Whether the inscriptionnotauthenticated is active or not | 
**EInscriptionnotauthenticatedResidenceType** | [**FieldEInscriptionnotauthenticatedResidenceType**](FieldEInscriptionnotauthenticatedResidenceType.md) |  | 
**TInscriptionnotauthenticatedChecklistnote** | **string** | The checklist note of the Inscriptionnotauthenticated | 
**DInscriptionnotauthenticatedSelleronlyretribution** | **string** | The amount retribution for the seller only of the Inscriptionnotauthenticated | 
**BInscriptionnotauthenticatedDraft** | **bool** | Whether the Inscriptionnotauthenticated is a draft or not | 

## Methods

### NewInscriptionnotauthenticatedResponseCompound

`func NewInscriptionnotauthenticatedResponseCompound(pkiInscriptionnotauthenticatedID int32, fkiInscriptionID int32, fkiTaxassignmentID int32, eInscriptionnotauthenticatedType FieldEInscriptionnotauthenticatedType, dInscriptionnotauthenticatedMortgageloan string, etInscriptionnotauthenticatedMortgagetype FieldEtInscriptionnotauthenticatedMortgagetype, dInscriptionnotauthenticatedTransactionprice string, eInscriptionnotauthenticatedRemunerationtype FieldEInscriptionnotauthenticatedRemunerationtype, dInscriptionnotauthenticatedRemuneration string, dInscriptionnotauthenticatedRemunerationsubtotal string, dInscriptionnotauthenticatedRemunerationtotal string, sInscriptionnotauthenticatedOffertopurchasenumber string, bInscriptionnotauthenticatedConditional bool, bInscriptionnotauthenticatedMortgageisreferenced bool, bInscriptionnotauthenticatedHomeowner bool, tInscriptionnotauthenticatedConditions string, iInscriptionnotauthenticatedOrder int32, bInscriptionnotauthenticatedIsactive bool, eInscriptionnotauthenticatedResidenceType FieldEInscriptionnotauthenticatedResidenceType, tInscriptionnotauthenticatedChecklistnote string, dInscriptionnotauthenticatedSelleronlyretribution string, bInscriptionnotauthenticatedDraft bool, ) *InscriptionnotauthenticatedResponseCompound`

NewInscriptionnotauthenticatedResponseCompound instantiates a new InscriptionnotauthenticatedResponseCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInscriptionnotauthenticatedResponseCompoundWithDefaults

`func NewInscriptionnotauthenticatedResponseCompoundWithDefaults() *InscriptionnotauthenticatedResponseCompound`

NewInscriptionnotauthenticatedResponseCompoundWithDefaults instantiates a new InscriptionnotauthenticatedResponseCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiInscriptionnotauthenticatedID

`func (o *InscriptionnotauthenticatedResponseCompound) GetPkiInscriptionnotauthenticatedID() int32`

GetPkiInscriptionnotauthenticatedID returns the PkiInscriptionnotauthenticatedID field if non-nil, zero value otherwise.

### GetPkiInscriptionnotauthenticatedIDOk

`func (o *InscriptionnotauthenticatedResponseCompound) GetPkiInscriptionnotauthenticatedIDOk() (*int32, bool)`

GetPkiInscriptionnotauthenticatedIDOk returns a tuple with the PkiInscriptionnotauthenticatedID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiInscriptionnotauthenticatedID

`func (o *InscriptionnotauthenticatedResponseCompound) SetPkiInscriptionnotauthenticatedID(v int32)`

SetPkiInscriptionnotauthenticatedID sets PkiInscriptionnotauthenticatedID field to given value.


### GetFkiInscriptionID

`func (o *InscriptionnotauthenticatedResponseCompound) GetFkiInscriptionID() int32`

GetFkiInscriptionID returns the FkiInscriptionID field if non-nil, zero value otherwise.

### GetFkiInscriptionIDOk

`func (o *InscriptionnotauthenticatedResponseCompound) GetFkiInscriptionIDOk() (*int32, bool)`

GetFkiInscriptionIDOk returns a tuple with the FkiInscriptionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiInscriptionID

`func (o *InscriptionnotauthenticatedResponseCompound) SetFkiInscriptionID(v int32)`

SetFkiInscriptionID sets FkiInscriptionID field to given value.


### GetFkiDepartmentID

`func (o *InscriptionnotauthenticatedResponseCompound) GetFkiDepartmentID() int32`

GetFkiDepartmentID returns the FkiDepartmentID field if non-nil, zero value otherwise.

### GetFkiDepartmentIDOk

`func (o *InscriptionnotauthenticatedResponseCompound) GetFkiDepartmentIDOk() (*int32, bool)`

GetFkiDepartmentIDOk returns a tuple with the FkiDepartmentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiDepartmentID

`func (o *InscriptionnotauthenticatedResponseCompound) SetFkiDepartmentID(v int32)`

SetFkiDepartmentID sets FkiDepartmentID field to given value.

### HasFkiDepartmentID

`func (o *InscriptionnotauthenticatedResponseCompound) HasFkiDepartmentID() bool`

HasFkiDepartmentID returns a boolean if a field has been set.

### GetSDepartmentNameX

`func (o *InscriptionnotauthenticatedResponseCompound) GetSDepartmentNameX() string`

GetSDepartmentNameX returns the SDepartmentNameX field if non-nil, zero value otherwise.

### GetSDepartmentNameXOk

`func (o *InscriptionnotauthenticatedResponseCompound) GetSDepartmentNameXOk() (*string, bool)`

GetSDepartmentNameXOk returns a tuple with the SDepartmentNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSDepartmentNameX

`func (o *InscriptionnotauthenticatedResponseCompound) SetSDepartmentNameX(v string)`

SetSDepartmentNameX sets SDepartmentNameX field to given value.

### HasSDepartmentNameX

`func (o *InscriptionnotauthenticatedResponseCompound) HasSDepartmentNameX() bool`

HasSDepartmentNameX returns a boolean if a field has been set.

### GetFkiFinancialinstitutionID

`func (o *InscriptionnotauthenticatedResponseCompound) GetFkiFinancialinstitutionID() int32`

GetFkiFinancialinstitutionID returns the FkiFinancialinstitutionID field if non-nil, zero value otherwise.

### GetFkiFinancialinstitutionIDOk

`func (o *InscriptionnotauthenticatedResponseCompound) GetFkiFinancialinstitutionIDOk() (*int32, bool)`

GetFkiFinancialinstitutionIDOk returns a tuple with the FkiFinancialinstitutionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiFinancialinstitutionID

`func (o *InscriptionnotauthenticatedResponseCompound) SetFkiFinancialinstitutionID(v int32)`

SetFkiFinancialinstitutionID sets FkiFinancialinstitutionID field to given value.

### HasFkiFinancialinstitutionID

`func (o *InscriptionnotauthenticatedResponseCompound) HasFkiFinancialinstitutionID() bool`

HasFkiFinancialinstitutionID returns a boolean if a field has been set.

### GetSFinancialinstitutionNameX

`func (o *InscriptionnotauthenticatedResponseCompound) GetSFinancialinstitutionNameX() string`

GetSFinancialinstitutionNameX returns the SFinancialinstitutionNameX field if non-nil, zero value otherwise.

### GetSFinancialinstitutionNameXOk

`func (o *InscriptionnotauthenticatedResponseCompound) GetSFinancialinstitutionNameXOk() (*string, bool)`

GetSFinancialinstitutionNameXOk returns a tuple with the SFinancialinstitutionNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSFinancialinstitutionNameX

`func (o *InscriptionnotauthenticatedResponseCompound) SetSFinancialinstitutionNameX(v string)`

SetSFinancialinstitutionNameX sets SFinancialinstitutionNameX field to given value.

### HasSFinancialinstitutionNameX

`func (o *InscriptionnotauthenticatedResponseCompound) HasSFinancialinstitutionNameX() bool`

HasSFinancialinstitutionNameX returns a boolean if a field has been set.

### GetFkiBuyercontractID

`func (o *InscriptionnotauthenticatedResponseCompound) GetFkiBuyercontractID() int32`

GetFkiBuyercontractID returns the FkiBuyercontractID field if non-nil, zero value otherwise.

### GetFkiBuyercontractIDOk

`func (o *InscriptionnotauthenticatedResponseCompound) GetFkiBuyercontractIDOk() (*int32, bool)`

GetFkiBuyercontractIDOk returns a tuple with the FkiBuyercontractID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiBuyercontractID

`func (o *InscriptionnotauthenticatedResponseCompound) SetFkiBuyercontractID(v int32)`

SetFkiBuyercontractID sets FkiBuyercontractID field to given value.

### HasFkiBuyercontractID

`func (o *InscriptionnotauthenticatedResponseCompound) HasFkiBuyercontractID() bool`

HasFkiBuyercontractID returns a boolean if a field has been set.

### GetSBuyercontractContract

`func (o *InscriptionnotauthenticatedResponseCompound) GetSBuyercontractContract() string`

GetSBuyercontractContract returns the SBuyercontractContract field if non-nil, zero value otherwise.

### GetSBuyercontractContractOk

`func (o *InscriptionnotauthenticatedResponseCompound) GetSBuyercontractContractOk() (*string, bool)`

GetSBuyercontractContractOk returns a tuple with the SBuyercontractContract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSBuyercontractContract

`func (o *InscriptionnotauthenticatedResponseCompound) SetSBuyercontractContract(v string)`

SetSBuyercontractContract sets SBuyercontractContract field to given value.

### HasSBuyercontractContract

`func (o *InscriptionnotauthenticatedResponseCompound) HasSBuyercontractContract() bool`

HasSBuyercontractContract returns a boolean if a field has been set.

### GetFkiMortgagesupplierID

`func (o *InscriptionnotauthenticatedResponseCompound) GetFkiMortgagesupplierID() int32`

GetFkiMortgagesupplierID returns the FkiMortgagesupplierID field if non-nil, zero value otherwise.

### GetFkiMortgagesupplierIDOk

`func (o *InscriptionnotauthenticatedResponseCompound) GetFkiMortgagesupplierIDOk() (*int32, bool)`

GetFkiMortgagesupplierIDOk returns a tuple with the FkiMortgagesupplierID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiMortgagesupplierID

`func (o *InscriptionnotauthenticatedResponseCompound) SetFkiMortgagesupplierID(v int32)`

SetFkiMortgagesupplierID sets FkiMortgagesupplierID field to given value.

### HasFkiMortgagesupplierID

`func (o *InscriptionnotauthenticatedResponseCompound) HasFkiMortgagesupplierID() bool`

HasFkiMortgagesupplierID returns a boolean if a field has been set.

### GetSMortgagesupplierNameX

`func (o *InscriptionnotauthenticatedResponseCompound) GetSMortgagesupplierNameX() string`

GetSMortgagesupplierNameX returns the SMortgagesupplierNameX field if non-nil, zero value otherwise.

### GetSMortgagesupplierNameXOk

`func (o *InscriptionnotauthenticatedResponseCompound) GetSMortgagesupplierNameXOk() (*string, bool)`

GetSMortgagesupplierNameXOk returns a tuple with the SMortgagesupplierNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSMortgagesupplierNameX

`func (o *InscriptionnotauthenticatedResponseCompound) SetSMortgagesupplierNameX(v string)`

SetSMortgagesupplierNameX sets SMortgagesupplierNameX field to given value.

### HasSMortgagesupplierNameX

`func (o *InscriptionnotauthenticatedResponseCompound) HasSMortgagesupplierNameX() bool`

HasSMortgagesupplierNameX returns a boolean if a field has been set.

### GetFkiTaxassignmentID

`func (o *InscriptionnotauthenticatedResponseCompound) GetFkiTaxassignmentID() int32`

GetFkiTaxassignmentID returns the FkiTaxassignmentID field if non-nil, zero value otherwise.

### GetFkiTaxassignmentIDOk

`func (o *InscriptionnotauthenticatedResponseCompound) GetFkiTaxassignmentIDOk() (*int32, bool)`

GetFkiTaxassignmentIDOk returns a tuple with the FkiTaxassignmentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiTaxassignmentID

`func (o *InscriptionnotauthenticatedResponseCompound) SetFkiTaxassignmentID(v int32)`

SetFkiTaxassignmentID sets FkiTaxassignmentID field to given value.


### GetSTaxassignmentDescriptionX

`func (o *InscriptionnotauthenticatedResponseCompound) GetSTaxassignmentDescriptionX() string`

GetSTaxassignmentDescriptionX returns the STaxassignmentDescriptionX field if non-nil, zero value otherwise.

### GetSTaxassignmentDescriptionXOk

`func (o *InscriptionnotauthenticatedResponseCompound) GetSTaxassignmentDescriptionXOk() (*string, bool)`

GetSTaxassignmentDescriptionXOk returns a tuple with the STaxassignmentDescriptionX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSTaxassignmentDescriptionX

`func (o *InscriptionnotauthenticatedResponseCompound) SetSTaxassignmentDescriptionX(v string)`

SetSTaxassignmentDescriptionX sets STaxassignmentDescriptionX field to given value.

### HasSTaxassignmentDescriptionX

`func (o *InscriptionnotauthenticatedResponseCompound) HasSTaxassignmentDescriptionX() bool`

HasSTaxassignmentDescriptionX returns a boolean if a field has been set.

### GetDtInscriptionnotauthenticatedTransactiondate

`func (o *InscriptionnotauthenticatedResponseCompound) GetDtInscriptionnotauthenticatedTransactiondate() string`

GetDtInscriptionnotauthenticatedTransactiondate returns the DtInscriptionnotauthenticatedTransactiondate field if non-nil, zero value otherwise.

### GetDtInscriptionnotauthenticatedTransactiondateOk

`func (o *InscriptionnotauthenticatedResponseCompound) GetDtInscriptionnotauthenticatedTransactiondateOk() (*string, bool)`

GetDtInscriptionnotauthenticatedTransactiondateOk returns a tuple with the DtInscriptionnotauthenticatedTransactiondate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtInscriptionnotauthenticatedTransactiondate

`func (o *InscriptionnotauthenticatedResponseCompound) SetDtInscriptionnotauthenticatedTransactiondate(v string)`

SetDtInscriptionnotauthenticatedTransactiondate sets DtInscriptionnotauthenticatedTransactiondate field to given value.

### HasDtInscriptionnotauthenticatedTransactiondate

`func (o *InscriptionnotauthenticatedResponseCompound) HasDtInscriptionnotauthenticatedTransactiondate() bool`

HasDtInscriptionnotauthenticatedTransactiondate returns a boolean if a field has been set.

### GetDtInscriptionnotauthenticatedTransactiondateReal

`func (o *InscriptionnotauthenticatedResponseCompound) GetDtInscriptionnotauthenticatedTransactiondateReal() string`

GetDtInscriptionnotauthenticatedTransactiondateReal returns the DtInscriptionnotauthenticatedTransactiondateReal field if non-nil, zero value otherwise.

### GetDtInscriptionnotauthenticatedTransactiondateRealOk

`func (o *InscriptionnotauthenticatedResponseCompound) GetDtInscriptionnotauthenticatedTransactiondateRealOk() (*string, bool)`

GetDtInscriptionnotauthenticatedTransactiondateRealOk returns a tuple with the DtInscriptionnotauthenticatedTransactiondateReal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtInscriptionnotauthenticatedTransactiondateReal

`func (o *InscriptionnotauthenticatedResponseCompound) SetDtInscriptionnotauthenticatedTransactiondateReal(v string)`

SetDtInscriptionnotauthenticatedTransactiondateReal sets DtInscriptionnotauthenticatedTransactiondateReal field to given value.

### HasDtInscriptionnotauthenticatedTransactiondateReal

`func (o *InscriptionnotauthenticatedResponseCompound) HasDtInscriptionnotauthenticatedTransactiondateReal() bool`

HasDtInscriptionnotauthenticatedTransactiondateReal returns a boolean if a field has been set.

### GetDtInscriptionnotauthenticatedDepositdate

`func (o *InscriptionnotauthenticatedResponseCompound) GetDtInscriptionnotauthenticatedDepositdate() string`

GetDtInscriptionnotauthenticatedDepositdate returns the DtInscriptionnotauthenticatedDepositdate field if non-nil, zero value otherwise.

### GetDtInscriptionnotauthenticatedDepositdateOk

`func (o *InscriptionnotauthenticatedResponseCompound) GetDtInscriptionnotauthenticatedDepositdateOk() (*string, bool)`

GetDtInscriptionnotauthenticatedDepositdateOk returns a tuple with the DtInscriptionnotauthenticatedDepositdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtInscriptionnotauthenticatedDepositdate

`func (o *InscriptionnotauthenticatedResponseCompound) SetDtInscriptionnotauthenticatedDepositdate(v string)`

SetDtInscriptionnotauthenticatedDepositdate sets DtInscriptionnotauthenticatedDepositdate field to given value.

### HasDtInscriptionnotauthenticatedDepositdate

`func (o *InscriptionnotauthenticatedResponseCompound) HasDtInscriptionnotauthenticatedDepositdate() bool`

HasDtInscriptionnotauthenticatedDepositdate returns a boolean if a field has been set.

### GetEInscriptionnotauthenticatedType

`func (o *InscriptionnotauthenticatedResponseCompound) GetEInscriptionnotauthenticatedType() FieldEInscriptionnotauthenticatedType`

GetEInscriptionnotauthenticatedType returns the EInscriptionnotauthenticatedType field if non-nil, zero value otherwise.

### GetEInscriptionnotauthenticatedTypeOk

`func (o *InscriptionnotauthenticatedResponseCompound) GetEInscriptionnotauthenticatedTypeOk() (*FieldEInscriptionnotauthenticatedType, bool)`

GetEInscriptionnotauthenticatedTypeOk returns a tuple with the EInscriptionnotauthenticatedType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEInscriptionnotauthenticatedType

`func (o *InscriptionnotauthenticatedResponseCompound) SetEInscriptionnotauthenticatedType(v FieldEInscriptionnotauthenticatedType)`

SetEInscriptionnotauthenticatedType sets EInscriptionnotauthenticatedType field to given value.


### GetDInscriptionnotauthenticatedMortgageloan

`func (o *InscriptionnotauthenticatedResponseCompound) GetDInscriptionnotauthenticatedMortgageloan() string`

GetDInscriptionnotauthenticatedMortgageloan returns the DInscriptionnotauthenticatedMortgageloan field if non-nil, zero value otherwise.

### GetDInscriptionnotauthenticatedMortgageloanOk

`func (o *InscriptionnotauthenticatedResponseCompound) GetDInscriptionnotauthenticatedMortgageloanOk() (*string, bool)`

GetDInscriptionnotauthenticatedMortgageloanOk returns a tuple with the DInscriptionnotauthenticatedMortgageloan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDInscriptionnotauthenticatedMortgageloan

`func (o *InscriptionnotauthenticatedResponseCompound) SetDInscriptionnotauthenticatedMortgageloan(v string)`

SetDInscriptionnotauthenticatedMortgageloan sets DInscriptionnotauthenticatedMortgageloan field to given value.


### GetEtInscriptionnotauthenticatedMortgagetype

`func (o *InscriptionnotauthenticatedResponseCompound) GetEtInscriptionnotauthenticatedMortgagetype() FieldEtInscriptionnotauthenticatedMortgagetype`

GetEtInscriptionnotauthenticatedMortgagetype returns the EtInscriptionnotauthenticatedMortgagetype field if non-nil, zero value otherwise.

### GetEtInscriptionnotauthenticatedMortgagetypeOk

`func (o *InscriptionnotauthenticatedResponseCompound) GetEtInscriptionnotauthenticatedMortgagetypeOk() (*FieldEtInscriptionnotauthenticatedMortgagetype, bool)`

GetEtInscriptionnotauthenticatedMortgagetypeOk returns a tuple with the EtInscriptionnotauthenticatedMortgagetype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtInscriptionnotauthenticatedMortgagetype

`func (o *InscriptionnotauthenticatedResponseCompound) SetEtInscriptionnotauthenticatedMortgagetype(v FieldEtInscriptionnotauthenticatedMortgagetype)`

SetEtInscriptionnotauthenticatedMortgagetype sets EtInscriptionnotauthenticatedMortgagetype field to given value.


### GetDInscriptionnotauthenticatedTransactionprice

`func (o *InscriptionnotauthenticatedResponseCompound) GetDInscriptionnotauthenticatedTransactionprice() string`

GetDInscriptionnotauthenticatedTransactionprice returns the DInscriptionnotauthenticatedTransactionprice field if non-nil, zero value otherwise.

### GetDInscriptionnotauthenticatedTransactionpriceOk

`func (o *InscriptionnotauthenticatedResponseCompound) GetDInscriptionnotauthenticatedTransactionpriceOk() (*string, bool)`

GetDInscriptionnotauthenticatedTransactionpriceOk returns a tuple with the DInscriptionnotauthenticatedTransactionprice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDInscriptionnotauthenticatedTransactionprice

`func (o *InscriptionnotauthenticatedResponseCompound) SetDInscriptionnotauthenticatedTransactionprice(v string)`

SetDInscriptionnotauthenticatedTransactionprice sets DInscriptionnotauthenticatedTransactionprice field to given value.


### GetEInscriptionnotauthenticatedRemunerationtype

`func (o *InscriptionnotauthenticatedResponseCompound) GetEInscriptionnotauthenticatedRemunerationtype() FieldEInscriptionnotauthenticatedRemunerationtype`

GetEInscriptionnotauthenticatedRemunerationtype returns the EInscriptionnotauthenticatedRemunerationtype field if non-nil, zero value otherwise.

### GetEInscriptionnotauthenticatedRemunerationtypeOk

`func (o *InscriptionnotauthenticatedResponseCompound) GetEInscriptionnotauthenticatedRemunerationtypeOk() (*FieldEInscriptionnotauthenticatedRemunerationtype, bool)`

GetEInscriptionnotauthenticatedRemunerationtypeOk returns a tuple with the EInscriptionnotauthenticatedRemunerationtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEInscriptionnotauthenticatedRemunerationtype

`func (o *InscriptionnotauthenticatedResponseCompound) SetEInscriptionnotauthenticatedRemunerationtype(v FieldEInscriptionnotauthenticatedRemunerationtype)`

SetEInscriptionnotauthenticatedRemunerationtype sets EInscriptionnotauthenticatedRemunerationtype field to given value.


### GetDInscriptionnotauthenticatedRemuneration

`func (o *InscriptionnotauthenticatedResponseCompound) GetDInscriptionnotauthenticatedRemuneration() string`

GetDInscriptionnotauthenticatedRemuneration returns the DInscriptionnotauthenticatedRemuneration field if non-nil, zero value otherwise.

### GetDInscriptionnotauthenticatedRemunerationOk

`func (o *InscriptionnotauthenticatedResponseCompound) GetDInscriptionnotauthenticatedRemunerationOk() (*string, bool)`

GetDInscriptionnotauthenticatedRemunerationOk returns a tuple with the DInscriptionnotauthenticatedRemuneration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDInscriptionnotauthenticatedRemuneration

`func (o *InscriptionnotauthenticatedResponseCompound) SetDInscriptionnotauthenticatedRemuneration(v string)`

SetDInscriptionnotauthenticatedRemuneration sets DInscriptionnotauthenticatedRemuneration field to given value.


### GetDInscriptionnotauthenticatedRemunerationsubtotal

`func (o *InscriptionnotauthenticatedResponseCompound) GetDInscriptionnotauthenticatedRemunerationsubtotal() string`

GetDInscriptionnotauthenticatedRemunerationsubtotal returns the DInscriptionnotauthenticatedRemunerationsubtotal field if non-nil, zero value otherwise.

### GetDInscriptionnotauthenticatedRemunerationsubtotalOk

`func (o *InscriptionnotauthenticatedResponseCompound) GetDInscriptionnotauthenticatedRemunerationsubtotalOk() (*string, bool)`

GetDInscriptionnotauthenticatedRemunerationsubtotalOk returns a tuple with the DInscriptionnotauthenticatedRemunerationsubtotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDInscriptionnotauthenticatedRemunerationsubtotal

`func (o *InscriptionnotauthenticatedResponseCompound) SetDInscriptionnotauthenticatedRemunerationsubtotal(v string)`

SetDInscriptionnotauthenticatedRemunerationsubtotal sets DInscriptionnotauthenticatedRemunerationsubtotal field to given value.


### GetDInscriptionnotauthenticatedRemunerationtotal

`func (o *InscriptionnotauthenticatedResponseCompound) GetDInscriptionnotauthenticatedRemunerationtotal() string`

GetDInscriptionnotauthenticatedRemunerationtotal returns the DInscriptionnotauthenticatedRemunerationtotal field if non-nil, zero value otherwise.

### GetDInscriptionnotauthenticatedRemunerationtotalOk

`func (o *InscriptionnotauthenticatedResponseCompound) GetDInscriptionnotauthenticatedRemunerationtotalOk() (*string, bool)`

GetDInscriptionnotauthenticatedRemunerationtotalOk returns a tuple with the DInscriptionnotauthenticatedRemunerationtotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDInscriptionnotauthenticatedRemunerationtotal

`func (o *InscriptionnotauthenticatedResponseCompound) SetDInscriptionnotauthenticatedRemunerationtotal(v string)`

SetDInscriptionnotauthenticatedRemunerationtotal sets DInscriptionnotauthenticatedRemunerationtotal field to given value.


### GetDtInscriptionnotauthenticatedCancellationdate

`func (o *InscriptionnotauthenticatedResponseCompound) GetDtInscriptionnotauthenticatedCancellationdate() string`

GetDtInscriptionnotauthenticatedCancellationdate returns the DtInscriptionnotauthenticatedCancellationdate field if non-nil, zero value otherwise.

### GetDtInscriptionnotauthenticatedCancellationdateOk

`func (o *InscriptionnotauthenticatedResponseCompound) GetDtInscriptionnotauthenticatedCancellationdateOk() (*string, bool)`

GetDtInscriptionnotauthenticatedCancellationdateOk returns a tuple with the DtInscriptionnotauthenticatedCancellationdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtInscriptionnotauthenticatedCancellationdate

`func (o *InscriptionnotauthenticatedResponseCompound) SetDtInscriptionnotauthenticatedCancellationdate(v string)`

SetDtInscriptionnotauthenticatedCancellationdate sets DtInscriptionnotauthenticatedCancellationdate field to given value.

### HasDtInscriptionnotauthenticatedCancellationdate

`func (o *InscriptionnotauthenticatedResponseCompound) HasDtInscriptionnotauthenticatedCancellationdate() bool`

HasDtInscriptionnotauthenticatedCancellationdate returns a boolean if a field has been set.

### GetDtInscriptionnotauthenticatedPossessiondate

`func (o *InscriptionnotauthenticatedResponseCompound) GetDtInscriptionnotauthenticatedPossessiondate() string`

GetDtInscriptionnotauthenticatedPossessiondate returns the DtInscriptionnotauthenticatedPossessiondate field if non-nil, zero value otherwise.

### GetDtInscriptionnotauthenticatedPossessiondateOk

`func (o *InscriptionnotauthenticatedResponseCompound) GetDtInscriptionnotauthenticatedPossessiondateOk() (*string, bool)`

GetDtInscriptionnotauthenticatedPossessiondateOk returns a tuple with the DtInscriptionnotauthenticatedPossessiondate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtInscriptionnotauthenticatedPossessiondate

`func (o *InscriptionnotauthenticatedResponseCompound) SetDtInscriptionnotauthenticatedPossessiondate(v string)`

SetDtInscriptionnotauthenticatedPossessiondate sets DtInscriptionnotauthenticatedPossessiondate field to given value.

### HasDtInscriptionnotauthenticatedPossessiondate

`func (o *InscriptionnotauthenticatedResponseCompound) HasDtInscriptionnotauthenticatedPossessiondate() bool`

HasDtInscriptionnotauthenticatedPossessiondate returns a boolean if a field has been set.

### GetSInscriptionnotauthenticatedOffertopurchasenumber

`func (o *InscriptionnotauthenticatedResponseCompound) GetSInscriptionnotauthenticatedOffertopurchasenumber() string`

GetSInscriptionnotauthenticatedOffertopurchasenumber returns the SInscriptionnotauthenticatedOffertopurchasenumber field if non-nil, zero value otherwise.

### GetSInscriptionnotauthenticatedOffertopurchasenumberOk

`func (o *InscriptionnotauthenticatedResponseCompound) GetSInscriptionnotauthenticatedOffertopurchasenumberOk() (*string, bool)`

GetSInscriptionnotauthenticatedOffertopurchasenumberOk returns a tuple with the SInscriptionnotauthenticatedOffertopurchasenumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSInscriptionnotauthenticatedOffertopurchasenumber

`func (o *InscriptionnotauthenticatedResponseCompound) SetSInscriptionnotauthenticatedOffertopurchasenumber(v string)`

SetSInscriptionnotauthenticatedOffertopurchasenumber sets SInscriptionnotauthenticatedOffertopurchasenumber field to given value.


### GetDtInscriptionnotauthenticatedNotaryscheduledate

`func (o *InscriptionnotauthenticatedResponseCompound) GetDtInscriptionnotauthenticatedNotaryscheduledate() string`

GetDtInscriptionnotauthenticatedNotaryscheduledate returns the DtInscriptionnotauthenticatedNotaryscheduledate field if non-nil, zero value otherwise.

### GetDtInscriptionnotauthenticatedNotaryscheduledateOk

`func (o *InscriptionnotauthenticatedResponseCompound) GetDtInscriptionnotauthenticatedNotaryscheduledateOk() (*string, bool)`

GetDtInscriptionnotauthenticatedNotaryscheduledateOk returns a tuple with the DtInscriptionnotauthenticatedNotaryscheduledate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtInscriptionnotauthenticatedNotaryscheduledate

`func (o *InscriptionnotauthenticatedResponseCompound) SetDtInscriptionnotauthenticatedNotaryscheduledate(v string)`

SetDtInscriptionnotauthenticatedNotaryscheduledate sets DtInscriptionnotauthenticatedNotaryscheduledate field to given value.

### HasDtInscriptionnotauthenticatedNotaryscheduledate

`func (o *InscriptionnotauthenticatedResponseCompound) HasDtInscriptionnotauthenticatedNotaryscheduledate() bool`

HasDtInscriptionnotauthenticatedNotaryscheduledate returns a boolean if a field has been set.

### GetDtInscriptionnotauthenticatedFinancingscheduledate

`func (o *InscriptionnotauthenticatedResponseCompound) GetDtInscriptionnotauthenticatedFinancingscheduledate() string`

GetDtInscriptionnotauthenticatedFinancingscheduledate returns the DtInscriptionnotauthenticatedFinancingscheduledate field if non-nil, zero value otherwise.

### GetDtInscriptionnotauthenticatedFinancingscheduledateOk

`func (o *InscriptionnotauthenticatedResponseCompound) GetDtInscriptionnotauthenticatedFinancingscheduledateOk() (*string, bool)`

GetDtInscriptionnotauthenticatedFinancingscheduledateOk returns a tuple with the DtInscriptionnotauthenticatedFinancingscheduledate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtInscriptionnotauthenticatedFinancingscheduledate

`func (o *InscriptionnotauthenticatedResponseCompound) SetDtInscriptionnotauthenticatedFinancingscheduledate(v string)`

SetDtInscriptionnotauthenticatedFinancingscheduledate sets DtInscriptionnotauthenticatedFinancingscheduledate field to given value.

### HasDtInscriptionnotauthenticatedFinancingscheduledate

`func (o *InscriptionnotauthenticatedResponseCompound) HasDtInscriptionnotauthenticatedFinancingscheduledate() bool`

HasDtInscriptionnotauthenticatedFinancingscheduledate returns a boolean if a field has been set.

### GetBInscriptionnotauthenticatedConditional

`func (o *InscriptionnotauthenticatedResponseCompound) GetBInscriptionnotauthenticatedConditional() bool`

GetBInscriptionnotauthenticatedConditional returns the BInscriptionnotauthenticatedConditional field if non-nil, zero value otherwise.

### GetBInscriptionnotauthenticatedConditionalOk

`func (o *InscriptionnotauthenticatedResponseCompound) GetBInscriptionnotauthenticatedConditionalOk() (*bool, bool)`

GetBInscriptionnotauthenticatedConditionalOk returns a tuple with the BInscriptionnotauthenticatedConditional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBInscriptionnotauthenticatedConditional

`func (o *InscriptionnotauthenticatedResponseCompound) SetBInscriptionnotauthenticatedConditional(v bool)`

SetBInscriptionnotauthenticatedConditional sets BInscriptionnotauthenticatedConditional field to given value.


### GetBInscriptionnotauthenticatedMortgageisreferenced

`func (o *InscriptionnotauthenticatedResponseCompound) GetBInscriptionnotauthenticatedMortgageisreferenced() bool`

GetBInscriptionnotauthenticatedMortgageisreferenced returns the BInscriptionnotauthenticatedMortgageisreferenced field if non-nil, zero value otherwise.

### GetBInscriptionnotauthenticatedMortgageisreferencedOk

`func (o *InscriptionnotauthenticatedResponseCompound) GetBInscriptionnotauthenticatedMortgageisreferencedOk() (*bool, bool)`

GetBInscriptionnotauthenticatedMortgageisreferencedOk returns a tuple with the BInscriptionnotauthenticatedMortgageisreferenced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBInscriptionnotauthenticatedMortgageisreferenced

`func (o *InscriptionnotauthenticatedResponseCompound) SetBInscriptionnotauthenticatedMortgageisreferenced(v bool)`

SetBInscriptionnotauthenticatedMortgageisreferenced sets BInscriptionnotauthenticatedMortgageisreferenced field to given value.


### GetBInscriptionnotauthenticatedHomeowner

`func (o *InscriptionnotauthenticatedResponseCompound) GetBInscriptionnotauthenticatedHomeowner() bool`

GetBInscriptionnotauthenticatedHomeowner returns the BInscriptionnotauthenticatedHomeowner field if non-nil, zero value otherwise.

### GetBInscriptionnotauthenticatedHomeownerOk

`func (o *InscriptionnotauthenticatedResponseCompound) GetBInscriptionnotauthenticatedHomeownerOk() (*bool, bool)`

GetBInscriptionnotauthenticatedHomeownerOk returns a tuple with the BInscriptionnotauthenticatedHomeowner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBInscriptionnotauthenticatedHomeowner

`func (o *InscriptionnotauthenticatedResponseCompound) SetBInscriptionnotauthenticatedHomeowner(v bool)`

SetBInscriptionnotauthenticatedHomeowner sets BInscriptionnotauthenticatedHomeowner field to given value.


### GetTInscriptionnotauthenticatedConditions

`func (o *InscriptionnotauthenticatedResponseCompound) GetTInscriptionnotauthenticatedConditions() string`

GetTInscriptionnotauthenticatedConditions returns the TInscriptionnotauthenticatedConditions field if non-nil, zero value otherwise.

### GetTInscriptionnotauthenticatedConditionsOk

`func (o *InscriptionnotauthenticatedResponseCompound) GetTInscriptionnotauthenticatedConditionsOk() (*string, bool)`

GetTInscriptionnotauthenticatedConditionsOk returns a tuple with the TInscriptionnotauthenticatedConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTInscriptionnotauthenticatedConditions

`func (o *InscriptionnotauthenticatedResponseCompound) SetTInscriptionnotauthenticatedConditions(v string)`

SetTInscriptionnotauthenticatedConditions sets TInscriptionnotauthenticatedConditions field to given value.


### GetDtInscriptionnotauthenticatedConditiondeadlinedate

`func (o *InscriptionnotauthenticatedResponseCompound) GetDtInscriptionnotauthenticatedConditiondeadlinedate() string`

GetDtInscriptionnotauthenticatedConditiondeadlinedate returns the DtInscriptionnotauthenticatedConditiondeadlinedate field if non-nil, zero value otherwise.

### GetDtInscriptionnotauthenticatedConditiondeadlinedateOk

`func (o *InscriptionnotauthenticatedResponseCompound) GetDtInscriptionnotauthenticatedConditiondeadlinedateOk() (*string, bool)`

GetDtInscriptionnotauthenticatedConditiondeadlinedateOk returns a tuple with the DtInscriptionnotauthenticatedConditiondeadlinedate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtInscriptionnotauthenticatedConditiondeadlinedate

`func (o *InscriptionnotauthenticatedResponseCompound) SetDtInscriptionnotauthenticatedConditiondeadlinedate(v string)`

SetDtInscriptionnotauthenticatedConditiondeadlinedate sets DtInscriptionnotauthenticatedConditiondeadlinedate field to given value.

### HasDtInscriptionnotauthenticatedConditiondeadlinedate

`func (o *InscriptionnotauthenticatedResponseCompound) HasDtInscriptionnotauthenticatedConditiondeadlinedate() bool`

HasDtInscriptionnotauthenticatedConditiondeadlinedate returns a boolean if a field has been set.

### GetIInscriptionnotauthenticatedOrder

`func (o *InscriptionnotauthenticatedResponseCompound) GetIInscriptionnotauthenticatedOrder() int32`

GetIInscriptionnotauthenticatedOrder returns the IInscriptionnotauthenticatedOrder field if non-nil, zero value otherwise.

### GetIInscriptionnotauthenticatedOrderOk

`func (o *InscriptionnotauthenticatedResponseCompound) GetIInscriptionnotauthenticatedOrderOk() (*int32, bool)`

GetIInscriptionnotauthenticatedOrderOk returns a tuple with the IInscriptionnotauthenticatedOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIInscriptionnotauthenticatedOrder

`func (o *InscriptionnotauthenticatedResponseCompound) SetIInscriptionnotauthenticatedOrder(v int32)`

SetIInscriptionnotauthenticatedOrder sets IInscriptionnotauthenticatedOrder field to given value.


### GetBInscriptionnotauthenticatedIsactive

`func (o *InscriptionnotauthenticatedResponseCompound) GetBInscriptionnotauthenticatedIsactive() bool`

GetBInscriptionnotauthenticatedIsactive returns the BInscriptionnotauthenticatedIsactive field if non-nil, zero value otherwise.

### GetBInscriptionnotauthenticatedIsactiveOk

`func (o *InscriptionnotauthenticatedResponseCompound) GetBInscriptionnotauthenticatedIsactiveOk() (*bool, bool)`

GetBInscriptionnotauthenticatedIsactiveOk returns a tuple with the BInscriptionnotauthenticatedIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBInscriptionnotauthenticatedIsactive

`func (o *InscriptionnotauthenticatedResponseCompound) SetBInscriptionnotauthenticatedIsactive(v bool)`

SetBInscriptionnotauthenticatedIsactive sets BInscriptionnotauthenticatedIsactive field to given value.


### GetEInscriptionnotauthenticatedResidenceType

`func (o *InscriptionnotauthenticatedResponseCompound) GetEInscriptionnotauthenticatedResidenceType() FieldEInscriptionnotauthenticatedResidenceType`

GetEInscriptionnotauthenticatedResidenceType returns the EInscriptionnotauthenticatedResidenceType field if non-nil, zero value otherwise.

### GetEInscriptionnotauthenticatedResidenceTypeOk

`func (o *InscriptionnotauthenticatedResponseCompound) GetEInscriptionnotauthenticatedResidenceTypeOk() (*FieldEInscriptionnotauthenticatedResidenceType, bool)`

GetEInscriptionnotauthenticatedResidenceTypeOk returns a tuple with the EInscriptionnotauthenticatedResidenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEInscriptionnotauthenticatedResidenceType

`func (o *InscriptionnotauthenticatedResponseCompound) SetEInscriptionnotauthenticatedResidenceType(v FieldEInscriptionnotauthenticatedResidenceType)`

SetEInscriptionnotauthenticatedResidenceType sets EInscriptionnotauthenticatedResidenceType field to given value.


### GetTInscriptionnotauthenticatedChecklistnote

`func (o *InscriptionnotauthenticatedResponseCompound) GetTInscriptionnotauthenticatedChecklistnote() string`

GetTInscriptionnotauthenticatedChecklistnote returns the TInscriptionnotauthenticatedChecklistnote field if non-nil, zero value otherwise.

### GetTInscriptionnotauthenticatedChecklistnoteOk

`func (o *InscriptionnotauthenticatedResponseCompound) GetTInscriptionnotauthenticatedChecklistnoteOk() (*string, bool)`

GetTInscriptionnotauthenticatedChecklistnoteOk returns a tuple with the TInscriptionnotauthenticatedChecklistnote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTInscriptionnotauthenticatedChecklistnote

`func (o *InscriptionnotauthenticatedResponseCompound) SetTInscriptionnotauthenticatedChecklistnote(v string)`

SetTInscriptionnotauthenticatedChecklistnote sets TInscriptionnotauthenticatedChecklistnote field to given value.


### GetDInscriptionnotauthenticatedSelleronlyretribution

`func (o *InscriptionnotauthenticatedResponseCompound) GetDInscriptionnotauthenticatedSelleronlyretribution() string`

GetDInscriptionnotauthenticatedSelleronlyretribution returns the DInscriptionnotauthenticatedSelleronlyretribution field if non-nil, zero value otherwise.

### GetDInscriptionnotauthenticatedSelleronlyretributionOk

`func (o *InscriptionnotauthenticatedResponseCompound) GetDInscriptionnotauthenticatedSelleronlyretributionOk() (*string, bool)`

GetDInscriptionnotauthenticatedSelleronlyretributionOk returns a tuple with the DInscriptionnotauthenticatedSelleronlyretribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDInscriptionnotauthenticatedSelleronlyretribution

`func (o *InscriptionnotauthenticatedResponseCompound) SetDInscriptionnotauthenticatedSelleronlyretribution(v string)`

SetDInscriptionnotauthenticatedSelleronlyretribution sets DInscriptionnotauthenticatedSelleronlyretribution field to given value.


### GetBInscriptionnotauthenticatedDraft

`func (o *InscriptionnotauthenticatedResponseCompound) GetBInscriptionnotauthenticatedDraft() bool`

GetBInscriptionnotauthenticatedDraft returns the BInscriptionnotauthenticatedDraft field if non-nil, zero value otherwise.

### GetBInscriptionnotauthenticatedDraftOk

`func (o *InscriptionnotauthenticatedResponseCompound) GetBInscriptionnotauthenticatedDraftOk() (*bool, bool)`

GetBInscriptionnotauthenticatedDraftOk returns a tuple with the BInscriptionnotauthenticatedDraft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBInscriptionnotauthenticatedDraft

`func (o *InscriptionnotauthenticatedResponseCompound) SetBInscriptionnotauthenticatedDraft(v bool)`

SetBInscriptionnotauthenticatedDraft sets BInscriptionnotauthenticatedDraft field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


