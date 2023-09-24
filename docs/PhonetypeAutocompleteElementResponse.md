# PhonetypeAutocompleteElementResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiPhonetypeID** | **int32** | The unique ID of the Phonetype.  Valid values:  |Value|Description| |-|-| |1|Office| |2|Home| |3|Mobile| |4|Fax| |5|Pager| |6|Toll Free| | 
**SPhonetypeNameX** | **string** | The name of the Phonetype in the language of the requester | 
**BPhonetypeIsactive** | **bool** | Whether the Phonetype is active or not | 

## Methods

### NewPhonetypeAutocompleteElementResponse

`func NewPhonetypeAutocompleteElementResponse(pkiPhonetypeID int32, sPhonetypeNameX string, bPhonetypeIsactive bool, ) *PhonetypeAutocompleteElementResponse`

NewPhonetypeAutocompleteElementResponse instantiates a new PhonetypeAutocompleteElementResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhonetypeAutocompleteElementResponseWithDefaults

`func NewPhonetypeAutocompleteElementResponseWithDefaults() *PhonetypeAutocompleteElementResponse`

NewPhonetypeAutocompleteElementResponseWithDefaults instantiates a new PhonetypeAutocompleteElementResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiPhonetypeID

`func (o *PhonetypeAutocompleteElementResponse) GetPkiPhonetypeID() int32`

GetPkiPhonetypeID returns the PkiPhonetypeID field if non-nil, zero value otherwise.

### GetPkiPhonetypeIDOk

`func (o *PhonetypeAutocompleteElementResponse) GetPkiPhonetypeIDOk() (*int32, bool)`

GetPkiPhonetypeIDOk returns a tuple with the PkiPhonetypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiPhonetypeID

`func (o *PhonetypeAutocompleteElementResponse) SetPkiPhonetypeID(v int32)`

SetPkiPhonetypeID sets PkiPhonetypeID field to given value.


### GetSPhonetypeNameX

`func (o *PhonetypeAutocompleteElementResponse) GetSPhonetypeNameX() string`

GetSPhonetypeNameX returns the SPhonetypeNameX field if non-nil, zero value otherwise.

### GetSPhonetypeNameXOk

`func (o *PhonetypeAutocompleteElementResponse) GetSPhonetypeNameXOk() (*string, bool)`

GetSPhonetypeNameXOk returns a tuple with the SPhonetypeNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSPhonetypeNameX

`func (o *PhonetypeAutocompleteElementResponse) SetSPhonetypeNameX(v string)`

SetSPhonetypeNameX sets SPhonetypeNameX field to given value.


### GetBPhonetypeIsactive

`func (o *PhonetypeAutocompleteElementResponse) GetBPhonetypeIsactive() bool`

GetBPhonetypeIsactive returns the BPhonetypeIsactive field if non-nil, zero value otherwise.

### GetBPhonetypeIsactiveOk

`func (o *PhonetypeAutocompleteElementResponse) GetBPhonetypeIsactiveOk() (*bool, bool)`

GetBPhonetypeIsactiveOk returns a tuple with the BPhonetypeIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBPhonetypeIsactive

`func (o *PhonetypeAutocompleteElementResponse) SetBPhonetypeIsactive(v bool)`

SetBPhonetypeIsactive sets BPhonetypeIsactive field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


