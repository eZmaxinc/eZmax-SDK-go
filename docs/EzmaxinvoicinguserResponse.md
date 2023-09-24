# EzmaxinvoicinguserResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzmaxinvoicinguserID** | Pointer to **int32** | The unique ID of the Ezmaxinvoicinguser | [optional] 
**FkiEzmaxinvoicingID** | Pointer to **int32** | The unique ID of the Ezmaxinvoicing | [optional] 
**FkiBillingentityinternalID** | **int32** | The unique ID of the Billingentityinternal. | 
**SBillingentityinternalDescriptionX** | **string** | The description of the Billingentityinternal in the language of the requester | 
**FkiUserID** | **int32** | The unique ID of the User | 
**IEzmaxinvoicinguserEzsigndocument** | **int32** | The number of ezsign documents | 
**BEzmaxinvoicinguserEzsignaccount** | **bool** | Whether there is an eZsign account | 
**BEzmaxinvoicinguserBillableezsign** | **bool** | Whether it is billable for eZsign | 
**EEzmaxinvoicinguserVariationezsign** | [**FieldEEzmaxinvoicinguserVariationezsign**](FieldEEzmaxinvoicinguserVariationezsign.md) |  | 

## Methods

### NewEzmaxinvoicinguserResponse

`func NewEzmaxinvoicinguserResponse(fkiBillingentityinternalID int32, sBillingentityinternalDescriptionX string, fkiUserID int32, iEzmaxinvoicinguserEzsigndocument int32, bEzmaxinvoicinguserEzsignaccount bool, bEzmaxinvoicinguserBillableezsign bool, eEzmaxinvoicinguserVariationezsign FieldEEzmaxinvoicinguserVariationezsign, ) *EzmaxinvoicinguserResponse`

NewEzmaxinvoicinguserResponse instantiates a new EzmaxinvoicinguserResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzmaxinvoicinguserResponseWithDefaults

`func NewEzmaxinvoicinguserResponseWithDefaults() *EzmaxinvoicinguserResponse`

NewEzmaxinvoicinguserResponseWithDefaults instantiates a new EzmaxinvoicinguserResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzmaxinvoicinguserID

`func (o *EzmaxinvoicinguserResponse) GetPkiEzmaxinvoicinguserID() int32`

GetPkiEzmaxinvoicinguserID returns the PkiEzmaxinvoicinguserID field if non-nil, zero value otherwise.

### GetPkiEzmaxinvoicinguserIDOk

`func (o *EzmaxinvoicinguserResponse) GetPkiEzmaxinvoicinguserIDOk() (*int32, bool)`

GetPkiEzmaxinvoicinguserIDOk returns a tuple with the PkiEzmaxinvoicinguserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzmaxinvoicinguserID

`func (o *EzmaxinvoicinguserResponse) SetPkiEzmaxinvoicinguserID(v int32)`

SetPkiEzmaxinvoicinguserID sets PkiEzmaxinvoicinguserID field to given value.

### HasPkiEzmaxinvoicinguserID

`func (o *EzmaxinvoicinguserResponse) HasPkiEzmaxinvoicinguserID() bool`

HasPkiEzmaxinvoicinguserID returns a boolean if a field has been set.

### GetFkiEzmaxinvoicingID

`func (o *EzmaxinvoicinguserResponse) GetFkiEzmaxinvoicingID() int32`

GetFkiEzmaxinvoicingID returns the FkiEzmaxinvoicingID field if non-nil, zero value otherwise.

### GetFkiEzmaxinvoicingIDOk

`func (o *EzmaxinvoicinguserResponse) GetFkiEzmaxinvoicingIDOk() (*int32, bool)`

GetFkiEzmaxinvoicingIDOk returns a tuple with the FkiEzmaxinvoicingID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzmaxinvoicingID

`func (o *EzmaxinvoicinguserResponse) SetFkiEzmaxinvoicingID(v int32)`

SetFkiEzmaxinvoicingID sets FkiEzmaxinvoicingID field to given value.

### HasFkiEzmaxinvoicingID

`func (o *EzmaxinvoicinguserResponse) HasFkiEzmaxinvoicingID() bool`

HasFkiEzmaxinvoicingID returns a boolean if a field has been set.

### GetFkiBillingentityinternalID

`func (o *EzmaxinvoicinguserResponse) GetFkiBillingentityinternalID() int32`

GetFkiBillingentityinternalID returns the FkiBillingentityinternalID field if non-nil, zero value otherwise.

### GetFkiBillingentityinternalIDOk

`func (o *EzmaxinvoicinguserResponse) GetFkiBillingentityinternalIDOk() (*int32, bool)`

GetFkiBillingentityinternalIDOk returns a tuple with the FkiBillingentityinternalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiBillingentityinternalID

`func (o *EzmaxinvoicinguserResponse) SetFkiBillingentityinternalID(v int32)`

SetFkiBillingentityinternalID sets FkiBillingentityinternalID field to given value.


### GetSBillingentityinternalDescriptionX

`func (o *EzmaxinvoicinguserResponse) GetSBillingentityinternalDescriptionX() string`

