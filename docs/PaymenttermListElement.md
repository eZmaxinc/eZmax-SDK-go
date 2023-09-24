# PaymenttermListElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiPaymenttermID** | **int32** | The unique ID of the Paymentterm | 
**SPaymenttermCode** | **string** | The code of the Paymentterm | 
**EPaymenttermType** | [**FieldEPaymenttermType**](FieldEPaymenttermType.md) |  | 
**IPaymenttermDay** | **int32** | The day of the Paymentterm | 
**SPaymenttermDescriptionX** | **string** | The description of the Paymentterm in the language of the requester | 
**BPaymenttermIsactive** | **bool** | Whether the Paymentterm is active or not | 

## Methods

### NewPaymenttermListElement

`func NewPaymenttermListElement(pkiPaymenttermID int32, sPaymenttermCode string, ePaymenttermType FieldEPaymenttermType, iPaymenttermDay int32, sPaymenttermDescriptionX string, bPaymenttermIsactive bool, ) *PaymenttermListElement`

NewPaymenttermListElement instantiates a new PaymenttermListElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymenttermListElementWithDefaults

`func NewPaymenttermListElementWithDefaults() *PaymenttermListElement`

NewPaymenttermListElementWithDefaults instantiates a new PaymenttermListElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiPaymenttermID

`func (o *PaymenttermListElement) GetPkiPaymenttermID() int32`

GetPkiPaymenttermID returns the PkiPaymenttermID field if non-nil, zero value otherwise.

### GetPkiPaymenttermIDOk

`func (o *PaymenttermListElement) GetPkiPaymenttermIDOk() (*int32, bool)`

GetPkiPaymenttermIDOk returns a tuple with the PkiPaymenttermID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiPaymenttermID

`func (o *PaymenttermListElement) SetPkiPaymenttermID(v int32)`

SetPkiPaymenttermID sets PkiPaymenttermID field to given value.


### GetSPaymenttermCode

`func (o *PaymenttermListElement) GetSPaymenttermCode() string`

GetSPaymenttermCode returns the SPaymenttermCode field if non-nil, zero value otherwise.

### GetSPaymenttermCodeOk

`func (o *PaymenttermListElement) GetSPaymenttermCodeOk() (*string, bool)`

GetSPaymenttermCodeOk returns a tuple with the SPaymenttermCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSPaymenttermCode

`func (o *PaymenttermListElement) SetSPaymenttermCode(v string)`

SetSPaymenttermCode sets SPaymenttermCode field to given value.


### GetEPaymenttermType

`func (o *PaymenttermListElement) GetEPaymenttermType() FieldEPaymenttermType`

GetEPaymenttermType returns the EPaymenttermType field if non-nil, zero value otherwise.

### GetEPaymenttermTypeOk

`func (o *PaymenttermListElement) GetEPaymenttermTypeOk() (*FieldEPaymenttermType, bool)`

GetEPaymenttermTypeOk returns a tuple with the EPaymenttermType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPaymenttermType

`func (o *PaymenttermListElement) SetEPaymenttermType(v FieldEPaymenttermType)`

SetEPaymenttermType sets EPaymenttermType field to given value.


### GetIPaymenttermDay

`func (o *PaymenttermListElement) GetIPaymenttermDay() int32`

GetIPaymenttermDay returns the IPaymenttermDay field if non-nil, zero value otherwise.

### GetIPaymenttermDayOk

`func (o *PaymenttermListElement) GetIPaymenttermDayOk() (*int32, bool)`

GetIPaymenttermDayOk returns a tuple with the IPaymenttermDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIPaymenttermDay

`func (o *PaymenttermListElement) SetIPaymenttermDay(v int32)`

SetIPaymenttermDay sets IPaymenttermDay field to given value.


### GetSPaymenttermDescriptionX

`func (o *PaymenttermListElement) GetSPaymenttermDescriptionX() string`

GetSPaymenttermDescriptionX returns the SPaymenttermDescriptionX field if non-nil, zero value otherwise.

### GetSPaymenttermDescriptionXOk

`func (o *PaymenttermListElement) GetSPaymenttermDescriptionXOk() (*string, bool)`

GetSPaymenttermDescriptionXOk returns a tuple with the SPaymenttermDescriptionX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSPaymenttermDescriptionX

`func (o *PaymenttermListElement) SetSPaymenttermDescriptionX(v string)`

SetSPaymenttermDescriptionX sets SPaymenttermDescriptionX field to given value.


### GetBPaymenttermIsactive

`func (o *PaymenttermListElement) GetBPaymenttermIsactive() bool`

GetBPaymenttermIsactive returns the BPaymenttermIsactive field if non-nil, zero value otherwise.

### GetBPaymenttermIsactiveOk

`func (o *PaymenttermListElement) GetBPaymenttermIsactiveOk() (*bool, bool)`

GetBPaymenttermIsactiveOk returns a tuple with the BPaymenttermIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBPaymenttermIsactive

`func (o *PaymenttermListElement) SetBPaymenttermIsactive(v bool)`

SetBPaymenttermIsactive sets BPaymenttermIsactive field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


