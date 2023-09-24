# EzmaxinvoicingsummaryglobalResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzmaxinvoicingsummaryglobalID** | Pointer to **int32** | The unique ID of the Ezmaxinvoicingsummaryglobal | [optional] 
**FkiEzmaxinvoicingID** | Pointer to **int32** | The unique ID of the Ezmaxinvoicing | [optional] 
**FkiEzmaxproductID** | **int32** | The unique ID of the Ezmaxproduct | 
**SEzmaxproductDescriptionX** | **string** | The description of the Ezmaxproduct in the language of the requester | 
**DtEzmaxinvoicingsummaryglobalStart** | **string** | The start date for the Ezmaxinvoicingsummaryglobal | 
**DtEzmaxinvoicingsummaryglobalEnd** | **string** | The end date for the Ezmaxinvoicingsummaryglobal | 
**IEzmaxinvoicingsummaryglobalDays** | **int32** | The number of days for the Ezmaxinvoicingsummaryglobal | 
**DEzmaxinvoicingsummaryglobalCountreal** | **string** | The count item calculated | 
**DEzmaxinvoicingsummaryglobalCountbilled** | **string** | The count item billed | 
**DEzmaxinvoicingsummaryglobalSubtotal** | **string** | The Ezmaxinvoicingsummaryglobal subtotal | 
**DEzmaxinvoicingsummaryglobalRebateamount** | **string** | The rebate amount for the Ezmaxinvoicingsummaryglobal | 
**DEzmaxinvoicingsummaryglobalRebatepercent** | **string** | The rebate percentage of the Ezmaxinvoicingsummaryglobal | 
**DEzmaxinvoicingsummaryglobalRebatetotal** | **string** | The rebate amount total for the Ezmaxinvoicingsummaryglobal | 
**DEzmaxinvoicingsummaryglobalTotal** | **string** | The Ezmaxinvoicingsummaryglobal total | 
**DEzmaxinvoicingsummaryglobalRepresentative** | Pointer to **string** | The amount of commission for the representative | [optional] 
**DEzmaxinvoicingsummaryglobalPartner** | Pointer to **string** | The amount of commission for the partner | [optional] 
**DEzmaxinvoicingsummaryglobalNet** | Pointer to **string** | The net amount of the Ezmaxinvoicingsummaryglobal | [optional] 
**BEzmaxinvoicingsummaryglobalAdjustment** | **bool** | Whether it is adjustment for the Ezmaxinvoicingsummaryglobal | 
**TEzmaxproductHelpX** | **string** | The help message of the Ezmaxproduct in the language of the requester | 

## Methods

### NewEzmaxinvoicingsummaryglobalResponse

`func NewEzmaxinvoicingsummaryglobalResponse(fkiEzmaxproductID int32, sEzmaxproductDescriptionX string, dtEzmaxinvoicingsummaryglobalStart string, dtEzmaxinvoicingsummaryglobalEnd string, iEzmaxinvoicingsummaryglobalDays int32, dEzmaxinvoicingsummaryglobalCountreal string, dEzmaxinvoicingsummaryglobalCountbilled string, dEzmaxinvoicingsummaryglobalSubtotal string, dEzmaxinvoicingsummaryglobalRebateamount string, dEzmaxinvoicingsummaryglobalRebatepercent string, dEzmaxinvoicingsummaryglobalRebatetotal string, dEzmaxinvoicingsummaryglobalTotal string, bEzmaxinvoicingsummaryglobalAdjustment bool, tEzmaxproductHelpX string, ) *EzmaxinvoicingsummaryglobalResponse`

NewEzmaxinvoicingsummaryglobalResponse instantiates a new EzmaxinvoicingsummaryglobalResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzmaxinvoicingsummaryglobalResponseWithDefaults

`func NewEzmaxinvoicingsummaryglobalResponseWithDefaults() *EzmaxinvoicingsummaryglobalResponse`

NewEzmaxinvoicingsummaryglobalResponseWithDefaults instantiates a new EzmaxinvoicingsummaryglobalResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzmaxinvoicingsummaryglobalID

