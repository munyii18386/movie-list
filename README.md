# PURPOSE: 
> Develop a fully functional application by utilizing basic knowledge, tools and techniques acquired in prior INFO classes. Gain first-hand knowledge on  how to implement a complete project from scratch that will result in a better understanding of the inner workings of web applications. The aim is to be comfortable building projects by implementing common techniques such as directory structures, URL routers and utilizing MVC patterns.

# SCOPE:

## Planning and design:
> - Learn how to leverage MVC pattern and designing program execution flow.
## Application Features:
> - Explore Web working principles with a focus on HTTP, DNS resolution flow.
>
> - Explore Websockets and understand how to allow browsers to have full-duplex communications       within servers.
>
> - Get a deeper understanding of routing and how to map URLs to processing logic.
>
>  - Understand  vulnerabilities and what precautions to take in order to prevent attacks.
>
## Logic Functionality:
> - Practice designing generic controllers, and how to handle requests and return responses after    inheriting from an object handler.
## Auxiliary Functionality:
> - Better understand common functionality such as log processing, information configuration etc.

# SCHEDULE:

|Week |Milestone                             |Objectives                                      |
|-----|--------------------------------------|------------------------------------------------|
|0    |Basic Setup and Configuration         |Identify project tools, software and/or platform                    |
|1    |Web foundation                        |Understand HTTP operating mechanism                                 |
|     |                                      |Review client side basics                                           |
|2    |Deployment and maintenance            |Generate logs and the process of logging                            |
|     |                                      |How to deal and manage runtime errors                               |
|3    |Error handling, debugging and testing |Design application with proper error handling, test cases etc       |
|4    |Text files                            |Practice how to produce and/or handle received text content, including but                                                                         not limited to strings, numbers, JSON, XML|
|5    |HTTP form                             |Deployment and configuration                                        |
|     |                                      |Explore the ability to communicate between clients and servers      |
|6    |Databases                             |Review CRUD                                                         |
|     |                                      |Analyze pros and cons of SQL vs NoSQl within project scope          |
|7    |Data storage and session              |Understand how to control the whole process of viewing web for users|
|8    |Web Services                          |Explore how to design web services that are platform independent    |
|     |                                      |Explore socket communication                                        |
|     |                                      |Explore REST architecture                                           |
|9    |Final Project                         |Review and apply any stretch goals as applicable                    |
|10   |Final Project                         |Review and apply any stretch goals as applicable                    |
# EXPECTED OUTCOME:
> At the end of the schedule it is expected that there will be a working application in place that meets the following criteria: A web API server. a fully functional client interface and  appropriate database usage

# EVALUATION/ASSESSMENT METHOD:
> Once a week, there will be scheduled  in-person hourly meeting. During this meeting, I will have to showcase or explain the progress of the project. To narrow down the scope of the meeting, I have itemized what will be reviewed and/or assessed.  Any feedback and/or suggestions will be incorporated into the next week. 

|Week |                                                                                                                 |
|-----|-----------------------------------------------------------------------------------------------------------------|
|1    |written, non-technical description of the project.                                                               |
|     |Technical Description of all the project components. A wireframe for theclient interface, an initial database            schema and an architectural diagram on the initial flow between client, server and database component            |
|     |Registered domain                                                                                                |
|     |Create and initialize a gitHub repo. Include the non-technical and technical description in the README file.             This file will serve as the table to content for the entire project                                              |
|2&3  |Initial deployment of a simple “Hello World” program to serve as confirmation of a working environment	. This         test case will have a dedicated end point                                                                         |
|     |Organize and keep track of events in one location for easy access using a custom process or a third party                library                                                                                                          |
|     |Write a simple test case to demonstrate working knowledge of how test cases can ensure quality and integrity.    |
|     |Each additional week, additional test cases specific to the project will be required.                            |
|     |Update the technical description in the README file to include a contingency backup and recovery plan            |
|4    |Use  text processing tools like XML, JSON, Regexp and templates to produce or handle received text content       |
|     |To demonstrate the use of text processing tools, write logs and error events to a static file.                   |
|5    |Create and deploy client pages that has at least 2-3 pages i.e a login page, a home page etc                     |
|     |Define API endpoints. Functionally is not required at this point but should map from server to client or server         to database as applicable. These endpoints can be comments in the file.                                           |
|6    |Refine the initial database schema if needed. Update the README file accordingly                                 |
|     |Create and deploy database(s) . Ensure tables are created and populate with one default row                      |
|7    |Reinforce security by ensuring confidential data is encrypted correctly                                          |
|     |Add functionality to ensure that use input is always validated and sanitized before storing the data. This could         be as simple as returning an error if a user does not provide valid input values.                                |
|     |At this point, the client should have a sign-in/login interface. The user should have an active session once             logged in. The  server should be able to handle authentication and sessions. Use applicable functionality and             storage (such as redis) to handle authentication and session requests.                                          |
|     |Re-deploy the client and server updates. Ensure the changes reflect on the domain. Measure of completion will be         determined by the ability to login and logout from the domain                                                    |
|8    |Establish a WebSocket connection that updates the client interface                                               |
|     |The client should be able to create a new connection. In the event of failure, the client should handle failure          elegantly.                                                                                                       |
|     |Connections should be stored in a thread-safe data structure                                                     |
|     |Include proper validation to ensure only valid or authenticated users receive the upgraded connection            |
|9&10 |No deliverables. Work on final touches on project                                                                |
|     |If no further work is required, submit final product for approval or demo                                        |

# RESOURCES:
> This project is designed to expand hands on experience beyond what was taught  in prior classes. There will be overlapping reference to material still  available via Canvas. In addition, the internet will be a valuable resource to get knowledge. I chose the following links because they either are guides/tutorials from the software or platforms I may be using or because in my prior experience the authors have been able to explain subject matters in simplified terms and as such I am able to apply the principles to my own projects. 

- https://astaxie.gitbooks.io/build-web-application-with-golang/en/
- https://docs.djangoproject.com/en/3.0/intro/tutorial01/
- https://www.tutorialspoint.com/nodejs/nodejs_express_framework.htm
- https://docs.docker.com/get-started/
- https://devcenter.heroku.com/start
- http://tutorials.jenkov.com/software-architecture/index.html
- https://www.tutorialspoint.com/software_architecture_design/introduction.htm
- https://www.tutorialspoint.com/design_pattern/mvc_pattern.htm
- https://www.rabbitmq.com/tutorials/tutorial-one-go.html
- https://docs.mongodb.com/manual/tutorial/ 
- https://redislabs.com/ebook/part-1-getting-started/chapter-1-getting-to-know-redis/
- http://www.mysqltutorial.org/
- https://docs.google.com/document/d/1WJqojo22sBHqSJED_waGpNHzvDKzyrQrmruJk_v69EM/edit

