
---
weight: 11
title: fabric
---
# fabric

## AddressAPI.GetAddress
```go
package main
import (
  "github.com/micro/clients/go/client"
  fabric_proto "/fabric/proto"
)
func main() {
  c := client.NewClient(nil)
  req := fabric_proto.GetAddressRequest{}
  rsp := fabric_proto.GetAddressResponse{}
  if err := c.Call("go.micro.srv.fabric", "AddressAPI.GetAddress", req, &rsp); err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(rsp)
}
```
```javascript
// To install "npm install --save @microhq/ng-client"
import { Component, OnInit } from "@angular/core";
import { ClientService } from "@microhq/ng-client";
@Component({
  selector: "app-example",
  templateUrl: "./example.component.html",
  styleUrls: ["./example.component.css"]
})
export class ExampleComponent implements OnInit {
  constructor(private mc: ClientService) {}
  ngOnInit() {
    this.mc
      .call("go.micro.srv.fabric", "AddressAPI.GetAddress", {})
      .then((response: any) => {
        console.log(response)
      });
  }
}
```
 Get Address allows a user to retrieve field level details of an address found using the SearchAddress API.
	**anzctl** :
	```
    anzctl selfservice address get -d &#39;
    {
    &#34;identifier&#34;: &#34;ABCD123456789&#34;
    }
    &#39; -o json
	```
### Request Parameters
Name |  Type | Description
--------- | --------- | ---------
id | string | 

### Response Parameters
Name |  Type | Description
--------- | --------- | ---------
address_identifiers | AddressResult | 


### Message AddressResult
Name |  Type | Description
--------- | --------- | ---------


### 
<aside class="success">
Remember — a happy kitten is an authenticated kitten!
</aside>

## PartyAPI.GetParty
```go
package main
import (
  "github.com/micro/clients/go/client"
  fabric_proto "/fabric/proto"
)
func main() {
  c := client.NewClient(nil)
  req := fabric_proto.GetPartyRequest{}
  rsp := fabric_proto.GetPartyResponse{}
  if err := c.Call("go.micro.srv.fabric", "PartyAPI.GetParty", req, &rsp); err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(rsp)
}
```
```javascript
// To install "npm install --save @microhq/ng-client"
import { Component, OnInit } from "@angular/core";
import { ClientService } from "@microhq/ng-client";
@Component({
  selector: "app-example",
  templateUrl: "./example.component.html",
  styleUrls: ["./example.component.css"]
})
export class ExampleComponent implements OnInit {
  constructor(private mc: ClientService) {}
  ngOnInit() {
    this.mc
      .call("go.micro.srv.fabric", "PartyAPI.GetParty", {})
      .then((response: any) => {
        console.log(response)
      });
  }
}
```
 Get Party allows a user to retrieve information stored in their OCV (One Customer View) profile via Fabric Self Service. The party information retrieved will be for the ocvID located in the JWT token presented for authentication.
 **anzctl** : `~$ anzctl selfservice party get -o json`
### Request Parameters
Name |  Type | Description
--------- | --------- | ---------

### Response Parameters
Name |  Type | Description
--------- | --------- | ---------
legal_name | Name | 
residential_address | Address | 
mailing_address | Address | 


### Message Name
Name |  Type | Description
--------- | --------- | ---------


### Message Address
Name |  Type | Description
--------- | --------- | ---------


### 
<aside class="success">
Remember — a happy kitten is an authenticated kitten!
</aside>

