---
title: "Configuration"
date: 2023-04-06T09:49:12-07:00
draft: false
toc: true
weight: 7
---

## Weaviate

By default KubeDR will launch a subprocess to run weaviate. However, if you want to run
Weaviate manually you can configure KubeDR by setting the url of your weaviate instance

```
kubedr config set dbConfig.url="http://WEAVIATE_HOST:PORT"
```