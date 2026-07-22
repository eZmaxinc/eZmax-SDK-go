# InscriptionnotauthenticatedListElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiInscriptionID** | **int32** | The unique ID of the Inscription. | 
**PkiInscriptionnotauthenticatedID** | Pointer to **int32** | The unique ID of the Inscriptionnotauthenticated. | [optional] 
**FkiInscriptiontypeID** | **int32** | The unique ID of the Inscriptiontype | 
**SInscriptiontypeNameX** | **string** | The name of the Inscriptiontype in the language of the requester | 
**FkiInscriptionbuildingtypeID** | **int32** | The unique ID of the Inscriptionbuildingtype | 
**SInscriptionbuildingtypeNameX** | **string** | The name of the Inscriptionbuildingtype in the language of the requester | 
**FkiInscriptioncategoryID** | **int32** | The unique ID of the Inscriptioncategory | 
**SInscriptioncategoryNameX** | **string** | The name of the Inscriptioncategory in the language of the requester | 
**FkiBuyercontractID** | Pointer to **int32** | The unique ID of the Buyercontract | [optional] 
**SBuyercontractContract** | Pointer to **string** | The number of the Buyercontract | [optional] 
**EInscriptionStep** | [**FieldEInscriptionStep**](FieldEInscriptionStep.md) |  | 
**EInscriptionType** | [**FieldEInscriptionType**](FieldEInscriptionType.md) |  | 
**SInscriptionCivicend** | **string** | The address civic end of the Inscription | 
**SInscriptionMLS** | Pointer to **string** | The mls of the Inscription | [optional] 
**SInscriptionContract** | Pointer to **string** | The sale contract number | [optional] 
**DInscriptionSaleprice** | **string** | The saleprice of the Inscription | 
**DInscriptionRentprice** | **string** | The rent price of the Inscription | 
**DtInscriptionDate** | Pointer to **string** | The date of the Inscription | [optional] 
**DtInscriptionExpirationdate** | Pointer to **string** | The expiration date of the Inscription | [optional] 
**DtInscriptionNotarydate** | Pointer to **string** | The notary date of the Inscription | [optional] 
**BInscriptionInspection** | Pointer to **bool** | Whether the inscription can be acces by an inspector | [optional] 
**BInscriptionIsactive** | **bool** | Whether the inscription is active or not | 
**BInscriptionArchived** | **bool** | Whether the inscription is archived or not | 
**DtInscriptionnotauthenticatedNotaryscheduledate** | Pointer to **string** | The notary schedule date of the Inscriptionnotauthenticated | [optional] 
**DtInscriptionnotauthenticatedTransactiondate** | Pointer to **string** | The transaction date of the Inscriptionnotauthenticated | [optional] 
**DtInscriptionnotauthenticatedTransactiondateReal** | Pointer to **string** | The real transactiondate of the Inscriptionnotauthenticated | [optional] 
**BInscriptionnotauthenticatedConditional** | Pointer to **bool** | Whether the inscriptionnotauthenticated is conditional | [optional] 
**BInscriptionnotauthenticatedIsactive** | Pointer to **bool** | Whether the inscriptionnotauthenticated is active or not | [optional] 
**BInscriptionnotauthenticatedDraft** | Pointer to **bool** | Whether the Inscriptionnotauthenticated is a draft or not | [optional] 
**SAddressCivic** | Pointer to **string** | The Civic number. | [optional] 
**SAddressStreet** | Pointer to **string** | The Street Name | [optional] 
**SAddressSuite** | Pointer to **string** | The Suite or appartment number | [optional] 
**SAddressCity** | Pointer to **string** | The City name | [optional] 
**SAddressZip** | Pointer to **string** | The Postal/Zip Code  The value must be entered without spaces | [optional] 
**FkiProvinceID** | Pointer to **int32** | The unique ID of the Province.  Here are some common values (Complete list must be retrieved from API):  |Value|Description| |-|-| |1|(Canada) Alberta |2|(Canada) British Columbia| |3|(Canada) Manitoba| |3|(Canada) Manitoba| |4|(Canada) New Brunswick| |5|(Canada) Newfoundland| |6|(Canada) Northwest Territories| |7|(Canada) Nova Scotia| |8|(Canada) Nunavut| |9|(Canada) Ontario| |10|(Canada) Prince Edward Island| |11|(Canada) Quebec| |12|(Canada) Saskatchewan| |13|(Canada) Yukon| |14|(United-States) Alabama| |15|(United-States) Alaska| |16|(United-States) Arizona| |17|(United-States) Arkansas| |18|(United-States) California| |19|(United-States) Colorado| |20|(United-States) Connecticut| |21|(United-States) Delaware| |22|(United-States) District of Columbia| |23|(United-States) Florida| |24|(United-States) Georgia| |25|(United-States) Hawaii| |26|(United-States) Idaho| |27|(United-States) Illinois| |28|(United-States) Indiana| |29|(United-States) Iowa| |30|(United-States) Kansas| |31|(United-States) Kentucky| |32|(United-States) Louisiane| |33|(United-States) Maine| |34|(United-States) Maryland| |35|(United-States) Massachusetts| |36|(United-States) Michigan| |37|(United-States) Minnesota| |38|(United-States) Mississippi| |39|(United-States) Missouri| |40|(United-States) Montana| |41|(United-States) Nebraska| |42|(United-States) Nevada| |43|(United-States) New Hampshire| |44|(United-States) New Jersey| |45|(United-States) New Mexico| |46|(United-States) New York| |47|(United-States) North Carolina| |48|(United-States) North Dakota| |49|(United-States) Ohio| |50|(United-States) Oklahoma| |51|(United-States) Oregon| |52|(United-States) Pennsylvania| |53|(United-States) Rhode Island| |54|(United-States) South Carolina| |55|(United-States) South Dakota| |56|(United-States) Tennessee| |57|(United-States) Texas| |58|(United-States) Utah| |60|(United-States) Vermont| |59|(United-States) Virginia| |61|(United-States) Washington| |62|(United-States) West Virginia| |63|(United-States) Wisconsin| |64|(United-States) Wyoming| | [optional] 
**SProvinceNameX** | Pointer to **string** | The name of the Province in the language of the requester | [optional] 
**FkiCountryID** | Pointer to **int32** | The unique ID of the Country.  Here are some common values (Complete list must be retrieved from API):  |Value|Description| |-|-| |1|Canada| |2|United-States| | [optional] 
**SCountryNameX** | Pointer to **string** | The name of the Country in the language of the requester | [optional] 
**SInscriptionnotauthenticatedOffertopurchasenumber** | **string** | The offer to purchase number of the Inscriptionnotauthenticated | 
**IInscriptionUnit** | **int32** | The number of unit for the Inscription | 

