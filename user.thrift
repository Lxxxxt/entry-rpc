namespace go user_service

struct User {
    1: required i64 id,
    2: required string username,
    3: required string email,
    4: required i32 age
}

service UserService {
    User GetUser(1:i64 userID),
    void SaveUser(1:User user),
    void DeleteUser(1:i64 userID)
}