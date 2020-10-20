# docbase-cli
## Overview
cliでDocBaseの記事を検索

## Install
### Mac
1. pecoのインストール
※ pecoインストールされている場合は実施不要
```sh
brew install peco
```
2. docbase-cliのインストール
```sh
version=v0.1.3
wget https://github.com/yasuyuki0321/docbase-cli/releases/download/$version/docbase-cli_darwin_x86_64.tar.gz
```
3. ファイルの解凍
```sh
tar zxvf docbase-cli_darwin_x86_64.tar.gz
```
4. ファイルをpathの通っているディレクトリに配置する

## Usage
- ↓に従い、DocBaseへのアクセストークンを作成する
[DocBase APIドキュメント \- DocBase](https://help.docbase.io/posts/45703)
- docbase-cliにアクセストークン等の情報を登録する
```sh
docbase init
```
登録する情報
  - アクセストークン
  - ユーザID
  - ドメイン
  ※ DocBaseにアクセスする際のURLのxxxの文字列(xxx.docbase.io)

## Note
- docbase open: docbaseのトップページの表示
- docbase init: ツールを使用するために必要な情報の設定
- docbase list: 自分が公開しているメモの一覧を表示
- docbase list -s private: 公開していないメモの一覧を表示
- docbase list -a アカウント名: アカウント名の公開しているメモの一覧を表示
- これまでに取得してたGOOD JOBS COUNT(グッジョブ!) とSTARS COUNT(スター)の合計数が確認できる
- TITLEを選択してEnterを押すことで、ブラウザに記事を表示することができる
- 一覧上部の '>' の部分に文字を入力することで一覧にフィルタを掛けることができる
