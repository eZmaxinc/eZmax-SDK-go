# SupplyAutocompleteElementResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiSupplyID** | **int32** | The unique ID of the Supply | 
**SSupplyDescriptionX** | **string** | The description of the Supply in the language of the requester | 
**BSupplyIsactive** | **bool** | Whether the supply is active or not | 

## Methods

### NewSupplyAutocompleteElementResponse

`func NewSupplyAutocompleteElementResponse(pkiSupplyID int32, sSupplyDescriptionX string, bSupplyIsactive bool, ) *SupplyAutocompleteElementResponse`

NewSupplyAutocompleteElementResponse instantiates a new SupplyAutocompleteElementResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupplyAutocompleteElementResponseWithDefaults

`func NewSupplyAutocompleteElementResponseWithDefaults() *SupplyAutocompleteElementResponse`

NewSupplyAutocompleteElementResponseWithDefaults instantiates a new SupplyAutocompleteElementResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiSupplyID

`func (o *SupplyAutocompleteElementResponse) GetPkiSupplyID() int32`

GetPkiSupplyID returns the PkiSupplyID field if non-nil, zero value otherwise.

### GetPkiSupplyIDOk

`func (o *SupplyAutocompleteElementResponse) GetPkiSupplyIDOk() (*int32, bool)`

GetPkiSupplyIDOk returns a tuple with the PkiSupplyID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiSupplyID

`func (o *SupplyAutocompleteElementResponse) SetPkiSupplyID(v int32)`

SetPkiSupplyID sets PkiSupplyID field to given value.


### GetSSupplyDescriptionX

`func (o *SupplyAutocompleteElementResponse) GetSSupplyDescriptionX() string`

GetSSupplyDescriptionX returns the SSupplyDescriptionX field if non-nil, zero value otherwise.

### GetSSupplyDescriptionXOk

`func (o *SupplyAutocompleteElementResponse) GetSSupplyDescriptionXOk() (*string, bool)`

GetSSupplyDescriptionXOk returns a tuple with the SSupplyDescriptionX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSSupplyDescriptionX

`func (o *SupplyAutocompleteElementResponse) SetSSupplyDescriptionX(v string)`

SetSSupplyDescriptionX sets SSupplyDescriptionX field to given value.


### GetBSupplyIsactive

`func (o *SupplyAutocompleteElementResponse) GetBSupplyIsactive() bool`

GetBSupplyIsactive returns the BSupplyIsactive field if non-nil, zero value otherwise.

### GetBSupplyIsactiveOk

`func (o *SupplyAutocompleteElementResponse) GetBSupplyIsactiveOk() (*bool, bool)`

GetBSupplyIsactiveOk returns a tuple with the BSupplyIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBSupplyIsactive

`func (o *SupplyAutocompleteElementResponse) SetBSupplyIsactive(v bool)`

SetBSupplyIsactive sets BSupplyIsactive field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


