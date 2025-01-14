# CustomEzsignfoldertypeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsignfoldertypeID** | **int32** | The unique ID of the Ezsignfoldertype. | 
**FkiFontIDAnnotation** | Pointer to **int32** | The unique ID of the Font | [optional] 
**FkiFontIDFormfield** | Pointer to **int32** | The unique ID of the Font | [optional] 
**FkiFontIDSignature** | Pointer to **int32** | The unique ID of the Font | [optional] 
**SEzsignfoldertypeNameX** | Pointer to **string** | The name of the Ezsignfoldertype in the language of the requester | [optional] 
**BEzsignfoldertypeSendproofezsignsigner** | Pointer to **bool** | Whether we send the proof in the email to Ezsignsigner | [optional] 
**BEzsignfoldertypeAllowdownloadattachmentezsignsigner** | Pointer to **bool** | Whether we allow the Ezsigndocument to be downloaded by an Ezsignsigner | [optional] 
**BEzsignfoldertypeAllowdownloadproofezsignsigner** | Pointer to **bool** | Whether we allow the proof to be downloaded by an Ezsignsigner | [optional] 
**BEzsignfoldertypeAutomaticsignature** | Pointer to **bool** | Whether we allow the automatic signature by an User | [optional] 
**BEzsignfoldertypeDelegate** | Pointer to **bool** | Wheter if delegation of signature is allowed to another user or not | [optional] 
**BEzsignfoldertypeDiscussion** | Pointer to **bool** | Wheter if creating a new Discussion is allowed or not | [optional] 
**BEzsignfoldertypeReassignezsignsigner** | Pointer to **bool** | Wheter if Reassignment of signature is allowed by a signatory to another signatory or not | [optional] 
**BEzsignfoldertypeReassignuser** | Pointer to **bool** | Wheter if Reassignment of signature is allowed by a user to a signatory or another user or not | [optional] 
**BEzsignfoldertypeReassigngroup** | Pointer to **bool** | Wheter if Reassignment of signatures of the groups to which the user belongs is authorized by a user to himself | [optional] 
**IEzsignfoldertypeDeadlinedays** | Pointer to **int32** | The number of days to get all Ezsignsignatures | [optional] 
**IEzsignfoldertypeFontsizeannotation** | Pointer to **int32** | Font size for annotations | [optional] 
**IEzsignfoldertypeFontsizeformfield** | Pointer to **int32** | Font size for form fields | [optional] 

## Methods

### NewCustomEzsignfoldertypeResponse

`func NewCustomEzsignfoldertypeResponse(pkiEzsignfoldertypeID int32, ) *CustomEzsignfoldertypeResponse`

NewCustomEzsignfoldertypeResponse instantiates a new CustomEzsignfoldertypeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomEzsignfoldertypeResponseWithDefaults

`func NewCustomEzsignfoldertypeResponseWithDefaults() *CustomEzsignfoldertypeResponse`

NewCustomEzsignfoldertypeResponseWithDefaults instantiates a new CustomEzsignfoldertypeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsignfoldertypeID

`func (o *CustomEzsignfoldertypeResponse) GetPkiEzsignfoldertypeID() int32`

GetPkiEzsignfoldertypeID returns the PkiEzsignfoldertypeID field if non-nil, zero value otherwise.

### GetPkiEzsignfoldertypeIDOk

`func (o *CustomEzsignfoldertypeResponse) GetPkiEzsignfoldertypeIDOk() (*int32, bool)`

GetPkiEzsignfoldertypeIDOk returns a tuple with the PkiEzsignfoldertypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsignfoldertypeID

`func (o *CustomEzsignfoldertypeResponse) SetPkiEzsignfoldertypeID(v int32)`

SetPkiEzsignfoldertypeID sets PkiEzsignfoldertypeID field to given value.


### GetFkiFontIDAnnotation

`func (o *CustomEzsignfoldertypeResponse) GetFkiFontIDAnnotation() int32`

GetFkiFontIDAnnotation returns the FkiFontIDAnnotation field if non-nil, zero value otherwise.

### GetFkiFontIDAnnotationOk

`func (o *CustomEzsignfoldertypeResponse) GetFkiFontIDAnnotationOk() (*int32, bool)`

GetFkiFontIDAnnotationOk returns a tuple with the FkiFontIDAnnotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiFontIDAnnotation

