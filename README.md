# Go AvoRed an E commerce App 
 

How To run Go Cobra CLI command

    go run main.go avored


How To Install and Create an certificate for localhost

     
     brew install mkcert
     mkcert -install
     
     cd DOCUMENT_ROOT_OF_YOUR_APP
     mkcert localhost



How to create a ent go model

    go run entgo.io/ent/cmd/ent init User

    // After adding or updating model we need to re generate the ./ent
    go generate ./ent
