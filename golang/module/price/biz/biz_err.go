package biz

// "" stands for "Normal"
var Price_STATUS = []string{"Normal", "Hot", "Deleted", ""}

//const USER_ERR_CANNOT_CREATE = "can not create new user"

// ProductID
const PRICE_ERR_PRODUCTID_NOT_FOUND = "The requested productid does not exist"
const PRICE_ERR_PRODUCTID_POSITIVE = "ProductID must be a positive integer"

// price
const PRICE_ERR_Price_NOT_NEGATIVE = "ProductID can't be negative floating point number"

// PriceID
const PRICE_ERR_PriceID_NOT_FOUND = "The requested priceid does not exist"
const PRICE_ERR_PriceID_POSITIVE = "The PriceID must be a positive intege"

// status
const PRICE_ERR_Status_NOT_APPROPIRATE = `Status must be a {"Normal", "Hot", "Deleted", "" }`
