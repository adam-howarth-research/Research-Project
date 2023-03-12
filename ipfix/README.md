# IPFIX Observation Creation

  ### Activity to Observation Mapper
  Node-activity events generated in the previous component must be converted to IPFIX for export and analysis

  ### Observation to ObservationDomain Mapper
  Observations must be associated with a location (ObservationPoint) where the observation was made. This instructs the cardinality of traffic observations and impacts analytic complexity of the detection function.