`func (o *CustomEzsignfoldertypeResponse) SetFkiFontIDAnnotation(v int32)`

SetFkiFontIDAnnotation sets FkiFontIDAnnotation field to given value.

### HasFkiFontIDAnnotation

`func (o *CustomEzsignfoldertypeResponse) HasFkiFontIDAnnotation() bool`

HasFkiFontIDAnnotation returns a boolean if a field has been set.

### GetFkiFontIDFormfield

`func (o *CustomEzsignfoldertypeResponse) GetFkiFontIDFormfield() int32`

GetFkiFontIDFormfield returns the FkiFontIDFormfield field if non-nil, zero value otherwise.

### GetFkiFontIDFormfieldOk

`func (o *CustomEzsignfoldertypeResponse) GetFkiFontIDFormfieldOk() (*int32, bool)`

GetFkiFontIDFormfieldOk returns a tuple with the FkiFontIDFormfield field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiFontIDFormfield

`func (o *CustomEzsignfoldertypeResponse) SetFkiFontIDFormfield(v int32)`

SetFkiFontIDFormfield sets FkiFontIDFormfield field to given value.

### HasFkiFontIDFormfield

`func (o *CustomEzsignfoldertypeResponse) HasFkiFontIDFormfield() bool`

HasFkiFontIDFormfield returns a boolean if a field has been set.

### GetFkiFontIDSignature

`func (o *CustomEzsignfoldertypeResponse) GetFkiFontIDSignature() int32`

GetFkiFontIDSignature returns the FkiFontIDSignature field if non-nil, zero value otherwise.

### GetFkiFontIDSignatureOk

`func (o *CustomEzsignfoldertypeResponse) GetFkiFontIDSignatureOk() (*int32, bool)`

GetFkiFontIDSignatureOk returns a tuple with the FkiFontIDSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiFontIDSignature

`func (o *CustomEzsignfoldertypeResponse) SetFkiFontIDSignature(v int32)`

SetFkiFontIDSignature sets FkiFontIDSignature field to given value.

### HasFkiFontIDSignature

`func (o *CustomEzsignfoldertypeResponse) HasFkiFontIDSignature() bool`

HasFkiFontIDSignature returns a boolean if a field has been set.

### GetSEzsignfoldertypeNameX

`func (o *CustomEzsignfoldertypeResponse) GetSEzsignfoldertypeNameX() string`

GetSEzsignfoldertypeNameX returns the SEzsignfoldertypeNameX field if non-nil, zero value otherwise.

### GetSEzsignfoldertypeNameXOk

`func (o *CustomEzsignfoldertypeResponse) GetSEzsignfoldertypeNameXOk() (*string, bool)`

GetSEzsignfoldertypeNameXOk returns a tuple with the SEzsignfoldertypeNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsignfoldertypeNameX

`func (o *CustomEzsignfoldertypeResponse) SetSEzsignfoldertypeNameX(v string)`

SetSEzsignfoldertypeNameX sets SEzsignfoldertypeNameX field to given value.

### HasSEzsignfoldertypeNameX

`func (o *CustomEzsignfoldertypeResponse) HasSEzsignfoldertypeNameX() bool`

HasSEzsignfoldertypeNameX returns a boolean if a field has been set.

### GetBEzsignfoldertypeSendproofezsignsigner

`func (o *CustomEzsignfoldertypeResponse) GetBEzsignfoldertypeSendproofezsignsigner() bool`

GetBEzsignfoldertypeSendproofezsignsigner returns the BEzsignfoldertypeSendproofezsignsigner field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeSendproofezsignsignerOk

`func (o *CustomEzsignfoldertypeResponse) GetBEzsignfoldertypeSendproofezsignsignerOk() (*bool, bool)`

GetBEzsignfoldertypeSendproofezsignsignerOk returns a tuple with the BEzsignfoldertypeSendproofezsignsigner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeSendproofezsignsigner

`func (o *CustomEzsignfoldertypeResponse) SetBEzsignfoldertypeSendproofezsignsigner(v bool)`

SetBEzsignfoldertypeSendproofezsignsigner sets BEzsignfoldertypeSendproofezsignsigner field to given value.

### HasBEzsignfoldertypeSendproofezsignsigner

`func (o *CustomEzsignfoldertypeResponse) HasBEzsignfoldertypeSendproofezsignsigner() bool`

