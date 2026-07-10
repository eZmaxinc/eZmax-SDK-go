# InscriptionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiInscriptionID** | **int32** | The unique ID of the Inscription. | 
**FkiCompanyID** | **int32** | The unique ID of the Company | 
**SCompanyNameX** | Pointer to **string** | The Name of the Company in the language of the requester | [optional] 
**FkiDepartmentID** | **int32** | The unique ID of the Department | 
**SDepartmentNameX** | Pointer to **string** | The Name of the Department in the language of the requester | [optional] 
**FkiRealestateboardID** | **int32** | The unique ID of the Realestateboard | 
**SRealestateboardNameX** | Pointer to **string** | The name of the Realestateboard | [optional] 
**FkiAddressID** | **int32** | The unique ID of the Address | 
**SAddress** | Pointer to **string** | The complete address in a single line | [optional] 
**FkiInscriptionbuildingtypeID** | **int32** | The unique ID of the Inscriptionbuildingtype | 
**SInscriptionbuildingtypeNameX** | Pointer to **string** | The name of the Inscriptionbuildingtype in the language of the requester | [optional] 
**FkiInscriptiontypeID** | **int32** | The unique ID of the Inscriptiontype | 
**SInscriptiontypeNameX** | Pointer to **string** | The name of the Inscriptiontype in the language of the requester | [optional] 
**FkiInscriptioncategoryID** | **int32** | The unique ID of the Inscriptioncategory | 
**SInscriptioncategoryNameX** | Pointer to **string** | The name of the Inscriptioncategory in the language of the requester | [optional] 
**EInscriptionStep** | [**FieldEInscriptionStep**](FieldEInscriptionStep.md) |  | 
**EInscriptionResidenceType** | [**FieldEInscriptionResidenceType**](FieldEInscriptionResidenceType.md) |  | 
**SInscriptionCivicend** | **string** | The civicend of the Inscription | 
**SInscriptionMLS** | **string** | The mls of the Inscription | 
**SInscriptionContract** | **string** | The sale contract number | 
**IInscriptionSellerdeclaration** | **int32** | The sellerdeclaration of the Inscription | 
**EInscriptionType** | [**FieldEInscriptionType**](FieldEInscriptionType.md) |  | 
**DInscriptionInitialsaleprice** | **string** | The initialsaleprice of the Inscription | 
**DInscriptionSaleprice** | **string** | The saleprice of the Inscription | 
**DInscriptionRentprice** | **string** | The rentprice of the Inscription | 
**EInscriptionRemunerationtype** | [**FieldEInscriptionRemunerationtype**](FieldEInscriptionRemunerationtype.md) |  | 
**EInscriptionRemunerationinscriptorsellertype** | [**FieldEInscriptionRemunerationinscriptorsellertype**](FieldEInscriptionRemunerationinscriptorsellertype.md) |  | 
**EInscriptionRemunerationreferencetype** | [**FieldEInscriptionRemunerationreferencetype**](FieldEInscriptionRemunerationreferencetype.md) |  | 
**EInscriptionRemunerationtotaltype** | [**FieldEInscriptionRemunerationtotaltype**](FieldEInscriptionRemunerationtotaltype.md) |  | 
**DInscriptionRemuneration** | **string** | The remuneration of the Inscription | 
**DInscriptionRemunerationinscriptorseller** | **string** | The remunerationinscriptorseller of the Inscription | 
**DInscriptionRemunerationreference** | **string** | The remunerationreference of the Inscription | 
**DInscriptionRemunerationtotal** | **string** | The remunerationtotal of the Inscription | 
**DInscriptionMortgagesold** | **string** | The mortgagesold of the Inscription | 
**DtInscriptionDate** | **string** | The date of the Inscription | 
**DtInscriptionCancellationdate** | **string** | The cancellationdate of the Inscription | 
**DtInscriptionInitialexpirationdate** | **string** | The initialexpirationdate of the Inscription | 
**DtInscriptionExpirationdate** | **string** | The expirationdate of the Inscription | 
**DtInscriptionNotarydate** | **string** | The notarydate of the Inscription | 
**DtInscriptionNotaryentereddate** | **string** | The notaryentereddate of the Inscription | 
**TInscriptionCadastre** | **string** | The cadastre of the Inscription | 
**BInscriptionReference** | **bool** | Whether if it&#39;s an reference | 
**BInscriptionInspection** | **bool** | Whether the inscription can be acces by an inspector | 
**BInscriptionIsactive** | **bool** | Whether the inscription is active or not | 
**TInscriptionChecklistnote** | **string** | The checklistnote of the Inscription | 
**BInscriptionNew** | **bool** | Whether if it&#39;s an new | 
**BInscriptionHomeowner** | **bool** | Whether if it&#39;s an homeowner | 
**BInscriptionArchived** | **bool** | Whether the inscription is archived or not | 
**BInscriptionLitigation** | **bool** | Whether if it&#39;s an litigation | 
**BInscriptionRepossession** | **bool** | Whether if it&#39;s an repossession | 
**BInscriptionIssolicitation** | **bool** | Whether if it&#39;s an issolicitation | 
**BInscriptionSalebyowner** | **bool** | Whether if it&#39;s an salebyowner | 
**BInscriptionSoldwithoutlegalwarranty** | **bool** | Whether if it&#39;s an soldwithoutlegalwarranty | 
**IInscriptionConstructionyear** | **int32** | The constructionyear of the Inscription | 
**IInscriptionUnit** | **int32** | The unit of the Inscription | 

## Methods

### NewInscriptionResponse

