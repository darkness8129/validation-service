## Architecture

The API is built using Clean Architecture. The principles of DIP (Dependency Inversion Principle) and DI (Dependency Injection) are utilized.

## Commands

- `docker-compose up` - to start the service using docker-compose on 8080 port

## Possible Improvements

1. **Entity and Storage layers:** I didn't create the entity and storage layers because there was no need for them. However, if in the future there's a requirement, for example, to interact with a database, then it will definitely be necessary to implement these layers.

2. **Errors:** Currently, errors are returned sequentially from the service. That is, if there are multiple errors in validation of the card, to discover the next error, you need to resolve the previous one. I think, if necessary, it could be made more convenient: not to stop the validation process as soon as one error is found, but to continue it returning all possible validation errors at once.

3. **Tests:** The `ccvalidation` package was covered with unit tests because there was no additional business logic at the service level. Again, if work with storage/api is added, it would be better to write integration tests at the service level (using mocks).

4. **Tokenization:** Passing card data in an untokenized form between services is not secure, so it's better to ensure that a token, rather than the card number in its plain form, is passed to the service input.
