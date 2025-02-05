# CurrencyAutocompleteElementResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiCurrencyID** | **int32** | The unique ID of the Currency. | 
**SCurrencyDescriptionX** | **string** | The description of the Currency in the language of the requester | 
**BCurrencyIsactive** | **bool** | Whether the Currency is active or not | 

## Methods

### NewCurrencyAutocompleteElementResponse

`func NewCurrencyAutocompleteElementResponse(pkiCurrencyID int32, sCurrencyDescriptionX string, bCurrencyIsactive bool, ) *CurrencyAutocompleteElementResponse`

NewCurrencyAutocompleteElementResponse instantiates a new CurrencyAutocompleteElementResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCurrencyAutocompleteElementResponseWithDefaults

`func NewCurrencyAutocompleteElementResponseWithDefaults() *CurrencyAutocompleteElementResponse`

NewCurrencyAutocompleteElementResponseWithDefaults instantiates a new CurrencyAutocompleteElementResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiCurrencyID

`func (o *CurrencyAutocompleteElementResponse) GetPkiCurrencyID() int32`

GetPkiCurrencyID returns the PkiCurrencyID field if non-nil, zero value otherwise.

### GetPkiCurrencyIDOk

`func (o *CurrencyAutocompleteElementResponse) GetPkiCurrencyIDOk() (*int32, bool)`

GetPkiCurrencyIDOk returns a tuple with the PkiCurrencyID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiCurrencyID

`func (o *CurrencyAutocompleteElementResponse) SetPkiCurrencyID(v int32)`

SetPkiCurrencyID sets PkiCurrencyID field to given value.


### GetSCurrencyDescriptionX

`func (o *CurrencyAutocompleteElementResponse) GetSCurrencyDescriptionX() string`

GetSCurrencyDescriptionX returns the SCurrencyDescriptionX field if non-nil, zero value otherwise.

### GetSCurrencyDescriptionXOk

`func (o *CurrencyAutocompleteElementResponse) GetSCurrencyDescriptionXOk() (*string, bool)`

GetSCurrencyDescriptionXOk returns a tuple with the SCurrencyDescriptionX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSCurrencyDescriptionX

`func (o *CurrencyAutocompleteElementResponse) SetSCurrencyDescriptionX(v string)`

SetSCurrencyDescriptionX sets SCurrencyDescriptionX field to given value.


### GetBCurrencyIsactive

`func (o *CurrencyAutocompleteElementResponse) GetBCurrencyIsactive() bool`

GetBCurrencyIsactive returns the BCurrencyIsactive field if non-nil, zero value otherwise.

### GetBCurrencyIsactiveOk

`func (o *CurrencyAutocompleteElementResponse) GetBCurrencyIsactiveOk() (*bool, bool)`

GetBCurrencyIsactiveOk returns a tuple with the BCurrencyIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBCurrencyIsactive

`func (o *CurrencyAutocompleteElementResponse) SetBCurrencyIsactive(v bool)`

SetBCurrencyIsactive sets BCurrencyIsactive field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


