# NotaryListElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiNotaryID** | **int32** | The unique ID of the Notary. | 
**FkiPaymentmethodID** | Pointer to **int32** | The unique ID of the Paymentmethod | [optional] 
**FkiProvinceID** | Pointer to **int32** | The unique ID of the Province.  Here are some common values (Complete list must be retrieved from API):  |Value|Description| |-|-| |1|(Canada) Alberta |2|(Canada) British Columbia| |3|(Canada) Manitoba| |3|(Canada) Manitoba| |4|(Canada) New Brunswick| |5|(Canada) Newfoundland| |6|(Canada) Northwest Territories| |7|(Canada) Nova Scotia| |8|(Canada) Nunavut| |9|(Canada) Ontario| |10|(Canada) Prince Edward Island| |11|(Canada) Quebec| |12|(Canada) Saskatchewan| |13|(Canada) Yukon| |14|(United-States) Alabama| |15|(United-States) Alaska| |16|(United-States) Arizona| |17|(United-States) Arkansas| |18|(United-States) California| |19|(United-States) Colorado| |20|(United-States) Connecticut| |21|(United-States) Delaware| |22|(United-States) District of Columbia| |23|(United-States) Florida| |24|(United-States) Georgia| |25|(United-States) Hawaii| |26|(United-States) Idaho| |27|(United-States) Illinois| |28|(United-States) Indiana| |29|(United-States) Iowa| |30|(United-States) Kansas| |31|(United-States) Kentucky| |32|(United-States) Louisiane| |33|(United-States) Maine| |34|(United-States) Maryland| |35|(United-States) Massachusetts| |36|(United-States) Michigan| |37|(United-States) Minnesota| |38|(United-States) Mississippi| |39|(United-States) Missouri| |40|(United-States) Montana| |41|(United-States) Nebraska| |42|(United-States) Nevada| |43|(United-States) New Hampshire| |44|(United-States) New Jersey| |45|(United-States) New Mexico| |46|(United-States) New York| |47|(United-States) North Carolina| |48|(United-States) North Dakota| |49|(United-States) Ohio| |50|(United-States) Oklahoma| |51|(United-States) Oregon| |52|(United-States) Pennsylvania| |53|(United-States) Rhode Island| |54|(United-States) South Carolina| |55|(United-States) South Dakota| |56|(United-States) Tennessee| |57|(United-States) Texas| |58|(United-States) Utah| |60|(United-States) Vermont| |59|(United-States) Virginia| |61|(United-States) Washington| |62|(United-States) West Virginia| |63|(United-States) Wisconsin| |64|(United-States) Wyoming| | [optional] 
**FkiCountryID** | Pointer to **int32** | The unique ID of the Country.  Here are some common values (Complete list must be retrieved from API):  |Value|Description| |-|-| |1|Canada| |2|United-States| | [optional] 
**SContactLastname** | Pointer to **string** | The Last name of the contact | [optional] 
**SContactFirstname** | Pointer to **string** | The First name of the contact | [optional] 
**SContactCompany** | Pointer to **string** | The Company name of the contact | [optional] 
**SAddressCivic** | Pointer to **string** | The Civic number. | [optional] 
**SAddressStreet** | Pointer to **string** | The Street Name | [optional] 
**SAddressSuite** | Pointer to **string** | The Suite or appartment number | [optional] 
**SAddressCity** | Pointer to **string** | The City name | [optional] 
**SPhoneE164** | Pointer to **string** | A phone number in E.164 Format | [optional] 
**SPaymentmethodDescriptionX** | Pointer to **string** | The description of the Paymentmethod in the language of the requester | [optional] 
**SProvinceNameX** | Pointer to **string** | The name of the Province in the language of the requester | [optional] 
**SCountryNameX** | Pointer to **string** | The name of the Country in the language of the requester | [optional] 

## Methods

### NewNotaryListElement

`func NewNotaryListElement(pkiNotaryID int32, ) *NotaryListElement`

NewNotaryListElement instantiates a new NotaryListElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotaryListElementWithDefaults

`func NewNotaryListElementWithDefaults() *NotaryListElement`

NewNotaryListElementWithDefaults instantiates a new NotaryListElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiNotaryID

`func (o *NotaryListElement) GetPkiNotaryID() int32`

GetPkiNotaryID returns the PkiNotaryID field if non-nil, zero value otherwise.

### GetPkiNotaryIDOk

`func (o *NotaryListElement) GetPkiNotaryIDOk() (*int32, bool)`

GetPkiNotaryIDOk returns a tuple with the PkiNotaryID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiNotaryID

`func (o *NotaryListElement) SetPkiNotaryID(v int32)`

SetPkiNotaryID sets PkiNotaryID field to given value.


### GetFkiPaymentmethodID

`func (o *NotaryListElement) GetFkiPaymentmethodID() int32`

