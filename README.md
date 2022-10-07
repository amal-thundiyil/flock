<p align='center'>
<img width="40%" src='./docs/images/logo.png'>
</p>

<h1>
<p align='center'>
flock - Distributed Job Scheduler
</p>
</h1>

<p align='center'>
<img src="https://github.com/amal-thundiyil/flock/actions/workflows/actions.yml/badge.svg">
<a href="https://github.com/amal-thundiyil/flock/blob/main/LICENSE"><img src="https://img.shields.io/badge/License-MIT-green.svg"></a>
<img src="https://visitor-badge.laobi.icu/badge?page_id=amal-thundiyil.flock">
</p>

## 📌 Introduction

Is your code taking too much time? Just flock it...

`flock` is a distributed system to run scheduled jobs against a server or a group of servers of any size. One of the machines is the leader and the others will be followers. If the leader fails or becomes unreachable, any other one will take over and reschedule all jobs to keep the system healthy.

## 🤖 Installation

```sh
make install-dev
```

## 👨‍💻️ Usage

```sh
flock up
# to specify file with desired config file
flock up -f `flock.json`
```

### Configuration sources

Settings can be specified in three ways (in order of precedence):

1. Command line arguments.
2. Environment variables starting with `FLOCK_`
3. flock.json config file

Config file will be loaded from the following paths:

```sh
$HOME/.flock/flock.json
```

Refer sample configuration file - [flock.yml](flock.yml).

## Description

Flock nodes can work in two modes, agents or servers.

A Flock agent is a cluster member that can handle job executions, run your scripts and return the resulting output to the server.

A Flock server is also a cluster member that send job execution queries to agents or other servers, so servers can execute jobs too.

Flock clusters have a leader, the leader is responsible of starting job execution queries in the cluster.

Any Flock agent or server acts as a cluster member and it's available to run scheduled jobs.

Default is for all nodes to execute each job. You can control what nodes run a job by specifying tags and a count of target nodes having this tag. This gives an unprecedented level of flexibility in runnig jobs across a cluster of any size and with any combination of machines you need.

All the execution responses will be gathered by the scheduler and stored in the database.
