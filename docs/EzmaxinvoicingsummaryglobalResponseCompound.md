# EzmaxinvoicingsummaryglobalResponseCompound

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
**AObjEzmaxinvoicingcommission** | Pointer to [**[]EzmaxinvoicingcommissionResponseCompound**](EzmaxinvoicingcommissionResponseCompound.md) |  | [optional] 

## Methods

### NewEzmaxinvoicingsummaryglobalResponseCompound

`func NewEzmaxinvoicingsummaryglobalResponseCompound(fkiEzmaxproductID int32, sEzmaxproductDescriptionX string, dtEzmaxinvoicingsummaryglobalStart string, dtEzmaxinvoicingsummaryglobalEnd string, iEzmaxinvoicingsummaryglobalDays int32, dEzmaxinvoicingsummaryglobalCountreal string, dEzmaxinvoicingsummaryglobalCountbilled string, dEzmaxinvoicingsummaryglobalSubtotal string, dEzmaxinvoicingsummaryglobalRebateamount string, dEzmaxinvoicingsummaryglobalRebatepercent string, dEzmaxinvoicingsummaryglobalRebatetotal string, dEzmaxinvoicingsummaryglobalTotal string, bEzmaxinvoicingsummaryglobalAdjustment bool, tEzmaxproductHelpX string, ) *EzmaxinvoicingsummaryglobalResponseCompound`

NewEzmaxinvoicingsummaryglobalResponseCompound instantiates a new EzmaxinvoicingsummaryglobalResponseCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzmaxinvoicingsummaryglobalResponseCompoundWithDefaults

`func NewEzmaxinvoicingsummaryglobalResponseCompoundWithDefaults() *EzmaxinvoicingsummaryglobalResponseCompound`

NewEzmaxinvoicingsummaryglobalResponseCompoundWithDefaults instantiates a new EzmaxinvoicingsummaryglobalResponseCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzmaxinvoicingsummaryglobalID

`func (o *EzmaxinvoicingsummaryglobalResponseCompound) GetPkiEzmaxinvoicingsummaryglobalID() int32`

GetPkiEzmaxinvoicingsummaryglobalID returns the PkiEzmaxinvoicingsummaryglobalID field if non-nil, zero value otherwise.

### GetPkiEzmaxinvoicingsummaryglobalIDOk

`func (o *EzmaxinvoicingsummaryglobalResponseCompound) GetPkiEzmaxinvoicingsummaryglobalIDOk() (*int32, bool)`

GetPkiEzmaxinvoicingsummaryglobalIDOk returns a tuple with the PkiEzmaxinvoicingsummaryglobalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzmaxinvoicingsummaryglobalID

`func (o *EzmaxinvoicingsummaryglobalResponseCompound) SetPkiEzmaxinvoicingsummaryglobalID(v int32)`

SetPkiEzmaxinvoicingsummaryglobalID sets PkiEzmaxinvoicingsummaryglobalID field to given value.

### HasPkiEzmaxinvoicingsummaryglobalID

`func (o *EzmaxinvoicingsummaryglobalResponseCompound) HasPkiEzmaxinvoicingsummaryglobalID() bool`

HasPkiEzmaxinvoicingsummaryglobalID returns a boolean if a field has been set.

### GetFkiEzmaxinvoicingID

`func (o *EzmaxinvoicingsummaryglobalResponseCompound) GetFkiEzmaxinvoicingID() int32`

GetFkiEzmaxinvoicingID returns the FkiEzmaxinvoicingID field if non-nil, zero value otherwise.

### GetFkiEzmaxinvoicingIDOk

`func (o *EzmaxinvoicingsummaryglobalResponseCompound) GetFkiEzmaxinvoicingIDOk() (*int32, bool)`

GetFkiEzmaxinvoicingIDOk returns a tuple with the FkiEzmaxinvoicingID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzmaxinvoicingID

`func (o *EzmaxinvoicingsummaryglobalResponseCompound) SetFkiEzmaxinvoicingID(v int32)`

SetFkiEzmaxinvoicingID sets FkiEzmaxinvoicingID field to given value.

### HasFkiEzmaxinvoicingID

`func (o *EzmaxinvoicingsummaryglobalResponseCompound) HasFkiEzmaxinvoicingID() bool`

HasFkiEzmaxinvoicingID returns a boolean if a field has been set.

### GetFkiEzmaxproductID

`func (o *EzmaxinvoicingsummaryglobalResponseCompound) GetFkiEzmaxproductID() int32`

GetFkiEzmaxproductID returns the FkiEzmaxproductID field if non-nil, zero value otherwise.

### GetFkiEzmaxproductIDOk

`func (o *EzmaxinvoicingsummaryglobalResponseCompound) GetFkiEzmaxproductIDOk() (*int32, bool)`

GetFkiEzmaxproductIDOk returns a tuple with the FkiEzmaxproductID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzmaxproductID

`func (o *EzmaxinvoicingsummaryglobalResponseCompound) SetFkiEzmaxproductID(v int32)`

SetFkiEzmaxproductID sets FkiEzmaxproductID field to given value.


### GetSEzmaxproductDescriptionX

`func (o *EzmaxinvoicingsummaryglobalResponseCompound) GetSEzmaxproductDescriptionX() string`

GetSEzmaxproductDescriptionX returns the SEzmaxproductDescriptionX field if non-nil, zero value otherwise.

### GetSEzmaxproductDescriptionXOk

`func (o *EzmaxinvoicingsummaryglobalResponseCompound) GetSEzmaxproductDescriptionXOk() (*string, bool)`

GetSEzmaxproductDescriptionXOk returns a tuple with the SEzmaxproductDescriptionX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzmaxproductDescriptionX

`func (o *EzmaxinvoicingsummaryglobalResponseCompound) SetSEzmaxproductDescriptionX(v string)`

SetSEzmaxproductDescriptionX sets SEzmaxproductDescriptionX field to given value.


### GetDtEzmaxinvoicingsummaryglobalStart

`func (o *EzmaxinvoicingsummaryglobalResponseCompound) GetDtEzmaxinvoicingsummaryglobalStart() string`

GetDtEzmaxinvoicingsummaryglobalStart returns the DtEzmaxinvoicingsummaryglobalStart field if non-nil, zero value otherwise.

### GetDtEzmaxinvoicingsummaryglobalStartOk

`func (o *EzmaxinvoicingsummaryglobalResponseCompound) GetDtEzmaxinvoicingsummaryglobalStartOk() (*string, bool)`

GetDtEzmaxinvoicingsummaryglobalStartOk returns a tuple with the DtEzmaxinvoicingsummaryglobalStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtEzmaxinvoicingsummaryglobalStart

`func (o *EzmaxinvoicingsummaryglobalResponseCompound) SetDtEzmaxinvoicingsummaryglobalStart(v string)`

SetDtEzmaxinvoicingsummaryglobalStart sets DtEzmaxinvoicingsummaryglobalStart field to given value.


### GetDtEzmaxinvoicingsummaryglobalEnd

`func (o *EzmaxinvoicingsummaryglobalResponseCompound) GetDtEzmaxinvoicingsummaryglobalEnd() string`

GetDtEzmaxinvoicingsummaryglobalEnd returns the DtEzmaxinvoicingsummaryglobalEnd field if non-nil, zero value otherwise.

### GetDtEzmaxinvoicingsummaryglobalEndOk

`func (o *EzmaxinvoicingsummaryglobalResponseCompound) GetDtEzmaxinvoicingsummaryglobalEndOk() (*string, bool)`

GetDtEzmaxinvoicingsummaryglobalEndOk returns a tuple with the DtEzmaxinvoicingsummaryglobalEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtEzmaxinvoicingsummaryglobalEnd

`func (o *EzmaxinvoicingsummaryglobalResponseCompound) SetDtEzmaxinvoicingsummaryglobalEnd(v string)`

SetDtEzmaxinvoicingsummaryglobalEnd sets DtEzmaxinvoicingsummaryglobalEnd field to given value.


