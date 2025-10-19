# BuyercontractListElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiBuyercontractID** | **int32** | The unique ID of the Buyercontract | 
**FkiInscriptiontypeID** | **int32** | The unique ID of the Inscriptiontype | 
**SInscriptiontypeNameX** | **string** | The name of the Inscriptiontype in the language of the requester | 
**EBuyercontractStep** | [**FieldEBuyercontractStep**](FieldEBuyercontractStep.md) |  | 
**DBuyercontractMinimumprice** | **string** | The minimumprice of the Buyercontract | 
**DBuyercontractMaximumprice** | **string** | The maximumprice of the Buyercontract | 
**EBuyercontractType** | [**FieldEBuyercontractType**](FieldEBuyercontractType.md) |  | 
**DtBuyercontractDate** | **string** | The date of the Buyercontract | 
**DtBuyercontractExpirationdate** | Pointer to **string** | The expirationdate of the Buyercontract | [optional] 
**BBuyercontractIsactive** | **bool** | Whether the buyercontract is active or not | 
**SBuyercontractBrokers** | **string** | The brokers&#39; name of the Buyercontract | 
**SBuyercontractBuyers** | **string** | The buyers&#39; name of the Buyercontract | 

## Methods

### NewBuyercontractListElement

`func NewBuyercontractListElement(pkiBuyercontractID int32, fkiInscriptiontypeID int32, sInscriptiontypeNameX string, eBuyercontractStep FieldEBuyercontractStep, dBuyercontractMinimumprice string, dBuyercontractMaximumprice string, eBuyercontractType FieldEBuyercontractType, dtBuyercontractDate string, bBuyercontractIsactive bool, sBuyercontractBrokers string, sBuyercontractBuyers string, ) *BuyercontractListElement`

NewBuyercontractListElement instantiates a new BuyercontractListElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuyercontractListElementWithDefaults

`func NewBuyercontractListElementWithDefaults() *BuyercontractListElement`

NewBuyercontractListElementWithDefaults instantiates a new BuyercontractListElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiBuyercontractID

`func (o *BuyercontractListElement) GetPkiBuyercontractID() int32`

GetPkiBuyercontractID returns the PkiBuyercontractID field if non-nil, zero value otherwise.

### GetPkiBuyercontractIDOk

`func (o *BuyercontractListElement) GetPkiBuyercontractIDOk() (*int32, bool)`

GetPkiBuyercontractIDOk returns a tuple with the PkiBuyercontractID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiBuyercontractID

`func (o *BuyercontractListElement) SetPkiBuyercontractID(v int32)`

SetPkiBuyercontractID sets PkiBuyercontractID field to given value.


### GetFkiInscriptiontypeID

`func (o *BuyercontractListElement) GetFkiInscriptiontypeID() int32`

GetFkiInscriptiontypeID returns the FkiInscriptiontypeID field if non-nil, zero value otherwise.

### GetFkiInscriptiontypeIDOk

`func (o *BuyercontractListElement) GetFkiInscriptiontypeIDOk() (*int32, bool)`

GetFkiInscriptiontypeIDOk returns a tuple with the FkiInscriptiontypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiInscriptiontypeID

`func (o *BuyercontractListElement) SetFkiInscriptiontypeID(v int32)`

SetFkiInscriptiontypeID sets FkiInscriptiontypeID field to given value.


### GetSInscriptiontypeNameX

`func (o *BuyercontractListElement) GetSInscriptiontypeNameX() string`

GetSInscriptiontypeNameX returns the SInscriptiontypeNameX field if non-nil, zero value otherwise.

### GetSInscriptiontypeNameXOk

`func (o *BuyercontractListElement) GetSInscriptiontypeNameXOk() (*string, bool)`

GetSInscriptiontypeNameXOk returns a tuple with the SInscriptiontypeNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSInscriptiontypeNameX

`func (o *BuyercontractListElement) SetSInscriptiontypeNameX(v string)`

SetSInscriptiontypeNameX sets SInscriptiontypeNameX field to given value.


### GetEBuyercontractStep

`func (o *BuyercontractListElement) GetEBuyercontractStep() FieldEBuyercontractStep`

GetEBuyercontractStep returns the EBuyercontractStep field if non-nil, zero value otherwise.

### GetEBuyercontractStepOk

`func (o *BuyercontractListElement) GetEBuyercontractStepOk() (*FieldEBuyercontractStep, bool)`

GetEBuyercontractStepOk returns a tuple with the EBuyercontractStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEBuyercontractStep

`func (o *BuyercontractListElement) SetEBuyercontractStep(v FieldEBuyercontractStep)`

SetEBuyercontractStep sets EBuyercontractStep field to given value.


### GetDBuyercontractMinimumprice

`func (o *BuyercontractListElement) GetDBuyercontractMinimumprice() string`

GetDBuyercontractMinimumprice returns the DBuyercontractMinimumprice field if non-nil, zero value otherwise.

### GetDBuyercontractMinimumpriceOk

`func (o *BuyercontractListElement) GetDBuyercontractMinimumpriceOk() (*string, bool)`

