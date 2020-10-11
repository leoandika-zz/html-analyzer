# Welcome to HTML Analyzer (Backend) Project!

A Backend-Only HTML Analyzer Project to analyze your web page, fully written using Go Language. This application is able to analyze the following attributes : 

- HTML Version of the page
- Page title
- Heading Level counts (From h1 to h6!)
- Number of Internal and External Links
- Number of Inaccessible Links (yep, you read that right!)
- Whether your page contains login form or not

## How to run
A provided makefile make everything easy! Simply run the following commands to do things : 
 1. Clone the project and change current directory to the project folder
 2. To build the project, just run `make build`
 3. When you're done buildidng, run `make run` to run the application. Remember, the default port used in this application is 8087. Watch out for any other applications running on the same port.
 4. To run the pre-written tests, run `make test`
 5. Last but not least, to remove the binary, run `make clean`
 
 ## How to contribute
 As they say, no project is perfect. This project also need many improvements. 
 Therefore, if you found anything you can improve, please don't hesitate to create an issue and make a pull request. 
 I will review it as soon as possible! :)