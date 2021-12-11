/*
  Warnings:

  - You are about to drop the column `text` on the `transactions` table. All the data in the column will be lost.
  - Added the required column `amount` to the `transactions` table without a default value. This is not possible if the table is not empty.

*/
-- RedefineTables
PRAGMA foreign_keys=OFF;
CREATE TABLE "new_transactions" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "amount" REAL NOT NULL,
    "created_at" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" DATETIME NOT NULL,
    "account_id" TEXT NOT NULL
);
INSERT INTO "new_transactions" ("account_id", "created_at", "id", "updated_at") SELECT "account_id", "created_at", "id", "updated_at" FROM "transactions";
DROP TABLE "transactions";
ALTER TABLE "new_transactions" RENAME TO "transactions";
PRAGMA foreign_key_check;
PRAGMA foreign_keys=ON;
