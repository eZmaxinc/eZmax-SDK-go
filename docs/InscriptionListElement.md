# InscriptionListElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiInscriptionID** | **int32** | The unique ID of the Inscription. | 
**PkiInscriptionnotauthenticatedID** | Pointer to **int32** | The unique ID of the Inscriptionnotauthenticated. | [optional] 
**FkiInscriptiontypeID** | **int32** | The unique ID of the Inscriptiontype | 
**SInscriptiontypeNameX** | **string** | The name of the Inscriptiontype in the language of the requester | 
**EInscriptionStep** | [**FieldEInscriptionStep**](FieldEInscriptionStep.md) |  | 
**SInscriptionCivicend** | **string** | The civicend of the Inscription | 
**SInscriptionMLS** | Pointer to **string** | The mls of the Inscription | [optional] 
**DInscriptionSaleprice** | **string** | The saleprice of the Inscription | 
**DInscriptionRentprice** | **string** | The rentprice of the Inscription | 
**DtInscriptionDate** | Pointer to **string** | The date of the Inscription | [optional] 
**DtInscriptionExpirationdate** | Pointer to **string** | The expirationdate of the Inscription | [optional] 
**DtInscriptionNotarydate** | Pointer to **string** | The notarydate of the Inscription | [optional] 
**BInscriptionIsactive** | **bool** | Whether the inscription is active or not | 
**BInscriptionArchived** | **bool** | Whether the inscription is archived or not | 
**BInscriptionInspection** | Pointer to **bool** | Whether the inscription can be acces by an inspector | [optional] 
**DtInscriptionnotauthenticatedNotaryscheduledate** | Pointer to **string** | The notaryscheduledate of the Inscriptionnotauthenticated | [optional] 
**DtInscriptionnotauthenticatedTransactiondate** | Pointer to **string** | The transactiondate of the Inscriptionnotauthenticated | [optional] 
**DtInscriptionnotauthenticatedTransactiondateReal** | Pointer to **string** | The transactiondatereal of the Inscriptionnotauthenticated | [optional] 
**BInscriptionnotauthenticatedConditional** | Pointer to **bool** | Whether the inscriptionnotauthenticated is conditional | [optional] 
**BInscriptionnotauthenticatedIsactive** | Pointer to **bool** | Whether the inscriptionnotauthenticated is active or not | [optional] 
**SAddressCivic** | Pointer to **string** | The Civic number. | [optional] 
**SAddressStreet** | Pointer to **string** | The Street Name | [optional] 
**SAddressSuite** | Pointer to **string** | The Suite or appartment number | [optional] 
**SAddressCity** | Pointer to **string** | The City name | [optional] 
**SAddressZip** | Pointer to **string** | The Postal/Zip Code  The value must be entered without spaces | [optional] 
**FkiProvinceID** | Pointer to **int32** | The unique ID of the Province.  Here are some common values (Complete list must be retrieved from API):  |Value|Description| |-|-| |1|(Canada) Alberta |2|(Canada) British Columbia| |3|(Canada) Manitoba| |3|(Canada) Manitoba| |4|(Canada) New Brunswick| |5|(Canada) Newfoundland| |6|(Canada) Northwest Territories| |7|(Canada) Nova Scotia| |8|(Canada) Nunavut| |9|(Canada) Ontario| |10|(Canada) Prince Edward Island| |11|(Canada) Quebec| |12|(Canada) Saskatchewan| |13|(Canada) Yukon| |14|(United-States) Alabama| |15|(United-States) Alaska| |16|(United-States) Arizona| |17|(United-States) Arkansas| |18|(United-States) California| |19|(United-States) Colorado| |20|(United-States) Connecticut| |21|(United-States) Delaware| |22|(United-States) District of Columbia| |23|(United-States) Florida| |24|(United-States) Georgia| |25|(United-States) Hawaii| |26|(United-States) Idaho| |27|(United-States) Illinois| |28|(United-States) Indiana| |29|(United-States) Iowa| |30|(United-States) Kansas| |31|(United-States) Kentucky| |32|(United-States) Louisiane| |33|(United-States) Maine| |34|(United-States) Maryland| |35|(United-States) Massachusetts| |36|(United-States) Michigan| |37|(United-States) Minnesota| |38|(United-States) Mississippi| |39|(United-States) Missouri| |40|(United-States) Montana| |41|(United-States) Nebraska| |42|(United-States) Nevada| |43|(United-States) New Hampshire| |44|(United-States) New Jersey| |45|(United-States) New Mexico| |46|(United-States) New York| |47|(United-States) North Carolina| |48|(United-States) North Dakota| |49|(United-States) Ohio| |50|(United-States) Oklahoma| |51|(United-States) Oregon| |52|(United-States) Pennsylvania| |53|(United-States) Rhode Island| |54|(United-States) South Carolina| |55|(United-States) South Dakota| |56|(United-States) Tennessee| |57|(United-States) Texas| |58|(United-States) Utah| |60|(United-States) Vermont| |59|(United-States) Virginia| |61|(United-States) Washington| |62|(United-States) West Virginia| |63|(United-States) Wisconsin| |64|(United-States) Wyoming| | [optional] 
**SProvinceNameX** | Pointer to **string** | The name of the Province in the language of the requester | [optional] 
**FkiCountryID** | Pointer to **int32** | The unique ID of the Country.  Here are some common values (Complete list must be retrieved from API):  |Value|Description| |-|-| |1|Canada| |2|United-States| | [optional] 
**SCountryNameX** | Pointer to **string** | The name of the Country in the language of the requester | [optional] 
**IInscriptionnotauthenticatedCanceled** | **int32** | The numbre of inscriptionnotauthenticated was canceled in this Inscription | 
**BAllowedCopyintoinscriptionedm** | **bool** | Whether we are allowed to copy into the Inscription EDM | 

