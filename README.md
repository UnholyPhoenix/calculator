[[_paths]]
== Paths

[[_well-known_openid-configuration_get]]
=== Discovery
....
GET /.well-known/openid-configuration
....


==== Description
Discovery endpoint for retrieving openid configuration specific for this service. Endpoint is public and available to everyone.


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|User claims.|<<_well-known_openid-configuration_get_response_200,Response 200>>
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_well-known_openid-configuration_get_response_200]]
**Response 200**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**authorization_endpoint** +
__optional__|URL of OpenID Provider authorization endpoints.|string
|**claim_types_supported** +
__optional__|Holds supported claim types.|< object > array
|**claims_parameter_supported** +
__optional__|Is claims parameter supported.|boolean
|**claims_supported** +
__optional__|Holds available claim names.|< object > array
|**grant_types_supported** +
__optional__|Holds available granty type for OpenID.|< object > array
|**id_token_alg_values_supported** +
__optional__|Holds available encryption algorithms.|< object > array
|**issuer** +
__optional__|Issuer identifier, this value is the same as the one returned in id_token.|string
|**jwks_uri** +
__optional__|URL of OpenID Provider certificate endpoint.|string
|**request_parameter_supported** +
__optional__|Is request parameter supported.|boolean
|**request_uri_parameter_supported** +
__optional__|Is request_uri parameter supported.|boolean
|**require_request_uri_registration** +
__optional__|Is request_uri at registration required.|boolean
|**response_types_supported** +
__optional__|Holds available response types for OpenID.|< object > array
|**revocation_endpoint** +
__optional__|URL of OpenID Provider revoke endpoint.|string
|**scopes_supported** +
__optional__|Holds OpenID supported scopes(oauth2 scopes are omitted).|< object > array
|**service_documentation** +
__optional__|URL to service documentation.|string
|**subject_types_supported** +
__optional__|Holds available subject types.|< object > array
|**token_endpoint** +
__optional__|URL of OpenID Provider token endpoint.|string
|**token_endpoint_auth_methods_supported** +
__optional__|Holds available authentication methods for token endpoint.|< object > array
|**userinfo_endpoint** +
__optional__|URL of OpenID Provider userinfo endpoint.|string
|===


==== Produces

* `application/json`


==== Tags

* OpenID_Connect


==== Example HTTP response

===== Response 200
[source,json]
----
{
  "issuer" : "https://api.mdi.ericsson.net",
  "authorization_endpoint" : "https://api.mdi.ericsson.net/v1/oauth2/oidc/authorize",
  "token_endpoint" : "https://api.mdi.ericsson.net/v1/oauth2/oidc/token",
  "userinfo_endpoint" : "https://api.mdi.ericsson.net/v1/oauth2/oidc/token",
  "revocation_endpoint" : "https://api.mdi.ericsson.net/v1/oauth2/oidc/revoke",
  "jwks_uri" : "https://api.mdi.ericsson.net/v1/oauth2/oidc/certs",
  "scopes_supported" : [ "openid", "profile", "email", "phone", "offline_access" ],
  "response_types_supported" : [ "code" ],
  "grant_types_supported" : [ "authorization_code" ],
  "subject_types_supported" : [ "pairwise" ],
  "id_token_alg_values_supported" : [ "RS256" ],
  "token_endpoint_auth_methods_supported" : [ "client_secret_basic", "client_secret_post" ],
  "claim_types_supported" : [ "normal" ],
  "claims_supported" : [ "name", "email", "email_verified", "phone_number", "phone_number_verified", "update_at" ],
  "service_documentation" : "https://developer.mobileconnect.io/docs/",
  "claims_parameter_supported" : false,
  "request_parameter_supported" : false,
  "request_uri_parameter_supported" : false,
  "require_request_uri_registration" : false
}
----


[[_application_provision_post]]
=== Create application
....
POST /application/provision
....


==== Description
Creates new application.


==== Parameters

[options="header", cols=".^2a,.^3a,.^4a"]
|===
|Type|Name|Schema
|**Body**|**Application request body** +
__optional__|<<_application_provision_post_application_request_body,Application request body>>
|===

[[_application_provision_post_application_request_body]]
**Application request body**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**backgroundUrl** +
__optional__|Application background image url.|string
|**category** +
__optional__|Application category|string
|**description** +
__optional__|Application's description|string
|**label** +
__required__|Application's name|string
|**logoUrl** +
__optional__|Application logo image url.|string
|**mcTrustLevel** +
__optional__|Application trust level. Posible levels are 0, 1, 2.|integer
|**platformUri** +
__optional__|Array of platform uri objects.|< <<_application_provision_post_platformuri,platformUri>> > array
|**provisionerForMno** +
__optional__|ID of MNO for which the application serves as provisioning app.|string
|**requireLoa** +
__optional__|Restricts which LoA is minimal for login (eg. loa2 or loa3). Empty string means that there is no restriction.|string
|**scope** +
__optional__|Array of scopes. Note: if not set, scope gets value openid by default.|< string > array
|**scopeDescription** +
__optional__|Description, reasoning about requested scopes|string
|**shortNames** +
__optional__|Array of application short names.|< string > array
|**url** +
__required__|URL of application's website|string
|**visibility** +
__optional__|Application visibility. Possible values: pri (private), pro (protected), pub (public)|string
|===

[[_application_provision_post_platformuri]]
**platformUri**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**type** +
__optional__|Platform name|string
|**url** +
__optional__|Platform uri|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**201**|A representation of an application.|<<_application_provision_post_response_201,Response 201>>
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_application_provision_post_response_400,Response 400>>
|**401**|Access denied. Authentication is required to proceed processing request.|No Content
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**404**|Not Found. The requested resource could not be found but may be available again in the future. Subsequent requests by the client are permissible.|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_application_provision_post_response_201]]
**Response 201**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**author** +
__optional__|Author info|<<_application_provision_post_author,author>>
|**createdAt** +
__optional__|Applications's create date.|string
|**description** +
__optional__|Applications's description.|string
|**id** +
__optional__|Application's id. +
**Pattern** : `"[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}"`|string
|**image** +
__optional__|Set of application image links with all available sizes.|<<_application_provision_post_image,image>>
|**label** +
__optional__|Application's name.|string
|**privileges** +
__optional__|List of privileges application has.|<<_application_provision_post_privileges,privileges>>
|**promotedImage** +
__optional__|Application's promoted image.|<<_application_provision_post_promotedimage,promotedImage>>
|**publisher** +
__optional__|Publisher info.|<<_application_provision_post_publisher,publisher>>
|**requireLoa** +
__optional__|Restricts which LoA is minimal for login (eg. loa2 or loa3). Empty string means that there is no restriction.|string
|**scope** +
__optional__|Array of scopes.|< string > array
|**scopeDescription** +
__optional__|Description, reasoning about requested scopes.|string
|**shortNames** +
__optional__|Array of application short names.|< string > array
|**url** +
__optional__|URL of application's website.|string
|===

[[_application_provision_post_author]]
**author**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**email** +
__optional__|Author's E-mail.|string
|**name** +
__optional__|Author's name.|string
|===

[[_application_provision_post_image]]
**image**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**120x120** +
__optional__|Image link|string
|**240x240** +
__optional__|Image link|string
|**45x45** +
__optional__|Image link|string
|**480x480** +
__optional__|Image link|string
|**60x60** +
__optional__|Image link|string
|**90x90** +
__optional__|Image link|string
|===

[[_application_provision_post_privileges]]
**privileges**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**authorizator** +
__optional__|Defines if application is authorizator.|boolean
|**autogrant** +
__optional__|Defines if application scopes can be autogranted.|boolean
|**provision** +
__optional__|Defines if application has provisioner privileges.|boolean
|**provisionDescriptor** +
__optional__|Defines if application has descriptor provision privilege.|boolean
|**spUserProvision** +
__optional__|Defines if application has privileges to provision users.|boolean
|**system** +
__optional__|Defines if it is system application.|boolean
|**trusted** +
__optional__|Defines if application is trusted.|boolean
|**ump** +
__optional__|Defines if application acts as an User Management Portal|boolean
|===

[[_application_provision_post_promotedimage]]
**promotedImage**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**282x174** +
__optional__|Image link|string
|===

[[_application_provision_post_publisher]]
**publisher**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**email** +
__optional__|Publisher's E-mail|string
|**name** +
__optional__|Publisher's name|string
|===

[[_application_provision_post_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request, label_missing_parameter, label_wrong_parameter, url_missing_parameter, url_wrong_parameter, visibility_wrong_parameter, category_wrong_parameter, platform_uri_type_wrong_parameter, require_loa_wrong_type_parameter, require_loa_invalid_format_parameter, mc_trust_level_parameter_format, mno_does_not_exist, provisioner_wrong_format)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* Provision


==== Example HTTP response

===== Response 201
[source,json]
----
{
  "id" : "dde9536c-2ca2-4fdc-a710-17ef0a58ede8",
  "label" : "FakeAnApp v2",
  "shortNames" : [ "FakeAnApp v2" ],
  "url" : "http://awesomeapp.example.com",
  "image" : {
    "45x45" : "http://mobileconnect.io/images/5a486dee-718d-4caf-8e10-c332024c4a32-45x45.jpg",
    "60x60" : "http://mobileconnect.io/images/5a486dee-718d-4caf-8e10-c332024c4a32-60x60.jpg",
    "90x90" : "http://mobileconnect.io/images/5a486dee-718d-4caf-8e10-c332024c4a32-90x90.jpg",
    "120x120" : "http://mobileconnect.io/images/5a486dee-718d-4caf-8e10-c332024c4a32-120x120.jpg",
    "240x240" : "http://mobileconnect.io/images/5a486dee-718d-4caf-8e10-c332024c4a32-240x240.jpg",
    "480x480" : "http://mobileconnect.io/images/5a486dee-718d-4caf-8e10-c332024c4a32-480x480.jpg"
  },
  "promotedImage" : {
    "282x174" : "http://mobileconnect.io/images/50c23757-9dfc-4736-af29-93dace422949-282x174.jpg"
  },
  "author" : {
    "name" : "null",
    "email" : "author@etal.io"
  },
  "createdAt" : "2017-10-02T11:55:58+0000",
  "description" : "Be a description",
  "requireLoa" : "",
  "scopeDescription" : "Reasoning about the scopes that this app is requesting",
  "scope" : [ "scope.r", "scope.w" ],
  "publisher" : {
    "name" : "Mr. Publisher",
    "email" : "author@etal.io"
  },
  "privileges" : {
    "trusted" : false,
    "authorizator" : false,
    "system" : false,
    "autogrant" : false,
    "ump" : false,
    "provision" : true,
    "provisionDescriptor" : false,
    "spUserProvision" : false
  },
  "links" : [ {
    "rel" : "author",
    "href" : "https://api.mdi.ericsson.net/v1/profile/dde9536c-2ca2-4fdc-a710-17ef0a58ede8"
  }, {
    "rel" : "application",
    "href" : "https://api.mdi.ericsson.net/v1/application/dde9536c-2ca2-4fdc-a710-17ef0a58ede8"
  }, {
    "rel" : "keys",
    "href" : "https://api.mdi.ericsson.net/v1/application/dde9536c-2ca2-4fdc-a710-17ef0a58ede8/keys"
  } ],
  "platforms" : [ "web", "ios", "android" ],
  "category" : {
    "id" : "3cf1cd0e-ddf3-4bd3-85e8-45ad9b99621d",
    "name" : "Books"
  },
  "platformUri" : [ {
    "type" : "web",
    "url" : "http:://www.app.dev"
  }, {
    "type" : "ios",
    "url" : "http:://www.app.dev1"
  }, {
    "type" : "android",
    "url" : "http:://www.app.dev2"
  } ],
  "featured" : true,
  "mcPrivileges" : {
    "trustLevel" : 2
  },
  "logoUrl" : "",
  "backgroundUrl" : "",
  "visibility" : "pub",
  "provisionerForMno" : "905538d4-fa92-4305-aa5b-38383b274544"
}
----


===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_application_applicationid_get]]
=== Get application resource
....
GET /application/{applicationId}
....


==== Description
Returns requested application.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Path**|**applicationId** +
__required__|Application id|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|A representation of an application.|<<_application_applicationid_get_response_200,Response 200>>
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_application_applicationid_get_response_400,Response 400>>
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**404**|Not Found. The requested resource could not be found but may be available again in the future. Subsequent requests by the client are permissible.|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_application_applicationid_get_response_200]]
**Response 200**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**author** +
__optional__|Author info|<<_application_applicationid_get_author,author>>
|**createdAt** +
__optional__|Applications's create date.|string
|**description** +
__optional__|Applications's description.|string
|**id** +
__optional__|Application's id. +
**Pattern** : `"[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}"`|string
|**image** +
__optional__|Set of application image links with all available sizes.|<<_application_applicationid_get_image,image>>
|**label** +
__optional__|Application's name.|string
|**privileges** +
__optional__|List of privileges application has.|<<_application_applicationid_get_privileges,privileges>>
|**promotedImage** +
__optional__|Application's promoted image.|<<_application_applicationid_get_promotedimage,promotedImage>>
|**publisher** +
__optional__|Publisher info.|<<_application_applicationid_get_publisher,publisher>>
|**requireLoa** +
__optional__|Restricts which LoA is minimal for login (eg. loa2 or loa3). Empty string means that there is no restriction.|string
|**scope** +
__optional__|Array of scopes.|< string > array
|**scopeDescription** +
__optional__|Description, reasoning about requested scopes.|string
|**shortNames** +
__optional__|Array of application short names.|< string > array
|**url** +
__optional__|URL of application's website.|string
|===

[[_application_applicationid_get_author]]
**author**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**email** +
__optional__|Author's E-mail.|string
|**name** +
__optional__|Author's name.|string
|===

[[_application_applicationid_get_image]]
**image**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**120x120** +
__optional__|Image link|string
|**240x240** +
__optional__|Image link|string
|**45x45** +
__optional__|Image link|string
|**480x480** +
__optional__|Image link|string
|**60x60** +
__optional__|Image link|string
|**90x90** +
__optional__|Image link|string
|===

[[_application_applicationid_get_privileges]]
**privileges**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**authorizator** +
__optional__|Defines if application is authorizator.|boolean
|**autogrant** +
__optional__|Defines if application scopes can be autogranted.|boolean
|**provision** +
__optional__|Defines if application has provisioner privileges.|boolean
|**provisionDescriptor** +
__optional__|Defines if application has descriptor provision privilege.|boolean
|**spUserProvision** +
__optional__|Defines if application has privileges to provision users.|boolean
|**system** +
__optional__|Defines if it is system application.|boolean
|**trusted** +
__optional__|Defines if application is trusted.|boolean
|**ump** +
__optional__|Defines if application acts as an User Management Portal|boolean
|===

[[_application_applicationid_get_promotedimage]]
**promotedImage**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**282x174** +
__optional__|Image link|string
|===

[[_application_applicationid_get_publisher]]
**publisher**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**email** +
__optional__|Publisher's E-mail|string
|**name** +
__optional__|Publisher's name|string
|===

[[_application_applicationid_get_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (no_route, bad_request, visibility_wrong_parameter, category_wrong_parameter, platform_uri_type_wrong_parameter, require_loa_wrong_type_parameter, require_loa_invalid_format_parameter)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* Application


==== Example HTTP response

===== Response 200
[source,json]
----
{
  "id" : "dde9536c-2ca2-4fdc-a710-17ef0a58ede8",
  "label" : "FakeAnApp v2",
  "shortNames" : [ "FakeAnApp v2" ],
  "url" : "http://awesomeapp.example.com",
  "image" : {
    "45x45" : "http://mobileconnect.io/images/5a486dee-718d-4caf-8e10-c332024c4a32-45x45.jpg",
    "60x60" : "http://mobileconnect.io/images/5a486dee-718d-4caf-8e10-c332024c4a32-60x60.jpg",
    "90x90" : "http://mobileconnect.io/images/5a486dee-718d-4caf-8e10-c332024c4a32-90x90.jpg",
    "120x120" : "http://mobileconnect.io/images/5a486dee-718d-4caf-8e10-c332024c4a32-120x120.jpg",
    "240x240" : "http://mobileconnect.io/images/5a486dee-718d-4caf-8e10-c332024c4a32-240x240.jpg",
    "480x480" : "http://mobileconnect.io/images/5a486dee-718d-4caf-8e10-c332024c4a32-480x480.jpg"
  },
  "promotedImage" : {
    "282x174" : "http://mobileconnect.io/images/50c23757-9dfc-4736-af29-93dace422949-282x174.jpg"
  },
  "author" : {
    "name" : "null",
    "email" : "author@etal.io"
  },
  "createdAt" : "2017-10-02T11:55:58+0000",
  "description" : "Be a description",
  "requireLoa" : "",
  "scopeDescription" : "Reasoning about the scopes that this app is requesting",
  "scope" : [ "scope.r", "scope.w" ],
  "publisher" : {
    "name" : "Mr. Publisher",
    "email" : "author@etal.io"
  },
  "privileges" : {
    "trusted" : false,
    "authorizator" : false,
    "system" : false,
    "autogrant" : false,
    "ump" : false,
    "provision" : true,
    "provisionDescriptor" : false,
    "spUserProvision" : false
  },
  "links" : [ {
    "rel" : "author",
    "href" : "https://api.mdi.ericsson.net/v1/profile/dde9536c-2ca2-4fdc-a710-17ef0a58ede8"
  }, {
    "rel" : "application",
    "href" : "https://api.mdi.ericsson.net/v1/application/dde9536c-2ca2-4fdc-a710-17ef0a58ede8"
  }, {
    "rel" : "keys",
    "href" : "https://api.mdi.ericsson.net/v1/application/dde9536c-2ca2-4fdc-a710-17ef0a58ede8/keys"
  } ],
  "platforms" : [ "web", "ios", "android" ],
  "category" : {
    "id" : "3cf1cd0e-ddf3-4bd3-85e8-45ad9b99621d",
    "name" : "Books"
  },
  "platformUri" : [ {
    "type" : "web",
    "url" : "http:://www.app.dev"
  }, {
    "type" : "ios",
    "url" : "http:://www.app.dev1"
  }, {
    "type" : "android",
    "url" : "http:://www.app.dev2"
  } ],
  "featured" : true,
  "mcPrivileges" : {
    "trustLevel" : 2
  },
  "logoUrl" : "",
  "backgroundUrl" : "",
  "visibility" : "pub",
  "provisionerForMno" : "905538d4-fa92-4305-aa5b-38383b274544"
}
----


===== Response 400
[source,json]
----
{
  "error" : "no_route",
  "error_description" : "Route does not exist. Posible Id missmatch."
}
----


[[_application_applicationid_put]]
=== Update application
....
PUT /application/{applicationId}
....


==== Description
Updates new application. Latest developers EULA must be accepted to allow own application modifications.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Path**|**applicationId** +
__required__|Application id|string
|**Body**|**Application request body** +
__optional__||<<_application_applicationid_put_application_request_body,Application request body>>
|===

[[_application_applicationid_put_application_request_body]]
**Application request body**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**backgroundUrl** +
__optional__|Application background image url.|string
|**category** +
__optional__|Application category|string
|**description** +
__optional__|Application's description|string
|**label** +
__required__|Application's name|string
|**logoUrl** +
__optional__|Application logo image url.|string
|**platformUri** +
__optional__|Array of platform uri objects.|< <<_application_applicationid_put_platformuri,platformUri>> > array
|**requireLoa** +
__optional__|Restricts which LoA is minimal for login (eg. loa2 or loa3). Empty string means that there is no restriction.|string
|**scope** +
__optional__|Array of scopes|< string > array
|**scopeDescription** +
__optional__|Description, reasoning about requested scopes|string
|**shortNames** +
__optional__|Array of application short names.|< string > array
|**url** +
__required__|URL of application's website|string
|**visibility** +
__optional__|Application visibility. Possible values: pri (private), pro (protected), pub (public)|string
|===

[[_application_applicationid_put_platformuri]]
**platformUri**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**type** +
__optional__|Platform name|string
|**url** +
__optional__|Platform uri|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**202**|A representation of an application.|<<_application_applicationid_put_response_202,Response 202>>
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_application_applicationid_put_response_400,Response 400>>
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_application_applicationid_put_response_202]]
**Response 202**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**author** +
__optional__|Author info|<<_application_applicationid_put_author,author>>
|**createdAt** +
__optional__|Applications's create date.|string
|**description** +
__optional__|Applications's description.|string
|**id** +
__optional__|Application's id. +
**Pattern** : `"[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}"`|string
|**image** +
__optional__|Set of application image links with all available sizes.|<<_application_applicationid_put_image,image>>
|**label** +
__optional__|Application's name.|string
|**privileges** +
__optional__|List of privileges application has.|<<_application_applicationid_put_privileges,privileges>>
|**promotedImage** +
__optional__|Application's promoted image.|<<_application_applicationid_put_promotedimage,promotedImage>>
|**publisher** +
__optional__|Publisher info.|<<_application_applicationid_put_publisher,publisher>>
|**requireLoa** +
__optional__|Restricts which LoA is minimal for login (eg. loa2 or loa3). Empty string means that there is no restriction.|string
|**scope** +
__optional__|Array of scopes.|< string > array
|**scopeDescription** +
__optional__|Description, reasoning about requested scopes.|string
|**shortNames** +
__optional__|Array of application short names.|< string > array
|**url** +
__optional__|URL of application's website.|string
|===

[[_application_applicationid_put_author]]
**author**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**email** +
__optional__|Author's E-mail.|string
|**name** +
__optional__|Author's name.|string
|===

[[_application_applicationid_put_image]]
**image**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**120x120** +
__optional__|Image link|string
|**240x240** +
__optional__|Image link|string
|**45x45** +
__optional__|Image link|string
|**480x480** +
__optional__|Image link|string
|**60x60** +
__optional__|Image link|string
|**90x90** +
__optional__|Image link|string
|===

[[_application_applicationid_put_privileges]]
**privileges**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**authorizator** +
__optional__|Defines if application is authorizator.|boolean
|**autogrant** +
__optional__|Defines if application scopes can be autogranted.|boolean
|**provision** +
__optional__|Defines if application has provisioner privileges.|boolean
|**provisionDescriptor** +
__optional__|Defines if application has descriptor provision privilege.|boolean
|**spUserProvision** +
__optional__|Defines if application has privileges to provision users.|boolean
|**system** +
__optional__|Defines if it is system application.|boolean
|**trusted** +
__optional__|Defines if application is trusted.|boolean
|**ump** +
__optional__|Defines if application acts as an User Management Portal|boolean
|===

[[_application_applicationid_put_promotedimage]]
**promotedImage**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**282x174** +
__optional__|Image link|string
|===

[[_application_applicationid_put_publisher]]
**publisher**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**email** +
__optional__|Publisher's E-mail|string
|**name** +
__optional__|Publisher's name|string
|===

[[_application_applicationid_put_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (no_route, bad_request, visibility_wrong_parameter, category_wrong_parameter, platform_uri_type_wrong_parameter, require_loa_wrong_type_parameter, require_loa_invalid_format_parameter)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* Application


==== Example HTTP response

===== Response 202
[source,json]
----
{
  "id" : "dde9536c-2ca2-4fdc-a710-17ef0a58ede8",
  "label" : "FakeAnApp v2",
  "shortNames" : [ "FakeAnApp v2" ],
  "url" : "http://awesomeapp.example.com",
  "image" : {
    "45x45" : "http://mobileconnect.io/images/5a486dee-718d-4caf-8e10-c332024c4a32-45x45.jpg",
    "60x60" : "http://mobileconnect.io/images/5a486dee-718d-4caf-8e10-c332024c4a32-60x60.jpg",
    "90x90" : "http://mobileconnect.io/images/5a486dee-718d-4caf-8e10-c332024c4a32-90x90.jpg",
    "120x120" : "http://mobileconnect.io/images/5a486dee-718d-4caf-8e10-c332024c4a32-120x120.jpg",
    "240x240" : "http://mobileconnect.io/images/5a486dee-718d-4caf-8e10-c332024c4a32-240x240.jpg",
    "480x480" : "http://mobileconnect.io/images/5a486dee-718d-4caf-8e10-c332024c4a32-480x480.jpg"
  },
  "promotedImage" : {
    "282x174" : "http://mobileconnect.io/images/50c23757-9dfc-4736-af29-93dace422949-282x174.jpg"
  },
  "author" : {
    "name" : "null",
    "email" : "author@etal.io"
  },
  "createdAt" : "2017-10-02T11:55:58+0000",
  "description" : "Be a description",
  "requireLoa" : "",
  "scopeDescription" : "Reasoning about the scopes that this app is requesting",
  "scope" : [ "scope.r", "scope.w" ],
  "publisher" : {
    "name" : "Mr. Publisher",
    "email" : "author@etal.io"
  },
  "privileges" : {
    "trusted" : false,
    "authorizator" : false,
    "system" : false,
    "autogrant" : false,
    "ump" : false,
    "provision" : true,
    "provisionDescriptor" : false,
    "spUserProvision" : false
  },
  "links" : [ {
    "rel" : "author",
    "href" : "https://api.mdi.ericsson.net/v1/profile/dde9536c-2ca2-4fdc-a710-17ef0a58ede8"
  }, {
    "rel" : "application",
    "href" : "https://api.mdi.ericsson.net/v1/application/dde9536c-2ca2-4fdc-a710-17ef0a58ede8"
  }, {
    "rel" : "keys",
    "href" : "https://api.mdi.ericsson.net/v1/application/dde9536c-2ca2-4fdc-a710-17ef0a58ede8/keys"
  } ],
  "platforms" : [ "web", "ios", "android" ],
  "category" : {
    "id" : "3cf1cd0e-ddf3-4bd3-85e8-45ad9b99621d",
    "name" : "Books"
  },
  "platformUri" : [ {
    "type" : "web",
    "url" : "http:://www.app.dev"
  }, {
    "type" : "ios",
    "url" : "http:://www.app.dev1"
  }, {
    "type" : "android",
    "url" : "http:://www.app.dev2"
  } ],
  "featured" : true,
  "mcPrivileges" : {
    "trustLevel" : 2
  },
  "logoUrl" : "",
  "backgroundUrl" : "",
  "visibility" : "pub",
  "provisionerForMno" : "905538d4-fa92-4305-aa5b-38383b274544"
}
----


===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_application_applicationid_delete]]
=== Remove an application
....
DELETE /application/{applicationId}
....


==== Description
Removes existing application. Latest developers EULA must be accepted to allow application removal.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Path**|**applicationId** +
__required__|Application id|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**204**|On success 204 code is returned.|No Content
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_application_applicationid_delete_response_400,Response 400>>
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_application_applicationid_delete_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* Application


==== Example HTTP response

===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_application_applicationid_image_post]]
=== Add image
....
POST /application/{applicationId}/image
....


==== Description
Adds image to application.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Path**|**applicationId** +
__required__|Application id|string
|**Body**|**Body should be the multipart/form-data of the attachment with key named 'image'. No Content Headers.** +
__optional__||object
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**201**|Image used by application API endpoint.|<<_application_applicationid_image_post_response_201,Response 201>>
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_application_applicationid_image_post_response_400,Response 400>>
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_application_applicationid_image_post_response_201]]
**Response 201**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**id** +
__optional__|Application image UUID|string
|**type** +
__optional__|Type of image|string
|**urls** +
__optional__|URLS of resized images|<<_application_applicationid_image_post_urls,urls>>
|===

[[_application_applicationid_image_post_urls]]
**urls**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**120x120** +
__optional__|URL of resized image 120x120|string
|**240x240** +
__optional__|URL of resized image 240x240|string
|**45x45** +
__optional__|URL of resized image 45x45|string
|**480x480** +
__optional__|URL of resized image 480x480|string
|**60x60** +
__optional__|URL of resized image 60x60|string
|**90x90** +
__optional__|URL of resized image 90x90|string
|===

[[_application_applicationid_image_post_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request, image_missing, image_wrong_type, image_wrong_format, image_too_big, images_not_resized)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* Application


==== Example HTTP response

===== Response 201
[source,json]
----
{
  "id" : "5a486dee-718d-4caf-8e10-c332024c4a32",
  "type" : "png",
  "urls" : {
    "45x45" : "http://cloudfront.net/images/applications/5a486dee-718d-4caf-8e10-c332024c4a32-45x45.png",
    "60x60" : "http://cloudfront.net/images/applications/5a486dee-718d-4caf-8e10-c332024c4a32-60x60.png",
    "90x90" : "http://cloudfront.net/images/applications/5a486dee-718d-4caf-8e10-c332024c4a32-90x90.png",
    "120x120" : "http://cloudfront.net/images/applications/5a486dee-718d-4caf-8e10-c332024c4a32-120x120.png",
    "240x240" : "http://cloudfront.net/images/applications/5a486dee-718d-4caf-8e10-c332024c4a32-240x240.png",
    "480x480" : "http://cloudfront.net/images/applications/5a486dee-718d-4caf-8e10-c332024c4a32-480x480.png"
  }
}
----


===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_application_applicationid_image_get]]
=== Returns image URL
....
GET /application/{applicationId}/image
....


==== Description
Returns image URL of specified size.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Path**|**applicationId** +
__required__|Application id|string
|**Query**|**size** +
__required__|Parameter for predefined image size. NOTE: Availible sizes are 45x45, 60x60, 90x90, 120x120, 240x240 and 480x480.|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|Image used by application API endpoint.|<<_application_applicationid_image_get_response_200,Response 200>>
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_application_applicationid_image_get_response_400,Response 400>>
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**404**|Not Found. The requested resource could not be found but may be available again in the future. Subsequent requests by the client are permissible.|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_application_applicationid_image_get_response_200]]
**Response 200**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**id** +
__optional__|Application image UUID|string
|**type** +
__optional__|Type of image|string
|**urls** +
__optional__|URLS of resized images|<<_application_applicationid_image_get_urls,urls>>
|===

[[_application_applicationid_image_get_urls]]
**urls**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**120x120** +
__optional__|URL of resized image 120x120|string
|**240x240** +
__optional__|URL of resized image 240x240|string
|**45x45** +
__optional__|URL of resized image 45x45|string
|**480x480** +
__optional__|URL of resized image 480x480|string
|**60x60** +
__optional__|URL of resized image 60x60|string
|**90x90** +
__optional__|URL of resized image 90x90|string
|===

[[_application_applicationid_image_get_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request, image_size_missing_parameter, invalid_image_size)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* Application


==== Example HTTP response

===== Response 200
[source,json]
----
{
  "id" : "5a486dee-718d-4caf-8e10-c332024c4a32",
  "type" : "png",
  "urls" : {
    "45x45" : "http://cloudfront.net/images/applications/5a486dee-718d-4caf-8e10-c332024c4a32-45x45.png",
    "60x60" : "http://cloudfront.net/images/applications/5a486dee-718d-4caf-8e10-c332024c4a32-60x60.png",
    "90x90" : "http://cloudfront.net/images/applications/5a486dee-718d-4caf-8e10-c332024c4a32-90x90.png",
    "120x120" : "http://cloudfront.net/images/applications/5a486dee-718d-4caf-8e10-c332024c4a32-120x120.png",
    "240x240" : "http://cloudfront.net/images/applications/5a486dee-718d-4caf-8e10-c332024c4a32-240x240.png",
    "480x480" : "http://cloudfront.net/images/applications/5a486dee-718d-4caf-8e10-c332024c4a32-480x480.png"
  }
}
----


===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_application_applicationid_image_delete]]
=== Remove application image
....
DELETE /application/{applicationId}/image
....


==== Description
Removes application image.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Path**|**applicationId** +
__required__|Application id|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**204**|On success 204 code is returned.|No Content
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_application_applicationid_image_delete_response_400,Response 400>>
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_application_applicationid_image_delete_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* Application


==== Example HTTP response

===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_application_applicationid_key_applicationkey_get]]
=== Get application key
....
GET /application/{applicationId}/key/{applicationKey}
....


==== Description
Returns information related to a requested application's key


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Path**|**applicationId** +
__required__|Application id|string
|**Path**|**applicationKey** +
__required__|Application key id|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|Authorization key and secret of application.|<<_application_applicationid_key_applicationkey_get_response_200,Response 200>>
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_application_applicationid_key_applicationkey_get_response_400,Response 400>>
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**404**|Not Found. The requested resource could not be found but may be available again in the future. Subsequent requests by the client are permissible.|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_application_applicationid_key_applicationkey_get_response_200]]
**Response 200**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**algorithm** +
__optional__|Key encryption algorithm|string
|**bundleId** +
__optional__|Bundle ID|string
|**client_notification_endpoint** +
__optional__|Token notification endpoint of the client|string
|**key** +
__required__|Client key|string
|**label** +
__optional__|Describes who keys are intended for|string
|**links** +
__optional__|Links associated with application key|< <<_application_applicationid_key_applicationkey_get_links,links>> > array
|**packageName** +
__optional__|Package name|string
|**platform** +
__optional__|Platform which application operates on|string
|**redirect_uris** +
__optional__|Array of redirection URIs|< string > array
|===

[[_application_applicationid_key_applicationkey_get_links]]
**links**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**href** +
__optional__|Link path|string
|**rel** +
__optional__|Link relation|string
|===

[[_application_applicationid_key_applicationkey_get_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* Application


==== Example HTTP response

===== Response 200
[source,json]
----
{
  "label" : "Developer Client",
  "key" : "123456789abcdefghijklmnoprstuvz1",
  "redirect_uris" : [ "http://example.com" ],
  "client_notification_endpoint" : "https://example.com/cb",
  "platform" : "web",
  "links" : [ {
    "rel" : "keys",
    "href" : "https://api.mdi.ericsson.net/v1/application/dde9536c-2ca2-4fdc-a710-17ef0a58ede8/keys"
  } ]
}
----


===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_application_applicationid_key_applicationkey_put]]
=== Update application key
....
PUT /application/{applicationId}/key/{applicationKey}
....


==== Description
Updates existing set of keys.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Path**|**applicationId** +
__required__|Application id|string
|**Path**|**applicationKey** +
__required__|Application key id|string
|**Body**|**Application key request body** +
__optional__||<<_application_applicationid_key_applicationkey_put_application_key_request_body,Application key request body>>
|===

[[_application_applicationid_key_applicationkey_put_application_key_request_body]]
**Application key request body**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**algorithm** +
__optional__|Holds encryption algorithm, if not set default 'RS256' is selected|string
|**bundleId** +
__optional__|Holds bundle id - iOS identifier|string
|**client_notification_endpoint** +
__optional__|Holds token notification endpoint of the client. It has to be HTTPS and it enables async callback flow for server-based authentication requests.|string
|**label** +
__required__|Describes who keys are intended for|string
|**packageName** +
__optional__|Holds package name - Android identifier|string
|**platform** +
__optional__|Holds application platform.|string
|**redirect_uris** +
__optional__|Holds array of redirect URIs for key.|< string > array
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**202**|Authorization key and secret of application.|<<_application_applicationid_key_applicationkey_put_response_202,Response 202>>
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_application_applicationid_key_applicationkey_put_response_400,Response 400>>
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**404**|Not Found. The requested resource could not be found but may be available again in the future. Subsequent requests by the client are permissible.|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_application_applicationid_key_applicationkey_put_response_202]]
**Response 202**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**algorithm** +
__optional__|Key encryption algorithm|string
|**bundleId** +
__optional__|Bundle ID|string
|**client_notification_endpoint** +
__optional__|Token notification endpoint of the client|string
|**key** +
__required__|Client key|string
|**label** +
__optional__|Describes who keys are intended for|string
|**links** +
__optional__|Links associated with application key|< <<_application_applicationid_key_applicationkey_put_links,links>> > array
|**packageName** +
__optional__|Package name|string
|**platform** +
__optional__|Platform which application operates on|string
|**redirect_uris** +
__optional__|Array of redirection URIs|< string > array
|===

[[_application_applicationid_key_applicationkey_put_links]]
**links**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**href** +
__optional__|Link path|string
|**rel** +
__optional__|Link relation|string
|===

[[_application_applicationid_key_applicationkey_put_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request, label_missing_parameter, platform_wrong_parameter, redirect_uris_missing_parameter, no_route, invalid_algorithm)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* Application


==== Example HTTP response

===== Response 202
[source,json]
----
{
  "label" : "Developer Client",
  "key" : "123456789abcdefghijklmnoprstuvz1",
  "redirect_uris" : [ "http://example.com" ],
  "client_notification_endpoint" : "https://example.com/cb",
  "platform" : "web",
  "links" : [ {
    "rel" : "keys",
    "href" : "https://api.mdi.ericsson.net/v1/application/dde9536c-2ca2-4fdc-a710-17ef0a58ede8/keys"
  } ]
}
----


===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_application_applicationid_key_applicationkey_delete]]
=== Revoke application key
....
DELETE /application/{applicationId}/key/{applicationKey}
....


==== Description
Removes existing set of keys.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Path**|**applicationId** +
__required__|Application id|string
|**Path**|**applicationKey** +
__required__|Application key id|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**204**|On success 204 code is returned.|No Content
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_application_applicationid_key_applicationkey_delete_response_400,Response 400>>
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**404**|Not Found. The requested resource could not be found but may be available again in the future. Subsequent requests by the client are permissible.|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_application_applicationid_key_applicationkey_delete_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request, no_route)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* Application


==== Example HTTP response

===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_application_applicationid_keys_post]]
=== Create application key
....
POST /application/{applicationId}/keys
....


==== Description
Creates a new set of keys.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Path**|**applicationId** +
__required__|Application id|string
|**Body**|**Application key request body** +
__optional__||<<_application_applicationid_keys_post_application_key_request_body,Application key request body>>
|===

[[_application_applicationid_keys_post_application_key_request_body]]
**Application key request body**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**algorithm** +
__optional__|Holds encryption algorithm, if not set default 'RS256' is selected|string
|**bundleId** +
__optional__|Holds bundle id - iOS identifier|string
|**client_notification_endpoint** +
__optional__|Holds token notification endpoint of the client. It has to be HTTPS and it enables async callback flow for server-based authentication requests.|string
|**label** +
__required__|Describes who keys are intended for|string
|**packageName** +
__optional__|Holds package name - Android identifier|string
|**platform** +
__required__|Holds application platform. Note: Required is not explicit, if not set default "web" is used.|string
|**redirect_uris** +
__required__|Holds array of redirect URIs for key. Note: Required is not explicit for mobile platforms, redirect uri is autogenerated to "mdi{keyId}://authentication".|< string > array
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**201**|Authorization key and secret of application.|<<_application_applicationid_keys_post_response_201,Response 201>>
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_application_applicationid_keys_post_response_400,Response 400>>
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**404**|Not Found. The requested resource could not be found but may be available again in the future. Subsequent requests by the client are permissible.|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_application_applicationid_keys_post_response_201]]
**Response 201**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**algorithm** +
__optional__|Key encryption algorithm|string
|**bundleId** +
__optional__|Bundle ID|string
|**client_notification_endpoint** +
__optional__|Token notification endpoint of the client|string
|**key** +
__optional__|Client key|string
|**label** +
__optional__|Describes who keys are intended for|string
|**links** +
__optional__|Links associated with application key|< <<_application_applicationid_keys_post_links,links>> > array
|**packageName** +
__optional__|Package name|string
|**redirect_uris** +
__optional__|Array of redirection URIs|< string > array
|**secret** +
__optional__|Client secret|string
|===

[[_application_applicationid_keys_post_links]]
**links**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**href** +
__optional__|Link path|string
|**rel** +
__optional__|Link relation|string
|===

[[_application_applicationid_keys_post_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request, label_missing_parameter, platform_wrong_parameter, redirect_uris_missing_parameter, no_route, invalid_algorithm)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* Application


==== Example HTTP response

===== Response 201
[source,json]
----
{
  "label" : "Developer Client",
  "key" : "123456789abcdefghijklmnoprstuvz1",
  "secret" : "123456789abcdefghijklmnoprstuvz1123456789abcdefghijklmnoprstuvz1",
  "redirect_uris" : [ "http://example.com" ],
  "client_notification_endpoint" : "https://example.com/cb",
  "platform" : "web",
  "links" : [ {
    "rel" : "keys",
    "href" : "https://api.mdi.ericsson.net/v1/application/dde9536c-2ca2-4fdc-a710-17ef0a58ede8/keys"
  } ]
}
----


===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_application_applicationid_keys_get]]
=== Get application keys collection
....
GET /application/{applicationId}/keys
....


==== Description
Returns application's key collection.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Path**|**applicationId** +
__required__|Application id|string
|**Query**|**limit** +
__optional__|Pagination parameter, indicates how many records should be returned by page. NOTE: Default value is 20 and Max value is 50|string
|**Query**|**page** +
__optional__|Pagination parameter, indicates which page is to be returned from database. NOTE: Default value is 1.|number
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|Collection of Application keys.|< <<_application_applicationid_keys_get_response_200,Response 200>> > array
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_application_applicationid_keys_get_response_400,Response 400>>
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_application_applicationid_keys_get_response_200]]
**Response 200**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**algorithm** +
__optional__|Key encryption algorithm|string
|**bundleId** +
__optional__|Bundle ID|string
|**client_notification_endpoint** +
__optional__|Token notification endpoint of the client|string
|**key** +
__required__|Client key|string
|**label** +
__optional__|Describes who keys are intended for|string
|**links** +
__optional__|Links associated with application key|< <<_application_applicationid_keys_get_links,links>> > array
|**packageName** +
__optional__|Package name|string
|**platform** +
__optional__|Platform which application operates on|string
|**redirect_uris** +
__optional__|Array of redirection URIs|< string > array
|===

[[_application_applicationid_keys_get_links]]
**links**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**href** +
__optional__|Link path|string
|**rel** +
__optional__|Link relation|string
|===

[[_application_applicationid_keys_get_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (wrong_page_parameter, wrong_limit_parameter, limit_exceeds_max)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* Application


==== Example HTTP response

===== Response 200
[source,json]
----
[ {
  "label" : "Developer Client",
  "key" : "123456789abcdefghijklmnoprstuvz1",
  "redirect_uris" : [ "http://example.com" ],
  "client_notification_endpoint" : "https://example.com/cb",
  "platform" : "web",
  "links" : [ {
    "rel" : "keys",
    "href" : "https://api.mdi.ericsson.net/v1/application/dde9536c-2ca2-4fdc-a710-17ef0a58ede8/keys"
  } ]
} ]
----


===== Response 400
[source,json]
----
{
  "error" : "wrong_page_parameter",
  "error_description" : "Parameter is not a positive integer."
}
----


[[_application_applicationid_promoted-image_post]]
=== Add promoted image
....
POST /application/{applicationId}/promoted-image
....


==== Description
Adds promoted image to application.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Path**|**applicationId** +
__required__|Application id|string
|**Body**|**Body should be the multipart/form-data of the attachment with key named 'image'. No Content Headers.** +
__optional__||object
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**201**|Promoted image used by application API endpoint.|<<_application_applicationid_promoted-image_post_response_201,Response 201>>
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_application_applicationid_promoted-image_post_response_400,Response 400>>
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_application_applicationid_promoted-image_post_response_201]]
**Response 201**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**id** +
__optional__|Application image UUID|string
|**type** +
__optional__|Type of image|string
|**urls** +
__optional__|URLs of resized images|<<_application_applicationid_promoted-image_post_urls,urls>>
|===

[[_application_applicationid_promoted-image_post_urls]]
**urls**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**282x174** +
__optional__|URL of resized image|string
|===

[[_application_applicationid_promoted-image_post_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request, image_missing, image_wrong_type, image_wrong_format, image_too_big, images_not_resized)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* Application


==== Example HTTP response

===== Response 201
[source,json]
----
{
  "id" : "5a486dee-718d-4caf-8e10-c332024c4a32",
  "type" : "png",
  "urls" : {
    "282x174" : "http://cloudfront.net/images/applications/5a486dee-718d-4caf-8e10-c332024c4a32-282x174.png"
  }
}
----


===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_application_applicationid_promoted-image_get]]
=== Returns promoted image URL
....
GET /application/{applicationId}/promoted-image
....


==== Description
Returns promoted image URL of specified size.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Path**|**applicationId** +
__required__|Application id|string
|**Query**|**size** +
__required__|Parameter for predefined image size. NOTE: Availible sizes are 282x174.|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|Promoted image used by application API endpoint.|<<_application_applicationid_promoted-image_get_response_200,Response 200>>
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_application_applicationid_promoted-image_get_response_400,Response 400>>
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**404**|Not Found. The requested resource could not be found but may be available again in the future. Subsequent requests by the client are permissible.|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_application_applicationid_promoted-image_get_response_200]]
**Response 200**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**id** +
__optional__|Application image UUID|string
|**type** +
__optional__|Type of image|string
|**urls** +
__optional__|URLs of resized images|<<_application_applicationid_promoted-image_get_urls,urls>>
|===

[[_application_applicationid_promoted-image_get_urls]]
**urls**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**282x174** +
__optional__|URL of resized image|string
|===

[[_application_applicationid_promoted-image_get_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request, image_size_missing_parameter, invalid_image_size)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* Application


==== Example HTTP response

===== Response 200
[source,json]
----
{
  "id" : "5a486dee-718d-4caf-8e10-c332024c4a32",
  "type" : "png",
  "urls" : {
    "282x174" : "http://cloudfront.net/images/applications/5a486dee-718d-4caf-8e10-c332024c4a32-282x174.png"
  }
}
----


===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_application_applicationid_promoted-image_delete]]
=== Remove application promoted image
....
DELETE /application/{applicationId}/promoted-image
....


==== Description
Removes application promoted image.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Path**|**applicationId** +
__required__|Application id|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**204**|On success 204 code is returned.|No Content
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_application_applicationid_promoted-image_delete_response_400,Response 400>>
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_application_applicationid_promoted-image_delete_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* Application


==== Example HTTP response

===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_application_applicationid_provision_put]]
=== Update application
....
PUT /application/{applicationId}/provision
....


==== Description
Updates data for an existing application.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Path**|**applicationId** +
__required__|Application id|string
|**Body**|**Application request body** +
__optional__||<<_application_applicationid_provision_put_application_request_body,Application request body>>
|===

[[_application_applicationid_provision_put_application_request_body]]
**Application request body**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**backgroundUrl** +
__optional__|Application background image url.|string
|**category** +
__optional__|Application category|string
|**description** +
__optional__|Application's description|string
|**label** +
__required__|Application's name|string
|**logoUrl** +
__optional__|Application logo image url.|string
|**mcTrustLevel** +
__optional__|Application trust level. Posible levels are 0, 1, 2.|integer
|**platformUri** +
__optional__|Array of platform uri objects.|< <<_application_applicationid_provision_put_platformuri,platformUri>> > array
|**provisionerForMno** +
__optional__|ID of MNO for which the application serves as provisioning app.|string
|**requireLoa** +
__optional__|Restricts which LoA is minimal for login (eg. loa2 or loa3). Empty string means that there is no restriction.|string
|**scope** +
__optional__|Array of scopes. Note: if not set, scope gets value openid by default.|< string > array
|**scopeDescription** +
__optional__|Description, reasoning about requested scopes|string
|**shortNames** +
__optional__|Array of application short names.|< string > array
|**url** +
__required__|URL of application's website|string
|**visibility** +
__optional__|Application visibility. Possible values: pri (private), pro (protected), pub (public)|string
|===

[[_application_applicationid_provision_put_platformuri]]
**platformUri**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**type** +
__optional__|Platform name|string
|**url** +
__optional__|Platform uri|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**202**|A representation of an application.|<<_application_applicationid_provision_put_response_202,Response 202>>
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_application_applicationid_provision_put_response_400,Response 400>>
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**404**|Not Found. The requested resource could not be found but may be available again in the future. Subsequent requests by the client are permissible.|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_application_applicationid_provision_put_response_202]]
**Response 202**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**author** +
__optional__|Author info|<<_application_applicationid_provision_put_author,author>>
|**createdAt** +
__optional__|Applications's create date.|string
|**description** +
__optional__|Applications's description.|string
|**id** +
__optional__|Application's id. +
**Pattern** : `"[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}"`|string
|**image** +
__optional__|Set of application image links with all available sizes.|<<_application_applicationid_provision_put_image,image>>
|**label** +
__optional__|Application's name.|string
|**privileges** +
__optional__|List of privileges application has.|<<_application_applicationid_provision_put_privileges,privileges>>
|**promotedImage** +
__optional__|Application's promoted image.|<<_application_applicationid_provision_put_promotedimage,promotedImage>>
|**publisher** +
__optional__|Publisher info.|<<_application_applicationid_provision_put_publisher,publisher>>
|**requireLoa** +
__optional__|Restricts which LoA is minimal for login (eg. loa2 or loa3). Empty string means that there is no restriction.|string
|**scope** +
__optional__|Array of scopes.|< string > array
|**scopeDescription** +
__optional__|Description, reasoning about requested scopes.|string
|**shortNames** +
__optional__|Array of application short names.|< string > array
|**url** +
__optional__|URL of application's website.|string
|===

[[_application_applicationid_provision_put_author]]
**author**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**email** +
__optional__|Author's E-mail.|string
|**name** +
__optional__|Author's name.|string
|===

[[_application_applicationid_provision_put_image]]
**image**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**120x120** +
__optional__|Image link|string
|**240x240** +
__optional__|Image link|string
|**45x45** +
__optional__|Image link|string
|**480x480** +
__optional__|Image link|string
|**60x60** +
__optional__|Image link|string
|**90x90** +
__optional__|Image link|string
|===

[[_application_applicationid_provision_put_privileges]]
**privileges**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**authorizator** +
__optional__|Defines if application is authorizator.|boolean
|**autogrant** +
__optional__|Defines if application scopes can be autogranted.|boolean
|**provision** +
__optional__|Defines if application has provisioner privileges.|boolean
|**provisionDescriptor** +
__optional__|Defines if application has descriptor provision privilege.|boolean
|**spUserProvision** +
__optional__|Defines if application has privileges to provision users.|boolean
|**system** +
__optional__|Defines if it is system application.|boolean
|**trusted** +
__optional__|Defines if application is trusted.|boolean
|**ump** +
__optional__|Defines if application acts as an User Management Portal|boolean
|===

[[_application_applicationid_provision_put_promotedimage]]
**promotedImage**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**282x174** +
__optional__|Image link|string
|===

[[_application_applicationid_provision_put_publisher]]
**publisher**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**email** +
__optional__|Publisher's E-mail|string
|**name** +
__optional__|Publisher's name|string
|===

[[_application_applicationid_provision_put_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request, label_missing_parameter, label_wrong_parameter, url_missing_parameter, url_wrong_parameter, visibility_wrong_parameter, category_wrong_parameter, platform_uri_type_wrong_parameter, require_loa_wrong_type_parameter, require_loa_invalid_format_parameter, mc_trust_level_parameter_format, mno_does_not_exist, provisioner_wrong_format)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* Provision


==== Example HTTP response

===== Response 202
[source,json]
----
{
  "id" : "dde9536c-2ca2-4fdc-a710-17ef0a58ede8",
  "label" : "FakeAnApp v2",
  "shortNames" : [ "FakeAnApp v2" ],
  "url" : "http://awesomeapp.example.com",
  "image" : {
    "45x45" : "http://mobileconnect.io/images/5a486dee-718d-4caf-8e10-c332024c4a32-45x45.jpg",
    "60x60" : "http://mobileconnect.io/images/5a486dee-718d-4caf-8e10-c332024c4a32-60x60.jpg",
    "90x90" : "http://mobileconnect.io/images/5a486dee-718d-4caf-8e10-c332024c4a32-90x90.jpg",
    "120x120" : "http://mobileconnect.io/images/5a486dee-718d-4caf-8e10-c332024c4a32-120x120.jpg",
    "240x240" : "http://mobileconnect.io/images/5a486dee-718d-4caf-8e10-c332024c4a32-240x240.jpg",
    "480x480" : "http://mobileconnect.io/images/5a486dee-718d-4caf-8e10-c332024c4a32-480x480.jpg"
  },
  "promotedImage" : {
    "282x174" : "http://mobileconnect.io/images/50c23757-9dfc-4736-af29-93dace422949-282x174.jpg"
  },
  "author" : {
    "name" : "null",
    "email" : "author@etal.io"
  },
  "createdAt" : "2017-10-02T11:55:58+0000",
  "description" : "Be a description",
  "requireLoa" : "",
  "scopeDescription" : "Reasoning about the scopes that this app is requesting",
  "scope" : [ "scope.r", "scope.w" ],
  "publisher" : {
    "name" : "Mr. Publisher",
    "email" : "author@etal.io"
  },
  "privileges" : {
    "trusted" : false,
    "authorizator" : false,
    "system" : false,
    "autogrant" : false,
    "ump" : false,
    "provision" : true,
    "provisionDescriptor" : false,
    "spUserProvision" : false
  },
  "links" : [ {
    "rel" : "author",
    "href" : "https://api.mdi.ericsson.net/v1/profile/dde9536c-2ca2-4fdc-a710-17ef0a58ede8"
  }, {
    "rel" : "application",
    "href" : "https://api.mdi.ericsson.net/v1/application/dde9536c-2ca2-4fdc-a710-17ef0a58ede8"
  }, {
    "rel" : "keys",
    "href" : "https://api.mdi.ericsson.net/v1/application/dde9536c-2ca2-4fdc-a710-17ef0a58ede8/keys"
  } ],
  "platforms" : [ "web", "ios", "android" ],
  "category" : {
    "id" : "3cf1cd0e-ddf3-4bd3-85e8-45ad9b99621d",
    "name" : "Books"
  },
  "platformUri" : [ {
    "type" : "web",
    "url" : "http:://www.app.dev"
  }, {
    "type" : "ios",
    "url" : "http:://www.app.dev1"
  }, {
    "type" : "android",
    "url" : "http:://www.app.dev2"
  } ],
  "featured" : true,
  "mcPrivileges" : {
    "trustLevel" : 2
  },
  "logoUrl" : "",
  "backgroundUrl" : "",
  "visibility" : "pub",
  "provisionerForMno" : "905538d4-fa92-4305-aa5b-38383b274544"
}
----


===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_application_applicationid_provision-key_post]]
=== Create application key
....
POST /application/{applicationId}/provision-key
....


==== Description
Creates new application key.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Path**|**applicationId** +
__required__|Application id|string
|**Body**|**Application key request body** +
__optional__||<<_application_applicationid_provision-key_post_application_key_request_body,Application key request body>>
|===

[[_application_applicationid_provision-key_post_application_key_request_body]]
**Application key request body**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**algorithm** +
__optional__|Holds encryption algorithm, if not set default 'RS256' is selected.|string
|**bundleId** +
__optional__|Holds bundle id - iOS identifier.|string
|**client_id** +
__optional__|Applicaton key id.|string
|**client_notification_endpoint** +
__optional__|Holds token notification endpoint of the client. It has to be HTTPS and it enables async callback flow for server-based authentication requests.|string
|**client_secret** +
__optional__|Application key secret. Note: If this parameter is not set, secret is generated.|string
|**deeplinkUri** +
__optional__|Deeplink URI associated with application key. It should be set only for Mobile Native applications capable of handling given deeplink.|string
|**label** +
__required__|Describes what keys are intended for.|string
|**packageName** +
__optional__|Holds package name - Android identifier.|string
|**platform** +
__optional__|Holds application platform. Note: Required is not explicit, if not set default "web" is used.|string
|**redirect_uris** +
__optional__|Holds array of redirect URIs for key. Note: Required is not explicit for mobile platforms, redirect uri is autogenerated to "mdi{keyId}://authentication".|< string > array
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**201**|A representation of an application key.|<<_application_applicationid_provision-key_post_response_201,Response 201>>
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_application_applicationid_provision-key_post_response_400,Response 400>>
|**401**|Access denied. Authentication is required to proceed processing request.|No Content
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**404**|Not Found. The requested resource could not be found but may be available again in the future. Subsequent requests by the client are permissible.|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_application_applicationid_provision-key_post_response_201]]
**Response 201**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**algorithm** +
__optional__|Key encryption algorithm|string
|**bundleId** +
__required__|Bundle ID|string
|**client_notification_endpoint** +
__optional__|Token notification endpoint of the client|string
|**deeplinkUri** +
__required__|Deeplink URI associated with application key. It should be set only for Mobile Native applications capable of handling given deeplink.|string
|**key** +
__optional__|Client key|string
|**label** +
__optional__|Describes who keys are intended for|string
|**packageName** +
__required__|Package name|string
|**platform** +
__optional__|Application platform. Note: Required is not explicit, if not set default "web" is used.|string
|**redirect_uris** +
__optional__|Array of redirection URIs|< string > array
|**secret** +
__optional__|Client secret|string
|===

[[_application_applicationid_provision-key_post_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request, label_missing_parameter, label_wrong_parameter, platform_wrong_parameter, invalid_algorithm, provision_client_id_missing_parameter, provision_client_id_wrong_parameter, provision_client_secret_wrong_parameter, application_key_duplicate, redirect_uris_missing_parameter, redirect_uris_wrong_format)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* Provision


==== Example HTTP response

===== Response 201
[source,json]
----
{
  "label" : "New test",
  "key" : "6167a276bfbe738d49bff262637fb34b",
  "platform" : "ios",
  "client_notification_endpoint" : "https://example.com/cb",
  "bundleId" : "ios.bundle.id",
  "packageName" : null,
  "algorithm" : "RS512",
  "deeplinkUri" : "",
  "secret" : "dd2aeac2775aa92b7b65d85c72744ade"
}
----


===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_application_applicationid_provision-key_applicationkey_put]]
=== Update application key
....
PUT /application/{applicationId}/provision-key/{applicationKey}
....


==== Description
Updates an existing key.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Path**|**applicationId** +
__required__|Application id|string
|**Path**|**applicationKey** +
__required__|Applicationkey id|string
|**Body**|**Application key request body** +
__optional__||<<_application_applicationid_provision-key_applicationkey_put_application_key_request_body,Application key request body>>
|===

[[_application_applicationid_provision-key_applicationkey_put_application_key_request_body]]
**Application key request body**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**algorithm** +
__optional__|Holds encryption algorithm, if not set default 'RS256' is selected.|string
|**bundleId** +
__optional__|Holds bundle id - iOS identifier.|string
|**client_id** +
__optional__|Applicaton key id.|string
|**client_notification_endpoint** +
__optional__|Holds token notification endpoint of the client. It has to be HTTPS and it enables async callback flow for server-based authentication requests.|string
|**client_secret** +
__optional__|Application key secret. Note: If this parameter is not set, secret is generated.|string
|**deeplinkUri** +
__optional__|Deeplink URI associated with application key. It should be set only for Mobile Native applications capable of handling given deeplink.|string
|**label** +
__required__|Describes what keys are intended for.|string
|**packageName** +
__optional__|Holds package name - Android identifier.|string
|**platform** +
__optional__|Holds application platform. Note: Required is not explicit, if not set default "web" is used.|string
|**redirect_uris** +
__optional__|Holds array of redirect URIs for key. Note: Required is not explicit for mobile platforms, redirect uri is autogenerated to "mdi{keyId}://authentication".|< string > array
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**202**|A representation of an application key.|<<_application_applicationid_provision-key_applicationkey_put_response_202,Response 202>>
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_application_applicationid_provision-key_applicationkey_put_response_400,Response 400>>
|**401**|Access denied. Authentication is required to proceed processing request.|No Content
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**404**|Not Found. The requested resource could not be found but may be available again in the future. Subsequent requests by the client are permissible.|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_application_applicationid_provision-key_applicationkey_put_response_202]]
**Response 202**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**algorithm** +
__optional__|Key encryption algorithm|string
|**bundleId** +
__required__|Bundle ID|string
|**client_notification_endpoint** +
__optional__|Token notification endpoint of the client|string
|**deeplinkUri** +
__required__|Deeplink URI associated with application key. It should be set only for Mobile Native applications capable of handling given deeplink.|string
|**key** +
__optional__|Client key|string
|**label** +
__optional__|Describes who keys are intended for|string
|**packageName** +
__required__|Package name|string
|**platform** +
__optional__|Application platform. Note: Required is not explicit, if not set default "web" is used.|string
|**redirect_uris** +
__optional__|Array of redirection URIs|< string > array
|===

[[_application_applicationid_provision-key_applicationkey_put_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request, label_missing_parameter, label_wrong_parameter, platform_wrong_parameter, invalid_algorithm, provision_client_id_missing_parameter, provision_client_id_wrong_parameter, provision_client_secret_wrong_parameter, application_key_duplicate, redirect_uris_missing_parameter, redirect_uris_wrong_format)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* Provision


==== Example HTTP response

===== Response 202
[source,json]
----
{
  "label" : "New updated test",
  "key" : "df3e3b9e060831e8fbc23d0f09951ca4",
  "platform" : "web",
  "redirect_uris" : [ "http://example.com" ],
  "client_notification_endpoint" : "https://example.com/cb",
  "packageName" : null,
  "algorithm" : "RS512",
  "deeplinkUri" : ""
}
----


===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_application_applicationid_user_post]]
=== Application user provisioning
....
POST /application/{applicationId}/user
....


==== Description
Application endpoint for provisioning a user profile. Basic Authorization is required.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Path**|**applicationId** +
__required__|Application id|string
|**Body**|**SP user provision request body** +
__optional__||<<_application_applicationid_user_post_sp_user_provision_request_body,SP user provision request body>>
|===

[[_application_applicationid_user_post_sp_user_provision_request_body]]
**SP user provision request body**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**email** +
__required__|User's email.|string
|**family_name** +
__optional__|User's family name|string
|**given_name** +
__optional__|User's given name|string
|**middle_name** +
__optional__|User's middle name|string
|**msisdn** +
__required__|User's MSISDN.|string
|**operator** +
__required__|ID of user's MNO.|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**201**|SP user provision response body.|<<_application_applicationid_user_post_response_201,Response 201>>
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_application_applicationid_user_post_response_400,Response 400>>
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**409**|Profile data conflict.|<<_application_applicationid_user_post_response_409,Response 409>>
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_application_applicationid_user_post_response_201]]
**Response 201**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**sub** +
__optional__|User's PCR.|string
|===

[[_application_applicationid_user_post_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request)
|**error_description** +
__optional__|string
|===

[[_application_applicationid_user_post_response_409]]
**Response 409**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**conflictingProperties** +
__required__|List of request parameters that do not match internal state.|< string > array
|===


==== Produces

* `application/json`


==== Tags

* Application


==== Example HTTP response

===== Response 201
[source,json]
----
{
  "sub" : "3b4a73812-7df7-46bf-b187-f725331868d056"
}
----


===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


===== Response 409
[source,json]
----
{
  "conflictingProperties" : [ "msisdn", "operator", "given_name" ]
}
----


[[_application_applicationid_user_pcr_delete]]
=== Application user deprovisioning
....
DELETE /application/{applicationId}/user/{pcr}
....


==== Description
Application endpoint for de-provisioning a user profile. Basic Authorization is required.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Path**|**applicationId** +
__required__|Application id|string
|**Path**|**pcr** +
__required__|Pseudonymous customer reference|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**204**|On success 204 code is returned.|No Content
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_application_applicationid_user_pcr_delete_response_400,Response 400>>
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_application_applicationid_user_pcr_delete_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request, no_route)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* Application


==== Example HTTP response

===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_applications_post]]
=== Create application
....
POST /applications
....


==== Description
Creates new application. Latest developers EULA must be accepted to allow application adding.


==== Parameters

[options="header", cols=".^2a,.^3a,.^4a"]
|===
|Type|Name|Schema
|**Body**|**Application request body** +
__optional__|<<_applications_post_application_request_body,Application request body>>
|===

[[_applications_post_application_request_body]]
**Application request body**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**backgroundUrl** +
__optional__|Application background image url.|string
|**category** +
__optional__|Application category|string
|**description** +
__optional__|Application's description|string
|**label** +
__required__|Application's name|string
|**logoUrl** +
__optional__|Application logo image url.|string
|**platformUri** +
__optional__|Array of platform uri objects.|< <<_applications_post_platformuri,platformUri>> > array
|**requireLoa** +
__optional__|Restricts which LoA is minimal for login (eg. loa2 or loa3). Empty string means that there is no restriction.|string
|**scope** +
__optional__|Array of scopes|< string > array
|**scopeDescription** +
__optional__|Description, reasoning about requested scopes|string
|**shortNames** +
__optional__|Array of application short names.|< string > array
|**url** +
__required__|URL of application's website|string
|**visibility** +
__optional__|Application visibility. Possible values: pri (private), pro (protected), pub (public)|string
|===

[[_applications_post_platformuri]]
**platformUri**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**type** +
__optional__|Platform name|string
|**url** +
__optional__|Platform uri|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**201**|A representation of an application.|<<_applications_post_response_201,Response 201>>
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_applications_post_response_400,Response 400>>
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_applications_post_response_201]]
**Response 201**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**author** +
__optional__|Author info|<<_applications_post_author,author>>
|**createdAt** +
__optional__|Applications's create date.|string
|**description** +
__optional__|Applications's description.|string
|**id** +
__optional__|Application's id. +
**Pattern** : `"[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}"`|string
|**image** +
__optional__|Set of application image links with all available sizes.|<<_applications_post_image,image>>
|**label** +
__optional__|Application's name.|string
|**privileges** +
__optional__|List of privileges application has.|<<_applications_post_privileges,privileges>>
|**promotedImage** +
__optional__|Application's promoted image.|<<_applications_post_promotedimage,promotedImage>>
|**publisher** +
__optional__|Publisher info.|<<_applications_post_publisher,publisher>>
|**requireLoa** +
__optional__|Restricts which LoA is minimal for login (eg. loa2 or loa3). Empty string means that there is no restriction.|string
|**scope** +
__optional__|Array of scopes.|< string > array
|**scopeDescription** +
__optional__|Description, reasoning about requested scopes.|string
|**shortNames** +
__optional__|Array of application short names.|< string > array
|**url** +
__optional__|URL of application's website.|string
|===

[[_applications_post_author]]
**author**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**email** +
__optional__|Author's E-mail.|string
|**name** +
__optional__|Author's name.|string
|===

[[_applications_post_image]]
**image**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**120x120** +
__optional__|Image link|string
|**240x240** +
__optional__|Image link|string
|**45x45** +
__optional__|Image link|string
|**480x480** +
__optional__|Image link|string
|**60x60** +
__optional__|Image link|string
|**90x90** +
__optional__|Image link|string
|===

[[_applications_post_privileges]]
**privileges**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**authorizator** +
__optional__|Defines if application is authorizator.|boolean
|**autogrant** +
__optional__|Defines if application scopes can be autogranted.|boolean
|**provision** +
__optional__|Defines if application has provisioner privileges.|boolean
|**provisionDescriptor** +
__optional__|Defines if application has descriptor provision privilege.|boolean
|**spUserProvision** +
__optional__|Defines if application has privileges to provision users.|boolean
|**system** +
__optional__|Defines if it is system application.|boolean
|**trusted** +
__optional__|Defines if application is trusted.|boolean
|**ump** +
__optional__|Defines if application acts as an User Management Portal|boolean
|===

[[_applications_post_promotedimage]]
**promotedImage**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**282x174** +
__optional__|Image link|string
|===

[[_applications_post_publisher]]
**publisher**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**email** +
__optional__|Publisher's E-mail|string
|**name** +
__optional__|Publisher's name|string
|===

[[_applications_post_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request, visibility_wrong_parameter, label_missing_parameter, url_missing_parameter, category_wrong_parameter, platform_uri_type_wrong_parameter, require_loa_wrong_type_parameter, require_loa_invalid_format_parameter)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* Application


==== Example HTTP response

===== Response 201
[source,json]
----
{
  "id" : "dde9536c-2ca2-4fdc-a710-17ef0a58ede8",
  "label" : "FakeAnApp v2",
  "shortNames" : [ "FakeAnApp v2" ],
  "url" : "http://awesomeapp.example.com",
  "image" : {
    "45x45" : "http://mobileconnect.io/images/5a486dee-718d-4caf-8e10-c332024c4a32-45x45.jpg",
    "60x60" : "http://mobileconnect.io/images/5a486dee-718d-4caf-8e10-c332024c4a32-60x60.jpg",
    "90x90" : "http://mobileconnect.io/images/5a486dee-718d-4caf-8e10-c332024c4a32-90x90.jpg",
    "120x120" : "http://mobileconnect.io/images/5a486dee-718d-4caf-8e10-c332024c4a32-120x120.jpg",
    "240x240" : "http://mobileconnect.io/images/5a486dee-718d-4caf-8e10-c332024c4a32-240x240.jpg",
    "480x480" : "http://mobileconnect.io/images/5a486dee-718d-4caf-8e10-c332024c4a32-480x480.jpg"
  },
  "promotedImage" : {
    "282x174" : "http://mobileconnect.io/images/50c23757-9dfc-4736-af29-93dace422949-282x174.jpg"
  },
  "author" : {
    "name" : "null",
    "email" : "author@etal.io"
  },
  "createdAt" : "2017-10-02T11:55:58+0000",
  "description" : "Be a description",
  "requireLoa" : "",
  "scopeDescription" : "Reasoning about the scopes that this app is requesting",
  "scope" : [ "scope.r", "scope.w" ],
  "publisher" : {
    "name" : "Mr. Publisher",
    "email" : "author@etal.io"
  },
  "privileges" : {
    "trusted" : false,
    "authorizator" : false,
    "system" : false,
    "autogrant" : false,
    "ump" : false,
    "provision" : true,
    "provisionDescriptor" : false,
    "spUserProvision" : false
  },
  "links" : [ {
    "rel" : "author",
    "href" : "https://api.mdi.ericsson.net/v1/profile/dde9536c-2ca2-4fdc-a710-17ef0a58ede8"
  }, {
    "rel" : "application",
    "href" : "https://api.mdi.ericsson.net/v1/application/dde9536c-2ca2-4fdc-a710-17ef0a58ede8"
  }, {
    "rel" : "keys",
    "href" : "https://api.mdi.ericsson.net/v1/application/dde9536c-2ca2-4fdc-a710-17ef0a58ede8/keys"
  } ],
  "platforms" : [ "web", "ios", "android" ],
  "category" : {
    "id" : "3cf1cd0e-ddf3-4bd3-85e8-45ad9b99621d",
    "name" : "Books"
  },
  "platformUri" : [ {
    "type" : "web",
    "url" : "http:://www.app.dev"
  }, {
    "type" : "ios",
    "url" : "http:://www.app.dev1"
  }, {
    "type" : "android",
    "url" : "http:://www.app.dev2"
  } ],
  "featured" : true,
  "mcPrivileges" : {
    "trustLevel" : 2
  },
  "logoUrl" : "",
  "backgroundUrl" : "",
  "visibility" : "pub",
  "provisionerForMno" : "905538d4-fa92-4305-aa5b-38383b274544"
}
----


===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_applications_get]]
=== Get application collection
....
GET /applications
....


==== Description
Returns collection of applications.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Query**|**author** +
__optional__|Parameter for filtering applications by author|string
|**Query**|**complete** +
__optional__|Parameter for showing also internal MDI apps that are normally hidden in this collection|boolean
|**Query**|**featured** +
__optional__|Parameter for filtering applications by featured tags|boolean
|**Query**|**key** +
__optional__|Parameter for filtering applications by client key|string
|**Query**|**limit** +
__optional__|Pagination parameter, indicates how many records should be returned by page. NOTE: Default value is 20 and Max value is 50|string
|**Query**|**operator** +
__optional__|Parameter for retrieving applications promoted by specific operator|string
|**Query**|**page** +
__optional__|Pagination parameter, indicates which page is to be returned from database. NOTE: Default value is 1.|number
|**Query**|**platform** +
__optional__|Parameter for filtering applications by platform|string
|**Query**|**q** +
__optional__|Parameter for searching by application names|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|Collection of Applications.|< <<_applications_get_response_200,Response 200>> > array
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_applications_get_response_400,Response 400>>
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_applications_get_response_200]]
**Response 200**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**author** +
__optional__|Author info|<<_applications_get_author,author>>
|**createdAt** +
__optional__|Applications's create date.|string
|**description** +
__optional__|Applications's description.|string
|**id** +
__optional__|Application's id. +
**Pattern** : `"[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}"`|string
|**image** +
__optional__|Set of application image links with all available sizes.|<<_applications_get_image,image>>
|**label** +
__optional__|Application's name.|string
|**privileges** +
__optional__|List of privileges application has.|<<_applications_get_privileges,privileges>>
|**promotedImage** +
__optional__|Application's promoted image.|<<_applications_get_promotedimage,promotedImage>>
|**publisher** +
__optional__|Publisher info.|<<_applications_get_publisher,publisher>>
|**requireLoa** +
__optional__|Restricts which LoA is minimal for login (eg. loa2 or loa3). Empty string means that there is no restriction.|string
|**scope** +
__optional__|Array of scopes.|< string > array
|**scopeDescription** +
__optional__|Description, reasoning about requested scopes.|string
|**shortNames** +
__optional__|Array of application short names.|< string > array
|**url** +
__optional__|URL of application's website.|string
|===

[[_applications_get_author]]
**author**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**email** +
__optional__|Author's E-mail.|string
|**name** +
__optional__|Author's name.|string
|===

[[_applications_get_image]]
**image**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**120x120** +
__optional__|Image link|string
|**240x240** +
__optional__|Image link|string
|**45x45** +
__optional__|Image link|string
|**480x480** +
__optional__|Image link|string
|**60x60** +
__optional__|Image link|string
|**90x90** +
__optional__|Image link|string
|===

[[_applications_get_privileges]]
**privileges**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**authorizator** +
__optional__|Defines if application is authorizator.|boolean
|**autogrant** +
__optional__|Defines if application scopes can be autogranted.|boolean
|**provision** +
__optional__|Defines if application has provisioner privileges.|boolean
|**provisionDescriptor** +
__optional__|Defines if application has descriptor provision privilege.|boolean
|**spUserProvision** +
__optional__|Defines if application has privileges to provision users.|boolean
|**system** +
__optional__|Defines if it is system application.|boolean
|**trusted** +
__optional__|Defines if application is trusted.|boolean
|**ump** +
__optional__|Defines if application acts as an User Management Portal|boolean
|===

[[_applications_get_promotedimage]]
**promotedImage**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**282x174** +
__optional__|Image link|string
|===

[[_applications_get_publisher]]
**publisher**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**email** +
__optional__|Publisher's E-mail|string
|**name** +
__optional__|Publisher's name|string
|===

[[_applications_get_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (wrong_page_parameter, wrong_limit_parameter, limit_exceeds_max)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* Application


==== Example HTTP response

===== Response 200
[source,json]
----
[ {
  "id" : "dde9536c-2ca2-4fdc-a710-17ef0a58ede8",
  "label" : "FakeAnApp v2",
  "shortNames" : [ "FakeAnApp v2" ],
  "url" : "http://awesomeapp.example.com",
  "description" : "Be a description",
  "scopeDescription" : "Reasoning about the scopes that this app is requesting",
  "scope" : [ "scope.r", "scope.w" ],
  "images" : [ {
    "id" : "5a486dee-718d-4caf-8e10-c332024c4a32",
    "type" : "icon-64x64",
    "href" : "5a486dee-718d-4caf-8e10-c332024c4a32.png"
  } ],
  "author" : {
    "name" : "null",
    "email" : "author@etal.io"
  },
  "publisher" : {
    "name" : "Mr. Publisher",
    "email" : "author@etal.io"
  },
  "trusted" : true,
  "links" : [ {
    "rel" : "author",
    "href" : "https://api.mdi.ericsson.net/v1/profile/dde9536c-2ca2-4fdc-a710-17ef0a58ede8"
  }, {
    "rel" : "application",
    "href" : "https://api.mdi.ericsson.net/v1/application/dde9536c-2ca2-4fdc-a710-17ef0a58ede8"
  }, {
    "rel" : "keys",
    "href" : "https://api.mdi.ericsson.net/v1/application/dde9536c-2ca2-4fdc-a710-17ef0a58ede8/keys"
  } ],
  "platforms" : [ "web", "ios", "android" ],
  "category" : {
    "id" : "3cf1cd0e-ddf3-4bd3-85e8-45ad9b99621d",
    "name" : "Books"
  },
  "platformUri" : [ {
    "type" : "web",
    "url" : "http:://www.app.dev"
  } ],
  "featured" : true
} ]
----


===== Response 400
[source,json]
----
{
  "error" : "wrong_page_parameter",
  "error_description" : "Parameter is not a positive integer."
}
----


[[_applications_featured_get]]
=== Get featured application collection
....
GET /applications/featured
....


==== Description
Returns featured application collection.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Query**|**limit** +
__optional__|Pagination parameter, indicates how many records should be returned by page. NOTE: Default value is 20 and Max value is 50|string
|**Query**|**page** +
__optional__|Pagination parameter, indicates which page is to be returned from database. NOTE: Default value is 1.|number
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|Collection of Applications.|< <<_applications_featured_get_response_200,Response 200>> > array
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_applications_featured_get_response_400,Response 400>>
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_applications_featured_get_response_200]]
**Response 200**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**author** +
__optional__|Author info|<<_applications_featured_get_author,author>>
|**createdAt** +
__optional__|Applications's create date.|string
|**description** +
__optional__|Applications's description.|string
|**id** +
__optional__|Application's id. +
**Pattern** : `"[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}"`|string
|**image** +
__optional__|Set of application image links with all available sizes.|<<_applications_featured_get_image,image>>
|**label** +
__optional__|Application's name.|string
|**privileges** +
__optional__|List of privileges application has.|<<_applications_featured_get_privileges,privileges>>
|**promotedImage** +
__optional__|Application's promoted image.|<<_applications_featured_get_promotedimage,promotedImage>>
|**publisher** +
__optional__|Publisher info.|<<_applications_featured_get_publisher,publisher>>
|**requireLoa** +
__optional__|Restricts which LoA is minimal for login (eg. loa2 or loa3). Empty string means that there is no restriction.|string
|**scope** +
__optional__|Array of scopes.|< string > array
|**scopeDescription** +
__optional__|Description, reasoning about requested scopes.|string
|**shortNames** +
__optional__|Array of application short names.|< string > array
|**url** +
__optional__|URL of application's website.|string
|===

[[_applications_featured_get_author]]
**author**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**email** +
__optional__|Author's E-mail.|string
|**name** +
__optional__|Author's name.|string
|===

[[_applications_featured_get_image]]
**image**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**120x120** +
__optional__|Image link|string
|**240x240** +
__optional__|Image link|string
|**45x45** +
__optional__|Image link|string
|**480x480** +
__optional__|Image link|string
|**60x60** +
__optional__|Image link|string
|**90x90** +
__optional__|Image link|string
|===

[[_applications_featured_get_privileges]]
**privileges**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**authorizator** +
__optional__|Defines if application is authorizator.|boolean
|**autogrant** +
__optional__|Defines if application scopes can be autogranted.|boolean
|**provision** +
__optional__|Defines if application has provisioner privileges.|boolean
|**provisionDescriptor** +
__optional__|Defines if application has descriptor provision privilege.|boolean
|**spUserProvision** +
__optional__|Defines if application has privileges to provision users.|boolean
|**system** +
__optional__|Defines if it is system application.|boolean
|**trusted** +
__optional__|Defines if application is trusted.|boolean
|**ump** +
__optional__|Defines if application acts as an User Management Portal|boolean
|===

[[_applications_featured_get_promotedimage]]
**promotedImage**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**282x174** +
__optional__|Image link|string
|===

[[_applications_featured_get_publisher]]
**publisher**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**email** +
__optional__|Publisher's E-mail|string
|**name** +
__optional__|Publisher's name|string
|===

[[_applications_featured_get_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (wrong_page_parameter, wrong_limit_parameter, limit_exceeds_max)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* Application


==== Example HTTP response

===== Response 200
[source,json]
----
[ {
  "id" : "dde9536c-2ca2-4fdc-a710-17ef0a58ede8",
  "label" : "FakeAnApp v2",
  "shortNames" : [ "FakeAnApp v2" ],
  "url" : "http://awesomeapp.example.com",
  "description" : "Be a description",
  "scopeDescription" : "Reasoning about the scopes that this app is requesting",
  "scope" : [ "scope.r", "scope.w" ],
  "images" : [ {
    "id" : "5a486dee-718d-4caf-8e10-c332024c4a32",
    "type" : "icon-64x64",
    "href" : "5a486dee-718d-4caf-8e10-c332024c4a32.png"
  } ],
  "author" : {
    "name" : "null",
    "email" : "author@etal.io"
  },
  "publisher" : {
    "name" : "Mr. Publisher",
    "email" : "author@etal.io"
  },
  "trusted" : true,
  "links" : [ {
    "rel" : "author",
    "href" : "https://api.mdi.ericsson.net/v1/profile/dde9536c-2ca2-4fdc-a710-17ef0a58ede8"
  }, {
    "rel" : "application",
    "href" : "https://api.mdi.ericsson.net/v1/application/dde9536c-2ca2-4fdc-a710-17ef0a58ede8"
  }, {
    "rel" : "keys",
    "href" : "https://api.mdi.ericsson.net/v1/application/dde9536c-2ca2-4fdc-a710-17ef0a58ede8/keys"
  } ],
  "platforms" : [ "web", "ios", "android" ],
  "category" : {
    "id" : "3cf1cd0e-ddf3-4bd3-85e8-45ad9b99621d",
    "name" : "Books"
  },
  "platformUri" : [ {
    "type" : "web",
    "url" : "http:://www.app.dev"
  } ],
  "featured" : true
} ]
----


===== Response 400
[source,json]
----
{
  "error" : "wrong_page_parameter",
  "error_description" : "Parameter is not a positive integer."
}
----


[[_applications_featured_put]]
=== Update featured application collection
....
PUT /applications/featured
....


==== Description
Updates featured application collection with an array of application IDs.


==== Parameters

[options="header", cols=".^2a,.^3a,.^4a"]
|===
|Type|Name|Schema
|**Body**|**Featured applications update request body** +
__optional__|< string > array
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**202**|Collection of Applications.|< <<_applications_featured_put_response_202,Response 202>> > array
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_applications_featured_put_response_400,Response 400>>
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_applications_featured_put_response_202]]
**Response 202**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**author** +
__optional__|Author info|<<_applications_featured_put_author,author>>
|**createdAt** +
__optional__|Applications's create date.|string
|**description** +
__optional__|Applications's description.|string
|**id** +
__optional__|Application's id. +
**Pattern** : `"[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}"`|string
|**image** +
__optional__|Set of application image links with all available sizes.|<<_applications_featured_put_image,image>>
|**label** +
__optional__|Application's name.|string
|**privileges** +
__optional__|List of privileges application has.|<<_applications_featured_put_privileges,privileges>>
|**promotedImage** +
__optional__|Application's promoted image.|<<_applications_featured_put_promotedimage,promotedImage>>
|**publisher** +
__optional__|Publisher info.|<<_applications_featured_put_publisher,publisher>>
|**requireLoa** +
__optional__|Restricts which LoA is minimal for login (eg. loa2 or loa3). Empty string means that there is no restriction.|string
|**scope** +
__optional__|Array of scopes.|< string > array
|**scopeDescription** +
__optional__|Description, reasoning about requested scopes.|string
|**shortNames** +
__optional__|Array of application short names.|< string > array
|**url** +
__optional__|URL of application's website.|string
|===

[[_applications_featured_put_author]]
**author**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**email** +
__optional__|Author's E-mail.|string
|**name** +
__optional__|Author's name.|string
|===

[[_applications_featured_put_image]]
**image**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**120x120** +
__optional__|Image link|string
|**240x240** +
__optional__|Image link|string
|**45x45** +
__optional__|Image link|string
|**480x480** +
__optional__|Image link|string
|**60x60** +
__optional__|Image link|string
|**90x90** +
__optional__|Image link|string
|===

[[_applications_featured_put_privileges]]
**privileges**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**authorizator** +
__optional__|Defines if application is authorizator.|boolean
|**autogrant** +
__optional__|Defines if application scopes can be autogranted.|boolean
|**provision** +
__optional__|Defines if application has provisioner privileges.|boolean
|**provisionDescriptor** +
__optional__|Defines if application has descriptor provision privilege.|boolean
|**spUserProvision** +
__optional__|Defines if application has privileges to provision users.|boolean
|**system** +
__optional__|Defines if it is system application.|boolean
|**trusted** +
__optional__|Defines if application is trusted.|boolean
|**ump** +
__optional__|Defines if application acts as an User Management Portal|boolean
|===

[[_applications_featured_put_promotedimage]]
**promotedImage**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**282x174** +
__optional__|Image link|string
|===

[[_applications_featured_put_publisher]]
**publisher**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**email** +
__optional__|Publisher's E-mail|string
|**name** +
__optional__|Publisher's name|string
|===

[[_applications_featured_put_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (not_an_array, application_could_not_be_found)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* Application


==== Example HTTP response

===== Response 202
[source,json]
----
[ {
  "id" : "dde9536c-2ca2-4fdc-a710-17ef0a58ede8",
  "label" : "FakeAnApp v2",
  "shortNames" : [ "FakeAnApp v2" ],
  "url" : "http://awesomeapp.example.com",
  "description" : "Be a description",
  "scopeDescription" : "Reasoning about the scopes that this app is requesting",
  "scope" : [ "scope.r", "scope.w" ],
  "images" : [ {
    "id" : "5a486dee-718d-4caf-8e10-c332024c4a32",
    "type" : "icon-64x64",
    "href" : "5a486dee-718d-4caf-8e10-c332024c4a32.png"
  } ],
  "author" : {
    "name" : "null",
    "email" : "author@etal.io"
  },
  "publisher" : {
    "name" : "Mr. Publisher",
    "email" : "author@etal.io"
  },
  "trusted" : true,
  "links" : [ {
    "rel" : "author",
    "href" : "https://api.mdi.ericsson.net/v1/profile/dde9536c-2ca2-4fdc-a710-17ef0a58ede8"
  }, {
    "rel" : "application",
    "href" : "https://api.mdi.ericsson.net/v1/application/dde9536c-2ca2-4fdc-a710-17ef0a58ede8"
  }, {
    "rel" : "keys",
    "href" : "https://api.mdi.ericsson.net/v1/application/dde9536c-2ca2-4fdc-a710-17ef0a58ede8/keys"
  } ],
  "platforms" : [ "web", "ios", "android" ],
  "category" : {
    "id" : "3cf1cd0e-ddf3-4bd3-85e8-45ad9b99621d",
    "name" : "Books"
  },
  "platformUri" : [ {
    "type" : "web",
    "url" : "http:://www.app.dev"
  } ],
  "featured" : true
} ]
----


===== Response 400
[source,json]
----
{
  "error" : "not_an_array",
  "error_description" : "Request content has incorrect form."
}
----


[[_applications_list_get]]
=== Application list report
....
GET /applications/list
....


==== Description
Application endpoint for retrieving list of applications. Authenticated access token is needed.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Query**|**complete** +
__optional__|Parameter for showing also internal MDI apps that are normally hidden in this collection.|boolean
|**Query**|**fields** +
__optional__|Whitespace separated string with application fields. If string is empty all fields are returned.|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|Applications response body.|<<_applications_list_get_response_200,Response 200>>
|**400**|Invalid request. For extra information response will contain extra information (error and extra_description) in json response.|No Content
|**401**|Access denied. Authentication is required to proceed processing request.|No Content
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_applications_list_get_response_200]]
**Response 200**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**applications** +
__optional__|Array of filtered application fields|< object > array
|===


==== Produces

* `application/json`


==== Tags

* Application


==== Example HTTP response

===== Response 200
[source,json]
----
{
  "applications" : [ {
    "uid" : "3b4a73812-7df7-46bf-b187-f725331868d056",
    "name" : "App1",
    "acl" : "",
    "visibility" : "pub"
  }, {
    "uid" : "268a3daea5-447a-4eda-99asf24-d6c7fb0f4f84",
    "name" : "App2",
    "acl" : "{\"mno\":[\"0191220d0f-6891-4f01-a87c-528b475ea1b42\"]}",
    "visibility" : "pub"
  } ]
}
----


[[_applications_platforms_get]]
=== Returns available platforms
....
GET /applications/platforms
....


==== Description
Returns available platforms.


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|Collection of Application platforms.|< string > array
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_applications_platforms_get_response_400,Response 400>>
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_applications_platforms_get_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* Application


==== Example HTTP response

===== Response 200
[source,json]
----
[ "web", "ios", "android" ]
----


===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_auth_account_init_post]]
=== Init transaction: account
....
POST /auth/account/init
....


==== Description
Endpoint for initializing account action.


==== Parameters

[options="header", cols=".^2a,.^3a,.^4a"]
|===
|Type|Name|Schema
|**Body**|**Authentication init request body** +
__optional__|<<_auth_account_init_post_authentication_init_request_body,Authentication init request body>>
|===

[[_auth_account_init_post_authentication_init_request_body]]
**Authentication init request body**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**actionId** +
__required__|Action id.|string
|**correlationId** +
__optional__|Correlation id used by frontend app to correlate transaction responses.|string
|**descriptorId** +
__required__|Descriptor id.|string
|**instanceIdentifier** +
__optional__|Identifier browser instance maing the request.|string
|**profileId** +
__required__|User's profile id.|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|Status response|<<_auth_account_init_post_response_200,Response 200>>
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_auth_account_init_post_response_400,Response 400>>
|**401**|Access denied. Authentication is required to proceed processing request.|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_auth_account_init_post_response_200]]
**Response 200**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**clientName** +
__optional__|Requesting SP's name returned when status is final and the transaction is of type authentication or authorization|string
|**correlationId** +
__optional__|CorrelationId sent on init request.|string
|**data** +
__optional__|Data that has to be used with next action - values of form fields, display fields and placeholders. See Pluggable authenticators documentation for more information.|<<_auth_account_init_post_data,data>>
|**errors** +
__optional__|Errors as result of previous transaction request. See Pluggable authenticators documentation for more information.|<<_auth_account_init_post_errors,errors>>
|**fingerprint** +
__optional__|Descriptor fingerprint. Should not change between requests. If it changes fresh descriptor has to be feched and action has to be restarted.|string
|**mainAction** +
__optional__|Transaction's main action, used to display confirmation screen.|string
|**next** +
__optional__|Array of next actions that have to be executed.|< string > array
|**nextTransaction** +
__optional__|Next action that has to be executed.|< string > array
|**redirectedTo** +
__optional__|Next action that has to be executed.|string
|**status** +
__required__|Current transaction status.|enum (ready, pending, done, bad_request, rejected, redirected, failed, closed, pendingElsewhere, handedOver, doneInAnotherTab, cancelled, expired)
|**transactionId** +
__required__|Transaction id.|string
|**type** +
__required__|Transaction type.|enum (simple, registration, registrationSaa, authentication, activation, account, action, consent, operation, administration, reclaimProfile)
|===

[[_auth_account_init_post_data]]
**data**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**$authenticator id** +
__optional__|Action specific data.|object
|===

[[_auth_account_init_post_errors]]
**errors**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**$authenticator id** +
__optional__|Action specific errors.|object
|===

[[_auth_account_init_post_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request, descriptorId_missing_parameter, profileId_missing_parameter, actionId_missing_parameter, profile_does_not_exist, invalid_request)
|**error_description** +
__optional__|string
|===


==== Consumes

* `application/json`


==== Produces

* `application/json`


==== Tags

* Descriptor


==== Example HTTP response

===== Response 200
[source,json]
----
{
  "status" : "ready",
  "type" : "authentication",
  "transactionId" : "78239500-74d6-11e6-acf7-1129c898746a",
  "next" : [ "sms::authenticateLoa2" ],
  "fingerprint" : "ba43edd03c1265f045a943aee8758d21",
  "mainAction" : "sms::authenticateLoa2",
  "correlationId" : "123456"
}
----


===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_auth_account_init_link_post]]
=== NewTransactionLink - Account
....
POST /auth/account/init/link
....


==== Description
Endpoint for generating account init public code.


==== Parameters

[options="header", cols=".^2a,.^3a,.^4a"]
|===
|Type|Name|Schema
|**Body**|**Authentication init request body** +
__optional__|<<_auth_account_init_link_post_authentication_init_request_body,Authentication init request body>>
|===

[[_auth_account_init_link_post_authentication_init_request_body]]
**Authentication init request body**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**actionId** +
__required__|Action id.|string
|**correlationId** +
__optional__|Correlation id used by frontend app to correlate transaction responses.|string
|**descriptorId** +
__required__|Descriptor id.|string
|**instanceIdentifier** +
__optional__|Identifier browser instance maing the request.|string
|**profileId** +
__required__|User's profile id.|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|Status response|<<_auth_account_init_link_post_response_200,Response 200>>
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_auth_account_init_link_post_response_400,Response 400>>
|**401**|Access denied. Authentication is required to proceed processing request.|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_auth_account_init_link_post_response_200]]
**Response 200**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**code** +
__required__|Code related to data stored in storage.|string
|**publicUrl** +
__required__|Public URL where new transaction is initiated from stored data.|string
|===

[[_auth_account_init_link_post_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request, descriptorId_missing_parameter, profileId_missing_parameter, actionId_missing_parameter, profile_does_not_exist, invalid_request)
|**error_description** +
__optional__|string
|===


==== Consumes

* `application/json`


==== Produces

* `application/json`


==== Tags

* Descriptor


==== Example HTTP response

===== Response 200
[source,json]
----
{
  "code" : "8c7067d9d774659e",
  "publicUrl" : "http://weblogin.etalio.dev/d/8c7067d9d774659e"
}
----


===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_auth_action_init_post]]
=== Init transaction: action
....
POST /auth/action/init
....


==== Description
Endpoint for initializing single action. NOTE: One of parameters 'msisdn’, ‘profileId’ or ‘sourceTransactionId’ is required in request.


==== Parameters

[options="header", cols=".^2a,.^3a,.^4a"]
|===
|Type|Name|Schema
|**Body**|**Authentication init request body** +
__optional__|<<_auth_action_init_post_authentication_init_request_body,Authentication init request body>>
|===

[[_auth_action_init_post_authentication_init_request_body]]
**Authentication init request body**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**actionId** +
__required__|Action id.|string
|**correlationId** +
__optional__|Correlation id used by frontend app to correlate transaction responses.|string
|**descriptorId** +
__required__|Descriptor id.|string
|**instanceIdentifier** +
__optional__|Identifier browser instance maing the request.|string
|**msisdn** +
__optional__|User's phone number.|string
|**profileId** +
__optional__|User's profile id.|string
|**sourceTransactionId** +
__optional__|Id of original transaction that demanded this action.|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|Status response|<<_auth_action_init_post_response_200,Response 200>>
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_auth_action_init_post_response_400,Response 400>>
|**401**|Access denied. Authentication is required to proceed processing request.|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_auth_action_init_post_response_200]]
**Response 200**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**clientName** +
__optional__|Requesting SP's name returned when status is final and the transaction is of type authentication or authorization|string
|**correlationId** +
__optional__|CorrelationId sent on init request.|string
|**data** +
__optional__|Data that has to be used with next action - values of form fields, display fields and placeholders. See Pluggable authenticators documentation for more information.|<<_auth_action_init_post_data,data>>
|**errors** +
__optional__|Errors as result of previous transaction request. See Pluggable authenticators documentation for more information.|<<_auth_action_init_post_errors,errors>>
|**fingerprint** +
__optional__|Descriptor fingerprint. Should not change between requests. If it changes fresh descriptor has to be feched and action has to be restarted.|string
|**mainAction** +
__optional__|Transaction's main action, used to display confirmation screen.|string
|**next** +
__optional__|Array of next actions that have to be executed.|< string > array
|**nextTransaction** +
__optional__|Next action that has to be executed.|< string > array
|**redirectedTo** +
__optional__|Next action that has to be executed.|string
|**status** +
__required__|Current transaction status.|enum (ready, pending, done, bad_request, rejected, redirected, failed, closed, pendingElsewhere, handedOver, doneInAnotherTab, cancelled, expired)
|**transactionId** +
__required__|Transaction id.|string
|**type** +
__required__|Transaction type.|enum (simple, registration, registrationSaa, authentication, activation, account, action, consent, operation, administration, reclaimProfile)
|===

[[_auth_action_init_post_data]]
**data**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**$authenticator id** +
__optional__|Action specific data.|object
|===

[[_auth_action_init_post_errors]]
**errors**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**$authenticator id** +
__optional__|Action specific errors.|object
|===

[[_auth_action_init_post_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request, descriptorId_missing_parameter, profileId_missing_parameter, actionId_missing_parameter, profile_does_not_exist, msisdn_not_exists, invalid_request)
|**error_description** +
__optional__|string
|===


==== Consumes

* `application/json`


==== Produces

* `application/json`


==== Tags

* Descriptor


==== Example HTTP response

===== Response 200
[source,json]
----
{
  "status" : "ready",
  "type" : "authentication",
  "transactionId" : "78239500-74d6-11e6-acf7-1129c898746a",
  "next" : [ "sms::authenticateLoa2" ],
  "fingerprint" : "ba43edd03c1265f045a943aee8758d21",
  "mainAction" : "sms::authenticateLoa2",
  "correlationId" : "123456"
}
----


===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_auth_activation_init_post]]
=== Init transaction: activation
....
POST /auth/activation/init
....


==== Description
Endpoint for initializing profile activation transaction. NOTE: One of parameters 'msisdn', 'pcr' or 'encryptedMsisdn' is required in request.


==== Parameters

[options="header", cols=".^2a,.^3a,.^4a"]
|===
|Type|Name|Schema
|**Body**|**Activation init request body** +
__optional__|<<_auth_activation_init_post_activation_init_request_body,Activation init request body>>
|===

[[_auth_activation_init_post_activation_init_request_body]]
**Activation init request body**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**correlationId** +
__optional__|Correlation id used by frontend app to correlate transaction responses.|string
|**encryptedMsisdn** +
__optional__|User's encrypted MSISDN. Hexadecimally encoded binary string. Required if 'msisdn' or 'pcr' is not provided.|string
|**forClientId** +
__required__|Client id of application that requested authorization.|string
|**instanceIdentifier** +
__optional__|Identifier browser instance maing the request.|string
|**internalEncrMsisdn** +
__optional__|Defines if MSISDN was encrypted with internal key instead of MNO's.|boolean
|**loa** +
__required__|Requested Level of Assurance.|enum (loa2, loa3)
|**msisdn** +
__required__|User's phone number. Required if 'pcr' or 'encryptedMsisdn' is not provided.|string
|**pcr** +
__optional__|User's PCR. Required if 'msisdn' or 'encryptedMsisdn' is not provided.|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|Status response|<<_auth_activation_init_post_response_200,Response 200>>
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_auth_activation_init_post_response_400,Response 400>>
|**401**|Access denied. Authentication is required to proceed processing request.|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_auth_activation_init_post_response_200]]
**Response 200**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**clientName** +
__optional__|Requesting SP's name returned when status is final and the transaction is of type authentication or authorization|string
|**correlationId** +
__optional__|CorrelationId sent on init request.|string
|**data** +
__optional__|Data that has to be used with next action - values of form fields, display fields and placeholders. See Pluggable authenticators documentation for more information.|<<_auth_activation_init_post_data,data>>
|**errors** +
__optional__|Errors as result of previous transaction request. See Pluggable authenticators documentation for more information.|<<_auth_activation_init_post_errors,errors>>
|**fingerprint** +
__optional__|Descriptor fingerprint. Should not change between requests. If it changes fresh descriptor has to be feched and action has to be restarted.|string
|**mainAction** +
__optional__|Transaction's main action, used to display confirmation screen.|string
|**next** +
__optional__|Array of next actions that have to be executed.|< string > array
|**nextTransaction** +
__optional__|Next action that has to be executed.|< string > array
|**redirectedTo** +
__optional__|Next action that has to be executed.|string
|**status** +
__required__|Current transaction status.|enum (ready, pending, done, bad_request, rejected, redirected, failed, closed, pendingElsewhere, handedOver, doneInAnotherTab, cancelled, expired)
|**transactionId** +
__required__|Transaction id.|string
|**type** +
__required__|Transaction type.|enum (simple, registration, registrationSaa, authentication, activation, account, action, consent, operation, administration, reclaimProfile)
|===

[[_auth_activation_init_post_data]]
**data**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**$authenticator id** +
__optional__|Action specific data.|object
|===

[[_auth_activation_init_post_errors]]
**errors**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**$authenticator id** +
__optional__|Action specific errors.|object
|===

[[_auth_activation_init_post_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request, msisdn_missing_parameter, msisdn_not_exists, forClientId_missing_parameter, loa_missing_parameter, invalid_loa, invalid_request)
|**error_description** +
__optional__|string
|===


==== Consumes

* `application/json`


==== Produces

* `application/json`


==== Tags

* Descriptor


==== Example HTTP response

===== Response 200
[source,json]
----
{
  "status" : "ready",
  "type" : "activation",
  "transactionId" : "ffbd0c80-7c23-4756-99c6-b27d8b970c35",
  "next" : [ "mdi::activation_activationValidate" ],
  "fingerprint" : "768929be5a6f5bae51706e58fe9b53f2"
}
----


===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_auth_administration_init_post]]
=== Init transaction: administration
....
POST /auth/administration/init
....


==== Description
Endpoint for initializing administration action. Only System apps can perform this type of action.


==== Parameters

[options="header", cols=".^2a,.^3a,.^4a"]
|===
|Type|Name|Schema
|**Body**|**Authentication init request body** +
__optional__|<<_auth_administration_init_post_authentication_init_request_body,Authentication init request body>>
|===

[[_auth_administration_init_post_authentication_init_request_body]]
**Authentication init request body**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**actionId** +
__required__|Action id.|string
|**descriptorId** +
__required__|Descriptor id.|string
|**operator** +
__required__|MNO id.|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|Status response|<<_auth_administration_init_post_response_200,Response 200>>
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_auth_administration_init_post_response_400,Response 400>>
|**401**|Access denied. Authentication is required to proceed processing request.|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_auth_administration_init_post_response_200]]
**Response 200**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**clientName** +
__optional__|Requesting SP's name returned when status is final and the transaction is of type authentication or authorization|string
|**correlationId** +
__optional__|CorrelationId sent on init request.|string
|**data** +
__optional__|Data that has to be used with next action - values of form fields, display fields and placeholders. See Pluggable authenticators documentation for more information.|<<_auth_administration_init_post_data,data>>
|**errors** +
__optional__|Errors as result of previous transaction request. See Pluggable authenticators documentation for more information.|<<_auth_administration_init_post_errors,errors>>
|**fingerprint** +
__optional__|Descriptor fingerprint. Should not change between requests. If it changes fresh descriptor has to be feched and action has to be restarted.|string
|**mainAction** +
__optional__|Transaction's main action, used to display confirmation screen.|string
|**next** +
__optional__|Array of next actions that have to be executed.|< string > array
|**nextTransaction** +
__optional__|Next action that has to be executed.|< string > array
|**redirectedTo** +
__optional__|Next action that has to be executed.|string
|**status** +
__required__|Current transaction status.|enum (ready, pending, done, bad_request, rejected, redirected, failed, closed, pendingElsewhere, handedOver, doneInAnotherTab, cancelled, expired)
|**transactionId** +
__required__|Transaction id.|string
|**type** +
__required__|Transaction type.|enum (simple, registration, registrationSaa, authentication, activation, account, action, consent, operation, administration, reclaimProfile)
|===

[[_auth_administration_init_post_data]]
**data**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**$authenticator id** +
__optional__|Action specific data.|object
|===

[[_auth_administration_init_post_errors]]
**errors**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**$authenticator id** +
__optional__|Action specific errors.|object
|===

[[_auth_administration_init_post_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request, operator_missing_parameter, descriptorId_missing_parameter, actionId_missing_parameter, operator_could_not_be_found, invalid_request, access_denied)
|**error_description** +
__optional__|string
|===


==== Consumes

* `application/json`


==== Produces

* `application/json`


==== Tags

* Descriptor


==== Example HTTP response

===== Response 200
[source,json]
----
{
  "status" : "ready",
  "type" : "authentication",
  "transactionId" : "78239500-74d6-11e6-acf7-1129c898746a",
  "next" : [ "sms::authenticateLoa2" ],
  "fingerprint" : "ba43edd03c1265f045a943aee8758d21",
  "mainAction" : "sms::authenticateLoa2",
  "correlationId" : "123456"
}
----


===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_auth_administration_init_link_post]]
=== NewTransactionLink - Administration
....
POST /auth/administration/init/link
....


==== Description
Endpoint for generating administration init public code.


==== Parameters

[options="header", cols=".^2a,.^3a,.^4a"]
|===
|Type|Name|Schema
|**Body**|**Authentication init request body** +
__optional__|<<_auth_administration_init_link_post_authentication_init_request_body,Authentication init request body>>
|===

[[_auth_administration_init_link_post_authentication_init_request_body]]
**Authentication init request body**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**actionId** +
__required__|Action id.|string
|**descriptorId** +
__required__|Descriptor id.|string
|**operator** +
__required__|MNO id.|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|Status response|<<_auth_administration_init_link_post_response_200,Response 200>>
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_auth_administration_init_link_post_response_400,Response 400>>
|**401**|Access denied. Authentication is required to proceed processing request.|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_auth_administration_init_link_post_response_200]]
**Response 200**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**code** +
__required__|Code related to data stored in storage.|string
|**publicUrl** +
__required__|Public URL where new transaction is initiated from stored data.|string
|===

[[_auth_administration_init_link_post_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request, operator_missing_parameter, descriptorId_missing_parameter, actionId_missing_parameter, operator_could_not_be_found, invalid_request, access_denied)
|**error_description** +
__optional__|string
|===


==== Consumes

* `application/json`


==== Produces

* `application/json`


==== Tags

* Descriptor


==== Example HTTP response

===== Response 200
[source,json]
----
{
  "code" : "8c7067d9d774659e",
  "publicUrl" : "http://weblogin.etalio.dev/d/8c7067d9d774659e"
}
----


===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_auth_ape_post]]
=== Authentication Policy Engine (APE)
....
POST /auth/ape
....


==== Description
Endpoint for retrieving lists of preferred and available authentication methods.


==== Parameters

[options="header", cols=".^2a,.^3a,.^4a"]
|===
|Type|Name|Schema
|**Body**|**APE request body** +
__optional__|<<_auth_ape_post_ape_request_body,APE request body>>
|===

[[_auth_ape_post_ape_request_body]]
**APE request body**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**encryptedMsisdn** +
__optional__|User's encrypted MSISDN. Hexadecimally encoded binary string.|string
|**forClientId** +
__optional__|Client id of application that requested authorization.|string
|**internalEncrMsisdn** +
__optional__|Defines if MSISDN was encrypted with internal key instead of MNO's.|boolean
|**msisdn** +
__optional__|User's phone number.|string
|**operator** +
__required__|MNO id. If 'msisdn' is specified it has to match MNO of user with provided msisdn.|string
|**pcr** +
__optional__|User's PCR.|string
|**type** +
__optional__|Flow type: authetnication, authorization, bcAuthentication, bcAuthorization|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|A representation of a descriptor.|< <<_auth_ape_post_response_200,Response 200>> > array
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_auth_ape_post_response_400,Response 400>>
|**401**|Access denied. Authentication is required to proceed processing request.|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_auth_ape_post_response_200]]
**Response 200**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**available** +
__required__|Available authneticators.|<<_auth_ape_post_available,available>>
|**preferred** +
__required__|Preferred authneticators.|<<_auth_ape_post_preferred,preferred>>
|===

[[_auth_ape_post_available]]
**available**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**loa2** +
__required__|Available loa2 authenticators.|< object > array
|**loa3** +
__required__|Available loa3 authenticator.|< object > array
|===

[[_auth_ape_post_preferred]]
**preferred**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**loa2** +
__required__|Preferred loa2 authenticator.|string
|**loa3** +
__required__|Preferred loa3 authenticator.|string
|===

[[_auth_ape_post_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request, invalid_operator, invalid_client_id, operator_missing_parameter, profile_does_not_exist)
|**error_description** +
__optional__|string
|===


==== Consumes

* `application/json`


==== Produces

* `application/json`


==== Tags

* Descriptor


==== Example HTTP response

===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_auth_authentication_init_post]]
=== Init transaction: authentication
....
POST /auth/authentication/init
....


==== Description
Endpoint for retrieving lists of preferred and available authentication methods. NOTE: One of parameters 'msisdn', 'pcr' or 'encryptedMsisdn' is required in request.


==== Parameters

[options="header", cols=".^2a,.^3a,.^4a"]
|===
|Type|Name|Schema
|**Body**|**Authentication init request body** +
__optional__|<<_auth_authentication_init_post_authentication_init_request_body,Authentication init request body>>
|===

[[_auth_authentication_init_post_authentication_init_request_body]]
**Authentication init request body**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**authenticators** +
__required__|Chosen authenticators to start authentication with.|< string > array
|**claims** +
__optional__|Premium information (name, address, family name,…).|object
|**correlationId** +
__optional__|Correlation id used by frontend app to correlate transaction responses.|string
|**encryptedMsisdn** +
__optional__|User's encrypted MSISDN. Hexadecimally encoded binary string. Required if 'msisdn' or 'pcr' is not provided.|string
|**forClientId** +
__required__|Client id of application that requested authorization.|string
|**instanceIdentifier** +
__optional__|Identifier browser instance maing the request.|string
|**internalEncrMsisdn** +
__optional__|Defines if MSISDN was encrypted with internal key instead of MNO's.|boolean
|**loa** +
__required__|Requested Level of Assurance.|enum (loa2, loa3)
|**locales** +
__optional__|Selected language.|string
|**mcVersion** +
__optional__|Mobile connect version.|enum (mc_v1.1, mc_v1.2)
|**msisdn** +
__required__|User's phone number. Required if 'pcr' or 'encryptedMsisdn' is not provided.|string
|**operator** +
__required__|MNO id. It has to match MNO of user with provided msisdn/pcr.|string
|**pcr** +
__optional__|User's PCR. Required if 'msisdn' or 'encryptedMsisdn' is not provided.|string
|**scopes** +
__required__|Scopes that application requires|< string > array
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|Authentication init response body|<<_auth_authentication_init_post_response_200,Response 200>>
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_auth_authentication_init_post_response_400,Response 400>>
|**401**|Access denied. Authentication is required to proceed processing request.|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_auth_authentication_init_post_response_200]]
**Response 200**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**clientName** +
__optional__|Requesting SP's name returned when status is final and the transaction is of type authentication or authorization|string
|**correlationId** +
__optional__|CorrelationId sent on init request.|string
|**data** +
__optional__|Data that has to be used with next action - values of form fields, display fields and placeholders. See Pluggable authenticators documentation for more information.|<<_auth_authentication_init_post_data,data>>
|**errors** +
__optional__|Errors as result of previous transaction request. See Pluggable authenticators documentation for more information.|<<_auth_authentication_init_post_errors,errors>>
|**fingerprint** +
__optional__|Descriptor fingerprint. Should not change between requests. If it changes fresh descriptor has to be feched and action has to be restarted.|string
|**mainAction** +
__optional__|Transaction's main action, used to display confirmation screen.|string
|**next** +
__optional__|Array of next actions that have to be executed.|< string > array
|**nextTransaction** +
__optional__|Next action that has to be executed.|< string > array
|**redirectedTo** +
__optional__|Next action that has to be executed.|string
|**status** +
__required__|Current transaction status.|enum (ready, pending, done, bad_request, rejected, redirected, failed, closed, pendingElsewhere, handedOver, doneInAnotherTab, cancelled, expired)
|**transactionId** +
__required__|Transaction id.|string
|**type** +
__required__|Transaction type.|enum (simple, registration, registrationSaa, authentication, activation, account, action, consent, operation, administration, reclaimProfile)
|===

[[_auth_authentication_init_post_data]]
**data**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**$authenticator id** +
__optional__|Action specific data.|object
|===

[[_auth_authentication_init_post_errors]]
**errors**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**$authenticator id** +
__optional__|Action specific errors.|object
|===

[[_auth_authentication_init_post_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request, msisdn_missing_parameter, operator_missing_parameter, loa_missing_parameter, authenticators_missing_parameter, forClientId_missing_parameter, msisdn_not_exists, invalid_operator, invalid_loa, invalid_request)
|**error_description** +
__optional__|string
|===


==== Consumes

* `application/json`


==== Produces

* `application/json`


==== Tags

* Descriptor


==== Example HTTP response

===== Response 200
[source,json]
----
{
  "status" : "ready",
  "type" : "authentication",
  "transactionId" : "78239500-74d6-11e6-acf7-1129c898746a",
  "next" : [ "sms::authenticateLoa2" ],
  "fingerprint" : "ba43edd03c1265f045a943aee8758d21",
  "mainAction" : "sms::authenticateLoa2",
  "correlationId" : "123456"
}
----


===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_auth_authentication_init_link_post]]
=== NewTransactionLink - Authentication
....
POST /auth/authentication/init/link
....


==== Description
Endpoint for generating authentication init public code. NOTE: One of parameters 'msisdn', 'pcr' or 'encryptedMsisdn' is required in request.


==== Parameters

[options="header", cols=".^2a,.^3a,.^4a"]
|===
|Type|Name|Schema
|**Body**|**Authentication init request body** +
__optional__|<<_auth_authentication_init_link_post_authentication_init_request_body,Authentication init request body>>
|===

[[_auth_authentication_init_link_post_authentication_init_request_body]]
**Authentication init request body**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**authenticators** +
__required__|Chosen authenticators to start authentication with.|< string > array
|**claims** +
__optional__|Premium information (name, address, family name,…).|object
|**correlationId** +
__optional__|Correlation id used by frontend app to correlate transaction responses.|string
|**encryptedMsisdn** +
__optional__|User's encrypted MSISDN. Hexadecimally encoded binary string. Required if 'msisdn' or 'pcr' is not provided.|string
|**forClientId** +
__required__|Client id of application that requested authorization.|string
|**instanceIdentifier** +
__optional__|Identifier browser instance maing the request.|string
|**internalEncrMsisdn** +
__optional__|Defines if MSISDN was encrypted with internal key instead of MNO's.|boolean
|**loa** +
__required__|Requested Level of Assurance.|enum (loa2, loa3)
|**locales** +
__optional__|Selected language.|string
|**mcVersion** +
__optional__|Mobile connect version.|enum (mc_v1.1, mc_v1.2)
|**msisdn** +
__required__|User's phone number. Required if 'pcr' or 'encryptedMsisdn' is not provided.|string
|**operator** +
__required__|MNO id. It has to match MNO of user with provided msisdn/pcr.|string
|**pcr** +
__optional__|User's PCR. Required if 'msisdn' or 'encryptedMsisdn' is not provided.|string
|**scopes** +
__required__|Scopes that application requires|< string > array
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|Status response|<<_auth_authentication_init_link_post_response_200,Response 200>>
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_auth_authentication_init_link_post_response_400,Response 400>>
|**401**|Access denied. Authentication is required to proceed processing request.|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_auth_authentication_init_link_post_response_200]]
**Response 200**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**code** +
__required__|Code related to data stored in storage.|string
|**publicUrl** +
__required__|Public URL where new transaction is initiated from stored data.|string
|===

[[_auth_authentication_init_link_post_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request, msisdn_missing_parameter, operator_missing_parameter, loa_missing_parameter, authenticators_missing_parameter, forClientId_missing_parameter, msisdn_not_exists, invalid_operator, invalid_loa, invalid_request)
|**error_description** +
__optional__|string
|===


==== Consumes

* `application/json`


==== Produces

* `application/json`


==== Tags

* Descriptor


==== Example HTTP response

===== Response 200
[source,json]
----
{
  "code" : "8c7067d9d774659e",
  "publicUrl" : "http://weblogin.etalio.dev/d/8c7067d9d774659e"
}
----


===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_auth_authorization_init_link_post]]
=== NewTransactionLink - Authorization
....
POST /auth/authorization/init/link
....


==== Description
Endpoint for generating authorization init public code. NOTE: One of parameters 'msisdn', 'pcr' or 'encryptedMsisdn' is required in request.


==== Parameters

[options="header", cols=".^2a,.^3a,.^4a"]
|===
|Type|Name|Schema
|**Body**|**Authentication init request body** +
__optional__|<<_auth_authorization_init_link_post_authentication_init_request_body,Authentication init request body>>
|===

[[_auth_authorization_init_link_post_authentication_init_request_body]]
**Authentication init request body**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**authenticators** +
__required__|Chosen authenticators to start authentication with.|< string > array
|**bindingMessage** +
__required__|Client provided plain text, 'reference or ID' to interlock consumption device and authorization device for a better user experience. The message will be displayed on consumption device and mobile device.|string
|**claims** +
__optional__|Premium information (name, address, family name).|object
|**clientName** +
__required__|Client name returned on 'doneInAnotherTab' status.|string
|**context** +
__required__|A Transaction/action based message displayed on authorization device and must provide by RP/client.|string
|**correlationId** +
__optional__|Correlation id used by frontend app to correlate transaction responses.|string
|**encryptedMsisdn** +
__optional__|User's encrypted MSISDN. Hexadecimally encoded binary string. Required if 'msisdn' or 'pcr' is not provided.|string
|**forClientId** +
__required__|Client id of application that requested authorization.|string
|**instanceIdentifier** +
__optional__|Identifier browser instance maing the request.|string
|**internalEncrMsisdn** +
__optional__|Defines if MSISDN was encrypted with internal key instead of MNO's.|boolean
|**loa** +
__required__|Requested Level of Assurance.|enum (loa2, loa3)
|**locales** +
__optional__|Selected language.|string
|**mcVersion** +
__required__|Mobile connect version.|enum (mc_v1.1, mc_v1.2)
|**msisdn** +
__optional__|User's phone number. Required if 'pcr' or 'encryptedMsisdn' is not provided.|string
|**operator** +
__required__|MNO id. It has to match MNO of user with provided msisdn/pcr.|string
|**pcr** +
__optional__|User's PCR. Required if 'msisdn' or 'encryptedMsisdn' is not provided.|string
|**scopes** +
__required__|Scopes that application requires|< string > array
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|Status response|<<_auth_authorization_init_link_post_response_200,Response 200>>
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_auth_authorization_init_link_post_response_400,Response 400>>
|**401**|Access denied. Authentication is required to proceed processing request.|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_auth_authorization_init_link_post_response_200]]
**Response 200**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**code** +
__required__|Code related to data stored in storage.|string
|**publicUrl** +
__required__|Public URL where new transaction is initiated from stored data.|string
|===

[[_auth_authorization_init_link_post_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request, msisdn_missing_parameter, operator_missing_parameter, loa_missing_parameter, authenticators_missing_parameter, forClientId_missing_parameter, msisdn_not_exists, invalid_operator, invalid_loa, invalid_request)
|**error_description** +
__optional__|string
|===


==== Consumes

* `application/json`


==== Produces

* `application/json`


==== Tags

* Descriptor


==== Example HTTP response

===== Response 200
[source,json]
----
{
  "code" : "8c7067d9d774659e",
  "publicUrl" : "http://weblogin.etalio.dev/d/8c7067d9d774659e"
}
----


===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_auth_cancel_post]]
=== Cancel transaction
....
POST /auth/cancel
....


==== Description
Endpoints for canceling transaction.


==== Parameters

[options="header", cols=".^2a,.^3a,.^4a"]
|===
|Type|Name|Schema
|**Body**|**Authentication init request body** +
__optional__|<<_auth_cancel_post_authentication_init_request_body,Authentication init request body>>
|===

[[_auth_cancel_post_authentication_init_request_body]]
**Authentication init request body**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**transactionId** +
__required__|Transaction id.|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|Status response|<<_auth_cancel_post_response_200,Response 200>>
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_auth_cancel_post_response_400,Response 400>>
|**401**|Access denied. Authentication is required to proceed processing request.|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_auth_cancel_post_response_200]]
**Response 200**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**clientName** +
__optional__|Requesting SP's name returned when status is final and the transaction is of type authentication or authorization|string
|**correlationId** +
__optional__|CorrelationId sent on init request.|string
|**data** +
__optional__|Data that has to be used with next action - values of form fields, display fields and placeholders. See Pluggable authenticators documentation for more information.|<<_auth_cancel_post_data,data>>
|**errors** +
__optional__|Errors as result of previous transaction request. See Pluggable authenticators documentation for more information.|<<_auth_cancel_post_errors,errors>>
|**fingerprint** +
__optional__|Descriptor fingerprint. Should not change between requests. If it changes fresh descriptor has to be feched and action has to be restarted.|string
|**mainAction** +
__optional__|Transaction's main action, used to display confirmation screen.|string
|**next** +
__optional__|Array of next actions that have to be executed.|< string > array
|**nextTransaction** +
__optional__|Next action that has to be executed.|< string > array
|**redirectedTo** +
__optional__|Next action that has to be executed.|string
|**status** +
__required__|Current transaction status.|enum (ready, pending, done, bad_request, rejected, redirected, failed, closed, pendingElsewhere, handedOver, doneInAnotherTab, cancelled, expired)
|**transactionId** +
__required__|Transaction id.|string
|**type** +
__required__|Transaction type.|enum (simple, registration, registrationSaa, authentication, activation, account, action, consent, operation, administration, reclaimProfile)
|===

[[_auth_cancel_post_data]]
**data**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**$authenticator id** +
__optional__|Action specific data.|object
|===

[[_auth_cancel_post_errors]]
**errors**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**$authenticator id** +
__optional__|Action specific errors.|object
|===

[[_auth_cancel_post_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request, transactionId_missing_parameter, access_denied, invalid_transactionId)
|**error_description** +
__optional__|string
|===


==== Consumes

* `application/json`


==== Produces

* `application/json`


==== Tags

* Descriptor


==== Example HTTP response

===== Response 200
[source,json]
----
{
  "status" : "cancelled",
  "type" : "authentication",
  "transactionId" : "78239500-74d6-11e6-acf7-1129c898746a",
  "fingerprint" : "ba43edd03c1265f045a943aee8758d21",
  "mainAction" : "sms::authenticateLoa2",
  "correlationId" : "123456",
  "clientName" : "Boxxy"
}
----


===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_auth_confirm_put]]
=== Confirm async action
....
PUT /auth/confirm
....


==== Description
Endpoint for confirming async action. Used by authenticator. Authorization with valid JWT token is required.


==== Parameters

[options="header", cols=".^2a,.^3a,.^4a"]
|===
|Type|Name|Schema
|**Body**|**Confirm action request body** +
__optional__|<<_auth_confirm_put_confirm_action_request_body,Confirm action request body>>
|===

[[_auth_confirm_put_confirm_action_request_body]]
**Confirm action request body**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**authenticationData** +
__optional__|Succesful authentication data.|object
|**authenticatorData** +
__optional__|authenticator data. Will be stored to users profile (linking data).|object
|**data** +
__optional__|Frontend action data.|object
|**errors** +
__optional__|Frontend action errors.|object
|**internalError** +
__optional__|Authneticator error.|string
|**nextAction** +
__optional__|Next action.|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|Confirm action response body|<<_auth_confirm_put_response_200,Response 200>>
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_auth_confirm_put_response_400,Response 400>>
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_auth_confirm_put_response_200]]
**Response 200**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**status** +
__optional__|Request status.|string
|===

[[_auth_confirm_put_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request)
|**error_description** +
__optional__|string
|===


==== Consumes

* `application/json`


==== Produces

* `application/json`


==== Tags

* Descriptor


==== Example HTTP response

===== Response 200
[source,json]
----
{
  "status" : "OK"
}
----


===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_auth_confirm_delete]]
=== Cancel async action
....
DELETE /auth/confirm
....


==== Description
Endpoint for canceling async action. Used by authenticator. Authorization with valid JWT token is required.


==== Parameters

[options="header", cols=".^2a,.^3a,.^4a"]
|===
|Type|Name|Schema
|**Body**|**Notification request body** +
__optional__|<<_auth_confirm_delete_notification_request_body,Notification request body>>
|===

[[_auth_confirm_delete_notification_request_body]]
**Notification request body**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**authenticationData** +
__optional__|Succesful authentication data.|object
|**authenticatorData** +
__optional__|authenticator data. Will be stored to users profile (linking data).|object
|**data** +
__optional__|Frontend action data.|object
|**errors** +
__optional__|Frontend action errors.|object
|**internalError** +
__optional__|Authneticator error.|string
|**nextAction** +
__optional__|Next action.|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|Notification response body|<<_auth_confirm_delete_response_200,Response 200>>
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_auth_confirm_delete_response_400,Response 400>>
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_auth_confirm_delete_response_200]]
**Response 200**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**status** +
__optional__|Request status.|string
|===

[[_auth_confirm_delete_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request)
|**error_description** +
__optional__|string
|===


==== Consumes

* `application/json`


==== Produces

* `application/json`


==== Tags

* Descriptor


==== Example HTTP response

===== Response 200
[source,json]
----
{
  "status" : "OK"
}
----


===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_auth_confirmpubliccode_post]]
=== Confirm public code
....
POST /auth/confirmPublicCode
....


==== Description
Endpoint for confirming public code.


==== Parameters

[options="header", cols=".^2a,.^3a,.^4a"]
|===
|Type|Name|Schema
|**Body**|**Authentication init request body** +
__optional__|<<_auth_confirmpubliccode_post_authentication_init_request_body,Authentication init request body>>
|===

[[_auth_confirmpubliccode_post_authentication_init_request_body]]
**Authentication init request body**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**code** +
__required__|Public code.|string
|**instanceIdentifier** +
__optional__|Identifier browser instance maing the request.|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|Status response|<<_auth_confirmpubliccode_post_response_200,Response 200>>
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_auth_confirmpubliccode_post_response_400,Response 400>>
|**401**|Access denied. Authentication is required to proceed processing request.|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_auth_confirmpubliccode_post_response_200]]
**Response 200**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**data** +
__optional__|Data that has to be used with next action - values of form fields, display fields and placeholders. See Pluggable authenticators documentation for more information.|<<_auth_confirmpubliccode_post_data,data>>
|**errors** +
__optional__|Errors as result of previous transaction request. See Pluggable authenticators documentation for more information.|<<_auth_confirmpubliccode_post_errors,errors>>
|**fingerprint** +
__optional__|Descriptor fingerprint. Should not change between requests. If it changes fresh descriptor has to be feched and action has to be restarted.|string
|**next** +
__optional__|Array of next actions that have to be executed.|< string > array
|**operator** +
__optional__|Operator id.|string
|**redirectedTo** +
__optional__|Next action that has to be executed.|string
|**status** +
__required__|Current transaction status.|enum (OK, ready, pending, done, bad_request, redirected, failed, closed, pendingElsewhere, handedOver, expired)
|**transactionId** +
__optional__|Transaction id.|string
|**type** +
__optional__|Transaction type.|enum (simple, registration, authentication, account, action, operation, administration, reclaimProfile)
|===

[[_auth_confirmpubliccode_post_data]]
**data**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**$authenticator id** +
__optional__|Action specific data.|object
|===

[[_auth_confirmpubliccode_post_errors]]
**errors**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**$authenticator id** +
__optional__|Action specific errors.|object
|===

[[_auth_confirmpubliccode_post_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request, code_missing_parameter, invalid_code, access_denied)
|**error_description** +
__optional__|string
|===


==== Consumes

* `application/json`


==== Produces

* `application/json`


==== Tags

* Descriptor


==== Example HTTP response

===== Response 200
[source,json]
----
{
  "status" : "ready",
  "type" : "authentication",
  "transactionId" : "78239500-74d6-11e6-acf7-1129c898746a",
  "next" : [ "sms::authenticateLoa2" ],
  "fingerprint" : "ba43edd03c1265f045a943aee8758d21",
  "operator" : "b2d70644-8e4d-4b6d-a803-a5bd35833163"
}
----


===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_auth_confirmtransactionlink_post]]
=== Confirm transactionLink public code
....
POST /auth/confirmTransactionLink
....


==== Description
Endpoint for confirming transactionLink public code. Endpoint will initiate new transaction and store it. Response will contain new transaction id, status and related actions for its type.


==== Parameters

[options="header", cols=".^2a,.^3a,.^4a"]
|===
|Type|Name|Schema
|**Body**|**Authentication init request body** +
__optional__|<<_auth_confirmtransactionlink_post_authentication_init_request_body,Authentication init request body>>
|===

[[_auth_confirmtransactionlink_post_authentication_init_request_body]]
**Authentication init request body**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**code** +
__required__|Public code.|string
|**instanceIdentifier** +
__optional__|Identifier browser instance maing the request.|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|Status response|<<_auth_confirmtransactionlink_post_response_200,Response 200>>
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_auth_confirmtransactionlink_post_response_400,Response 400>>
|**401**|Access denied. Authentication is required to proceed processing request.|No Content
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_auth_confirmtransactionlink_post_response_200]]
**Response 200**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**data** +
__optional__|Data that has to be used with next action - values of form fields, display fields and placeholders. See Pluggable authenticators documentation for more information.|<<_auth_confirmtransactionlink_post_data,data>>
|**errors** +
__optional__|Errors as result of previous transaction request. See Pluggable authenticators documentation for more information.|<<_auth_confirmtransactionlink_post_errors,errors>>
|**fingerprint** +
__optional__|Descriptor fingerprint. Should not change between requests. If it changes fresh descriptor has to be feched and action has to be restarted.|string
|**next** +
__optional__|Array of next actions that have to be executed.|< string > array
|**operator** +
__optional__|Operator id.|string
|**redirectedTo** +
__optional__|Next action that has to be executed.|string
|**status** +
__required__|Current transaction status.|enum (OK, ready, pending, done, bad_request, redirected, failed, closed, pendingElsewhere, handedOver, expired)
|**transactionId** +
__optional__|Transaction id.|string
|**type** +
__optional__|Transaction type.|enum (simple, registration, authentication, account, action, operation, administration, reclaimProfile)
|===

[[_auth_confirmtransactionlink_post_data]]
**data**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**$authenticator id** +
__optional__|Action specific data.|object
|===

[[_auth_confirmtransactionlink_post_errors]]
**errors**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**$authenticator id** +
__optional__|Action specific errors.|object
|===

[[_auth_confirmtransactionlink_post_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request, code_missing_parameter, invalid_code, access_denied)
|**error_description** +
__optional__|string
|===


==== Consumes

* `application/json`


==== Produces

* `application/json`


==== Tags

* Descriptor


==== Example HTTP response

===== Response 200
[source,json]
----
{
  "status" : "ready",
  "type" : "authentication",
  "transactionId" : "78239500-74d6-11e6-acf7-1129c898746a",
  "next" : [ "sms::authenticateLoa2" ],
  "fingerprint" : "ba43edd03c1265f045a943aee8758d21",
  "operator" : "b2d70644-8e4d-4b6d-a803-a5bd35833163"
}
----


===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_auth_consent_init_post]]
=== Init transaction: consent
....
POST /auth/consent/init
....


==== Description
Endpoint for initializing profile consent transaction. NOTE: One of parameters 'msisdn', 'pcr' or 'encryptedMsisdn' is required in request.


==== Parameters

[options="header", cols=".^2a,.^3a,.^4a"]
|===
|Type|Name|Schema
|**Body**|**Consent init request body** +
__optional__|<<_auth_consent_init_post_consent_init_request_body,Consent init request body>>
|===

[[_auth_consent_init_post_consent_init_request_body]]
**Consent init request body**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**authenticators** +
__required__|Chosen authenticators to start authentication with.|< string > array
|**clientName** +
__required__|Client name returned on 'doneInAnotherTab' status.|string
|**correlationId** +
__optional__|Correlation id used by frontend app to correlate transaction responses.|string
|**encryptedMsisdn** +
__optional__|User's encrypted MSISDN. Hexadecimally encoded binary string. Required if 'msisdn' or 'pcr' is not provided.|string
|**forClientId** +
__required__|Client id of application that requested authorization.|string
|**instanceIdentifier** +
__optional__|Identifier browser instance maing the request.|string
|**internalEncrMsisdn** +
__optional__|Defines if MSISDN was encrypted with internal key instead of MNO's.|boolean
|**msisdn** +
__required__|User's phone number. Required if 'pcr' or 'encryptedMsisdn' is not provided.|string
|**operator** +
__required__|MNO id. It has to match MNO of user with provided msisdn/pcr.|string
|**pcr** +
__optional__|User's PCR. Required if 'msisdn' or 'encryptedMsisdn' is not provided.|string
|**scopes** +
__required__|Scopes that application requires|< string > array
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|Status response|<<_auth_consent_init_post_response_200,Response 200>>
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_auth_consent_init_post_response_400,Response 400>>
|**401**|Access denied. Authentication is required to proceed processing request.|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_auth_consent_init_post_response_200]]
**Response 200**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**clientName** +
__optional__|Requesting SP's name returned when status is final and the transaction is of type authentication or authorization|string
|**correlationId** +
__optional__|CorrelationId sent on init request.|string
|**data** +
__optional__|Data that has to be used with next action - values of form fields, display fields and placeholders. See Pluggable authenticators documentation for more information.|<<_auth_consent_init_post_data,data>>
|**errors** +
__optional__|Errors as result of previous transaction request. See Pluggable authenticators documentation for more information.|<<_auth_consent_init_post_errors,errors>>
|**fingerprint** +
__optional__|Descriptor fingerprint. Should not change between requests. If it changes fresh descriptor has to be feched and action has to be restarted.|string
|**mainAction** +
__optional__|Transaction's main action, used to display confirmation screen.|string
|**next** +
__optional__|Array of next actions that have to be executed.|< string > array
|**nextTransaction** +
__optional__|Next action that has to be executed.|< string > array
|**redirectedTo** +
__optional__|Next action that has to be executed.|string
|**status** +
__required__|Current transaction status.|enum (ready, pending, done, bad_request, rejected, redirected, failed, closed, pendingElsewhere, handedOver, doneInAnotherTab, cancelled, expired)
|**transactionId** +
__required__|Transaction id.|string
|**type** +
__required__|Transaction type.|enum (simple, registration, registrationSaa, authentication, activation, account, action, consent, operation, administration, reclaimProfile)
|===

[[_auth_consent_init_post_data]]
**data**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**$authenticator id** +
__optional__|Action specific data.|object
|===

[[_auth_consent_init_post_errors]]
**errors**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**$authenticator id** +
__optional__|Action specific errors.|object
|===

[[_auth_consent_init_post_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request, msisdn_missing_parameter, operator_missing_parameter, authenticators_missing_parameter, forClientId_missing_parameter, msisdn_not_exists, invalid_operator, invalid_request)
|**error_description** +
__optional__|string
|===


==== Consumes

* `application/json`


==== Produces

* `application/json`


==== Tags

* Descriptor


==== Example HTTP response

===== Response 200
[source,json]
----
{
  "status" : "pending",
  "type" : "consent",
  "transactionId" : "fbbd3453-0f95-4f1f-b71b-046de26792b4",
  "next" : [ "authenticator::actionName" ],
  "data" : {
    "authenticator" : {
      "clientName" : "ConsentApp"
    }
  },
  "fingerprint" : "75ebd2945694da5cbf16d242a9c479"
}
----


===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_auth_describe_netopid_get]]
=== Get descriptor
....
GET /auth/describe/{netOpId}
....


==== Description
Endpoint for getting full descriptor for MNO. It returns array of all descriptors available for MNO. If MNO doesn't have any authenticators added only internal 'mdi' descriptor is returned.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Path**|**netOpId** +
__required__|Mobile network operator id|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|A representation of a descriptor.|< <<_auth_describe_netopid_get_response_200,Response 200>> > array
|**401**|Access denied. Authentication is required to proceed processing request.|No Content
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**404**|Not Found. The requested resource could not be found but may be available again in the future. Subsequent requests by the client are permissible.|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_auth_describe_netopid_get_response_200]]
**Response 200**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**account** +
__optional__|List of actions for acccount management.|< <<_auth_describe_netopid_get_account,account>> > array
|**actions** +
__optional__|All actions.|<<_auth_describe_netopid_get_actions,actions>>
|**administration** +
__optional__|List of actions with coresponding labels|< <<_auth_describe_netopid_get_administration,administration>> > array
|**autolinking** +
__optional__||string
|**description** +
__optional__||string
|**id** +
__optional__||string
|**msisdnValidation** +
__optional__||string
|**msisdnValidationOtp** +
__optional__||string
|**name** +
__optional__||string
|**operation** +
__optional__|List of actions with coresponding labels|< <<_auth_describe_netopid_get_operation,operation>> > array
|**reclaimProfile** +
__optional__||string
|**userNotification** +
__optional__||string
|===

[[_auth_describe_netopid_get_account]]
**account**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**action** +
__required__|ID of an action.|string
|**label** +
__required__|Label.|string
|**when** +
__optional__|Indicates when action should be available to user.|enum (enabled, disabled)
|===

[[_auth_describe_netopid_get_actions]]
**actions**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**$action_object** +
__optional__|Structure of actions is presented in documentation for Pluggable authenticators.|object
|===

[[_auth_describe_netopid_get_administration]]
**administration**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**action** +
__required__|ID of an action|string
|**label** +
__required__|Localized label|string
|===

[[_auth_describe_netopid_get_operation]]
**operation**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**action** +
__required__|ID of an action|string
|**label** +
__required__|Localized label|string
|===


==== Produces

* `application/json`


==== Tags

* Descriptor


==== Example HTTP response

===== Response 200
[source,json]
----
[ {
  "id" : "mdi",
  "name" : "MDI Internal definition",
  "description" : "MDI Internal definition",
  "actions" : {
    "registration_loa2" : {
      "type" : "form",
      "fields" : [ {
        "name" : "name",
        "type" : "textInput",
        "label" : "Full name",
        "example" : "John Doe",
        "required" : true
      }, {
        "name" : "email",
        "type" : "emailInput",
        "label" : "E-mail",
        "example" : "john.doe@mobileconnect.io",
        "required" : true
      }, {
        "name" : "terms",
        "type" : "termsandconditions",
        "required" : true
      } ]
    },
    "account_changeName" : {
      "init" : {
        "type" : "action"
      },
      "type" : "form",
      "fields" : [ {
        "name" : "given_name",
        "type" : "textInput",
        "label" : "First name",
        "example" : "John",
        "required" : true
      }, {
        "name" : "family_name",
        "type" : "textInput",
        "label" : "Last name",
        "example" : "Doe",
        "required" : true
      }, {
        "name" : "middle_name",
        "type" : "textInput",
        "label" : "Middle name",
        "required" : false
      } ],
      "confirmationScreen" : [ {
        "name" : "nameTitle",
        "type" : "header",
        "value" : "Account updated"
      }, {
        "name" : "nameMessage",
        "type" : "text",
        "value" : "Your name was changed."
      } ]
    }
  },
  "account" : [ {
    "label" : "Change name",
    "action" : "account_changeName"
  } ]
}, {
  "id" : "sms",
  "name" : "SMS authenticator",
  "description" : "SMS authenticator",
  "msisdnValidation" : "authenticate",
  "actions" : {
    "authenticate" : {
      "type" : "action",
      "requiredData" : [ "msisdn", "publicCode" ],
      "canRedo" : true,
      "pendingScreen" : [ {
        "name" : "sms_sent_header",
        "type" : "header",
        "value" : "We have sent you a text"
      }, {
        "name" : "sms_sent_notif",
        "type" : "text",
        "value" : "Please check your mobile device and click on the link to complete the authentication. Your mobile number is {msisdn}."
      }, {
        "name" : "sms_sent_didnt",
        "type" : "text",
        "value" : "Didn't receive a text?"
      } ]
    }
  }
} ]
----


[[_auth_keys_get]]
=== Get keys
....
GET /auth/keys
....


==== Description
Endpoint for retrieving public key for decoding/validation of JWT token. Used by authenticator.


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|Keys response|<<_auth_keys_get_response_200,Response 200>>
|**401**|Access denied. Authentication is required to proceed processing request.|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_auth_keys_get_response_200]]
**Response 200**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**$timestamp** +
__optional__|Public key. $timestamp is key id.|string
|===


==== Consumes

* `application/json`


==== Produces

* `application/json`


==== Tags

* Descriptor


==== Example HTTP response

===== Response 200
[source,json]
----
{
  "1462539521" : "-----BEGIN PUBLIC KEY-----\nMIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEAqzwQRlVDXVpPGQrhpWaU\n9NnEEYNBPj8roCidXXK43JmGLhAeXFTKZGMyTowNQTLHg7FnQ+MpFxisdfvuBA15\nCtRP2vuor4IdwejGofxMIV4N9SzAQdGxXXjeo4Tcnmk1dGU0ApKhM3kVo6bZpQRF\n8EKSgTEW5Q1kJIWmBWdo/aqcNE7MCDapSAlROVra0foVBDXp52CisnCPv5o0TcHM\nLyEbP4imHyOyRzoKc4CRMs2VnxurKO3BN13YjrMM0kIIodLIQimmFM/DKPmED1Vm\nusdd+/GRytWSiezLJlquXFOALjc/qzZzfFWJkK3UQ75zQUDBQL0vPSQhmeR5N3xk\n3mKNrTzonHegLjAl4zckroz103RpH5mbx57FNvVDWAJ/gSCMFK7jdkKhzbLKMIcD\nxotGw5pAcRq8ZyRvabr/w9j4c1HXNrk3AJS0DzpfeqFzL8gpopB10b1OhCV9zFaP\nn5e2Ok8kltXXyunoPLnWbPb3fi4R1zvr5UW6+ag/03YZGh4RLYVzazSA6H3yFfGK\nYJkvQE6Jul0hjBBEKlBCQYXPgrQZB1KzF6FDJVXW6iVzeOqxfULaBJKtAUIUjGKC\na0bqCGM1hAex4VFKZArpYcDyJ44cG5zh/oNWB/MZRGUR+vTJw0N2c66CvNRPhk6d\nGQB0ToFfnRwJO1MMP/imKYUCAwEAAQ==\n-----END PUBLIC KEY-----"
}
----


[[_auth_notification_post]]
=== Send notification
....
POST /auth/notification
....


==== Description
Endpoint for sending notification to user. Used by authenticators. Authorization with valid JWT token is required.


==== Parameters

[options="header", cols=".^2a,.^3a,.^4a"]
|===
|Type|Name|Schema
|**Body**|**Notification request body** +
__optional__|<<_auth_notification_post_notification_request_body,Notification request body>>
|===

[[_auth_notification_post_notification_request_body]]
**Notification request body**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**text** +
__required__|Notification text to be sent to user.|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|Notification response body|boolean
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_auth_notification_post_response_400,Response 400>>
|**401**|Access denied. Authentication is required to proceed processing request.|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_auth_notification_post_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request, invalid_transaction_type)
|**error_description** +
__optional__|string
|===


==== Consumes

* `application/json`


==== Produces

* `application/json`


==== Tags

* Descriptor


==== Example HTTP response

===== Response 200
[source,json]
----
true
----


===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_auth_operation_init_post]]
=== Init transaction: operation
....
POST /auth/operation/init
....


==== Description
Endpoint for initializing operation action. Only System apps can perform this type of action.


==== Parameters

[options="header", cols=".^2a,.^3a,.^4a"]
|===
|Type|Name|Schema
|**Body**|**Authentication init request body** +
__optional__|<<_auth_operation_init_post_authentication_init_request_body,Authentication init request body>>
|===

[[_auth_operation_init_post_authentication_init_request_body]]
**Authentication init request body**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**actionId** +
__required__|Action id.|string
|**descriptorId** +
__required__|Descriptor id.|string
|**operator** +
__required__|MNO id.|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|Status response|<<_auth_operation_init_post_response_200,Response 200>>
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_auth_operation_init_post_response_400,Response 400>>
|**401**|Access denied. Authentication is required to proceed processing request.|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_auth_operation_init_post_response_200]]
**Response 200**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**clientName** +
__optional__|Requesting SP's name returned when status is final and the transaction is of type authentication or authorization|string
|**correlationId** +
__optional__|CorrelationId sent on init request.|string
|**data** +
__optional__|Data that has to be used with next action - values of form fields, display fields and placeholders. See Pluggable authenticators documentation for more information.|<<_auth_operation_init_post_data,data>>
|**errors** +
__optional__|Errors as result of previous transaction request. See Pluggable authenticators documentation for more information.|<<_auth_operation_init_post_errors,errors>>
|**fingerprint** +
__optional__|Descriptor fingerprint. Should not change between requests. If it changes fresh descriptor has to be feched and action has to be restarted.|string
|**mainAction** +
__optional__|Transaction's main action, used to display confirmation screen.|string
|**next** +
__optional__|Array of next actions that have to be executed.|< string > array
|**nextTransaction** +
__optional__|Next action that has to be executed.|< string > array
|**redirectedTo** +
__optional__|Next action that has to be executed.|string
|**status** +
__required__|Current transaction status.|enum (ready, pending, done, bad_request, rejected, redirected, failed, closed, pendingElsewhere, handedOver, doneInAnotherTab, cancelled, expired)
|**transactionId** +
__required__|Transaction id.|string
|**type** +
__required__|Transaction type.|enum (simple, registration, registrationSaa, authentication, activation, account, action, consent, operation, administration, reclaimProfile)
|===

[[_auth_operation_init_post_data]]
**data**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**$authenticator id** +
__optional__|Action specific data.|object
|===

[[_auth_operation_init_post_errors]]
**errors**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**$authenticator id** +
__optional__|Action specific errors.|object
|===

[[_auth_operation_init_post_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request, operator_missing_parameter, descriptorId_missing_parameter, actionId_missing_parameter, operator_could_not_be_found, invalid_request, access_denied)
|**error_description** +
__optional__|string
|===


==== Consumes

* `application/json`


==== Produces

* `application/json`


==== Tags

* Descriptor


==== Example HTTP response

===== Response 200
[source,json]
----
{
  "status" : "ready",
  "type" : "authentication",
  "transactionId" : "78239500-74d6-11e6-acf7-1129c898746a",
  "next" : [ "sms::authenticateLoa2" ],
  "fingerprint" : "ba43edd03c1265f045a943aee8758d21",
  "mainAction" : "sms::authenticateLoa2",
  "correlationId" : "123456"
}
----


===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_auth_operation_init_link_post]]
=== NewTransactionLink - Operation
....
POST /auth/operation/init/link
....


==== Description
Endpoint for generating operation init public code.


==== Parameters

[options="header", cols=".^2a,.^3a,.^4a"]
|===
|Type|Name|Schema
|**Body**|**Authentication init request body** +
__optional__|<<_auth_operation_init_link_post_authentication_init_request_body,Authentication init request body>>
|===

[[_auth_operation_init_link_post_authentication_init_request_body]]
**Authentication init request body**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**actionId** +
__required__|Action id.|string
|**descriptorId** +
__required__|Descriptor id.|string
|**operator** +
__required__|MNO id.|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|Status response|<<_auth_operation_init_link_post_response_200,Response 200>>
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_auth_operation_init_link_post_response_400,Response 400>>
|**401**|Access denied. Authentication is required to proceed processing request.|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_auth_operation_init_link_post_response_200]]
**Response 200**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**code** +
__required__|Code related to data stored in storage.|string
|**publicUrl** +
__required__|Public URL where new transaction is initiated from stored data.|string
|===

[[_auth_operation_init_link_post_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request, operator_missing_parameter, descriptorId_missing_parameter, actionId_missing_parameter, operator_could_not_be_found, invalid_request, access_denied)
|**error_description** +
__optional__|string
|===


==== Consumes

* `application/json`


==== Produces

* `application/json`


==== Tags

* Descriptor


==== Example HTTP response

===== Response 200
[source,json]
----
{
  "code" : "8c7067d9d774659e",
  "publicUrl" : "http://weblogin.etalio.dev/d/8c7067d9d774659e"
}
----


===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_auth_reclaimprofile_init_post]]
=== Init transaction: reclaimProfile
....
POST /auth/reclaimProfile/init
....


==== Description
Endpoint for initializing reclaimProfile action.


==== Parameters

[options="header", cols=".^2a,.^3a,.^4a"]
|===
|Type|Name|Schema
|**Body**|**Authentication init request body** +
__optional__|<<_auth_reclaimprofile_init_post_authentication_init_request_body,Authentication init request body>>
|===

[[_auth_reclaimprofile_init_post_authentication_init_request_body]]
**Authentication init request body**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**correlationId** +
__optional__|Correlation id used by frontend app to correlate transaction responses.|string
|**instanceIdentifier** +
__optional__|Identifier browser instance maing the request.|string
|**msisdn** +
__required__|User's phone number.|string
|**reclaimCode** +
__required__|Reclaim profile code.|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|Status response|<<_auth_reclaimprofile_init_post_response_200,Response 200>>
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_auth_reclaimprofile_init_post_response_400,Response 400>>
|**401**|Access denied. Authentication is required to proceed processing request.|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_auth_reclaimprofile_init_post_response_200]]
**Response 200**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**clientName** +
__optional__|Requesting SP's name returned when status is final and the transaction is of type authentication or authorization|string
|**correlationId** +
__optional__|CorrelationId sent on init request.|string
|**data** +
__optional__|Data that has to be used with next action - values of form fields, display fields and placeholders. See Pluggable authenticators documentation for more information.|<<_auth_reclaimprofile_init_post_data,data>>
|**errors** +
__optional__|Errors as result of previous transaction request. See Pluggable authenticators documentation for more information.|<<_auth_reclaimprofile_init_post_errors,errors>>
|**fingerprint** +
__optional__|Descriptor fingerprint. Should not change between requests. If it changes fresh descriptor has to be feched and action has to be restarted.|string
|**mainAction** +
__optional__|Transaction's main action, used to display confirmation screen.|string
|**next** +
__optional__|Array of next actions that have to be executed.|< string > array
|**nextTransaction** +
__optional__|Next action that has to be executed.|< string > array
|**redirectedTo** +
__optional__|Next action that has to be executed.|string
|**status** +
__required__|Current transaction status.|enum (ready, pending, done, bad_request, rejected, redirected, failed, closed, pendingElsewhere, handedOver, doneInAnotherTab, cancelled, expired)
|**transactionId** +
__required__|Transaction id.|string
|**type** +
__required__|Transaction type.|enum (simple, registration, registrationSaa, authentication, activation, account, action, consent, operation, administration, reclaimProfile)
|===

[[_auth_reclaimprofile_init_post_data]]
**data**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**$authenticator id** +
__optional__|Action specific data.|object
|===

[[_auth_reclaimprofile_init_post_errors]]
**errors**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**$authenticator id** +
__optional__|Action specific errors.|object
|===

[[_auth_reclaimprofile_init_post_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request, msisdn_missing_parameter, reclaimCode_missing_parameter, invalid_code, invalid_request)
|**error_description** +
__optional__|string
|===


==== Consumes

* `application/json`


==== Produces

* `application/json`


==== Tags

* Descriptor


==== Example HTTP response

===== Response 200
[source,json]
----
{
  "status" : "ready",
  "type" : "authentication",
  "transactionId" : "78239500-74d6-11e6-acf7-1129c898746a",
  "next" : [ "sms::authenticateLoa2" ],
  "fingerprint" : "ba43edd03c1265f045a943aee8758d21",
  "mainAction" : "sms::authenticateLoa2",
  "correlationId" : "123456"
}
----


===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_auth_reclaimprofile_init_link_post]]
=== NewTransactionLink - ReclaimProfile
....
POST /auth/reclaimProfile/init/link
....


==== Description
Endpoint for generating reclaimProfile init public code.


==== Parameters

[options="header", cols=".^2a,.^3a,.^4a"]
|===
|Type|Name|Schema
|**Body**|**Authentication init request body** +
__optional__|<<_auth_reclaimprofile_init_link_post_authentication_init_request_body,Authentication init request body>>
|===

[[_auth_reclaimprofile_init_link_post_authentication_init_request_body]]
**Authentication init request body**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**correlationId** +
__optional__|Correlation id used by frontend app to correlate transaction responses.|string
|**instanceIdentifier** +
__optional__|Identifier browser instance maing the request.|string
|**msisdn** +
__required__|User's phone number.|string
|**reclaimCode** +
__required__|Reclaim profile code.|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|Status response|<<_auth_reclaimprofile_init_link_post_response_200,Response 200>>
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_auth_reclaimprofile_init_link_post_response_400,Response 400>>
|**401**|Access denied. Authentication is required to proceed processing request.|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_auth_reclaimprofile_init_link_post_response_200]]
**Response 200**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**code** +
__required__|Code related to data stored in storage.|string
|**publicUrl** +
__required__|Public URL where new transaction is initiated from stored data.|string
|===

[[_auth_reclaimprofile_init_link_post_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request, msisdn_missing_parameter, reclaimCode_missing_parameter, invalid_code, invalid_request)
|**error_description** +
__optional__|string
|===


==== Consumes

* `application/json`


==== Produces

* `application/json`


==== Tags

* Descriptor


==== Example HTTP response

===== Response 200
[source,json]
----
{
  "code" : "8c7067d9d774659e",
  "publicUrl" : "http://weblogin.etalio.dev/d/8c7067d9d774659e"
}
----


===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_auth_registration_init_post]]
=== Init transaction: registration
....
POST /auth/registration/init
....


==== Description
Endpoint for initializing user registration transaction. NOTE: One of parameters 'msisdn' or 'encryptedMsisdn' is required in request.


==== Parameters

[options="header", cols=".^2a,.^3a,.^4a"]
|===
|Type|Name|Schema
|**Body**|**Authentication init request body** +
__optional__|<<_auth_registration_init_post_authentication_init_request_body,Authentication init request body>>
|===

[[_auth_registration_init_post_authentication_init_request_body]]
**Authentication init request body**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**correlationId** +
__optional__|Correlation id used by frontend app to correlate transaction responses.|string
|**encryptedMsisdn** +
__optional__|User's encrypted MSISDN. Hexadecimally encoded binary string. Required if 'msisdn' is not provided.|string
|**forClientId** +
__required__|Client id of application that requested authorization.|string
|**instanceIdentifier** +
__optional__|Identifier browser instance maing the request.|string
|**loa** +
__required__|Requested Level of Assurance.|enum (loa2, loa3)
|**msisdn** +
__required__|User's phone number.|string
|**operator** +
__required__|MNO id.|string
|**scopes** +
__optional__|Scopes that application requires|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|Status response|<<_auth_registration_init_post_response_200,Response 200>>
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_auth_registration_init_post_response_400,Response 400>>
|**401**|Access denied. Authentication is required to proceed processing request.|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_auth_registration_init_post_response_200]]
**Response 200**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**clientName** +
__optional__|Requesting SP's name returned when status is final and the transaction is of type authentication or authorization|string
|**correlationId** +
__optional__|CorrelationId sent on init request.|string
|**data** +
__optional__|Data that has to be used with next action - values of form fields, display fields and placeholders. See Pluggable authenticators documentation for more information.|<<_auth_registration_init_post_data,data>>
|**errors** +
__optional__|Errors as result of previous transaction request. See Pluggable authenticators documentation for more information.|<<_auth_registration_init_post_errors,errors>>
|**fingerprint** +
__optional__|Descriptor fingerprint. Should not change between requests. If it changes fresh descriptor has to be feched and action has to be restarted.|string
|**mainAction** +
__optional__|Transaction's main action, used to display confirmation screen.|string
|**next** +
__optional__|Array of next actions that have to be executed.|< string > array
|**nextTransaction** +
__optional__|Next action that has to be executed.|< string > array
|**redirectedTo** +
__optional__|Next action that has to be executed.|string
|**status** +
__required__|Current transaction status.|enum (ready, pending, done, bad_request, rejected, redirected, failed, closed, pendingElsewhere, handedOver, doneInAnotherTab, cancelled, expired)
|**transactionId** +
__required__|Transaction id.|string
|**type** +
__required__|Transaction type.|enum (simple, registration, registrationSaa, authentication, activation, account, action, consent, operation, administration, reclaimProfile)
|===

[[_auth_registration_init_post_data]]
**data**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**$authenticator id** +
__optional__|Action specific data.|object
|===

[[_auth_registration_init_post_errors]]
**errors**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**$authenticator id** +
__optional__|Action specific errors.|object
|===

[[_auth_registration_init_post_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request, msisdn_missing_parameter, operator_missing_parameter, loa_missing_parameter, forClientId_missing_parameter, operator_could_not_be_found, invalid_loa, invalid_request)
|**error_description** +
__optional__|string
|===


==== Consumes

* `application/json`


==== Produces

* `application/json`


==== Tags

* Descriptor


==== Example HTTP response

===== Response 200
[source,json]
----
{
  "status" : "ready",
  "type" : "authentication",
  "transactionId" : "78239500-74d6-11e6-acf7-1129c898746a",
  "next" : [ "sms::authenticateLoa2" ],
  "fingerprint" : "ba43edd03c1265f045a943aee8758d21",
  "mainAction" : "sms::authenticateLoa2",
  "correlationId" : "123456"
}
----


===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_auth_registration_init_link_post]]
=== NewTransactionLink - Registration
....
POST /auth/registration/init/link
....


==== Description
Endpoint for generating registration init public code. NOTE: One of parameters 'msisdn’ or ‘encryptedMsisdn’ is required in request.


==== Parameters

[options="header", cols=".^2a,.^3a,.^4a"]
|===
|Type|Name|Schema
|**Body**|**Authentication init request body** +
__optional__|<<_auth_registration_init_link_post_authentication_init_request_body,Authentication init request body>>
|===

[[_auth_registration_init_link_post_authentication_init_request_body]]
**Authentication init request body**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**correlationId** +
__optional__|Correlation id used by frontend app to correlate transaction responses.|string
|**encryptedMsisdn** +
__optional__|User's encrypted MSISDN. Hexadecimally encoded binary string. Required if 'msisdn' is not provided.|string
|**forClientId** +
__required__|Client id of application that requested authorization.|string
|**instanceIdentifier** +
__optional__|Identifier browser instance maing the request.|string
|**loa** +
__required__|Requested Level of Assurance.|enum (loa2, loa3)
|**msisdn** +
__required__|User's phone number.|string
|**operator** +
__required__|MNO id.|string
|**scopes** +
__optional__|Scopes that application requires|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|Status response|<<_auth_registration_init_link_post_response_200,Response 200>>
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_auth_registration_init_link_post_response_400,Response 400>>
|**401**|Access denied. Authentication is required to proceed processing request.|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_auth_registration_init_link_post_response_200]]
**Response 200**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**code** +
__required__|Code related to data stored in storage.|string
|**publicUrl** +
__required__|Public URL where new transaction is initiated from stored data.|string
|===

[[_auth_registration_init_link_post_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request, msisdn_missing_parameter, operator_missing_parameter, loa_missing_parameter, forClientId_missing_parameter, operator_could_not_be_found, invalid_loa, invalid_request)
|**error_description** +
__optional__|string
|===


==== Consumes

* `application/json`


==== Produces

* `application/json`


==== Tags

* Descriptor


==== Example HTTP response

===== Response 200
[source,json]
----
{
  "code" : "8c7067d9d774659e",
  "publicUrl" : "http://weblogin.etalio.dev/d/8c7067d9d774659e"
}
----


===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_auth_status_post]]
=== Status
....
POST /auth/status
....


==== Description
Endpoint for retrieving transaction status.


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|Status response|<<_auth_status_post_response_200,Response 200>>
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_auth_status_post_response_400,Response 400>>
|**401**|Access denied. Authentication is required to proceed processing request.|No Content
|**404**|Not Found. The requested resource could not be found but may be available again in the future. Subsequent requests by the client are permissible.|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_auth_status_post_response_200]]
**Response 200**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**clientName** +
__optional__|Requesting SP's name returned when status is final and the transaction is of type authentication or authorization|string
|**correlationId** +
__optional__|CorrelationId sent on init request.|string
|**data** +
__optional__|Data that has to be used with next action - values of form fields, display fields and placeholders. See Pluggable authenticators documentation for more information.|<<_auth_status_post_data,data>>
|**errors** +
__optional__|Errors as result of previous transaction request. See Pluggable authenticators documentation for more information.|<<_auth_status_post_errors,errors>>
|**fingerprint** +
__optional__|Descriptor fingerprint. Should not change between requests. If it changes fresh descriptor has to be feched and action has to be restarted.|string
|**mainAction** +
__optional__|Transaction's main action, used to display confirmation screen.|string
|**next** +
__optional__|Array of next actions that have to be executed.|< string > array
|**nextTransaction** +
__optional__|Next action that has to be executed.|< string > array
|**redirectedTo** +
__optional__|Next action that has to be executed.|string
|**status** +
__required__|Current transaction status.|enum (ready, pending, done, bad_request, rejected, redirected, failed, closed, pendingElsewhere, handedOver, doneInAnotherTab, cancelled, expired)
|**transactionId** +
__required__|Transaction id.|string
|**type** +
__required__|Transaction type.|enum (simple, registration, registrationSaa, authentication, activation, account, action, consent, operation, administration, reclaimProfile)
|===

[[_auth_status_post_data]]
**data**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**$authenticator id** +
__optional__|Action specific data.|object
|===

[[_auth_status_post_errors]]
**errors**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**$authenticator id** +
__optional__|Action specific errors.|object
|===

[[_auth_status_post_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request, transactionId_missing_parameter, access_denied, invalid_transactionId)
|**error_description** +
__optional__|string
|===


==== Consumes

* `application/json`


==== Produces

* `application/json`


==== Tags

* Descriptor


==== Example HTTP response

===== Response 200
[source,json]
----
{
  "status" : "ready",
  "type" : "authentication",
  "transactionId" : "78239500-74d6-11e6-acf7-1129c898746a",
  "next" : [ "sms::authenticateLoa2" ],
  "fingerprint" : "ba43edd03c1265f045a943aee8758d21",
  "mainAction" : "sms::authenticateLoa2",
  "correlationId" : "123456"
}
----


===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_auth_transactiontype_post]]
=== Transaction
....
POST /auth/{transactionType}
....


==== Description
Endpoints for executing transaction actions.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Path**|**transactionType** +
__required__|transactionType must be one of transaction types: 'authentication', 'registration', 'account', 'operation', 'administration', 'action', 'reclaimProfile'.|string
|**Body**|**Authentication init request body** +
__optional__||<<_auth_transactiontype_post_authentication_init_request_body,Authentication init request body>>
|===

[[_auth_transactiontype_post_authentication_init_request_body]]
**Authentication init request body**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**$descriptorId** +
__optional__|Object with action data for authneticator with $descriptorId. Structure depends on current action. For more information see documentation for Pluggable authenticators.|object
|**transactionId** +
__required__|Transaction id returned in previous calls (on same endpoint or init endpoint).|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|Status response|<<_auth_transactiontype_post_response_200,Response 200>>
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_auth_transactiontype_post_response_400,Response 400>>
|**401**|Access denied. Authentication is required to proceed processing request.|No Content
|**404**|Not Found. The requested resource could not be found but may be available again in the future. Subsequent requests by the client are permissible.|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_auth_transactiontype_post_response_200]]
**Response 200**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**clientName** +
__optional__|Requesting SP's name returned when status is final and the transaction is of type authentication or authorization|string
|**correlationId** +
__optional__|CorrelationId sent on init request.|string
|**data** +
__optional__|Data that has to be used with next action - values of form fields, display fields and placeholders. See Pluggable authenticators documentation for more information.|<<_auth_transactiontype_post_data,data>>
|**errors** +
__optional__|Errors as result of previous transaction request. See Pluggable authenticators documentation for more information.|<<_auth_transactiontype_post_errors,errors>>
|**fingerprint** +
__optional__|Descriptor fingerprint. Should not change between requests. If it changes fresh descriptor has to be feched and action has to be restarted.|string
|**mainAction** +
__optional__|Transaction's main action, used to display confirmation screen.|string
|**next** +
__optional__|Array of next actions that have to be executed.|< string > array
|**nextTransaction** +
__optional__|Next action that has to be executed.|< string > array
|**redirectedTo** +
__optional__|Next action that has to be executed.|string
|**status** +
__required__|Current transaction status.|enum (ready, pending, done, bad_request, rejected, redirected, failed, closed, pendingElsewhere, handedOver, doneInAnotherTab, cancelled, expired)
|**transactionId** +
__required__|Transaction id.|string
|**type** +
__required__|Transaction type.|enum (simple, registration, registrationSaa, authentication, activation, account, action, consent, operation, administration, reclaimProfile)
|===

[[_auth_transactiontype_post_data]]
**data**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**$authenticator id** +
__optional__|Action specific data.|object
|===

[[_auth_transactiontype_post_errors]]
**errors**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**$authenticator id** +
__optional__|Action specific errors.|object
|===

[[_auth_transactiontype_post_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request, transactionId_missing_parameter, invalid_transactionId)
|**error_description** +
__optional__|string
|===


==== Consumes

* `application/json`


==== Produces

* `application/json`


==== Tags

* Descriptor


==== Example HTTP response

===== Response 200
[source,json]
----
{
  "status" : "ready",
  "type" : "authentication",
  "transactionId" : "78239500-74d6-11e6-acf7-1129c898746a",
  "next" : [ "sms::authenticateLoa2" ],
  "fingerprint" : "ba43edd03c1265f045a943aee8758d21",
  "mainAction" : "sms::authenticateLoa2",
  "correlationId" : "123456"
}
----


===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_auth_transactiontype_close_post]]
=== Close transaction
....
POST /auth/{transactionType}/close
....


==== Description
Endpoints for closing transaction. Transaction can only be closed when current status is 'done'. In order to successfully complete transaction it has to be closed otherwise changes that were done with this transaction will not be comitted.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Path**|**transactionType** +
__required__|transactionType must be one of transaction types: 'authentication', 'registration', 'account', 'operation', 'administration', 'action', 'reclaimProfile'.|string
|**Body**|**Authentication init request body** +
__optional__||<<_auth_transactiontype_close_post_authentication_init_request_body,Authentication init request body>>
|===

[[_auth_transactiontype_close_post_authentication_init_request_body]]
**Authentication init request body**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**transactionId** +
__required__|Transaction id.|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|Status response|<<_auth_transactiontype_close_post_response_200,Response 200>>
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_auth_transactiontype_close_post_response_400,Response 400>>
|**401**|Access denied. Authentication is required to proceed processing request.|No Content
|**404**|Not Found. The requested resource could not be found but may be available again in the future. Subsequent requests by the client are permissible.|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_auth_transactiontype_close_post_response_200]]
**Response 200**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**authenticated** +
__optional__|Flag informing that current token got authenitcated as a result of transaction.|boolean
|**clientName** +
__optional__|Requesting SP's name returned when status is final and the transaction is of type authentication or authorization|string
|**correlationId** +
__optional__|CorrelationId sent on init request.|string
|**mainAction** +
__optional__|Transaction's main action, used to display confirmation screen.|string
|**status** +
__required__|Current transaction status.|enum (closed, handedOver, doneInAnotherTab, expired)
|**transactionId** +
__required__|Transaction id.|string
|**type** +
__required__|Transaction type.|enum (simple, registration, registrationSaa, authentication, authorization, account, action, operation, administration, reclaimProfile)
|===

[[_auth_transactiontype_close_post_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request, transactionId_missing_parameter, invalid_transactionId, invalid_transaction_status)
|**error_description** +
__optional__|string
|===


==== Consumes

* `application/json`


==== Produces

* `application/json`


==== Tags

* Descriptor


==== Example HTTP response

===== Response 200
[source,json]
----
{
  "status" : "closed",
  "type" : "account",
  "transactionId" : "22a83710-75c7-11e6-80c1-1129c898746a",
  "correlationId" : "1234",
  "mainAction" : "mdi::updateProfile"
}
----


===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_auth_transactiontype_forceclose_post]]
=== Force close transaction
....
POST /auth/{transactionType}/forceClose
....


==== Description
Endpoints for force closing transaction - closing transaction by token of instance in charge that is no longer control token but can close transaction on behalf of other tokens in this instance. Transaction can only be force closed when current status is 'doneInAnotherTab'. In order to successfully complete transaction it has to be closed otherwise changes that were done with this transaction will not be comitted.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Path**|**transactionType** +
__required__|transactionType must be one of transaction types: 'authentication', 'registration', 'account', 'operation', 'administration', 'action', 'reclaimProfile'.|string
|**Body**|**Authentication init request body** +
__optional__||<<_auth_transactiontype_forceclose_post_authentication_init_request_body,Authentication init request body>>
|===

[[_auth_transactiontype_forceclose_post_authentication_init_request_body]]
**Authentication init request body**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**transactionId** +
__required__|Transaction id.|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|Status response|<<_auth_transactiontype_forceclose_post_response_200,Response 200>>
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_auth_transactiontype_forceclose_post_response_400,Response 400>>
|**401**|Access denied. Authentication is required to proceed processing request.|No Content
|**404**|Not Found. The requested resource could not be found but may be available again in the future. Subsequent requests by the client are permissible.|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_auth_transactiontype_forceclose_post_response_200]]
**Response 200**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**authenticated** +
__optional__|Flag informing that current token got authenitcated as a result of transaction.|boolean
|**clientName** +
__optional__|Requesting SP's name returned when status is final and the transaction is of type authentication or authorization|string
|**correlationId** +
__optional__|CorrelationId sent on init request.|string
|**mainAction** +
__optional__|Transaction's main action, used to display confirmation screen.|string
|**status** +
__required__|Current transaction status.|enum (closed, handedOver, doneInAnotherTab, expired)
|**transactionId** +
__required__|Transaction id.|string
|**type** +
__required__|Transaction type.|enum (simple, registration, registrationSaa, authentication, authorization, account, action, operation, administration, reclaimProfile)
|===

[[_auth_transactiontype_forceclose_post_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request, transactionId_missing_parameter, invalid_transactionId, invalid_transaction_status)
|**error_description** +
__optional__|string
|===


==== Consumes

* `application/json`


==== Produces

* `application/json`


==== Tags

* Descriptor


==== Example HTTP response

===== Response 200
[source,json]
----
{
  "status" : "closed",
  "type" : "account",
  "transactionId" : "22a83710-75c7-11e6-80c1-1129c898746a",
  "correlationId" : "1234",
  "mainAction" : "mdi::updateProfile"
}
----


===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_authenticator_linking_confirm_code_post]]
=== Get authenticator linking status
....
POST /authenticator/linking/confirm/{code}
....


==== Description
Endpoint for getting status of authenticator linking transaction.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Path**|**code** +
__required__|Public code|string
|**Body**|**Authenticator linking confirmation endpoint.** +
__optional__||<<_authenticator_linking_confirm_code_post_authenticator_linking_confirmation_endpoint,Authenticator linking confirmation endpoint.>>
|===

[[_authenticator_linking_confirm_code_post_authenticator_linking_confirmation_endpoint]]
**Authenticator linking confirmation endpoint.**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**authenticator** +
__required__|Authenticator identification string.|string
|**authenticator_data** +
__optional__|Optional authenticator specific data to be stored to finish linking process, e.g. authenticator's user ID, name etc..|object
|**transaction_id** +
__optional__|Transaction id that should be specified upon linkink initialization|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|Authenticator linking confirmation endpoint response.|<<_authenticator_linking_confirm_code_post_response_200,Response 200>>
|**400**|Possible reasons for error: app could not be found, bad request, limit exeeds max, page does not exist, wrong page parameters,…|No Content
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**404**|Not Found. The requested resource could not be found but may be available again in the future. Subsequent requests by the client are permissible.|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_authenticator_linking_confirm_code_post_response_200]]
**Response 200**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**status** +
__required__|Status OK indicator.|string
|===


==== Produces

* `application/json`


==== Tags

* Authenticator


==== Example HTTP response

===== Response 200
[source,json]
----
{
  "status" : "OK"
}
----


[[_authenticator_linking_status_code_get]]
=== Get authenticator linking status
....
GET /authenticator/linking/status/{code}
....


==== Description
Endpoint for getting status of authenticator linking transaction.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Path**|**code** +
__required__|Public code|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|Authenticator linking status endpoint response.|< <<_authenticator_linking_status_code_get_response_200,Response 200>> > array
|**400**|Possible reasons for error: app could not be found, bad request, limit exeeds max, page does not exist, wrong page parameters,…|No Content
|**401**|Access denied. Authentication is required to proceed processing request.|No Content
|**404**|Not Found. The requested resource could not be found but may be available again in the future. Subsequent requests by the client are permissible.|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_authenticator_linking_status_code_get_response_200]]
**Response 200**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**authenticator_name** +
__optional__|Name of linked authenticator.|string
|**status** +
__required__|Status OK/pending indicator.|string
|===


==== Produces

* `application/json`


==== Tags

* Authenticator


[[_categories_get]]
=== Get categories collection
....
GET /categories
....


==== Description
Returns categories collection.


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|Collection of Categories.|< <<_categories_get_response_200,Response 200>> > array
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_categories_get_response_200]]
**Response 200**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**id** +
__optional__|Category's ID +
**Pattern** : `"[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}"`|string
|**label** +
__optional__|Category name|string
|**name** +
__optional__|Category name - deprecated|string
|===


==== Produces

* `application/json`


==== Tags

* Categories


==== Example HTTP response

===== Response 200
[source,json]
----
[ {
  "id" : "3cf1cd0e-ddf3-4bd3-85e8-45ad9b99621d",
  "name" : "Books",
  "label" : "Books"
} ]
----


[[_countries_get]]
=== Get countries collection
....
GET /countries
....


==== Description
Returns countries collection.


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|Collection of Countries.|< <<_countries_get_response_200,Response 200>> > array
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_countries_get_response_200]]
**Response 200**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**code** +
__optional__|Country ITU-T E.164 phone code|string
|**iso** +
__optional__|Country ISO 3166-1-Alpha-2 code|string
|**label** +
__optional__|Country name|string
|===


==== Produces

* `application/json`


==== Tags

* Countries


==== Example HTTP response

===== Response 200
[source,json]
----
[ {
  "iso" : "af",
  "code" : "93",
  "label" : "Afghanistan"
} ]
----


[[_discovery_post]]
=== Discovery
....
POST /discovery
....


==== Description
Discovery endpoint for retrieving operator endpoints. Basic Authorization is required. Can be used with trusted applications access token and client_id parameter. NOTE: query parameters are required in next combinations: msisdn, encryptedMsisdn and internalEncryptedMsisdn OR pcr and clientId.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Query**|**clientId** +
__required__|Client id, required for verifying PCR|string
|**Query**|**encryptedMsisdn** +
__required__|Encrypted msisdn|string
|**Query**|**ignoreExistingProfile** +
__optional__|Discover by ignoring existing profile. Used for re-registreing user.|boolean
|**Query**|**internalEncryptedMsisdn** +
__required__|Internal encryptedMsisdn msisdn|string
|**Query**|**mnoId** +
__optional__|Network operator uid. Used for bulk MNO migrations checks, churned users checks.|string
|**Query**|**msisdn** +
__required__|Phone number|string
|**Query**|**pcr** +
__required__|Pseudonymous customer reference|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|Internal discovery response.|<<_discovery_post_response_200,Response 200>>
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_discovery_post_response_400,Response 400>>
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_discovery_post_response_200]]
**Response 200**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**mnoId** +
__optional__|Discovered MNO uid for the user data (msisdn, encryptedMno, internalEncryptedMno, pcr…).|string
|**operatorSelector** +
__optional__|List of available netop-s when discovery status is 'operatorSelector'|< <<_discovery_post_operatorselector,operatorSelector>> > array
|**previousMnoId** +
__optional__|When migration is required, previousMnoId is returned together with current mnoId.|string
|**status** +
__optional__|Internal discovery status|enum (ok, registrationRequired, nonMdiOperator, operatorSelector, unknown)
|===

[[_discovery_post_operatorselector]]
**operatorSelector**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**mnoId** +
__required__|Network operator uid.|string
|**mnoName** +
__required__|Network operator name.|string
|===

[[_discovery_post_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* Discovery


==== Example HTTP response

===== Response 200
[source,json]
----
{
  "status" : "ok",
  "mnoId" : "e406cd96-6e3d-48fe-8a97-5f1498f158d7"
}
----


===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_hc_get]]
=== MCXG Cassandra DB health-check
....
GET /hc
....


==== Description
Health Check for MCXG Cassandra DB.


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|On success 200 code is returned.|No Content
|**400**|Possible reasons for error: app could not be found, bad request, limit exeeds max, page does not exist, wrong page parameters,…|No Content
|**404**|Not Found. The requested resource could not be found but may be available again in the future. Subsequent requests by the client are permissible.|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===


==== Produces

* `application/json`


==== Tags

* HealthChecks


[[_hc_descriptors_get]]
=== Descriptor health-check
....
GET /hc/descriptors
....


==== Description
Health Check PAM descriptors.


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|On success 200 code is returned.|No Content
|**400**|Possible reasons for error: app could not be found, bad request, limit exeeds max, page does not exist, wrong page parameters,…|No Content
|**404**|Not Found. The requested resource could not be found but may be available again in the future. Subsequent requests by the client are permissible.|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===


==== Produces

* `application/json`


==== Tags

* HealthChecks


[[_msisdn_post]]
=== Check MSISDN for MT procedure
....
POST /msisdn
....


==== Description
Check if MSISDN already registered with MobileConnect.


==== Parameters

[options="header", cols=".^2a,.^3a,.^4a"]
|===
|Type|Name|Schema
|**Body**|**Msisdn request body** +
__optional__|<<_msisdn_post_msisdn_request_body,Msisdn request body>>
|===

[[_msisdn_post_msisdn_request_body]]
**Msisdn request body**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**msisdn** +
__required__|Phone number|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|Msisdn check response body.|<<_msisdn_post_response_200,Response 200>>
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_msisdn_post_response_400,Response 400>>
|**401**|Access denied. Authentication is required to proceed processing request.|No Content
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**404**|Not Found. The requested resource could not be found but may be available again in the future. Subsequent requests by the client are permissible.|No Content
|**429**|Too Many Requests. The user or application has send too many requests in a given amount of time. If field 'blocked_for' is present in the response, it contains an information about how long the request will remain blocked (in seconds).|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_msisdn_post_response_200]]
**Response 200**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**status** +
__required__|Status|string
|===

[[_msisdn_post_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request, msisdn_missing_parameter, invalid_msisdn)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* Msisdn


==== Example HTTP response

===== Response 200
[source,json]
----
{
  "status" : "OK"
}
----


===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_msisdn_msisdn_mo_create_post]]
=== Create verification code for MO procedure
....
POST /msisdn (msisdn_mo_create)
....


==== Description
Generates verification code for given msisdn.


==== Parameters

[options="header", cols=".^2a,.^3a,.^4a"]
|===
|Type|Name|Schema
|**Body**|**Request body for MO code creation** +
__optional__|<<_msisdn_msisdn_mo_create_post_request_body_for_mo_code_creation,Request body for MO code creation>>
|===

[[_msisdn_msisdn_mo_create_post_request_body_for_mo_code_creation]]
**Request body for MO code creation**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**msisdn** +
__required__|Phone number|string
|**responseCode** +
__required__|Create code for given phone number|boolean
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**201**|Msisdn check MO response code.|<<_msisdn_msisdn_mo_create_post_response_201,Response 201>>
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_msisdn_msisdn_mo_create_post_response_400,Response 400>>
|**401**|Access denied. Authentication is required to proceed processing request.|No Content
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**404**|Not Found. The requested resource could not be found but may be available again in the future. Subsequent requests by the client are permissible.|No Content
|**429**|Too Many Requests. The user or application has send too many requests in a given amount of time. If field 'blocked_for' is present in the response, it contains an information about how long the request will remain blocked (in seconds).|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_msisdn_msisdn_mo_create_post_response_201]]
**Response 201**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**responseCode** +
__required__|MO code|string
|===

[[_msisdn_msisdn_mo_create_post_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request, msisdn_already_exists)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* Msisdn


==== Example HTTP response

===== Response 201
[source,json]
----
{
  "responseCode" : "xafg"
}
----


===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_msisdn_msisdn_mo_verify_post]]
=== Verify code for MO procedure
....
POST /msisdn (msisdn_mo_verify)
....


==== Description
Verify code for given MSISDN.


==== Parameters

[options="header", cols=".^2a,.^3a,.^4a"]
|===
|Type|Name|Schema
|**Body**|**Request body for MO code verification** +
__optional__|<<_msisdn_msisdn_mo_verify_post_request_body_for_mo_code_verification,Request body for MO code verification>>
|===

[[_msisdn_msisdn_mo_verify_post_request_body_for_mo_code_verification]]
**Request body for MO code verification**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**moSmsCode** +
__optional__|MO Code|string
|**msisdn** +
__required__|Phone number|string
|**responseCode** +
__optional__|MO code (deprecated)|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|Msisdn check response body.|<<_msisdn_msisdn_mo_verify_post_response_200,Response 200>>
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_msisdn_msisdn_mo_verify_post_response_400,Response 400>>
|**401**|Access denied. Authentication is required to proceed processing request.|No Content
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**404**|Not Found. The requested resource could not be found but may be available again in the future. Subsequent requests by the client are permissible.|No Content
|**429**|Too Many Requests. The user or application has send too many requests in a given amount of time. If field 'blocked_for' is present in the response, it contains an information about how long the request will remain blocked (in seconds).|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_msisdn_msisdn_mo_verify_post_response_200]]
**Response 200**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**status** +
__required__|Status|string
|===

[[_msisdn_msisdn_mo_verify_post_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request, invalid_code_format, invalid_code)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* Msisdn


==== Example HTTP response

===== Response 200
[source,json]
----
{
  "status" : "OK"
}
----


===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_msisdn_msisdn_mt_create_post]]
=== Create verification code for MT or opt-in procedure
....
POST /msisdn (msisdn_mt_create)
....


==== Description
Generate verification code for given MSISDN.


==== Parameters

[options="header", cols=".^2a,.^3a,.^4a"]
|===
|Type|Name|Schema
|**Body**|**Request body for MT code creation/verification** +
__optional__|<<_msisdn_msisdn_mt_create_post_request_body_for_mt_code_creation_verification,Request body for MT code creation/verification>>
|===

[[_msisdn_msisdn_mt_create_post_request_body_for_mt_code_creation_verification]]
**Request body for MT code creation/verification**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**code** +
__optional__|Request SMS MT Code (obsolete)|boolean
|**msisdn** +
__required__|Phone number|string
|**operator** +
__optional__|Operator ID|string
|**requestCode** +
__required__|Request SMS MT registration code|boolean
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**201**|Start MT or opt-in procedure request body.|<<_msisdn_msisdn_mt_create_post_response_201,Response 201>>
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_msisdn_msisdn_mt_create_post_response_400,Response 400>>
|**401**|Access denied. Authentication is required to proceed processing request.|No Content
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**404**|Not Found. The requested resource could not be found but may be available again in the future. Subsequent requests by the client are permissible.|No Content
|**429**|Too Many Requests. The user or application has send too many requests in a given amount of time. If field 'blocked_for' is present in the response, it contains an information about how long the request will remain blocked (in seconds).|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_msisdn_msisdn_mt_create_post_response_201]]
**Response 201**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**additionalData** +
__optional__|Additional information about the code being sent to the provided MSISDN.|< <<_msisdn_msisdn_mt_create_post_additionaldata,additionalData>> > array
|**status** +
__required__|Status|string
|===

[[_msisdn_msisdn_mt_create_post_additionaldata]]
**additionalData**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**name** +
__optional__|Holds the name of the field to be used in verification request. Contains either 'mtSmsCode' or 'optInCode'.|string
|**required** +
__optional__|Carries info about this field being required or not.|boolean
|**type** +
__optional__|Field value type.|string
|===

[[_msisdn_msisdn_mt_create_post_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request, msisdn_already_exists, operator_could_not_be_found, inactive_operator, cannot_register_operator)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* Msisdn


==== Example HTTP response

===== Response 201
[source,json]
----
{
  "status" : "OK",
  "additionalData" : [ {
    "name" : "mtSmsCode",
    "type" : "smsCode",
    "required" : true
  } ]
}
----


===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_msisdn_msisdn_mt_verify_post]]
=== Verify code for MT procedure
....
POST /msisdn (msisdn_mt_verify)
....


==== Description
Verify code for given MSISDN.


==== Parameters

[options="header", cols=".^2a,.^3a,.^4a"]
|===
|Type|Name|Schema
|**Body**|**Request body for MT code verification** +
__optional__|<<_msisdn_msisdn_mt_verify_post_request_body_for_mt_code_verification,Request body for MT code verification>>
|===

[[_msisdn_msisdn_mt_verify_post_request_body_for_mt_code_verification]]
**Request body for MT code verification**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**msisdn** +
__required__|Phone number|string
|**mtSmsCode** +
__optional__|MT Code (newer alternative to requestCode)|string
|**operator** +
__optional__|Operator ID|string
|**requestCode** +
__optional__|MT Code (deprecated)|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|Msisdn check response body.|<<_msisdn_msisdn_mt_verify_post_response_200,Response 200>>
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_msisdn_msisdn_mt_verify_post_response_400,Response 400>>
|**401**|Access denied. Authentication is required to proceed processing request.|No Content
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**404**|Not Found. The requested resource could not be found but may be available again in the future. Subsequent requests by the client are permissible.|No Content
|**429**|Too Many Requests. The user or application has send too many requests in a given amount of time. If field 'blocked_for' is present in the response, it contains an information about how long the request will remain blocked (in seconds).|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_msisdn_msisdn_mt_verify_post_response_200]]
**Response 200**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**status** +
__required__|Status|string
|===

[[_msisdn_msisdn_mt_verify_post_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request, invalid_code_format, invalid_code)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* Msisdn


==== Example HTTP response

===== Response 200
[source,json]
----
{
  "status" : "OK"
}
----


===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_msisdn_msisdn_opt_in_verify_post]]
=== Verify code for opt-in procedure
....
POST /msisdn (msisdn_opt_in_verify)
....


==== Description
Verify opt-in code for given MSISDN.


==== Parameters

[options="header", cols=".^2a,.^3a,.^4a"]
|===
|Type|Name|Schema
|**Body**|**Request body opt-in code verification** +
__optional__|<<_msisdn_msisdn_opt_in_verify_post_request_body_opt-in_code_verification,Request body opt-in code verification>>
|===

[[_msisdn_msisdn_opt_in_verify_post_request_body_opt-in_code_verification]]
**Request body opt-in code verification**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**msisdn** +
__required__|Phone number|string
|**operator** +
__required__|Operator ID|string
|**optInCode** +
__required__|Opt-in SMS registration code|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|Msisdn check response body.|<<_msisdn_msisdn_opt_in_verify_post_response_200,Response 200>>
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_msisdn_msisdn_opt_in_verify_post_response_400,Response 400>>
|**401**|Access denied. Authentication is required to proceed processing request.|No Content
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**404**|Not Found. The requested resource could not be found but may be available again in the future. Subsequent requests by the client are permissible.|No Content
|**429**|Too Many Requests. The user or application has send too many requests in a given amount of time. If field 'blocked_for' is present in the response, it contains an information about how long the request will remain blocked (in seconds).|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_msisdn_msisdn_opt_in_verify_post_response_200]]
**Response 200**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**status** +
__required__|Status|string
|===

[[_msisdn_msisdn_opt_in_verify_post_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request, optInCode_missing_parameter, invalid_code, operator_missing_parameter, operator_could_not_be_found, inactive_operator, cannot_register_operator)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* Msisdn


==== Example HTTP response

===== Response 200
[source,json]
----
{
  "status" : "OK"
}
----


===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_netops_post]]
=== Add operator
....
POST /netops
....


==== Description
Adds operator.


==== Parameters

[options="header", cols=".^2a,.^3a,.^4a"]
|===
|Type|Name|Schema
|**Body**|**Network operator request body** +
__optional__|<<_netops_post_network_operator_request_body,Network operator request body>>
|===

[[_netops_post_network_operator_request_body]]
**Network operator request body**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**acl** +
__optional__|List of MSISDN prefix patterns allowed to register with this MNO (BETA only).|< string > array
|**apePreference** +
__optional__|APE preference|<<_netops_post_apepreference,apePreference>>
|**brandname** +
__optional__|Contains brandname|string
|**cacheExpiry** +
__optional__|MCXG cache expiry time (in seconds) for profile data from this MNO. Default cache expiry is set to one week.|integer
|**contactPerson** +
__optional__|Contact person in CU. May contain anything (a name, a user ID, e-mail address, MSISDN, etc.).|string
|**countryCode** +
__optional__|Network operator country calling code|number
|**featuredBlacklist** +
__optional__|List of featured applications not to be displayed when this operator is in use.|< string > array
|**hniSet** +
__optional__|Network operator's array of Home Network Identities (MCC + MNC combinations)|< string > array
|**integrated** +
__optional__|Is Network operator part of Mobile Connect Accelerator network, possible values are HOST - is integrated, BETA - in BETA stage, INACTIVE - MNO is unnasigned, INTERNATIONAL - grandfather node.|string
|**iso** +
__optional__|Network operator ISO 3166 country code|string
|**label** +
__required__|Network operator name|string
|**mnp** +
__optional__|Network operator has mobile number portability|boolean
|**ndc** +
__optional__|List of national destination codes (NDC)|< string > array
|**policy** +
__optional__|Authentication policy containing prioritized list of LoA2 and LoA3 authentication methods.|<<_netops_post_policy,policy>>
|**promotedApps** +
__optional__|Promoted applications for network operator|< string > array
|**servingOperator** +
__optional__|Network operator serving name|string
|**shortRegistration** +
__optional__|Enables short registration on this MNO when the value is true.|boolean
|**showcase** +
__optional__|Contains showcase header|string
|**smsGateway** +
__optional__|SMS gateway to be used by network operator|string
|**timezone** +
__required__|Timezone of the MNO. The format must be in form Area/Location such as Europe/London. Required when adding MNO|string
|**trialEnd** +
__optional__|Trial period end time as an ISO8601 formatted date, used only when MNO is updated to integration state BETA. Default (three months) is used if not provided.|string
|**trialNotificationList** +
__optional__|Contacts that will be notified when trial is ending. Required when integrated is BETA.|<<_netops_post_trialnotificationlist,trialNotificationList>>
|**ttl** +
__optional__|Network operator transaction TTL configuration|<<_netops_post_ttl,ttl>>
|**umpHost** +
__optional__|User Management Portal url|string
|**ussdGateway** +
__optional__|USSD gateway to be used by network operator|string
|**visible** +
__optional__|Is Network operator visible|boolean
|===

[[_netops_post_apepreference]]
**apePreference**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**loa2** +
__optional__|Preferred authentication method for LoA2|string
|**loa3** +
__optional__|Preferred authentication method for LoA3|string
|===

[[_netops_post_policy]]
**policy**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**loa2** +
__required__|Prioritized list of LoA2 authentication methods.|< string > array
|**loa3** +
__required__|Prioritized list of LoA3 authentication methods.|< <<_netops_post_policy_loa3,loa3>> > array
|===

[[_netops_post_policy_loa3]]
**loa3**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**primary** +
__optional__|Primary authentication method - currently only 'password' is supported|string
|**secondary** +
__optional__|Secondary (confirmatory) authentication method|string
|===

[[_netops_post_trialnotificationlist]]
**trialNotificationList**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**email** +
__required__|Email addresses.|< string > array
|**msisdn** +
__required__|Phone number and name of the person.|< <<_netops_post_trialnotificationlist_msisdn,msisdn>> > array
|===

[[_netops_post_trialnotificationlist_msisdn]]
**msisdn**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**msisdn** +
__required__|Phone number.|string
|**name** +
__required__|Anything that describes the msisdn (first name, last name, …). It can be also empty.|string
|===

[[_netops_post_ttl]]
**ttl**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**authentication** +
__optional__|Authentication transaction|integer
|**authorization** +
__optional__|Authorization transaction|integer
|**bcAuthentication** +
__optional__|Backchannel authentication transaction|integer
|**bcAuthorization** +
__optional__|Backchannel authorization transaction|integer
|**registration** +
__optional__|Registration transaction|integer
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**201**|A representation of a netop.|<<_netops_post_response_201,Response 201>>
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_netops_post_response_400,Response 400>>
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_netops_post_response_201]]
**Response 201**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**EULA** +
__optional__|Contains MNO EULA (End User Licence Agreement) version. If this property is missing from response it means that MNO does not have it's own individual EULA, so default Mobile Connect Digital Identity EULA is used instead.|<<_netops_post_eula,EULA>>
|**acl** +
__optional__|List of MSISDN prefix patterns allowed to register with this MNO (BETA only).|< string > array
|**ape_preference** +
__optional__|APE settings for the MNO|<<_netops_post_ape_preference,ape_preference>>
|**brandname** +
__optional__|Contains brandname|string
|**cacheExpiry** +
__optional__|MCXG cache expiry time (in seconds) for profile data from this MNO.|integer
|**contactPerson** +
__optional__|Contact person in CU. May contain anything (a name, a user ID, e-mail address, MSISDN, etc.).|string
|**countryCode** +
__optional__|Network operator country calling code|number
|**featuredBlacklist** +
__optional__|List of featured applications not to be displayed when this operator is in use.|< string > array
|**hni** +
__optional__|NetOp's array of Home Network Identities (MCC + MNC combinations)|< string > array
|**id** +
__optional__|Network operator's ID +
**Pattern** : `"[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}"`|string
|**image** +
__optional__|Set of netop logo's|<<_netops_post_image,image>>
|**integrated** +
__optional__|Is Network operator part of Mobile Connect Accelerator network, possible values are HOST - is integrated, BETA - in BETA stage, INACTIVE - MNO is unnasigned, INTERNATIONAL - grandfather node.|string
|**iso** +
__optional__|Network operator ISO 3166 country code|string
|**label** +
__optional__|Network operator name|string
|**logo** +
__optional__|NetOp's logo (deprecated)|string
|**minimumUmpLoa** +
__optional__|Defines the LOA level to use on User Management Portal|string
|**mnp** +
__optional__|Network operator has mobile number portability|boolean
|**ndc** +
__optional__|List of national destination codes (NDC)|< string > array
|**policy** +
__optional__|Authentication policy containing prioritized list of LoA2 and LoA3 authentication methods.|<<_netops_post_policy,policy>>
|**promotedApps** +
__optional__|Promoted applications for network operator|< string > array
|**servingOperator** +
__optional__|Network operator serving name|string
|**shortRegistration** +
__optional__|Short Registration on this MNO is enabled when the value is true.|boolean
|**showcase** +
__optional__|Contains showcase header|string
|**smsGateway** +
__optional__|SMS gateway to be used by network operator|string
|**timezone** +
__optional__|Timezone of the MNO. The format must be in form Area/Location such as Europe/London.|string
|**trialEnd** +
__optional__|Trial period end time (ISO8601).|string
|**trialNotificationList** +
__optional__|Contacts that will be notified when trial is ending.|<<_netops_post_trialnotificationlist,trialNotificationList>>
|**trialStart** +
__optional__|Trial period beginning time (ISO8601).|string
|**ttl** +
__optional__|Network operator transaction TTL configuration|<<_netops_post_ttl,ttl>>
|**type** +
__optional__|Operator type (mcx or mckGlobal)|string
|**umpHost** +
__optional__|User Management Portal url|string
|**ussdGateway** +
__optional__|USSD gateway to be used by network operator|string
|**visible** +
__optional__|Is Network operator visible|boolean
|===

[[_netops_post_eula]]
**EULA**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**issuer** +
__optional__|MNO EULA issuer. If issuer is different than MNO ID, then EULA belongs to MNO host.|string
|**version** +
__optional__|MNO EULA version|string
|===

[[_netops_post_ape_preference]]
**ape_preference**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**loa2** +
__optional__|Preferred authentication method for LoA2|string
|**loa3** +
__optional__|Preferred authentication method for LoA3|string
|===

[[_netops_post_image]]
**image**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**120x120** +
__optional__|Image link|string
|**240x240** +
__optional__|Image link|string
|**500x500** +
__optional__|Image link|string
|**60x60** +
__optional__|Image link|string
|===

[[_netops_post_policy]]
**policy**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**loa2** +
__optional__|Prioritized list of LoA2 authentication methods.|< string > array
|**loa3** +
__optional__|Prioritized list of LoA3 authentication methods.|< <<_netops_post_policy_loa3,loa3>> > array
|===

[[_netops_post_policy_loa3]]
**loa3**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**primary** +
__optional__|Primary authentication method|string
|**secondary** +
__optional__|Secondary (confirmatory) authentication method|string
|===

[[_netops_post_trialnotificationlist]]
**trialNotificationList**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**email** +
__optional__|Email addresses.|< string > array
|**msisdn** +
__optional__|Phone number and name of the person.|< <<_netops_post_trialnotificationlist_msisdn,msisdn>> > array
|===

[[_netops_post_trialnotificationlist_msisdn]]
**msisdn**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**msisdn** +
__optional__|Phone number.|string
|**name** +
__optional__|Anything that describes the msisdn (first name, last name, …). It can be also empty.|string
|===

[[_netops_post_ttl]]
**ttl**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**authentication** +
__optional__|Authentication transaction|integer
|**authorization** +
__optional__|Authorization transaction|integer
|**bcAuthentication** +
__optional__|Backchannel authentication transaction|integer
|**bcAuthorization** +
__optional__|Backchannel authorization transaction|integer
|**registration** +
__optional__|Registration transaction|integer
|===

[[_netops_post_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request, invalid_hniSet, invalid_ndc, label_missing_parameter, image_wrong_type, application_could_not_be_found, image_wrong_format, showcase_invalid_parameter, brandname_invalid_parameter, invalid_policy, invalid_acl, invalid_ussdGateway, invalid_cacheExpiry, invalid_trialEnd, invalid_shortRegistration, invalid_trialNotificationList, invalid_timezone, timezone_missing_parameter, trialNotificationList_missing_parameter, invalid_minimumUmpLoa)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* Netops


==== Example HTTP response

===== Response 201
[source,json]
----
{
  "id" : "e406cd96-6e3d-48fe-8a97-5f1498f158d7",
  "label" : "Mobitel",
  "iso" : "si",
  "type" : "mcx",
  "countryCode" : 386,
  "servingOperator" : "Mobitel",
  "integrated" : "BETA",
  "hni" : [ "293.41", "293.64" ],
  "ndc" : [ "33" ],
  "mnp" : false,
  "smsGateway" : "IPX",
  "ussdGateway" : "mobitel_ussd",
  "promotedApps" : [ "00000000-de30-da1a-0000-db7b5bb0ac61" ],
  "featuredBlacklist" : [ "00000000-de30-da1a-0000-adb342bcc881" ],
  "showcase" : "Showcase header",
  "brandname" : "Commercial brandname",
  "EULA" : {
    "issuer" : "e406cd96-6e3d-48fe-8a97-5f1498f158d7",
    "version" : "1.0"
  },
  "policy" : {
    "loa2" : [ "password", "sms" ],
    "loa3" : [ {
      "primary" : "password",
      "secondary" : "sms"
    }, {
      "primary" : "password",
      "secondary" : "push"
    } ]
  },
  "acl" : [ "9355" ],
  "cacheExpiry" : 172800,
  "shortRegistration" : true,
  "trialNotificationList" : {
    "msisdn" : [ {
      "name" : "FirstName LastName1",
      "msisdn" : "+38640333771"
    }, {
      "name" : "FirstName LastName2",
      "msisdn" : "+38640333772"
    } ],
    "email" : [ "example1@example.org", "example2@example.org" ]
  },
  "timezone" : "Europe/London",
  "minimumUmpLoa" : "loa3",
  "ttl" : {
    "registration" : 300,
    "authentication" : 180,
    "bcAuthentication" : 180,
    "authorization" : 180,
    "bcAuthorization" : 180
  },
  "umpHost" : "https://www.eaxmpleportal.com"
}
----


===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_netops_get]]
=== Get netops collection
....
GET /netops
....


==== Description
Returns netops collection.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Query**|**hni** +
__optional__|Parameter for filtering network operators by HNI set (combination of MCC and MNC). Format: MCC.MNC, Example: 110.44.|string
|**Query**|**integrated** +
__optional__|Parameter for filtering network operators by MDI integration state. Valid values: 'BETA', 'HOST', 'INTERNATIONAL' and 'INACTIVE'.|string
|**Query**|**iso** +
__optional__|Parameter for filtering network operators by ISO country code (lower or upper case).|string
|**Query**|**limit** +
__optional__|Pagination parameter, indicates how many records should be returned by page. NOTE: Default value is 20 and Max value is 50|string
|**Query**|**msisdn** +
__optional__|Parameter for filtering network operators by country code from provided MSISDN.|string
|**Query**|**page** +
__optional__|Pagination parameter, indicates which page is to be returned from database. NOTE: Default value is 1.|number
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|Collection of Net operators.|< <<_netops_get_response_200,Response 200>> > array
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_netops_get_response_400,Response 400>>
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_netops_get_response_200]]
**Response 200**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**EULA** +
__optional__|Contains MNO EULA (End User Licence Agreement) version. If this property is missing from response it means that MNO does not have it's own individual EULA, so default Mobile Connect Digital Identity EULA is used instead.|<<_netops_get_eula,EULA>>
|**acl** +
__optional__|List of MSISDN prefix patterns allowed to register with this MNO (BETA only).|< string > array
|**ape_preference** +
__optional__|APE settings for the MNO|<<_netops_get_ape_preference,ape_preference>>
|**brandname** +
__optional__|Contains brandname|string
|**cacheExpiry** +
__optional__|MCXG cache expiry time (in seconds) for profile data from this MNO.|integer
|**contactPerson** +
__optional__|Contact person in CU. May contain anything (a name, a user ID, e-mail address, MSISDN, etc.).|string
|**countryCode** +
__optional__|Network operator country calling code|number
|**featuredBlacklist** +
__optional__|List of featured applications not to be displayed when this operator is in use.|< string > array
|**hni** +
__optional__|NetOp's array of Home Network Identities (MCC + MNC combinations)|< string > array
|**id** +
__optional__|Network operator's ID +
**Pattern** : `"[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}"`|string
|**image** +
__optional__|Set of netop logo's|<<_netops_get_image,image>>
|**integrated** +
__optional__|Is Network operator part of Mobile Connect Accelerator network, possible values are HOST - is integrated, BETA - in BETA stage, INACTIVE - MNO is unnasigned, INTERNATIONAL - grandfather node.|string
|**iso** +
__optional__|Network operator ISO 3166 country code|string
|**label** +
__optional__|Network operator name|string
|**logo** +
__optional__|NetOp's logo (deprecated)|string
|**minimumUmpLoa** +
__optional__|Defines the LOA level to use on User Management Portal|string
|**mnp** +
__optional__|Network operator has mobile number portability|boolean
|**ndc** +
__optional__|List of national destination codes (NDC)|< string > array
|**policy** +
__optional__|Authentication policy containing prioritized list of LoA2 and LoA3 authentication methods.|<<_netops_get_policy,policy>>
|**promotedApps** +
__optional__|Promoted applications for network operator|< string > array
|**servingOperator** +
__optional__|Network operator serving name|string
|**shortRegistration** +
__optional__|Short Registration on this MNO is enabled when the value is true.|boolean
|**showcase** +
__optional__|Contains showcase header|string
|**smsGateway** +
__optional__|SMS gateway to be used by network operator|string
|**timezone** +
__optional__|Timezone of the MNO. The format must be in form Area/Location such as Europe/London.|string
|**trialEnd** +
__optional__|Trial period end time (ISO8601).|string
|**trialNotificationList** +
__optional__|Contacts that will be notified when trial is ending.|<<_netops_get_trialnotificationlist,trialNotificationList>>
|**trialStart** +
__optional__|Trial period beginning time (ISO8601).|string
|**ttl** +
__optional__|Network operator transaction TTL configuration|<<_netops_get_ttl,ttl>>
|**type** +
__optional__|Operator type (mcx or mckGlobal)|string
|**umpHost** +
__optional__|User Management Portal url|string
|**ussdGateway** +
__optional__|USSD gateway to be used by network operator|string
|**visible** +
__optional__|Is Network operator visible|boolean
|===

[[_netops_get_eula]]
**EULA**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**issuer** +
__optional__|MNO EULA issuer. If issuer is different than MNO ID, then EULA belongs to MNO host.|string
|**version** +
__optional__|MNO EULA version|string
|===

[[_netops_get_ape_preference]]
**ape_preference**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**loa2** +
__optional__|Preferred authentication method for LoA2|string
|**loa3** +
__optional__|Preferred authentication method for LoA3|string
|===

[[_netops_get_image]]
**image**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**120x120** +
__optional__|Image link|string
|**240x240** +
__optional__|Image link|string
|**500x500** +
__optional__|Image link|string
|**60x60** +
__optional__|Image link|string
|===

[[_netops_get_policy]]
**policy**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**loa2** +
__optional__|Prioritized list of LoA2 authentication methods.|< string > array
|**loa3** +
__optional__|Prioritized list of LoA3 authentication methods.|< <<_netops_get_policy_loa3,loa3>> > array
|===

[[_netops_get_policy_loa3]]
**loa3**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**primary** +
__optional__|Primary authentication method|string
|**secondary** +
__optional__|Secondary (confirmatory) authentication method|string
|===

[[_netops_get_trialnotificationlist]]
**trialNotificationList**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**email** +
__optional__|Email addresses.|< string > array
|**msisdn** +
__optional__|Phone number and name of the person.|< <<_netops_get_trialnotificationlist_msisdn,msisdn>> > array
|===

[[_netops_get_trialnotificationlist_msisdn]]
**msisdn**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**msisdn** +
__optional__|Phone number.|string
|**name** +
__optional__|Anything that describes the msisdn (first name, last name, …). It can be also empty.|string
|===

[[_netops_get_ttl]]
**ttl**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**authentication** +
__optional__|Authentication transaction|integer
|**authorization** +
__optional__|Authorization transaction|integer
|**bcAuthentication** +
__optional__|Backchannel authentication transaction|integer
|**bcAuthorization** +
__optional__|Backchannel authorization transaction|integer
|**registration** +
__optional__|Registration transaction|integer
|===

[[_netops_get_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request, wrong_page_parameter, wrong_limit_parameter, limit_exceeds_max, invalid_msisdn, invalid_hniSet, invalid_iso, invalid_integrated)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* Netops


==== Example HTTP response

===== Response 200
[source,json]
----
[ {
  "id" : "e406cd96-6e3d-48fe-8a97-5f1498f158d7",
  "label" : "Mobitel",
  "servingOperator" : "Mobitel",
  "iso" : "si",
  "countryCode" : 386,
  "hni" : [ "293.41", "293.64" ]
} ]
----


===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_netops_put]]
=== Update operator
....
PUT /netops
....


==== Description
Update operator.


==== Parameters

[options="header", cols=".^2a,.^3a,.^4a"]
|===
|Type|Name|Schema
|**Body**|**Network operator request body** +
__optional__|<<_netops_put_network_operator_request_body,Network operator request body>>
|===

[[_netops_put_network_operator_request_body]]
**Network operator request body**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**acl** +
__optional__|List of MSISDN prefix patterns allowed to register with this MNO (BETA only).|< string > array
|**apePreference** +
__optional__|APE preference|<<_netops_put_apepreference,apePreference>>
|**brandname** +
__optional__|Contains brandname|string
|**cacheExpiry** +
__optional__|MCXG cache expiry time (in seconds) for profile data from this MNO. Default cache expiry is set to one week.|integer
|**contactPerson** +
__optional__|Contact person in CU. May contain anything (a name, a user ID, e-mail address, MSISDN, etc.).|string
|**countryCode** +
__optional__|Network operator country calling code|number
|**featuredBlacklist** +
__optional__|List of featured applications not to be displayed when this operator is in use.|< string > array
|**hniSet** +
__optional__|Network operator's array of Home Network Identities (MCC + MNC combinations)|< string > array
|**integrated** +
__optional__|Is Network operator part of Mobile Connect Accelerator network, possible values are HOST - is integrated, BETA - in BETA stage, INACTIVE - MNO is unnasigned, INTERNATIONAL - grandfather node.|string
|**iso** +
__optional__|Network operator ISO 3166 country code|string
|**label** +
__required__|Network operator name|string
|**mnp** +
__optional__|Network operator has mobile number portability|boolean
|**ndc** +
__optional__|List of national destination codes (NDC)|< string > array
|**policy** +
__optional__|Authentication policy containing prioritized list of LoA2 and LoA3 authentication methods.|<<_netops_put_policy,policy>>
|**promotedApps** +
__optional__|Promoted applications for network operator|< string > array
|**servingOperator** +
__optional__|Network operator serving name|string
|**shortRegistration** +
__optional__|Enables short registration on this MNO when the value is true.|boolean
|**showcase** +
__optional__|Contains showcase header|string
|**smsGateway** +
__optional__|SMS gateway to be used by network operator|string
|**timezone** +
__required__|Timezone of the MNO. The format must be in form Area/Location such as Europe/London. Required when adding MNO|string
|**trialEnd** +
__optional__|Trial period end time as an ISO8601 formatted date, used only when MNO is updated to integration state BETA. Default (three months) is used if not provided.|string
|**trialNotificationList** +
__optional__|Contacts that will be notified when trial is ending. Required when integrated is BETA.|<<_netops_put_trialnotificationlist,trialNotificationList>>
|**ttl** +
__optional__|Network operator transaction TTL configuration|<<_netops_put_ttl,ttl>>
|**umpHost** +
__optional__|User Management Portal url|string
|**ussdGateway** +
__optional__|USSD gateway to be used by network operator|string
|**visible** +
__optional__|Is Network operator visible|boolean
|===

[[_netops_put_apepreference]]
**apePreference**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**loa2** +
__optional__|Preferred authentication method for LoA2|string
|**loa3** +
__optional__|Preferred authentication method for LoA3|string
|===

[[_netops_put_policy]]
**policy**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**loa2** +
__required__|Prioritized list of LoA2 authentication methods.|< string > array
|**loa3** +
__required__|Prioritized list of LoA3 authentication methods.|< <<_netops_put_policy_loa3,loa3>> > array
|===

[[_netops_put_policy_loa3]]
**loa3**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**primary** +
__optional__|Primary authentication method - currently only 'password' is supported|string
|**secondary** +
__optional__|Secondary (confirmatory) authentication method|string
|===

[[_netops_put_trialnotificationlist]]
**trialNotificationList**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**email** +
__required__|Email addresses.|< string > array
|**msisdn** +
__required__|Phone number and name of the person.|< <<_netops_put_trialnotificationlist_msisdn,msisdn>> > array
|===

[[_netops_put_trialnotificationlist_msisdn]]
**msisdn**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**msisdn** +
__required__|Phone number.|string
|**name** +
__required__|Anything that describes the msisdn (first name, last name, …). It can be also empty.|string
|===

[[_netops_put_ttl]]
**ttl**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**authentication** +
__optional__|Authentication transaction|integer
|**authorization** +
__optional__|Authorization transaction|integer
|**bcAuthentication** +
__optional__|Backchannel authentication transaction|integer
|**bcAuthorization** +
__optional__|Backchannel authorization transaction|integer
|**registration** +
__optional__|Registration transaction|integer
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**202**|A representation of a netop.|<<_netops_put_response_202,Response 202>>
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_netops_put_response_400,Response 400>>
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_netops_put_response_202]]
**Response 202**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**EULA** +
__optional__|Contains MNO EULA (End User Licence Agreement) version. If this property is missing from response it means that MNO does not have it's own individual EULA, so default Mobile Connect Digital Identity EULA is used instead.|<<_netops_put_eula,EULA>>
|**acl** +
__optional__|List of MSISDN prefix patterns allowed to register with this MNO (BETA only).|< string > array
|**ape_preference** +
__optional__|APE settings for the MNO|<<_netops_put_ape_preference,ape_preference>>
|**brandname** +
__optional__|Contains brandname|string
|**cacheExpiry** +
__optional__|MCXG cache expiry time (in seconds) for profile data from this MNO.|integer
|**contactPerson** +
__optional__|Contact person in CU. May contain anything (a name, a user ID, e-mail address, MSISDN, etc.).|string
|**countryCode** +
__optional__|Network operator country calling code|number
|**featuredBlacklist** +
__optional__|List of featured applications not to be displayed when this operator is in use.|< string > array
|**hni** +
__optional__|NetOp's array of Home Network Identities (MCC + MNC combinations)|< string > array
|**id** +
__optional__|Network operator's ID +
**Pattern** : `"[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}"`|string
|**image** +
__optional__|Set of netop logo's|<<_netops_put_image,image>>
|**integrated** +
__optional__|Is Network operator part of Mobile Connect Accelerator network, possible values are HOST - is integrated, BETA - in BETA stage, INACTIVE - MNO is unnasigned, INTERNATIONAL - grandfather node.|string
|**iso** +
__optional__|Network operator ISO 3166 country code|string
|**label** +
__optional__|Network operator name|string
|**logo** +
__optional__|NetOp's logo (deprecated)|string
|**minimumUmpLoa** +
__optional__|Defines the LOA level to use on User Management Portal|string
|**mnp** +
__optional__|Network operator has mobile number portability|boolean
|**ndc** +
__optional__|List of national destination codes (NDC)|< string > array
|**policy** +
__optional__|Authentication policy containing prioritized list of LoA2 and LoA3 authentication methods.|<<_netops_put_policy,policy>>
|**promotedApps** +
__optional__|Promoted applications for network operator|< string > array
|**servingOperator** +
__optional__|Network operator serving name|string
|**shortRegistration** +
__optional__|Short Registration on this MNO is enabled when the value is true.|boolean
|**showcase** +
__optional__|Contains showcase header|string
|**smsGateway** +
__optional__|SMS gateway to be used by network operator|string
|**timezone** +
__optional__|Timezone of the MNO. The format must be in form Area/Location such as Europe/London.|string
|**trialEnd** +
__optional__|Trial period end time (ISO8601).|string
|**trialNotificationList** +
__optional__|Contacts that will be notified when trial is ending.|<<_netops_put_trialnotificationlist,trialNotificationList>>
|**trialStart** +
__optional__|Trial period beginning time (ISO8601).|string
|**ttl** +
__optional__|Network operator transaction TTL configuration|<<_netops_put_ttl,ttl>>
|**type** +
__optional__|Operator type (mcx or mckGlobal)|string
|**umpHost** +
__optional__|User Management Portal url|string
|**ussdGateway** +
__optional__|USSD gateway to be used by network operator|string
|**visible** +
__optional__|Is Network operator visible|boolean
|===

[[_netops_put_eula]]
**EULA**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**issuer** +
__optional__|MNO EULA issuer. If issuer is different than MNO ID, then EULA belongs to MNO host.|string
|**version** +
__optional__|MNO EULA version|string
|===

[[_netops_put_ape_preference]]
**ape_preference**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**loa2** +
__optional__|Preferred authentication method for LoA2|string
|**loa3** +
__optional__|Preferred authentication method for LoA3|string
|===

[[_netops_put_image]]
**image**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**120x120** +
__optional__|Image link|string
|**240x240** +
__optional__|Image link|string
|**500x500** +
__optional__|Image link|string
|**60x60** +
__optional__|Image link|string
|===

[[_netops_put_policy]]
**policy**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**loa2** +
__optional__|Prioritized list of LoA2 authentication methods.|< string > array
|**loa3** +
__optional__|Prioritized list of LoA3 authentication methods.|< <<_netops_put_policy_loa3,loa3>> > array
|===

[[_netops_put_policy_loa3]]
**loa3**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**primary** +
__optional__|Primary authentication method|string
|**secondary** +
__optional__|Secondary (confirmatory) authentication method|string
|===

[[_netops_put_trialnotificationlist]]
**trialNotificationList**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**email** +
__optional__|Email addresses.|< string > array
|**msisdn** +
__optional__|Phone number and name of the person.|< <<_netops_put_trialnotificationlist_msisdn,msisdn>> > array
|===

[[_netops_put_trialnotificationlist_msisdn]]
**msisdn**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**msisdn** +
__optional__|Phone number.|string
|**name** +
__optional__|Anything that describes the msisdn (first name, last name, …). It can be also empty.|string
|===

[[_netops_put_ttl]]
**ttl**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**authentication** +
__optional__|Authentication transaction|integer
|**authorization** +
__optional__|Authorization transaction|integer
|**bcAuthentication** +
__optional__|Backchannel authentication transaction|integer
|**bcAuthorization** +
__optional__|Backchannel authorization transaction|integer
|**registration** +
__optional__|Registration transaction|integer
|===

[[_netops_put_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request, invalid_hniSet, invalid_ndc, image_wrong_format, image_wrong_type, application_could_not_be_found, showcase_invalid_parameter, brandname_invalid_parameter, integrated_invalid_parameter, invalid_policy, invalid_acl, invalid_ussdGateway, invalid_cacheExpiry, invalid_trialEnd, invalid_shortRegistration, invalid_trialNotificationList, invalid_timezone, timezone_missing_parameter, trialNotificationList_missing_parameter, invalid_minimumUmpLoa)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* Netops


==== Example HTTP response

===== Response 202
[source,json]
----
{
  "id" : "e406cd96-6e3d-48fe-8a97-5f1498f158d7",
  "label" : "Mobitel",
  "iso" : "si",
  "type" : "mcx",
  "countryCode" : 386,
  "servingOperator" : "Mobitel",
  "integrated" : "BETA",
  "hni" : [ "293.41", "293.64" ],
  "ndc" : [ "33" ],
  "mnp" : false,
  "smsGateway" : "IPX",
  "ussdGateway" : "mobitel_ussd",
  "promotedApps" : [ "00000000-de30-da1a-0000-db7b5bb0ac61" ],
  "featuredBlacklist" : [ "00000000-de30-da1a-0000-adb342bcc881" ],
  "showcase" : "Showcase header",
  "brandname" : "Commercial brandname",
  "EULA" : {
    "issuer" : "e406cd96-6e3d-48fe-8a97-5f1498f158d7",
    "version" : "1.0"
  },
  "policy" : {
    "loa2" : [ "password", "sms" ],
    "loa3" : [ {
      "primary" : "password",
      "secondary" : "sms"
    }, {
      "primary" : "password",
      "secondary" : "push"
    } ]
  },
  "acl" : [ "9355" ],
  "cacheExpiry" : 172800,
  "shortRegistration" : true,
  "trialNotificationList" : {
    "msisdn" : [ {
      "name" : "FirstName LastName1",
      "msisdn" : "+38640333771"
    }, {
      "name" : "FirstName LastName2",
      "msisdn" : "+38640333772"
    } ],
    "email" : [ "example1@example.org", "example2@example.org" ]
  },
  "timezone" : "Europe/London",
  "minimumUmpLoa" : "loa3",
  "ttl" : {
    "registration" : 300,
    "authentication" : 180,
    "bcAuthentication" : 180,
    "authorization" : 180,
    "bcAuthorization" : 180
  },
  "umpHost" : "https://www.eaxmpleportal.com"
}
----


===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_netops-groups_post]]
=== Add operator group
....
POST /netops-groups
....


==== Description
Adds operator group.


==== Parameters

[options="header", cols=".^2a,.^3a,.^4a"]
|===
|Type|Name|Schema
|**Body**|**Operator group request body** +
__optional__|<<_netops-groups_post_operator_group_request_body,Operator group request body>>
|===

[[_netops-groups_post_operator_group_request_body]]
**Operator group request body**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**label** +
__required__|Network operator's group name|string
|**netOps** +
__required__|Array of netops in netop group|< string > array
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**201**|A representation of a network operator group.|<<_netops-groups_post_response_201,Response 201>>
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_netops-groups_post_response_400,Response 400>>
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_netops-groups_post_response_201]]
**Response 201**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**id** +
__optional__|Network operator ID +
**Pattern** : `"[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}"`|string
|**label** +
__optional__|Network operator group name|string
|**netOps** +
__optional__|Array of network operators in group|< string > array
|===

[[_netops-groups_post_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request, operator_could_not_be_found, label_missing_parameter, netOps_missing_parameter)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* NetopsGroups


==== Example HTTP response

===== Response 201
[source,json]
----
{
  "_id" : "088b5c1b-21bb-47dc-b02d-eba4273a451f",
  "label" : "Vodaphone",
  "netOps" : [ "60a81b9b-1289-4366-ad1a-bc8114f9640d" ]
}
----


===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_netops-groups_get]]
=== Get operator groups collection
....
GET /netops-groups
....


==== Description
Returns operator groups collection.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Query**|**limit** +
__optional__|Pagination parameter, indicates how many records should be returned by page. NOTE: Default value is 20 and Max value is 50|string
|**Query**|**page** +
__optional__|Pagination parameter, indicates which page is to be returned from database. NOTE: Default value is 1.|number
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|Collection of Net operator groups.|< <<_netops-groups_get_response_200,Response 200>> > array
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_netops-groups_get_response_400,Response 400>>
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_netops-groups_get_response_200]]
**Response 200**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**id** +
__optional__|Network operator ID +
**Pattern** : `"[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}"`|string
|**label** +
__optional__|Network operator group name|string
|**netOps** +
__optional__|Array of network operators in group|< string > array
|===

[[_netops-groups_get_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request, wrong_page_parameter, wrong_limit_parameter, limit_exceeds_max)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* NetopsGroups


==== Example HTTP response

===== Response 200
[source,json]
----
[ {
  "_id" : "088b5c1b-21bb-47dc-b02d-eba4273a451f",
  "groupName" : "Vodaphone",
  "netOps" : [ "60a81b9b-1289-4366-ad1a-bc8114f9640d" ]
} ]
----


===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_netops-groups_netopgroupid_get]]
=== Get network operator group
....
GET /netops-groups/{netOpGroupId}
....


==== Description
Returns operator group data.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Path**|**netOpGroupId** +
__required__|Network operator group id|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|A representation of a network operator group.|<<_netops-groups_netopgroupid_get_response_200,Response 200>>
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_netops-groups_netopgroupid_get_response_400,Response 400>>
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**404**|Not Found. The requested resource could not be found but may be available again in the future. Subsequent requests by the client are permissible.|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_netops-groups_netopgroupid_get_response_200]]
**Response 200**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**id** +
__optional__|Network operator ID +
**Pattern** : `"[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}"`|string
|**label** +
__optional__|Network operator group name|string
|**netOps** +
__optional__|Array of network operators in group|< string > array
|===

[[_netops-groups_netopgroupid_get_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request, no_route)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* NetopsGroups


==== Example HTTP response

===== Response 200
[source,json]
----
{
  "_id" : "088b5c1b-21bb-47dc-b02d-eba4273a451f",
  "label" : "Vodaphone",
  "netOps" : [ "60a81b9b-1289-4366-ad1a-bc8114f9640d" ]
}
----


===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_netops-groups_netopgroupid_put]]
=== Update operator group
....
PUT /netops-groups/{netOpGroupId}
....


==== Description
Update operator group.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Path**|**netOpGroupId** +
__required__|Network operator group id|string
|**Body**|**Operator group request body** +
__optional__||<<_netops-groups_netopgroupid_put_operator_group_request_body,Operator group request body>>
|===

[[_netops-groups_netopgroupid_put_operator_group_request_body]]
**Operator group request body**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**label** +
__required__|Network operator's group name|string
|**netOps** +
__required__|Array of netops in netop group|< string > array
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**202**|A representation of a network operator group.|<<_netops-groups_netopgroupid_put_response_202,Response 202>>
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_netops-groups_netopgroupid_put_response_400,Response 400>>
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**404**|Not Found. The requested resource could not be found but may be available again in the future. Subsequent requests by the client are permissible.|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_netops-groups_netopgroupid_put_response_202]]
**Response 202**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**id** +
__optional__|Network operator ID +
**Pattern** : `"[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}"`|string
|**label** +
__optional__|Network operator group name|string
|**netOps** +
__optional__|Array of network operators in group|< string > array
|===

[[_netops-groups_netopgroupid_put_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request, operator_could_not_be_found)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* NetopsGroups


==== Example HTTP response

===== Response 202
[source,json]
----
{
  "_id" : "088b5c1b-21bb-47dc-b02d-eba4273a451f",
  "label" : "Vodaphone",
  "netOps" : [ "60a81b9b-1289-4366-ad1a-bc8114f9640d" ]
}
----


===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_netops-groups_netopgroupid_delete]]
=== Remove an operator group
....
DELETE /netops-groups/{netOpGroupId}
....


==== Description
Removes existing operator group.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Path**|**netOpGroupId** +
__required__|Network operator group id|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**204**|On success 204 code is returned.|No Content
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_netops-groups_netopgroupid_delete_response_400,Response 400>>
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**404**|Not Found. The requested resource could not be found but may be available again in the future. Subsequent requests by the client are permissible.|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_netops-groups_netopgroupid_delete_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* NetopsGroups


==== Example HTTP response

===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_netops_list_get]]
=== Operator list report
....
GET /netops/list
....


==== Description
Netop endpoint for retrieving list of operators. Authenticated access token is needed.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Query**|**fields** +
__optional__|Whitespace separated string with operator fields. If string is empty all fields are returned.|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|Network operator response body.|<<_netops_list_get_response_200,Response 200>>
|**400**|Possible reasons for error: app could not be found, bad request, limit exeeds max, page does not exist, wrong page parameters,…|No Content
|**401**|Access denied. Authentication is required to proceed processing request.|No Content
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**404**|Not Found. The requested resource could not be found but may be available again in the future. Subsequent requests by the client are permissible.|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_netops_list_get_response_200]]
**Response 200**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**operators** +
__optional__|Array of filtered operator fields|< object > array
|===


==== Produces

* `application/json`


==== Tags

* Netops


==== Example HTTP response

===== Response 200
[source,json]
----
{
  "operators" : [ {
    "uid" : "9055338d4-fa92-4305-aa5b-38383b274544",
    "name" : "SPRINT Spectrum L.P.- PA"
  }, {
    "uid" : "04460944f-832d-4e1a-9ce2-8e537f328d59",
    "name" : "Keystone Wireless LLC"
  }, {
    "uid" : "c5ef4e46-0d33-43599-9a38-647b62f2d94c",
    "name" : "Hi3G Access AB"
  } ]
}
----


[[_netops_mnp_iso_put]]
=== Update mnp per country
....
PUT /netops/mnp/{iso}
....


==== Description
Update mobile number portability per country.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Path**|**iso** +
__required__|iso country code|string
|**Body**|**Update mobile number portability per country** +
__optional__||<<_netops_mnp_iso_put_update_mobile_number_portability_per_country,Update mobile number portability per country>>
|===

[[_netops_mnp_iso_put_update_mobile_number_portability_per_country]]
**Update mobile number portability per country**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**mnp** +
__required__|Network operators has mobile number portability|boolean
|**ndc** +
__optional__|List of national destination codes (NDC)|< string > array
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**202**|Update mobile number portability per country response body.|<<_netops_mnp_iso_put_response_202,Response 202>>
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_netops_mnp_iso_put_response_400,Response 400>>
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_netops_mnp_iso_put_response_202]]
**Response 202**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**status** +
__required__|Status OK.|string
|===

[[_netops_mnp_iso_put_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request, invalid_hniSet, invalid_ndc, image_wrong_format, image_wrong_type, application_could_not_be_found, showcase_invalid_parameter, brandname_invalid_parameter, integrated_invalid_parameter, invalid_policy, invalid_acl, invalid_ussdGateway, invalid_cacheExpiry, invalid_trialEnd, invalid_shortRegistration, invalid_trialNotificationList, invalid_timezone, timezone_missing_parameter, trialNotificationList_missing_parameter, invalid_minimumUmpLoa)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* Netops


==== Example HTTP response

===== Response 202
[source,json]
----
{
  "status" : "OK"
}
----


===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_netops_netopid_get]]
=== Get network operator
....
GET /netops/{netOpId}
....


==== Description
Returns operator data.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Path**|**netOpId** +
__required__|Network operator id|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|A representation of a netop.|<<_netops_netopid_get_response_200,Response 200>>
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_netops_netopid_get_response_400,Response 400>>
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**404**|Not Found. The requested resource could not be found but may be available again in the future. Subsequent requests by the client are permissible.|No Content
|===

[[_netops_netopid_get_response_200]]
**Response 200**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**EULA** +
__optional__|Contains MNO EULA (End User Licence Agreement) version. If this property is missing from response it means that MNO does not have it's own individual EULA, so default Mobile Connect Digital Identity EULA is used instead.|<<_netops_netopid_get_eula,EULA>>
|**acl** +
__optional__|List of MSISDN prefix patterns allowed to register with this MNO (BETA only).|< string > array
|**ape_preference** +
__optional__|APE settings for the MNO|<<_netops_netopid_get_ape_preference,ape_preference>>
|**brandname** +
__optional__|Contains brandname|string
|**cacheExpiry** +
__optional__|MCXG cache expiry time (in seconds) for profile data from this MNO.|integer
|**contactPerson** +
__optional__|Contact person in CU. May contain anything (a name, a user ID, e-mail address, MSISDN, etc.).|string
|**countryCode** +
__optional__|Network operator country calling code|number
|**featuredBlacklist** +
__optional__|List of featured applications not to be displayed when this operator is in use.|< string > array
|**hni** +
__optional__|NetOp's array of Home Network Identities (MCC + MNC combinations)|< string > array
|**id** +
__optional__|Network operator's ID +
**Pattern** : `"[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}"`|string
|**image** +
__optional__|Set of netop logo's|<<_netops_netopid_get_image,image>>
|**integrated** +
__optional__|Is Network operator part of Mobile Connect Accelerator network, possible values are HOST - is integrated, BETA - in BETA stage, INACTIVE - MNO is unnasigned, INTERNATIONAL - grandfather node.|string
|**iso** +
__optional__|Network operator ISO 3166 country code|string
|**label** +
__optional__|Network operator name|string
|**logo** +
__optional__|NetOp's logo (deprecated)|string
|**minimumUmpLoa** +
__optional__|Defines the LOA level to use on User Management Portal|string
|**mnp** +
__optional__|Network operator has mobile number portability|boolean
|**ndc** +
__optional__|List of national destination codes (NDC)|< string > array
|**policy** +
__optional__|Authentication policy containing prioritized list of LoA2 and LoA3 authentication methods.|<<_netops_netopid_get_policy,policy>>
|**promotedApps** +
__optional__|Promoted applications for network operator|< string > array
|**servingOperator** +
__optional__|Network operator serving name|string
|**shortRegistration** +
__optional__|Short Registration on this MNO is enabled when the value is true.|boolean
|**showcase** +
__optional__|Contains showcase header|string
|**smsGateway** +
__optional__|SMS gateway to be used by network operator|string
|**timezone** +
__optional__|Timezone of the MNO. The format must be in form Area/Location such as Europe/London.|string
|**trialEnd** +
__optional__|Trial period end time (ISO8601).|string
|**trialNotificationList** +
__optional__|Contacts that will be notified when trial is ending.|<<_netops_netopid_get_trialnotificationlist,trialNotificationList>>
|**trialStart** +
__optional__|Trial period beginning time (ISO8601).|string
|**ttl** +
__optional__|Network operator transaction TTL configuration|<<_netops_netopid_get_ttl,ttl>>
|**type** +
__optional__|Operator type (mcx or mckGlobal)|string
|**umpHost** +
__optional__|User Management Portal url|string
|**ussdGateway** +
__optional__|USSD gateway to be used by network operator|string
|**visible** +
__optional__|Is Network operator visible|boolean
|===

[[_netops_netopid_get_eula]]
**EULA**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**issuer** +
__optional__|MNO EULA issuer. If issuer is different than MNO ID, then EULA belongs to MNO host.|string
|**version** +
__optional__|MNO EULA version|string
|===

[[_netops_netopid_get_ape_preference]]
**ape_preference**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**loa2** +
__optional__|Preferred authentication method for LoA2|string
|**loa3** +
__optional__|Preferred authentication method for LoA3|string
|===

[[_netops_netopid_get_image]]
**image**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**120x120** +
__optional__|Image link|string
|**240x240** +
__optional__|Image link|string
|**500x500** +
__optional__|Image link|string
|**60x60** +
__optional__|Image link|string
|===

[[_netops_netopid_get_policy]]
**policy**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**loa2** +
__optional__|Prioritized list of LoA2 authentication methods.|< string > array
|**loa3** +
__optional__|Prioritized list of LoA3 authentication methods.|< <<_netops_netopid_get_policy_loa3,loa3>> > array
|===

[[_netops_netopid_get_policy_loa3]]
**loa3**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**primary** +
__optional__|Primary authentication method|string
|**secondary** +
__optional__|Secondary (confirmatory) authentication method|string
|===

[[_netops_netopid_get_trialnotificationlist]]
**trialNotificationList**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**email** +
__optional__|Email addresses.|< string > array
|**msisdn** +
__optional__|Phone number and name of the person.|< <<_netops_netopid_get_trialnotificationlist_msisdn,msisdn>> > array
|===

[[_netops_netopid_get_trialnotificationlist_msisdn]]
**msisdn**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**msisdn** +
__optional__|Phone number.|string
|**name** +
__optional__|Anything that describes the msisdn (first name, last name, …). It can be also empty.|string
|===

[[_netops_netopid_get_ttl]]
**ttl**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**authentication** +
__optional__|Authentication transaction|integer
|**authorization** +
__optional__|Authorization transaction|integer
|**bcAuthentication** +
__optional__|Backchannel authentication transaction|integer
|**bcAuthorization** +
__optional__|Backchannel authorization transaction|integer
|**registration** +
__optional__|Registration transaction|integer
|===

[[_netops_netopid_get_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request, no_route)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* Netops


==== Example HTTP response

===== Response 200
[source,json]
----
{
  "id" : "e406cd96-6e3d-48fe-8a97-5f1498f158d7",
  "label" : "Mobitel",
  "iso" : "si",
  "type" : "mcx",
  "countryCode" : 386,
  "servingOperator" : "Mobitel",
  "integrated" : "BETA",
  "hni" : [ "293.41", "293.64" ],
  "ndc" : [ "33" ],
  "mnp" : false,
  "smsGateway" : "IPX",
  "ussdGateway" : "mobitel_ussd",
  "promotedApps" : [ "00000000-de30-da1a-0000-db7b5bb0ac61" ],
  "featuredBlacklist" : [ "00000000-de30-da1a-0000-adb342bcc881" ],
  "showcase" : "Showcase header",
  "brandname" : "Commercial brandname",
  "EULA" : {
    "issuer" : "e406cd96-6e3d-48fe-8a97-5f1498f158d7",
    "version" : "1.0"
  },
  "policy" : {
    "loa2" : [ "password", "sms" ],
    "loa3" : [ {
      "primary" : "password",
      "secondary" : "sms"
    }, {
      "primary" : "password",
      "secondary" : "push"
    } ]
  },
  "acl" : [ "9355" ],
  "cacheExpiry" : 172800,
  "shortRegistration" : true,
  "trialNotificationList" : {
    "msisdn" : [ {
      "name" : "FirstName LastName1",
      "msisdn" : "+38640333771"
    }, {
      "name" : "FirstName LastName2",
      "msisdn" : "+38640333772"
    } ],
    "email" : [ "example1@example.org", "example2@example.org" ]
  },
  "timezone" : "Europe/London",
  "minimumUmpLoa" : "loa3",
  "ttl" : {
    "registration" : 300,
    "authentication" : 180,
    "bcAuthentication" : 180,
    "authorization" : 180,
    "bcAuthorization" : 180
  },
  "umpHost" : "https://www.eaxmpleportal.com"
}
----


===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_netops_netopid_delete]]
=== Remove a net. operator
....
DELETE /netops/{netOpId}
....


==== Description
Removes existing net. operator.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Path**|**netOpId** +
__required__|Network operator id|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**204**|On success 204 code is returned.|No Content
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_netops_netopid_delete_response_400,Response 400>>
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_netops_netopid_delete_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* Netops


==== Example HTTP response

===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_netops_netopid_application-whitelist_post]]
=== Add applications to MNOs applications whitelist
....
POST /netops/{netOpId}/application-whitelist
....


==== Description
Add applications to mobile network operator's applications whitelist.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Path**|**netOpId** +
__required__|Network operator id|string
|**Body**|**Add applications to MNOs applications whitelist request body** +
__optional__||< string > array
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**201**|Add applications to MNOs applications whitelist response body.|<<_netops_netopid_application-whitelist_post_response_201,Response 201>>
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_netops_netopid_application-whitelist_post_response_400,Response 400>>
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**404**|Not Found. The requested resource could not be found but may be available again in the future. Subsequent requests by the client are permissible.|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_netops_netopid_application-whitelist_post_response_201]]
**Response 201**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**added** +
__required__|Added application IDs|< string > array
|===

[[_netops_netopid_application-whitelist_post_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request, invalid_app_whitelist)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* Netops


==== Example HTTP response

===== Response 201
[source,json]
----
{
  "added" : [ "399d3214-8e55-4e85-bb42-f7d8abcc1fbf", "20066abe-04b6-4c9f-8704-5e41df3af6ce", "d9553ebe-a2fc-425f-b4a1-2be3f96915a1", "806dd531-cc65-4388-9ed3-ec7ca10be8e0" ]
}
----


===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_netops_netopid_application-whitelist_get]]
=== Retrieve MNO whitelisted applications
....
GET /netops/{netOpId}/application-whitelist
....


==== Description
Retrieves mobile network operator's whitelisted applications collection.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Path**|**netOpId** +
__required__|Network operator id|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|Collection of Applications.|< <<_netops_netopid_application-whitelist_get_response_200,Response 200>> > array
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_netops_netopid_application-whitelist_get_response_400,Response 400>>
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**404**|Not Found. The requested resource could not be found but may be available again in the future. Subsequent requests by the client are permissible.|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_netops_netopid_application-whitelist_get_response_200]]
**Response 200**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**author** +
__optional__|Author info|<<_netops_netopid_application-whitelist_get_author,author>>
|**createdAt** +
__optional__|Applications's create date.|string
|**description** +
__optional__|Applications's description.|string
|**id** +
__optional__|Application's id. +
**Pattern** : `"[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}"`|string
|**image** +
__optional__|Set of application image links with all available sizes.|<<_netops_netopid_application-whitelist_get_image,image>>
|**label** +
__optional__|Application's name.|string
|**privileges** +
__optional__|List of privileges application has.|<<_netops_netopid_application-whitelist_get_privileges,privileges>>
|**promotedImage** +
__optional__|Application's promoted image.|<<_netops_netopid_application-whitelist_get_promotedimage,promotedImage>>
|**publisher** +
__optional__|Publisher info.|<<_netops_netopid_application-whitelist_get_publisher,publisher>>
|**requireLoa** +
__optional__|Restricts which LoA is minimal for login (eg. loa2 or loa3). Empty string means that there is no restriction.|string
|**scope** +
__optional__|Array of scopes.|< string > array
|**scopeDescription** +
__optional__|Description, reasoning about requested scopes.|string
|**shortNames** +
__optional__|Array of application short names.|< string > array
|**url** +
__optional__|URL of application's website.|string
|===

[[_netops_netopid_application-whitelist_get_author]]
**author**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**email** +
__optional__|Author's E-mail.|string
|**name** +
__optional__|Author's name.|string
|===

[[_netops_netopid_application-whitelist_get_image]]
**image**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**120x120** +
__optional__|Image link|string
|**240x240** +
__optional__|Image link|string
|**45x45** +
__optional__|Image link|string
|**480x480** +
__optional__|Image link|string
|**60x60** +
__optional__|Image link|string
|**90x90** +
__optional__|Image link|string
|===

[[_netops_netopid_application-whitelist_get_privileges]]
**privileges**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**authorizator** +
__optional__|Defines if application is authorizator.|boolean
|**autogrant** +
__optional__|Defines if application scopes can be autogranted.|boolean
|**provision** +
__optional__|Defines if application has provisioner privileges.|boolean
|**provisionDescriptor** +
__optional__|Defines if application has descriptor provision privilege.|boolean
|**spUserProvision** +
__optional__|Defines if application has privileges to provision users.|boolean
|**system** +
__optional__|Defines if it is system application.|boolean
|**trusted** +
__optional__|Defines if application is trusted.|boolean
|**ump** +
__optional__|Defines if application acts as an User Management Portal|boolean
|===

[[_netops_netopid_application-whitelist_get_promotedimage]]
**promotedImage**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**282x174** +
__optional__|Image link|string
|===

[[_netops_netopid_application-whitelist_get_publisher]]
**publisher**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**email** +
__optional__|Publisher's E-mail|string
|**name** +
__optional__|Publisher's name|string
|===

[[_netops_netopid_application-whitelist_get_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* Netops


==== Example HTTP response

===== Response 200
[source,json]
----
[ {
  "id" : "dde9536c-2ca2-4fdc-a710-17ef0a58ede8",
  "label" : "FakeAnApp v2",
  "shortNames" : [ "FakeAnApp v2" ],
  "url" : "http://awesomeapp.example.com",
  "description" : "Be a description",
  "scopeDescription" : "Reasoning about the scopes that this app is requesting",
  "scope" : [ "scope.r", "scope.w" ],
  "images" : [ {
    "id" : "5a486dee-718d-4caf-8e10-c332024c4a32",
    "type" : "icon-64x64",
    "href" : "5a486dee-718d-4caf-8e10-c332024c4a32.png"
  } ],
  "author" : {
    "name" : "null",
    "email" : "author@etal.io"
  },
  "publisher" : {
    "name" : "Mr. Publisher",
    "email" : "author@etal.io"
  },
  "trusted" : true,
  "links" : [ {
    "rel" : "author",
    "href" : "https://api.mdi.ericsson.net/v1/profile/dde9536c-2ca2-4fdc-a710-17ef0a58ede8"
  }, {
    "rel" : "application",
    "href" : "https://api.mdi.ericsson.net/v1/application/dde9536c-2ca2-4fdc-a710-17ef0a58ede8"
  }, {
    "rel" : "keys",
    "href" : "https://api.mdi.ericsson.net/v1/application/dde9536c-2ca2-4fdc-a710-17ef0a58ede8/keys"
  } ],
  "platforms" : [ "web", "ios", "android" ],
  "category" : {
    "id" : "3cf1cd0e-ddf3-4bd3-85e8-45ad9b99621d",
    "name" : "Books"
  },
  "platformUri" : [ {
    "type" : "web",
    "url" : "http:://www.app.dev"
  } ],
  "featured" : true
} ]
----


===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_netops_netopid_application-whitelist_put]]
=== Update MNOs list of whitelisted applications
....
PUT /netops/{netOpId}/application-whitelist
....


==== Description
Updates applications from the mobile network operator's list of whitelisted applications.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Path**|**netOpId** +
__required__|Network operator id|string
|**Body**|**Update MNOs whitelisted applications request body** +
__optional__||< string > array
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**202**|Update MNOs whitelisted applications response body.|<<_netops_netopid_application-whitelist_put_response_202,Response 202>>
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_netops_netopid_application-whitelist_put_response_400,Response 400>>
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**404**|Not Found. The requested resource could not be found but may be available again in the future. Subsequent requests by the client are permissible.|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_netops_netopid_application-whitelist_put_response_202]]
**Response 202**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**whitelisted_apps** +
__required__|Whitelisted application IDs|< string > array
|===

[[_netops_netopid_application-whitelist_put_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request, invalid_app_whitelist)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* Netops


==== Example HTTP response

===== Response 202
[source,json]
----
{
  "whitelisted_apps" : [ "399d3214-8e55-4e85-bb42-f7d8abcc1fbf", "20066abe-04b6-4c9f-8704-5e41df3af6ce", "d9553ebe-a2fc-425f-b4a1-2be3f96915a1", "806dd531-cc65-4388-9ed3-ec7ca10be8e0" ]
}
----


===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_netops_netopid_application-whitelist_delete]]
=== Remove applications from MNOs applications whitelist
....
DELETE /netops/{netOpId}/application-whitelist
....


==== Description
Remove applications to mobile network operator's applications whitelist.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Path**|**netOpId** +
__required__|Network operator id|string
|**Body**|**Remove applications from MNOs applications whitelist request body** +
__optional__||< string > array
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**204**|On success 204 code is returned.|No Content
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_netops_netopid_application-whitelist_delete_response_400,Response 400>>
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**404**|Not Found. The requested resource could not be found but may be available again in the future. Subsequent requests by the client are permissible.|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_netops_netopid_application-whitelist_delete_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request, invalid_app_whitelist)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* Netops


==== Example HTTP response

===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_netops_netopid_application-whitelist-ids_get]]
=== Retrieve MNO whitelisted application IDs
....
GET /netops/{netOpId}/application-whitelist-ids
....


==== Description
Retrieves array of mobile network operator's whitelisted application IDs.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Path**|**netOpId** +
__required__|Network operator id|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|Retrieve MNO whitelisted application IDs response body.|< string > array
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_netops_netopid_application-whitelist-ids_get_response_400,Response 400>>
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**404**|Not Found. The requested resource could not be found but may be available again in the future. Subsequent requests by the client are permissible.|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_netops_netopid_application-whitelist-ids_get_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request, invalid_app_whitelist)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* Netops


==== Example HTTP response

===== Response 200
[source,json]
----
[ "399d3214-8e55-4e85-bb42-f7d8abcc1fbf", "20066abe-04b6-4c9f-8704-5e41df3af6ce", "d9553ebe-a2fc-425f-b4a1-2be3f96915a1", "806dd531-cc65-4388-9ed3-ec7ca10be8e0" ]
----


===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_netops_netopid_descriptors_post]]
=== Add descriptor link
....
POST /netops/{netOpId}/descriptors
....


==== Description
Adds authenticator's descriptor link to MNO.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Path**|**netOpId** +
__required__|Network operator id|string
|**Body**|**Descriptor request body** +
__optional__||<<_netops_netopid_descriptors_post_descriptor_request_body,Descriptor request body>>
|===

[[_netops_netopid_descriptors_post_descriptor_request_body]]
**Descriptor request body**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**url** +
__required__|Authenticator's descriptor url.|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**201**|Response with descriptor link data.|<<_netops_netopid_descriptors_post_response_201,Response 201>>
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_netops_netopid_descriptors_post_response_400,Response 400>>
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**404**|Not Found. The requested resource could not be found but may be available again in the future. Subsequent requests by the client are permissible.|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_netops_netopid_descriptors_post_response_201]]
**Response 201**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**createdAt** +
__required__|Time when authneticator was added to MNO.|string
|**descriptorId** +
__required__|Authenticator's descriptor id.|string
|**url** +
__required__|Authenticator's descriptor url.|string
|===

[[_netops_netopid_descriptors_post_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request, version_missing_parameter, text_missing_parameter, invalid_version, invalid_text)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* Netops


==== Example HTTP response

===== Response 201
[source,json]
----
{
  "descriptorId" : "sms",
  "url" : "http://localhost/v1-auth-sms",
  "createdAt" : "2016-09-09T10:35:59+0000"
}
----


===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_netops_netopid_descriptors_get]]
=== Get descriptor links collection
....
GET /netops/{netOpId}/descriptors
....


==== Description
Returns array of authenticator's descriptor link that are added to MNO.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Path**|**netOpId** +
__required__|Network operator id|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|Collection of descriptor link.|< <<_netops_netopid_descriptors_get_response_200,Response 200>> > array
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_netops_netopid_descriptors_get_response_400,Response 400>>
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**404**|Not Found. The requested resource could not be found but may be available again in the future. Subsequent requests by the client are permissible.|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_netops_netopid_descriptors_get_response_200]]
**Response 200**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**createdAt** +
__required__|Time when authneticator was added to MNO.|string
|**descriptorId** +
__required__|Authenticator's descriptor id.|string
|**url** +
__required__|Authenticator's descriptor url.|string
|===

[[_netops_netopid_descriptors_get_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* Netops


==== Example HTTP response

===== Response 200
[source,json]
----
[ {
  "descriptorId" : "sms",
  "url" : "http://localhost/v1-auth-sms",
  "createdAt" : "2016-09-09T10:35:59+0000"
}, {
  "descriptorId" : "password",
  "url" : "http://localhost/v1-auth-password",
  "createdAt" : "2016-09-09T10:30:17+0000"
} ]
----


===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_netops_netopid_descriptors_delete]]
=== Delete descriptor link
....
DELETE /netops/{netOpId}/descriptors
....


==== Description
Delete authenticator's descriptor link.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Path**|**netOpId** +
__required__|Network operator id|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**204**|On success 204 code is returned.|No Content
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_netops_netopid_descriptors_delete_response_400,Response 400>>
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**404**|Not Found. The requested resource could not be found but may be available again in the future. Subsequent requests by the client are permissible.|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_netops_netopid_descriptors_delete_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* Netops


==== Example HTTP response

===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_netops_netopid_descriptors_discover_get]]
=== Discover descriptors
....
GET /netops/{netOpId}/descriptors/discover
....


==== Description
Returns descriptor links of authenticators that are available on the network.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Path**|**netOpId** +
__required__|Network operator id|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|Array of descriptor links.|< string > array
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_netops_netopid_descriptors_discover_get_response_400,Response 400>>
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**404**|Not Found. The requested resource could not be found but may be available again in the future. Subsequent requests by the client are permissible.|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_netops_netopid_descriptors_discover_get_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* Netops


==== Example HTTP response

===== Response 200
[source,json]
----
[ "http://localhost/v1-auth-sms", "http://localhost/v1-auth-password" ]
----


===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_netops_netopid_descriptors_descriptorid_get]]
=== Get descriptor link
....
GET /netops/{netOpId}/descriptors/{descriptorId}
....


==== Description
Returns authenticator's descriptor link if it is added to MNO.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Path**|**descriptorId** +
__required__|Descriptor id|string
|**Path**|**netOpId** +
__required__|Network operator id|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|Response with descriptor link data.|<<_netops_netopid_descriptors_descriptorid_get_response_200,Response 200>>
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_netops_netopid_descriptors_descriptorid_get_response_400,Response 400>>
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**404**|Not Found. The requested resource could not be found but may be available again in the future. Subsequent requests by the client are permissible.|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_netops_netopid_descriptors_descriptorid_get_response_200]]
**Response 200**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**createdAt** +
__required__|Time when authneticator was added to MNO.|string
|**descriptorId** +
__required__|Authenticator's descriptor id.|string
|**url** +
__required__|Authenticator's descriptor url.|string
|===

[[_netops_netopid_descriptors_descriptorid_get_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* Netops


==== Example HTTP response

===== Response 200
[source,json]
----
{
  "descriptorId" : "sms",
  "url" : "http://localhost/v1-auth-sms",
  "createdAt" : "2016-09-09T10:35:59+0000"
}
----


===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_netops_netopid_eula_post]]
=== Add MNO EULA
....
POST /netops/{netOpId}/eula
....


==== Description
Creates mobile network operator's End User Licence Agreement. If content in request is not escaped it will be converted in backend.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Path**|**netOpId** +
__required__|Network operator id|string
|**Body**|**MNO EULA request body** +
__optional__||<<_netops_netopid_eula_post_mno_eula_request_body,MNO EULA request body>>
|===

[[_netops_netopid_eula_post_mno_eula_request_body]]
**MNO EULA request body**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**text** +
__required__|MNO EULA content|string
|**version** +
__required__|MNO EULA version|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**201**|MNO EULA response body.|<<_netops_netopid_eula_post_response_201,Response 201>>
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_netops_netopid_eula_post_response_400,Response 400>>
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**404**|Not Found. The requested resource could not be found but may be available again in the future. Subsequent requests by the client are permissible.|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_netops_netopid_eula_post_response_201]]
**Response 201**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**issuer** +
__required__|MNO EULA issuer. If issuer is different than MNO ID, then EULA belongs to MNO host.|string
|**text** +
__required__|MNO EULA content|string
|**version** +
__required__|MNO EULA version|string
|===

[[_netops_netopid_eula_post_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request, version_missing_parameter, text_missing_parameter, invalid_version, invalid_text)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* Netops


==== Example HTTP response

===== Response 201
[source,json]
----
{
  "issuer" : "e406cd96-6e3d-48fe-8a97-5f1498f158d7",
  "version" : "1.0",
  "text" : "&lt;html&gt;ANY PERSON OR ENTITY (“User”) DOWNLOADING, INSTALLING, OR OTHERWISE USING THE PUBLIC BETA OF THE MOBILE CONNECT MOBILE IDENTITY APPLICATION (the “App”) AND SERVICES (jointly called the “SERVICES”) AGREES TO AND IS BOUND BY THE FOLLOWING TERMS AND CONDITIONS (the “Agreement”). IF THE USER DOES NOT AGREE TO THE TERMS OF THIS AGREEMENT, THE USER IS NOT PERMITTED TO DOWNLOAD, INSTALL OR USE THE SERVICES.&lt;/html&gt;"
}
----


===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_netops_netopid_eula_get]]
=== Retrieve MNO EULA
....
GET /netops/{netOpId}/eula
....


==== Description
Retrieves mobile network operator's End User Licence Agreement. Content is in escaped HTML format.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Path**|**netOpId** +
__required__|Network operator id|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|MNO EULA response body.|<<_netops_netopid_eula_get_response_200,Response 200>>
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**404**|Not Found. The requested resource could not be found but may be available again in the future. Subsequent requests by the client are permissible.|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_netops_netopid_eula_get_response_200]]
**Response 200**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**issuer** +
__required__|MNO EULA issuer. If issuer is different than MNO ID, then EULA belongs to MNO host.|string
|**text** +
__required__|MNO EULA content|string
|**version** +
__required__|MNO EULA version|string
|===


==== Produces

* `application/json`


==== Tags

* Netops


==== Example HTTP response

===== Response 200
[source,json]
----
{
  "issuer" : "e406cd96-6e3d-48fe-8a97-5f1498f158d7",
  "version" : "1.0",
  "text" : "&lt;html&gt;ANY PERSON OR ENTITY (“User”) DOWNLOADING, INSTALLING, OR OTHERWISE USING THE PUBLIC BETA OF THE MOBILE CONNECT MOBILE IDENTITY APPLICATION (the “App”) AND SERVICES (jointly called the “SERVICES”) AGREES TO AND IS BOUND BY THE FOLLOWING TERMS AND CONDITIONS (the “Agreement”). IF THE USER DOES NOT AGREE TO THE TERMS OF THIS AGREEMENT, THE USER IS NOT PERMITTED TO DOWNLOAD, INSTALL OR USE THE SERVICES.&lt;/html&gt;"
}
----


[[_netops_netopid_eula_delete]]
=== Delete MNO EULA
....
DELETE /netops/{netOpId}/eula
....


==== Description
Removes existing MNO EULA.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Path**|**netOpId** +
__required__|Network operator id|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**204**|On success 204 code is returned.|No Content
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===


==== Produces

* `application/json`


==== Tags

* Netops


[[_netops_netopid_logo_post]]
=== Add network operator logo
....
POST /netops/{netOpId}/logo
....


==== Description
Adds logo to network operator.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Path**|**netOpId** +
__required__|Network operator id|string
|**Body**|**Body should be the multipart/form-data of the attachment with key named 'image'** +
__optional__||object
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**201**|Image logo used by netop API endpoint.|<<_netops_netopid_logo_post_response_201,Response 201>>
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_netops_netopid_logo_post_response_400,Response 400>>
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_netops_netopid_logo_post_response_201]]
**Response 201**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**id** +
__optional__|Netop logo UUID|string
|**type** +
__optional__|Type of image|string
|**urls** +
__optional__|Location of resized images|<<_netops_netopid_logo_post_urls,urls>>
|===

[[_netops_netopid_logo_post_urls]]
**urls**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**120x120** +
__optional__|URL of a resized image|string
|**240x240** +
__optional__|URL of a resized image|string
|**500x500** +
__optional__|URL of a resized image|string
|**60x60** +
__optional__|URL of a resized image|string
|===

[[_netops_netopid_logo_post_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request, image_missing, image_wrong_type, image_wrong_format, image_too_big, images_not_resized)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* Netops


==== Example HTTP response

===== Response 201
[source,json]
----
{
  "id" : "5a486dee-718d-4caf-8e10-c332024c4a32",
  "type" : "png",
  "url" : {
    "60x60" : "http://cloudfront.net/images/applications/5a486dee-718d-4caf-8e10-c332024c4a32-60x60.png",
    "120x120" : "http://cloudfront.net/images/applications/5a486dee-718d-4caf-8e10-c332024c4a32-120x120.png",
    "240x240" : "http://cloudfront.net/images/applications/5a486dee-718d-4caf-8e10-c332024c4a32-240x240.png",
    "500x500" : "http://cloudfront.net/images/applications/5a486dee-718d-4caf-8e10-c332024c4a32-500x500.png"
  }
}
----


===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_netops_netopid_logo_get]]
=== Returns logo URL
....
GET /netops/{netOpId}/logo
....


==== Description
Returns logo URL of specified size.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Path**|**netOpId** +
__required__|Network operator id|string
|**Query**|**size** +
__required__|Parameter for predefined image size. NOTE: Availible sizes are 45x45, 60x60, 120x120, 240x240 and 480x480.|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|Image logo used by netop API endpoint.|<<_netops_netopid_logo_get_response_200,Response 200>>
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_netops_netopid_logo_get_response_400,Response 400>>
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_netops_netopid_logo_get_response_200]]
**Response 200**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**id** +
__optional__|Netop logo UUID|string
|**type** +
__optional__|Type of image|string
|**urls** +
__optional__|Location of resized images|<<_netops_netopid_logo_get_urls,urls>>
|===

[[_netops_netopid_logo_get_urls]]
**urls**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**120x120** +
__optional__|URL of a resized image|string
|**240x240** +
__optional__|URL of a resized image|string
|**500x500** +
__optional__|URL of a resized image|string
|**60x60** +
__optional__|URL of a resized image|string
|===

[[_netops_netopid_logo_get_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request, image_size_missing_parameter, invalid_image_size)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* Netops


==== Example HTTP response

===== Response 200
[source,json]
----
{
  "id" : "5a486dee-718d-4caf-8e10-c332024c4a32",
  "type" : "png",
  "url" : {
    "60x60" : "http://cloudfront.net/images/applications/5a486dee-718d-4caf-8e10-c332024c4a32-60x60.png",
    "120x120" : "http://cloudfront.net/images/applications/5a486dee-718d-4caf-8e10-c332024c4a32-120x120.png",
    "240x240" : "http://cloudfront.net/images/applications/5a486dee-718d-4caf-8e10-c332024c4a32-240x240.png",
    "500x500" : "http://cloudfront.net/images/applications/5a486dee-718d-4caf-8e10-c332024c4a32-500x500.png"
  }
}
----


===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_netops_netopid_logo_delete]]
=== Remove network operator logo
....
DELETE /netops/{netOpId}/logo
....


==== Description
Removes network operator logo.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Path**|**netOpId** +
__required__|Network operator id|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**204**|On success 204 code is returned.|No Content
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_netops_netopid_logo_delete_response_400,Response 400>>
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_netops_netopid_logo_delete_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* Netops


==== Example HTTP response

===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_oauth2_oidc_ape_post]]
=== Authentication Policy Engine (APE) - HTTP POST
....
POST /oauth2/oidc/ape
....


==== Description
Endpoint for retrieving prioritized list of authentication methods. Priority list is set by user's mobile network operator. Only trusted applications with (at least) client_credentials access token can access this endpoint. If user with provided MSISDN does not exist, 404 Not Found is returned.


==== Parameters

[options="header", cols=".^2a,.^3a,.^4a"]
|===
|Type|Name|Schema
|**Body**|**Authentication request body (APE)** +
__optional__|<<_oauth2_oidc_ape_post_authentication_request_body_ape,Authentication request body (APE)>>
|===

[[_oauth2_oidc_ape_post_authentication_request_body_ape]]
**Authentication request body (APE)**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**client_id** +
__required__|Application's client id|string
|**msisdn** +
__required__|Phone number|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|APE priority list response.|<<_oauth2_oidc_ape_post_response_200,Response 200>>
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_oauth2_oidc_ape_post_response_400,Response 400>>
|**401**|Access denied. Authentication is required to proceed processing request.|No Content
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_oauth2_oidc_ape_post_response_200]]
**Response 200**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**loa2** +
__optional__|LOA2 (Level of assurance 2) priority list|< string > array
|**loa3** +
__optional__|LOA3 (Level of assurance 3) priority list|< <<_oauth2_oidc_ape_post_loa3,loa3>> > array
|===

[[_oauth2_oidc_ape_post_loa3]]
**loa3**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**primary** +
__optional__|Authentication method (currently only password is supported)|string
|**secondary** +
__optional__|Two-factor authentication method (can be one of following: ussd, push or sms). Null if primary is LoA3 (mepin).|string
|===

[[_oauth2_oidc_ape_post_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request, msisdn_missing_parameter, invalid_msisdn, profile_does_not_exist, inactive_operator, app_not_whitelisted, invalid_client_id, profile_suspended)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* OpenID_Connect


==== Example HTTP response

===== Response 200
[source,json]
----
{
  "loa2" : [ "ussd", "push", "password" ],
  "loa3" : [ {
    "primary" : "password",
    "secondary" : "ussd"
  }, {
    "primary" : "password",
    "secondary" : "sms"
  }, {
    "primary" : "password",
    "secondary" : "push"
  } ]
}
----


===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_oauth2_oidc_ape_get]]
=== Authentication Policy Engine (APE) - HTTP GET (Recommended)
....
GET /oauth2/oidc/ape
....


==== Description
Endpoint for retrieving prioritized list of authentication methods. Priority list is set by user's mobile network operator. Only trusted applications with (at least) client_credentials access token can access this endpoint. If user with provided MSISDN does not exist, 404 Not Found is returned.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Query**|**client_id** +
__required__|Application's client id|string
|**Query**|**msisdn** +
__required__|Contains MSISDN to return prioritized authentication policy list for.|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|APE priority list response.|<<_oauth2_oidc_ape_get_response_200,Response 200>>
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_oauth2_oidc_ape_get_response_400,Response 400>>
|**401**|Access denied. Authentication is required to proceed processing request.|No Content
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**404**|Not Found. The requested resource could not be found but may be available again in the future. Subsequent requests by the client are permissible.|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_oauth2_oidc_ape_get_response_200]]
**Response 200**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**loa2** +
__optional__|LOA2 (Level of assurance 2) priority list|< string > array
|**loa3** +
__optional__|LOA3 (Level of assurance 3) priority list|< <<_oauth2_oidc_ape_get_loa3,loa3>> > array
|===

[[_oauth2_oidc_ape_get_loa3]]
**loa3**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**primary** +
__optional__|Authentication method (currently only password is supported)|string
|**secondary** +
__optional__|Two-factor authentication method (can be one of following: ussd, push or sms). Null if primary is LoA3 (mepin).|string
|===

[[_oauth2_oidc_ape_get_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request, msisdn_missing_parameter, invalid_msisdn, profile_does_not_exist, inactive_operator, app_not_whitelisted, invalid_client_id, profile_suspended)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* OpenID_Connect


==== Example HTTP response

===== Response 200
[source,json]
----
{
  "loa2" : [ "ussd", "push", "password" ],
  "loa3" : [ {
    "primary" : "password",
    "secondary" : "ussd"
  }, {
    "primary" : "password",
    "secondary" : "sms"
  }, {
    "primary" : "password",
    "secondary" : "push"
  } ]
}
----


===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_oauth2_oidc_authenticate_confirm_code_code_post]]
=== Confirm authentication code
....
POST /oauth2/oidc/authenticate/confirm/code/{code}
....


==== Description
Endpoint for confirming authentication code. JSON payload may differ depending on authenticator.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Path**|**code** +
__required__|Public code|string
|**Body**|**Authentication code confirmation request** +
__optional__||<<_oauth2_oidc_authenticate_confirm_code_code_post_authentication_code_confirmation_request,Authentication code confirmation request>>
|===

[[_oauth2_oidc_authenticate_confirm_code_code_post_authentication_code_confirmation_request]]
**Authentication code confirmation request**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**authentication_data** +
__optional__|Holds details about authentication.|<<_oauth2_oidc_authenticate_confirm_code_code_post_authentication_data,authentication_data>>
|**transaction_id** +
__optional__|Transaction ID parameter of transaction to be confirmed.|string
|===

[[_oauth2_oidc_authenticate_confirm_code_code_post_authentication_data]]
**authentication_data**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**acr** +
__optional__|Authentication acr value.|string
|**amr** +
__optional__|Authentication amr values. Allowed values: "OK", "OTP", "BIOM", "DEV_PIN", "SIM_PIN", "UID_PWD", "HDR".|< string > array
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|Authentication response.|<<_oauth2_oidc_authenticate_confirm_code_code_post_response_200,Response 200>>
|**400**|Possible reasons for error: app could not be found, bad request, limit exeeds max, page does not exist, wrong page parameters,…|No Content
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**404**|Not Found. The requested resource could not be found but may be available again in the future. Subsequent requests by the client are permissible.|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_oauth2_oidc_authenticate_confirm_code_code_post_response_200]]
**Response 200**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**status** +
__optional__|Request status.|string
|===


==== Produces

* `application/json`


==== Tags

* OpenID_Connect


==== Example HTTP response

===== Response 200
[source,json]
----
{
  "status" : "OK"
}
----


[[_oauth2_oidc_authenticate_loa3_post]]
=== Authenticate user (LoA3)
....
POST /oauth2/oidc/authenticate/loa3
....


==== Description
Endpoint for two-factor authentication (LoA3 - Level of assurance 3). First factor is always password, but second factor must be one of the following authentication methods: USSD, push notification or SMS. Secondary authentication needs to be confirmed afterwards by using suitable confirmation endpoint. Process is completed by polling Authentication status endpoint.


==== Parameters

[options="header", cols=".^2a,.^3a,.^4a"]
|===
|Type|Name|Schema
|**Body**|**Authentication request body (LoA3)** +
__optional__|<<_oauth2_oidc_authenticate_loa3_post_authentication_request_body_loa3,Authentication request body (LoA3)>>
|===

[[_oauth2_oidc_authenticate_loa3_post_authentication_request_body_loa3]]
**Authentication request body (LoA3)**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**msisdn** +
__required__|Phone number|string
|**password** +
__optional__|User's password - required when primary method is password.|string
|**primary** +
__required__|Primary authentication factor. Must have one of following values: 'password', 'sms', 'ussd' or 'push'.|string
|**secondary** +
__required__|Secondary authentication factor. Must have one of following values: 'password', 'sms', 'ussd' or 'push'.|string
|**text** +
__optional__|USSD message text - required when secondary authentication method is USSD. To be used ONLY with USSD secondary authentication method.|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**201**|Authentication response|<<_oauth2_oidc_authenticate_loa3_post_response_201,Response 201>>
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_oauth2_oidc_authenticate_loa3_post_response_400,Response 400>>
|**401**|Access denied. Authentication is required to proceed processing request.|No Content
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**429**|Too Many Requests. The user or application has send too many requests in a given amount of time. If field 'blocked_for' is present in the response, it contains an information about how long the request will remain blocked (in seconds).|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_oauth2_oidc_authenticate_loa3_post_response_201]]
**Response 201**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**authenticator_data** +
__optional__|Holds additional data needed that might be needed for complete authentication depending on authenticator.|<<_oauth2_oidc_authenticate_loa3_post_authenticator_data,authenticator_data>>
|**expire_at** +
__required__|Holds authentication attempt validity time in ISO8601 format (UTC timezone).|string
|**status** +
__required__|Status|string
|===

[[_oauth2_oidc_authenticate_loa3_post_authenticator_data]]
**authenticator_data**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**push_status** +
__optional__|Indicates if push notification was sent successfully, values: "ok", "error". Used by Nok Nok authenticator.|string
|**qr_code** +
__optional__|Base 64 encoded QR code image. Used by Nok Nok authenticator as fallback mechanism to wake up Nok Nok App in case push notification was not delivered.|string
|===

[[_oauth2_oidc_authenticate_loa3_post_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request, msisdn_missing_parameter, invalid_msisdn, password_missing_parameter, invalid_password, primary_missing_parameter, secondary_missing_parameter, invalid_primary, invalid_secondary, text_missing_parameter, invalid_text, profile_does_not_exist, user_already_authenticated, invalid_authentication_method, inactive_operator, text_is_ussd_only, app_not_whitelisted, profile_suspended)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* OpenID_Connect


==== Example HTTP response

===== Response 201
[source,json]
----
{
  "status" : "OK",
  "expire_at" : "2016-02-26T14:39:34+0000",
  "authenticator_data" : null
}
----


===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_oauth2_oidc_authenticate_mepin_post]]
=== Authenticate user (MePIN LoA2)
....
POST /oauth2/oidc/authenticate/mepin
....


==== Description
Endpoint for authenticating user by using MePIN LoA2.


==== Parameters

[options="header", cols=".^2a,.^3a,.^4a"]
|===
|Type|Name|Schema
|**Body**|**Authentication request body** +
__optional__|<<_oauth2_oidc_authenticate_mepin_post_authentication_request_body,Authentication request body>>
|===

[[_oauth2_oidc_authenticate_mepin_post_authentication_request_body]]
**Authentication request body**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**msisdn** +
__required__|Phone number|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**201**|Authentication response|<<_oauth2_oidc_authenticate_mepin_post_response_201,Response 201>>
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_oauth2_oidc_authenticate_mepin_post_response_400,Response 400>>
|**401**|Access denied. Authentication is required to proceed processing request.|No Content
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**429**|Too Many Requests. The user or application has send too many requests in a given amount of time. If field 'blocked_for' is present in the response, it contains an information about how long the request will remain blocked (in seconds).|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_oauth2_oidc_authenticate_mepin_post_response_201]]
**Response 201**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**authenticator_data** +
__optional__|Holds additional data needed that might be needed for complete authentication depending on authenticator.|<<_oauth2_oidc_authenticate_mepin_post_authenticator_data,authenticator_data>>
|**expire_at** +
__required__|Holds authentication attempt validity time in ISO8601 format (UTC timezone).|string
|**status** +
__required__|Status|string
|===

[[_oauth2_oidc_authenticate_mepin_post_authenticator_data]]
**authenticator_data**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**push_status** +
__optional__|Indicates if push notification was sent successfully, values: "ok", "error". Used by Nok Nok authenticator.|string
|**qr_code** +
__optional__|Base 64 encoded QR code image. Used by Nok Nok authenticator as fallback mechanism to wake up Nok Nok App in case push notification was not delivered.|string
|===

[[_oauth2_oidc_authenticate_mepin_post_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request, msisdn_missing_parameter, invalid_msisdn, msisdn_mismatch, profile_does_not_exist, user_already_authenticated, invalid_authentication_method, inactive_operator, app_not_whitelisted, profile_suspended)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* OpenID_Connect


==== Example HTTP response

===== Response 201
[source,json]
----
{
  "status" : "OK",
  "expire_at" : "2016-02-26T14:39:34+0000",
  "authenticator_data" : null
}
----


===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_oauth2_oidc_authenticate_noknok_post]]
=== Authenticate user (Nok Nok LoA2 endpoint)
....
POST /oauth2/oidc/authenticate/noknok
....


==== Description
Endpoint for authenticating user with Nok Nok.


==== Parameters

[options="header", cols=".^2a,.^3a,.^4a"]
|===
|Type|Name|Schema
|**Body**|**Authentication request body** +
__optional__|<<_oauth2_oidc_authenticate_noknok_post_authentication_request_body,Authentication request body>>
|===

[[_oauth2_oidc_authenticate_noknok_post_authentication_request_body]]
**Authentication request body**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**msisdn** +
__required__|Phone number|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**201**|Authentication response|<<_oauth2_oidc_authenticate_noknok_post_response_201,Response 201>>
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_oauth2_oidc_authenticate_noknok_post_response_400,Response 400>>
|**401**|Access denied. Authentication is required to proceed processing request.|No Content
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**429**|Too Many Requests. The user or application has send too many requests in a given amount of time. If field 'blocked_for' is present in the response, it contains an information about how long the request will remain blocked (in seconds).|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_oauth2_oidc_authenticate_noknok_post_response_201]]
**Response 201**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**authenticator_data** +
__optional__|Holds additional data needed that might be needed for complete authentication depending on authenticator.|<<_oauth2_oidc_authenticate_noknok_post_authenticator_data,authenticator_data>>
|**expire_at** +
__required__|Holds authentication attempt validity time in ISO8601 format (UTC timezone).|string
|**status** +
__required__|Status|string
|===

[[_oauth2_oidc_authenticate_noknok_post_authenticator_data]]
**authenticator_data**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**push_status** +
__optional__|Indicates if push notification was sent successfully, values: "ok", "error". Used by Nok Nok authenticator.|string
|**qr_code** +
__optional__|Base 64 encoded QR code image. Used by Nok Nok authenticator as fallback mechanism to wake up Nok Nok App in case push notification was not delivered.|string
|===

[[_oauth2_oidc_authenticate_noknok_post_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request, msisdn_missing_parameter, invalid_msisdn, msisdn_mismatch, profile_does_not_exist, user_already_authenticated, invalid_authentication_method, inactive_operator, app_not_whitelisted, profile_suspended)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* OpenID_Connect


==== Example HTTP response

===== Response 201
[source,json]
----
{
  "status" : "OK",
  "expire_at" : "2016-02-26T13:47:32+0000",
  "authenticator_data" : {
    "push_status" : "ok",
    "qr_code" : "iVBORw0KGgoAAAANSUhEUgAAASwAAAEsAQAAAABRBrPYAAADZUlEQVR42u2aQa6cQAxEjViw5Ah9E7gY0iBxMbhJH4HlLEZ0qsogMclfRFm1FUZK9OfzWDjtssvuWPmbz24P9mAP9h9i2cza/G72trwbfO02/MRPfuOJhcDwZylv65fSlf2VjS8QsPfoT+vHSOSylU/q1v6Tyga2O2yy94B/gkAYCYT2MZzYnEBMKRp2MLe2MpuRRaSvEgdTvtnQF2ZZe2Uec/DPtKwUc9UPCPf3v34oDlVi/rGG59S4/ufUUUE/lOg6scyHVD3LLtkBqcYiwKS75VvFGJTRg1CCoX2wAPOI9hfT71UiYPk97i37HiqVMm/ET3oLxawEwQZvGiaM57RQRpO6YAqBsWtnZFlbEK5RN5DM4YKfLQLGOot8O1hx0boTqjDUwsB5WDkChv6NLEN8n7OMrTiijicGh3g7rJoxa66H6N+F4b7U2FmPSw6BXSWWZqRn9WLS6YUVX0Ng3i/Wni4KNQsxU/BMOkgmxcBGzRE0tEPvblD55q+GwCB4eg72vZU2dizSvyajTwxMc4TmOo1HUAucFY8NvfCVY2DyIdS6kVAPgXjcIaYQWLfyYGgJER874K4S9rFvK1UxVhgVzomCPwvXxEihlsPaGJidUxD6N30hsSQrhfn0NhnVjHWHe/I5XfMpte5Jt+QImC8INJCaai+DRBFos+pxBEzGg+HOppHO2/nmpXiyCFi+Osc5GbmB0riq8SgAxq3SpP2SHnIqnVS4kkn/AbDM3GK5wkzNrWsnKzVSQd+9vl6seK9eKfiB4a40gmwpjDRFwLJmauTWuZyRv2VDpEO8D3f1YjQenm/6DWuWJaWftk8hsMuJz5wjFlVcc7Wcdyr1Y9KIrwoKtc7W3a3XiUXAmG8zXZSuIbT71jWEffuQujFOFP6rzQdrqb4t0n8ErKj2yrzyiLJ6oW8OttuupmpMXZvtwzc07gZ1L3cu8qvHMA+xzjb7uai0RrejumL52tXUi/lV9LX0W7TS1/UWne6tUVaMZR8mZEE2TkaD21hdas0WArvFd95ncWYle9wn8ZoxF4pfyXFLYH5/Mtu3Xakd41znN+zqHL681AItDsYvydfecoN6y+75VjPm8hi9XGkt8GGQMCNHv5QQmG7Yz3sgbmg4FHn/3spvxaFS7Pk/Zg/2YA/2r9gvh94akdUzhOQAAAAASUVORK5CYII="
  }
}
----


===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_oauth2_oidc_authenticate_password_confirm_post]]
=== Validate password (LOA3)
....
POST /oauth2/oidc/authenticate/password/confirm
....


==== Description
Endpoint for validating password as a secondary authentication LOA3 method, when authentication code does not get delivered to the trusted application by any other authentication method and as a consequence no other authentication confirmation endpoint validates password. Authentication status endpoint needs to be called upon successful password validation to complete the authentication request.


==== Parameters

[options="header", cols=".^2a,.^3a,.^4a"]
|===
|Type|Name|Schema
|**Body**|**LOA3 password validation** +
__optional__|<<_oauth2_oidc_authenticate_password_confirm_post_loa3_password_validation,LOA3 password validation>>
|===

[[_oauth2_oidc_authenticate_password_confirm_post_loa3_password_validation]]
**LOA3 password validation**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**password** +
__required__|Contains password to be validated.|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**201**|Authentication response.|<<_oauth2_oidc_authenticate_password_confirm_post_response_201,Response 201>>
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_oauth2_oidc_authenticate_password_confirm_post_response_400,Response 400>>
|**401**|Access denied. Authentication is required to proceed processing request.|No Content
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_oauth2_oidc_authenticate_password_confirm_post_response_201]]
**Response 201**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**status** +
__optional__|Request status.|string
|===

[[_oauth2_oidc_authenticate_password_confirm_post_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request, password_missing_parameter, invalid_password)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* OpenID_Connect


==== Example HTTP response

===== Response 201
[source,json]
----
{
  "status" : "OK"
}
----


===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_oauth2_oidc_authenticate_push_post]]
=== Authenticate user (push notification)
....
POST /oauth2/oidc/authenticate/push
....


==== Description
Endpoint for authenticating user by using push notifications.


==== Parameters

[options="header", cols=".^2a,.^3a,.^4a"]
|===
|Type|Name|Schema
|**Body**|**Authentication request body** +
__optional__|<<_oauth2_oidc_authenticate_push_post_authentication_request_body,Authentication request body>>
|===

[[_oauth2_oidc_authenticate_push_post_authentication_request_body]]
**Authentication request body**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**msisdn** +
__required__|Phone number|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**201**|Authentication response|<<_oauth2_oidc_authenticate_push_post_response_201,Response 201>>
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_oauth2_oidc_authenticate_push_post_response_400,Response 400>>
|**401**|Access denied. Authentication is required to proceed processing request.|No Content
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**429**|Too Many Requests. The user or application has send too many requests in a given amount of time. If field 'blocked_for' is present in the response, it contains an information about how long the request will remain blocked (in seconds).|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_oauth2_oidc_authenticate_push_post_response_201]]
**Response 201**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**authenticator_data** +
__optional__|Holds additional data needed that might be needed for complete authentication depending on authenticator.|<<_oauth2_oidc_authenticate_push_post_authenticator_data,authenticator_data>>
|**expire_at** +
__required__|Holds authentication attempt validity time in ISO8601 format (UTC timezone).|string
|**status** +
__required__|Status|string
|===

[[_oauth2_oidc_authenticate_push_post_authenticator_data]]
**authenticator_data**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**push_status** +
__optional__|Indicates if push notification was sent successfully, values: "ok", "error". Used by Nok Nok authenticator.|string
|**qr_code** +
__optional__|Base 64 encoded QR code image. Used by Nok Nok authenticator as fallback mechanism to wake up Nok Nok App in case push notification was not delivered.|string
|===

[[_oauth2_oidc_authenticate_push_post_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request, msisdn_missing_parameter, invalid_msisdn, msisdn_mismatch, profile_does_not_exist, user_already_authenticated, invalid_authentication_method, inactive_operator, app_not_whitelisted, profile_suspended)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* OpenID_Connect


==== Example HTTP response

===== Response 201
[source,json]
----
{
  "status" : "OK",
  "expire_at" : "2016-02-26T14:39:34+0000",
  "authenticator_data" : null
}
----


===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_oauth2_oidc_authenticate_push_confirm_post]]
=== Confirm user authentication (push notification)
....
POST /oauth2/oidc/authenticate/push/confirm
....


==== Description
Endpoint for confirming received push notification code.


==== Parameters

[options="header", cols=".^2a,.^3a,.^4a"]
|===
|Type|Name|Schema
|**Body**|**Push notification code confirmation request** +
__optional__|<<_oauth2_oidc_authenticate_push_confirm_post_push_notification_code_confirmation_request,Push notification code confirmation request>>
|===

[[_oauth2_oidc_authenticate_push_confirm_post_push_notification_code_confirmation_request]]
**Push notification code confirmation request**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**code** +
__required__|Contains received push notification code.|string
|**deviceId** +
__required__|Contains ID of device where request is being sent from.|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|Authentication response|<<_oauth2_oidc_authenticate_push_confirm_post_response_200,Response 200>>
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_oauth2_oidc_authenticate_push_confirm_post_response_400,Response 400>>
|**401**|Access denied. Authentication is required to proceed processing request.|No Content
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**429**|Too Many Requests. The user or application has send too many requests in a given amount of time. If field 'blocked_for' is present in the response, it contains an information about how long the request will remain blocked (in seconds).|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_oauth2_oidc_authenticate_push_confirm_post_response_200]]
**Response 200**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**status** +
__required__|Status|string
|===

[[_oauth2_oidc_authenticate_push_confirm_post_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request, deviceId_missing_parameter, code_missing_parameter, invalid_code)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* OpenID_Connect


==== Example HTTP response

===== Response 200
[source,json]
----
{
  "status" : "OK"
}
----


===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_oauth2_oidc_authenticate_sms_post]]
=== Authenticate user (SMS)
....
POST /oauth2/oidc/authenticate/sms
....


==== Description
Endpoint for authenticating user by using SMS. SMS message contains authentication confirmation URL.


==== Parameters

[options="header", cols=".^2a,.^3a,.^4a"]
|===
|Type|Name|Schema
|**Body**|**Authentication request body** +
__optional__|<<_oauth2_oidc_authenticate_sms_post_authentication_request_body,Authentication request body>>
|===

[[_oauth2_oidc_authenticate_sms_post_authentication_request_body]]
**Authentication request body**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**msisdn** +
__required__|Phone number|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**201**|Authentication response|<<_oauth2_oidc_authenticate_sms_post_response_201,Response 201>>
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_oauth2_oidc_authenticate_sms_post_response_400,Response 400>>
|**401**|Access denied. Authentication is required to proceed processing request.|No Content
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**429**|Too Many Requests. The user or application has send too many requests in a given amount of time. If field 'blocked_for' is present in the response, it contains an information about how long the request will remain blocked (in seconds).|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_oauth2_oidc_authenticate_sms_post_response_201]]
**Response 201**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**authenticator_data** +
__optional__|Holds additional data needed that might be needed for complete authentication depending on authenticator.|<<_oauth2_oidc_authenticate_sms_post_authenticator_data,authenticator_data>>
|**expire_at** +
__required__|Holds authentication attempt validity time in ISO8601 format (UTC timezone).|string
|**status** +
__required__|Status|string
|===

[[_oauth2_oidc_authenticate_sms_post_authenticator_data]]
**authenticator_data**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**push_status** +
__optional__|Indicates if push notification was sent successfully, values: "ok", "error". Used by Nok Nok authenticator.|string
|**qr_code** +
__optional__|Base 64 encoded QR code image. Used by Nok Nok authenticator as fallback mechanism to wake up Nok Nok App in case push notification was not delivered.|string
|===

[[_oauth2_oidc_authenticate_sms_post_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request, msisdn_missing_parameter, invalid_msisdn, msisdn_mismatch, profile_does_not_exist, user_already_authenticated, invalid_authentication_method, inactive_operator, app_not_whitelisted, profile_suspended)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* OpenID_Connect


==== Example HTTP response

===== Response 201
[source,json]
----
{
  "status" : "OK",
  "expire_at" : "2016-02-26T14:39:34+0000",
  "authenticator_data" : null
}
----


===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_oauth2_oidc_authenticate_sms_confirm_post]]
=== Confirm user authentication (SMS with URL)
....
POST /oauth2/oidc/authenticate/sms/confirm
....


==== Description
Endpoint for confirming code, received by SMS.


==== Parameters

[options="header", cols=".^2a,.^3a,.^4a"]
|===
|Type|Name|Schema
|**Body**|**SMS (URL) code confirmation request** +
__optional__|<<_oauth2_oidc_authenticate_sms_confirm_post_sms_url_code_confirmation_request,SMS (URL) code confirmation request>>
|===

[[_oauth2_oidc_authenticate_sms_confirm_post_sms_url_code_confirmation_request]]
**SMS (URL) code confirmation request**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**code** +
__required__|Contains code received from SMS (URL).|string
|**password** +
__optional__|Contains profile password - must be included only if password is secondary authentication method.|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|Authentication response|<<_oauth2_oidc_authenticate_sms_confirm_post_response_200,Response 200>>
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_oauth2_oidc_authenticate_sms_confirm_post_response_400,Response 400>>
|**401**|Access denied. Authentication is required to proceed processing request.|No Content
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_oauth2_oidc_authenticate_sms_confirm_post_response_200]]
**Response 200**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**status** +
__required__|Status|string
|===

[[_oauth2_oidc_authenticate_sms_confirm_post_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request, code_missing_parameter, invalid_code, password_missing_parameter, profile_does_not_exist, invalid_password)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* OpenID_Connect


==== Example HTTP response

===== Response 200
[source,json]
----
{
  "status" : "OK"
}
----


===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_oauth2_oidc_authenticate_status_get]]
=== Validate password (LOA3)
....
GET /oauth2/oidc/authenticate/status
....


==== Description
Endpoint for polling current authentication status. Status can be either 'non-authenticated', 'pending' or 'authenticated'.


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|Authentication status response.|<<_oauth2_oidc_authenticate_status_get_response_200,Response 200>>
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_oauth2_oidc_authenticate_status_get_response_400,Response 400>>
|**401**|Access denied. Authentication is required to proceed processing request.|No Content
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_oauth2_oidc_authenticate_status_get_response_200]]
**Response 200**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**required_action** +
__optional__|Action required for authentication to be completed. If field contains value 'wait', no action is required.|string
|**status** +
__required__|Status: 'non-authenticated', 'pending' or 'authenticated'|string
|===

[[_oauth2_oidc_authenticate_status_get_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (app_not_whitelisted)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* OpenID_Connect


==== Example HTTP response

===== Response 200
[source,json]
----
{
  "status" : "authenticated",
  "required_action" : "wait"
}
----


===== Response 400
[source,json]
----
{
  "error" : "app_not_whitelisted",
  "error_description" : "Operator does not permit using Mobile Connect for this application."
}
----


[[_oauth2_oidc_authenticate_ussd_post]]
=== Authenticate user (USSD)
....
POST /oauth2/oidc/authenticate/ussd
....


==== Description
Endpoint for authenticating user by using USSD message.


==== Parameters

[options="header", cols=".^2a,.^3a,.^4a"]
|===
|Type|Name|Schema
|**Body**|**Authentication request body** +
__optional__|<<_oauth2_oidc_authenticate_ussd_post_authentication_request_body,Authentication request body>>
|===

[[_oauth2_oidc_authenticate_ussd_post_authentication_request_body]]
**Authentication request body**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**msisdn** +
__required__|Phone number|string
|**text** +
__required__|USSD message text|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**201**|Authentication response|<<_oauth2_oidc_authenticate_ussd_post_response_201,Response 201>>
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_oauth2_oidc_authenticate_ussd_post_response_400,Response 400>>
|**401**|Access denied. Authentication is required to proceed processing request.|No Content
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**429**|Too Many Requests. The user or application has send too many requests in a given amount of time. If field 'blocked_for' is present in the response, it contains an information about how long the request will remain blocked (in seconds).|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_oauth2_oidc_authenticate_ussd_post_response_201]]
**Response 201**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**authenticator_data** +
__optional__|Holds additional data needed that might be needed for complete authentication depending on authenticator.|<<_oauth2_oidc_authenticate_ussd_post_authenticator_data,authenticator_data>>
|**expire_at** +
__required__|Holds authentication attempt validity time in ISO8601 format (UTC timezone).|string
|**status** +
__required__|Status|string
|===

[[_oauth2_oidc_authenticate_ussd_post_authenticator_data]]
**authenticator_data**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**push_status** +
__optional__|Indicates if push notification was sent successfully, values: "ok", "error". Used by Nok Nok authenticator.|string
|**qr_code** +
__optional__|Base 64 encoded QR code image. Used by Nok Nok authenticator as fallback mechanism to wake up Nok Nok App in case push notification was not delivered.|string
|===

[[_oauth2_oidc_authenticate_ussd_post_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request, msisdn_missing_parameter, invalid_msisdn, msisdn_mismatch, text_missing_parameter, invalid_text, profile_does_not_exist, user_already_authenticated, invalid_authentication_method, inactive_operator, app_not_whitelisted, profile_suspended)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* OpenID_Connect


==== Example HTTP response

===== Response 201
[source,json]
----
{
  "status" : "OK",
  "expire_at" : "2016-02-26T14:39:34+0000",
  "authenticator_data" : null
}
----


===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_oauth2_oidc_authenticate_ussd_confirm_post]]
=== Confirm user authentication (USSD)
....
POST /oauth2/oidc/authenticate/ussd/confirm
....


==== Description
Endpoint for confirming authentication by code received via USSD. (Locked to USSD provider's IP).


==== Parameters

[options="header", cols=".^2a,.^3a,.^4a"]
|===
|Type|Name|Schema
|**Body**|**USSD confirmation request** +
__optional__|<<_oauth2_oidc_authenticate_ussd_confirm_post_ussd_confirmation_request,USSD confirmation request>>
|===

[[_oauth2_oidc_authenticate_ussd_confirm_post_ussd_confirmation_request]]
**USSD confirmation request**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**code** +
__required__|Contains received CSRF code.|string
|**msisdn** +
__required__|Contains user's MSISDN.|string
|**response** +
__required__|Contains user's response - 'true' if user accepted authentication and 'false' if it was rejected.|boolean
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|Authentication response|<<_oauth2_oidc_authenticate_ussd_confirm_post_response_200,Response 200>>
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_oauth2_oidc_authenticate_ussd_confirm_post_response_400,Response 400>>
|**401**|Access denied. Authentication is required to proceed processing request.|No Content
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_oauth2_oidc_authenticate_ussd_confirm_post_response_200]]
**Response 200**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**status** +
__required__|Status|string
|===

[[_oauth2_oidc_authenticate_ussd_confirm_post_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request, msisdn_missing_parameter, invalid_msisdn, code_missing_parameter, invalid_code, response_missing_parameterinvalid_response, invalid_response)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* OpenID_Connect


==== Example HTTP response

===== Response 200
[source,json]
----
{
  "status" : "OK"
}
----


===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_oauth2_oidc_authorize_post]]
=== (Authorize - Authorization Code, Authorize implicit flow - Token Id Token, Authorize implicit flow - Id Token)
....
POST /oauth2/oidc/authorize
....


==== Description
(Authorize-Authorization Code): Authorization endpoint for OAuth2 requests, if 'openid' scope requested 'id_token' is included to token. (Authorize implicit flow-Token Id Token): OpenID Connect supported implicit flow for retreiving access token and id_token. Refresh token is not returned in this flow. (Authorize implicit flow-Id Token): OpenID Connect endpoint to get 'id_token', access_token is not returned in this flow.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Body**|**Request authorization code** +
__optional__|Request authorization bodys. NOTE: only one json obejct will be returned and not arrayof objects. Array shows all possible request bodys. Field 'nonce' is required for (Authorize implicit flow-Token Id Token) and (Authorize implicit flow-Id Token).|<<_oauth2_oidc_authorize_post_request_authorization_code,Request authorization code>>
|===

[[_oauth2_oidc_authorize_post_request_authorization_code]]
**Request authorization code**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**bundle_id** +
__optional__|REQUIRED or package_name: iOS identificator.|string
|**client_id** +
__required__|Contains application key.|string
|**nonce** +
__required__|String value used to associate a Client session with an ID Token, and to mitigate replay attacks.|string
|**redirect_uri** +
__optional__|The location URI where the user should be returned to after they approve access for your app. URI must be pre-pregistered with the client (can have multiple). If no redirect_uri is provided, client redirect uri is used.|string
|**requestCode** +
__optional__|Verification code for MT SMS flow. This is the code user gets from the SMS sent from Mobile Connect Accelerator.|string
|**responseCode** +
__optional__|Verification code for MO SMS flow. This is the code sent via SMS from user to Mobile Connect Accelerator.|string
|**response_type** +
__required__|Set to code indicating that an authorization code will be returned to the application after the user approves the authorization request|string
|**scope** +
__optional__|The data your application is requesting access to. Send a space (' ') separated list of scope types. If you do not specify the scope, 'profile.anonymous.r' is used. 'openid' scope must be used to specify openid connect request.|string
|**state** +
__required__|A unique value used by your application in order to prevent cross-site request forgery (CSRF) attacks on your implementation. Value will be truncated to 255 characters and returned as is in the response.|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|Authorize implicit flow responses are shown below in array. NOTE: Only one Json object is returned (not array). Name of key indicates response example.|No Content
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string. OAuth2 error will be provided in JSON response. Additional information can be found at http://tools.ietf.org/html/rfc6749#section-5.2|<<_oauth2_oidc_authorize_post_response_400,Response 400>>
|**401**|Access denied. Authentication is required to proceed processing request.|No Content
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_oauth2_oidc_authorize_post_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (invalid_request, invalid_client, invalid_grant, unauthorized_client, unsupported_grant_type, invalid_scope, profile_expired, redirect_uri_mismatch, consent_required, eula_not_accepted, app_not_whitelisted)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* OpenID_Connect


==== Example HTTP response

===== Response 400
[source,json]
----
{
  "error" : "invalid_request",
  "error_description" : "The request is missing a required parameter, includes an unsupported parameter value (other than grant type), repeats a parameter, includes multiple credentials, utilizes more than one mechanism for authenticating the client, or is otherwise malformed."
}
----


[[_oauth2_oidc_bc-authorize_post]]
=== Backchannel authorization request
....
POST /oauth2/oidc/bc-authorize
....


==== Description
Authorization endpoint for server-based requests.


==== Parameters

[options="header", cols=".^2a,.^3a,.^4a"]
|===
|Type|Name|Schema
|**Body**|**Request backchannel authorization** +
__optional__|<<_oauth2_oidc_bc-authorize_post_request_backchannel_authorization,Request backchannel authorization>>
|===

[[_oauth2_oidc_bc-authorize_post_request_backchannel_authorization]]
**Request backchannel authorization**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**acr_values** +
__required__|Space-separated string that specifies the Authentication Context Class Reference which translates to LoA level. Valid values in MDI: 3 (LoA3), 2 (LoA2).|string
|**binding_message** +
__optional__|Client provided plain text, 'reference or ID' to interlock consumption device and authorization device for a better user experience. The message will be displayed on consumption device and mobile device.|string
|**client_name** +
__optional__|A short name to identify RP/Client application and must be displayed on authentication/ authorisation device.|string
|**client_notification_token** +
__optional__|It is a unique id provided by the RP that will be used by the OpenID Provider as a bearer token to authenticate the callback request to send the tokens to the RP (Notification Mode).|string
|**context** +
__optional__|A Transaction/action based message displayed on authorization device and must provide by RP/client.|string
|**login_hint** +
__required__|An indication to the ID GW/Authorization Server on what ID to use for login. The recommendation is that the value matches the value used in Discovery. The login_hint can contain the MSISDN, encrypted MSISDN or PCR. The format should be as MSISDN:<Value>, ENCR_MSISDN:<Value> and PCR:<value>|string
|**nonce** +
__optional__|String value used to associate a Client session with an ID Token, and to mitigate replay attacks.|string
|**scope** +
__required__|Space delimited list of requested OIDC Mobile Connect Profile scopes.|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|Backchannel authorization response.|<<_oauth2_oidc_bc-authorize_post_response_200,Response 200>>
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string. OAuth2 error will be provided in JSON response. Additional information can be found at http://tools.ietf.org/html/rfc6749#section-5.2 .|<<_oauth2_oidc_bc-authorize_post_response_400,Response 400>>
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_oauth2_oidc_bc-authorize_post_response_200]]
**Response 200**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**auth_req_id** +
__required__|It is a unique id to identify the authentication request (transaction) made by the RP. The auth_req_id will be included too in the token notification when Notification mode is used to allow the RP correlate the authentication request and the received tokens.|string
|**expires_in** +
__required__|Expiration time of the 'auth_req_id' in seconds since the auth_request was received.|integer
|===

[[_oauth2_oidc_bc-authorize_post_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (invalid_request, invalid_client, invalid_grant, unauthorized_client, unsupported_grant_type, invalid_scope, profile_expired, redirect_uri_mismatch, consent_required, eula_not_accepted, app_not_whitelisted)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* OpenID_Connect


==== Example HTTP response

===== Response 200
[source,json]
----
{
  "auth_req_id" : "73523d40-d5c2-11e6-a0d7-bfe9287ef782",
  "expires_in" : 3600
}
----


===== Response 400
[source,json]
----
{
  "error" : "invalid_request",
  "error_description" : "The request is missing a required parameter, includes an unsupported parameter value (other than grant type), repeats a parameter, includes multiple credentials, utilizes more than one mechanism for authenticating the client, or is otherwise malformed."
}
----


[[_oauth2_oidc_certs_get]]
=== Signature Keys
....
GET /oauth2/oidc/certs
....


==== Description
Endpoint for retrieving public key for signature verification.


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|Collection of JWKs.|<<_oauth2_oidc_certs_get_response_200,Response 200>>
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_oauth2_oidc_certs_get_response_200]]
**Response 200**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**keys** +
__optional__|Array of keys|< <<_oauth2_oidc_certs_get_keys,keys>> > array
|===

[[_oauth2_oidc_certs_get_keys]]
**keys**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**alg** +
__optional__|Algorithm|string
|**exp** +
__optional__|Base64url encoded exponent value|string
|**kid** +
__optional__|Key ID|string
|**mod** +
__optional__|Base64url encoded modulus value|string
|**use** +
__optional__|Public Key Use|string
|===


==== Produces

* `application/json`


==== Tags

* OpenID_Connect


==== Example HTTP response

===== Response 200
[source,json]
----
{
  "keys" : [ {
    "kid" : "f31ae35369ef5cdf63a88a6d6fc614b7",
    "alg" : "RSA",
    "use" : "sig",
    "mod" : "wGj9ofu6sJPDnRxa6ZvmSpKPdgCBNp3hYr6KMVIzEGvcHKRXxCDjks8fC9uOVjb7Wv1C1Bjzqvd0whkw6E93NNXi9i6QFm2tcvTwuFoGKBbNi7574RO8veF988nwateKUDmLwvNgc78QFgChSUPYSM6lJLTE31FfzWCcVJsaLks",
    "exp" : "AQAB"
  }, {
    "kid" : "1c52dd3aed239b40e06944a9c16c9942",
    "alg" : "RSA",
    "use" : "sig",
    "mod" : "spuCQv8qLJtjY4dauV2TYbdyZA4rpKH3ajC1gTXR3YNP3UPpEN5NWaR93EbrXqsac6zlz_SDzyuKKOKTTSqV1YGnMdjFEFjMp5yowCsDl3nXKtcxLvVmMN8esOnyOeeFM20ttCg6YoXF2O7AmqW7wgbzs-P_43EjO6iWr98UbjE",
    "exp" : "AQAB"
  } ]
}
----


[[_oauth2_oidc_mc_confirm_post]]
=== Confirm user authentication (SMS with URL) - Mobile Connect
....
POST /oauth2/oidc/mc/confirm
....


==== Description
Endpoint for confirming code, received by SMS.


==== Parameters

[options="header", cols=".^2a,.^3a,.^4a"]
|===
|Type|Name|Schema
|**Body**|**SMS (URL) code confirmation request** +
__optional__|<<_oauth2_oidc_mc_confirm_post_sms_url_code_confirmation_request,SMS (URL) code confirmation request>>
|===

[[_oauth2_oidc_mc_confirm_post_sms_url_code_confirmation_request]]
**SMS (URL) code confirmation request**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**code** +
__required__|Contains code received from SMS (URL).|string
|**password** +
__optional__|Contains profile password - must be included only if password is secondary authentication method.|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**201**|Authentication response.|<<_oauth2_oidc_mc_confirm_post_response_201,Response 201>>
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_oauth2_oidc_mc_confirm_post_response_400,Response 400>>
|**401**|Access denied. Authentication is required to proceed processing request.|No Content
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_oauth2_oidc_mc_confirm_post_response_201]]
**Response 201**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**authenticator_data** +
__optional__|Holds additional data needed that might be needed for complete authentication depending on authenticator.|<<_oauth2_oidc_mc_confirm_post_authenticator_data,authenticator_data>>
|**expire_at** +
__required__|Holds authentication attempt validity time in ISO8601 format (UTC timezone).|string
|**status** +
__required__|Status|string
|===

[[_oauth2_oidc_mc_confirm_post_authenticator_data]]
**authenticator_data**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**push_status** +
__optional__|Indicates if push notification was sent successfully, values: "ok", "error". Used by Nok Nok authenticator.|string
|**qr_code** +
__optional__|Base 64 encoded QR code image. Used by Nok Nok authenticator as fallback mechanism to wake up Nok Nok App in case push notification was not delivered.|string
|===

[[_oauth2_oidc_mc_confirm_post_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request, code_missing_parameter, invalid_code, password_missing_parameter, invalid_password)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* OpenID_Connect


==== Example HTTP response

===== Response 201
[source,json]
----
{
  "status" : "OK",
  "expire_at" : "2016-02-26T14:39:34+0000",
  "authenticator_data" : null
}
----


===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_oauth2_oidc_mc_register_sms_post]]
=== Authenticate user (SMS) for on the fly registration - Mobile Connect
....
POST /oauth2/oidc/mc/register/sms
....


==== Description
Endpoint for authenticating user by using SMS. SMS message contains authentication confirmation URL.


==== Parameters

[options="header", cols=".^2a,.^3a,.^4a"]
|===
|Type|Name|Schema
|**Body**|**Authentication request body for MC registration** +
__optional__|<<_oauth2_oidc_mc_register_sms_post_authentication_request_body_for_mc_registration,Authentication request body for MC registration>>
|===

[[_oauth2_oidc_mc_register_sms_post_authentication_request_body_for_mc_registration]]
**Authentication request body for MC registration**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**msisdn** +
__required__|Phone number|string
|**operator** +
__required__|Mobile operator ID +
**Pattern** : `"[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}"`|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**201**|Authentication response.|<<_oauth2_oidc_mc_register_sms_post_response_201,Response 201>>
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_oauth2_oidc_mc_register_sms_post_response_400,Response 400>>
|**401**|Access denied. Authentication is required to proceed processing request.|No Content
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_oauth2_oidc_mc_register_sms_post_response_201]]
**Response 201**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**authenticator_data** +
__optional__|Holds additional data needed that might be needed for complete authentication depending on authenticator.|<<_oauth2_oidc_mc_register_sms_post_authenticator_data,authenticator_data>>
|**expire_at** +
__required__|Holds authentication attempt validity time in ISO8601 format (UTC timezone).|string
|**status** +
__required__|Status|string
|===

[[_oauth2_oidc_mc_register_sms_post_authenticator_data]]
**authenticator_data**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**push_status** +
__optional__|Indicates if push notification was sent successfully, values: "ok", "error". Used by Nok Nok authenticator.|string
|**qr_code** +
__optional__|Base 64 encoded QR code image. Used by Nok Nok authenticator as fallback mechanism to wake up Nok Nok App in case push notification was not delivered.|string
|===

[[_oauth2_oidc_mc_register_sms_post_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request, msisdn_missing_parameter, operator_missing_parameter, invalid_msisdn, msisdn_mismatch, profile_already_exists, user_already_authenticated, invalid_authentication_method)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* OpenID_Connect


==== Example HTTP response

===== Response 201
[source,json]
----
{
  "status" : "OK",
  "expire_at" : "2016-02-26T14:39:34+0000",
  "authenticator_data" : null
}
----


===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_oauth2_oidc_mc_status_get]]
=== Check authentication status - Mobile Connect
....
GET /oauth2/oidc/mc/status
....


==== Description
Endpoint for polling current authentication status. Status can be either 'non-authenticated', 'pending' or 'authenticated'.


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|Authentication status response.|<<_oauth2_oidc_mc_status_get_response_200,Response 200>>
|**401**|Access denied. Authentication is required to proceed processing request.|No Content
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_oauth2_oidc_mc_status_get_response_200]]
**Response 200**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**required_action** +
__optional__|Action required for authentication to be completed. If field contains value 'wait', no action is required.|string
|**status** +
__required__|Status: 'non-authenticated', 'pending' or 'authenticated'|string
|===


==== Produces

* `application/json`


==== Tags

* OpenID_Connect


==== Example HTTP response

===== Response 200
[source,json]
----
{
  "status" : "authenticated",
  "required_action" : "wait"
}
----


[[_oauth2_oidc_premiuminfo_get]]
=== Premiuminfo
....
GET /oauth2/oidc/premiuminfo
....


==== Description
Premiuminfo endpoint for retrieving premium information about the user.


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|User claims.|<<_oauth2_oidc_premiuminfo_get_response_200,Response 200>>
|**400**|OIDC server responded with false when validating request. OAuth2 error will be provided in JSON response.|No Content
|**401**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_oauth2_oidc_premiuminfo_get_response_200]]
**Response 200**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**email** +
__optional__|User's email.|string
|**email_verified** +
__optional__|Is email verified.|boolean
|**name** +
__optional__|User's name.|string
|**phone_number** +
__optional__|User's msisdn.|string
|**phone_number_verified** +
__optional__|Is msisdn verified.|boolean
|**sub** +
__optional__|Unique End-User identifier.|string
|**update_at** +
__optional__|Last modification date of profile.|string
|===


==== Produces

* `application/json`


==== Tags

* OpenID_Connect


==== Example HTTP response

===== Response 200
[source,json]
----
{
  "sub" : "5f90512d-972d-4def-bf90-9ef0ef2e5d2d",
  "email" : "test@test.com"
}
----


[[_oauth2_oidc_revoke_post]]
=== Token - Revoke
....
POST /oauth2/oidc/revoke
....


==== Description
Revoke access or refresh token.


==== Parameters

[options="header", cols=".^2a,.^3a,.^4a"]
|===
|Type|Name|Schema
|**Body**|**Revoke token request parameters** +
__optional__|<<_oauth2_oidc_revoke_post_revoke_token_request_parameters,Revoke token request parameters>>
|===

[[_oauth2_oidc_revoke_post_revoke_token_request_parameters]]
**Revoke token request parameters**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**token** +
__required__|The issued token to revoke.|string
|**token_type_hint** +
__required__|A hint about the type of the token submitted for revocation, can be one of access_token or refresh_token.|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**204**|On success 204 code is returned.|No Content
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_oauth2_oidc_revoke_post_response_400,Response 400>>
|**401**|Access denied. Authentication is required to proceed processing request.|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_oauth2_oidc_revoke_post_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (invalid_token)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/x-www-form-urlencoded`


==== Tags

* OpenID_Connect


==== Example HTTP response

===== Response 400
[source,json]
----
{
  "error" : "invalid_token",
  "error_description" : "Missing parameter token or parameter token_type_hint or token_type_hint is not array. Partameter token_type_hint must contain refresh_token and access_token."
}
----


[[_oauth2_oidc_token_post]]
=== (Token - Authorization Code, Token - Client Credentials, Token - Password, Token - Refresh token)
....
POST /oauth2/oidc/token
....


==== Description
(Token-Authorization Code): Retrieve tokens from authorization code with client credentials. 'id_token' is returned if requested with 'openid' scope. (Token-Client Credentials): Client credentials grant type for retreiving TRUSTED access token. Basic Authorization is required. (Token-Password): Password grant type for retrieiving TRUSTED profile related access token. (Token-Refresh token): Refresh Token grant type for retreiving access token by sending refresh token in exchange.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Body**|**Request token using authorization code** +
__optional__|Request token bodys. NOTE: only one json obejct will be returned and not arrayof objects. Array shows all possible request bodys. All field required combinations are shown in the example requests.|<<_oauth2_oidc_token_post_request_token_using_authorization_code,Request token using authorization code>>
|===

[[_oauth2_oidc_token_post_request_token_using_authorization_code]]
**Request token using authorization code**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**code** +
__required__|Authorization code obtained in process.|string
|**for_client_id** +
__optional__|Client Id of authenticating application|string
|**grant_type** +
__required__|Grant type that is performed, authorization_code|string
|**password** +
__required__|User's password|string
|**redirect_uri** +
__required__|Redirect URI must be included if it was included in the Authorization step. The values of both must be the same for the request to succeed.|string
|**refresh_token** +
__required__|The refresh token issued with the expired access token.|string
|**username** +
__required__|User's msisdn|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|Token endpoint response.|<<_oauth2_oidc_token_post_response_200,Response 200>>
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_oauth2_oidc_token_post_response_400,Response 400>>
|**401**|Access denied. Authentication is required to proceed processing request.|No Content
|**429**|Too Many Requests. The user or application has send too many requests in a given amount of time. If field 'blocked_for' is present in the response, it contains an information about how long the request will remain blocked (in seconds).|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_oauth2_oidc_token_post_response_200]]
**Response 200**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**access_token** +
__optional__|Issued access token for client.|string
|**expires_in** +
__optional__|Token expiration time in seconds.|number
|**id_token** +
__optional__|OPTIONAL: Base64 encoded JWT token, returned if requested with openid scope.|string
|**refresh_token** +
__optional__|OPTIONAL: Not returned in 'client_credentials' grant type.|string
|**scope** +
__optional__|Scopes requested to the process and linked to access token.|string
|**token_type** +
__optional__|Token type, 'Bearer'.|string
|===

[[_oauth2_oidc_token_post_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request, invalid_request, invalid_client, invalid_grant, unauthorized_client, unsupported_grant_type, invalid_scope, profile_expired, redirect_uri_mismatch, invalid_authentication_method, inactive_operator, app_not_whitelisted, profile_suspended)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/x-www-form-urlencoded`


==== Tags

* OpenID_Connect


==== Example HTTP response

===== Response 200
[source,json]
----
{
  "access_token" : "30a60cea8f2c7382c79245434b01c0d068261695",
  "refresh_token" : "7bb1ab3cc45b29f70c9d51a1e305877733c3ce3b",
  "expires_in" : 3600,
  "token_type" : "Bearer",
  "scope" : "profile.anonymous.r",
  "id_token" : "eyJ0eXAiOiJKV1MiLCJhbGciOiJSUzI1NiIsImtpZCI6IjE0Yzc2NmJmZTk5YTk2MGNjMDljZDY3Yzk1YTBjMTA3In0.eyJpc3MiOiJodHRwOi8vZXRhbGlvLmxvY2FsIiwic3ViIjoiN2QwYjRlZmEtODlhMC00OTA1LWI0YjctYzY4NjJkZjU3MTVmIiwiYXVkIjoiMDM0M2FkODQ1MThjMTI2ZDM5NjI3MDg1NzJiODE0MTkiLCJpYXQiOjE0MDc4NDQ0NzIsImV4cCI6MTQwNzg0ODA3MiwiYXV0aF90aW1lIjoxNDA3ODQ0NDcyLCJub25jZSI6Im4tMFM2X1d6QTJNaiJ9.lWiupcIQuXgnsihRH29U9sgCqxVIRZagRg6V3D1l2gn4pfvwr7HeDAj6jpN1G-tr3bR5aZPfAjpMndcajPwrQ_ysfqHRUf5u3BgxdV-PGxZ3Zu-9ON3Khtn97psdLqtz1zAxmnNNylu58rcMnwo0nZwXRB42773vWpuTGpx9Pmc"
}
----


===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_oauth2_oidc_userinfo_get]]
=== Userinfo
....
GET /oauth2/oidc/userinfo
....


==== Description
Userinfo endpoint for retrieving openid connect claims about the user. Claims are requested through scopes, it is a resource protected endpoint.


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|User claims.|<<_oauth2_oidc_userinfo_get_response_200,Response 200>>
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string. OAuth2 error will be provided in JSON response. Additional information can be found at http://tools.ietf.org/html/rfc6749#section-5.2|<<_oauth2_oidc_userinfo_get_response_400,Response 400>>
|**401**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_oauth2_oidc_userinfo_get_response_200]]
**Response 200**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**email** +
__optional__|User's email.|string
|**email_verified** +
__optional__|Is email verified.|boolean
|**name** +
__optional__|User's name.|string
|**phone_number** +
__optional__|User's msisdn.|string
|**phone_number_verified** +
__optional__|Is msisdn verified.|boolean
|**sub** +
__optional__|Unique End-User identifier.|string
|**update_at** +
__optional__|Last modification date of profile.|string
|===

[[_oauth2_oidc_userinfo_get_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (insufficient_scope)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* OpenID_Connect


==== Example HTTP response

===== Response 200
[source,json]
----
{
  "sub" : "b5733564-e0f4-4381-939b-3a0104a36ff5",
  "name" : "Kevin S. Archey",
  "email" : "kevinsarchey@dayrep.com",
  "email_verified" : true,
  "phone_number" : "+38640777777",
  "phone_number_verified" : true
}
----


===== Response 400
[source,json]
----
{
  "error" : "insufficient_scope",
  "error_description" : "The request requires higher privileges than provided by the access token."
}
----


[[_profile_get]]
=== Retrieve profile by MSISDN
....
GET /profile
....


==== Description
Returns requested profile.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Query**|**msisdn** +
__required__|Parameter for retrieving profile with provided MSISDN.|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|Authenticator linking status endpoint response.|< <<_profile_get_response_200,Response 200>> > array
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_profile_get_response_400,Response 400>>
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**404**|Not Found. The requested resource could not be found but may be available again in the future. Subsequent requests by the client are permissible.|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_profile_get_response_200]]
**Response 200**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**agreements** +
__optional__|Contains profile agreements. Available in scope.|<<_profile_get_agreements,agreements>>
|**apePreference** +
__optional__|List of prefered and available authenticators.|<<_profile_get_apepreference,apePreference>>
|**device** +
__optional__|Push notification device info|<<_profile_get_device,device>>
|**email** +
__optional__|Contains current address, possible pending address and verification status. Available in profile.email.r scope.|<<_profile_get_email,email>>
|**id** +
__optional__|Profile's ID. Available in profile.basic.r scope. +
**Pattern** : `"[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}"`|string
|**image** +
__optional__|Set of profile image links with all available sizes. Available in profile.basic.r scope.|<<_profile_get_image,image>>
|**links** +
__optional__|Profile related links|< <<_profile_get_links,links>> > array
|**msisdn** +
__optional__|Contains country code, national number (nn) and phone number. Available in profile.msisdn.r scope.|<<_profile_get_msisdn,msisdn>>
|**name** +
__optional__|Contains profile's name. Available in profile.basic.r scope.|string
|**operator** +
__optional__|Operator ID. Available in profile.operator.r scope. +
**Pattern** : `"[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}"`|string
|**status** +
__optional__|Profile account statuses|<<_profile_get_status,status>>
|===

[[_profile_get_agreements]]
**agreements**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**EULA** +
__optional__|Is Mobile Connect Accelerator End User Licence Agreement accepted|boolean
|**developerEULA** +
__optional__|Is developer EULA accepted|boolean
|===

[[_profile_get_apepreference]]
**apePreference**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**available** +
__optional__|Prefered authentication methods|<<_profile_get_apepreference_available,available>>
|**prefered** +
__optional__|Prefered authentication methods|<<_profile_get_apepreference_prefered,prefered>>
|===

[[_profile_get_apepreference_available]]
**available**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**loa2** +
__optional__|available LoA2 authentication methods.|< string > array
|**loa3** +
__optional__|available LoA2 authentication methods.|< string > array
|===

[[_profile_get_apepreference_prefered]]
**prefered**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**loa2** +
__optional__|Prefered LoA2 authentication method.|string
|**loa3** +
__optional__|Prefered LoA3 authentication method.|string
|===

[[_profile_get_device]]
**device**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**ID** +
__optional__|Push notification device registration ID|string
|**type** +
__optional__|Push notification device type: 'android' or 'ios'|string
|===

[[_profile_get_email]]
**email**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**address** +
__optional__|Currently used E-mail address|string
|**pending** +
__optional__|E-mail address waiting to be verified|string
|**verified** +
__optional__||boolean
|===

[[_profile_get_image]]
**image**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**120x120** +
__optional__|Image link|string
|**240x240** +
__optional__|Image link|string
|**60x60** +
__optional__|Image link|string
|**960x960** +
__optional__|Image link|string
|===

[[_profile_get_links]]
**links**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**href** +
__optional__|Link path|string
|**rel** +
__optional__|Link relation|string
|===

[[_profile_get_msisdn]]
**msisdn**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**country_code** +
__optional__|number
|**nn** +
__optional__|string
|**number** +
__optional__|string
|===

[[_profile_get_status]]
**status**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**suspended** +
__optional__|True if profile is currently suspended.|boolean
|===

[[_profile_get_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request, msisdn_missing_parameter)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* Profile


==== Example HTTP response

===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_profile_delete]]
=== Delete profile
....
DELETE /profile
....


==== Description
Instantly removes profile without TTL.


==== Parameters

[options="header", cols=".^2a,.^3a,.^4a"]
|===
|Type|Name|Schema
|**Body**|**Delete profile request body** +
__optional__|<<_profile_delete_delete_profile_request_body,Delete profile request body>>
|===

[[_profile_delete_delete_profile_request_body]]
**Delete profile request body**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**msisdn** +
__required__|User's MSISDN|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**204**|On success 204 code is returned.|No Content
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_profile_delete_response_400,Response 400>>
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**404**|Not Found. The requested resource could not be found but may be available again in the future. Subsequent requests by the client are permissible.|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_profile_delete_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* Profile


==== Example HTTP response

===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_profile_claim_post]]
=== Claim existing profile
....
POST /profile/claim
....


==== Description
Claim existing profile with new msisdn, e-mail code and SMS verification code.


==== Parameters

[options="header", cols=".^2a,.^3a,.^4a"]
|===
|Type|Name|Schema
|**Body**|**Msisdn claim request body** +
__optional__|<<_profile_claim_post_msisdn_claim_request_body,Msisdn claim request body>>
|===

[[_profile_claim_post_msisdn_claim_request_body]]
**Msisdn claim request body**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**emailCode** +
__required__|Contains e-mail verification code.|string
|**msisdn** +
__required__|User's new MSISDN|string
|**password** +
__required__|Contains old profile password.|string
|**requestCode** +
__optional__|Verification code for MT SMS flow. This is the code user gets from the SMS sent from Mobile Connect Accelerator.|string
|**responseCode** +
__optional__|Verification code for MO SMS flow. This is the code sent via SMS from user to Mobile Connect Accelerator.|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**201**|Msisdn claim response body|<<_profile_claim_post_response_201,Response 201>>
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_profile_claim_post_response_400,Response 400>>
|**401**|Access denied. Authentication is required to proceed processing request.|No Content
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_profile_claim_post_response_201]]
**Response 201**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**status** +
__required__|Status|string
|===

[[_profile_claim_post_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* Profile


==== Example HTTP response

===== Response 201
[source,json]
----
{
  "status" : "OK"
}
----


===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_profile_delete_ttl_get]]
=== Profile delete TTL
....
GET /profile/delete/ttl
....


==== Description
Get profile deletion time in hours.


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|TTL for profile deletion.|<<_profile_delete_ttl_get_response_200,Response 200>>
|**401**|Access denied. Authentication is required to proceed processing request.|No Content
|**405**|Method not allowed. A request method is not supported for the requested resource.|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_profile_delete_ttl_get_response_200]]
**Response 200**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**ttl** +
__required__|TTL for profile deletion in hours.|integer
|===


==== Produces

* `application/json`


==== Tags

* Profile


==== Example HTTP response

===== Response 200
[source,json]
----
{
  "ttl" : 24
}
----


[[_profile_loa3upgrade_code_get]]
=== GET missing LoA3 Upgrade
....
GET /profile/loa3upgrade/{code}
....


==== Description
Returns required and optional fields for profile upgrade to LoA3.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Path**|**code** +
__required__|Public code|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**202**|Returns required and optional fields that can be upgraded to make profile compatible with applications that require loa3 login with password.|<<_profile_loa3upgrade_code_get_response_202,Response 202>>
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_profile_loa3upgrade_code_get_response_400,Response 400>>
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_profile_loa3upgrade_code_get_response_202]]
**Response 202**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**optional** +
__optional__|Holds a list of optional fields that can be set during upgrade the profile to loa3|< string > array
|**required** +
__optional__|Holds a list of mandatory fields to upgrade the profile to loa3|< string > array
|===

[[_profile_loa3upgrade_code_get_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request, invalid_code)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* Profile


==== Example HTTP response

===== Response 202
[source,json]
----
{
  "required" : [ "password", "email" ],
  "optional" : [ "name" ]
}
----


===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_profile_loa3upgrade_code_put]]
=== Upgrade profile to LoA3
....
PUT /profile/loa3upgrade/{code}
....


==== Description
Upgrade profile from LoA2 to LoA3.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Path**|**code** +
__required__|Public code|string
|**Body**|**Upgrading the profile to LoA3** +
__optional__||<<_profile_loa3upgrade_code_put_upgrading_the_profile_to_loa3,Upgrading the profile to LoA3>>
|===

[[_profile_loa3upgrade_code_put_upgrading_the_profile_to_loa3]]
**Upgrading the profile to LoA3**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**email** +
__optional__|Contains new profile e-mail address.|string
|**family_name** +
__optional__|Contains Profile last name.|string
|**given_name** +
__optional__|Contains Profile first name.|string
|**middle_name** +
__optional__|Contains Profile middle name.|string
|**password** +
__optional__|Contains Profile password.|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**202**|Returns status whether the operation succeeded or not.|<<_profile_loa3upgrade_code_put_response_202,Response 202>>
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_profile_loa3upgrade_code_put_response_400,Response 400>>
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_profile_loa3upgrade_code_put_response_202]]
**Response 202**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**status** +
__optional__|Reports status of the executed action. On success the value should be OK|string
|===

[[_profile_loa3upgrade_code_put_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request, invalid_code, password_too_short, password_missing_parameter, email_missing_parameter, invalid_email_format)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* Profile


==== Example HTTP response

===== Response 202
[source,json]
----
{
  "status" : "OK"
}
----


===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_profile_me_get]]
=== Retrieve personal profile
....
GET /profile/me
....


==== Description
Returns current profile.


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|Retrieve personal profile response body.|< <<_profile_me_get_response_200,Response 200>> > array
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_profile_me_get_response_400,Response 400>>
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_profile_me_get_response_200]]
**Response 200**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**agreements** +
__optional__|Contains profile agreements. Available in scope.|<<_profile_me_get_agreements,agreements>>
|**apePreference** +
__optional__|List of prefered and available authenticators.|<<_profile_me_get_apepreference,apePreference>>
|**device** +
__optional__|Push notification device info|<<_profile_me_get_device,device>>
|**email** +
__optional__|Contains current address, possible pending address and verification status. Available in profile.email.r scope.|<<_profile_me_get_email,email>>
|**id** +
__optional__|Profile's ID. Available in profile.basic.r scope. +
**Pattern** : `"[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}"`|string
|**image** +
__optional__|Set of profile image links with all available sizes. Available in profile.basic.r scope.|<<_profile_me_get_image,image>>
|**links** +
__optional__|Profile related links|< <<_profile_me_get_links,links>> > array
|**msisdn** +
__optional__|Contains country code, national number (nn) and phone number. Available in profile.msisdn.r scope.|<<_profile_me_get_msisdn,msisdn>>
|**name** +
__optional__|Contains profile's name. Available in profile.basic.r scope.|string
|**operator** +
__optional__|Operator ID. Available in profile.operator.r scope. +
**Pattern** : `"[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}"`|string
|**status** +
__optional__|Profile account statuses|<<_profile_me_get_status,status>>
|===

[[_profile_me_get_agreements]]
**agreements**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**EULA** +
__optional__|Is Mobile Connect Accelerator End User Licence Agreement accepted|boolean
|**developerEULA** +
__optional__|Is developer EULA accepted|boolean
|===

[[_profile_me_get_apepreference]]
**apePreference**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**available** +
__optional__|Prefered authentication methods|<<_profile_me_get_apepreference_available,available>>
|**prefered** +
__optional__|Prefered authentication methods|<<_profile_me_get_apepreference_prefered,prefered>>
|===

[[_profile_me_get_apepreference_available]]
**available**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**loa2** +
__optional__|available LoA2 authentication methods.|< string > array
|**loa3** +
__optional__|available LoA2 authentication methods.|< string > array
|===

[[_profile_me_get_apepreference_prefered]]
**prefered**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**loa2** +
__optional__|Prefered LoA2 authentication method.|string
|**loa3** +
__optional__|Prefered LoA3 authentication method.|string
|===

[[_profile_me_get_device]]
**device**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**ID** +
__optional__|Push notification device registration ID|string
|**type** +
__optional__|Push notification device type: 'android' or 'ios'|string
|===

[[_profile_me_get_email]]
**email**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**address** +
__optional__|Currently used E-mail address|string
|**pending** +
__optional__|E-mail address waiting to be verified|string
|**verified** +
__optional__||boolean
|===

[[_profile_me_get_image]]
**image**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**120x120** +
__optional__|Image link|string
|**240x240** +
__optional__|Image link|string
|**60x60** +
__optional__|Image link|string
|**960x960** +
__optional__|Image link|string
|===

[[_profile_me_get_links]]
**links**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**href** +
__optional__|Link path|string
|**rel** +
__optional__|Link relation|string
|===

[[_profile_me_get_msisdn]]
**msisdn**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**country_code** +
__optional__|number
|**nn** +
__optional__|string
|**number** +
__optional__|string
|===

[[_profile_me_get_status]]
**status**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**suspended** +
__optional__|True if profile is currently suspended.|boolean
|===

[[_profile_me_get_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* Profile


==== Example HTTP response

===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_profile_password-reset_post]]
=== Send E-mail with forgotten password reset link
....
POST /profile/password-reset
....


==== Description
Sends E-mail with forgotten password reset link to user.


==== Parameters

[options="header", cols=".^2a,.^3a,.^4a"]
|===
|Type|Name|Schema
|**Body**|**Reset password for given msisdn** +
__optional__|<<_profile_password-reset_post_reset_password_for_given_msisdn,Reset password for given msisdn>>
|===

[[_profile_password-reset_post_reset_password_for_given_msisdn]]
**Reset password for given msisdn**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**msisdn** +
__required__|Contains phone number|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|On success 200 code is returned.|No Content
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_profile_password-reset_post_response_400,Response 400>>
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_profile_password-reset_post_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request, msisdn_missing_parameter, profile_does_not_exist, email_is_not_verifired)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* Profile


==== Example HTTP response

===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_profile_password_reset_get]]
=== Profile password reset
....
GET /profile/password/reset
....


==== Description
Resets password for profile. Parameter emailCode is required.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Query**|**emailCode** +
__required__|E-mail code used for reseting password.|string
|**Query**|**msisdn** +
__optional__|Contains phone number.|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|Reset password for given msisdn.|<<_profile_password_reset_get_response_200,Response 200>>
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_profile_password_reset_get_response_400,Response 400>>
|**401**|Access denied. Authentication is required to proceed processing request.|No Content
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_profile_password_reset_get_response_200]]
**Response 200**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**msisdn** +
__optional__|Contains phone number|string
|===

[[_profile_password_reset_get_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request, emailCode_missing_parameter, invalid_email_code)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* Profile


==== Example HTTP response

===== Response 200
[source,json]
----
{
  "msisdn" : "+38651000000"
}
----


===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_profile_password_reset_put]]
=== Change profile password
....
PUT /profile/password/reset
....


==== Description
Changes profile password.


==== Parameters

[options="header", cols=".^2a,.^3a,.^4a"]
|===
|Type|Name|Schema
|**Body**|**Reset password for given msisdn** +
__optional__|<<_profile_password_reset_put_reset_password_for_given_msisdn,Reset password for given msisdn>>
|===

[[_profile_password_reset_put_reset_password_for_given_msisdn]]
**Reset password for given msisdn**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**emailCode** +
__required__|Contains email code|string
|**fresh** +
__required__|Contains new profile password|string
|**requestCode** +
__required__|Verification code for MT SMS flow. This is the code user gets from the SMS sent from Mobile Connect Accelerator.|string
|**responseCode** +
__optional__|Verification code for MO SMS flow. This is the code sent via SMS from user to Mobile Connect Accelerator.|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**202**|Status message on successful update.|<<_profile_password_reset_put_response_202,Response 202>>
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_profile_password_reset_put_response_400,Response 400>>
|**401**|Access denied. Authentication is required to proceed processing request.|No Content
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_profile_password_reset_put_response_202]]
**Response 202**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**status** +
__optional__|string
|===

[[_profile_password_reset_put_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request, emailCode_missing_parameter, invalid_email_code, fresh_missing_parameter, password_too_short, requestCode_missing_parameter, responseCode_missing_parameter, invalid_code, requestCode_and_responseCode_parameters)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* Profile


==== Example HTTP response

===== Response 202
[source,json]
----
{
  "status" : "OK"
}
----


===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_profile_profileid_get]]
=== Retrieve profile by ID
....
GET /profile/{profileId}
....


==== Description
Returns requested profile.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Path**|**profileId** +
__required__|Profile id|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|Authenticator linking status endpoint response.|< <<_profile_profileid_get_response_200,Response 200>> > array
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_profile_profileid_get_response_400,Response 400>>
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_profile_profileid_get_response_200]]
**Response 200**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**agreements** +
__optional__|Contains profile agreements. Available in scope.|<<_profile_profileid_get_agreements,agreements>>
|**apePreference** +
__optional__|List of prefered and available authenticators.|<<_profile_profileid_get_apepreference,apePreference>>
|**device** +
__optional__|Push notification device info|<<_profile_profileid_get_device,device>>
|**email** +
__optional__|Contains current address, possible pending address and verification status. Available in profile.email.r scope.|<<_profile_profileid_get_email,email>>
|**id** +
__optional__|Profile's ID. Available in profile.basic.r scope. +
**Pattern** : `"[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}"`|string
|**image** +
__optional__|Set of profile image links with all available sizes. Available in profile.basic.r scope.|<<_profile_profileid_get_image,image>>
|**links** +
__optional__|Profile related links|< <<_profile_profileid_get_links,links>> > array
|**msisdn** +
__optional__|Contains country code, national number (nn) and phone number. Available in profile.msisdn.r scope.|<<_profile_profileid_get_msisdn,msisdn>>
|**name** +
__optional__|Contains profile's name. Available in profile.basic.r scope.|string
|**operator** +
__optional__|Operator ID. Available in profile.operator.r scope. +
**Pattern** : `"[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}"`|string
|**status** +
__optional__|Profile account statuses|<<_profile_profileid_get_status,status>>
|===

[[_profile_profileid_get_agreements]]
**agreements**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**EULA** +
__optional__|Is Mobile Connect Accelerator End User Licence Agreement accepted|boolean
|**developerEULA** +
__optional__|Is developer EULA accepted|boolean
|===

[[_profile_profileid_get_apepreference]]
**apePreference**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**available** +
__optional__|Prefered authentication methods|<<_profile_profileid_get_apepreference_available,available>>
|**prefered** +
__optional__|Prefered authentication methods|<<_profile_profileid_get_apepreference_prefered,prefered>>
|===

[[_profile_profileid_get_apepreference_available]]
**available**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**loa2** +
__optional__|available LoA2 authentication methods.|< string > array
|**loa3** +
__optional__|available LoA2 authentication methods.|< string > array
|===

[[_profile_profileid_get_apepreference_prefered]]
**prefered**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**loa2** +
__optional__|Prefered LoA2 authentication method.|string
|**loa3** +
__optional__|Prefered LoA3 authentication method.|string
|===

[[_profile_profileid_get_device]]
**device**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**ID** +
__optional__|Push notification device registration ID|string
|**type** +
__optional__|Push notification device type: 'android' or 'ios'|string
|===

[[_profile_profileid_get_email]]
**email**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**address** +
__optional__|Currently used E-mail address|string
|**pending** +
__optional__|E-mail address waiting to be verified|string
|**verified** +
__optional__||boolean
|===

[[_profile_profileid_get_image]]
**image**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**120x120** +
__optional__|Image link|string
|**240x240** +
__optional__|Image link|string
|**60x60** +
__optional__|Image link|string
|**960x960** +
__optional__|Image link|string
|===

[[_profile_profileid_get_links]]
**links**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**href** +
__optional__|Link path|string
|**rel** +
__optional__|Link relation|string
|===

[[_profile_profileid_get_msisdn]]
**msisdn**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**country_code** +
__optional__|number
|**nn** +
__optional__|string
|**number** +
__optional__|string
|===

[[_profile_profileid_get_status]]
**status**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**suspended** +
__optional__|True if profile is currently suspended.|boolean
|===

[[_profile_profileid_get_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* Profile


==== Example HTTP response

===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_profile_profileid_put]]
=== Update profile
....
PUT /profile/{profileId}
....


==== Description
Updates profile.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Path**|**profileId** +
__required__|Profile id|string
|**Body**|**Updating Profile request body** +
__optional__||<<_profile_profileid_put_updating_profile_request_body,Updating Profile request body>>
|===

[[_profile_profileid_put_updating_profile_request_body]]
**Updating Profile request body**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**EUA** +
__optional__|Accept Mobile Connect Accelerator End User Licence Agreement *DEPRECATED*|boolean
|**EULA** +
__optional__|Accept Mobile Connect Digital Identity End User Licence Agreement|<<_profile_profileid_put_eula,EULA>>
|**apePreference** +
__optional__|List of prefered authenticators. Can only be updated with 'provision' privileges.|<<_profile_profileid_put_apepreference,apePreference>>
|**developerEULA** +
__optional__|Accept developer EULA|boolean
|**device** +
__optional__|Push notification device info|<<_profile_profileid_put_device,device>>
|**email** +
__optional__|Contains new profile e-mail address.|string
|**emailVerified** +
__optional__|Flag for setting new email as verified. Only available for provisioning apps.|boolean
|**family_name** +
__optional__|Contains Profile last name.|string
|**given_name** +
__optional__|Contains Profile first name.|string
|**middle_name** +
__optional__|Contains Profile middle name.|string
|**moSmsCode** +
__optional__|Verification code for MO SMS flow. This is the code sent via SMS from user to Mobile Connect Accelerator.|string
|**msisdn** +
__optional__|Contains phone number and verification code. Sending msisdn as string is deprecated, msisdn must be the same as current used in profile.|<<_profile_profileid_put_msisdn,msisdn>>
|**mtSmsCode** +
__optional__|Verification code for MT SMS flow. This is the code user gets from the SMS sent from Mobile Connect Accelerator|string
|**operator** +
__optional__|Operator ID +
**Pattern** : `"[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}"`|string
|**optInCode** +
__optional__|Opt-in verification code|string
|**password** +
__optional__|Contains Profile password. Updating is possible with profile.password.w scope.|<<_profile_profileid_put_password,password>>
|===

[[_profile_profileid_put_eula]]
**EULA**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**issuer** +
__optional__|EULA issuer; Value can be either 'mdi' for default Mobile Connect Digital Identity EULA or MNO unique ID for custom MNO EULA.|string
|**version** +
__optional__|EULA version|string
|===

[[_profile_profileid_put_apepreference]]
**apePreference**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**loa2** +
__optional__|Prefered LoA2 authentication method.|string
|**loa3** +
__optional__|Prefered LoA3 authentication method.|string
|===

[[_profile_profileid_put_device]]
**device**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**ID** +
__optional__|Push notification device registration ID|string
|**type** +
__optional__|Push notification device type: 'android' or 'ios'|string
|===

[[_profile_profileid_put_msisdn]]
**msisdn**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**fresh** +
__required__|New msisdn to be replaced|string
|**requestCode** +
__optional__|Verification code for MT SMS flow. This is the code user gets from the SMS sent from Mobile Connect Accelerator.|string
|**responseCode** +
__optional__|Verification code for MO SMS flow. This is the code sent via SMS from user to Mobile Connect Accelerator.|string
|===

[[_profile_profileid_put_password]]
**password**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**fresh** +
__required__|New password|string
|**old** +
__required__|Old password|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**202**|Authenticator linking status endpoint response.|< <<_profile_profileid_put_response_202,Response 202>> > array
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_profile_profileid_put_response_400,Response 400>>
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**404**|Not Found. The requested resource could not be found but may be available again in the future. Subsequent requests by the client are permissible.|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_profile_profileid_put_response_202]]
**Response 202**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**agreements** +
__optional__|Contains profile agreements. Available in scope.|<<_profile_profileid_put_agreements,agreements>>
|**apePreference** +
__optional__|List of prefered and available authenticators.|<<_profile_profileid_put_apepreference,apePreference>>
|**device** +
__optional__|Push notification device info|<<_profile_profileid_put_device,device>>
|**email** +
__optional__|Contains current address, possible pending address and verification status. Available in profile.email.r scope.|<<_profile_profileid_put_email,email>>
|**id** +
__optional__|Profile's ID. Available in profile.basic.r scope. +
**Pattern** : `"[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}"`|string
|**image** +
__optional__|Set of profile image links with all available sizes. Available in profile.basic.r scope.|<<_profile_profileid_put_image,image>>
|**links** +
__optional__|Profile related links|< <<_profile_profileid_put_links,links>> > array
|**msisdn** +
__optional__|Contains country code, national number (nn) and phone number. Available in profile.msisdn.r scope.|<<_profile_profileid_put_msisdn,msisdn>>
|**name** +
__optional__|Contains profile's name. Available in profile.basic.r scope.|string
|**operator** +
__optional__|Operator ID. Available in profile.operator.r scope. +
**Pattern** : `"[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}"`|string
|**status** +
__optional__|Profile account statuses|<<_profile_profileid_put_status,status>>
|===

[[_profile_profileid_put_agreements]]
**agreements**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**EULA** +
__optional__|Is Mobile Connect Accelerator End User Licence Agreement accepted|boolean
|**developerEULA** +
__optional__|Is developer EULA accepted|boolean
|===

[[_profile_profileid_put_apepreference]]
**apePreference**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**available** +
__optional__|Prefered authentication methods|<<_profile_profileid_put_apepreference_available,available>>
|**prefered** +
__optional__|Prefered authentication methods|<<_profile_profileid_put_apepreference_prefered,prefered>>
|===

[[_profile_profileid_put_apepreference_available]]
**available**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**loa2** +
__optional__|available LoA2 authentication methods.|< string > array
|**loa3** +
__optional__|available LoA2 authentication methods.|< string > array
|===

[[_profile_profileid_put_apepreference_prefered]]
**prefered**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**loa2** +
__optional__|Prefered LoA2 authentication method.|string
|**loa3** +
__optional__|Prefered LoA3 authentication method.|string
|===

[[_profile_profileid_put_device]]
**device**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**ID** +
__optional__|Push notification device registration ID|string
|**type** +
__optional__|Push notification device type: 'android' or 'ios'|string
|===

[[_profile_profileid_put_email]]
**email**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**address** +
__optional__|Currently used E-mail address|string
|**pending** +
__optional__|E-mail address waiting to be verified|string
|**verified** +
__optional__||boolean
|===

[[_profile_profileid_put_image]]
**image**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**120x120** +
__optional__|Image link|string
|**240x240** +
__optional__|Image link|string
|**60x60** +
__optional__|Image link|string
|**960x960** +
__optional__|Image link|string
|===

[[_profile_profileid_put_links]]
**links**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**href** +
__optional__|Link path|string
|**rel** +
__optional__|Link relation|string
|===

[[_profile_profileid_put_msisdn]]
**msisdn**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**country_code** +
__optional__|number
|**nn** +
__optional__|string
|**number** +
__optional__|string
|===

[[_profile_profileid_put_status]]
**status**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**suspended** +
__optional__|True if profile is currently suspended.|boolean
|===

[[_profile_profileid_put_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request, changing_msisdn_not_supported, old_missing_parameter, fresh_missing_parameter, operator_could_not_be_found, invalid_email_format, old_password_is_incorect, new_password_missing_parameter, password_too_short, profile_already_exists, msisdn_already_exists, operator_2f_disabled, required_eula_mdi, required_eula_mno, EUA_only_accepts_true, invalid_device_id, invalid_device_type, inactive_operator, inactive_operator, cannot_register_operator, optInCode_missing_parameter, mtSmsCode_missing_parameter, invalid_code)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* Profile


==== Example HTTP response

===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_profile_profileid_delete]]
=== Removes profile
....
DELETE /profile/{profileId}
....


==== Description
Removes profile after retention period(30 days).


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Path**|**profileId** +
__required__|Profile id|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**204**|On success 204 code is returned.|No Content
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_profile_profileid_delete_response_400,Response 400>>
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**404**|Not Found. The requested resource could not be found but may be available again in the future. Subsequent requests by the client are permissible.|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_profile_profileid_delete_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* Profile


==== Example HTTP response

===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_profile_profileid_application_applicationid_notifications_get]]
=== Get profile application notifications
....
GET /profile/{profileId}/application/{applicationId}/notifications
....


==== Description
Returns collection of profile application notifications.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Path**|**applicationId** +
__required__|Application id|string
|**Path**|**profileId** +
__required__|Profile id|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|Collection of user's notifications for selected application.|< <<_profile_profileid_application_applicationid_notifications_get_response_200,Response 200>> > array
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_profile_profileid_application_applicationid_notifications_get_response_200]]
**Response 200**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**application** +
__optional__|A representation of an application +
**Example** : `{
  "id" : "dde9536c-2ca2-4fdc-a710-17ef0a58ede8",
  "label" : "FakeAnApp v2",
  "shortNames" : [ "FakeAnApp v2" ],
  "url" : "http://awesomeapp.example.com",
  "image" : {
    "45x45" : "http://mobileconnect.io/images/5a486dee-718d-4caf-8e10-c332024c4a32-45x45.jpg",
    "60x60" : "http://mobileconnect.io/images/5a486dee-718d-4caf-8e10-c332024c4a32-60x60.jpg",
    "90x90" : "http://mobileconnect.io/images/5a486dee-718d-4caf-8e10-c332024c4a32-90x90.jpg",
    "120x120" : "http://mobileconnect.io/images/5a486dee-718d-4caf-8e10-c332024c4a32-120x120.jpg",
    "240x240" : "http://mobileconnect.io/images/5a486dee-718d-4caf-8e10-c332024c4a32-240x240.jpg",
    "480x480" : "http://mobileconnect.io/images/5a486dee-718d-4caf-8e10-c332024c4a32-480x480.jpg"
  },
  "promotedImage" : {
    "282x174" : "http://mobileconnect.io/images/50c23757-9dfc-4736-af29-93dace422949-282x174.jpg"
  },
  "author" : {
    "name" : "null",
    "email" : "author@etal.io"
  },
  "createdAt" : "2017-10-02T11:55:58+0000",
  "description" : "Be a description",
  "requireLoa" : "",
  "scopeDescription" : "Reasoning about the scopes that this app is requesting",
  "scope" : [ "scope.r", "scope.w" ],
  "publisher" : {
    "name" : "Mr. Publisher",
    "email" : "author@etal.io"
  },
  "privileges" : {
    "trusted" : false,
    "authorizator" : false,
    "system" : false,
    "autogrant" : false,
    "ump" : false,
    "provision" : true,
    "provisionDescriptor" : false,
    "spUserProvision" : false
  },
  "links" : [ {
    "rel" : "author",
    "href" : "https://api.mdi.ericsson.net/v1/profile/dde9536c-2ca2-4fdc-a710-17ef0a58ede8"
  }, {
    "rel" : "application",
    "href" : "https://api.mdi.ericsson.net/v1/application/dde9536c-2ca2-4fdc-a710-17ef0a58ede8"
  }, {
    "rel" : "keys",
    "href" : "https://api.mdi.ericsson.net/v1/application/dde9536c-2ca2-4fdc-a710-17ef0a58ede8/keys"
  } ],
  "platforms" : [ "web", "ios", "android" ],
  "category" : {
    "id" : "3cf1cd0e-ddf3-4bd3-85e8-45ad9b99621d",
    "name" : "Books"
  },
  "platformUri" : [ {
    "type" : "web",
    "url" : "http:://www.app.dev"
  }, {
    "type" : "ios",
    "url" : "http:://www.app.dev1"
  }, {
    "type" : "android",
    "url" : "http:://www.app.dev2"
  } ],
  "featured" : true,
  "mcPrivileges" : {
    "trustLevel" : 2
  },
  "logoUrl" : "",
  "backgroundUrl" : "",
  "visibility" : "pub",
  "provisionerForMno" : "905538d4-fa92-4305-aa5b-38383b274544"
}`|<<_profile_profileid_application_applicationid_notifications_get_application,application>>
|**authorId** +
__optional__|Profile ID of notification author +
**Pattern** : `"[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}"`|string
|**date** +
__optional__|Date of the notification|string
|**id** +
__optional__|Notification ID +
**Pattern** : `"[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}"`|string
|**ip** +
__optional__|IP number of notification author|string
|**label** +
__optional__|Notification type|string
|**notification** +
__optional__|Notification type - deprecated|enum (login, loginFail, logout, register, passwordChange, grant, grantDeny, revoke, auth2factorSent, auth2factorFail, auth2factorSuccess)
|===

[[_profile_profileid_application_applicationid_notifications_get_application]]
**application**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**author** +
__optional__|Author info|<<_profile_profileid_application_applicationid_notifications_get_application_author,author>>
|**createdAt** +
__optional__|Applications's create date.|string
|**description** +
__optional__|Applications's description.|string
|**id** +
__optional__|Application's id. +
**Pattern** : `"[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}"`|string
|**image** +
__optional__|Set of application image links with all available sizes.|<<_profile_profileid_application_applicationid_notifications_get_application_image,image>>
|**label** +
__optional__|Application's name.|string
|**privileges** +
__optional__|List of privileges application has.|<<_profile_profileid_application_applicationid_notifications_get_application_privileges,privileges>>
|**promotedImage** +
__optional__|Application's promoted image.|<<_profile_profileid_application_applicationid_notifications_get_application_promotedimage,promotedImage>>
|**publisher** +
__optional__|Publisher info.|<<_profile_profileid_application_applicationid_notifications_get_application_publisher,publisher>>
|**requireLoa** +
__optional__|Restricts which LoA is minimal for login (eg. loa2 or loa3). Empty string means that there is no restriction.|string
|**scope** +
__optional__|Array of scopes.|< string > array
|**scopeDescription** +
__optional__|Description, reasoning about requested scopes.|string
|**shortNames** +
__optional__|Array of application short names.|< string > array
|**url** +
__optional__|URL of application's website.|string
|===

[[_profile_profileid_application_applicationid_notifications_get_application_author]]
**author**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**email** +
__optional__|Author's E-mail.|string
|**name** +
__optional__|Author's name.|string
|===

[[_profile_profileid_application_applicationid_notifications_get_application_image]]
**image**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**120x120** +
__optional__|Image link|string
|**240x240** +
__optional__|Image link|string
|**45x45** +
__optional__|Image link|string
|**480x480** +
__optional__|Image link|string
|**60x60** +
__optional__|Image link|string
|**90x90** +
__optional__|Image link|string
|===

[[_profile_profileid_application_applicationid_notifications_get_application_privileges]]
**privileges**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**authorizator** +
__optional__|Defines if application is authorizator.|boolean
|**autogrant** +
__optional__|Defines if application scopes can be autogranted.|boolean
|**provision** +
__optional__|Defines if application has provisioner privileges.|boolean
|**provisionDescriptor** +
__optional__|Defines if application has descriptor provision privilege.|boolean
|**spUserProvision** +
__optional__|Defines if application has privileges to provision users.|boolean
|**system** +
__optional__|Defines if it is system application.|boolean
|**trusted** +
__optional__|Defines if application is trusted.|boolean
|**ump** +
__optional__|Defines if application acts as an User Management Portal|boolean
|===

[[_profile_profileid_application_applicationid_notifications_get_application_promotedimage]]
**promotedImage**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**282x174** +
__optional__|Image link|string
|===

[[_profile_profileid_application_applicationid_notifications_get_application_publisher]]
**publisher**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**email** +
__optional__|Publisher's E-mail|string
|**name** +
__optional__|Publisher's name|string
|===


==== Produces

* `application/json`


==== Tags

* Profile


==== Example HTTP response

===== Response 200
[source,json]
----
[ {
  "id" : "a3a3ddbc-2a2c-4ea7-ba65-5cb04b8e2344",
  "authorId" : "dde9536c-2ca2-4fdc-a710-17ef0a58ede8",
  "date" : "2013-01-01T00:00:00+00:00",
  "type" : "login",
  "ip" : "127.0.0.1",
  "application" : {
    "_id" : "7e71ea8d-933c-41ab-ad3e-c22c94f6ab28",
    "name" : "HereBe Dragons",
    "author" : {
      "name" : "John Author",
      "email" : "author@example.tld"
    },
    "publisher" : {
      "name" : "John Publisher",
      "email" : "publisher@example.tld"
    },
    "url" : "http://awesomeapp.example.com",
    "description" : "Be a description",
    "scopeDescription" : "Reasoning about the scopes that this app is requesting",
    "scope" : [ "profile.basic.r", "profile.basic.w" ],
    "trusted" : false,
    "category" : {
      "id" : "3cf1cd0e-ddf3-4bd3-85e8-45ad9b99621d",
      "name" : "Books"
    }
  }
} ]
----


[[_profile_profileid_applications_favorites_post]]
=== Add favorite granted apps
....
POST /profile/{profileId}/applications/favorites
....


==== Description
Adds granted applications to list of users favorite applications. On success full list of users favorite applications is returned.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Path**|**profileId** +
__required__|Profile id|string
|**Body**|**Contains applications IDs** +
__optional__||<<_profile_profileid_applications_favorites_post_contains_applications_ids,Contains applications IDs>>
|===

[[_profile_profileid_applications_favorites_post_contains_applications_ids]]
**Contains applications IDs**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**applications** +
__required__|Contains applications IDs|< string > array
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**201**|Full list of users favorite applications.|<<_profile_profileid_applications_favorites_post_response_201,Response 201>>
|**400**|Possible reasons for error: app could not be found, bad request, limit exeeds max, page does not exist, wrong page parameters,…|No Content
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_profile_profileid_applications_favorites_post_response_201]]
**Response 201**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**applications** +
__required__|Contains applications IDs|< string > array
|===


==== Produces

* `application/json`


==== Tags

* Profile


==== Example HTTP response

===== Response 201
[source,json]
----
{
  "applications" : [ "4accd0ef-9534-4668-b6f4-5f345fb8c011", "cebc4e1d-2c91-447c-b5b5-8054bf007cdb", "e8001897-e47a-478a-818e-70bbe09f77ae" ]
}
----


[[_profile_profileid_applications_favorites_get]]
=== Read favorite granted apps
....
GET /profile/{profileId}/applications/favorites
....


==== Description
Retrieves users favorite apps.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Path**|**profileId** +
__required__|Profile id|string
|**Query**|**limit** +
__optional__|Pagination parameter, indicates how many records should be returned by page. NOTE: Default value is 20 and Max value is 50|string
|**Query**|**page** +
__optional__|Pagination parameter, indicates which page is to be returned from database. NOTE: Default value is 1.|number
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|Collection of Applications.|< <<_profile_profileid_applications_favorites_get_response_200,Response 200>> > array
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_profile_profileid_applications_favorites_get_response_400,Response 400>>
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_profile_profileid_applications_favorites_get_response_200]]
**Response 200**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**author** +
__optional__|Author info|<<_profile_profileid_applications_favorites_get_author,author>>
|**createdAt** +
__optional__|Applications's create date.|string
|**description** +
__optional__|Applications's description.|string
|**id** +
__optional__|Application's id. +
**Pattern** : `"[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}"`|string
|**image** +
__optional__|Set of application image links with all available sizes.|<<_profile_profileid_applications_favorites_get_image,image>>
|**label** +
__optional__|Application's name.|string
|**privileges** +
__optional__|List of privileges application has.|<<_profile_profileid_applications_favorites_get_privileges,privileges>>
|**promotedImage** +
__optional__|Application's promoted image.|<<_profile_profileid_applications_favorites_get_promotedimage,promotedImage>>
|**publisher** +
__optional__|Publisher info.|<<_profile_profileid_applications_favorites_get_publisher,publisher>>
|**requireLoa** +
__optional__|Restricts which LoA is minimal for login (eg. loa2 or loa3). Empty string means that there is no restriction.|string
|**scope** +
__optional__|Array of scopes.|< string > array
|**scopeDescription** +
__optional__|Description, reasoning about requested scopes.|string
|**shortNames** +
__optional__|Array of application short names.|< string > array
|**url** +
__optional__|URL of application's website.|string
|===

[[_profile_profileid_applications_favorites_get_author]]
**author**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**email** +
__optional__|Author's E-mail.|string
|**name** +
__optional__|Author's name.|string
|===

[[_profile_profileid_applications_favorites_get_image]]
**image**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**120x120** +
__optional__|Image link|string
|**240x240** +
__optional__|Image link|string
|**45x45** +
__optional__|Image link|string
|**480x480** +
__optional__|Image link|string
|**60x60** +
__optional__|Image link|string
|**90x90** +
__optional__|Image link|string
|===

[[_profile_profileid_applications_favorites_get_privileges]]
**privileges**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**authorizator** +
__optional__|Defines if application is authorizator.|boolean
|**autogrant** +
__optional__|Defines if application scopes can be autogranted.|boolean
|**provision** +
__optional__|Defines if application has provisioner privileges.|boolean
|**provisionDescriptor** +
__optional__|Defines if application has descriptor provision privilege.|boolean
|**spUserProvision** +
__optional__|Defines if application has privileges to provision users.|boolean
|**system** +
__optional__|Defines if it is system application.|boolean
|**trusted** +
__optional__|Defines if application is trusted.|boolean
|**ump** +
__optional__|Defines if application acts as an User Management Portal|boolean
|===

[[_profile_profileid_applications_favorites_get_promotedimage]]
**promotedImage**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**282x174** +
__optional__|Image link|string
|===

[[_profile_profileid_applications_favorites_get_publisher]]
**publisher**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**email** +
__optional__|Publisher's E-mail|string
|**name** +
__optional__|Publisher's name|string
|===

[[_profile_profileid_applications_favorites_get_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (wrong_page_parameter, wrong_limit_parameter, promoted_must_be_bool, featured_must_be_bool, limit_exceeds_max)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* Profile


==== Example HTTP response

===== Response 200
[source,json]
----
[ {
  "id" : "dde9536c-2ca2-4fdc-a710-17ef0a58ede8",
  "label" : "FakeAnApp v2",
  "shortNames" : [ "FakeAnApp v2" ],
  "url" : "http://awesomeapp.example.com",
  "description" : "Be a description",
  "scopeDescription" : "Reasoning about the scopes that this app is requesting",
  "scope" : [ "scope.r", "scope.w" ],
  "images" : [ {
    "id" : "5a486dee-718d-4caf-8e10-c332024c4a32",
    "type" : "icon-64x64",
    "href" : "5a486dee-718d-4caf-8e10-c332024c4a32.png"
  } ],
  "author" : {
    "name" : "null",
    "email" : "author@etal.io"
  },
  "publisher" : {
    "name" : "Mr. Publisher",
    "email" : "author@etal.io"
  },
  "trusted" : true,
  "links" : [ {
    "rel" : "author",
    "href" : "https://api.mdi.ericsson.net/v1/profile/dde9536c-2ca2-4fdc-a710-17ef0a58ede8"
  }, {
    "rel" : "application",
    "href" : "https://api.mdi.ericsson.net/v1/application/dde9536c-2ca2-4fdc-a710-17ef0a58ede8"
  }, {
    "rel" : "keys",
    "href" : "https://api.mdi.ericsson.net/v1/application/dde9536c-2ca2-4fdc-a710-17ef0a58ede8/keys"
  } ],
  "platforms" : [ "web", "ios", "android" ],
  "category" : {
    "id" : "3cf1cd0e-ddf3-4bd3-85e8-45ad9b99621d",
    "name" : "Books"
  },
  "platformUri" : [ {
    "type" : "web",
    "url" : "http:://www.app.dev"
  } ],
  "featured" : true
} ]
----


===== Response 400
[source,json]
----
{
  "error" : "wrong_page_parameter",
  "error_description" : "Parameter is not a positive integer."
}
----


[[_profile_profileid_applications_favorites_delete]]
=== Remove favorite granted apps
....
DELETE /profile/{profileId}/applications/favorites
....


==== Description
Removes users granted apps from list of favorite apps.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Path**|**profileId** +
__required__|Profile id|string
|**Body**|**Contains applications IDs** +
__optional__||<<_profile_profileid_applications_favorites_delete_contains_applications_ids,Contains applications IDs>>
|===

[[_profile_profileid_applications_favorites_delete_contains_applications_ids]]
**Contains applications IDs**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**applications** +
__required__|Contains applications IDs|< string > array
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**204**|On success 204 code is returned.|No Content
|**400**|Possible reasons for error: app could not be found, bad request, limit exeeds max, page does not exist, wrong page parameters,…|No Content
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===


==== Produces

* `application/json`


==== Tags

* Profile


[[_profile_profileid_authenticator_mepin_post]]
=== Initiate MePIN linking
....
POST /profile/{profileId}/authenticator/mepin
....


==== Description
Initiates MePIN linking process.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Path**|**profileId** +
__required__|Profile id|string
|**Body**|**MePIN account linking data** +
__optional__||<<_profile_profileid_authenticator_mepin_post_mepin_account_linking_data,MePIN account linking data>>
|===

[[_profile_profileid_authenticator_mepin_post_mepin_account_linking_data]]
**MePIN account linking data**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**nickname** +
__required__|Contains MePIN account nickname.|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**201**|MePIN account linking response with access code.|<<_profile_profileid_authenticator_mepin_post_response_201,Response 201>>
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_profile_profileid_authenticator_mepin_post_response_400,Response 400>>
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**429**|Too Many Requests. The user or application has send too many requests in a given amount of time. If field 'blocked_for' is present in the response, it contains an information about how long the request will remain blocked (in seconds).|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_profile_profileid_authenticator_mepin_post_response_201]]
**Response 201**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**authenticator_data** +
__optional__|Holds authenticator issued data needed for user to complete the linking.|<<_profile_profileid_authenticator_mepin_post_authenticator_data,authenticator_data>>
|**code** +
__optional__|Code allowing to fetch status of linking transaction.|string
|**expire_at** +
__optional__|Holds linking attempt validity time in ISO8601 format (UTC timezone).|string
|**status** +
__optional__|Status OK indicator.|string
|===

[[_profile_profileid_authenticator_mepin_post_authenticator_data]]
**authenticator_data**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**access_code** +
__optional__|MePIN linking access code.|string
|===

[[_profile_profileid_authenticator_mepin_post_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (invalid_nickname)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* Profile


==== Example HTTP response

===== Response 201
[source,json]
----
{
  "status" : "OK",
  "code" : "faeea010970b120722c1dd47a015d480",
  "expire_at" : "2016-02-26T13:11:59+0000",
  "authenticator_data" : {
    "access_code" : "562098105"
  }
}
----


===== Response 400
[source,json]
----
{
  "error" : "invalid_nickname",
  "error_description" : "Parameter nickname must be a string."
}
----


[[_profile_profileid_authenticator_mepin_get]]
=== MePIN account link status
....
GET /profile/{profileId}/authenticator/mepin
....


==== Description
Returns MePIN account link status.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Path**|**profileId** +
__required__|Profile id|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|MePIN account link status|<<_profile_profileid_authenticator_mepin_get_response_200,Response 200>>
|**401**|Access denied. Authentication is required to proceed processing request.|No Content
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_profile_profileid_authenticator_mepin_get_response_200]]
**Response 200**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**expires_at** +
__optional__|Holds expire time of linking attempt in ISO8601 format (UTC timezone). Returned only when status is 'pending'.|string
|**status** +
__optional__|One of three values: "not-linked", "pending", "linked".|string
|===


==== Produces

* `application/json`


==== Tags

* Profile


==== Example HTTP response

===== Response 200
[source,json]
----
{
  "status" : "pending",
  "expires_at" : "2015-10-07T14:37:32+0000"
}
----


[[_profile_profileid_authenticator_mepin_delete]]
=== Unlink MePIN account
....
DELETE /profile/{profileId}/authenticator/mepin
....


==== Description
Unlinks MePIN account by deleting all MePIN information from authenticators field in profile.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Path**|**profileId** +
__required__|Profile id|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**204**|On success 204 code is returned.|No Content
|**401**|Access denied. Authentication is required to proceed processing request.|No Content
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===


==== Produces

* `application/json`


==== Tags

* Profile


[[_profile_profileid_authenticator_noknok_post]]
=== Initiate Nok Nok FIDO authenticator registration
....
POST /profile/{profileId}/authenticator/noknok
....


==== Description
Initate Nok Nok FIDO authenticator registration.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Path**|**profileId** +
__required__|Profile id|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**201**|Nok Nok FIDO authenticator registration with QR code data.|<<_profile_profileid_authenticator_noknok_post_response_201,Response 201>>
|**400**|Possible reasons for error: app could not be found, bad request, limit exeeds max, page does not exist, wrong page parameters,…|No Content
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**429**|Too Many Requests. The user or application has send too many requests in a given amount of time. If field 'blocked_for' is present in the response, it contains an information about how long the request will remain blocked (in seconds).|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_profile_profileid_authenticator_noknok_post_response_201]]
**Response 201**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**authenticator_data** +
__optional__|Holds authenticator issued data needed for user to complete the linking.|<<_profile_profileid_authenticator_noknok_post_authenticator_data,authenticator_data>>
|**code** +
__optional__|Code allowing to fetch status of linking transaction.|string
|**expire_at** +
__optional__|Holds linking attempt validity time in ISO8601 format (UTC timezone).|string
|**status** +
__optional__|Status OK indicator.|string
|===

[[_profile_profileid_authenticator_noknok_post_authenticator_data]]
**authenticator_data**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**qr_code** +
__optional__|Base64 encoded QR code image for initializing Nok Nok app for FIDO Authenticator registration.|string
|===


==== Produces

* `application/json`


==== Tags

* Profile


==== Example HTTP response

===== Response 201
[source,json]
----
{
  "status" : "OK",
  "code" : "c1ac5892b6ad7e9ef327d43d538f5354",
  "expire_at" : "2016-02-26T13:15:11+0000",
  "authenticator_data" : {
    "qr_code" : "iVBORw0KGgoAAAANSUhEUgAAASwAAAEsAQAAAABRBrPYAAADZ0lEQVR42u2aTa6jQAyEjbJgyRG4CbkYEpFyMXITjsCSRYSnqgwzkLzFaFZuDZFelHR/LJz2T9n9zP/mNduFXdiF/YfYZGa3ye5z3y6dDb7c/dHWr3mYFuxYERj+npOvjXvt/jZs9lavzcPwgHbzY4s1wF6wzyo8UI8NNscGD3T4CQrCViwtFb7TyNZfcXYFYV6PdnN8epjcr35hpxxM/oYla2Hkk5vDZB0/fbtlUkymLVj/evtODjmxeOmcap9xYtWMPEZzf0jRObGpRrS4MW4Q9fA3hr6yF+Lm8INkxrDEojGbwlyB4rXy2P1oaWIsTidy1gvlI+wb8agzE5SBdSh+2Ix1lg+FjDaQhQvA4GUwDSfG3IvXUm01HWUQyawETC/WC0TLncdmkYXhb6cfJDE2xVLFGOmsbxn/2tNTXgSGw+pNWpYqyhk8oU0gDs/+lhnD6SBQHOa+Dcem3Huj0z2sBIxpd5C/tSFj6XQy0pQOCsAmJtt2idDHOs1tZa4kYRGYRB8Oa2sr4Hmb+6kWFoJRO1kD+1aEOu2b9EmtRgkYWyGsSwhaEzXdiX2k6MyY6l5UDjqdJKFajS0L58fYBW2VnJuhrPbJwTH35sUgoJrnXsQZMsTYrlJPva0ILPxNLba6UtrM3GtnSzNjUB/G0RKifpTrsa8bY17Ql4H5GFE/uCY0HBXE0IC9t5eAbTmrC03eoxbqnNgZvfwgVxJjDqxnL7qFDJuJmFv6emwrcmMxrXy0CnPFP7+axgdFYBZjb9a9UZmKEr2C+50rYGpsmxI8J7VCEuYcfky12u4SMKTYd6tosU0IqsvQYR1b7LzYpEmTZhw4IiYuedmsBHzS5GkxVz+kgGfBaNWf6mIFZ2dWAhb1m9ESlQNedgt5pa8lYPA3eJl1+zWEq6+jDlE6KAJjBWTAb2GueYEU4tnfMmNWaf4aE7OBapA3KaYUdijiibGJ0eLyrX2Gr3ETr1ia42VQXizmlhqzqpvbc69uGj9Hrzkx3WdtvYWGZX9S8WnclBmLXpT/FYA3qdrwt482NjMWY9b1990uQ0a60Ow43syO+T7D72P0Gv7m52uv5JiklK/xwDOK33JKXJmx/YZdh6W2Ii7nOEz+dMukWNywS0qpcvAN2oSGf1/Ep8Su/zG7sAu7sH/FfgFDc5xl5Z8QAgAAAABJRU5ErkJggg=="
  }
}
----


[[_profile_profileid_authenticator_noknok_get]]
=== Returns Nok Nok authenticator link status
....
GET /profile/{profileId}/authenticator/noknok
....


==== Description
Returns Nok Nok authenticator link status and list of registered FIDO Authenticators.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Path**|**profileId** +
__required__|Profile id|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|MePIN account link status|<<_profile_profileid_authenticator_noknok_get_response_200,Response 200>>
|**401**|Access denied. Authentication is required to proceed processing request.|No Content
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_profile_profileid_authenticator_noknok_get_response_200]]
**Response 200**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**authenticators** +
__optional__|Holds list of registered FIDO Authenticators for this user.|< <<_profile_profileid_authenticator_noknok_get_authenticators,authenticators>> > array
|**status** +
__required__|One of two values: "not-linked", "linked".|string
|===

[[_profile_profileid_authenticator_noknok_get_authenticators]]
**authenticators**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**created_at** +
__required__|FIDO Authenticator registration time in ISO8601 format (UTC timezone)|string
|**description** +
__required__|FIDO Authenticator label.|string
|**handle** +
__required__|FIDO Authenticator registration identifier.|string
|**loa** +
__required__|List of LoA levels satisfied by this authenticator.|< string > array
|===


==== Produces

* `application/json`


==== Tags

* Profile


==== Example HTTP response

===== Response 200
[source,json]
----
{
  "status" : "linked",
  "authenticators" : [ {
    "description" : "Android Fingerprint Authenticator",
    "created_at" : "2016-02-26T14:57:25+0000",
    "handle" : "WyJ1YWZfMS4wIiwiNGU0ZSM0MDEwIiwiWmpCLW50M0JQYlFCYjFvY2lsWm02LVBmZHJzWXJhOXlqVUVDTWJXVjFYOCJd",
    "loa" : [ "loa2", "loa3" ]
  } ]
}
----


[[_profile_profileid_authenticator_noknok_delete]]
=== Deregister Nok Nok FIDO Authenticator
....
DELETE /profile/{profileId}/authenticator/noknok
....


==== Description
Deregister Nok Nok FIDO Authenticator, either deregister all FIDO Authenticators thus unlink Nok Nok or deregister FIDO Authenticator specified by handle.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Path**|**profileId** +
__required__|Profile id|string
|**Body**|**Nok Nok FIDO Authenticator deregistering** +
__optional__||<<_profile_profileid_authenticator_noknok_delete_nok_nok_fido_authenticator_deregistering,Nok Nok FIDO Authenticator deregistering>>
|===

[[_profile_profileid_authenticator_noknok_delete_nok_nok_fido_authenticator_deregistering]]
**Nok Nok FIDO Authenticator deregistering**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**handle** +
__optional__|FIDO authenticator handle parameter for deregistering single FIDO Authenticator.|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**204**|On success 204 code is returned.|No Content
|**401**|Access denied. Authentication is required to proceed processing request.|No Content
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===


==== Produces

* `application/json`


==== Tags

* Profile


[[_profile_profileid_authenticators_get]]
=== Get linked authenticators
....
GET /profile/{profileId}/authenticators
....


==== Description
Returns list of linked authenticators (descriptors).


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Path**|**profileId** +
__required__|Profile id|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|Response body.|< string > array
|**401**|Access denied. Authentication is required to proceed processing request.|No Content
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===


==== Produces

* `application/json`


==== Tags

* Profile


==== Example HTTP response

===== Response 200
[source,json]
----
[ "sms", "password" ]
----


[[_profile_profileid_email-verify_post]]
=== Profile E-mail verification
....
POST /profile/{profileId}/email-verify
....


==== Description
Verifies Profile E-mail by using code sent to Profile owner via E-mail message.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Path**|**profileId** +
__required__|Profile id|string
|**Body**|**E-mail verification request body** +
__optional__||<<_profile_profileid_email-verify_post_e-mail_verification_request_body,E-mail verification request body>>
|===

[[_profile_profileid_email-verify_post_e-mail_verification_request_body]]
**E-mail verification request body**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**code** +
__required__|E-mail verification code received by E-mail.|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|Message on successfully verified E-mail.|<<_profile_profileid_email-verify_post_response_200,Response 200>>
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_profile_profileid_email-verify_post_response_400,Response 400>>
|**401**|Access denied. Authentication is required to proceed processing request.|No Content
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_profile_profileid_email-verify_post_response_200]]
**Response 200**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**status** +
__optional__|string
|===

[[_profile_profileid_email-verify_post_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request, code_missing_parameter, invalid_email_code, email_does_not_exist)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* Profile


==== Example HTTP response

===== Response 200
[source,json]
----
{
  "status" : "OK"
}
----


===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_profile_profileid_grant_applicationkey_get]]
=== Grant - Resource
....
GET /profile/{profileId}/grant/{applicationKey}
....


==== Description
Returns profile grant resource for application key.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Path**|**applicationKey** +
__required__|Application key id|string
|**Path**|**profileId** +
__required__|Profile id|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|Representation of granted application|<<_profile_profileid_grant_applicationkey_get_response_200,Response 200>>
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**404**|Not Found. The requested resource could not be found but may be available again in the future. Subsequent requests by the client are permissible.|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_profile_profileid_grant_applicationkey_get_response_200]]
**Response 200**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**key** +
__optional__|Holds the key that grant was requested for.|string
|**last_login_time** +
__optional__|Holds last login time with this grant.|string
|**rel_profile** +
__optional__|Holds the profile ID of grant.|string
|**scope** +
__optional__|Holds user granted scopes for application.|string
|===


==== Produces

* `application/json`


==== Tags

* Profile


==== Example HTTP response

===== Response 200
[source,json]
----
{
  "rel_profile" : "051a2754-dfd8-4c52-b982-4d48cbaa056b",
  "key" : "536700d3d2c65b245cc887aa2b4d79ab",
  "scope" : "profile.basic.rw, application.rw, openid, offline_access",
  "last_login_time" : "2015-10-07T14:37:32+0000"
}
----


[[_profile_profileid_grant_applicationkey_put]]
=== Grant - Update
....
PUT /profile/{profileId}/grant/{applicationKey}
....


==== Description
Updates profile's granted scopes for given application client key.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Path**|**applicationKey** +
__required__|Application key id|string
|**Path**|**profileId** +
__required__|Profile id|string
|**Body**|**Grant application to profile** +
__optional__||<<_profile_profileid_grant_applicationkey_put_grant_application_to_profile,Grant application to profile>>
|===

[[_profile_profileid_grant_applicationkey_put_grant_application_to_profile]]
**Grant application to profile**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**scope** +
__required__|Application required and approved scopes.|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**202**|Representation of granted application|<<_profile_profileid_grant_applicationkey_put_response_202,Response 202>>
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_profile_profileid_grant_applicationkey_put_response_400,Response 400>>
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**404**|Not Found. The requested resource could not be found but may be available again in the future. Subsequent requests by the client are permissible.|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_profile_profileid_grant_applicationkey_put_response_202]]
**Response 202**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**key** +
__optional__|Holds the key that grant was requested for.|string
|**last_login_time** +
__optional__|Holds last login time with this grant.|string
|**rel_profile** +
__optional__|Holds the profile ID of grant.|string
|**scope** +
__optional__|Holds user granted scopes for application.|string
|===

[[_profile_profileid_grant_applicationkey_put_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request, scope_missing_parameter)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* Profile


==== Example HTTP response

===== Response 202
[source,json]
----
{
  "rel_profile" : "051a2754-dfd8-4c52-b982-4d48cbaa056b",
  "key" : "536700d3d2c65b245cc887aa2b4d79ab",
  "scope" : "profile.basic.rw, application.rw, openid, offline_access",
  "last_login_time" : "2015-10-07T14:37:32+0000"
}
----


===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_profile_profileid_grant_applicationkey_delete]]
=== Grant - Revoke
....
DELETE /profile/{profileId}/grant/{applicationKey}
....


==== Description
Revokes access grant for an application, revokes only current used key for the profile grant.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Path**|**applicationKey** +
__required__|Application key id|string
|**Path**|**profileId** +
__required__|Profile id|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**204**|On success 204 code is returned.|No Content
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_profile_profileid_grant_applicationkey_delete_response_400,Response 400>>
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_profile_profileid_grant_applicationkey_delete_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* Profile


==== Example HTTP response

===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_profile_profileid_grant_applicationkey_logout_post]]
=== Grant - Logout
....
POST /profile/{profileId}/grant/{applicationKey}/logout
....


==== Description
Revokes all access and refresh tokens associated with the grant.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Path**|**applicationKey** +
__required__|Application key id|string
|**Path**|**profileId** +
__required__|Profile id|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|Profile grant logout|<<_profile_profileid_grant_applicationkey_logout_post_response_200,Response 200>>
|**400**|Possible reasons for error: app could not be found, bad request, limit exeeds max, page does not exist, wrong page parameters,…|No Content
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_profile_profileid_grant_applicationkey_logout_post_response_200]]
**Response 200**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**status** +
__optional__|Status OK indicator.|string
|===


==== Produces

* `application/json`


==== Tags

* Profile


==== Example HTTP response

===== Response 200
[source,json]
----
{
  "status" : "OK"
}
----


[[_profile_profileid_grants_post]]
=== Grant - Add
....
POST /profile/{profileId}/grants
....


==== Description
Grants application with requested scopes for profile.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Path**|**profileId** +
__required__|Profile id|string
|**Body**|**Grant application to profile** +
__optional__||<<_profile_profileid_grants_post_grant_application_to_profile,Grant application to profile>>
|===

[[_profile_profileid_grants_post_grant_application_to_profile]]
**Grant application to profile**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**key** +
__required__|Application key for profile grant.|string
|**scope** +
__required__|Application required and approved scopes.|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**201**|Representation of granted application|<<_profile_profileid_grants_post_response_201,Response 201>>
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_profile_profileid_grants_post_response_400,Response 400>>
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**404**|Not Found. The requested resource could not be found but may be available again in the future. Subsequent requests by the client are permissible.|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_profile_profileid_grants_post_response_201]]
**Response 201**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**key** +
__optional__|Holds the key that grant was requested for.|string
|**last_login_time** +
__optional__|Holds last login time with this grant.|string
|**rel_profile** +
__optional__|Holds the profile ID of grant.|string
|**scope** +
__optional__|Holds user granted scopes for application.|string
|===

[[_profile_profileid_grants_post_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request, key_missing_parameter, scope_missing_parameter, profile_grant_exists)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* Profile


==== Example HTTP response

===== Response 201
[source,json]
----
{
  "rel_profile" : "051a2754-dfd8-4c52-b982-4d48cbaa056b",
  "key" : "536700d3d2c65b245cc887aa2b4d79ab",
  "scope" : "profile.basic.rw, application.rw, openid, offline_access",
  "last_login_time" : "2015-10-07T14:37:32+0000"
}
----


===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_profile_profileid_grants_get]]
=== Grant - Add
....
GET /profile/{profileId}/grants
....


==== Description
Grants application with requested scopes for profile.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Path**|**profileId** +
__required__|Profile id|string
|**Query**|**limit** +
__optional__|Pagination parameter, indicates how many records should be returned by page. NOTE: Default value is 20 and Max value is 50|string
|**Query**|**page** +
__optional__|Pagination parameter, indicates which page is to be returned from database. NOTE: Default value is 1.|number
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|List of applications granted to profile.|< <<_profile_profileid_grants_get_response_200,Response 200>> > array
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_profile_profileid_grants_get_response_400,Response 400>>
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**404**|Not Found. The requested resource could not be found but may be available again in the future. Subsequent requests by the client are permissible.|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_profile_profileid_grants_get_response_200]]
**Response 200**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**clients** +
__optional__|Holds a list of granted application clients.|< <<_profile_profileid_grants_get_clients,clients>> > array
|**id** +
__optional__|Holds granted application ID.|string
|**image** +
__optional__|Holds a set of granted application image links with all available sizes.|<<_profile_profileid_grants_get_image,image>>
|**label** +
__optional__|Holds granted application label.|string
|===

[[_profile_profileid_grants_get_clients]]
**clients**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**key** +
__optional__|Holds granted client key (Trusted access only).|string
|**label** +
__optional__|Holds granted client key label.|string
|**last_login_time** +
__optional__|Holds last login time with this grant.|string
|**platform** +
__optional__|Holds granted client platform.|string
|**scope** +
__optional__|Holds granted scopes for application (Trusted access only).|string
|===

[[_profile_profileid_grants_get_image]]
**image**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**120x120** +
__optional__|Image link|string
|**240x240** +
__optional__|Image link|string
|**45x45** +
__optional__|Image link|string
|**480x480** +
__optional__|Image link|string
|**60x60** +
__optional__|Image link|string
|**90x90** +
__optional__|Image link|string
|===

[[_profile_profileid_grants_get_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (wrong_page_parameter, wrong_limit_parameter, limit_exceeds_max)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* Profile


==== Example HTTP response

===== Response 200
[source,json]
----
[ {
  "id" : "26e609a8-613a-4249-b4a5-99ea2151fcb3",
  "label" : "Nuclear FPS game",
  "image" : {
    "45x45" : "http://mobileconnect.io/images/nuclear-45x45.png",
    "60x60" : "http://mobileconnect.io/images/nuclear-60x60.png",
    "90x90" : "http://mobileconnect.io/images/nuclear-90x90.png",
    "120x120" : "http://mobileconnect.io/images/nuclear-120x120.png",
    "240x240" : "http://mobileconnect.io/images/nuclear-240x240.png",
    "480x480" : "http://mobileconnect.io/images/nuclear-480x480.png"
  },
  "clients" : [ {
    "label" : "Nuclear (android)",
    "platform" : "android",
    "key" : "671a1c8dedd619e28a5c089ce4b3ec7a",
    "scope" : "profile.basic.rw, openid, offline_access",
    "last_login_time" : "2015-9-07T14:37:32+0000"
  } ]
}, {
  "id" : "13614105-2043-44aa-a332-9716f70bb0f7",
  "label" : "MyStore application",
  "image" : {
    "45x45" : "http://mobileconnect.io/images/mystore-45x45.png",
    "60x60" : "http://mobileconnect.io/images/mystore-60x60.png",
    "90x90" : "http://mobileconnect.io/images/mystore-90x90.png",
    "120x120" : "http://mobileconnect.io/images/mystore-120x120.png",
    "240x240" : "http://mobileconnect.io/images/mystore-240x240.png",
    "480x480" : "http://mobileconnect.io/images/mystore-480x480.png"
  },
  "clients" : [ {
    "label" : "MyStore application (web)",
    "platform" : "web",
    "key" : "b401aa956824fbbc8c066e4d7a7cffc2",
    "scope" : "profile.basic.rw, openid, phone",
    "last_login_time" : "2015-10-07T14:37:32+0000"
  }, {
    "label" : "MyStore application (android)",
    "platform" : "android",
    "key" : "36cd60eace548469c3b3684e18a404c9",
    "scope" : "profile.basic.rw, openid, phone",
    "last_login_time" : "2015-10-07T14:27:32+0000"
  }, {
    "label" : "MyStore application (iOS)",
    "platform" : "ios",
    "key" : "2255ab3e935d179793c76a172459309a",
    "scope" : "profile.basic.rw, openid, phone",
    "last_login_time" : "2015-10-07T11:33:32+0000"
  } ]
} ]
----


===== Response 400
[source,json]
----
{
  "error" : "wrong_page_parameter",
  "error_description" : "Parameter is not a positive integer."
}
----


[[_profile_profileid_image_post]]
=== Add profile image
....
POST /profile/{profileId}/image
....


==== Description
Adds image to profile.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Path**|**profileId** +
__required__|Profile id|string
|**Body**|**Msisdn claim request body** +
__optional__|Body should be the multipart/form-data of the attachment with key named 'image'. No Content Headers.|<<_profile_profileid_image_post_msisdn_claim_request_body,Msisdn claim request body>>
|===

[[_profile_profileid_image_post_msisdn_claim_request_body]]
**Msisdn claim request body**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**files** +
__optional__|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**201**|Image used by Profile|<<_profile_profileid_image_post_response_201,Response 201>>
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_profile_profileid_image_post_response_400,Response 400>>
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_profile_profileid_image_post_response_201]]
**Response 201**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**id** +
__optional__|Profile image ID|string
|**type** +
__optional__|Type of image|string
|**urls** +
__optional__|Location of resized images|<<_profile_profileid_image_post_urls,urls>>
|===

[[_profile_profileid_image_post_urls]]
**urls**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**120x120** +
__optional__|URL of a resized image|string
|**240x240** +
__optional__|URL of a resized image|string
|**60x60** +
__optional__|URL of a resized image|string
|**960x960** +
__optional__|URL of a resized image|string
|===

[[_profile_profileid_image_post_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request, image_missing, image_wrong_type, image_wrong_format, image_too_big, images_not_resized)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* Profile


==== Example HTTP response

===== Response 201
[source,json]
----
{
  "id" : "051a2754-dfd8-4c52-b982-4d48cbaa056b",
  "type" : "png",
  "urls" : {
    "60x60" : "http://cloudfront.net/image/profiles/051a2754-dfd8-4c52-b982-4d48cbaa056b-60x60.png",
    "120x120" : "http://cloudfront.net/image/profiles/051a2754-dfd8-4c52-b982-4d48cbaa056b-120x120.png",
    "240x240" : "http://cloudfront.net/image/profiles/051a2754-dfd8-4c52-b982-4d48cbaa056b-240x240.png",
    "960x960" : "http://cloudfront.net/image/profiles/051a2754-dfd8-4c52-b982-4d48cbaa056b-960x960.png"
  }
}
----


===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_profile_profileid_image_get]]
=== Returns image URL
....
GET /profile/{profileId}/image
....


==== Description
Returns image URL of specified size. Parameter size is required.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Path**|**profileId** +
__required__|Profile id|string
|**Query**|**size** +
__required__|Parameter for predefined image size. Available values are 60x60, 120x120, 240x240 and 960x960.|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|Image used by Profile|<<_profile_profileid_image_get_response_200,Response 200>>
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_profile_profileid_image_get_response_400,Response 400>>
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**404**|Not Found. The requested resource could not be found but may be available again in the future. Subsequent requests by the client are permissible.|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_profile_profileid_image_get_response_200]]
**Response 200**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**id** +
__optional__|Profile image ID|string
|**type** +
__optional__|Type of image|string
|**urls** +
__optional__|Location of resized images|<<_profile_profileid_image_get_urls,urls>>
|===

[[_profile_profileid_image_get_urls]]
**urls**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**120x120** +
__optional__|URL of a resized image|string
|**240x240** +
__optional__|URL of a resized image|string
|**60x60** +
__optional__|URL of a resized image|string
|**960x960** +
__optional__|URL of a resized image|string
|===

[[_profile_profileid_image_get_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request, image_size_missing_parameter, invalid_image_size)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* Profile


==== Example HTTP response

===== Response 200
[source,json]
----
{
  "id" : "051a2754-dfd8-4c52-b982-4d48cbaa056b",
  "type" : "png",
  "urls" : {
    "60x60" : "http://cloudfront.net/image/profiles/051a2754-dfd8-4c52-b982-4d48cbaa056b-60x60.png",
    "120x120" : "http://cloudfront.net/image/profiles/051a2754-dfd8-4c52-b982-4d48cbaa056b-120x120.png",
    "240x240" : "http://cloudfront.net/image/profiles/051a2754-dfd8-4c52-b982-4d48cbaa056b-240x240.png",
    "960x960" : "http://cloudfront.net/image/profiles/051a2754-dfd8-4c52-b982-4d48cbaa056b-960x960.png"
  }
}
----


===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_profile_profileid_image_delete]]
=== Remove profile image
....
DELETE /profile/{profileId}/image
....


==== Description
Removes image from profile.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Path**|**profileId** +
__required__|Profile id|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**204**|On success 204 code is returned.|No Content
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_profile_profileid_image_delete_response_400,Response 400>>
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**404**|Not Found. The requested resource could not be found but may be available again in the future. Subsequent requests by the client are permissible.|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_profile_profileid_image_delete_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* Profile


==== Example HTTP response

===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_profile_profileid_login-status_applicationkey_post]]
=== Grant - Login status
....
POST /profile/{profileId}/login-status/{applicationKey}
....


==== Description
Returns current user's login status for client ID.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Path**|**applicationKey** +
__required__|Application key id|string
|**Path**|**profileId** +
__required__|Profile id|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|Profile grant login statust|<<_profile_profileid_login-status_applicationkey_post_response_200,Response 200>>
|**400**|Possible reasons for error: app could not be found, bad request, limit exeeds max, page does not exist, wrong page parameters,…|No Content
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_profile_profileid_login-status_applicationkey_post_response_200]]
**Response 200**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**loggedIn** +
__optional__|Login status.|boolean
|===


==== Produces

* `application/json`


==== Tags

* Profile


==== Example HTTP response

===== Response 200
[source,json]
----
{
  "loggedIn" : true
}
----


[[_profile_profileid_notifications_get]]
=== Get profile notifications
....
GET /profile/{profileId}/notifications
....


==== Description
Returns collection of profile notifications.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Path**|**profileId** +
__required__|Profile id|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|Collection of user's notifications.|< <<_profile_profileid_notifications_get_response_200,Response 200>> > array
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_profile_profileid_notifications_get_response_200]]
**Response 200**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**application** +
__optional__|A representation of an application +
**Example** : `{
  "id" : "dde9536c-2ca2-4fdc-a710-17ef0a58ede8",
  "label" : "FakeAnApp v2",
  "shortNames" : [ "FakeAnApp v2" ],
  "url" : "http://awesomeapp.example.com",
  "image" : {
    "45x45" : "http://mobileconnect.io/images/5a486dee-718d-4caf-8e10-c332024c4a32-45x45.jpg",
    "60x60" : "http://mobileconnect.io/images/5a486dee-718d-4caf-8e10-c332024c4a32-60x60.jpg",
    "90x90" : "http://mobileconnect.io/images/5a486dee-718d-4caf-8e10-c332024c4a32-90x90.jpg",
    "120x120" : "http://mobileconnect.io/images/5a486dee-718d-4caf-8e10-c332024c4a32-120x120.jpg",
    "240x240" : "http://mobileconnect.io/images/5a486dee-718d-4caf-8e10-c332024c4a32-240x240.jpg",
    "480x480" : "http://mobileconnect.io/images/5a486dee-718d-4caf-8e10-c332024c4a32-480x480.jpg"
  },
  "promotedImage" : {
    "282x174" : "http://mobileconnect.io/images/50c23757-9dfc-4736-af29-93dace422949-282x174.jpg"
  },
  "author" : {
    "name" : "null",
    "email" : "author@etal.io"
  },
  "createdAt" : "2017-10-02T11:55:58+0000",
  "description" : "Be a description",
  "requireLoa" : "",
  "scopeDescription" : "Reasoning about the scopes that this app is requesting",
  "scope" : [ "scope.r", "scope.w" ],
  "publisher" : {
    "name" : "Mr. Publisher",
    "email" : "author@etal.io"
  },
  "privileges" : {
    "trusted" : false,
    "authorizator" : false,
    "system" : false,
    "autogrant" : false,
    "ump" : false,
    "provision" : true,
    "provisionDescriptor" : false,
    "spUserProvision" : false
  },
  "links" : [ {
    "rel" : "author",
    "href" : "https://api.mdi.ericsson.net/v1/profile/dde9536c-2ca2-4fdc-a710-17ef0a58ede8"
  }, {
    "rel" : "application",
    "href" : "https://api.mdi.ericsson.net/v1/application/dde9536c-2ca2-4fdc-a710-17ef0a58ede8"
  }, {
    "rel" : "keys",
    "href" : "https://api.mdi.ericsson.net/v1/application/dde9536c-2ca2-4fdc-a710-17ef0a58ede8/keys"
  } ],
  "platforms" : [ "web", "ios", "android" ],
  "category" : {
    "id" : "3cf1cd0e-ddf3-4bd3-85e8-45ad9b99621d",
    "name" : "Books"
  },
  "platformUri" : [ {
    "type" : "web",
    "url" : "http:://www.app.dev"
  }, {
    "type" : "ios",
    "url" : "http:://www.app.dev1"
  }, {
    "type" : "android",
    "url" : "http:://www.app.dev2"
  } ],
  "featured" : true,
  "mcPrivileges" : {
    "trustLevel" : 2
  },
  "logoUrl" : "",
  "backgroundUrl" : "",
  "visibility" : "pub",
  "provisionerForMno" : "905538d4-fa92-4305-aa5b-38383b274544"
}`|<<_profile_profileid_notifications_get_application,application>>
|**authorId** +
__optional__|Profile ID of notification author +
**Pattern** : `"[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}"`|string
|**date** +
__optional__|Date of the notification|string
|**id** +
__optional__|Notification ID +
**Pattern** : `"[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}"`|string
|**ip** +
__optional__|IP number of notification author|string
|**label** +
__optional__|Notification type|string
|**notification** +
__optional__|Notification type - deprecated|enum (login, loginFail, logout, register, passwordChange, grant, grantDeny, revoke, auth2factorSent, auth2factorFail, auth2factorSuccess)
|===

[[_profile_profileid_notifications_get_application]]
**application**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**author** +
__optional__|Author info|<<_profile_profileid_notifications_get_application_author,author>>
|**createdAt** +
__optional__|Applications's create date.|string
|**description** +
__optional__|Applications's description.|string
|**id** +
__optional__|Application's id. +
**Pattern** : `"[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}"`|string
|**image** +
__optional__|Set of application image links with all available sizes.|<<_profile_profileid_notifications_get_application_image,image>>
|**label** +
__optional__|Application's name.|string
|**privileges** +
__optional__|List of privileges application has.|<<_profile_profileid_notifications_get_application_privileges,privileges>>
|**promotedImage** +
__optional__|Application's promoted image.|<<_profile_profileid_notifications_get_application_promotedimage,promotedImage>>
|**publisher** +
__optional__|Publisher info.|<<_profile_profileid_notifications_get_application_publisher,publisher>>
|**requireLoa** +
__optional__|Restricts which LoA is minimal for login (eg. loa2 or loa3). Empty string means that there is no restriction.|string
|**scope** +
__optional__|Array of scopes.|< string > array
|**scopeDescription** +
__optional__|Description, reasoning about requested scopes.|string
|**shortNames** +
__optional__|Array of application short names.|< string > array
|**url** +
__optional__|URL of application's website.|string
|===

[[_profile_profileid_notifications_get_application_author]]
**author**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**email** +
__optional__|Author's E-mail.|string
|**name** +
__optional__|Author's name.|string
|===

[[_profile_profileid_notifications_get_application_image]]
**image**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**120x120** +
__optional__|Image link|string
|**240x240** +
__optional__|Image link|string
|**45x45** +
__optional__|Image link|string
|**480x480** +
__optional__|Image link|string
|**60x60** +
__optional__|Image link|string
|**90x90** +
__optional__|Image link|string
|===

[[_profile_profileid_notifications_get_application_privileges]]
**privileges**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**authorizator** +
__optional__|Defines if application is authorizator.|boolean
|**autogrant** +
__optional__|Defines if application scopes can be autogranted.|boolean
|**provision** +
__optional__|Defines if application has provisioner privileges.|boolean
|**provisionDescriptor** +
__optional__|Defines if application has descriptor provision privilege.|boolean
|**spUserProvision** +
__optional__|Defines if application has privileges to provision users.|boolean
|**system** +
__optional__|Defines if it is system application.|boolean
|**trusted** +
__optional__|Defines if application is trusted.|boolean
|**ump** +
__optional__|Defines if application acts as an User Management Portal|boolean
|===

[[_profile_profileid_notifications_get_application_promotedimage]]
**promotedImage**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**282x174** +
__optional__|Image link|string
|===

[[_profile_profileid_notifications_get_application_publisher]]
**publisher**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**email** +
__optional__|Publisher's E-mail|string
|**name** +
__optional__|Publisher's name|string
|===


==== Produces

* `application/json`


==== Tags

* Profile


==== Example HTTP response

===== Response 200
[source,json]
----
[ {
  "id" : "a3a3ddbc-2a2c-4ea7-ba65-5cb04b8e2344",
  "authorId" : "dde9536c-2ca2-4fdc-a710-17ef0a58ede8",
  "date" : "2013-01-01T00:00:00+00:00",
  "type" : "login",
  "ip" : "127.0.0.1",
  "application" : {
    "_id" : "7e71ea8d-933c-41ab-ad3e-c22c94f6ab28",
    "name" : "HereBe Dragons",
    "author" : {
      "name" : "John Author",
      "email" : "author@example.tld"
    },
    "publisher" : {
      "name" : "John Publisher",
      "email" : "publisher@example.tld"
    },
    "url" : "http://awesomeapp.example.com",
    "description" : "Be a description",
    "scopeDescription" : "Reasoning about the scopes that this app is requesting",
    "scope" : [ "profile.basic.r", "profile.basic.w" ],
    "trusted" : false,
    "category" : {
      "id" : "3cf1cd0e-ddf3-4bd3-85e8-45ad9b99621d",
      "name" : "Books"
    }
  }
} ]
----


[[_profile_profileid_suspend_post]]
=== Suspend profile
....
POST /profile/{profileId}/suspend
....


==== Description
Suspends profile.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Path**|**profileId** +
__required__|Profile id|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**202**|On success 202 code is returned.|No Content
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===


==== Produces

* `application/json`


==== Tags

* Profile


[[_profile_profileid_suspend_delete]]
=== Unsuspend profile
....
DELETE /profile/{profileId}/suspend
....


==== Description
Unsuspends profile.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Path**|**profileId** +
__required__|Profile id|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**202**|On success 202 code is returned.|No Content
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===


==== Produces

* `application/json`


==== Tags

* Profile


[[_profiles_post]]
=== Add new profile
....
POST /profiles
....


==== Description
Adds new profile.


==== Parameters

[options="header", cols=".^2a,.^3a,.^4a"]
|===
|Type|Name|Schema
|**Body**|**Adding Profile request body** +
__optional__|<<_profiles_post_adding_profile_request_body,Adding Profile request body>>
|===

[[_profiles_post_adding_profile_request_body]]
**Adding Profile request body**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**EUA** +
__optional__|Accept Mobile Connect Accelerator End User Licence Agreement *DEPRECATED*|boolean
|**EULA** +
__optional__|Accept Mobile Connect Digital Identity End User Licence Agreement|<<_profiles_post_eula,EULA>>
|**developerEULA** +
__optional__|Accept developer EULA|boolean
|**device** +
__optional__|Push notification device info|<<_profiles_post_device,device>>
|**email** +
__optional__|Contains E-mail address.|string
|**emailVerified** +
__optional__|Flag for setting new email as verified. Only available for provisioning apps.|boolean
|**family_name** +
__optional__|Contains last name.|string
|**given_name** +
__optional__|Contains first name.|string
|**middle_name** +
__optional__|Contains middle name.|string
|**moSmsCode** +
__optional__|Verification code for MO SMS flow. This is the code sent via SMS from user to Mobile Connect Accelerator.|string
|**msisdn** +
__required__|Contains phone number|string
|**mtSmsCode** +
__optional__|Verification code for MT SMS flow. This is the code user gets from the SMS sent from Mobile Connect Accelerator|string
|**operator** +
__required__|Operator ID +
**Pattern** : `"[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}"`|string
|**optInCode** +
__optional__|Opt-in verification code|string
|**password** +
__optional__|Contains user's password and code.|<<_profiles_post_password,password>>
|**shortRegistration** +
__optional__|Flag for LoA2 (short) registration. If it is set to true, password is not required for registration.|boolean
|===

[[_profiles_post_eula]]
**EULA**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**issuer** +
__optional__|EULA issuer; Value can be either 'mdi' for default Mobile Connect Digital Identity EULA or MNO unique ID for custom MNO EULA.|string
|**version** +
__optional__|EULA version|string
|===

[[_profiles_post_device]]
**device**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**ID** +
__optional__|Push notification device registration ID|string
|**type** +
__optional__|Push notification device type: 'android' or 'ios'|string
|===

[[_profiles_post_password]]
**password**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**code** +
__optional__|Verification code (obsolete in favor of requestCode)|string
|**fresh** +
__required__|Contains new profile password|string
|**requestCode** +
__optional__|Verification code for MT SMS flow. This is the code user gets from the SMS sent from Mobile Connect Accelerator|string
|**responseCode** +
__optional__|Verification code for MO SMS flow. This is the code sent via SMS from user to Mobile Connect Accelerator.|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**201**|Authenticator linking status endpoint response.|< <<_profiles_post_response_201,Response 201>> > array
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_profiles_post_response_400,Response 400>>
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_profiles_post_response_201]]
**Response 201**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**agreements** +
__optional__|Contains profile agreements. Available in scope.|<<_profiles_post_agreements,agreements>>
|**apePreference** +
__optional__|List of prefered and available authenticators.|<<_profiles_post_apepreference,apePreference>>
|**device** +
__optional__|Push notification device info|<<_profiles_post_device,device>>
|**email** +
__optional__|Contains current address, possible pending address and verification status. Available in profile.email.r scope.|<<_profiles_post_email,email>>
|**id** +
__optional__|Profile's ID. Available in profile.basic.r scope. +
**Pattern** : `"[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}"`|string
|**image** +
__optional__|Set of profile image links with all available sizes. Available in profile.basic.r scope.|<<_profiles_post_image,image>>
|**links** +
__optional__|Profile related links|< <<_profiles_post_links,links>> > array
|**msisdn** +
__optional__|Contains country code, national number (nn) and phone number. Available in profile.msisdn.r scope.|<<_profiles_post_msisdn,msisdn>>
|**name** +
__optional__|Contains profile's name. Available in profile.basic.r scope.|string
|**operator** +
__optional__|Operator ID. Available in profile.operator.r scope. +
**Pattern** : `"[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}"`|string
|**status** +
__optional__|Profile account statuses|<<_profiles_post_status,status>>
|===

[[_profiles_post_agreements]]
**agreements**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**EULA** +
__optional__|Is Mobile Connect Accelerator End User Licence Agreement accepted|boolean
|**developerEULA** +
__optional__|Is developer EULA accepted|boolean
|===

[[_profiles_post_apepreference]]
**apePreference**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**available** +
__optional__|Prefered authentication methods|<<_profiles_post_apepreference_available,available>>
|**prefered** +
__optional__|Prefered authentication methods|<<_profiles_post_apepreference_prefered,prefered>>
|===

[[_profiles_post_apepreference_available]]
**available**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**loa2** +
__optional__|available LoA2 authentication methods.|< string > array
|**loa3** +
__optional__|available LoA2 authentication methods.|< string > array
|===

[[_profiles_post_apepreference_prefered]]
**prefered**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**loa2** +
__optional__|Prefered LoA2 authentication method.|string
|**loa3** +
__optional__|Prefered LoA3 authentication method.|string
|===

[[_profiles_post_device]]
**device**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**ID** +
__optional__|Push notification device registration ID|string
|**type** +
__optional__|Push notification device type: 'android' or 'ios'|string
|===

[[_profiles_post_email]]
**email**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**address** +
__optional__|Currently used E-mail address|string
|**pending** +
__optional__|E-mail address waiting to be verified|string
|**verified** +
__optional__||boolean
|===

[[_profiles_post_image]]
**image**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**120x120** +
__optional__|Image link|string
|**240x240** +
__optional__|Image link|string
|**60x60** +
__optional__|Image link|string
|**960x960** +
__optional__|Image link|string
|===

[[_profiles_post_links]]
**links**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**href** +
__optional__|Link path|string
|**rel** +
__optional__|Link relation|string
|===

[[_profiles_post_msisdn]]
**msisdn**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**country_code** +
__optional__|number
|**nn** +
__optional__|string
|**number** +
__optional__|string
|===

[[_profiles_post_status]]
**status**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**suspended** +
__optional__|True if profile is currently suspended.|boolean
|===

[[_profiles_post_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request, profile_already_exists, invalid_msisdn, operator_could_not_be_found, invalid_email_format, old_missing_parameter, old_password_is_incorect, new_password_missing_parameter, password_too_short, operator_2f_disabled, required_eula_mdi, required_eula_mno, EUA_only_accepts_true, invalid_device_id, invalid_device_type, inactive_operator, cannot_register_operator, optInCode_missing_parameter, mtSmsCode_missing_parameter, invalid_code, invalid_shortRegistration, short_registration_not_supported)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* Profile


==== Example HTTP response

===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_profiles_get]]
=== Retrieve list of profiles
....
GET /profiles
....


==== Description
Returns the list of all profiles.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Query**|**limit** +
__optional__|Pagination parameter, indicates how many records should be returned by page. NOTE: Default value is 20 and Max value is 50|string
|**Query**|**page** +
__optional__|Pagination parameter, indicates which page is to be returned from database. NOTE: Default value is 1.|number
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|Authenticator linking status endpoint response.|< < <<_profiles_get_response_200,Response 200>> > array > array
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_profiles_get_response_400,Response 400>>
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**404**|Not Found. The requested resource could not be found but may be available again in the future. Subsequent requests by the client are permissible.|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_profiles_get_response_200]]
**Response 200**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**agreements** +
__optional__|Contains profile agreements. Available in scope.|<<_profiles_get_agreements,agreements>>
|**apePreference** +
__optional__|List of prefered and available authenticators.|<<_profiles_get_apepreference,apePreference>>
|**device** +
__optional__|Push notification device info|<<_profiles_get_device,device>>
|**email** +
__optional__|Contains current address, possible pending address and verification status. Available in profile.email.r scope.|<<_profiles_get_email,email>>
|**id** +
__optional__|Profile's ID. Available in profile.basic.r scope. +
**Pattern** : `"[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}"`|string
|**image** +
__optional__|Set of profile image links with all available sizes. Available in profile.basic.r scope.|<<_profiles_get_image,image>>
|**links** +
__optional__|Profile related links|< <<_profiles_get_links,links>> > array
|**msisdn** +
__optional__|Contains country code, national number (nn) and phone number. Available in profile.msisdn.r scope.|<<_profiles_get_msisdn,msisdn>>
|**name** +
__optional__|Contains profile's name. Available in profile.basic.r scope.|string
|**operator** +
__optional__|Operator ID. Available in profile.operator.r scope. +
**Pattern** : `"[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}"`|string
|**status** +
__optional__|Profile account statuses|<<_profiles_get_status,status>>
|===

[[_profiles_get_agreements]]
**agreements**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**EULA** +
__optional__|Is Mobile Connect Accelerator End User Licence Agreement accepted|boolean
|**developerEULA** +
__optional__|Is developer EULA accepted|boolean
|===

[[_profiles_get_apepreference]]
**apePreference**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**available** +
__optional__|Prefered authentication methods|<<_profiles_get_apepreference_available,available>>
|**prefered** +
__optional__|Prefered authentication methods|<<_profiles_get_apepreference_prefered,prefered>>
|===

[[_profiles_get_apepreference_available]]
**available**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**loa2** +
__optional__|available LoA2 authentication methods.|< string > array
|**loa3** +
__optional__|available LoA2 authentication methods.|< string > array
|===

[[_profiles_get_apepreference_prefered]]
**prefered**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**loa2** +
__optional__|Prefered LoA2 authentication method.|string
|**loa3** +
__optional__|Prefered LoA3 authentication method.|string
|===

[[_profiles_get_device]]
**device**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**ID** +
__optional__|Push notification device registration ID|string
|**type** +
__optional__|Push notification device type: 'android' or 'ios'|string
|===

[[_profiles_get_email]]
**email**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**address** +
__optional__|Currently used E-mail address|string
|**pending** +
__optional__|E-mail address waiting to be verified|string
|**verified** +
__optional__||boolean
|===

[[_profiles_get_image]]
**image**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**120x120** +
__optional__|Image link|string
|**240x240** +
__optional__|Image link|string
|**60x60** +
__optional__|Image link|string
|**960x960** +
__optional__|Image link|string
|===

[[_profiles_get_links]]
**links**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**href** +
__optional__|Link path|string
|**rel** +
__optional__|Link relation|string
|===

[[_profiles_get_msisdn]]
**msisdn**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**country_code** +
__optional__|number
|**nn** +
__optional__|string
|**number** +
__optional__|string
|===

[[_profiles_get_status]]
**status**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**suspended** +
__optional__|True if profile is currently suspended.|boolean
|===

[[_profiles_get_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (wrong_page_parameter, wrong_limit_parameter, limit_exceeds_max)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* Profile


==== Example HTTP response

===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_profiles_delete]]
=== Delete profiles
....
DELETE /profiles
....


==== Description
Instantly removes multiple profiles without TTL.


==== Parameters

[options="header", cols=".^2a,.^3a,.^4a"]
|===
|Type|Name|Schema
|**Body**|**Delete multiple profiles request body** +
__optional__|<<_profiles_delete_delete_multiple_profiles_request_body,Delete multiple profiles request body>>
|===

[[_profiles_delete_delete_multiple_profiles_request_body]]
**Delete multiple profiles request body**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**msisdns** +
__required__|User's batch of MSISDN's|< string > array
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**204**|On success 204 code is returned.|No Content
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_profiles_delete_response_400,Response 400>>
|**401**|Access denied. Authentication is required to proceed processing request.|No Content
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**429**|Too Many Requests. The user or application has send too many requests in a given amount of time. If field 'blocked_for' is present in the response, it contains an information about how long the request will remain blocked (in seconds).|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_profiles_delete_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request, input_parameter_limit)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* Profile


==== Example HTTP response

===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_provision_applications_get]]
=== Get application collection for provision app
....
GET /provision/applications
....


==== Description
Returns collection of applications.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Query**|**author** +
__optional__|Parameter for filtering applications by author|string
|**Query**|**complete** +
__optional__|Parameter for showing also internal MDI apps that are normally hidden in this collection|boolean
|**Query**|**featured** +
__optional__|Parameter for filtering applications by featured tags|boolean
|**Query**|**key** +
__optional__|Parameter for filtering applications by client key|string
|**Query**|**limit** +
__optional__|Pagination parameter, indicates how many records should be returned by page. NOTE: Default value is 20 and Max value is 50|string
|**Query**|**operator** +
__optional__|Parameter for retrieving applications promoted by specific operator|string
|**Query**|**page** +
__optional__|Pagination parameter, indicates which page is to be returned from database. NOTE: Default value is 1.|number
|**Query**|**platform** +
__optional__|Parameter for filtering applications by platform|string
|**Query**|**q** +
__optional__|Parameter for searching by application names|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|Collection of Applications.|< <<_provision_applications_get_response_200,Response 200>> > array
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_provision_applications_get_response_400,Response 400>>
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_provision_applications_get_response_200]]
**Response 200**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**author** +
__optional__|Author info|<<_provision_applications_get_author,author>>
|**createdAt** +
__optional__|Applications's create date.|string
|**description** +
__optional__|Applications's description.|string
|**id** +
__optional__|Application's id. +
**Pattern** : `"[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}"`|string
|**image** +
__optional__|Set of application image links with all available sizes.|<<_provision_applications_get_image,image>>
|**label** +
__optional__|Application's name.|string
|**privileges** +
__optional__|List of privileges application has.|<<_provision_applications_get_privileges,privileges>>
|**promotedImage** +
__optional__|Application's promoted image.|<<_provision_applications_get_promotedimage,promotedImage>>
|**publisher** +
__optional__|Publisher info.|<<_provision_applications_get_publisher,publisher>>
|**requireLoa** +
__optional__|Restricts which LoA is minimal for login (eg. loa2 or loa3). Empty string means that there is no restriction.|string
|**scope** +
__optional__|Array of scopes.|< string > array
|**scopeDescription** +
__optional__|Description, reasoning about requested scopes.|string
|**shortNames** +
__optional__|Array of application short names.|< string > array
|**url** +
__optional__|URL of application's website.|string
|===

[[_provision_applications_get_author]]
**author**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**email** +
__optional__|Author's E-mail.|string
|**name** +
__optional__|Author's name.|string
|===

[[_provision_applications_get_image]]
**image**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**120x120** +
__optional__|Image link|string
|**240x240** +
__optional__|Image link|string
|**45x45** +
__optional__|Image link|string
|**480x480** +
__optional__|Image link|string
|**60x60** +
__optional__|Image link|string
|**90x90** +
__optional__|Image link|string
|===

[[_provision_applications_get_privileges]]
**privileges**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**authorizator** +
__optional__|Defines if application is authorizator.|boolean
|**autogrant** +
__optional__|Defines if application scopes can be autogranted.|boolean
|**provision** +
__optional__|Defines if application has provisioner privileges.|boolean
|**provisionDescriptor** +
__optional__|Defines if application has descriptor provision privilege.|boolean
|**spUserProvision** +
__optional__|Defines if application has privileges to provision users.|boolean
|**system** +
__optional__|Defines if it is system application.|boolean
|**trusted** +
__optional__|Defines if application is trusted.|boolean
|**ump** +
__optional__|Defines if application acts as an User Management Portal|boolean
|===

[[_provision_applications_get_promotedimage]]
**promotedImage**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**282x174** +
__optional__|Image link|string
|===

[[_provision_applications_get_publisher]]
**publisher**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**email** +
__optional__|Publisher's E-mail|string
|**name** +
__optional__|Publisher's name|string
|===

[[_provision_applications_get_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (wrong_page_parameter, wrong_limit_parameter, limit_exceeds_max)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* Application


==== Example HTTP response

===== Response 200
[source,json]
----
[ {
  "id" : "dde9536c-2ca2-4fdc-a710-17ef0a58ede8",
  "label" : "FakeAnApp v2",
  "shortNames" : [ "FakeAnApp v2" ],
  "url" : "http://awesomeapp.example.com",
  "description" : "Be a description",
  "scopeDescription" : "Reasoning about the scopes that this app is requesting",
  "scope" : [ "scope.r", "scope.w" ],
  "images" : [ {
    "id" : "5a486dee-718d-4caf-8e10-c332024c4a32",
    "type" : "icon-64x64",
    "href" : "5a486dee-718d-4caf-8e10-c332024c4a32.png"
  } ],
  "author" : {
    "name" : "null",
    "email" : "author@etal.io"
  },
  "publisher" : {
    "name" : "Mr. Publisher",
    "email" : "author@etal.io"
  },
  "trusted" : true,
  "links" : [ {
    "rel" : "author",
    "href" : "https://api.mdi.ericsson.net/v1/profile/dde9536c-2ca2-4fdc-a710-17ef0a58ede8"
  }, {
    "rel" : "application",
    "href" : "https://api.mdi.ericsson.net/v1/application/dde9536c-2ca2-4fdc-a710-17ef0a58ede8"
  }, {
    "rel" : "keys",
    "href" : "https://api.mdi.ericsson.net/v1/application/dde9536c-2ca2-4fdc-a710-17ef0a58ede8/keys"
  } ],
  "platforms" : [ "web", "ios", "android" ],
  "category" : {
    "id" : "3cf1cd0e-ddf3-4bd3-85e8-45ad9b99621d",
    "name" : "Books"
  },
  "platformUri" : [ {
    "type" : "web",
    "url" : "http:://www.app.dev"
  } ],
  "featured" : true
} ]
----


===== Response 400
[source,json]
----
{
  "error" : "wrong_page_parameter",
  "error_description" : "Parameter is not a positive integer."
}
----


[[_provision_netops_netopid_get]]
=== Get network operator for provision app
....
GET /provision/netops/{netOpId}
....


==== Description
Returns operator data.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Path**|**netOpId** +
__required__|Network operator id|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|A representation of a netop.|<<_provision_netops_netopid_get_response_200,Response 200>>
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_provision_netops_netopid_get_response_400,Response 400>>
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**404**|Not Found. The requested resource could not be found but may be available again in the future. Subsequent requests by the client are permissible.|No Content
|===

[[_provision_netops_netopid_get_response_200]]
**Response 200**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**EULA** +
__optional__|Contains MNO EULA (End User Licence Agreement) version. If this property is missing from response it means that MNO does not have it's own individual EULA, so default Mobile Connect Digital Identity EULA is used instead.|<<_provision_netops_netopid_get_eula,EULA>>
|**acl** +
__optional__|List of MSISDN prefix patterns allowed to register with this MNO (BETA only).|< string > array
|**ape_preference** +
__optional__|APE settings for the MNO|<<_provision_netops_netopid_get_ape_preference,ape_preference>>
|**brandname** +
__optional__|Contains brandname|string
|**cacheExpiry** +
__optional__|MCXG cache expiry time (in seconds) for profile data from this MNO.|integer
|**contactPerson** +
__optional__|Contact person in CU. May contain anything (a name, a user ID, e-mail address, MSISDN, etc.).|string
|**countryCode** +
__optional__|Network operator country calling code|number
|**featuredBlacklist** +
__optional__|List of featured applications not to be displayed when this operator is in use.|< string > array
|**hni** +
__optional__|NetOp's array of Home Network Identities (MCC + MNC combinations)|< string > array
|**id** +
__optional__|Network operator's ID +
**Pattern** : `"[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}"`|string
|**image** +
__optional__|Set of netop logo's|<<_provision_netops_netopid_get_image,image>>
|**integrated** +
__optional__|Is Network operator part of Mobile Connect Accelerator network, possible values are HOST - is integrated, BETA - in BETA stage, INACTIVE - MNO is unnasigned, INTERNATIONAL - grandfather node.|string
|**iso** +
__optional__|Network operator ISO 3166 country code|string
|**label** +
__optional__|Network operator name|string
|**logo** +
__optional__|NetOp's logo (deprecated)|string
|**minimumUmpLoa** +
__optional__|Defines the LOA level to use on User Management Portal|string
|**mnp** +
__optional__|Network operator has mobile number portability|boolean
|**ndc** +
__optional__|List of national destination codes (NDC)|< string > array
|**policy** +
__optional__|Authentication policy containing prioritized list of LoA2 and LoA3 authentication methods.|<<_provision_netops_netopid_get_policy,policy>>
|**promotedApps** +
__optional__|Promoted applications for network operator|< string > array
|**servingOperator** +
__optional__|Network operator serving name|string
|**shortRegistration** +
__optional__|Short Registration on this MNO is enabled when the value is true.|boolean
|**showcase** +
__optional__|Contains showcase header|string
|**smsGateway** +
__optional__|SMS gateway to be used by network operator|string
|**timezone** +
__optional__|Timezone of the MNO. The format must be in form Area/Location such as Europe/London.|string
|**trialEnd** +
__optional__|Trial period end time (ISO8601).|string
|**trialNotificationList** +
__optional__|Contacts that will be notified when trial is ending.|<<_provision_netops_netopid_get_trialnotificationlist,trialNotificationList>>
|**trialStart** +
__optional__|Trial period beginning time (ISO8601).|string
|**ttl** +
__optional__|Network operator transaction TTL configuration|<<_provision_netops_netopid_get_ttl,ttl>>
|**type** +
__optional__|Operator type (mcx or mckGlobal)|string
|**umpHost** +
__optional__|User Management Portal url|string
|**ussdGateway** +
__optional__|USSD gateway to be used by network operator|string
|**visible** +
__optional__|Is Network operator visible|boolean
|===

[[_provision_netops_netopid_get_eula]]
**EULA**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**issuer** +
__optional__|MNO EULA issuer. If issuer is different than MNO ID, then EULA belongs to MNO host.|string
|**version** +
__optional__|MNO EULA version|string
|===

[[_provision_netops_netopid_get_ape_preference]]
**ape_preference**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**loa2** +
__optional__|Preferred authentication method for LoA2|string
|**loa3** +
__optional__|Preferred authentication method for LoA3|string
|===

[[_provision_netops_netopid_get_image]]
**image**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**120x120** +
__optional__|Image link|string
|**240x240** +
__optional__|Image link|string
|**500x500** +
__optional__|Image link|string
|**60x60** +
__optional__|Image link|string
|===

[[_provision_netops_netopid_get_policy]]
**policy**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**loa2** +
__optional__|Prioritized list of LoA2 authentication methods.|< string > array
|**loa3** +
__optional__|Prioritized list of LoA3 authentication methods.|< <<_provision_netops_netopid_get_policy_loa3,loa3>> > array
|===

[[_provision_netops_netopid_get_policy_loa3]]
**loa3**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**primary** +
__optional__|Primary authentication method|string
|**secondary** +
__optional__|Secondary (confirmatory) authentication method|string
|===

[[_provision_netops_netopid_get_trialnotificationlist]]
**trialNotificationList**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**email** +
__optional__|Email addresses.|< string > array
|**msisdn** +
__optional__|Phone number and name of the person.|< <<_provision_netops_netopid_get_trialnotificationlist_msisdn,msisdn>> > array
|===

[[_provision_netops_netopid_get_trialnotificationlist_msisdn]]
**msisdn**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**msisdn** +
__optional__|Phone number.|string
|**name** +
__optional__|Anything that describes the msisdn (first name, last name, …). It can be also empty.|string
|===

[[_provision_netops_netopid_get_ttl]]
**ttl**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**authentication** +
__optional__|Authentication transaction|integer
|**authorization** +
__optional__|Authorization transaction|integer
|**bcAuthentication** +
__optional__|Backchannel authentication transaction|integer
|**bcAuthorization** +
__optional__|Backchannel authorization transaction|integer
|**registration** +
__optional__|Registration transaction|integer
|===

[[_provision_netops_netopid_get_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request, no_route)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* Netops


==== Example HTTP response

===== Response 200
[source,json]
----
{
  "id" : "e406cd96-6e3d-48fe-8a97-5f1498f158d7",
  "label" : "Mobitel",
  "iso" : "si",
  "type" : "mcx",
  "countryCode" : 386,
  "servingOperator" : "Mobitel",
  "integrated" : "BETA",
  "hni" : [ "293.41", "293.64" ],
  "ndc" : [ "33" ],
  "mnp" : false,
  "smsGateway" : "IPX",
  "ussdGateway" : "mobitel_ussd",
  "promotedApps" : [ "00000000-de30-da1a-0000-db7b5bb0ac61" ],
  "featuredBlacklist" : [ "00000000-de30-da1a-0000-adb342bcc881" ],
  "showcase" : "Showcase header",
  "brandname" : "Commercial brandname",
  "EULA" : {
    "issuer" : "e406cd96-6e3d-48fe-8a97-5f1498f158d7",
    "version" : "1.0"
  },
  "policy" : {
    "loa2" : [ "password", "sms" ],
    "loa3" : [ {
      "primary" : "password",
      "secondary" : "sms"
    }, {
      "primary" : "password",
      "secondary" : "push"
    } ]
  },
  "acl" : [ "9355" ],
  "cacheExpiry" : 172800,
  "shortRegistration" : true,
  "trialNotificationList" : {
    "msisdn" : [ {
      "name" : "FirstName LastName1",
      "msisdn" : "+38640333771"
    }, {
      "name" : "FirstName LastName2",
      "msisdn" : "+38640333772"
    } ],
    "email" : [ "example1@example.org", "example2@example.org" ]
  },
  "timezone" : "Europe/London",
  "minimumUmpLoa" : "loa3",
  "ttl" : {
    "registration" : 300,
    "authentication" : 180,
    "bcAuthentication" : 180,
    "authorization" : 180,
    "bcAuthorization" : 180
  },
  "umpHost" : "https://www.eaxmpleportal.com"
}
----


===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_provision_profile_get]]
=== Retrieve profile by MSISDN for provision app
....
GET /provision/profile
....


==== Description
Returns requested profile.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Query**|**msisdn** +
__required__|Parameter for retrieving profile with provided MSISDN.|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|Authenticator linking status endpoint response.|< <<_provision_profile_get_response_200,Response 200>> > array
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_provision_profile_get_response_400,Response 400>>
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**404**|Not Found. The requested resource could not be found but may be available again in the future. Subsequent requests by the client are permissible.|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_provision_profile_get_response_200]]
**Response 200**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**agreements** +
__optional__|Contains profile agreements. Available in scope.|<<_provision_profile_get_agreements,agreements>>
|**apePreference** +
__optional__|List of prefered and available authenticators.|<<_provision_profile_get_apepreference,apePreference>>
|**device** +
__optional__|Push notification device info|<<_provision_profile_get_device,device>>
|**email** +
__optional__|Contains current address, possible pending address and verification status. Available in profile.email.r scope.|<<_provision_profile_get_email,email>>
|**id** +
__optional__|Profile's ID. Available in profile.basic.r scope. +
**Pattern** : `"[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}"`|string
|**image** +
__optional__|Set of profile image links with all available sizes. Available in profile.basic.r scope.|<<_provision_profile_get_image,image>>
|**links** +
__optional__|Profile related links|< <<_provision_profile_get_links,links>> > array
|**msisdn** +
__optional__|Contains country code, national number (nn) and phone number. Available in profile.msisdn.r scope.|<<_provision_profile_get_msisdn,msisdn>>
|**name** +
__optional__|Contains profile's name. Available in profile.basic.r scope.|string
|**operator** +
__optional__|Operator ID. Available in profile.operator.r scope. +
**Pattern** : `"[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}"`|string
|**status** +
__optional__|Profile account statuses|<<_provision_profile_get_status,status>>
|===

[[_provision_profile_get_agreements]]
**agreements**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**EULA** +
__optional__|Is Mobile Connect Accelerator End User Licence Agreement accepted|boolean
|**developerEULA** +
__optional__|Is developer EULA accepted|boolean
|===

[[_provision_profile_get_apepreference]]
**apePreference**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**available** +
__optional__|Prefered authentication methods|<<_provision_profile_get_apepreference_available,available>>
|**prefered** +
__optional__|Prefered authentication methods|<<_provision_profile_get_apepreference_prefered,prefered>>
|===

[[_provision_profile_get_apepreference_available]]
**available**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**loa2** +
__optional__|available LoA2 authentication methods.|< string > array
|**loa3** +
__optional__|available LoA2 authentication methods.|< string > array
|===

[[_provision_profile_get_apepreference_prefered]]
**prefered**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**loa2** +
__optional__|Prefered LoA2 authentication method.|string
|**loa3** +
__optional__|Prefered LoA3 authentication method.|string
|===

[[_provision_profile_get_device]]
**device**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**ID** +
__optional__|Push notification device registration ID|string
|**type** +
__optional__|Push notification device type: 'android' or 'ios'|string
|===

[[_provision_profile_get_email]]
**email**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**address** +
__optional__|Currently used E-mail address|string
|**pending** +
__optional__|E-mail address waiting to be verified|string
|**verified** +
__optional__||boolean
|===

[[_provision_profile_get_image]]
**image**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**120x120** +
__optional__|Image link|string
|**240x240** +
__optional__|Image link|string
|**60x60** +
__optional__|Image link|string
|**960x960** +
__optional__|Image link|string
|===

[[_provision_profile_get_links]]
**links**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**href** +
__optional__|Link path|string
|**rel** +
__optional__|Link relation|string
|===

[[_provision_profile_get_msisdn]]
**msisdn**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**country_code** +
__optional__|number
|**nn** +
__optional__|string
|**number** +
__optional__|string
|===

[[_provision_profile_get_status]]
**status**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**suspended** +
__optional__|True if profile is currently suspended.|boolean
|===

[[_provision_profile_get_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request, msisdn_missing_parameter)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* Profile


==== Example HTTP response

===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_provision_profile_delete]]
=== Delete profile for provision app
....
DELETE /provision/profile
....


==== Description
Instantly removes profile without TTL.


==== Parameters

[options="header", cols=".^2a,.^3a,.^4a"]
|===
|Type|Name|Schema
|**Body**|**Delete profile request body** +
__optional__|<<_provision_profile_delete_delete_profile_request_body,Delete profile request body>>
|===

[[_provision_profile_delete_delete_profile_request_body]]
**Delete profile request body**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**msisdn** +
__required__|User's MSISDN|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**204**|On success 204 code is returned.|No Content
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_provision_profile_delete_response_400,Response 400>>
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**404**|Not Found. The requested resource could not be found but may be available again in the future. Subsequent requests by the client are permissible.|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_provision_profile_delete_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* Profile


==== Example HTTP response

===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_provision_profile_profileid_get]]
=== Retrieve profile by ID for provision app
....
GET /provision/profile/{profileId}
....


==== Description
Returns requested profile.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Path**|**profileId** +
__required__|Profile id|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|Authenticator linking status endpoint response.|< <<_provision_profile_profileid_get_response_200,Response 200>> > array
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_provision_profile_profileid_get_response_400,Response 400>>
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_provision_profile_profileid_get_response_200]]
**Response 200**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**agreements** +
__optional__|Contains profile agreements. Available in scope.|<<_provision_profile_profileid_get_agreements,agreements>>
|**apePreference** +
__optional__|List of prefered and available authenticators.|<<_provision_profile_profileid_get_apepreference,apePreference>>
|**device** +
__optional__|Push notification device info|<<_provision_profile_profileid_get_device,device>>
|**email** +
__optional__|Contains current address, possible pending address and verification status. Available in profile.email.r scope.|<<_provision_profile_profileid_get_email,email>>
|**id** +
__optional__|Profile's ID. Available in profile.basic.r scope. +
**Pattern** : `"[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}"`|string
|**image** +
__optional__|Set of profile image links with all available sizes. Available in profile.basic.r scope.|<<_provision_profile_profileid_get_image,image>>
|**links** +
__optional__|Profile related links|< <<_provision_profile_profileid_get_links,links>> > array
|**msisdn** +
__optional__|Contains country code, national number (nn) and phone number. Available in profile.msisdn.r scope.|<<_provision_profile_profileid_get_msisdn,msisdn>>
|**name** +
__optional__|Contains profile's name. Available in profile.basic.r scope.|string
|**operator** +
__optional__|Operator ID. Available in profile.operator.r scope. +
**Pattern** : `"[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}"`|string
|**status** +
__optional__|Profile account statuses|<<_provision_profile_profileid_get_status,status>>
|===

[[_provision_profile_profileid_get_agreements]]
**agreements**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**EULA** +
__optional__|Is Mobile Connect Accelerator End User Licence Agreement accepted|boolean
|**developerEULA** +
__optional__|Is developer EULA accepted|boolean
|===

[[_provision_profile_profileid_get_apepreference]]
**apePreference**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**available** +
__optional__|Prefered authentication methods|<<_provision_profile_profileid_get_apepreference_available,available>>
|**prefered** +
__optional__|Prefered authentication methods|<<_provision_profile_profileid_get_apepreference_prefered,prefered>>
|===

[[_provision_profile_profileid_get_apepreference_available]]
**available**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**loa2** +
__optional__|available LoA2 authentication methods.|< string > array
|**loa3** +
__optional__|available LoA2 authentication methods.|< string > array
|===

[[_provision_profile_profileid_get_apepreference_prefered]]
**prefered**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**loa2** +
__optional__|Prefered LoA2 authentication method.|string
|**loa3** +
__optional__|Prefered LoA3 authentication method.|string
|===

[[_provision_profile_profileid_get_device]]
**device**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**ID** +
__optional__|Push notification device registration ID|string
|**type** +
__optional__|Push notification device type: 'android' or 'ios'|string
|===

[[_provision_profile_profileid_get_email]]
**email**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**address** +
__optional__|Currently used E-mail address|string
|**pending** +
__optional__|E-mail address waiting to be verified|string
|**verified** +
__optional__||boolean
|===

[[_provision_profile_profileid_get_image]]
**image**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**120x120** +
__optional__|Image link|string
|**240x240** +
__optional__|Image link|string
|**60x60** +
__optional__|Image link|string
|**960x960** +
__optional__|Image link|string
|===

[[_provision_profile_profileid_get_links]]
**links**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**href** +
__optional__|Link path|string
|**rel** +
__optional__|Link relation|string
|===

[[_provision_profile_profileid_get_msisdn]]
**msisdn**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**country_code** +
__optional__|number
|**nn** +
__optional__|string
|**number** +
__optional__|string
|===

[[_provision_profile_profileid_get_status]]
**status**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**suspended** +
__optional__|True if profile is currently suspended.|boolean
|===

[[_provision_profile_profileid_get_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* Profile


==== Example HTTP response

===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_provision_profile_profileid_put]]
=== Update profile for provision app
....
PUT /provision/profile/{profileId}
....


==== Description
Updates profile.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Path**|**profileId** +
__required__|Profile id|string
|**Body**|**Updating Profile request body** +
__optional__||<<_provision_profile_profileid_put_updating_profile_request_body,Updating Profile request body>>
|===

[[_provision_profile_profileid_put_updating_profile_request_body]]
**Updating Profile request body**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**EUA** +
__optional__|Accept Mobile Connect Accelerator End User Licence Agreement *DEPRECATED*|boolean
|**EULA** +
__optional__|Accept Mobile Connect Digital Identity End User Licence Agreement|<<_provision_profile_profileid_put_eula,EULA>>
|**apePreference** +
__optional__|List of prefered authenticators. Can only be updated with 'provision' privileges.|<<_provision_profile_profileid_put_apepreference,apePreference>>
|**developerEULA** +
__optional__|Accept developer EULA|boolean
|**device** +
__optional__|Push notification device info|<<_provision_profile_profileid_put_device,device>>
|**email** +
__optional__|Contains new profile e-mail address.|string
|**emailVerified** +
__optional__|Flag for setting new email as verified. Only available for provisioning apps.|boolean
|**family_name** +
__optional__|Contains Profile last name.|string
|**given_name** +
__optional__|Contains Profile first name.|string
|**middle_name** +
__optional__|Contains Profile middle name.|string
|**moSmsCode** +
__optional__|Verification code for MO SMS flow. This is the code sent via SMS from user to Mobile Connect Accelerator.|string
|**msisdn** +
__optional__|Contains phone number and verification code. Sending msisdn as string is deprecated, msisdn must be the same as current used in profile.|<<_provision_profile_profileid_put_msisdn,msisdn>>
|**mtSmsCode** +
__optional__|Verification code for MT SMS flow. This is the code user gets from the SMS sent from Mobile Connect Accelerator|string
|**operator** +
__optional__|Operator ID +
**Pattern** : `"[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}"`|string
|**optInCode** +
__optional__|Opt-in verification code|string
|**password** +
__optional__|Contains Profile password. Updating is possible with profile.password.w scope.|<<_provision_profile_profileid_put_password,password>>
|===

[[_provision_profile_profileid_put_eula]]
**EULA**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**issuer** +
__optional__|EULA issuer; Value can be either 'mdi' for default Mobile Connect Digital Identity EULA or MNO unique ID for custom MNO EULA.|string
|**version** +
__optional__|EULA version|string
|===

[[_provision_profile_profileid_put_apepreference]]
**apePreference**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**loa2** +
__optional__|Prefered LoA2 authentication method.|string
|**loa3** +
__optional__|Prefered LoA3 authentication method.|string
|===

[[_provision_profile_profileid_put_device]]
**device**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**ID** +
__optional__|Push notification device registration ID|string
|**type** +
__optional__|Push notification device type: 'android' or 'ios'|string
|===

[[_provision_profile_profileid_put_msisdn]]
**msisdn**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**fresh** +
__required__|New msisdn to be replaced|string
|**requestCode** +
__optional__|Verification code for MT SMS flow. This is the code user gets from the SMS sent from Mobile Connect Accelerator.|string
|**responseCode** +
__optional__|Verification code for MO SMS flow. This is the code sent via SMS from user to Mobile Connect Accelerator.|string
|===

[[_provision_profile_profileid_put_password]]
**password**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**fresh** +
__required__|New password|string
|**old** +
__required__|Old password|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**202**|Authenticator linking status endpoint response.|< <<_provision_profile_profileid_put_response_202,Response 202>> > array
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_provision_profile_profileid_put_response_400,Response 400>>
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**404**|Not Found. The requested resource could not be found but may be available again in the future. Subsequent requests by the client are permissible.|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_provision_profile_profileid_put_response_202]]
**Response 202**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**agreements** +
__optional__|Contains profile agreements. Available in scope.|<<_provision_profile_profileid_put_agreements,agreements>>
|**apePreference** +
__optional__|List of prefered and available authenticators.|<<_provision_profile_profileid_put_apepreference,apePreference>>
|**device** +
__optional__|Push notification device info|<<_provision_profile_profileid_put_device,device>>
|**email** +
__optional__|Contains current address, possible pending address and verification status. Available in profile.email.r scope.|<<_provision_profile_profileid_put_email,email>>
|**id** +
__optional__|Profile's ID. Available in profile.basic.r scope. +
**Pattern** : `"[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}"`|string
|**image** +
__optional__|Set of profile image links with all available sizes. Available in profile.basic.r scope.|<<_provision_profile_profileid_put_image,image>>
|**links** +
__optional__|Profile related links|< <<_provision_profile_profileid_put_links,links>> > array
|**msisdn** +
__optional__|Contains country code, national number (nn) and phone number. Available in profile.msisdn.r scope.|<<_provision_profile_profileid_put_msisdn,msisdn>>
|**name** +
__optional__|Contains profile's name. Available in profile.basic.r scope.|string
|**operator** +
__optional__|Operator ID. Available in profile.operator.r scope. +
**Pattern** : `"[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}"`|string
|**status** +
__optional__|Profile account statuses|<<_provision_profile_profileid_put_status,status>>
|===

[[_provision_profile_profileid_put_agreements]]
**agreements**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**EULA** +
__optional__|Is Mobile Connect Accelerator End User Licence Agreement accepted|boolean
|**developerEULA** +
__optional__|Is developer EULA accepted|boolean
|===

[[_provision_profile_profileid_put_apepreference]]
**apePreference**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**available** +
__optional__|Prefered authentication methods|<<_provision_profile_profileid_put_apepreference_available,available>>
|**prefered** +
__optional__|Prefered authentication methods|<<_provision_profile_profileid_put_apepreference_prefered,prefered>>
|===

[[_provision_profile_profileid_put_apepreference_available]]
**available**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**loa2** +
__optional__|available LoA2 authentication methods.|< string > array
|**loa3** +
__optional__|available LoA2 authentication methods.|< string > array
|===

[[_provision_profile_profileid_put_apepreference_prefered]]
**prefered**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**loa2** +
__optional__|Prefered LoA2 authentication method.|string
|**loa3** +
__optional__|Prefered LoA3 authentication method.|string
|===

[[_provision_profile_profileid_put_device]]
**device**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**ID** +
__optional__|Push notification device registration ID|string
|**type** +
__optional__|Push notification device type: 'android' or 'ios'|string
|===

[[_provision_profile_profileid_put_email]]
**email**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**address** +
__optional__|Currently used E-mail address|string
|**pending** +
__optional__|E-mail address waiting to be verified|string
|**verified** +
__optional__||boolean
|===

[[_provision_profile_profileid_put_image]]
**image**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**120x120** +
__optional__|Image link|string
|**240x240** +
__optional__|Image link|string
|**60x60** +
__optional__|Image link|string
|**960x960** +
__optional__|Image link|string
|===

[[_provision_profile_profileid_put_links]]
**links**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**href** +
__optional__|Link path|string
|**rel** +
__optional__|Link relation|string
|===

[[_provision_profile_profileid_put_msisdn]]
**msisdn**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**country_code** +
__optional__|number
|**nn** +
__optional__|string
|**number** +
__optional__|string
|===

[[_provision_profile_profileid_put_status]]
**status**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**suspended** +
__optional__|True if profile is currently suspended.|boolean
|===

[[_provision_profile_profileid_put_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request, changing_msisdn_not_supported, old_missing_parameter, fresh_missing_parameter, operator_could_not_be_found, invalid_email_format, old_password_is_incorect, new_password_missing_parameter, password_too_short, profile_already_exists, msisdn_already_exists, operator_2f_disabled, required_eula_mdi, required_eula_mno, EUA_only_accepts_true, invalid_device_id, invalid_device_type, inactive_operator, inactive_operator, cannot_register_operator, optInCode_missing_parameter, mtSmsCode_missing_parameter, invalid_code)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* Profile


==== Example HTTP response

===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_provision_profile_profileid_delete]]
=== Removes profile for provision app
....
DELETE /provision/profile/{profileId}
....


==== Description
Removes profile after retention period(30 days).


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Path**|**profileId** +
__required__|Profile id|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**204**|On success 204 code is returned.|No Content
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_provision_profile_profileid_delete_response_400,Response 400>>
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**404**|Not Found. The requested resource could not be found but may be available again in the future. Subsequent requests by the client are permissible.|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_provision_profile_profileid_delete_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* Profile


==== Example HTTP response

===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_provision_profile_profileid_applications_favorites_post]]
=== dd favorite granted apps for provision app
....
POST /provision/profile/{profileId}/applications/favorites
....


==== Description
Adds granted applications to list of users favorite applications. On success full list of users favorite applications is returned.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Path**|**profileId** +
__required__|Profile id|string
|**Body**|**Contains applications IDs** +
__optional__||<<_provision_profile_profileid_applications_favorites_post_contains_applications_ids,Contains applications IDs>>
|===

[[_provision_profile_profileid_applications_favorites_post_contains_applications_ids]]
**Contains applications IDs**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**applications** +
__required__|Contains applications IDs|< string > array
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**201**|Full list of users favorite applications.|<<_provision_profile_profileid_applications_favorites_post_response_201,Response 201>>
|**400**|Possible reasons for error: app could not be found, bad request, limit exeeds max, page does not exist, wrong page parameters,…|No Content
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_provision_profile_profileid_applications_favorites_post_response_201]]
**Response 201**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**applications** +
__required__|Contains applications IDs|< string > array
|===


==== Produces

* `application/json`


==== Tags

* Profile


==== Example HTTP response

===== Response 201
[source,json]
----
{
  "applications" : [ "4accd0ef-9534-4668-b6f4-5f345fb8c011", "cebc4e1d-2c91-447c-b5b5-8054bf007cdb", "e8001897-e47a-478a-818e-70bbe09f77ae" ]
}
----


[[_provision_profile_profileid_applications_favorites_get]]
=== Read favorite granted apps for provision app
....
GET /provision/profile/{profileId}/applications/favorites
....


==== Description
Retrieves users favorite apps.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Path**|**profileId** +
__required__|Profile id|string
|**Query**|**limit** +
__optional__|Pagination parameter, indicates how many records should be returned by page. NOTE: Default value is 20 and Max value is 50|string
|**Query**|**page** +
__optional__|Pagination parameter, indicates which page is to be returned from database. NOTE: Default value is 1.|number
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|Collection of Applications.|< <<_provision_profile_profileid_applications_favorites_get_response_200,Response 200>> > array
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_provision_profile_profileid_applications_favorites_get_response_400,Response 400>>
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_provision_profile_profileid_applications_favorites_get_response_200]]
**Response 200**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**author** +
__optional__|Author info|<<_provision_profile_profileid_applications_favorites_get_author,author>>
|**createdAt** +
__optional__|Applications's create date.|string
|**description** +
__optional__|Applications's description.|string
|**id** +
__optional__|Application's id. +
**Pattern** : `"[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}"`|string
|**image** +
__optional__|Set of application image links with all available sizes.|<<_provision_profile_profileid_applications_favorites_get_image,image>>
|**label** +
__optional__|Application's name.|string
|**privileges** +
__optional__|List of privileges application has.|<<_provision_profile_profileid_applications_favorites_get_privileges,privileges>>
|**promotedImage** +
__optional__|Application's promoted image.|<<_provision_profile_profileid_applications_favorites_get_promotedimage,promotedImage>>
|**publisher** +
__optional__|Publisher info.|<<_provision_profile_profileid_applications_favorites_get_publisher,publisher>>
|**requireLoa** +
__optional__|Restricts which LoA is minimal for login (eg. loa2 or loa3). Empty string means that there is no restriction.|string
|**scope** +
__optional__|Array of scopes.|< string > array
|**scopeDescription** +
__optional__|Description, reasoning about requested scopes.|string
|**shortNames** +
__optional__|Array of application short names.|< string > array
|**url** +
__optional__|URL of application's website.|string
|===

[[_provision_profile_profileid_applications_favorites_get_author]]
**author**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**email** +
__optional__|Author's E-mail.|string
|**name** +
__optional__|Author's name.|string
|===

[[_provision_profile_profileid_applications_favorites_get_image]]
**image**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**120x120** +
__optional__|Image link|string
|**240x240** +
__optional__|Image link|string
|**45x45** +
__optional__|Image link|string
|**480x480** +
__optional__|Image link|string
|**60x60** +
__optional__|Image link|string
|**90x90** +
__optional__|Image link|string
|===

[[_provision_profile_profileid_applications_favorites_get_privileges]]
**privileges**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**authorizator** +
__optional__|Defines if application is authorizator.|boolean
|**autogrant** +
__optional__|Defines if application scopes can be autogranted.|boolean
|**provision** +
__optional__|Defines if application has provisioner privileges.|boolean
|**provisionDescriptor** +
__optional__|Defines if application has descriptor provision privilege.|boolean
|**spUserProvision** +
__optional__|Defines if application has privileges to provision users.|boolean
|**system** +
__optional__|Defines if it is system application.|boolean
|**trusted** +
__optional__|Defines if application is trusted.|boolean
|**ump** +
__optional__|Defines if application acts as an User Management Portal|boolean
|===

[[_provision_profile_profileid_applications_favorites_get_promotedimage]]
**promotedImage**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**282x174** +
__optional__|Image link|string
|===

[[_provision_profile_profileid_applications_favorites_get_publisher]]
**publisher**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**email** +
__optional__|Publisher's E-mail|string
|**name** +
__optional__|Publisher's name|string
|===

[[_provision_profile_profileid_applications_favorites_get_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (wrong_page_parameter, wrong_limit_parameter, promoted_must_be_bool, featured_must_be_bool, limit_exceeds_max)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* Profile


==== Example HTTP response

===== Response 200
[source,json]
----
[ {
  "id" : "dde9536c-2ca2-4fdc-a710-17ef0a58ede8",
  "label" : "FakeAnApp v2",
  "shortNames" : [ "FakeAnApp v2" ],
  "url" : "http://awesomeapp.example.com",
  "description" : "Be a description",
  "scopeDescription" : "Reasoning about the scopes that this app is requesting",
  "scope" : [ "scope.r", "scope.w" ],
  "images" : [ {
    "id" : "5a486dee-718d-4caf-8e10-c332024c4a32",
    "type" : "icon-64x64",
    "href" : "5a486dee-718d-4caf-8e10-c332024c4a32.png"
  } ],
  "author" : {
    "name" : "null",
    "email" : "author@etal.io"
  },
  "publisher" : {
    "name" : "Mr. Publisher",
    "email" : "author@etal.io"
  },
  "trusted" : true,
  "links" : [ {
    "rel" : "author",
    "href" : "https://api.mdi.ericsson.net/v1/profile/dde9536c-2ca2-4fdc-a710-17ef0a58ede8"
  }, {
    "rel" : "application",
    "href" : "https://api.mdi.ericsson.net/v1/application/dde9536c-2ca2-4fdc-a710-17ef0a58ede8"
  }, {
    "rel" : "keys",
    "href" : "https://api.mdi.ericsson.net/v1/application/dde9536c-2ca2-4fdc-a710-17ef0a58ede8/keys"
  } ],
  "platforms" : [ "web", "ios", "android" ],
  "category" : {
    "id" : "3cf1cd0e-ddf3-4bd3-85e8-45ad9b99621d",
    "name" : "Books"
  },
  "platformUri" : [ {
    "type" : "web",
    "url" : "http:://www.app.dev"
  } ],
  "featured" : true
} ]
----


===== Response 400
[source,json]
----
{
  "error" : "wrong_page_parameter",
  "error_description" : "Parameter is not a positive integer."
}
----


[[_provision_profile_profileid_applications_favorites_delete]]
=== Remove favorite granted apps for provision app
....
DELETE /provision/profile/{profileId}/applications/favorites
....


==== Description
Removes users granted apps from list of favorite apps.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Path**|**profileId** +
__required__|Profile id|string
|**Body**|**Contains applications IDs** +
__optional__||<<_provision_profile_profileid_applications_favorites_delete_contains_applications_ids,Contains applications IDs>>
|===

[[_provision_profile_profileid_applications_favorites_delete_contains_applications_ids]]
**Contains applications IDs**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**applications** +
__required__|Contains applications IDs|< string > array
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**204**|On success 204 code is returned.|No Content
|**400**|Possible reasons for error: app could not be found, bad request, limit exeeds max, page does not exist, wrong page parameters,…|No Content
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===


==== Produces

* `application/json`


==== Tags

* Profile


[[_provision_profile_profileid_authenticators_authenticator_post]]
=== Update authenticator data
....
POST /provision/profile/{profileId}/authenticators/{authenticator}
....


==== Description
Manual provision of authenticator data for profile. Requires 'provision' privileges. {authenticator} is authenticator ID and it has to be added to users operator.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Path**|**authenticator** +
__required__|Authenticator id|string
|**Path**|**profileId** +
__required__|Profile id|string
|**Body**|**Authenticator data to be saved in profile.** +
__optional__||<<_provision_profile_profileid_authenticators_authenticator_post_authenticator_data_to_be_saved_in_profile,Authenticator data to be saved in profile.>>
|===

[[_provision_profile_profileid_authenticators_authenticator_post_authenticator_data_to_be_saved_in_profile]]
**Authenticator data to be saved in profile.**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**data** +
__optional__|Data passed to authenticator for processing. Can have any number of additional properties.|object
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**201**|On success 201 code is returned.|No Content
|**400**|Possible reasons for error: app could not be found, bad request, limit exeeds max, page does not exist, wrong page parameters,…|No Content
|**401**|Access denied. Authentication is required to proceed processing request.|No Content
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===


==== Produces

* `application/json`


==== Tags

* Profile


[[_provision_profile_profileid_authenticators_authenticator_delete]]
=== Delete authenticator data
....
DELETE /provision/profile/{profileId}/authenticators/{authenticator}
....


==== Description
Manual delete of authenticator data for profile. Requires 'provision' privileges. {authenticator} is authenticator ID and it has to be added to users operator.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Path**|**authenticator** +
__required__|Authenticator id|string
|**Path**|**profileId** +
__required__|Profile id|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**204**|On success 204 code is returned.|No Content
|**400**|Possible reasons for error: app could not be found, bad request, limit exeeds max, page does not exist, wrong page parameters,…|No Content
|**401**|Access denied. Authentication is required to proceed processing request.|No Content
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===


==== Produces

* `application/json`


==== Tags

* Profile


[[_provision_profile_profileid_grant_applicationkey_get]]
=== Grant - Resource for provision app
....
GET /provision/profile/{profileId}/grant/{applicationKey}
....


==== Description
Grants application with requested scopes for profile.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Path**|**profileId** +
__required__|Profile id|string
|**Query**|**limit** +
__optional__|Pagination parameter, indicates how many records should be returned by page. NOTE: Default value is 20 and Max value is 50|string
|**Query**|**page** +
__optional__|Pagination parameter, indicates which page is to be returned from database. NOTE: Default value is 1.|number
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|List of applications granted to profile.|< <<_provision_profile_profileid_grant_applicationkey_get_response_200,Response 200>> > array
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_provision_profile_profileid_grant_applicationkey_get_response_400,Response 400>>
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**404**|Not Found. The requested resource could not be found but may be available again in the future. Subsequent requests by the client are permissible.|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_provision_profile_profileid_grant_applicationkey_get_response_200]]
**Response 200**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**clients** +
__optional__|Holds a list of granted application clients.|< <<_provision_profile_profileid_grant_applicationkey_get_clients,clients>> > array
|**id** +
__optional__|Holds granted application ID.|string
|**image** +
__optional__|Holds a set of granted application image links with all available sizes.|<<_provision_profile_profileid_grant_applicationkey_get_image,image>>
|**label** +
__optional__|Holds granted application label.|string
|===

[[_provision_profile_profileid_grant_applicationkey_get_clients]]
**clients**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**key** +
__optional__|Holds granted client key (Trusted access only).|string
|**label** +
__optional__|Holds granted client key label.|string
|**last_login_time** +
__optional__|Holds last login time with this grant.|string
|**platform** +
__optional__|Holds granted client platform.|string
|**scope** +
__optional__|Holds granted scopes for application (Trusted access only).|string
|===

[[_provision_profile_profileid_grant_applicationkey_get_image]]
**image**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**120x120** +
__optional__|Image link|string
|**240x240** +
__optional__|Image link|string
|**45x45** +
__optional__|Image link|string
|**480x480** +
__optional__|Image link|string
|**60x60** +
__optional__|Image link|string
|**90x90** +
__optional__|Image link|string
|===

[[_provision_profile_profileid_grant_applicationkey_get_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (wrong_page_parameter, wrong_limit_parameter, limit_exceeds_max)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* Profile


==== Example HTTP response

===== Response 200
[source,json]
----
[ {
  "id" : "26e609a8-613a-4249-b4a5-99ea2151fcb3",
  "label" : "Nuclear FPS game",
  "image" : {
    "45x45" : "http://mobileconnect.io/images/nuclear-45x45.png",
    "60x60" : "http://mobileconnect.io/images/nuclear-60x60.png",
    "90x90" : "http://mobileconnect.io/images/nuclear-90x90.png",
    "120x120" : "http://mobileconnect.io/images/nuclear-120x120.png",
    "240x240" : "http://mobileconnect.io/images/nuclear-240x240.png",
    "480x480" : "http://mobileconnect.io/images/nuclear-480x480.png"
  },
  "clients" : [ {
    "label" : "Nuclear (android)",
    "platform" : "android",
    "key" : "671a1c8dedd619e28a5c089ce4b3ec7a",
    "scope" : "profile.basic.rw, openid, offline_access",
    "last_login_time" : "2015-9-07T14:37:32+0000"
  } ]
}, {
  "id" : "13614105-2043-44aa-a332-9716f70bb0f7",
  "label" : "MyStore application",
  "image" : {
    "45x45" : "http://mobileconnect.io/images/mystore-45x45.png",
    "60x60" : "http://mobileconnect.io/images/mystore-60x60.png",
    "90x90" : "http://mobileconnect.io/images/mystore-90x90.png",
    "120x120" : "http://mobileconnect.io/images/mystore-120x120.png",
    "240x240" : "http://mobileconnect.io/images/mystore-240x240.png",
    "480x480" : "http://mobileconnect.io/images/mystore-480x480.png"
  },
  "clients" : [ {
    "label" : "MyStore application (web)",
    "platform" : "web",
    "key" : "b401aa956824fbbc8c066e4d7a7cffc2",
    "scope" : "profile.basic.rw, openid, phone",
    "last_login_time" : "2015-10-07T14:37:32+0000"
  }, {
    "label" : "MyStore application (android)",
    "platform" : "android",
    "key" : "36cd60eace548469c3b3684e18a404c9",
    "scope" : "profile.basic.rw, openid, phone",
    "last_login_time" : "2015-10-07T14:27:32+0000"
  }, {
    "label" : "MyStore application (iOS)",
    "platform" : "ios",
    "key" : "2255ab3e935d179793c76a172459309a",
    "scope" : "profile.basic.rw, openid, phone",
    "last_login_time" : "2015-10-07T11:33:32+0000"
  } ]
} ]
----


===== Response 400
[source,json]
----
{
  "error" : "wrong_page_parameter",
  "error_description" : "Parameter is not a positive integer."
}
----


[[_provision_profile_profileid_image_post]]
=== Add profile image for provision app
....
POST /provision/profile/{profileId}/image
....


==== Description
Adds image to profile.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Path**|**profileId** +
__required__|Profile id|string
|**Body**|**Msisdn claim request body** +
__optional__|Body should be the multipart/form-data of the attachment with key named 'image'. No Content Headers.|<<_provision_profile_profileid_image_post_msisdn_claim_request_body,Msisdn claim request body>>
|===

[[_provision_profile_profileid_image_post_msisdn_claim_request_body]]
**Msisdn claim request body**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**files** +
__optional__|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**201**|Image used by Profile|<<_provision_profile_profileid_image_post_response_201,Response 201>>
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_provision_profile_profileid_image_post_response_400,Response 400>>
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_provision_profile_profileid_image_post_response_201]]
**Response 201**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**id** +
__optional__|Profile image ID|string
|**type** +
__optional__|Type of image|string
|**urls** +
__optional__|Location of resized images|<<_provision_profile_profileid_image_post_urls,urls>>
|===

[[_provision_profile_profileid_image_post_urls]]
**urls**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**120x120** +
__optional__|URL of a resized image|string
|**240x240** +
__optional__|URL of a resized image|string
|**60x60** +
__optional__|URL of a resized image|string
|**960x960** +
__optional__|URL of a resized image|string
|===

[[_provision_profile_profileid_image_post_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request, image_missing, image_wrong_type, image_wrong_format, image_too_big, images_not_resized)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* Profile


==== Example HTTP response

===== Response 201
[source,json]
----
{
  "id" : "051a2754-dfd8-4c52-b982-4d48cbaa056b",
  "type" : "png",
  "urls" : {
    "60x60" : "http://cloudfront.net/image/profiles/051a2754-dfd8-4c52-b982-4d48cbaa056b-60x60.png",
    "120x120" : "http://cloudfront.net/image/profiles/051a2754-dfd8-4c52-b982-4d48cbaa056b-120x120.png",
    "240x240" : "http://cloudfront.net/image/profiles/051a2754-dfd8-4c52-b982-4d48cbaa056b-240x240.png",
    "960x960" : "http://cloudfront.net/image/profiles/051a2754-dfd8-4c52-b982-4d48cbaa056b-960x960.png"
  }
}
----


===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_provision_profile_profileid_image_get]]
=== Returns image URL for provision app
....
GET /provision/profile/{profileId}/image
....


==== Description
Returns image URL of specified size. Parameter size is required.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Path**|**profileId** +
__required__|Profile id|string
|**Query**|**size** +
__required__|Parameter for predefined image size. Available values are 60x60, 120x120, 240x240 and 960x960.|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|Image used by Profile|<<_provision_profile_profileid_image_get_response_200,Response 200>>
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_provision_profile_profileid_image_get_response_400,Response 400>>
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**404**|Not Found. The requested resource could not be found but may be available again in the future. Subsequent requests by the client are permissible.|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_provision_profile_profileid_image_get_response_200]]
**Response 200**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**id** +
__optional__|Profile image ID|string
|**type** +
__optional__|Type of image|string
|**urls** +
__optional__|Location of resized images|<<_provision_profile_profileid_image_get_urls,urls>>
|===

[[_provision_profile_profileid_image_get_urls]]
**urls**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**120x120** +
__optional__|URL of a resized image|string
|**240x240** +
__optional__|URL of a resized image|string
|**60x60** +
__optional__|URL of a resized image|string
|**960x960** +
__optional__|URL of a resized image|string
|===

[[_provision_profile_profileid_image_get_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request, image_size_missing_parameter, invalid_image_size)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* Profile


==== Example HTTP response

===== Response 200
[source,json]
----
{
  "id" : "051a2754-dfd8-4c52-b982-4d48cbaa056b",
  "type" : "png",
  "urls" : {
    "60x60" : "http://cloudfront.net/image/profiles/051a2754-dfd8-4c52-b982-4d48cbaa056b-60x60.png",
    "120x120" : "http://cloudfront.net/image/profiles/051a2754-dfd8-4c52-b982-4d48cbaa056b-120x120.png",
    "240x240" : "http://cloudfront.net/image/profiles/051a2754-dfd8-4c52-b982-4d48cbaa056b-240x240.png",
    "960x960" : "http://cloudfront.net/image/profiles/051a2754-dfd8-4c52-b982-4d48cbaa056b-960x960.png"
  }
}
----


===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_provision_profile_profileid_image_delete]]
=== Remove profile image for provision app
....
DELETE /provision/profile/{profileId}/image
....


==== Description
Removes image from profile.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Path**|**profileId** +
__required__|Profile id|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**204**|On success 204 code is returned.|No Content
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_provision_profile_profileid_image_delete_response_400,Response 400>>
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**404**|Not Found. The requested resource could not be found but may be available again in the future. Subsequent requests by the client are permissible.|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_provision_profile_profileid_image_delete_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* Profile


==== Example HTTP response

===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_provision_profile_profileid_notifications_get]]
=== Get profile notifications for provision app
....
GET /provision/profile/{profileId}/notifications
....


==== Description
Returns collection of profile notifications.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Path**|**profileId** +
__required__|Profile id|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|Collection of user's notifications.|< <<_provision_profile_profileid_notifications_get_response_200,Response 200>> > array
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_provision_profile_profileid_notifications_get_response_200]]
**Response 200**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**application** +
__optional__|A representation of an application +
**Example** : `{
  "id" : "dde9536c-2ca2-4fdc-a710-17ef0a58ede8",
  "label" : "FakeAnApp v2",
  "shortNames" : [ "FakeAnApp v2" ],
  "url" : "http://awesomeapp.example.com",
  "image" : {
    "45x45" : "http://mobileconnect.io/images/5a486dee-718d-4caf-8e10-c332024c4a32-45x45.jpg",
    "60x60" : "http://mobileconnect.io/images/5a486dee-718d-4caf-8e10-c332024c4a32-60x60.jpg",
    "90x90" : "http://mobileconnect.io/images/5a486dee-718d-4caf-8e10-c332024c4a32-90x90.jpg",
    "120x120" : "http://mobileconnect.io/images/5a486dee-718d-4caf-8e10-c332024c4a32-120x120.jpg",
    "240x240" : "http://mobileconnect.io/images/5a486dee-718d-4caf-8e10-c332024c4a32-240x240.jpg",
    "480x480" : "http://mobileconnect.io/images/5a486dee-718d-4caf-8e10-c332024c4a32-480x480.jpg"
  },
  "promotedImage" : {
    "282x174" : "http://mobileconnect.io/images/50c23757-9dfc-4736-af29-93dace422949-282x174.jpg"
  },
  "author" : {
    "name" : "null",
    "email" : "author@etal.io"
  },
  "createdAt" : "2017-10-02T11:55:58+0000",
  "description" : "Be a description",
  "requireLoa" : "",
  "scopeDescription" : "Reasoning about the scopes that this app is requesting",
  "scope" : [ "scope.r", "scope.w" ],
  "publisher" : {
    "name" : "Mr. Publisher",
    "email" : "author@etal.io"
  },
  "privileges" : {
    "trusted" : false,
    "authorizator" : false,
    "system" : false,
    "autogrant" : false,
    "ump" : false,
    "provision" : true,
    "provisionDescriptor" : false,
    "spUserProvision" : false
  },
  "links" : [ {
    "rel" : "author",
    "href" : "https://api.mdi.ericsson.net/v1/profile/dde9536c-2ca2-4fdc-a710-17ef0a58ede8"
  }, {
    "rel" : "application",
    "href" : "https://api.mdi.ericsson.net/v1/application/dde9536c-2ca2-4fdc-a710-17ef0a58ede8"
  }, {
    "rel" : "keys",
    "href" : "https://api.mdi.ericsson.net/v1/application/dde9536c-2ca2-4fdc-a710-17ef0a58ede8/keys"
  } ],
  "platforms" : [ "web", "ios", "android" ],
  "category" : {
    "id" : "3cf1cd0e-ddf3-4bd3-85e8-45ad9b99621d",
    "name" : "Books"
  },
  "platformUri" : [ {
    "type" : "web",
    "url" : "http:://www.app.dev"
  }, {
    "type" : "ios",
    "url" : "http:://www.app.dev1"
  }, {
    "type" : "android",
    "url" : "http:://www.app.dev2"
  } ],
  "featured" : true,
  "mcPrivileges" : {
    "trustLevel" : 2
  },
  "logoUrl" : "",
  "backgroundUrl" : "",
  "visibility" : "pub",
  "provisionerForMno" : "905538d4-fa92-4305-aa5b-38383b274544"
}`|<<_provision_profile_profileid_notifications_get_application,application>>
|**authorId** +
__optional__|Profile ID of notification author +
**Pattern** : `"[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}"`|string
|**date** +
__optional__|Date of the notification|string
|**id** +
__optional__|Notification ID +
**Pattern** : `"[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}"`|string
|**ip** +
__optional__|IP number of notification author|string
|**label** +
__optional__|Notification type|string
|**notification** +
__optional__|Notification type - deprecated|enum (login, loginFail, logout, register, passwordChange, grant, grantDeny, revoke, auth2factorSent, auth2factorFail, auth2factorSuccess)
|===

[[_provision_profile_profileid_notifications_get_application]]
**application**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**author** +
__optional__|Author info|<<_provision_profile_profileid_notifications_get_application_author,author>>
|**createdAt** +
__optional__|Applications's create date.|string
|**description** +
__optional__|Applications's description.|string
|**id** +
__optional__|Application's id. +
**Pattern** : `"[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}"`|string
|**image** +
__optional__|Set of application image links with all available sizes.|<<_provision_profile_profileid_notifications_get_application_image,image>>
|**label** +
__optional__|Application's name.|string
|**privileges** +
__optional__|List of privileges application has.|<<_provision_profile_profileid_notifications_get_application_privileges,privileges>>
|**promotedImage** +
__optional__|Application's promoted image.|<<_provision_profile_profileid_notifications_get_application_promotedimage,promotedImage>>
|**publisher** +
__optional__|Publisher info.|<<_provision_profile_profileid_notifications_get_application_publisher,publisher>>
|**requireLoa** +
__optional__|Restricts which LoA is minimal for login (eg. loa2 or loa3). Empty string means that there is no restriction.|string
|**scope** +
__optional__|Array of scopes.|< string > array
|**scopeDescription** +
__optional__|Description, reasoning about requested scopes.|string
|**shortNames** +
__optional__|Array of application short names.|< string > array
|**url** +
__optional__|URL of application's website.|string
|===

[[_provision_profile_profileid_notifications_get_application_author]]
**author**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**email** +
__optional__|Author's E-mail.|string
|**name** +
__optional__|Author's name.|string
|===

[[_provision_profile_profileid_notifications_get_application_image]]
**image**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**120x120** +
__optional__|Image link|string
|**240x240** +
__optional__|Image link|string
|**45x45** +
__optional__|Image link|string
|**480x480** +
__optional__|Image link|string
|**60x60** +
__optional__|Image link|string
|**90x90** +
__optional__|Image link|string
|===

[[_provision_profile_profileid_notifications_get_application_privileges]]
**privileges**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**authorizator** +
__optional__|Defines if application is authorizator.|boolean
|**autogrant** +
__optional__|Defines if application scopes can be autogranted.|boolean
|**provision** +
__optional__|Defines if application has provisioner privileges.|boolean
|**provisionDescriptor** +
__optional__|Defines if application has descriptor provision privilege.|boolean
|**spUserProvision** +
__optional__|Defines if application has privileges to provision users.|boolean
|**system** +
__optional__|Defines if it is system application.|boolean
|**trusted** +
__optional__|Defines if application is trusted.|boolean
|**ump** +
__optional__|Defines if application acts as an User Management Portal|boolean
|===

[[_provision_profile_profileid_notifications_get_application_promotedimage]]
**promotedImage**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**282x174** +
__optional__|Image link|string
|===

[[_provision_profile_profileid_notifications_get_application_publisher]]
**publisher**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**email** +
__optional__|Publisher's E-mail|string
|**name** +
__optional__|Publisher's name|string
|===


==== Produces

* `application/json`


==== Tags

* Profile


==== Example HTTP response

===== Response 200
[source,json]
----
[ {
  "id" : "a3a3ddbc-2a2c-4ea7-ba65-5cb04b8e2344",
  "authorId" : "dde9536c-2ca2-4fdc-a710-17ef0a58ede8",
  "date" : "2013-01-01T00:00:00+00:00",
  "type" : "login",
  "ip" : "127.0.0.1",
  "application" : {
    "_id" : "7e71ea8d-933c-41ab-ad3e-c22c94f6ab28",
    "name" : "HereBe Dragons",
    "author" : {
      "name" : "John Author",
      "email" : "author@example.tld"
    },
    "publisher" : {
      "name" : "John Publisher",
      "email" : "publisher@example.tld"
    },
    "url" : "http://awesomeapp.example.com",
    "description" : "Be a description",
    "scopeDescription" : "Reasoning about the scopes that this app is requesting",
    "scope" : [ "profile.basic.r", "profile.basic.w" ],
    "trusted" : false,
    "category" : {
      "id" : "3cf1cd0e-ddf3-4bd3-85e8-45ad9b99621d",
      "name" : "Books"
    }
  }
} ]
----


[[_provision_profile_profileid_suspend_post]]
=== Suspend profile for provision app
....
POST /provision/profile/{profileId}/suspend
....


==== Description
Suspends profile.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Path**|**profileId** +
__required__|Profile id|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**202**|On success 202 code is returned.|No Content
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===


==== Produces

* `application/json`


==== Tags

* Profile


[[_provision_profile_profileid_suspend_delete]]
=== Unsuspend profile for provision app
....
DELETE /provision/profile/{profileId}/suspend
....


==== Description
Unsuspends profile.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Path**|**profileId** +
__required__|Profile id|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**202**|On success 202 code is returned.|No Content
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===


==== Produces

* `application/json`


==== Tags

* Profile


[[_provision_profiles_post]]
=== Add new profile with provision app
....
POST /provision/profiles
....


==== Description
Adds new profile.


==== Parameters

[options="header", cols=".^2a,.^3a,.^4a"]
|===
|Type|Name|Schema
|**Body**|**Adding Profile request body** +
__optional__|<<_provision_profiles_post_adding_profile_request_body,Adding Profile request body>>
|===

[[_provision_profiles_post_adding_profile_request_body]]
**Adding Profile request body**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**EUA** +
__optional__|Accept Mobile Connect Accelerator End User Licence Agreement *DEPRECATED*|boolean
|**EULA** +
__optional__|Accept Mobile Connect Digital Identity End User Licence Agreement|<<_provision_profiles_post_eula,EULA>>
|**developerEULA** +
__optional__|Accept developer EULA|boolean
|**device** +
__optional__|Push notification device info|<<_provision_profiles_post_device,device>>
|**email** +
__optional__|Contains E-mail address.|string
|**emailVerified** +
__optional__|Flag for setting new email as verified. Only available for provisioning apps.|boolean
|**family_name** +
__optional__|Contains last name.|string
|**given_name** +
__optional__|Contains first name.|string
|**middle_name** +
__optional__|Contains middle name.|string
|**moSmsCode** +
__optional__|Verification code for MO SMS flow. This is the code sent via SMS from user to Mobile Connect Accelerator.|string
|**msisdn** +
__required__|Contains phone number|string
|**mtSmsCode** +
__optional__|Verification code for MT SMS flow. This is the code user gets from the SMS sent from Mobile Connect Accelerator|string
|**operator** +
__required__|Operator ID +
**Pattern** : `"[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}"`|string
|**optInCode** +
__optional__|Opt-in verification code|string
|**password** +
__optional__|Contains user's password and code.|<<_provision_profiles_post_password,password>>
|**shortRegistration** +
__optional__|Flag for LoA2 (short) registration. If it is set to true, password is not required for registration.|boolean
|===

[[_provision_profiles_post_eula]]
**EULA**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**issuer** +
__optional__|EULA issuer; Value can be either 'mdi' for default Mobile Connect Digital Identity EULA or MNO unique ID for custom MNO EULA.|string
|**version** +
__optional__|EULA version|string
|===

[[_provision_profiles_post_device]]
**device**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**ID** +
__optional__|Push notification device registration ID|string
|**type** +
__optional__|Push notification device type: 'android' or 'ios'|string
|===

[[_provision_profiles_post_password]]
**password**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**code** +
__optional__|Verification code (obsolete in favor of requestCode)|string
|**fresh** +
__required__|Contains new profile password|string
|**requestCode** +
__optional__|Verification code for MT SMS flow. This is the code user gets from the SMS sent from Mobile Connect Accelerator|string
|**responseCode** +
__optional__|Verification code for MO SMS flow. This is the code sent via SMS from user to Mobile Connect Accelerator.|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**201**|Authenticator linking status endpoint response.|< <<_provision_profiles_post_response_201,Response 201>> > array
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_provision_profiles_post_response_400,Response 400>>
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_provision_profiles_post_response_201]]
**Response 201**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**agreements** +
__optional__|Contains profile agreements. Available in scope.|<<_provision_profiles_post_agreements,agreements>>
|**apePreference** +
__optional__|List of prefered and available authenticators.|<<_provision_profiles_post_apepreference,apePreference>>
|**device** +
__optional__|Push notification device info|<<_provision_profiles_post_device,device>>
|**email** +
__optional__|Contains current address, possible pending address and verification status. Available in profile.email.r scope.|<<_provision_profiles_post_email,email>>
|**id** +
__optional__|Profile's ID. Available in profile.basic.r scope. +
**Pattern** : `"[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}"`|string
|**image** +
__optional__|Set of profile image links with all available sizes. Available in profile.basic.r scope.|<<_provision_profiles_post_image,image>>
|**links** +
__optional__|Profile related links|< <<_provision_profiles_post_links,links>> > array
|**msisdn** +
__optional__|Contains country code, national number (nn) and phone number. Available in profile.msisdn.r scope.|<<_provision_profiles_post_msisdn,msisdn>>
|**name** +
__optional__|Contains profile's name. Available in profile.basic.r scope.|string
|**operator** +
__optional__|Operator ID. Available in profile.operator.r scope. +
**Pattern** : `"[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}"`|string
|**status** +
__optional__|Profile account statuses|<<_provision_profiles_post_status,status>>
|===

[[_provision_profiles_post_agreements]]
**agreements**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**EULA** +
__optional__|Is Mobile Connect Accelerator End User Licence Agreement accepted|boolean
|**developerEULA** +
__optional__|Is developer EULA accepted|boolean
|===

[[_provision_profiles_post_apepreference]]
**apePreference**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**available** +
__optional__|Prefered authentication methods|<<_provision_profiles_post_apepreference_available,available>>
|**prefered** +
__optional__|Prefered authentication methods|<<_provision_profiles_post_apepreference_prefered,prefered>>
|===

[[_provision_profiles_post_apepreference_available]]
**available**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**loa2** +
__optional__|available LoA2 authentication methods.|< string > array
|**loa3** +
__optional__|available LoA2 authentication methods.|< string > array
|===

[[_provision_profiles_post_apepreference_prefered]]
**prefered**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**loa2** +
__optional__|Prefered LoA2 authentication method.|string
|**loa3** +
__optional__|Prefered LoA3 authentication method.|string
|===

[[_provision_profiles_post_device]]
**device**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**ID** +
__optional__|Push notification device registration ID|string
|**type** +
__optional__|Push notification device type: 'android' or 'ios'|string
|===

[[_provision_profiles_post_email]]
**email**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**address** +
__optional__|Currently used E-mail address|string
|**pending** +
__optional__|E-mail address waiting to be verified|string
|**verified** +
__optional__||boolean
|===

[[_provision_profiles_post_image]]
**image**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**120x120** +
__optional__|Image link|string
|**240x240** +
__optional__|Image link|string
|**60x60** +
__optional__|Image link|string
|**960x960** +
__optional__|Image link|string
|===

[[_provision_profiles_post_links]]
**links**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**href** +
__optional__|Link path|string
|**rel** +
__optional__|Link relation|string
|===

[[_provision_profiles_post_msisdn]]
**msisdn**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**country_code** +
__optional__|number
|**nn** +
__optional__|string
|**number** +
__optional__|string
|===

[[_provision_profiles_post_status]]
**status**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**suspended** +
__optional__|True if profile is currently suspended.|boolean
|===

[[_provision_profiles_post_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request, profile_already_exists, invalid_msisdn, operator_could_not_be_found, invalid_email_format, old_missing_parameter, old_password_is_incorect, new_password_missing_parameter, password_too_short, operator_2f_disabled, required_eula_mdi, required_eula_mno, EUA_only_accepts_true, invalid_device_id, invalid_device_type, inactive_operator, cannot_register_operator, optInCode_missing_parameter, mtSmsCode_missing_parameter, invalid_code, invalid_shortRegistration, short_registration_not_supported)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* Profile


==== Example HTTP response

===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_provision_profiles_delete]]
=== Delete profiles from provisioned
....
DELETE /provision/profiles
....


==== Description
Instantly removes multiple profiles without TTL.


==== Parameters

[options="header", cols=".^2a,.^3a,.^4a"]
|===
|Type|Name|Schema
|**Body**|**Delete multiple profiles request body** +
__optional__|<<_provision_profiles_delete_delete_multiple_profiles_request_body,Delete multiple profiles request body>>
|===

[[_provision_profiles_delete_delete_multiple_profiles_request_body]]
**Delete multiple profiles request body**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**msisdns** +
__required__|User's batch of MSISDN's|< string > array
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**204**|On success 204 code is returned.|No Content
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_provision_profiles_delete_response_400,Response 400>>
|**401**|Access denied. Authentication is required to proceed processing request.|No Content
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**429**|Too Many Requests. The user or application has send too many requests in a given amount of time. If field 'blocked_for' is present in the response, it contains an information about how long the request will remain blocked (in seconds).|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_provision_profiles_delete_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request, input_parameter_limit)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* Profile


==== Example HTTP response

===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_provision_profiles_profileid_grant_applicationkey_get]]
=== Grants - Collection for provision app
....
GET /provision/profiles/{profileId}/grant/{applicationKey}
....


==== Description
Returns profile grant resource for application key.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Path**|**applicationKey** +
__required__|Application key id|string
|**Path**|**profileId** +
__required__|Profile id|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|Representation of granted application|<<_provision_profiles_profileid_grant_applicationkey_get_response_200,Response 200>>
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**404**|Not Found. The requested resource could not be found but may be available again in the future. Subsequent requests by the client are permissible.|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_provision_profiles_profileid_grant_applicationkey_get_response_200]]
**Response 200**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**key** +
__optional__|Holds the key that grant was requested for.|string
|**last_login_time** +
__optional__|Holds last login time with this grant.|string
|**rel_profile** +
__optional__|Holds the profile ID of grant.|string
|**scope** +
__optional__|Holds user granted scopes for application.|string
|===


==== Produces

* `application/json`


==== Tags

* Profile


==== Example HTTP response

===== Response 200
[source,json]
----
{
  "rel_profile" : "051a2754-dfd8-4c52-b982-4d48cbaa056b",
  "key" : "536700d3d2c65b245cc887aa2b4d79ab",
  "scope" : "profile.basic.rw, application.rw, openid, offline_access",
  "last_login_time" : "2015-10-07T14:37:32+0000"
}
----


[[_provision_profiles_profileid_grant_applicationkey_delete]]
=== Grant - Revoke for provision app
....
DELETE /provision/profiles/{profileId}/grant/{applicationKey}
....


==== Description
Revokes access grant for an application, revokes only current used key for the profile grant.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Path**|**applicationKey** +
__required__|Application key id|string
|**Path**|**profileId** +
__required__|Profile id|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**204**|On success 204 code is returned.|No Content
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_provision_profiles_profileid_grant_applicationkey_delete_response_400,Response 400>>
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_provision_profiles_profileid_grant_applicationkey_delete_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* Profile


==== Example HTTP response

===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_reports_developers_get]]
=== Developers collection report
....
GET /reports/developers
....


==== Description
Reports endpoint for retrieving collection of developers. Authenticated access token is needed, only system apps can access the endpoint.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Query**|**endTime** +
__optional__|Registration time filtering start date ISO8601 format UTC|string
|**Query**|**iso** +
__optional__|Developers' country filtering lowercase|string
|**Query**|**limit** +
__optional__|Pagination parameter, indicates how many records should be returned by page. NOTE: Default value is 20 and Max value is 50|string
|**Query**|**page** +
__optional__|Pagination parameter, indicates which page is to be returned from database. NOTE: Default value is 1.|number
|**Query**|**startTime** +
__optional__|Registration time filtering start date ISO8601 format UTC|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|Collection of Developers.|< <<_reports_developers_get_response_200,Response 200>> > array
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_reports_developers_get_response_400,Response 400>>
|**401**|Access denied. Authentication is required to proceed processing request.|No Content
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_reports_developers_get_response_200]]
**Response 200**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**country** +
__optional__|Contains name of the country determined based on msisdn.|string
|**developerRegistration** +
__optional__|Contains date of registration as a developer (accepting developerEULA for the first time) ISO8601 format UTC.|string
|**email** +
__optional__|Contains verified email address.|string
|**latestEULA** +
__optional__|Indicates if the latest developerEULA was accepted.|boolean
|**msisdn** +
__optional__|Contains developers's msisdn.|string
|**name** +
__optional__|Contains profile's name.|string
|===

[[_reports_developers_get_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request, wrong_page_parameter, wrong_limit_parameter, limit_exceeds_max, invalid_date_format, invalid_iso)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* Reports


==== Example HTTP response

===== Response 200
[source,json]
----
[ {
  "name" : "Profile name",
  "country" : "Slovenia",
  "msisdn" : "+38611111111",
  "email" : "email@email.com",
  "developerRegistration" : "2015-10-08T09:44:51+0000",
  "latestEULA" : true
}, {
  "name" : "Another profile name",
  "country" : "Canada",
  "msisdn" : "+16471111111",
  "email" : null,
  "developerRegistration" : "2015-10-09T11:22:33+0000",
  "latestEULA" : false
} ]
----


===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_reports_developers_cache_clear_external_attr_desc_get]]
=== Clear external attributes descriptor cache
....
GET /reports/developers (cache_clear_external_attr_desc)
....


==== Description
Clears external attributes plugin descriptor cache.


==== Parameters

[options="header", cols=".^2a,.^3a,.^9a,.^4a"]
|===
|Type|Name|Description|Schema
|**Query**|**mnoId** +
__required__|MNO ID or 'all' for all MNOs|string
|===


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|On success 200 code is returned.|No Content
|**400**|Invalid input. For extra information response will contain extra information (error and error_description) in json response. Error description will be a translated string.|<<_reports_developers_cache_clear_external_attr_desc_get_response_400,Response 400>>
|**404**|Not Found. The requested resource could not be found but may be available again in the future. Subsequent requests by the client are permissible.|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_reports_developers_cache_clear_external_attr_desc_get_response_400]]
**Response 400**

[options="header", cols=".^3a,.^4a"]
|===
|Name|Schema
|**error** +
__optional__|enum (bad_request)
|**error_description** +
__optional__|string
|===


==== Produces

* `application/json`


==== Tags

* Internal


==== Example HTTP response

===== Response 400
[source,json]
----
{
  "error" : "bad_request",
  "error_description" : "Parameters or data missing."
}
----


[[_scopes_get]]
=== Get scopes collection
....
GET /scopes
....


==== Description
Returns scopes collection.


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|Collection of scopes.|< <<_scopes_get_response_200,Response 200>> > array
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_scopes_get_response_200]]
**Response 200**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**description** +
__optional__|Contains scope description. (obsolete, pending removal)|string
|**label** +
__optional__|Contains scope label.|string
|**name** +
__optional__|Contains scope name.|string
|===


==== Produces

* `application/json`


==== Tags

* Scope


==== Example HTTP response

===== Response 200
[source,json]
----
[ {
  "name" : "profile.anonymous.r",
  "label" : "view your profile information",
  "description" : "view your profile information"
} ]
----


[[_status_get]]
=== Mobile Connect Accelerator API status
....
GET /status
....


==== Description
Returns Mobile Connect Accelerator API status.


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|API status.|<<_status_get_response_200,Response 200>>
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_status_get_response_200]]
**Response 200**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**body** +
__optional__|Human readable status description|string
|**created_on** +
__optional__|URL of application's website|string
|**status** +
__optional__|Status|string
|===


==== Produces

* `application/json`


==== Tags

* Status


==== Example HTTP response

===== Response 200
[source,json]
----
{
  "status" : "good",
  "body" : "Service fully operational",
  "created_on" : "2014-01-07T12:00:00Z"
}
----


[[_version_get]]
=== Get Mobile Connect Accelerator application version information
....
GET /version
....


==== Description
Returns Mobile Connect Accelerator application version information, changelog.


==== Responses

[options="header", cols=".^2a,.^14a,.^4a"]
|===
|HTTP Code|Description|Schema
|**200**|Collection of versions.|< <<_version_get_response_200,Response 200>> > array
|**403**|Access denied. The request was a valid request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference. On servers where authentication is required, this commonly means that the provided credentials were successfully authenticated but that the credentials still do not grant the client permission to access the resource (e.g., a recognized user attempting to access restricted content).|No Content
|**500**|Internal error. Something went wrong when storing to database.|No Content
|===

[[_version_get_response_200]]
**Response 200**

[options="header", cols=".^3a,.^11a,.^4a"]
|===
|Name|Description|Schema
|**EULA** +
__optional__|EULA (End User Licence Agreement) version|string
|**changelog** +
__optional__|Application changelog for selected version.|string
|**currentVersion** +
__optional__|Current application version number.|string
|**minVersion** +
__optional__|Minimum application version number.|string
|**platform** +
__optional__|Application platform id.|string
|===


==== Produces

* `application/json`


==== Tags

* Version


==== Example HTTP response

===== Response 200
[source,json]
----
[ {
  "platform" : "web",
  "currentVersion" : "1.01",
  "minVersion" : "0.01",
  "changelog" : "Lorem ipsum web.",
  "EULA" : "1.0"
} ]
----



