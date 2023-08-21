---
title: "AI Safety"
date: 2023-04-06T09:49:12-07:00
draft: false
toc: true
weight: 6
---

KubeDR has two different mechanisms to ensure it uses tools safely and doesn't perform destructive actions.
* Human approval of actions
* IAM and RBAC

By default KubeDR relies on human review of all actions. While this is safe, it creates a burden
on the human operator. Furthermore, by preventing KubeDR from operating autonomously we limit KubeDR
to operating at the speed of the human operator.

The second approach is to rely on IAM and RBAC to restrict the actions KubeDR can take. Relying
on IAM and RBAC makes it easy to reason about the potential impact of KubeDR actions. If KubeDR
is sufficiently restricted you may deem KubeDR sufficiently safe to operate autonomously.

This document explains how to restrict KubeDR using IAM and RBAC for various tools.


## kubectl

To restrict KubeDR's Kubernetes privileges we rely on [RBAC](https://kubernetes.io/docs/reference/access-authn-authz/rbac/).
We create a [service account](manifests/kubedr_rbac.yaml) to be used to execute `kubectl`. This service account is then bound to a 
[ClusterRole](manifests/kubedr_rbac.yaml) to grant KubeDR permissions sufficient for troubleshooting.

Follow the steps below to restrict KubeDR's privileges when running kubectl.

1. Create a copy of [kubedr_rbac.yaml](manifests.kubedr_rbac.yaml)
1. Edit the **kubedr-readonly** ClusterRole as desired
   * This role is designed to give KubeDR readonly access to Kubernetes so that it can troubleshoot problems but not make changes
   * You may want to remove the ability to read secrets if you rely on secrets and are worried about exfiltration of secrets
1. (Optionally) replace the **ClusterRoleBinding** with one or more **RoleBindings** if you only want KubeDR to have access to certain namespaces.
1. Apply the manifests
   ```
   kubectl apply -f kubedr_rbac.yaml
   ```
1. Edit your `${HOME}/.kubedr/config.yaml` to include a kubectl section like the one below

   ```
   ...
   tools:
    kubectl:
        serviceAccount:
            name: kubedr
            namespace: kubedr
        autoRun: false
   ```

   * If you changed the name or namespace of the serviceAccount be sure to set it correctly in your account

1. Run KubeDR as you normally would

   * KubeDR will use the [TokenRequest API](https://kubernetes.io/docs/reference/kubernetes-api/authentication-resources/token-request-v1/) to obtain a token for the service account and then use it to run `kubectl`

1. If you want to run without human confirmation you can set **autoRun** to true in `${HOME}/.kubedr/config.yaml` 

   ```
   ...
   tools:
    kubectl:
        serviceAccount:
            name: kubedr
            namespace: kubedr
        autoRun: true
   ```

## References

* [Blog on AI Safety By Simon Willson](https://simonwillison.net/2023/Apr/14/worst-that-can-happen/)
* [Prompt injection: What's the worst that can happen?](https://simonwillison.net/2023/Apr/14/worst-that-can-happen/)