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

## Usage

1. Download the latest release from the [releases page](https://github.com/jlewi/kubedr/releases)

1. If you don't already have an OpenAI API key visit [https://openai.com/](https://openai.com/) to 
   obtain one
   
1. Configure kubedr to use your OpenAI API Key

   ```
   kubedr config set apiKeyFile=/path/to/your/api.key.file
   ```

   * `api.key.file` should be a text file containing your OpenAI APIKey 

1. Check whether network traffic can reach an endpoint

   ```
   kubedr diagnose traffic -n ${NAMESPACE} {NAME}
   ```

1. kubedr will respond with a hypothesis, prediction, and experiment. kubedr will 
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

1. kubedr will run the experiment and then analyze the results and do one of the following

   * If kubedr has reached a conclusion it will provide the conclusion and exit
   * Generate a new hypothesis, prediction, and experiment to futher diagnose the issue

# More Information

See the [FAQ](/faq) for more information.

</div>