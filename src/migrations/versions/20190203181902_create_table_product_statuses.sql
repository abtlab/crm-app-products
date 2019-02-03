CREATE TABLE ProductStatuses
(
  Id        UUID        NOT NULL,
  AccountId UUID        NOT NULL,
  Name      VARCHAR(64) NOT NULL
);

ALTER TABLE ProductStatuses
  ADD CONSTRAINT PK_ProductStatuses_Id PRIMARY KEY (Id),
  ADD CONSTRAINT UQ_ProductStatuses_AccountId_Name UNIQUE (AccountId, Name);

CREATE UNIQUE INDEX IX_ProductStatuses_AccountId_Name ON ProductStatuses (AccountId, Name);