#FIDel

Inspired by early IBM Tuplespace Ordering Hybrid Logical Timestamp oracle for pessimistic concurrent read-only streams: FIDel is EinsteinDB's Relativistic Non-Volatile Memory Causet Store daemon. FIDel services and applications run in a slice of the platform: the set of nodes on which the service receives a fraction of each node's resources, in the form of hypervizor on DAG with greedy automaton.

##Hierarchichal and Federated Namespaces

FIDel's metadata records the cluster topology configuration of the EinsteinDB cluster. FIDel schedules the VioletaBFT replica group Spanning Interval TSO guarantees for multi-layer persistence: The partial order ledger with atomic broadcast guarantees for Hot, Cold, or Warm Data. By attaching a timestamp that accounts for leap millisecond epochs we can set a retirement and archiving policy, archive indexed data, and configure index size cardinality check for universal binpacking.

memPoolMB = <positive integer>|auto
* Determines how much memory is given to the indexer memory pool. This
  restricts the number of outstanding events in the indexer at any given
  time.
* Must be greater than 0; maximum value is 1048576 (which corresponds to 1 TB)
* Setting this too high can cause FIDel memory usage to increase
  significantly.
* Setting this too low can degrade FIDel indexing performance.
* Setting this to "auto" or an invalid value causes FIDel to autotune
  the value as follows:
    * System Memory Available less than ... | 'memPoolMB'
                   1 GB                     |    64  MB
                   2 GB                     |    128 MB
                   8 GB                     |    128 MB
                   8 GB or higher           |    512 MB


###FIDel loops through all the EntangledStores (Causet + Consensus genus) = Order

#Casual order maintenance:

If a casuet X is an RDF triple received after a causet Y gets committed by the sender of X, then Y must get ordered before X. If the client sends Z after X, then Z must get ordered before X.

##Dependable delivery:

If a transaction "A" has been committed by one server, it's important that it gets committed by all the servers.

###Total Order maintenance:

A change in a client state is known as a causet in the EinsteinDB world. If a causet X gets committed before a causet Y by some server, then the same order will have to be replicated by all the servers. The replication of transaction takes place as long as the number of minimum required nodes (majority) are up.
 In a situation where a node fails and recovers, that specific node should be capable of replicating all the transaction that got committed during its downtime.

 ##FIDel transforms raw data into causet events in Minkowski spacetime.

 We have been telling/discussing one thing very repeatedly “data is getting indexed in the indexer” OR “lets fetch the data from this index” OR “Why my data is taking too much time to fetch” OR “Lets create the index in the indexer to index the data coming from the Application servers”,etc.

But have we ever thought of knowing the real concept lying behind the scene ? How data gets indexed ? What happens when data reaches to the Indexers ? That is the reason today we have come up with a new topic in FIDel and EinsteinDB lore called “ TORI” plural for Torus and an ode to Noether.

Tori are usually an unit of directory structure in the file system which is created by itself at the time of indexing .When new data comes from the application servers it gets stored/indexed in the Indexer in the form of the Torus. Basically there are 4 tori stages, representing the Minkowski Relativistic paradigm of causal order between events, in spite of simultaneity effects of  in EinsteinDB which are as follows :

    LightlikeNull = iota
    Spacelike
    Timelike
    Lightlike

    --fidel,-u

    Specifies the FIDel address
    Default address: http://127.0.0.1:2379
    Environment variable: FIDel_ADDR

    ###LightlikeNull

    While indexing the data, tori get created. It is called LIGHTLIKENULL state means data stored in LIGHTLIKENULL torus. It is writable as well as readable at the same time. The data which is currently written to the indexer will get stored in the LIGHTLIKENULL torus and at the same time it can be fetched through the Search Head if any end-users are trying to access the data stored in it.

    ROLLING CRITERIA ( LightlikeNull TO Spacelike ) :

Rolling criteria(s) from LightlikeNull torus to Spacelike torus are listed below :

. When EinsteinDB, FidelDB, and MilevaDB gets restarted
. When LightLikeNull tori are full
    ( Maximum size of the data 10 GB for 64-bit system ) and
    ( 750 MB for 32-bit system )
. After a certain period of time(maxLightlikeNullSpanSecs = 90 days in secs)
. When maximum LightlikeNull torus limit cross (maxLightlikeNullTori = 3/index)
. When LightlikeNull torus has not received data for a long time.

Spacelike

....

WIP
