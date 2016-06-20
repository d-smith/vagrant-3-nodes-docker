## Three Nodes with Docker

This project boots 3 vagrant hosts within the same private network, 
provisioning Docker on each of the nodes. At this point in time
the release candidate of docker (1.12 rc2) is installed to allow
forming a swarm cluster.

To create a cluster, run docker swarm init on one of the nodes, then
docker swarm join on the others (referencing the starting node).

Note this installation assumes the following plugs are installed:

<pre>
vagrant plugin install vagrant-proxyconf
vagrant plugin install vagrant-hostmanager
</pre>
