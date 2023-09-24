# CustomEzsignfoldersignerassociationstatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FkiEzsignfoldersignerassociationID** | **int32** | The unique ID of the Ezsignfoldersignerassociation | 
**SEzsignfoldersignerassociationstatusLastname** | Pointer to **string** | The last name of the Ezsignsigner | [optional] 
**SEzsignfoldersignerassociationstatusFirstname** | Pointer to **string** | The first name of the Ezsignsigner | [optional] 
**SEzsignfoldersignerassociationstatusDescriptionX** | Pointer to **string** | The description of the Ezsignsigner | [optional] 
**AObjEzsignsignaturestatus** | [**[]CustomEzsignsignaturestatusResponse**](CustomEzsignsignaturestatusResponse.md) |  | 

## Methods

### NewCustomEzsignfoldersignerassociationstatusResponse

`func NewCustomEzsignfoldersignerassociationstatusResponse(fkiEzsignfoldersignerassociationID int32, aObjEzsignsignaturestatus []CustomEzsignsignaturestatusResponse, ) *CustomEzsignfoldersignerassociationstatusResponse`

NewCustomEzsignfoldersignerassociationstatusResponse instantiates a new CustomEzsignfoldersignerassociationstatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomEzsignfoldersignerassociationstatusResponseWithDefaults

`func NewCustomEzsignfoldersignerassociationstatusResponseWithDefaults() *CustomEzsignfoldersignerassociationstatusResponse`

NewCustomEzsignfoldersignerassociationstatusResponseWithDefaults instantiates a new CustomEzsignfoldersignerassociationstatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFkiEzsignfoldersignerassociationID

`func (o *CustomEzsignfoldersignerassociationstatusResponse) GetFkiEzsignfoldersignerassociationID() int32`

GetFkiEzsignfoldersignerassociationID returns the FkiEzsignfoldersignerassociationID field if non-nil, zero value otherwise.

### GetFkiEzsignfoldersignerassociationIDOk

`func (o *CustomEzsignfoldersignerassociationstatusResponse) GetFkiEzsignfoldersignerassociationIDOk() (*int32, bool)`

GetFkiEzsignfoldersignerassociationIDOk returns a tuple with the FkiEzsignfoldersignerassociationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignfoldersignerassociationID

`func (o *CustomEzsignfoldersignerassociationstatusResponse) SetFkiEzsignfoldersignerassociationID(v int32)`

SetFkiEzsignfoldersignerassociationID sets FkiEzsignfoldersignerassociationID field to given value.


### GetSEzsignfoldersignerassociationstatusLastname

`func (o *CustomEzsignfoldersignerassociationstatusResponse) GetSEzsignfoldersignerassociationstatusLastname() string`

GetSEzsignfoldersignerassociationstatusLastname returns the SEzsignfoldersignerassociationstatusLastname field if non-nil, zero value otherwise.

### GetSEzsignfoldersignerassociationstatusLastnameOk

`func (o *CustomEzsignfoldersignerassociationstatusResponse) GetSEzsignfoldersignerassociationstatusLastnameOk() (*string, bool)`

GetSEzsignfoldersignerassociationstatusLastnameOk returns a tuple with the SEzsignfoldersignerassociationstatusLastname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsignfoldersignerassociationstatusLastname

`func (o *CustomEzsignfoldersignerassociationstatusResponse) SetSEzsignfoldersignerassociationstatusLastname(v string)`

SetSEzsignfoldersignerassociationstatusLastname sets SEzsignfoldersignerassociationstatusLastname field to given value.

### HasSEzsignfoldersignerassociationstatusLastname

`func (o *CustomEzsignfoldersignerassociationstatusResponse) HasSEzsignfoldersignerassociationstatusLastname() bool`

HasSEzsignfoldersignerassociationstatusLastname returns a boolean if a field has been set.

### GetSEzsignfoldersignerassociationstatusFirstname

`func (o *CustomEzsignfoldersignerassociationstatusResponse) GetSEzsignfoldersignerassociationstatusFirstname() string`

GetSEzsignfoldersignerassociationstatusFirstname returns the SEzsignfoldersignerassociationstatusFirstname field if non-nil, zero value otherwise.

### GetSEzsignfoldersignerassociationstatusFirstnameOk

`func (o *CustomEzsignfoldersignerassociationstatusResponse) GetSEzsignfoldersignerassociationstatusFirstnameOk() (*string, bool)`

GetSEzsignfoldersignerassociationstatusFirstnameOk returns a tuple with the SEzsignfoldersignerassociationstatusFirstname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsignfoldersignerassociationstatusFirstname

`func (o *CustomEzsignfoldersignerassociationstatusResponse) SetSEzsignfoldersignerassociationstatusFirstname(v string)`

SetSEzsignfoldersignerassociationstatusFirstname sets SEzsignfoldersignerassociationstatusFirstname field to given value.

### HasSEzsignfoldersignerassociationstatusFirstname

`func (o *CustomEzsignfoldersignerassociationstatusResponse) HasSEzsignfoldersignerassociationstatusFirstname() bool`

HasSEzsignfoldersignerassociationstatusFirstname returns a boolean if a field has been set.

### GetSEzsignfoldersignerassociationstatusDescriptionX

`func (o *CustomEzsignfoldersignerassociationstatusResponse) GetSEzsignfoldersignerassociationstatusDescriptionX() string`

GetSEzsignfoldersignerassociationstatusDescriptionX returns the SEzsignfoldersignerassociationstatusDescriptionX field if non-nil, zero value otherwise.

### GetSEzsignfoldersignerassociationstatusDescriptionXOk

`func (o *CustomEzsignfoldersignerassociationstatusResponse) GetSEzsignfoldersignerassociationstatusDescriptionXOk() (*string, bool)`

GetSEzsignfoldersignerassociationstatusDescriptionXOk returns a tuple with the SEzsignfoldersignerassociationstatusDescriptionX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsignfoldersignerassociationstatusDescriptionX

`func (o *CustomEzsignfoldersignerassociationstatusResponse) SetSEzsignfoldersignerassociationstatusDescriptionX(v string)`

SetSEzsignfoldersignerassociationstatusDescriptionX sets SEzsignfoldersignerassociationstatusDescriptionX field to given value.

### HasSEzsignfoldersignerassociationstatusDescriptionX

`func (o *CustomEzsignfoldersignerassociationstatusResponse) HasSEzsignfoldersignerassociationstatusDescriptionX() bool`

HasSEzsignfoldersignerassociationstatusDescriptionX returns a boolean if a field has been set.

### GetAObjEzsignsignaturestatus

`func (o *CustomEzsignfoldersignerassociationstatusResponse) GetAObjEzsignsignaturestatus() []CustomEzsignsignaturestatusResponse`

GetAObjEzsignsignaturestatus returns the AObjEzsignsignaturestatus field if non-nil, zero value otherwise.

### GetAObjEzsignsignaturestatusOk

`func (o *CustomEzsignfoldersignerassociationstatusResponse) GetAObjEzsignsignaturestatusOk() (*[]CustomEzsignsignaturestatusResponse, bool)`

GetAObjEzsignsignaturestatusOk returns a tuple with the AObjEzsignsignaturestatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjEzsignsignaturestatus

`func (o *CustomEzsignfoldersignerassociationstatusResponse) SetAObjEzsignsignaturestatus(v []CustomEzsignsignaturestatusResponse)`

SetAObjEzsignsignaturestatus sets AObjEzsignsignaturestatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


