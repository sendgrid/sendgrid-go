# Event

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventName** | [**EventName**](EventName.md) | Name of event |
**Processed** | **string** | The date when the event was processed |
**Reason** | **string** | Explanation of what caused the message to be \"bounced\", \"deferred\", or \"blocked\". Usually contains error message from the server - e.g. message from gmail why mail was deferred. |[optional] 
**AttemptNum** | **int32** | Used with \"deferred\" events to indicate the attempt number out of 10. One \"deferred\" entry will exists under events array for each time a message was deferred with error message from the server.  |[optional] 
**Url** | **string** | Used with \"clicked\" event to indicate which url the user clicked. |[optional] 
**BounceType** | [**BounceType**](BounceType.md) | Use to distinguish between types of bounces |
**HttpUserAgent** | **string** | Client recipient used to click or open message |
**MxServer** | **string** | The MX server that received the email. For example, mx.gmail.com |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


