# EzmaxinvoicingsummaryexternalResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzmaxinvoicingsummaryexternalID** | Pointer to **int32** | The unique ID of the Ezmaxinvoicingsummaryexternal | [optional] 
**FkiEzmaxinvoicingID** | Pointer to **int32** | The unique ID of the Ezmaxinvoicing | [optional] 
**FkiBillingentityexternalID** | **int32** | The unique ID of the Billingentityexternal | 
**SBillingentityexternalDescription** | **string** | The description of the Billingentityexternal | 
**SEzmaxinvoicingsummaryexternalDescription** | **string** | The description of the Ezmaxinvoicingsummaryexternal | 

## Methods

### NewEzmaxinvoicingsummaryexternalResponse

`func NewEzmaxinvoicingsummaryexternalResponse(fkiBillingentityexternalID int32, sBillingentityexternalDescription string, sEzmaxinvoicingsummaryexternalDescription string, ) *EzmaxinvoicingsummaryexternalResponse`

NewEzmaxinvoicingsummaryexternalResponse instantiates a new EzmaxinvoicingsummaryexternalResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzmaxinvoicingsummaryexternalResponseWithDefaults

`func NewEzmaxinvoicingsummaryexternalResponseWithDefaults() *EzmaxinvoicingsummaryexternalResponse`

NewEzmaxinvoicingsummaryexternalResponseWithDefaults instantiates a new EzmaxinvoicingsummaryexternalResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzmaxinvoicingsummaryexternalID

`func (o *EzmaxinvoicingsummaryexternalResponse) GetPkiEzmaxinvoicingsummaryexternalID() int32`

GetPkiEzmaxinvoicingsummaryexternalID returns the PkiEzmaxinvoicingsummaryexternalID field if non-nil, zero value otherwise.

### GetPkiEzmaxinvoicingsummaryexternalIDOk

`func (o *EzmaxinvoicingsummaryexternalResponse) GetPkiEzmaxinvoicingsummaryexternalIDOk() (*int32, bool)`

GetPkiEzmaxinvoicingsummaryexternalIDOk returns a tuple with the PkiEzmaxinvoicingsummaryexternalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzmaxinvoicingsummaryexternalID

`func (o *EzmaxinvoicingsummaryexternalResponse) SetPkiEzmaxinvoicingsummaryexternalID(v int32)`

SetPkiEzmaxinvoicingsummaryexternalID sets PkiEzmaxinvoicingsummaryexternalID field to given value.

### HasPkiEzmaxinvoicingsummaryexternalID

`func (o *EzmaxinvoicingsummaryexternalResponse) HasPkiEzmaxinvoicingsummaryexternalID() bool`

HasPkiEzmaxinvoicingsummaryexternalID returns a boolean if a field has been set.

### GetFkiEzmaxinvoicingID

`func (o *EzmaxinvoicingsummaryexternalResponse) GetFkiEzmaxinvoicingID() int32`

GetFkiEzmaxinvoicingID returns the FkiEzmaxinvoicingID field if non-nil, zero value otherwise.

### GetFkiEzmaxinvoicingIDOk

`func (o *EzmaxinvoicingsummaryexternalResponse) GetFkiEzmaxinvoicingIDOk() (*int32, bool)`

GetFkiEzmaxinvoicingIDOk returns a tuple with the FkiEzmaxinvoicingID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzmaxinvoicingID

`func (o *EzmaxinvoicingsummaryexternalResponse) SetFkiEzmaxinvoicingID(v int32)`

SetFkiEzmaxinvoicingID sets FkiEzmaxinvoicingID field to given value.

### HasFkiEzmaxinvoicingID

`func (o *EzmaxinvoicingsummaryexternalResponse) HasFkiEzmaxinvoicingID() bool`

HasFkiEzmaxinvoicingID returns a boolean if a field has been set.

### GetFkiBillingentityexternalID

`func (o *EzmaxinvoicingsummaryexternalResponse) GetFkiBillingentityexternalID() int32`

GetFkiBillingentityexternalID returns the FkiBillingentityexternalID field if non-nil, zero value otherwise.

### GetFkiBillingentityexternalIDOk

`func (o *EzmaxinvoicingsummaryexternalResponse) GetFkiBillingentityexternalIDOk() (*int32, bool)`

GetFkiBillingentityexternalIDOk returns a tuple with the FkiBillingentityexternalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiBillingentityexternalID

`func (o *EzmaxinvoicingsummaryexternalResponse) SetFkiBillingentityexternalID(v int32)`

SetFkiBillingentityexternalID sets FkiBillingentityexternalID field to given value.


### GetSBillingentityexternalDescription

`func (o *EzmaxinvoicingsummaryexternalResponse) GetSBillingentityexternalDescription() string`

GetSBillingentityexternalDescription returns the SBillingentityexternalDescription field if non-nil, zero value otherwise.

### GetSBillingentityexternalDescriptionOk

`func (o *EzmaxinvoicingsummaryexternalResponse) GetSBillingentityexternalDescriptionOk() (*string, bool)`

GetSBillingentityexternalDescriptionOk returns a tuple with the SBillingentityexternalDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSBillingentityexternalDescription

`func (o *EzmaxinvoicingsummaryexternalResponse) SetSBillingentityexternalDescription(v string)`

SetSBillingentityexternalDescription sets SBillingentityexternalDescription field to given value.


### GetSEzmaxinvoicingsummaryexternalDescription

`func (o *EzmaxinvoicingsummaryexternalResponse) GetSEzmaxinvoicingsummaryexternalDescription() string`

GetSEzmaxinvoicingsummaryexternalDescription returns the SEzmaxinvoicingsummaryexternalDescription field if non-nil, zero value otherwise.

### GetSEzmaxinvoicingsummaryexternalDescriptionOk

`func (o *EzmaxinvoicingsummaryexternalResponse) GetSEzmaxinvoicingsummaryexternalDescriptionOk() (*string, bool)`

GetSEzmaxinvoicingsummaryexternalDescriptionOk returns a tuple with the SEzmaxinvoicingsummaryexternalDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzmaxinvoicingsummaryexternalDescription

`func (o *EzmaxinvoicingsummaryexternalResponse) SetSEzmaxinvoicingsummaryexternalDescription(v string)`

SetSEzmaxinvoicingsummaryexternalDescription sets SEzmaxinvoicingsummaryexternalDescription field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


