# InscriptionResponseCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiInscriptionID** | **int32** | The unique ID of the Inscription. | 
**FkiDepartmentID** | Pointer to **int32** | The unique ID of the Department | [optional] 
**SDepartmentNameX** | Pointer to **string** | The Name of the Department in the language of the requester | [optional] 
**FkiRealestateboardID** | **int32** | The unique ID of the Realestateboard | 
**SRealestateboardNameX** | Pointer to **string** | The name of the Realestateboard | [optional] 
**FkiAddressID** | **int32** | The unique ID of the Address | 
**ObjAddress** | Pointer to [**AddressResponseCompound**](AddressResponseCompound.md) |  | [optional] 
**FkiInscriptionbuildingtypeID** | **int32** | The unique ID of the Inscriptionbuildingtype | 
**SInscriptionbuildingtypeNameX** | Pointer to **string** | The name of the Inscriptionbuildingtype in the language of the requester | [optional] 
**FkiInscriptiontypeID** | **int32** | The unique ID of the Inscriptiontype | 
**SInscriptiontypeNameX** | Pointer to **string** | The name of the Inscriptiontype in the language of the requester | [optional] 
**FkiInscriptioncategoryID** | **int32** | The unique ID of the Inscriptioncategory | 
**SInscriptioncategoryNameX** | Pointer to **string** | The name of the Inscriptioncategory in the language of the requester | [optional] 
**EInscriptionStep** | [**FieldEInscriptionStep**](FieldEInscriptionStep.md) |  | 
**EInscriptionResidenceType** | [**FieldEInscriptionResidenceType**](FieldEInscriptionResidenceType.md) |  | 
**SInscriptionCivicend** | **string** | The address civic end of the Inscription | 
**SInscriptionMLS** | Pointer to **string** | The mls of the Inscription | [optional] 
**SInscriptionContract** | **string** | The sale contract number | 
**IInscriptionSellerdeclaration** | **int32** | The seller declaration number of the Inscription | 
**EInscriptionType** | [**FieldEInscriptionType**](FieldEInscriptionType.md) |  | 
**DInscriptionInitialsaleprice** | **string** | The initial sale price of the Inscription | 
**DInscriptionSaleprice** | **string** | The saleprice of the Inscription | 
**DInscriptionRentprice** | **string** | The rent price of the Inscription | 
**EInscriptionRemunerationtype** | [**FieldEInscriptionRemunerationtype**](FieldEInscriptionRemunerationtype.md) |  | 
**EInscriptionRemunerationinscriptorsellertype** | [**FieldEInscriptionRemunerationinscriptorsellertype**](FieldEInscriptionRemunerationinscriptorsellertype.md) |  | 
**EInscriptionRemunerationreferencetype** | [**FieldEInscriptionRemunerationreferencetype**](FieldEInscriptionRemunerationreferencetype.md) |  | 
**EInscriptionRemunerationtotaltype** | [**FieldEInscriptionRemunerationtotaltype**](FieldEInscriptionRemunerationtotaltype.md) |  | 
**DInscriptionRemuneration** | **string** | The remuneration amount of the Inscription | 
**DInscriptionRemunerationinscriptorseller** | **string** | The remuneration amount for the inscriptor or seller of the Inscription | 
**DInscriptionRemunerationreference** | **string** | The remuneration amount for the reference of the Inscription | 
**DInscriptionRemunerationtotal** | **string** | The remuneration amount total of the Inscription | 
**DInscriptionMortgagesold** | **string** | The balande for the mortgage of the Inscription | 
**DtInscriptionDate** | Pointer to **string** | The date of the Inscription | [optional] 
**DtInscriptionCancellationdate** | Pointer to **string** | The cancellation date of the Inscription | [optional] 
**DtInscriptionInitialexpirationdate** | Pointer to **string** | The initial expiration date of the Inscription | [optional] 
**DtInscriptionExpirationdate** | Pointer to **string** | The expiration date of the Inscription | [optional] 
**DtInscriptionNotarydate** | Pointer to **string** | The notary date of the Inscription | [optional] 
**DtInscriptionNotaryentereddate** | Pointer to **string** | The notary entered date of the Inscription | [optional] 
**TInscriptionCadastre** | **string** | The cadastre of the Inscription | 
**BInscriptionReference** | **bool** | Whether if it&#39;s an reference | 
**BInscriptionInspection** | **bool** | Whether the inscription can be acces by an inspector | 
**BInscriptionIsactive** | **bool** | Whether the inscription is active or not | 
**TInscriptionChecklistnote** | **string** | The checklist note of the Inscription | 
**BInscriptionNew** | **bool** | Whether if it&#39;s an new | 
**BInscriptionHomeowner** | **bool** | Whether if it&#39;s an homeowner | 
**BInscriptionArchived** | **bool** | Whether the inscription is archived or not | 
**BInscriptionLitigation** | **bool** | Whether if it&#39;s an litigation | 
**BInscriptionRepossession** | **bool** | Whether if it&#39;s an repossession | 
**BInscriptionIssolicitation** | **bool** | Whether if it&#39;s a solicitation | 
**BInscriptionSalebyowner** | **bool** | Whether if it&#39;s a sale by the owner | 
**BInscriptionSoldwithoutlegalwarranty** | **bool** | Whether if it&#39;s sold without the legal warranty | 
**IInscriptionConstructionyear** | **int32** | The construction year of the Inscription | 
**IInscriptionUnit** | **int32** | The number of unit for the Inscription | 
**ObjAudit** | Pointer to [**CommonAudit**](CommonAudit.md) |  | [optional] 

## Methods

### NewInscriptionResponseCompound

