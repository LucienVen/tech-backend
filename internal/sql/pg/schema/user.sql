CREATE TABLE users (
    id UUID PRIMARY KEY,
    -- 教职人员ID，使用UUID类型
    username VARCHAR(255) NOT NULL,
    -- 用户名
    nick_name VARCHAR(255),
    -- 昵称
    passwd VARCHAR(255) NOT NULL,
    -- 密码（已加密）
    phone VARCHAR(20),
    -- 联系方式
    email VARCHAR(255),
    -- 邮箱地址
    status SMALLINT NOT NULL DEFAULT 0,
    -- 用户状态：0-未激活，1-正常，2-禁用，3-注销，9-已删除
    is_delete BIGINT DEFAULT 0,
    -- 是否删除（逻辑删除标记）
    creator VARCHAR(255),
    -- 创建者
    updater VARCHAR(255),
    -- 更新者
    create_time BIGINT NOT NULL,
    -- 创建时间
    update_time BIGINT NOT NULL,
    -- 更新时间
    CONSTRAINT users_email_unique UNIQUE(email) -- 如果需要邮箱唯一约束
);
-- 为表添加注释
COMMENT ON TABLE users IS '存储用户基本信息，包括用户名、昵称、邮箱、密码、状态等。';
-- 为列添加注释
COMMENT ON COLUMN users.id IS '用户的唯一标识符，使用UUID类型。';
COMMENT ON COLUMN users.username IS '用户的用户名，不能为空。';
COMMENT ON COLUMN users.nick_name IS '用户的昵称。';
COMMENT ON COLUMN users.passwd IS '用户的加密密码。';
COMMENT ON COLUMN users.phone IS '用户的联系方式。';
COMMENT ON COLUMN users.email IS '用户的邮箱地址。';
COMMENT ON COLUMN users.status IS '用户状态：0-未激活，1-正常，2-禁用，3-注销，9-已删除。';
COMMENT ON COLUMN users.is_delete IS '逻辑删除标记，0表示未删除，1表示已删除。';
COMMENT ON COLUMN users.creator IS '记录该用户的创建者。';
COMMENT ON COLUMN users.updater IS '记录该用户的更新者。';
COMMENT ON COLUMN users.create_time IS '记录该用户的创建时间。';
COMMENT ON COLUMN users.update_time IS '记录该用户的更新时间。';