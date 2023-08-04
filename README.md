# kubedr
kubedr is an AI to troubleshoot Kubernetes &amp; Cloud issues

## Setup

First, checkout the submodule that contains the theme:

```
(cd roboweb/themes && git submodule update --init --recursive)
```

## Preview the site locally

```
cd roboweb && hugo server -D
```

The `-D` option causes drafts to be published.

## Build the static site

```
cd roboweb && hugo -D
```

The static site site will be in `roboweb/public` directory

## Publishing

Publishing happens through [this workflow](./github/workflows/gh-pages.yaml).

# References

[Hugo Quick Start](https://gohugo.io/getting-started/quick-start/)