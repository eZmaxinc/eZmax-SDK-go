# PaymenttermRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiPaymenttermID** | Pointer to **int32** | The unique ID of the Paymentterm | [optional] 
**SPaymenttermCode** | **string** | The code of the Paymentterm | 
**EPaymenttermType** | [**FieldEPaymenttermType**](FieldEPaymenttermType.md) |  | 
**IPaymenttermDay** | **int32** | The day of the Paymentterm | 
**ObjPaymenttermDescription** | [**MultilingualPaymenttermDescription**](MultilingualPaymenttermDescription.md) |  | 
**BPaymenttermIsactive** | **bool** | Whether the Paymentterm is active or not | 

## Methods

### NewPaymenttermRequest

`func NewPaymenttermRequest(sPaymenttermCode string, ePaymenttermType FieldEPaymenttermType, iPaymenttermDay int32, objPaymenttermDescription MultilingualPaymenttermDescription, bPaymenttermIsactive bool, ) *PaymenttermRequest`

NewPaymenttermRequest instantiates a new PaymenttermRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymenttermRequestWithDefaults

`func NewPaymenttermRequestWithDefaults() *PaymenttermRequest`

NewPaymenttermRequestWithDefaults instantiates a new PaymenttermRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiPaymenttermID

`func (o *PaymenttermRequest) GetPkiPaymenttermID() int32`

GetPkiPaymenttermID returns the PkiPaymenttermID field if non-nil, zero value otherwise.

### GetPkiPaymenttermIDOk

`func (o *PaymenttermRequest) GetPkiPaymenttermIDOk() (*int32, bool)`

GetPkiPaymenttermIDOk returns a tuple with the PkiPaymenttermID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiPaymenttermID

`func (o *PaymenttermRequest) SetPkiPaymenttermID(v int32)`

SetPkiPaymenttermID sets PkiPaymenttermID field to given value.

### HasPkiPaymenttermID

`func (o *PaymenttermRequest) HasPkiPaymenttermID() bool`

HasPkiPaymenttermID returns a boolean if a field has been set.

### GetSPaymenttermCode

`func (o *PaymenttermRequest) GetSPaymenttermCode() string`

GetSPaymenttermCode returns the SPaymenttermCode field if non-nil, zero value otherwise.

### GetSPaymenttermCodeOk

`func (o *PaymenttermRequest) GetSPaymenttermCodeOk() (*string, bool)`

GetSPaymenttermCodeOk returns a tuple with the SPaymenttermCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSPaymenttermCode

`func (o *PaymenttermRequest) SetSPaymenttermCode(v string)`

SetSPaymenttermCode sets SPaymenttermCode field to given value.


### GetEPaymenttermType

`func (o *PaymenttermRequest) GetEPaymenttermType() FieldEPaymenttermType`

GetEPaymenttermType returns the EPaymenttermType field if non-nil, zero value otherwise.

### GetEPaymenttermTypeOk

`func (o *PaymenttermRequest) GetEPaymenttermTypeOk() (*FieldEPaymenttermType, bool)`

GetEPaymenttermTypeOk returns a tuple with the EPaymenttermType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPaymenttermType

`func (o *PaymenttermRequest) SetEPaymenttermType(v FieldEPaymenttermType)`

SetEPaymenttermType sets EPaymenttermType field to given value.


### GetIPaymenttermDay

`func (o *PaymenttermRequest) GetIPaymenttermDay() int32`

GetIPaymenttermDay returns the IPaymenttermDay field if non-nil, zero value otherwise.

### GetIPaymenttermDayOk

`func (o *PaymenttermRequest) GetIPaymenttermDayOk() (*int32, bool)`

GetIPaymenttermDayOk returns a tuple with the IPaymenttermDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIPaymenttermDay

`func (o *PaymenttermRequest) SetIPaymenttermDay(v int32)`

SetIPaymenttermDay sets IPaymenttermDay field to given value.


### GetObjPaymenttermDescription

`func (o *PaymenttermRequest) GetObjPaymenttermDescription() MultilingualPaymenttermDescription`

GetObjPaymenttermDescription returns the ObjPaymenttermDescription field if non-nil, zero value otherwise.

### GetObjPaymenttermDescriptionOk

`func (o *PaymenttermRequest) GetObjPaymenttermDescriptionOk() (*MultilingualPaymenttermDescription, bool)`

GetObjPaymenttermDescriptionOk returns a tuple with the ObjPaymenttermDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjPaymenttermDescription

`func (o *PaymenttermRequest) SetObjPaymenttermDescription(v MultilingualPaymenttermDescription)`

SetObjPaymenttermDescription sets ObjPaymenttermDescription field to given value.


### GetBPaymenttermIsactive

`func (o *PaymenttermRequest) GetBPaymenttermIsactive() bool`

GetBPaymenttermIsactive returns the BPaymenttermIsactive field if non-nil, zero value otherwise.

### GetBPaymenttermIsactiveOk

`func (o *PaymenttermRequest) GetBPaymenttermIsactiveOk() (*bool, bool)`

GetBPaymenttermIsactiveOk returns a tuple with the BPaymenttermIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBPaymenttermIsactive

`func (o *PaymenttermRequest) SetBPaymenttermIsactive(v bool)`

SetBPaymenttermIsactive sets BPaymenttermIsactive field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


