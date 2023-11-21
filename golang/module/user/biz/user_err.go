package biz

var USER_GENDER = []string{"Male", "Female", "Others"}

const USER_ERR_CANNOT_CREATE = "can not create new user"

// user column constraints
// Account
const USER_ERR_Account_UNIQUE = "Account already exists"
const USER_PATTERN_Account = `^[^\s]{8,100}$`
const USER_ERR_PATTERN_Account = `Maximum 100 characters, minimum 8 characters, can’t have space`

// Password
const USER_PATTERN_Password = `^(?=.*[a-z])(?=.*[A-Z])(?=.*[@#$%^&+=])(?=\S+$).{8,100}$`
const USER_ERR_PATTERN_Password = "Maximum 100 characters and minimum 8 character , can’t have space, at least ( 1 uppercase, 1 lowercase, 1 special character)"

// Name
const USER_ERR_Name_MINIMUM_1 = "Name can't be blank"
const USER_PATTERN_Name = `^[a-zA-ZĐđÀ-ỹ].{0,49}$`
const USER_ERR_PATTERN_Name = "Can’t have special character and number"

// Gender
const USER_ERR_Gender = "just in (“Male”, “Female”, “Others”)"

// Mail
const USER_PATTERN_Mail = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
const USER_ERR_PATTERN_Mail = "Mail wrong fomrat"

// Birth
const USER_ERR_Birth = "Birth date wrong fomrat"

//Image
//Phone
//Role
//Expenditure
//NumOrder
