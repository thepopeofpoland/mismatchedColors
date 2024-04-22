# mismatchedColors
A repository for the final full-stack website project for a web development course.

updated this file

## Frontend

## Backend
This is the side of the project where all of the go is. This is used to create endpoints for web requests to edit the database.

### /
This is a test endpoint that uses a POST request that writes []byte("hello") as an output.

### /getrandomcolors
This is an endpoint that uses a POST request to return all of the colors in the RandomColors table in the MismatchedColors.db. It returns them as 
``` JSON
{"ColorID": "int", 
"ColorName": "text", 
"ColorHex": "#FFFFFF"}
```
The ColorHex should (so long as people are responsible with inserting new data) output text that starts with # and then the six characters. For example, #FFFFFF. 

### /insertrandomcolor
Uses a POST request to insert a new color into the RandomColors table in the MismatchedColors.db. A body is needed to put in the post request that includes
``` JSON
{
    "ColorName": "text",
    "ColorHex": "#FFFFFF"
}
```

It then returns the color that was inserted into the table:

``` JSON
{
    "ColorID": "int",
    "ColorName": "text",
    "ColorHex": "#FFFFFF"
}
```