## Methods

### NewInscriptionnotauthenticatedListElement

`func NewInscriptionnotauthenticatedListElement(pkiInscriptionID int32, fkiInscriptiontypeID int32, sInscriptiontypeNameX string, fkiInscriptionbuildingtypeID int32, sInscriptionbuildingtypeNameX string, fkiInscriptioncategoryID int32, sInscriptioncategoryNameX string, eInscriptionStep FieldEInscriptionStep, eInscriptionType FieldEInscriptionType, sInscriptionCivicend string, dInscriptionSaleprice string, dInscriptionRentprice string, bInscriptionIsactive bool, bInscriptionArchived bool, sInscriptionnotauthenticatedOffertopurchasenumber string, iInscriptionUnit int32, ) *InscriptionnotauthenticatedListElement`

NewInscriptionnotauthenticatedListElement instantiates a new InscriptionnotauthenticatedListElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInscriptionnotauthenticatedListElementWithDefaults

`func NewInscriptionnotauthenticatedListElementWithDefaults() *InscriptionnotauthenticatedListElement`

NewInscriptionnotauthenticatedListElementWithDefaults instantiates a new InscriptionnotauthenticatedListElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiInscriptionID

`func (o *InscriptionnotauthenticatedListElement) GetPkiInscriptionID() int32`

GetPkiInscriptionID returns the PkiInscriptionID field if non-nil, zero value otherwise.

### GetPkiInscriptionIDOk

`func (o *InscriptionnotauthenticatedListElement) GetPkiInscriptionIDOk() (*int32, bool)`

GetPkiInscriptionIDOk returns a tuple with the PkiInscriptionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiInscriptionID

`func (o *InscriptionnotauthenticatedListElement) SetPkiInscriptionID(v int32)`

SetPkiInscriptionID sets PkiInscriptionID field to given value.


### GetPkiInscriptionnotauthenticatedID

`func (o *InscriptionnotauthenticatedListElement) GetPkiInscriptionnotauthenticatedID() int32`

GetPkiInscriptionnotauthenticatedID returns the PkiInscriptionnotauthenticatedID field if non-nil, zero value otherwise.

### GetPkiInscriptionnotauthenticatedIDOk

`func (o *InscriptionnotauthenticatedListElement) GetPkiInscriptionnotauthenticatedIDOk() (*int32, bool)`

GetPkiInscriptionnotauthenticatedIDOk returns a tuple with the PkiInscriptionnotauthenticatedID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiInscriptionnotauthenticatedID

`func (o *InscriptionnotauthenticatedListElement) SetPkiInscriptionnotauthenticatedID(v int32)`

SetPkiInscriptionnotauthenticatedID sets PkiInscriptionnotauthenticatedID field to given value.

### HasPkiInscriptionnotauthenticatedID

`func (o *InscriptionnotauthenticatedListElement) HasPkiInscriptionnotauthenticatedID() bool`

HasPkiInscriptionnotauthenticatedID returns a boolean if a field has been set.

### GetFkiInscriptiontypeID

`func (o *InscriptionnotauthenticatedListElement) GetFkiInscriptiontypeID() int32`

GetFkiInscriptiontypeID returns the FkiInscriptiontypeID field if non-nil, zero value otherwise.

### GetFkiInscriptiontypeIDOk

`func (o *InscriptionnotauthenticatedListElement) GetFkiInscriptiontypeIDOk() (*int32, bool)`

GetFkiInscriptiontypeIDOk returns a tuple with the FkiInscriptiontypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiInscriptiontypeID

`func (o *InscriptionnotauthenticatedListElement) SetFkiInscriptiontypeID(v int32)`

SetFkiInscriptiontypeID sets FkiInscriptiontypeID field to given value.


### GetSInscriptiontypeNameX

