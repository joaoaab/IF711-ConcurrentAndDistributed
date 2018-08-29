import numpy as np
import matplotlib.pyplot as plt
from matplotlib.ticker import MaxNLocator
from collections import namedtuple


avg_tcp = 88.52175
stdev_tcp = 7.779

avg_udp = 12.4695
stdev_udp = 1.279

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