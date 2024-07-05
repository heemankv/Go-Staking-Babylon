# Research

#### [Bookmarks](utils/Bookmarks.md) contains a summary of all the websites I visited.

## Pretext :

I am an EVM based developer, majorly all my knowledge about Blockchain has been through the Ethereum and it's L2's lens, After which I started reading about on Cosmos and Modular Blockchains, Understanding % Developing on top of bitcoin was surely a mind opening task for me, where in I learnt a lot of concepts and gain deeper insights on concepts I had already know, I think this project was a good way for me to catch up with bitcoin (long over due) and and get deeper into Defi. Thank you.

## Important !!!

- babbylonchain's testnet-4 had 3 caps, all of which are not over, i.e no testnet.
  - Staking not possible using cli / [dapp](https://btcstaking.testnet.babylonchain.io/).
- babylonchains's discord server is full! and since testnet is over there is no proper access to sBTC & tBBN tokens.
- babylonchain has not provided network information of [bbn-test-4](https://github.com/babylonchain/networks/tree/main/bbn-test-4) like they have for [bbn-test-3](https://github.com/babylonchain/networks/tree/main/bbn-test-3)
- even though on the [official docs](https://docs.babylonchain.io/docs/user-guides/btc-staking-testnet/become-btc-staker) btc-staker is given under testnet-3, this [staking-backend.md](https://github.com/babylonchain/networks/blob/main/bbn-test-4/integration/staking-backend.md) ensures that btc-staker can still be used for staking on testnet-4.
- unlike RPC documentation available for [bitcoind](https://bitcoincore.org/en/doc/26.0.0/), there is no RPC information availabe for babylonchain's [stakerd](https://github.com/babylonchain/btc-staker).
- given lack of network information and sunset of testnet-4, the transactions are likely to fail.

---

## Information Dump :

##### This is an information dump that I collected over the course of 3 days I worked in this Project. These might not be Industry Standard definition but these helped me to get past the hurdles I faced.

### Executive knowledge (My Observations)

1. **bbn-test-3 vs bbn-test-4** :
    - testnet 3 did not have a transaction indexer and so in testnet 3 users had to execute transaction in two chains, `signet bitcoin` and `babylon-testnet-3`, users had to stake their sBTC on to a staking script provide on signet by babylonchain and then use it's txnOutput and self attest that on babylon for indexing and validation of state.
    - This [medium article](https://medium.com/babylonchain-io/bitcoin-staking-guide-for-babylon-testnet-7c0fe4fffa95) on Testnet 3 proves my point of 2 chain interaction on Testnet 3 and this [Video on Babylon Testnet 4 ](https://www.youtube.com/watch?v=hg5u2PVl9nw) along with [this system design](https://github.com/babylonchain/networks/blob/main/bbn-test-4/integration/assets/system-detailed.png) proves my point of a single interaction in Testnet 4.


2. **Ideally the Staking should have gone through** :
    - I believe that the staking should not fail, rather it's acceptance by finality providers could fail, staking deals with majorly putting sBTC on bitcoin staking script and atleast this transaction should have gone through.

3. **staking-api service** :
    - api service that is provided to the babylon dashboard, this api provides config data to the frontend, can only use to get finality providers.

4. **btcd RpcClient is outdated** : doesn't have lots of functions of v26.0.0 \
    - still has functions of v0.17.0 (listAccounts- removed in v0.18.0).
    - had to use RawRequest for major of the calls I made.

5. **stakerService from btc-staker is not buildable** : this led me to use abstract RPC calls. \
    - could not use the RPC client since it's import was not buildable.
    - see branch `failure/build-stakerd-rpc-client`.

6. **staking with bitcoind** : btc-staker doesn't provide enough public functions
    - to do create raw staking transaction bitcoind has all the needed functions usable, but `btc-staker` doesn't provide enough functions.
    - see branch `failure/staking-with-bitcoind`.



### Context (Only Writing what I didn't know earlier)

0. **Staking on Babylon** : How it works \

- Staking Transaction on Signet : locks the Bitcoin to stake in a Bitcoin staking script.
- Submit Pre-Signed Staking Covenant and On-demand Unbonding Transaction for unlocking bitcoin before time.
- Pre-Signed Unbonding covenant transaction submission.
- Linking the signet and babylon testnet wallets together for rewards.
- [Read about covenant here](https://github.com/babylonchain/covenant-emulator/)

1. **Sats** : Like GWEI but for bitcoin. \
   Satoshis, often abbreviated to “sats”, are the atomic unit of Bitcoin, named after Bitcoin's creator, Satoshi Nakamoto.
   A single Bitcoin represents 100 million satoshis (1 BTC = 100,000,000 sats)

2. **Bitcoind** : Like mantraD, babylonD, stakerD, etc. \
   Bitcoind is a (daemon) program that implements the Bitcoin protocol for remote procedure call (RPC Server) use.

3. **Bitcoin-cli** : Interaction layer on top of `bitcoind`.
   Allows to do read and write transactions on top of `bitcoind`.

4. **Babylond** : cli for babylon chain.
   After setting up `bitcoind` we have to set `babylond` for key generation to interact with `stakerd`.

5. **BTC-staker** : [Github Repo](https://github.com/babylonchain/btc-staker)
   This is the staking script that babylonchain provides for staking, it has `stakerd` and `staker-cli`, needs to be configured `stakerd.config` available in `utils` folder.

6. **Finality Providers** :
   The Finality Provider Daemon is responsible for monitoring for new Babylon blocks, committing public randomness for the blocks it intends to provide finality signatures for, and submitting finality signatures.

7. **stakerd and bitcoind has access to pvtKey** : EVM : just like hardat \
   they already have access to account's pvt key, so they can internally get transactions signed.

8. **UTXO** : Unspent Transaction Output \
   [Read here](https://www.investopedia.com/terms/u/utxo.asp)

9. **Signet & sBTC** : Signet is latest testnet for bitcoin. \
   Like Test-Ether is used in testnets similarly sBTC is used to develop and test apps on top of bitcoin.

10. **Taproot address** : Still not very clear.
    P2TR : A pay-to-taproot (P2TR), also known as a Taproot or Bech32m address, is the most recent and advanced bitcoin address format. Taproot introduces more advanced security, privacy, flexibility and scaling to bitcoin.

11. **Timelock Script** : like a smart contract. | A spending condition placed on an output.
    A timelock allows a Bitcoin transaction to be created such that the recipient of the outputs cannot spend them for a specified time

### Doubts I had :

1. BTC address keeps on changing, How do I track a particular BTC address ? | I am looking for something like Watch Only Wallets Supported on EVM.

## External Issues :

1. Staking Window is closed.
2. Blocker: Caution pt 2: Discord server is completely filled, resisting me to get tokens easily.


## To Understand :

- [TapRoot Upgrades](https://github.com/bitcoin/bips/blob/master/bip-0341.mediawiki)
- [Schnorr signatures](https://github.com/bitcoin/bips/blob/master/bip-0340.mediawiki)
---

# Here is a step wise break down on how to proceed with the code.

## Part 0 : Initial Setup

### Set up a Bitcoin node

- Use [btc-staker readme](https://github.com/babylonchain/btc-staker?tab=readme-ov-file#2-setting-up-a-bitcoin-node) to setup a node properly
- Local Reference Command (to be run in ~/Work/assignments/bitquid/bitcoin-26.0 ) \
  `./bin/bitcoind -deprecatedrpc=create_bdb,listaccounts -signet -server -txindex -rpcport=38332 -rpcuser=dexterhv -rpcpassword=verma`

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

1. Create an RPC client.
2. `listwallets` : Check available wallets, it should return the wallets that were setup.
3. `getaddressesbylabel` : Get the addresses of the given label, by default use the first address.
4. in Loop : break if balance > 0.0005 ([Staking Limits](https://medium.com/babylonchain-io/babylon-bitcoin-staking-testnet-4-launch-3c7fe3979827))

- `getbalances` : Use this to get the trusted, untrusted_pending balance of the user.

5. Implement Staking logic after the loop.

### Future Scopes (Extendibility) :

- Using `getblockchaininfo` verify that `blockchaind` server is connected to `signet`.
- After pt 2, user `getaddressesbylabel` to choose from addresses to tract.

### Scoped Out :

- Possibility of using xPubKey to track the entire wallet, rather we are only tracing one of the reciever address of the wallet.

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

1. make `Post` call to `<baseURL>/babylon_finality_providers` to get the list of finality providers available.

- it is expected that this call will fail / not give expected output (see Important-6).

2. make `Get` call to [staking-api.testnet.babylon](https://staking-api.testnet.babylonchain.io/v1/finality-providers) to get list of finality providers.
3. choose any finality provider at random from the response and extract it's `btc_pk`.

#### Staking Transansaction

##### Way-1 - Using stakerd

1. make a `Post` call to `<baseURL>/stake` with params :

- stakerAddress : user address, selected in Part 1, pt 3.
- stakingAmount : in btc, threshold > 0.0005.
- fpBtcPks : finality Provider's btc_pk extracted above.
- stakingTimeBlocks.
- it is likely that the txn will fail. (see Important-6).

2. Parse any response and show to user.

##### Way-2 - Using bitcoind

1. create `staker_key` in the bitcoind wallet
2. create unfunded and not signed staking transaction using the `BuildV0IdentifiableStakingOutputsAndTx` function
3. serialize the unfunded and not signed staking transaction to `staking_transaction_hex`
4. using the go client, call `fundrawtransaction` "staking_transaction_hex" to retrieve `funded_staking_transaction_hex`. The bitcoind wallet will automatically choose unspent outputs to fund this transaction.
5. using the go client, call `signrawtransactionwithwallet` "funded_staking_transaction_hex". This call will sign all inputs of the transaction and return `signed_staking_transaction_hex`.
6. using the go client, call `sendrawtransaction` "signed_staking_transaction_hex"

###### Docs needed :

1. [Staking Transaction Via Bitcoind](https://github.com/babylonchain/babylon/blob/add420f074751cf53edea5b7a55cca3d34291f5b/docs/transaction-impl-spec.md#observable-staking-transactions)
2. [Staking Parameters](https://github.com/babylonchain/networks/tree/main/bbn-test-4/parameters)
3. [BuildV0IdentifiableStakingOutputsAndTx](https://github.com/babylonchain/babylon/blob/add420f074751cf53edea5b7a55cca3d34291f5b/btcstaking/identifiable_staking.go?plain=1#L231)

### Future Scopes (Extendibility) :

- With the right stakerd config, there should not be any need for staker-api.
- With the right stakerd config, the stake transaction should work.

---

## TODO list :

1. `BuildV0IdentifiableStakingOutputsAndTx` : Implementation Steps - Staking Transaction - Way 2
2. `stakingTimeBlocks` : Implementation Steps - Staking Transaction - Way 1

### Youtube Videos I saw :

1. [Bitcoin Basics](https://www.youtube.com/watch?v=DW4_zDSufhQ)
2. [Bitcoin Address Change Explained](https://www.youtube.com/watch?v=IzBxxeGSkIo)
3. [Sparrow Wallet](https://www.youtube.com/watch?v=7JJkLW4SHKQ)
4. [Hashed Time Locked Contracts](https://youtu.be/9PSjgELvQTo?t=84)
5. [Bitcoind on Mac](https://www.youtube.com/watch?v=hMo4QeHVXiI)
6. [Babylon Testnet 4 UI explained](https://www.youtube.com/watch?v=hg5u2PVl9nw)
