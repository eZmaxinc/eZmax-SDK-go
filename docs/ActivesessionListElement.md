# ActivesessionListElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiActivesessionID** | **int32** | The unique ID of the Activesession | 
**FkiUserID** | **int32** | The unique ID of the User | 
**FkiComputerID** | **int32** | The unique ID of the Computer | 
**FkiCompanyID** | **int32** | The unique ID of the Company | 
**FkiDepartmentID** | **int32** | The unique ID of the Department | 
**SCompanyNameX** | **string** | The Name of the Company in the language of the requester | 
**SDepartmentNameX** | **string** | The Name of the Department in the language of the requester | 
**SActivesessionLoginname** | **string** | The loginname of the Activesession | 
**SComputerDescription** | **string** | The description of the Computer | 
**DtActivesessionFirsthit** | **string** | The first hit of the Activesession | 
**DtActivesessionLasthit** | **string** | The last hit of the Activesession | 
**SActivesessionIP** | **string** | Represent an IP address. | 

## Methods

### NewActivesessionListElement

`func NewActivesessionListElement(pkiActivesessionID int32, fkiUserID int32, fkiComputerID int32, fkiCompanyID int32, fkiDepartmentID int32, sCompanyNameX string, sDepartmentNameX string, sActivesessionLoginname string, sComputerDescription string, dtActivesessionFirsthit string, dtActivesessionLasthit string, sActivesessionIP string, ) *ActivesessionListElement`

NewActivesessionListElement instantiates a new ActivesessionListElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivesessionListElementWithDefaults

`func NewActivesessionListElementWithDefaults() *ActivesessionListElement`

NewActivesessionListElementWithDefaults instantiates a new ActivesessionListElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiActivesessionID

`func (o *ActivesessionListElement) GetPkiActivesessionID() int32`

GetPkiActivesessionID returns the PkiActivesessionID field if non-nil, zero value otherwise.

### GetPkiActivesessionIDOk

`func (o *ActivesessionListElement) GetPkiActivesessionIDOk() (*int32, bool)`

GetPkiActivesessionIDOk returns a tuple with the PkiActivesessionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiActivesessionID

`func (o *ActivesessionListElement) SetPkiActivesessionID(v int32)`

SetPkiActivesessionID sets PkiActivesessionID field to given value.


### GetFkiUserID

`func (o *ActivesessionListElement) GetFkiUserID() int32`

GetFkiUserID returns the FkiUserID field if non-nil, zero value otherwise.

### GetFkiUserIDOk

`func (o *ActivesessionListElement) GetFkiUserIDOk() (*int32, bool)`

GetFkiUserIDOk returns a tuple with the FkiUserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUserID

`func (o *ActivesessionListElement) SetFkiUserID(v int32)`

SetFkiUserID sets FkiUserID field to given value.


### GetFkiComputerID

`func (o *ActivesessionListElement) GetFkiComputerID() int32`

GetFkiComputerID returns the FkiComputerID field if non-nil, zero value otherwise.

### GetFkiComputerIDOk

`func (o *ActivesessionListElement) GetFkiComputerIDOk() (*int32, bool)`

GetFkiComputerIDOk returns a tuple with the FkiComputerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiComputerID

`func (o *ActivesessionListElement) SetFkiComputerID(v int32)`

SetFkiComputerID sets FkiComputerID field to given value.


### GetFkiCompanyID

`func (o *ActivesessionListElement) GetFkiCompanyID() int32`

GetFkiCompanyID returns the FkiCompanyID field if non-nil, zero value otherwise.

### GetFkiCompanyIDOk

`func (o *ActivesessionListElement) GetFkiCompanyIDOk() (*int32, bool)`

GetFkiCompanyIDOk returns a tuple with the FkiCompanyID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiCompanyID

`func (o *ActivesessionListElement) SetFkiCompanyID(v int32)`

SetFkiCompanyID sets FkiCompanyID field to given value.


### GetFkiDepartmentID

`func (o *ActivesessionListElement) GetFkiDepartmentID() int32`

GetFkiDepartmentID returns the FkiDepartmentID field if non-nil, zero value otherwise.

### GetFkiDepartmentIDOk

`func (o *ActivesessionListElement) GetFkiDepartmentIDOk() (*int32, bool)`

GetFkiDepartmentIDOk returns a tuple with the FkiDepartmentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiDepartmentID

