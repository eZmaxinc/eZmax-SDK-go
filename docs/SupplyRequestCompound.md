# SupplyRequestCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiSupplyID** | Pointer to **int32** | The unique ID of the Supply | [optional] 
**FkiGlaccountID** | Pointer to **int32** | The unique ID of the Glaccount | [optional] 
**FkiGlaccountcontainerID** | Pointer to **int32** | The unique ID of the Glaccountcontainer | [optional] 
**FkiVariableexpenseID** | **int32** | The unique ID of the Variableexpense | 
**SSupplyCode** | **string** | The code of the Supply | 
**ObjSupplyDescription** | [**MultilingualSupplyDescription**](MultilingualSupplyDescription.md) |  | 
**DSupplyUnitprice** | **string** | The unit price of the Supply | 
**BSupplyIsactive** | **bool** | Whether the supply is active or not | 
**BSupplyVariableprice** | **bool** | Whether if the price is variable | 

## Methods

### NewSupplyRequestCompound

`func NewSupplyRequestCompound(fkiVariableexpenseID int32, sSupplyCode string, objSupplyDescription MultilingualSupplyDescription, dSupplyUnitprice string, bSupplyIsactive bool, bSupplyVariableprice bool, ) *SupplyRequestCompound`

NewSupplyRequestCompound instantiates a new SupplyRequestCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupplyRequestCompoundWithDefaults

`func NewSupplyRequestCompoundWithDefaults() *SupplyRequestCompound`

NewSupplyRequestCompoundWithDefaults instantiates a new SupplyRequestCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiSupplyID

`func (o *SupplyRequestCompound) GetPkiSupplyID() int32`

GetPkiSupplyID returns the PkiSupplyID field if non-nil, zero value otherwise.

### GetPkiSupplyIDOk

`func (o *SupplyRequestCompound) GetPkiSupplyIDOk() (*int32, bool)`

GetPkiSupplyIDOk returns a tuple with the PkiSupplyID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiSupplyID

`func (o *SupplyRequestCompound) SetPkiSupplyID(v int32)`

SetPkiSupplyID sets PkiSupplyID field to given value.

### HasPkiSupplyID

`func (o *SupplyRequestCompound) HasPkiSupplyID() bool`

HasPkiSupplyID returns a boolean if a field has been set.

### GetFkiGlaccountID

`func (o *SupplyRequestCompound) GetFkiGlaccountID() int32`

GetFkiGlaccountID returns the FkiGlaccountID field if non-nil, zero value otherwise.

### GetFkiGlaccountIDOk

`func (o *SupplyRequestCompound) GetFkiGlaccountIDOk() (*int32, bool)`

GetFkiGlaccountIDOk returns a tuple with the FkiGlaccountID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiGlaccountID

`func (o *SupplyRequestCompound) SetFkiGlaccountID(v int32)`

SetFkiGlaccountID sets FkiGlaccountID field to given value.

### HasFkiGlaccountID

`func (o *SupplyRequestCompound) HasFkiGlaccountID() bool`

HasFkiGlaccountID returns a boolean if a field has been set.

### GetFkiGlaccountcontainerID

`func (o *SupplyRequestCompound) GetFkiGlaccountcontainerID() int32`

GetFkiGlaccountcontainerID returns the FkiGlaccountcontainerID field if non-nil, zero value otherwise.

### GetFkiGlaccountcontainerIDOk

`func (o *SupplyRequestCompound) GetFkiGlaccountcontainerIDOk() (*int32, bool)`

GetFkiGlaccountcontainerIDOk returns a tuple with the FkiGlaccountcontainerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiGlaccountcontainerID

`func (o *SupplyRequestCompound) SetFkiGlaccountcontainerID(v int32)`

SetFkiGlaccountcontainerID sets FkiGlaccountcontainerID field to given value.

### HasFkiGlaccountcontainerID

`func (o *SupplyRequestCompound) HasFkiGlaccountcontainerID() bool`

