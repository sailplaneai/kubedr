---
title: "Logging"
date: 2023-04-06T09:49:12-07:00
draft: false
toc: true
weight: 7
---

This page describes how to customize KubeDR's logging.

## Filesystem

By default KubeDR creates a new timestampped file containing the logs each time it is invoked.
You can configure the directory where these files are stored by setting `logging.logDir`
in `${HOME}/.kubedr/config.yaml` e.g.

```
logging:
    logDir: /path/to/log/directory
```

## Google Cloud Logging

If you have a GCP account you can configure KubeDR to write logs to 
[Google Cloud Logging](https://cloud.google.com/logging)

### Enabling Cloud Logging

To enable Cloud Logging edit your `${HOME}/.kubedr/config.yaml` 
to set `project` to the GCP project you want to log to. e.g.

```
logging:
    gcp:
        project: <YOUR PROJECT>
```

You must have the role [roles/logging.logWriter](https://cloud.google.com/iam/docs/understanding-roles#logging-roles)
in order to write logs to the project.

KubeDR uses [application default credentials(ADC)](https://cloud.google.com/docs/authentication/application-default-credentials) to access Google Cloud Logging. 
You must set ADC by running the command below

```
gcloud auth application-default login
```

### Accessing Cloud Logs

Once you have enabled Cloud Logging all KubeDR logs will be streamed
to Cloud Logging. Each run is tagged with a unique TraceID which is
stored in the labels of each log entry. This can be used to fetch
the logs for a particular run using a query like the following.

```
labels.source="kubedr"
labels.traceid="9e0010be-49e5-11ee-96c4-6ea4a1656548"
resource.labels.project_id="chat-acme"
```

Be sure to set the correct values of `traceid` and `project_id`.

KubeDR will print out the link for the logs at the start of its run
for easy access.

#### Console Messages

In order to filter the logs to what was shown to the user
you can add `jsonPayload.console="true"` to your query e.g.

```
labels.source="kubedr"
labels.traceid="9e0010be-49e5-11ee-96c4-6ea4a1656548"
resource.labels.project_id="chat-acme"
jsonPayload.console="true"
```
