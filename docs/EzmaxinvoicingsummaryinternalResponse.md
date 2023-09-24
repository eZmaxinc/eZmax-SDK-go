# EzmaxinvoicingsummaryinternalResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzmaxinvoicingsummaryinternalID** | Pointer to **int32** | The unique ID of the Ezmaxinvoicingsummaryinternal | [optional] 
**ObjEzmaxinvoicingsummaryinternalDescription** | [**MultilingualEzmaxinvoicingsummaryinternalDescription**](MultilingualEzmaxinvoicingsummaryinternalDescription.md) |  | 
**SEzmaxinvoicingsummaryinternalDescriptionX** | **string** | The Ezmaxinvoicingsummaryinternal description in the language of the requester | 
**FkiEzmaxinvoicingID** | Pointer to **int32** | The unique ID of the Ezmaxinvoicing | [optional] 
**FkiBillingentityinternalID** | **int32** | The unique ID of the Billingentityinternal. | 
**SBillingentityinternalDescriptionX** | **string** | The description of the Billingentityinternal in the language of the requester | 

## Methods

### NewEzmaxinvoicingsummaryinternalResponse

`func NewEzmaxinvoicingsummaryinternalResponse(objEzmaxinvoicingsummaryinternalDescription MultilingualEzmaxinvoicingsummaryinternalDescription, sEzmaxinvoicingsummaryinternalDescriptionX string, fkiBillingentityinternalID int32, sBillingentityinternalDescriptionX string, ) *EzmaxinvoicingsummaryinternalResponse`

NewEzmaxinvoicingsummaryinternalResponse instantiates a new EzmaxinvoicingsummaryinternalResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzmaxinvoicingsummaryinternalResponseWithDefaults

`func NewEzmaxinvoicingsummaryinternalResponseWithDefaults() *EzmaxinvoicingsummaryinternalResponse`

NewEzmaxinvoicingsummaryinternalResponseWithDefaults instantiates a new EzmaxinvoicingsummaryinternalResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzmaxinvoicingsummaryinternalID

`func (o *EzmaxinvoicingsummaryinternalResponse) GetPkiEzmaxinvoicingsummaryinternalID() int32`

GetPkiEzmaxinvoicingsummaryinternalID returns the PkiEzmaxinvoicingsummaryinternalID field if non-nil, zero value otherwise.

### GetPkiEzmaxinvoicingsummaryinternalIDOk

`func (o *EzmaxinvoicingsummaryinternalResponse) GetPkiEzmaxinvoicingsummaryinternalIDOk() (*int32, bool)`

GetPkiEzmaxinvoicingsummaryinternalIDOk returns a tuple with the PkiEzmaxinvoicingsummaryinternalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzmaxinvoicingsummaryinternalID

`func (o *EzmaxinvoicingsummaryinternalResponse) SetPkiEzmaxinvoicingsummaryinternalID(v int32)`

SetPkiEzmaxinvoicingsummaryinternalID sets PkiEzmaxinvoicingsummaryinternalID field to given value.

### HasPkiEzmaxinvoicingsummaryinternalID

`func (o *EzmaxinvoicingsummaryinternalResponse) HasPkiEzmaxinvoicingsummaryinternalID() bool`

HasPkiEzmaxinvoicingsummaryinternalID returns a boolean if a field has been set.

### GetObjEzmaxinvoicingsummaryinternalDescription

`func (o *EzmaxinvoicingsummaryinternalResponse) GetObjEzmaxinvoicingsummaryinternalDescription() MultilingualEzmaxinvoicingsummaryinternalDescription`

GetObjEzmaxinvoicingsummaryinternalDescription returns the ObjEzmaxinvoicingsummaryinternalDescription field if non-nil, zero value otherwise.

### GetObjEzmaxinvoicingsummaryinternalDescriptionOk

`func (o *EzmaxinvoicingsummaryinternalResponse) GetObjEzmaxinvoicingsummaryinternalDescriptionOk() (*MultilingualEzmaxinvoicingsummaryinternalDescription, bool)`

GetObjEzmaxinvoicingsummaryinternalDescriptionOk returns a tuple with the ObjEzmaxinvoicingsummaryinternalDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjEzmaxinvoicingsummaryinternalDescription

