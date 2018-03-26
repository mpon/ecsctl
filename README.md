# ecswalk

Show information for Amazon Elastic Container Service(ECS) like the AWS management console.

## Demo

![_screenshot/services.gif](_screenshot/services.gif)

![_screenshot/tasks.gif](_screenshot/tasks.gif)


## Usage

### Get Information from ECS Interactively

* get ECS services by selecting cluster interactively

```console
$ ecswalk walk services
```

* get ECS tasks by selecting cluster and service interactively

```console
$ ecswalk walk tasks
```

### Get Information from ECS

* get ECS clusters

```console
$ ecswalk get clusters
```

* get ECS services for specified ECS cluster

```console
$ ecswalk get service -c default
```

* get ECS tasks for specified ECS cluster and service

```console
$ ecswalk get tasks -c default -s web-service
```

### TODO: Run ECS task

* [ ] run a task with running serivce task definition
* [ ] polling cloudwatch logs and task status

```console
$ ecswalk run --c default --s web-service echo hello
```


## Options

* you can set aws configure profile by `.ecswalk.yaml`

```yaml
profile: my-aws
```