`func (o *InscriptionnotauthenticatedListElement) GetSInscriptiontypeNameX() string`

GetSInscriptiontypeNameX returns the SInscriptiontypeNameX field if non-nil, zero value otherwise.

### GetSInscriptiontypeNameXOk

`func (o *InscriptionnotauthenticatedListElement) GetSInscriptiontypeNameXOk() (*string, bool)`

GetSInscriptiontypeNameXOk returns a tuple with the SInscriptiontypeNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSInscriptiontypeNameX

`func (o *InscriptionnotauthenticatedListElement) SetSInscriptiontypeNameX(v string)`

SetSInscriptiontypeNameX sets SInscriptiontypeNameX field to given value.


### GetFkiInscriptionbuildingtypeID

`func (o *InscriptionnotauthenticatedListElement) GetFkiInscriptionbuildingtypeID() int32`

GetFkiInscriptionbuildingtypeID returns the FkiInscriptionbuildingtypeID field if non-nil, zero value otherwise.

### GetFkiInscriptionbuildingtypeIDOk

`func (o *InscriptionnotauthenticatedListElement) GetFkiInscriptionbuildingtypeIDOk() (*int32, bool)`

GetFkiInscriptionbuildingtypeIDOk returns a tuple with the FkiInscriptionbuildingtypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiInscriptionbuildingtypeID

`func (o *InscriptionnotauthenticatedListElement) SetFkiInscriptionbuildingtypeID(v int32)`

SetFkiInscriptionbuildingtypeID sets FkiInscriptionbuildingtypeID field to given value.


### GetSInscriptionbuildingtypeNameX

`func (o *InscriptionnotauthenticatedListElement) GetSInscriptionbuildingtypeNameX() string`

GetSInscriptionbuildingtypeNameX returns the SInscriptionbuildingtypeNameX field if non-nil, zero value otherwise.

### GetSInscriptionbuildingtypeNameXOk

`func (o *InscriptionnotauthenticatedListElement) GetSInscriptionbuildingtypeNameXOk() (*string, bool)`

GetSInscriptionbuildingtypeNameXOk returns a tuple with the SInscriptionbuildingtypeNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSInscriptionbuildingtypeNameX

`func (o *InscriptionnotauthenticatedListElement) SetSInscriptionbuildingtypeNameX(v string)`

SetSInscriptionbuildingtypeNameX sets SInscriptionbuildingtypeNameX field to given value.


### GetFkiInscriptioncategoryID

`func (o *InscriptionnotauthenticatedListElement) GetFkiInscriptioncategoryID() int32`

GetFkiInscriptioncategoryID returns the FkiInscriptioncategoryID field if non-nil, zero value otherwise.

### GetFkiInscriptioncategoryIDOk

`func (o *InscriptionnotauthenticatedListElement) GetFkiInscriptioncategoryIDOk() (*int32, bool)`

GetFkiInscriptioncategoryIDOk returns a tuple with the FkiInscriptioncategoryID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiInscriptioncategoryID

`func (o *InscriptionnotauthenticatedListElement) SetFkiInscriptioncategoryID(v int32)`

SetFkiInscriptioncategoryID sets FkiInscriptioncategoryID field to given value.


### GetSInscriptioncategoryNameX

`func (o *InscriptionnotauthenticatedListElement) GetSInscriptioncategoryNameX() string`

GetSInscriptioncategoryNameX returns the SInscriptioncategoryNameX field if non-nil, zero value otherwise.

### GetSInscriptioncategoryNameXOk

`func (o *InscriptionnotauthenticatedListElement) GetSInscriptioncategoryNameXOk() (*string, bool)`

GetSInscriptioncategoryNameXOk returns a tuple with the SInscriptioncategoryNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSInscriptioncategoryNameX

`func (o *InscriptionnotauthenticatedListElement) SetSInscriptioncategoryNameX(v string)`

SetSInscriptioncategoryNameX sets SInscriptioncategoryNameX field to given value.


### GetFkiBuyercontractID

`func (o *InscriptionnotauthenticatedListElement) GetFkiBuyercontractID() int32`

GetFkiBuyercontractID returns the FkiBuyercontractID field if non-nil, zero value otherwise.

### GetFkiBuyercontractIDOk

`func (o *InscriptionnotauthenticatedListElement) GetFkiBuyercontractIDOk() (*int32, bool)`

GetFkiBuyercontractIDOk returns a tuple with the FkiBuyercontractID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiBuyercontractID

`func (o *InscriptionnotauthenticatedListElement) SetFkiBuyercontractID(v int32)`

SetFkiBuyercontractID sets FkiBuyercontractID field to given value.

### HasFkiBuyercontractID

`func (o *InscriptionnotauthenticatedListElement) HasFkiBuyercontractID() bool`

HasFkiBuyercontractID returns a boolean if a field has been set.

### GetSBuyercontractContract

`func (o *InscriptionnotauthenticatedListElement) GetSBuyercontractContract() string`

GetSBuyercontractContract returns the SBuyercontractContract field if non-nil, zero value otherwise.

### GetSBuyercontractContractOk

`func (o *InscriptionnotauthenticatedListElement) GetSBuyercontractContractOk() (*string, bool)`

GetSBuyercontractContractOk returns a tuple with the SBuyercontractContract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSBuyercontractContract

