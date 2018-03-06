# 前準備

    go get -u google.golang.org/grpc
    コンパイラのインストール
    https://github.com/google/protobuf/releases
    protoc-3.5.1-osx-x86_64.zip
    Binフォルダ内のprotocをpathが通ったフォルダへ放り込む
    Protocのgoプラグインをインストール
    go get -u github.com/golang/protobuf/protoc-gen-go
    ProtocのDoc生成プラグインをインストール
    go get -u github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc
