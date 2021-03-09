# API概要
* 認証
    * `POST /signup`: ユーザーを新規登録
    * `POST /login`: ログイン処理、JWTトークンを取得
* API(JWTによる認証が必要)
    * `GET /api/user`: ログイン中のユーザーに関する情報を取得
    ```
    $ curl -X GET -H "Authorization: Bearer (token)" localhost:8080/api/user
    {   
        "ID":1,
        "CreatedAt":"2021-03-09T09:36:55.75Z",
        "UpdatedAt":"2021-03-09T09:36:55.787Z",
        "DeletedAt":null,
        "name":"ドラえもん",
        "password":"",
        "email":"dora22@gmail.com",
        "Friends":  [
            {
                "ID":2,
                "CreatedAt":
                "2021-03-09T09:36:55.762Z",
                "UpdatedAt":"2021-03-09T09:36:55.795Z",
                "DeletedAt":null,"name":"野比のび太",
                "password":"",
                "email":"",
                "Friends":null,
                "ManagingGroups":null,
                "PerticipatingGroups":null
            },{
                "ID":6,
                "CreatedAt":"2021-03-09T09:36:55.78Z",
                "UpdatedAt":"2021-03-09T09:36:55.817Z",
                "DeletedAt":null,
                "name":"野比のび助",
                "password":"",
                "email":"",
                "Friends":null,
                "ManagingGroups":null,
                "PerticipatingGroups":null
            },{
                "ID":7,
                "CreatedAt":"2021-03-09T09:36:55.784Z",
                "UpdatedAt":"2021-03-09T09:36:55.784Z",
                "DeletedAt":null,
                "name":"野比玉子",
                "password":"",
                "email":"",
                "Friends":null,
                "ManagingGroups":null,
                "PerticipatingGroups":null
            }
        ],
        "ManagingGroups":[],
        "PerticipatingGroups":  [
            {
                "ID":1,
                "CreatedAt":"2021-03-09T09:36:55.825Z",
                "UpdatedAt":"2021-03-09T09:36:55.825Z",
                "DeletedAt":null,
                "name":"野比家",
                "Managers":null,
                "Members":null
            },{
                "ID":2,
                "CreatedAt":"2021-03-09T09:36:55.854Z",
                "UpdatedAt":"2021-03-09T09:36:55.854Z",
                "DeletedAt":null,
                "name":"空き地に集まる会",
                "Managers":null,
                "Members":null
            }
        ]
    }
    ```
    * `GET /api/group/:group_id`: 指定のグループに関する情報を取得
    ```
    $ curl -X GET -H "Authorization: Bearer (token)" localhost:8080/api/group/1
    {
        "ID":1,
        "CreatedAt":"2021-03-09T09:36:55.825Z",
        "UpdatedAt":"2021-03-09T09:36:55.825Z",
        "DeletedAt":null,
        "name":"野比家",
        "Managers": [
            {
                "ID":6,
                "CreatedAt":"2021-03-09T09:36:55.78Z",
                "UpdatedAt":"2021-03-09T09:36:55.817Z",
                "DeletedAt":null,
                "name":"野比のび助",
                "password":"",
                "email":"",
                "Friends":null,
                "ManagingGroups":null,
                "PerticipatingGroups":null
            },{
                "ID":7,
                "CreatedAt":"2021-03-09T09:36:55.784Z",
                "UpdatedAt":"2021-03-09T09:36:55.784Z",
                "DeletedAt":null,
                "name":"野比玉子",
                "password":"",
                "email":"",
                "Friends":null,
                "ManagingGroups":null,
                "PerticipatingGroups":null
            }
        ],
        "Members":  [
            {
                "ID":1,
                "CreatedAt":"2021-03-09T09:36:55.75Z",
                "UpdatedAt":"2021-03-09T09:36:55.787Z",
                "DeletedAt":null,
                "name":"ドラえもん",
                "password":"",
                "email":"",
                "Friends":null,
                "ManagingGroups":null,
                "PerticipatingGroups":null
            },{
                "ID":2,
                "CreatedAt":"2021-03-09T09:36:55.762Z",
                "UpdatedAt":"2021-03-09T09:36:55.795Z",
                "DeletedAt":null,
                "name":"野比のび太",
                "password":"",
                "email":"",
                "Friends":null,
                "ManagingGroups":null,
                "PerticipatingGroups":null
            }
        ]
    }
    ```
    * `GET /api/directmessage/:user_id`: 指定のユーザーとのダイレクトメッセージを受信
    * `POST /api/directmessage/:user_id`: ダイレクトメッセージを送信
    * `GET /api/groupmessage/:group_id`: 指定のグループのグループメッセージを受信
    * `POST /api/groupmessage/:group_id`: 指定のグループのグループメッセージを送信
    * `GET /api/healthrecord`: 自分の健康記録を取得
    * `POST /api/healthrecord`: 自分の健康記録を送信