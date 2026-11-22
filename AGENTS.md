# GoBitfinex (bitfinex-api-go)

## 專案是什麼

Bitfinex REST 與 WebSocket Go API SDK（module `github.com/wanth1997/GoBitfinex`，remote `wanth1997/GoBitfinex`），同時提供 `v2/rest`、`v2/websocket` 與向後相容的 legacy `v1` 套件（`go 1.21`，`toolchain go1.26.10`）。

## 架構速覽

- `v2/rest/`：REST v2 client（23 個 service，掛載於 `rest.Client`）。
- `v2/websocket/`：WebSocket v2 client（連線管理、心跳、自動重連、orderbook checksum、頻道訂閱）。
- `pkg/models/`：27 個共用資料模型子套件；`pkg/mux/`：連線多工器；`pkg/convert/`、`pkg/utils/`：型別轉換與 Nonce 產生器。
- `v1/`：Legacy REST + WebSocket client（凍結不再新增功能；所有公開 REST service 方法皆帶 `V1` 前綴，如 `V1Get`、`V1All`）。
- `tests/integration/{v1,v2}`：整合與 mock WS 測試；`examples/`：使用範例。

## 驗證命令

```bash
go build ./...
go vet ./v2/... ./pkg/...
go test ./v1/... ./v2/... ./pkg/...
```

（`tests/integration/` 內含連網 live 測試，未經確認勿對真實 API 憑證執行；legacy `v1/websocket.go:181` 既有 `unreachable code` vet 警示。）

## 專案特有規則與禁忌

- Module 路徑固定為 `github.com/wanth1997/GoBitfinex`（本機資料夾名為 `bitfinex-api-go`，import 不可混用舊名）。
- `v1/` 僅維持相容修復，公開 REST service 方法一律保留 `V1` 前綴；新功能只加在 `v2/`。
- 新增 `v2/rest` endpoint 時，優先在 `pkg/models/` 補齊對應型別與單元測試，避免擴大裸 `[]interface{}` 回傳面。

## 相關文件

- `docs/project-brief.md` — 專案全貌與待辦
- `docs/work-log.md` — 工作記錄（最新在最上方）
- `docs/rest_v2.md`、`docs/ws_v2.md`、`docs/v1.md` — API 參考索引（完整型別與註解以 `go doc` 為準）
