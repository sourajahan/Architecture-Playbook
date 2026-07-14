# Conceptualizing Spark's DAG approach vs MapReduce's Sequential steps
# MapReduce forces writing to disk at every step (Slow).
# Spark keeps data in Memory (Fast).

def optimized_pipeline(data_rdd):
    # This creates a "Logical Plan" (The DAG)
    # The actual execution happens only when we trigger an "Action"
    return data_rdd.filter(lambda x: x.is_valid()) \
                    .map(lambda x: x.to_revenue()) \
                    .reduceByKey(lambda a, b: a + b) 
                    # Everything stays in RAM until this 'reduce' action triggers.
