# design choice

I initially designed that the offline database to be a sqlite file but I was wrong.

## reasons why not sqlite 

- doesn't have a robust type system to meet the design requirements.
- will not enforce a RBAC at the database engine level.
- might not be sufficient for multiple reads and writes acreoss multiple go services. this risks data corruption.
- is just a file and will therefore be relatively easy for anyone to read and write without the app and authentication.