`func (o *EzmaxinvoicingsummaryglobalResponse) GetPkiEzmaxinvoicingsummaryglobalID() int32`

GetPkiEzmaxinvoicingsummaryglobalID returns the PkiEzmaxinvoicingsummaryglobalID field if non-nil, zero value otherwise.

### GetPkiEzmaxinvoicingsummaryglobalIDOk

`func (o *EzmaxinvoicingsummaryglobalResponse) GetPkiEzmaxinvoicingsummaryglobalIDOk() (*int32, bool)`

GetPkiEzmaxinvoicingsummaryglobalIDOk returns a tuple with the PkiEzmaxinvoicingsummaryglobalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzmaxinvoicingsummaryglobalID

`func (o *EzmaxinvoicingsummaryglobalResponse) SetPkiEzmaxinvoicingsummaryglobalID(v int32)`

SetPkiEzmaxinvoicingsummaryglobalID sets PkiEzmaxinvoicingsummaryglobalID field to given value.

### HasPkiEzmaxinvoicingsummaryglobalID

`func (o *EzmaxinvoicingsummaryglobalResponse) HasPkiEzmaxinvoicingsummaryglobalID() bool`

HasPkiEzmaxinvoicingsummaryglobalID returns a boolean if a field has been set.

### GetFkiEzmaxinvoicingID

`func (o *EzmaxinvoicingsummaryglobalResponse) GetFkiEzmaxinvoicingID() int32`

GetFkiEzmaxinvoicingID returns the FkiEzmaxinvoicingID field if non-nil, zero value otherwise.

### GetFkiEzmaxinvoicingIDOk

`func (o *EzmaxinvoicingsummaryglobalResponse) GetFkiEzmaxinvoicingIDOk() (*int32, bool)`

GetFkiEzmaxinvoicingIDOk returns a tuple with the FkiEzmaxinvoicingID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzmaxinvoicingID

`func (o *EzmaxinvoicingsummaryglobalResponse) SetFkiEzmaxinvoicingID(v int32)`

SetFkiEzmaxinvoicingID sets FkiEzmaxinvoicingID field to given value.

### HasFkiEzmaxinvoicingID

`func (o *EzmaxinvoicingsummaryglobalResponse) HasFkiEzmaxinvoicingID() bool`

HasFkiEzmaxinvoicingID returns a boolean if a field has been set.

### GetFkiEzmaxproductID

`func (o *EzmaxinvoicingsummaryglobalResponse) GetFkiEzmaxproductID() int32`

GetFkiEzmaxproductID returns the FkiEzmaxproductID field if non-nil, zero value otherwise.

### GetFkiEzmaxproductIDOk

`func (o *EzmaxinvoicingsummaryglobalResponse) GetFkiEzmaxproductIDOk() (*int32, bool)`

GetFkiEzmaxproductIDOk returns a tuple with the FkiEzmaxproductID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzmaxproductID

`func (o *EzmaxinvoicingsummaryglobalResponse) SetFkiEzmaxproductID(v int32)`

SetFkiEzmaxproductID sets FkiEzmaxproductID field to given value.


### GetSEzmaxproductDescriptionX

`func (o *EzmaxinvoicingsummaryglobalResponse) GetSEzmaxproductDescriptionX() string`

GetSEzmaxproductDescriptionX returns the SEzmaxproductDescriptionX field if non-nil, zero value otherwise.

### GetSEzmaxproductDescriptionXOk

`func (o *EzmaxinvoicingsummaryglobalResponse) GetSEzmaxproductDescriptionXOk() (*string, bool)`

GetSEzmaxproductDescriptionXOk returns a tuple with the SEzmaxproductDescriptionX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzmaxproductDescriptionX

`func (o *EzmaxinvoicingsummaryglobalResponse) SetSEzmaxproductDescriptionX(v string)`

SetSEzmaxproductDescriptionX sets SEzmaxproductDescriptionX field to given value.


### GetDtEzmaxinvoicingsummaryglobalStart

`func (o *EzmaxinvoicingsummaryglobalResponse) GetDtEzmaxinvoicingsummaryglobalStart() string`

