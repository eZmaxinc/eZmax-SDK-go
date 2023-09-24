# CustomFormDataSignerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FkiEzsignfoldersignerassociationID** | **int32** | The unique ID of the Ezsignfoldersignerassociation | 
**FkiUserID** | Pointer to **int32** | The unique ID of the User | [optional] 
**SContactFirstname** | **string** | The First name of the contact | 
**SContactLastname** | **string** | The Last name of the contact | 
**AObjEzsignformfieldgroup** | [**[]CustomFormDataEzsignformfieldgroupResponse**](CustomFormDataEzsignformfieldgroupResponse.md) |  | 

## Methods

### NewCustomFormDataSignerResponse

`func NewCustomFormDataSignerResponse(fkiEzsignfoldersignerassociationID int32, sContactFirstname string, sContactLastname string, aObjEzsignformfieldgroup []CustomFormDataEzsignformfieldgroupResponse, ) *CustomFormDataSignerResponse`

NewCustomFormDataSignerResponse instantiates a new CustomFormDataSignerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomFormDataSignerResponseWithDefaults

`func NewCustomFormDataSignerResponseWithDefaults() *CustomFormDataSignerResponse`

NewCustomFormDataSignerResponseWithDefaults instantiates a new CustomFormDataSignerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFkiEzsignfoldersignerassociationID

`func (o *CustomFormDataSignerResponse) GetFkiEzsignfoldersignerassociationID() int32`

GetFkiEzsignfoldersignerassociationID returns the FkiEzsignfoldersignerassociationID field if non-nil, zero value otherwise.

### GetFkiEzsignfoldersignerassociationIDOk

`func (o *CustomFormDataSignerResponse) GetFkiEzsignfoldersignerassociationIDOk() (*int32, bool)`

GetFkiEzsignfoldersignerassociationIDOk returns a tuple with the FkiEzsignfoldersignerassociationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignfoldersignerassociationID

`func (o *CustomFormDataSignerResponse) SetFkiEzsignfoldersignerassociationID(v int32)`

SetFkiEzsignfoldersignerassociationID sets FkiEzsignfoldersignerassociationID field to given value.


### GetFkiUserID

`func (o *CustomFormDataSignerResponse) GetFkiUserID() int32`

GetFkiUserID returns the FkiUserID field if non-nil, zero value otherwise.

### GetFkiUserIDOk

`func (o *CustomFormDataSignerResponse) GetFkiUserIDOk() (*int32, bool)`

GetFkiUserIDOk returns a tuple with the FkiUserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUserID

`func (o *CustomFormDataSignerResponse) SetFkiUserID(v int32)`

SetFkiUserID sets FkiUserID field to given value.

### HasFkiUserID

`func (o *CustomFormDataSignerResponse) HasFkiUserID() bool`

HasFkiUserID returns a boolean if a field has been set.

### GetSContactFirstname

`func (o *CustomFormDataSignerResponse) GetSContactFirstname() string`

GetSContactFirstname returns the SContactFirstname field if non-nil, zero value otherwise.

### GetSContactFirstnameOk

`func (o *CustomFormDataSignerResponse) GetSContactFirstnameOk() (*string, bool)`

GetSContactFirstnameOk returns a tuple with the SContactFirstname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSContactFirstname

`func (o *CustomFormDataSignerResponse) SetSContactFirstname(v string)`

SetSContactFirstname sets SContactFirstname field to given value.


### GetSContactLastname

`func (o *CustomFormDataSignerResponse) GetSContactLastname() string`

GetSContactLastname returns the SContactLastname field if non-nil, zero value otherwise.

### GetSContactLastnameOk

`func (o *CustomFormDataSignerResponse) GetSContactLastnameOk() (*string, bool)`

GetSContactLastnameOk returns a tuple with the SContactLastname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSContactLastname

`func (o *CustomFormDataSignerResponse) SetSContactLastname(v string)`

SetSContactLastname sets SContactLastname field to given value.


### GetAObjEzsignformfieldgroup

`func (o *CustomFormDataSignerResponse) GetAObjEzsignformfieldgroup() []CustomFormDataEzsignformfieldgroupResponse`

GetAObjEzsignformfieldgroup returns the AObjEzsignformfieldgroup field if non-nil, zero value otherwise.

### GetAObjEzsignformfieldgroupOk

`func (o *CustomFormDataSignerResponse) GetAObjEzsignformfieldgroupOk() (*[]CustomFormDataEzsignformfieldgroupResponse, bool)`

GetAObjEzsignformfieldgroupOk returns a tuple with the AObjEzsignformfieldgroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjEzsignformfieldgroup

`func (o *CustomFormDataSignerResponse) SetAObjEzsignformfieldgroup(v []CustomFormDataEzsignformfieldgroupResponse)`

SetAObjEzsignformfieldgroup sets AObjEzsignformfieldgroup field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


