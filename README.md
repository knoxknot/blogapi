# Building a Blog API in Go  
---
This repository describes building a blog API to post, retrieve, update and delete contents. I have also written an [article](https://medium.com/@nwoyesamuelc/in-memory-data-store-api-development-with-go-49c76518485e) on medium that ties to this repository.

Building an Image
---
change into the blogapi directory  
**Run**: 
<code>
docker build -t knoxknot/blogapi:$TAG
</code>

Testing the Endpoints Using Curl  
---
<code>

- curl -X GET http://localhost:8080/api/v1/articles | jq .          #to retrieve all articles.  
- curl -X GET http://localhost:8080/api/v1/articles/2 | jq .        #to retrieve a specific article by id.
- curl -X POST http://localhost:8080/api/v1/articles -d '{"ID": "4","Title": "Go: Clean, Readable and Fast","Content": "It is clean in design, easy to read  and outputs very fast."}' -H "Content-Type: application/json" | jq .       #to create an article.
- curl -X PUT http://localhost:8080/api/v1/articles/4 -d '{"ID":"4","Title":"Go is Clean and Readable","Content":"It is clean in design, easy to read  and outputs very fast."}' -H "Content-Type: application/json" | jq .        #to update a specific article with a given id
- curl -X DELETE http://localhost:8080/api/v1/articles/4 | jq .           #to delete an article by id 

</code>