`func NewInscriptionResponse(pkiInscriptionID int32, fkiCompanyID int32, fkiDepartmentID int32, fkiRealestateboardID int32, fkiAddressID int32, fkiInscriptionbuildingtypeID int32, fkiInscriptiontypeID int32, fkiInscriptioncategoryID int32, eInscriptionStep FieldEInscriptionStep, eInscriptionResidenceType FieldEInscriptionResidenceType, sInscriptionCivicend string, sInscriptionMLS string, sInscriptionContract string, iInscriptionSellerdeclaration int32, eInscriptionType FieldEInscriptionType, dInscriptionInitialsaleprice string, dInscriptionSaleprice string, dInscriptionRentprice string, eInscriptionRemunerationtype FieldEInscriptionRemunerationtype, eInscriptionRemunerationinscriptorsellertype FieldEInscriptionRemunerationinscriptorsellertype, eInscriptionRemunerationreferencetype FieldEInscriptionRemunerationreferencetype, eInscriptionRemunerationtotaltype FieldEInscriptionRemunerationtotaltype, dInscriptionRemuneration string, dInscriptionRemunerationinscriptorseller string, dInscriptionRemunerationreference string, dInscriptionRemunerationtotal string, dInscriptionMortgagesold string, dtInscriptionDate string, dtInscriptionCancellationdate string, dtInscriptionInitialexpirationdate string, dtInscriptionExpirationdate string, dtInscriptionNotarydate string, dtInscriptionNotaryentereddate string, tInscriptionCadastre string, bInscriptionReference bool, bInscriptionInspection bool, bInscriptionIsactive bool, tInscriptionChecklistnote string, bInscriptionNew bool, bInscriptionHomeowner bool, bInscriptionArchived bool, bInscriptionLitigation bool, bInscriptionRepossession bool, bInscriptionIssolicitation bool, bInscriptionSalebyowner bool, bInscriptionSoldwithoutlegalwarranty bool, iInscriptionConstructionyear int32, iInscriptionUnit int32, ) *InscriptionResponse`

NewInscriptionResponse instantiates a new InscriptionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInscriptionResponseWithDefaults

`func NewInscriptionResponseWithDefaults() *InscriptionResponse`

NewInscriptionResponseWithDefaults instantiates a new InscriptionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiInscriptionID

`func (o *InscriptionResponse) GetPkiInscriptionID() int32`

GetPkiInscriptionID returns the PkiInscriptionID field if non-nil, zero value otherwise.

### GetPkiInscriptionIDOk

`func (o *InscriptionResponse) GetPkiInscriptionIDOk() (*int32, bool)`

GetPkiInscriptionIDOk returns a tuple with the PkiInscriptionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiInscriptionID

`func (o *InscriptionResponse) SetPkiInscriptionID(v int32)`

SetPkiInscriptionID sets PkiInscriptionID field to given value.


### GetFkiCompanyID

`func (o *InscriptionResponse) GetFkiCompanyID() int32`

GetFkiCompanyID returns the FkiCompanyID field if non-nil, zero value otherwise.

### GetFkiCompanyIDOk

`func (o *InscriptionResponse) GetFkiCompanyIDOk() (*int32, bool)`

GetFkiCompanyIDOk returns a tuple with the FkiCompanyID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiCompanyID

`func (o *InscriptionResponse) SetFkiCompanyID(v int32)`

SetFkiCompanyID sets FkiCompanyID field to given value.


### GetSCompanyNameX

`func (o *InscriptionResponse) GetSCompanyNameX() string`

GetSCompanyNameX returns the SCompanyNameX field if non-nil, zero value otherwise.

### GetSCompanyNameXOk

`func (o *InscriptionResponse) GetSCompanyNameXOk() (*string, bool)`

GetSCompanyNameXOk returns a tuple with the SCompanyNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSCompanyNameX

`func (o *InscriptionResponse) SetSCompanyNameX(v string)`

SetSCompanyNameX sets SCompanyNameX field to given value.

### HasSCompanyNameX

`func (o *InscriptionResponse) HasSCompanyNameX() bool`

HasSCompanyNameX returns a boolean if a field has been set.

### GetFkiDepartmentID

`func (o *InscriptionResponse) GetFkiDepartmentID() int32`

GetFkiDepartmentID returns the FkiDepartmentID field if non-nil, zero value otherwise.

### GetFkiDepartmentIDOk

`func (o *InscriptionResponse) GetFkiDepartmentIDOk() (*int32, bool)`

GetFkiDepartmentIDOk returns a tuple with the FkiDepartmentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiDepartmentID

`func (o *InscriptionResponse) SetFkiDepartmentID(v int32)`

SetFkiDepartmentID sets FkiDepartmentID field to given value.


### GetSDepartmentNameX

`func (o *InscriptionResponse) GetSDepartmentNameX() string`

GetSDepartmentNameX returns the SDepartmentNameX field if non-nil, zero value otherwise.

### GetSDepartmentNameXOk

`func (o *InscriptionResponse) GetSDepartmentNameXOk() (*string, bool)`

GetSDepartmentNameXOk returns a tuple with the SDepartmentNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSDepartmentNameX

`func (o *InscriptionResponse) SetSDepartmentNameX(v string)`

SetSDepartmentNameX sets SDepartmentNameX field to given value.

### HasSDepartmentNameX

`func (o *InscriptionResponse) HasSDepartmentNameX() bool`

HasSDepartmentNameX returns a boolean if a field has been set.

### GetFkiRealestateboardID

`func (o *InscriptionResponse) GetFkiRealestateboardID() int32`

GetFkiRealestateboardID returns the FkiRealestateboardID field if non-nil, zero value otherwise.

### GetFkiRealestateboardIDOk

`func (o *InscriptionResponse) GetFkiRealestateboardIDOk() (*int32, bool)`

GetFkiRealestateboardIDOk returns a tuple with the FkiRealestateboardID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiRealestateboardID

`func (o *InscriptionResponse) SetFkiRealestateboardID(v int32)`

