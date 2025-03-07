# CustomerAutocompleteElementResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiCustomerID** | **int32** | The unique ID of the Customer. | 
**SCustomerName** | **string** | The name of the Customer | 
**BCustomerIsactive** | **bool** | Whether the customer is active or not | 

## Methods

### NewCustomerAutocompleteElementResponse

`func NewCustomerAutocompleteElementResponse(pkiCustomerID int32, sCustomerName string, bCustomerIsactive bool, ) *CustomerAutocompleteElementResponse`

NewCustomerAutocompleteElementResponse instantiates a new CustomerAutocompleteElementResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerAutocompleteElementResponseWithDefaults

`func NewCustomerAutocompleteElementResponseWithDefaults() *CustomerAutocompleteElementResponse`

NewCustomerAutocompleteElementResponseWithDefaults instantiates a new CustomerAutocompleteElementResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiCustomerID

`func (o *CustomerAutocompleteElementResponse) GetPkiCustomerID() int32`

GetPkiCustomerID returns the PkiCustomerID field if non-nil, zero value otherwise.

### GetPkiCustomerIDOk

`func (o *CustomerAutocompleteElementResponse) GetPkiCustomerIDOk() (*int32, bool)`

GetPkiCustomerIDOk returns a tuple with the PkiCustomerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiCustomerID

`func (o *CustomerAutocompleteElementResponse) SetPkiCustomerID(v int32)`

SetPkiCustomerID sets PkiCustomerID field to given value.


### GetSCustomerName

`func (o *CustomerAutocompleteElementResponse) GetSCustomerName() string`

GetSCustomerName returns the SCustomerName field if non-nil, zero value otherwise.

### GetSCustomerNameOk

`func (o *CustomerAutocompleteElementResponse) GetSCustomerNameOk() (*string, bool)`

GetSCustomerNameOk returns a tuple with the SCustomerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSCustomerName

`func (o *CustomerAutocompleteElementResponse) SetSCustomerName(v string)`

SetSCustomerName sets SCustomerName field to given value.


### GetBCustomerIsactive

`func (o *CustomerAutocompleteElementResponse) GetBCustomerIsactive() bool`

GetBCustomerIsactive returns the BCustomerIsactive field if non-nil, zero value otherwise.

### GetBCustomerIsactiveOk

`func (o *CustomerAutocompleteElementResponse) GetBCustomerIsactiveOk() (*bool, bool)`

GetBCustomerIsactiveOk returns a tuple with the BCustomerIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBCustomerIsactive

`func (o *CustomerAutocompleteElementResponse) SetBCustomerIsactive(v bool)`

SetBCustomerIsactive sets BCustomerIsactive field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


