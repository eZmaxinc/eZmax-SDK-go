# ContacttitleAutocompleteElementResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiContacttitleID** | **int32** | The unique ID of the Contacttitle.  Valid values:  |Value|Description| |-|-| |1|Ms.| |2|Mr.| |4|(Blank)| |5|Me (For Notaries)| | 
**SContacttitleNameX** | **string** | The name of the Contacttitle in the language of the requester | 

## Methods

### NewContacttitleAutocompleteElementResponse

`func NewContacttitleAutocompleteElementResponse(pkiContacttitleID int32, sContacttitleNameX string, ) *ContacttitleAutocompleteElementResponse`

NewContacttitleAutocompleteElementResponse instantiates a new ContacttitleAutocompleteElementResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContacttitleAutocompleteElementResponseWithDefaults

`func NewContacttitleAutocompleteElementResponseWithDefaults() *ContacttitleAutocompleteElementResponse`

NewContacttitleAutocompleteElementResponseWithDefaults instantiates a new ContacttitleAutocompleteElementResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiContacttitleID

`func (o *ContacttitleAutocompleteElementResponse) GetPkiContacttitleID() int32`

GetPkiContacttitleID returns the PkiContacttitleID field if non-nil, zero value otherwise.

### GetPkiContacttitleIDOk

`func (o *ContacttitleAutocompleteElementResponse) GetPkiContacttitleIDOk() (*int32, bool)`

GetPkiContacttitleIDOk returns a tuple with the PkiContacttitleID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiContacttitleID

`func (o *ContacttitleAutocompleteElementResponse) SetPkiContacttitleID(v int32)`

SetPkiContacttitleID sets PkiContacttitleID field to given value.


### GetSContacttitleNameX

`func (o *ContacttitleAutocompleteElementResponse) GetSContacttitleNameX() string`

GetSContacttitleNameX returns the SContacttitleNameX field if non-nil, zero value otherwise.

### GetSContacttitleNameXOk

`func (o *ContacttitleAutocompleteElementResponse) GetSContacttitleNameXOk() (*string, bool)`

GetSContacttitleNameXOk returns a tuple with the SContacttitleNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSContacttitleNameX

`func (o *ContacttitleAutocompleteElementResponse) SetSContacttitleNameX(v string)`

SetSContacttitleNameX sets SContacttitleNameX field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


