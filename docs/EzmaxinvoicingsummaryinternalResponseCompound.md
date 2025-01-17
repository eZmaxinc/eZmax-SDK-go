# EzmaxinvoicingsummaryinternalResponseCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzmaxinvoicingsummaryinternalID** | Pointer to **int32** | The unique ID of the Ezmaxinvoicingsummaryinternal | [optional] 
**ObjEzmaxinvoicingsummaryinternalDescription** | [**MultilingualEzmaxinvoicingsummaryinternalDescription**](MultilingualEzmaxinvoicingsummaryinternalDescription.md) |  | 
**SEzmaxinvoicingsummaryinternalDescriptionX** | **string** | The Ezmaxinvoicingsummaryinternal description in the language of the requester | 
**FkiEzmaxinvoicingID** | Pointer to **int32** | The unique ID of the Ezmaxinvoicing | [optional] 
**FkiBillingentityinternalID** | **int32** | The unique ID of the Billingentityinternal. | 
**SBillingentityinternalDescriptionX** | **string** | The description of the Billingentityinternal in the language of the requester | 
**AObjEzmaxinvoicingsummaryinternaldetail** | [**[]EzmaxinvoicingsummaryinternaldetailResponseCompound**](EzmaxinvoicingsummaryinternaldetailResponseCompound.md) |  | 

## Methods

### NewEzmaxinvoicingsummaryinternalResponseCompound

`func NewEzmaxinvoicingsummaryinternalResponseCompound(objEzmaxinvoicingsummaryinternalDescription MultilingualEzmaxinvoicingsummaryinternalDescription, sEzmaxinvoicingsummaryinternalDescriptionX string, fkiBillingentityinternalID int32, sBillingentityinternalDescriptionX string, aObjEzmaxinvoicingsummaryinternaldetail []EzmaxinvoicingsummaryinternaldetailResponseCompound, ) *EzmaxinvoicingsummaryinternalResponseCompound`

NewEzmaxinvoicingsummaryinternalResponseCompound instantiates a new EzmaxinvoicingsummaryinternalResponseCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzmaxinvoicingsummaryinternalResponseCompoundWithDefaults

`func NewEzmaxinvoicingsummaryinternalResponseCompoundWithDefaults() *EzmaxinvoicingsummaryinternalResponseCompound`

NewEzmaxinvoicingsummaryinternalResponseCompoundWithDefaults instantiates a new EzmaxinvoicingsummaryinternalResponseCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzmaxinvoicingsummaryinternalID

`func (o *EzmaxinvoicingsummaryinternalResponseCompound) GetPkiEzmaxinvoicingsummaryinternalID() int32`

GetPkiEzmaxinvoicingsummaryinternalID returns the PkiEzmaxinvoicingsummaryinternalID field if non-nil, zero value otherwise.

### GetPkiEzmaxinvoicingsummaryinternalIDOk

`func (o *EzmaxinvoicingsummaryinternalResponseCompound) GetPkiEzmaxinvoicingsummaryinternalIDOk() (*int32, bool)`

GetPkiEzmaxinvoicingsummaryinternalIDOk returns a tuple with the PkiEzmaxinvoicingsummaryinternalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzmaxinvoicingsummaryinternalID

`func (o *EzmaxinvoicingsummaryinternalResponseCompound) SetPkiEzmaxinvoicingsummaryinternalID(v int32)`

SetPkiEzmaxinvoicingsummaryinternalID sets PkiEzmaxinvoicingsummaryinternalID field to given value.

### HasPkiEzmaxinvoicingsummaryinternalID

`func (o *EzmaxinvoicingsummaryinternalResponseCompound) HasPkiEzmaxinvoicingsummaryinternalID() bool`

HasPkiEzmaxinvoicingsummaryinternalID returns a boolean if a field has been set.

### GetObjEzmaxinvoicingsummaryinternalDescription

`func (o *EzmaxinvoicingsummaryinternalResponseCompound) GetObjEzmaxinvoicingsummaryinternalDescription() MultilingualEzmaxinvoicingsummaryinternalDescription`

GetObjEzmaxinvoicingsummaryinternalDescription returns the ObjEzmaxinvoicingsummaryinternalDescription field if non-nil, zero value otherwise.

### GetObjEzmaxinvoicingsummaryinternalDescriptionOk

`func (o *EzmaxinvoicingsummaryinternalResponseCompound) GetObjEzmaxinvoicingsummaryinternalDescriptionOk() (*MultilingualEzmaxinvoicingsummaryinternalDescription, bool)`

GetObjEzmaxinvoicingsummaryinternalDescriptionOk returns a tuple with the ObjEzmaxinvoicingsummaryinternalDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjEzmaxinvoicingsummaryinternalDescription

`func (o *EzmaxinvoicingsummaryinternalResponseCompound) SetObjEzmaxinvoicingsummaryinternalDescription(v MultilingualEzmaxinvoicingsummaryinternalDescription)`

SetObjEzmaxinvoicingsummaryinternalDescription sets ObjEzmaxinvoicingsummaryinternalDescription field to given value.


### GetSEzmaxinvoicingsummaryinternalDescriptionX