GetDtEzmaxinvoicingsummaryglobalStart returns the DtEzmaxinvoicingsummaryglobalStart field if non-nil, zero value otherwise.

### GetDtEzmaxinvoicingsummaryglobalStartOk

`func (o *EzmaxinvoicingsummaryglobalResponse) GetDtEzmaxinvoicingsummaryglobalStartOk() (*string, bool)`

GetDtEzmaxinvoicingsummaryglobalStartOk returns a tuple with the DtEzmaxinvoicingsummaryglobalStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtEzmaxinvoicingsummaryglobalStart

`func (o *EzmaxinvoicingsummaryglobalResponse) SetDtEzmaxinvoicingsummaryglobalStart(v string)`

SetDtEzmaxinvoicingsummaryglobalStart sets DtEzmaxinvoicingsummaryglobalStart field to given value.


### GetDtEzmaxinvoicingsummaryglobalEnd

`func (o *EzmaxinvoicingsummaryglobalResponse) GetDtEzmaxinvoicingsummaryglobalEnd() string`

GetDtEzmaxinvoicingsummaryglobalEnd returns the DtEzmaxinvoicingsummaryglobalEnd field if non-nil, zero value otherwise.

### GetDtEzmaxinvoicingsummaryglobalEndOk

`func (o *EzmaxinvoicingsummaryglobalResponse) GetDtEzmaxinvoicingsummaryglobalEndOk() (*string, bool)`

GetDtEzmaxinvoicingsummaryglobalEndOk returns a tuple with the DtEzmaxinvoicingsummaryglobalEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtEzmaxinvoicingsummaryglobalEnd

`func (o *EzmaxinvoicingsummaryglobalResponse) SetDtEzmaxinvoicingsummaryglobalEnd(v string)`

SetDtEzmaxinvoicingsummaryglobalEnd sets DtEzmaxinvoicingsummaryglobalEnd field to given value.


### GetIEzmaxinvoicingsummaryglobalDays

`func (o *EzmaxinvoicingsummaryglobalResponse) GetIEzmaxinvoicingsummaryglobalDays() int32`

GetIEzmaxinvoicingsummaryglobalDays returns the IEzmaxinvoicingsummaryglobalDays field if non-nil, zero value otherwise.

### GetIEzmaxinvoicingsummaryglobalDaysOk

`func (o *EzmaxinvoicingsummaryglobalResponse) GetIEzmaxinvoicingsummaryglobalDaysOk() (*int32, bool)`

GetIEzmaxinvoicingsummaryglobalDaysOk returns a tuple with the IEzmaxinvoicingsummaryglobalDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzmaxinvoicingsummaryglobalDays

`func (o *EzmaxinvoicingsummaryglobalResponse) SetIEzmaxinvoicingsummaryglobalDays(v int32)`

SetIEzmaxinvoicingsummaryglobalDays sets IEzmaxinvoicingsummaryglobalDays field to given value.


### GetDEzmaxinvoicingsummaryglobalCountreal

`func (o *EzmaxinvoicingsummaryglobalResponse) GetDEzmaxinvoicingsummaryglobalCountreal() string`

GetDEzmaxinvoicingsummaryglobalCountreal returns the DEzmaxinvoicingsummaryglobalCountreal field if non-nil, zero value otherwise.

### GetDEzmaxinvoicingsummaryglobalCountrealOk

`func (o *EzmaxinvoicingsummaryglobalResponse) GetDEzmaxinvoicingsummaryglobalCountrealOk() (*string, bool)`

GetDEzmaxinvoicingsummaryglobalCountrealOk returns a tuple with the DEzmaxinvoicingsummaryglobalCountreal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDEzmaxinvoicingsummaryglobalCountreal

`func (o *EzmaxinvoicingsummaryglobalResponse) SetDEzmaxinvoicingsummaryglobalCountreal(v string)`

SetDEzmaxinvoicingsummaryglobalCountreal sets DEzmaxinvoicingsummaryglobalCountreal field to given value.


### GetDEzmaxinvoicingsummaryglobalCountbilled

`func (o *EzmaxinvoicingsummaryglobalResponse) GetDEzmaxinvoicingsummaryglobalCountbilled() string`

