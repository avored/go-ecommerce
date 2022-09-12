# Go E commerce Backend

We are using entgo package as an orm for GO

How to generate a schema

    go run entgo.io/ent/cmd/ent init AdminUser
    go generate ./ent



Oauth2 set up in progress

`http://localhost:8080/oauth2/token?grant_type=client_credentials&client_id=000000&client_secret=999999&scope=read
