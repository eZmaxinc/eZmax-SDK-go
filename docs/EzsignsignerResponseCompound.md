# EzsignsignerResponseCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjContact** | [**EzsignsignerResponseCompoundContact**](EzsignsignerResponseCompoundContact.md) |  | 
**PkiEzsignsignerID** | **int32** | The unique ID of the Ezsignsigner | 
**FkiTaxassignmentID** | **int32** | The unique ID of the Taxassignment.  Valid values:  |Value|Description| |-|-| |1|No tax| |2|GST| |3|HST (ON)| |4|HST (NB)| |5|HST (NS)| |6|HST (NL)| |7|HST (PE)| |8|GST + QST (QC)| |9|GST + QST (QC) Non-Recoverable| |10|GST + PST (BC)| |11|GST + PST (SK)| |12|GST + RST (MB)| |13|GST + PST (BC) Non-Recoverable| |14|GST + PST (SK) Non-Recoverable| |15|GST + RST (MB) Non-Recoverable| | 
**FkiSecretquestionID** | Pointer to **int32** | The unique ID of the Secretquestion.  Valid values:  |Value|Description| |-|-| |1|The name of the hospital in which you were born| |2|The name of your grade school| |3|The last name of your favorite teacher| |4|Your favorite sports team| |5|Your favorite TV show| |6|Your favorite movie| |7|The name of the street on which you grew up| |8|The name of your first employer| |9|Your first car| |10|Your favorite food| |11|The name of your first pet| |12|Favorite musician/band| |13|What instrument you play| |14|Your father&#39;s middle name| |15|Your mother&#39;s maiden name| |16|Name of your eldest child| |17|Your spouse&#39;s middle name| |18|Favorite restaurant| |19|Childhood nickname| |20|Favorite vacation destination| |21|Your boat&#39;s name| |22|Date of Birth (YYYY-MM-DD)| | [optional] 
**FkiUserlogintypeID** | **int32** | The unique ID of the Userlogintype | 
**SUserlogintypeDescriptionX** | **string** | The description of the Userlogintype in the language of the requester | 

## Methods

### NewEzsignsignerResponseCompound

`func NewEzsignsignerResponseCompound(objContact EzsignsignerResponseCompoundContact, pkiEzsignsignerID int32, fkiTaxassignmentID int32, fkiUserlogintypeID int32, sUserlogintypeDescriptionX string, ) *EzsignsignerResponseCompound`

NewEzsignsignerResponseCompound instantiates a new EzsignsignerResponseCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsignsignerResponseCompoundWithDefaults

`func NewEzsignsignerResponseCompoundWithDefaults() *EzsignsignerResponseCompound`

NewEzsignsignerResponseCompoundWithDefaults instantiates a new EzsignsignerResponseCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjContact

`func (o *EzsignsignerResponseCompound) GetObjContact() EzsignsignerResponseCompoundContact`

GetObjContact returns the ObjContact field if non-nil, zero value otherwise.

### GetObjContactOk

`func (o *EzsignsignerResponseCompound) GetObjContactOk() (*EzsignsignerResponseCompoundContact, bool)`

GetObjContactOk returns a tuple with the ObjContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjContact

`func (o *EzsignsignerResponseCompound) SetObjContact(v EzsignsignerResponseCompoundContact)`

SetObjContact sets ObjContact field to given value.


### GetPkiEzsignsignerID

`func (o *EzsignsignerResponseCompound) GetPkiEzsignsignerID() int32`

GetPkiEzsignsignerID returns the PkiEzsignsignerID field if non-nil, zero value otherwise.

### GetPkiEzsignsignerIDOk

`func (o *EzsignsignerResponseCompound) GetPkiEzsignsignerIDOk() (*int32, bool)`

GetPkiEzsignsignerIDOk returns a tuple with the PkiEzsignsignerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsignsignerID

`func (o *EzsignsignerResponseCompound) SetPkiEzsignsignerID(v int32)`

SetPkiEzsignsignerID sets PkiEzsignsignerID field to given value.


### GetFkiTaxassignmentID

`func (o *EzsignsignerResponseCompound) GetFkiTaxassignmentID() int32`

GetFkiTaxassignmentID returns the FkiTaxassignmentID field if non-nil, zero value otherwise.

### GetFkiTaxassignmentIDOk

`func (o *EzsignsignerResponseCompound) GetFkiTaxassignmentIDOk() (*int32, bool)`

GetFkiTaxassignmentIDOk returns a tuple with the FkiTaxassignmentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiTaxassignmentID

`func (o *EzsignsignerResponseCompound) SetFkiTaxassignmentID(v int32)`

SetFkiTaxassignmentID sets FkiTaxassignmentID field to given value.


### GetFkiSecretquestionID

`func (o *EzsignsignerResponseCompound) GetFkiSecretquestionID() int32`

GetFkiSecretquestionID returns the FkiSecretquestionID field if non-nil, zero value otherwise.

### GetFkiSecretquestionIDOk

`func (o *EzsignsignerResponseCompound) GetFkiSecretquestionIDOk() (*int32, bool)`

GetFkiSecretquestionIDOk returns a tuple with the FkiSecretquestionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiSecretquestionID

`func (o *EzsignsignerResponseCompound) SetFkiSecretquestionID(v int32)`

SetFkiSecretquestionID sets FkiSecretquestionID field to given value.

### HasFkiSecretquestionID

`func (o *EzsignsignerResponseCompound) HasFkiSecretquestionID() bool`

HasFkiSecretquestionID returns a boolean if a field has been set.

### GetFkiUserlogintypeID

`func (o *EzsignsignerResponseCompound) GetFkiUserlogintypeID() int32`

GetFkiUserlogintypeID returns the FkiUserlogintypeID field if non-nil, zero value otherwise.

### GetFkiUserlogintypeIDOk

`func (o *EzsignsignerResponseCompound) GetFkiUserlogintypeIDOk() (*int32, bool)`

GetFkiUserlogintypeIDOk returns a tuple with the FkiUserlogintypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUserlogintypeID

`func (o *EzsignsignerResponseCompound) SetFkiUserlogintypeID(v int32)`

SetFkiUserlogintypeID sets FkiUserlogintypeID field to given value.


### GetSUserlogintypeDescriptionX

`func (o *EzsignsignerResponseCompound) GetSUserlogintypeDescriptionX() string`

GetSUserlogintypeDescriptionX returns the SUserlogintypeDescriptionX field if non-nil, zero value otherwise.

### GetSUserlogintypeDescriptionXOk

`func (o *EzsignsignerResponseCompound) GetSUserlogintypeDescriptionXOk() (*string, bool)`

GetSUserlogintypeDescriptionXOk returns a tuple with the SUserlogintypeDescriptionX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSUserlogintypeDescriptionX

`func (o *EzsignsignerResponseCompound) SetSUserlogintypeDescriptionX(v string)`

SetSUserlogintypeDescriptionX sets SUserlogintypeDescriptionX field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

