# EzsigntemplatepackageListElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsigntemplatepackageID** | **int32** | The unique ID of the Ezsigntemplatepackage | 
**FkiDepartmentID** | **NullableInt32** | The unique ID of the Department. | 
**FkiTeamID** | **NullableInt32** | The unique ID of the Team | 
**FkiEzsignfoldertypeID** | **NullableInt32** | The unique ID of the Ezsignfoldertype. | 
**FkiLanguageID** | **int32** | The unique ID of the Language.  Valid values:  |Value|Description| |-|-| |1|French| |2|English| | 
**EEzsigntemplatepackageType** | [**FieldEEzsigntemplatepackageType**](FieldEEzsigntemplatepackageType.md) |  | 
**SEzsigntemplatepackageDescription** | **string** | The description of the Ezsigntemplatepackage | 
**BEzsigntemplatepackageIsactive** | **bool** | Whether the Ezsigntemplatepackage is active or not | 
**IEzsigntemplatepackagemembership** | **int32** | The total number of Ezsigntemplatepackagemembership in the Ezsigntemplatepackage | 

## Methods

### NewEzsigntemplatepackageListElement

`func NewEzsigntemplatepackageListElement(pkiEzsigntemplatepackageID int32, fkiDepartmentID NullableInt32, fkiTeamID NullableInt32, fkiEzsignfoldertypeID NullableInt32, fkiLanguageID int32, eEzsigntemplatepackageType FieldEEzsigntemplatepackageType, sEzsigntemplatepackageDescription string, bEzsigntemplatepackageIsactive bool, iEzsigntemplatepackagemembership int32, ) *EzsigntemplatepackageListElement`

NewEzsigntemplatepackageListElement instantiates a new EzsigntemplatepackageListElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsigntemplatepackageListElementWithDefaults

`func NewEzsigntemplatepackageListElementWithDefaults() *EzsigntemplatepackageListElement`

NewEzsigntemplatepackageListElementWithDefaults instantiates a new EzsigntemplatepackageListElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsigntemplatepackageID

`func (o *EzsigntemplatepackageListElement) GetPkiEzsigntemplatepackageID() int32`

GetPkiEzsigntemplatepackageID returns the PkiEzsigntemplatepackageID field if non-nil, zero value otherwise.

### GetPkiEzsigntemplatepackageIDOk

`func (o *EzsigntemplatepackageListElement) GetPkiEzsigntemplatepackageIDOk() (*int32, bool)`

GetPkiEzsigntemplatepackageIDOk returns a tuple with the PkiEzsigntemplatepackageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsigntemplatepackageID

`func (o *EzsigntemplatepackageListElement) SetPkiEzsigntemplatepackageID(v int32)`

SetPkiEzsigntemplatepackageID sets PkiEzsigntemplatepackageID field to given value.


### GetFkiDepartmentID

`func (o *EzsigntemplatepackageListElement) GetFkiDepartmentID() int32`

GetFkiDepartmentID returns the FkiDepartmentID field if non-nil, zero value otherwise.

### GetFkiDepartmentIDOk

`func (o *EzsigntemplatepackageListElement) GetFkiDepartmentIDOk() (*int32, bool)`

GetFkiDepartmentIDOk returns a tuple with the FkiDepartmentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiDepartmentID

`func (o *EzsigntemplatepackageListElement) SetFkiDepartmentID(v int32)`

SetFkiDepartmentID sets FkiDepartmentID field to given value.


### SetFkiDepartmentIDNil

`func (o *EzsigntemplatepackageListElement) SetFkiDepartmentIDNil(b bool)`

 SetFkiDepartmentIDNil sets the value for FkiDepartmentID to be an explicit nil

### UnsetFkiDepartmentID
`func (o *EzsigntemplatepackageListElement) UnsetFkiDepartmentID()`

UnsetFkiDepartmentID ensures that no value is present for FkiDepartmentID, not even an explicit nil
### GetFkiTeamID

`func (o *EzsigntemplatepackageListElement) GetFkiTeamID() int32`

GetFkiTeamID returns the FkiTeamID field if non-nil, zero value otherwise.

### GetFkiTeamIDOk

`func (o *EzsigntemplatepackageListElement) GetFkiTeamIDOk() (*int32, bool)`

GetFkiTeamIDOk returns a tuple with the FkiTeamID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiTeamID

`func (o *EzsigntemplatepackageListElement) SetFkiTeamID(v int32)`

SetFkiTeamID sets FkiTeamID field to given value.


### SetFkiTeamIDNil

`func (o *EzsigntemplatepackageListElement) SetFkiTeamIDNil(b bool)`

 SetFkiTeamIDNil sets the value for FkiTeamID to be an explicit nil

### UnsetFkiTeamID
`func (o *EzsigntemplatepackageListElement) UnsetFkiTeamID()`

UnsetFkiTeamID ensures that no value is present for FkiTeamID, not even an explicit nil
### GetFkiEzsignfoldertypeID

`func (o *EzsigntemplatepackageListElement) GetFkiEzsignfoldertypeID() int32`