### GetIEzmaxinvoicingsummaryglobalDays

`func (o *EzmaxinvoicingsummaryglobalResponseCompound) GetIEzmaxinvoicingsummaryglobalDays() int32`

GetIEzmaxinvoicingsummaryglobalDays returns the IEzmaxinvoicingsummaryglobalDays field if non-nil, zero value otherwise.

### GetIEzmaxinvoicingsummaryglobalDaysOk

`func (o *EzmaxinvoicingsummaryglobalResponseCompound) GetIEzmaxinvoicingsummaryglobalDaysOk() (*int32, bool)`

GetIEzmaxinvoicingsummaryglobalDaysOk returns a tuple with the IEzmaxinvoicingsummaryglobalDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzmaxinvoicingsummaryglobalDays

`func (o *EzmaxinvoicingsummaryglobalResponseCompound) SetIEzmaxinvoicingsummaryglobalDays(v int32)`

SetIEzmaxinvoicingsummaryglobalDays sets IEzmaxinvoicingsummaryglobalDays field to given value.


### GetDEzmaxinvoicingsummaryglobalCountreal

`func (o *EzmaxinvoicingsummaryglobalResponseCompound) GetDEzmaxinvoicingsummaryglobalCountreal() string`

GetDEzmaxinvoicingsummaryglobalCountreal returns the DEzmaxinvoicingsummaryglobalCountreal field if non-nil, zero value otherwise.

### GetDEzmaxinvoicingsummaryglobalCountrealOk

`func (o *EzmaxinvoicingsummaryglobalResponseCompound) GetDEzmaxinvoicingsummaryglobalCountrealOk() (*string, bool)`

GetDEzmaxinvoicingsummaryglobalCountrealOk returns a tuple with the DEzmaxinvoicingsummaryglobalCountreal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDEzmaxinvoicingsummaryglobalCountreal

`func (o *EzmaxinvoicingsummaryglobalResponseCompound) SetDEzmaxinvoicingsummaryglobalCountreal(v string)`

SetDEzmaxinvoicingsummaryglobalCountreal sets DEzmaxinvoicingsummaryglobalCountreal field to given value.


### GetDEzmaxinvoicingsummaryglobalCountbilled

`func (o *EzmaxinvoicingsummaryglobalResponseCompound) GetDEzmaxinvoicingsummaryglobalCountbilled() string`

GetDEzmaxinvoicingsummaryglobalCountbilled returns the DEzmaxinvoicingsummaryglobalCountbilled field if non-nil, zero value otherwise.

### GetDEzmaxinvoicingsummaryglobalCountbilledOk

`func (o *EzmaxinvoicingsummaryglobalResponseCompound) GetDEzmaxinvoicingsummaryglobalCountbilledOk() (*string, bool)`

GetDEzmaxinvoicingsummaryglobalCountbilledOk returns a tuple with the DEzmaxinvoicingsummaryglobalCountbilled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDEzmaxinvoicingsummaryglobalCountbilled

`func (o *EzmaxinvoicingsummaryglobalResponseCompound) SetDEzmaxinvoicingsummaryglobalCountbilled(v string)`

SetDEzmaxinvoicingsummaryglobalCountbilled sets DEzmaxinvoicingsummaryglobalCountbilled field to given value.


### GetDEzmaxinvoicingsummaryglobalSubtotal

`func (o *EzmaxinvoicingsummaryglobalResponseCompound) GetDEzmaxinvoicingsummaryglobalSubtotal() string`

GetDEzmaxinvoicingsummaryglobalSubtotal returns the DEzmaxinvoicingsummaryglobalSubtotal field if non-nil, zero value otherwise.

### GetDEzmaxinvoicingsummaryglobalSubtotalOk

`func (o *EzmaxinvoicingsummaryglobalResponseCompound) GetDEzmaxinvoicingsummaryglobalSubtotalOk() (*string, bool)`

GetDEzmaxinvoicingsummaryglobalSubtotalOk returns a tuple with the DEzmaxinvoicingsummaryglobalSubtotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDEzmaxinvoicingsummaryglobalSubtotal

