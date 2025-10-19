# OtherincomeListElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiOtherincomeID** | **int32** | The unique ID of the Otherincome | 
**FkiOtherincometypeID** | **int32** | The unique ID of the Otherincometype | 
**SOtherincometypeDescriptionX** | **string** | The description of the Otherincometype in the language of the requester | 
**SOtherincomeDescription** | **string** | The description of the Otherincome | 
**EOtherincomeRemunerationtype** | [**FieldEOtherincomeRemunerationtype**](FieldEOtherincomeRemunerationtype.md) |  | 
**DOtherincomeRemunerationsubtotal** | **string** | The remuneration subtotal of the Otherincome | 
**DOtherincomeRemunerationtaxes** | Pointer to **string** | The remuneration total taxes of the Otherincome | [optional] 
**DOtherincomeRemunerationtotal** | Pointer to **string** | The remuneration total of the Otherincome | [optional] 
**DtOtherincomePaid** | **string** | The paid of the Otherincome | 
**BOtherincomeIsactive** | **bool** | Whether the otherincome is active or not | 

## Methods

### NewOtherincomeListElement

`func NewOtherincomeListElement(pkiOtherincomeID int32, fkiOtherincometypeID int32, sOtherincometypeDescriptionX string, sOtherincomeDescription string, eOtherincomeRemunerationtype FieldEOtherincomeRemunerationtype, dOtherincomeRemunerationsubtotal string, dtOtherincomePaid string, bOtherincomeIsactive bool, ) *OtherincomeListElement`

NewOtherincomeListElement instantiates a new OtherincomeListElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOtherincomeListElementWithDefaults

`func NewOtherincomeListElementWithDefaults() *OtherincomeListElement`

NewOtherincomeListElementWithDefaults instantiates a new OtherincomeListElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiOtherincomeID

`func (o *OtherincomeListElement) GetPkiOtherincomeID() int32`

GetPkiOtherincomeID returns the PkiOtherincomeID field if non-nil, zero value otherwise.

### GetPkiOtherincomeIDOk

`func (o *OtherincomeListElement) GetPkiOtherincomeIDOk() (*int32, bool)`

GetPkiOtherincomeIDOk returns a tuple with the PkiOtherincomeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiOtherincomeID

`func (o *OtherincomeListElement) SetPkiOtherincomeID(v int32)`

SetPkiOtherincomeID sets PkiOtherincomeID field to given value.


### GetFkiOtherincometypeID

`func (o *OtherincomeListElement) GetFkiOtherincometypeID() int32`

GetFkiOtherincometypeID returns the FkiOtherincometypeID field if non-nil, zero value otherwise.

### GetFkiOtherincometypeIDOk

`func (o *OtherincomeListElement) GetFkiOtherincometypeIDOk() (*int32, bool)`

GetFkiOtherincometypeIDOk returns a tuple with the FkiOtherincometypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiOtherincometypeID

`func (o *OtherincomeListElement) SetFkiOtherincometypeID(v int32)`

SetFkiOtherincometypeID sets FkiOtherincometypeID field to given value.


### GetSOtherincometypeDescriptionX

`func (o *OtherincomeListElement) GetSOtherincometypeDescriptionX() string`

GetSOtherincometypeDescriptionX returns the SOtherincometypeDescriptionX field if non-nil, zero value otherwise.

### GetSOtherincometypeDescriptionXOk

`func (o *OtherincomeListElement) GetSOtherincometypeDescriptionXOk() (*string, bool)`

GetSOtherincometypeDescriptionXOk returns a tuple with the SOtherincometypeDescriptionX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSOtherincometypeDescriptionX

`func (o *OtherincomeListElement) SetSOtherincometypeDescriptionX(v string)`

SetSOtherincometypeDescriptionX sets SOtherincometypeDescriptionX field to given value.


### GetSOtherincomeDescription

`func (o *OtherincomeListElement) GetSOtherincomeDescription() string`

GetSOtherincomeDescription returns the SOtherincomeDescription field if non-nil, zero value otherwise.

### GetSOtherincomeDescriptionOk

`func (o *OtherincomeListElement) GetSOtherincomeDescriptionOk() (*string, bool)`

GetSOtherincomeDescriptionOk returns a tuple with the SOtherincomeDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSOtherincomeDescription

`func (o *OtherincomeListElement) SetSOtherincomeDescription(v string)`

SetSOtherincomeDescription sets SOtherincomeDescription field to given value.


### GetEOtherincomeRemunerationtype

`func (o *OtherincomeListElement) GetEOtherincomeRemunerationtype() FieldEOtherincomeRemunerationtype`

GetEOtherincomeRemunerationtype returns the EOtherincomeRemunerationtype field if non-nil, zero value otherwise.

