# RejectedoffertopurchaseListElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiRejectedoffertopurchaseID** | **int32** | The unique ID of the Rejectedoffertopurchase | 
**SRejectedoffertopurchaseNumber** | **string** | The number of the Rejectedoffertopurchase | 
**DtRejectedoffertopurchaseDate** | **string** | The date of the Rejectedoffertopurchase | 
**BRejectedoffertopurchaseIsactive** | **bool** | Whether the rejectedoffertopurchase is active or not | 
**DtCreatedDate** | **string** | The date and time at which the object was created | 
**SAddressCivic** | Pointer to **string** | The Civic number. | [optional] 
**SAddressStreet** | Pointer to **string** | The Street Name | [optional] 
**SAddressSuite** | Pointer to **string** | The Suite or appartment number | [optional] 
**SAddressCity** | Pointer to **string** | The City name | [optional] 
**SAddressZip** | Pointer to **string** | The Postal/Zip Code  The value must be entered without spaces | [optional] 
**FkiProvinceID** | Pointer to **int32** | The unique ID of the Province.  Here are some common values (Complete list must be retrieved from API):  |Value|Description| |-|-| |1|(Canada) Alberta |2|(Canada) British Columbia| |3|(Canada) Manitoba| |3|(Canada) Manitoba| |4|(Canada) New Brunswick| |5|(Canada) Newfoundland| |6|(Canada) Northwest Territories| |7|(Canada) Nova Scotia| |8|(Canada) Nunavut| |9|(Canada) Ontario| |10|(Canada) Prince Edward Island| |11|(Canada) Quebec| |12|(Canada) Saskatchewan| |13|(Canada) Yukon| |14|(United-States) Alabama| |15|(United-States) Alaska| |16|(United-States) Arizona| |17|(United-States) Arkansas| |18|(United-States) California| |19|(United-States) Colorado| |20|(United-States) Connecticut| |21|(United-States) Delaware| |22|(United-States) District of Columbia| |23|(United-States) Florida| |24|(United-States) Georgia| |25|(United-States) Hawaii| |26|(United-States) Idaho| |27|(United-States) Illinois| |28|(United-States) Indiana| |29|(United-States) Iowa| |30|(United-States) Kansas| |31|(United-States) Kentucky| |32|(United-States) Louisiane| |33|(United-States) Maine| |34|(United-States) Maryland| |35|(United-States) Massachusetts| |36|(United-States) Michigan| |37|(United-States) Minnesota| |38|(United-States) Mississippi| |39|(United-States) Missouri| |40|(United-States) Montana| |41|(United-States) Nebraska| |42|(United-States) Nevada| |43|(United-States) New Hampshire| |44|(United-States) New Jersey| |45|(United-States) New Mexico| |46|(United-States) New York| |47|(United-States) North Carolina| |48|(United-States) North Dakota| |49|(United-States) Ohio| |50|(United-States) Oklahoma| |51|(United-States) Oregon| |52|(United-States) Pennsylvania| |53|(United-States) Rhode Island| |54|(United-States) South Carolina| |55|(United-States) South Dakota| |56|(United-States) Tennessee| |57|(United-States) Texas| |58|(United-States) Utah| |60|(United-States) Vermont| |59|(United-States) Virginia| |61|(United-States) Washington| |62|(United-States) West Virginia| |63|(United-States) Wisconsin| |64|(United-States) Wyoming| | [optional] 
**SProvinceNameX** | Pointer to **string** | The name of the Province in the language of the requester | [optional] 
**FkiCountryID** | Pointer to **int32** | The unique ID of the Country.  Here are some common values (Complete list must be retrieved from API):  |Value|Description| |-|-| |1|Canada| |2|United-States| | [optional] 
**SCountryNameX** | Pointer to **string** | The name of the Country in the language of the requester | [optional] 
**BRejectedoffertopurchaseLinkedtoinscription** | **bool** | Indicate if the Rejectedoffertopurchase is linked to an inscription | 

## Methods

### NewRejectedoffertopurchaseListElement