`func NewInscriptionResponseCompound(pkiInscriptionID int32, fkiRealestateboardID int32, fkiAddressID int32, fkiInscriptionbuildingtypeID int32, fkiInscriptiontypeID int32, fkiInscriptioncategoryID int32, eInscriptionStep FieldEInscriptionStep, eInscriptionResidenceType FieldEInscriptionResidenceType, sInscriptionCivicend string, sInscriptionContract string, iInscriptionSellerdeclaration int32, eInscriptionType FieldEInscriptionType, dInscriptionInitialsaleprice string, dInscriptionSaleprice string, dInscriptionRentprice string, eInscriptionRemunerationtype FieldEInscriptionRemunerationtype, eInscriptionRemunerationinscriptorsellertype FieldEInscriptionRemunerationinscriptorsellertype, eInscriptionRemunerationreferencetype FieldEInscriptionRemunerationreferencetype, eInscriptionRemunerationtotaltype FieldEInscriptionRemunerationtotaltype, dInscriptionRemuneration string, dInscriptionRemunerationinscriptorseller string, dInscriptionRemunerationreference string, dInscriptionRemunerationtotal string, dInscriptionMortgagesold string, tInscriptionCadastre string, bInscriptionReference bool, bInscriptionInspection bool, bInscriptionIsactive bool, tInscriptionChecklistnote string, bInscriptionNew bool, bInscriptionHomeowner bool, bInscriptionArchived bool, bInscriptionLitigation bool, bInscriptionRepossession bool, bInscriptionIssolicitation bool, bInscriptionSalebyowner bool, bInscriptionSoldwithoutlegalwarranty bool, iInscriptionConstructionyear int32, iInscriptionUnit int32, ) *InscriptionResponseCompound`

NewInscriptionResponseCompound instantiates a new InscriptionResponseCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInscriptionResponseCompoundWithDefaults

`func NewInscriptionResponseCompoundWithDefaults() *InscriptionResponseCompound`

NewInscriptionResponseCompoundWithDefaults instantiates a new InscriptionResponseCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiInscriptionID

`func (o *InscriptionResponseCompound) GetPkiInscriptionID() int32`

GetPkiInscriptionID returns the PkiInscriptionID field if non-nil, zero value otherwise.

### GetPkiInscriptionIDOk

`func (o *InscriptionResponseCompound) GetPkiInscriptionIDOk() (*int32, bool)`

GetPkiInscriptionIDOk returns a tuple with the PkiInscriptionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiInscriptionID

`func (o *InscriptionResponseCompound) SetPkiInscriptionID(v int32)`

SetPkiInscriptionID sets PkiInscriptionID field to given value.


### GetFkiDepartmentID

`func (o *InscriptionResponseCompound) GetFkiDepartmentID() int32`

GetFkiDepartmentID returns the FkiDepartmentID field if non-nil, zero value otherwise.

### GetFkiDepartmentIDOk

`func (o *InscriptionResponseCompound) GetFkiDepartmentIDOk() (*int32, bool)`

GetFkiDepartmentIDOk returns a tuple with the FkiDepartmentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiDepartmentID

`func (o *InscriptionResponseCompound) SetFkiDepartmentID(v int32)`

SetFkiDepartmentID sets FkiDepartmentID field to given value.

### HasFkiDepartmentID

`func (o *InscriptionResponseCompound) HasFkiDepartmentID() bool`

HasFkiDepartmentID returns a boolean if a field has been set.

### GetSDepartmentNameX

`func (o *InscriptionResponseCompound) GetSDepartmentNameX() string`

GetSDepartmentNameX returns the SDepartmentNameX field if non-nil, zero value otherwise.

### GetSDepartmentNameXOk

`func (o *InscriptionResponseCompound) GetSDepartmentNameXOk() (*string, bool)`

GetSDepartmentNameXOk returns a tuple with the SDepartmentNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSDepartmentNameX

`func (o *InscriptionResponseCompound) SetSDepartmentNameX(v string)`

SetSDepartmentNameX sets SDepartmentNameX field to given value.

### HasSDepartmentNameX

`func (o *InscriptionResponseCompound) HasSDepartmentNameX() bool`

HasSDepartmentNameX returns a boolean if a field has been set.

### GetFkiRealestateboardID

`func (o *InscriptionResponseCompound) GetFkiRealestateboardID() int32`

GetFkiRealestateboardID returns the FkiRealestateboardID field if non-nil, zero value otherwise.

### GetFkiRealestateboardIDOk

`func (o *InscriptionResponseCompound) GetFkiRealestateboardIDOk() (*int32, bool)`

GetFkiRealestateboardIDOk returns a tuple with the FkiRealestateboardID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiRealestateboardID

`func (o *InscriptionResponseCompound) SetFkiRealestateboardID(v int32)`

SetFkiRealestateboardID sets FkiRealestateboardID field to given value.


### GetSRealestateboardNameX

`func (o *InscriptionResponseCompound) GetSRealestateboardNameX() string`

GetSRealestateboardNameX returns the SRealestateboardNameX field if non-nil, zero value otherwise.

### GetSRealestateboardNameXOk

`func (o *InscriptionResponseCompound) GetSRealestateboardNameXOk() (*string, bool)`

GetSRealestateboardNameXOk returns a tuple with the SRealestateboardNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSRealestateboardNameX

`func (o *InscriptionResponseCompound) SetSRealestateboardNameX(v string)`

SetSRealestateboardNameX sets SRealestateboardNameX field to given value.

### HasSRealestateboardNameX

`func (o *InscriptionResponseCompound) HasSRealestateboardNameX() bool`

HasSRealestateboardNameX returns a boolean if a field has been set.

### GetFkiAddressID

`func (o *InscriptionResponseCompound) GetFkiAddressID() int32`

GetFkiAddressID returns the FkiAddressID field if non-nil, zero value otherwise.

### GetFkiAddressIDOk

`func (o *InscriptionResponseCompound) GetFkiAddressIDOk() (*int32, bool)`

GetFkiAddressIDOk returns a tuple with the FkiAddressID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiAddressID

`func (o *InscriptionResponseCompound) SetFkiAddressID(v int32)`

SetFkiAddressID sets FkiAddressID field to given value.


### GetObjAddress

`func (o *InscriptionResponseCompound) GetObjAddress() AddressResponseCompound`

GetObjAddress returns the ObjAddress field if non-nil, zero value otherwise.

### GetObjAddressOk

