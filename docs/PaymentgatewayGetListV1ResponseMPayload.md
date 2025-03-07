# PaymentgatewayGetListV1ResponseMPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IRowReturned** | **int32** | The number of rows returned | 
**IRowFiltered** | **int32** | The number of rows matching your filters (if any) or the total number of rows | 
**AObjPaymentgateway** | [**[]PaymentgatewayListElement**](PaymentgatewayListElement.md) |  | 

## Methods

### NewPaymentgatewayGetListV1ResponseMPayload

`func NewPaymentgatewayGetListV1ResponseMPayload(iRowReturned int32, iRowFiltered int32, aObjPaymentgateway []PaymentgatewayListElement, ) *PaymentgatewayGetListV1ResponseMPayload`

NewPaymentgatewayGetListV1ResponseMPayload instantiates a new PaymentgatewayGetListV1ResponseMPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentgatewayGetListV1ResponseMPayloadWithDefaults

`func NewPaymentgatewayGetListV1ResponseMPayloadWithDefaults() *PaymentgatewayGetListV1ResponseMPayload`

NewPaymentgatewayGetListV1ResponseMPayloadWithDefaults instantiates a new PaymentgatewayGetListV1ResponseMPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIRowReturned

`func (o *PaymentgatewayGetListV1ResponseMPayload) GetIRowReturned() int32`

GetIRowReturned returns the IRowReturned field if non-nil, zero value otherwise.

### GetIRowReturnedOk

`func (o *PaymentgatewayGetListV1ResponseMPayload) GetIRowReturnedOk() (*int32, bool)`

GetIRowReturnedOk returns a tuple with the IRowReturned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIRowReturned

`func (o *PaymentgatewayGetListV1ResponseMPayload) SetIRowReturned(v int32)`

SetIRowReturned sets IRowReturned field to given value.


### GetIRowFiltered

`func (o *PaymentgatewayGetListV1ResponseMPayload) GetIRowFiltered() int32`

GetIRowFiltered returns the IRowFiltered field if non-nil, zero value otherwise.

### GetIRowFilteredOk

`func (o *PaymentgatewayGetListV1ResponseMPayload) GetIRowFilteredOk() (*int32, bool)`

GetIRowFilteredOk returns a tuple with the IRowFiltered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIRowFiltered

`func (o *PaymentgatewayGetListV1ResponseMPayload) SetIRowFiltered(v int32)`

SetIRowFiltered sets IRowFiltered field to given value.


### GetAObjPaymentgateway

`func (o *PaymentgatewayGetListV1ResponseMPayload) GetAObjPaymentgateway() []PaymentgatewayListElement`

GetAObjPaymentgateway returns the AObjPaymentgateway field if non-nil, zero value otherwise.

### GetAObjPaymentgatewayOk

`func (o *PaymentgatewayGetListV1ResponseMPayload) GetAObjPaymentgatewayOk() (*[]PaymentgatewayListElement, bool)`

GetAObjPaymentgatewayOk returns a tuple with the AObjPaymentgateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjPaymentgateway

`func (o *PaymentgatewayGetListV1ResponseMPayload) SetAObjPaymentgateway(v []PaymentgatewayListElement)`

SetAObjPaymentgateway sets AObjPaymentgateway field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


