# SupplyResponseCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiSupplyID** | **int32** | The unique ID of the Supply | 
**FkiGlaccountID** | Pointer to **int32** | The unique ID of the Glaccount | [optional] 
**FkiGlaccountcontainerID** | Pointer to **int32** | The unique ID of the Glaccountcontainer | [optional] 
**FkiVariableexpenseID** | **int32** | The unique ID of the Variableexpense | 
**SSupplyCode** | **string** | The code of the Supply | 
**ObjSupplyDescription** | [**MultilingualSupplyDescription**](MultilingualSupplyDescription.md) |  | 
**DSupplyUnitprice** | **string** | The unit price of the Supply | 
**BSupplyIsactive** | **bool** | Whether the supply is active or not | 
**BSupplyVariableprice** | **bool** | Whether if the price is variable | 
**SGlaccountDescriptionX** | Pointer to **string** | The Description for the Glaccount in the language of the requester | [optional] 
**SGlaccountcontainerLongdescriptionX** | Pointer to **string** | The Description for the Glaccountcontainer in the language of the requester | [optional] 
**SVariableexpenseDescriptionX** | Pointer to **string** | The description of the Variableexpense in the language of the requester | [optional] 

## Methods

### NewSupplyResponseCompound

`func NewSupplyResponseCompound(pkiSupplyID int32, fkiVariableexpenseID int32, sSupplyCode string, objSupplyDescription MultilingualSupplyDescription, dSupplyUnitprice string, bSupplyIsactive bool, bSupplyVariableprice bool, ) *SupplyResponseCompound`

NewSupplyResponseCompound instantiates a new SupplyResponseCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupplyResponseCompoundWithDefaults

`func NewSupplyResponseCompoundWithDefaults() *SupplyResponseCompound`

NewSupplyResponseCompoundWithDefaults instantiates a new SupplyResponseCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiSupplyID

`func (o *SupplyResponseCompound) GetPkiSupplyID() int32`

GetPkiSupplyID returns the PkiSupplyID field if non-nil, zero value otherwise.

### GetPkiSupplyIDOk

`func (o *SupplyResponseCompound) GetPkiSupplyIDOk() (*int32, bool)`

GetPkiSupplyIDOk returns a tuple with the PkiSupplyID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiSupplyID

`func (o *SupplyResponseCompound) SetPkiSupplyID(v int32)`

SetPkiSupplyID sets PkiSupplyID field to given value.


### GetFkiGlaccountID

`func (o *SupplyResponseCompound) GetFkiGlaccountID() int32`

GetFkiGlaccountID returns the FkiGlaccountID field if non-nil, zero value otherwise.

### GetFkiGlaccountIDOk

`func (o *SupplyResponseCompound) GetFkiGlaccountIDOk() (*int32, bool)`

GetFkiGlaccountIDOk returns a tuple with the FkiGlaccountID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiGlaccountID

`func (o *SupplyResponseCompound) SetFkiGlaccountID(v int32)`

SetFkiGlaccountID sets FkiGlaccountID field to given value.

### HasFkiGlaccountID

`func (o *SupplyResponseCompound) HasFkiGlaccountID() bool`

HasFkiGlaccountID returns a boolean if a field has been set.

### GetFkiGlaccountcontainerID

`func (o *SupplyResponseCompound) GetFkiGlaccountcontainerID() int32`

GetFkiGlaccountcontainerID returns the FkiGlaccountcontainerID field if non-nil, zero value otherwise.

### GetFkiGlaccountcontainerIDOk

`func (o *SupplyResponseCompound) GetFkiGlaccountcontainerIDOk() (*int32, bool)`

GetFkiGlaccountcontainerIDOk returns a tuple with the FkiGlaccountcontainerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiGlaccountcontainerID

`func (o *SupplyResponseCompound) SetFkiGlaccountcontainerID(v int32)`

SetFkiGlaccountcontainerID sets FkiGlaccountcontainerID field to given value.

### HasFkiGlaccountcontainerID

`func (o *SupplyResponseCompound) HasFkiGlaccountcontainerID() bool`

HasFkiGlaccountcontainerID returns a boolean if a field has been set.

### GetFkiVariableexpenseID

`func (o *SupplyResponseCompound) GetFkiVariableexpenseID() int32`

GetFkiVariableexpenseID returns the FkiVariableexpenseID field if non-nil, zero value otherwise.

### GetFkiVariableexpenseIDOk

`func (o *SupplyResponseCompound) GetFkiVariableexpenseIDOk() (*int32, bool)`

GetFkiVariableexpenseIDOk returns a tuple with the FkiVariableexpenseID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiVariableexpenseID

`func (o *SupplyResponseCompound) SetFkiVariableexpenseID(v int32)`

SetFkiVariableexpenseID sets FkiVariableexpenseID field to given value.


### GetSSupplyCode

`func (o *SupplyResponseCompound) GetSSupplyCode() string`

GetSSupplyCode returns the SSupplyCode field if non-nil, zero value otherwise.

### GetSSupplyCodeOk

`func (o *SupplyResponseCompound) GetSSupplyCodeOk() (*string, bool)`

GetSSupplyCodeOk returns a tuple with the SSupplyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSSupplyCode

`func (o *SupplyResponseCompound) SetSSupplyCode(v string)`

SetSSupplyCode sets SSupplyCode field to given value.


### GetObjSupplyDescription

`func (o *SupplyResponseCompound) GetObjSupplyDescription() MultilingualSupplyDescription`

