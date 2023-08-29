---
title: "Assistant"
date: 2023-04-06T09:49:12-07:00
draft: false
toc: true
weight: 6
---

This page describes how to use KubeDR to assist you with various tasks.

## Ask KubeDR for assistance

To ask KubeDR for assistance with a task you use the assist sub-command

```
kubedr assist <task or question>
```

For example, if you are using [flux](https://fluxcd.io/) for GitOps we might use
kubedr to determine the latest release has been successfully applied.

```
kubectl assist "Are flux applied resources up to date in namespace sampleapp?"
```

Or we might ask for help with kubectl

```
kubectl assist "What GKE cluster is the Kubernetes config using?"
```

## GoalBooks

KubeDR can be tought how to perform various actions using GoalBooks. 
GoalBooks are very similar to [runbooks](/docs/runbooks). A GoalBook is a
YAML file providing tuples of (thought, goal, action) which can 
be used to accomplish various actions. 

```
apiVersion: kubedr.ai/v1alpha1
kind: GoalBook
metadata:
  name: kubectl
  namespace: kubectl
entries:
- thought: In order to get information about the kubernetes context we need to know the current-context
  goal: Get the current-context by running `kubectl config current-context` this will return the name of the current context
  action: Run `kubectl config current-context`
- thought: To get the current kubernetes cluster we can run `kubectl config get-contexts` and parse the line with the current context. The current selected context will be marked with a `*`
  goal: Get the current kubernetes cluster by running `kubectl config view get-contexts` and extract the cluster column.
  action: Run `kubectl config get-contexts`
```

More examples of GoalBooks can be found in the [kubedr repository](https://github.com/jlewi/kubedr/tree/main/runbooks).

The YAML file is intended to be an intermediary format into which your runbooks and documentation can be converted so
as to be understood by Kube Doctor.

In the future Kube Doctor will help you automatically generate these YAML files from your existing runbooks and documentation.
For now you have to manually author them.
