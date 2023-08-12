title: "Getting Started"
date: 2023-04-06T09:49:12-07:00
draft: false
toc: true
---


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

