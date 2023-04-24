import pandas as pd
import matplotlib.pyplot as plt

def read_docker_stats_csv(file_path):
    column_names = ['Container', 'Name', 'CPU %', 'Mem Usage', 'Net I/O', 'Block I/O', 'Mem %', 'PIDs', 'start', 'end']
    df = pd.read_csv(file_path, names=column_names, skiprows=1)
    df['CPU %'] = df['CPU %'].str.rstrip('%').astype('float')
    df['Mem %'] = df['Mem %'].str.rstrip('%').astype('float')
    df['Time'] = (df['start'] + df['end']) / 2
    return df

def plot_cpu_usage(df1, df2, df3, df4, container_name, node_edge_counts, min_start_time1, min_start_time2, min_start_time3, min_start_time4):
    container_df1 = df1[df1['Name'] == container_name]
    container_df2 = df2[df2['Name'] == container_name]
    container_df3 = df3[df3['Name'] == container_name]
    container_df4 = df4[df4['Name'] == container_name]
    
    plt.figure()
    plt.plot(container_df1['Time'] - min_start_time1, container_df1['CPU %'], label=f"{node_edge_counts[0]} Nodes, {node_edge_counts[1]} Edges")
    plt.plot(container_df2['Time'] - min_start_time2, container_df2['CPU %'], label=f"{node_edge_counts[2]} Nodes, {node_edge_counts[3]} Edges", alpha=.7)
    plt.plot(container_df3['Time'] - min_start_time3, container_df3['CPU %'], label=f"{node_edge_counts[4]} Nodes, {node_edge_counts[5]} Edges", alpha=.7)
    plt.plot(container_df4['Time'] - min_start_time4, container_df4['CPU %'], label=f"{node_edge_counts[6]} Nodes, {node_edge_counts[7]} Edges", alpha=.7)

    plt.xlabel('Iteration')
    plt.ylabel('CPU Usage (%)')
    plt.title(f'{container_name} - Docker CPU Usage')
    plt.legend()
    plt.xticks(rotation=45)
    plt.tight_layout()
    plt.savefig(f'{container_name}_cpu_usage_comparison.png')


def plot_memory_usage(df1, df2, df3, df4, container_name, node_edge_counts, min_start_time1, min_start_time2, min_start_time3, min_start_time4):
    container_df1 = df1[df1['Name'] == container_name]
    container_df2 = df2[df2['Name'] == container_name]
    container_df3 = df3[df3['Name'] == container_name]
    container_df4 = df4[df4['Name'] == container_name]

    plt.figure()
    plt.plot(container_df1['Time'] - min_start_time1, container_df1['Mem %'], label=f"{node_edge_counts[0]} Nodes, {node_edge_counts[1]} Edges")
    plt.plot(container_df2['Time'] - min_start_time2, container_df2['Mem %'], label=f"{node_edge_counts[2]} Nodes, {node_edge_counts[3]} Edges", alpha=.7)
    plt.plot(container_df3['Time'] - min_start_time3, container_df3['Mem %'], label=f"{node_edge_counts[4]} Nodes, {node_edge_counts[5]} Edges", alpha=.7)
    plt.plot(container_df4['Time'] - min_start_time4, container_df4['Mem %'], label=f"{node_edge_counts[6]} Nodes, {node_edge_counts[7]} Edges", alpha=.7)

    plt.xlabel('Iteration')
    plt.ylabel('Memory Usage (%)')
    plt.title(f'{container_name} - Docker Memory Usage')
    plt.legend()
    plt.xticks(rotation=45)
    plt.tight_layout()
    plt.savefig(f'{container_name}_memory_usage_comparison.png')