GetFkiPaymentmethodID returns the FkiPaymentmethodID field if non-nil, zero value otherwise.

### GetFkiPaymentmethodIDOk

`func (o *NotaryListElement) GetFkiPaymentmethodIDOk() (*int32, bool)`

GetFkiPaymentmethodIDOk returns a tuple with the FkiPaymentmethodID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiPaymentmethodID

`func (o *NotaryListElement) SetFkiPaymentmethodID(v int32)`

SetFkiPaymentmethodID sets FkiPaymentmethodID field to given value.

### HasFkiPaymentmethodID

`func (o *NotaryListElement) HasFkiPaymentmethodID() bool`

HasFkiPaymentmethodID returns a boolean if a field has been set.

### GetFkiProvinceID

`func (o *NotaryListElement) GetFkiProvinceID() int32`

GetFkiProvinceID returns the FkiProvinceID field if non-nil, zero value otherwise.

### GetFkiProvinceIDOk

`func (o *NotaryListElement) GetFkiProvinceIDOk() (*int32, bool)`

GetFkiProvinceIDOk returns a tuple with the FkiProvinceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiProvinceID

`func (o *NotaryListElement) SetFkiProvinceID(v int32)`

SetFkiProvinceID sets FkiProvinceID field to given value.

### HasFkiProvinceID

`func (o *NotaryListElement) HasFkiProvinceID() bool`

HasFkiProvinceID returns a boolean if a field has been set.

### GetFkiCountryID

`func (o *NotaryListElement) GetFkiCountryID() int32`

GetFkiCountryID returns the FkiCountryID field if non-nil, zero value otherwise.

### GetFkiCountryIDOk

`func (o *NotaryListElement) GetFkiCountryIDOk() (*int32, bool)`

GetFkiCountryIDOk returns a tuple with the FkiCountryID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiCountryID

`func (o *NotaryListElement) SetFkiCountryID(v int32)`

SetFkiCountryID sets FkiCountryID field to given value.

### HasFkiCountryID

`func (o *NotaryListElement) HasFkiCountryID() bool`

HasFkiCountryID returns a boolean if a field has been set.

### GetSContactLastname

`func (o *NotaryListElement) GetSContactLastname() string`

GetSContactLastname returns the SContactLastname field if non-nil, zero value otherwise.

### GetSContactLastnameOk

`func (o *NotaryListElement) GetSContactLastnameOk() (*string, bool)`

GetSContactLastnameOk returns a tuple with the SContactLastname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSContactLastname

`func (o *NotaryListElement) SetSContactLastname(v string)`

SetSContactLastname sets SContactLastname field to given value.

### HasSContactLastname

`func (o *NotaryListElement) HasSContactLastname() bool`

HasSContactLastname returns a boolean if a field has been set.

### GetSContactFirstname

`func (o *NotaryListElement) GetSContactFirstname() string`

GetSContactFirstname returns the SContactFirstname field if non-nil, zero value otherwise.

### GetSContactFirstnameOk

`func (o *NotaryListElement) GetSContactFirstnameOk() (*string, bool)`

GetSContactFirstnameOk returns a tuple with the SContactFirstname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSContactFirstname

`func (o *NotaryListElement) SetSContactFirstname(v string)`

SetSContactFirstname sets SContactFirstname field to given value.

### HasSContactFirstname

`func (o *NotaryListElement) HasSContactFirstname() bool`

HasSContactFirstname returns a boolean if a field has been set.

### GetSContactCompany

`func (o *NotaryListElement) GetSContactCompany() string`

GetSContactCompany returns the SContactCompany field if non-nil, zero value otherwise.

### GetSContactCompanyOk

`func (o *NotaryListElement) GetSContactCompanyOk() (*string, bool)`

GetSContactCompanyOk returns a tuple with the SContactCompany field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSContactCompany

`func (o *NotaryListElement) SetSContactCompany(v string)`

SetSContactCompany sets SContactCompany field to given value.

### HasSContactCompany

`func (o *NotaryListElement) HasSContactCompany() bool`

HasSContactCompany returns a boolean if a field has been set.

### GetSAddressCivic

`func (o *NotaryListElement) GetSAddressCivic() string`

GetSAddressCivic returns the SAddressCivic field if non-nil, zero value otherwise.

### GetSAddressCivicOk

`func (o *NotaryListElement) GetSAddressCivicOk() (*string, bool)`

GetSAddressCivicOk returns a tuple with the SAddressCivic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAddressCivic

`func (o *NotaryListElement) SetSAddressCivic(v string)`

SetSAddressCivic sets SAddressCivic field to given value.

### HasSAddressCivic

`func (o *NotaryListElement) HasSAddressCivic() bool`

HasSAddressCivic returns a boolean if a field has been set.

### GetSAddressStreet

`func (o *NotaryListElement) GetSAddressStreet() string`