`func (o *EzmaxinvoicingsummaryinternalResponse) SetObjEzmaxinvoicingsummaryinternalDescription(v MultilingualEzmaxinvoicingsummaryinternalDescription)`

SetObjEzmaxinvoicingsummaryinternalDescription sets ObjEzmaxinvoicingsummaryinternalDescription field to given value.


### GetSEzmaxinvoicingsummaryinternalDescriptionX

`func (o *EzmaxinvoicingsummaryinternalResponse) GetSEzmaxinvoicingsummaryinternalDescriptionX() string`

GetSEzmaxinvoicingsummaryinternalDescriptionX returns the SEzmaxinvoicingsummaryinternalDescriptionX field if non-nil, zero value otherwise.

### GetSEzmaxinvoicingsummaryinternalDescriptionXOk

`func (o *EzmaxinvoicingsummaryinternalResponse) GetSEzmaxinvoicingsummaryinternalDescriptionXOk() (*string, bool)`

GetSEzmaxinvoicingsummaryinternalDescriptionXOk returns a tuple with the SEzmaxinvoicingsummaryinternalDescriptionX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzmaxinvoicingsummaryinternalDescriptionX

`func (o *EzmaxinvoicingsummaryinternalResponse) SetSEzmaxinvoicingsummaryinternalDescriptionX(v string)`

SetSEzmaxinvoicingsummaryinternalDescriptionX sets SEzmaxinvoicingsummaryinternalDescriptionX field to given value.


### GetFkiEzmaxinvoicingID

`func (o *EzmaxinvoicingsummaryinternalResponse) GetFkiEzmaxinvoicingID() int32`

GetFkiEzmaxinvoicingID returns the FkiEzmaxinvoicingID field if non-nil, zero value otherwise.

### GetFkiEzmaxinvoicingIDOk

`func (o *EzmaxinvoicingsummaryinternalResponse) GetFkiEzmaxinvoicingIDOk() (*int32, bool)`

GetFkiEzmaxinvoicingIDOk returns a tuple with the FkiEzmaxinvoicingID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzmaxinvoicingID

`func (o *EzmaxinvoicingsummaryinternalResponse) SetFkiEzmaxinvoicingID(v int32)`

SetFkiEzmaxinvoicingID sets FkiEzmaxinvoicingID field to given value.

### HasFkiEzmaxinvoicingID

`func (o *EzmaxinvoicingsummaryinternalResponse) HasFkiEzmaxinvoicingID() bool`

HasFkiEzmaxinvoicingID returns a boolean if a field has been set.

### GetFkiBillingentityinternalID

`func (o *EzmaxinvoicingsummaryinternalResponse) GetFkiBillingentityinternalID() int32`

GetFkiBillingentityinternalID returns the FkiBillingentityinternalID field if non-nil, zero value otherwise.

### GetFkiBillingentityinternalIDOk

`func (o *EzmaxinvoicingsummaryinternalResponse) GetFkiBillingentityinternalIDOk() (*int32, bool)`

GetFkiBillingentityinternalIDOk returns a tuple with the FkiBillingentityinternalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiBillingentityinternalID

`func (o *EzmaxinvoicingsummaryinternalResponse) SetFkiBillingentityinternalID(v int32)`

SetFkiBillingentityinternalID sets FkiBillingentityinternalID field to given value.


### GetSBillingentityinternalDescriptionX

`func (o *EzmaxinvoicingsummaryinternalResponse) GetSBillingentityinternalDescriptionX() string`

GetSBillingentityinternalDescriptionX returns the SBillingentityinternalDescriptionX field if non-nil, zero value otherwise.

### GetSBillingentityinternalDescriptionXOk

`func (o *EzmaxinvoicingsummaryinternalResponse) GetSBillingentityinternalDescriptionXOk() (*string, bool)`

GetSBillingentityinternalDescriptionXOk returns a tuple with the SBillingentityinternalDescriptionX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSBillingentityinternalDescriptionX

`func (o *EzmaxinvoicingsummaryinternalResponse) SetSBillingentityinternalDescriptionX(v string)`

SetSBillingentityinternalDescriptionX sets SBillingentityinternalDescriptionX field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


