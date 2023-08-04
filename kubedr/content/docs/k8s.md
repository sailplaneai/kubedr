---
title: "Kubernetes"
date: 2023-04-06T09:49:12-07:00
draft: false
toc: true
---

This article shows you how to start JupyterLab with the Roboweb extension on
Kubernetes using our public docker image

* [us-west1-docker.pkg.dev/roboweb-images/roboweb-images/kubeflow/jupyterlab:latest](https://console.cloud.google.com/artifacts/docker/roboweb-images/us-west1/roboweb-images/kubeflow%2Fjupyterlab)

  * Note the Docker image is publicly available but to browse the registry you must sign in with a Google account.

This docker image contains JupyterLab with the Roboweb extension already installed.
It is based on [Kubeflow's JupyterLab image](https://www.kubeflow.org/docs/components/notebooks/container-images/#custom-images)
and thus should be deployable inside Kubeflow.

If you don't have Kubeflow, you can use the minimal Kubernetes manifest in
[jlewi/roboweb-docs/k8s/manifests](https://github.com/jlewi/roboweb-docs/tree/main/k8s/manifests).

Clone the [jlewi/roboweb-docs](https://github.com/jlewi/roboweb-docs.git) repository

```
git clone https://github.com/jlewi/roboweb-docs.git
```

**Note** As is this manifest isn't intended for doing meaningful work inside JupyterLab. In particular, it doesn't
use persistent disk to store your notebooks so any work will be lost when the
pod restarts. For a more robust notebook deployment we recommend

* Using [Kubeflow Notebooks](https://www.kubeflow.org/docs/components/notebooks/)
* Customizing the manifests to meet your needs

## Deploying JupyterLab on Kubernetes

To deploy it run

```
cd roboweb-docs
kubectl create namespace roboweb
kustomize build k8s/manifests | kubectl apply -f
```

This will deploy the notebook in the namespace `roboweb`.

To use a different namespace edit the `k8s/manifests/kustomization.yaml` file.

## Accessing It

To access JupyterLab port-forward to the pod

```
kubectl -n roboweb port-forward notebook-0 8888:8888
```

Then open a browser to `http://localhost:8888` to access JupyterLab.

## Using Roboweb in JupyterLab

Refer to the directions in the [quickstart](/docs/quickstart/) for how to use Roboweb in JupyterLab.

You can skip the steps to install Roboweb since the docker image already has it installed.