SetFkiRealestateboardID sets FkiRealestateboardID field to given value.


### GetSRealestateboardNameX

`func (o *InscriptionResponse) GetSRealestateboardNameX() string`

GetSRealestateboardNameX returns the SRealestateboardNameX field if non-nil, zero value otherwise.

### GetSRealestateboardNameXOk

`func (o *InscriptionResponse) GetSRealestateboardNameXOk() (*string, bool)`

GetSRealestateboardNameXOk returns a tuple with the SRealestateboardNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSRealestateboardNameX

`func (o *InscriptionResponse) SetSRealestateboardNameX(v string)`

SetSRealestateboardNameX sets SRealestateboardNameX field to given value.

### HasSRealestateboardNameX

`func (o *InscriptionResponse) HasSRealestateboardNameX() bool`

HasSRealestateboardNameX returns a boolean if a field has been set.

### GetFkiAddressID

`func (o *InscriptionResponse) GetFkiAddressID() int32`

GetFkiAddressID returns the FkiAddressID field if non-nil, zero value otherwise.

### GetFkiAddressIDOk

`func (o *InscriptionResponse) GetFkiAddressIDOk() (*int32, bool)`

GetFkiAddressIDOk returns a tuple with the FkiAddressID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiAddressID

`func (o *InscriptionResponse) SetFkiAddressID(v int32)`

SetFkiAddressID sets FkiAddressID field to given value.


### GetSAddress

`func (o *InscriptionResponse) GetSAddress() string`

GetSAddress returns the SAddress field if non-nil, zero value otherwise.

### GetSAddressOk

`func (o *InscriptionResponse) GetSAddressOk() (*string, bool)`

GetSAddressOk returns a tuple with the SAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAddress

`func (o *InscriptionResponse) SetSAddress(v string)`

SetSAddress sets SAddress field to given value.

### HasSAddress

`func (o *InscriptionResponse) HasSAddress() bool`

HasSAddress returns a boolean if a field has been set.

### GetFkiInscriptionbuildingtypeID

`func (o *InscriptionResponse) GetFkiInscriptionbuildingtypeID() int32`

GetFkiInscriptionbuildingtypeID returns the FkiInscriptionbuildingtypeID field if non-nil, zero value otherwise.

### GetFkiInscriptionbuildingtypeIDOk

`func (o *InscriptionResponse) GetFkiInscriptionbuildingtypeIDOk() (*int32, bool)`

GetFkiInscriptionbuildingtypeIDOk returns a tuple with the FkiInscriptionbuildingtypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiInscriptionbuildingtypeID

`func (o *InscriptionResponse) SetFkiInscriptionbuildingtypeID(v int32)`

SetFkiInscriptionbuildingtypeID sets FkiInscriptionbuildingtypeID field to given value.


### GetSInscriptionbuildingtypeNameX

`func (o *InscriptionResponse) GetSInscriptionbuildingtypeNameX() string`

GetSInscriptionbuildingtypeNameX returns the SInscriptionbuildingtypeNameX field if non-nil, zero value otherwise.

### GetSInscriptionbuildingtypeNameXOk

`func (o *InscriptionResponse) GetSInscriptionbuildingtypeNameXOk() (*string, bool)`

GetSInscriptionbuildingtypeNameXOk returns a tuple with the SInscriptionbuildingtypeNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSInscriptionbuildingtypeNameX

`func (o *InscriptionResponse) SetSInscriptionbuildingtypeNameX(v string)`

SetSInscriptionbuildingtypeNameX sets SInscriptionbuildingtypeNameX field to given value.

### HasSInscriptionbuildingtypeNameX

`func (o *InscriptionResponse) HasSInscriptionbuildingtypeNameX() bool`

HasSInscriptionbuildingtypeNameX returns a boolean if a field has been set.

### GetFkiInscriptiontypeID

`func (o *InscriptionResponse) GetFkiInscriptiontypeID() int32`

GetFkiInscriptiontypeID returns the FkiInscriptiontypeID field if non-nil, zero value otherwise.

### GetFkiInscriptiontypeIDOk

`func (o *InscriptionResponse) GetFkiInscriptiontypeIDOk() (*int32, bool)`

GetFkiInscriptiontypeIDOk returns a tuple with the FkiInscriptiontypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiInscriptiontypeID

`func (o *InscriptionResponse) SetFkiInscriptiontypeID(v int32)`

SetFkiInscriptiontypeID sets FkiInscriptiontypeID field to given value.


### GetSInscriptiontypeNameX

`func (o *InscriptionResponse) GetSInscriptiontypeNameX() string`

GetSInscriptiontypeNameX returns the SInscriptiontypeNameX field if non-nil, zero value otherwise.

### GetSInscriptiontypeNameXOk

`func (o *InscriptionResponse) GetSInscriptiontypeNameXOk() (*string, bool)`

GetSInscriptiontypeNameXOk returns a tuple with the SInscriptiontypeNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSInscriptiontypeNameX

`func (o *InscriptionResponse) SetSInscriptiontypeNameX(v string)`

SetSInscriptiontypeNameX sets SInscriptiontypeNameX field to given value.

### HasSInscriptiontypeNameX

`func (o *InscriptionResponse) HasSInscriptiontypeNameX() bool`

HasSInscriptiontypeNameX returns a boolean if a field has been set.

### GetFkiInscriptioncategoryID

`func (o *InscriptionResponse) GetFkiInscriptioncategoryID() int32`

GetFkiInscriptioncategoryID returns the FkiInscriptioncategoryID field if non-nil, zero value otherwise.

### GetFkiInscriptioncategoryIDOk

`func (o *InscriptionResponse) GetFkiInscriptioncategoryIDOk() (*int32, bool)`

GetFkiInscriptioncategoryIDOk returns a tuple with the FkiInscriptioncategoryID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiInscriptioncategoryID

