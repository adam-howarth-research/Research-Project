# Research-Project
This project contains the code required for conducting the experimental phase of  DEVELOPING A COST FRAMEWORK FOR DYNAMIC NETWORK SECURITY MONITORING: AN ANALYSIS OF BUSINESS PROCESS INDUCED NETWORK RECONFIGURATION

### Behavior Induction and Event Creation

  #### Entity Activity Generation
  Nodes on the virtual network must use the network like the nodes observed in the real-world networks the simulations are based on. This includes the services/protocols used as well as the volume and inter-activity time deltas. This software component will generate a series of node activities, which will be used to create the observational record input for detection algorithms.
 
### IPFIX Observation Creation

  #### Activity to Observation Mapper
  Node-activity events generated in the previous component must be converted to IPFIX for export and analysis

  #### Observation to ObservationDomain Mapper
  Observations must be associated with a location (ObservationPoint) where the observation was made. This instructs the cardinality of traffic observations and impacts analytic complexity of the detection function.
 
### Network Security Detection Analytics

  #### Analytic Consumers
  Endpoints must be created that consume flow records and run detection analytics against them. This generic execution container will run DDOS, Data Exfiltration and a general depth first search analytics based on methods found in the literature review. This container through the use of performance profiling data will log their resource utilization.
 
### Reporting and Result Aggregation

  #### Analytic Resouce Utilization Reporting
  Reporting node will format and log experimental output for reporting in the final thesis project. 

