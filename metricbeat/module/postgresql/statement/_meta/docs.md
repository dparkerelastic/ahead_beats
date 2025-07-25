This is the `statement` metricset of the PostgreSQL module.

This module collects information from the `pg_stat_statements` view, that keeps track of planning and execution statistics of all SQL statements executed by the server.

`pg_stat_statements` is included by an additional module in PostgreSQL. This module requires additional shared memory, and is disabled by default.

You can enable it by adding this module to the configuration as a shared preloaded library.

```
shared_preload_libraries = 'pg_stat_statements'
pg_stat_statements.max = 10000
pg_stat_statements.track = all
```

::::{note}
Preloading this library in your server will increase the memory usage of your PostgreSQL server. Use it with care.
::::


Once the server is started with this module, it starts collecting statistics about all statements executed. To make these statistics available in the `pg_stat_statements` view, the following statement needs to be executed in the server:

```sql
CREATE EXTENSION pg_stat_statements;
```

You can read more about the available options for this module in the [official documentation](https://www.postgresql.org/docs/13/pgstatstatements.html).

::::{note}
The PostgreSQL module of Filebeat is also able to collect information about statements executed in the server from its logs. You may chose which one is better for your needings. An important difference is that the Metricbeat module collects aggregated information when the statement is executed several times, but cannot know when each statement was executed. This information can be obtained from logs.
::::

