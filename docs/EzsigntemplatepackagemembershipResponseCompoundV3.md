# EzsigntemplatepackagemembershipResponseCompoundV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsigntemplatepackagemembershipID** | **int32** | The unique ID of the Ezsigntemplatepackagemembership | 
**FkiEzsigntemplatepackageID** | **int32** | The unique ID of the Ezsigntemplatepackage | 
**FkiEzsigntemplateID** | **int32** | The unique ID of the Ezsigntemplate | 
**IEzsigntemplatepackagemembershipOrder** | **int32** | The order in which the Ezsigntemplate will be imported when using an Ezsigntemplatepackage. | 
**ObjEzsigntemplate** | [**EzsigntemplateResponseCompoundV4**](EzsigntemplateResponseCompoundV4.md) |  | 
**AObjEzsigntemplatepackagesignermembership** | [**[]EzsigntemplatepackagesignermembershipResponseCompound**](EzsigntemplatepackagesignermembershipResponseCompound.md) |  | 

## Methods

### NewEzsigntemplatepackagemembershipResponseCompoundV3

`func NewEzsigntemplatepackagemembershipResponseCompoundV3(pkiEzsigntemplatepackagemembershipID int32, fkiEzsigntemplatepackageID int32, fkiEzsigntemplateID int32, iEzsigntemplatepackagemembershipOrder int32, objEzsigntemplate EzsigntemplateResponseCompoundV4, aObjEzsigntemplatepackagesignermembership []EzsigntemplatepackagesignermembershipResponseCompound, ) *EzsigntemplatepackagemembershipResponseCompoundV3`

NewEzsigntemplatepackagemembershipResponseCompoundV3 instantiates a new EzsigntemplatepackagemembershipResponseCompoundV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsigntemplatepackagemembershipResponseCompoundV3WithDefaults

`func NewEzsigntemplatepackagemembershipResponseCompoundV3WithDefaults() *EzsigntemplatepackagemembershipResponseCompoundV3`

NewEzsigntemplatepackagemembershipResponseCompoundV3WithDefaults instantiates a new EzsigntemplatepackagemembershipResponseCompoundV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsigntemplatepackagemembershipID

`func (o *EzsigntemplatepackagemembershipResponseCompoundV3) GetPkiEzsigntemplatepackagemembershipID() int32`

GetPkiEzsigntemplatepackagemembershipID returns the PkiEzsigntemplatepackagemembershipID field if non-nil, zero value otherwise.

### GetPkiEzsigntemplatepackagemembershipIDOk

`func (o *EzsigntemplatepackagemembershipResponseCompoundV3) GetPkiEzsigntemplatepackagemembershipIDOk() (*int32, bool)`

GetPkiEzsigntemplatepackagemembershipIDOk returns a tuple with the PkiEzsigntemplatepackagemembershipID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsigntemplatepackagemembershipID

`func (o *EzsigntemplatepackagemembershipResponseCompoundV3) SetPkiEzsigntemplatepackagemembershipID(v int32)`

SetPkiEzsigntemplatepackagemembershipID sets PkiEzsigntemplatepackagemembershipID field to given value.


### GetFkiEzsigntemplatepackageID

`func (o *EzsigntemplatepackagemembershipResponseCompoundV3) GetFkiEzsigntemplatepackageID() int32`

GetFkiEzsigntemplatepackageID returns the FkiEzsigntemplatepackageID field if non-nil, zero value otherwise.

### GetFkiEzsigntemplatepackageIDOk

`func (o *EzsigntemplatepackagemembershipResponseCompoundV3) GetFkiEzsigntemplatepackageIDOk() (*int32, bool)`

GetFkiEzsigntemplatepackageIDOk returns a tuple with the FkiEzsigntemplatepackageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsigntemplatepackageID

`func (o *EzsigntemplatepackagemembershipResponseCompoundV3) SetFkiEzsigntemplatepackageID(v int32)`

