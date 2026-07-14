# Conceptualizing Spark's DAG approach vs MapReduce's Sequential steps
# Principal-grade Spark code avoids unnecessary shuffling.

def optimized_analytics_pipeline(df):
    # 1. Declarative logic (We tell Spark WHAT to do, not HOW)
    # Spark builds a logical plan (DAG) here.
    
    pipeline = df.filter(df["event_type"] == "purchase") \
                 .select("user_id", "amount", "timestamp") \
                 .groupBy("user_id") \
                 .agg({"amount": "sum"})
    
    # 2. Lazy Execution
    # Nothing is computed yet. Spark optimizes the graph to avoid extra I/O.
    
    return pipeline.write.parquet("s3://analytics-output/") 
    # Action triggers the execution.
