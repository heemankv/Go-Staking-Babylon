# Here is a step wise break down on how to proceed with the code.

## Part 0 : Initial Setup 

### Set up a Bitcoin node
- Use [btc-staker readme](https://github.com/babylonchain/btc-staker?tab=readme-ov-file#2-setting-up-a-bitcoin-node) to setup a node properly
- Local Reference Command (to be run in ~/Work/assignments/bitquid/bitcoin-26.0 ) \
  ``` ./bin/bitcoind -deprecatedrpc=create_bdb,listaccounts -signet -server -txindex -rpcport=38332 -rpcuser=dexterhv -rpcpassword=verma ```



### Command
./bin/bitcoind -deprecatedrpc=accounts -signet -server -txindex -rpcport=38332 -rpcuser=dexterhv -rpcpassword=verma -rpcallowip=0.0.0.0/0



## Part 1 : Bitcoin Deposit Monitoring

### Versioning : 
- Bitcoin Core 26.0.0

### Docs : 
- [Bitcoin Core 26.0.0](https://bitcoincore.org/en/doc/26.0.0)

### Assumptions 
- bitcoind server is running and it's localhost link is known.
- Account is created by the bitcoind and is accessible.
- RPC User, Password and URL are known.

### Input Arguments
- User Wallet Label to choose (if any) 


### Implementation Steps
1) Create an RPC client.
2) `listwallets` : Check available wallets, it should return the wallets that were setup.
3) `getaddressesbylabel` : Get the addresses of the given label, by default use the first address.
4) in Loop : break if balance > 0.0005 ([Staking Limits](https://medium.com/babylonchain-io/babylon-bitcoin-staking-testnet-4-launch-3c7fe3979827))
  - `getbalances` : Use this to get the trusted, untrusted_pending balance of the user.
5) Implement Staking logic after the loop.


### Future Scopes (Extendibility) :
- Using `getblockchaininfo` verify that `blockchaind` server is connected to `signet`.
- After pt 2, user `getaddressesbylabel` to choose from addresses to tract.

### Scoped Out :
- Possibility of using xPubKey to track the entire wallet, rather we are only tracing one of the reciever address of the wallet.

---

# Important !!!
- babbylonchain's testnet-4 had 3 caps, all of which are not over, i.e no testnet.
  - Staking not possible using cli / [dapp](https://btcstaking.testnet.babylonchain.io/).
- babylonchains's discord server is full! and since testnet is over there is no proper access to sBTC & tBBN tokens.
- babylonchain has not provided network information of [bbn-test-4](https://github.com/babylonchain/networks/tree/main/bbn-test-4) like they have for [bbn-test-3](https://github.com/babylonchain/networks/tree/main/bbn-test-3)
- even though on the [official docs](https://docs.babylonchain.io/docs/user-guides/btc-staking-testnet/become-btc-staker) btc-staker is given under testnet-3, this [staking-backend.md](https://github.com/babylonchain/networks/blob/main/bbn-test-4/integration/staking-backend.md) ensures that btc-staker can still be used for staking on testnet-4.
- unlike RPC documentation available for [bitcoind](https://bitcoincore.org/en/doc/26.0.0/), there is no RPC information availabe for babylonchain's [stakerd](https://github.com/babylonchain/btc-staker).
- given lack of network information and sunset of testnet-4, the transactions are likely to fail.


---


## Part 2 : Detection of Finality Provider & Staking Transaction :

### Versioning : 
- btc-staker v0.3.0

### Docs : 
- [btc-staker v0.3.0](https://docs.babylonchain.io/docs/user-guides/btc-staking-testnet/become-btc-staker#5-staking-operations-with-stakercli)
- [staking-api-service](https://github.com/babylonchain/staking-api-service/blob/dev/docs/swagger.json)

### Assumptions 
- stakerd server is running and it's localhost link is known.
- babylonchaind is used to created keychain.
- stakerd config file (~./stakerd/stakerd.conf) is configured correctly.
- sBTC account containing sBTC has more than 0.0005 sBTC.
- babylonchain account containing tBBN has enough tokens for transaction.

### Implementation Steps 
#### Detection of Finality Provider
1) make `Post` call to `<baseURL>/babylon_finality_providers` to get the list of finality providers available.
  - it is expected that this call will fail / not give expected output (see Important-6).
2) make `Get` call to [staking-api.testnet.babylon](https://staking-api.testnet.babylonchain.io/v1/finality-providers) to get list of finality providers.
3) choose any finality provider at random from the response and extract it's `btc_pk`.

#### Staking Transansaction 
##### Way-1 - Using stakerd
1) make a `Post` call to `<baseURL>/stake` with params : 
  - stakerAddress : user address, selected in Part 1, pt 3.
  - stakingAmount : in btc, threshold > 0.0005.
  - fpBtcPks : finality Provider's btc_pk extracted above.
  - stakingTimeBlocks.
  - it is likely that the txn will fail. (see Important-6).
2) Parse any response and show to user.

##### Way-2 - Using bitcoind

  1) create `staker_key` in the bitcoind wallet
  2) create unfunded and not signed staking transaction using the `BuildV0IdentifiableStakingOutputsAndTx` function 
  3) serialize the unfunded and not signed staking transaction to `staking_transaction_hex`
  4) using the go client, call `fundrawtransaction` "staking_transaction_hex" to retrieve `funded_staking_transaction_hex`. The bitcoind wallet will automatically choose unspent outputs to fund this transaction.
  5) using the go client, call `signrawtransactionwithwallet` "funded_staking_transaction_hex". This call will sign all inputs of the transaction and return `signed_staking_transaction_hex`.
  6) using the go client, call `sendrawtransaction` "signed_staking_transaction_hex"


###### Docs needed : 
1) [Staking Transaction Via Bitcoind](https://github.com/babylonchain/babylon/blob/add420f074751cf53edea5b7a55cca3d34291f5b/docs/transaction-impl-spec.md#observable-staking-transactions)
2) [Staking Parameters](https://github.com/babylonchain/networks/tree/main/bbn-test-4/parameters)
3) [BuildV0IdentifiableStakingOutputsAndTx](https://github.com/babylonchain/babylon/blob/add420f074751cf53edea5b7a55cca3d34291f5b/btcstaking/identifiable_staking.go?plain=1#L231)




### Future Scopes (Extendibility) :
- With the right stakerd config, there should not be any need for staker-api.
- With the right stakerd config, the stake transaction should work.


---

## TODO list :
1) `BuildV0IdentifiableStakingOutputsAndTx` : Implementation Steps - Staking Transaction - Way 2
2) `stakingTimeBlocks` : Implementation Steps - Staking Transaction - Way 1