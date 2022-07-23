# bench

## 構成

Scenario_xxx
  workerとしてgoroutineで動作できるもの
job_xxx
  Action+Validationのまとまり
  終わるまでブロックする
  Scenarioの中で利用する

## 想定負荷の流れ

- SaaS管理者は1人 (1 worker)
  - admin billingを順番に見る
    - 整合性チェック、テナントのworkerのbillingに左右されるのでちょっと難しそう
    - 何かしらのチェックは入れたいので頑張りどころ
      - 大会をfinishしたタイミングでそのtenantのbillingは確定するので、それを記憶するsync.Mapはどうだろう
      - ただしfinishしてから反映まで1秒の猶予があり、その1秒の間にどうなるかは不明
  - 見終わったらテナントを1増やして新規テナントシナリオを増やす

- 増えたテナントで大会が開催される(tenant worker)
  - 大会を追加
    - playerを追加
    - CSV入稿（下記の増加想定）
    - 大会のfinish
    - rankingの確認
  - tenant billingが返ってくる
  - 大会の追加に戻る

- 初期データテナントで整合性チェックをする
  - ここはあまり負荷は増えない
  - 巨大テナント(id=1) 1 worker
    - ranking, score
    - billing
  - 人気テナント(id=2~20?)(破壊的シナリオNG)
    - ranking, score
    - billing
  - のんびりテナント(id=21~40?)(破壊的シナリオOK)
    - playerを失格にして確認
    - billing

## シナリオ一覧

- SaaS管理者: AdminBilling
- 新規テナント: OrganizerNewTenant
- 既存巨大テナント(id=1): PopularTenant(heavry)
- 既存人気テナント: PopularTenant
- 既存のんびりテナント 破壊的操作OK: PeacefulTenant
- 管理者請求額確認: AdminBillingValidate
- テナント請求額確認: TenantBillingValidate

## CSV入稿について

benchから入稿されるCSVは、入稿される度に後ろに行数が増えていく
最後に入稿されたCSVが有効

# メモ

なんとかActionでリクエストを作って送って返ってきたresをValidateResponseで検証してるんだけど、この2つの関数に関係がないのでリクエスト開始から結果取得完了までの時刻(レスポンスタイム)を元になにかするのができない
n秒超えたらタイムアウトではないけど0点、みたいな調整がやりづらいのでそこだけ作りを変えたい気持ち
リクエストを送るctxをwrapして、そこでリクエスト送信時にメタデータを入れてvalidateにもctxを渡してそれをみれるようにするとか
