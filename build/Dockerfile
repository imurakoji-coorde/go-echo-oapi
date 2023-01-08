# goバージョン
FROM golang:1.19.1
# ワーキングディレクトリの設定
WORKDIR /go/src/
# ホストのファイルをコンテナの作業ディレクトリに移行
#ADD . /go/src/

EXPOSE 8080

RUN go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.11.0
RUN go install github.com/cosmtrek/air@latest
CMD ["air"]