## Methods

### NewInscriptionListElement

`func NewInscriptionListElement(pkiInscriptionID int32, fkiInscriptiontypeID int32, sInscriptiontypeNameX string, eInscriptionStep FieldEInscriptionStep, sInscriptionCivicend string, dInscriptionSaleprice string, dInscriptionRentprice string, bInscriptionIsactive bool, bInscriptionArchived bool, iInscriptionnotauthenticatedCanceled int32, bAllowedCopyintoinscriptionedm bool, ) *InscriptionListElement`

NewInscriptionListElement instantiates a new InscriptionListElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInscriptionListElementWithDefaults

`func NewInscriptionListElementWithDefaults() *InscriptionListElement`

NewInscriptionListElementWithDefaults instantiates a new InscriptionListElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiInscriptionID

`func (o *InscriptionListElement) GetPkiInscriptionID() int32`

GetPkiInscriptionID returns the PkiInscriptionID field if non-nil, zero value otherwise.

### GetPkiInscriptionIDOk

`func (o *InscriptionListElement) GetPkiInscriptionIDOk() (*int32, bool)`

GetPkiInscriptionIDOk returns a tuple with the PkiInscriptionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiInscriptionID

`func (o *InscriptionListElement) SetPkiInscriptionID(v int32)`

SetPkiInscriptionID sets PkiInscriptionID field to given value.


### GetPkiInscriptionnotauthenticatedID

`func (o *InscriptionListElement) GetPkiInscriptionnotauthenticatedID() int32`

GetPkiInscriptionnotauthenticatedID returns the PkiInscriptionnotauthenticatedID field if non-nil, zero value otherwise.

### GetPkiInscriptionnotauthenticatedIDOk

`func (o *InscriptionListElement) GetPkiInscriptionnotauthenticatedIDOk() (*int32, bool)`

GetPkiInscriptionnotauthenticatedIDOk returns a tuple with the PkiInscriptionnotauthenticatedID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiInscriptionnotauthenticatedID

`func (o *InscriptionListElement) SetPkiInscriptionnotauthenticatedID(v int32)`

SetPkiInscriptionnotauthenticatedID sets PkiInscriptionnotauthenticatedID field to given value.

### HasPkiInscriptionnotauthenticatedID

`func (o *InscriptionListElement) HasPkiInscriptionnotauthenticatedID() bool`

HasPkiInscriptionnotauthenticatedID returns a boolean if a field has been set.

### GetFkiInscriptiontypeID

`func (o *InscriptionListElement) GetFkiInscriptiontypeID() int32`

GetFkiInscriptiontypeID returns the FkiInscriptiontypeID field if non-nil, zero value otherwise.

### GetFkiInscriptiontypeIDOk

`func (o *InscriptionListElement) GetFkiInscriptiontypeIDOk() (*int32, bool)`

GetFkiInscriptiontypeIDOk returns a tuple with the FkiInscriptiontypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiInscriptiontypeID

`func (o *InscriptionListElement) SetFkiInscriptiontypeID(v int32)`

SetFkiInscriptiontypeID sets FkiInscriptiontypeID field to given value.


### GetSInscriptiontypeNameX

`func (o *InscriptionListElement) GetSInscriptiontypeNameX() string`

GetSInscriptiontypeNameX returns the SInscriptiontypeNameX field if non-nil, zero value otherwise.

### GetSInscriptiontypeNameXOk

`func (o *InscriptionListElement) GetSInscriptiontypeNameXOk() (*string, bool)`

GetSInscriptiontypeNameXOk returns a tuple with the SInscriptiontypeNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSInscriptiontypeNameX

`func (o *InscriptionListElement) SetSInscriptiontypeNameX(v string)`

SetSInscriptiontypeNameX sets SInscriptiontypeNameX field to given value.


### GetEInscriptionStep

`func (o *InscriptionListElement) GetEInscriptionStep() FieldEInscriptionStep`

GetEInscriptionStep returns the EInscriptionStep field if non-nil, zero value otherwise.

### GetEInscriptionStepOk

`func (o *InscriptionListElement) GetEInscriptionStepOk() (*FieldEInscriptionStep, bool)`

GetEInscriptionStepOk returns a tuple with the EInscriptionStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEInscriptionStep

`func (o *InscriptionListElement) SetEInscriptionStep(v FieldEInscriptionStep)`

SetEInscriptionStep sets EInscriptionStep field to given value.


### GetSInscriptionCivicend

`func (o *InscriptionListElement) GetSInscriptionCivicend() string`

GetSInscriptionCivicend returns the SInscriptionCivicend field if non-nil, zero value otherwise.

### GetSInscriptionCivicendOk

`func (o *InscriptionListElement) GetSInscriptionCivicendOk() (*string, bool)`

GetSInscriptionCivicendOk returns a tuple with the SInscriptionCivicend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSInscriptionCivicend

`func (o *InscriptionListElement) SetSInscriptionCivicend(v string)`

SetSInscriptionCivicend sets SInscriptionCivicend field to given value.


### GetSInscriptionMLS

`func (o *InscriptionListElement) GetSInscriptionMLS() string`

GetSInscriptionMLS returns the SInscriptionMLS field if non-nil, zero value otherwise.

### GetSInscriptionMLSOk

`func (o *InscriptionListElement) GetSInscriptionMLSOk() (*string, bool)`

GetSInscriptionMLSOk returns a tuple with the SInscriptionMLS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSInscriptionMLS

`func (o *InscriptionListElement) SetSInscriptionMLS(v string)`

SetSInscriptionMLS sets SInscriptionMLS field to given value.

### HasSInscriptionMLS

`func (o *InscriptionListElement) HasSInscriptionMLS() bool`

HasSInscriptionMLS returns a boolean if a field has been set.

### GetDInscriptionSaleprice

`func (o *InscriptionListElement) GetDInscriptionSaleprice() string`

