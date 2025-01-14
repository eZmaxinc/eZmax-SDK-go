# SupplyListElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiSupplyID** | **int32** | The unique ID of the Supply | 
**FkiGlaccountID** | Pointer to **int32** | The unique ID of the Glaccount | [optional] 
**FkiGlaccountcontainerID** | Pointer to **int32** | The unique ID of the Glaccountcontainer | [optional] 
**FkiVariableexpenseID** | **int32** | The unique ID of the Variableexpense | 
**SSupplyCode** | **string** | The code of the Supply | 
**SSupplyDescriptionX** | **string** | The description of the Supply in the language of the requester | 
**DSupplyUnitprice** | **string** | The unit price of the Supply | 
**BSupplyIsactive** | **bool** | Whether the supply is active or not | 
**BSupplyVariableprice** | **bool** | Whether if the price is variable | 
**SGlaccountDescriptionX** | Pointer to **string** | The Description for the Glaccount in the language of the requester | [optional] 
**SGlaccountcontainerLongdescriptionX** | Pointer to **string** | The Description for the Glaccountcontainer in the language of the requester | [optional] 
**SVariableexpenseDescriptionX** | Pointer to **string** | The description of the Variableexpense in the language of the requester | [optional] 

## Methods

### NewSupplyListElement

`func NewSupplyListElement(pkiSupplyID int32, fkiVariableexpenseID int32, sSupplyCode string, sSupplyDescriptionX string, dSupplyUnitprice string, bSupplyIsactive bool, bSupplyVariableprice bool, ) *SupplyListElement`

NewSupplyListElement instantiates a new SupplyListElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupplyListElementWithDefaults

`func NewSupplyListElementWithDefaults() *SupplyListElement`

NewSupplyListElementWithDefaults instantiates a new SupplyListElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiSupplyID

`func (o *SupplyListElement) GetPkiSupplyID() int32`

GetPkiSupplyID returns the PkiSupplyID field if non-nil, zero value otherwise.

### GetPkiSupplyIDOk

`func (o *SupplyListElement) GetPkiSupplyIDOk() (*int32, bool)`

GetPkiSupplyIDOk returns a tuple with the PkiSupplyID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiSupplyID

`func (o *SupplyListElement) SetPkiSupplyID(v int32)`

SetPkiSupplyID sets PkiSupplyID field to given value.


### GetFkiGlaccountID

`func (o *SupplyListElement) GetFkiGlaccountID() int32`

GetFkiGlaccountID returns the FkiGlaccountID field if non-nil, zero value otherwise.

### GetFkiGlaccountIDOk

`func (o *SupplyListElement) GetFkiGlaccountIDOk() (*int32, bool)`

GetFkiGlaccountIDOk returns a tuple with the FkiGlaccountID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiGlaccountID

`func (o *SupplyListElement) SetFkiGlaccountID(v int32)`

SetFkiGlaccountID sets FkiGlaccountID field to given value.

### HasFkiGlaccountID

`func (o *SupplyListElement) HasFkiGlaccountID() bool`

HasFkiGlaccountID returns a boolean if a field has been set.

### GetFkiGlaccountcontainerID

`func (o *SupplyListElement) GetFkiGlaccountcontainerID() int32`

GetFkiGlaccountcontainerID returns the FkiGlaccountcontainerID field if non-nil, zero value otherwise.

### GetFkiGlaccountcontainerIDOk

`func (o *SupplyListElement) GetFkiGlaccountcontainerIDOk() (*int32, bool)`

GetFkiGlaccountcontainerIDOk returns a tuple with the FkiGlaccountcontainerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiGlaccountcontainerID

`func (o *SupplyListElement) SetFkiGlaccountcontainerID(v int32)`

SetFkiGlaccountcontainerID sets FkiGlaccountcontainerID field to given value.

### HasFkiGlaccountcontainerID

`func (o *SupplyListElement) HasFkiGlaccountcontainerID() bool`

HasFkiGlaccountcontainerID returns a boolean if a field has been set.

### GetFkiVariableexpenseID

`func (o *SupplyListElement) GetFkiVariableexpenseID() int32`

GetFkiVariableexpenseID returns the FkiVariableexpenseID field if non-nil, zero value otherwise.

### GetFkiVariableexpenseIDOk

`func (o *SupplyListElement) GetFkiVariableexpenseIDOk() (*int32, bool)`

GetFkiVariableexpenseIDOk returns a tuple with the FkiVariableexpenseID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiVariableexpenseID

`func (o *SupplyListElement) SetFkiVariableexpenseID(v int32)`

SetFkiVariableexpenseID sets FkiVariableexpenseID field to given value.


### GetSSupplyCode

`func (o *SupplyListElement) GetSSupplyCode() string`

GetSSupplyCode returns the SSupplyCode field if non-nil, zero value otherwise.

### GetSSupplyCodeOk

`func (o *SupplyListElement) GetSSupplyCodeOk() (*string, bool)`

GetSSupplyCodeOk returns a tuple with the SSupplyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSSupplyCode

`func (o *SupplyListElement) SetSSupplyCode(v string)`

SetSSupplyCode sets SSupplyCode field to given value.


### GetSSupplyDescriptionX

`func (o *SupplyListElement) GetSSupplyDescriptionX() string`

