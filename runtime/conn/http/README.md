# HTTP GQL-server Connector

The HTTP Connector provides a connection to a GQL-server over an HTTP connection. The connector implements a simple HTTP-based protocol that enables the TCK Runtime to submit requests to a GQL-server, and receive responses. This connection and protocol is not intended to be used for anything other than servicing TCK Runtime requests.

The GQL-server is required to support the following three HTTP resources. The URLs documented are prefixed with the base URL, which is different for each environment.

## Login

Requests made against a GQL-implementation must be authenticated. To authenticate a request, must include an _authentication ID_ as part of the request parameters. The following HTTP resource is used by clients to retrieve an authentication ID for a given principal, using a given password.

**URL Structure** `/auth/id`

**Method** POST

**Request**
```
{
    "principal_id": "tckuser",
    "password": "****"
}
```
**Response**
```
{
    "auth_id": "**************"
}
```
**Status**

| Code | Description                   |
|------|-------------------------------|
| 201  | **Created** authentication ID |
| 400 | Bad Request |
| 500 | Internal Server Error |

## Create Session

Create a new GQL-session.

**URL Structure** `/session`

**Method** POST

**Request Body** Empty

**Response Body**
```
{
    "session_id": "********"
}
```
**Status**

| Code | Description            |
|------|------------------------|
| 201  | **Created** session ID |
| 400 | Bad Request            |
| 500 | Internal Server Error  |

## Delete Session

Delete a GQL-session. The <session_id> is replaced with the session_id.

**URL Structure** `/session/<session_id>`

**Method** DELETE

**Request Headers**

| Name | Value     |
| ---- |-----------|
| Authorization | <auth_id> |

**Request Body** Empty

**Response Body** Empty

**Status**

| Code | Description            |
|------|------------------------|
| 200  | **OK** session deleted |
| 400  | Bad Request            |
| 500  | Internal Server Error  |

## Submit a request

Submit a request to the GQL-server. The <session_id> is replaced with the session_id.

**URL Structure** `/request`

**Method** POST

**Request Headers**

| Name | Value        |
| ---- |--------------|
| Authorization | <auth_id>    |
| Session | <session_id> |

### Binding Table Result Example

**Request Body**
The request is a multipart request. The first part is a unicode text string containing the query. The second part is a JSON document containing parameters.
```
Part 1 =======================================

MATCH (p:Person) WHERE p.name = $name RETURN p, p.name AS name, p.age AS age

Part 2 =======================================

{
    "name": "John Doe"
}
```
**Response Headers**

| Name | Value | Description                                  |
| ---- | ----- |----------------------------------------------|
| Content-Type | application/table+json | The response contains a binding table result |

**Response Body**
The format of the binding table is as follows: First, record_type defines the name and type of each column in the table. Second is a list (JSON array) of rows. Each row contains a list of columns, one for each entry in the record_type, and in the same order. The property named "gid" is the global identifier for the object, which in this example is a NODE.
```
{
    "gql_status": "00000",
    "result": {
        "record_type": [
            {
                "field_name": "p",
                "field_type": "NODE"
            },
            {
                "field_name": "name",
                "field_type": "STRING"
            },
            {
                "field_name": "age",
                "field_type": "INT",
            }
        ],
        [
            [
                {
                    "gid": "xxx",
                    "labels": ["Person"],
                    "properties": [
                        {
                            "type": "STRING",
                            "name": "name",
                            "value": "John Doe"
                        },
                        {
                            "type": "INT",
                            "name": "age",
                            "value": 35
                        }
                    ]
                },
                {
                    "value": "John Doe",
                },
                {
                    "value": 35
                }
            ]
        ]
    }
}
```

**Status**

| Code | Description                         |
|------|-------------------------------------|
| 201  | **OK** request created and executed |
| 400  | Bad Request                         |
| 500  | Internal Server Error               |

### Value Result Example

**Request Body**
The request is a multipart request. The first part is a unicode text string containing the query. The second part is a JSON document containing parameters.
```
Part 1 =======================================

MATCH (p:Person) RETURN COUNT(p)

Part 2 =======================================
```
**Response Headers**

| Name | Value                  | Description                          |
| ---- |------------------------|--------------------------------------|
| Content-Type | application/value+json | The response contains a value result |

**Response Body**
```
{
    "gql_status": "00000",
    "result": {
        "type": "INT",
        "value": 1
    }
}
```

**Status**

| Code | Description                         |
|------|-------------------------------------|
| 201  | **OK** request created and executed |
| 400  | Bad Request                         |
| 500  | Internal Server Error               |

### Omitted Result Example

**Request Body**
The request is a multipart request. The first part is a unicode text string containing the query. The second part is a JSON document containing parameters.
```
Part 1 =======================================

INSERT (p:Person {name: "Jimmy Dean", age: 101})

Part 2 =======================================
```
**Response Headers**

| Name | Value                    | Description                             |
| ---- |--------------------------|-----------------------------------------|
| Content-Type | application/omitted+json | The response contains an omitted result |

**Response Body**
```
{
    "gql_status": "00001",
}
```

**Status**

| Code | Description                         |
|------|-------------------------------------|
| 201  | **OK** request created and executed |
| 400  | Bad Request                         |
| 500  | Internal Server Error               |
