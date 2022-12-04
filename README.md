# kubemerge

Script to merge mutltiple Kubeconfig file into one


## Usage

In order to merge kubeconfig files

You'll need to run command as :

```shell
    kubemerge <full-path-of-config1> [ <full-path-of-config2> ......]
```

## Install

```shell
curl -O github.com/datagenx/kubemerge/kubemerge \
    && sudo mv kubemerge /usr/local/bin \
    && sudo chmod +x /usr/local/bin/kubemerge
```
