# SupplierListElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiSupplierID** | **int32** | The unique ID of the Supplier. | 
**FkiPaymentmethodID** | Pointer to **int32** | The unique ID of the Paymentmethod | [optional] 
**SSupplierName** | **string** | The name of the Supplier | 
**SSupplierCode** | **string** | The code of the Supplier | 
**SSupplierAccount** | **string** | The account of the Supplier | 
**BSupplierIsactive** | **bool** | Whether the supplier is active or not | 
**SPhoneE164** | Pointer to **string** | A phone number in E.164 Format | [optional] 
**SEmailAddress** | Pointer to **string** | The email address. | [optional] 
**SAddressCivic** | Pointer to **string** | The Civic number. | [optional] 
**SAddressStreet** | Pointer to **string** | The Street Name | [optional] 
**SAddressSuite** | Pointer to **string** | The Suite or appartment number | [optional] 
**SAddressCity** | Pointer to **string** | The City name | [optional] 
**SAddressZip** | Pointer to **string** | The Postal/Zip Code  The value must be entered without spaces | [optional] 
**FkiProvinceID** | Pointer to **int32** | The unique ID of the Province.  Here are some common values (Complete list must be retrieved from API):  |Value|Description| |-|-| |1|(Canada) Alberta |2|(Canada) British Columbia| |3|(Canada) Manitoba| |3|(Canada) Manitoba| |4|(Canada) New Brunswick| |5|(Canada) Newfoundland| |6|(Canada) Northwest Territories| |7|(Canada) Nova Scotia| |8|(Canada) Nunavut| |9|(Canada) Ontario| |10|(Canada) Prince Edward Island| |11|(Canada) Quebec| |12|(Canada) Saskatchewan| |13|(Canada) Yukon| |14|(United-States) Alabama| |15|(United-States) Alaska| |16|(United-States) Arizona| |17|(United-States) Arkansas| |18|(United-States) California| |19|(United-States) Colorado| |20|(United-States) Connecticut| |21|(United-States) Delaware| |22|(United-States) District of Columbia| |23|(United-States) Florida| |24|(United-States) Georgia| |25|(United-States) Hawaii| |26|(United-States) Idaho| |27|(United-States) Illinois| |28|(United-States) Indiana| |29|(United-States) Iowa| |30|(United-States) Kansas| |31|(United-States) Kentucky| |32|(United-States) Louisiane| |33|(United-States) Maine| |34|(United-States) Maryland| |35|(United-States) Massachusetts| |36|(United-States) Michigan| |37|(United-States) Minnesota| |38|(United-States) Mississippi| |39|(United-States) Missouri| |40|(United-States) Montana| |41|(United-States) Nebraska| |42|(United-States) Nevada| |43|(United-States) New Hampshire| |44|(United-States) New Jersey| |45|(United-States) New Mexico| |46|(United-States) New York| |47|(United-States) North Carolina| |48|(United-States) North Dakota| |49|(United-States) Ohio| |50|(United-States) Oklahoma| |51|(United-States) Oregon| |52|(United-States) Pennsylvania| |53|(United-States) Rhode Island| |54|(United-States) South Carolina| |55|(United-States) South Dakota| |56|(United-States) Tennessee| |57|(United-States) Texas| |58|(United-States) Utah| |60|(United-States) Vermont| |59|(United-States) Virginia| |61|(United-States) Washington| |62|(United-States) West Virginia| |63|(United-States) Wisconsin| |64|(United-States) Wyoming| | [optional] 
**SProvinceNameX** | Pointer to **string** | The name of the Province in the language of the requester | [optional] 
**FkiCountryID** | Pointer to **int32** | The unique ID of the Country.  Here are some common values (Complete list must be retrieved from API):  |Value|Description| |-|-| |1|Canada| |2|United-States| | [optional] 
**SCountryNameX** | Pointer to **string** | The name of the Country in the language of the requester | [optional] 
**SPaymentmethodDescriptionX** | Pointer to **string** | The description of the Paymentmethod in the language of the requester | [optional] 
**SElectronicfundstransferbankaccountTransit** | Pointer to **string** | The transit of the Electronicfundstransferbankaccount | [optional] 
**SElectronicfundstransferbankaccountInstitution** | Pointer to **string** | The institution of the Electronicfundstransferbankaccount | [optional] 
**SElectronicfundstransferbankaccountAccount** | Pointer to **string** | The account of the Electronicfundstransferbankaccount | [optional] 
**SGlaccountcontainerLongcode** | **string** | The Code for the Glaccountcontainer | 
**SGlaccountcontainerLongdescriptionX** | **string** | The Description for the Glaccountcontainer in the language of the requester | 

