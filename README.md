oxford-face-client: A client App for [oxford-face](http://github.com/kkdai/oxford-face) Golang package
======================
[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/kkdai/oxford-face-client/master/LICENSE) [![Build Status](https://travis-ci.org/kkdai/oxford-face-client.svg)](https://travis-ci.org/kkdai/oxford-face-client)

![](https://camo.githubusercontent.com/71f5a0a445474d8412fad833c73be9b57f66c2c8/68747470733a2f2f7777772e70726f6a6563746f78666f72642e61692f696d616765732f6272696768742f666163652f466163654150492d4d61696e2e706e67)

It is a simple client to help you to maniplate the [Project Oxford](https://www.projectoxford.ai/) [Face API](https://www.projectoxford.ai/face). No need to remember complicate Face ID anymore, just input URL or image path to add our local list and compare with index.

Feature
--------------


Install
--------------

```
go get -u -x github.com/kkdai/oxford-face-client
```

Usage
---------------------

```
oxford-face-client   [OPTIONS]
```



Options
---------------

- `-k` Project Oxford API key. It is must item.
- `-v` Display verbose


Interactive Command
---------------

It support command line interactive command as follow:

- `a`: Display next page aticles.
- `s`: Open content folder in finder.
- `l`: Display all store face id in our temp list.
- `v`: Toggle debug information.
- `q`: Exist current application.

Examples
---------------

TBC
     



Contribute
---------------

Please open up an issue on GitHub before you put a lot efforts on pull request.
The code submitting to PR must be filtered with `gofmt`


Project52
---------------

It is one of my [project 52](https://github.com/kkdai/project52).


License
---------------

This package is licensed under MIT license. See LICENSE for details.