`func (o *EzmaxinvoicingsummaryglobalResponseCompound) SetDEzmaxinvoicingsummaryglobalSubtotal(v string)`

SetDEzmaxinvoicingsummaryglobalSubtotal sets DEzmaxinvoicingsummaryglobalSubtotal field to given value.


### GetDEzmaxinvoicingsummaryglobalRebateamount

`func (o *EzmaxinvoicingsummaryglobalResponseCompound) GetDEzmaxinvoicingsummaryglobalRebateamount() string`

GetDEzmaxinvoicingsummaryglobalRebateamount returns the DEzmaxinvoicingsummaryglobalRebateamount field if non-nil, zero value otherwise.

### GetDEzmaxinvoicingsummaryglobalRebateamountOk

`func (o *EzmaxinvoicingsummaryglobalResponseCompound) GetDEzmaxinvoicingsummaryglobalRebateamountOk() (*string, bool)`

GetDEzmaxinvoicingsummaryglobalRebateamountOk returns a tuple with the DEzmaxinvoicingsummaryglobalRebateamount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDEzmaxinvoicingsummaryglobalRebateamount

`func (o *EzmaxinvoicingsummaryglobalResponseCompound) SetDEzmaxinvoicingsummaryglobalRebateamount(v string)`

SetDEzmaxinvoicingsummaryglobalRebateamount sets DEzmaxinvoicingsummaryglobalRebateamount field to given value.


### GetDEzmaxinvoicingsummaryglobalRebatepercent

`func (o *EzmaxinvoicingsummaryglobalResponseCompound) GetDEzmaxinvoicingsummaryglobalRebatepercent() string`

GetDEzmaxinvoicingsummaryglobalRebatepercent returns the DEzmaxinvoicingsummaryglobalRebatepercent field if non-nil, zero value otherwise.

### GetDEzmaxinvoicingsummaryglobalRebatepercentOk

`func (o *EzmaxinvoicingsummaryglobalResponseCompound) GetDEzmaxinvoicingsummaryglobalRebatepercentOk() (*string, bool)`

GetDEzmaxinvoicingsummaryglobalRebatepercentOk returns a tuple with the DEzmaxinvoicingsummaryglobalRebatepercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDEzmaxinvoicingsummaryglobalRebatepercent

`func (o *EzmaxinvoicingsummaryglobalResponseCompound) SetDEzmaxinvoicingsummaryglobalRebatepercent(v string)`

SetDEzmaxinvoicingsummaryglobalRebatepercent sets DEzmaxinvoicingsummaryglobalRebatepercent field to given value.


### GetDEzmaxinvoicingsummaryglobalRebatetotal

`func (o *EzmaxinvoicingsummaryglobalResponseCompound) GetDEzmaxinvoicingsummaryglobalRebatetotal() string`

GetDEzmaxinvoicingsummaryglobalRebatetotal returns the DEzmaxinvoicingsummaryglobalRebatetotal field if non-nil, zero value otherwise.

### GetDEzmaxinvoicingsummaryglobalRebatetotalOk

`func (o *EzmaxinvoicingsummaryglobalResponseCompound) GetDEzmaxinvoicingsummaryglobalRebatetotalOk() (*string, bool)`

GetDEzmaxinvoicingsummaryglobalRebatetotalOk returns a tuple with the DEzmaxinvoicingsummaryglobalRebatetotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDEzmaxinvoicingsummaryglobalRebatetotal

`func (o *EzmaxinvoicingsummaryglobalResponseCompound) SetDEzmaxinvoicingsummaryglobalRebatetotal(v string)`

SetDEzmaxinvoicingsummaryglobalRebatetotal sets DEzmaxinvoicingsummaryglobalRebatetotal field to given value.


### GetDEzmaxinvoicingsummaryglobalTotal

`func (o *EzmaxinvoicingsummaryglobalResponseCompound) GetDEzmaxinvoicingsummaryglobalTotal() string`

GetDEzmaxinvoicingsummaryglobalTotal returns the DEzmaxinvoicingsummaryglobalTotal field if non-nil, zero value otherwise.

