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

- `a {FACE_ADDR}`: Display next page aticles. The `FACE_ADDR ` could be URL or path. 
	- ex: `a https://oxfordportal.blob.core.windows.net/face/demov1/verification1-1.jpg`
	- ex: `a test_data/verification1-1.jpg`
- `s {FACE_INDEX}`: Check the Similarity of target index to other. The `FACE_INDEX ` is the target face index of our list.
	- ex: `s 1`: If list has two items, will compare the similarity of index 1 and index 0.
	- ex: `s 0`: If list has three items, it will compare with 0->1, 0->2 similarity.
- `c {FACE1_INDEX} {FACE2_INDEX}`: Will Verify if `FACE1_INDEX` and `FACE2_INDEX` is the same face or not.
	- ex: `c 0 1` Check the first and second face
	- ex: `c 1 3` Check the second the 4th face.
- `l`: Display all store face id in our temp list.
- `v`: Toggle debug information.
- `q`: Exist current application.

Examples
---------------

```
Command:( A:Add Face S:Check Similarity C:Check input two face  V:Verbose G:Read Q:exit)

// Add first image from URL
:>a https://oxfordportal.blob.core.windows.net/face/demov1/verification1-1.jpg
New Face: id= 157675c5-ca4c-4882-ab05-72b14b0d0194


// Add second image from URL
:>a https://oxfordportal.blob.core.windows.net/face/demov1/verification1-2.jpg
New Face: id= 5a236d4a-e07c-4f9b-aaca-c13e4b5c8fc0


// Check first image and second image if it is identical face.
:>c 0 1
Is identical? true  confidence: 0.66216


//Check similarity of index "0" to other
:>s 0
Most similar in index: 1  faceid:  5a236d4a-e07c-4f9b-aaca-c13e4b5c8fc0  confidence:  0.662164748


//Add face using file path
:>a test_data/verification1-1.jpg
New Face: id= f7eb6931-835d-4d68-8272-2e8c0064c870


// List all faces
:>l
Index 	| Face ID 		| From
===============================================================
0 	 157675c5-ca4c-4882-ab05-72b14b0d0194 	 https://oxfordportal.blob.core.windows.net/face/demov1/verification1-1.jpg
1 	 5a236d4a-e07c-4f9b-aaca-c13e4b5c8fc0 	 https://oxfordportal.blob.core.windows.net/face/demov1/verification1-2.jpg
2 	 f7eb6931-835d-4d68-8272-2e8c0064c870 	 test_data/verification1-1.jpg


```
     

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
