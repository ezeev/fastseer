
# FastSeer Admin API
My API usually works as expected.

Table of Contents

1. [Rules Management API](#v1/rules)

<a name="v1/rules"></a>

## v1/rules

| Specification | Value |
|-----|-----|
| Resource Path | /v1/rules |
| API Version | 1.0.0 |
| BasePath for the API | https://shopify-app.fastseer.com/api/ |
| Consumes | application/json |
| Produces |  |



### Operations


| Resource Path | Operation | Description |
|-----|-----|-----|
| /rules | [PUT](#putRules) | Adds or updates rules based. You can pass as many rules to be updated as desired. The id attribute is the unique key. If the id does not exist, then a new rule is created, otherwise the existing rule is overwritten. |
| /rules | [DELETE](#deleteRules) | Deletes rules based on ids. You can pass as many id params as desired. |
| /rules | [GET](#getRules) | Retrieves orders for given customer defined by customer ID |



<a name="putRules"></a>

#### API: /rules (PUT)


Adds or updates rules based. You can pass as many rules to be updated as desired. The id attribute is the unique key. If the id does not exist, then a new rule is created, otherwise the existing rule is overwritten.



| Param Name | Param Type | Data Type | Description | Required? |
|-----|-----|-----|-----|-----|
| body | body | [SearchRuleList](#github.com.ezeev.fastseer.SearchRuleList) | A json array of SearchRules (see SearchRule model) | Yes |
| shop | query | string | Your shop url i.e. fastseer.myshopify.com | Yes |
| hmac | query | string | Secure hmac string from Shopify admin | Yes |
| locale | query | string | locale for the shop. i.e. en | Yes |
| timestamp | query | int | Timestamp for hmac auth | Yes |


| Code | Type | Model | Message |
|-----|-----|-----|-----|
| 200 | array | [MessageResponse](#github.com.ezeev.fastseer.MessageResponse) |  |
| 400 | object | [ErrorResponse](#github.com.ezeev.fastseer.ErrorResponse) | Error when a delete fails |


<a name="deleteRules"></a>

#### API: /rules (DELETE)


Deletes rules based on ids. You can pass as many id params as desired.



| Param Name | Param Type | Data Type | Description | Required? |
|-----|-----|-----|-----|-----|
| id | query | string | id of the rule to delete | Yes |
| shop | query | string | Your shop url i.e. fastseer.myshopify.com | Yes |
| hmac | query | string | Secure hmac string from Shopify admin | Yes |
| locale | query | string | locale for the shop. i.e. en | Yes |
| timestamp | query | int | Timestamp for hmac auth | Yes |


| Code | Type | Model | Message |
|-----|-----|-----|-----|
| 200 | array | [MessageResponse](#github.com.ezeev.fastseer.MessageResponse) |  |
| 400 | object | [ErrorResponse](#github.com.ezeev.fastseer.ErrorResponse) | Error when a delete fails |


<a name="getRules"></a>

#### API: /rules (GET)


Retrieves orders for given customer defined by customer ID



| Param Name | Param Type | Data Type | Description | Required? |
|-----|-----|-----|-----|-----|
| q | query | string | Rules query | Yes |
| shop | query | string | Your shop url i.e. fastseer.myshopify.com | Yes |
| hmac | query | string | Secure hmac string from Shopify admin | Yes |
| locale | query | string | locale for the shop. i.e. en | Yes |
| timestamp | query | int | Timestamp for hmac auth | Yes |


| Code | Type | Model | Message |
|-----|-----|-----|-----|
| 200 | array | [SearchRule](#github.com.ezeev.fastseer.rules.SearchRule) |  |
| 400 | object | [ErrorResponse](#github.com.ezeev.fastseer.ErrorResponse) | Error when a rules query fails |




### Models

<a name="github.com.ezeev.fastseer.ErrorResponse"></a>

#### ErrorResponse

| Field Name (alphabetical) | Field Type | Description |
|-----|-----|-----|
| error | error |  |
| message | string |  |

<a name="github.com.ezeev.fastseer.MessageResponse"></a>

#### MessageResponse

| Field Name (alphabetical) | Field Type | Description |
|-----|-----|-----|
| message | string |  |

<a name="github.com.ezeev.fastseer.SearchRuleList"></a>

#### SearchRuleList

| Field Name (alphabetical) | Field Type | Description |
|-----|-----|-----|

<a name="github.com.ezeev.fastseer.rules.SearchRule"></a>

#### SearchRule

| Field Name (alphabetical) | Field Type | Description |
|-----|-----|-----|
| actAddBqs_ss | array |  |
| actAddFacetFields_ss | array |  |
| actAddFqs_ss | array |  |
| actReplaceQuery_s | string |  |
| containsAnyQueryTriggers_txt | array |  |
| containsFqs_ss | array |  |
| id | string |  |
| matchQueryTriggers_ss | array |  |
| name_s | string |  |
| order_i | int |  |
| tags_ss | array |  |