GetSAddressStreet returns the SAddressStreet field if non-nil, zero value otherwise.

### GetSAddressStreetOk

`func (o *NotaryListElement) GetSAddressStreetOk() (*string, bool)`

GetSAddressStreetOk returns a tuple with the SAddressStreet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAddressStreet

`func (o *NotaryListElement) SetSAddressStreet(v string)`

SetSAddressStreet sets SAddressStreet field to given value.

### HasSAddressStreet

`func (o *NotaryListElement) HasSAddressStreet() bool`

HasSAddressStreet returns a boolean if a field has been set.

### GetSAddressSuite

`func (o *NotaryListElement) GetSAddressSuite() string`

GetSAddressSuite returns the SAddressSuite field if non-nil, zero value otherwise.

### GetSAddressSuiteOk

`func (o *NotaryListElement) GetSAddressSuiteOk() (*string, bool)`

GetSAddressSuiteOk returns a tuple with the SAddressSuite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAddressSuite

`func (o *NotaryListElement) SetSAddressSuite(v string)`

SetSAddressSuite sets SAddressSuite field to given value.

### HasSAddressSuite

`func (o *NotaryListElement) HasSAddressSuite() bool`

HasSAddressSuite returns a boolean if a field has been set.

### GetSAddressCity

`func (o *NotaryListElement) GetSAddressCity() string`

GetSAddressCity returns the SAddressCity field if non-nil, zero value otherwise.

### GetSAddressCityOk

`func (o *NotaryListElement) GetSAddressCityOk() (*string, bool)`

GetSAddressCityOk returns a tuple with the SAddressCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAddressCity

`func (o *NotaryListElement) SetSAddressCity(v string)`

SetSAddressCity sets SAddressCity field to given value.

### HasSAddressCity

`func (o *NotaryListElement) HasSAddressCity() bool`

HasSAddressCity returns a boolean if a field has been set.

### GetSPhoneE164

`func (o *NotaryListElement) GetSPhoneE164() string`

GetSPhoneE164 returns the SPhoneE164 field if non-nil, zero value otherwise.

### GetSPhoneE164Ok

`func (o *NotaryListElement) GetSPhoneE164Ok() (*string, bool)`

GetSPhoneE164Ok returns a tuple with the SPhoneE164 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSPhoneE164

`func (o *NotaryListElement) SetSPhoneE164(v string)`

SetSPhoneE164 sets SPhoneE164 field to given value.

### HasSPhoneE164

`func (o *NotaryListElement) HasSPhoneE164() bool`

HasSPhoneE164 returns a boolean if a field has been set.

### GetSPaymentmethodDescriptionX

`func (o *NotaryListElement) GetSPaymentmethodDescriptionX() string`

GetSPaymentmethodDescriptionX returns the SPaymentmethodDescriptionX field if non-nil, zero value otherwise.

### GetSPaymentmethodDescriptionXOk

`func (o *NotaryListElement) GetSPaymentmethodDescriptionXOk() (*string, bool)`

GetSPaymentmethodDescriptionXOk returns a tuple with the SPaymentmethodDescriptionX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSPaymentmethodDescriptionX

`func (o *NotaryListElement) SetSPaymentmethodDescriptionX(v string)`

SetSPaymentmethodDescriptionX sets SPaymentmethodDescriptionX field to given value.

### HasSPaymentmethodDescriptionX

`func (o *NotaryListElement) HasSPaymentmethodDescriptionX() bool`

HasSPaymentmethodDescriptionX returns a boolean if a field has been set.

### GetSProvinceNameX

`func (o *NotaryListElement) GetSProvinceNameX() string`

GetSProvinceNameX returns the SProvinceNameX field if non-nil, zero value otherwise.

### GetSProvinceNameXOk

`func (o *NotaryListElement) GetSProvinceNameXOk() (*string, bool)`

GetSProvinceNameXOk returns a tuple with the SProvinceNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSProvinceNameX

`func (o *NotaryListElement) SetSProvinceNameX(v string)`

SetSProvinceNameX sets SProvinceNameX field to given value.

### HasSProvinceNameX

`func (o *NotaryListElement) HasSProvinceNameX() bool`

HasSProvinceNameX returns a boolean if a field has been set.

### GetSCountryNameX

`func (o *NotaryListElement) GetSCountryNameX() string`

GetSCountryNameX returns the SCountryNameX field if non-nil, zero value otherwise.

### GetSCountryNameXOk

`func (o *NotaryListElement) GetSCountryNameXOk() (*string, bool)`

GetSCountryNameXOk returns a tuple with the SCountryNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSCountryNameX

`func (o *NotaryListElement) SetSCountryNameX(v string)`

SetSCountryNameX sets SCountryNameX field to given value.

### HasSCountryNameX

`func (o *NotaryListElement) HasSCountryNameX() bool`

HasSCountryNameX returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


