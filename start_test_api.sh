#!/bin/bash

echo "-------- Test case execution started --------"
tag_var=~@qa
if [ "$#" -eq  "0" ]
   then
     echo "No tags supplied. Running tests for all the scenarios"  
 else
     tag_var=$1
     echo "Tag Name supplied. Using $tag_var as tag"
 fi
echo "---------------------------------------------"

export GO_ENV="test"
GO111MODULE=auto
cd functional_test && godog --format=progress --tags=$tag_var