GetDInscriptionSaleprice returns the DInscriptionSaleprice field if non-nil, zero value otherwise.

### GetDInscriptionSalepriceOk

`func (o *InscriptionListElement) GetDInscriptionSalepriceOk() (*string, bool)`

GetDInscriptionSalepriceOk returns a tuple with the DInscriptionSaleprice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDInscriptionSaleprice

`func (o *InscriptionListElement) SetDInscriptionSaleprice(v string)`

SetDInscriptionSaleprice sets DInscriptionSaleprice field to given value.


### GetDInscriptionRentprice

`func (o *InscriptionListElement) GetDInscriptionRentprice() string`

GetDInscriptionRentprice returns the DInscriptionRentprice field if non-nil, zero value otherwise.

### GetDInscriptionRentpriceOk

`func (o *InscriptionListElement) GetDInscriptionRentpriceOk() (*string, bool)`

GetDInscriptionRentpriceOk returns a tuple with the DInscriptionRentprice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDInscriptionRentprice

`func (o *InscriptionListElement) SetDInscriptionRentprice(v string)`

SetDInscriptionRentprice sets DInscriptionRentprice field to given value.


### GetDtInscriptionDate

`func (o *InscriptionListElement) GetDtInscriptionDate() string`

GetDtInscriptionDate returns the DtInscriptionDate field if non-nil, zero value otherwise.

### GetDtInscriptionDateOk

`func (o *InscriptionListElement) GetDtInscriptionDateOk() (*string, bool)`

GetDtInscriptionDateOk returns a tuple with the DtInscriptionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtInscriptionDate

`func (o *InscriptionListElement) SetDtInscriptionDate(v string)`

SetDtInscriptionDate sets DtInscriptionDate field to given value.

### HasDtInscriptionDate

`func (o *InscriptionListElement) HasDtInscriptionDate() bool`

HasDtInscriptionDate returns a boolean if a field has been set.

### GetDtInscriptionExpirationdate

`func (o *InscriptionListElement) GetDtInscriptionExpirationdate() string`

GetDtInscriptionExpirationdate returns the DtInscriptionExpirationdate field if non-nil, zero value otherwise.

### GetDtInscriptionExpirationdateOk

`func (o *InscriptionListElement) GetDtInscriptionExpirationdateOk() (*string, bool)`

GetDtInscriptionExpirationdateOk returns a tuple with the DtInscriptionExpirationdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtInscriptionExpirationdate

`func (o *InscriptionListElement) SetDtInscriptionExpirationdate(v string)`

SetDtInscriptionExpirationdate sets DtInscriptionExpirationdate field to given value.

### HasDtInscriptionExpirationdate

`func (o *InscriptionListElement) HasDtInscriptionExpirationdate() bool`

HasDtInscriptionExpirationdate returns a boolean if a field has been set.

### GetDtInscriptionNotarydate

`func (o *InscriptionListElement) GetDtInscriptionNotarydate() string`

GetDtInscriptionNotarydate returns the DtInscriptionNotarydate field if non-nil, zero value otherwise.

### GetDtInscriptionNotarydateOk

`func (o *InscriptionListElement) GetDtInscriptionNotarydateOk() (*string, bool)`

GetDtInscriptionNotarydateOk returns a tuple with the DtInscriptionNotarydate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtInscriptionNotarydate

`func (o *InscriptionListElement) SetDtInscriptionNotarydate(v string)`

SetDtInscriptionNotarydate sets DtInscriptionNotarydate field to given value.

### HasDtInscriptionNotarydate

`func (o *InscriptionListElement) HasDtInscriptionNotarydate() bool`

HasDtInscriptionNotarydate returns a boolean if a field has been set.

### GetBInscriptionIsactive

`func (o *InscriptionListElement) GetBInscriptionIsactive() bool`

GetBInscriptionIsactive returns the BInscriptionIsactive field if non-nil, zero value otherwise.

