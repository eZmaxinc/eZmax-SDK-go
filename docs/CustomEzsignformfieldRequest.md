# CustomEzsignformfieldRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsignformfieldID** | Pointer to **int32** | The unique ID of the Ezsignformfield | [optional] 
**SEzsignformfieldLabel** | Pointer to **string** | The Label for the Ezsignformfield | [optional] 
**BEzsignformfieldSelected** | Pointer to **bool** | Whether the Ezsignformfield is selected or not by default.  This can only be set if eEzsignformfieldgroupType is **Checkbox** or **Radio** | [optional] 
**SEzsignformfieldEnteredvalue** | Pointer to **string** | This is the value enterred for the Ezsignformfield  This can only be set if eEzsignformfieldgroupType is **Dropdown**, **Text** or **Textarea**  You can use the codes below and they will be replaced at signature time.    | Code | Description | Example | | ------------------------- | ------------ | ------------ | | {sUserFirstname} | The first name of the contact | John | | {sUserLastname} | The last name of the contact | Doe | | {sUserJobtitle} | The job title | Sales Representative | | {sEmailAddress} | The email address | email@example.com | | {sPhoneE164} | A phone number in E.164 Format | +15149901516 | | {sPhoneE164Cell} | A phone number in E.164 Format | +15149901516 | | [optional] 

## Methods

### NewCustomEzsignformfieldRequest

`func NewCustomEzsignformfieldRequest() *CustomEzsignformfieldRequest`

NewCustomEzsignformfieldRequest instantiates a new CustomEzsignformfieldRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomEzsignformfieldRequestWithDefaults

`func NewCustomEzsignformfieldRequestWithDefaults() *CustomEzsignformfieldRequest`

NewCustomEzsignformfieldRequestWithDefaults instantiates a new CustomEzsignformfieldRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsignformfieldID

`func (o *CustomEzsignformfieldRequest) GetPkiEzsignformfieldID() int32`

GetPkiEzsignformfieldID returns the PkiEzsignformfieldID field if non-nil, zero value otherwise.

### GetPkiEzsignformfieldIDOk

`func (o *CustomEzsignformfieldRequest) GetPkiEzsignformfieldIDOk() (*int32, bool)`

GetPkiEzsignformfieldIDOk returns a tuple with the PkiEzsignformfieldID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsignformfieldID

`func (o *CustomEzsignformfieldRequest) SetPkiEzsignformfieldID(v int32)`

SetPkiEzsignformfieldID sets PkiEzsignformfieldID field to given value.

### HasPkiEzsignformfieldID

`func (o *CustomEzsignformfieldRequest) HasPkiEzsignformfieldID() bool`

HasPkiEzsignformfieldID returns a boolean if a field has been set.

### GetSEzsignformfieldLabel

`func (o *CustomEzsignformfieldRequest) GetSEzsignformfieldLabel() string`

GetSEzsignformfieldLabel returns the SEzsignformfieldLabel field if non-nil, zero value otherwise.

### GetSEzsignformfieldLabelOk

`func (o *CustomEzsignformfieldRequest) GetSEzsignformfieldLabelOk() (*string, bool)`

GetSEzsignformfieldLabelOk returns a tuple with the SEzsignformfieldLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsignformfieldLabel

`func (o *CustomEzsignformfieldRequest) SetSEzsignformfieldLabel(v string)`

SetSEzsignformfieldLabel sets SEzsignformfieldLabel field to given value.

### HasSEzsignformfieldLabel

`func (o *CustomEzsignformfieldRequest) HasSEzsignformfieldLabel() bool`

HasSEzsignformfieldLabel returns a boolean if a field has been set.

### GetBEzsignformfieldSelected

`func (o *CustomEzsignformfieldRequest) GetBEzsignformfieldSelected() bool`

GetBEzsignformfieldSelected returns the BEzsignformfieldSelected field if non-nil, zero value otherwise.

### GetBEzsignformfieldSelectedOk

`func (o *CustomEzsignformfieldRequest) GetBEzsignformfieldSelectedOk() (*bool, bool)`

GetBEzsignformfieldSelectedOk returns a tuple with the BEzsignformfieldSelected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzsignformfieldSelected

`func (o *CustomEzsignformfieldRequest) SetBEzsignformfieldSelected(v bool)`

SetBEzsignformfieldSelected sets BEzsignformfieldSelected field to given value.

### HasBEzsignformfieldSelected

`func (o *CustomEzsignformfieldRequest) HasBEzsignformfieldSelected() bool`

HasBEzsignformfieldSelected returns a boolean if a field has been set.

### GetSEzsignformfieldEnteredvalue

`func (o *CustomEzsignformfieldRequest) GetSEzsignformfieldEnteredvalue() string`

GetSEzsignformfieldEnteredvalue returns the SEzsignformfieldEnteredvalue field if non-nil, zero value otherwise.

### GetSEzsignformfieldEnteredvalueOk

`func (o *CustomEzsignformfieldRequest) GetSEzsignformfieldEnteredvalueOk() (*string, bool)`

GetSEzsignformfieldEnteredvalueOk returns a tuple with the SEzsignformfieldEnteredvalue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsignformfieldEnteredvalue

`func (o *CustomEzsignformfieldRequest) SetSEzsignformfieldEnteredvalue(v string)`

SetSEzsignformfieldEnteredvalue sets SEzsignformfieldEnteredvalue field to given value.

### HasSEzsignformfieldEnteredvalue

`func (o *CustomEzsignformfieldRequest) HasSEzsignformfieldEnteredvalue() bool`

HasSEzsignformfieldEnteredvalue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


