/* name: RegisterUser :exec */
INSERT INTO users (id, username, password, email) VALUES (?, ?, ?, ?);

/* name: CheckUsernameExistence :one */
SELECT COUNT(*) FROM users WHERE username = ?;

/* name: CheckEmailExistence :one */
SELECT COUNT(*) FROM users WHERE email = ?;

/* name: LoginUser :one */
SELECT id, password
FROM users
WHERE username = ?;

/* name: GetUserDetails :one */
SELECT id, username, password, email, phone, address
FROM users
WHERE id = ?;

/* name: UpdateUserDetails :exec */
UPDATE users
SET
    email = ?,
    password = ?,
    phone = ?,
    address = ?
WHERE id = ?;