`func (o *EzmaxinvoicingsummaryinternalResponseCompound) GetSEzmaxinvoicingsummaryinternalDescriptionX() string`

GetSEzmaxinvoicingsummaryinternalDescriptionX returns the SEzmaxinvoicingsummaryinternalDescriptionX field if non-nil, zero value otherwise.

### GetSEzmaxinvoicingsummaryinternalDescriptionXOk

`func (o *EzmaxinvoicingsummaryinternalResponseCompound) GetSEzmaxinvoicingsummaryinternalDescriptionXOk() (*string, bool)`

GetSEzmaxinvoicingsummaryinternalDescriptionXOk returns a tuple with the SEzmaxinvoicingsummaryinternalDescriptionX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzmaxinvoicingsummaryinternalDescriptionX

`func (o *EzmaxinvoicingsummaryinternalResponseCompound) SetSEzmaxinvoicingsummaryinternalDescriptionX(v string)`

SetSEzmaxinvoicingsummaryinternalDescriptionX sets SEzmaxinvoicingsummaryinternalDescriptionX field to given value.


### GetFkiEzmaxinvoicingID

`func (o *EzmaxinvoicingsummaryinternalResponseCompound) GetFkiEzmaxinvoicingID() int32`

GetFkiEzmaxinvoicingID returns the FkiEzmaxinvoicingID field if non-nil, zero value otherwise.

### GetFkiEzmaxinvoicingIDOk

`func (o *EzmaxinvoicingsummaryinternalResponseCompound) GetFkiEzmaxinvoicingIDOk() (*int32, bool)`

GetFkiEzmaxinvoicingIDOk returns a tuple with the FkiEzmaxinvoicingID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzmaxinvoicingID

`func (o *EzmaxinvoicingsummaryinternalResponseCompound) SetFkiEzmaxinvoicingID(v int32)`

SetFkiEzmaxinvoicingID sets FkiEzmaxinvoicingID field to given value.

### HasFkiEzmaxinvoicingID

`func (o *EzmaxinvoicingsummaryinternalResponseCompound) HasFkiEzmaxinvoicingID() bool`

HasFkiEzmaxinvoicingID returns a boolean if a field has been set.

### GetFkiBillingentityinternalID

`func (o *EzmaxinvoicingsummaryinternalResponseCompound) GetFkiBillingentityinternalID() int32`

GetFkiBillingentityinternalID returns the FkiBillingentityinternalID field if non-nil, zero value otherwise.

### GetFkiBillingentityinternalIDOk

`func (o *EzmaxinvoicingsummaryinternalResponseCompound) GetFkiBillingentityinternalIDOk() (*int32, bool)`

GetFkiBillingentityinternalIDOk returns a tuple with the FkiBillingentityinternalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiBillingentityinternalID

`func (o *EzmaxinvoicingsummaryinternalResponseCompound) SetFkiBillingentityinternalID(v int32)`

SetFkiBillingentityinternalID sets FkiBillingentityinternalID field to given value.


### GetSBillingentityinternalDescriptionX

`func (o *EzmaxinvoicingsummaryinternalResponseCompound) GetSBillingentityinternalDescriptionX() string`

GetSBillingentityinternalDescriptionX returns the SBillingentityinternalDescriptionX field if non-nil, zero value otherwise.

### GetSBillingentityinternalDescriptionXOk

`func (o *EzmaxinvoicingsummaryinternalResponseCompound) GetSBillingentityinternalDescriptionXOk() (*string, bool)`

GetSBillingentityinternalDescriptionXOk returns a tuple with the SBillingentityinternalDescriptionX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSBillingentityinternalDescriptionX

`func (o *EzmaxinvoicingsummaryinternalResponseCompound) SetSBillingentityinternalDescriptionX(v string)`

SetSBillingentityinternalDescriptionX sets SBillingentityinternalDescriptionX field to given value.


### GetAObjEzmaxinvoicingsummaryinternaldetail

`func (o *EzmaxinvoicingsummaryinternalResponseCompound) GetAObjEzmaxinvoicingsummaryinternaldetail() []EzmaxinvoicingsummaryinternaldetailResponseCompound`

GetAObjEzmaxinvoicingsummaryinternaldetail returns the AObjEzmaxinvoicingsummaryinternaldetail field if non-nil, zero value otherwise.

### GetAObjEzmaxinvoicingsummaryinternaldetailOk

`func (o *EzmaxinvoicingsummaryinternalResponseCompound) GetAObjEzmaxinvoicingsummaryinternaldetailOk() (*[]EzmaxinvoicingsummaryinternaldetailResponseCompound, bool)`

GetAObjEzmaxinvoicingsummaryinternaldetailOk returns a tuple with the AObjEzmaxinvoicingsummaryinternaldetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjEzmaxinvoicingsummaryinternaldetail

`func (o *EzmaxinvoicingsummaryinternalResponseCompound) SetAObjEzmaxinvoicingsummaryinternaldetail(v []EzmaxinvoicingsummaryinternaldetailResponseCompound)`

SetAObjEzmaxinvoicingsummaryinternaldetail sets AObjEzmaxinvoicingsummaryinternaldetail field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


