# GO-APIS

## Environment file 

Needed variables for running on local:
```
DB_USER=
DB_PWD=
DB_ADDRESS=
DB_PORT=
DB_NAME=
```

## Architecture
Following clean architecture strictly, each service have a module. Currently, only "student" module have been implemented. <br/>

Each module have three layer and 1 entity (Model) including:transport, bussiness (namely biz) and storage. Each layer must be independently.

* **Transport:**<br/>Recieve HTTP request from client and parse data, response JSON format to client. Then, the rest of task is delegated for Bussiness layer.

* **Bussiness** <br/> This layer implement logic or mathametic of services. This layer do not call dicretly to database but continue to delegate for Storage layer.

* **Storage** <br/> Storage layer is responsible for storing and retrieving data from database or filesystem.
## APIS brief

I versionized the APIs, so each endpoint starting with name of version
 

|                  **uri**                 | **Method** | **Operation** |                     **description**                     | **Resquest Body**     |        **Response body**       |
|:----------------------------------------:|:----------:|:-------------:|:-------------------------------------------------------:|-----------------------|:------------------------------:|
| /v1/student/[:id]                        | GET        | Read          | Get a student by id                                     | N/A                   | a json format of data          |
| /v1/student/                             | POST       | Create        | Create a new student                                    | a json format of data | code of response               |
| /v1/student/[:id]                        | PATH       | Update        | Update a student by id                                  | a json format of data | code of response               |
| /v1/student/?page=[number]&lmit=[number] | GET        | Read          | Get a list of student according to page and limit param | N/A                   | a json format of array of data |
| /v1/student/[:id]                        | DELETE     | Delete        | Delete a student by id                                  | N/A                   | code of response               |


### Example:
**Get a student**
* ***url:***
```
<domain>/v1/student/1
```
* ***Response body***
  * ***Code success: 200***<br/>
    ***Content:
    ```json
    {
    "data": {
        "id": 9,
        "fullname": "Stephen Hawking",
        "sex": "male",
        "phone": 0,
        "email": "abc@mail.com"
      }
    }
    ```
* ***Error response***
  * **Code:** 404 NOT FOUND <br />
    **Content:** `{ error : "User doesn't exist" }`