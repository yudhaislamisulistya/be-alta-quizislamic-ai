---
markmap:
  colorFreezeLevel: 2
---

# API Routes Markmap

## authentications

- POST /authentications/login
- POST /authentications/forgot-password
- POST /authentications/change-password

## users

- GET /users
- PUT /users/:id/password
- GET /users/filter
- GET /users/gender/:gender
- GET /users/method/:method
- GET /users/verified-email/:status
- GET /users/birth-year/:year
- GET /users/empty-profile-photo
- GET /users/token-expired
- GET /users/token-verified-email/:token
- GET /users/joined-date-range
- GET /users/search
- GET /users/sort
- GET /users/pagination
- GET /users/:id
- POST /users
- PUT /users/:id
- DELETE /users/:id
- GET /users/verification-email
- POST /users/verification-email

## questions

- POST /questions
- POST /questions/multiple-choice
- POST /questions/true-false
- POST /questions/fill-in
- GET /questions
- GET /questions/search
- GET /questions/sort
- GET /questions/pagination
- GET /questions/all
- GET /questions/:user_id/:quiz_id

## questions-categories

- GET /questions-categories
- GET /questions-categories/name/:name
- GET /questions-categories/search
- GET /questions-categories/sort
- GET /questions-categories/pagination
- GET /questions-categories/:id
- POST /questions-categories
- PUT /questions-categories/:id
- DELETE /questions-categories/:id

## wallets

- GET /wallets
- GET /wallets/pagination
- GET /wallets/sort
- GET /wallets/:id
- POST /wallets
- PUT /wallets/:id
- DELETE /wallets/:id
- POST /wallets/send

## quizzes

- GET /quizzes
- GET /quizzes/search
- GET /quizzes/pagination
- GET /quizzes/sort
- GET /quizzes/:id
- GET /quizzes/:user_id
- GET /quizzes/:user_id/:quiz_id
- POST /quizzes
- PUT /quizzes/:id
- DELETE /quizzes/:id

## levels

- GET /levels
- GET /levels/search
- GET /levels/pagination
- GET /levels/sort
- GET /levels/:id
- POST /levels
- PUT /levels/:id
- DELETE /levels/:id

## packages

- GET /packages
- GET /packages/search
- GET /packages/pagination
- GET /packages/sort
- GET /packages/:id
- POST /packages
- PUT /packages/:id
- DELETE /packages/:id

## quiz-histories

- GET /quiz-histories
- GET /quiz-histories/pagination
- GET /quiz-histories/sort
- GET /quiz-histories/score
- GET /quiz-histories/attempt-date-range
- GET /quiz-histories/:id
- GET /quiz-histories/user/:id
- GET /quiz-histories/quizzes/:id
- POST /quiz-histories
- PUT /quiz-histories/:id
- DELETE /quiz-histories/:id

## quiz-answers

- GET /quiz-answers
- GET /quiz-answers/pagination
- GET /quiz-answers/sort
- GET /quiz-answers/filter
- GET /quiz-answers/:id
- GET /quiz-answers/quiz-histories/:id
- GET /quiz-answers/questions/:id
- POST /quiz-answers
- PUT /quiz-answers/:id
- DELETE /quiz-answers/:id

## package-histories

- GET /package-histories
- GET /package-histories/search
- GET /package-histories/pagination
- GET /package-histories/sort
- GET /package-histories/filter
- GET /package-histories/:id
- GET /quiz-histories/transaction-date-range
- GET /package-histories/packages/:id
- GET /package-histories/users/:id
- POST /package-histories
- PUT /package-histories/:id
- DELETE /package-histories/:id

## wallet-transactions

- GET /wallet-transactions
- GET /wallet-transactions/pagination
- GET /wallet-transactions/sort
- GET /wallet-transactions/filter
- GET /wallet-transactions/amount
- GET /wallet-transactions/transaction-date-range
- GET /wallet-transactions/wallets/:id
- GET /wallet-transactions/:id
- POST /wallet-transactions
- PUT /wallet-transactions/:id
- DELETE /wallet-transactions/:id

## quiz-reviews

- GET /quiz-reviews
- GET /quiz-reviews/search
- GET /quiz-reviews/pagination
- GET /quiz-reviews/sort
- GET /quiz-reviews/filter
- GET /quiz-reviews/:id
- GET /quiz-reviews/quizzes/:id
- GET /quiz-reviews/users/:id
- POST /quiz-reviews
- PUT /quiz-reviews/:id
- DELETE /quiz-reviews/:id

## activity-logs

- GET /activity-logs
- GET /activity-logs/pagination
- GET /activity-logs/sort
- GET /activity-logs/filter
- GET /activity-logs/:id
- GET /activity-logs/users/:id
- POST /activity-logs
- PUT /activity-logs/:id
- DELETE /activity-logs/:id

## notifications

- GET /notifications
- GET /notifications/pagination
- GET /notifications/sort
- GET /notifications/filter
- GET /notifications/:id
- GET /notifications/users/:id
- POST /notifications
- PUT /notifications/:id
- DELETE /notifications/:id
