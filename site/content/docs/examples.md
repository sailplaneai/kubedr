---
title: "Examples"
date: 2023-04-06T09:49:12-07:00
draft: false
toc: true
weight: 4
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
kubedr diagnose problem "Why can't network traffic reach the service esp-echo in namespace kubedr-examples"
```

Every troubleshooting session is different but typically kubedr will run through some combination of the following steps

* Issuing the command `kubectl -n kubedr-examples get svc esp-echo -o yaml` to check the service
* Checking for matching pods by doing `kubectl -n kubedr-examples get pods -l app=echo -o jsonpath="{.items[*].status.phase}"`
* Issuing the command `kubectl -n kubedr-examples describe svc esp-echo` to check for any pods
* Issuing the command `kubectl -n kubedr-examples get deployment -l app=echo -o yaml` to check for deployments matching the selector

A typical conclusion might look like the following 

```
The problem is that there is no Deployment available for the 'esp-echo' service in the 'kubedr-examples' namespace. Consequently, no pods are running that could handle the network traffic directed to the 'esp-echo' service. You would need to create a Deployment that correctly selects the 'esp-echo' service based on labels and ensures that the necessary pods are running.
```

In this case KubeDr correctly determines that the problem is that there are no deployments available because the service's selector and deployment's labels don't match.

## Workload Identity Example

Create the service account.

```
kubectl apply -f examples/workload_identity.yaml
```

Now we can use KubeDr to diagnose problems with workload identity.

```
kubedr diagnose problem "why can't service account esp-echo in namespace kubedr-examples access GCS buckets"
```

kubedr will typically do the following

* Fetch the service account to check for the workload identity annotation `iam.gke.io/gcp-service-account`
* Get the desired GCP service account from the `iam.gke.io/gcp-service-account`
* Try to fetch the IAM service account

Based on this, kubedr will conclude that the service account is missing and possibly suggest a command to fix it by creating the service account.

To further explore using Kubedr you can try creating the service account and fixing this problem. This will only work if you have permission to 
create service accounts.

```
gcloud iam service-accounts create kubedr-example --description="Service Account for KubeDr examples" --display-name="kubedr-example"
```

Update the service account to use the newly created service account

```
kubect -n kubedr-examples edit serviceaccount esp-echo
```

Set the value of the annotation `iam.gke.io/gcp-service-account` to  `kubedr-example@${PROJECT}.iam.gserviceaccount.com`. Be sure to
substitute in the name of your project for $PROJECT.

Now rerun kubedr

```
kubedr diagnose problem "why can't service account esp-echo in namespace kubedr-examples access GCS buckets"
```

kubedr should correctly determine that the GCP service account exists but is missing the correct rolebinding to
allow the Kubernetes service account to assume its identity.


## Cleanup

After you are done trying out KubeDr its a good idea to delete the namespace to cleanup any examples you created.

```
kubectl delete namespace 
```
