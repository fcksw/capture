## 每日行情表，存储每日筛选后的股票
CREATE TABLE `stock_quote_daily_info` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `symbol` varchar(256) DEFAULT NULL COMMENT '股票代码（symbol）',
  `name` varchar(45) DEFAULT NULL COMMENT '股票名称（name）',
  `current_price` varchar(45) DEFAULT NULL COMMENT '当前价格（current）',
  `chg` varchar(45) DEFAULT NULL COMMENT '涨跌额（chg）',
  `percent` varchar(45) DEFAULT NULL COMMENT '涨跌幅百分比（percent）',
  `current_year_percent` varchar(45) DEFAULT NULL COMMENT '年初至今涨跌（current_year_percent）',
  `volume` varchar(45) DEFAULT NULL COMMENT '成交量（volume）',
  `amount` varchar(45) DEFAULT NULL COMMENT '成交额（amount）',
  `turnover_rate` varchar(45) DEFAULT NULL COMMENT '换手率（turnover_rate）',
  `pe_ttm` varchar(45) DEFAULT NULL COMMENT '市盈率（pe_ttm）',
  `dividend_yield` varchar(45) DEFAULT NULL COMMENT '股息率（dividend_yield）',
  `market_capital` varchar(45) DEFAULT NULL COMMENT '市值（market_capital）',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='股票行情表'