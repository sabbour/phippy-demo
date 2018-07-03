# the phippy demo

it's kind of a long story, but this repo will eventually contain something interesting. 

## prerequisites

everything has prerequisites. for this one, create yourself a dockerhub or azure container registry instance. then, log into it from the command line. **then**, make sure you have a Kubernetes cluster set up and that your client is set up to talk to it.

your AKS cluster needs to have been configured with [HTTP Application Routing](https://docs.microsoft.com/en-us/azure/aks/http-application-routing).

if you want to select which registry to push to, use the command below:

```sh
draft config set registry <your docker hub or Azure Container Registry>
```

## good? 

to get started, do this:

```bash
git clone https://github.com/bradygmsft/phippy-demo.git
cd phippy-demo
cd parrot
draft up
cd ..
cd captainkube
draft up
cd ..
cd nodeapp
draft up
cd ..
```

## all done? 

test your deployment by running this:

```bash
kubectl get svc --namespace phippy
```

you should see something like this

```sh
NAME                  TYPE           CLUSTER-IP     EXTERNAL-IP     PORT(S)     AGE
parrot-parrot         ClusterIP      10.0.89.91     <none>          80/TCP      1m
kubernetes            ClusterIP      10.0.0.1       <none>          443/TCP     2d
nodeapp-nodeapp       ClusterIP      10.0.236.198   <none>          80/TCP      2m
```

## allow inbound access

the cluster has a default inbound DNS address. to get it, execute the `az` CLI command:

```bash
az aks show -n <your aks cluster> -g <your resource group>
```

find the `httpApplicationRouting` segment of the JSON payload and copy the `HTTPApplicationRoutingZoneName` value:

```json
"httpApplicationRouting": {
  "config": {
    "HTTPApplicationRoutingZoneName": "05500238a78f.westeurope.aksapp.io"
  },
  "enabled": true
}
```

paste in the DNS to `parrot\charts\values.yaml` so that it looks like this, replacing the example value with your own:

```yaml
ingress:
  enabled: true  # make sure it is enabled
  basedomain: 05500238a78f.westeurope.aksapp.io # replace this
```

paste in the DNS to `nodeapp\charts\values.yaml` so that it looks like this, replacing the example value with your own:

```yaml
ingress:
  enabled: true # make sure it is enabled
  basedomain: 05500238a78f.westeurope.aksapp.io # replace this
```

once you do this, the dotnet app should respond on `parrot.<your clusters specific dns zone>`, and the Node.js app on `nodeapp.<your clusters specific dns zone>`. for example: `parrot.05500238a78f.westeurope.aksapp.io`

to deploy the changes, run these commands:

```bash
cd parrot
draft up
cd ..
cd nodeapp
draft up
```

test your deployment by running this:

```bash
kubectl get ingress --namespace phippy
```

you should see something like this

```sh
NAME               HOSTS                                      ADDRESS          PORTS     AGE
parrot-parrot      parrot.05500238a78f.westeurope.aksapp.io   52.136.252.253   80        1m
nodeapp-nodeapp    nodeapp.05500238a78f.westeurope.aksapp.io  52.136.252.253   80        2m
```

give your parrot a visit by heading over to its hostname, for example [http://parrot.05500238a78f.westeurope.aksapp.io](http://parrot.05500238a78f.westeurope.aksapp.io)

you will see a happy bunch.

scale your nodeapp by running this:

```bash
kubectl scale deployment/nodeapp-nodeapp --replicas 3 --namespace phippy
```

and watch more brady ninjas come to life!

## issues?

we all have them. for starters, [this](https://github.com/bradygmsft/phippy-demo/issues/1). if you see issues as you're exploring, create them, then send us a pull request to resolve them. or one or the other. you know...

contribute, don't complain.