# API概要
* 認証
    * `POST /signup`: ユーザーを新規登録
    * `POST /login`: ログイン処理、JWTトークンを取得
* API(JWTによる認証が必要)
    * `GET /api/user`: ログイン中のユーザーに関する情報を取得
    * `GET /api/group/:group_id`: 指定のグループに関する情報を取得
    * `GET /api/directmessage/:user_id`: 指定のユーザーとのダイレクトメッセージを受信
    * `POST /api/directmessage/:user_id`: ダイレクトメッセージを送信
    * `GET /api/groupmessage/:group_id`: 指定のグループのグループメッセージを受信
    * `POST /api/groupmessage/:group_id`: 指定のグループのグループメッセージを送信
    * `GET /api/healthrecord`: 自分の健康記録を取得
    * `POST /api/healthrecord`: 自分の健康記録を送信