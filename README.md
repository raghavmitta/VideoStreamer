# Video Streamer

## **Description**

Video Streamer is project developed using **Go** as programming languages and **MySql** as database . It uses **Youtube V3 API** to fetch video data based on given query and then after performing **ETL** operation , it exposes it to user using various **REST API**

## **Guide To Run Project Using Docker**


 1. First clone the repo on your local system using [link](https://github.com/raghavmitta/VideoStreamer.git)
 2. Ensure you have Docker Desktop installed in your system and is in running state
 3. Open config.yml and change config according to your requirement, here is sample config
 

    server:  
  port: 8081  
database:  
  driver_name: mysql  
  host_name: root:password@tcp(db:3306)  
  db_name: test  
pagination:  
  page_size: 10  (no. of response in a single page)
ticker:  
  time: 60  (time after which Youtube API should be consumed asynchronously  in seconds)
api:  
  keys: [MULTIPLE API ACCESS KEYS]  
  query: tutorial  (key to search Youtube V3 API)

 5. Open Command prompt and navigate to location which contains **docker-compose** file
 6. Run command :`docker-compose up` 
 7. Now, you should be able to see the output in terminal

## **API Description**
**GetAllUsers**: It returns all the video data sorted in descending order by publishing time
**Link**: http://localhost:8080/all

**ExactSearch**:It returns data after searching it based on query provided  in title and description ,sorted in descending order by publishing time
**Link**:http://localhost:8080/search/{query}

**PartialSearch**:It searches videos containing partial match for the search query in either video title or description. 
It sort based on no. words it matched with words in the query and if two videos have same matches then it sort them based on publishing time
**Link**:http://localhost:8080/partial-search/{query}
