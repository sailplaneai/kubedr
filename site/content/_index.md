---
title: "Kubernetes Doctor (kubedr)"
date: 2023-03-02T15:43:09-08:00
draft: false
toc: true
---

KubeDoctor is a lightweight agent that helps troubleshoot issues in your Kubernetes clusters. You install a CLI tool, connect it to your cluster, and then you can ask it to solve issues with your networking, pods, PVs, etc. Behind the scenes, KubeDoctor uses a technique called ["Scientific Debugging,"](https://arxiv.org/abs/2304.02195) which tests multiple hypotheses about a failure by observing the state of your system using kubectl commands and log queries. We help you debug issues inexpensively, using compute instead of human hours.

The agent is useful off-the-shelf. It comes preloaded with an ever-growing knowledge base of common failures, which you can extend by adding your runbooks to its knowledge base.

Here's a short demo 👇

<div style="position: relative; padding-bottom: 57.324840764331206%; height: 0;"><iframe src="https://www.loom.com/embed/0af9c20dd8494791adbadf4ef5f485b5" frameborder="0" webkitallowfullscreen mozallowfullscreen allowfullscreen style="position: absolute; top: 0; left: 0; width: 100%; height: 100%;"></iframe></div>

<div style="text-align: left">

# More Information

See the [FAQ](/faq) for more information.


</div>