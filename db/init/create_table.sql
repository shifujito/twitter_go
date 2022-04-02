-- --ユーザーの作成
-- CREATE USER shion;
-- --DBの作成
-- CREATE DATABASE twitter;
--ユーザーにDBの権限をまとめて付与
-- GRANT ALL PRIVILEGES ON *.* TO 'shion'@'%';
--ユーザーを切り替え
-- \c docker
--テーブルを作成
CREATE TABLE users (
  id integer,
  name varchar(30)
);
--テーブルにデータを挿入
INSERT INTO users VALUES (1,'The Very Hungry Caterpillar');
