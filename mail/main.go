package main

import (
	"fmt"
	"io"
	"net/mail"
	"strings"
)

var mailMsg string

func main() {
	// strings.NewReader 関数を使用して、msg 文字列から新しい Reader オブジェクトを作成
	// Reader オブジェクトは、io.Reader インターフェースを実装しており、メールを読み込むために使用される
	r := strings.NewReader(mailMsg)

	// mail.ReadMessage 関数を使用して、メールを読み込み、メールオブジェクトを返す
	m, _ := mail.ReadMessage(r)

	// Message 構造体の Header フィールドを header 変数に代入。これにより、メールのヘッダー情報にアクセスできる
	header := m.Header

	// Header の Fromを出力
	fmt.Println("From:", header.Get("From"))

	body, _ := io.ReadAll(m.Body)
	// Body を出力
	fmt.Printf("%s", body)
}

Delivered-To: このヘッダーは、メールが最終的に配信された受信者のメールアドレスを示します。

Received: このヘッダーは、メールが経由した各サーバーによって追加され、メールの転送経路を示します。複数のReceivedヘッダーが存在することがあります。

ARC-Seal, ARC-Message-Signature, ARC-Authentication-Results: これらはAuthenticated Received Chain (ARC)に関連するヘッダーで、メールが転送される過程での認証結果を保持します。

Return-Path: このヘッダーは、バウンスメッセージ（配信不能通知）が送られる際の返信先アドレスを示します。

Received-SPF: Sender Policy Framework (SPF)の検証結果を示し、メールが正当な送信元から送られたかどうかを確認します。

Authentication-Results: メールが認証プロセス（DKIM、SPF、DMARCなど）を通過したかどうかの結果を示します。

DKIM-Signature: DomainKeys Identified Mail (DKIM)の署名で、メールの内容が送信者によって署名され、途中で改ざんされていないことを保証します。

X-Google-DKIM-Signature: Googleが追加したDKIM署名で、Gmailがメールを処理したことを示します。

X-Gm-Message-State: Gmailによって追加されたヘッダーで、メールの状態に関する情報を含みます。

X-Google-Smtp-Source: Gmailによって追加されたヘッダーで、メールの送信に関する追加情報を含みます。

MIME-Version: MIME (Multipurpose Internet Mail Extensions)のバージョンを示し、メールがテキスト以外のコンテンツを含むことを示します。

From: メールの送信者の名前とメールアドレスを示します。

Date: メールが送信された日時を示します。

Message-ID: メールに一意に割り当てられたIDで、メールの追跡に使用されます。

Subject: メールの件名を示します。

To: メールの宛先アドレスを示します。

Content-Type: メールの本文や添付ファイルの形式（MIMEタイプ）を示します。

Content-Disposition: 添付ファイルの表示方法やファイル名を示します。

Content-Transfer-Encoding: メールの本文や添付ファイルがどのようにエンコードされているかを示します。

Content-ID: 添付ファイルに関連付けられた一意のIDで、メール内で参照する際に使用されます。

X-Attachment-Id: Gmailが添付ファイルに割り当てたIDで、メール内での添付ファイルの参照に使用されます。