### GetBInscriptionIsactiveOk

`func (o *InscriptionListElement) GetBInscriptionIsactiveOk() (*bool, bool)`

GetBInscriptionIsactiveOk returns a tuple with the BInscriptionIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBInscriptionIsactive

`func (o *InscriptionListElement) SetBInscriptionIsactive(v bool)`

SetBInscriptionIsactive sets BInscriptionIsactive field to given value.


### GetBInscriptionArchived

`func (o *InscriptionListElement) GetBInscriptionArchived() bool`

GetBInscriptionArchived returns the BInscriptionArchived field if non-nil, zero value otherwise.

### GetBInscriptionArchivedOk

`func (o *InscriptionListElement) GetBInscriptionArchivedOk() (*bool, bool)`

GetBInscriptionArchivedOk returns a tuple with the BInscriptionArchived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBInscriptionArchived

`func (o *InscriptionListElement) SetBInscriptionArchived(v bool)`

SetBInscriptionArchived sets BInscriptionArchived field to given value.


### GetBInscriptionInspection

`func (o *InscriptionListElement) GetBInscriptionInspection() bool`

GetBInscriptionInspection returns the BInscriptionInspection field if non-nil, zero value otherwise.

### GetBInscriptionInspectionOk

`func (o *InscriptionListElement) GetBInscriptionInspectionOk() (*bool, bool)`

GetBInscriptionInspectionOk returns a tuple with the BInscriptionInspection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBInscriptionInspection

`func (o *InscriptionListElement) SetBInscriptionInspection(v bool)`

SetBInscriptionInspection sets BInscriptionInspection field to given value.

### HasBInscriptionInspection

`func (o *InscriptionListElement) HasBInscriptionInspection() bool`

HasBInscriptionInspection returns a boolean if a field has been set.

### GetDtInscriptionnotauthenticatedNotaryscheduledate

`func (o *InscriptionListElement) GetDtInscriptionnotauthenticatedNotaryscheduledate() string`

GetDtInscriptionnotauthenticatedNotaryscheduledate returns the DtInscriptionnotauthenticatedNotaryscheduledate field if non-nil, zero value otherwise.

### GetDtInscriptionnotauthenticatedNotaryscheduledateOk

`func (o *InscriptionListElement) GetDtInscriptionnotauthenticatedNotaryscheduledateOk() (*string, bool)`

GetDtInscriptionnotauthenticatedNotaryscheduledateOk returns a tuple with the DtInscriptionnotauthenticatedNotaryscheduledate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtInscriptionnotauthenticatedNotaryscheduledate

`func (o *InscriptionListElement) SetDtInscriptionnotauthenticatedNotaryscheduledate(v string)`

SetDtInscriptionnotauthenticatedNotaryscheduledate sets DtInscriptionnotauthenticatedNotaryscheduledate field to given value.

### HasDtInscriptionnotauthenticatedNotaryscheduledate

`func (o *InscriptionListElement) HasDtInscriptionnotauthenticatedNotaryscheduledate() bool`

HasDtInscriptionnotauthenticatedNotaryscheduledate returns a boolean if a field has been set.

### GetDtInscriptionnotauthenticatedTransactiondate

`func (o *InscriptionListElement) GetDtInscriptionnotauthenticatedTransactiondate() string`

GetDtInscriptionnotauthenticatedTransactiondate returns the DtInscriptionnotauthenticatedTransactiondate field if non-nil, zero value otherwise.

### GetDtInscriptionnotauthenticatedTransactiondateOk

`func (o *InscriptionListElement) GetDtInscriptionnotauthenticatedTransactiondateOk() (*string, bool)`

GetDtInscriptionnotauthenticatedTransactiondateOk returns a tuple with the DtInscriptionnotauthenticatedTransactiondate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtInscriptionnotauthenticatedTransactiondate

`func (o *InscriptionListElement) SetDtInscriptionnotauthenticatedTransactiondate(v string)`

