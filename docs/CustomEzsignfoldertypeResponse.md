# CustomEzsignfoldertypeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsignfoldertypeID** | **int32** | The unique ID of the Ezsignfoldertype. | 
**SEzsignfoldertypeNameX** | Pointer to **string** | The name of the Ezsignfoldertype in the language of the requester | [optional] 
**BEzsignfoldertypeIncludeproofsigner** | Pointer to **bool** | Whether we include the proof with the signed Ezsigndocument for Ezsignsigners | [optional] 
**BEzsignfoldertypeIncludeproofuser** | Pointer to **bool** | Whether we include the proof with the signed Ezsigndocument for users | [optional] 
**BEzsignfoldertypeDelegate** | Pointer to **bool** | Wheter if delegation of signature is allowed to another user or not | [optional] 
**BEzsignfoldertypeReassign** | Pointer to **bool** | Wheter if Reassignment of signature is allowed to another signatory or not | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


