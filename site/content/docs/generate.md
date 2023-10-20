---
title: "Generate"
date: 2023-10-19T14:29:00-04:00
draft: false
toc: true
weight: 6
---

This page describes how to use KubeDR to generate Kustomize filter functions based on an API specification. 

## Ask KubeDR to generate a filter function 

To ask KubeDR to generate a filter function use the `generate` sub-command

```
kubedr generate <path API specification>
```

For example, to build a filter function which changes prefix of all docker images you can use [imageprefix.go](imageprefix.go)

```
kubedr generate imageprefix.go
```

If the file is in a GitHub repo, the agent will generate a commit to generate the filter function implementation. If not, the agent will simply add the go files to implement the function. 