GetDEzmaxinvoicingsummaryglobalCountbilled returns the DEzmaxinvoicingsummaryglobalCountbilled field if non-nil, zero value otherwise.

### GetDEzmaxinvoicingsummaryglobalCountbilledOk

`func (o *EzmaxinvoicingsummaryglobalResponse) GetDEzmaxinvoicingsummaryglobalCountbilledOk() (*string, bool)`

GetDEzmaxinvoicingsummaryglobalCountbilledOk returns a tuple with the DEzmaxinvoicingsummaryglobalCountbilled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDEzmaxinvoicingsummaryglobalCountbilled

`func (o *EzmaxinvoicingsummaryglobalResponse) SetDEzmaxinvoicingsummaryglobalCountbilled(v string)`

SetDEzmaxinvoicingsummaryglobalCountbilled sets DEzmaxinvoicingsummaryglobalCountbilled field to given value.


### GetDEzmaxinvoicingsummaryglobalSubtotal

`func (o *EzmaxinvoicingsummaryglobalResponse) GetDEzmaxinvoicingsummaryglobalSubtotal() string`

GetDEzmaxinvoicingsummaryglobalSubtotal returns the DEzmaxinvoicingsummaryglobalSubtotal field if non-nil, zero value otherwise.

### GetDEzmaxinvoicingsummaryglobalSubtotalOk

`func (o *EzmaxinvoicingsummaryglobalResponse) GetDEzmaxinvoicingsummaryglobalSubtotalOk() (*string, bool)`

GetDEzmaxinvoicingsummaryglobalSubtotalOk returns a tuple with the DEzmaxinvoicingsummaryglobalSubtotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDEzmaxinvoicingsummaryglobalSubtotal

`func (o *EzmaxinvoicingsummaryglobalResponse) SetDEzmaxinvoicingsummaryglobalSubtotal(v string)`

SetDEzmaxinvoicingsummaryglobalSubtotal sets DEzmaxinvoicingsummaryglobalSubtotal field to given value.


### GetDEzmaxinvoicingsummaryglobalRebateamount

`func (o *EzmaxinvoicingsummaryglobalResponse) GetDEzmaxinvoicingsummaryglobalRebateamount() string`

GetDEzmaxinvoicingsummaryglobalRebateamount returns the DEzmaxinvoicingsummaryglobalRebateamount field if non-nil, zero value otherwise.

### GetDEzmaxinvoicingsummaryglobalRebateamountOk

`func (o *EzmaxinvoicingsummaryglobalResponse) GetDEzmaxinvoicingsummaryglobalRebateamountOk() (*string, bool)`

GetDEzmaxinvoicingsummaryglobalRebateamountOk returns a tuple with the DEzmaxinvoicingsummaryglobalRebateamount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDEzmaxinvoicingsummaryglobalRebateamount

`func (o *EzmaxinvoicingsummaryglobalResponse) SetDEzmaxinvoicingsummaryglobalRebateamount(v string)`

SetDEzmaxinvoicingsummaryglobalRebateamount sets DEzmaxinvoicingsummaryglobalRebateamount field to given value.


### GetDEzmaxinvoicingsummaryglobalRebatepercent

`func (o *EzmaxinvoicingsummaryglobalResponse) GetDEzmaxinvoicingsummaryglobalRebatepercent() string`

GetDEzmaxinvoicingsummaryglobalRebatepercent returns the DEzmaxinvoicingsummaryglobalRebatepercent field if non-nil, zero value otherwise.

### GetDEzmaxinvoicingsummaryglobalRebatepercentOk

`func (o *EzmaxinvoicingsummaryglobalResponse) GetDEzmaxinvoicingsummaryglobalRebatepercentOk() (*string, bool)`

GetDEzmaxinvoicingsummaryglobalRebatepercentOk returns a tuple with the DEzmaxinvoicingsummaryglobalRebatepercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDEzmaxinvoicingsummaryglobalRebatepercent

`func (o *EzmaxinvoicingsummaryglobalResponse) SetDEzmaxinvoicingsummaryglobalRebatepercent(v string)`

