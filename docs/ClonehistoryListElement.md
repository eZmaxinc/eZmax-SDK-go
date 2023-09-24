# ClonehistoryListElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiClonehistoryID** | **int32** | The unique ID of the Clonehistory | 
**FkiUserIDCloning** | **int32** | The unique ID of the User | 
**FkiUserIDCloned** | **int32** | The unique ID of the User | 
**DtClonehistoryFirsthit** | **string** | The firsthit of the Clonehistory | 
**DtClonehistoryLasthit** | Pointer to **string** | The lasthit of the Clonehistory | [optional] 
**SUserLoginnameCloning** | **string** | The login name of the User. | 
**SUserFirstnameCloning** | **string** | The first name of the user | 
**SUserLastnameCloning** | **string** | The last name of the user | 
**SUserLoginnameCloned** | **string** | The login name of the User. | 
**SUserFirstnameCloned** | **string** | The first name of the user | 
**SUserLastnameCloned** | **string** | The last name of the user | 

## Methods

### NewClonehistoryListElement

`func NewClonehistoryListElement(pkiClonehistoryID int32, fkiUserIDCloning int32, fkiUserIDCloned int32, dtClonehistoryFirsthit string, sUserLoginnameCloning string, sUserFirstnameCloning string, sUserLastnameCloning string, sUserLoginnameCloned string, sUserFirstnameCloned string, sUserLastnameCloned string, ) *ClonehistoryListElement`

NewClonehistoryListElement instantiates a new ClonehistoryListElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClonehistoryListElementWithDefaults

`func NewClonehistoryListElementWithDefaults() *ClonehistoryListElement`

NewClonehistoryListElementWithDefaults instantiates a new ClonehistoryListElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiClonehistoryID

`func (o *ClonehistoryListElement) GetPkiClonehistoryID() int32`

GetPkiClonehistoryID returns the PkiClonehistoryID field if non-nil, zero value otherwise.

### GetPkiClonehistoryIDOk

`func (o *ClonehistoryListElement) GetPkiClonehistoryIDOk() (*int32, bool)`

GetPkiClonehistoryIDOk returns a tuple with the PkiClonehistoryID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiClonehistoryID

`func (o *ClonehistoryListElement) SetPkiClonehistoryID(v int32)`

SetPkiClonehistoryID sets PkiClonehistoryID field to given value.


### GetFkiUserIDCloning

`func (o *ClonehistoryListElement) GetFkiUserIDCloning() int32`

GetFkiUserIDCloning returns the FkiUserIDCloning field if non-nil, zero value otherwise.

### GetFkiUserIDCloningOk

`func (o *ClonehistoryListElement) GetFkiUserIDCloningOk() (*int32, bool)`

GetFkiUserIDCloningOk returns a tuple with the FkiUserIDCloning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUserIDCloning

`func (o *ClonehistoryListElement) SetFkiUserIDCloning(v int32)`

SetFkiUserIDCloning sets FkiUserIDCloning field to given value.


### GetFkiUserIDCloned

`func (o *ClonehistoryListElement) GetFkiUserIDCloned() int32`

GetFkiUserIDCloned returns the FkiUserIDCloned field if non-nil, zero value otherwise.

### GetFkiUserIDClonedOk

`func (o *ClonehistoryListElement) GetFkiUserIDClonedOk() (*int32, bool)`

GetFkiUserIDClonedOk returns a tuple with the FkiUserIDCloned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUserIDCloned

`func (o *ClonehistoryListElement) SetFkiUserIDCloned(v int32)`

SetFkiUserIDCloned sets FkiUserIDCloned field to given value.


### GetDtClonehistoryFirsthit

`func (o *ClonehistoryListElement) GetDtClonehistoryFirsthit() string`

GetDtClonehistoryFirsthit returns the DtClonehistoryFirsthit field if non-nil, zero value otherwise.

### GetDtClonehistoryFirsthitOk

`func (o *ClonehistoryListElement) GetDtClonehistoryFirsthitOk() (*string, bool)`

GetDtClonehistoryFirsthitOk returns a tuple with the DtClonehistoryFirsthit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtClonehistoryFirsthit

`func (o *ClonehistoryListElement) SetDtClonehistoryFirsthit(v string)`

SetDtClonehistoryFirsthit sets DtClonehistoryFirsthit field to given value.


### GetDtClonehistoryLasthit

`func (o *ClonehistoryListElement) GetDtClonehistoryLasthit() string`

GetDtClonehistoryLasthit returns the DtClonehistoryLasthit field if non-nil, zero value otherwise.