SetFkiEzsigntemplatepackageID sets FkiEzsigntemplatepackageID field to given value.


### GetFkiEzsigntemplateID

`func (o *EzsigntemplatepackagemembershipResponseCompoundV3) GetFkiEzsigntemplateID() int32`

GetFkiEzsigntemplateID returns the FkiEzsigntemplateID field if non-nil, zero value otherwise.

### GetFkiEzsigntemplateIDOk

`func (o *EzsigntemplatepackagemembershipResponseCompoundV3) GetFkiEzsigntemplateIDOk() (*int32, bool)`

GetFkiEzsigntemplateIDOk returns a tuple with the FkiEzsigntemplateID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsigntemplateID

`func (o *EzsigntemplatepackagemembershipResponseCompoundV3) SetFkiEzsigntemplateID(v int32)`

SetFkiEzsigntemplateID sets FkiEzsigntemplateID field to given value.


### GetIEzsigntemplatepackagemembershipOrder

`func (o *EzsigntemplatepackagemembershipResponseCompoundV3) GetIEzsigntemplatepackagemembershipOrder() int32`

GetIEzsigntemplatepackagemembershipOrder returns the IEzsigntemplatepackagemembershipOrder field if non-nil, zero value otherwise.

### GetIEzsigntemplatepackagemembershipOrderOk

`func (o *EzsigntemplatepackagemembershipResponseCompoundV3) GetIEzsigntemplatepackagemembershipOrderOk() (*int32, bool)`

GetIEzsigntemplatepackagemembershipOrderOk returns a tuple with the IEzsigntemplatepackagemembershipOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplatepackagemembershipOrder

`func (o *EzsigntemplatepackagemembershipResponseCompoundV3) SetIEzsigntemplatepackagemembershipOrder(v int32)`

SetIEzsigntemplatepackagemembershipOrder sets IEzsigntemplatepackagemembershipOrder field to given value.


### GetObjEzsigntemplate

`func (o *EzsigntemplatepackagemembershipResponseCompoundV3) GetObjEzsigntemplate() EzsigntemplateResponseCompoundV4`

GetObjEzsigntemplate returns the ObjEzsigntemplate field if non-nil, zero value otherwise.

### GetObjEzsigntemplateOk

`func (o *EzsigntemplatepackagemembershipResponseCompoundV3) GetObjEzsigntemplateOk() (*EzsigntemplateResponseCompoundV4, bool)`

GetObjEzsigntemplateOk returns a tuple with the ObjEzsigntemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjEzsigntemplate

`func (o *EzsigntemplatepackagemembershipResponseCompoundV3) SetObjEzsigntemplate(v EzsigntemplateResponseCompoundV4)`

SetObjEzsigntemplate sets ObjEzsigntemplate field to given value.


### GetAObjEzsigntemplatepackagesignermembership

`func (o *EzsigntemplatepackagemembershipResponseCompoundV3) GetAObjEzsigntemplatepackagesignermembership() []EzsigntemplatepackagesignermembershipResponseCompound`

GetAObjEzsigntemplatepackagesignermembership returns the AObjEzsigntemplatepackagesignermembership field if non-nil, zero value otherwise.

### GetAObjEzsigntemplatepackagesignermembershipOk

`func (o *EzsigntemplatepackagemembershipResponseCompoundV3) GetAObjEzsigntemplatepackagesignermembershipOk() (*[]EzsigntemplatepackagesignermembershipResponseCompound, bool)`

GetAObjEzsigntemplatepackagesignermembershipOk returns a tuple with the AObjEzsigntemplatepackagesignermembership field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjEzsigntemplatepackagesignermembership

`func (o *EzsigntemplatepackagemembershipResponseCompoundV3) SetAObjEzsigntemplatepackagesignermembership(v []EzsigntemplatepackagesignermembershipResponseCompound)`

SetAObjEzsigntemplatepackagesignermembership sets AObjEzsigntemplatepackagesignermembership field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