def plot_cpu_usage_rolling_avg(df1, df2, df3, df4, container_name, node_edge_counts, min_start_time1, min_start_time2, min_start_time3, min_start_time4, window_size=100):
    container_df1 = df1[df1['Name'] == container_name]
    container_df2 = df2[df2['Name'] == container_name]
    container_df3 = df3[df3['Name'] == container_name]
    container_df4 = df4[df4['Name'] == container_name]

    rolling_avg1 = container_df1['CPU %'].rolling(window=window_size).mean()
    rolling_avg2 = container_df2['CPU %'].rolling(window=window_size).mean()
    rolling_avg3 = container_df3['CPU %'].rolling(window=window_size).mean()
    rolling_avg4 = container_df4['CPU %'].rolling(window=window_size).mean()

    plt.figure()
    plt.plot(container_df1['Time'] - min_start_time1, rolling_avg1, label=f"{node_edge_counts[0]} Nodes, {node_edge_counts[1]} Edges")
    plt.plot(container_df2['Time'] - min_start_time2, rolling_avg2, label=f"{node_edge_counts[2]} Nodes, {node_edge_counts[3]} Edges", alpha=.7)
    plt.plot(container_df3['Time'] - min_start_time3, rolling_avg3, label=f"{node_edge_counts[4]} Nodes, {node_edge_counts[5]} Edges", alpha=.7)
    plt.plot(container_df4['Time'] - min_start_time4, rolling_avg4, label=f"{node_edge_counts[6]} Nodes, {node_edge_counts[7]} Edges", alpha=.7)

    plt.xlabel('Iteration')
    plt.ylabel('CPU Usage Rolling Average (%)')
    plt.title(f'{container_name} - Docker CPU Usage Rolling Average')
    plt.legend()
    plt.xticks(rotation=45)
    plt.tight_layout()
    plt.savefig(f'{container_name}_cpu_usage_rolling_avg_comparison.png')
def plot_memory_usage_rolling_avg(df1, df2, df3, df4, container_name, node_edge_counts, min_start_time1, min_start_time2, min_start_time3, min_start_time4, window_size=100):
    container_df1 = df1[df1['Name'] == container_name]
    container_df2 = df2[df2['Name'] == container_name]
    container_df3 = df3[df3['Name'] == container_name]
    container_df4 = df4[df4['Name'] == container_name]

    rolling_avg1 = container_df1['Mem %'].rolling(window=window_size).mean()
    rolling_avg2 = container_df2['Mem %'].rolling(window=window_size).mean()
    rolling_avg3 = container_df3['Mem %'].rolling(window=window_size).mean()
    rolling_avg4 = container_df4['Mem %'].rolling(window=window_size).mean()

    plt.figure()
    plt.plot(container_df1['Time'] - min_start_time1, rolling_avg1, label=f"{node_edge_counts[0]} Nodes, {node_edge_counts[1]} Edges")
    plt.plot(container_df2['Time'] - min_start_time2, rolling_avg2, label=f"{node_edge_counts[2]} Nodes, {node_edge_counts[3]} Edges", alpha=.7)
    plt.plot(container_df3['Time'] - min_start_time3, rolling_avg3, label=f"{node_edge_counts[4]} Nodes, {node_edge_counts[5]} Edges", alpha=.7)
    plt.plot(container_df4['Time'] - min_start_time4, rolling_avg4, label=f"{node_edge_counts[6]} Nodes, {node_edge_counts[7]} Edges", alpha=.7)

    plt.xlabel('Iteration')
    plt.ylabel('Memory Usage Rolling Average (%)')
    plt.title(f'{container_name} - Docker Memory Usage Rolling Average')
    plt.legend()
    plt.xticks(rotation=45)
    plt.tight_layout()
    plt.savefig(f'{container_name}_memory_usage_rolling_avg_comparison.png')