`func (o *InscriptionResponseCompound) GetObjAddressOk() (*AddressResponseCompound, bool)`

GetObjAddressOk returns a tuple with the ObjAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjAddress

`func (o *InscriptionResponseCompound) SetObjAddress(v AddressResponseCompound)`

SetObjAddress sets ObjAddress field to given value.

### HasObjAddress

`func (o *InscriptionResponseCompound) HasObjAddress() bool`

HasObjAddress returns a boolean if a field has been set.

### GetFkiInscriptionbuildingtypeID

`func (o *InscriptionResponseCompound) GetFkiInscriptionbuildingtypeID() int32`

GetFkiInscriptionbuildingtypeID returns the FkiInscriptionbuildingtypeID field if non-nil, zero value otherwise.

### GetFkiInscriptionbuildingtypeIDOk

`func (o *InscriptionResponseCompound) GetFkiInscriptionbuildingtypeIDOk() (*int32, bool)`

GetFkiInscriptionbuildingtypeIDOk returns a tuple with the FkiInscriptionbuildingtypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiInscriptionbuildingtypeID

`func (o *InscriptionResponseCompound) SetFkiInscriptionbuildingtypeID(v int32)`

SetFkiInscriptionbuildingtypeID sets FkiInscriptionbuildingtypeID field to given value.


### GetSInscriptionbuildingtypeNameX

`func (o *InscriptionResponseCompound) GetSInscriptionbuildingtypeNameX() string`

GetSInscriptionbuildingtypeNameX returns the SInscriptionbuildingtypeNameX field if non-nil, zero value otherwise.

### GetSInscriptionbuildingtypeNameXOk

`func (o *InscriptionResponseCompound) GetSInscriptionbuildingtypeNameXOk() (*string, bool)`

GetSInscriptionbuildingtypeNameXOk returns a tuple with the SInscriptionbuildingtypeNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSInscriptionbuildingtypeNameX

`func (o *InscriptionResponseCompound) SetSInscriptionbuildingtypeNameX(v string)`

SetSInscriptionbuildingtypeNameX sets SInscriptionbuildingtypeNameX field to given value.

### HasSInscriptionbuildingtypeNameX

`func (o *InscriptionResponseCompound) HasSInscriptionbuildingtypeNameX() bool`

HasSInscriptionbuildingtypeNameX returns a boolean if a field has been set.

### GetFkiInscriptiontypeID

`func (o *InscriptionResponseCompound) GetFkiInscriptiontypeID() int32`

GetFkiInscriptiontypeID returns the FkiInscriptiontypeID field if non-nil, zero value otherwise.

### GetFkiInscriptiontypeIDOk

`func (o *InscriptionResponseCompound) GetFkiInscriptiontypeIDOk() (*int32, bool)`

GetFkiInscriptiontypeIDOk returns a tuple with the FkiInscriptiontypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiInscriptiontypeID

`func (o *InscriptionResponseCompound) SetFkiInscriptiontypeID(v int32)`

SetFkiInscriptiontypeID sets FkiInscriptiontypeID field to given value.


### GetSInscriptiontypeNameX

`func (o *InscriptionResponseCompound) GetSInscriptiontypeNameX() string`

GetSInscriptiontypeNameX returns the SInscriptiontypeNameX field if non-nil, zero value otherwise.

### GetSInscriptiontypeNameXOk

`func (o *InscriptionResponseCompound) GetSInscriptiontypeNameXOk() (*string, bool)`

GetSInscriptiontypeNameXOk returns a tuple with the SInscriptiontypeNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSInscriptiontypeNameX

`func (o *InscriptionResponseCompound) SetSInscriptiontypeNameX(v string)`

SetSInscriptiontypeNameX sets SInscriptiontypeNameX field to given value.

### HasSInscriptiontypeNameX

`func (o *InscriptionResponseCompound) HasSInscriptiontypeNameX() bool`

HasSInscriptiontypeNameX returns a boolean if a field has been set.

### GetFkiInscriptioncategoryID

`func (o *InscriptionResponseCompound) GetFkiInscriptioncategoryID() int32`

GetFkiInscriptioncategoryID returns the FkiInscriptioncategoryID field if non-nil, zero value otherwise.

### GetFkiInscriptioncategoryIDOk

`func (o *InscriptionResponseCompound) GetFkiInscriptioncategoryIDOk() (*int32, bool)`

GetFkiInscriptioncategoryIDOk returns a tuple with the FkiInscriptioncategoryID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiInscriptioncategoryID

`func (o *InscriptionResponseCompound) SetFkiInscriptioncategoryID(v int32)`

SetFkiInscriptioncategoryID sets FkiInscriptioncategoryID field to given value.


### GetSInscriptioncategoryNameX

`func (o *InscriptionResponseCompound) GetSInscriptioncategoryNameX() string`

GetSInscriptioncategoryNameX returns the SInscriptioncategoryNameX field if non-nil, zero value otherwise.

### GetSInscriptioncategoryNameXOk

`func (o *InscriptionResponseCompound) GetSInscriptioncategoryNameXOk() (*string, bool)`

GetSInscriptioncategoryNameXOk returns a tuple with the SInscriptioncategoryNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSInscriptioncategoryNameX

`func (o *InscriptionResponseCompound) SetSInscriptioncategoryNameX(v string)`

SetSInscriptioncategoryNameX sets SInscriptioncategoryNameX field to given value.

### HasSInscriptioncategoryNameX

`func (o *InscriptionResponseCompound) HasSInscriptioncategoryNameX() bool`

HasSInscriptioncategoryNameX returns a boolean if a field has been set.

### GetEInscriptionStep

`func (o *InscriptionResponseCompound) GetEInscriptionStep() FieldEInscriptionStep`

GetEInscriptionStep returns the EInscriptionStep field if non-nil, zero value otherwise.

### GetEInscriptionStepOk

`func (o *InscriptionResponseCompound) GetEInscriptionStepOk() (*FieldEInscriptionStep, bool)`

GetEInscriptionStepOk returns a tuple with the EInscriptionStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEInscriptionStep

