---
title: "Troubleshooting"
date: 2023-04-06T09:49:12-07:00
draft: false
toc: true
weight: 8
---

## Timeout waiting for weaviate client to connect to DB

When KubeDR runs it will block waiting to establish a connection to the Weaviate database.
While it is waiting it will show messages like the following.

```
2023/09/05 16:52:04 Weaviate not yet up. Waiting for another second.
2023/09/05 16:52:05 Weaviate not yet up. Waiting for another second.
```

If it is unable to connect it will eventually timeout with an error like the following

```
Error starting weaviate weaviate did not start up in 1m0s. Either the Weaviate URL "http://127.0.0.1:53912/v1" is wrong or Weaviate did not start up in the interval given in 'startupTimeout'
```

There are two possible causes. There is already a weaviate process running and this is
preventing KubeDR from starting a new weaviate process. 

To check for such processes run `ps -ef | grep weaviate` e.g

```
>ps -ef | grep weaviate                                       
  501 97365   716   0  4:51PM ttys001    0:01.28 weaviate --scheme=http --port 52889
```

Kill the running process by running

```
pkill -x weaviate
```

The other possible cause is that you set `dbConfig.url` to point to an existing instance of weaviate but aren't
actually running weaviate at that URL. To check your configuration you can run

```
kubedr config get
```

If `dbConfig.url` is set to a non empty string there are two possible ways to fix the connection issue

1. Unset `dbConfig.url`

   ```
   kubedr config set dbConfig.url=""
   ```

   * unsetting it will cause kubedr to launch weaviate

1. Start weaviate at the URL you set


## Weaviate Error Unable to Resolve node name to host

When running KubeDR you get an error like the following

```
explorer: get class: vector search: object vector search at index goals: remote shard 2a4JrhIUoDS4: resolve node name "some.host.name" to host
Failed to query runbook entries
```

This is an issue with Weaviate not obtaining a stable hostname (refer to the [docs](https://weaviate.io/developers/weaviate/more-resources/migration-guide#why-is-a-data-migration-necessary)).

You can fix this by configuring KubeDR to explicitly set the name through the environment variable `CLUSTER_HOSTNAME`.
Edit `${HOME}/.kubedr/config.yaml` and set the environment variable in the `DBCONFIG` section e.g.

```
apiVersion: ""
kind: ""
dbConfig:
    path: ""
    url: ""
    env:
        - name: CLUSTER_HOSTNAME
          value: some.host.name
...
```

Set the value to the value in the error message.

**Important** Since KubeDR is starting Weaviate in a subprocess simply setting `CLUSTER_HOSTNAME` in your environment won't work.
