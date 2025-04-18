ALTER TABLE student
ADD COLUMN is_delete TINYINT DEFAULT 0 COMMENT '是否删除（逻辑删除标记）',
ADD COLUMN creator VARCHAR(100) DEFAULT NULL COMMENT '创建者',
ADD COLUMN updater VARCHAR(100) DEFAULT NULL COMMENT '更新者',
ADD COLUMN create_time int DEFAULT NULL COMMENT '创建时间',
ADD COLUMN update_time int DEFAULT NULL COMMENT '更新时间';


-- 修改表AUTO_INCREMENT
alter table class AUTO_INCREMENT=0;