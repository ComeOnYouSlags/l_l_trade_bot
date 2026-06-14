# requirements

require (
	github.com/google/uuid v1.4.0
	github.com/gorilla/websocket v1.5.1
	github.com/stretchr/testify v1.8.4
)



# data log format

DATE - TIME - PAIR - PRICE - MEAN - TREND

01.01.1970 - 19:59:59 - usdt/btc - 60001.0000000001 - 59899.0000000001 - 2.56


TREND


slope = ((n*sum(x*y))-(sum(x)*(sum(y))))/(n*sum(x^2)-(sum(x))^2)

slope > 0 - ascending
slope < 0 - discending
|slope| = trend power


FOLLOW THE TREND STRATEGY

PRICE > MEAN AND slope > 0 - BUY
PRICE < MEAN AND clope < 0 - SELL



# TESTS

go test ./tests/ -v