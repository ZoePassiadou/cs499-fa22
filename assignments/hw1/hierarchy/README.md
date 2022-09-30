Answer part 2 here.

First, we can see that our local DRAM uses two CPUs(numa 0 and numa 1) as we can also see from the specification of the machine (c220g1) , (Intel E5-2630 v3 8-core CPUs at 2.40 GHz x2).

As we can see from our graph, the latency of a numa node for sendind packets to itsef is 93.5ns for numa 0 and 92.3ns for numa 1.This is noticably smaller than the latency of the numa nodes sending packets to eachother (133.8ns from numa 0 to numa 1 and 133.7ns from numa 1 to numa 0) this is obvious considering the distance.
However we can  say the opposite for the bandwidth here, bandwidth form numa 0 to itself is 48956.9MB/s, bandwidth from numa 1 to numa 1 is 49016.3 MB/s, bandwidth from numa 0 to numa is 25411.3MB/s and from numa 1 to numa 0 is 25428MB/s. The rate of transefer to the same CPU is larger because it communicates with the RAM directly, however, from one CPU to another the bandwidth is decreased because of the limited bandwidth of CPU's bus speed (2*8 GT/s = 16GT/s ≈ 32GB/s). 


We can also see that the local disk latency (2300000ns) is obviously much larger than the local DRAM latency, this was expected as ram latency is generally small and disk latency is generally high. In contrast, the bandwidth of the local disk is much smaller since by nature, HDDS are way slower than RAM.

However the difference between local and remote disk latency is small because the network latency is
almost insignificant compared to the hard disk. In contrast, the latency of the remote DRAM is much larger compared to local DRAM
because the latency of ram is generally small so the network latency plays a significant difference.

When it comes to the bandwidth, we can see a significant difference between transfer rate to the numa nodes of the local DRAM compared
to the numa nodes of the remote DRAM, this is logical considering the limited bandwidth of the network.

However the bandwidth of the local and remote disk is the same because the local disk bandwidth is smaller than the
network bandwidth.