`func (o *InscriptionnotauthenticatedListElement) SetSBuyercontractContract(v string)`

SetSBuyercontractContract sets SBuyercontractContract field to given value.

### HasSBuyercontractContract

`func (o *InscriptionnotauthenticatedListElement) HasSBuyercontractContract() bool`

HasSBuyercontractContract returns a boolean if a field has been set.

### GetEInscriptionStep

`func (o *InscriptionnotauthenticatedListElement) GetEInscriptionStep() FieldEInscriptionStep`

GetEInscriptionStep returns the EInscriptionStep field if non-nil, zero value otherwise.

### GetEInscriptionStepOk

`func (o *InscriptionnotauthenticatedListElement) GetEInscriptionStepOk() (*FieldEInscriptionStep, bool)`

GetEInscriptionStepOk returns a tuple with the EInscriptionStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEInscriptionStep

`func (o *InscriptionnotauthenticatedListElement) SetEInscriptionStep(v FieldEInscriptionStep)`

SetEInscriptionStep sets EInscriptionStep field to given value.


### GetEInscriptionType

`func (o *InscriptionnotauthenticatedListElement) GetEInscriptionType() FieldEInscriptionType`

GetEInscriptionType returns the EInscriptionType field if non-nil, zero value otherwise.

### GetEInscriptionTypeOk

`func (o *InscriptionnotauthenticatedListElement) GetEInscriptionTypeOk() (*FieldEInscriptionType, bool)`

GetEInscriptionTypeOk returns a tuple with the EInscriptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEInscriptionType

`func (o *InscriptionnotauthenticatedListElement) SetEInscriptionType(v FieldEInscriptionType)`

SetEInscriptionType sets EInscriptionType field to given value.


### GetSInscriptionCivicend

`func (o *InscriptionnotauthenticatedListElement) GetSInscriptionCivicend() string`

GetSInscriptionCivicend returns the SInscriptionCivicend field if non-nil, zero value otherwise.

### GetSInscriptionCivicendOk

`func (o *InscriptionnotauthenticatedListElement) GetSInscriptionCivicendOk() (*string, bool)`

GetSInscriptionCivicendOk returns a tuple with the SInscriptionCivicend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSInscriptionCivicend

`func (o *InscriptionnotauthenticatedListElement) SetSInscriptionCivicend(v string)`

SetSInscriptionCivicend sets SInscriptionCivicend field to given value.


### GetSInscriptionMLS

`func (o *InscriptionnotauthenticatedListElement) GetSInscriptionMLS() string`

GetSInscriptionMLS returns the SInscriptionMLS field if non-nil, zero value otherwise.

### GetSInscriptionMLSOk

`func (o *InscriptionnotauthenticatedListElement) GetSInscriptionMLSOk() (*string, bool)`

GetSInscriptionMLSOk returns a tuple with the SInscriptionMLS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSInscriptionMLS

`func (o *InscriptionnotauthenticatedListElement) SetSInscriptionMLS(v string)`

SetSInscriptionMLS sets SInscriptionMLS field to given value.

### HasSInscriptionMLS

`func (o *InscriptionnotauthenticatedListElement) HasSInscriptionMLS() bool`

HasSInscriptionMLS returns a boolean if a field has been set.

### GetSInscriptionContract

`func (o *InscriptionnotauthenticatedListElement) GetSInscriptionContract() string`

GetSInscriptionContract returns the SInscriptionContract field if non-nil, zero value otherwise.

### GetSInscriptionContractOk

`func (o *InscriptionnotauthenticatedListElement) GetSInscriptionContractOk() (*string, bool)`

GetSInscriptionContractOk returns a tuple with the SInscriptionContract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSInscriptionContract

`func (o *InscriptionnotauthenticatedListElement) SetSInscriptionContract(v string)`

SetSInscriptionContract sets SInscriptionContract field to given value.

### HasSInscriptionContract

`func (o *InscriptionnotauthenticatedListElement) HasSInscriptionContract() bool`

HasSInscriptionContract returns a boolean if a field has been set.

### GetDInscriptionSaleprice

`func (o *InscriptionnotauthenticatedListElement) GetDInscriptionSaleprice() string`

GetDInscriptionSaleprice returns the DInscriptionSaleprice field if non-nil, zero value otherwise.

### GetDInscriptionSalepriceOk

`func (o *InscriptionnotauthenticatedListElement) GetDInscriptionSalepriceOk() (*string, bool)`

GetDInscriptionSalepriceOk returns a tuple with the DInscriptionSaleprice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDInscriptionSaleprice

`func (o *InscriptionnotauthenticatedListElement) SetDInscriptionSaleprice(v string)`

SetDInscriptionSaleprice sets DInscriptionSaleprice field to given value.


### GetDInscriptionRentprice

`func (o *InscriptionnotauthenticatedListElement) GetDInscriptionRentprice() string`

GetDInscriptionRentprice returns the DInscriptionRentprice field if non-nil, zero value otherwise.

### GetDInscriptionRentpriceOk

`func (o *InscriptionnotauthenticatedListElement) GetDInscriptionRentpriceOk() (*string, bool)`

GetDInscriptionRentpriceOk returns a tuple with the DInscriptionRentprice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDInscriptionRentprice

`func (o *InscriptionnotauthenticatedListElement) SetDInscriptionRentprice(v string)`

