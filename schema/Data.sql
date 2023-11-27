use  SpiderShop;

-- Dữ liệu cho bảng Cate
INSERT INTO `Cate` (`Name`, `Desc`, `Image`, `Status`, `Revenue`, `NumWareHouse`, `NumOrder`) VALUES
('Category E', 'Description A', 'image_a.jpg', 'DELETED', 1000.50, 50, 20),
('Category F', 'Description B', 'image_b.jpg', 'DELETED', 800.75, 30, 15),
('Category G', 'Description C', 'image_c.jpg', 'DELETED', 1200.25, 40, 25),
('Category H', 'Description D', 'image_d.jpg', 'DELETED', 900.30, 25, 18);

-- Dữ liệu cho bảng Product
INSERT INTO `Product` (`CateID`, `Name`, `Desc`, `Image`, `Status`, `createdAt`, `updatedAt`) VALUES
(1, 'Product 1', 'Product 1 description', 'product_image_1.jpg', 'Normal', NOW(), NOW()),
(1, 'Product 2', 'Product 2 description', 'product_image_2.jpg', 'Hot', NOW(), NOW()),
(2, 'Product 3', 'Product 3 description', 'product_image_3.jpg', 'Normal', NOW(), NOW()),
(3, 'Product 4', 'Product 4 description', 'product_image_4.jpg', 'Hot', NOW(), NOW()),
(4, 'Product 5', 'Product 5 description', 'product_image_5.jpg', 'Normal', NOW(), NOW());

explain select CateID from Cate where Status = 'Deleted';

explain select * from Product
where CateID not in (select CateID from Cate where Status = 'Deleted');