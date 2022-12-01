# Simple Virtual Machine

Simple Virtual Machine writen in Go
<hr>

## Introduction

A Simple build and manage Virtual Machine in KVM 

## How To Use 

```bash 
svm init # Setup Profile
svm img -url <image name> # download Cloud image 
svm create --name <vm_name> --img <image name> # create vm 
```


## Build Binary file

```bash
make builds
```
