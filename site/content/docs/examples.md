---
title: "Examples"
date: 2023-04-06T09:49:12-07:00
draft: false
toc: true
---

This page provides examples of troubleshooting various problems using KubeDr. 

You will need access to a Kubernetes cluster in order to deploy the example
resources. If you don't have a Kubernetes cluster refer to the 
[Kubernetes Documentation](https://kubernetes.io/docs/setup/) to learn about
the many ways to run Kubernetes.

## Prerequisites

1. Clone the [kubedr](https://github.com/jlewi/kubedr) repository to obtain
   the Kubernetes manifests

   ```
   git clone https://github.com/jlewi/kubedr git_kubedr
   ```

1. If you haven't already indexed the [predefined runbooks](https://github.com/jlewi/kubedr/tree/main/runbooks) 
   index them now as the examples depend on them

   ```
   kubedr index dir git_kubedr/runbooks
   ```

   * Indexing a runbook is an idempotent operation so running the above operation is a null op if you've already indexed the runbooks

1. Create the namespace `kubedr-examples` as that's used by the examples.

   ```
   kubectl create namespace kubedr-examples
   ```

## Networking Example

Networking is arguably one of the most easily misconfigured things since it often involves multiple layers (e.g. Cloud loadbalancers, service mesh, application configuration).

In this example we deploy a simple echo service and then troubleshoot why traffic can't reach it.

Deploy the application.

```
kubectl apply -f ${REPO}/examples/service_labels/manifests.yaml
```

Verify that the pods are deployed.

```
kubectl get pods
No resources found in autobuilder namespace.
kubectl -n kubedr-examples  get pods
NAME                        READY   STATUS    RESTARTS   AGE
esp-echo-764cfbb49c-k68dq   1/1     Running   0          39s
```

Try to port-forward to the service and verify that traffic isn't getting through

Setup port-forwarding

```
kubectl -n kubedr-examples  port-forward service/esp-echo 8080:80
```

Now try to send a curl request

```
curl localhost:8080/hello
curl: (7) Failed to connect to localhost port 8080 after 3 ms: Connection refused
```

Since the request is failing lets use KubeDr to diagnose the problem

```
./kubedr diagnose problem "Why can't network traffic reach the service esp-echo in namespace kubedr-examples"
```

## Cleanup

After you are done trying out KubeDr its a good idea to delete the namespace to cleanup any examples you created.

```
kubectl delete namespace 
```
