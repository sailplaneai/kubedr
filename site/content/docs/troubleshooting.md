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
