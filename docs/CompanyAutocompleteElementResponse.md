# CompanyAutocompleteElementResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiCompanyID** | **int32** | The unique ID of the Company | 
**SCompanyNameX** | **string** | The Name of the Company in the language of the requester | 
**BCompanyIsactive** | **bool** | Whether the Company is active or not | 

## Methods

### NewCompanyAutocompleteElementResponse

`func NewCompanyAutocompleteElementResponse(pkiCompanyID int32, sCompanyNameX string, bCompanyIsactive bool, ) *CompanyAutocompleteElementResponse`

NewCompanyAutocompleteElementResponse instantiates a new CompanyAutocompleteElementResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyAutocompleteElementResponseWithDefaults

`func NewCompanyAutocompleteElementResponseWithDefaults() *CompanyAutocompleteElementResponse`

NewCompanyAutocompleteElementResponseWithDefaults instantiates a new CompanyAutocompleteElementResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiCompanyID

`func (o *CompanyAutocompleteElementResponse) GetPkiCompanyID() int32`

GetPkiCompanyID returns the PkiCompanyID field if non-nil, zero value otherwise.

### GetPkiCompanyIDOk

`func (o *CompanyAutocompleteElementResponse) GetPkiCompanyIDOk() (*int32, bool)`

GetPkiCompanyIDOk returns a tuple with the PkiCompanyID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiCompanyID

`func (o *CompanyAutocompleteElementResponse) SetPkiCompanyID(v int32)`

SetPkiCompanyID sets PkiCompanyID field to given value.


### GetSCompanyNameX

`func (o *CompanyAutocompleteElementResponse) GetSCompanyNameX() string`

GetSCompanyNameX returns the SCompanyNameX field if non-nil, zero value otherwise.

### GetSCompanyNameXOk

`func (o *CompanyAutocompleteElementResponse) GetSCompanyNameXOk() (*string, bool)`

GetSCompanyNameXOk returns a tuple with the SCompanyNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSCompanyNameX

`func (o *CompanyAutocompleteElementResponse) SetSCompanyNameX(v string)`

SetSCompanyNameX sets SCompanyNameX field to given value.


### GetBCompanyIsactive

`func (o *CompanyAutocompleteElementResponse) GetBCompanyIsactive() bool`

GetBCompanyIsactive returns the BCompanyIsactive field if non-nil, zero value otherwise.

### GetBCompanyIsactiveOk

`func (o *CompanyAutocompleteElementResponse) GetBCompanyIsactiveOk() (*bool, bool)`

GetBCompanyIsactiveOk returns a tuple with the BCompanyIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBCompanyIsactive

`func (o *CompanyAutocompleteElementResponse) SetBCompanyIsactive(v bool)`

SetBCompanyIsactive sets BCompanyIsactive field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