GetObjSupplyDescription returns the ObjSupplyDescription field if non-nil, zero value otherwise.

### GetObjSupplyDescriptionOk

`func (o *SupplyResponseCompound) GetObjSupplyDescriptionOk() (*MultilingualSupplyDescription, bool)`

GetObjSupplyDescriptionOk returns a tuple with the ObjSupplyDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjSupplyDescription

`func (o *SupplyResponseCompound) SetObjSupplyDescription(v MultilingualSupplyDescription)`

SetObjSupplyDescription sets ObjSupplyDescription field to given value.


### GetDSupplyUnitprice

`func (o *SupplyResponseCompound) GetDSupplyUnitprice() string`

GetDSupplyUnitprice returns the DSupplyUnitprice field if non-nil, zero value otherwise.

### GetDSupplyUnitpriceOk

`func (o *SupplyResponseCompound) GetDSupplyUnitpriceOk() (*string, bool)`

GetDSupplyUnitpriceOk returns a tuple with the DSupplyUnitprice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDSupplyUnitprice

`func (o *SupplyResponseCompound) SetDSupplyUnitprice(v string)`

SetDSupplyUnitprice sets DSupplyUnitprice field to given value.


### GetBSupplyIsactive

`func (o *SupplyResponseCompound) GetBSupplyIsactive() bool`

GetBSupplyIsactive returns the BSupplyIsactive field if non-nil, zero value otherwise.

### GetBSupplyIsactiveOk

`func (o *SupplyResponseCompound) GetBSupplyIsactiveOk() (*bool, bool)`

GetBSupplyIsactiveOk returns a tuple with the BSupplyIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBSupplyIsactive

`func (o *SupplyResponseCompound) SetBSupplyIsactive(v bool)`

SetBSupplyIsactive sets BSupplyIsactive field to given value.


### GetBSupplyVariableprice

`func (o *SupplyResponseCompound) GetBSupplyVariableprice() bool`

GetBSupplyVariableprice returns the BSupplyVariableprice field if non-nil, zero value otherwise.

### GetBSupplyVariablepriceOk

`func (o *SupplyResponseCompound) GetBSupplyVariablepriceOk() (*bool, bool)`

GetBSupplyVariablepriceOk returns a tuple with the BSupplyVariableprice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBSupplyVariableprice

`func (o *SupplyResponseCompound) SetBSupplyVariableprice(v bool)`

SetBSupplyVariableprice sets BSupplyVariableprice field to given value.


### GetSGlaccountDescriptionX

`func (o *SupplyResponseCompound) GetSGlaccountDescriptionX() string`

GetSGlaccountDescriptionX returns the SGlaccountDescriptionX field if non-nil, zero value otherwise.

### GetSGlaccountDescriptionXOk

`func (o *SupplyResponseCompound) GetSGlaccountDescriptionXOk() (*string, bool)`

GetSGlaccountDescriptionXOk returns a tuple with the SGlaccountDescriptionX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSGlaccountDescriptionX

`func (o *SupplyResponseCompound) SetSGlaccountDescriptionX(v string)`

SetSGlaccountDescriptionX sets SGlaccountDescriptionX field to given value.

### HasSGlaccountDescriptionX

`func (o *SupplyResponseCompound) HasSGlaccountDescriptionX() bool`

HasSGlaccountDescriptionX returns a boolean if a field has been set.

### GetSGlaccountcontainerLongdescriptionX

`func (o *SupplyResponseCompound) GetSGlaccountcontainerLongdescriptionX() string`

GetSGlaccountcontainerLongdescriptionX returns the SGlaccountcontainerLongdescriptionX field if non-nil, zero value otherwise.

### GetSGlaccountcontainerLongdescriptionXOk

`func (o *SupplyResponseCompound) GetSGlaccountcontainerLongdescriptionXOk() (*string, bool)`

GetSGlaccountcontainerLongdescriptionXOk returns a tuple with the SGlaccountcontainerLongdescriptionX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSGlaccountcontainerLongdescriptionX

`func (o *SupplyResponseCompound) SetSGlaccountcontainerLongdescriptionX(v string)`

SetSGlaccountcontainerLongdescriptionX sets SGlaccountcontainerLongdescriptionX field to given value.

### HasSGlaccountcontainerLongdescriptionX

`func (o *SupplyResponseCompound) HasSGlaccountcontainerLongdescriptionX() bool`

HasSGlaccountcontainerLongdescriptionX returns a boolean if a field has been set.

### GetSVariableexpenseDescriptionX

`func (o *SupplyResponseCompound) GetSVariableexpenseDescriptionX() string`

GetSVariableexpenseDescriptionX returns the SVariableexpenseDescriptionX field if non-nil, zero value otherwise.

### GetSVariableexpenseDescriptionXOk

`func (o *SupplyResponseCompound) GetSVariableexpenseDescriptionXOk() (*string, bool)`

GetSVariableexpenseDescriptionXOk returns a tuple with the SVariableexpenseDescriptionX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSVariableexpenseDescriptionX

`func (o *SupplyResponseCompound) SetSVariableexpenseDescriptionX(v string)`

SetSVariableexpenseDescriptionX sets SVariableexpenseDescriptionX field to given value.

### HasSVariableexpenseDescriptionX

`func (o *SupplyResponseCompound) HasSVariableexpenseDescriptionX() bool`

HasSVariableexpenseDescriptionX returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