GetFkiEzsignfoldertypeID returns the FkiEzsignfoldertypeID field if non-nil, zero value otherwise.

### GetFkiEzsignfoldertypeIDOk

`func (o *EzsigntemplatepackageListElement) GetFkiEzsignfoldertypeIDOk() (*int32, bool)`

GetFkiEzsignfoldertypeIDOk returns a tuple with the FkiEzsignfoldertypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignfoldertypeID

`func (o *EzsigntemplatepackageListElement) SetFkiEzsignfoldertypeID(v int32)`

SetFkiEzsignfoldertypeID sets FkiEzsignfoldertypeID field to given value.


### SetFkiEzsignfoldertypeIDNil

`func (o *EzsigntemplatepackageListElement) SetFkiEzsignfoldertypeIDNil(b bool)`

 SetFkiEzsignfoldertypeIDNil sets the value for FkiEzsignfoldertypeID to be an explicit nil

### UnsetFkiEzsignfoldertypeID
`func (o *EzsigntemplatepackageListElement) UnsetFkiEzsignfoldertypeID()`

UnsetFkiEzsignfoldertypeID ensures that no value is present for FkiEzsignfoldertypeID, not even an explicit nil
### GetFkiLanguageID

`func (o *EzsigntemplatepackageListElement) GetFkiLanguageID() int32`

GetFkiLanguageID returns the FkiLanguageID field if non-nil, zero value otherwise.

### GetFkiLanguageIDOk

`func (o *EzsigntemplatepackageListElement) GetFkiLanguageIDOk() (*int32, bool)`

GetFkiLanguageIDOk returns a tuple with the FkiLanguageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiLanguageID

`func (o *EzsigntemplatepackageListElement) SetFkiLanguageID(v int32)`

SetFkiLanguageID sets FkiLanguageID field to given value.


### GetEEzsigntemplatepackageType

`func (o *EzsigntemplatepackageListElement) GetEEzsigntemplatepackageType() FieldEEzsigntemplatepackageType`

GetEEzsigntemplatepackageType returns the EEzsigntemplatepackageType field if non-nil, zero value otherwise.

### GetEEzsigntemplatepackageTypeOk

`func (o *EzsigntemplatepackageListElement) GetEEzsigntemplatepackageTypeOk() (*FieldEEzsigntemplatepackageType, bool)`

GetEEzsigntemplatepackageTypeOk returns a tuple with the EEzsigntemplatepackageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigntemplatepackageType

`func (o *EzsigntemplatepackageListElement) SetEEzsigntemplatepackageType(v FieldEEzsigntemplatepackageType)`

SetEEzsigntemplatepackageType sets EEzsigntemplatepackageType field to given value.


### GetSEzsigntemplatepackageDescription

`func (o *EzsigntemplatepackageListElement) GetSEzsigntemplatepackageDescription() string`

GetSEzsigntemplatepackageDescription returns the SEzsigntemplatepackageDescription field if non-nil, zero value otherwise.

### GetSEzsigntemplatepackageDescriptionOk

`func (o *EzsigntemplatepackageListElement) GetSEzsigntemplatepackageDescriptionOk() (*string, bool)`

GetSEzsigntemplatepackageDescriptionOk returns a tuple with the SEzsigntemplatepackageDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigntemplatepackageDescription

`func (o *EzsigntemplatepackageListElement) SetSEzsigntemplatepackageDescription(v string)`

SetSEzsigntemplatepackageDescription sets SEzsigntemplatepackageDescription field to given value.


### GetBEzsigntemplatepackageIsactive

`func (o *EzsigntemplatepackageListElement) GetBEzsigntemplatepackageIsactive() bool`

GetBEzsigntemplatepackageIsactive returns the BEzsigntemplatepackageIsactive field if non-nil, zero value otherwise.

### GetBEzsigntemplatepackageIsactiveOk

`func (o *EzsigntemplatepackageListElement) GetBEzsigntemplatepackageIsactiveOk() (*bool, bool)`

GetBEzsigntemplatepackageIsactiveOk returns a tuple with the BEzsigntemplatepackageIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsigntemplatepackageIsactive

`func (o *EzsigntemplatepackageListElement) SetBEzsigntemplatepackageIsactive(v bool)`

SetBEzsigntemplatepackageIsactive sets BEzsigntemplatepackageIsactive field to given value.


### GetIEzsigntemplatepackagemembership

`func (o *EzsigntemplatepackageListElement) GetIEzsigntemplatepackagemembership() int32`

GetIEzsigntemplatepackagemembership returns the IEzsigntemplatepackagemembership field if non-nil, zero value otherwise.

### GetIEzsigntemplatepackagemembershipOk

`func (o *EzsigntemplatepackageListElement) GetIEzsigntemplatepackagemembershipOk() (*int32, bool)`

GetIEzsigntemplatepackagemembershipOk returns a tuple with the IEzsigntemplatepackagemembership field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigntemplatepackagemembership

`func (o *EzsigntemplatepackageListElement) SetIEzsigntemplatepackagemembership(v int32)`

SetIEzsigntemplatepackagemembership sets IEzsigntemplatepackagemembership field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