`func (o *InscriptionResponseCompound) SetEInscriptionStep(v FieldEInscriptionStep)`

SetEInscriptionStep sets EInscriptionStep field to given value.


### GetEInscriptionResidenceType

`func (o *InscriptionResponseCompound) GetEInscriptionResidenceType() FieldEInscriptionResidenceType`

GetEInscriptionResidenceType returns the EInscriptionResidenceType field if non-nil, zero value otherwise.

### GetEInscriptionResidenceTypeOk

`func (o *InscriptionResponseCompound) GetEInscriptionResidenceTypeOk() (*FieldEInscriptionResidenceType, bool)`

GetEInscriptionResidenceTypeOk returns a tuple with the EInscriptionResidenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEInscriptionResidenceType

`func (o *InscriptionResponseCompound) SetEInscriptionResidenceType(v FieldEInscriptionResidenceType)`

SetEInscriptionResidenceType sets EInscriptionResidenceType field to given value.


### GetSInscriptionCivicend

`func (o *InscriptionResponseCompound) GetSInscriptionCivicend() string`

GetSInscriptionCivicend returns the SInscriptionCivicend field if non-nil, zero value otherwise.

### GetSInscriptionCivicendOk

`func (o *InscriptionResponseCompound) GetSInscriptionCivicendOk() (*string, bool)`

GetSInscriptionCivicendOk returns a tuple with the SInscriptionCivicend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSInscriptionCivicend

`func (o *InscriptionResponseCompound) SetSInscriptionCivicend(v string)`

SetSInscriptionCivicend sets SInscriptionCivicend field to given value.


### GetSInscriptionMLS

`func (o *InscriptionResponseCompound) GetSInscriptionMLS() string`

GetSInscriptionMLS returns the SInscriptionMLS field if non-nil, zero value otherwise.

### GetSInscriptionMLSOk

`func (o *InscriptionResponseCompound) GetSInscriptionMLSOk() (*string, bool)`

GetSInscriptionMLSOk returns a tuple with the SInscriptionMLS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSInscriptionMLS

`func (o *InscriptionResponseCompound) SetSInscriptionMLS(v string)`

SetSInscriptionMLS sets SInscriptionMLS field to given value.

### HasSInscriptionMLS

`func (o *InscriptionResponseCompound) HasSInscriptionMLS() bool`

HasSInscriptionMLS returns a boolean if a field has been set.

### GetSInscriptionContract

`func (o *InscriptionResponseCompound) GetSInscriptionContract() string`

GetSInscriptionContract returns the SInscriptionContract field if non-nil, zero value otherwise.

### GetSInscriptionContractOk

`func (o *InscriptionResponseCompound) GetSInscriptionContractOk() (*string, bool)`

GetSInscriptionContractOk returns a tuple with the SInscriptionContract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSInscriptionContract

`func (o *InscriptionResponseCompound) SetSInscriptionContract(v string)`

SetSInscriptionContract sets SInscriptionContract field to given value.


### GetIInscriptionSellerdeclaration

`func (o *InscriptionResponseCompound) GetIInscriptionSellerdeclaration() int32`

GetIInscriptionSellerdeclaration returns the IInscriptionSellerdeclaration field if non-nil, zero value otherwise.

### GetIInscriptionSellerdeclarationOk

`func (o *InscriptionResponseCompound) GetIInscriptionSellerdeclarationOk() (*int32, bool)`

GetIInscriptionSellerdeclarationOk returns a tuple with the IInscriptionSellerdeclaration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIInscriptionSellerdeclaration

`func (o *InscriptionResponseCompound) SetIInscriptionSellerdeclaration(v int32)`

SetIInscriptionSellerdeclaration sets IInscriptionSellerdeclaration field to given value.


### GetEInscriptionType

`func (o *InscriptionResponseCompound) GetEInscriptionType() FieldEInscriptionType`

GetEInscriptionType returns the EInscriptionType field if non-nil, zero value otherwise.

### GetEInscriptionTypeOk

`func (o *InscriptionResponseCompound) GetEInscriptionTypeOk() (*FieldEInscriptionType, bool)`

GetEInscriptionTypeOk returns a tuple with the EInscriptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEInscriptionType

`func (o *InscriptionResponseCompound) SetEInscriptionType(v FieldEInscriptionType)`

SetEInscriptionType sets EInscriptionType field to given value.


### GetDInscriptionInitialsaleprice

`func (o *InscriptionResponseCompound) GetDInscriptionInitialsaleprice() string`

GetDInscriptionInitialsaleprice returns the DInscriptionInitialsaleprice field if non-nil, zero value otherwise.

### GetDInscriptionInitialsalepriceOk

`func (o *InscriptionResponseCompound) GetDInscriptionInitialsalepriceOk() (*string, bool)`

GetDInscriptionInitialsalepriceOk returns a tuple with the DInscriptionInitialsaleprice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDInscriptionInitialsaleprice

`func (o *InscriptionResponseCompound) SetDInscriptionInitialsaleprice(v string)`

SetDInscriptionInitialsaleprice sets DInscriptionInitialsaleprice field to given value.


### GetDInscriptionSaleprice

`func (o *InscriptionResponseCompound) GetDInscriptionSaleprice() string`

GetDInscriptionSaleprice returns the DInscriptionSaleprice field if non-nil, zero value otherwise.

### GetDInscriptionSalepriceOk

`func (o *InscriptionResponseCompound) GetDInscriptionSalepriceOk() (*string, bool)`

GetDInscriptionSalepriceOk returns a tuple with the DInscriptionSaleprice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDInscriptionSaleprice

`func (o *InscriptionResponseCompound) SetDInscriptionSaleprice(v string)`

SetDInscriptionSaleprice sets DInscriptionSaleprice field to given value.


### GetDInscriptionRentprice

`func (o *InscriptionResponseCompound) GetDInscriptionRentprice() string`

GetDInscriptionRentprice returns the DInscriptionRentprice field if non-nil, zero value otherwise.

### GetDInscriptionRentpriceOk

`func (o *InscriptionResponseCompound) GetDInscriptionRentpriceOk() (*string, bool)`

