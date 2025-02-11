CREATE TABLE monthly_costs (
  cost_type_id BIGINT UNSIGNED NOT NULL COMMENT 'コスト種別ID',
  cost_year INT UNSIGNED NOT NULL COMMENT '年',
  cost_month INT UNSIGNED NOT NULL COMMENT '月',
  cost_yen INT UNSIGNED NOT NULL COMMENT '料金(円)',
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '作成日時(UTC)',
  updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時(UTC)',

  PRIMARY KEY (cost_type_id, cost_year, cost_month),
  CONSTRAINT fk_monthly_costs_cost_type_id FOREIGN KEY (cost_type_id) REFERENCES cost_types (id)
)
COMMENT '月次コスト'
;
