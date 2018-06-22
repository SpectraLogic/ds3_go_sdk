These instructions are for how to generate the Go SDK documentation using `godoc` and `wget`. You need 2 terminals to generate the documentation.

In terminal 1, navigate to the go SDK folder
`cd /ds3_go_sdk`

Launch the godoc documenation as a website from terminal 1 using the command:
`godoc -http=:8080`


In terminal 2, navigate to the docs folder
`cd /ds3_go_sdk/docs`

Create a folder for the target documentation, and then cd into it.
```
mkdir vX.X.X
cd vX.X.X
```

Use wget to caputre the generated code with the command:
`wget -r -l 20 -e robots=off -np -E -p -k -nH --cut-dirs=4 http://localhost:8080/pkg/github.com/SpectraLogic/ds3_go_sdk`

This may take a minute.

FYI the options used wget perform the following function:
-r downloads files recursively
-l specify recursive maximum depth level (i.e. 20 arbitrarily chosen)
-e robots=off execute without robots protection off
-np does not ascend to the parent directory
-N only retrieves files if newer than local
-E adds extension .html to html files if they do not have extension
-p downloads all files for each page, including css, js, and images
-k converts links to relative paths
-nH does not generate host prefix directory (i.e. no localhost:8080 folder)
--cut-dirs=4 does not generate the first few layer of folders