GetDInscriptionRentpriceOk returns a tuple with the DInscriptionRentprice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDInscriptionRentprice

`func (o *InscriptionResponseCompound) SetDInscriptionRentprice(v string)`

SetDInscriptionRentprice sets DInscriptionRentprice field to given value.


### GetEInscriptionRemunerationtype

`func (o *InscriptionResponseCompound) GetEInscriptionRemunerationtype() FieldEInscriptionRemunerationtype`

GetEInscriptionRemunerationtype returns the EInscriptionRemunerationtype field if non-nil, zero value otherwise.

### GetEInscriptionRemunerationtypeOk

`func (o *InscriptionResponseCompound) GetEInscriptionRemunerationtypeOk() (*FieldEInscriptionRemunerationtype, bool)`

GetEInscriptionRemunerationtypeOk returns a tuple with the EInscriptionRemunerationtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEInscriptionRemunerationtype

`func (o *InscriptionResponseCompound) SetEInscriptionRemunerationtype(v FieldEInscriptionRemunerationtype)`

SetEInscriptionRemunerationtype sets EInscriptionRemunerationtype field to given value.


### GetEInscriptionRemunerationinscriptorsellertype

`func (o *InscriptionResponseCompound) GetEInscriptionRemunerationinscriptorsellertype() FieldEInscriptionRemunerationinscriptorsellertype`

GetEInscriptionRemunerationinscriptorsellertype returns the EInscriptionRemunerationinscriptorsellertype field if non-nil, zero value otherwise.

### GetEInscriptionRemunerationinscriptorsellertypeOk

`func (o *InscriptionResponseCompound) GetEInscriptionRemunerationinscriptorsellertypeOk() (*FieldEInscriptionRemunerationinscriptorsellertype, bool)`

GetEInscriptionRemunerationinscriptorsellertypeOk returns a tuple with the EInscriptionRemunerationinscriptorsellertype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEInscriptionRemunerationinscriptorsellertype

`func (o *InscriptionResponseCompound) SetEInscriptionRemunerationinscriptorsellertype(v FieldEInscriptionRemunerationinscriptorsellertype)`

SetEInscriptionRemunerationinscriptorsellertype sets EInscriptionRemunerationinscriptorsellertype field to given value.


### GetEInscriptionRemunerationreferencetype

`func (o *InscriptionResponseCompound) GetEInscriptionRemunerationreferencetype() FieldEInscriptionRemunerationreferencetype`

GetEInscriptionRemunerationreferencetype returns the EInscriptionRemunerationreferencetype field if non-nil, zero value otherwise.

### GetEInscriptionRemunerationreferencetypeOk

`func (o *InscriptionResponseCompound) GetEInscriptionRemunerationreferencetypeOk() (*FieldEInscriptionRemunerationreferencetype, bool)`

GetEInscriptionRemunerationreferencetypeOk returns a tuple with the EInscriptionRemunerationreferencetype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEInscriptionRemunerationreferencetype

`func (o *InscriptionResponseCompound) SetEInscriptionRemunerationreferencetype(v FieldEInscriptionRemunerationreferencetype)`

SetEInscriptionRemunerationreferencetype sets EInscriptionRemunerationreferencetype field to given value.


### GetEInscriptionRemunerationtotaltype

`func (o *InscriptionResponseCompound) GetEInscriptionRemunerationtotaltype() FieldEInscriptionRemunerationtotaltype`

GetEInscriptionRemunerationtotaltype returns the EInscriptionRemunerationtotaltype field if non-nil, zero value otherwise.

### GetEInscriptionRemunerationtotaltypeOk

`func (o *InscriptionResponseCompound) GetEInscriptionRemunerationtotaltypeOk() (*FieldEInscriptionRemunerationtotaltype, bool)`

GetEInscriptionRemunerationtotaltypeOk returns a tuple with the EInscriptionRemunerationtotaltype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEInscriptionRemunerationtotaltype

`func (o *InscriptionResponseCompound) SetEInscriptionRemunerationtotaltype(v FieldEInscriptionRemunerationtotaltype)`

SetEInscriptionRemunerationtotaltype sets EInscriptionRemunerationtotaltype field to given value.


### GetDInscriptionRemuneration

`func (o *InscriptionResponseCompound) GetDInscriptionRemuneration() string`

GetDInscriptionRemuneration returns the DInscriptionRemuneration field if non-nil, zero value otherwise.

### GetDInscriptionRemunerationOk

`func (o *InscriptionResponseCompound) GetDInscriptionRemunerationOk() (*string, bool)`

GetDInscriptionRemunerationOk returns a tuple with the DInscriptionRemuneration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDInscriptionRemuneration

`func (o *InscriptionResponseCompound) SetDInscriptionRemuneration(v string)`

SetDInscriptionRemuneration sets DInscriptionRemuneration field to given value.


### GetDInscriptionRemunerationinscriptorseller

`func (o *InscriptionResponseCompound) GetDInscriptionRemunerationinscriptorseller() string`

GetDInscriptionRemunerationinscriptorseller returns the DInscriptionRemunerationinscriptorseller field if non-nil, zero value otherwise.

### GetDInscriptionRemunerationinscriptorsellerOk

`func (o *InscriptionResponseCompound) GetDInscriptionRemunerationinscriptorsellerOk() (*string, bool)`

GetDInscriptionRemunerationinscriptorsellerOk returns a tuple with the DInscriptionRemunerationinscriptorseller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDInscriptionRemunerationinscriptorseller

`func (o *InscriptionResponseCompound) SetDInscriptionRemunerationinscriptorseller(v string)`

SetDInscriptionRemunerationinscriptorseller sets DInscriptionRemunerationinscriptorseller field to given value.


### GetDInscriptionRemunerationreference

`func (o *InscriptionResponseCompound) GetDInscriptionRemunerationreference() string`

GetDInscriptionRemunerationreference returns the DInscriptionRemunerationreference field if non-nil, zero value otherwise.

### GetDInscriptionRemunerationreferenceOk

