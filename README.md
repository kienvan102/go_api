# APIs brief

## Get a student

  Returns json data about a single student.

* **URL**

  v1/students/

* **Method:**

  `GET`
  
*  **URL Params**

   **Required:**
 
   `N/A`

* **Data Params**

  None

* **Success Response:**

  * **Code:** 200 <br />
    **Content:** `{ id : <numbder of id>, name : strings of name }`
 
* **Error Response:**

  * **Code:** 404 NOT FOUND <br />
    **Content:** `{ error : "User doesn't exist" }`

* **Sample Call:**

  ```
   <url>/v1/students/
  ```
<!-- 
  OR

  * **Code:** 401 UNAUTHORIZED <br />
    **Content:** `{ error : "You are unauthorized to make this request." }` -->

<!-- * **Sample Call:**

  ```javascript
    $.ajax({
      url: "/users/1",
      dataType: "json",
      type : "GET",
      success : function(r) {
        console.log(r);
      }
    }); -->
  <!-- ``` -->

## Get list of student
Returns json data about list of students.

* **URL**

  v1/students/

  `example`


* **Method:**

  `GET`
  
*  **URL Params**

   **Required:**
 
    `id=[integer]`

* **Data Params**

  Limit
  Page

* **Success Response:**

  * **Code:** 200 <br />
    **Content:** `{ id : <numbder of id>, name : strings of name }`
 
* **Error Response:**

  * **Code:** 404 NOT FOUND <br />
    **Content:** `{ error : "User doesn't exist" }`

* **Sample Call:**

  ```
   <url>/v1/students/?Limit=2&page=2
  ```
