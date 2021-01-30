# SecretDB master

## 概要

Cosmos SDK を利用して秘匿化ブロックチェーンデータベースを作成する．

- コンセンサスアルゴリズム : Tendermint
- 秘匿化 : Intel SGX + Graphene
- データベース : MongoDB + MapReduce

***秘匿化は未実装***

## 構成

```bash
SecretDB-master/
├── app
│   ├── app.go
│   ├── export.go
│   └── prefix.go
├── cmd : コマンド関連
│   ├── secretdbcli
│   │   └── main.go
│   └── secretdbd
│       ├── genaccounts.go
│       └── main.go
├── config.yml
├── go.mod
├── init.sh : 初期化用スクリプト
└── x
 └── secretdb : 本体
    ├── abci.go
    ├── client
    │   ├── cli : secretcbcli
    │   │   ├── query.go
    │   │   ├── queryBlockHash.go
    │   │   ├── queryItem.go
    │   │   ├── queryOperatorPubkey.go
    │   │   ├── tx.go
    │   │   ├── txBlockHash.go
    │   │   ├── txItem.go
    │   │   └── utils.go
    │   └── rest : REST
    │       ├── queryBlockHash.go
    │       ├── queryItem.go
    │       ├── queryOperatorPubkey.go
    │       ├── rest.go
    │       ├── txBlockHash.go
    │       └── txItem.go
    ├── genesis.go
    ├── handler.go : メッセージハンドラ
    ├── handlerMsgCreateBlockHash.go
    ├── handlerMsgDeleteBlockHash.go
    ├── handlerMsgDeleteItem.go
    ├── handlerMsgDeleteItems.go
    ├── handlerMsgSetBlockHash.go
    ├── handlerMsgStoreItem.go
    ├── handlerMsgUpdateItem.go
    ├── handlerMsgUpdateItems.go
    ├── keeper : データベースとのやり取りを管理
    │   ├── blockHash.go
    │   ├── item.go
    │   ├── keeper.go
    │   ├── operatorPubkey.go
    │   ├── params.go
    │   ├── querier.go
    │   └── utils.go
    ├── module.go
    ├── spec
    │   └── README.md
    └── types : 構造体等の定義
        ├── codec.go
        ├── errors.go
        ├── events.go
        ├── expected_keepers.go
        ├── genesis.go
        ├── key.go
        ├── msg.go
        ├── msgCreateBlockHash.go
        ├── msgDeleteBlockHash.go
        ├── msgDeleteItem.go
        ├── msgDeleteItems.go
        ├── msgSetBlockHash.go
        ├── msgStoreItem.go
        ├── msgUpdateItem.go
        ├── msgUpdateItems.go
        ├── params.go
        ├── querier.go
        ├── typeBlockHash.go
        ├── typeItem.go
        └── types.go

```

## インストール

```bash
go install ./cmd/secretdbd
go install ./cmd/secretdbcli
```

## 設定

```$HOME/.secretdbd/config/app.toml```に以下を設定

```bash
validator-name = "validator"
kayring-backend = "test"
keyring-password = "password of your keyring"
child-count = 1
child-uri = ["tcp://..."]
child-chainid = ["test-child1"]
```

## 起動方法

```bash
# リセット
secretdbd unsafe-reset-all

# genesis.jsonの作成
secretdbd init mynode --chain-id test-master

# CLIの設定
secretdbcli config keyring-backend test
secretdbcli config chain-id test-master
secretdbcli config output json
secretdbcli config indent true
secretdbcli config trust-node true

# 鍵の生成
secretdbcli keys add operator
secretdbcli keys add validator
secretdbcli keys add user1
secretdbcli keys add user2

# genesis.jsonにアカウントを追加
secretdbd add-genesis-account $(secretdbcli keys show operator -a) 1000token
secretdbd add-genesis-account $(secretdbcli keys show validator -a) 1000token,100000000stake
secretdbd add-genesis-account $(secretdbcli keys show user1 -a) 1000token
secretdbd add-genesis-account $(secretdbcli keys show user2 -a) 1000token

# genesis.jsonに初期トランザクションを追加
secretdbd gentx --name validator --keyring-backend test-master
secretdbd collect-gentxs

# 起動
secretdbd start --rpc.laddr "tcp://0.0.0.0:26657"
```