SetDtInscriptionnotauthenticatedTransactiondate sets DtInscriptionnotauthenticatedTransactiondate field to given value.

### HasDtInscriptionnotauthenticatedTransactiondate

`func (o *InscriptionListElement) HasDtInscriptionnotauthenticatedTransactiondate() bool`

HasDtInscriptionnotauthenticatedTransactiondate returns a boolean if a field has been set.

### GetDtInscriptionnotauthenticatedTransactiondateReal

`func (o *InscriptionListElement) GetDtInscriptionnotauthenticatedTransactiondateReal() string`

GetDtInscriptionnotauthenticatedTransactiondateReal returns the DtInscriptionnotauthenticatedTransactiondateReal field if non-nil, zero value otherwise.

### GetDtInscriptionnotauthenticatedTransactiondateRealOk

`func (o *InscriptionListElement) GetDtInscriptionnotauthenticatedTransactiondateRealOk() (*string, bool)`

GetDtInscriptionnotauthenticatedTransactiondateRealOk returns a tuple with the DtInscriptionnotauthenticatedTransactiondateReal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtInscriptionnotauthenticatedTransactiondateReal

`func (o *InscriptionListElement) SetDtInscriptionnotauthenticatedTransactiondateReal(v string)`

SetDtInscriptionnotauthenticatedTransactiondateReal sets DtInscriptionnotauthenticatedTransactiondateReal field to given value.

### HasDtInscriptionnotauthenticatedTransactiondateReal

`func (o *InscriptionListElement) HasDtInscriptionnotauthenticatedTransactiondateReal() bool`

HasDtInscriptionnotauthenticatedTransactiondateReal returns a boolean if a field has been set.

### GetBInscriptionnotauthenticatedConditional

`func (o *InscriptionListElement) GetBInscriptionnotauthenticatedConditional() bool`

GetBInscriptionnotauthenticatedConditional returns the BInscriptionnotauthenticatedConditional field if non-nil, zero value otherwise.

### GetBInscriptionnotauthenticatedConditionalOk

`func (o *InscriptionListElement) GetBInscriptionnotauthenticatedConditionalOk() (*bool, bool)`

GetBInscriptionnotauthenticatedConditionalOk returns a tuple with the BInscriptionnotauthenticatedConditional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBInscriptionnotauthenticatedConditional

`func (o *InscriptionListElement) SetBInscriptionnotauthenticatedConditional(v bool)`

SetBInscriptionnotauthenticatedConditional sets BInscriptionnotauthenticatedConditional field to given value.

### HasBInscriptionnotauthenticatedConditional

`func (o *InscriptionListElement) HasBInscriptionnotauthenticatedConditional() bool`

HasBInscriptionnotauthenticatedConditional returns a boolean if a field has been set.

### GetBInscriptionnotauthenticatedIsactive

`func (o *InscriptionListElement) GetBInscriptionnotauthenticatedIsactive() bool`

GetBInscriptionnotauthenticatedIsactive returns the BInscriptionnotauthenticatedIsactive field if non-nil, zero value otherwise.

### GetBInscriptionnotauthenticatedIsactiveOk

`func (o *InscriptionListElement) GetBInscriptionnotauthenticatedIsactiveOk() (*bool, bool)`

GetBInscriptionnotauthenticatedIsactiveOk returns a tuple with the BInscriptionnotauthenticatedIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBInscriptionnotauthenticatedIsactive

`func (o *InscriptionListElement) SetBInscriptionnotauthenticatedIsactive(v bool)`

SetBInscriptionnotauthenticatedIsactive sets BInscriptionnotauthenticatedIsactive field to given value.

### HasBInscriptionnotauthenticatedIsactive

`func (o *InscriptionListElement) HasBInscriptionnotauthenticatedIsactive() bool`

HasBInscriptionnotauthenticatedIsactive returns a boolean if a field has been set.

### GetSAddressCivic