HasBEzsignfoldertypeSendproofezsignsigner returns a boolean if a field has been set.

### GetBEzsignfoldertypeAllowdownloadattachmentezsignsigner

`func (o *CustomEzsignfoldertypeResponse) GetBEzsignfoldertypeAllowdownloadattachmentezsignsigner() bool`

GetBEzsignfoldertypeAllowdownloadattachmentezsignsigner returns the BEzsignfoldertypeAllowdownloadattachmentezsignsigner field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeAllowdownloadattachmentezsignsignerOk

`func (o *CustomEzsignfoldertypeResponse) GetBEzsignfoldertypeAllowdownloadattachmentezsignsignerOk() (*bool, bool)`

GetBEzsignfoldertypeAllowdownloadattachmentezsignsignerOk returns a tuple with the BEzsignfoldertypeAllowdownloadattachmentezsignsigner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeAllowdownloadattachmentezsignsigner

`func (o *CustomEzsignfoldertypeResponse) SetBEzsignfoldertypeAllowdownloadattachmentezsignsigner(v bool)`

SetBEzsignfoldertypeAllowdownloadattachmentezsignsigner sets BEzsignfoldertypeAllowdownloadattachmentezsignsigner field to given value.

### HasBEzsignfoldertypeAllowdownloadattachmentezsignsigner

`func (o *CustomEzsignfoldertypeResponse) HasBEzsignfoldertypeAllowdownloadattachmentezsignsigner() bool`

HasBEzsignfoldertypeAllowdownloadattachmentezsignsigner returns a boolean if a field has been set.

### GetBEzsignfoldertypeAllowdownloadproofezsignsigner

`func (o *CustomEzsignfoldertypeResponse) GetBEzsignfoldertypeAllowdownloadproofezsignsigner() bool`

GetBEzsignfoldertypeAllowdownloadproofezsignsigner returns the BEzsignfoldertypeAllowdownloadproofezsignsigner field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeAllowdownloadproofezsignsignerOk

`func (o *CustomEzsignfoldertypeResponse) GetBEzsignfoldertypeAllowdownloadproofezsignsignerOk() (*bool, bool)`

GetBEzsignfoldertypeAllowdownloadproofezsignsignerOk returns a tuple with the BEzsignfoldertypeAllowdownloadproofezsignsigner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeAllowdownloadproofezsignsigner

`func (o *CustomEzsignfoldertypeResponse) SetBEzsignfoldertypeAllowdownloadproofezsignsigner(v bool)`

SetBEzsignfoldertypeAllowdownloadproofezsignsigner sets BEzsignfoldertypeAllowdownloadproofezsignsigner field to given value.

### HasBEzsignfoldertypeAllowdownloadproofezsignsigner

`func (o *CustomEzsignfoldertypeResponse) HasBEzsignfoldertypeAllowdownloadproofezsignsigner() bool`

HasBEzsignfoldertypeAllowdownloadproofezsignsigner returns a boolean if a field has been set.

### GetBEzsignfoldertypeAutomaticsignature

`func (o *CustomEzsignfoldertypeResponse) GetBEzsignfoldertypeAutomaticsignature() bool`

GetBEzsignfoldertypeAutomaticsignature returns the BEzsignfoldertypeAutomaticsignature field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeAutomaticsignatureOk

`func (o *CustomEzsignfoldertypeResponse) GetBEzsignfoldertypeAutomaticsignatureOk() (*bool, bool)`

GetBEzsignfoldertypeAutomaticsignatureOk returns a tuple with the BEzsignfoldertypeAutomaticsignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeAutomaticsignature

`func (o *CustomEzsignfoldertypeResponse) SetBEzsignfoldertypeAutomaticsignature(v bool)`

SetBEzsignfoldertypeAutomaticsignature sets BEzsignfoldertypeAutomaticsignature field to given value.

### HasBEzsignfoldertypeAutomaticsignature

`func (o *CustomEzsignfoldertypeResponse) HasBEzsignfoldertypeAutomaticsignature() bool`

HasBEzsignfoldertypeAutomaticsignature returns a boolean if a field has been set.

### GetBEzsignfoldertypeDelegate

`func (o *CustomEzsignfoldertypeResponse) GetBEzsignfoldertypeDelegate() bool`