SetDEzmaxinvoicingsummaryglobalRebatepercent sets DEzmaxinvoicingsummaryglobalRebatepercent field to given value.


### GetDEzmaxinvoicingsummaryglobalRebatetotal

`func (o *EzmaxinvoicingsummaryglobalResponse) GetDEzmaxinvoicingsummaryglobalRebatetotal() string`

GetDEzmaxinvoicingsummaryglobalRebatetotal returns the DEzmaxinvoicingsummaryglobalRebatetotal field if non-nil, zero value otherwise.

### GetDEzmaxinvoicingsummaryglobalRebatetotalOk

`func (o *EzmaxinvoicingsummaryglobalResponse) GetDEzmaxinvoicingsummaryglobalRebatetotalOk() (*string, bool)`

GetDEzmaxinvoicingsummaryglobalRebatetotalOk returns a tuple with the DEzmaxinvoicingsummaryglobalRebatetotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDEzmaxinvoicingsummaryglobalRebatetotal

`func (o *EzmaxinvoicingsummaryglobalResponse) SetDEzmaxinvoicingsummaryglobalRebatetotal(v string)`

SetDEzmaxinvoicingsummaryglobalRebatetotal sets DEzmaxinvoicingsummaryglobalRebatetotal field to given value.


### GetDEzmaxinvoicingsummaryglobalTotal

`func (o *EzmaxinvoicingsummaryglobalResponse) GetDEzmaxinvoicingsummaryglobalTotal() string`

GetDEzmaxinvoicingsummaryglobalTotal returns the DEzmaxinvoicingsummaryglobalTotal field if non-nil, zero value otherwise.

### GetDEzmaxinvoicingsummaryglobalTotalOk

`func (o *EzmaxinvoicingsummaryglobalResponse) GetDEzmaxinvoicingsummaryglobalTotalOk() (*string, bool)`

GetDEzmaxinvoicingsummaryglobalTotalOk returns a tuple with the DEzmaxinvoicingsummaryglobalTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDEzmaxinvoicingsummaryglobalTotal

`func (o *EzmaxinvoicingsummaryglobalResponse) SetDEzmaxinvoicingsummaryglobalTotal(v string)`

SetDEzmaxinvoicingsummaryglobalTotal sets DEzmaxinvoicingsummaryglobalTotal field to given value.


### GetDEzmaxinvoicingsummaryglobalRepresentative

`func (o *EzmaxinvoicingsummaryglobalResponse) GetDEzmaxinvoicingsummaryglobalRepresentative() string`

GetDEzmaxinvoicingsummaryglobalRepresentative returns the DEzmaxinvoicingsummaryglobalRepresentative field if non-nil, zero value otherwise.

### GetDEzmaxinvoicingsummaryglobalRepresentativeOk

`func (o *EzmaxinvoicingsummaryglobalResponse) GetDEzmaxinvoicingsummaryglobalRepresentativeOk() (*string, bool)`

GetDEzmaxinvoicingsummaryglobalRepresentativeOk returns a tuple with the DEzmaxinvoicingsummaryglobalRepresentative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDEzmaxinvoicingsummaryglobalRepresentative

`func (o *EzmaxinvoicingsummaryglobalResponse) SetDEzmaxinvoicingsummaryglobalRepresentative(v string)`

SetDEzmaxinvoicingsummaryglobalRepresentative sets DEzmaxinvoicingsummaryglobalRepresentative field to given value.

### HasDEzmaxinvoicingsummaryglobalRepresentative

`func (o *EzmaxinvoicingsummaryglobalResponse) HasDEzmaxinvoicingsummaryglobalRepresentative() bool`

HasDEzmaxinvoicingsummaryglobalRepresentative returns a boolean if a field has been set.

### GetDEzmaxinvoicingsummaryglobalPartner

`func (o *EzmaxinvoicingsummaryglobalResponse) GetDEzmaxinvoicingsummaryglobalPartner() string`

GetDEzmaxinvoicingsummaryglobalPartner returns the DEzmaxinvoicingsummaryglobalPartner field if non-nil, zero value otherwise.

### GetDEzmaxinvoicingsummaryglobalPartnerOk

