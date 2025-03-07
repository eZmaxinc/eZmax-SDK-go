# CustomCreditcardRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FksCreditcardtokenID** | **string** | The creditcard token identifier | 
**SCreditcardCVV** | **string** | The creditcard card CVV | 
**ObjCreditcarddetail** | [**CreditcarddetailRequest**](CreditcarddetailRequest.md) |  | 

## Methods

### NewCustomCreditcardRequest

`func NewCustomCreditcardRequest(fksCreditcardtokenID string, sCreditcardCVV string, objCreditcarddetail CreditcarddetailRequest, ) *CustomCreditcardRequest`

NewCustomCreditcardRequest instantiates a new CustomCreditcardRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomCreditcardRequestWithDefaults

`func NewCustomCreditcardRequestWithDefaults() *CustomCreditcardRequest`

NewCustomCreditcardRequestWithDefaults instantiates a new CustomCreditcardRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFksCreditcardtokenID

`func (o *CustomCreditcardRequest) GetFksCreditcardtokenID() string`

GetFksCreditcardtokenID returns the FksCreditcardtokenID field if non-nil, zero value otherwise.

### GetFksCreditcardtokenIDOk

`func (o *CustomCreditcardRequest) GetFksCreditcardtokenIDOk() (*string, bool)`

GetFksCreditcardtokenIDOk returns a tuple with the FksCreditcardtokenID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFksCreditcardtokenID

`func (o *CustomCreditcardRequest) SetFksCreditcardtokenID(v string)`

SetFksCreditcardtokenID sets FksCreditcardtokenID field to given value.


### GetSCreditcardCVV

`func (o *CustomCreditcardRequest) GetSCreditcardCVV() string`

GetSCreditcardCVV returns the SCreditcardCVV field if non-nil, zero value otherwise.

### GetSCreditcardCVVOk

`func (o *CustomCreditcardRequest) GetSCreditcardCVVOk() (*string, bool)`

GetSCreditcardCVVOk returns a tuple with the SCreditcardCVV field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSCreditcardCVV

`func (o *CustomCreditcardRequest) SetSCreditcardCVV(v string)`

SetSCreditcardCVV sets SCreditcardCVV field to given value.


### GetObjCreditcarddetail

`func (o *CustomCreditcardRequest) GetObjCreditcarddetail() CreditcarddetailRequest`

GetObjCreditcarddetail returns the ObjCreditcarddetail field if non-nil, zero value otherwise.

### GetObjCreditcarddetailOk

`func (o *CustomCreditcardRequest) GetObjCreditcarddetailOk() (*CreditcarddetailRequest, bool)`

GetObjCreditcarddetailOk returns a tuple with the ObjCreditcarddetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjCreditcarddetail

`func (o *CustomCreditcardRequest) SetObjCreditcarddetail(v CreditcarddetailRequest)`

SetObjCreditcarddetail sets ObjCreditcarddetail field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