def calculate_second_half_average(df1, df2, df3, df4):
    # calculate the mid-point time for each dataset
    mid_time1 = (df1['start'].max() + df1['end'].min()) / 2
    mid_time2 = (df2['start'].max() + df2['end'].min()) / 2
    mid_time3 = (df3['start'].max() + df3['end'].min()) / 2
    mid_time4 = (df4['start'].max() + df4['end'].min()) / 2
    
    unique_container_names = set(df1['Name'].unique()).union(set(df2['Name'].unique())).union(set(df3['Name'].unique())).union(set(df4['Name'].unique()))
    
    for container_name in unique_container_names:
        container_df1 = df1[df1['Name'] == container_name]
        container_df2 = df2[df2['Name'] == container_name]
        container_df3 = df3[df3['Name'] == container_name]
        container_df4 = df4[df4['Name'] == container_name]
        
        # filter each dataset to include only entries after the mid-point time
        container_df1_second_half = container_df1[container_df1['start'] > mid_time1]
        container_df2_second_half = container_df2[container_df2['start'] > mid_time2]
        container_df3_second_half = container_df3[container_df3['start'] > mid_time3]
        container_df4_second_half = container_df4[container_df4['start'] > mid_time4]
        
        # calculate the average CPU and memory usage for each dataset in the second half
        avg_cpu_container_df1 = container_df1_second_half['CPU %'].mean()
        avg_mem_container_df1 = container_df1_second_half['Mem %'].mean()
        avg_cpu_container_df2 = container_df2_second_half['CPU %'].mean()
        avg_mem_container_df2 = container_df2_second_half['Mem %'].mean()
        avg_cpu_container_df3 = container_df3_second_half['CPU %'].mean()
        avg_mem_container_df3 = container_df3_second_half['Mem %'].mean()
        avg_cpu_container_df4 = container_df4_second_half['CPU %'].mean()
        avg_mem_container_df4 = container_df4_second_half['Mem %'].mean()
        
        # print the results
        print(f"Average CPU usage in the second half of {container_name} - dataset 1: {avg_cpu_container_df1:.2f}%")
        print(f"Average memory usage in the second half of {container_name} - dataset 1: {avg_mem_container_df1:.2f}%")
        print(f"Average CPU usage in the second half of {container_name} - dataset 2: {avg_cpu_container_df2:.2f}%")
        print(f"Average memory usage in the second half of {container_name} - dataset 2: {avg_mem_container_df2:.2f}%")
        print(f"Average CPU usage in the second half of {container_name} - dataset 3: {avg_cpu_container_df3:.2f}%")
        print(f"Average memory usage in the second half of {container_name} - dataset 3: {avg_mem_container_df3:.2f}%")
        print(f"Average CPU usage in the second half of {container_name} - dataset 4: {avg_cpu_container_df4:.2f}%")

def main():
    file_path1 = input("Enter the path to the first CSV file: ")
    df1 = read_docker_stats_csv(file_path1)

    file_path2 = input("Enter the path to the second CSV file: ")
    df2 = read_docker_stats_csv(file_path2)
    
    file_path3 = input("Enter the path to the third CSV file: ")
    df3 = read_docker_stats_csv(file_path3)

    file_path4 = input("Enter the path to the fourth CSV file: ")
    df4 = read_docker_stats_csv(file_path4)
    
    node_count1 = int(input("Enter the node count for the first dataset: "))
    edge_count1 = int(input("Enter the edge count for the first dataset: "))
    node_count2 = int(input("Enter the node count for the second dataset: "))
    edge_count2 = int(input("Enter the edge count for the second dataset: "))
    node_count3 = int(input("Enter the node count for the third dataset: "))
    edge_count3 = int(input("Enter the edge count for the third dataset: "))
    node_count4 = int(input("Enter the node count for the fourth dataset: "))
    edge_count4 = int(input("Enter the edge count for the fourth dataset: "))
    
    min_start_time1 = df1['start'].min()
    min_start_time2 = df2['start'].min()
    min_start_time3 = df3['start'].min()
    min_start_time4 = df4['start'].min()

    unique_container_names = set(df1['Name'].unique()).union(set(df2['Name'].unique())).union(set(df3['Name'].unique())).union(set(df4['Name'].unique()))

    for container_name in unique_container_names:
        plot_cpu_usage(df1, df2, df3, df4, container_name, (node_count1, edge_count1, node_count2, edge_count2, node_count3, edge_count3, node_count4, edge_count4), min_start_time1, min_start_time2, min_start_time3, min_start_time4)
        plot_memory_usage(df1, df2, df3, df4, container_name, (node_count1, edge_count1, node_count2, edge_count2, node_count3, edge_count3, node_count4, edge_count4), min_start_time1, min_start_time2, min_start_time3, min_start_time4)
        plot_cpu_usage_rolling_avg(df1, df2, df3, df4, container_name, (node_count1, edge_count1, node_count2, edge_count2, node_count3, edge_count3, node_count4, edge_count4), min_start_time1, min_start_time2, min_start_time3, min_start_time4)
        plot_memory_usage_rolling_avg(df1, df2, df3, df4, container_name, (node_count1, edge_count1, node_count2, edge_count2, node_count3, edge_count3, node_count4, edge_count4), min_start_time1, min_start_time2, min_start_time3, min_start_time4)

    print("Graphs saved for each container as [container_name]_cpu_usage_comparison.png, [container_name]_memory_usage_comparison.png, [container_name]_cpu_usage_rolling_avg_comparison.png, and [container_name]_memory_usage_rolling_avg_comparison.png")
    calculate_second_half_average(df1, df2, df3, df4)
if __name__ == "__main__":
    main()


