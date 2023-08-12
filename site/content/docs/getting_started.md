---
title: "Getting Started"
date: 2023-04-06T09:49:12-07:00
draft: false
toc: true
---

Follow the instructions below to get started with KubeDr

## Download KubeDr

1. Download the latest release from the [releases page](https://github.com/jlewi/kubedr/releases)


## Download Weaviate

1. Download the latest release of [weaviate](https://weaviate.io/) from the [releases page](https://github.com/weaviate/weaviate/releases)

1. Ensure the weaviate binary is on your path and accessible as `weavate` 

   * KubeDr runs weaviate locally in order to index and query runbooks


## Configure KubeDr

1. If you don't already have an OpenAI API key visit [https://openai.com/](https://openai.com/) to 
   obtain one
   
1. Configure kubedr to use your OpenAI API Key

   ```
   kubedr config set apiKeyFile=/path/to/your/api.key.file
   ```

   * `api.key.file` should be a text file containing your OpenAI APIKey 

1. Create the index for runbooks
   
   kubedr index create

1. Clone KubeDr's public runbooks

   ```
   git clone https://github.com/jlewi/kubedr /tmp/kubedr
   ```

1. Index the runbooks

   ```
   kubedr index dir /tmp/kubedr/runbooks
   ```

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
## Diagnose a problem

You can diagnose a problem using the **diagnose** command 

```
kubedr diagnose problem ${PROBLEM}
```

${PROBLEM} should be a natural language description of the problem you want to diagnose. Here are some examples

* "service account esp-echo in namespace sampleapp can't access GCS buckets"
* "network traffic can't reach the service esp-echo in namespace sampleapp"

It is important to write clear problem descriptions. Here are some tips

* For a problem dealing with a Kubernetes resource include the following information

  * Type of resource
  * Namespace of resource
  * Name of resource

* For a problem dealing with a Google Cloud resource include the following information

  * Type of resource
  * Project of resource
  * Zone or region of resource (if applicable)
  * Name of resource


kubedr will respond with a hypothesis, prediction, and experiment. kubedr will 
ask for confirmation to run the experiment. Type `yes` to run the experiment

```
kubedr will 
Use the scientific method to determine if network traffic can access the service named hydros in namespace hydros. Say 'Everything looks good' if no problems can be found.

kubedr is thinking...
Hypothesis: The Kubernetes Service "hydros" might not be properly set up in the "hydros" namespace, causing network traffic to be unable to access the service.

Prediction: If the Service "hydros" is correctly configured and functioning correctly, the command 'kubectl -n hydros get service hydros' should return information about the service.

Experiment: Run the following command - 'kubectl -n hydros get service hydros' to check if the service exists and is properly set up in the namespace.

kubedr would like to execute the following command:
kubectl -n hydros get service hydros

Do you want to run the command (yes/no)[no]: yes
```

kubedr will run the experiment and then analyze the results and do one of the following

   * If kubedr has reached a conclusion it will provide the conclusion and exit
   * Generate a new hypothesis, prediction, and experiment to futher diagnose the issue

