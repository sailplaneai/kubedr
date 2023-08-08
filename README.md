# KubeDoctor

KubeDoctor is a lightweight agent that helps troubleshoot issues in your Kubernetes clusters. You install a CLI tool, connect it to your cluster, and then you can ask it to solve issues with your networking, pods, PVs, etc. Behind the scenes, KubeDoctor uses a technique called ["Scientific Debugging,"](https://arxiv.org/abs/2304.02195) which tests multiple hypotheses about a failure by observing the state of your system using kubectl commands and log queries. We help you debug issues inexpensively, using compute instead of human hours.

The agent is useful off-the-shelf. It comes preloaded with an ever-growing knowledge base of common failures, which you can extend by adding your runbooks to its knowledge base.

For a short demo and usage instructions visite the website

[www.kubedr.ai](https://www.kubedr.ai)