`func NewRejectedoffertopurchaseListElement(pkiRejectedoffertopurchaseID int32, sRejectedoffertopurchaseNumber string, dtRejectedoffertopurchaseDate string, bRejectedoffertopurchaseIsactive bool, dtCreatedDate string, bRejectedoffertopurchaseLinkedtoinscription bool, ) *RejectedoffertopurchaseListElement`

NewRejectedoffertopurchaseListElement instantiates a new RejectedoffertopurchaseListElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRejectedoffertopurchaseListElementWithDefaults

`func NewRejectedoffertopurchaseListElementWithDefaults() *RejectedoffertopurchaseListElement`

NewRejectedoffertopurchaseListElementWithDefaults instantiates a new RejectedoffertopurchaseListElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiRejectedoffertopurchaseID

`func (o *RejectedoffertopurchaseListElement) GetPkiRejectedoffertopurchaseID() int32`

GetPkiRejectedoffertopurchaseID returns the PkiRejectedoffertopurchaseID field if non-nil, zero value otherwise.

### GetPkiRejectedoffertopurchaseIDOk

`func (o *RejectedoffertopurchaseListElement) GetPkiRejectedoffertopurchaseIDOk() (*int32, bool)`

GetPkiRejectedoffertopurchaseIDOk returns a tuple with the PkiRejectedoffertopurchaseID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiRejectedoffertopurchaseID

`func (o *RejectedoffertopurchaseListElement) SetPkiRejectedoffertopurchaseID(v int32)`

SetPkiRejectedoffertopurchaseID sets PkiRejectedoffertopurchaseID field to given value.


### GetSRejectedoffertopurchaseNumber

`func (o *RejectedoffertopurchaseListElement) GetSRejectedoffertopurchaseNumber() string`

GetSRejectedoffertopurchaseNumber returns the SRejectedoffertopurchaseNumber field if non-nil, zero value otherwise.

### GetSRejectedoffertopurchaseNumberOk

`func (o *RejectedoffertopurchaseListElement) GetSRejectedoffertopurchaseNumberOk() (*string, bool)`

GetSRejectedoffertopurchaseNumberOk returns a tuple with the SRejectedoffertopurchaseNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSRejectedoffertopurchaseNumber

`func (o *RejectedoffertopurchaseListElement) SetSRejectedoffertopurchaseNumber(v string)`

SetSRejectedoffertopurchaseNumber sets SRejectedoffertopurchaseNumber field to given value.


### GetDtRejectedoffertopurchaseDate

`func (o *RejectedoffertopurchaseListElement) GetDtRejectedoffertopurchaseDate() string`

GetDtRejectedoffertopurchaseDate returns the DtRejectedoffertopurchaseDate field if non-nil, zero value otherwise.

### GetDtRejectedoffertopurchaseDateOk

`func (o *RejectedoffertopurchaseListElement) GetDtRejectedoffertopurchaseDateOk() (*string, bool)`

GetDtRejectedoffertopurchaseDateOk returns a tuple with the DtRejectedoffertopurchaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtRejectedoffertopurchaseDate

`func (o *RejectedoffertopurchaseListElement) SetDtRejectedoffertopurchaseDate(v string)`

SetDtRejectedoffertopurchaseDate sets DtRejectedoffertopurchaseDate field to given value.


### GetBRejectedoffertopurchaseIsactive

`func (o *RejectedoffertopurchaseListElement) GetBRejectedoffertopurchaseIsactive() bool`

GetBRejectedoffertopurchaseIsactive returns the BRejectedoffertopurchaseIsactive field if non-nil, zero value otherwise.

### GetBRejectedoffertopurchaseIsactiveOk

`func (o *RejectedoffertopurchaseListElement) GetBRejectedoffertopurchaseIsactiveOk() (*bool, bool)`

GetBRejectedoffertopurchaseIsactiveOk returns a tuple with the BRejectedoffertopurchaseIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBRejectedoffertopurchaseIsactive

`func (o *RejectedoffertopurchaseListElement) SetBRejectedoffertopurchaseIsactive(v bool)`

SetBRejectedoffertopurchaseIsactive sets BRejectedoffertopurchaseIsactive field to given value.


### GetDtCreatedDate

`func (o *RejectedoffertopurchaseListElement) GetDtCreatedDate() string`