### GetEOtherincomeRemunerationtypeOk

`func (o *OtherincomeListElement) GetEOtherincomeRemunerationtypeOk() (*FieldEOtherincomeRemunerationtype, bool)`

GetEOtherincomeRemunerationtypeOk returns a tuple with the EOtherincomeRemunerationtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEOtherincomeRemunerationtype

`func (o *OtherincomeListElement) SetEOtherincomeRemunerationtype(v FieldEOtherincomeRemunerationtype)`

SetEOtherincomeRemunerationtype sets EOtherincomeRemunerationtype field to given value.


### GetDOtherincomeRemunerationsubtotal

`func (o *OtherincomeListElement) GetDOtherincomeRemunerationsubtotal() string`

GetDOtherincomeRemunerationsubtotal returns the DOtherincomeRemunerationsubtotal field if non-nil, zero value otherwise.

### GetDOtherincomeRemunerationsubtotalOk

`func (o *OtherincomeListElement) GetDOtherincomeRemunerationsubtotalOk() (*string, bool)`

GetDOtherincomeRemunerationsubtotalOk returns a tuple with the DOtherincomeRemunerationsubtotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDOtherincomeRemunerationsubtotal

`func (o *OtherincomeListElement) SetDOtherincomeRemunerationsubtotal(v string)`

SetDOtherincomeRemunerationsubtotal sets DOtherincomeRemunerationsubtotal field to given value.


### GetDOtherincomeRemunerationtaxes

`func (o *OtherincomeListElement) GetDOtherincomeRemunerationtaxes() string`

GetDOtherincomeRemunerationtaxes returns the DOtherincomeRemunerationtaxes field if non-nil, zero value otherwise.

### GetDOtherincomeRemunerationtaxesOk

`func (o *OtherincomeListElement) GetDOtherincomeRemunerationtaxesOk() (*string, bool)`

GetDOtherincomeRemunerationtaxesOk returns a tuple with the DOtherincomeRemunerationtaxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDOtherincomeRemunerationtaxes

`func (o *OtherincomeListElement) SetDOtherincomeRemunerationtaxes(v string)`

SetDOtherincomeRemunerationtaxes sets DOtherincomeRemunerationtaxes field to given value.

### HasDOtherincomeRemunerationtaxes

`func (o *OtherincomeListElement) HasDOtherincomeRemunerationtaxes() bool`

HasDOtherincomeRemunerationtaxes returns a boolean if a field has been set.

### GetDOtherincomeRemunerationtotal

`func (o *OtherincomeListElement) GetDOtherincomeRemunerationtotal() string`

GetDOtherincomeRemunerationtotal returns the DOtherincomeRemunerationtotal field if non-nil, zero value otherwise.

### GetDOtherincomeRemunerationtotalOk

`func (o *OtherincomeListElement) GetDOtherincomeRemunerationtotalOk() (*string, bool)`

GetDOtherincomeRemunerationtotalOk returns a tuple with the DOtherincomeRemunerationtotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDOtherincomeRemunerationtotal

`func (o *OtherincomeListElement) SetDOtherincomeRemunerationtotal(v string)`

SetDOtherincomeRemunerationtotal sets DOtherincomeRemunerationtotal field to given value.

### HasDOtherincomeRemunerationtotal

`func (o *OtherincomeListElement) HasDOtherincomeRemunerationtotal() bool`

HasDOtherincomeRemunerationtotal returns a boolean if a field has been set.

### GetDtOtherincomePaid

`func (o *OtherincomeListElement) GetDtOtherincomePaid() string`

GetDtOtherincomePaid returns the DtOtherincomePaid field if non-nil, zero value otherwise.

### GetDtOtherincomePaidOk

`func (o *OtherincomeListElement) GetDtOtherincomePaidOk() (*string, bool)`

GetDtOtherincomePaidOk returns a tuple with the DtOtherincomePaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtOtherincomePaid

`func (o *OtherincomeListElement) SetDtOtherincomePaid(v string)`

SetDtOtherincomePaid sets DtOtherincomePaid field to given value.


### GetBOtherincomeIsactive

`func (o *OtherincomeListElement) GetBOtherincomeIsactive() bool`

GetBOtherincomeIsactive returns the BOtherincomeIsactive field if non-nil, zero value otherwise.

### GetBOtherincomeIsactiveOk

`func (o *OtherincomeListElement) GetBOtherincomeIsactiveOk() (*bool, bool)`

GetBOtherincomeIsactiveOk returns a tuple with the BOtherincomeIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBOtherincomeIsactive

`func (o *OtherincomeListElement) SetBOtherincomeIsactive(v bool)`

SetBOtherincomeIsactive sets BOtherincomeIsactive field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


