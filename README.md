# go-hubspot-api

A Golang Wrapper for the [Hubspot API](https://developers.hubspot.com/docs/api/overview)

# Install

```console
$ go get github.com/dinistavares/go-hubspot-api
```

# Usage

Import the Hubspot package.

```go
import "github.com/dinistavares/go-hubspot-api"
```

Create a new Hubspot Client and use the provided services.

```go
  client := hubspot.New()

  // Get a Hubspot contact
  contact, _, err := s.client.Contacts.Get("8396967887", nil)
```

## Authenticate

```go
import (
  "github.com/dinistavares/go-hubspot-api"
)

func main(){
  accessToken := "xxxxxxx"

  client := hubspot.New()
  client.Authenticate(accessToken)

  // Get a Hubspot contact
  contact, _, err := s.client.Contacts.Get("8396967887", nil)
}
```

# Examples

### Contacts

**Search contact by email**
```go
  // Create search filter
  filter := hubspot.SearchFilter{
    PropertyName: "email",
    Value:        "jane.doe@acme.com",
    Operator:     hubspot.SearchOperatorEquals,
  }

  // Set which contact properties should be returned
  properties := []string{"firstname", "lastname", "email", "phone", "listMemberships"}

  // Create search request and set properties
  search := hubspot.SearchRequest{
    Properties:   &properties,
  }

  // Create filter group and add filter
  filterGroup := search.CreateNewFilterGroup()
  filterGroup.AddFilter(filter)

  // Search contacs
  data, _, err := s.client.Contacts.Search(&search)
```

### Associations

**Get all deals associated with a contact**
```go
  // Get all deals associated to the contact with ID '8396967887'
  associationIDs, _, err := s.client.Associations.List(hubspot.ObjectTypeIDContacts, "8396967887", hubspot.ObjectTypeIDDeals, nil)
```

### Deals

**Get a deal by ID**
```go
  // Define which properties should be returned
  properties := []string{"dealId", "deal_currency_code", "dealstage", "hs_priority"}

  // Create and set query options
  opts := hubspot.Query()
  opts.SetProperties(properties)

  // Get deal with ID '14970082502'
  data, _, err := s.client.Deals.Get("14970082502", opts)
```