`func (o *InscriptionListElement) GetSAddressCivic() string`

GetSAddressCivic returns the SAddressCivic field if non-nil, zero value otherwise.

### GetSAddressCivicOk

`func (o *InscriptionListElement) GetSAddressCivicOk() (*string, bool)`

GetSAddressCivicOk returns a tuple with the SAddressCivic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAddressCivic

`func (o *InscriptionListElement) SetSAddressCivic(v string)`

SetSAddressCivic sets SAddressCivic field to given value.

### HasSAddressCivic

`func (o *InscriptionListElement) HasSAddressCivic() bool`

HasSAddressCivic returns a boolean if a field has been set.

### GetSAddressStreet

`func (o *InscriptionListElement) GetSAddressStreet() string`

GetSAddressStreet returns the SAddressStreet field if non-nil, zero value otherwise.

### GetSAddressStreetOk

`func (o *InscriptionListElement) GetSAddressStreetOk() (*string, bool)`

GetSAddressStreetOk returns a tuple with the SAddressStreet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAddressStreet

`func (o *InscriptionListElement) SetSAddressStreet(v string)`

SetSAddressStreet sets SAddressStreet field to given value.

### HasSAddressStreet

`func (o *InscriptionListElement) HasSAddressStreet() bool`

HasSAddressStreet returns a boolean if a field has been set.

### GetSAddressSuite

`func (o *InscriptionListElement) GetSAddressSuite() string`

GetSAddressSuite returns the SAddressSuite field if non-nil, zero value otherwise.

### GetSAddressSuiteOk

`func (o *InscriptionListElement) GetSAddressSuiteOk() (*string, bool)`

GetSAddressSuiteOk returns a tuple with the SAddressSuite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAddressSuite

`func (o *InscriptionListElement) SetSAddressSuite(v string)`

SetSAddressSuite sets SAddressSuite field to given value.

### HasSAddressSuite

`func (o *InscriptionListElement) HasSAddressSuite() bool`

HasSAddressSuite returns a boolean if a field has been set.

### GetSAddressCity

`func (o *InscriptionListElement) GetSAddressCity() string`

GetSAddressCity returns the SAddressCity field if non-nil, zero value otherwise.

### GetSAddressCityOk

`func (o *InscriptionListElement) GetSAddressCityOk() (*string, bool)`

GetSAddressCityOk returns a tuple with the SAddressCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAddressCity

`func (o *InscriptionListElement) SetSAddressCity(v string)`

SetSAddressCity sets SAddressCity field to given value.

### HasSAddressCity

`func (o *InscriptionListElement) HasSAddressCity() bool`

HasSAddressCity returns a boolean if a field has been set.

### GetSAddressZip

`func (o *InscriptionListElement) GetSAddressZip() string`

GetSAddressZip returns the SAddressZip field if non-nil, zero value otherwise.

### GetSAddressZipOk

`func (o *InscriptionListElement) GetSAddressZipOk() (*string, bool)`

GetSAddressZipOk returns a tuple with the SAddressZip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAddressZip

`func (o *InscriptionListElement) SetSAddressZip(v string)`

SetSAddressZip sets SAddressZip field to given value.

### HasSAddressZip

`func (o *InscriptionListElement) HasSAddressZip() bool`

HasSAddressZip returns a boolean if a field has been set.

### GetFkiProvinceID

`func (o *InscriptionListElement) GetFkiProvinceID() int32`

GetFkiProvinceID returns the FkiProvinceID field if non-nil, zero value otherwise.

### GetFkiProvinceIDOk

`func (o *InscriptionListElement) GetFkiProvinceIDOk() (*int32, bool)`

GetFkiProvinceIDOk returns a tuple with the FkiProvinceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiProvinceID

`func (o *InscriptionListElement) SetFkiProvinceID(v int32)`

SetFkiProvinceID sets FkiProvinceID field to given value.

### HasFkiProvinceID

