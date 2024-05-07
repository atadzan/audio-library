## Simple audio-library app which manages audio tracks

### Technology stack
* PostgresSQL 
* minIO

### Make sure you have installed following tools
* make 
* docker, docker-compose
* golang-migrate

### How to run app ?
#### Steps
* run `make dependencies-up` to run dependency technologies
* run `make apply-migration` to apply init database migration
* run `make run` to run application

### Track creation process
1. Upload track media file
2. After successful media file upload, it will respond path. Add this path in `Create track` endpoint params 

### Endpoints
* [Sign up](#sign-up)
* [Sign-in](#sign-in)
* [Get profile](#get-profile)
* [Upload track](#upload-track)
* [Create track](#create-track)
* [List tracks](#list-tracks)
* [Like track](#like-track)
* [Revert like](#revert-like)
* [List favourite tracks](#list-favourite-tracks)

#### **Sign up**

* **URL**

  /v1/auth/signUp

* **Method**

  `POST`

* **Headers**

  None

* **URL Params**

  None

* **Data Params**

    ```
    Content-Type: application/json
    {
       "fullname"  string
       "login"     string
       "password"  string 
    }
    ```

* **Success Response:**

    * **Code**: 200
    * **Content**:
      ```
      {
        "message"        "success"
        "token"          "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOjEsImlhdCI6MTcxNDkxNzcyMH0.3NkuFQ09_XhFQ91PDVLOQkHDXHW-pXjTnpBvN5qv90M"      
      }
      ```

* **Error Response**
    * **Code**: 400
    * **Description**: if input param is invalid
    * **Content**:
      ``` 
      {
        "message": "invalid input params"
      }
      ``` 

    * **Code**: 500
    * **Description**: something went wrong in server
    * **Content**:
       ``` 
       {
         "message": "internal server error"
       }
       ``` 

#### **Sign in**

* **URL**

  /v1/auth/signIn

* **Method**

  `POST`

* **Headers**

  None

* **URL Params**

  None

* **Data Params**

    ```
    Content-Type: application/json
    {
       "login"     string
       "password"  string 
    }
    ```

* **Success Response:**

    * **Code**: 200
    * **Content**:
      ```
      {
        "message"        "success"
        "token"          "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOjEsImlhdCI6MTcxNDkxNzcyMH0.3NkuFQ09_XhFQ91PDVLOQkHDXHW-pXjTnpBvN5qv90M"      
      }
      ```

* **Error Response**
    * **Code**: 400
    * **Description**: if input param is invalid
    * **Content**:
      ``` 
      {
        "message": "invalid input params"
      }
      ``` 

    * **Code**: 400
    * **Description**: if user not found
    * **Content**:
      ``` 
      {
        "message": "user not found"
      }
      ``` 

    * **Code**: 500
    * **Description**: something went wrong in server
    * **Content**:
       ``` 
       {
         "message": "internal server error"
       }
       ``` 

#### **Get profile**

* **URL**

  /v1/profile

* **Method**

  `GET`

* **Headers**

  Authorization Bearer `token`

* **URL Params**

  None

* **Data Params**
    
    None

* **Success Response:**

    * **Code**: 200
    * **Content**:
      ```
      {
        "fullname"  "Test user fullname'
        "login"     "test-user"
        "createdAt" "2024-05-06T23:02:59Z"
      }
      ```

* **Error Response**
    * **Code**: 400
    * **Description**: if user not found
    * **Content**:
      ``` 
      {
        "message": "user not found"
      }
      ``` 

    * **Code**: 500
    * **Description**: something went wrong in server
    * **Content**:
       ``` 
       {
         "message": "internal server error"
       }
       ```

#### **Upload track**

* **URL**

  /v1/tracks/upload

* **Method**

  `POST`

* **Headers**

  Authorization Bearer `token`

* **URL Params**

  None

* **Data Params**

     ```
    Content-Type: multipart/form-data
    {
       "file"     file
    }
    ```

* **Success Response:**

    * **Code**: 200
    * **Content**:
      ```
      {
         "message": "success",
         "path":    "ce67e596-8b3f-4ff8-baed-19c1e04292ab.mp3"
      }
      ```

* **Error Response**
    * **Code**: 400
    * **Description**: if input param is invalid
    * **Content**:
      ``` 
      {
        "message": "invalid input params"
      }
      ```

    * **Code**: 500
    * **Description**: something went wrong in server
    * **Content**:
       ``` 
       {
         "message": "internal server error"
       }
       ``` 




#### **Create track**

* **URL**

  /v1/tracks

* **Method**

  `POST`

* **Headers**

  Authorization Bearer `token`

* **URL Params**

  None

* **Data Params**

     ```
    Content-Type: application/json
    {
        title  string
        artist string
        genre  string 
        path   string
    }
    ```

* **Success Response:**

    * **Code**: 200
    * **Content**:
      ```
      {
         "message": "success",
      }
      ```

* **Error Response**
    * **Code**: 400
    * **Description**: if input param is invalid
    * **Content**:
      ``` 
      {
        "message": "invalid input params"
      }
      ```

    * **Code**: 500
    * **Description**: something went wrong in server
    * **Content**:
       ``` 
       {
         "message": "internal server error"
       }
       ``` 




#### **List tracks**

* **URL**

  /v1/tracks

* **Method**

  `GET`

* **Headers**

  Authorization Bearer `token`

* **URL Params**

  * page int
  * limit int // default 10

* **Data Params**

    None

* **Success Response:**

    * **Code**: 200
    * **Content**:
  ```
  {
     "id": 1,
     "title": "Test track",
     "artist": "Test artist",
     "genre": "test genre",
     "path": "http://localhost:9001/tracks/ce67e596-8b3f-4ff8-baed-19c1e04292ab.mp3",
     "uploader": "Test-user",
     "releaseDate": "2024-05-06T23:02:59Z"
  }
  ```

* **Error Response**
    * **Code**: 400
    * **Description**: if input param is invalid
    * **Content**:
      ``` 
      {
        "message": "invalid input params"
      }
      ```

    * **Code**: 500
    * **Description**: something went wrong in server
    * **Content**:
       ``` 
       {
         "message": "internal server error"
       }
       ``` 




#### **Like track**

* **URL**

  /v1/tracks/:id/like

* **Method**

  `POST`

* **Headers**

  Authorization Bearer `token`

* **URL Params**

    * id int
  
* **Data Params**

  None

* **Success Response:**

    * **Code**: 200
    * **Content**:
      ```
      {
         "message": "success",
      }
      ```

* **Error Response**
    * **Code**: 400
    * **Description**: if input param is invalid
    * **Content**:
      ``` 
      {
        "message": "invalid input params"
      }
      ```

    * **Code**: 400
    * **Description**: if operation failed
    * **Content**:
      ``` 
      {
        "message": "operation failed"
      }
      ```

    * **Code**: 500
    * **Description**: something went wrong in server
    * **Content**:
       ``` 
       {
         "message": "internal server error"
       }
       ``` 




#### **Revert like**

* **URL**

  /v1/tracks/:id/revertLike

* **Method**

  `POST`

* **Headers**

  Authorization Bearer `token`

* **URL Params**

    * id int

* **Data Params**

  None

* **Success Response:**

    * **Code**: 200
    * **Content**:
      ```
      {
         "message": "success",
      }
      ```

* **Error Response**
    * **Code**: 400
    * **Description**: if input param is invalid
    * **Content**:
      ``` 
      {
        "message": "invalid input params"
      }
      ```

    * **Code**: 400
    * **Description**: if operation failed
    * **Content**:
      ``` 
      {
        "message": "operation failed"
      }
      ```

    * **Code**: 500
    * **Description**: something went wrong in server
    * **Content**:
       ``` 
       {
         "message": "internal server error"
       }
       ``` 




#### **List favourite tracks**

* **URL**

  /v1/tracks/favourite

* **Method**

  `GET`

* **Headers**

  Authorization Bearer `token`

* **URL Params**

    * page int
    * limit int // default 10

* **Data Params**

  None

* **Success Response:**

    * **Code**: 200
* **Content**:
  ```
  {
     "id": 1,
     "title": "Test track",
     "artist": "Test artist",
     "genre": "test genre",
     "path": "http://localhost:9001/tracks/ce67e596-8b3f-4ff8-baed-19c1e04292ab.mp3",
     "uploader": "Test-user",
     "releaseDate": "2024-05-06T23:02:59Z"
  }
  ```

* **Error Response**
    * **Code**: 400
    * **Description**: if input param is invalid
    * **Content**:
      ``` 
      {
        "message": "invalid input params"
      }
      ```

    * **Code**: 500
    * **Description**: something went wrong in server
    * **Content**:
       ``` 
       {
         "message": "internal server error"
       }
       ``` 


