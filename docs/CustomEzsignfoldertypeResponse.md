# CustomEzsignfoldertypeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsignfoldertypeID** | **int32** | The unique ID of the Ezsignfoldertype. | 
**SEzsignfoldertypeNameX** | Pointer to **string** | The name of the Ezsignfoldertype in the language of the requester | [optional] 
**BEzsignfoldertypeSendproofezsignsigner** | Pointer to **bool** | Whether we send the proof in the email to Ezsignsigner | [optional] 
**BEzsignfoldertypeIncludeproofsigner** | Pointer to **bool** | THIS FIELD WILL BE DELETED. Whether we include the proof with the signed Ezsigndocument for Ezsignsigners | [optional] 
**BEzsignfoldertypeIncludeproofuser** | Pointer to **bool** | Whether we include the proof with the signed Ezsigndocument for users | [optional] 
**BEzsignfoldertypeAllowdownloadattachmentezsignsigner** | Pointer to **bool** | Whether we allow the Ezsigndocument to be downloaded by an Ezsignsigner | [optional] 
**BEzsignfoldertypeAllowdownloadproofezsignsigner** | Pointer to **bool** | Whether we allow the proof to be downloaded by an Ezsignsigner | [optional] 
**BEzsignfoldertypeDelegate** | Pointer to **bool** | Wheter if delegation of signature is allowed to another user or not | [optional] 
**BEzsignfoldertypeReassign** | Pointer to **bool** | Wheter if Reassignment of signature is allowed to another signatory or not | [optional] 
**BEzsignfoldertypeReassignezsignsigner** | Pointer to **bool** | Wheter if Reassignment of signature is allowed by a signatory to another signatory or not | [optional] 
**BEzsignfoldertypeReassignuser** | Pointer to **bool** | Wheter if Reassignment of signature is allowed by a user to a signatory or another user or not | [optional] 

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

### GetBEzsignfoldertypeIncludeproofsigner

`func (o *CustomEzsignfoldertypeResponse) GetBEzsignfoldertypeIncludeproofsigner() bool`

GetBEzsignfoldertypeIncludeproofsigner returns the BEzsignfoldertypeIncludeproofsigner field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeIncludeproofsignerOk

`func (o *CustomEzsignfoldertypeResponse) GetBEzsignfoldertypeIncludeproofsignerOk() (*bool, bool)`

GetBEzsignfoldertypeIncludeproofsignerOk returns a tuple with the BEzsignfoldertypeIncludeproofsigner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeIncludeproofsigner

`func (o *CustomEzsignfoldertypeResponse) SetBEzsignfoldertypeIncludeproofsigner(v bool)`

SetBEzsignfoldertypeIncludeproofsigner sets BEzsignfoldertypeIncludeproofsigner field to given value.

### HasBEzsignfoldertypeIncludeproofsigner

`func (o *CustomEzsignfoldertypeResponse) HasBEzsignfoldertypeIncludeproofsigner() bool`

HasBEzsignfoldertypeIncludeproofsigner returns a boolean if a field has been set.

### GetBEzsignfoldertypeIncludeproofuser

`func (o *CustomEzsignfoldertypeResponse) GetBEzsignfoldertypeIncludeproofuser() bool`

GetBEzsignfoldertypeIncludeproofuser returns the BEzsignfoldertypeIncludeproofuser field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeIncludeproofuserOk

`func (o *CustomEzsignfoldertypeResponse) GetBEzsignfoldertypeIncludeproofuserOk() (*bool, bool)`

GetBEzsignfoldertypeIncludeproofuserOk returns a tuple with the BEzsignfoldertypeIncludeproofuser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeIncludeproofuser

`func (o *CustomEzsignfoldertypeResponse) SetBEzsignfoldertypeIncludeproofuser(v bool)`

SetBEzsignfoldertypeIncludeproofuser sets BEzsignfoldertypeIncludeproofuser field to given value.

### HasBEzsignfoldertypeIncludeproofuser

`func (o *CustomEzsignfoldertypeResponse) HasBEzsignfoldertypeIncludeproofuser() bool`

HasBEzsignfoldertypeIncludeproofuser returns a boolean if a field has been set.

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

### GetBEzsignfoldertypeReassign

`func (o *CustomEzsignfoldertypeResponse) GetBEzsignfoldertypeReassign() bool`

GetBEzsignfoldertypeReassign returns the BEzsignfoldertypeReassign field if non-nil, zero value otherwise.

### GetBEzsignfoldertypeReassignOk

`func (o *CustomEzsignfoldertypeResponse) GetBEzsignfoldertypeReassignOk() (*bool, bool)`

GetBEzsignfoldertypeReassignOk returns a tuple with the BEzsignfoldertypeReassign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignfoldertypeReassign

`func (o *CustomEzsignfoldertypeResponse) SetBEzsignfoldertypeReassign(v bool)`

SetBEzsignfoldertypeReassign sets BEzsignfoldertypeReassign field to given value.

### HasBEzsignfoldertypeReassign

`func (o *CustomEzsignfoldertypeResponse) HasBEzsignfoldertypeReassign() bool`

HasBEzsignfoldertypeReassign returns a boolean if a field has been set.

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