GetDtCreatedDate returns the DtCreatedDate field if non-nil, zero value otherwise.

### GetDtCreatedDateOk

`func (o *RejectedoffertopurchaseListElement) GetDtCreatedDateOk() (*string, bool)`

GetDtCreatedDateOk returns a tuple with the DtCreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtCreatedDate

`func (o *RejectedoffertopurchaseListElement) SetDtCreatedDate(v string)`

SetDtCreatedDate sets DtCreatedDate field to given value.


### GetSAddressCivic

`func (o *RejectedoffertopurchaseListElement) GetSAddressCivic() string`

GetSAddressCivic returns the SAddressCivic field if non-nil, zero value otherwise.

### GetSAddressCivicOk

`func (o *RejectedoffertopurchaseListElement) GetSAddressCivicOk() (*string, bool)`

GetSAddressCivicOk returns a tuple with the SAddressCivic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAddressCivic

`func (o *RejectedoffertopurchaseListElement) SetSAddressCivic(v string)`

SetSAddressCivic sets SAddressCivic field to given value.

### HasSAddressCivic

`func (o *RejectedoffertopurchaseListElement) HasSAddressCivic() bool`

HasSAddressCivic returns a boolean if a field has been set.

### GetSAddressStreet

`func (o *RejectedoffertopurchaseListElement) GetSAddressStreet() string`

GetSAddressStreet returns the SAddressStreet field if non-nil, zero value otherwise.

### GetSAddressStreetOk

`func (o *RejectedoffertopurchaseListElement) GetSAddressStreetOk() (*string, bool)`

GetSAddressStreetOk returns a tuple with the SAddressStreet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAddressStreet

`func (o *RejectedoffertopurchaseListElement) SetSAddressStreet(v string)`

SetSAddressStreet sets SAddressStreet field to given value.

### HasSAddressStreet

`func (o *RejectedoffertopurchaseListElement) HasSAddressStreet() bool`

HasSAddressStreet returns a boolean if a field has been set.

### GetSAddressSuite

`func (o *RejectedoffertopurchaseListElement) GetSAddressSuite() string`

GetSAddressSuite returns the SAddressSuite field if non-nil, zero value otherwise.

### GetSAddressSuiteOk

`func (o *RejectedoffertopurchaseListElement) GetSAddressSuiteOk() (*string, bool)`

GetSAddressSuiteOk returns a tuple with the SAddressSuite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAddressSuite

`func (o *RejectedoffertopurchaseListElement) SetSAddressSuite(v string)`

SetSAddressSuite sets SAddressSuite field to given value.

### HasSAddressSuite

`func (o *RejectedoffertopurchaseListElement) HasSAddressSuite() bool`

HasSAddressSuite returns a boolean if a field has been set.

### GetSAddressCity

`func (o *RejectedoffertopurchaseListElement) GetSAddressCity() string`

GetSAddressCity returns the SAddressCity field if non-nil, zero value otherwise.

### GetSAddressCityOk

`func (o *RejectedoffertopurchaseListElement) GetSAddressCityOk() (*string, bool)`

GetSAddressCityOk returns a tuple with the SAddressCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAddressCity

`func (o *RejectedoffertopurchaseListElement) SetSAddressCity(v string)`

SetSAddressCity sets SAddressCity field to given value.

### HasSAddressCity

`func (o *RejectedoffertopurchaseListElement) HasSAddressCity() bool`

HasSAddressCity returns a boolean if a field has been set.

### GetSAddressZip

`func (o *RejectedoffertopurchaseListElement) GetSAddressZip() string`

GetSAddressZip returns the SAddressZip field if non-nil, zero value otherwise.

### GetSAddressZipOk

`func (o *RejectedoffertopurchaseListElement) GetSAddressZipOk() (*string, bool)`

GetSAddressZipOk returns a tuple with the SAddressZip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAddressZip

`func (o *RejectedoffertopurchaseListElement) SetSAddressZip(v string)`

SetSAddressZip sets SAddressZip field to given value.

### HasSAddressZip

`func (o *RejectedoffertopurchaseListElement) HasSAddressZip() bool`

HasSAddressZip returns a boolean if a field has been set.

