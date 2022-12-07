#!/bin/bash
name_branch=$1

git checkout $name_branch

git pull origin $name_branch

git status