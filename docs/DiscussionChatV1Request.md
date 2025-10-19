# DiscussionChatV1Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FkiDiscussionID** | Pointer to **int32** | The unique ID of the Discussion | [optional] 
**EDiscussionRobot** | [**FieldEDiscussionRobot**](FieldEDiscussionRobot.md) |  | 
**TDiscussionMessage** | **string** | The Message of the Discussion | 

## Methods

### NewDiscussionChatV1Request

`func NewDiscussionChatV1Request(eDiscussionRobot FieldEDiscussionRobot, tDiscussionMessage string, ) *DiscussionChatV1Request`

NewDiscussionChatV1Request instantiates a new DiscussionChatV1Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscussionChatV1RequestWithDefaults

`func NewDiscussionChatV1RequestWithDefaults() *DiscussionChatV1Request`

NewDiscussionChatV1RequestWithDefaults instantiates a new DiscussionChatV1Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFkiDiscussionID

`func (o *DiscussionChatV1Request) GetFkiDiscussionID() int32`

GetFkiDiscussionID returns the FkiDiscussionID field if non-nil, zero value otherwise.

### GetFkiDiscussionIDOk

`func (o *DiscussionChatV1Request) GetFkiDiscussionIDOk() (*int32, bool)`

GetFkiDiscussionIDOk returns a tuple with the FkiDiscussionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiDiscussionID

`func (o *DiscussionChatV1Request) SetFkiDiscussionID(v int32)`

SetFkiDiscussionID sets FkiDiscussionID field to given value.

### HasFkiDiscussionID

`func (o *DiscussionChatV1Request) HasFkiDiscussionID() bool`

HasFkiDiscussionID returns a boolean if a field has been set.

### GetEDiscussionRobot

`func (o *DiscussionChatV1Request) GetEDiscussionRobot() FieldEDiscussionRobot`

GetEDiscussionRobot returns the EDiscussionRobot field if non-nil, zero value otherwise.

### GetEDiscussionRobotOk

`func (o *DiscussionChatV1Request) GetEDiscussionRobotOk() (*FieldEDiscussionRobot, bool)`

GetEDiscussionRobotOk returns a tuple with the EDiscussionRobot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEDiscussionRobot

`func (o *DiscussionChatV1Request) SetEDiscussionRobot(v FieldEDiscussionRobot)`

SetEDiscussionRobot sets EDiscussionRobot field to given value.


### GetTDiscussionMessage

`func (o *DiscussionChatV1Request) GetTDiscussionMessage() string`

GetTDiscussionMessage returns the TDiscussionMessage field if non-nil, zero value otherwise.

### GetTDiscussionMessageOk

`func (o *DiscussionChatV1Request) GetTDiscussionMessageOk() (*string, bool)`

GetTDiscussionMessageOk returns a tuple with the TDiscussionMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTDiscussionMessage

`func (o *DiscussionChatV1Request) SetTDiscussionMessage(v string)`

SetTDiscussionMessage sets TDiscussionMessage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


