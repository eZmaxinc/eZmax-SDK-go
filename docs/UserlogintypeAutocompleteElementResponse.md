# UserlogintypeAutocompleteElementResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiUserlogintypeID** | **int32** | The unique ID of the Userlogintype  Valid values:  |Value|Description|Detail| |-|-|-| |1|**Email Only**|The Ezsignsigner will receive a secure link by email| |2|**Email and phone or SMS**|The Ezsignsigner will receive a secure link by email and will need to authenticate using SMS or Phone call. **Additional fee applies**| |3|**Email and secret question**|The Ezsignsigner will receive a secure link by email and will need to authenticate using a predefined question and answer| |4|**In person only**|The Ezsignsigner will only be able to sign \&quot;In-Person\&quot; and there won&#39;t be any authentication. No email will be sent for invitation to sign. Make sure you evaluate the risk of signature denial and at minimum, we recommend you use a handwritten signature type| |5|**In person with phone or SMS**|The Ezsignsigner will only be able to sign \&quot;In-Person\&quot; and will need to authenticate using SMS or Phone call. No email will be sent for invitation to sign. **Additional fee applies**| |6|**Embedded**|The Ezsignsigner will only be able to sign in the embedded solution. No email will be sent for invitation to sign. **Additional fee applies**|   |7|**Embedded with phone or SMS**|The Ezsignsigner will only be able to sign in the embedded solution and will need to authenticate using SMS or Phone call. No email will be sent for invitation to sign. **Additional fee applies**|   |8|**No validation**|The Ezsignsigner will not receive an email and won&#39;t have to validate his connection using 2 factor. **Additional fee applies**|      |9|**Sms only**|The Ezsignsigner will not receive an email but will will need to authenticate using SMS. **Additional fee applies**|      | 
**SUserlogintypeDescriptionX** | **string** | The description of the Userlogintype in the language of the requester | 
**BUserlogintypeIsactive** | **bool** | Whether the Userlogintype is active or not | 

## Methods

### NewUserlogintypeAutocompleteElementResponse

`func NewUserlogintypeAutocompleteElementResponse(pkiUserlogintypeID int32, sUserlogintypeDescriptionX string, bUserlogintypeIsactive bool, ) *UserlogintypeAutocompleteElementResponse`

NewUserlogintypeAutocompleteElementResponse instantiates a new UserlogintypeAutocompleteElementResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserlogintypeAutocompleteElementResponseWithDefaults

`func NewUserlogintypeAutocompleteElementResponseWithDefaults() *UserlogintypeAutocompleteElementResponse`

NewUserlogintypeAutocompleteElementResponseWithDefaults instantiates a new UserlogintypeAutocompleteElementResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiUserlogintypeID

`func (o *UserlogintypeAutocompleteElementResponse) GetPkiUserlogintypeID() int32`

GetPkiUserlogintypeID returns the PkiUserlogintypeID field if non-nil, zero value otherwise.

### GetPkiUserlogintypeIDOk

`func (o *UserlogintypeAutocompleteElementResponse) GetPkiUserlogintypeIDOk() (*int32, bool)`

GetPkiUserlogintypeIDOk returns a tuple with the PkiUserlogintypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiUserlogintypeID

`func (o *UserlogintypeAutocompleteElementResponse) SetPkiUserlogintypeID(v int32)`

SetPkiUserlogintypeID sets PkiUserlogintypeID field to given value.


### GetSUserlogintypeDescriptionX

`func (o *UserlogintypeAutocompleteElementResponse) GetSUserlogintypeDescriptionX() string`

GetSUserlogintypeDescriptionX returns the SUserlogintypeDescriptionX field if non-nil, zero value otherwise.

### GetSUserlogintypeDescriptionXOk

`func (o *UserlogintypeAutocompleteElementResponse) GetSUserlogintypeDescriptionXOk() (*string, bool)`

GetSUserlogintypeDescriptionXOk returns a tuple with the SUserlogintypeDescriptionX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSUserlogintypeDescriptionX

`func (o *UserlogintypeAutocompleteElementResponse) SetSUserlogintypeDescriptionX(v string)`

SetSUserlogintypeDescriptionX sets SUserlogintypeDescriptionX field to given value.


### GetBUserlogintypeIsactive

`func (o *UserlogintypeAutocompleteElementResponse) GetBUserlogintypeIsactive() bool`

GetBUserlogintypeIsactive returns the BUserlogintypeIsactive field if non-nil, zero value otherwise.

### GetBUserlogintypeIsactiveOk

`func (o *UserlogintypeAutocompleteElementResponse) GetBUserlogintypeIsactiveOk() (*bool, bool)`

GetBUserlogintypeIsactiveOk returns a tuple with the BUserlogintypeIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBUserlogintypeIsactive

`func (o *UserlogintypeAutocompleteElementResponse) SetBUserlogintypeIsactive(v bool)`

SetBUserlogintypeIsactive sets BUserlogintypeIsactive field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


