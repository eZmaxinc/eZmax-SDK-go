# BuyercontractResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiBuyercontractID** | **int32** | The unique ID of the Buyercontract | 
**FkiInscriptiontypeID** | **int32** | The unique ID of the Inscriptiontype | 
**EBuyercontractStep** | [**FieldEBuyercontractStep**](FieldEBuyercontractStep.md) |  | 
**DBuyercontractMinimumprice** | **string** | The minimum price of the Buyercontract | 
**DBuyercontractMaximumprice** | **string** | The maximum price of the Buyercontract | 
**EBuyercontractType** | [**FieldEBuyercontractType**](FieldEBuyercontractType.md) |  | 
**SBuyercontractContract** | Pointer to **string** | The number of the Buyercontract | [optional] 
**DtBuyercontractDate** | **string** | The date of the Buyercontract | 
**DtBuyercontractExpirationdate** | Pointer to **string** | The expiration date of the Buyercontract | [optional] 
**DBuyercontractRemuneration** | Pointer to **string** | The remuneration of the Buyercontract | [optional] 
**EBuyercontractRemunerationtype** | Pointer to [**FieldEBuyercontractRemunerationtype**](FieldEBuyercontractRemunerationtype.md) |  | [optional] 
**BBuyercontractLitigation** | Pointer to **bool** | Whether if it&#39;s an litigation | [optional] 
**BBuyercontractIsactive** | **bool** | Whether the buyercontract is active or not | 

## Methods

### NewBuyercontractResponse

`func NewBuyercontractResponse(pkiBuyercontractID int32, fkiInscriptiontypeID int32, eBuyercontractStep FieldEBuyercontractStep, dBuyercontractMinimumprice string, dBuyercontractMaximumprice string, eBuyercontractType FieldEBuyercontractType, dtBuyercontractDate string, bBuyercontractIsactive bool, ) *BuyercontractResponse`

NewBuyercontractResponse instantiates a new BuyercontractResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuyercontractResponseWithDefaults

`func NewBuyercontractResponseWithDefaults() *BuyercontractResponse`

NewBuyercontractResponseWithDefaults instantiates a new BuyercontractResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiBuyercontractID

`func (o *BuyercontractResponse) GetPkiBuyercontractID() int32`

GetPkiBuyercontractID returns the PkiBuyercontractID field if non-nil, zero value otherwise.

### GetPkiBuyercontractIDOk

`func (o *BuyercontractResponse) GetPkiBuyercontractIDOk() (*int32, bool)`

GetPkiBuyercontractIDOk returns a tuple with the PkiBuyercontractID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiBuyercontractID

`func (o *BuyercontractResponse) SetPkiBuyercontractID(v int32)`

SetPkiBuyercontractID sets PkiBuyercontractID field to given value.


### GetFkiInscriptiontypeID

`func (o *BuyercontractResponse) GetFkiInscriptiontypeID() int32`

GetFkiInscriptiontypeID returns the FkiInscriptiontypeID field if non-nil, zero value otherwise.

### GetFkiInscriptiontypeIDOk

`func (o *BuyercontractResponse) GetFkiInscriptiontypeIDOk() (*int32, bool)`

GetFkiInscriptiontypeIDOk returns a tuple with the FkiInscriptiontypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiInscriptiontypeID

`func (o *BuyercontractResponse) SetFkiInscriptiontypeID(v int32)`

SetFkiInscriptiontypeID sets FkiInscriptiontypeID field to given value.


### GetEBuyercontractStep

`func (o *BuyercontractResponse) GetEBuyercontractStep() FieldEBuyercontractStep`

GetEBuyercontractStep returns the EBuyercontractStep field if non-nil, zero value otherwise.

### GetEBuyercontractStepOk

`func (o *BuyercontractResponse) GetEBuyercontractStepOk() (*FieldEBuyercontractStep, bool)`

GetEBuyercontractStepOk returns a tuple with the EBuyercontractStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEBuyercontractStep

`func (o *BuyercontractResponse) SetEBuyercontractStep(v FieldEBuyercontractStep)`

