import matplotlib.pyplot as plt

network_sizes = ['100N-1000E', '1000N-10000E', '10000N-100000E', '100000N-1000000E']
containers = ['Baseline', 'KLDDoS', 'PCR', 'BFS']

cpu_usage = [
    [0.05, 0.05, 0.13, 0.40],
    [0.03, 0.17, 0.85, 8.40],
    [0.01, 0.02, 0.03, 0.09],
    [0.09, 0.70, 3.12, 5.43]
]

memory_usage = [
    [0.05, 0.06, 0.05, 0.05],
    [0.06, 0.16, 0.89, 8.11],
    [0.06, 0.11, 0.11, 0.54],
    [0.11, 0.17, 0.44, 0.32]
]

x = range(len(network_sizes))

# Define vertical line positions
vline_positions = [0.5, 1.5, 2.5]
vline_labels = ['17 Stores\n(239N-10740E)', '34 Stores\n(410N-20999E)', '100 Stores\n(1083N-59725E)']

# Plot CPU usage
fig, ax = plt.subplots()
for i, container in enumerate(containers):
    ax.plot(x, cpu_usage[i], label=container)

for pos, label in zip(vline_positions, vline_labels):
    ax.axvline(pos, color='gray', linestyle='--', alpha=0.5)
    ax.text(pos, ax.get_ylim()[1] * 0.7, label, rotation=90, ha='right', va='top', color='gray', alpha=0.7)

plt.xticks(x, network_sizes)
plt.xlabel("Network Size")
plt.ylabel("Average CPU Usage (%)")
plt.title("Average CPU Usage by Container and Network Size")
plt.legend()
plt.show()

# Plot Memory usage
fig, ax = plt.subplots()
for i, container in enumerate(containers):
    ax.plot(x, memory_usage[i], label=container)

for pos, label in zip(vline_positions, vline_labels):
    ax.axvline(pos, color='gray', linestyle='--', alpha=0.5)
    ax.text(pos, ax.get_ylim()[1] * 0.7, label, rotation=90, ha='right', va='top', color='gray', alpha=0.7)

plt.xticks(x, network_sizes)
plt.xlabel("Network Size")
plt.ylabel("Average Memory Usage (%)")
plt.title("Average Memory Usage by Container and Network Size")
plt.legend()
plt.show()
