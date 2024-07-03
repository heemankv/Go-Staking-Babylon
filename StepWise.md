# Here is a step wise break down on how to proceed with the code.

## Part 0 : Initial Setup 

### Set up a Bitcoin node
- Use [btc-staker readme](https://github.com/babylonchain/btc-staker?tab=readme-ov-file#2-setting-up-a-bitcoin-node) to setup a node properly
- Local Reference Command (to be run in ~/Work/assignments/bitquid/bitcoin-26.0 ) \
  ``` ./bin/bitcoind -deprecatedrpc=create_bdb,listaccounts -signet -server -txindex -rpcport=38332 -rpcuser=dexterhv -rpcpassword=verma ```



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

---


## Part 2 : Detection of Finality Provider & Staking Transaction :

### Versioning : 
- btc-staker v0.3.0

### Docs : 
- [btc-staker v0.3.0](https://docs.babylonchain.io/docs/user-guides/btc-staking-testnet/become-btc-staker#5-staking-operations-with-stakercli)

### Assumptions 
- stakerd server is running and it's localhost link is known.
- babylonchaind is used to created keychain.
- stakerd config file (~./stakerd/stakerd.conf) is configured correctly.

### Implementation Steps
1) Validate 
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