GetSSupplyDescriptionX returns the SSupplyDescriptionX field if non-nil, zero value otherwise.

### GetSSupplyDescriptionXOk

`func (o *SupplyListElement) GetSSupplyDescriptionXOk() (*string, bool)`

GetSSupplyDescriptionXOk returns a tuple with the SSupplyDescriptionX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSSupplyDescriptionX

`func (o *SupplyListElement) SetSSupplyDescriptionX(v string)`

SetSSupplyDescriptionX sets SSupplyDescriptionX field to given value.


### GetDSupplyUnitprice

`func (o *SupplyListElement) GetDSupplyUnitprice() string`

GetDSupplyUnitprice returns the DSupplyUnitprice field if non-nil, zero value otherwise.

### GetDSupplyUnitpriceOk

`func (o *SupplyListElement) GetDSupplyUnitpriceOk() (*string, bool)`

GetDSupplyUnitpriceOk returns a tuple with the DSupplyUnitprice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDSupplyUnitprice

`func (o *SupplyListElement) SetDSupplyUnitprice(v string)`

SetDSupplyUnitprice sets DSupplyUnitprice field to given value.


### GetBSupplyIsactive

`func (o *SupplyListElement) GetBSupplyIsactive() bool`

GetBSupplyIsactive returns the BSupplyIsactive field if non-nil, zero value otherwise.

### GetBSupplyIsactiveOk

`func (o *SupplyListElement) GetBSupplyIsactiveOk() (*bool, bool)`

GetBSupplyIsactiveOk returns a tuple with the BSupplyIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBSupplyIsactive

`func (o *SupplyListElement) SetBSupplyIsactive(v bool)`

SetBSupplyIsactive sets BSupplyIsactive field to given value.


### GetBSupplyVariableprice

`func (o *SupplyListElement) GetBSupplyVariableprice() bool`

GetBSupplyVariableprice returns the BSupplyVariableprice field if non-nil, zero value otherwise.

### GetBSupplyVariablepriceOk

`func (o *SupplyListElement) GetBSupplyVariablepriceOk() (*bool, bool)`

GetBSupplyVariablepriceOk returns a tuple with the BSupplyVariableprice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBSupplyVariableprice

`func (o *SupplyListElement) SetBSupplyVariableprice(v bool)`

SetBSupplyVariableprice sets BSupplyVariableprice field to given value.


### GetSGlaccountDescriptionX

`func (o *SupplyListElement) GetSGlaccountDescriptionX() string`

GetSGlaccountDescriptionX returns the SGlaccountDescriptionX field if non-nil, zero value otherwise.

### GetSGlaccountDescriptionXOk

`func (o *SupplyListElement) GetSGlaccountDescriptionXOk() (*string, bool)`

GetSGlaccountDescriptionXOk returns a tuple with the SGlaccountDescriptionX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSGlaccountDescriptionX

`func (o *SupplyListElement) SetSGlaccountDescriptionX(v string)`

SetSGlaccountDescriptionX sets SGlaccountDescriptionX field to given value.

### HasSGlaccountDescriptionX

`func (o *SupplyListElement) HasSGlaccountDescriptionX() bool`

HasSGlaccountDescriptionX returns a boolean if a field has been set.

### GetSGlaccountcontainerLongdescriptionX

`func (o *SupplyListElement) GetSGlaccountcontainerLongdescriptionX() string`

GetSGlaccountcontainerLongdescriptionX returns the SGlaccountcontainerLongdescriptionX field if non-nil, zero value otherwise.

### GetSGlaccountcontainerLongdescriptionXOk

`func (o *SupplyListElement) GetSGlaccountcontainerLongdescriptionXOk() (*string, bool)`

GetSGlaccountcontainerLongdescriptionXOk returns a tuple with the SGlaccountcontainerLongdescriptionX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSGlaccountcontainerLongdescriptionX

`func (o *SupplyListElement) SetSGlaccountcontainerLongdescriptionX(v string)`

SetSGlaccountcontainerLongdescriptionX sets SGlaccountcontainerLongdescriptionX field to given value.

### HasSGlaccountcontainerLongdescriptionX

`func (o *SupplyListElement) HasSGlaccountcontainerLongdescriptionX() bool`

HasSGlaccountcontainerLongdescriptionX returns a boolean if a field has been set.

### GetSVariableexpenseDescriptionX

`func (o *SupplyListElement) GetSVariableexpenseDescriptionX() string`

GetSVariableexpenseDescriptionX returns the SVariableexpenseDescriptionX field if non-nil, zero value otherwise.

### GetSVariableexpenseDescriptionXOk

`func (o *SupplyListElement) GetSVariableexpenseDescriptionXOk() (*string, bool)`

GetSVariableexpenseDescriptionXOk returns a tuple with the SVariableexpenseDescriptionX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSVariableexpenseDescriptionX

`func (o *SupplyListElement) SetSVariableexpenseDescriptionX(v string)`

SetSVariableexpenseDescriptionX sets SVariableexpenseDescriptionX field to given value.

### HasSVariableexpenseDescriptionX

`func (o *SupplyListElement) HasSVariableexpenseDescriptionX() bool`

HasSVariableexpenseDescriptionX returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


