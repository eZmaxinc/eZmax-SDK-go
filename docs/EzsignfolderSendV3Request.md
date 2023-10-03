# EzsignfolderSendV3Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TEzsignfolderMessage** | Pointer to **string** | A custom text message that will be added to the email sent. | [optional] 
**DtEzsignfolderDelayedsenddate** | Pointer to **string** | The date and time at which the Ezsignfolder will be sent in the future. | [optional] 
**AFkiEzsignfoldersignerassociationID** | **[]int32** |  | 

## Methods

### NewEzsignfolderSendV3Request

`func NewEzsignfolderSendV3Request(aFkiEzsignfoldersignerassociationID []int32, ) *EzsignfolderSendV3Request`

NewEzsignfolderSendV3Request instantiates a new EzsignfolderSendV3Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsignfolderSendV3RequestWithDefaults

`func NewEzsignfolderSendV3RequestWithDefaults() *EzsignfolderSendV3Request`

NewEzsignfolderSendV3RequestWithDefaults instantiates a new EzsignfolderSendV3Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTEzsignfolderMessage

`func (o *EzsignfolderSendV3Request) GetTEzsignfolderMessage() string`

GetTEzsignfolderMessage returns the TEzsignfolderMessage field if non-nil, zero value otherwise.

### GetTEzsignfolderMessageOk

`func (o *EzsignfolderSendV3Request) GetTEzsignfolderMessageOk() (*string, bool)`

GetTEzsignfolderMessageOk returns a tuple with the TEzsignfolderMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTEzsignfolderMessage

`func (o *EzsignfolderSendV3Request) SetTEzsignfolderMessage(v string)`

SetTEzsignfolderMessage sets TEzsignfolderMessage field to given value.

### HasTEzsignfolderMessage

`func (o *EzsignfolderSendV3Request) HasTEzsignfolderMessage() bool`

HasTEzsignfolderMessage returns a boolean if a field has been set.

### GetDtEzsignfolderDelayedsenddate

`func (o *EzsignfolderSendV3Request) GetDtEzsignfolderDelayedsenddate() string`

GetDtEzsignfolderDelayedsenddate returns the DtEzsignfolderDelayedsenddate field if non-nil, zero value otherwise.

### GetDtEzsignfolderDelayedsenddateOk

`func (o *EzsignfolderSendV3Request) GetDtEzsignfolderDelayedsenddateOk() (*string, bool)`

GetDtEzsignfolderDelayedsenddateOk returns a tuple with the DtEzsignfolderDelayedsenddate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtEzsignfolderDelayedsenddate

`func (o *EzsignfolderSendV3Request) SetDtEzsignfolderDelayedsenddate(v string)`

SetDtEzsignfolderDelayedsenddate sets DtEzsignfolderDelayedsenddate field to given value.

### HasDtEzsignfolderDelayedsenddate

`func (o *EzsignfolderSendV3Request) HasDtEzsignfolderDelayedsenddate() bool`

HasDtEzsignfolderDelayedsenddate returns a boolean if a field has been set.

### GetAFkiEzsignfoldersignerassociationID

`func (o *EzsignfolderSendV3Request) GetAFkiEzsignfoldersignerassociationID() []int32`

GetAFkiEzsignfoldersignerassociationID returns the AFkiEzsignfoldersignerassociationID field if non-nil, zero value otherwise.

### GetAFkiEzsignfoldersignerassociationIDOk

`func (o *EzsignfolderSendV3Request) GetAFkiEzsignfoldersignerassociationIDOk() (*[]int32, bool)`

GetAFkiEzsignfoldersignerassociationIDOk returns a tuple with the AFkiEzsignfoldersignerassociationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAFkiEzsignfoldersignerassociationID

`func (o *EzsignfolderSendV3Request) SetAFkiEzsignfoldersignerassociationID(v []int32)`

SetAFkiEzsignfoldersignerassociationID sets AFkiEzsignfoldersignerassociationID field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


