Raw notes - looks like the overlay network was not working on kernel
version of the ubuntu 14 version I was using... using wily did the 
trick (ami-05384865 in us-west-1).

The service was also exposed on a different port than I expected...

More experimentation to nail this down is needed...


<pre>
ubuntu@ip-10-0-0-195:~$ docker swarm init
Swarm initialized: current node (cv97xx6cmlfxmdaue9pjcw9ql) is now a manager.
ubuntu@ip-10-0-0-195:~$ docker network create -d overlay net1
16yq3f4vn5rddhwgesvvinjiq
ubuntu@ip-10-0-0-195:~$ docker service create --network net1 --name ping --publish 3000/tcp dasmith/ping
8joctah1e446qqn9qombrv53r
ubuntu@ip-10-0-0-195:~$ 
ubuntu@ip-10-0-0-195:~$ 
ubuntu@ip-10-0-0-195:~$ 
ubuntu@ip-10-0-0-195:~$ docker service ls
ID            NAME  REPLICAS  IMAGE         COMMAND
8joctah1e446  ping  1/1       dasmith/ping  
ubuntu@ip-10-0-0-195:~$ docker service tasks
"docker service tasks" requires exactly 1 argument(s).
See 'docker service tasks --help'.

Usage:  docker service tasks [OPTIONS] SERVICE

List the tasks of a service
ubuntu@ip-10-0-0-195:~$ docker service tasks ping
ID                         NAME    SERVICE  IMAGE         LAST STATE          DESIRED STATE  NODE
5xlllv3i9nx1buwceft47xye7  ping.1  ping     dasmith/ping  Running 13 seconds  Running        ip-10-0-0-195
ubuntu@ip-10-0-0-195:~$ curl localhost:3000
curl: (7) Failed to connect to localhost port 3000: Connection refused
ubuntu@ip-10-0-0-195:~$ docker service inspect ping
[
    {
        "ID": "8joctah1e446qqn9qombrv53r",
        "Version": {
            "Index": 25
        },
        "CreatedAt": "2016-06-24T19:47:21.837530239Z",
        "UpdatedAt": "2016-06-24T19:47:21.855023643Z",
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
                    "Target": "16yq3f4vn5rddhwgesvvinjiq"
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
            "Spec": {},
            "Ports": [
                {
                    "Protocol": "tcp",
                    "TargetPort": 3000,
                    "PublishedPort": 30000
                }
            ],
            "VirtualIPs": [
                {
                    "NetworkID": "88ezpa24l5bci9p324gnwzykf",
                    "Addr": "10.255.0.6/16"
                },
                {
                    "NetworkID": "16yq3f4vn5rddhwgesvvinjiq",
                    "Addr": "10.0.0.2/24"
                }
            ]
        }
    }
]
ubuntu@ip-10-0-0-195:~$ curl localhost:30000
PING

</pre>



## Three Nodes with Docker

(Note - this is in progress - seeing some problems with the overlay
network and nodes talking to each other, thought I saw it working
but now it isn't)

Ok... boot up three nodes on amazon, or wherever...

Pick one to be the master and initialize the swarm:

<pre>
ubuntu@ip-10-0-0-166:~$ docker swarm init
Swarm initialized: current node (8sqns2kanreyoxb08047iejo9) is now a manager.
</pre>

Log into the other nodes and join the swarm.

<pre>
ubuntu@ip-10-0-0-165:~$ docker swarm join ip-10-0-0-166:2377
This node joined a Swarm as a worker.
</pre>

On the master take a look at the nodes

<pre>
ubuntu@ip-10-0-0-166:~$ docker node ls
ID                           NAME           MEMBERSHIP  STATUS  AVAILABILITY  MANAGER STATUS
0fp18aur4lrvqogwer3iu57v5    ip-10-0-0-165  Accepted    Ready   Active        
8sqns2kanreyoxb08047iejo9 *  ip-10-0-0-166  Accepted    Ready   Active        Leader
dyqhnmi5riz89bh64w62l2fim    ip-10-0-0-167  Accepted    Ready   Active        
</pre>

Now create an overlay network for the swarm so the swarm hosts can
talk to each other.

<pre>
ubuntu@ip-10-0-0-166:~$ docker network create -d overlay app1net
1olux091glxc8e96e8o4ro08c
ubuntu@ip-10-0-0-166:~$ docker network ls
NETWORK ID          NAME                DRIVER              SCOPE
1olux091glxc        app1net             overlay             swarm               
daedc112ec4c        bridge              bridge              local               
3179713dd616        docker_gwbridge     bridge              local               
88fdf9b1b168        host                host                local               
al6u5f1a5la7        ingress             overlay             swarm               
95e05c9e318b        none                null                local               
</pre>

On one of the clients:

<pre>
ubuntu@ip-10-0-0-165:~$ docker network ls
NETWORK ID          NAME                DRIVER              SCOPE
477a680c8d0f        bridge              bridge              local               
5ca1c860621f        docker_gwbridge     bridge              local               
50f71faf235e        host                host                local               
al6u5f1a5la7        ingress             overlay             swarm               
2738e5cba853        none                null                local               
</pre>

Ok - create some services

<pre>
ubuntu@ip-10-0-0-166:~$ docker service create --name pingsvc --network app1net --publish 3000:3000 dasmith/ping
egatd0ayb974tjsavrs5kvz6j
ubuntu@ip-10-0-0-166:~$ docker service create --name pongsvc --network app1net --publish 4000:4000 dasmith/pong
8o92p3gao5twb8qfoji6l38b0
ubuntu@ip-10-0-0-166:~$ docker service create --name pingpong --network app1net --publish 8080:8080 dasmith/pingpong
2epjct0zu53mwo9wuik8oa2yh
ubuntu@ip-10-0-0-166:~$ docker service ls
ID            NAME      REPLICAS  IMAGE             COMMAND
2epjct0zu53m  pingpong  0/1       dasmith/pingpong  
8o92p3gao5tw  pongsvc   0/1       dasmith/pong      
egatd0ayb974  pingsvc   0/1       dasmith/ping      


</pre>

So... what's running on the master?

<pre>
ubuntu@ip-10-0-0-166:~$ docker ps
CONTAINER ID        IMAGE                 COMMAND             CREATED             STATUS              PORTS               NAMES
187da1efd4a8        dasmith/ping:latest   "/opt/ping"         3 minutes ago       Up 3 minutes        3000/tcp            ping.1.3z2w2al70qwmg91xm16tjshky
</pre>

Let's try the ping pong service on the master node:

<pre>



### Vagrant

Note the vagrant config is currently suspect - I can't intracluster
traffic over the overlay to work as it currently stands.

Note this installation assumes the following plugins are installed:

<pre>
vagrant plugin install vagrant-proxyconf
vagrant plugin install vagrant-hostmanager
</pre>