`func (o *InscriptionResponseCompound) GetDInscriptionRemunerationreferenceOk() (*string, bool)`

GetDInscriptionRemunerationreferenceOk returns a tuple with the DInscriptionRemunerationreference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDInscriptionRemunerationreference

`func (o *InscriptionResponseCompound) SetDInscriptionRemunerationreference(v string)`

SetDInscriptionRemunerationreference sets DInscriptionRemunerationreference field to given value.


### GetDInscriptionRemunerationtotal

`func (o *InscriptionResponseCompound) GetDInscriptionRemunerationtotal() string`

GetDInscriptionRemunerationtotal returns the DInscriptionRemunerationtotal field if non-nil, zero value otherwise.

### GetDInscriptionRemunerationtotalOk

`func (o *InscriptionResponseCompound) GetDInscriptionRemunerationtotalOk() (*string, bool)`

GetDInscriptionRemunerationtotalOk returns a tuple with the DInscriptionRemunerationtotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDInscriptionRemunerationtotal

`func (o *InscriptionResponseCompound) SetDInscriptionRemunerationtotal(v string)`

SetDInscriptionRemunerationtotal sets DInscriptionRemunerationtotal field to given value.


### GetDInscriptionMortgagesold

`func (o *InscriptionResponseCompound) GetDInscriptionMortgagesold() string`

GetDInscriptionMortgagesold returns the DInscriptionMortgagesold field if non-nil, zero value otherwise.

### GetDInscriptionMortgagesoldOk

`func (o *InscriptionResponseCompound) GetDInscriptionMortgagesoldOk() (*string, bool)`

GetDInscriptionMortgagesoldOk returns a tuple with the DInscriptionMortgagesold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDInscriptionMortgagesold

`func (o *InscriptionResponseCompound) SetDInscriptionMortgagesold(v string)`

SetDInscriptionMortgagesold sets DInscriptionMortgagesold field to given value.


### GetDtInscriptionDate

`func (o *InscriptionResponseCompound) GetDtInscriptionDate() string`

GetDtInscriptionDate returns the DtInscriptionDate field if non-nil, zero value otherwise.

### GetDtInscriptionDateOk

`func (o *InscriptionResponseCompound) GetDtInscriptionDateOk() (*string, bool)`

GetDtInscriptionDateOk returns a tuple with the DtInscriptionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtInscriptionDate

`func (o *InscriptionResponseCompound) SetDtInscriptionDate(v string)`

SetDtInscriptionDate sets DtInscriptionDate field to given value.

### HasDtInscriptionDate

`func (o *InscriptionResponseCompound) HasDtInscriptionDate() bool`

HasDtInscriptionDate returns a boolean if a field has been set.

### GetDtInscriptionCancellationdate

`func (o *InscriptionResponseCompound) GetDtInscriptionCancellationdate() string`

GetDtInscriptionCancellationdate returns the DtInscriptionCancellationdate field if non-nil, zero value otherwise.

### GetDtInscriptionCancellationdateOk

`func (o *InscriptionResponseCompound) GetDtInscriptionCancellationdateOk() (*string, bool)`

GetDtInscriptionCancellationdateOk returns a tuple with the DtInscriptionCancellationdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtInscriptionCancellationdate

`func (o *InscriptionResponseCompound) SetDtInscriptionCancellationdate(v string)`

SetDtInscriptionCancellationdate sets DtInscriptionCancellationdate field to given value.

### HasDtInscriptionCancellationdate

`func (o *InscriptionResponseCompound) HasDtInscriptionCancellationdate() bool`

HasDtInscriptionCancellationdate returns a boolean if a field has been set.

### GetDtInscriptionInitialexpirationdate

`func (o *InscriptionResponseCompound) GetDtInscriptionInitialexpirationdate() string`

GetDtInscriptionInitialexpirationdate returns the DtInscriptionInitialexpirationdate field if non-nil, zero value otherwise.

### GetDtInscriptionInitialexpirationdateOk

`func (o *InscriptionResponseCompound) GetDtInscriptionInitialexpirationdateOk() (*string, bool)`

GetDtInscriptionInitialexpirationdateOk returns a tuple with the DtInscriptionInitialexpirationdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtInscriptionInitialexpirationdate

`func (o *InscriptionResponseCompound) SetDtInscriptionInitialexpirationdate(v string)`

SetDtInscriptionInitialexpirationdate sets DtInscriptionInitialexpirationdate field to given value.

### HasDtInscriptionInitialexpirationdate

`func (o *InscriptionResponseCompound) HasDtInscriptionInitialexpirationdate() bool`

HasDtInscriptionInitialexpirationdate returns a boolean if a field has been set.

### GetDtInscriptionExpirationdate

`func (o *InscriptionResponseCompound) GetDtInscriptionExpirationdate() string`

GetDtInscriptionExpirationdate returns the DtInscriptionExpirationdate field if non-nil, zero value otherwise.

### GetDtInscriptionExpirationdateOk

`func (o *InscriptionResponseCompound) GetDtInscriptionExpirationdateOk() (*string, bool)`

GetDtInscriptionExpirationdateOk returns a tuple with the DtInscriptionExpirationdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtInscriptionExpirationdate

`func (o *InscriptionResponseCompound) SetDtInscriptionExpirationdate(v string)`

SetDtInscriptionExpirationdate sets DtInscriptionExpirationdate field to given value.

### HasDtInscriptionExpirationdate

`func (o *InscriptionResponseCompound) HasDtInscriptionExpirationdate() bool`

HasDtInscriptionExpirationdate returns a boolean if a field has been set.

### GetDtInscriptionNotarydate

`func (o *InscriptionResponseCompound) GetDtInscriptionNotarydate() string`

GetDtInscriptionNotarydate returns the DtInscriptionNotarydate field if non-nil, zero value otherwise.

### GetDtInscriptionNotarydateOk

`func (o *InscriptionResponseCompound) GetDtInscriptionNotarydateOk() (*string, bool)`

GetDtInscriptionNotarydateOk returns a tuple with the DtInscriptionNotarydate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtInscriptionNotarydate

`func (o *InscriptionResponseCompound) SetDtInscriptionNotarydate(v string)`

SetDtInscriptionNotarydate sets DtInscriptionNotarydate field to given value.

### HasDtInscriptionNotarydate

`func (o *InscriptionResponseCompound) HasDtInscriptionNotarydate() bool`

HasDtInscriptionNotarydate returns a boolean if a field has been set.

### GetDtInscriptionNotaryentereddate

`func (o *InscriptionResponseCompound) GetDtInscriptionNotaryentereddate() string`

GetDtInscriptionNotaryentereddate returns the DtInscriptionNotaryentereddate field if non-nil, zero value otherwise.

### GetDtInscriptionNotaryentereddateOk

`func (o *InscriptionResponseCompound) GetDtInscriptionNotaryentereddateOk() (*string, bool)`

GetDtInscriptionNotaryentereddateOk returns a tuple with the DtInscriptionNotaryentereddate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtInscriptionNotaryentereddate

`func (o *InscriptionResponseCompound) SetDtInscriptionNotaryentereddate(v string)`

SetDtInscriptionNotaryentereddate sets DtInscriptionNotaryentereddate field to given value.

### HasDtInscriptionNotaryentereddate

`func (o *InscriptionResponseCompound) HasDtInscriptionNotaryentereddate() bool`

HasDtInscriptionNotaryentereddate returns a boolean if a field has been set.

### GetTInscriptionCadastre

`func (o *InscriptionResponseCompound) GetTInscriptionCadastre() string`

GetTInscriptionCadastre returns the TInscriptionCadastre field if non-nil, zero value otherwise.

### GetTInscriptionCadastreOk

`func (o *InscriptionResponseCompound) GetTInscriptionCadastreOk() (*string, bool)`

GetTInscriptionCadastreOk returns a tuple with the TInscriptionCadastre field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTInscriptionCadastre

`func (o *InscriptionResponseCompound) SetTInscriptionCadastre(v string)`

SetTInscriptionCadastre sets TInscriptionCadastre field to given value.


### GetBInscriptionReference

`func (o *InscriptionResponseCompound) GetBInscriptionReference() bool`

GetBInscriptionReference returns the BInscriptionReference field if non-nil, zero value otherwise.

### GetBInscriptionReferenceOk

`func (o *InscriptionResponseCompound) GetBInscriptionReferenceOk() (*bool, bool)`

GetBInscriptionReferenceOk returns a tuple with the BInscriptionReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBInscriptionReference

`func (o *InscriptionResponseCompound) SetBInscriptionReference(v bool)`

SetBInscriptionReference sets BInscriptionReference field to given value.


### GetBInscriptionInspection

`func (o *InscriptionResponseCompound) GetBInscriptionInspection() bool`

GetBInscriptionInspection returns the BInscriptionInspection field if non-nil, zero value otherwise.

### GetBInscriptionInspectionOk

`func (o *InscriptionResponseCompound) GetBInscriptionInspectionOk() (*bool, bool)`

GetBInscriptionInspectionOk returns a tuple with the BInscriptionInspection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBInscriptionInspection

`func (o *InscriptionResponseCompound) SetBInscriptionInspection(v bool)`

SetBInscriptionInspection sets BInscriptionInspection field to given value.


### GetBInscriptionIsactive

`func (o *InscriptionResponseCompound) GetBInscriptionIsactive() bool`

GetBInscriptionIsactive returns the BInscriptionIsactive field if non-nil, zero value otherwise.

### GetBInscriptionIsactiveOk

`func (o *InscriptionResponseCompound) GetBInscriptionIsactiveOk() (*bool, bool)`

GetBInscriptionIsactiveOk returns a tuple with the BInscriptionIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBInscriptionIsactive

`func (o *InscriptionResponseCompound) SetBInscriptionIsactive(v bool)`

SetBInscriptionIsactive sets BInscriptionIsactive field to given value.


### GetTInscriptionChecklistnote

`func (o *InscriptionResponseCompound) GetTInscriptionChecklistnote() string`

GetTInscriptionChecklistnote returns the TInscriptionChecklistnote field if non-nil, zero value otherwise.

### GetTInscriptionChecklistnoteOk

`func (o *InscriptionResponseCompound) GetTInscriptionChecklistnoteOk() (*string, bool)`

GetTInscriptionChecklistnoteOk returns a tuple with the TInscriptionChecklistnote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTInscriptionChecklistnote

`func (o *InscriptionResponseCompound) SetTInscriptionChecklistnote(v string)`

SetTInscriptionChecklistnote sets TInscriptionChecklistnote field to given value.


### GetBInscriptionNew

`func (o *InscriptionResponseCompound) GetBInscriptionNew() bool`

GetBInscriptionNew returns the BInscriptionNew field if non-nil, zero value otherwise.

### GetBInscriptionNewOk

`func (o *InscriptionResponseCompound) GetBInscriptionNewOk() (*bool, bool)`

GetBInscriptionNewOk returns a tuple with the BInscriptionNew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBInscriptionNew

`func (o *InscriptionResponseCompound) SetBInscriptionNew(v bool)`

SetBInscriptionNew sets BInscriptionNew field to given value.


### GetBInscriptionHomeowner

`func (o *InscriptionResponseCompound) GetBInscriptionHomeowner() bool`

GetBInscriptionHomeowner returns the BInscriptionHomeowner field if non-nil, zero value otherwise.

### GetBInscriptionHomeownerOk

`func (o *InscriptionResponseCompound) GetBInscriptionHomeownerOk() (*bool, bool)`

GetBInscriptionHomeownerOk returns a tuple with the BInscriptionHomeowner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBInscriptionHomeowner

`func (o *InscriptionResponseCompound) SetBInscriptionHomeowner(v bool)`

SetBInscriptionHomeowner sets BInscriptionHomeowner field to given value.


### GetBInscriptionArchived

`func (o *InscriptionResponseCompound) GetBInscriptionArchived() bool`

GetBInscriptionArchived returns the BInscriptionArchived field if non-nil, zero value otherwise.