`func (o *ActivesessionListElement) SetFkiDepartmentID(v int32)`

SetFkiDepartmentID sets FkiDepartmentID field to given value.


### GetSCompanyNameX

`func (o *ActivesessionListElement) GetSCompanyNameX() string`

GetSCompanyNameX returns the SCompanyNameX field if non-nil, zero value otherwise.

### GetSCompanyNameXOk

`func (o *ActivesessionListElement) GetSCompanyNameXOk() (*string, bool)`

GetSCompanyNameXOk returns a tuple with the SCompanyNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSCompanyNameX

`func (o *ActivesessionListElement) SetSCompanyNameX(v string)`

SetSCompanyNameX sets SCompanyNameX field to given value.


### GetSDepartmentNameX

`func (o *ActivesessionListElement) GetSDepartmentNameX() string`

GetSDepartmentNameX returns the SDepartmentNameX field if non-nil, zero value otherwise.

### GetSDepartmentNameXOk

`func (o *ActivesessionListElement) GetSDepartmentNameXOk() (*string, bool)`

GetSDepartmentNameXOk returns a tuple with the SDepartmentNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSDepartmentNameX

`func (o *ActivesessionListElement) SetSDepartmentNameX(v string)`

SetSDepartmentNameX sets SDepartmentNameX field to given value.


### GetSActivesessionLoginname

`func (o *ActivesessionListElement) GetSActivesessionLoginname() string`

GetSActivesessionLoginname returns the SActivesessionLoginname field if non-nil, zero value otherwise.

### GetSActivesessionLoginnameOk

`func (o *ActivesessionListElement) GetSActivesessionLoginnameOk() (*string, bool)`

GetSActivesessionLoginnameOk returns a tuple with the SActivesessionLoginname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSActivesessionLoginname

`func (o *ActivesessionListElement) SetSActivesessionLoginname(v string)`

SetSActivesessionLoginname sets SActivesessionLoginname field to given value.


### GetSComputerDescription

`func (o *ActivesessionListElement) GetSComputerDescription() string`

GetSComputerDescription returns the SComputerDescription field if non-nil, zero value otherwise.

### GetSComputerDescriptionOk

`func (o *ActivesessionListElement) GetSComputerDescriptionOk() (*string, bool)`

GetSComputerDescriptionOk returns a tuple with the SComputerDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSComputerDescription

`func (o *ActivesessionListElement) SetSComputerDescription(v string)`

SetSComputerDescription sets SComputerDescription field to given value.


### GetDtActivesessionFirsthit

`func (o *ActivesessionListElement) GetDtActivesessionFirsthit() string`

GetDtActivesessionFirsthit returns the DtActivesessionFirsthit field if non-nil, zero value otherwise.

### GetDtActivesessionFirsthitOk

`func (o *ActivesessionListElement) GetDtActivesessionFirsthitOk() (*string, bool)`

GetDtActivesessionFirsthitOk returns a tuple with the DtActivesessionFirsthit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtActivesessionFirsthit

`func (o *ActivesessionListElement) SetDtActivesessionFirsthit(v string)`

SetDtActivesessionFirsthit sets DtActivesessionFirsthit field to given value.


### GetDtActivesessionLasthit

`func (o *ActivesessionListElement) GetDtActivesessionLasthit() string`

GetDtActivesessionLasthit returns the DtActivesessionLasthit field if non-nil, zero value otherwise.

### GetDtActivesessionLasthitOk

`func (o *ActivesessionListElement) GetDtActivesessionLasthitOk() (*string, bool)`

GetDtActivesessionLasthitOk returns a tuple with the DtActivesessionLasthit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtActivesessionLasthit

`func (o *ActivesessionListElement) SetDtActivesessionLasthit(v string)`

SetDtActivesessionLasthit sets DtActivesessionLasthit field to given value.


### GetSActivesessionIP

`func (o *ActivesessionListElement) GetSActivesessionIP() string`

GetSActivesessionIP returns the SActivesessionIP field if non-nil, zero value otherwise.

### GetSActivesessionIPOk

`func (o *ActivesessionListElement) GetSActivesessionIPOk() (*string, bool)`

GetSActivesessionIPOk returns a tuple with the SActivesessionIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSActivesessionIP

`func (o *ActivesessionListElement) SetSActivesessionIP(v string)`

SetSActivesessionIP sets SActivesessionIP field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