SetDInscriptionRentprice sets DInscriptionRentprice field to given value.


### GetDtInscriptionDate

`func (o *InscriptionnotauthenticatedListElement) GetDtInscriptionDate() string`

GetDtInscriptionDate returns the DtInscriptionDate field if non-nil, zero value otherwise.

### GetDtInscriptionDateOk

`func (o *InscriptionnotauthenticatedListElement) GetDtInscriptionDateOk() (*string, bool)`

GetDtInscriptionDateOk returns a tuple with the DtInscriptionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtInscriptionDate

`func (o *InscriptionnotauthenticatedListElement) SetDtInscriptionDate(v string)`

SetDtInscriptionDate sets DtInscriptionDate field to given value.

### HasDtInscriptionDate

`func (o *InscriptionnotauthenticatedListElement) HasDtInscriptionDate() bool`

HasDtInscriptionDate returns a boolean if a field has been set.

### GetDtInscriptionExpirationdate

`func (o *InscriptionnotauthenticatedListElement) GetDtInscriptionExpirationdate() string`

GetDtInscriptionExpirationdate returns the DtInscriptionExpirationdate field if non-nil, zero value otherwise.

### GetDtInscriptionExpirationdateOk

`func (o *InscriptionnotauthenticatedListElement) GetDtInscriptionExpirationdateOk() (*string, bool)`

GetDtInscriptionExpirationdateOk returns a tuple with the DtInscriptionExpirationdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtInscriptionExpirationdate

`func (o *InscriptionnotauthenticatedListElement) SetDtInscriptionExpirationdate(v string)`

SetDtInscriptionExpirationdate sets DtInscriptionExpirationdate field to given value.

### HasDtInscriptionExpirationdate

`func (o *InscriptionnotauthenticatedListElement) HasDtInscriptionExpirationdate() bool`

HasDtInscriptionExpirationdate returns a boolean if a field has been set.

### GetDtInscriptionNotarydate

`func (o *InscriptionnotauthenticatedListElement) GetDtInscriptionNotarydate() string`

GetDtInscriptionNotarydate returns the DtInscriptionNotarydate field if non-nil, zero value otherwise.

### GetDtInscriptionNotarydateOk

`func (o *InscriptionnotauthenticatedListElement) GetDtInscriptionNotarydateOk() (*string, bool)`

GetDtInscriptionNotarydateOk returns a tuple with the DtInscriptionNotarydate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtInscriptionNotarydate

`func (o *InscriptionnotauthenticatedListElement) SetDtInscriptionNotarydate(v string)`

SetDtInscriptionNotarydate sets DtInscriptionNotarydate field to given value.

### HasDtInscriptionNotarydate

`func (o *InscriptionnotauthenticatedListElement) HasDtInscriptionNotarydate() bool`

HasDtInscriptionNotarydate returns a boolean if a field has been set.

### GetBInscriptionInspection

`func (o *InscriptionnotauthenticatedListElement) GetBInscriptionInspection() bool`

GetBInscriptionInspection returns the BInscriptionInspection field if non-nil, zero value otherwise.

### GetBInscriptionInspectionOk

`func (o *InscriptionnotauthenticatedListElement) GetBInscriptionInspectionOk() (*bool, bool)`

GetBInscriptionInspectionOk returns a tuple with the BInscriptionInspection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBInscriptionInspection

`func (o *InscriptionnotauthenticatedListElement) SetBInscriptionInspection(v bool)`

SetBInscriptionInspection sets BInscriptionInspection field to given value.

### HasBInscriptionInspection

`func (o *InscriptionnotauthenticatedListElement) HasBInscriptionInspection() bool`

HasBInscriptionInspection returns a boolean if a field has been set.

### GetBInscriptionIsactive

`func (o *InscriptionnotauthenticatedListElement) GetBInscriptionIsactive() bool`

GetBInscriptionIsactive returns the BInscriptionIsactive field if non-nil, zero value otherwise.

### GetBInscriptionIsactiveOk

`func (o *InscriptionnotauthenticatedListElement) GetBInscriptionIsactiveOk() (*bool, bool)`

GetBInscriptionIsactiveOk returns a tuple with the BInscriptionIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBInscriptionIsactive

`func (o *InscriptionnotauthenticatedListElement) SetBInscriptionIsactive(v bool)`

SetBInscriptionIsactive sets BInscriptionIsactive field to given value.


### GetBInscriptionArchived

`func (o *InscriptionnotauthenticatedListElement) GetBInscriptionArchived() bool`

GetBInscriptionArchived returns the BInscriptionArchived field if non-nil, zero value otherwise.

### GetBInscriptionArchivedOk

`func (o *InscriptionnotauthenticatedListElement) GetBInscriptionArchivedOk() (*bool, bool)`

GetBInscriptionArchivedOk returns a tuple with the BInscriptionArchived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBInscriptionArchived

`func (o *InscriptionnotauthenticatedListElement) SetBInscriptionArchived(v bool)`

SetBInscriptionArchived sets BInscriptionArchived field to given value.


### GetDtInscriptionnotauthenticatedNotaryscheduledate

`func (o *InscriptionnotauthenticatedListElement) GetDtInscriptionnotauthenticatedNotaryscheduledate() string`