GetDBuyercontractMinimumpriceOk returns a tuple with the DBuyercontractMinimumprice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDBuyercontractMinimumprice

`func (o *BuyercontractListElement) SetDBuyercontractMinimumprice(v string)`

SetDBuyercontractMinimumprice sets DBuyercontractMinimumprice field to given value.


### GetDBuyercontractMaximumprice

`func (o *BuyercontractListElement) GetDBuyercontractMaximumprice() string`

GetDBuyercontractMaximumprice returns the DBuyercontractMaximumprice field if non-nil, zero value otherwise.

### GetDBuyercontractMaximumpriceOk

`func (o *BuyercontractListElement) GetDBuyercontractMaximumpriceOk() (*string, bool)`

GetDBuyercontractMaximumpriceOk returns a tuple with the DBuyercontractMaximumprice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDBuyercontractMaximumprice

`func (o *BuyercontractListElement) SetDBuyercontractMaximumprice(v string)`

SetDBuyercontractMaximumprice sets DBuyercontractMaximumprice field to given value.


### GetEBuyercontractType

`func (o *BuyercontractListElement) GetEBuyercontractType() FieldEBuyercontractType`

GetEBuyercontractType returns the EBuyercontractType field if non-nil, zero value otherwise.

### GetEBuyercontractTypeOk

`func (o *BuyercontractListElement) GetEBuyercontractTypeOk() (*FieldEBuyercontractType, bool)`

GetEBuyercontractTypeOk returns a tuple with the EBuyercontractType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEBuyercontractType

`func (o *BuyercontractListElement) SetEBuyercontractType(v FieldEBuyercontractType)`

SetEBuyercontractType sets EBuyercontractType field to given value.


### GetDtBuyercontractDate

`func (o *BuyercontractListElement) GetDtBuyercontractDate() string`

GetDtBuyercontractDate returns the DtBuyercontractDate field if non-nil, zero value otherwise.

### GetDtBuyercontractDateOk

`func (o *BuyercontractListElement) GetDtBuyercontractDateOk() (*string, bool)`

GetDtBuyercontractDateOk returns a tuple with the DtBuyercontractDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtBuyercontractDate

`func (o *BuyercontractListElement) SetDtBuyercontractDate(v string)`

SetDtBuyercontractDate sets DtBuyercontractDate field to given value.


### GetDtBuyercontractExpirationdate

`func (o *BuyercontractListElement) GetDtBuyercontractExpirationdate() string`

GetDtBuyercontractExpirationdate returns the DtBuyercontractExpirationdate field if non-nil, zero value otherwise.

### GetDtBuyercontractExpirationdateOk

`func (o *BuyercontractListElement) GetDtBuyercontractExpirationdateOk() (*string, bool)`

GetDtBuyercontractExpirationdateOk returns a tuple with the DtBuyercontractExpirationdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtBuyercontractExpirationdate

`func (o *BuyercontractListElement) SetDtBuyercontractExpirationdate(v string)`

SetDtBuyercontractExpirationdate sets DtBuyercontractExpirationdate field to given value.

### HasDtBuyercontractExpirationdate

`func (o *BuyercontractListElement) HasDtBuyercontractExpirationdate() bool`

HasDtBuyercontractExpirationdate returns a boolean if a field has been set.

### GetBBuyercontractIsactive

`func (o *BuyercontractListElement) GetBBuyercontractIsactive() bool`

GetBBuyercontractIsactive returns the BBuyercontractIsactive field if non-nil, zero value otherwise.

### GetBBuyercontractIsactiveOk

`func (o *BuyercontractListElement) GetBBuyercontractIsactiveOk() (*bool, bool)`

GetBBuyercontractIsactiveOk returns a tuple with the BBuyercontractIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBBuyercontractIsactive

`func (o *BuyercontractListElement) SetBBuyercontractIsactive(v bool)`

SetBBuyercontractIsactive sets BBuyercontractIsactive field to given value.


### GetSBuyercontractBrokers

`func (o *BuyercontractListElement) GetSBuyercontractBrokers() string`

GetSBuyercontractBrokers returns the SBuyercontractBrokers field if non-nil, zero value otherwise.

### GetSBuyercontractBrokersOk

`func (o *BuyercontractListElement) GetSBuyercontractBrokersOk() (*string, bool)`

GetSBuyercontractBrokersOk returns a tuple with the SBuyercontractBrokers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSBuyercontractBrokers

`func (o *BuyercontractListElement) SetSBuyercontractBrokers(v string)`

SetSBuyercontractBrokers sets SBuyercontractBrokers field to given value.


### GetSBuyercontractBuyers

`func (o *BuyercontractListElement) GetSBuyercontractBuyers() string`

GetSBuyercontractBuyers returns the SBuyercontractBuyers field if non-nil, zero value otherwise.

### GetSBuyercontractBuyersOk

`func (o *BuyercontractListElement) GetSBuyercontractBuyersOk() (*string, bool)`

GetSBuyercontractBuyersOk returns a tuple with the SBuyercontractBuyers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSBuyercontractBuyers

`func (o *BuyercontractListElement) SetSBuyercontractBuyers(v string)`

SetSBuyercontractBuyers sets SBuyercontractBuyers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