## PreferencesAPI.GetPreferences
```go
package main
import (
  "github.com/micro/clients/go/client"
  fabric_proto "/fabric/proto"
)
func main() {
  c := client.NewClient(nil)
  req := fabric_proto.GetPreferencesRequest{}
  rsp := fabric_proto.GetPreferencesResponse{}
  if err := c.Call("go.micro.srv.fabric", "PreferencesAPI.GetPreferences", req, &rsp); err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(rsp)
}
```
```javascript
// To install "npm install --save @microhq/ng-client"
import { Component, OnInit } from "@angular/core";
import { ClientService } from "@microhq/ng-client";
@Component({
  selector: "app-example",
  templateUrl: "./example.component.html",
  styleUrls: ["./example.component.css"]
})
export class ExampleComponent implements OnInit {
  constructor(private mc: ClientService) {}
  ngOnInit() {
    this.mc
      .call("go.micro.srv.fabric", "PreferencesAPI.GetPreferences", {})
      .then((response: any) => {
        console.log(response)
      });
  }
}
```
 Get Preferences allows for storage and retrieval of Fabric user preferences, in particular, related to user facing features of the ANZx mobile app.
 **anzctl** : `~$ anzctl selfservice preferences get -o json
### Request Parameters
Name |  Type | Description
--------- | --------- | ---------

### Response Parameters
Name |  Type | Description
--------- | --------- | ---------
preferences | Preferences |  The current preferences value for the given user/persona.
etag | string |  This is a hash of the value of the returned preferences. This must be stored by the client and sent when using the UpdatePreferences endpoint.


### Message Preferences
Name |  Type | Description
--------- | --------- | ---------


### 
<aside class="success">
Remember — a happy kitten is an authenticated kitten!
</aside>

## ProfilePictureAPI.GetProfilePicture
```go
package main
import (
  "github.com/micro/clients/go/client"
  fabric_proto "/fabric/proto"
)
func main() {
  c := client.NewClient(nil)
  req := fabric_proto.GetProfilePictureRequest{}
  rsp := fabric_proto.GetProfilePictureResponse{}
  if err := c.Call("go.micro.srv.fabric", "ProfilePictureAPI.GetProfilePicture", req, &rsp); err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(rsp)
}
```
```javascript
// To install "npm install --save @microhq/ng-client"
import { Component, OnInit } from "@angular/core";
import { ClientService } from "@microhq/ng-client";
@Component({
  selector: "app-example",
  templateUrl: "./example.component.html",
  styleUrls: ["./example.component.css"]
})
export class ExampleComponent implements OnInit {
  constructor(private mc: ClientService) {}
  ngOnInit() {
    this.mc
      .call("go.micro.srv.fabric", "ProfilePictureAPI.GetProfilePicture", {})
      .then((response: any) => {
        console.log(response)
      });
  }
}
```
 Get Profile Picture allows an ANZ user to retrieve a URL to their profile picture via Fabric Self Service.
 **anzctl** : `~$ anzctl selfservice profile getpicture -o json
### Request Parameters
Name |  Type | Description
--------- | --------- | ---------

### Response Parameters
Name |  Type | Description
--------- | --------- | ---------
url | string | 


### 
<aside class="success">
Remember — a happy kitten is an authenticated kitten!
</aside>

## ProfilePictureAPI.LinkProfilePicture
```go
package main
import (
  "github.com/micro/clients/go/client"
  fabric_proto "/fabric/proto"
)
func main() {
  c := client.NewClient(nil)
  req := fabric_proto.LinkProfilePictureRequest{}
  rsp := fabric_proto.LinkProfilePictureResponse{}
  if err := c.Call("go.micro.srv.fabric", "ProfilePictureAPI.LinkProfilePicture", req, &rsp); err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(rsp)
}
```
```javascript
// To install "npm install --save @microhq/ng-client"
import { Component, OnInit } from "@angular/core";
import { ClientService } from "@microhq/ng-client";
@Component({
  selector: "app-example",
  templateUrl: "./example.component.html",
  styleUrls: ["./example.component.css"]
})
export class ExampleComponent implements OnInit {
  constructor(private mc: ClientService) {}
  ngOnInit() {
    this.mc
      .call("go.micro.srv.fabric", "ProfilePictureAPI.LinkProfilePicture", {})
      .then((response: any) => {
        console.log(response)
      });
  }
}
```
 Link Profile Picture allows an ANZ user to store the identifier for a linked profile picture against their persona in Fabric Self Service.
	**anzctl** :
	```
	anzctl selfservice profile linkpicture -d &#39;
	{
	&#34;identifier&#34;: &#34;9fff987a-22f2-11eb-adc1-0243da120001&#34;
	}&#39; -o json
	```
### Request Parameters
Name |  Type | Description
--------- | --------- | ---------
identifier | string |  The unique identifier of the picture to be linked to a persona

### Response Parameters
Name |  Type | Description
--------- | --------- | ---------


### 
<aside class="success">
Remember — a happy kitten is an authenticated kitten!
</aside>

## AddressAPI.SearchAddress
```go
package main
import (
  "github.com/micro/clients/go/client"
  fabric_proto "/fabric/proto"
)
func main() {
  c := client.NewClient(nil)
  req := fabric_proto.SearchAddressRequest{}
  rsp := fabric_proto.SearchAddressResponse{}
  if err := c.Call("go.micro.srv.fabric", "AddressAPI.SearchAddress", req, &rsp); err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(rsp)
}
```
```javascript
// To install "npm install --save @microhq/ng-client"
import { Component, OnInit } from "@angular/core";
import { ClientService } from "@microhq/ng-client";
@Component({
  selector: "app-example",
  templateUrl: "./example.component.html",
  styleUrls: ["./example.component.css"]
})
export class ExampleComponent implements OnInit {
  constructor(private mc: ClientService) {}
  ngOnInit() {
    this.mc
      .call("go.micro.srv.fabric", "AddressAPI.SearchAddress", {})
      .then((response: any) => {
        console.log(response)
      });
  }
}
```
 Search Address allows a user to search for an Australian postal address using an underlying service called QAS (Quick Address Service), that returns a list of candidate results.
	**anzctl** :
	```
	anzctl selfservice address search -d &#39;
	{
	&#34;query&#34;: &#34;15 Smith Street&#34;,
	&#34;limit&#34;: 5
	}
	&#39; -o json
	```
