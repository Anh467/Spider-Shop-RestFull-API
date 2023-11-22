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
	updatedAt datetime default current_timestamp
);

-- ########################### dbo.Cate ###########################
create table `Cate` (
	`CateID` int auto_increment unique not null,
    `Name` nvarchar(50) not null,
    `Desc` nvarchar(300) default "",
    `Image` varchar(200) default "",
    `Status` varchar(10) default "Normal",
    `Revenue` double default 0,
    `NumWareHouse` int default 0,
    `NumOrder` int default 0
)
-- tracking