HasFkiGlaccountcontainerID returns a boolean if a field has been set.

### GetFkiVariableexpenseID

`func (o *SupplyRequestCompound) GetFkiVariableexpenseID() int32`

GetFkiVariableexpenseID returns the FkiVariableexpenseID field if non-nil, zero value otherwise.

### GetFkiVariableexpenseIDOk

`func (o *SupplyRequestCompound) GetFkiVariableexpenseIDOk() (*int32, bool)`

GetFkiVariableexpenseIDOk returns a tuple with the FkiVariableexpenseID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiVariableexpenseID

`func (o *SupplyRequestCompound) SetFkiVariableexpenseID(v int32)`

SetFkiVariableexpenseID sets FkiVariableexpenseID field to given value.


### GetSSupplyCode

`func (o *SupplyRequestCompound) GetSSupplyCode() string`

GetSSupplyCode returns the SSupplyCode field if non-nil, zero value otherwise.

### GetSSupplyCodeOk

`func (o *SupplyRequestCompound) GetSSupplyCodeOk() (*string, bool)`

GetSSupplyCodeOk returns a tuple with the SSupplyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSSupplyCode

`func (o *SupplyRequestCompound) SetSSupplyCode(v string)`

SetSSupplyCode sets SSupplyCode field to given value.


### GetObjSupplyDescription

`func (o *SupplyRequestCompound) GetObjSupplyDescription() MultilingualSupplyDescription`

GetObjSupplyDescription returns the ObjSupplyDescription field if non-nil, zero value otherwise.

### GetObjSupplyDescriptionOk

`func (o *SupplyRequestCompound) GetObjSupplyDescriptionOk() (*MultilingualSupplyDescription, bool)`

GetObjSupplyDescriptionOk returns a tuple with the ObjSupplyDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjSupplyDescription

`func (o *SupplyRequestCompound) SetObjSupplyDescription(v MultilingualSupplyDescription)`

SetObjSupplyDescription sets ObjSupplyDescription field to given value.


### GetDSupplyUnitprice

`func (o *SupplyRequestCompound) GetDSupplyUnitprice() string`

GetDSupplyUnitprice returns the DSupplyUnitprice field if non-nil, zero value otherwise.

### GetDSupplyUnitpriceOk

`func (o *SupplyRequestCompound) GetDSupplyUnitpriceOk() (*string, bool)`

GetDSupplyUnitpriceOk returns a tuple with the DSupplyUnitprice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDSupplyUnitprice

`func (o *SupplyRequestCompound) SetDSupplyUnitprice(v string)`

SetDSupplyUnitprice sets DSupplyUnitprice field to given value.


### GetBSupplyIsactive

`func (o *SupplyRequestCompound) GetBSupplyIsactive() bool`

GetBSupplyIsactive returns the BSupplyIsactive field if non-nil, zero value otherwise.

### GetBSupplyIsactiveOk

`func (o *SupplyRequestCompound) GetBSupplyIsactiveOk() (*bool, bool)`

GetBSupplyIsactiveOk returns a tuple with the BSupplyIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBSupplyIsactive

`func (o *SupplyRequestCompound) SetBSupplyIsactive(v bool)`

SetBSupplyIsactive sets BSupplyIsactive field to given value.


### GetBSupplyVariableprice

`func (o *SupplyRequestCompound) GetBSupplyVariableprice() bool`

GetBSupplyVariableprice returns the BSupplyVariableprice field if non-nil, zero value otherwise.

### GetBSupplyVariablepriceOk

`func (o *SupplyRequestCompound) GetBSupplyVariablepriceOk() (*bool, bool)`

GetBSupplyVariablepriceOk returns a tuple with the BSupplyVariableprice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBSupplyVariableprice

`func (o *SupplyRequestCompound) SetBSupplyVariableprice(v bool)`

SetBSupplyVariableprice sets BSupplyVariableprice field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


