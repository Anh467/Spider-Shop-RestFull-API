package biz

var PRODUCT_STATUS = []string{"Normal", "Hot", "Deleted"}

// const USER_ERR_CANNOT_CREATE = "can not create new user"
const PRODUCT_ERR_TABLE_NO_CHANGE = "There is no change in the table product"

// PRODUCt
const PRODUCT_ERR_Temporarily_deleted = "Temporarily deleted"

// ProductID
const PRODUCT_ERR_PRODUCT_ID_NOT_EXIST = "Product ID not exist"
const PRODUCT_ERR_PRODUCT_ID_POSITIVE = "Product ID can't be negative and 0"

// CateID
const PRODUCT_ERR_CateID_NOT_EXIST = "Category ID not found in category list"
const PRODUCT_ERR_CateID_POSITIVE = "Category ID can't be negative"

// Name
const USER_PATTERN_Name = `^[a-zA-ZĐđÀ-ỹ].{0,49}$`
const USER_ERR_PATTERN_Name = "Can’t have special character and number"

// status
const CATE_ERR_Status_EXIST = "Status of category hust in Deleted, Normal, Hot"