## Methods

### NewSupplierListElement

`func NewSupplierListElement(pkiSupplierID int32, sSupplierName string, sSupplierCode string, sSupplierAccount string, bSupplierIsactive bool, sGlaccountcontainerLongcode string, sGlaccountcontainerLongdescriptionX string, ) *SupplierListElement`

NewSupplierListElement instantiates a new SupplierListElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupplierListElementWithDefaults

`func NewSupplierListElementWithDefaults() *SupplierListElement`

NewSupplierListElementWithDefaults instantiates a new SupplierListElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiSupplierID

`func (o *SupplierListElement) GetPkiSupplierID() int32`

GetPkiSupplierID returns the PkiSupplierID field if non-nil, zero value otherwise.

### GetPkiSupplierIDOk

`func (o *SupplierListElement) GetPkiSupplierIDOk() (*int32, bool)`

GetPkiSupplierIDOk returns a tuple with the PkiSupplierID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiSupplierID

`func (o *SupplierListElement) SetPkiSupplierID(v int32)`

SetPkiSupplierID sets PkiSupplierID field to given value.


### GetFkiPaymentmethodID

`func (o *SupplierListElement) GetFkiPaymentmethodID() int32`

GetFkiPaymentmethodID returns the FkiPaymentmethodID field if non-nil, zero value otherwise.

### GetFkiPaymentmethodIDOk

`func (o *SupplierListElement) GetFkiPaymentmethodIDOk() (*int32, bool)`

GetFkiPaymentmethodIDOk returns a tuple with the FkiPaymentmethodID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiPaymentmethodID

`func (o *SupplierListElement) SetFkiPaymentmethodID(v int32)`

SetFkiPaymentmethodID sets FkiPaymentmethodID field to given value.

### HasFkiPaymentmethodID

`func (o *SupplierListElement) HasFkiPaymentmethodID() bool`

HasFkiPaymentmethodID returns a boolean if a field has been set.

### GetSSupplierName

`func (o *SupplierListElement) GetSSupplierName() string`

GetSSupplierName returns the SSupplierName field if non-nil, zero value otherwise.

### GetSSupplierNameOk

`func (o *SupplierListElement) GetSSupplierNameOk() (*string, bool)`

GetSSupplierNameOk returns a tuple with the SSupplierName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSSupplierName

`func (o *SupplierListElement) SetSSupplierName(v string)`

SetSSupplierName sets SSupplierName field to given value.


### GetSSupplierCode

`func (o *SupplierListElement) GetSSupplierCode() string`

GetSSupplierCode returns the SSupplierCode field if non-nil, zero value otherwise.

### GetSSupplierCodeOk

`func (o *SupplierListElement) GetSSupplierCodeOk() (*string, bool)`

GetSSupplierCodeOk returns a tuple with the SSupplierCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSSupplierCode

`func (o *SupplierListElement) SetSSupplierCode(v string)`

SetSSupplierCode sets SSupplierCode field to given value.


### GetSSupplierAccount

`func (o *SupplierListElement) GetSSupplierAccount() string`

GetSSupplierAccount returns the SSupplierAccount field if non-nil, zero value otherwise.

### GetSSupplierAccountOk

`func (o *SupplierListElement) GetSSupplierAccountOk() (*string, bool)`

GetSSupplierAccountOk returns a tuple with the SSupplierAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSSupplierAccount

`func (o *SupplierListElement) SetSSupplierAccount(v string)`

SetSSupplierAccount sets SSupplierAccount field to given value.


### GetBSupplierIsactive

`func (o *SupplierListElement) GetBSupplierIsactive() bool`

GetBSupplierIsactive returns the BSupplierIsactive field if non-nil, zero value otherwise.

### GetBSupplierIsactiveOk

`func (o *SupplierListElement) GetBSupplierIsactiveOk() (*bool, bool)`

GetBSupplierIsactiveOk returns a tuple with the BSupplierIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBSupplierIsactive

`func (o *SupplierListElement) SetBSupplierIsactive(v bool)`

SetBSupplierIsactive sets BSupplierIsactive field to given value.


### GetSPhoneE164

