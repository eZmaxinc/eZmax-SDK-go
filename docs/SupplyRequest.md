# SupplyRequest

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

### NewSupplyRequest

`func NewSupplyRequest(fkiVariableexpenseID int32, sSupplyCode string, objSupplyDescription MultilingualSupplyDescription, dSupplyUnitprice string, bSupplyIsactive bool, bSupplyVariableprice bool, ) *SupplyRequest`

NewSupplyRequest instantiates a new SupplyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupplyRequestWithDefaults

`func NewSupplyRequestWithDefaults() *SupplyRequest`

NewSupplyRequestWithDefaults instantiates a new SupplyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiSupplyID

`func (o *SupplyRequest) GetPkiSupplyID() int32`

GetPkiSupplyID returns the PkiSupplyID field if non-nil, zero value otherwise.

### GetPkiSupplyIDOk

`func (o *SupplyRequest) GetPkiSupplyIDOk() (*int32, bool)`

GetPkiSupplyIDOk returns a tuple with the PkiSupplyID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiSupplyID

`func (o *SupplyRequest) SetPkiSupplyID(v int32)`

SetPkiSupplyID sets PkiSupplyID field to given value.

### HasPkiSupplyID

`func (o *SupplyRequest) HasPkiSupplyID() bool`

HasPkiSupplyID returns a boolean if a field has been set.

### GetFkiGlaccountID

`func (o *SupplyRequest) GetFkiGlaccountID() int32`

GetFkiGlaccountID returns the FkiGlaccountID field if non-nil, zero value otherwise.

### GetFkiGlaccountIDOk

`func (o *SupplyRequest) GetFkiGlaccountIDOk() (*int32, bool)`

GetFkiGlaccountIDOk returns a tuple with the FkiGlaccountID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiGlaccountID

`func (o *SupplyRequest) SetFkiGlaccountID(v int32)`

SetFkiGlaccountID sets FkiGlaccountID field to given value.

### HasFkiGlaccountID

`func (o *SupplyRequest) HasFkiGlaccountID() bool`

HasFkiGlaccountID returns a boolean if a field has been set.

### GetFkiGlaccountcontainerID

`func (o *SupplyRequest) GetFkiGlaccountcontainerID() int32`

GetFkiGlaccountcontainerID returns the FkiGlaccountcontainerID field if non-nil, zero value otherwise.

### GetFkiGlaccountcontainerIDOk

`func (o *SupplyRequest) GetFkiGlaccountcontainerIDOk() (*int32, bool)`

GetFkiGlaccountcontainerIDOk returns a tuple with the FkiGlaccountcontainerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiGlaccountcontainerID

`func (o *SupplyRequest) SetFkiGlaccountcontainerID(v int32)`

SetFkiGlaccountcontainerID sets FkiGlaccountcontainerID field to given value.

### HasFkiGlaccountcontainerID

`func (o *SupplyRequest) HasFkiGlaccountcontainerID() bool`

HasFkiGlaccountcontainerID returns a boolean if a field has been set.

### GetFkiVariableexpenseID

`func (o *SupplyRequest) GetFkiVariableexpenseID() int32`

GetFkiVariableexpenseID returns the FkiVariableexpenseID field if non-nil, zero value otherwise.

### GetFkiVariableexpenseIDOk

`func (o *SupplyRequest) GetFkiVariableexpenseIDOk() (*int32, bool)`

GetFkiVariableexpenseIDOk returns a tuple with the FkiVariableexpenseID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiVariableexpenseID

`func (o *SupplyRequest) SetFkiVariableexpenseID(v int32)`

SetFkiVariableexpenseID sets FkiVariableexpenseID field to given value.


### GetSSupplyCode

`func (o *SupplyRequest) GetSSupplyCode() string`

GetSSupplyCode returns the SSupplyCode field if non-nil, zero value otherwise.

### GetSSupplyCodeOk

`func (o *SupplyRequest) GetSSupplyCodeOk() (*string, bool)`

GetSSupplyCodeOk returns a tuple with the SSupplyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSSupplyCode

`func (o *SupplyRequest) SetSSupplyCode(v string)`

SetSSupplyCode sets SSupplyCode field to given value.


### GetObjSupplyDescription

`func (o *SupplyRequest) GetObjSupplyDescription() MultilingualSupplyDescription`

GetObjSupplyDescription returns the ObjSupplyDescription field if non-nil, zero value otherwise.

### GetObjSupplyDescriptionOk

`func (o *SupplyRequest) GetObjSupplyDescriptionOk() (*MultilingualSupplyDescription, bool)`

GetObjSupplyDescriptionOk returns a tuple with the ObjSupplyDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjSupplyDescription

`func (o *SupplyRequest) SetObjSupplyDescription(v MultilingualSupplyDescription)`

SetObjSupplyDescription sets ObjSupplyDescription field to given value.


### GetDSupplyUnitprice

`func (o *SupplyRequest) GetDSupplyUnitprice() string`

GetDSupplyUnitprice returns the DSupplyUnitprice field if non-nil, zero value otherwise.

### GetDSupplyUnitpriceOk

`func (o *SupplyRequest) GetDSupplyUnitpriceOk() (*string, bool)`

GetDSupplyUnitpriceOk returns a tuple with the DSupplyUnitprice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDSupplyUnitprice

`func (o *SupplyRequest) SetDSupplyUnitprice(v string)`

SetDSupplyUnitprice sets DSupplyUnitprice field to given value.


### GetBSupplyIsactive

`func (o *SupplyRequest) GetBSupplyIsactive() bool`

GetBSupplyIsactive returns the BSupplyIsactive field if non-nil, zero value otherwise.

### GetBSupplyIsactiveOk

`func (o *SupplyRequest) GetBSupplyIsactiveOk() (*bool, bool)`

GetBSupplyIsactiveOk returns a tuple with the BSupplyIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBSupplyIsactive

`func (o *SupplyRequest) SetBSupplyIsactive(v bool)`

SetBSupplyIsactive sets BSupplyIsactive field to given value.


### GetBSupplyVariableprice

`func (o *SupplyRequest) GetBSupplyVariableprice() bool`

GetBSupplyVariableprice returns the BSupplyVariableprice field if non-nil, zero value otherwise.

### GetBSupplyVariablepriceOk

`func (o *SupplyRequest) GetBSupplyVariablepriceOk() (*bool, bool)`

GetBSupplyVariablepriceOk returns a tuple with the BSupplyVariableprice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBSupplyVariableprice

`func (o *SupplyRequest) SetBSupplyVariableprice(v bool)`

SetBSupplyVariableprice sets BSupplyVariableprice field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


