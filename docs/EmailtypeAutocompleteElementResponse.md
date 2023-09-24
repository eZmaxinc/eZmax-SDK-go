# EmailtypeAutocompleteElementResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEmailtypeID** | **int32** | The unique ID of the Emailtype.  Valid values:  |Value|Description| |-|-| |1|Office| |2|Home| | 
**SEmailtypeNameX** | **string** | The name of the Emailtype in the language of the requester | 
**BEmailtypeIsactive** | **bool** | Whether the Emailtype is active or not | 

## Methods

### NewEmailtypeAutocompleteElementResponse

`func NewEmailtypeAutocompleteElementResponse(pkiEmailtypeID int32, sEmailtypeNameX string, bEmailtypeIsactive bool, ) *EmailtypeAutocompleteElementResponse`

NewEmailtypeAutocompleteElementResponse instantiates a new EmailtypeAutocompleteElementResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailtypeAutocompleteElementResponseWithDefaults

`func NewEmailtypeAutocompleteElementResponseWithDefaults() *EmailtypeAutocompleteElementResponse`

NewEmailtypeAutocompleteElementResponseWithDefaults instantiates a new EmailtypeAutocompleteElementResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEmailtypeID

`func (o *EmailtypeAutocompleteElementResponse) GetPkiEmailtypeID() int32`

GetPkiEmailtypeID returns the PkiEmailtypeID field if non-nil, zero value otherwise.

### GetPkiEmailtypeIDOk

`func (o *EmailtypeAutocompleteElementResponse) GetPkiEmailtypeIDOk() (*int32, bool)`

GetPkiEmailtypeIDOk returns a tuple with the PkiEmailtypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEmailtypeID

`func (o *EmailtypeAutocompleteElementResponse) SetPkiEmailtypeID(v int32)`

SetPkiEmailtypeID sets PkiEmailtypeID field to given value.


### GetSEmailtypeNameX

`func (o *EmailtypeAutocompleteElementResponse) GetSEmailtypeNameX() string`

GetSEmailtypeNameX returns the SEmailtypeNameX field if non-nil, zero value otherwise.

### GetSEmailtypeNameXOk

`func (o *EmailtypeAutocompleteElementResponse) GetSEmailtypeNameXOk() (*string, bool)`

GetSEmailtypeNameXOk returns a tuple with the SEmailtypeNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEmailtypeNameX

`func (o *EmailtypeAutocompleteElementResponse) SetSEmailtypeNameX(v string)`

SetSEmailtypeNameX sets SEmailtypeNameX field to given value.


### GetBEmailtypeIsactive

`func (o *EmailtypeAutocompleteElementResponse) GetBEmailtypeIsactive() bool`

GetBEmailtypeIsactive returns the BEmailtypeIsactive field if non-nil, zero value otherwise.

### GetBEmailtypeIsactiveOk

`func (o *EmailtypeAutocompleteElementResponse) GetBEmailtypeIsactiveOk() (*bool, bool)`

GetBEmailtypeIsactiveOk returns a tuple with the BEmailtypeIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEmailtypeIsactive

`func (o *EmailtypeAutocompleteElementResponse) SetBEmailtypeIsactive(v bool)`

SetBEmailtypeIsactive sets BEmailtypeIsactive field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