SetEBuyercontractStep sets EBuyercontractStep field to given value.


### GetDBuyercontractMinimumprice

`func (o *BuyercontractResponse) GetDBuyercontractMinimumprice() string`

GetDBuyercontractMinimumprice returns the DBuyercontractMinimumprice field if non-nil, zero value otherwise.

### GetDBuyercontractMinimumpriceOk

`func (o *BuyercontractResponse) GetDBuyercontractMinimumpriceOk() (*string, bool)`

GetDBuyercontractMinimumpriceOk returns a tuple with the DBuyercontractMinimumprice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDBuyercontractMinimumprice

`func (o *BuyercontractResponse) SetDBuyercontractMinimumprice(v string)`

SetDBuyercontractMinimumprice sets DBuyercontractMinimumprice field to given value.


### GetDBuyercontractMaximumprice

`func (o *BuyercontractResponse) GetDBuyercontractMaximumprice() string`

GetDBuyercontractMaximumprice returns the DBuyercontractMaximumprice field if non-nil, zero value otherwise.

### GetDBuyercontractMaximumpriceOk

`func (o *BuyercontractResponse) GetDBuyercontractMaximumpriceOk() (*string, bool)`

GetDBuyercontractMaximumpriceOk returns a tuple with the DBuyercontractMaximumprice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDBuyercontractMaximumprice

`func (o *BuyercontractResponse) SetDBuyercontractMaximumprice(v string)`

SetDBuyercontractMaximumprice sets DBuyercontractMaximumprice field to given value.


### GetEBuyercontractType

`func (o *BuyercontractResponse) GetEBuyercontractType() FieldEBuyercontractType`

GetEBuyercontractType returns the EBuyercontractType field if non-nil, zero value otherwise.

### GetEBuyercontractTypeOk

`func (o *BuyercontractResponse) GetEBuyercontractTypeOk() (*FieldEBuyercontractType, bool)`

GetEBuyercontractTypeOk returns a tuple with the EBuyercontractType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEBuyercontractType

`func (o *BuyercontractResponse) SetEBuyercontractType(v FieldEBuyercontractType)`

SetEBuyercontractType sets EBuyercontractType field to given value.


### GetSBuyercontractContract

`func (o *BuyercontractResponse) GetSBuyercontractContract() string`

GetSBuyercontractContract returns the SBuyercontractContract field if non-nil, zero value otherwise.

### GetSBuyercontractContractOk

`func (o *BuyercontractResponse) GetSBuyercontractContractOk() (*string, bool)`

GetSBuyercontractContractOk returns a tuple with the SBuyercontractContract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSBuyercontractContract

`func (o *BuyercontractResponse) SetSBuyercontractContract(v string)`

SetSBuyercontractContract sets SBuyercontractContract field to given value.

### HasSBuyercontractContract

`func (o *BuyercontractResponse) HasSBuyercontractContract() bool`

HasSBuyercontractContract returns a boolean if a field has been set.

### GetDtBuyercontractDate

`func (o *BuyercontractResponse) GetDtBuyercontractDate() string`

GetDtBuyercontractDate returns the DtBuyercontractDate field if non-nil, zero value otherwise.

### GetDtBuyercontractDateOk

`func (o *BuyercontractResponse) GetDtBuyercontractDateOk() (*string, bool)`

GetDtBuyercontractDateOk returns a tuple with the DtBuyercontractDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtBuyercontractDate

`func (o *BuyercontractResponse) SetDtBuyercontractDate(v string)`

SetDtBuyercontractDate sets DtBuyercontractDate field to given value.


### GetDtBuyercontractExpirationdate

`func (o *BuyercontractResponse) GetDtBuyercontractExpirationdate() string`

GetDtBuyercontractExpirationdate returns the DtBuyercontractExpirationdate field if non-nil, zero value otherwise.

### GetDtBuyercontractExpirationdateOk

`func (o *BuyercontractResponse) GetDtBuyercontractExpirationdateOk() (*string, bool)`