`func (o *InscriptionResponse) SetFkiInscriptioncategoryID(v int32)`

SetFkiInscriptioncategoryID sets FkiInscriptioncategoryID field to given value.


### GetSInscriptioncategoryNameX

`func (o *InscriptionResponse) GetSInscriptioncategoryNameX() string`

GetSInscriptioncategoryNameX returns the SInscriptioncategoryNameX field if non-nil, zero value otherwise.

### GetSInscriptioncategoryNameXOk

`func (o *InscriptionResponse) GetSInscriptioncategoryNameXOk() (*string, bool)`

GetSInscriptioncategoryNameXOk returns a tuple with the SInscriptioncategoryNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSInscriptioncategoryNameX

`func (o *InscriptionResponse) SetSInscriptioncategoryNameX(v string)`

SetSInscriptioncategoryNameX sets SInscriptioncategoryNameX field to given value.

### HasSInscriptioncategoryNameX

`func (o *InscriptionResponse) HasSInscriptioncategoryNameX() bool`

HasSInscriptioncategoryNameX returns a boolean if a field has been set.

### GetEInscriptionStep

`func (o *InscriptionResponse) GetEInscriptionStep() FieldEInscriptionStep`

GetEInscriptionStep returns the EInscriptionStep field if non-nil, zero value otherwise.

### GetEInscriptionStepOk

`func (o *InscriptionResponse) GetEInscriptionStepOk() (*FieldEInscriptionStep, bool)`

GetEInscriptionStepOk returns a tuple with the EInscriptionStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEInscriptionStep

`func (o *InscriptionResponse) SetEInscriptionStep(v FieldEInscriptionStep)`

SetEInscriptionStep sets EInscriptionStep field to given value.


### GetEInscriptionResidenceType

`func (o *InscriptionResponse) GetEInscriptionResidenceType() FieldEInscriptionResidenceType`

GetEInscriptionResidenceType returns the EInscriptionResidenceType field if non-nil, zero value otherwise.

### GetEInscriptionResidenceTypeOk

`func (o *InscriptionResponse) GetEInscriptionResidenceTypeOk() (*FieldEInscriptionResidenceType, bool)`

GetEInscriptionResidenceTypeOk returns a tuple with the EInscriptionResidenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEInscriptionResidenceType

`func (o *InscriptionResponse) SetEInscriptionResidenceType(v FieldEInscriptionResidenceType)`

SetEInscriptionResidenceType sets EInscriptionResidenceType field to given value.


### GetSInscriptionCivicend

`func (o *InscriptionResponse) GetSInscriptionCivicend() string`

GetSInscriptionCivicend returns the SInscriptionCivicend field if non-nil, zero value otherwise.

### GetSInscriptionCivicendOk

`func (o *InscriptionResponse) GetSInscriptionCivicendOk() (*string, bool)`

GetSInscriptionCivicendOk returns a tuple with the SInscriptionCivicend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSInscriptionCivicend

`func (o *InscriptionResponse) SetSInscriptionCivicend(v string)`

SetSInscriptionCivicend sets SInscriptionCivicend field to given value.


### GetSInscriptionMLS

`func (o *InscriptionResponse) GetSInscriptionMLS() string`

GetSInscriptionMLS returns the SInscriptionMLS field if non-nil, zero value otherwise.

### GetSInscriptionMLSOk

`func (o *InscriptionResponse) GetSInscriptionMLSOk() (*string, bool)`

GetSInscriptionMLSOk returns a tuple with the SInscriptionMLS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSInscriptionMLS

`func (o *InscriptionResponse) SetSInscriptionMLS(v string)`

SetSInscriptionMLS sets SInscriptionMLS field to given value.


### GetSInscriptionContract

`func (o *InscriptionResponse) GetSInscriptionContract() string`

GetSInscriptionContract returns the SInscriptionContract field if non-nil, zero value otherwise.

### GetSInscriptionContractOk

`func (o *InscriptionResponse) GetSInscriptionContractOk() (*string, bool)`

GetSInscriptionContractOk returns a tuple with the SInscriptionContract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSInscriptionContract

`func (o *InscriptionResponse) SetSInscriptionContract(v string)`

SetSInscriptionContract sets SInscriptionContract field to given value.


### GetIInscriptionSellerdeclaration

`func (o *InscriptionResponse) GetIInscriptionSellerdeclaration() int32`

GetIInscriptionSellerdeclaration returns the IInscriptionSellerdeclaration field if non-nil, zero value otherwise.

### GetIInscriptionSellerdeclarationOk

`func (o *InscriptionResponse) GetIInscriptionSellerdeclarationOk() (*int32, bool)`

GetIInscriptionSellerdeclarationOk returns a tuple with the IInscriptionSellerdeclaration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIInscriptionSellerdeclaration

`func (o *InscriptionResponse) SetIInscriptionSellerdeclaration(v int32)`

SetIInscriptionSellerdeclaration sets IInscriptionSellerdeclaration field to given value.


### GetEInscriptionType

`func (o *InscriptionResponse) GetEInscriptionType() FieldEInscriptionType`

GetEInscriptionType returns the EInscriptionType field if non-nil, zero value otherwise.

### GetEInscriptionTypeOk

`func (o *InscriptionResponse) GetEInscriptionTypeOk() (*FieldEInscriptionType, bool)`

GetEInscriptionTypeOk returns a tuple with the EInscriptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEInscriptionType

`func (o *InscriptionResponse) SetEInscriptionType(v FieldEInscriptionType)`

SetEInscriptionType sets EInscriptionType field to given value.


### GetDInscriptionInitialsaleprice

`func (o *InscriptionResponse) GetDInscriptionInitialsaleprice() string`

GetDInscriptionInitialsaleprice returns the DInscriptionInitialsaleprice field if non-nil, zero value otherwise.

