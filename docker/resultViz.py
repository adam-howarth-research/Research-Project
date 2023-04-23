import pandas as pd
import matplotlib.pyplot as plt

def read_docker_stats_csv(file_path):
    column_names = ['Container', 'Name', 'CPU %', 'Mem Usage', 'Net I/O', 'Block I/O', 'Mem %', 'PIDs', 'start', 'end']
    df = pd.read_csv(file_path, names=column_names, skiprows=1)
    df['CPU %'] = df['CPU %'].str.rstrip('%').astype('float')
    df['Mem %'] = df['Mem %'].str.rstrip('%').astype('float')
    df['Time'] = (df['start'] + df['end']) / 2
    return df

def plot_cpu_usage(df1, df2, container_name, node_edge_counts, min_start_time1, min_start_time2):
    container_df1 = df1[df1['Name'] == container_name]
    container_df2 = df2[df2['Name'] == container_name]
    
    plt.figure()
    plt.plot(container_df1['Time'] - min_start_time1, container_df1['CPU %'], label=f"{node_edge_counts[0]} Nodes, {node_edge_counts[1]} Edges")
    plt.plot(container_df2['Time'] - min_start_time2, container_df2['CPU %'], label=f"{node_edge_counts[2]} Nodes, {node_edge_counts[3]} Edges")
    
    plt.xlabel('Time (average of start and end, relative to the earliest start of each dataset)')
    plt.ylabel('CPU Usage (%)')
    plt.title(f'{container_name} - Docker CPU Usage')
    plt.legend()
    plt.xticks(rotation=45)
    plt.tight_layout()
    plt.savefig(f'{container_name}_cpu_usage_comparison.png')

def plot_memory_usage(df1, df2, container_name, node_edge_counts, min_start_time1, min_start_time2):
    container_df1 = df1[df1['Name'] == container_name]
    container_df2 = df2[df2['Name'] == container_name]
    
    plt.figure()
    plt.plot(container_df1['Time'] - min_start_time1, container_df1['Mem %'], label=f"{node_edge_counts[0]} Nodes, {node_edge_counts[1]} Edges")
    plt.plot(container_df2['Time'] - min_start_time2, container_df2['Mem %'], label=f"{node_edge_counts[2]} Nodes, {node_edge_counts[3]} Edges")
    
    plt.xlabel('Time (average of start and end, relative to the earliest start of each dataset)')
    plt.ylabel('Memory Usage (%)')
    plt.title(f'{container_name} - Docker Memory Usage')
    plt.legend()
    plt.xticks(rotation=45)
    plt.tight_layout()
    plt.savefig(f'{container_name}_memory_usage_comparison.png')

def main():
    file_path1 = input("Enter the path to the first CSV file: ")
    df1 = read_docker_stats_csv(file_path1)

    file_path2 = input("Enter the path to the second CSV file: ")
    df2 = read_docker_stats_csv(file_path2)
    
    node_count1 = int(input("Enter the node count for the first dataset: "))
    edge_count1 = int(input("Enter the edge count for the first dataset: "))
    node_count2 = int(input("Enter the node count for the second dataset: "))
    edge_count2 = int(input("Enter the edge count for the second dataset: "))
    
    min_start_time1 = df1['start'].min
import pandas as pd
import matplotlib.pyplot as plt

def read_docker_stats_csv(file_path):
    column_names = ['Container', 'Name', 'CPU %', 'Mem Usage', 'Net I/O', 'Block I/O', 'Mem %', 'PIDs', 'start', 'end']
    df = pd.read_csv(file_path, names=column_names, skiprows=1)
    df['CPU %'] = df['CPU %'].str.rstrip('%').astype('float')
    df['Mem %'] = df['Mem %'].str.rstrip('%').astype('float')
    df['Time'] = (df['start'] + df['end']) / 2
    return df

def plot_cpu_usage(df1, df2, container_name, node_edge_counts, min_start_time1, min_start_time2):
    container_df1 = df1[df1['Name'] == container_name]
    container_df2 = df2[df2['Name'] == container_name]
    
    plt.figure()
    plt.plot(container_df1['Time'] - min_start_time1, container_df1['CPU %'], label=f"{node_edge_counts[0]} Nodes, {node_edge_counts[1]} Edges")
    plt.plot(container_df2['Time'] - min_start_time2, container_df2['CPU %'], label=f"{node_edge_counts[2]} Nodes, {node_edge_counts[3]} Edges")
    
    plt.xlabel('Time (average of start and end, relative to the earliest start of each dataset)')
    plt.ylabel('CPU Usage (%)')
    plt.title(f'{container_name} - Docker CPU Usage')
    plt.legend()
    plt.xticks(rotation=45)
    plt.tight_layout()
    plt.savefig(f'{container_name}_cpu_usage_comparison.png')

def plot_memory_usage(df1, df2, container_name, node_edge_counts, min_start_time1, min_start_time2):
    container_df1 = df1[df1['Name'] == container_name]
    container_df2 = df2[df2['Name'] == container_name]
    
    plt.figure()
    plt.plot(container_df1['Time'] - min_start_time1, container_df1['Mem %'], label=f"{node_edge_counts[0]} Nodes, {node_edge_counts[1]} Edges")
    plt.plot(container_df2['Time'] - min_start_time2, container_df2['Mem %'], label=f"{node_edge_counts[2]} Nodes, {node_edge_counts[3]} Edges")
    
    plt.xlabel('Time (average of start and end, relative to the earliest start of each dataset)')
    plt.ylabel('Memory Usage (%)')
    plt.title(f'{container_name} - Docker Memory Usage')
    plt.legend()
    plt.xticks(rotation=45)
    plt.tight_layout()
    plt.savefig(f'{container_name}_memory_usage_comparison.png')

def main():
    file_path1 = input("Enter the path to the first CSV file: ")
    df1 = read_docker_stats_csv(file_path1)

    file_path2 = input("Enter the path to the second CSV file: ")
    df2 = read_docker_stats_csv(file_path2)
    
    node_count1 = int(input("Enter the node count for the first dataset: "))
    edge_count1 = int(input("Enter the edge count for the first dataset: "))
    node_count2 = int(input("Enter the node count for the second dataset: "))
    edge_count2 = int(input("Enter the edge count for the second dataset: "))
    
    min_start_time1 = df1['start'].min()
    min_start_time2 = df2['start'].min()

    unique_container_names = set(df1['Name'].unique()).union(set(df2['Name'].unique()))

    for container_name in unique_container_names:
        plot_cpu_usage(df1, df2, container_name, (node_count1, edge_count1, node_count2, edge_count2), min_start_time1, min_start_time2)
        plot_memory_usage(df1, df2, container_name, (node_count1, edge_count1, node_count2, edge_count2), min_start_time1, min_start_time2)

    print("Graphs saved for each container as [container_name]_cpu_usage_comparison.png and [container_name]_memory_usage_comparison.png")

if __name__ == "__main__":
    main()

