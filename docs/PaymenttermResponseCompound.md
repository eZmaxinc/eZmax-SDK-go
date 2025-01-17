# PaymenttermResponseCompound

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

### NewPaymenttermResponseCompound

`func NewPaymenttermResponseCompound(pkiPaymenttermID int32, sPaymenttermCode string, ePaymenttermType FieldEPaymenttermType, iPaymenttermDay int32, objPaymenttermDescription MultilingualPaymenttermDescription, bPaymenttermIsactive bool, objAudit CommonAudit, ) *PaymenttermResponseCompound`

NewPaymenttermResponseCompound instantiates a new PaymenttermResponseCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymenttermResponseCompoundWithDefaults

`func NewPaymenttermResponseCompoundWithDefaults() *PaymenttermResponseCompound`

NewPaymenttermResponseCompoundWithDefaults instantiates a new PaymenttermResponseCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiPaymenttermID

`func (o *PaymenttermResponseCompound) GetPkiPaymenttermID() int32`

GetPkiPaymenttermID returns the PkiPaymenttermID field if non-nil, zero value otherwise.

### GetPkiPaymenttermIDOk

`func (o *PaymenttermResponseCompound) GetPkiPaymenttermIDOk() (*int32, bool)`

GetPkiPaymenttermIDOk returns a tuple with the PkiPaymenttermID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiPaymenttermID

`func (o *PaymenttermResponseCompound) SetPkiPaymenttermID(v int32)`

SetPkiPaymenttermID sets PkiPaymenttermID field to given value.


### GetSPaymenttermCode

`func (o *PaymenttermResponseCompound) GetSPaymenttermCode() string`

GetSPaymenttermCode returns the SPaymenttermCode field if non-nil, zero value otherwise.

### GetSPaymenttermCodeOk

`func (o *PaymenttermResponseCompound) GetSPaymenttermCodeOk() (*string, bool)`

GetSPaymenttermCodeOk returns a tuple with the SPaymenttermCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSPaymenttermCode

`func (o *PaymenttermResponseCompound) SetSPaymenttermCode(v string)`

SetSPaymenttermCode sets SPaymenttermCode field to given value.


### GetEPaymenttermType

`func (o *PaymenttermResponseCompound) GetEPaymenttermType() FieldEPaymenttermType`

GetEPaymenttermType returns the EPaymenttermType field if non-nil, zero value otherwise.

### GetEPaymenttermTypeOk

`func (o *PaymenttermResponseCompound) GetEPaymenttermTypeOk() (*FieldEPaymenttermType, bool)`

GetEPaymenttermTypeOk returns a tuple with the EPaymenttermType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPaymenttermType

`func (o *PaymenttermResponseCompound) SetEPaymenttermType(v FieldEPaymenttermType)`

SetEPaymenttermType sets EPaymenttermType field to given value.


### GetIPaymenttermDay

`func (o *PaymenttermResponseCompound) GetIPaymenttermDay() int32`

GetIPaymenttermDay returns the IPaymenttermDay field if non-nil, zero value otherwise.

### GetIPaymenttermDayOk

`func (o *PaymenttermResponseCompound) GetIPaymenttermDayOk() (*int32, bool)`

GetIPaymenttermDayOk returns a tuple with the IPaymenttermDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIPaymenttermDay

`func (o *PaymenttermResponseCompound) SetIPaymenttermDay(v int32)`

SetIPaymenttermDay sets IPaymenttermDay field to given value.


### GetObjPaymenttermDescription

`func (o *PaymenttermResponseCompound) GetObjPaymenttermDescription() MultilingualPaymenttermDescription`

GetObjPaymenttermDescription returns the ObjPaymenttermDescription field if non-nil, zero value otherwise.

### GetObjPaymenttermDescriptionOk

`func (o *PaymenttermResponseCompound) GetObjPaymenttermDescriptionOk() (*MultilingualPaymenttermDescription, bool)`

GetObjPaymenttermDescriptionOk returns a tuple with the ObjPaymenttermDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjPaymenttermDescription

`func (o *PaymenttermResponseCompound) SetObjPaymenttermDescription(v MultilingualPaymenttermDescription)`

SetObjPaymenttermDescription sets ObjPaymenttermDescription field to given value.


### GetBPaymenttermIsactive

`func (o *PaymenttermResponseCompound) GetBPaymenttermIsactive() bool`

GetBPaymenttermIsactive returns the BPaymenttermIsactive field if non-nil, zero value otherwise.

### GetBPaymenttermIsactiveOk

`func (o *PaymenttermResponseCompound) GetBPaymenttermIsactiveOk() (*bool, bool)`

GetBPaymenttermIsactiveOk returns a tuple with the BPaymenttermIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBPaymenttermIsactive

`func (o *PaymenttermResponseCompound) SetBPaymenttermIsactive(v bool)`

SetBPaymenttermIsactive sets BPaymenttermIsactive field to given value.


### GetObjAudit

`func (o *PaymenttermResponseCompound) GetObjAudit() CommonAudit`

GetObjAudit returns the ObjAudit field if non-nil, zero value otherwise.

### GetObjAuditOk

`func (o *PaymenttermResponseCompound) GetObjAuditOk() (*CommonAudit, bool)`

GetObjAuditOk returns a tuple with the ObjAudit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjAudit

`func (o *PaymenttermResponseCompound) SetObjAudit(v CommonAudit)`

SetObjAudit sets ObjAudit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