### GetDEzmaxinvoicingsummaryglobalTotalOk

`func (o *EzmaxinvoicingsummaryglobalResponseCompound) GetDEzmaxinvoicingsummaryglobalTotalOk() (*string, bool)`

GetDEzmaxinvoicingsummaryglobalTotalOk returns a tuple with the DEzmaxinvoicingsummaryglobalTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDEzmaxinvoicingsummaryglobalTotal

`func (o *EzmaxinvoicingsummaryglobalResponseCompound) SetDEzmaxinvoicingsummaryglobalTotal(v string)`

SetDEzmaxinvoicingsummaryglobalTotal sets DEzmaxinvoicingsummaryglobalTotal field to given value.


### GetDEzmaxinvoicingsummaryglobalRepresentative

`func (o *EzmaxinvoicingsummaryglobalResponseCompound) GetDEzmaxinvoicingsummaryglobalRepresentative() string`

GetDEzmaxinvoicingsummaryglobalRepresentative returns the DEzmaxinvoicingsummaryglobalRepresentative field if non-nil, zero value otherwise.

### GetDEzmaxinvoicingsummaryglobalRepresentativeOk

`func (o *EzmaxinvoicingsummaryglobalResponseCompound) GetDEzmaxinvoicingsummaryglobalRepresentativeOk() (*string, bool)`

GetDEzmaxinvoicingsummaryglobalRepresentativeOk returns a tuple with the DEzmaxinvoicingsummaryglobalRepresentative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDEzmaxinvoicingsummaryglobalRepresentative

`func (o *EzmaxinvoicingsummaryglobalResponseCompound) SetDEzmaxinvoicingsummaryglobalRepresentative(v string)`

SetDEzmaxinvoicingsummaryglobalRepresentative sets DEzmaxinvoicingsummaryglobalRepresentative field to given value.

### HasDEzmaxinvoicingsummaryglobalRepresentative

`func (o *EzmaxinvoicingsummaryglobalResponseCompound) HasDEzmaxinvoicingsummaryglobalRepresentative() bool`

HasDEzmaxinvoicingsummaryglobalRepresentative returns a boolean if a field has been set.

### GetDEzmaxinvoicingsummaryglobalPartner

`func (o *EzmaxinvoicingsummaryglobalResponseCompound) GetDEzmaxinvoicingsummaryglobalPartner() string`

GetDEzmaxinvoicingsummaryglobalPartner returns the DEzmaxinvoicingsummaryglobalPartner field if non-nil, zero value otherwise.

### GetDEzmaxinvoicingsummaryglobalPartnerOk

`func (o *EzmaxinvoicingsummaryglobalResponseCompound) GetDEzmaxinvoicingsummaryglobalPartnerOk() (*string, bool)`

GetDEzmaxinvoicingsummaryglobalPartnerOk returns a tuple with the DEzmaxinvoicingsummaryglobalPartner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDEzmaxinvoicingsummaryglobalPartner

`func (o *EzmaxinvoicingsummaryglobalResponseCompound) SetDEzmaxinvoicingsummaryglobalPartner(v string)`

SetDEzmaxinvoicingsummaryglobalPartner sets DEzmaxinvoicingsummaryglobalPartner field to given value.

### HasDEzmaxinvoicingsummaryglobalPartner

`func (o *EzmaxinvoicingsummaryglobalResponseCompound) HasDEzmaxinvoicingsummaryglobalPartner() bool`

HasDEzmaxinvoicingsummaryglobalPartner returns a boolean if a field has been set.

### GetDEzmaxinvoicingsummaryglobalNet

`func (o *EzmaxinvoicingsummaryglobalResponseCompound) GetDEzmaxinvoicingsummaryglobalNet() string`

GetDEzmaxinvoicingsummaryglobalNet returns the DEzmaxinvoicingsummaryglobalNet field if non-nil, zero value otherwise.

### GetDEzmaxinvoicingsummaryglobalNetOk

`func (o *EzmaxinvoicingsummaryglobalResponseCompound) GetDEzmaxinvoicingsummaryglobalNetOk() (*string, bool)`