GetDtInscriptionnotauthenticatedNotaryscheduledate returns the DtInscriptionnotauthenticatedNotaryscheduledate field if non-nil, zero value otherwise.

### GetDtInscriptionnotauthenticatedNotaryscheduledateOk

`func (o *InscriptionnotauthenticatedListElement) GetDtInscriptionnotauthenticatedNotaryscheduledateOk() (*string, bool)`

GetDtInscriptionnotauthenticatedNotaryscheduledateOk returns a tuple with the DtInscriptionnotauthenticatedNotaryscheduledate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtInscriptionnotauthenticatedNotaryscheduledate

`func (o *InscriptionnotauthenticatedListElement) SetDtInscriptionnotauthenticatedNotaryscheduledate(v string)`

SetDtInscriptionnotauthenticatedNotaryscheduledate sets DtInscriptionnotauthenticatedNotaryscheduledate field to given value.

### HasDtInscriptionnotauthenticatedNotaryscheduledate

`func (o *InscriptionnotauthenticatedListElement) HasDtInscriptionnotauthenticatedNotaryscheduledate() bool`

HasDtInscriptionnotauthenticatedNotaryscheduledate returns a boolean if a field has been set.

### GetDtInscriptionnotauthenticatedTransactiondate

`func (o *InscriptionnotauthenticatedListElement) GetDtInscriptionnotauthenticatedTransactiondate() string`

GetDtInscriptionnotauthenticatedTransactiondate returns the DtInscriptionnotauthenticatedTransactiondate field if non-nil, zero value otherwise.

### GetDtInscriptionnotauthenticatedTransactiondateOk

`func (o *InscriptionnotauthenticatedListElement) GetDtInscriptionnotauthenticatedTransactiondateOk() (*string, bool)`

GetDtInscriptionnotauthenticatedTransactiondateOk returns a tuple with the DtInscriptionnotauthenticatedTransactiondate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtInscriptionnotauthenticatedTransactiondate

`func (o *InscriptionnotauthenticatedListElement) SetDtInscriptionnotauthenticatedTransactiondate(v string)`

SetDtInscriptionnotauthenticatedTransactiondate sets DtInscriptionnotauthenticatedTransactiondate field to given value.

### HasDtInscriptionnotauthenticatedTransactiondate

`func (o *InscriptionnotauthenticatedListElement) HasDtInscriptionnotauthenticatedTransactiondate() bool`

HasDtInscriptionnotauthenticatedTransactiondate returns a boolean if a field has been set.

### GetDtInscriptionnotauthenticatedTransactiondateReal

`func (o *InscriptionnotauthenticatedListElement) GetDtInscriptionnotauthenticatedTransactiondateReal() string`

GetDtInscriptionnotauthenticatedTransactiondateReal returns the DtInscriptionnotauthenticatedTransactiondateReal field if non-nil, zero value otherwise.

### GetDtInscriptionnotauthenticatedTransactiondateRealOk

`func (o *InscriptionnotauthenticatedListElement) GetDtInscriptionnotauthenticatedTransactiondateRealOk() (*string, bool)`

GetDtInscriptionnotauthenticatedTransactiondateRealOk returns a tuple with the DtInscriptionnotauthenticatedTransactiondateReal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtInscriptionnotauthenticatedTransactiondateReal

`func (o *InscriptionnotauthenticatedListElement) SetDtInscriptionnotauthenticatedTransactiondateReal(v string)`

SetDtInscriptionnotauthenticatedTransactiondateReal sets DtInscriptionnotauthenticatedTransactiondateReal field to given value.

### HasDtInscriptionnotauthenticatedTransactiondateReal

`func (o *InscriptionnotauthenticatedListElement) HasDtInscriptionnotauthenticatedTransactiondateReal() bool`

HasDtInscriptionnotauthenticatedTransactiondateReal returns a boolean if a field has been set.

### GetBInscriptionnotauthenticatedConditional

`func (o *InscriptionnotauthenticatedListElement) GetBInscriptionnotauthenticatedConditional() bool`

GetBInscriptionnotauthenticatedConditional returns the BInscriptionnotauthenticatedConditional field if non-nil, zero value otherwise.

### GetBInscriptionnotauthenticatedConditionalOk

`func (o *InscriptionnotauthenticatedListElement) GetBInscriptionnotauthenticatedConditionalOk() (*bool, bool)`

GetBInscriptionnotauthenticatedConditionalOk returns a tuple with the BInscriptionnotauthenticatedConditional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBInscriptionnotauthenticatedConditional

`func (o *InscriptionnotauthenticatedListElement) SetBInscriptionnotauthenticatedConditional(v bool)`

SetBInscriptionnotauthenticatedConditional sets BInscriptionnotauthenticatedConditional field to given value.

### HasBInscriptionnotauthenticatedConditional

`func (o *InscriptionnotauthenticatedListElement) HasBInscriptionnotauthenticatedConditional() bool`

HasBInscriptionnotauthenticatedConditional returns a boolean if a field has been set.

### GetBInscriptionnotauthenticatedIsactive

`func (o *InscriptionnotauthenticatedListElement) GetBInscriptionnotauthenticatedIsactive() bool`

GetBInscriptionnotauthenticatedIsactive returns the BInscriptionnotauthenticatedIsactive field if non-nil, zero value otherwise.