`func (o *SupplierListElement) GetSPhoneE164() string`

GetSPhoneE164 returns the SPhoneE164 field if non-nil, zero value otherwise.

### GetSPhoneE164Ok

`func (o *SupplierListElement) GetSPhoneE164Ok() (*string, bool)`

GetSPhoneE164Ok returns a tuple with the SPhoneE164 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSPhoneE164

`func (o *SupplierListElement) SetSPhoneE164(v string)`

SetSPhoneE164 sets SPhoneE164 field to given value.

### HasSPhoneE164

`func (o *SupplierListElement) HasSPhoneE164() bool`

HasSPhoneE164 returns a boolean if a field has been set.

### GetSEmailAddress

`func (o *SupplierListElement) GetSEmailAddress() string`

GetSEmailAddress returns the SEmailAddress field if non-nil, zero value otherwise.

### GetSEmailAddressOk

`func (o *SupplierListElement) GetSEmailAddressOk() (*string, bool)`

GetSEmailAddressOk returns a tuple with the SEmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEmailAddress

`func (o *SupplierListElement) SetSEmailAddress(v string)`

SetSEmailAddress sets SEmailAddress field to given value.

### HasSEmailAddress

`func (o *SupplierListElement) HasSEmailAddress() bool`

HasSEmailAddress returns a boolean if a field has been set.

### GetSAddressCivic

`func (o *SupplierListElement) GetSAddressCivic() string`

GetSAddressCivic returns the SAddressCivic field if non-nil, zero value otherwise.

### GetSAddressCivicOk

`func (o *SupplierListElement) GetSAddressCivicOk() (*string, bool)`

GetSAddressCivicOk returns a tuple with the SAddressCivic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAddressCivic

`func (o *SupplierListElement) SetSAddressCivic(v string)`

SetSAddressCivic sets SAddressCivic field to given value.

### HasSAddressCivic

`func (o *SupplierListElement) HasSAddressCivic() bool`

HasSAddressCivic returns a boolean if a field has been set.

### GetSAddressStreet

`func (o *SupplierListElement) GetSAddressStreet() string`

GetSAddressStreet returns the SAddressStreet field if non-nil, zero value otherwise.

### GetSAddressStreetOk

`func (o *SupplierListElement) GetSAddressStreetOk() (*string, bool)`

GetSAddressStreetOk returns a tuple with the SAddressStreet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAddressStreet

`func (o *SupplierListElement) SetSAddressStreet(v string)`

SetSAddressStreet sets SAddressStreet field to given value.

### HasSAddressStreet

`func (o *SupplierListElement) HasSAddressStreet() bool`

HasSAddressStreet returns a boolean if a field has been set.

### GetSAddressSuite

`func (o *SupplierListElement) GetSAddressSuite() string`

GetSAddressSuite returns the SAddressSuite field if non-nil, zero value otherwise.

### GetSAddressSuiteOk

`func (o *SupplierListElement) GetSAddressSuiteOk() (*string, bool)`

GetSAddressSuiteOk returns a tuple with the SAddressSuite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAddressSuite

`func (o *SupplierListElement) SetSAddressSuite(v string)`

SetSAddressSuite sets SAddressSuite field to given value.

### HasSAddressSuite

`func (o *SupplierListElement) HasSAddressSuite() bool`

HasSAddressSuite returns a boolean if a field has been set.

### GetSAddressCity

`func (o *SupplierListElement) GetSAddressCity() string`

GetSAddressCity returns the SAddressCity field if non-nil, zero value otherwise.

### GetSAddressCityOk

`func (o *SupplierListElement) GetSAddressCityOk() (*string, bool)`

GetSAddressCityOk returns a tuple with the SAddressCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAddressCity

`func (o *SupplierListElement) SetSAddressCity(v string)`

SetSAddressCity sets SAddressCity field to given value.

### HasSAddressCity

`func (o *SupplierListElement) HasSAddressCity() bool`

HasSAddressCity returns a boolean if a field has been set.

### GetSAddressZip

`func (o *SupplierListElement) GetSAddressZip() string`

GetSAddressZip returns the SAddressZip field if non-nil, zero value otherwise.

### GetSAddressZipOk

`func (o *SupplierListElement) GetSAddressZipOk() (*string, bool)`

GetSAddressZipOk returns a tuple with the SAddressZip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAddressZip

`func (o *SupplierListElement) SetSAddressZip(v string)`

