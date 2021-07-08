# Cartoon Redirect

This application redirects requests to official cartoon websites, which I've found useful for sharing links on Slack, for example.

## How to use

Just run and access the URL. The latest version is currently running at https://cartoon-redirect.herokuapp.com.

There are two ways to customize the output, specifying the cartoon and/or asking for the latest or random cartoon.

To choose a specific cartoon, append one of these to the URL:
```
/dilbert
/calvin
/garfield
/peanuts
/xkcd
```

To get a random cartoon, add `?q=random` to the URL.


### Examples:
```
# This will return the latest of any one of the available cartoons:
https://cartoon-redirect.herokuapp.com
```
```
# This will return the latest Dilbert cartoon:
https://cartoon-redirect.herokuapp.com/dilbert
```
```
# This will return a Peanuts cartoon from a random date:
https://cartoon-redirect.herokuapp.com/peanuts?q=random
```
```
# It's also possible to get all random, surprise me!
https://cartoon-redirect.herokuapp.com?q=random
```