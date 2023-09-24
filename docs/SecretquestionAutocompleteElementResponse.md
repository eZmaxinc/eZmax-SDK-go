# SecretquestionAutocompleteElementResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SSecretquestionTextX** | **string** | The text of the Secretquestion in the language of the requester | 
**PkiSecretquestionID** | **int32** | The unique ID of the Secretquestion.  Valid values:  |Value|Description| |-|-| |1|The name of the hospital in which you were born| |2|The name of your grade school| |3|The last name of your favorite teacher| |4|Your favorite sports team| |5|Your favorite TV show| |6|Your favorite movie| |7|The name of the street on which you grew up| |8|The name of your first employer| |9|Your first car| |10|Your favorite food| |11|The name of your first pet| |12|Favorite musician/band| |13|What instrument you play| |14|Your father&#39;s middle name| |15|Your mother&#39;s maiden name| |16|Name of your eldest child| |17|Your spouse&#39;s middle name| |18|Favorite restaurant| |19|Childhood nickname| |20|Favorite vacation destination| |21|Your boat&#39;s name| |22|Date of Birth (YYYY-MM-DD)| | 
**BSecretquestionIsactive** | **bool** | Whether the Secretquestion is active or not | 

## Methods

### NewSecretquestionAutocompleteElementResponse

`func NewSecretquestionAutocompleteElementResponse(sSecretquestionTextX string, pkiSecretquestionID int32, bSecretquestionIsactive bool, ) *SecretquestionAutocompleteElementResponse`

NewSecretquestionAutocompleteElementResponse instantiates a new SecretquestionAutocompleteElementResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretquestionAutocompleteElementResponseWithDefaults

`func NewSecretquestionAutocompleteElementResponseWithDefaults() *SecretquestionAutocompleteElementResponse`

NewSecretquestionAutocompleteElementResponseWithDefaults instantiates a new SecretquestionAutocompleteElementResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSSecretquestionTextX

`func (o *SecretquestionAutocompleteElementResponse) GetSSecretquestionTextX() string`

GetSSecretquestionTextX returns the SSecretquestionTextX field if non-nil, zero value otherwise.

### GetSSecretquestionTextXOk

`func (o *SecretquestionAutocompleteElementResponse) GetSSecretquestionTextXOk() (*string, bool)`

GetSSecretquestionTextXOk returns a tuple with the SSecretquestionTextX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSSecretquestionTextX

`func (o *SecretquestionAutocompleteElementResponse) SetSSecretquestionTextX(v string)`

SetSSecretquestionTextX sets SSecretquestionTextX field to given value.


### GetPkiSecretquestionID

`func (o *SecretquestionAutocompleteElementResponse) GetPkiSecretquestionID() int32`

GetPkiSecretquestionID returns the PkiSecretquestionID field if non-nil, zero value otherwise.

### GetPkiSecretquestionIDOk

`func (o *SecretquestionAutocompleteElementResponse) GetPkiSecretquestionIDOk() (*int32, bool)`

GetPkiSecretquestionIDOk returns a tuple with the PkiSecretquestionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiSecretquestionID

`func (o *SecretquestionAutocompleteElementResponse) SetPkiSecretquestionID(v int32)`

SetPkiSecretquestionID sets PkiSecretquestionID field to given value.


### GetBSecretquestionIsactive

`func (o *SecretquestionAutocompleteElementResponse) GetBSecretquestionIsactive() bool`

GetBSecretquestionIsactive returns the BSecretquestionIsactive field if non-nil, zero value otherwise.

### GetBSecretquestionIsactiveOk

`func (o *SecretquestionAutocompleteElementResponse) GetBSecretquestionIsactiveOk() (*bool, bool)`

GetBSecretquestionIsactiveOk returns a tuple with the BSecretquestionIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBSecretquestionIsactive

`func (o *SecretquestionAutocompleteElementResponse) SetBSecretquestionIsactive(v bool)`

SetBSecretquestionIsactive sets BSecretquestionIsactive field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


