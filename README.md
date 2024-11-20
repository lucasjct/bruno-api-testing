# bruno-api-testing
How to test API using Bruno    

> [!TIP]
>  Sometimes, it is a good idea to create tests in the same language used by the team, it can increase the collaboration and improve quality.   In this case, Bruno may be more appropriate for exploratory API testing.    
> [Check an example how to test api with Golang](test/readme.md)   


Bruno is an open-source alternative to API testing. We can use Bruno instead of Postman, Insomnia, etc. It is available for Linux, Windows and MacOS. To download it, access [Bruno](https://www.usebruno.com/).

With Bruno we can manage the requests with facility, in separated files. There is one file for each request. In Postman, there is a single JSON file for all tests and all requests. For example, if we change just one value of variables, we have that generate a new collection and replace the whole JSON file. Normally, this is the process available on free plan.    

<p align="center">
  <img src="./image/one-file-for-each-request.png" />
</p>



Instead of, storing collections in the cloud, and create account to use the toll, Bruno has used the filesystem to be abble we use the version control system, like Git, and all teams have access to API.  

> [!TIP]  
> [Read a Bruno Manifesto](https://docs.usebruno.com/introduction/manifesto)   


##  Bru lang    

Is a domain-specific language (DSL) with readable syntax and it is easy to maintain. From Bruno UI, automatically has created the files with the extension `.bru` for each request in your directory.

Let's see an example, first create a test using Bruno UI:  

<p align="center">
  <img src="./image/bruno-get.png" />
</p>



The file `.bru`, with syntax Bru Lang generated from the test above.  

Now, we can see the bru lang:   

<p align="center">
  <img src="./image/bru-lang-get.png" />
</p>


If the request was POST method, will be added a new field in the file:   


<p align="center">
  <img src="./image/payload.png" />
</p>  



## Product Area     

With tools like Bruno, we can test functionality, business rules, and expected responses, assertions, and use this information in prototypes. But it can be possible if we have good documentation. Then, we can write descriptions for each request, because Bruno has accepted markdown.    

Example in the UI:    


<p align="center">
  <img src="./image/doc.png" />
</p> 



Like the example below, in the file `.bru` the field in the doc has been created.  

<p align="center">
  <img src="./image/docfile.png" />
</p> 



## CLI    

We can run tests using CLI. It is a great feature to CI/CD and can be implemented in pipelines. See a example here: [pipeline with Bruno](https://github.com/lucasjct/bruno-api-testing/actions/workflows/run-test-bruno.yml).  

Just install Node, and Bruno CLI using the command: ` npm install -g @usebruno/cli` . Then, execute the tests using `bru run`. See the results below:   


<p align="center">
  <img src="./image/cli.png" />
</p>  


See all options, run `bru run --help`.   


> [!TIP]  
> [Check the documentation to CLI specification](https://docs.usebruno.com/bru-cli/overview)  