### GetBInscriptionArchivedOk

`func (o *InscriptionResponseCompound) GetBInscriptionArchivedOk() (*bool, bool)`

GetBInscriptionArchivedOk returns a tuple with the BInscriptionArchived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBInscriptionArchived

`func (o *InscriptionResponseCompound) SetBInscriptionArchived(v bool)`

SetBInscriptionArchived sets BInscriptionArchived field to given value.


### GetBInscriptionLitigation

`func (o *InscriptionResponseCompound) GetBInscriptionLitigation() bool`

GetBInscriptionLitigation returns the BInscriptionLitigation field if non-nil, zero value otherwise.

### GetBInscriptionLitigationOk

`func (o *InscriptionResponseCompound) GetBInscriptionLitigationOk() (*bool, bool)`

GetBInscriptionLitigationOk returns a tuple with the BInscriptionLitigation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBInscriptionLitigation

`func (o *InscriptionResponseCompound) SetBInscriptionLitigation(v bool)`

SetBInscriptionLitigation sets BInscriptionLitigation field to given value.


### GetBInscriptionRepossession

`func (o *InscriptionResponseCompound) GetBInscriptionRepossession() bool`

GetBInscriptionRepossession returns the BInscriptionRepossession field if non-nil, zero value otherwise.

### GetBInscriptionRepossessionOk

`func (o *InscriptionResponseCompound) GetBInscriptionRepossessionOk() (*bool, bool)`

GetBInscriptionRepossessionOk returns a tuple with the BInscriptionRepossession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBInscriptionRepossession

`func (o *InscriptionResponseCompound) SetBInscriptionRepossession(v bool)`

SetBInscriptionRepossession sets BInscriptionRepossession field to given value.


### GetBInscriptionIssolicitation

`func (o *InscriptionResponseCompound) GetBInscriptionIssolicitation() bool`

GetBInscriptionIssolicitation returns the BInscriptionIssolicitation field if non-nil, zero value otherwise.

### GetBInscriptionIssolicitationOk

`func (o *InscriptionResponseCompound) GetBInscriptionIssolicitationOk() (*bool, bool)`

GetBInscriptionIssolicitationOk returns a tuple with the BInscriptionIssolicitation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBInscriptionIssolicitation

`func (o *InscriptionResponseCompound) SetBInscriptionIssolicitation(v bool)`

SetBInscriptionIssolicitation sets BInscriptionIssolicitation field to given value.


### GetBInscriptionSalebyowner

`func (o *InscriptionResponseCompound) GetBInscriptionSalebyowner() bool`

GetBInscriptionSalebyowner returns the BInscriptionSalebyowner field if non-nil, zero value otherwise.

### GetBInscriptionSalebyownerOk

`func (o *InscriptionResponseCompound) GetBInscriptionSalebyownerOk() (*bool, bool)`

GetBInscriptionSalebyownerOk returns a tuple with the BInscriptionSalebyowner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBInscriptionSalebyowner

`func (o *InscriptionResponseCompound) SetBInscriptionSalebyowner(v bool)`

SetBInscriptionSalebyowner sets BInscriptionSalebyowner field to given value.


### GetBInscriptionSoldwithoutlegalwarranty

`func (o *InscriptionResponseCompound) GetBInscriptionSoldwithoutlegalwarranty() bool`

GetBInscriptionSoldwithoutlegalwarranty returns the BInscriptionSoldwithoutlegalwarranty field if non-nil, zero value otherwise.

### GetBInscriptionSoldwithoutlegalwarrantyOk

`func (o *InscriptionResponseCompound) GetBInscriptionSoldwithoutlegalwarrantyOk() (*bool, bool)`

GetBInscriptionSoldwithoutlegalwarrantyOk returns a tuple with the BInscriptionSoldwithoutlegalwarranty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBInscriptionSoldwithoutlegalwarranty

`func (o *InscriptionResponseCompound) SetBInscriptionSoldwithoutlegalwarranty(v bool)`

SetBInscriptionSoldwithoutlegalwarranty sets BInscriptionSoldwithoutlegalwarranty field to given value.


### GetIInscriptionConstructionyear

`func (o *InscriptionResponseCompound) GetIInscriptionConstructionyear() int32`

GetIInscriptionConstructionyear returns the IInscriptionConstructionyear field if non-nil, zero value otherwise.

### GetIInscriptionConstructionyearOk

`func (o *InscriptionResponseCompound) GetIInscriptionConstructionyearOk() (*int32, bool)`

GetIInscriptionConstructionyearOk returns a tuple with the IInscriptionConstructionyear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIInscriptionConstructionyear

`func (o *InscriptionResponseCompound) SetIInscriptionConstructionyear(v int32)`

SetIInscriptionConstructionyear sets IInscriptionConstructionyear field to given value.


### GetIInscriptionUnit

`func (o *InscriptionResponseCompound) GetIInscriptionUnit() int32`

GetIInscriptionUnit returns the IInscriptionUnit field if non-nil, zero value otherwise.

### GetIInscriptionUnitOk

`func (o *InscriptionResponseCompound) GetIInscriptionUnitOk() (*int32, bool)`

GetIInscriptionUnitOk returns a tuple with the IInscriptionUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIInscriptionUnit

`func (o *InscriptionResponseCompound) SetIInscriptionUnit(v int32)`

SetIInscriptionUnit sets IInscriptionUnit field to given value.


### GetObjAudit

`func (o *InscriptionResponseCompound) GetObjAudit() CommonAudit`

GetObjAudit returns the ObjAudit field if non-nil, zero value otherwise.

### GetObjAuditOk

`func (o *InscriptionResponseCompound) GetObjAuditOk() (*CommonAudit, bool)`

GetObjAuditOk returns a tuple with the ObjAudit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjAudit

`func (o *InscriptionResponseCompound) SetObjAudit(v CommonAudit)`

SetObjAudit sets ObjAudit field to given value.

### HasObjAudit

`func (o *InscriptionResponseCompound) HasObjAudit() bool`

HasObjAudit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


