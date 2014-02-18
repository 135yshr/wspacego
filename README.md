# _wspacego_


_whitespace language interpreter_

* master [![Build Status](https://travis-ci.org/135yshr/wspacego.png?branch=master)](https://travis-ci.org/135yshr/wspacego)
* develop [![Build Status](https://travis-ci.org/135yshr/wspacego.png?branch=develop)](https://travis-ci.org/135yshr/wspacego)

## Project Setup

_How do I, as a developer, start working on the project?_ 

1. _go get github.com/r7kamura/gospel_
1. _go get github.com/135yshr/wspacego_

## Testing

_Using the [gospel](https://github.com/r7kamura/gospel) to the test framework._

### Unit Tests

1. `go test` or `go test ./...`

### Integration Tests

1. _commit to the repository_
2. _Travis ci takes care of running the test automatically after a while._

## Deploying

### _How to setup the deployment environment_

- _Please install the execution environment of the golang._
- _Use the command go get, please get the gospel from github._

### _How to deploy_

1. `go install`

## Troubleshooting & Useful Tools

_error git push_

> remote: Invalid username or password.
> fatal: Authentication failed for 'https://github.com/135yshr/wspacego/'
> 
> - change push repository  
> `git remote set-url origin git@github.com:135yshr/wspacego.git`

## Change history

- _relese version 1.0_
- _I have moved to github from gist_
- _I've split the file for each function_

## License
Copyright &copy; 2014 [135yshr](https://twitter.com/135yshr)
Licensed under the [Apache License, Version 2.0][Apache]
 
[Apache]: http://www.apache.org/licenses/LICENSE-2.0
[MIT]: http://www.opensource.org/licenses/mit-license.php
[GPL]: http://www.gnu.org/licenses/gpl.html