### GetBInscriptionnotauthenticatedIsactiveOk

`func (o *InscriptionnotauthenticatedListElement) GetBInscriptionnotauthenticatedIsactiveOk() (*bool, bool)`

GetBInscriptionnotauthenticatedIsactiveOk returns a tuple with the BInscriptionnotauthenticatedIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBInscriptionnotauthenticatedIsactive

`func (o *InscriptionnotauthenticatedListElement) SetBInscriptionnotauthenticatedIsactive(v bool)`

SetBInscriptionnotauthenticatedIsactive sets BInscriptionnotauthenticatedIsactive field to given value.

### HasBInscriptionnotauthenticatedIsactive

`func (o *InscriptionnotauthenticatedListElement) HasBInscriptionnotauthenticatedIsactive() bool`

HasBInscriptionnotauthenticatedIsactive returns a boolean if a field has been set.

### GetBInscriptionnotauthenticatedDraft

`func (o *InscriptionnotauthenticatedListElement) GetBInscriptionnotauthenticatedDraft() bool`

GetBInscriptionnotauthenticatedDraft returns the BInscriptionnotauthenticatedDraft field if non-nil, zero value otherwise.

### GetBInscriptionnotauthenticatedDraftOk

`func (o *InscriptionnotauthenticatedListElement) GetBInscriptionnotauthenticatedDraftOk() (*bool, bool)`

GetBInscriptionnotauthenticatedDraftOk returns a tuple with the BInscriptionnotauthenticatedDraft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBInscriptionnotauthenticatedDraft

`func (o *InscriptionnotauthenticatedListElement) SetBInscriptionnotauthenticatedDraft(v bool)`

SetBInscriptionnotauthenticatedDraft sets BInscriptionnotauthenticatedDraft field to given value.

### HasBInscriptionnotauthenticatedDraft

`func (o *InscriptionnotauthenticatedListElement) HasBInscriptionnotauthenticatedDraft() bool`

HasBInscriptionnotauthenticatedDraft returns a boolean if a field has been set.

### GetSAddressCivic

`func (o *InscriptionnotauthenticatedListElement) GetSAddressCivic() string`

GetSAddressCivic returns the SAddressCivic field if non-nil, zero value otherwise.

### GetSAddressCivicOk

`func (o *InscriptionnotauthenticatedListElement) GetSAddressCivicOk() (*string, bool)`

GetSAddressCivicOk returns a tuple with the SAddressCivic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAddressCivic

`func (o *InscriptionnotauthenticatedListElement) SetSAddressCivic(v string)`

SetSAddressCivic sets SAddressCivic field to given value.

### HasSAddressCivic

`func (o *InscriptionnotauthenticatedListElement) HasSAddressCivic() bool`

HasSAddressCivic returns a boolean if a field has been set.

### GetSAddressStreet

`func (o *InscriptionnotauthenticatedListElement) GetSAddressStreet() string`

GetSAddressStreet returns the SAddressStreet field if non-nil, zero value otherwise.

### GetSAddressStreetOk

`func (o *InscriptionnotauthenticatedListElement) GetSAddressStreetOk() (*string, bool)`

GetSAddressStreetOk returns a tuple with the SAddressStreet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAddressStreet

`func (o *InscriptionnotauthenticatedListElement) SetSAddressStreet(v string)`

SetSAddressStreet sets SAddressStreet field to given value.

### HasSAddressStreet

`func (o *InscriptionnotauthenticatedListElement) HasSAddressStreet() bool`

HasSAddressStreet returns a boolean if a field has been set.

### GetSAddressSuite

`func (o *InscriptionnotauthenticatedListElement) GetSAddressSuite() string`

GetSAddressSuite returns the SAddressSuite field if non-nil, zero value otherwise.

### GetSAddressSuiteOk

`func (o *InscriptionnotauthenticatedListElement) GetSAddressSuiteOk() (*string, bool)`

GetSAddressSuiteOk returns a tuple with the SAddressSuite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAddressSuite

`func (o *InscriptionnotauthenticatedListElement) SetSAddressSuite(v string)`

SetSAddressSuite sets SAddressSuite field to given value.

### HasSAddressSuite

`func (o *InscriptionnotauthenticatedListElement) HasSAddressSuite() bool`

HasSAddressSuite returns a boolean if a field has been set.

### GetSAddressCity

`func (o *InscriptionnotauthenticatedListElement) GetSAddressCity() string`

GetSAddressCity returns the SAddressCity field if non-nil, zero value otherwise.

### GetSAddressCityOk

`func (o *InscriptionnotauthenticatedListElement) GetSAddressCityOk() (*string, bool)`

GetSAddressCityOk returns a tuple with the SAddressCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAddressCity

`func (o *InscriptionnotauthenticatedListElement) SetSAddressCity(v string)`

SetSAddressCity sets SAddressCity field to given value.

### HasSAddressCity

`func (o *InscriptionnotauthenticatedListElement) HasSAddressCity() bool`

HasSAddressCity returns a boolean if a field has been set.

### GetSAddressZip

`func (o *InscriptionnotauthenticatedListElement) GetSAddressZip() string`

GetSAddressZip returns the SAddressZip field if non-nil, zero value otherwise.

### GetSAddressZipOk

`func (o *InscriptionnotauthenticatedListElement) GetSAddressZipOk() (*string, bool)`