`func (o *EzmaxinvoicingsummaryglobalResponse) GetDEzmaxinvoicingsummaryglobalPartnerOk() (*string, bool)`

GetDEzmaxinvoicingsummaryglobalPartnerOk returns a tuple with the DEzmaxinvoicingsummaryglobalPartner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDEzmaxinvoicingsummaryglobalPartner

`func (o *EzmaxinvoicingsummaryglobalResponse) SetDEzmaxinvoicingsummaryglobalPartner(v string)`

SetDEzmaxinvoicingsummaryglobalPartner sets DEzmaxinvoicingsummaryglobalPartner field to given value.

### HasDEzmaxinvoicingsummaryglobalPartner

`func (o *EzmaxinvoicingsummaryglobalResponse) HasDEzmaxinvoicingsummaryglobalPartner() bool`

HasDEzmaxinvoicingsummaryglobalPartner returns a boolean if a field has been set.

### GetDEzmaxinvoicingsummaryglobalNet

`func (o *EzmaxinvoicingsummaryglobalResponse) GetDEzmaxinvoicingsummaryglobalNet() string`

GetDEzmaxinvoicingsummaryglobalNet returns the DEzmaxinvoicingsummaryglobalNet field if non-nil, zero value otherwise.

### GetDEzmaxinvoicingsummaryglobalNetOk

`func (o *EzmaxinvoicingsummaryglobalResponse) GetDEzmaxinvoicingsummaryglobalNetOk() (*string, bool)`

GetDEzmaxinvoicingsummaryglobalNetOk returns a tuple with the DEzmaxinvoicingsummaryglobalNet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDEzmaxinvoicingsummaryglobalNet

`func (o *EzmaxinvoicingsummaryglobalResponse) SetDEzmaxinvoicingsummaryglobalNet(v string)`

SetDEzmaxinvoicingsummaryglobalNet sets DEzmaxinvoicingsummaryglobalNet field to given value.

### HasDEzmaxinvoicingsummaryglobalNet

`func (o *EzmaxinvoicingsummaryglobalResponse) HasDEzmaxinvoicingsummaryglobalNet() bool`

HasDEzmaxinvoicingsummaryglobalNet returns a boolean if a field has been set.

### GetBEzmaxinvoicingsummaryglobalAdjustment

`func (o *EzmaxinvoicingsummaryglobalResponse) GetBEzmaxinvoicingsummaryglobalAdjustment() bool`

GetBEzmaxinvoicingsummaryglobalAdjustment returns the BEzmaxinvoicingsummaryglobalAdjustment field if non-nil, zero value otherwise.

### GetBEzmaxinvoicingsummaryglobalAdjustmentOk

`func (o *EzmaxinvoicingsummaryglobalResponse) GetBEzmaxinvoicingsummaryglobalAdjustmentOk() (*bool, bool)`

GetBEzmaxinvoicingsummaryglobalAdjustmentOk returns a tuple with the BEzmaxinvoicingsummaryglobalAdjustment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzmaxinvoicingsummaryglobalAdjustment

`func (o *EzmaxinvoicingsummaryglobalResponse) SetBEzmaxinvoicingsummaryglobalAdjustment(v bool)`

SetBEzmaxinvoicingsummaryglobalAdjustment sets BEzmaxinvoicingsummaryglobalAdjustment field to given value.


### GetTEzmaxproductHelpX

`func (o *EzmaxinvoicingsummaryglobalResponse) GetTEzmaxproductHelpX() string`

GetTEzmaxproductHelpX returns the TEzmaxproductHelpX field if non-nil, zero value otherwise.

### GetTEzmaxproductHelpXOk

`func (o *EzmaxinvoicingsummaryglobalResponse) GetTEzmaxproductHelpXOk() (*string, bool)`

GetTEzmaxproductHelpXOk returns a tuple with the TEzmaxproductHelpX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTEzmaxproductHelpX

`func (o *EzmaxinvoicingsummaryglobalResponse) SetTEzmaxproductHelpX(v string)`

SetTEzmaxproductHelpX sets TEzmaxproductHelpX field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