GetDtBuyercontractExpirationdateOk returns a tuple with the DtBuyercontractExpirationdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtBuyercontractExpirationdate

`func (o *BuyercontractResponse) SetDtBuyercontractExpirationdate(v string)`

SetDtBuyercontractExpirationdate sets DtBuyercontractExpirationdate field to given value.

### HasDtBuyercontractExpirationdate

`func (o *BuyercontractResponse) HasDtBuyercontractExpirationdate() bool`

HasDtBuyercontractExpirationdate returns a boolean if a field has been set.

### GetDBuyercontractRemuneration

`func (o *BuyercontractResponse) GetDBuyercontractRemuneration() string`

GetDBuyercontractRemuneration returns the DBuyercontractRemuneration field if non-nil, zero value otherwise.

### GetDBuyercontractRemunerationOk

`func (o *BuyercontractResponse) GetDBuyercontractRemunerationOk() (*string, bool)`

GetDBuyercontractRemunerationOk returns a tuple with the DBuyercontractRemuneration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDBuyercontractRemuneration

`func (o *BuyercontractResponse) SetDBuyercontractRemuneration(v string)`

SetDBuyercontractRemuneration sets DBuyercontractRemuneration field to given value.

### HasDBuyercontractRemuneration

`func (o *BuyercontractResponse) HasDBuyercontractRemuneration() bool`

HasDBuyercontractRemuneration returns a boolean if a field has been set.

### GetEBuyercontractRemunerationtype

`func (o *BuyercontractResponse) GetEBuyercontractRemunerationtype() FieldEBuyercontractRemunerationtype`

GetEBuyercontractRemunerationtype returns the EBuyercontractRemunerationtype field if non-nil, zero value otherwise.

### GetEBuyercontractRemunerationtypeOk

`func (o *BuyercontractResponse) GetEBuyercontractRemunerationtypeOk() (*FieldEBuyercontractRemunerationtype, bool)`

GetEBuyercontractRemunerationtypeOk returns a tuple with the EBuyercontractRemunerationtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEBuyercontractRemunerationtype

`func (o *BuyercontractResponse) SetEBuyercontractRemunerationtype(v FieldEBuyercontractRemunerationtype)`

SetEBuyercontractRemunerationtype sets EBuyercontractRemunerationtype field to given value.

### HasEBuyercontractRemunerationtype

`func (o *BuyercontractResponse) HasEBuyercontractRemunerationtype() bool`

HasEBuyercontractRemunerationtype returns a boolean if a field has been set.

### GetBBuyercontractLitigation

`func (o *BuyercontractResponse) GetBBuyercontractLitigation() bool`

GetBBuyercontractLitigation returns the BBuyercontractLitigation field if non-nil, zero value otherwise.

### GetBBuyercontractLitigationOk

`func (o *BuyercontractResponse) GetBBuyercontractLitigationOk() (*bool, bool)`

GetBBuyercontractLitigationOk returns a tuple with the BBuyercontractLitigation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBBuyercontractLitigation

`func (o *BuyercontractResponse) SetBBuyercontractLitigation(v bool)`

SetBBuyercontractLitigation sets BBuyercontractLitigation field to given value.

### HasBBuyercontractLitigation

`func (o *BuyercontractResponse) HasBBuyercontractLitigation() bool`

HasBBuyercontractLitigation returns a boolean if a field has been set.

### GetBBuyercontractIsactive

`func (o *BuyercontractResponse) GetBBuyercontractIsactive() bool`

GetBBuyercontractIsactive returns the BBuyercontractIsactive field if non-nil, zero value otherwise.

### GetBBuyercontractIsactiveOk

`func (o *BuyercontractResponse) GetBBuyercontractIsactiveOk() (*bool, bool)`

GetBBuyercontractIsactiveOk returns a tuple with the BBuyercontractIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBBuyercontractIsactive

`func (o *BuyercontractResponse) SetBBuyercontractIsactive(v bool)`

SetBBuyercontractIsactive sets BBuyercontractIsactive field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


