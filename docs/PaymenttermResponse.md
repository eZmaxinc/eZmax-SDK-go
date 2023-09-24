# PaymenttermResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiPaymenttermID** | **int32** | The unique ID of the Paymentterm | 
**SPaymenttermCode** | **string** | The code of the Paymentterm | 
**EPaymenttermType** | [**FieldEPaymenttermType**](FieldEPaymenttermType.md) |  | 
**IPaymenttermDay** | **int32** | The day of the Paymentterm | 
**ObjPaymenttermDescription** | [**MultilingualPaymenttermDescription**](MultilingualPaymenttermDescription.md) |  | 
**BPaymenttermIsactive** | **bool** | Whether the Paymentterm is active or not | 
**ObjAudit** | [**CommonAudit**](CommonAudit.md) |  | 

## Methods

### NewPaymenttermResponse

`func NewPaymenttermResponse(pkiPaymenttermID int32, sPaymenttermCode string, ePaymenttermType FieldEPaymenttermType, iPaymenttermDay int32, objPaymenttermDescription MultilingualPaymenttermDescription, bPaymenttermIsactive bool, objAudit CommonAudit, ) *PaymenttermResponse`

NewPaymenttermResponse instantiates a new PaymenttermResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymenttermResponseWithDefaults

`func NewPaymenttermResponseWithDefaults() *PaymenttermResponse`

NewPaymenttermResponseWithDefaults instantiates a new PaymenttermResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiPaymenttermID

`func (o *PaymenttermResponse) GetPkiPaymenttermID() int32`

GetPkiPaymenttermID returns the PkiPaymenttermID field if non-nil, zero value otherwise.

### GetPkiPaymenttermIDOk

`func (o *PaymenttermResponse) GetPkiPaymenttermIDOk() (*int32, bool)`

GetPkiPaymenttermIDOk returns a tuple with the PkiPaymenttermID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiPaymenttermID

`func (o *PaymenttermResponse) SetPkiPaymenttermID(v int32)`

SetPkiPaymenttermID sets PkiPaymenttermID field to given value.


### GetSPaymenttermCode

`func (o *PaymenttermResponse) GetSPaymenttermCode() string`

GetSPaymenttermCode returns the SPaymenttermCode field if non-nil, zero value otherwise.

### GetSPaymenttermCodeOk

`func (o *PaymenttermResponse) GetSPaymenttermCodeOk() (*string, bool)`

GetSPaymenttermCodeOk returns a tuple with the SPaymenttermCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSPaymenttermCode

`func (o *PaymenttermResponse) SetSPaymenttermCode(v string)`

SetSPaymenttermCode sets SPaymenttermCode field to given value.


### GetEPaymenttermType

`func (o *PaymenttermResponse) GetEPaymenttermType() FieldEPaymenttermType`

GetEPaymenttermType returns the EPaymenttermType field if non-nil, zero value otherwise.

### GetEPaymenttermTypeOk

`func (o *PaymenttermResponse) GetEPaymenttermTypeOk() (*FieldEPaymenttermType, bool)`

GetEPaymenttermTypeOk returns a tuple with the EPaymenttermType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPaymenttermType

`func (o *PaymenttermResponse) SetEPaymenttermType(v FieldEPaymenttermType)`

SetEPaymenttermType sets EPaymenttermType field to given value.


### GetIPaymenttermDay

`func (o *PaymenttermResponse) GetIPaymenttermDay() int32`

GetIPaymenttermDay returns the IPaymenttermDay field if non-nil, zero value otherwise.

### GetIPaymenttermDayOk

`func (o *PaymenttermResponse) GetIPaymenttermDayOk() (*int32, bool)`

GetIPaymenttermDayOk returns a tuple with the IPaymenttermDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIPaymenttermDay

`func (o *PaymenttermResponse) SetIPaymenttermDay(v int32)`

SetIPaymenttermDay sets IPaymenttermDay field to given value.


### GetObjPaymenttermDescription

`func (o *PaymenttermResponse) GetObjPaymenttermDescription() MultilingualPaymenttermDescription`

GetObjPaymenttermDescription returns the ObjPaymenttermDescription field if non-nil, zero value otherwise.

### GetObjPaymenttermDescriptionOk

`func (o *PaymenttermResponse) GetObjPaymenttermDescriptionOk() (*MultilingualPaymenttermDescription, bool)`

GetObjPaymenttermDescriptionOk returns a tuple with the ObjPaymenttermDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjPaymenttermDescription

`func (o *PaymenttermResponse) SetObjPaymenttermDescription(v MultilingualPaymenttermDescription)`

SetObjPaymenttermDescription sets ObjPaymenttermDescription field to given value.


### GetBPaymenttermIsactive

`func (o *PaymenttermResponse) GetBPaymenttermIsactive() bool`

GetBPaymenttermIsactive returns the BPaymenttermIsactive field if non-nil, zero value otherwise.

### GetBPaymenttermIsactiveOk

`func (o *PaymenttermResponse) GetBPaymenttermIsactiveOk() (*bool, bool)`

GetBPaymenttermIsactiveOk returns a tuple with the BPaymenttermIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBPaymenttermIsactive

`func (o *PaymenttermResponse) SetBPaymenttermIsactive(v bool)`

SetBPaymenttermIsactive sets BPaymenttermIsactive field to given value.


### GetObjAudit

`func (o *PaymenttermResponse) GetObjAudit() CommonAudit`

GetObjAudit returns the ObjAudit field if non-nil, zero value otherwise.

### GetObjAuditOk

`func (o *PaymenttermResponse) GetObjAuditOk() (*CommonAudit, bool)`

GetObjAuditOk returns a tuple with the ObjAudit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjAudit

`func (o *PaymenttermResponse) SetObjAudit(v CommonAudit)`

SetObjAudit sets ObjAudit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