### GetDInscriptionInitialsalepriceOk

`func (o *InscriptionResponse) GetDInscriptionInitialsalepriceOk() (*string, bool)`

GetDInscriptionInitialsalepriceOk returns a tuple with the DInscriptionInitialsaleprice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDInscriptionInitialsaleprice

`func (o *InscriptionResponse) SetDInscriptionInitialsaleprice(v string)`

SetDInscriptionInitialsaleprice sets DInscriptionInitialsaleprice field to given value.


### GetDInscriptionSaleprice

`func (o *InscriptionResponse) GetDInscriptionSaleprice() string`

GetDInscriptionSaleprice returns the DInscriptionSaleprice field if non-nil, zero value otherwise.

### GetDInscriptionSalepriceOk

`func (o *InscriptionResponse) GetDInscriptionSalepriceOk() (*string, bool)`

GetDInscriptionSalepriceOk returns a tuple with the DInscriptionSaleprice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDInscriptionSaleprice

`func (o *InscriptionResponse) SetDInscriptionSaleprice(v string)`

SetDInscriptionSaleprice sets DInscriptionSaleprice field to given value.


### GetDInscriptionRentprice

`func (o *InscriptionResponse) GetDInscriptionRentprice() string`

GetDInscriptionRentprice returns the DInscriptionRentprice field if non-nil, zero value otherwise.

### GetDInscriptionRentpriceOk

`func (o *InscriptionResponse) GetDInscriptionRentpriceOk() (*string, bool)`

GetDInscriptionRentpriceOk returns a tuple with the DInscriptionRentprice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDInscriptionRentprice

`func (o *InscriptionResponse) SetDInscriptionRentprice(v string)`

SetDInscriptionRentprice sets DInscriptionRentprice field to given value.


### GetEInscriptionRemunerationtype

`func (o *InscriptionResponse) GetEInscriptionRemunerationtype() FieldEInscriptionRemunerationtype`

GetEInscriptionRemunerationtype returns the EInscriptionRemunerationtype field if non-nil, zero value otherwise.

### GetEInscriptionRemunerationtypeOk

`func (o *InscriptionResponse) GetEInscriptionRemunerationtypeOk() (*FieldEInscriptionRemunerationtype, bool)`

GetEInscriptionRemunerationtypeOk returns a tuple with the EInscriptionRemunerationtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEInscriptionRemunerationtype

`func (o *InscriptionResponse) SetEInscriptionRemunerationtype(v FieldEInscriptionRemunerationtype)`

SetEInscriptionRemunerationtype sets EInscriptionRemunerationtype field to given value.


### GetEInscriptionRemunerationinscriptorsellertype

`func (o *InscriptionResponse) GetEInscriptionRemunerationinscriptorsellertype() FieldEInscriptionRemunerationinscriptorsellertype`

GetEInscriptionRemunerationinscriptorsellertype returns the EInscriptionRemunerationinscriptorsellertype field if non-nil, zero value otherwise.

### GetEInscriptionRemunerationinscriptorsellertypeOk

`func (o *InscriptionResponse) GetEInscriptionRemunerationinscriptorsellertypeOk() (*FieldEInscriptionRemunerationinscriptorsellertype, bool)`

GetEInscriptionRemunerationinscriptorsellertypeOk returns a tuple with the EInscriptionRemunerationinscriptorsellertype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEInscriptionRemunerationinscriptorsellertype

`func (o *InscriptionResponse) SetEInscriptionRemunerationinscriptorsellertype(v FieldEInscriptionRemunerationinscriptorsellertype)`

SetEInscriptionRemunerationinscriptorsellertype sets EInscriptionRemunerationinscriptorsellertype field to given value.


### GetEInscriptionRemunerationreferencetype

`func (o *InscriptionResponse) GetEInscriptionRemunerationreferencetype() FieldEInscriptionRemunerationreferencetype`

GetEInscriptionRemunerationreferencetype returns the EInscriptionRemunerationreferencetype field if non-nil, zero value otherwise.

### GetEInscriptionRemunerationreferencetypeOk

`func (o *InscriptionResponse) GetEInscriptionRemunerationreferencetypeOk() (*FieldEInscriptionRemunerationreferencetype, bool)`

GetEInscriptionRemunerationreferencetypeOk returns a tuple with the EInscriptionRemunerationreferencetype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEInscriptionRemunerationreferencetype

`func (o *InscriptionResponse) SetEInscriptionRemunerationreferencetype(v FieldEInscriptionRemunerationreferencetype)`

SetEInscriptionRemunerationreferencetype sets EInscriptionRemunerationreferencetype field to given value.


### GetEInscriptionRemunerationtotaltype

`func (o *InscriptionResponse) GetEInscriptionRemunerationtotaltype() FieldEInscriptionRemunerationtotaltype`

GetEInscriptionRemunerationtotaltype returns the EInscriptionRemunerationtotaltype field if non-nil, zero value otherwise.

### GetEInscriptionRemunerationtotaltypeOk

`func (o *InscriptionResponse) GetEInscriptionRemunerationtotaltypeOk() (*FieldEInscriptionRemunerationtotaltype, bool)`

GetEInscriptionRemunerationtotaltypeOk returns a tuple with the EInscriptionRemunerationtotaltype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEInscriptionRemunerationtotaltype

`func (o *InscriptionResponse) SetEInscriptionRemunerationtotaltype(v FieldEInscriptionRemunerationtotaltype)`

SetEInscriptionRemunerationtotaltype sets EInscriptionRemunerationtotaltype field to given value.


### GetDInscriptionRemuneration

`func (o *InscriptionResponse) GetDInscriptionRemuneration() string`

GetDInscriptionRemuneration returns the DInscriptionRemuneration field if non-nil, zero value otherwise.

