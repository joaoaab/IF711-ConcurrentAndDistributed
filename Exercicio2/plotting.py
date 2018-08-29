import numpy as np
import matplotlib.pyplot as plt
from matplotlib.ticker import MaxNLocator
from collections import namedtuple


avg_tcp = 940.4
stdev_tcp = 77.36

avg_udp = 107.478
stdev_udp = 4.8226

protocols = ["tcp", "udp"]

x_pos = np.arange(len(protocols))
avgs = [avg_tcp, avg_udp]
stds = [stdev_tcp, stdev_udp]
fig, ax = plt.subplots()
ax.bar(x_pos, avgs, yerr=stds, align='center', alpha=0.5, ecolor="black", capsize = 10)
ax.set_ylabel("Time to send and receive(ms)")
ax.set_xticks(x_pos)
ax.set_xticklabels(protocols)
ax.set_title('Time to send and receive N=10000 in Milliseconds')
ax.yaxis.grid(True)

plt.show()