## Setup

First, checkout the submodule that contains the theme:

```
(cd site/themes && git submodule update --init --recursive)
```

## Preview the site locally

```
cd site && hugo server -D
```

The `-D` option causes drafts to be published.

## Build the static site

```
hugo -D
```

The static site site will be in `roboweb/public` directory

## Publishing

Publishing happens through [this workflow](./github/workflows/gh-pages.yaml).

# References

[docsy-example repository](https://github.com/google/docsy-example)
[Hugo Modules](https://gohugo.io/hugo-modules/)
[Hugo Quick Start](https://gohugo.io/getting-started/quick-start/)
[Docsy Migrate to Hugo Modules](https://www.docsy.dev/docs/updating/convert-site-to-module/)