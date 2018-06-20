# the phippy demo

it's kind of a long story, but this repo will eventually contain something interesting. 

## prerequisites

everything has prerequisites. for this one, create yourself a dockerhub or azure container registry instance. then, log into it from the command line. **then**, make sure you have a Kubernetes cluster set up and that your client is set up to talk to it. 

## good? 

to get started, do this:

```bash
git clone https://github.com/bradygmsft/phippy-demo.git
cd phippy-demo
cd dotnetapp
draft create
draft up
cd ..
cd nodeapp
draft create
draft up
cd ..
```

## all done? 

test your deployment by running this:

```bash
kubectl get svc
```

you should see something like this, if you had a clean cluster prior to the deployment. 

```
NAME                 TYPE        CLUSTER-IP     EXTERNAL-IP   PORT(S)    AGE
dotnetapp-csharp     ClusterIP   10.0.190.92   <none>        8080/TCP   2m
kubernetes           ClusterIP   10.0.0.1       <none>        443/TCP    1d
nodeapp-javascript   ClusterIP   10.0.43.125    <none>        8080/TCP   1m
```

## allow inbound access 

open the `nodeapp\charts\values.yaml` file and the `dotnetapp\charts\values.yaml` file. note the `service` property in each of these:

```yaml
service:
  name: node
  type: ClusterIP
  externalPort: 8080
  internalPort: 8080
```

change `nodeapp\charts\values.yaml` to look like this:

```yaml
service:
  name: node
  type: LoadBalancer
  externalPort: 5001
  internalPort: 3000
```

change `dotnetapp\charts\values.yaml` to look like this:

```yaml
service:
  name: dotnetcore
  type: LoadBalancer
  externalPort: 5000
  internalPort: 80
```

once you do this, the dotnet app should respond on port 5000, and the Node.js app on 5001. to deploy the changes, run these commands:

```bash
cd dotnetapp
draft up
cd ..
cd nodeapp
draft up
cd ..
kubectl get svc
```

when the command executes you'll see a table of all the ip addresses for each of the sites:

 

## issues? 

we all have them. for starters, [this](https://github.com/bradygmsft/phippy-demo/issues/1). if you see issues as you're exploring, create them, then send us a pull request to resolve them. or one or the other. you know...

contribute, don't complain. 