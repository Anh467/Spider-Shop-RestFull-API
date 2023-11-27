create database SpiderShop;
use  SpiderShop;
-- ########################### dbo.User ###########################
create table User(
    UserID int primary key auto_increment not null unique,
    Account varchar(100) unique not null,
    Password varchar(255) not null, -- Consider using a larger size for password hash
    Name nvarchar(50) not null,
    Image varchar(200) default "",
    Gender varchar(10) default "Others",
    Phone varchar(20) default "",
    Mail varchar(100) default "" check (Mail = '' or Mail like '%@%.%'), -- Basic email format validation
    Birth datetime default current_timestamp,
    Role varchar(10) default "User",
    Expenditure float8 default 0,
    NumOrder int default 0,
    createdAt datetime default current_timestamp,
	updatedAt datetime default current_timestamp on update current_timestamp
);

-- ########################### dbo.Cate ###########################
create table `Cate` (
	`CateID` int primary key auto_increment unique not null,
    `Name` nvarchar(50) not null,
    `Desc` nvarchar(300) default "",
    `Image` varchar(200) default "",
    `Status` varchar(10) default "Normal",
    `Revenue` double default 0,
    `NumWareHouse` int default 0,
    `NumOrder` int default 0
);
-- ########################### dbo.Product ###########################
create table `Product`(
	`ProductID` int primary key auto_increment not null,
    `CateID` int,
    `Name` nvarchar(50) not null,
    `Desc` nvarchar(300) default "",
    `Image` varchar(200) default "",
    `Status` varchar(10) default "Normal",
    createdAt datetime default current_timestamp,
	updatedAt datetime default current_timestamp on update current_timestamp,
    INDEX idx_CateID (CateID),
    constraint CHK_Status_Product check(Status in ('Normal', 'Hot', 'Deleted')),
    CONSTRAINT FK_Product_Category FOREIGN KEY (CateID) REFERENCES Cate(CateID)
);
-- ###############	############ dbo.Tracking ###########################
alter table User
drop index idx_UserID_User;
-- $$$$$$$$$$$$$$$$$$$$$$$$$$$ dbo.User $$$$$$$$$$$$$$$$$$$$$$$$$$$
alter table User
add index idx_Account_User(Account);
-- $$$$$$$$$$$$$$$$$$$$$$$$$$$ dbo.Cate $$$$$$$$$$$$$$$$$$$$$$$$$$$
alter table Cate
add CONSTRAINT CHK_Status CHECK (Status IN ('Normal', 'Hot', 'Deleted'));
alter table Cate
add index idx_Status_User(Status);
-- $$$$$$$$$$$$$$$$$$$$$$$$$$$ dbo.Product $$$$$$$$$$$$$$$$$$$$$$$$$$$
alter table Product
add index idx_Status_User(Status);
