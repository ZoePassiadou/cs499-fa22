Answer part 2 here.

First, we can see that our local DRAM uses two CPUs(numa 0 and numa 1). 
As we can see from our graph, the latency of a numa node for sendind packets
to itsef is smaller rather than sending to the other numa node, this is obvious considering the distance.
We can also say the same for the bandwidth here.

We can also see that the local disk latency is obviously much larger than the local DRAM latency, this was expected as ram latency is generally small. Also explaing the differnecr between the bandwidth of local DRAM and local disk as ram is very fast.

However the difference between local and remote disk latency is small because the network latency is
almost insignificant compared to the hard disk. In contrast, the latency of the remote DRAM is much larger compared to local DRAM
because the latency of ram is generally small so the network latency plays a significant difference.

When it comes to the bandwidth, we can see a significant difference between transfer rate to the numa nodes of the local DRAM compared
to the numa nodes of the remote DRAM, this is logical considering the limited bandwidth of the network.

However the bandwidth of the local and remote disk is the same because the local disk bandwidth is smaller than the
network bandwidth.