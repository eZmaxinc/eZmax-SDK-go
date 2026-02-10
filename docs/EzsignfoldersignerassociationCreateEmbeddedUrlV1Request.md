# EzsignfoldersignerassociationCreateEmbeddedUrlV1Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SReturnUrl** | Pointer to **string** | The return Url to redirect after the signing is completed  **Warning** Due to the potential for Ezsignsigners to cancel redirection, close their browser post-signing, or spoof the landing URL, it&#39;s advisable not to solely depend on the sReturnUrl for accurate status within your integration.  Once the Ezsignsigner finishes, they are directed back to your application. Your application can retain transaction state details by either storing data in a cookie or incorporating query parameters in the sReturnUrl. For example: https://www.example.com/sReturnUrl?sSessionID&#x3D;ABC123  The actual url that will be called will have an extra url parameter appended to give details about the process. The possible values are listed in the table below. For example: https://www.example.com/sReturnUrl?sSessionID&#x3D;ABC123&amp;eEzsignEvent&#x3D;CompletedEzsignfolder   |**Query parameters appended**| |---| |eEzsignEvent|   |**eEzsignEvent**|**Description**| |---|---| |SessionTimeout|The session timed out| |SessionLogout|The Ezsignsigner signed out| |DeclinedTermOfUse|The Ezsignsigner refused the terms| |DeclinedSign|The Ezsignsigner refused to sign| |Reassigned|The Ezsignsigner reassigned his signatures to someone else| |CompletedStep|The Ezsignsigner completed his step. There is other signatures to complete the Ezsigndocument| |CompletedEzsignfolder|The Ezsignfolder is completed. Everyone signed their signatures| | [optional] 
**SIframedomain** | Pointer to **string** | Domain protection for the iFrame | [optional] 
**BIsIframe** | Pointer to **bool** | Whether the url would be in an iFrame or not | [optional] 

## Methods

### NewEzsignfoldersignerassociationCreateEmbeddedUrlV1Request

`func NewEzsignfoldersignerassociationCreateEmbeddedUrlV1Request() *EzsignfoldersignerassociationCreateEmbeddedUrlV1Request`

NewEzsignfoldersignerassociationCreateEmbeddedUrlV1Request instantiates a new EzsignfoldersignerassociationCreateEmbeddedUrlV1Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsignfoldersignerassociationCreateEmbeddedUrlV1RequestWithDefaults

`func NewEzsignfoldersignerassociationCreateEmbeddedUrlV1RequestWithDefaults() *EzsignfoldersignerassociationCreateEmbeddedUrlV1Request`

NewEzsignfoldersignerassociationCreateEmbeddedUrlV1RequestWithDefaults instantiates a new EzsignfoldersignerassociationCreateEmbeddedUrlV1Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSReturnUrl

`func (o *EzsignfoldersignerassociationCreateEmbeddedUrlV1Request) GetSReturnUrl() string`

GetSReturnUrl returns the SReturnUrl field if non-nil, zero value otherwise.

### GetSReturnUrlOk

`func (o *EzsignfoldersignerassociationCreateEmbeddedUrlV1Request) GetSReturnUrlOk() (*string, bool)`

GetSReturnUrlOk returns a tuple with the SReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSReturnUrl

`func (o *EzsignfoldersignerassociationCreateEmbeddedUrlV1Request) SetSReturnUrl(v string)`

SetSReturnUrl sets SReturnUrl field to given value.

### HasSReturnUrl

`func (o *EzsignfoldersignerassociationCreateEmbeddedUrlV1Request) HasSReturnUrl() bool`

HasSReturnUrl returns a boolean if a field has been set.

### GetSIframedomain

`func (o *EzsignfoldersignerassociationCreateEmbeddedUrlV1Request) GetSIframedomain() string`

GetSIframedomain returns the SIframedomain field if non-nil, zero value otherwise.

### GetSIframedomainOk

`func (o *EzsignfoldersignerassociationCreateEmbeddedUrlV1Request) GetSIframedomainOk() (*string, bool)`

GetSIframedomainOk returns a tuple with the SIframedomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSIframedomain

`func (o *EzsignfoldersignerassociationCreateEmbeddedUrlV1Request) SetSIframedomain(v string)`

SetSIframedomain sets SIframedomain field to given value.

### HasSIframedomain

`func (o *EzsignfoldersignerassociationCreateEmbeddedUrlV1Request) HasSIframedomain() bool`

HasSIframedomain returns a boolean if a field has been set.

### GetBIsIframe

`func (o *EzsignfoldersignerassociationCreateEmbeddedUrlV1Request) GetBIsIframe() bool`

GetBIsIframe returns the BIsIframe field if non-nil, zero value otherwise.

### GetBIsIframeOk

`func (o *EzsignfoldersignerassociationCreateEmbeddedUrlV1Request) GetBIsIframeOk() (*bool, bool)`

GetBIsIframeOk returns a tuple with the BIsIframe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBIsIframe

`func (o *EzsignfoldersignerassociationCreateEmbeddedUrlV1Request) SetBIsIframe(v bool)`

SetBIsIframe sets BIsIframe field to given value.

### HasBIsIframe

`func (o *EzsignfoldersignerassociationCreateEmbeddedUrlV1Request) HasBIsIframe() bool`

HasBIsIframe returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


