# bruno-api-testing
Demo how to test API using Bruno   

Bruno is an open-source alternative for API testing. We can use Bruno instead of Postman, Insomnia, etc.  

Using Bruno we can manage the requests, with facility, in separate files. There is one file for each request. In Postman, there is a unique JSON file for all tests. For example, if we change just the value of variables, we have that generate a new collection and replace the whole JSON file. Normally this is the process.    


![each-file](./image/one-file-for-each-request.png)


Instead of, storing collections in the cloud, and should be necessary to create an account to use the toll, Bruno uses the file system locally and we can use the version control system and all teams have access to API.  

> [!TIP]  
> [Read a Bruno Manifesto](https://docs.usebruno.com/introduction/manifesto)   


##  Bru lang    

Is a DSL with readable syntax and it is easy to maintain. From Bruno UI, automatically creates a file with the extension `.bru` for each request in your collection directory.

Let's see an example, first create a test using Bruno UI:  

![bruno-ui](./image/bruno-get.png)



The file `.bru`, with syntax Bru Lang generated from the test above.  

Now, we can see the bru lang:   

![bru-lang](./image/bru-lang-get.png)


If the request was POST method, will be added a new field in the file:   
![payload](./image/payload.png)  



## Product Area     

With tools like Bruno, we can test functionality, business rules, and expected responses, and use this information in prototypes. But it can be possible if we have good documentation. Then, we can write descriptions for each request, Bruno accepting markdown.    

Example in the UI:  

![docs-ui](./image/doc.png)   

Like the example below, in the file `.bru` the field doc has been created.  

![docs-in-file](./image/docfile.png) 

## CLI    

We can run tests using CLI. It is a great feature for CI/CD.  

Just install node, and Bruno CLI using the command: ` npm install -g @usebruno/cli` . Then, execute the tests using `bru run`. See the results below:   

![docs-in-file](./image/docfile.png)   

For all options, run `bru run --help`.   


> [!TIP]  
> [Check all possibilities](https://docs.usebruno.com/bru-cli/overview)  