### GetDInscriptionRemunerationOk

`func (o *InscriptionResponse) GetDInscriptionRemunerationOk() (*string, bool)`

GetDInscriptionRemunerationOk returns a tuple with the DInscriptionRemuneration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDInscriptionRemuneration

`func (o *InscriptionResponse) SetDInscriptionRemuneration(v string)`

SetDInscriptionRemuneration sets DInscriptionRemuneration field to given value.


### GetDInscriptionRemunerationinscriptorseller

`func (o *InscriptionResponse) GetDInscriptionRemunerationinscriptorseller() string`

GetDInscriptionRemunerationinscriptorseller returns the DInscriptionRemunerationinscriptorseller field if non-nil, zero value otherwise.

### GetDInscriptionRemunerationinscriptorsellerOk

`func (o *InscriptionResponse) GetDInscriptionRemunerationinscriptorsellerOk() (*string, bool)`

GetDInscriptionRemunerationinscriptorsellerOk returns a tuple with the DInscriptionRemunerationinscriptorseller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDInscriptionRemunerationinscriptorseller

`func (o *InscriptionResponse) SetDInscriptionRemunerationinscriptorseller(v string)`

SetDInscriptionRemunerationinscriptorseller sets DInscriptionRemunerationinscriptorseller field to given value.


### GetDInscriptionRemunerationreference

`func (o *InscriptionResponse) GetDInscriptionRemunerationreference() string`

GetDInscriptionRemunerationreference returns the DInscriptionRemunerationreference field if non-nil, zero value otherwise.

### GetDInscriptionRemunerationreferenceOk

`func (o *InscriptionResponse) GetDInscriptionRemunerationreferenceOk() (*string, bool)`

GetDInscriptionRemunerationreferenceOk returns a tuple with the DInscriptionRemunerationreference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDInscriptionRemunerationreference

`func (o *InscriptionResponse) SetDInscriptionRemunerationreference(v string)`

SetDInscriptionRemunerationreference sets DInscriptionRemunerationreference field to given value.


### GetDInscriptionRemunerationtotal

`func (o *InscriptionResponse) GetDInscriptionRemunerationtotal() string`

GetDInscriptionRemunerationtotal returns the DInscriptionRemunerationtotal field if non-nil, zero value otherwise.

### GetDInscriptionRemunerationtotalOk

`func (o *InscriptionResponse) GetDInscriptionRemunerationtotalOk() (*string, bool)`

GetDInscriptionRemunerationtotalOk returns a tuple with the DInscriptionRemunerationtotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDInscriptionRemunerationtotal

`func (o *InscriptionResponse) SetDInscriptionRemunerationtotal(v string)`

SetDInscriptionRemunerationtotal sets DInscriptionRemunerationtotal field to given value.


### GetDInscriptionMortgagesold

`func (o *InscriptionResponse) GetDInscriptionMortgagesold() string`

GetDInscriptionMortgagesold returns the DInscriptionMortgagesold field if non-nil, zero value otherwise.

### GetDInscriptionMortgagesoldOk

`func (o *InscriptionResponse) GetDInscriptionMortgagesoldOk() (*string, bool)`

GetDInscriptionMortgagesoldOk returns a tuple with the DInscriptionMortgagesold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDInscriptionMortgagesold

`func (o *InscriptionResponse) SetDInscriptionMortgagesold(v string)`

SetDInscriptionMortgagesold sets DInscriptionMortgagesold field to given value.


### GetDtInscriptionDate

`func (o *InscriptionResponse) GetDtInscriptionDate() string`

GetDtInscriptionDate returns the DtInscriptionDate field if non-nil, zero value otherwise.

### GetDtInscriptionDateOk

`func (o *InscriptionResponse) GetDtInscriptionDateOk() (*string, bool)`

GetDtInscriptionDateOk returns a tuple with the DtInscriptionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtInscriptionDate

`func (o *InscriptionResponse) SetDtInscriptionDate(v string)`

SetDtInscriptionDate sets DtInscriptionDate field to given value.


### GetDtInscriptionCancellationdate

`func (o *InscriptionResponse) GetDtInscriptionCancellationdate() string`

GetDtInscriptionCancellationdate returns the DtInscriptionCancellationdate field if non-nil, zero value otherwise.

### GetDtInscriptionCancellationdateOk

`func (o *InscriptionResponse) GetDtInscriptionCancellationdateOk() (*string, bool)`

GetDtInscriptionCancellationdateOk returns a tuple with the DtInscriptionCancellationdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtInscriptionCancellationdate

`func (o *InscriptionResponse) SetDtInscriptionCancellationdate(v string)`

SetDtInscriptionCancellationdate sets DtInscriptionCancellationdate field to given value.


### GetDtInscriptionInitialexpirationdate

`func (o *InscriptionResponse) GetDtInscriptionInitialexpirationdate() string`

GetDtInscriptionInitialexpirationdate returns the DtInscriptionInitialexpirationdate field if non-nil, zero value otherwise.

### GetDtInscriptionInitialexpirationdateOk

`func (o *InscriptionResponse) GetDtInscriptionInitialexpirationdateOk() (*string, bool)`

GetDtInscriptionInitialexpirationdateOk returns a tuple with the DtInscriptionInitialexpirationdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtInscriptionInitialexpirationdate

`func (o *InscriptionResponse) SetDtInscriptionInitialexpirationdate(v string)`

SetDtInscriptionInitialexpirationdate sets DtInscriptionInitialexpirationdate field to given value.


### GetDtInscriptionExpirationdate

`func (o *InscriptionResponse) GetDtInscriptionExpirationdate() string`

GetDtInscriptionExpirationdate returns the DtInscriptionExpirationdate field if non-nil, zero value otherwise.

### GetDtInscriptionExpirationdateOk