### Request Parameters
Name |  Type | Description
--------- | --------- | ---------
query | string | 
limit | int32 | 

### Response Parameters
Name |  Type | Description
--------- | --------- | ---------
address_identifiers | AddressIdentifier | 


### Message AddressIdentifier
Name |  Type | Description
--------- | --------- | ---------


### 
<aside class="success">
Remember — a happy kitten is an authenticated kitten!
</aside>

## PartyAPI.UpdateParty
```go
package main
import (
  "github.com/micro/clients/go/client"
  fabric_proto "/fabric/proto"
)
func main() {
  c := client.NewClient(nil)
  req := fabric_proto.UpdatePartyRequest{}
  rsp := fabric_proto.UpdatePartyResponse{}
  if err := c.Call("go.micro.srv.fabric", "PartyAPI.UpdateParty", req, &rsp); err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(rsp)
}
```
```javascript
// To install "npm install --save @microhq/ng-client"
import { Component, OnInit } from "@angular/core";
import { ClientService } from "@microhq/ng-client";
@Component({
  selector: "app-example",
  templateUrl: "./example.component.html",
  styleUrls: ["./example.component.css"]
})
export class ExampleComponent implements OnInit {
  constructor(private mc: ClientService) {}
  ngOnInit() {
    this.mc
      .call("go.micro.srv.fabric", "PartyAPI.UpdateParty", {})
      .then((response: any) => {
        console.log(response)
      });
  }
}
```
 Update Party allows a user to update information stored in their OCV (One Customer View) profile via Fabric Self Service. The party information updated will be for the ocvID located in the JWT token presented for authentication.
 **anzctl** :
 ```
 anzctl selfservice party update -d &#39;
 {
 PAYLOAD TBD
 }
 &#39; -o json
 ```
### Request Parameters
Name |  Type | Description
--------- | --------- | ---------
residential_address | Address | 
mailing_address | Address | 

### Response Parameters
Name |  Type | Description
--------- | --------- | ---------


### Message Address
Name |  Type | Description
--------- | --------- | ---------


### 
<aside class="success">
Remember — a happy kitten is an authenticated kitten!
</aside>

## PreferencesAPI.UpdatePreferences
```go
package main
import (
  "github.com/micro/clients/go/client"
  fabric_proto "/fabric/proto"
)
func main() {
  c := client.NewClient(nil)
  req := fabric_proto.UpdatePreferencesRequest{}
  rsp := fabric_proto.UpdatePreferencesResponse{}
  if err := c.Call("go.micro.srv.fabric", "PreferencesAPI.UpdatePreferences", req, &rsp); err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(rsp)
}
```
```javascript
// To install "npm install --save @microhq/ng-client"
import { Component, OnInit } from "@angular/core";
import { ClientService } from "@microhq/ng-client";
@Component({
  selector: "app-example",
  templateUrl: "./example.component.html",
  styleUrls: ["./example.component.css"]
})
export class ExampleComponent implements OnInit {
  constructor(private mc: ClientService) {}
  ngOnInit() {
    this.mc
      .call("go.micro.srv.fabric", "PreferencesAPI.UpdatePreferences", {})
      .then((response: any) => {
        console.log(response)
      });
  }
}
```
 Update Preferences allows an ANZ user to update key/value based preferences stored in Fabric Self Service.
 **anzctl** :
 ```
	anzctl selfservice preferences update -d &#39;
	{
	&#34;preferences&#34;: {
	&#34;values&#34;: {
	&#34;key&#34;: {
	&#34;stringList&#34;: {
	&#34;list&#34;: [&#34;a&#34;, &#34;b&#34;, &#34;c&#34;]
	}
	},
	&#34;anotherKey&#34;: {
	&#34;stringValue&#34;: &#34;value&#34;
	},
	&#34;aThirdKey&#34;: {
	&#34;booleanList&#34;: {
	&#34;list&#34;: [true, false, true]
	}
	}
	}
	},
	&#34;etag&#34;: &#34;eTagFromLastKnownPreferences&#34;
	}
	&#39; -o json
 ```
### Request Parameters
Name |  Type | Description
--------- | --------- | ---------
preferences | Preferences | 
etag | string |  This hash must match the hash of the preferences current stored server side (returned from GetPreferences) or this request will be rejected.

### Response Parameters
Name |  Type | Description
--------- | --------- | ---------
etag | string |  This is a hash of the value of the newly stored preferences and must be stored by the client and sent when next using the UpdatePreferences endpoint.


### Message Preferences
Name |  Type | Description
--------- | --------- | ---------


### 
<aside class="success">
Remember — a happy kitten is an authenticated kitten!
</aside>