`func (o *InscriptionListElement) HasFkiProvinceID() bool`

HasFkiProvinceID returns a boolean if a field has been set.

### GetSProvinceNameX

`func (o *InscriptionListElement) GetSProvinceNameX() string`

GetSProvinceNameX returns the SProvinceNameX field if non-nil, zero value otherwise.

### GetSProvinceNameXOk

`func (o *InscriptionListElement) GetSProvinceNameXOk() (*string, bool)`

GetSProvinceNameXOk returns a tuple with the SProvinceNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSProvinceNameX

`func (o *InscriptionListElement) SetSProvinceNameX(v string)`

SetSProvinceNameX sets SProvinceNameX field to given value.

### HasSProvinceNameX

`func (o *InscriptionListElement) HasSProvinceNameX() bool`

HasSProvinceNameX returns a boolean if a field has been set.

### GetFkiCountryID

`func (o *InscriptionListElement) GetFkiCountryID() int32`

GetFkiCountryID returns the FkiCountryID field if non-nil, zero value otherwise.

### GetFkiCountryIDOk

`func (o *InscriptionListElement) GetFkiCountryIDOk() (*int32, bool)`

GetFkiCountryIDOk returns a tuple with the FkiCountryID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiCountryID

`func (o *InscriptionListElement) SetFkiCountryID(v int32)`

SetFkiCountryID sets FkiCountryID field to given value.

### HasFkiCountryID

`func (o *InscriptionListElement) HasFkiCountryID() bool`

HasFkiCountryID returns a boolean if a field has been set.

### GetSCountryNameX

`func (o *InscriptionListElement) GetSCountryNameX() string`

GetSCountryNameX returns the SCountryNameX field if non-nil, zero value otherwise.

### GetSCountryNameXOk

`func (o *InscriptionListElement) GetSCountryNameXOk() (*string, bool)`

GetSCountryNameXOk returns a tuple with the SCountryNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSCountryNameX

`func (o *InscriptionListElement) SetSCountryNameX(v string)`

SetSCountryNameX sets SCountryNameX field to given value.

### HasSCountryNameX

`func (o *InscriptionListElement) HasSCountryNameX() bool`

HasSCountryNameX returns a boolean if a field has been set.

### GetIInscriptionnotauthenticatedCanceled

`func (o *InscriptionListElement) GetIInscriptionnotauthenticatedCanceled() int32`

GetIInscriptionnotauthenticatedCanceled returns the IInscriptionnotauthenticatedCanceled field if non-nil, zero value otherwise.

### GetIInscriptionnotauthenticatedCanceledOk

`func (o *InscriptionListElement) GetIInscriptionnotauthenticatedCanceledOk() (*int32, bool)`

GetIInscriptionnotauthenticatedCanceledOk returns a tuple with the IInscriptionnotauthenticatedCanceled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIInscriptionnotauthenticatedCanceled

`func (o *InscriptionListElement) SetIInscriptionnotauthenticatedCanceled(v int32)`

SetIInscriptionnotauthenticatedCanceled sets IInscriptionnotauthenticatedCanceled field to given value.


### GetBAllowedCopyintoinscriptionedm

`func (o *InscriptionListElement) GetBAllowedCopyintoinscriptionedm() bool`

GetBAllowedCopyintoinscriptionedm returns the BAllowedCopyintoinscriptionedm field if non-nil, zero value otherwise.

### GetBAllowedCopyintoinscriptionedmOk

`func (o *InscriptionListElement) GetBAllowedCopyintoinscriptionedmOk() (*bool, bool)`

GetBAllowedCopyintoinscriptionedmOk returns a tuple with the BAllowedCopyintoinscriptionedm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBAllowedCopyintoinscriptionedm

`func (o *InscriptionListElement) SetBAllowedCopyintoinscriptionedm(v bool)`

SetBAllowedCopyintoinscriptionedm sets BAllowedCopyintoinscriptionedm field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