GetDEzmaxinvoicingsummaryglobalNetOk returns a tuple with the DEzmaxinvoicingsummaryglobalNet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDEzmaxinvoicingsummaryglobalNet

`func (o *EzmaxinvoicingsummaryglobalResponseCompound) SetDEzmaxinvoicingsummaryglobalNet(v string)`

SetDEzmaxinvoicingsummaryglobalNet sets DEzmaxinvoicingsummaryglobalNet field to given value.

### HasDEzmaxinvoicingsummaryglobalNet

`func (o *EzmaxinvoicingsummaryglobalResponseCompound) HasDEzmaxinvoicingsummaryglobalNet() bool`

HasDEzmaxinvoicingsummaryglobalNet returns a boolean if a field has been set.

### GetBEzmaxinvoicingsummaryglobalAdjustment

`func (o *EzmaxinvoicingsummaryglobalResponseCompound) GetBEzmaxinvoicingsummaryglobalAdjustment() bool`

GetBEzmaxinvoicingsummaryglobalAdjustment returns the BEzmaxinvoicingsummaryglobalAdjustment field if non-nil, zero value otherwise.

### GetBEzmaxinvoicingsummaryglobalAdjustmentOk

`func (o *EzmaxinvoicingsummaryglobalResponseCompound) GetBEzmaxinvoicingsummaryglobalAdjustmentOk() (*bool, bool)`

GetBEzmaxinvoicingsummaryglobalAdjustmentOk returns a tuple with the BEzmaxinvoicingsummaryglobalAdjustment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzmaxinvoicingsummaryglobalAdjustment

`func (o *EzmaxinvoicingsummaryglobalResponseCompound) SetBEzmaxinvoicingsummaryglobalAdjustment(v bool)`

SetBEzmaxinvoicingsummaryglobalAdjustment sets BEzmaxinvoicingsummaryglobalAdjustment field to given value.


### GetTEzmaxproductHelpX

`func (o *EzmaxinvoicingsummaryglobalResponseCompound) GetTEzmaxproductHelpX() string`

GetTEzmaxproductHelpX returns the TEzmaxproductHelpX field if non-nil, zero value otherwise.

### GetTEzmaxproductHelpXOk

`func (o *EzmaxinvoicingsummaryglobalResponseCompound) GetTEzmaxproductHelpXOk() (*string, bool)`

GetTEzmaxproductHelpXOk returns a tuple with the TEzmaxproductHelpX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTEzmaxproductHelpX

`func (o *EzmaxinvoicingsummaryglobalResponseCompound) SetTEzmaxproductHelpX(v string)`

SetTEzmaxproductHelpX sets TEzmaxproductHelpX field to given value.


### GetAObjEzmaxinvoicingcommission

`func (o *EzmaxinvoicingsummaryglobalResponseCompound) GetAObjEzmaxinvoicingcommission() []EzmaxinvoicingcommissionResponseCompound`

GetAObjEzmaxinvoicingcommission returns the AObjEzmaxinvoicingcommission field if non-nil, zero value otherwise.

### GetAObjEzmaxinvoicingcommissionOk

`func (o *EzmaxinvoicingsummaryglobalResponseCompound) GetAObjEzmaxinvoicingcommissionOk() (*[]EzmaxinvoicingcommissionResponseCompound, bool)`

GetAObjEzmaxinvoicingcommissionOk returns a tuple with the AObjEzmaxinvoicingcommission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjEzmaxinvoicingcommission

`func (o *EzmaxinvoicingsummaryglobalResponseCompound) SetAObjEzmaxinvoicingcommission(v []EzmaxinvoicingcommissionResponseCompound)`

SetAObjEzmaxinvoicingcommission sets AObjEzmaxinvoicingcommission field to given value.

### HasAObjEzmaxinvoicingcommission

`func (o *EzmaxinvoicingsummaryglobalResponseCompound) HasAObjEzmaxinvoicingcommission() bool`

HasAObjEzmaxinvoicingcommission returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