GetBEzsignfoldertypeDelegate returns the BEzsignfoldertypeDelegate field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeDelegateOk

`func (o *CustomEzsignfoldertypeResponse) GetBEzsignfoldertypeDelegateOk() (*bool, bool)`

GetBEzsignfoldertypeDelegateOk returns a tuple with the BEzsignfoldertypeDelegate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeDelegate

`func (o *CustomEzsignfoldertypeResponse) SetBEzsignfoldertypeDelegate(v bool)`

SetBEzsignfoldertypeDelegate sets BEzsignfoldertypeDelegate field to given value.

### HasBEzsignfoldertypeDelegate

`func (o *CustomEzsignfoldertypeResponse) HasBEzsignfoldertypeDelegate() bool`

HasBEzsignfoldertypeDelegate returns a boolean if a field has been set.

### GetBEzsignfoldertypeDiscussion

`func (o *CustomEzsignfoldertypeResponse) GetBEzsignfoldertypeDiscussion() bool`

GetBEzsignfoldertypeDiscussion returns the BEzsignfoldertypeDiscussion field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeDiscussionOk

`func (o *CustomEzsignfoldertypeResponse) GetBEzsignfoldertypeDiscussionOk() (*bool, bool)`

GetBEzsignfoldertypeDiscussionOk returns a tuple with the BEzsignfoldertypeDiscussion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeDiscussion

`func (o *CustomEzsignfoldertypeResponse) SetBEzsignfoldertypeDiscussion(v bool)`

SetBEzsignfoldertypeDiscussion sets BEzsignfoldertypeDiscussion field to given value.

### HasBEzsignfoldertypeDiscussion

`func (o *CustomEzsignfoldertypeResponse) HasBEzsignfoldertypeDiscussion() bool`

HasBEzsignfoldertypeDiscussion returns a boolean if a field has been set.

### GetBEzsignfoldertypeReassignezsignsigner

`func (o *CustomEzsignfoldertypeResponse) GetBEzsignfoldertypeReassignezsignsigner() bool`

GetBEzsignfoldertypeReassignezsignsigner returns the BEzsignfoldertypeReassignezsignsigner field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeReassignezsignsignerOk

`func (o *CustomEzsignfoldertypeResponse) GetBEzsignfoldertypeReassignezsignsignerOk() (*bool, bool)`

GetBEzsignfoldertypeReassignezsignsignerOk returns a tuple with the BEzsignfoldertypeReassignezsignsigner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeReassignezsignsigner

`func (o *CustomEzsignfoldertypeResponse) SetBEzsignfoldertypeReassignezsignsigner(v bool)`

SetBEzsignfoldertypeReassignezsignsigner sets BEzsignfoldertypeReassignezsignsigner field to given value.

### HasBEzsignfoldertypeReassignezsignsigner

`func (o *CustomEzsignfoldertypeResponse) HasBEzsignfoldertypeReassignezsignsigner() bool`

HasBEzsignfoldertypeReassignezsignsigner returns a boolean if a field has been set.

### GetBEzsignfoldertypeReassignuser

`func (o *CustomEzsignfoldertypeResponse) GetBEzsignfoldertypeReassignuser() bool`

GetBEzsignfoldertypeReassignuser returns the BEzsignfoldertypeReassignuser field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeReassignuserOk

`func (o *CustomEzsignfoldertypeResponse) GetBEzsignfoldertypeReassignuserOk() (*bool, bool)`

GetBEzsignfoldertypeReassignuserOk returns a tuple with the BEzsignfoldertypeReassignuser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeReassignuser

`func (o *CustomEzsignfoldertypeResponse) SetBEzsignfoldertypeReassignuser(v bool)`

SetBEzsignfoldertypeReassignuser sets BEzsignfoldertypeReassignuser field to given value.

### HasBEzsignfoldertypeReassignuser

`func (o *CustomEzsignfoldertypeResponse) HasBEzsignfoldertypeReassignuser() bool`

HasBEzsignfoldertypeReassignuser returns a boolean if a field has been set.

### GetBEzsignfoldertypeReassigngroup

`func (o *CustomEzsignfoldertypeResponse) GetBEzsignfoldertypeReassigngroup() bool`

GetBEzsignfoldertypeReassigngroup returns the BEzsignfoldertypeReassigngroup field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeReassigngroupOk

`func (o *CustomEzsignfoldertypeResponse) GetBEzsignfoldertypeReassigngroupOk() (*bool, bool)`

GetBEzsignfoldertypeReassigngroupOk returns a tuple with the BEzsignfoldertypeReassigngroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeReassigngroup

`func (o *CustomEzsignfoldertypeResponse) SetBEzsignfoldertypeReassigngroup(v bool)`

SetBEzsignfoldertypeReassigngroup sets BEzsignfoldertypeReassigngroup field to given value.

### HasBEzsignfoldertypeReassigngroup

`func (o *CustomEzsignfoldertypeResponse) HasBEzsignfoldertypeReassigngroup() bool`

HasBEzsignfoldertypeReassigngroup returns a boolean if a field has been set.

### GetIEzsignfoldertypeDeadlinedays

`func (o *CustomEzsignfoldertypeResponse) GetIEzsignfoldertypeDeadlinedays() int32`

GetIEzsignfoldertypeDeadlinedays returns the IEzsignfoldertypeDeadlinedays field if non-nil, zero value otherwise.

### GetIEzsignfoldertypeDeadlinedaysOk

`func (o *CustomEzsignfoldertypeResponse) GetIEzsignfoldertypeDeadlinedaysOk() (*int32, bool)`

GetIEzsignfoldertypeDeadlinedaysOk returns a tuple with the IEzsignfoldertypeDeadlinedays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignfoldertypeDeadlinedays

`func (o *CustomEzsignfoldertypeResponse) SetIEzsignfoldertypeDeadlinedays(v int32)`

SetIEzsignfoldertypeDeadlinedays sets IEzsignfoldertypeDeadlinedays field to given value.

### HasIEzsignfoldertypeDeadlinedays

`func (o *CustomEzsignfoldertypeResponse) HasIEzsignfoldertypeDeadlinedays() bool`

HasIEzsignfoldertypeDeadlinedays returns a boolean if a field has been set.

### GetIEzsignfoldertypeFontsizeannotation

`func (o *CustomEzsignfoldertypeResponse) GetIEzsignfoldertypeFontsizeannotation() int32`

GetIEzsignfoldertypeFontsizeannotation returns the IEzsignfoldertypeFontsizeannotation field if non-nil, zero value otherwise.

### GetIEzsignfoldertypeFontsizeannotationOk

`func (o *CustomEzsignfoldertypeResponse) GetIEzsignfoldertypeFontsizeannotationOk() (*int32, bool)`

GetIEzsignfoldertypeFontsizeannotationOk returns a tuple with the IEzsignfoldertypeFontsizeannotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignfoldertypeFontsizeannotation

`func (o *CustomEzsignfoldertypeResponse) SetIEzsignfoldertypeFontsizeannotation(v int32)`

SetIEzsignfoldertypeFontsizeannotation sets IEzsignfoldertypeFontsizeannotation field to given value.

### HasIEzsignfoldertypeFontsizeannotation

`func (o *CustomEzsignfoldertypeResponse) HasIEzsignfoldertypeFontsizeannotation() bool`

HasIEzsignfoldertypeFontsizeannotation returns a boolean if a field has been set.

### GetIEzsignfoldertypeFontsizeformfield

`func (o *CustomEzsignfoldertypeResponse) GetIEzsignfoldertypeFontsizeformfield() int32`

GetIEzsignfoldertypeFontsizeformfield returns the IEzsignfoldertypeFontsizeformfield field if non-nil, zero value otherwise.

### GetIEzsignfoldertypeFontsizeformfieldOk

`func (o *CustomEzsignfoldertypeResponse) GetIEzsignfoldertypeFontsizeformfieldOk() (*int32, bool)`

GetIEzsignfoldertypeFontsizeformfieldOk returns a tuple with the IEzsignfoldertypeFontsizeformfield field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsignfoldertypeFontsizeformfield

`func (o *CustomEzsignfoldertypeResponse) SetIEzsignfoldertypeFontsizeformfield(v int32)`

SetIEzsignfoldertypeFontsizeformfield sets IEzsignfoldertypeFontsizeformfield field to given value.

### HasIEzsignfoldertypeFontsizeformfield

`func (o *CustomEzsignfoldertypeResponse) HasIEzsignfoldertypeFontsizeformfield() bool`

HasIEzsignfoldertypeFontsizeformfield returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