### GetDtClonehistoryLasthitOk

`func (o *ClonehistoryListElement) GetDtClonehistoryLasthitOk() (*string, bool)`

GetDtClonehistoryLasthitOk returns a tuple with the DtClonehistoryLasthit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtClonehistoryLasthit

`func (o *ClonehistoryListElement) SetDtClonehistoryLasthit(v string)`

SetDtClonehistoryLasthit sets DtClonehistoryLasthit field to given value.

### HasDtClonehistoryLasthit

`func (o *ClonehistoryListElement) HasDtClonehistoryLasthit() bool`

HasDtClonehistoryLasthit returns a boolean if a field has been set.

### GetSUserLoginnameCloning

`func (o *ClonehistoryListElement) GetSUserLoginnameCloning() string`

GetSUserLoginnameCloning returns the SUserLoginnameCloning field if non-nil, zero value otherwise.

### GetSUserLoginnameCloningOk

`func (o *ClonehistoryListElement) GetSUserLoginnameCloningOk() (*string, bool)`

GetSUserLoginnameCloningOk returns a tuple with the SUserLoginnameCloning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSUserLoginnameCloning

`func (o *ClonehistoryListElement) SetSUserLoginnameCloning(v string)`

SetSUserLoginnameCloning sets SUserLoginnameCloning field to given value.


### GetSUserFirstnameCloning

`func (o *ClonehistoryListElement) GetSUserFirstnameCloning() string`

GetSUserFirstnameCloning returns the SUserFirstnameCloning field if non-nil, zero value otherwise.

### GetSUserFirstnameCloningOk

`func (o *ClonehistoryListElement) GetSUserFirstnameCloningOk() (*string, bool)`

GetSUserFirstnameCloningOk returns a tuple with the SUserFirstnameCloning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSUserFirstnameCloning

`func (o *ClonehistoryListElement) SetSUserFirstnameCloning(v string)`

SetSUserFirstnameCloning sets SUserFirstnameCloning field to given value.


### GetSUserLastnameCloning

`func (o *ClonehistoryListElement) GetSUserLastnameCloning() string`

GetSUserLastnameCloning returns the SUserLastnameCloning field if non-nil, zero value otherwise.

### GetSUserLastnameCloningOk

`func (o *ClonehistoryListElement) GetSUserLastnameCloningOk() (*string, bool)`

GetSUserLastnameCloningOk returns a tuple with the SUserLastnameCloning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSUserLastnameCloning

`func (o *ClonehistoryListElement) SetSUserLastnameCloning(v string)`

SetSUserLastnameCloning sets SUserLastnameCloning field to given value.


### GetSUserLoginnameCloned

`func (o *ClonehistoryListElement) GetSUserLoginnameCloned() string`

GetSUserLoginnameCloned returns the SUserLoginnameCloned field if non-nil, zero value otherwise.

### GetSUserLoginnameClonedOk

`func (o *ClonehistoryListElement) GetSUserLoginnameClonedOk() (*string, bool)`

GetSUserLoginnameClonedOk returns a tuple with the SUserLoginnameCloned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSUserLoginnameCloned

`func (o *ClonehistoryListElement) SetSUserLoginnameCloned(v string)`

SetSUserLoginnameCloned sets SUserLoginnameCloned field to given value.


### GetSUserFirstnameCloned

`func (o *ClonehistoryListElement) GetSUserFirstnameCloned() string`

GetSUserFirstnameCloned returns the SUserFirstnameCloned field if non-nil, zero value otherwise.

### GetSUserFirstnameClonedOk

`func (o *ClonehistoryListElement) GetSUserFirstnameClonedOk() (*string, bool)`

GetSUserFirstnameClonedOk returns a tuple with the SUserFirstnameCloned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSUserFirstnameCloned

`func (o *ClonehistoryListElement) SetSUserFirstnameCloned(v string)`

SetSUserFirstnameCloned sets SUserFirstnameCloned field to given value.


### GetSUserLastnameCloned

`func (o *ClonehistoryListElement) GetSUserLastnameCloned() string`

GetSUserLastnameCloned returns the SUserLastnameCloned field if non-nil, zero value otherwise.

### GetSUserLastnameClonedOk

`func (o *ClonehistoryListElement) GetSUserLastnameClonedOk() (*string, bool)`

GetSUserLastnameClonedOk returns a tuple with the SUserLastnameCloned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSUserLastnameCloned

`func (o *ClonehistoryListElement) SetSUserLastnameCloned(v string)`

SetSUserLastnameCloned sets SUserLastnameCloned field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