SetSAddressZip sets SAddressZip field to given value.

### HasSAddressZip

`func (o *SupplierListElement) HasSAddressZip() bool`

HasSAddressZip returns a boolean if a field has been set.

### GetFkiProvinceID

`func (o *SupplierListElement) GetFkiProvinceID() int32`

GetFkiProvinceID returns the FkiProvinceID field if non-nil, zero value otherwise.

### GetFkiProvinceIDOk

`func (o *SupplierListElement) GetFkiProvinceIDOk() (*int32, bool)`

GetFkiProvinceIDOk returns a tuple with the FkiProvinceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiProvinceID

`func (o *SupplierListElement) SetFkiProvinceID(v int32)`

SetFkiProvinceID sets FkiProvinceID field to given value.

### HasFkiProvinceID

`func (o *SupplierListElement) HasFkiProvinceID() bool`

HasFkiProvinceID returns a boolean if a field has been set.

### GetSProvinceNameX

`func (o *SupplierListElement) GetSProvinceNameX() string`

GetSProvinceNameX returns the SProvinceNameX field if non-nil, zero value otherwise.

### GetSProvinceNameXOk

`func (o *SupplierListElement) GetSProvinceNameXOk() (*string, bool)`

GetSProvinceNameXOk returns a tuple with the SProvinceNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSProvinceNameX

`func (o *SupplierListElement) SetSProvinceNameX(v string)`

SetSProvinceNameX sets SProvinceNameX field to given value.

### HasSProvinceNameX

`func (o *SupplierListElement) HasSProvinceNameX() bool`

HasSProvinceNameX returns a boolean if a field has been set.

### GetFkiCountryID

`func (o *SupplierListElement) GetFkiCountryID() int32`

GetFkiCountryID returns the FkiCountryID field if non-nil, zero value otherwise.

### GetFkiCountryIDOk

`func (o *SupplierListElement) GetFkiCountryIDOk() (*int32, bool)`

GetFkiCountryIDOk returns a tuple with the FkiCountryID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiCountryID

`func (o *SupplierListElement) SetFkiCountryID(v int32)`

SetFkiCountryID sets FkiCountryID field to given value.

### HasFkiCountryID

`func (o *SupplierListElement) HasFkiCountryID() bool`

HasFkiCountryID returns a boolean if a field has been set.

### GetSCountryNameX

`func (o *SupplierListElement) GetSCountryNameX() string`

GetSCountryNameX returns the SCountryNameX field if non-nil, zero value otherwise.

### GetSCountryNameXOk

`func (o *SupplierListElement) GetSCountryNameXOk() (*string, bool)`

GetSCountryNameXOk returns a tuple with the SCountryNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSCountryNameX

`func (o *SupplierListElement) SetSCountryNameX(v string)`

SetSCountryNameX sets SCountryNameX field to given value.

### HasSCountryNameX

`func (o *SupplierListElement) HasSCountryNameX() bool`

HasSCountryNameX returns a boolean if a field has been set.

### GetSPaymentmethodDescriptionX

`func (o *SupplierListElement) GetSPaymentmethodDescriptionX() string`

GetSPaymentmethodDescriptionX returns the SPaymentmethodDescriptionX field if non-nil, zero value otherwise.

### GetSPaymentmethodDescriptionXOk

`func (o *SupplierListElement) GetSPaymentmethodDescriptionXOk() (*string, bool)`

GetSPaymentmethodDescriptionXOk returns a tuple with the SPaymentmethodDescriptionX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSPaymentmethodDescriptionX

`func (o *SupplierListElement) SetSPaymentmethodDescriptionX(v string)`

SetSPaymentmethodDescriptionX sets SPaymentmethodDescriptionX field to given value.

### HasSPaymentmethodDescriptionX

`func (o *SupplierListElement) HasSPaymentmethodDescriptionX() bool`

HasSPaymentmethodDescriptionX returns a boolean if a field has been set.

### GetSElectronicfundstransferbankaccountTransit

`func (o *SupplierListElement) GetSElectronicfundstransferbankaccountTransit() string`

GetSElectronicfundstransferbankaccountTransit returns the SElectronicfundstransferbankaccountTransit field if non-nil, zero value otherwise.

### GetSElectronicfundstransferbankaccountTransitOk

`func (o *SupplierListElement) GetSElectronicfundstransferbankaccountTransitOk() (*string, bool)`

