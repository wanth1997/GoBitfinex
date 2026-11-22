# GoBitfinex

## 目標
Bitfinex 交易所的官方 Go 語言 API 客戶端，提供 REST 與 WebSocket 兩種介面，支援 v1 (legacy) 與 v2 版本。

## 架構 / 技術棧
- **語言：** Go 1.21（`go.mod` 釘選 `toolchain go1.26.10`）
- **WebSocket：** gorilla/websocket（v2 主要）、gobwas/ws（pkg/mux）
- **結構：**
  - `v1/` — Legacy REST + WebSocket API
  - `v2/rest/` — REST API v2（23 個 service，約 70 個 endpoint method）
  - `v2/websocket/` — WebSocket API v2（含重連、心跳、訂閱管理）
  - `pkg/models/` — 27 個 domain model 套件
  - `pkg/utils/` — Nonce 產生器
  - `pkg/mux/` — 訊息多工器
- **測試：** stretchr/testify、httptest mock

## 目前狀態
- v1 API 功能完整但為 legacy，所有 REST service 公開方法已加上 `V1` 前綴以區分版本
- v2 REST endpoint 覆蓋率約 95%（對比 Bitfinex 官方文件 73 個 endpoint）
- v2 WebSocket 功能完整（含重連、心跳、訂閱管理）
- 已完成 P0-P3 全面程式碼品質修復（2026-04-01）與 `toolchain go1.26.10` 安全升級（2026-10-23）

## 重要決策
- 保留 v1 向後相容但不再新增功能，所有 v1 REST service 方法加上 `V1` 前綴明確標示版本
- WebSocket 使用 gorilla/websocket 而非 gobwas/ws（gobwas 僅用於 mux 模組）
- 最低語言相容宣告為 Go 1.21（使用 `io.ReadAll` 等現代 API），並於 `go.mod` 釘選 `toolchain go1.26.10`
- v2 新增 endpoint 暫以 `[]interface{}` 回傳，待後續逐步補齊 typed model 層

## 待辦 / 下一步
- 為新增的 v2 endpoint（account, movements, alerts, settings 等）建立 typed model
- 為新增的 v2 endpoint 撰寫 unit test
- 評估是否將 `go.mod` 最低語言版本與舊版相依套件（`gorilla/websocket`、`stretchr/testify`、已停更的 `op/go-logging`）現代化
- v2/rest Client 的 23 個 service 可考慮拆分為更小的 interface（Interface Segregation）
