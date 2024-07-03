

### Doubts I have :
1) BTC address keeps on changing, How do I Track a particular BTC address ? | I am looking for something like Watch Only Wallets Supported on EVM.
2) If I actually have to get data on a BTC address that is having transfers, and I want to test it on my own, I need funds or a testing chain, How do I do that ? 



### Websites I interacted with :
1) https://bitcoincore.org/en/doc/27.0.0/





### Youtube Videos I Saw : 
1) https://www.youtube.com/watch?v=DW4_zDSufhQ
2) https://www.youtube.com/watch?v=IzBxxeGSkIo
3) https://www.youtube.com/watch?v=7JJkLW4SHKQ
4) https://www.youtube.com/watch?v=7JJkLW4SHKQ 





### Webistes I read from :
1) https://komodoplatform.com/en/academy/bitcoin-native-segwit-vs-taproot/



### Pointers I learnt :

1) What are sats :
Satoshis, often abbreviated to “sats”, are the atomic unit of Bitcoin, named after Bitcoin's creator, Satoshi Nakamoto.
A single Bitcoin represents 100 million satoshis (1 BTC = 100,000,000 sats)

2) What is Bitcoind :
Bitcoind is a (daemon) program that implements the Bitcoin protocol for remote procedure call (RPC Server) use.

3) tBTC : Testing BTC, used by developers and testers to check out their app.

4) Taproot address - P2TR : A pay-to-taproot (P2TR), also known as a Taproot or Bech32m address, is the most recent and advanced bitcoin address format. Taproot introduces more advanced security, privacy, flexibility and scaling to bitcoin. Like SegWit, Taproot addresses are opt-in and not currently widely supported. Taproot adoption can be tracked here.


## Executions I am now clear about : 
1) Clearing Doubt 2 : I will use the tBTC environment or Bitcoin testnet and use faucets to fund this btc wallet, and hence detect this funding with a script!
2) Clearing Doubt 1 : Since now I have a testnet btc environment I can create my own wallet that will have tBTC funds, This address will not be changed.


## Systems to use :
1) tBTC environment. | More clearly -> Bitcoin's latest testnet environment : Signet
2) bitcoind | bitcoin-cli, to spin up a testnet bitcoin node. 
3) Sparrow Wallet to manage my testnet wallet, Using same seed phrase in OKX to test my wallet.

## Going merry !
1) Proving that tBTC is a valid direction to proceed in :
  - make a test Wallet in Sparrow and fund it once with tBTC.
  - run the go script to get the transactions of this wallet, if the txn is confirmed then it should show here.
  - fund the wallet again with tBTC and that should show in the go script's rerun if the transaction is successful.
2) Getting list of Finality Provider :
  - use this : https://staking-api.testnet.babylonchain.io/swagger/index.html#/default/get_v1_finality_providers


## Caution : 
1) It is possible that the btc in testnet might be too less (0.0005 sBTC) to stake.
2) I need tBBN to pay for transactions.
3) Babylon Server is currently full, need tBTC and TBBN tokens, will need to ask for advice.

## External Issues :
1) Staking Window is closed.
2) Blocker: Caution pt 2: Discord server is completely filled, resisting me to get tokens easily