`func (o *InscriptionResponse) GetDtInscriptionExpirationdateOk() (*string, bool)`

GetDtInscriptionExpirationdateOk returns a tuple with the DtInscriptionExpirationdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtInscriptionExpirationdate

`func (o *InscriptionResponse) SetDtInscriptionExpirationdate(v string)`

SetDtInscriptionExpirationdate sets DtInscriptionExpirationdate field to given value.


### GetDtInscriptionNotarydate

`func (o *InscriptionResponse) GetDtInscriptionNotarydate() string`

GetDtInscriptionNotarydate returns the DtInscriptionNotarydate field if non-nil, zero value otherwise.

### GetDtInscriptionNotarydateOk

`func (o *InscriptionResponse) GetDtInscriptionNotarydateOk() (*string, bool)`

GetDtInscriptionNotarydateOk returns a tuple with the DtInscriptionNotarydate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtInscriptionNotarydate

`func (o *InscriptionResponse) SetDtInscriptionNotarydate(v string)`

SetDtInscriptionNotarydate sets DtInscriptionNotarydate field to given value.


### GetDtInscriptionNotaryentereddate

`func (o *InscriptionResponse) GetDtInscriptionNotaryentereddate() string`

GetDtInscriptionNotaryentereddate returns the DtInscriptionNotaryentereddate field if non-nil, zero value otherwise.

### GetDtInscriptionNotaryentereddateOk

`func (o *InscriptionResponse) GetDtInscriptionNotaryentereddateOk() (*string, bool)`

GetDtInscriptionNotaryentereddateOk returns a tuple with the DtInscriptionNotaryentereddate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtInscriptionNotaryentereddate

`func (o *InscriptionResponse) SetDtInscriptionNotaryentereddate(v string)`

SetDtInscriptionNotaryentereddate sets DtInscriptionNotaryentereddate field to given value.


### GetTInscriptionCadastre

`func (o *InscriptionResponse) GetTInscriptionCadastre() string`

GetTInscriptionCadastre returns the TInscriptionCadastre field if non-nil, zero value otherwise.

### GetTInscriptionCadastreOk

`func (o *InscriptionResponse) GetTInscriptionCadastreOk() (*string, bool)`

GetTInscriptionCadastreOk returns a tuple with the TInscriptionCadastre field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTInscriptionCadastre

`func (o *InscriptionResponse) SetTInscriptionCadastre(v string)`

SetTInscriptionCadastre sets TInscriptionCadastre field to given value.


### GetBInscriptionReference

`func (o *InscriptionResponse) GetBInscriptionReference() bool`

GetBInscriptionReference returns the BInscriptionReference field if non-nil, zero value otherwise.

### GetBInscriptionReferenceOk

`func (o *InscriptionResponse) GetBInscriptionReferenceOk() (*bool, bool)`

GetBInscriptionReferenceOk returns a tuple with the BInscriptionReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBInscriptionReference

`func (o *InscriptionResponse) SetBInscriptionReference(v bool)`

SetBInscriptionReference sets BInscriptionReference field to given value.


### GetBInscriptionInspection

`func (o *InscriptionResponse) GetBInscriptionInspection() bool`

GetBInscriptionInspection returns the BInscriptionInspection field if non-nil, zero value otherwise.

### GetBInscriptionInspectionOk

`func (o *InscriptionResponse) GetBInscriptionInspectionOk() (*bool, bool)`

GetBInscriptionInspectionOk returns a tuple with the BInscriptionInspection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBInscriptionInspection

`func (o *InscriptionResponse) SetBInscriptionInspection(v bool)`

SetBInscriptionInspection sets BInscriptionInspection field to given value.


### GetBInscriptionIsactive

`func (o *InscriptionResponse) GetBInscriptionIsactive() bool`

GetBInscriptionIsactive returns the BInscriptionIsactive field if non-nil, zero value otherwise.

### GetBInscriptionIsactiveOk

`func (o *InscriptionResponse) GetBInscriptionIsactiveOk() (*bool, bool)`

GetBInscriptionIsactiveOk returns a tuple with the BInscriptionIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBInscriptionIsactive

`func (o *InscriptionResponse) SetBInscriptionIsactive(v bool)`

SetBInscriptionIsactive sets BInscriptionIsactive field to given value.


### GetTInscriptionChecklistnote

`func (o *InscriptionResponse) GetTInscriptionChecklistnote() string`

GetTInscriptionChecklistnote returns the TInscriptionChecklistnote field if non-nil, zero value otherwise.

### GetTInscriptionChecklistnoteOk

`func (o *InscriptionResponse) GetTInscriptionChecklistnoteOk() (*string, bool)`

GetTInscriptionChecklistnoteOk returns a tuple with the TInscriptionChecklistnote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTInscriptionChecklistnote

`func (o *InscriptionResponse) SetTInscriptionChecklistnote(v string)`

SetTInscriptionChecklistnote sets TInscriptionChecklistnote field to given value.


### GetBInscriptionNew

`func (o *InscriptionResponse) GetBInscriptionNew() bool`

GetBInscriptionNew returns the BInscriptionNew field if non-nil, zero value otherwise.

### GetBInscriptionNewOk

`func (o *InscriptionResponse) GetBInscriptionNewOk() (*bool, bool)`

GetBInscriptionNewOk returns a tuple with the BInscriptionNew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBInscriptionNew

`func (o *InscriptionResponse) SetBInscriptionNew(v bool)`

SetBInscriptionNew sets BInscriptionNew field to given value.


### GetBInscriptionHomeowner

`func (o *InscriptionResponse) GetBInscriptionHomeowner() bool`

GetBInscriptionHomeowner returns the BInscriptionHomeowner field if non-nil, zero value otherwise.

### GetBInscriptionHomeownerOk

