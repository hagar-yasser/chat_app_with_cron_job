# Chat App

## Ruby version

ruby-2.5.8

## Rails version

rails 5.0.0

## Go version

go1.16

## Definitions

- Application: An application resembles a user's account and can contain several chats and is identified with an autogenerated token.

## Models

- Application (token,name,chats_count)
- Chat (number,messages_count)
- Message (number,body)

## Functionalities

- creating/updating applications
- creating chats within applications
- creating/updating messages within chats
- searching through messages within a chat

## Assumptions

1. Newly created chats and messages are concentrated in hot/trendy applications or chats and not normally distributed.

- The cron job updating the database depends on the the fact that changes are most pobably concentrated in a few amount of applications or chats
  Otherwise if the changes where normally distributed using a cron job will only be a delay without reducing any stress from the DB server.

## Endpoints

- GET http://localhost:3000/applications/{app_token}
  1. gets a specific application
  2. returns the application object
- POST http://localhost:3000/applications
  1. creates a new application
  2. returns the application object created
- PUT http://localhost:3000/applications/{app_token}

  1. updates the name of a specific application with the _name_ field in the JSON request body
  2. returns the application object updated

- GET http://localhost:3000/applications/{app_token}/chats
  1. gets chats of a specific application
  2. returns an array of chat objects
- POST http://localhost:3000/applications/{app_token}/chats

  1. creates a new chat within a specific application
  2. returns the chat object created

- GET http://localhost:3000/applications/{app_token}/chats/{chat_no}/messages
  1. gets messages of a specific chat
  2. returns an array of message objects
- POST http://localhost:3000/applications/{app_token}/chats/{chat_no}/messages
  1. creates a new message in a specific chat with a body as with the _body_ field in the JSON request body
  2. returns the message object created
- PUT http://localhost:3000/applications/{app_token}/chats/{chat_no}/messages/{message_no}

  1. updates the body of a specific message with the _body_ field in the JSON request body
  2. returns the message object updated

- GET http://localhost:3000/applications/{app_token}/chats/{chat_no}/messages/search
  1. returns matches with the _query_ field in the JSON request body in the messages of a specific chat/returns an array of messages in the chat specified that have a match with the query

## Services

- app : service creating and running the main rails application for routing and background asynchronous jobs
- mysql:5.7
- redis:7
- elasticsearch:8.3.2
- sidekiq v6
- gocron : service creating and running the cron job

## How to run

using the docker-compose.yml file in the root path of the repository run : docker-compose up

## Notes

I notice the Error UnixHTTPConnectionPool(host='localhost', port=None): Read timed out. (read timeout=60). frequently running docker-compose up. I haven't yet found a practical solution for it. However, if this problem was encountered try the following steps.

### Fixes

#### Incase the ram exhaustion or the timeout error were encoutered try doing the following steps

1. Restart docker (systemctl restart docker)
2. Close the browsers open
3. Run docker-compose down
4. Run docker-compose down -v
5. Retry

## Future Work

- Add caching layer that stores the applications, chats, or messages created/read recently. Adding a caching layer would:-

  1. Decrease the latency of the response and reduce the SQL queries to the DB.
  2. Mitigate the inconsistency risk caused by creating chats and messages asynchronously and reading them before they were created.

- Separating the request frontend server from the app server by separating the routing functionalities from the backend logic. This can be done through:-
  1. Creating 2 applications one that receives HTTP requests from the user and the other that processes the backend logic.
  2. Creating a Message Broker server
  3. The request frontend server would process the HTTP requests and forwards a message in the message queue of the app server
  4. This way the app server can be scaled separate from the request frontend server depending on the load of the application
  5. A load balancer can be used to divide the requests between the running app servers
