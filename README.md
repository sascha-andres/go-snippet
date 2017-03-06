# go-snippet

[![Go Report Card](https://goreportcard.com/badge/github.com/sascha-andres/go-snippet)](https://goreportcard.com/report/github.com/sascha-andres/go-snippet) [![codebeat badge](https://codebeat.co/badges/7fc9e22e-807e-4742-b129-d6062397beef)](https://codebeat.co/projects/github-com-sascha-andres-go-snippet) 
## What is go-snippet?

You can define snippets which are inserted at files at places you want them to be.

## Why go-snippet?

I have a bunch of bash scripts taht often use common functions. Instead of including all or having copies of a global include file in different repositories I wanted to be able to choose what to include and have an easier way to keep it in sync.

## Give me an example

Suppose you write a lot of bash scripts and want headers in them. Instead of repeating the `echo`s you write a function. As you do this over and over again, this is a good thing for a snippet.

So let's assume we have a snippet like this asved as header.sh in your snippetdir:

    function header () {
      echo
      echo "*** $1 ***"
      echo
    }

You would then write a bash script like this:

    #! /bin/bash

    #### header.sh ###

    header "Hello world!"

A call to `go-snippet bash --file myfile.sh` would generate the following file:

    #! /bin/bash

    #### header.sh ###
    function header () {
      echo
      echo "*** $1 ***"
      echo
    }
    #### header.sh ###

    header "Hello world!"

Here we go, a working script! Suppose you do not like to have stars but minus signs, so you change the snippet to:

    function header () {
      echo
      echo "--- $1 ---"
      echo
    }

Run the tooling again and you get:

    #! /bin/bash

    #### header.sh ###
    function header () {
      echo
      echo "--- $1 ---"
      echo
    }
    #### header.sh ###

    header "Hello world!"