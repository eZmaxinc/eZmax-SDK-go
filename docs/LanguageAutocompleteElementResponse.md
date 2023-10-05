# LanguageAutocompleteElementResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiLanguageID** | **int32** | The unique ID of the Language.  Valid values:  |Value|Description| |-|-| |1|French| |2|English| | 
**SLanguageNameX** | **string** | The Name of the Language in the language of the requester | 
**BLanguageIsactive** | **bool** | Whether the Language is active or not | 

## Methods

### NewLanguageAutocompleteElementResponse

`func NewLanguageAutocompleteElementResponse(pkiLanguageID int32, sLanguageNameX string, bLanguageIsactive bool, ) *LanguageAutocompleteElementResponse`

NewLanguageAutocompleteElementResponse instantiates a new LanguageAutocompleteElementResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLanguageAutocompleteElementResponseWithDefaults

`func NewLanguageAutocompleteElementResponseWithDefaults() *LanguageAutocompleteElementResponse`

NewLanguageAutocompleteElementResponseWithDefaults instantiates a new LanguageAutocompleteElementResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiLanguageID

`func (o *LanguageAutocompleteElementResponse) GetPkiLanguageID() int32`

GetPkiLanguageID returns the PkiLanguageID field if non-nil, zero value otherwise.

### GetPkiLanguageIDOk

`func (o *LanguageAutocompleteElementResponse) GetPkiLanguageIDOk() (*int32, bool)`

GetPkiLanguageIDOk returns a tuple with the PkiLanguageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiLanguageID

`func (o *LanguageAutocompleteElementResponse) SetPkiLanguageID(v int32)`

SetPkiLanguageID sets PkiLanguageID field to given value.


### GetSLanguageNameX

`func (o *LanguageAutocompleteElementResponse) GetSLanguageNameX() string`

GetSLanguageNameX returns the SLanguageNameX field if non-nil, zero value otherwise.

### GetSLanguageNameXOk

`func (o *LanguageAutocompleteElementResponse) GetSLanguageNameXOk() (*string, bool)`

GetSLanguageNameXOk returns a tuple with the SLanguageNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSLanguageNameX

`func (o *LanguageAutocompleteElementResponse) SetSLanguageNameX(v string)`

SetSLanguageNameX sets SLanguageNameX field to given value.


### GetBLanguageIsactive

`func (o *LanguageAutocompleteElementResponse) GetBLanguageIsactive() bool`

GetBLanguageIsactive returns the BLanguageIsactive field if non-nil, zero value otherwise.

### GetBLanguageIsactiveOk

`func (o *LanguageAutocompleteElementResponse) GetBLanguageIsactiveOk() (*bool, bool)`

GetBLanguageIsactiveOk returns a tuple with the BLanguageIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBLanguageIsactive

`func (o *LanguageAutocompleteElementResponse) SetBLanguageIsactive(v bool)`

SetBLanguageIsactive sets BLanguageIsactive field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


