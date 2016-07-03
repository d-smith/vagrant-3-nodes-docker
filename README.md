## Swarm Experiment

### AWS

The ubuntu wily image has the networking support needed for an overlay network to work across
cluster members (i.e. ubuntu/images/hvm/ubuntu-wily-15.10-amd64-server-20160222 (ami-05384865))

Set up the swarm

<pre>
ubuntu@ip-10-0-0-88:~$ docker swarm init
ubuntu@ip-10-0-0-87:~$ docker swarm join ip-10-0-0-88:2377
ubuntu@ip-10-0-0-89:~$ docker swarm join ip-10-0-0-88:2377


ubuntu@ip-10-0-0-88:~$ docker service create --network net1 --name ping --publish 3000/tcp dasmith/ping
ubuntu@ip-10-0-0-88:~$ docker service inspect ping
[
    {
        "ID": "ezbdjxzs0hejekszdf03pgtq3",
        "Version": {
            "Index": 25
        },
        "CreatedAt": "2016-07-03T14:28:01.223979743Z",
        "UpdatedAt": "2016-07-03T14:28:01.22694425Z",
        "Spec": {
            "Name": "ping",
            "TaskTemplate": {
                "ContainerSpec": {
                    "Image": "dasmith/ping"
                },
                "Resources": {
                    "Limits": {},
                    "Reservations": {}
                },
                "RestartPolicy": {
                    "Condition": "any",
                    "MaxAttempts": 0
                },
                "Placement": {}
            },
            "Mode": {
                "Replicated": {
                    "Replicas": 1
                }
            },
            "UpdateConfig": {},
            "Networks": [
                {
                    "Target": "6mih4mfqvbls6aipbjtp3wizi"
                }
            ],
            "EndpointSpec": {
                "Mode": "vip",
                "Ports": [
                    {
                        "Protocol": "tcp",
                        "TargetPort": 3000
                    }
                ]
            }
        },
        "Endpoint": {
            "Spec": {
                "Mode": "vip",
                "Ports": [
                    {
                        "Protocol": "tcp",
                        "TargetPort": 3000
                    }
                ]
            },
            "Ports": [
                {
                    "Protocol": "tcp",
                    "TargetPort": 3000,
                    "PublishedPort": 30000
                }
            ],
            "VirtualIPs": [
                {
                    "NetworkID": "awnru4fqy5yne4vwakt3z8qsk",
                    "Addr": "10.255.0.6/16"
                },
                {
                    "NetworkID": "6mih4mfqvbls6aipbjtp3wizi",
                    "Addr": "10.0.0.2/24"
                }
            ]
        }
    }
]

ubuntu@ip-10-0-0-88:~$ docker service create --network net1 --name pong --publish 4000/tcp dasmith/pong
ubuntu@ip-10-0-0-88:~$ docker service create --network net1 --name pingpong --publish 8080/tcp dasmith/pingpong

ubuntu@ip-10-0-0-88:~$ docker service inspect pingpong
[
    {
        "ID": "1ii7cqm3obx5qpv09xewnu5jk",
        "Version": {
            "Index": 41
        },
        "CreatedAt": "2016-07-03T14:32:40.822936415Z",
        "UpdatedAt": "2016-07-03T14:32:40.825559289Z",
        "Spec": {
            "Name": "pingpong",
            "TaskTemplate": {
                "ContainerSpec": {
                    "Image": "dasmith/pingpong"
                },
                "Resources": {
                    "Limits": {},
                    "Reservations": {}
                },
                "RestartPolicy": {
                    "Condition": "any",
                    "MaxAttempts": 0
                },
                "Placement": {}
            },
            "Mode": {
                "Replicated": {
                    "Replicas": 1
                }
            },
            "UpdateConfig": {},
            "Networks": [
                {
                    "Target": "6mih4mfqvbls6aipbjtp3wizi"
                }
            ],
            "EndpointSpec": {
                "Mode": "vip",
                "Ports": [
                    {
                        "Protocol": "tcp",
                        "TargetPort": 8080
                    }
                ]
            }
        },
        "Endpoint": {
            "Spec": {
                "Mode": "vip",
                "Ports": [
                    {
                        "Protocol": "tcp",
                        "TargetPort": 8080
                    }
                ]
            },
            "Ports": [
                {
                    "Protocol": "tcp",
                    "TargetPort": 8080,
                    "PublishedPort": 30002
                }
            ],
            "VirtualIPs": [
                {
                    "NetworkID": "awnru4fqy5yne4vwakt3z8qsk",
                    "Addr": "10.255.0.10/16"
                },
                {
                    "NetworkID": "6mih4mfqvbls6aipbjtp3wizi",
                    "Addr": "10.0.0.6/24"
                }
            ]
        }
    }
]
ubuntu@ip-10-0-0-88:~$ curl localhost:30002
Get http://pingsvc:3000: dial tcp: lookup pingsvc on 127.0.0.11:53: read udp 127.0.0.1:48560->127.0.0.11:53: i/o timeout

