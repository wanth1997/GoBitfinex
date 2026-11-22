# Bitfinex REST V2 API Reference

```go
import "github.com/wanth1997/GoBitfinex/v2/rest"
```

`rest.Client` exposes 23 services covering public market data and authenticated account, order, position, funding, and wallet operations.

For full struct fields, parameter types, and doc comments, run:

```bash
go doc -all github.com/wanth1997/GoBitfinex/v2/rest
go doc github.com/wanth1997/GoBitfinex/v2/rest.OrderService
```

## Client Initialization

```go
client := rest.NewClient()
authClient := rest.NewClient().Credentials(apiKey, apiSecret)
```

Additional constructors: `NewClientWithURL`, `NewClientWithHTTPClient`, `NewClientWithURLHTTPClient`, `NewClientWithHttpDo`, `NewClientWithURLHttpDo`, `NewClientWithURLNonce`, `NewClientWithURLHttpDoNonce`, `NewClientWithSynchronousNonce`, `NewClientWithSynchronousURLNonce`.

## Services (`*rest.Client`)

| Field | Type | Methods |
|---|---|---|
| `Account` | `AccountService` | `UserInfo()`, `Summary()`, `MarginInfo(key)`, `AvailableBalance(symbol, dir, rate, orderType)`, `LoginsHistory()`, `AuditHistory()` |
| `Alerts` | `AlertService` | `All(alertType)`, `Set(alertType, symbol, price)`, `Delete(symbol, price)` |
| `Book` | `BookService` | `All(symbol, precision, priceLevels)` |
| `Candles` | `CandleService` | `History(symbol, resolution)`, `HistoryWithQuery(symbol, resolution, start, end, limit, sort)`, `Last(symbol, resolution)` |
| `Currencies` | `CurrenciesService` | `Conf(label, symbol, unit, explorer, pairs)` |
| `Derivatives` | `DerivativesService` | `CollateralLimits(symbol)` |
| `Funding` | `FundingService` | `Offers(symbol)`, `OfferHistory(symbol)`, `SubmitOffer(fo)`, `CancelOffer(fc)`, `CancelAllOffers()`, `Loans(symbol)`, `LoansHistory(symbol)`, `Credits(symbol)`, `CreditsHistory(symbol)`, `Trades(symbol)`, `Info(key)`, `KeepFunding(args)`, `AutoRenew(currency, amount, rate, period, enabled)`, `Close(id, fundingType)` |
| `Invoice` | `InvoiceService` | `GenerateInvoice(payload)` |
| `Ledgers` | `LedgerService` | `Ledgers(currency, start, end, max)` |
| `Liquidations` | `LiquidationsService` | `History()` |
| `Market` | `MarketService` | `AveragePrice(pld)`, `ForeignExchangeRate(pld)` |
| `Movements` | `MovementsService` | `History(currency)`, `AllHistory()` |
| `Orders` | `OrderService` | `All()`, `AllActive()`, `AllHistory()`, `GetByOrderId(orderID)`, `GetBySymbol(symbol)`, `GetHistoryByOrderId(orderID)`, `GetHistoryBySymbol(symbol)`, `OrderTrades(symbol, orderID)`, `SubmitOrder(onr)`, `SubmitUpdateOrder(our)`, `SubmitCancelOrder(oc)`, `CancelOrderMulti(args)`, `OrderMultiOp(ops)`, `OrderNewMultiOp(onr)`, `OrderUpdateMultiOp(our)`, `CancelOrderMultiOp(orderID)`, `CancelOrdersMultiOp(ids)` |
| `Platform` | `PlatformService` | `Status()` |
| `Positions` | `PositionService` | `All()`, `History()`, `Audit()`, `Snapshot()`, `Claim(cp)`, `Increase(symbol, amount)`, `IncreaseInfo(symbol, amount)` |
| `Rankings` | `RankingsService` | `History(key, timeFrame, symbol, limit)` |
| `Settings` | `SettingsService` | `Read(keys)`, `Write(settings)`, `Delete(keys)` |
| `Stats` | `StatsService` | `PositionLast(symbol, side)`, `PositionHistory(symbol, side)`, `FundingLast(symbol)`, `FundingHistory(symbol)`, `CreditSizeLast(symbol, side)`, `CreditSizeHistory(symbol, side)`, `SymbolCreditSizeLast(fundingSymbol, tradingSymbol)`, `SymbolCreditSizeHistory(fundingSymbol, tradingSymbol)` |
| `Status` | `StatusService` | `DerivativeStatus(symbol)`, `DerivativeStatusMulti(symbols)`, `DerivativeStatusAll()`, `DerivativeStatusHistory(key)` |
| `Tickers` | `TickerService` | `Get(symbol)`, `GetMulti(symbols)`, `All()` |
| `TickersHistory` | `TickerHistoryService` | `Get(pld)` |
| `Trades` | `TradeService` | `AccountAll()`, `AccountAllWithSymbol(symbol)`, `AccountHistoryWithQuery(...)`, `PublicHistoryWithQuery(...)` |
| `Wallet` | `WalletService` | `Wallet()`, `Transfer(from, to, currency, currencyTo, amount)`, `DepositAddress(wallet, method)`, `CreateDepositAddress(wallet, method)`, `Withdraw(wallet, method, amount, address)`, `SetCollateral(symbol, amount)` |