GetSElectronicfundstransferbankaccountTransitOk returns a tuple with the SElectronicfundstransferbankaccountTransit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSElectronicfundstransferbankaccountTransit

`func (o *SupplierListElement) SetSElectronicfundstransferbankaccountTransit(v string)`

SetSElectronicfundstransferbankaccountTransit sets SElectronicfundstransferbankaccountTransit field to given value.

### HasSElectronicfundstransferbankaccountTransit

`func (o *SupplierListElement) HasSElectronicfundstransferbankaccountTransit() bool`

HasSElectronicfundstransferbankaccountTransit returns a boolean if a field has been set.

### GetSElectronicfundstransferbankaccountInstitution

`func (o *SupplierListElement) GetSElectronicfundstransferbankaccountInstitution() string`

GetSElectronicfundstransferbankaccountInstitution returns the SElectronicfundstransferbankaccountInstitution field if non-nil, zero value otherwise.

### GetSElectronicfundstransferbankaccountInstitutionOk

`func (o *SupplierListElement) GetSElectronicfundstransferbankaccountInstitutionOk() (*string, bool)`

GetSElectronicfundstransferbankaccountInstitutionOk returns a tuple with the SElectronicfundstransferbankaccountInstitution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSElectronicfundstransferbankaccountInstitution

`func (o *SupplierListElement) SetSElectronicfundstransferbankaccountInstitution(v string)`

SetSElectronicfundstransferbankaccountInstitution sets SElectronicfundstransferbankaccountInstitution field to given value.

### HasSElectronicfundstransferbankaccountInstitution

`func (o *SupplierListElement) HasSElectronicfundstransferbankaccountInstitution() bool`

HasSElectronicfundstransferbankaccountInstitution returns a boolean if a field has been set.

### GetSElectronicfundstransferbankaccountAccount

`func (o *SupplierListElement) GetSElectronicfundstransferbankaccountAccount() string`

GetSElectronicfundstransferbankaccountAccount returns the SElectronicfundstransferbankaccountAccount field if non-nil, zero value otherwise.

### GetSElectronicfundstransferbankaccountAccountOk

`func (o *SupplierListElement) GetSElectronicfundstransferbankaccountAccountOk() (*string, bool)`

GetSElectronicfundstransferbankaccountAccountOk returns a tuple with the SElectronicfundstransferbankaccountAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSElectronicfundstransferbankaccountAccount

`func (o *SupplierListElement) SetSElectronicfundstransferbankaccountAccount(v string)`

SetSElectronicfundstransferbankaccountAccount sets SElectronicfundstransferbankaccountAccount field to given value.

### HasSElectronicfundstransferbankaccountAccount

`func (o *SupplierListElement) HasSElectronicfundstransferbankaccountAccount() bool`

HasSElectronicfundstransferbankaccountAccount returns a boolean if a field has been set.

### GetSGlaccountcontainerLongcode

`func (o *SupplierListElement) GetSGlaccountcontainerLongcode() string`

GetSGlaccountcontainerLongcode returns the SGlaccountcontainerLongcode field if non-nil, zero value otherwise.

### GetSGlaccountcontainerLongcodeOk

`func (o *SupplierListElement) GetSGlaccountcontainerLongcodeOk() (*string, bool)`

GetSGlaccountcontainerLongcodeOk returns a tuple with the SGlaccountcontainerLongcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSGlaccountcontainerLongcode

`func (o *SupplierListElement) SetSGlaccountcontainerLongcode(v string)`

SetSGlaccountcontainerLongcode sets SGlaccountcontainerLongcode field to given value.


### GetSGlaccountcontainerLongdescriptionX

`func (o *SupplierListElement) GetSGlaccountcontainerLongdescriptionX() string`

GetSGlaccountcontainerLongdescriptionX returns the SGlaccountcontainerLongdescriptionX field if non-nil, zero value otherwise.

### GetSGlaccountcontainerLongdescriptionXOk

`func (o *SupplierListElement) GetSGlaccountcontainerLongdescriptionXOk() (*string, bool)`

GetSGlaccountcontainerLongdescriptionXOk returns a tuple with the SGlaccountcontainerLongdescriptionX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSGlaccountcontainerLongdescriptionX

`func (o *SupplierListElement) SetSGlaccountcontainerLongdescriptionX(v string)`

SetSGlaccountcontainerLongdescriptionX sets SGlaccountcontainerLongdescriptionX field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