ubuntu@ip-10-0-0-88:~$ docker service create --network net1 --name pingsvc --publish 3000/tcp dasmith/ping
a7hs285dyrws1qtc65xm3t1k0
ubuntu@ip-10-0-0-88:~$ docker service create --network net1 --name pongsvc --publish 4000/tcp dasmith/pong
5k0vn3j3zqe5qtoxc6ovjnne7

ubuntu@ip-10-0-0-88:~$ curl localhost:30002
PING
PONG

ubuntu@ip-10-0-0-88:~$ docker ps
CONTAINER ID        IMAGE                 COMMAND             CREATED             STATUS              PORTS               NAMES
1118ec79768a        dasmith/pong:latest   "/opt/pong"         16 minutes ago      Up 16 minutes                           pongsvc.1.5v14ta82mhtct3h79dnhrl8yl
11328efe6c9c        dasmith/ping:latest   "/opt/ping"         22 minutes ago      Up 22 minutes       3000/tcp            ping.1.455ghhrxkomgt2q50sf5kp1vi

ubuntu@ip-10-0-0-87:~$ docker ps
CONTAINER ID        IMAGE                     COMMAND             CREATED             STATUS              PORTS               NAMES
7ef050be7a30        dasmith/ping:latest       "/opt/ping"         17 minutes ago      Up 17 minutes       3000/tcp            pingsvc.1.3nkpqqblwiisg5mb2cdfjthce
a796115a70f8        dasmith/pingpong:latest   "/opt/pingpong"     18 minutes ago      Up 18 minutes       8080/tcp            pingpong.1.8y7gsj4xck22m2rmrmxnf021f

ubuntu@ip-10-0-0-89:~$ docker ps
CONTAINER ID        IMAGE                 COMMAND             CREATED             STATUS              PORTS               NAMES
9d1bbdc7fdb6        dasmith/pong:latest   "/opt/pong"         21 minutes ago      Up 21 minutes                           pong.1.771o43q5og0bxrhdkpmga3si2

ubuntu@ip-10-0-0-89:~$ curl ip-10-0-0-88:30002
PING
PONG
</pre>


From the above, the curl from worker 89 to leader 88 on the advertised service port (88) has the master route 
the request to worker 87 where pingpong is running. Pingpong makes calls to pingsvc (running on 87) and pongsvc 
(running on 88).

## Vagrant

Still trying to get this to work...

* Ubuntu 14 out of the box doesn't have the kernel support needed for the overlay network
* Updating the kernel lets the services boot given the overlay network, but the routing does
not appear to work
* The wily image does not start/work with Vagrant


### Vagrant

Note the vagrant config is currently suspect - I can't intracluster
traffic over the overlay to work as it currently stands.

Note this installation assumes the following plugins are installed:

<pre>
vagrant plugin install vagrant-proxyconf
vagrant plugin install vagrant-hostmanager
vagrant plugin install vagrant-reload 
</pre>

