import pandas as pd
import matplotlib.pyplot as plt

def read_docker_stats_csv(file_path):
    column_names = ['Container', 'Name', 'CPU %', 'Mem Usage', 'Net I/O', 'Block I/O', 'Mem %', 'PIDs', 'start', 'end']
    df = pd.read_csv(file_path, names=column_names, skiprows=1)
    df['CPU %'] = df['CPU %'].str.rstrip('%').astype('float')
    df['Mem %'] = df['Mem %'].str.rstrip('%').astype('float')
    df['Time'] = (df['start'] + df['end']) / 2
    return df

def plot_cpu_usage(datasets: list, container_name: str):

    plt.figure()


    for dataset in datasets:
        df = dataset[0]
        n_c = dataset[1]
        e_c = dataset[2]
        min_time = dataset[3]

        container_df = df[df['Name'] == container_name]

   
        plt.plot(container_df['Time'] - min_time, container_df['CPU %'], label=f"{n_c} Nodes, {e_c} Edges", alpha = .5)


    plt.xlabel('Iteration')
    plt.ylabel('CPU Usage (%)')
    plt.title(f'{container_name} - Docker CPU Usage')
    plt.legend()
    plt.xticks(rotation=45)
    plt.tight_layout()
    plt.savefig(f'{container_name}_cpu_usage_comparison.png')


def plot_memory_usage(datasets: list, container_name: str):

    plt.figure()


    for dataset in datasets:
        df = dataset[0]
        n_c = dataset[1]
        e_c = dataset[2]
        min_time = dataset[3]
        container_df = df[df['Name'] == container_name]

        plt.plot(container_df['Time'] - min_time, container_df['Mem %'], label=f"{n_c} Nodes, {e_c} Edges", alpha = .5)


    plt.xlabel('Iteration')
    plt.ylabel('Memory Usage (%)')
    plt.title(f'{container_name} - Docker Memory Usage')
    plt.legend()
    plt.xticks(rotation=45)
    plt.tight_layout()
    plt.savefig(f'{container_name}_memory_usage_comparison.png')


def plot_cpu_usage_rolling_avg(datasets: list, container_name: str, window_size=100):

    plt.figure()

    for dataset in datasets:
        df = dataset[0]
        n_c = dataset[1]
        e_c = dataset[2]
        min_time = dataset[3]
        container_df = df[df['Name'] == container_name]
        rolling_avg = container_df['CPU %'].rolling(window=window_size).mean()
        plt.plot(container_df['Time'] - min_time, rolling_avg, label=f"{n_c} Nodes, {e_c} Edges", alpha = .5)

    plt.xlabel('Iteration')
    plt.ylabel('CPU Usage Rolling Average (%)')
    plt.title(f'{container_name} - Docker CPU Usage Rolling Average')
    plt.legend()
    plt.xticks(rotation=45)
    plt.tight_layout()
    plt.savefig(f'{container_name}_cpu_usage_rolling_avg_comparison.png')

def plot_memory_usage_rolling_avg(datasets: list, container_name: str, window_size=100):

    plt.figure()

    for dataset in datasets:
        df = dataset[0]
        n_c = dataset[1]
        e_c = dataset[2]
        min_time = dataset[3]
        container_df = df[df['Name'] == container_name]
        rolling_avg = container_df['Mem %'].rolling(window=window_size).mean()
        plt.plot(container_df['Time'] - min_time, rolling_avg, label=f"{n_c} Nodes, {e_c} Edges", alpha = .5)

    plt.xlabel('Iteration')
    plt.ylabel('Memory Usage Rolling Average (%)')
    plt.title(f'{container_name} - Docker Memory Usage Rolling Average')
    plt.legend()
    plt.xticks(rotation=45)
    plt.tight_layout()
    plt.savefig(f'{container_name}_memory_usage_rolling_avg_comparison.png')


def calculate_second_half_average(datasets: list, container_name):

    for dataset in datasets:
        df = dataset[0]
        n_c = dataset[1]
        e_c = dataset[2]
        mid_time = (df['start'].max() + df['end'].min()) / 2
        container_df = df[df['Name'] == container_name]
        
         # filter each dataset to include only entries after the mid-point time
        container_df_second_half = container_df[container_df['start'] > mid_time]

        avg_cpu_container_df = container_df_second_half['CPU %'].mean()
        avg_mem_container_df = container_df_second_half['Mem %'].mean()
    
        # print the results
        print(f"Average CPU usage of {container_name} - dataset 1: {avg_cpu_container_df:.2f}%")
        print(f"Average memory usage of {container_name} - dataset 1: {avg_mem_container_df:.2f}%")

def main():
    file_path1 = input("Enter the path to the first CSV file: ")
    df1 = read_docker_stats_csv(file_path1)
    df1NE = file_path1.split('.')[0].split('_')[2:4]

    file_path2 = input("Enter the path to the second CSV file: ")
    df2 = read_docker_stats_csv(file_path2)
    df2NE = file_path2.split('.')[0].split('_')[2:4]

    file_path3 = input("Enter the path to the third CSV file: ")
    df3 = read_docker_stats_csv(file_path3)
    df3NE = file_path3.split('.')[0].split('_')[2:4]

    file_path4 = input("Enter the path to the fourth CSV file: ")
    df4 = read_docker_stats_csv(file_path4)
    df4NE = file_path4.split('.')[0].split('_')[2:4]

    file_path5 = input("Enter the path to the fifth CSV file: ")
    df5 = read_docker_stats_csv(file_path5)
    df5NE = file_path5.split('.')[0].split('_')[2:4]
    
    node_count1 = int(df1NE[0])
    edge_count1 = int(df1NE[1])
    node_count2 = int(df2NE[0])
    edge_count2 = int(df2NE[1])
    node_count3 = int(df3NE[0])
    edge_count3 = int(df3NE[1])
    node_count4 = int(df4NE[0])
    edge_count4 = int(df4NE[1])
    node_count5 = int(df5NE[0])
    edge_count5 = int(df5NE[1])

    
    min_start_time1 = df1['start'].min()
    min_start_time2 = df2['start'].min()
    min_start_time3 = df3['start'].min()
    min_start_time4 = df4['start'].min()
    min_start_time5 = df5['start'].min()

    dataset1 = (df1, node_count1, edge_count1, min_start_time1)
    dataset2 = (df2, node_count2, edge_count2, min_start_time2)
    dataset3 = (df3, node_count3, edge_count3, min_start_time3)
    dataset4 = (df4, node_count4, edge_count4, min_start_time4)
    dataset5 = (df5, node_count5, edge_count5, min_start_time5)

    datasets = [dataset1, dataset2, dataset3, dataset4, dataset5]

    unique_container_names = set(df1['Name'].unique()).union(set(df2['Name'].unique())).union(set(df3['Name'].unique())).union(set(df4['Name'].unique())).union(set(df5['Name'].unique()))

    for container_name in unique_container_names:
        plot_cpu_usage(datasets, container_name)
        plot_memory_usage(datasets, container_name)
        plot_cpu_usage_rolling_avg(datasets, container_name)
        plot_memory_usage_rolling_avg(datasets, container_name)
        calculate_second_half_average(datasets, container_name)

    print("Graphs saved for each container as [container_name]_cpu_usage_comparison.png, [container_name]_memory_usage_comparison.png, [container_name]_cpu_usage_rolling_avg_comparison.png, and [container_name]_memory_usage_rolling_avg_comparison.png")
if __name__ == "__main__":
    main()


