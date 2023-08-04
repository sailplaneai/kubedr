---
title: "Exploratory Programming With AI"
date: 2023-03-02T15:43:09-08:00
draft: false
toc: true
---

Roboweb is an AI assistant for [exploratory programming](https://en.wikipedia.org/wiki/Exploratory_programming). It embeds OpenAI's [ChatGPT](https://chat.gpt.ai/)
in JupyterLab to create an ideal environment for exploratory programming.

Here's a short demo 👇

<div style="position: relative; padding-bottom: 57.324840764331206%; height: 0;"><iframe src="https://www.loom.com/embed/5ab26344119440a5b70ce4094827d9e3" frameborder="0" webkitallowfullscreen mozallowfullscreen allowfullscreen style="position: absolute; top: 0; left: 0; width: 100%; height: 100%;"></iframe></div>

<div style="text-align: left">

# Installation

### Docker (Recommended)

Make sure you have [Docker Desktop](https://www.docker.com/products/docker-desktop/) installed.

```bash
docker run -it --rm -p 8888:8888 \
 -v "$PWD":/home/jovyan/work hamelsmu/roboweb
```

This will serve Jupyter Lab on port `8888`, with a link that will be printed to your terminal.

---

### pip

It's recommended to setup a virtual environment first.  In this example, we use conda, but feel free to use whatever you want.

```bash
## Create a virtual environment
conda create -n robo -c conda-forge python=3.9 jupyterlab
# activate the virtual environment
conda activate robo
# install the extensions
python -m pip install roboweb-extension roboweb-server
# enable the server extension
python -m jupyter serverextension enable --py roboweb_server
# run the server
jupyter lab
```

---

### Kubernetes

You can deploy this app on Kubernetes using [these instructions](docs/k8s).


## Usage


### 1. Sign In / Create An Account

In the right hand navigation window click the roboweb icon (which is a small robot). This will open up the roboweb extension in the right half of the screen.

{{< figure src="/docs/toolbar.png" alt="Roboweb toolbar">}}

Sign in with Google or create an account.  Accounts help you keep track of your chats and also allow you to retrieve them later.

{{< figure src="/docs/signin.png" alt="SignIn Page">}}

### 2. Add Your OpenAPI Key

Once you are logged in, click the settings button in the navigation bar inside Roboweb's JupyterLab extension.

{{< figure src="/docs/settings_bottom.png" alt="SignIn Page">}}

On the settings page, enter your OpenAI API key and then click `Save`.  Your API key is stored in your browser’s local storage and is never transmitted to Roboweb’s servers.

{{< figure src="/docs/apikey.png" alt="Settings">}}


### 3. Using the extension to fix code

If an error is detected in a cell, roboweb will automatically ask you if you would like to fix the error.

{{< figure src="/docs/error.png" caption="Roboweb makes it easy to fix errors." alt="Codebox">}}

If you click the `Fix detected errors` button, roboweb will tell you how to fix it:

{{< figure src="/docs/fix.png" caption="Roboweb makes it easy to fix errors." alt="Codebox">}}

# More Information

See the [FAQ](/faq) for more information.


</div>