GetSAddressZipOk returns a tuple with the SAddressZip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAddressZip

`func (o *InscriptionnotauthenticatedListElement) SetSAddressZip(v string)`

SetSAddressZip sets SAddressZip field to given value.

### HasSAddressZip

`func (o *InscriptionnotauthenticatedListElement) HasSAddressZip() bool`

HasSAddressZip returns a boolean if a field has been set.

### GetFkiProvinceID

`func (o *InscriptionnotauthenticatedListElement) GetFkiProvinceID() int32`

GetFkiProvinceID returns the FkiProvinceID field if non-nil, zero value otherwise.

### GetFkiProvinceIDOk

`func (o *InscriptionnotauthenticatedListElement) GetFkiProvinceIDOk() (*int32, bool)`

GetFkiProvinceIDOk returns a tuple with the FkiProvinceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiProvinceID

`func (o *InscriptionnotauthenticatedListElement) SetFkiProvinceID(v int32)`

SetFkiProvinceID sets FkiProvinceID field to given value.

### HasFkiProvinceID

`func (o *InscriptionnotauthenticatedListElement) HasFkiProvinceID() bool`

HasFkiProvinceID returns a boolean if a field has been set.

### GetSProvinceNameX

`func (o *InscriptionnotauthenticatedListElement) GetSProvinceNameX() string`

GetSProvinceNameX returns the SProvinceNameX field if non-nil, zero value otherwise.

### GetSProvinceNameXOk

`func (o *InscriptionnotauthenticatedListElement) GetSProvinceNameXOk() (*string, bool)`

GetSProvinceNameXOk returns a tuple with the SProvinceNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSProvinceNameX

`func (o *InscriptionnotauthenticatedListElement) SetSProvinceNameX(v string)`

SetSProvinceNameX sets SProvinceNameX field to given value.

### HasSProvinceNameX

`func (o *InscriptionnotauthenticatedListElement) HasSProvinceNameX() bool`

HasSProvinceNameX returns a boolean if a field has been set.

### GetFkiCountryID

`func (o *InscriptionnotauthenticatedListElement) GetFkiCountryID() int32`

GetFkiCountryID returns the FkiCountryID field if non-nil, zero value otherwise.

### GetFkiCountryIDOk

`func (o *InscriptionnotauthenticatedListElement) GetFkiCountryIDOk() (*int32, bool)`

GetFkiCountryIDOk returns a tuple with the FkiCountryID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiCountryID

`func (o *InscriptionnotauthenticatedListElement) SetFkiCountryID(v int32)`

SetFkiCountryID sets FkiCountryID field to given value.

### HasFkiCountryID

`func (o *InscriptionnotauthenticatedListElement) HasFkiCountryID() bool`

HasFkiCountryID returns a boolean if a field has been set.

### GetSCountryNameX

`func (o *InscriptionnotauthenticatedListElement) GetSCountryNameX() string`

GetSCountryNameX returns the SCountryNameX field if non-nil, zero value otherwise.

### GetSCountryNameXOk

`func (o *InscriptionnotauthenticatedListElement) GetSCountryNameXOk() (*string, bool)`

GetSCountryNameXOk returns a tuple with the SCountryNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSCountryNameX

`func (o *InscriptionnotauthenticatedListElement) SetSCountryNameX(v string)`

SetSCountryNameX sets SCountryNameX field to given value.

### HasSCountryNameX

`func (o *InscriptionnotauthenticatedListElement) HasSCountryNameX() bool`

HasSCountryNameX returns a boolean if a field has been set.

### GetSInscriptionnotauthenticatedOffertopurchasenumber

`func (o *InscriptionnotauthenticatedListElement) GetSInscriptionnotauthenticatedOffertopurchasenumber() string`

GetSInscriptionnotauthenticatedOffertopurchasenumber returns the SInscriptionnotauthenticatedOffertopurchasenumber field if non-nil, zero value otherwise.

### GetSInscriptionnotauthenticatedOffertopurchasenumberOk

`func (o *InscriptionnotauthenticatedListElement) GetSInscriptionnotauthenticatedOffertopurchasenumberOk() (*string, bool)`

GetSInscriptionnotauthenticatedOffertopurchasenumberOk returns a tuple with the SInscriptionnotauthenticatedOffertopurchasenumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSInscriptionnotauthenticatedOffertopurchasenumber

`func (o *InscriptionnotauthenticatedListElement) SetSInscriptionnotauthenticatedOffertopurchasenumber(v string)`

SetSInscriptionnotauthenticatedOffertopurchasenumber sets SInscriptionnotauthenticatedOffertopurchasenumber field to given value.


### GetIInscriptionUnit

`func (o *InscriptionnotauthenticatedListElement) GetIInscriptionUnit() int32`

GetIInscriptionUnit returns the IInscriptionUnit field if non-nil, zero value otherwise.

### GetIInscriptionUnitOk

`func (o *InscriptionnotauthenticatedListElement) GetIInscriptionUnitOk() (*int32, bool)`

GetIInscriptionUnitOk returns a tuple with the IInscriptionUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIInscriptionUnit

`func (o *InscriptionnotauthenticatedListElement) SetIInscriptionUnit(v int32)`

SetIInscriptionUnit sets IInscriptionUnit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


