# PaymentmethodAutocompleteElementResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiPaymentmethodID** | **int32** | The unique ID of the Paymentmethod | 
**SPaymentmethodDescriptionX** | **string** | The description of the Paymentmethod in the language of the requester | 
**BPaymentmethodIsactive** | **bool** | Whether the Paymentmethod is active or not | 

## Methods

### NewPaymentmethodAutocompleteElementResponse

`func NewPaymentmethodAutocompleteElementResponse(pkiPaymentmethodID int32, sPaymentmethodDescriptionX string, bPaymentmethodIsactive bool, ) *PaymentmethodAutocompleteElementResponse`

NewPaymentmethodAutocompleteElementResponse instantiates a new PaymentmethodAutocompleteElementResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentmethodAutocompleteElementResponseWithDefaults

`func NewPaymentmethodAutocompleteElementResponseWithDefaults() *PaymentmethodAutocompleteElementResponse`

NewPaymentmethodAutocompleteElementResponseWithDefaults instantiates a new PaymentmethodAutocompleteElementResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiPaymentmethodID

`func (o *PaymentmethodAutocompleteElementResponse) GetPkiPaymentmethodID() int32`

GetPkiPaymentmethodID returns the PkiPaymentmethodID field if non-nil, zero value otherwise.

### GetPkiPaymentmethodIDOk

`func (o *PaymentmethodAutocompleteElementResponse) GetPkiPaymentmethodIDOk() (*int32, bool)`

GetPkiPaymentmethodIDOk returns a tuple with the PkiPaymentmethodID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiPaymentmethodID

`func (o *PaymentmethodAutocompleteElementResponse) SetPkiPaymentmethodID(v int32)`

SetPkiPaymentmethodID sets PkiPaymentmethodID field to given value.


### GetSPaymentmethodDescriptionX

`func (o *PaymentmethodAutocompleteElementResponse) GetSPaymentmethodDescriptionX() string`

GetSPaymentmethodDescriptionX returns the SPaymentmethodDescriptionX field if non-nil, zero value otherwise.

### GetSPaymentmethodDescriptionXOk

`func (o *PaymentmethodAutocompleteElementResponse) GetSPaymentmethodDescriptionXOk() (*string, bool)`

GetSPaymentmethodDescriptionXOk returns a tuple with the SPaymentmethodDescriptionX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSPaymentmethodDescriptionX

`func (o *PaymentmethodAutocompleteElementResponse) SetSPaymentmethodDescriptionX(v string)`

SetSPaymentmethodDescriptionX sets SPaymentmethodDescriptionX field to given value.


### GetBPaymentmethodIsactive

`func (o *PaymentmethodAutocompleteElementResponse) GetBPaymentmethodIsactive() bool`

GetBPaymentmethodIsactive returns the BPaymentmethodIsactive field if non-nil, zero value otherwise.

### GetBPaymentmethodIsactiveOk

`func (o *PaymentmethodAutocompleteElementResponse) GetBPaymentmethodIsactiveOk() (*bool, bool)`

GetBPaymentmethodIsactiveOk returns a tuple with the BPaymentmethodIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBPaymentmethodIsactive

`func (o *PaymentmethodAutocompleteElementResponse) SetBPaymentmethodIsactive(v bool)`

SetBPaymentmethodIsactive sets BPaymentmethodIsactive field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


