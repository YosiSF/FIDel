
## FIDel

FIDel is an IPFS implementation of EinsteinDB with some additional features such as the ability to federate data from multiple IPFS nodes into one or more EinsteinDB nodes and the ability to host EinsteinDB on an IPFS node.

## Directory structure
- `fidel`: The fidel daemon running as a service.
- `tests`: Unit tests for fidel.
- `tools`: Contains scripts for building, testing and running fidel.
- `bin`: Contains the fidel binary after compilation.
- `src`: Contains the source code for fidel.
- `docs`: Contains the documentation for fidel.
- `config`: Contains the configuration files for fidel.
- `data`: Contains the data for fidel.
- `logs`: Contains the logs for fidel.
- `tmp`: Contains temporary files for fidel.
- `vendor`: Contains the vendor libraries for fidel.
- `.gitignore`: Contains the .gitignore file for fidel.
- `.gitattributes`: Contains the .gitattributes file for fidel.
- `.travis.yml`: Contains the .travis.yml file for fidel.
- `.editorconfig`: Contains the .editorconfig file for fidel.
- `.gitlab-ci.yml`: Contains the .gitlab-ci.yml file for fidel.

For more details, please refer to [FIDel documentation](https://github.com/YosiSF/FIDel/tree/master/docs)

##Theory
FIDel is a federated IPFS which support data federation and storage management. FIDel contains two components: a modified IPFS node, and a set of user-space tools to manage federated IPFS nodes.

## Modified IPFS Components
The original [IPFS](https://github.com/ipfs) is an open source project that aims to make the web faster, safer, and more open. It stores files in such a way as to permit retrieval by content rather than location or filename. The technique used is based on BitTorrent technology but without relying on trackers (which are centralized). Instead it uses distributed hash tables (DHTs), similar to those used by torrents making users both clients and servers in one decentralized network using only their internet connection for communication with other peers in the network. To be able to retrieve content from anywhere this network does not rely on any particular server being available; instead anyone who has previously accessed the file acts as part of the distribution chain when someone else comes along looking for it again later.[[1]](#references)

The FIDel component modification includes three parts: `pinning mechanism`, `federator` and `storage management`.

### Pinning Mechanism
Pinning mechanism allows us to pin data in IPFS node. We will use this mechanism to pin files in local storage or remote storage(the federated nodes).

### Federator
The federator is a process that runs on the FIDel node and controls all of the federation operations. It interacts with other components via REST APIs, `determine which data need to be stored locally` and `which data should be stored on remote nodes`. The current implementation only supports one-to-one federation between two IPFS nodes but we hope that this can be extended in future implementations.

### Storage Manager
A storage manager is added into every FIDel node, it has three main functionalities:
1.	Record disk usage periodically;
2.	Determine which blocks are not critical for later retrieval;
3.	Move non-critical blocks from local storage onto remote nodes when space running out; Reclaim space by cleaning up those moved blocks if necessary (for example the disk usage lowered); Move back some important blocks from remote node when necessary (for example the disk usage increased);

## User Space Tools Modification

In addition to modifying existing components of IPFS, new user space tools have been created as part of FIDel.
We have created a new command to allow the user to specify which other IPFS nodes will be used for storing data and which ones are going to store this node’s data. In addition, another tool has been added that provides information about the current state of data in each node as well as disk usage statistics for both local and remote storage. Both tools can be accessed via REST API calls or using a command line interface that is similar to `ipfs` commands but with their own set of parameters specific to federation operations.

## Building FIDel
FIDel requires Go 1.8+ (1.9 recommended). To build fidel:
  ```sh 
make all   //Note: please use "sudo make all" if you want run your binaries system-wide 
```

Once built, you can find the compiled binary under bin/fidel directory named “fidel” (or “fideld” on Windows), depending on whether it was compiled in debug mode or not. The binary is already linked with libsnappy so there is no need for an external dependency except when statically compiling it into your OS distribution image (for example Yocto).
The FIDel binary can be started using the following command:
  ```sh   
fidel daemon   //Note: please use "sudo fidel daemon" if you want run your binaries system-wide 
```

You can find more information about building in [BUILD.md](https://github.com/YosiSF/FIDel/blob/master/docs/build.md) file included in documentation directory or visit our online documentation at http://edbdoc.org/.

## Testing FIDel

We have included unit tests for IPFS components that we modified, to run those tests execute the below command from source root directory (note that you need to install go test tool before running these commands):
  ```sh   
make check   //Note: please use "sudo make check" if you want run your binaries system-wide 
```

FIDel's metadata records the cluster topology configuration of the EinsteinDB cluster. FIDel schedules the VioletaBFT replica group Spanning Interval TSO guarantees for multi-layer persistence: The partial order ledger with atomic broadcast guarantees for Hot, Cold, or Warm Data. By attaching a timestamp that accounts for leap millisecond epochs we can set a retirement and archiving policy, archive indexed data, and configure index size cardinality check for universal binpacking. The rational for this is discussed in the TSO paper.

## FIDel REST APIs
FIDel exposes a set of [REST APIs](https://github.com/YosiSF/FIDel/tree/master/docs) to manipulate data and configure parameters on remote nodes. For example, you can use them to add files or fetch their content from another node without having direct access to it’s file system (which could be very useful when using IPFS as a CDN). You can also view disk usage statistics by calling specific endpoints that return values in JSON format. This shouldn't be confused with graphQL (a method for retrieving data similar to SQL queries but only against one endpoint).  
With FIDel you are also able to retrieve information about your cluster topology configuration and other status details such as network connection status between peers or pinning history of each block stored locally and remotely using commands similar those available in original IPFS implementation but with different parameters: 	
1.[ipfs-stat]() - print out the current stats from ipfs daemon  
2.[ipfs-block-stat]() - get stats for an individual block   
3.[ipfs-object-pinned]() - get the pinned objects from ipfs daemon    
4.[ipfs-pinned]() - get the blocks that are pinned either recursively or indirectly in ipfs
5.	[ipfs-config](https://github.com/YosiSF/FIDel/tree/master/docs) – display and modify config settings

## FIDel CLI Commands
We also created a new command to allow users to specify which other IPFS nodes will be used for storing data, we call this `fidel add`. This is similar to original `add` command but with different parameters. For example:
```sh 
./bin$ fidel add --remote <remote_storage_address> --type=all /home//testfile  //Note: please use "sudo fidel add" if you want run your binaries system-wide 
```

`--remote` parameter specifies which remote IPFS node will store our file (if not specified it’s assumed that current node is going to store all files locally). The second parameter after `–type` tells us what type of storage we want for given file (i.e. hot, warm or

Inspired by early IBM Tuplespace Ordering Hybrid Logical Timestamp oracle for pessimistic concurrent read-only streams: FIDel is EinsteinDB's Relativistic Non-Volatile Memory Causet Store daemon. FIDel services and applications run in a slice of the platform: the set of nodes on which the service receives a fraction of each node's resources, in the form of hypervizor on DAG with greedy automaton.

## Hierarchichal and Federated Namespaces

Time-series data is structured by tenant, application and namespace. Namespaces are the content addressable object names for Hot, Cold or Warm Data objects. The database storage manager uses these hierarchichal namespaces to guarantee consistency across all of the FIDel nodes in a cluster.
For example: `/fidel/edb/tenant1/appname1`    
In this case we have created new directory named `tenant1` under IPFS path `fidel->edb`, then inside it another one called “appname” that contains our files (for now only single file but more could be added if necessary). It is also possible to create multiple directories with different tenants (in this case directories would be named differently) and each tenant can have many applications running on it (each having their own set of data stored in folders under particular application name). This model should work fine in most cases where you need to store your data separately depending on what kind of information it represents but it’s also possible to group them together into same buckets if needed (i.e. there might not always be a need for multi-level hierarchy as in given example – instead you can put everything directly under edb directory).
FIDel's metadata records the cluster topology configuration of the EinsteinDB cluster. FIDel schedules the VioletaBFT replica group Spanning Interval TSO guarantees for multi-layer persistence: The partial order ledger with atomic broadcast guarantees for Hot, Cold, or Warm Data. By attaching a timestamp that accounts for leap millisecond epochs we can set a retirement and archiving policy, archive indexed data, and configure index size cardinality check for universal binpacking 
by a FIDel replica  are tagged in the kernel and subsequently classified to the EinsteinDB token bucket. The htb queuing discipline then provides each child token bucket with its configured rate, 
and fairly distributes the excess capacity from the root to the children that can use it in proportion to their rates. 
events in Minkowski spacetime.
long time.


## Federated Storage
FIDel supports both local and federated storage in a distributed fashion. FIDel clusters can be configured with either Hot, Cold or Warm Data. This allows for the efficient operations of Fixed-Record Append Only Logs (FOAL). The FOAL maintains a meta-data record of all objects stored across multiple nodes. Like IPFS, FIDel is content addressable which facilitates storing data in any number of locations within the cluster (or outside it). These namespaces are maintained using VioletaBFT’s Spanning Interval TSO guarantees.

## Retention Policy
Retention policy determines how long each block will be kept on particular node as well as when to move them onto remote storages if such need arises or only keep part of files locally and rest on another node(s) without having to wait for space running out before making this decision (so this way you can easily configure your application to retrieve most important files from local storage while everything else would be fetched from other nodes just like it happens in CDNs). We added some configuration parameters that allow us to specify retention policy:  
1.[--retain]() - sets expiration time for given file (in seconds)  
2.[--remote]() - sets the address of remote IPFS node where data will be stored when it’s time is expired locally  
3.[--type]() - defines what kind of storage we want for given file (i.e. hot, warm or other)

## Cluster Topology Configuration
FIDel's metadata records the cluster topology configuration of the EinsteinDB cluster. FIDel schedules VioletaBFT replica group Spanning Interval TSO guarantees for multi-layer persistence: The partial order ledger with atomic broadcast guarantees for Hot, Cold, or Warm Data. By attaching a timestamp that accounts for leap millisecond epochs we can set a retirement and archiving policy, archive indexed data, and configure index size cardinality check for universal binpacking. The rational for this is discussed in the TSO paper.

## References
1.[IPFS](https://github.com/ipfs): https://github.com/ipfs/go-ipfs

## References
1.[IPFS](https://github.com/ipfs): https://github.com/ipfs/go-ipfs/blob/master/docs/introduction.md  
2.[Yunqi Community](http://yqh.aliyun.com): http://yqh.aliyun.com/article/489  
3.[Spanning Interval TSO](http://www.vldb.org/pvldb/vol8/p1566-calder.pdf): http://www.vldb.org/pvldb/vol8/p1566-calder.pdf  
4.[VioletaBFT: Practical State Machine Replication for the Cloud Era](https://github.com). https://arxiv.org/abs180301954


## License
[Apache 2 license](https://github.com/YosiSF).