`func (o *InscriptionResponse) GetBInscriptionHomeownerOk() (*bool, bool)`

GetBInscriptionHomeownerOk returns a tuple with the BInscriptionHomeowner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBInscriptionHomeowner

`func (o *InscriptionResponse) SetBInscriptionHomeowner(v bool)`

SetBInscriptionHomeowner sets BInscriptionHomeowner field to given value.


### GetBInscriptionArchived

`func (o *InscriptionResponse) GetBInscriptionArchived() bool`

GetBInscriptionArchived returns the BInscriptionArchived field if non-nil, zero value otherwise.

### GetBInscriptionArchivedOk

`func (o *InscriptionResponse) GetBInscriptionArchivedOk() (*bool, bool)`

GetBInscriptionArchivedOk returns a tuple with the BInscriptionArchived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBInscriptionArchived

`func (o *InscriptionResponse) SetBInscriptionArchived(v bool)`

SetBInscriptionArchived sets BInscriptionArchived field to given value.


### GetBInscriptionLitigation

`func (o *InscriptionResponse) GetBInscriptionLitigation() bool`

GetBInscriptionLitigation returns the BInscriptionLitigation field if non-nil, zero value otherwise.

### GetBInscriptionLitigationOk

`func (o *InscriptionResponse) GetBInscriptionLitigationOk() (*bool, bool)`

GetBInscriptionLitigationOk returns a tuple with the BInscriptionLitigation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBInscriptionLitigation

`func (o *InscriptionResponse) SetBInscriptionLitigation(v bool)`

SetBInscriptionLitigation sets BInscriptionLitigation field to given value.


### GetBInscriptionRepossession

`func (o *InscriptionResponse) GetBInscriptionRepossession() bool`

GetBInscriptionRepossession returns the BInscriptionRepossession field if non-nil, zero value otherwise.

### GetBInscriptionRepossessionOk

`func (o *InscriptionResponse) GetBInscriptionRepossessionOk() (*bool, bool)`

GetBInscriptionRepossessionOk returns a tuple with the BInscriptionRepossession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBInscriptionRepossession

`func (o *InscriptionResponse) SetBInscriptionRepossession(v bool)`

SetBInscriptionRepossession sets BInscriptionRepossession field to given value.


### GetBInscriptionIssolicitation

`func (o *InscriptionResponse) GetBInscriptionIssolicitation() bool`

GetBInscriptionIssolicitation returns the BInscriptionIssolicitation field if non-nil, zero value otherwise.

### GetBInscriptionIssolicitationOk

`func (o *InscriptionResponse) GetBInscriptionIssolicitationOk() (*bool, bool)`

GetBInscriptionIssolicitationOk returns a tuple with the BInscriptionIssolicitation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBInscriptionIssolicitation

`func (o *InscriptionResponse) SetBInscriptionIssolicitation(v bool)`

SetBInscriptionIssolicitation sets BInscriptionIssolicitation field to given value.


### GetBInscriptionSalebyowner

`func (o *InscriptionResponse) GetBInscriptionSalebyowner() bool`

GetBInscriptionSalebyowner returns the BInscriptionSalebyowner field if non-nil, zero value otherwise.

### GetBInscriptionSalebyownerOk

`func (o *InscriptionResponse) GetBInscriptionSalebyownerOk() (*bool, bool)`

GetBInscriptionSalebyownerOk returns a tuple with the BInscriptionSalebyowner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBInscriptionSalebyowner

`func (o *InscriptionResponse) SetBInscriptionSalebyowner(v bool)`

SetBInscriptionSalebyowner sets BInscriptionSalebyowner field to given value.


### GetBInscriptionSoldwithoutlegalwarranty

`func (o *InscriptionResponse) GetBInscriptionSoldwithoutlegalwarranty() bool`

GetBInscriptionSoldwithoutlegalwarranty returns the BInscriptionSoldwithoutlegalwarranty field if non-nil, zero value otherwise.

### GetBInscriptionSoldwithoutlegalwarrantyOk

`func (o *InscriptionResponse) GetBInscriptionSoldwithoutlegalwarrantyOk() (*bool, bool)`

GetBInscriptionSoldwithoutlegalwarrantyOk returns a tuple with the BInscriptionSoldwithoutlegalwarranty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBInscriptionSoldwithoutlegalwarranty

`func (o *InscriptionResponse) SetBInscriptionSoldwithoutlegalwarranty(v bool)`

SetBInscriptionSoldwithoutlegalwarranty sets BInscriptionSoldwithoutlegalwarranty field to given value.


### GetIInscriptionConstructionyear

`func (o *InscriptionResponse) GetIInscriptionConstructionyear() int32`

GetIInscriptionConstructionyear returns the IInscriptionConstructionyear field if non-nil, zero value otherwise.

### GetIInscriptionConstructionyearOk

`func (o *InscriptionResponse) GetIInscriptionConstructionyearOk() (*int32, bool)`

GetIInscriptionConstructionyearOk returns a tuple with the IInscriptionConstructionyear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIInscriptionConstructionyear

`func (o *InscriptionResponse) SetIInscriptionConstructionyear(v int32)`

SetIInscriptionConstructionyear sets IInscriptionConstructionyear field to given value.


### GetIInscriptionUnit

`func (o *InscriptionResponse) GetIInscriptionUnit() int32`

GetIInscriptionUnit returns the IInscriptionUnit field if non-nil, zero value otherwise.

### GetIInscriptionUnitOk

`func (o *InscriptionResponse) GetIInscriptionUnitOk() (*int32, bool)`

GetIInscriptionUnitOk returns a tuple with the IInscriptionUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIInscriptionUnit

`func (o *InscriptionResponse) SetIInscriptionUnit(v int32)`

SetIInscriptionUnit sets IInscriptionUnit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


