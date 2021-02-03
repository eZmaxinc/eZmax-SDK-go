# EzsignsignatureRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FkiEzsignfoldersignerassociationID** | **int32** | A reference to a valid Ezsignfoldersignerassociation.  That value is returned after a successful Ezsignfoldersignerassociation Creation.  | 
**IEzsignpagePagenumber** | **int32** | The page number in the document where to apply the signature | 
**IEzsignsignatureX** | **int32** | The X coordinate (Horizontal) where to put the signature block on the page.  Coordinate is calculated at 100dpi (dot per inch). So for example, if you want to put the signature block 2 inches from the left border of the page, you would use \&quot;200\&quot; for the X coordinate. | 
**IEzsignsignatureY** | **int32** | The Y coordinate (Vertical) where to put the signature block on the page.  Coordinate is calculated at 100dpi (dot per inch). So for example, if you want to put the signature block 3 inches from the top border of the page, you would use \&quot;300\&quot; for the Y coordinate. | 
**IEzsignsignatureStep** | **int32** | The step when the Ezsignsigner will be invited to sign.  For example, if you say iEzsignsignatureStep&#x3D;2, that block of signature will be available for signature only after ALL the signatures in step 1 are completed. | 
**EEzsignsignatureType** | **string** | The type of signature required.  1. **Acknowledgement** is for an acknowledgment of receipt. 2. **Handwritten** is for a handwritten kind of signature where users needs to \&quot;draw\&quot; their signature on screen. 3. **Initials** is a simple \&quot;click to add initials\&quot; block. 4. **Name** is a simple \&quot;Click to sign\&quot; block. This is the most common block of signature. | 
**FkiEzsigndocumentID** | **int32** | A reference to a valid Ezsigndocument.  That value is returned after a successful Ezsigndocumentation Creation. | 

## Methods

### NewEzsignsignatureRequest

`func NewEzsignsignatureRequest(fkiEzsignfoldersignerassociationID int32, iEzsignpagePagenumber int32, iEzsignsignatureX int32, iEzsignsignatureY int32, iEzsignsignatureStep int32, eEzsignsignatureType string, fkiEzsigndocumentID int32, ) *EzsignsignatureRequest`

NewEzsignsignatureRequest instantiates a new EzsignsignatureRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsignsignatureRequestWithDefaults

`func NewEzsignsignatureRequestWithDefaults() *EzsignsignatureRequest`

NewEzsignsignatureRequestWithDefaults instantiates a new EzsignsignatureRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFkiEzsignfoldersignerassociationID

`func (o *EzsignsignatureRequest) GetFkiEzsignfoldersignerassociationID() int32`

GetFkiEzsignfoldersignerassociationID returns the FkiEzsignfoldersignerassociationID field if non-nil, zero value otherwise.

### GetFkiEzsignfoldersignerassociationIDOk

`func (o *EzsignsignatureRequest) GetFkiEzsignfoldersignerassociationIDOk() (*int32, bool)`

GetFkiEzsignfoldersignerassociationIDOk returns a tuple with the FkiEzsignfoldersignerassociationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignfoldersignerassociationID

`func (o *EzsignsignatureRequest) SetFkiEzsignfoldersignerassociationID(v int32)`

SetFkiEzsignfoldersignerassociationID sets FkiEzsignfoldersignerassociationID field to given value.


### GetIEzsignpagePagenumber

`func (o *EzsignsignatureRequest) GetIEzsignpagePagenumber() int32`

GetIEzsignpagePagenumber returns the IEzsignpagePagenumber field if non-nil, zero value otherwise.

### GetIEzsignpagePagenumberOk

`func (o *EzsignsignatureRequest) GetIEzsignpagePagenumberOk() (*int32, bool)`

GetIEzsignpagePagenumberOk returns a tuple with the IEzsignpagePagenumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignpagePagenumber

`func (o *EzsignsignatureRequest) SetIEzsignpagePagenumber(v int32)`

SetIEzsignpagePagenumber sets IEzsignpagePagenumber field to given value.


### GetIEzsignsignatureX

`func (o *EzsignsignatureRequest) GetIEzsignsignatureX() int32`

GetIEzsignsignatureX returns the IEzsignsignatureX field if non-nil, zero value otherwise.

### GetIEzsignsignatureXOk

`func (o *EzsignsignatureRequest) GetIEzsignsignatureXOk() (*int32, bool)`

GetIEzsignsignatureXOk returns a tuple with the IEzsignsignatureX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignsignatureX

`func (o *EzsignsignatureRequest) SetIEzsignsignatureX(v int32)`

SetIEzsignsignatureX sets IEzsignsignatureX field to given value.


### GetIEzsignsignatureY

`func (o *EzsignsignatureRequest) GetIEzsignsignatureY() int32`

GetIEzsignsignatureY returns the IEzsignsignatureY field if non-nil, zero value otherwise.

### GetIEzsignsignatureYOk

`func (o *EzsignsignatureRequest) GetIEzsignsignatureYOk() (*int32, bool)`

GetIEzsignsignatureYOk returns a tuple with the IEzsignsignatureY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignsignatureY

`func (o *EzsignsignatureRequest) SetIEzsignsignatureY(v int32)`

SetIEzsignsignatureY sets IEzsignsignatureY field to given value.


### GetIEzsignsignatureStep

`func (o *EzsignsignatureRequest) GetIEzsignsignatureStep() int32`

GetIEzsignsignatureStep returns the IEzsignsignatureStep field if non-nil, zero value otherwise.

### GetIEzsignsignatureStepOk

`func (o *EzsignsignatureRequest) GetIEzsignsignatureStepOk() (*int32, bool)`

GetIEzsignsignatureStepOk returns a tuple with the IEzsignsignatureStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignsignatureStep

`func (o *EzsignsignatureRequest) SetIEzsignsignatureStep(v int32)`

SetIEzsignsignatureStep sets IEzsignsignatureStep field to given value.


### GetEEzsignsignatureType

`func (o *EzsignsignatureRequest) GetEEzsignsignatureType() string`

GetEEzsignsignatureType returns the EEzsignsignatureType field if non-nil, zero value otherwise.

### GetEEzsignsignatureTypeOk

`func (o *EzsignsignatureRequest) GetEEzsignsignatureTypeOk() (*string, bool)`

GetEEzsignsignatureTypeOk returns a tuple with the EEzsignsignatureType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsignsignatureType

`func (o *EzsignsignatureRequest) SetEEzsignsignatureType(v string)`

SetEEzsignsignatureType sets EEzsignsignatureType field to given value.


### GetFkiEzsigndocumentID

`func (o *EzsignsignatureRequest) GetFkiEzsigndocumentID() int32`

GetFkiEzsigndocumentID returns the FkiEzsigndocumentID field if non-nil, zero value otherwise.

### GetFkiEzsigndocumentIDOk

`func (o *EzsignsignatureRequest) GetFkiEzsigndocumentIDOk() (*int32, bool)`

GetFkiEzsigndocumentIDOk returns a tuple with the FkiEzsigndocumentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsigndocumentID

`func (o *EzsignsignatureRequest) SetFkiEzsigndocumentID(v int32)`

SetFkiEzsigndocumentID sets FkiEzsigndocumentID field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