### GetFkiProvinceID

`func (o *RejectedoffertopurchaseListElement) GetFkiProvinceID() int32`

GetFkiProvinceID returns the FkiProvinceID field if non-nil, zero value otherwise.

### GetFkiProvinceIDOk

`func (o *RejectedoffertopurchaseListElement) GetFkiProvinceIDOk() (*int32, bool)`

GetFkiProvinceIDOk returns a tuple with the FkiProvinceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiProvinceID

`func (o *RejectedoffertopurchaseListElement) SetFkiProvinceID(v int32)`

SetFkiProvinceID sets FkiProvinceID field to given value.

### HasFkiProvinceID

`func (o *RejectedoffertopurchaseListElement) HasFkiProvinceID() bool`

HasFkiProvinceID returns a boolean if a field has been set.

### GetSProvinceNameX

`func (o *RejectedoffertopurchaseListElement) GetSProvinceNameX() string`

GetSProvinceNameX returns the SProvinceNameX field if non-nil, zero value otherwise.

### GetSProvinceNameXOk

`func (o *RejectedoffertopurchaseListElement) GetSProvinceNameXOk() (*string, bool)`

GetSProvinceNameXOk returns a tuple with the SProvinceNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSProvinceNameX

`func (o *RejectedoffertopurchaseListElement) SetSProvinceNameX(v string)`

SetSProvinceNameX sets SProvinceNameX field to given value.

### HasSProvinceNameX

`func (o *RejectedoffertopurchaseListElement) HasSProvinceNameX() bool`

HasSProvinceNameX returns a boolean if a field has been set.

### GetFkiCountryID

`func (o *RejectedoffertopurchaseListElement) GetFkiCountryID() int32`

GetFkiCountryID returns the FkiCountryID field if non-nil, zero value otherwise.

### GetFkiCountryIDOk

`func (o *RejectedoffertopurchaseListElement) GetFkiCountryIDOk() (*int32, bool)`

GetFkiCountryIDOk returns a tuple with the FkiCountryID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiCountryID

`func (o *RejectedoffertopurchaseListElement) SetFkiCountryID(v int32)`

SetFkiCountryID sets FkiCountryID field to given value.

### HasFkiCountryID

`func (o *RejectedoffertopurchaseListElement) HasFkiCountryID() bool`

HasFkiCountryID returns a boolean if a field has been set.

### GetSCountryNameX

`func (o *RejectedoffertopurchaseListElement) GetSCountryNameX() string`

GetSCountryNameX returns the SCountryNameX field if non-nil, zero value otherwise.

### GetSCountryNameXOk

`func (o *RejectedoffertopurchaseListElement) GetSCountryNameXOk() (*string, bool)`

GetSCountryNameXOk returns a tuple with the SCountryNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSCountryNameX

`func (o *RejectedoffertopurchaseListElement) SetSCountryNameX(v string)`

SetSCountryNameX sets SCountryNameX field to given value.

### HasSCountryNameX

`func (o *RejectedoffertopurchaseListElement) HasSCountryNameX() bool`

HasSCountryNameX returns a boolean if a field has been set.

### GetBRejectedoffertopurchaseLinkedtoinscription

`func (o *RejectedoffertopurchaseListElement) GetBRejectedoffertopurchaseLinkedtoinscription() bool`

GetBRejectedoffertopurchaseLinkedtoinscription returns the BRejectedoffertopurchaseLinkedtoinscription field if non-nil, zero value otherwise.

### GetBRejectedoffertopurchaseLinkedtoinscriptionOk

`func (o *RejectedoffertopurchaseListElement) GetBRejectedoffertopurchaseLinkedtoinscriptionOk() (*bool, bool)`

GetBRejectedoffertopurchaseLinkedtoinscriptionOk returns a tuple with the BRejectedoffertopurchaseLinkedtoinscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBRejectedoffertopurchaseLinkedtoinscription

`func (o *RejectedoffertopurchaseListElement) SetBRejectedoffertopurchaseLinkedtoinscription(v bool)`

SetBRejectedoffertopurchaseLinkedtoinscription sets BRejectedoffertopurchaseLinkedtoinscription field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


