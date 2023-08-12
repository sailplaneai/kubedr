---
title: "Overview"
date: 2023-04-06T09:49:12-07:00
draft: false
toc: true
weight: 1
---

KubeDr is a CLI that uses AI to troubleshoot and fix Kubernetes and Cloud problems. 

KubeDr uses a technique called ["Scientific Debugging,"](https://arxiv.org/abs/2304.02195) to
generate the following
  
   * A hypothesis about why a problem is occuring
   * A prediction that can be used to test whether the hypothesis is true
   * An experiment to test the prediction

After generating a hypothesis, prediction, and experiment, KubeDr will ask for permission to run the experiment. Based on the experiment's output
KubeDr will evaluate the prediction and hypothesis. KubeDr will then use that information
to either generate new hypotheses, predictions, and experiments or draw a conclusion about the 
source of the problem.

KubeDr's prior knowledge about Cloud and Kubernetes can be augmented with runbooks. A
[runbook](https://github.com/jlewi/kubedr/tree/main/runbooks) is just a YAML file
containing tuples of (hypothesis, prediction, experiment). KubeDr can automatically
generate these YAML files from your existing runbooks.

## Why Use KubeDr?

Infrastructure, platform, and app teams confront the growing complexity of Cloud & Kubernetes on a daily basis.
This complexity makes it very difficult for a single engineer to fully understand a company's stack and not
just the layer they work on daily. This problem is particular acute when trying to diagnose a problem
whose root cause may be in a different part of the stack.

KubeDr solves this problem by using AI to synthesis information and generate troubleshooting
steps that can operate across the stack. KubeDr automates troubleshooting reducing toil 
and time to problem resolution. 

## Who is KubeDr for?

**Infrastructure engineers** who deal with Kubernetes and Cloud on a daily basis. KubeDr can
help infrastructure engineers quickly debug and fix these issues.

**Platform teams** who want to provide their customers with tools to diagnose and fix problems with the platform.

**App teams** who need help troubleshooting with Cloud/Kubernetes or their internal developer platform.

## Where should I go next?

Follow the [Getting Started](/docs/getting_started/) guide to download and use KubeDr.