GetSBillingentityinternalDescriptionX returns the SBillingentityinternalDescriptionX field if non-nil, zero value otherwise.

### GetSBillingentityinternalDescriptionXOk

`func (o *EzmaxinvoicinguserResponse) GetSBillingentityinternalDescriptionXOk() (*string, bool)`

GetSBillingentityinternalDescriptionXOk returns a tuple with the SBillingentityinternalDescriptionX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSBillingentityinternalDescriptionX

`func (o *EzmaxinvoicinguserResponse) SetSBillingentityinternalDescriptionX(v string)`

SetSBillingentityinternalDescriptionX sets SBillingentityinternalDescriptionX field to given value.


### GetFkiUserID

`func (o *EzmaxinvoicinguserResponse) GetFkiUserID() int32`

GetFkiUserID returns the FkiUserID field if non-nil, zero value otherwise.

### GetFkiUserIDOk

`func (o *EzmaxinvoicinguserResponse) GetFkiUserIDOk() (*int32, bool)`

GetFkiUserIDOk returns a tuple with the FkiUserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUserID

`func (o *EzmaxinvoicinguserResponse) SetFkiUserID(v int32)`

SetFkiUserID sets FkiUserID field to given value.


### GetIEzmaxinvoicinguserEzsigndocument

`func (o *EzmaxinvoicinguserResponse) GetIEzmaxinvoicinguserEzsigndocument() int32`

GetIEzmaxinvoicinguserEzsigndocument returns the IEzmaxinvoicinguserEzsigndocument field if non-nil, zero value otherwise.

### GetIEzmaxinvoicinguserEzsigndocumentOk

`func (o *EzmaxinvoicinguserResponse) GetIEzmaxinvoicinguserEzsigndocumentOk() (*int32, bool)`

GetIEzmaxinvoicinguserEzsigndocumentOk returns a tuple with the IEzmaxinvoicinguserEzsigndocument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzmaxinvoicinguserEzsigndocument

`func (o *EzmaxinvoicinguserResponse) SetIEzmaxinvoicinguserEzsigndocument(v int32)`

SetIEzmaxinvoicinguserEzsigndocument sets IEzmaxinvoicinguserEzsigndocument field to given value.


### GetBEzmaxinvoicinguserEzsignaccount

`func (o *EzmaxinvoicinguserResponse) GetBEzmaxinvoicinguserEzsignaccount() bool`

GetBEzmaxinvoicinguserEzsignaccount returns the BEzmaxinvoicinguserEzsignaccount field if non-nil, zero value otherwise.

### GetBEzmaxinvoicinguserEzsignaccountOk

`func (o *EzmaxinvoicinguserResponse) GetBEzmaxinvoicinguserEzsignaccountOk() (*bool, bool)`

GetBEzmaxinvoicinguserEzsignaccountOk returns a tuple with the BEzmaxinvoicinguserEzsignaccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzmaxinvoicinguserEzsignaccount

`func (o *EzmaxinvoicinguserResponse) SetBEzmaxinvoicinguserEzsignaccount(v bool)`

SetBEzmaxinvoicinguserEzsignaccount sets BEzmaxinvoicinguserEzsignaccount field to given value.


### GetBEzmaxinvoicinguserBillableezsign

`func (o *EzmaxinvoicinguserResponse) GetBEzmaxinvoicinguserBillableezsign() bool`

GetBEzmaxinvoicinguserBillableezsign returns the BEzmaxinvoicinguserBillableezsign field if non-nil, zero value otherwise.

### GetBEzmaxinvoicinguserBillableezsignOk

`func (o *EzmaxinvoicinguserResponse) GetBEzmaxinvoicinguserBillableezsignOk() (*bool, bool)`

GetBEzmaxinvoicinguserBillableezsignOk returns a tuple with the BEzmaxinvoicinguserBillableezsign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzmaxinvoicinguserBillableezsign

`func (o *EzmaxinvoicinguserResponse) SetBEzmaxinvoicinguserBillableezsign(v bool)`

SetBEzmaxinvoicinguserBillableezsign sets BEzmaxinvoicinguserBillableezsign field to given value.


### GetEEzmaxinvoicinguserVariationezsign

`func (o *EzmaxinvoicinguserResponse) GetEEzmaxinvoicinguserVariationezsign() FieldEEzmaxinvoicinguserVariationezsign`

GetEEzmaxinvoicinguserVariationezsign returns the EEzmaxinvoicinguserVariationezsign field if non-nil, zero value otherwise.

### GetEEzmaxinvoicinguserVariationezsignOk

`func (o *EzmaxinvoicinguserResponse) GetEEzmaxinvoicinguserVariationezsignOk() (*FieldEEzmaxinvoicinguserVariationezsign, bool)`

GetEEzmaxinvoicinguserVariationezsignOk returns a tuple with the EEzmaxinvoicinguserVariationezsign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzmaxinvoicinguserVariationezsign

`func (o *EzmaxinvoicinguserResponse) SetEEzmaxinvoicinguserVariationezsign(v FieldEEzmaxinvoicinguserVariationezsign)`

SetEEzmaxinvoicinguserVariationezsign sets EEzmaxinvoicinguserVariationezsign field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


