# EzsignfoldersignerassociationCreateEmbeddedUrlV2Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SReturnUrl** | Pointer to **string** | The return Url to redirect after the signing is completed  **Warning** Due to the potential for Ezsignsigners to cancel redirection, close their browser post-signing, or spoof the landing URL, it&#39;s advisable not to solely depend on the sReturnUrl for accurate status within your integration.  Once the Ezsignsigner finishes, they are directed back to your application. Your application can retain transaction state details by either storing data in a cookie or incorporating query parameters in the sReturnUrl. For example: https://www.example.com/sReturnUrl?sSessionID&#x3D;ABC123  The actual url that will be called will have an extra url parameter appended to give details about the process. The possible values are listed in the table below. For example: https://www.example.com/sReturnUrl?sSessionID&#x3D;ABC123&amp;eEzsignEvent&#x3D;CompletedEzsignfolder   |**Query parameters appended**| |---| |eEzsignEvent|   |**eEzsignEvent**|**Description**| |---|---| |SessionTimeout|The session timed out| |SessionLogout|The Ezsignsigner signed out| |DeclinedTermOfUse|The Ezsignsigner refused the terms| |DeclinedSign|The Ezsignsigner refused to sign| |Reassigned|The Ezsignsigner reassigned his signatures to someone else| |CompletedStep|The Ezsignsigner completed his step. There is other signatures to complete the Ezsigndocument| |CompletedEzsignfolder|The Ezsignfolder is completed. Everyone signed their signatures| | [optional] 
**SIframeOrigin** | Pointer to **string** | This is used to fill the header frame ancestor to allow the application to embed the Ezsign signature page | [optional] 

## Methods

### NewEzsignfoldersignerassociationCreateEmbeddedUrlV2Request

`func NewEzsignfoldersignerassociationCreateEmbeddedUrlV2Request() *EzsignfoldersignerassociationCreateEmbeddedUrlV2Request`

NewEzsignfoldersignerassociationCreateEmbeddedUrlV2Request instantiates a new EzsignfoldersignerassociationCreateEmbeddedUrlV2Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsignfoldersignerassociationCreateEmbeddedUrlV2RequestWithDefaults

`func NewEzsignfoldersignerassociationCreateEmbeddedUrlV2RequestWithDefaults() *EzsignfoldersignerassociationCreateEmbeddedUrlV2Request`

NewEzsignfoldersignerassociationCreateEmbeddedUrlV2RequestWithDefaults instantiates a new EzsignfoldersignerassociationCreateEmbeddedUrlV2Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSReturnUrl

`func (o *EzsignfoldersignerassociationCreateEmbeddedUrlV2Request) GetSReturnUrl() string`

GetSReturnUrl returns the SReturnUrl field if non-nil, zero value otherwise.

### GetSReturnUrlOk

`func (o *EzsignfoldersignerassociationCreateEmbeddedUrlV2Request) GetSReturnUrlOk() (*string, bool)`

GetSReturnUrlOk returns a tuple with the SReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSReturnUrl

`func (o *EzsignfoldersignerassociationCreateEmbeddedUrlV2Request) SetSReturnUrl(v string)`

SetSReturnUrl sets SReturnUrl field to given value.

### HasSReturnUrl

`func (o *EzsignfoldersignerassociationCreateEmbeddedUrlV2Request) HasSReturnUrl() bool`

HasSReturnUrl returns a boolean if a field has been set.

### GetSIframeOrigin

`func (o *EzsignfoldersignerassociationCreateEmbeddedUrlV2Request) GetSIframeOrigin() string`

GetSIframeOrigin returns the SIframeOrigin field if non-nil, zero value otherwise.

### GetSIframeOriginOk

`func (o *EzsignfoldersignerassociationCreateEmbeddedUrlV2Request) GetSIframeOriginOk() (*string, bool)`

GetSIframeOriginOk returns a tuple with the SIframeOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSIframeOrigin

`func (o *EzsignfoldersignerassociationCreateEmbeddedUrlV2Request) SetSIframeOrigin(v string)`

SetSIframeOrigin sets SIframeOrigin field to given value.

### HasSIframeOrigin

`func (o *EzsignfoldersignerassociationCreateEmbeddedUrlV2Request) HasSIframeOrigin() bool`

HasSIframeOrigin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


