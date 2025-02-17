/*
 Navicat Premium Data Transfer

 Source Server         : localhost_POSTGRES
 Source Server Type    : PostgreSQL
 Source Server Version : 140012
 Source Host           : localhost:5432
 Source Catalog        : manajemen_pengetahuan
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 140012
 File Encoding         : 65001

 Date: 20/06/2024 17:17:28
*/


-- ----------------------------
-- Sequence structure for article_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."article_id_seq";
CREATE SEQUENCE "public"."article_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;
ALTER SEQUENCE "public"."article_id_seq" OWNER TO "postgres";

-- ----------------------------
-- Sequence structure for user_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."user_id_seq";
CREATE SEQUENCE "public"."user_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;
ALTER SEQUENCE "public"."user_id_seq" OWNER TO "postgres";

-- ----------------------------
-- Table structure for article
-- ----------------------------
DROP TABLE IF EXISTS "public"."article";
CREATE TABLE "public"."article" (
  "id" int4 NOT NULL DEFAULT nextval('article_id_seq'::regclass),
  "user_id" int4,
  "title" varchar COLLATE "pg_catalog"."default",
  "content" text COLLATE "pg_catalog"."default",
  "updated_at" int4,
  "created_at" int4
)
;
ALTER TABLE "public"."article" OWNER TO "postgres";

-- ----------------------------
-- Records of article
-- ----------------------------
BEGIN;
INSERT INTO "public"."article" VALUES (2, 2, 'Article Title', 'Hello world!', 1717507931, 1717507931);
COMMIT;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS "public"."user";
CREATE TABLE "public"."user" (
  "id" int4 NOT NULL DEFAULT nextval('user_id_seq'::regclass),
  "email" varchar COLLATE "pg_catalog"."default",
  "password" varchar COLLATE "pg_catalog"."default",
  "name" varchar COLLATE "pg_catalog"."default",
  "updated_at" int4,
  "created_at" int4
)
;
ALTER TABLE "public"."user" OWNER TO "postgres";

-- ----------------------------
-- Records of user
-- ----------------------------
BEGIN;
INSERT INTO "public"."user" VALUES (2, 'test@test.com', '$2a$10$Y2pbNwTxsvqo418AGerJfO/QTPWPFcnYG4aSJnWtzFV81UC7lXM0C', 'testing', 1717507901, 1717507901);
COMMIT;

-- ----------------------------
-- Function structure for created_at_column
-- ----------------------------
DROP FUNCTION IF EXISTS "public"."created_at_column"();
CREATE OR REPLACE FUNCTION "public"."created_at_column"()
  RETURNS "pg_catalog"."trigger" AS $BODY$

BEGIN
	NEW.updated_at = EXTRACT(EPOCH FROM NOW());
	NEW.created_at = EXTRACT(EPOCH FROM NOW());
    RETURN NEW;
END;

$BODY$
  LANGUAGE plpgsql VOLATILE
  COST 100;
ALTER FUNCTION "public"."created_at_column"() OWNER TO "postgres";

-- ----------------------------
-- Function structure for update_at_column
-- ----------------------------
DROP FUNCTION IF EXISTS "public"."update_at_column"();
CREATE OR REPLACE FUNCTION "public"."update_at_column"()
  RETURNS "pg_catalog"."trigger" AS $BODY$

BEGIN
    NEW.updated_at = EXTRACT(EPOCH FROM NOW());
    RETURN NEW;
END;

$BODY$
  LANGUAGE plpgsql VOLATILE
  COST 100;
ALTER FUNCTION "public"."update_at_column"() OWNER TO "postgres";

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."article_id_seq"
OWNED BY "public"."article"."id";
SELECT setval('"public"."article_id_seq"', 3, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."user_id_seq"
OWNED BY "public"."user"."id";
SELECT setval('"public"."user_id_seq"', 3, true);

-- ----------------------------
-- Triggers structure for table article
-- ----------------------------
CREATE TRIGGER "create_article_created_at" BEFORE INSERT ON "public"."article"
FOR EACH ROW
EXECUTE PROCEDURE "public"."created_at_column"();
CREATE TRIGGER "update_article_updated_at" BEFORE UPDATE ON "public"."article"
FOR EACH ROW
EXECUTE PROCEDURE "public"."update_at_column"();

-- ----------------------------
-- Primary Key structure for table article
-- ----------------------------
ALTER TABLE "public"."article" ADD CONSTRAINT "article_id" PRIMARY KEY ("id");

-- ----------------------------
-- Triggers structure for table user
-- ----------------------------
CREATE TRIGGER "create_user_created_at" BEFORE INSERT ON "public"."user"
FOR EACH ROW
EXECUTE PROCEDURE "public"."created_at_column"();
CREATE TRIGGER "update_user_updated_at" BEFORE UPDATE ON "public"."user"
FOR EACH ROW
EXECUTE PROCEDURE "public"."update_at_column"();

-- ----------------------------
-- Primary Key structure for table user
-- ----------------------------
ALTER TABLE "public"."user" ADD CONSTRAINT "user_id" PRIMARY KEY ("id");

-- ----------------------------
-- Foreign Keys structure for table article
-- ----------------------------
ALTER TABLE "public"."article" ADD CONSTRAINT "article_user_id" FOREIGN KEY ("user_id") REFERENCES "public"."user" ("id") ON DELETE CASCADE ON UPDATE CASCADE;
