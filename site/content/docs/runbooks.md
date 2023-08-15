---
title: "Runbooks"
date: 2023-04-06T09:49:12-07:00
draft: false
toc: true
weight: 5
---

This page describes how to create and index runbooks. Runbooks allow you to customize the agent
for your infrastructure.

## What is a Runbook

A runbook is a YAML file providing tuples of (hypothesis, prediction, experiment) which can 
be used to troubleshoot various problems. An example is shown below

```
apiVersion: kubedr.ai/v1alpha1
kind: RunBook
metadata:
  name: iam
  namespace: gke
entries:
  - hypothesis: The service account's IAM roles are missing the desired permissions
    prediction: The Google Cloud IAM Role ${ROLE} doesn't include permission ${PERMISSION} 
    experiment: Run `gcloud iam roles describe  roles/${ROLE}` and check whether the desired permission is missing
  - hypothesis: The kubernetes service account doesn't have the requested permissions because workload identity is not enabled on the cluster
    prediction: The workload identity pool isn't set on the cluster
    experiment: Run `gcloud container clusters describe ${CLUSTER_NAME} --location=${REGION} --format="value(workloadIdentityConfig.workloadPool)`
    references:
      - https://cloud.google.com/kubernetes-engine/docs/troubleshooting/troubleshooting-security#workload-identity
```

More examples of runbooks can be found in the [kubedr repository](https://github.com/jlewi/kubedr/tree/main/runbooks).

The YAML file is intended to be an intermediary format into which your runbooks and documentation can be converted so
as to be understood by Kube Doctore.

In the future Kube Doctor will help you automatically generate these YAML files from your existing runbooks and documentation.
For now you have to manually author them.

## Index the runbooks

1. To index a directory containing one or more run the `index dir` command.

    ```
    kubedr index dir ${DIRECTORY}
    ```

   * This will recourse over the directory and process any YAML files containing Runbooks
   * Indexing a runbook is an idempotent operation so running the above operation is a null op if you've already indexed the runbooks


1. Verify the runbooks were indexed

   ```
   kubedr index list
   ```

   * This command should output a table listing all the tuples (hypothesis, prediction, experiment) kubedr has indexed

   ```
   +--------------------------------------+--------------+------------+--------------------------------+---------------------------------------------------------------------+-------------------------------------------------------+
   |                  ID                  |     NAME     | NAMESPACE  |           HYPOTHESIS           |                             PREDICTION                              |                      EXPERIMENT                       |
   +--------------------------------------+--------------+------------+--------------------------------+---------------------------------------------------------------------+-------------------------------------------------------+
   | 1569347c-855c-38a6-34bd-d5964b61a4c4 | iam-0        | gke        | The service account's IAM      | The Google Cloud IAM Role                                           | Run `gcloud iam roles describe                        |
   |                                      |              |            | roles are missing the desired  | ${ROLE} doesn't include                                             |  roles/${ROLE}` and check                             |
   |                                      |              |            | permissions                    | permission ${PERMISSION}                                            | whether the desired permission                        |
   |                                      |              |            |                                |                                                                     | is missing                                            |
   | 1675ddd2-5c30-a8fe-9875-d656c720912b | iam-1        | gke        | The kubernetes service account | The workload identity pool                                          | Run `gcloud container clusters describe               |
   |                                      |              |            | doesn't have the requested     | isn't set on the cluster                                            | ${CLUSTER_NAME} --location=${REGION}                  |
   |                                      |              |            | permissions because workload   |                                                                     | --format="value(workloadIdentityConfig.workloadPool)` |
   |                                      |              |            | identity is not enabled on the |                                                                     |                                                       |

   ...
   ```